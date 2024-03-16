//go:build http || all
// +build http all

package app

import (
	"context"
	"errors"
	"io"
	"io/fs"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	nt "github.com/nervatura/nervatura/service/pkg/nervatura"
	srv "github.com/nervatura/nervatura/service/pkg/service"
	ut "github.com/nervatura/nervatura/service/pkg/utils"
)

func Test_httpServer_StartService(t *testing.T) {
	type fields struct {
		app        *App
		mux        *chi.Mux
		service    srv.HTTPService
		admin      srv.AdminService
		result     string
		server     *http.Server
		tlsEnabled bool
		readAll    func(r io.Reader) ([]byte, error)
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "start_ok",
			fields: fields{
				app: &App{
					config: nt.IM{
						"version":                 "test",
						"NT_HTTP_TLS_ENABLED":     true,
						"NT_HTTP_PORT":            int64(-1),
						"NT_HTTP_READ_TIMEOUT":    float64(30),
						"NT_HTTP_WRITE_TIMEOUT":   float64(30),
						"NT_TLS_CERT_FILE":        "../data/x509/server_cert.pem",
						"NT_TLS_KEY_FILE":         "../data/x509/server_key.pem",
						"NT_TOKEN_PUBLIC_KEY_URL": "https://www.googleapis.com/oauth2/v1/certs",
						"NT_CLIENT_CONFIG":        "../data/test_client_config.json",

						"NT_CORS_ENABLED":           true,
						"NT_CORS_ALLOW_ORIGINS":     strings.Split("*", ","),
						"NT_CORS_ALLOW_METHODS":     strings.Split("GET,POST,DELETE,OPTIONS", ","),
						"NT_CORS_ALLOW_HEADERS":     strings.Split("Accept,Authorization,Content-Type,X-CSRF-Token,X-Api-Key", ","),
						"NT_CORS_EXPOSE_HEADERS":    strings.Split("", ","),
						"NT_CORS_ALLOW_CREDENTIALS": false,
						"NT_CORS_MAX_AGE":           int64(0),

						"NT_SECURITY_ENABLED":                    true,
						"NT_SECURITY_ALLOWED_HOSTS":              make([]string, 0),
						"NT_SECURITY_HOSTS_PROXY_HEADERS":        make([]string, 0),
						"NT_SECURITY_ALLOWED_HOSTS_ARE_REGEX":    false,
						"NT_SECURITY_SSL_REDIRECT":               false,
						"NT_SECURITY_SSL_TEMPORARY_REDIRECT":     false,
						"NT_SECURITY_SSL_HOST":                   "",
						"NT_SECURITY_PROXY_HEADERS":              make([]string, 0),
						"NT_SECURITY_STS_SECONDS":                int64(0),
						"NT_SECURITY_STS_INCLUDE_SUBDOMAINS":     false,
						"NT_SECURITY_STS_PRELOAD":                false,
						"NT_SECURITY_FORCE_STS_HEADER":           false,
						"NT_SECURITY_FRAME_DENY":                 false,
						"NT_SECURITY_CUSTOM_FRAME_OPTIONS_VALUE": "",
						"NT_SECURITY_CONTENT_TYPE_NOSNIFF":       false,
						"NT_SECURITY_BROWSER_XSS_FILTER":         false,
						"NT_SECURITY_CONTENT_SECURITY_POLICY":    "",
						"NT_SECURITY_PUBLIC_KEY":                 "",
						"NT_SECURITY_REFERRER_POLICY":            "",
						"NT_SECURITY_FEATURE_POLICY":             "",
						"NT_SECURITY_EXPECT_CT_HEADER":           "",
						"NT_SECURITY_DEVELOPMENT":                false,

						"NT_HTTP_HOME": "/admin",
					},
					infoLog:   log.New(os.Stdout, "INFO: ", log.LstdFlags),
					errorLog:  log.New(os.Stdout, "ERROR: ", log.LstdFlags),
					tokenKeys: make(map[string]map[string]string),
				},
				readAll: io.ReadAll,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &httpServer{
				app:        tt.fields.app,
				mux:        tt.fields.mux,
				service:    tt.fields.service,
				admin:      tt.fields.admin,
				result:     tt.fields.result,
				server:     tt.fields.server,
				tlsEnabled: tt.fields.tlsEnabled,
				readAll:    tt.fields.readAll,
			}
			if err := s.StartService(); (err != nil) != tt.wantErr {
				t.Errorf("httpServer.StartService() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_httpServer_setPublicKeys(t *testing.T) {
	type fields struct {
		app        *App
		mux        *chi.Mux
		service    srv.HTTPService
		admin      srv.AdminService
		result     string
		server     *http.Server
		tlsEnabled bool
		readAll    func(r io.Reader) ([]byte, error)
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "url_error",
			fields: fields{
				app: &App{
					config: nt.IM{
						"version":                 "test",
						"NT_TOKEN_PUBLIC_KEY_URL": "httpsx://www.googleapis.coma",
					},
					infoLog:   log.New(os.Stdout, "INFO: ", log.LstdFlags),
					errorLog:  log.New(os.Stdout, "ERROR: ", log.LstdFlags),
					tokenKeys: make(map[string]map[string]string),
				},
				readAll: io.ReadAll,
			},
		},
		{
			name: "body_error",
			fields: fields{
				app: &App{
					config: nt.IM{
						"version":                 "test",
						"NT_TOKEN_PUBLIC_KEY_URL": "https://www.google.com",
					},
					infoLog:   log.New(os.Stdout, "INFO: ", log.LstdFlags),
					errorLog:  log.New(os.Stdout, "ERROR: ", log.LstdFlags),
					tokenKeys: make(map[string]map[string]string),
				},
				readAll: func(r io.Reader) ([]byte, error) {
					return nil, errors.New("error")
				},
			},
		},
		{
			name: "body_ok",
			fields: fields{
				app: &App{
					config: nt.IM{
						"version":                 "test",
						"NT_TOKEN_PUBLIC_KEY_URL": "https://www.google.com",
					},
					infoLog:   log.New(os.Stdout, "INFO: ", log.LstdFlags),
					errorLog:  log.New(os.Stdout, "ERROR: ", log.LstdFlags),
					tokenKeys: make(map[string]map[string]string),
				},
				readAll: io.ReadAll,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &httpServer{
				app:        tt.fields.app,
				mux:        tt.fields.mux,
				service:    tt.fields.service,
				admin:      tt.fields.admin,
				result:     tt.fields.result,
				server:     tt.fields.server,
				tlsEnabled: tt.fields.tlsEnabled,
				readAll:    tt.fields.readAll,
			}
			s.setPublicKeys()
		})
	}
}

func Test_httpServer_startServer(t *testing.T) {
	type fields struct {
		app        *App
		mux        *chi.Mux
		service    srv.HTTPService
		admin      srv.AdminService
		result     string
		server     *http.Server
		tlsEnabled bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "tls_disabled",
			fields: fields{
				app: &App{
					config: nt.IM{
						"version":               "test",
						"NT_HTTP_TLS_ENABLED":   false,
						"NT_HTTP_PORT":          int64(-1),
						"NT_HTTP_READ_TIMEOUT":  float64(30),
						"NT_HTTP_WRITE_TIMEOUT": float64(30),
						"NT_TLS_CERT_FILE":      "",
						"NT_TLS_KEY_FILE":       "",
					},
					infoLog: log.New(os.Stdout, "INFO: ", log.LstdFlags),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &httpServer{
				app:        tt.fields.app,
				mux:        tt.fields.mux,
				service:    tt.fields.service,
				admin:      tt.fields.admin,
				result:     tt.fields.result,
				server:     tt.fields.server,
				tlsEnabled: tt.fields.tlsEnabled,
			}
			if err := s.startServer(); (err != nil) != tt.wantErr {
				t.Errorf("httpServer.startServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_httpServer_StopService(t *testing.T) {
	type fields struct {
		app        *App
		mux        *chi.Mux
		service    srv.HTTPService
		admin      srv.AdminService
		result     string
		server     *http.Server
		tlsEnabled bool
	}
	type args struct {
		ctx interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "stop_service",
			fields: fields{
				app: &App{
					infoLog: log.New(os.Stdout, "INFO: ", log.LstdFlags),
				},
				server: &http.Server{},
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
		{
			name:    "nil_server",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &httpServer{
				app:        tt.fields.app,
				mux:        tt.fields.mux,
				service:    tt.fields.service,
				admin:      tt.fields.admin,
				result:     tt.fields.result,
				server:     tt.fields.server,
				tlsEnabled: tt.fields.tlsEnabled,
			}
			if err := s.StopService(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("httpServer.StopService() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_httpServer_Results(t *testing.T) {
	type fields struct {
		app        *App
		mux        *chi.Mux
		service    srv.HTTPService
		admin      srv.AdminService
		result     string
		server     *http.Server
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
			s := &httpServer{
				app:        tt.fields.app,
				mux:        tt.fields.mux,
				service:    tt.fields.service,
				admin:      tt.fields.admin,
				result:     tt.fields.result,
				server:     tt.fields.server,
				tlsEnabled: tt.fields.tlsEnabled,
			}
			if got := s.Results(); got != tt.want {
				t.Errorf("httpServer.Results() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_httpServer_ConnectApp(t *testing.T) {
	type fields struct {
		app        *App
		mux        *chi.Mux
		service    srv.HTTPService
		admin      srv.AdminService
		result     string
		server     *http.Server
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
			s := &httpServer{
				app:        tt.fields.app,
				mux:        tt.fields.mux,
				service:    tt.fields.service,
				admin:      tt.fields.admin,
				result:     tt.fields.result,
				server:     tt.fields.server,
				tlsEnabled: tt.fields.tlsEnabled,
			}
			s.ConnectApp(tt.args.app)
		})
	}
}

func Test_httpServer_tokenAuth(t *testing.T) {
	type fields struct {
		app        *App
		mux        *chi.Mux
		service    srv.HTTPService
		admin      srv.AdminService
		result     string
		server     *http.Server
		tlsEnabled bool
		tokenLogin func(r *http.Request) (ctx context.Context, err error)
	}
	type args struct {
		next http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   http.Handler
	}{
		{
			name: "ok",
			fields: fields{
				service: srv.HTTPService{},
				tokenLogin: func(r *http.Request) (ctx context.Context, err error) {
					return context.Background(), nil
				},
			},
			args: args{
				next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			},
		},
		{
			name: "error",
			fields: fields{
				service: srv.HTTPService{},
				tokenLogin: func(r *http.Request) (ctx context.Context, err error) {
					return nil, errors.New("error")
				},
			},
			args: args{
				next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &httpServer{
				app:        tt.fields.app,
				mux:        tt.fields.mux,
				service:    tt.fields.service,
				admin:      tt.fields.admin,
				result:     tt.fields.result,
				server:     tt.fields.server,
				tlsEnabled: tt.fields.tlsEnabled,
				tokenLogin: tt.fields.tokenLogin,
			}
			s.tokenAuth(tt.args.next).ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("POST", "/", nil))
		})
	}
}

func Test_httpServer_homeRoute(t *testing.T) {
	type fields struct {
		app        *App
		mux        *chi.Mux
		service    srv.HTTPService
		admin      srv.AdminService
		result     string
		server     *http.Server
		tlsEnabled bool
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
			name: "admin_home",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/", nil),
			},
			fields: fields{
				app: &App{
					config: nt.IM{
						"NT_HTTP_HOME": "/admin",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &httpServer{
				app:        tt.fields.app,
				mux:        tt.fields.mux,
				service:    tt.fields.service,
				admin:      tt.fields.admin,
				result:     tt.fields.result,
				server:     tt.fields.server,
				tlsEnabled: tt.fields.tlsEnabled,
			}
			s.homeRoute(tt.args.w, tt.args.r)
		})
	}
}

func Test_httpServer_fileServer(t *testing.T) {
	type fields struct {
		app        *App
		mux        *chi.Mux
		service    srv.HTTPService
		admin      srv.AdminService
		result     string
		server     *http.Server
		tlsEnabled bool
		readAll    func(r io.Reader) ([]byte, error)
		tokenLogin func(r *http.Request) (ctx context.Context, err error)
	}
	type args struct {
		path string
		root http.FileSystem
	}
	var publicFS, _ = fs.Sub(ut.Public, "static")
	mux := chi.NewRouter()
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "fs_error",
			fields: fields{
				app: &App{
					errorLog: log.New(os.Stdout, "ERROR: ", log.LstdFlags),
				},
			},
			args: args{
				path: "{error}",
			},
		},
		{
			name: "serve",
			fields: fields{
				mux: mux,
			},
			args: args{
				path: "/",
				root: http.FS(publicFS),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &httpServer{
				app:        tt.fields.app,
				mux:        tt.fields.mux,
				service:    tt.fields.service,
				admin:      tt.fields.admin,
				result:     tt.fields.result,
				server:     tt.fields.server,
				tlsEnabled: tt.fields.tlsEnabled,
				readAll:    tt.fields.readAll,
				tokenLogin: tt.fields.tokenLogin,
			}
			s.fileServer(tt.args.path, tt.args.root)
		})
	}
}
