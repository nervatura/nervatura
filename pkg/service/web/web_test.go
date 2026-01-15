package web

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"path"
	"strings"
	"testing"
	"time"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	cp "github.com/nervatura/nervatura/v6/pkg/client/web/component"
	cls "github.com/nervatura/nervatura/v6/pkg/client/web/service"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	st "github.com/nervatura/nervatura/v6/pkg/static"
	td "github.com/nervatura/nervatura/v6/test/driver"
	"golang.org/x/oauth2"
)

func TestClientAuth(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ok",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/", nil),
			},
		},
	}
	cls := &cls.ClientService{
		Config: cu.IM{},
		Session: &api.SessionService{
			Config: api.SessionConfig{
				Method: md.SessionMethodMemory,
			},
		},
		UI:      cp.NewClientComponent(map[string]cu.SM{}),
		ReadAll: io.ReadAll,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.ClientServiceCtxKey, cls)
			ClientAuth(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestClientAuthEvent(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	r := httptest.NewRequest("POST", "/", nil)
	r.Header.Set("X-Session-Token", "SES012345")
	tests := []struct {
		name string
		args args
	}{
		{
			name: "missing",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/", nil),
			},
		},
		{
			name: "ok",
			args: args{
				w: httptest.NewRecorder(),
				r: r,
			},
		},
	}
	cls := &cls.ClientService{
		Config: cu.IM{},
		Session: &api.SessionService{
			Config: api.SessionConfig{
				Method: md.SessionMethodMemory,
			},
		},
		UI:      cp.NewClientComponent(map[string]cu.SM{}),
		ReadAll: io.ReadAll,
	}
	cls.Session.SaveSession("SES012345", &ct.Client{
		Ticket: ct.Ticket{
			SessionID: "SES012345",
		},
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.ClientServiceCtxKey, cls)
			ClientAuthEvent(httptest.NewRecorder(), tt.args.r.WithContext(ctx))
		})
	}
}

func TestClientSession(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	req := httptest.NewRequest("GET", "/session/SessionID", nil)
	req.SetPathValue("session_id", "SES012345")
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ok",
			args: args{
				w: httptest.NewRecorder(),
				r: req,
			},
		},
		{
			name: "missing",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/session/SessionID", nil),
			},
		},
	}
	cls := &cls.ClientService{
		Config: cu.IM{},
		Session: &api.SessionService{
			Config: api.SessionConfig{
				Method: md.SessionMethodMemory,
			},
		},
		UI:      cp.NewClientComponent(map[string]cu.SM{}),
		ReadAll: io.ReadAll,
	}
	cls.Session.SaveSession("SES012345", &ct.Client{
		Ticket: ct.Ticket{
			SessionID: "SES012345",
			User:      cu.IM{},
			Expiry:    time.Now().Add(time.Duration(1) * time.Hour),
		},
		BaseComponent: ct.BaseComponent{
			Data: cu.IM{
				"request_data": cu.IM{
					"code": "code",
					"name": "name",
				},
			},
		},
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.ClientServiceCtxKey, cls)
			ClientSession(httptest.NewRecorder(), tt.args.r.WithContext(ctx))
		})
	}
}

