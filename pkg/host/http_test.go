//go:build http || all
// +build http all

package server

import (
	"bytes"
	"context"
	"errors"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
)

func Test_httpServer_StartServer(t *testing.T) {
	type fields struct {
		config     cu.IM
		appLog     *slog.Logger
		mux        *http.ServeMux
		server     *http.Server
		tlsEnabled bool
		result     string
	}
	type args struct {
		config    cu.IM
		interrupt chan os.Signal
		ctx       context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"start_ok",
			fields{},
			args{
				config: cu.IM{
					"version":               "test",
					"NT_HTTP_TLS_ENABLED":   true,
					"NT_HTTP_PORT":          int64(-1),
					"NT_HTTP_READ_TIMEOUT":  float64(30),
					"NT_HTTP_WRITE_TIMEOUT": float64(30),
					"NT_TLS_CERT_FILE":      "../../data/test_server_cert.pem",
					"NT_TLS_KEY_FILE":       "../../data/test_server_key.pem",
					"NT_MCP_ENABLED":        true,

					"NT_CORS_ENABLED":           true,
					"NT_CORS_ALLOW_ORIGINS":     strings.Split("*", ","),
					"NT_CORS_ALLOW_METHODS":     strings.Split("GET,POST,DELETE,OPTIONS", ","),
					"NT_CORS_ALLOW_HEADERS":     strings.Split("Accept,Authorization,Content-Type,X-CSRF-Token,X-Api-Key", ","),
					"NT_CORS_EXPOSE_HEADERS":    strings.Split("", ","),
					"NT_CSRF_TRUSTED_ORIGINS":   strings.Split("", ","),
					"NT_CORS_ALLOW_CREDENTIALS": false,
					"NT_CORS_MAX_AGE":           int64(0),
				},
				interrupt: make(chan os.Signal),
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &httpServer{
				config:     tt.fields.config,
				appLog:     tt.fields.appLog,
				mux:        tt.fields.mux,
				server:     tt.fields.server,
				tlsEnabled: tt.fields.tlsEnabled,
				result:     tt.fields.result,
			}
			appLogOut := &bytes.Buffer{}
			httpLogOut := &bytes.Buffer{}
			if err := s.StartServer(tt.args.config, appLogOut, httpLogOut, tt.args.interrupt, tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("httpServer.StartServer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_httpServer_StopServer(t *testing.T) {
	s := &httpServer{}
	interrupt := make(chan os.Signal)
	ctx := context.Background()
	go func() {
		s.StartServer(cu.IM{
			"version":                   "test",
			"NT_HTTP_PORT":              8080,
			"NT_CORS_ENABLED":           true,
			"NT_CORS_ALLOW_ORIGINS":     strings.Split("*", ","),
			"NT_CORS_ALLOW_METHODS":     strings.Split("GET,POST,DELETE,OPTIONS", ","),
			"NT_CORS_ALLOW_HEADERS":     strings.Split("Accept,Authorization,Content-Type,X-CSRF-Token,X-Api-Key", ","),
			"NT_CORS_EXPOSE_HEADERS":    strings.Split("", ","),
			"NT_CORS_ALLOW_CREDENTIALS": false,
			"NT_CORS_MAX_AGE":           int64(0),
			"NT_CSRF_TRUSTED_ORIGINS":   strings.Split("", ","),
		}, &bytes.Buffer{}, &bytes.Buffer{}, interrupt, ctx)
	}()
	time.Sleep(1 * time.Second)
	s.StopServer(ctx)

	s = &httpServer{}
	s.StopServer(ctx)
}

func Test_httpServer_Results(t *testing.T) {
	type fields struct {
		config     cu.IM
		appLog     *slog.Logger
		mux        *http.ServeMux
		server     *http.Server
		tlsEnabled bool
		result     string
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
			s := &httpServer{
				config:     tt.fields.config,
				appLog:     tt.fields.appLog,
				mux:        tt.fields.mux,
				server:     tt.fields.server,
				tlsEnabled: tt.fields.tlsEnabled,
				result:     tt.fields.result,
			}
			if got := s.Results(); got != tt.want {
				t.Errorf("httpServer.Results() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_httpServer_homeRoute(t *testing.T) {
	type fields struct {
		config     cu.IM
		appLog     *slog.Logger
		mux        *http.ServeMux
		server     *http.Server
		tlsEnabled bool
		result     string
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "home",
			fields: fields{
				config: cu.IM{"NT_HTTP_HOME": "/"},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/", nil),
			},
		},
		{
			name: "home_redirect",
			fields: fields{
				config: cu.IM{"NT_HTTP_HOME": "/auth/"},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/", nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &httpServer{
				config:     tt.fields.config,
				appLog:     tt.fields.appLog,
				mux:        tt.fields.mux,
				server:     tt.fields.server,
				tlsEnabled: tt.fields.tlsEnabled,
				result:     tt.fields.result,
			}
			s.homeRoute(tt.args.w, tt.args.r)
		})
	}
}

func Test_httpServer_envList(t *testing.T) {
	type fields struct {
		config     cu.IM
		appLog     *slog.Logger
		mux        *http.ServeMux
		server     *http.Server
		tlsEnabled bool
		result     string
	}
	os.Setenv("NT_ALIAS_KALEVALA", "sqlite5://file:../data/empty.db")
	tests := []struct {
		name   string
		fields fields
		want   []cu.IM
	}{
		{
			name: "ok",
			fields: fields{
				config: cu.IM{
					"NT_API_KEY": "EXAMPLE_API_KEY",
				},
			},
			want: []cu.IM{
				{"envkey": "NT_ALIAS_KALEVALA", "envvalue": "sqlite5://file:../data/empty.db"},
				{"envkey": "NT_API_KEY", "envvalue": "EXAMPLE_API_KEY"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &httpServer{
				config:     tt.fields.config,
				appLog:     tt.fields.appLog,
				mux:        tt.fields.mux,
				server:     tt.fields.server,
				tlsEnabled: tt.fields.tlsEnabled,
				result:     tt.fields.result,
			}
			if got := s.envList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("httpServer.envList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_httpServer_configRoute(t *testing.T) {
	type fields struct {
		config     cu.IM
		appLog     *slog.Logger
		mux        *http.ServeMux
		server     *http.Server
		tlsEnabled bool
		result     string
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	req := httptest.NewRequest("GET", "/config/EXAMPLE_API_KEY", nil)
	req.SetPathValue("secKey", "EXAMPLE_API_KEY")
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "missing_sec_key",
			fields: fields{
				config: cu.IM{"NT_TASK_SEC_KEY": "EXAMPLE_API_KEY"},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/", nil),
			},
		},
		{
			name: "ok",
			fields: fields{
				config: cu.IM{"NT_TASK_SEC_KEY": "EXAMPLE_API_KEY"},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: req,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &httpServer{
				config:     tt.fields.config,
				appLog:     tt.fields.appLog,
				mux:        tt.fields.mux,
				server:     tt.fields.server,
				tlsEnabled: tt.fields.tlsEnabled,
				result:     tt.fields.result,
			}
			s.configRoute(tt.args.w, tt.args.r)
		})
	}
}

func Test_httpServer_headerClient(t *testing.T) {
	type fields struct {
		config     cu.IM
		appLog     *slog.Logger
		mux        *http.ServeMux
		server     *http.Server
		tlsEnabled bool
		result     string
	}
	type args struct {
		next http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "ok",
			fields: fields{
				config: cu.IM{"NT_API_KEY": "EXAMPLE_API_KEY"},
			},
			args: args{
				next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &httpServer{
				config:     tt.fields.config,
				appLog:     tt.fields.appLog,
				mux:        tt.fields.mux,
				server:     tt.fields.server,
				tlsEnabled: tt.fields.tlsEnabled,
				result:     tt.fields.result,
			}
			s.headerClient(tt.args.next).ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("GET", "/", nil))
		})
	}
}

func Test_httpServer_headerAuth(t *testing.T) {
	type fields struct {
		config     cu.IM
		appLog     *slog.Logger
		mux        *http.ServeMux
		server     *http.Server
		tlsEnabled bool
		result     string
	}
	type args struct {
		next http.Handler
		req  *http.Request
	}
	req_api := httptest.NewRequest("GET", "/", nil)
	req_api.Header.Set("X-API-KEY", "EXAMPLE_API_KEY")
	req_token := httptest.NewRequest("GET", "/", nil)
	req_token.Header.Set("Authorization", "Bearer EXAMPLE_TOKEN")
	req_login := httptest.NewRequest("GET", "/auth/login", nil)
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "api_key",
			fields: fields{
				config: cu.IM{"NT_API_KEY": "EXAMPLE_API_KEY"},
			},
			args: args{
				next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
				}),
				req: req_api,
			},
		},
		{
			name: "token",
			fields: fields{
				config: cu.IM{
					"NT_API_KEY": "EXAMPLE_API_KEY",
					"tokenKeys":  []cu.SM{{"type": "public", "value": "EXAMPLE_TOKEN"}},
				},
				appLog: slog.New(slog.NewTextHandler(os.Stdout, nil)),
			},
			args: args{
				next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
				}),
				req: req_token,
			},
		},
		{
			name: "login",
			fields: fields{
				config: cu.IM{},
			},
			args: args{
				next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
				req:  req_login,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &httpServer{
				config:     tt.fields.config,
				appLog:     tt.fields.appLog,
				mux:        tt.fields.mux,
				server:     tt.fields.server,
				tlsEnabled: tt.fields.tlsEnabled,
				result:     tt.fields.result,
			}
			s.headerAuth(tt.args.next).ServeHTTP(httptest.NewRecorder(), tt.args.req)
		})
	}
}

func Test_httpServer_headerSession(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		next http.Handler
	}{
		{
			name: "ok",
			next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var s httpServer
			s.headerSession(tt.next).ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("GET", "/", nil))
		})
	}
}

