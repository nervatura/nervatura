//go:build http || all
// +build http all

package app

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	nt "github.com/nervatura/nervatura/service/pkg/nervatura"
	srv "github.com/nervatura/nervatura/service/pkg/service"
)

func Test_httpServer_StartService(t *testing.T) {
	type fields struct {
		app        *App
		mux        *chi.Mux
		service    srv.HTTPService
		admin      srv.AdminService
		client     srv.ClientService
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
			name: "start_ok",
			fields: fields{
				app: &App{
					config: nt.IM{
						"version":                  "test",
						"NT_HTTP_TLS_ENABLED":      true,
						"NT_HTTP_PORT":             int64(-1),
						"NT_HTTP_READ_TIMEOUT":     float64(30),
						"NT_HTTP_WRITE_TIMEOUT":    float64(30),
						"NT_TLS_CERT_FILE":         "../data/x509/server_cert.pem",
						"NT_TLS_KEY_FILE":          "../data/x509/server_key.pem",
						"NT_TOKEN_PUBLIC_KEY_TYPE": "RSA",
						"NT_TOKEN_PUBLIC_KEY_URL":  "https://www.googleapis.com/oauth2/v1/certs",

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
				client:     tt.fields.client,
				result:     tt.fields.result,
				server:     tt.fields.server,
				tlsEnabled: tt.fields.tlsEnabled,
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
		client     srv.ClientService
		result     string
		server     *http.Server
		tlsEnabled bool
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
						"version":                  "test",
						"NT_TOKEN_PUBLIC_KEY_TYPE": "RSA",
						"NT_TOKEN_PUBLIC_KEY_URL":  "httpsx://www.googleapis.coma",
					},
					infoLog:   log.New(os.Stdout, "INFO: ", log.LstdFlags),
					errorLog:  log.New(os.Stdout, "ERROR: ", log.LstdFlags),
					tokenKeys: make(map[string]map[string]string),
				},
			},
		},
		{
			name: "body_error",
			fields: fields{
				app: &App{
					config: nt.IM{
						"version":                  "test",
						"NT_TOKEN_PUBLIC_KEY_TYPE": "RSA",
						"NT_TOKEN_PUBLIC_KEY_URL":  "https://www.google.com",
					},
					infoLog:   log.New(os.Stdout, "INFO: ", log.LstdFlags),
					errorLog:  log.New(os.Stdout, "ERROR: ", log.LstdFlags),
					tokenKeys: make(map[string]map[string]string),
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
				client:     tt.fields.client,
				result:     tt.fields.result,
				server:     tt.fields.server,
				tlsEnabled: tt.fields.tlsEnabled,
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
		client     srv.ClientService
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
				client:     tt.fields.client,
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
		client     srv.ClientService
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
				client:     tt.fields.client,
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
		client     srv.ClientService
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
				client:     tt.fields.client,
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
		client     srv.ClientService
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
				client:     tt.fields.client,
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
		client     srv.ClientService
		result     string
		server     *http.Server
		tlsEnabled bool
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &httpServer{
				app:        tt.fields.app,
				mux:        tt.fields.mux,
				service:    tt.fields.service,
				admin:      tt.fields.admin,
				client:     tt.fields.client,
				result:     tt.fields.result,
				server:     tt.fields.server,
				tlsEnabled: tt.fields.tlsEnabled,
			}
			if got := s.tokenAuth(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("httpServer.tokenAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_httpServer_fileServer(t *testing.T) {
	type fields struct {
		app        *App
		mux        *chi.Mux
		service    srv.HTTPService
		admin      srv.AdminService
		client     srv.ClientService
		result     string
		server     *http.Server
		tlsEnabled bool
	}
	type args struct {
		path string
		root http.FileSystem
	}
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
			name: "fs_redirect",
			fields: fields{
				mux: chi.NewRouter(),
			},
			args: args{
				path: "/redirect",
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
				client:     tt.fields.client,
				result:     tt.fields.result,
				server:     tt.fields.server,
				tlsEnabled: tt.fields.tlsEnabled,
			}
			s.fileServer(tt.args.path, tt.args.root)
		})
	}
}

func Test_httpServer_homeRoute(t *testing.T) {
	type fields struct {
		app        *App
		mux        *chi.Mux
		service    srv.HTTPService
		admin      srv.AdminService
		client     srv.ClientService
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
				client:     tt.fields.client,
				result:     tt.fields.result,
				server:     tt.fields.server,
				tlsEnabled: tt.fields.tlsEnabled,
			}
			s.homeRoute(tt.args.w, tt.args.r)
		})
	}
}

func Test_httpServer_adminRoute(t *testing.T) {
	type fields struct {
		app        *App
		mux        *chi.Mux
		service    srv.HTTPService
		admin      srv.AdminService
		client     srv.ClientService
		result     string
		server     *http.Server
		tlsEnabled bool
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	formReq := func(params url.Values) *http.Request {
		req := httptest.NewRequest("POST", "/admin", strings.NewReader(params.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		return req
	}
	as := srv.AdminService{
		Config: nt.IM{
			"NT_API_KEY": "TEST_API_KEY",
		},
		GetNervaStore: func(database string) *nt.NervaStore {
			return nil
		},
		GetTokenKeys: func() map[string]map[string]string {
			return nil
		},
	}
	as.LoadTemplates()
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "database",
			fields: fields{
				admin: as,
			},
			args: args{
				w: httptest.NewRecorder(),
				r: formReq(url.Values{
					"formID": []string{"database"},
				}),
			},
		},
		{
			name: "admin",
			fields: fields{
				admin: as,
			},
			args: args{
				w: httptest.NewRecorder(),
				r: formReq(url.Values{
					"formID": []string{"admin"},
				}),
			},
		},
		{
			name: "menu",
			fields: fields{
				admin: as,
			},
			args: args{
				w: httptest.NewRecorder(),
				r: formReq(url.Values{
					"formID": []string{"menu"},
				}),
			},
		},
		{
			name: "client",
			fields: fields{
				admin: as,
			},
			args: args{
				w: httptest.NewRecorder(),
				r: formReq(url.Values{
					"formID": []string{"menu"},
					"menu":   []string{"client"},
				}),
			},
		},
		{
			name: "docs",
			fields: fields{
				admin: as,
			},
			args: args{
				w: httptest.NewRecorder(),
				r: formReq(url.Values{
					"formID": []string{"menu"},
					"menu":   []string{"docs"},
				}),
			},
		},
		{
			name: "login",
			fields: fields{
				admin: as,
			},
			args: args{
				w: httptest.NewRecorder(),
				r: formReq(url.Values{
					"formID": []string{"login"},
				}),
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
				client:     tt.fields.client,
				result:     tt.fields.result,
				server:     tt.fields.server,
				tlsEnabled: tt.fields.tlsEnabled,
			}
			s.adminRoute(tt.args.w, tt.args.r)
		})
	}
}