func TestClientSessionEvent(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	cls1 := &cls.ClientService{
		Config: cu.IM{},
		AppLog: slog.Default(),
		NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
			return &api.DataStore{
				Db:     &td.TestDriver{Config: cu.IM{}},
				Config: config,
				AppLog: appLog,
			}
		},
		Session: &api.SessionService{
			Config: api.SessionConfig{
				Method: md.SessionMethodMemory,
			},
		},
		UI:      cp.NewClientComponent(map[string]cu.SM{}),
		ReadAll: io.ReadAll,
	}
	cls1.Session.SaveSession("SES012345", &ct.Client{
		Ticket: ct.Ticket{
			SessionID: "SES012345",
			User:      cu.IM{},
			Expiry:    time.Now().Add(time.Duration(1) * time.Hour),
		},
	})

	cls2 := &cls.ClientService{
		Config: cu.IM{},
		AppLog: slog.Default(),
		NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
			return &api.DataStore{
				Db:     &td.TestDriver{Config: cu.IM{}},
				Config: config,
				AppLog: appLog,
			}
		},
		Session: &api.SessionService{
			Config: api.SessionConfig{
				Method: md.SessionMethodMemory,
			},
		},
		ReadAll: func(r io.Reader) (data []byte, err error) {
			return nil, errors.New("error")
		},
	}
	cls2.Session.SaveSession("SES012345", &ct.Client{
		Ticket: ct.Ticket{
			SessionID: "SES012345",
			User:      cu.IM{},
			Expiry:    time.Now().Add(time.Duration(1) * time.Hour),
		},
	})

	tests := []struct {
		name    string
		args    args
		content string
		token   string
		cls     *cls.ClientService
	}{
		{
			name: "invalid",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/", nil),
			},
			content: "application/x-www-form-urlencoded",
			token:   "INVALID",
			cls:     cls1,
		},
		{
			name: "valid",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/", nil),
			},
			content: "application/x-www-form-urlencoded",
			token:   "SES012345",
			cls:     cls1,
		},
		{
			name: "triggerEvent_err",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/", nil),
			},
			content: "application/json",
			token:   "SES012345",
			cls:     cls2,
		},
	}

	for _, tt := range tests {
		tt.args.r.Header.Set("X-Session-Token", tt.token)
		tt.args.r.Header.Set("Content-Type", tt.content)
		ctx := context.WithValue(context.Background(), md.ClientServiceCtxKey, tt.cls)
		t.Run(tt.name, func(t *testing.T) {
			ClientSessionEvent(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestClientSessionCreate(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w  http.ResponseWriter
		r  *http.Request
		db api.DataDriver
	}{
		{
			name: "convert error",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest("POST", "/", strings.NewReader("invalid")),
			db:   &td.TestDriver{Config: cu.IM{}},
		},
		{
			name: "ok",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest("POST", "/", strings.NewReader(`{"database": "test", "username": "test", "lang": "en", "theme": "light", "module": "search", "request_id": "1234567890"}`)),
			db: &td.TestDriver{Config: cu.IM{
				"Connection": func() struct {
					Alias     string
					Connected bool
					Engine    string
				} {
					return struct {
						Alias     string
						Connected bool
						Engine    string
					}{
						Alias:     "test",
						Connected: true,
						Engine:    "sqlite",
					}
				},
				"QuerySQL": func(sqlString string) ([]cu.IM, error) {
					return []cu.IM{{"id": 1, "name": "test"}}, nil
				},
				"Query": func(queries []md.Query) ([]cu.IM, error) {
					return []cu.IM{{"id": 1, "code": "test", "user_name": "test", "user_group": "GROUP_ADMIN"}}, nil
				},
			}},
		},
		{
			name: "api key error",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest("POST", "/", strings.NewReader(`{"api_key": "invalid"}`)),
			db: &td.TestDriver{Config: cu.IM{
				"Connection": func() struct {
					Alias     string
					Connected bool
					Engine    string
				} {
					return struct {
						Alias     string
						Connected bool
						Engine    string
					}{
						Alias:     "test",
						Connected: true,
						Engine:    "sqlite",
					}
				},
				"QuerySQL": func(sqlString string) ([]cu.IM, error) {
					return []cu.IM{{"id": 1, "name": "test"}}, nil
				},
				"Query": func(queries []md.Query) ([]cu.IM, error) {
					return []cu.IM{}, errors.New("error")
				},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.ClientServiceCtxKey, &cls.ClientService{
				Config: cu.IM{},
				AppLog: slog.Default(),
				Session: &api.SessionService{
					Config: api.SessionConfig{
						Method: md.SessionMethodMemory,
					},
				},
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:     tt.db,
						Config: config,
						AppLog: appLog,
						ConvertToType: func(data interface{}, result any) (err error) {
							return nil
						},
						ConvertFromByte: func(data []byte, result interface{}) error {
							return cu.ConvertFromByte(data, result)
						},
						ConvertToByte: func(data interface{}) ([]byte, error) {
							return cu.ConvertToByte(data)
						},
						ReadAll: func(r io.Reader) ([]byte, error) {
							return io.ReadAll(r)
						},
					}
				},
				UI: cp.NewClientComponent(map[string]cu.SM{}),
			})
			ClientSessionCreate(tt.w, tt.r.WithContext(ctx))
		})
	}
}

func TestClientAuthCallback(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		w.Header().Set("Content-Type", "application/json")
		code := r.Form.Get("code")
		if code == "access_token" {
			w.Write([]byte(`{
				"access_token": "90d64460d14870c08c81352a05dedd3465940a7c.90d64460d14870c08c81352a05dedd3465940a7c",
				"scope": "email",
				"token_type": "bearer",
				"expires_in": 3600
			}`))
			return
		}
		w.Write([]byte(`{
			"access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiAiMTIzNDU2Nzg5MCIsImlhdCI6IDE1MTYyMzkwMjIsImVtYWlsIjoidXNlckBlbWFpbC5jb20ifQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
			"id_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiAiMTIzNDU2Nzg5MCIsImlhdCI6IDE1MTYyMzkwMjIsImVtYWlsIjoidXNlckBlbWFpbC5jb20ifQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
			"scope": "email",
			"token_type": "bearer",
			"expires_in": 3600
		}`))
	}))
	defer ts.Close()
	ses1 := &api.SessionService{
		Config: api.SessionConfig{
			Method: md.SessionMethodMemory,
		},
		Conn: &td.TestDriver{Config: cu.IM{}},
	}
	ses1.SaveSession("SessionID", &ct.Client{BaseComponent: ct.BaseComponent{
		Data: cu.IM{
			"auth_config": "invalid",
			"verifier":    "VE0123",
		},
	},
		Ticket: ct.Ticket{
			SessionID: "SessionID",
			User:      cu.IM{},
			Expiry:    time.Now().Add(time.Duration(1) * time.Hour),
		}})
	ses2 := &api.SessionService{
		Config: api.SessionConfig{
			Method: md.SessionMethodMemory,
		},
		Conn: &td.TestDriver{Config: cu.IM{}},
	}
	ses2.SaveSession("SessionID", &ct.Client{
		BaseComponent: ct.BaseComponent{
			Data: cu.IM{
				"auth_config": "testaut",
				"verifier":    "VE0123",
			},
		},
		Ticket: ct.Ticket{
			SessionID: "SessionID",
			User:      cu.IM{},
			Expiry:    time.Now().Add(time.Duration(1) * time.Hour),
		},
	})
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w          http.ResponseWriter
		r          *http.Request
		session    *api.SessionService
		authConfig *oauth2.Config
		config     cu.IM
	}{
		{
			name: "not_found",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest("GET", "/callback?state=SessionID", nil),
			session: &api.SessionService{
				Config: api.SessionConfig{
					Method: md.SessionMethodMemory,
				},
			},
			config: cu.IM{},
		},
		{
			name:       "invalid_token",
			w:          httptest.NewRecorder(),
			r:          httptest.NewRequest("GET", "/callback?state=SessionID", nil),
			session:    ses1,
			authConfig: &oauth2.Config{},
			config:     cu.IM{},
		},
		{
			name:    "access_token",
			w:       httptest.NewRecorder(),
			r:       httptest.NewRequest("GET", "/callback?state=SessionID&code=access_token", nil),
			session: ses2,
			authConfig: &oauth2.Config{
				ClientID:     "TEST_CLIENT_ID",
				ClientSecret: "TEST_CLIENT_SECRET",
				RedirectURL:  "REDIRECT_URL",
				Scopes:       []string{"email"},
				Endpoint: oauth2.Endpoint{
					AuthURL:  "testAuthCodeURL",
					TokenURL: ts.URL,
				},
			},
			config: cu.IM{},
		},
		{
			name:    "id_token",
			w:       httptest.NewRecorder(),
			r:       httptest.NewRequest("GET", "/callback?state=SessionID&code=id_token", nil),
			session: ses2,
			authConfig: &oauth2.Config{
				ClientID:     "TEST_CLIENT_ID",
				ClientSecret: "TEST_CLIENT_SECRET",
				RedirectURL:  "REDIRECT_URL",
				Scopes:       []string{"email"},
				Endpoint: oauth2.Endpoint{
					AuthURL:  "testAuthCodeURL",
					TokenURL: ts.URL,
				},
			},
			config: cu.IM{
				"Connection": func() struct {
					Alias     string
					Connected bool
					Engine    string
				} {
					return struct {
						Alias     string
						Connected bool
						Engine    string
					}{
						Alias:     "test",
						Connected: true,
						Engine:    "sqlite",
					}
				},
				"QuerySQL": func(sqlString string) ([]cu.IM, error) {
					return []cu.IM{{"id": 1, "name": "test"}}, nil
				},
				"Query": func(queries []md.Query) ([]cu.IM, error) {
					return []cu.IM{}, errors.New("error")
				},
			},
		},
		{
			name:    "success",
			w:       httptest.NewRecorder(),
			r:       httptest.NewRequest("GET", "/callback?state=SessionID&code=id_token", nil),
			session: ses2,
			authConfig: &oauth2.Config{
				ClientID:     "TEST_CLIENT_ID",
				ClientSecret: "TEST_CLIENT_SECRET",
				RedirectURL:  "REDIRECT_URL",
				Scopes:       []string{"email"},
				Endpoint: oauth2.Endpoint{
					AuthURL:  "testAuthCodeURL",
					TokenURL: ts.URL,
				},
			},
			config: cu.IM{
				"Connection": func() struct {
					Alias     string
					Connected bool
					Engine    string
				} {
					return struct {
						Alias     string
						Connected bool
						Engine    string
					}{
						Alias:     "test",
						Connected: true,
						Engine:    "sqlite",
					}
				},
				"QuerySQL": func(sqlString string) ([]cu.IM, error) {
					return []cu.IM{{"id": 1, "name": "test"}}, nil
				},
				"Query": func(queries []md.Query) ([]cu.IM, error) {
					if queries[0].From == "user_view" {
						return []cu.IM{{"id": 1, "code": "test", "user_name": "test", "user_group": "GROUP_ADMIN"}}, nil
					}
					return []cu.IM{{"value": ""}}, nil
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.ClientServiceCtxKey, &cls.ClientService{
				Config:     cu.IM{},
				AppLog:     slog.Default(),
				Session:    tt.session,
				AuthConfig: tt.authConfig,
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:     &td.TestDriver{Config: tt.config},
						Config: config,
						AppLog: appLog,
						ConvertFromByte: func(b []byte, i interface{}) error {
							_ = cu.ConvertFromByte([]byte("{\"email\":\"user@email.com\"}"), i)
							return nil
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(cu.IM{})
						},
						ConvertToType: func(data interface{}, result any) (err error) {
							return nil
						},
					}
				},
				UI: cp.NewClientComponent(map[string]cu.SM{}),
			})
			ClientAuthCallback(tt.w, tt.r.WithContext(ctx))
		})
	}
}