func Test_httpServer_mcpVerify(t *testing.T) {
	type fields struct {
		config     cu.IM
		appLog     *slog.Logger
		mux        *http.ServeMux
		server     *http.Server
		session    *api.SessionService
		tlsEnabled bool
		result     string
		memSession map[string]md.MemoryStore
	}
	type args struct {
		ctx   context.Context
		token string
		req   *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				config: cu.IM{"NT_API_KEY": "EXAMPLE_API_KEY"},
				appLog: slog.New(slog.NewTextHandler(os.Stdout, nil)),
			},
			args: args{
				ctx:   context.Background(),
				token: "EXAMPLE_API_KEY",
				req:   httptest.NewRequest("GET", "/", nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &httpServer{
				config:     tt.fields.config,
				appLog:     tt.fields.appLog,
				mux:        tt.fields.mux,
				server:     tt.fields.server,
				session:    tt.fields.session,
				tlsEnabled: tt.fields.tlsEnabled,
				result:     tt.fields.result,
				memSession: tt.fields.memSession,
			}
			_, err := s.mcpVerify(tt.args.ctx, tt.args.token, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("httpServer.mcpVerify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_httpServer_loadPrompts(t *testing.T) {
	type fields struct {
		config          cu.IM
		appLog          *slog.Logger
		mux             *http.ServeMux
		server          *http.Server
		session         *api.SessionService
		tlsEnabled      bool
		result          string
		memSession      map[string]md.MemoryStore
		ReadFile        func(name string) ([]byte, error)
		ConvertFromByte func(data []byte, result interface{}) error
		StaticReadFile  func(name string) ([]byte, error)
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "file_not_found",
			fields: fields{
				config: cu.IM{"NT_MCP_PROMPT": "../../data/not_found.json"},
				appLog: slog.New(slog.NewTextHandler(os.Stdout, nil)),
				ReadFile: func(name string) ([]byte, error) {
					return nil, errors.New("file not found")
				},
				ConvertFromByte: cu.ConvertFromByte,
				StaticReadFile: func(name string) ([]byte, error) {
					return []byte{}, nil
				},
			},
		},
		{
			name: "convert_error",
			fields: fields{
				config: cu.IM{"NT_MCP_PROMPT": "../../data/prompt.json"},
				appLog: slog.New(slog.NewTextHandler(os.Stdout, nil)),
				ReadFile: func(name string) ([]byte, error) {
					return []byte(`{"prompts": [{"name": "test", "description": "test", "prompt": "test"}]}`), nil
				},
				ConvertFromByte: func(data []byte, result interface{}) error {
					return errors.New("convert error")
				},
				StaticReadFile: func(name string) ([]byte, error) {
					return []byte{}, nil
				},
			},
		},
		{
			name: "static_file_not_found",
			fields: fields{
				config: cu.IM{"NT_MCP_PROMPT": "../../data/prompt.json"},
				appLog: slog.New(slog.NewTextHandler(os.Stdout, nil)),
				ReadFile: func(name string) ([]byte, error) {
					return nil, errors.New("file not found")
				},
				ConvertFromByte: cu.ConvertFromByte,
				StaticReadFile: func(name string) ([]byte, error) {
					return nil, errors.New("file not found")
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &httpServer{
				config:          tt.fields.config,
				appLog:          tt.fields.appLog,
				mux:             tt.fields.mux,
				server:          tt.fields.server,
				session:         tt.fields.session,
				tlsEnabled:      tt.fields.tlsEnabled,
				result:          tt.fields.result,
				memSession:      tt.fields.memSession,
				ReadFile:        tt.fields.ReadFile,
				ConvertFromByte: tt.fields.ConvertFromByte,
				StaticReadFile:  tt.fields.StaticReadFile,
			}
			s.loadPrompts()
		})
	}
}

func Test_httpServer_loadResources(t *testing.T) {
	type fields struct {
		config          cu.IM
		appLog          *slog.Logger
		mux             *http.ServeMux
		server          *http.Server
		session         *api.SessionService
		tlsEnabled      bool
		result          string
		memSession      map[string]md.MemoryStore
		ReadFile        func(name string) ([]byte, error)
		StaticReadFile  func(name string) ([]byte, error)
		ConvertFromByte func(data []byte, result any) error
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "file_not_found",
			fields: fields{
				config: cu.IM{"NT_MCP_RESOURCE": "../../data/not_found.json"},
				appLog: slog.New(slog.NewTextHandler(os.Stdout, nil)),
				ReadFile: func(name string) ([]byte, error) {
					return nil, errors.New("file not found")
				},
				ConvertFromByte: cu.ConvertFromByte,
				StaticReadFile: func(name string) ([]byte, error) {
					return []byte{}, nil
				},
			},
		},
		{
			name: "convert_error",
			fields: fields{
				config: cu.IM{"NT_MCP_RESOURCE": "../../data/resource.json"},
				appLog: slog.New(slog.NewTextHandler(os.Stdout, nil)),
				ReadFile: func(name string) ([]byte, error) {
					return []byte(`{"resources": [{"name": "test", "description": "test", "resource": "test"}]}`), nil
				},
				ConvertFromByte: func(data []byte, result interface{}) error {
					return errors.New("convert error")
				},
				StaticReadFile: func(name string) ([]byte, error) {
					return []byte{}, nil
				},
			},
		},
		{
			name: "static_file_not_found",
			fields: fields{
				config: cu.IM{"NT_MCP_RESOURCE": "../../data/resource.json"},
				appLog: slog.New(slog.NewTextHandler(os.Stdout, nil)),
				ReadFile: func(name string) ([]byte, error) {
					return nil, errors.New("file not found")
				},
				ConvertFromByte: cu.ConvertFromByte,
				StaticReadFile: func(name string) ([]byte, error) {
					return nil, errors.New("file not found")
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &httpServer{
				config:          tt.fields.config,
				appLog:          tt.fields.appLog,
				mux:             tt.fields.mux,
				server:          tt.fields.server,
				session:         tt.fields.session,
				tlsEnabled:      tt.fields.tlsEnabled,
				result:          tt.fields.result,
				memSession:      tt.fields.memSession,
				ReadFile:        tt.fields.ReadFile,
				StaticReadFile:  tt.fields.StaticReadFile,
				ConvertFromByte: tt.fields.ConvertFromByte,
			}
			s.loadResources()
		})
	}
}

func Test_httpServer_loadLabels(t *testing.T) {
	type fields struct {
		config          cu.IM
		appLog          *slog.Logger
		mux             *http.ServeMux
		server          *http.Server
		session         *api.SessionService
		tlsEnabled      bool
		result          string
		memSession      map[string]md.MemoryStore
		ReadFile        func(name string) ([]byte, error)
		StaticReadFile  func(name string) ([]byte, error)
		ConvertFromByte func(data []byte, result any) error
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "file_not_found",
			fields: fields{
				config: cu.IM{"NT_CLIENT_LABELS": "../../data/not_found.json"},
				appLog: slog.New(slog.NewTextHandler(os.Stdout, nil)),
				ReadFile: func(name string) ([]byte, error) {
					return nil, errors.New("file not found")
				},
				ConvertFromByte: cu.ConvertFromByte,
				StaticReadFile: func(name string) ([]byte, error) {
					return []byte{}, nil
				},
			},
		},
		{
			name: "convert_error",
			fields: fields{
				config: cu.IM{"NT_CLIENT_LABELS": "../../data/client_labels.json"},
				appLog: slog.New(slog.NewTextHandler(os.Stdout, nil)),
				ReadFile: func(name string) ([]byte, error) {
					return []byte(`{}`), nil
				},
				ConvertFromByte: func(data []byte, result interface{}) error {
					return errors.New("convert error")
				},
				StaticReadFile: func(name string) ([]byte, error) {
					return []byte{}, nil
				},
			},
		},
		{
			name: "success",
			fields: fields{
				config: cu.IM{"NT_CLIENT_LABELS": "../../data/client_labels.json"},
				appLog: slog.New(slog.NewTextHandler(os.Stdout, nil)),
				ReadFile: func(name string) ([]byte, error) {
					return []byte(`{}`), nil
				},
				ConvertFromByte: func(data []byte, result interface{}) error {
					return nil
				},
				StaticReadFile: func(name string) ([]byte, error) {
					return []byte{}, nil
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &httpServer{
				config:          tt.fields.config,
				appLog:          tt.fields.appLog,
				mux:             tt.fields.mux,
				server:          tt.fields.server,
				session:         tt.fields.session,
				tlsEnabled:      tt.fields.tlsEnabled,
				result:          tt.fields.result,
				memSession:      tt.fields.memSession,
				ReadFile:        tt.fields.ReadFile,
				StaticReadFile:  tt.fields.StaticReadFile,
				ConvertFromByte: tt.fields.ConvertFromByte,
			}
			s.loadLabels()
		})
	}
}
