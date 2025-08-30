//go:build grpc || all
// +build grpc all

package server

import (
	"bytes"
	"context"
	"errors"
	"log/slog"
	"os"
	"testing"
	"time"

	cu "github.com/nervatura/component/pkg/util"
	srv "github.com/nervatura/nervatura/v6/pkg/service/grpc"
	pb "github.com/nervatura/nervatura/v6/pkg/service/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func Test_grpcServer_StartServer(t *testing.T) {
	type fields struct {
		config     cu.IM
		appLog     *slog.Logger
		server     *grpc.Server
		service    srv.GService
		result     string
		tlsEnabled bool
		auth       func(authorization []string, parent context.Context) (ctx context.Context, err error)
	}
	type args struct {
		config    cu.IM
		interrupt chan os.Signal
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "cert_ok",
			args: args{
				config: cu.IM{
					"version":             "test",
					"NT_GRPC_TLS_ENABLED": true,
					"NT_GRPC_PORT":        int64(-1),
					"NT_TLS_CERT_FILE":    "../../data/test_server_cert.pem",
					"NT_TLS_KEY_FILE":     "../../data/test_server_key.pem",
				},
				interrupt: make(chan os.Signal),
			},
			wantErr: true,
		},
		{
			name: "cert_error",
			args: args{
				config: cu.IM{
					"version":             "test",
					"NT_GRPC_TLS_ENABLED": true,
					"NT_GRPC_PORT":        int64(-1),
					"NT_TLS_CERT_FILE":    "server_cert.pem",
					"NT_TLS_KEY_FILE":     "server_key.pem",
				},
				interrupt: make(chan os.Signal),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &grpcServer{
				config:     tt.fields.config,
				appLog:     tt.fields.appLog,
				server:     tt.fields.server,
				service:    tt.fields.service,
				result:     tt.fields.result,
				tlsEnabled: tt.fields.tlsEnabled,
				auth:       tt.fields.auth,
			}
			appLogOut := &bytes.Buffer{}
			httpLogOut := &bytes.Buffer{}
			if err := s.StartServer(tt.args.config, appLogOut, httpLogOut, tt.args.interrupt); (err != nil) != tt.wantErr {
				t.Errorf("grpcServer.StartServer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_grpcServer_StopServer(t *testing.T) {
	s := &grpcServer{}
	interrupt := make(chan os.Signal)
	go func() {
		s.StartServer(cu.IM{
			"version":             "test",
			"NT_GRPC_TLS_ENABLED": true,
			"NT_GRPC_PORT":        9200,
		}, &bytes.Buffer{}, &bytes.Buffer{}, interrupt)
	}()
	time.Sleep(1 * time.Second)
	s.StopServer(context.Background())
}

func Test_grpcServer_Results(t *testing.T) {
	type fields struct {
		config     cu.IM
		appLog     *slog.Logger
		server     *grpc.Server
		service    srv.GService
		result     string
		tlsEnabled bool
		auth       func(authorization []string, parent context.Context) (ctx context.Context, err error)
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "result",
			fields: fields{
				result: "value",
			},
			want: "value",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &grpcServer{
				config:     tt.fields.config,
				appLog:     tt.fields.appLog,
				server:     tt.fields.server,
				service:    tt.fields.service,
				result:     tt.fields.result,
				tlsEnabled: tt.fields.tlsEnabled,
				auth:       tt.fields.auth,
			}
			if got := s.Results(); got != tt.want {
				t.Errorf("grpcServer.Results() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_grpcServer_tokenAuth(t *testing.T) {
	type fields struct {
		config     cu.IM
		appLog     *slog.Logger
		server     *grpc.Server
		service    srv.GService
		result     string
		tlsEnabled bool
		auth       func(authorization []string, parent context.Context) (ctx context.Context, err error)
	}
	type args struct {
		ctx     context.Context
		req     interface{}
		info    *grpc.UnaryServerInfo
		handler grpc.UnaryHandler
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "auth_login",
			args: args{
				ctx: context.Background(),
				req: nil,
				info: &grpc.UnaryServerInfo{
					FullMethod: pb.API_AuthLogin_FullMethodName,
				},
				handler: func(ctx context.Context, req interface{}) (interface{}, error) {
					return nil, nil
				},
			},
			wantErr: false,
		},
		{
			name: "api_key_ok",
			fields: fields{
				service: srv.GService{
					Config: cu.IM{
						"NT_API_KEY": "TEST_API_KEY",
					},
				},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.Pairs("x-api-key", "TEST_API_KEY")),
				req: nil,
				info: &grpc.UnaryServerInfo{
					FullMethod: pb.API_CurrencyGet_FullMethodName,
				},
				handler: func(ctx context.Context, req interface{}) (interface{}, error) {
					return nil, nil
				},
			},
			wantErr: false,
		},
		{
			name: "api_key_error",
			fields: fields{
				service: srv.GService{
					Config: cu.IM{
						"NT_API_KEY": "TEST_API_KEY",
					},
				},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.Pairs("x-api-key", "TEST_API_KEY_ERROR")),
				req: nil,
				info: &grpc.UnaryServerInfo{
					FullMethod: pb.API_CurrencyGet_FullMethodName,
				},
				handler: func(ctx context.Context, req interface{}) (interface{}, error) {
					return nil, nil
				},
			},
			wantErr: true,
		},
		{
			name: "missing_metadata",
			args: args{
				ctx: context.Background(),
				req: nil,
				info: &grpc.UnaryServerInfo{
					FullMethod: pb.API_CurrencyGet_FullMethodName,
				},
			},
			wantErr: true,
		},
		{
			name: "invalid_token",
			fields: fields{
				auth: func(authorization []string, parent context.Context) (ctx context.Context, err error) {
					return nil, errors.New("invalid token")
				},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.Pairs("authorization", "INVALID_TOKEN")),
				req: nil,
				info: &grpc.UnaryServerInfo{
					FullMethod: pb.API_CurrencyGet_FullMethodName,
				},
				handler: func(ctx context.Context, req interface{}) (interface{}, error) {
					return nil, nil
				},
			},
			wantErr: true,
		},
		{
			name: "token_ok",
			fields: fields{
				auth: func(authorization []string, parent context.Context) (ctx context.Context, err error) {
					return context.Background(), nil
				},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.Pairs("authorization", "TEST_TOKEN")),
				req: nil,
				info: &grpc.UnaryServerInfo{
					FullMethod: pb.API_CurrencyGet_FullMethodName,
				},
				handler: func(ctx context.Context, req interface{}) (interface{}, error) {
					return nil, nil
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &grpcServer{
				config:     tt.fields.config,
				appLog:     tt.fields.appLog,
				server:     tt.fields.server,
				service:    tt.fields.service,
				result:     tt.fields.result,
				tlsEnabled: tt.fields.tlsEnabled,
				auth:       tt.fields.auth,
			}
			_, err := s.tokenAuth(tt.args.ctx, tt.args.req, tt.args.info, tt.args.handler)
			if (err != nil) != tt.wantErr {
				t.Errorf("grpcServer.tokenAuth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