func TestClientExportBrowser(t *testing.T) {
	req := httptest.NewRequest("GET", "/export/SessionID", nil)
	req.SetPathValue("session_id", "SessionID")

	ses1 := &api.SessionService{
		Config: api.SessionConfig{
			Method: md.SessionMethodMemory,
		},
		Conn: &td.TestDriver{Config: cu.IM{}},
	}
	ses1.SaveSession("SessionID", &ct.Client{
		BaseComponent: ct.BaseComponent{
			Data: cu.IM{
				"editor": cu.IM{
					"key":  "patient",
					"view": "submission",
				},
			},
		},
	})

	ses2 := &api.SessionService{
		Config: api.SessionConfig{
			Method: md.SessionMethodMemory,
		},
		Conn: &td.TestDriver{Config: cu.IM{}},
	}
	ses2.SaveSession("SessionID", &ct.Client{
		Ticket: ct.Ticket{
			SessionID: "SessionID",
			User:      cu.IM{},
			Expiry:    time.Now().Add(time.Duration(1) * time.Hour),
		},
		BaseComponent: ct.BaseComponent{
			Data: cu.IM{
				"search": cu.IM{
					"view":     "customer",
					"customer": cu.IM{},
					"rows": []cu.IM{
						{"id": 12345, "patient": "Name", "missing": "value"},
					},
				},
			},
		},
		Lang:            "en",
		CustomFunctions: cp.NewClientComponent(map[string]cu.SM{}),
	})
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w       http.ResponseWriter
		r       *http.Request
		session *api.SessionService
	}{
		{
			name:    "session_err",
			w:       httptest.NewRecorder(),
			r:       req,
			session: ses1,
		},
		{
			name:    "ok",
			w:       httptest.NewRecorder(),
			r:       req,
			session: ses2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.ClientServiceCtxKey, &cls.ClientService{
				Config:  cu.IM{},
				AppLog:  slog.Default(),
				Session: tt.session,
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:     &td.TestDriver{Config: cu.IM{}},
						Config: config,
						AppLog: appLog,
					}
				},
				UI: cp.NewClientComponent(map[string]cu.SM{}),
			})
			ClientExportBrowser(tt.w, tt.r.WithContext(ctx))
		})
	}
}

