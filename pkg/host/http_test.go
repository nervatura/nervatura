//go:build http || all
// +build http all

package server

import (
	"bytes"
	"context"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	cu "github.com/nervatura/component/pkg/util"
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
			if err := s.StartServer(tt.args.config, appLogOut, httpLogOut, tt.args.interrupt); (err != nil) != tt.wantErr {
				t.Errorf("httpServer.StartServer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_httpServer_StopServer(t *testing.T) {
	s := &httpServer{}
	interrupt := make(chan os.Signal)
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
		}, &bytes.Buffer{}, &bytes.Buffer{}, interrupt)
	}()
	time.Sleep(1 * time.Second)
	s.StopServer(context.Background())

	s = &httpServer{}
	s.StopServer(context.Background())
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

func Test_httpServer_headerAPI(t *testing.T) {
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
			s.headerAPI(tt.args.next).ServeHTTP(httptest.NewRecorder(), tt.args.req)
		})
	}
}
