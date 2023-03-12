//go:build grpc || all
// +build grpc all

package app

import (
	"context"
	"errors"
	"log"
	"os"
	"reflect"
	"testing"

	nt "github.com/nervatura/nervatura/service/pkg/nervatura"
	srv "github.com/nervatura/nervatura/service/pkg/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func Test_rpcServer_StartService(t *testing.T) {
	type fields struct {
		app        *App
		service    srv.RPCService
		result     string
		server     *grpc.Server
		tlsEnabled bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr string
	}{
		{
			name: "cert_ok",
			fields: fields{
				app: &App{
					config: nt.IM{
						"version":             "test",
						"NT_GRPC_TLS_ENABLED": true,
						"NT_GRPC_PORT":        int64(-1),
						"NT_TLS_CERT_FILE":    "../data/x509/server_cert.pem",
						"NT_TLS_KEY_FILE":     "../data/x509/server_key.pem",
					},
					infoLog:  log.New(os.Stdout, "INFO: ", log.LstdFlags),
					errorLog: log.New(os.Stdout, "ERROR: ", log.LstdFlags),
				},
			},
			wantErr: "listen tcp: address -1: invalid port",
		},
		{
			name: "cert_error",
			fields: fields{
				app: &App{
					config: nt.IM{
						"version":             "test",
						"NT_GRPC_TLS_ENABLED": true,
						"NT_GRPC_PORT":        int64(-1),
						"NT_TLS_CERT_FILE":    "server_cert.pem",
						"NT_TLS_KEY_FILE":     "server_key.pem",
					},
					infoLog:  log.New(os.Stdout, "INFO: ", log.LstdFlags),
					errorLog: log.New(os.Stdout, "ERROR: ", log.LstdFlags),
				},
			},
			wantErr: "open server_cert.pem: no such file or directory",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &rpcServer{
				app:        tt.fields.app,
				service:    tt.fields.service,
				result:     tt.fields.result,
				server:     tt.fields.server,
				tlsEnabled: tt.fields.tlsEnabled,
			}
			if err := s.StartService(); err.Error() != tt.wantErr {
				t.Errorf("rpcServer.StartService() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_rpcServer_StopService(t *testing.T) {
	type fields struct {
		app        *App
		service    srv.RPCService
		result     string
		server     *grpc.Server
		tlsEnabled bool
	}
	type args struct {
		in0 interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "stop_ok",
			fields: fields{
				app: &App{
					infoLog: log.New(os.Stdout, "INFO: ", log.LstdFlags),
				},
				server: grpc.NewServer(),
			},
			wantErr: false,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &rpcServer{
				app:        tt.fields.app,
				service:    tt.fields.service,
				result:     tt.fields.result,
				server:     tt.fields.server,
				tlsEnabled: tt.fields.tlsEnabled,
			}
			if err := s.StopService(tt.args.in0); (err != nil) != tt.wantErr {
				t.Errorf("rpcServer.StopService() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_rpcServer_ConnectApp(t *testing.T) {
	type fields struct {
		app        *App
		service    srv.RPCService
		result     string
		server     *grpc.Server
		tlsEnabled bool
	}
	type args struct {
		app interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "app",
			args: args{
				app: &App{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &rpcServer{
				app:        tt.fields.app,
				service:    tt.fields.service,
				result:     tt.fields.result,
				server:     tt.fields.server,
				tlsEnabled: tt.fields.tlsEnabled,
			}
			s.ConnectApp(tt.args.app)
		})
	}
}

func Test_rpcServer_Results(t *testing.T) {
	type fields struct {
		app        *App
		service    srv.RPCService
		result     string
		server     *grpc.Server
		tlsEnabled bool
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
			s := &rpcServer{
				app:        tt.fields.app,
				service:    tt.fields.service,
				result:     tt.fields.result,
				server:     tt.fields.server,
				tlsEnabled: tt.fields.tlsEnabled,
			}
			if got := s.Results(); got != tt.want {
				t.Errorf("rpcServer.Results() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rpcServer_tokenAuth(t *testing.T) {
	type fields struct {
		app        *App
		service    srv.RPCService
		result     string
		server     *grpc.Server
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
		want    interface{}
		wantErr bool
	}{
		{
			name: "login_ok",
			args: args{
				ctx: context.Background(),
				info: &grpc.UnaryServerInfo{
					FullMethod: "/nervatura.API/UserLogin",
				},
				handler: func(ctx context.Context, req interface{}) (interface{}, error) {
					return "OK", nil
				},
			},
			want:    "OK",
			wantErr: false,
		},
		{
			name: "api_key_ok",
			fields: fields{
				service: srv.RPCService{
					Config: nt.IM{
						"NT_API_KEY": "TEST_API_KEY",
					},
					GetNervaStore: func(database string) *nt.NervaStore {
						return nil
					},
				},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.Pairs("X-Api-Key", "TEST_API_KEY")),
				info: &grpc.UnaryServerInfo{
					FullMethod: "/nervatura.API/DatabaseCreate",
				},
				handler: func(ctx context.Context, req interface{}) (interface{}, error) {
					return "OK", nil
				},
			},
			want:    "OK",
			wantErr: false,
		},
		{
			name: "api_key_missing",
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.Pairs("Hejehuja", "TEST_API_KEY")),
				info: &grpc.UnaryServerInfo{
					FullMethod: "/nervatura.API/DatabaseCreate",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "metadata_missing",
			args: args{
				ctx: context.Background(),
				info: &grpc.UnaryServerInfo{
					FullMethod: "/nervatura.API/DatabaseCreate",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "token_valid",
			fields: fields{
				auth: func(authorization []string, parent context.Context) (ctx context.Context, err error) {
					return context.Background(), nil
				},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.Pairs("Authorization", "TEST_TOKEN")),
				info: &grpc.UnaryServerInfo{
					FullMethod: "/nervatura.API/TokenRefresh",
				},
				handler: func(ctx context.Context, req interface{}) (interface{}, error) {
					return "OK", nil
				},
			},
			want:    "OK",
			wantErr: false,
		},
		{
			name: "token_invalid",
			fields: fields{
				auth: func(authorization []string, parent context.Context) (ctx context.Context, err error) {
					return nil, errors.New("error")
				},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.Pairs("Authorization", "TEST_TOKEN")),
				info: &grpc.UnaryServerInfo{
					FullMethod: "/nervatura.API/TokenRefresh",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &rpcServer{
				app:        tt.fields.app,
				service:    tt.fields.service,
				result:     tt.fields.result,
				server:     tt.fields.server,
				tlsEnabled: tt.fields.tlsEnabled,
				auth:       tt.fields.auth,
			}
			got, err := s.tokenAuth(tt.args.ctx, tt.args.req, tt.args.info, tt.args.handler)
			if (err != nil) != tt.wantErr {
				t.Errorf("rpcServer.tokenAuth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rpcServer.tokenAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}