func TestClientExportReport(t *testing.T) {
	req1 := httptest.NewRequest("GET", "/session/export/report/SessionID?output=pdf&inline=true", nil)
	req2 := httptest.NewRequest("GET", "/session/export/report/SessionID?output=pdf", nil)
	req2.SetPathValue("session_id", "SessionID")
	req3 := httptest.NewRequest("GET", "/session/export/report/SessionID?output=pdf&queue=QUEUE_CODE", nil)
	req3.SetPathValue("session_id", "SessionID")
	req4 := httptest.NewRequest("GET", "/session/export/report/SessionID?output=csv&export=true", nil)
	req4.SetPathValue("session_id", "SessionID")
	ses1 := &api.SessionService{
		Config: api.SessionConfig{
			Method: md.SessionMethodMemory,
		},
		Conn: &td.TestDriver{Config: cu.IM{}},
	}
	ses1.SaveSession("SessionID", &ct.Client{
		Ticket: ct.Ticket{
			SessionID: "SessionID",
			User:      cu.IM{},
			Expiry:    time.Now().Add(time.Duration(1) * time.Hour),
			Database:  "test",
		},
		BaseComponent: ct.BaseComponent{
			Data: cu.IM{
				"editor": cu.IM{
					"shortcut": cu.IM{
						"report_key": "test",
					},
					"params": cu.IM{},
				},
				"modal": cu.IM{
					"data": cu.IM{
						"template":    "test",
						"orientation": "portrait",
						"paper_size":  "A4",
						"code":        "test",
					},
				},
			},
		},
	})
	pdf_json, _ := st.Report.ReadFile(path.Join("template", "ntr_customer_en.json"))
	csv_json, _ := st.Report.ReadFile(path.Join("template", "csv_custpos_en.json"))
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w       http.ResponseWriter
		r       *http.Request
		session *api.SessionService
		config  cu.IM
	}{
		{
			name:    "session_err",
			w:       httptest.NewRecorder(),
			r:       req1,
			session: ses1,
			config:  cu.IM{},
		},
		{
			name:    "output_err",
			w:       httptest.NewRecorder(),
			r:       req2,
			session: ses1,
			config:  cu.IM{},
		},
		{
			name:    "ok_pdf",
			w:       httptest.NewRecorder(),
			r:       req2,
			session: ses1,
			config: cu.IM{
				"NT_ALIAS_TEST": "test",
				"Query": func(queries []md.Query) ([]cu.IM, error) {
					return []cu.IM{{"id": 1, "data": cu.IM{"file_type": "FILE_PDF", "template": string(pdf_json)}}}, nil
				},
				"QuerySQL": func(sqlString string) ([]cu.IM, error) {
					return []cu.IM{{"id": 1, "data": cu.IM{"file_type": "FILE_PDF"}}}, nil
				},
			},
		},
		{
			name:    "ok_export",
			w:       httptest.NewRecorder(),
			r:       req4,
			session: ses1,
			config: cu.IM{
				"NT_ALIAS_TEST": "test",
				"Query": func(queries []md.Query) ([]cu.IM, error) {
					return []cu.IM{{"id": 1, "data": cu.IM{"file_type": "FILE_PDF", "template": string(pdf_json)}}}, nil
				},
				"QuerySQL": func(sqlString string) ([]cu.IM, error) {
					return []cu.IM{{"id": 1, "data": cu.IM{"file_type": "FILE_PDF"}}}, nil
				},
			},
		},
		{
			name:    "ok_csv",
			w:       httptest.NewRecorder(),
			r:       req2,
			session: ses1,
			config: cu.IM{
				"NT_ALIAS_TEST": "test",
				"Query": func(queries []md.Query) ([]cu.IM, error) {
					return []cu.IM{{"id": 1, "data": cu.IM{"file_type": "FILE_CSV", "template": string(csv_json)}}}, nil
				},
				"QuerySQL": func(sqlString string) ([]cu.IM, error) {
					return []cu.IM{{"id": 1, "data": cu.IM{"file_type": "FILE_CSV"}}}, nil
				},
			},
		},
		{
			name:    "ok_pdf_queue",
			w:       httptest.NewRecorder(),
			r:       req3,
			session: ses1,
			config: cu.IM{
				"NT_ALIAS_TEST": "test",
				"Query": func(queries []md.Query) ([]cu.IM, error) {
					return []cu.IM{{"id": 1, "data": cu.IM{"file_type": "FILE_PDF", "template": string(pdf_json)}}}, nil
				},
				"QuerySQL": func(sqlString string) ([]cu.IM, error) {
					return []cu.IM{{"id": 1, "data": cu.IM{"file_type": "FILE_PDF"}}}, nil
				},
			},
		},
		{
			name:    "error_pdf_queue",
			w:       httptest.NewRecorder(),
			r:       req3,
			session: ses1,
			config: cu.IM{
				"NT_ALIAS_TEST": "test",
				"Query": func(queries []md.Query) ([]cu.IM, error) {
					return []cu.IM{}, errors.New("error")
				},
				"QuerySQL": func(sqlString string) ([]cu.IM, error) {
					return []cu.IM{{"id": 1, "data": cu.IM{"file_type": "FILE_PDF"}}}, nil
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.ClientServiceCtxKey, &cls.ClientService{
				Config:  tt.config,
				AppLog:  slog.Default(),
				Session: tt.session,
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:              &td.TestDriver{Config: tt.config},
						Config:          config,
						AppLog:          appLog,
						ConvertFromByte: cu.ConvertFromByte,
						ConvertToByte: func(v any) ([]byte, error) {
							return nil, nil
						},
						ConvertToType: func(data interface{}, result any) (err error) {
							return nil
						},
					}
				},
				UI: cp.NewClientComponent(map[string]cu.SM{}),
			})
			ClientExportReport(tt.w, tt.r.WithContext(ctx))
		})
	}
}

