//go:build grpc || all
// +build grpc all

package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log/slog"
	"net"
	"os"
	"time"

	cu "github.com/nervatura/component/pkg/util"
	srv "github.com/nervatura/nervatura/v6/pkg/service/grpc"
	pb "github.com/nervatura/nervatura/v6/pkg/service/grpc/proto"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// grpcServer - Nervatura gRpc server
type grpcServer struct {
	config     cu.IM
	appLog     *slog.Logger
	server     *grpc.Server
	service    srv.GService
	result     string
	tlsEnabled bool
	auth       func(authorization []string, parent context.Context) (ctx context.Context, err error)
}

func init() {
	registerHost("grpc", &grpcServer{})
}

func (s *grpcServer) StartServer(config cu.IM, appLogOut, httpLogOut io.Writer, interrupt chan os.Signal) error {
	s.config = config
	s.appLog = slog.New(slog.NewJSONHandler(appLogOut, nil))
	s.service = srv.GService{
		Config:          config,
		AppLog:          s.appLog,
		DecodeAuthToken: ut.TokenDecode,
	}
	s.auth = s.service.TokenAuth

	var cred credentials.TransportCredentials
	if cu.ToBoolean(s.config["NT_GRPC_TLS_ENABLED"], false) {
		if cu.ToString(s.config["NT_TLS_CERT_FILE"], "") != "" && cu.ToString(s.config["NT_TLS_KEY_FILE"], "") != "" {
			cert, err := tls.LoadX509KeyPair(cu.ToString(s.config["NT_TLS_CERT_FILE"], ""), cu.ToString(s.config["NT_TLS_KEY_FILE"], ""))
			if err != nil {
				s.appLog.Error("error loading TLS certificate and key", "error", err)
				return err
			}
			cred = credentials.NewServerTLSFromCert(&cert)
			s.tlsEnabled = true
		}
	}

	s.server = grpc.NewServer(
		grpc.Creds(cred),
		// MaxConnectionAge is just to avoid long connection, to facilitate load balancing
		// MaxConnectionAgeGrace will torn them, default to infinity
		grpc.KeepaliveParams(keepalive.ServerParameters{MaxConnectionAge: 2 * time.Minute}),
		grpc.UnaryInterceptor(s.tokenAuth),
	)

	pb.RegisterAPIServer(s.server, &s.service)
	// Register reflection service on gRPC server.
	reflection.Register(s.server)

	s.appLog.Info(fmt.Sprintf("gRPC server serving at: %d. SSL/TLS authentication: %v.",
		cu.ToInteger(s.config["NT_GRPC_PORT"], 0), s.tlsEnabled))
	addr := fmt.Sprintf(":%d", cu.ToInteger(s.config["NT_GRPC_PORT"], 0))
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		s.appLog.Error("error starting gRPC server", "error", err)
		return err
	}
	return s.server.Serve(ln)
}

func (s *grpcServer) StopServer(ctx context.Context) error {
	if s.server != nil {
		s.appLog.Info("stopping gRPC server")
		s.server.GracefulStop()
	}
	return nil
}

func (s *grpcServer) Results() string {
	return s.result
}

func (s *grpcServer) tokenAuth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	switch info.FullMethod {
	case pb.API_AuthLogin_FullMethodName, pb.API_TokenDecode_FullMethodName:
		return handler(ctx, req)
	default:
		mt, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.InvalidArgument, "missing metadata")
		}

		if len(mt["x-api-key"]) > 0 {
			ictx, err := s.service.ApiKeyAuth(mt["x-api-key"], ctx)
			if err != nil {
				return nil, status.Errorf(codes.Unauthenticated, "invalid api key")
			}
			return handler(ictx, req)
		}

		ictx, err := s.auth(mt["authorization"], ctx)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "invalid token")
		}

		return handler(ictx, req)
	}
}