func TestClientTemplateEditor(t *testing.T) {
	req1 := httptest.NewRequest("GET", "/editor/SessionID/test", nil)
	req1.SetPathValue("session_id", "SessionID")
	req1.SetPathValue("code", "test")
	ses1 := &api.SessionService{
		Config: api.SessionConfig{
			Method: md.SessionMethodMemory,
		},
		Conn: &td.TestDriver{Config: cu.IM{}},
	}
	ses1.SaveSession("SessionID", &ct.Client{
		Ticket: ct.Ticket{
			SessionID: "SessionID",
			User:      cu.IM{},
			Expiry:    time.Now().Add(time.Duration(1) * time.Hour),
			Database:  "test",
		},
	})
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w          http.ResponseWriter
		r          *http.Request
		config     cu.IM
		session    *api.SessionService
		tokenError error
	}{
		{
			name:    "success",
			w:       httptest.NewRecorder(),
			r:       req1,
			session: ses1,
			config: cu.IM{
				"Query": func(queries []md.Query) ([]cu.IM, error) {
					return []cu.IM{{"id": 1, "data": cu.IM{"file_type": "FILE_PDF"}}}, nil
				},
			},
			tokenError: nil,
		},
		{
			name:    "token_error",
			w:       httptest.NewRecorder(),
			r:       req1,
			session: ses1,
			config: cu.IM{
				"Query": func(queries []md.Query) ([]cu.IM, error) {
					return []cu.IM{{"id": 1, "data": cu.IM{"file_type": "FILE_PDF"}}}, nil
				},
			},
			tokenError: errors.New("token error"),
		},
		{
			name:    "query_error",
			w:       httptest.NewRecorder(),
			r:       req1,
			session: ses1,
			config: cu.IM{
				"Query": func(queries []md.Query) ([]cu.IM, error) {
					return []cu.IM{}, errors.New("query error")
				},
			},
			tokenError: nil,
		},
		{
			name:    "session_error",
			w:       httptest.NewRecorder(),
			r:       httptest.NewRequest("POST", "/editor/SessionID/test", nil),
			session: ses1,
			config: cu.IM{
				"Query": func(queries []md.Query) ([]cu.IM, error) {
					return []cu.IM{}, nil
				},
			},
			tokenError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.ClientServiceCtxKey, &cls.ClientService{
				Config:  tt.config,
				AppLog:  slog.Default(),
				Session: tt.session,
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:              &td.TestDriver{Config: tt.config},
						Config:          config,
						AppLog:          appLog,
						ConvertFromByte: cu.ConvertFromByte,
						ConvertToByte: func(v any) ([]byte, error) {
							return nil, nil
						},
						ConvertToType: func(data interface{}, result any) (err error) {
							return nil
						},
						CreateLoginToken: func(user cu.SM, config cu.IM) (string, error) {
							return "token", tt.tokenError
						},
					}
				},
				UI: cp.NewClientComponent(map[string]cu.SM{}),
			})
			ClientTemplateEditor(tt.w, tt.r.WithContext(ctx))
		})
	}
}
