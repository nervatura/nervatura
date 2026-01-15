package http

import (
	"bytes"
	"context"
	"errors"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	td "github.com/nervatura/nervatura/v6/test/driver"
)

func TestProtectedResource(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/oauth-protected-resource", nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.ConfigCtxKey, cu.IM{})
			ProtectedResource(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestAuthorizationServer(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w http.ResponseWriter
		r *http.Request
	}{
		{
			name: "success",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest("GET", "/oauth-authorization-server", nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.ConfigCtxKey, cu.IM{})
			AuthorizationServer(tt.w, tt.r.WithContext(ctx))
		})
	}
}

func TestOpenIDConfiguration(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w http.ResponseWriter
		r *http.Request
	}{
		{
			name: "success",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest("GET", "/openid-configuration", nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.ConfigCtxKey, cu.IM{})
			OpenIDConfiguration(tt.w, tt.r.WithContext(ctx))
		})
	}
}

func TestJwks(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w http.ResponseWriter
		r *http.Request
	}{
		{
			name: "success",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest("GET", "/jwks.json", nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.ConfigCtxKey, cu.IM{})
			Jwks(tt.w, tt.r.WithContext(ctx))
		})
	}
}

func TestOAuthAuthorization(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w http.ResponseWriter
		r *http.Request
	}{
		{
			name: "success",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest("GET", "/oauth-authorization", nil),
		},
		{
			name: "invalid_response_type",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest("GET", "/oauth-authorization?response_type=invalid", nil),
		},
		{
			name: "invalid_redirect_uri",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest("GET", "/oauth-authorization?redirect_uri=invalid", nil),
		},
		{
			name: "invalid_client_id",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest("GET", "/oauth-authorization?client_id=invalid", nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.ConfigCtxKey, cu.IM{})
			ctx = context.WithValue(ctx, md.SessionServiceCtxKey, &api.SessionService{})
			OAuthAuthorization(tt.w, tt.r.WithContext(ctx))
		})
	}
}

func TestOAuthToken(t *testing.T) {
	req1 := httptest.NewRequest("POST", "/oauth/token", nil)
	req1.PostForm = url.Values{}
	req1.PostForm.Add("grant_type", "authorization_code")
	req1.PostForm.Add("code", "SES012345")
	req1.PostForm.Add("client_id", "test")
	ses := &api.SessionService{
		Config: api.SessionConfig{
			Method: md.SessionMethodMemory,
		},
		Conn: &td.TestDriver{Config: cu.IM{}},
	}
	ses.SaveSession("SES012345", cu.IM{})

	req2 := httptest.NewRequest("POST", "/oauth/token", nil)
	req2.PostForm = url.Values{}
	req2.PostForm.Add("grant_type", "authorization_code")
	req2.PostForm.Add("code", "SES012346")
	req2.PostForm.Add("client_id", "test")

	req3 := httptest.NewRequest("POST", "/oauth/token", nil)
	req3.PostForm = url.Values{}
	req3.PostForm.Add("grant_type", "authorization_code")
	req3.PostForm.Add("code", "SES012346")
	req3.PostForm.Add("client_id", "test2")

	req4 := httptest.NewRequest("POST", "/oauth/token", nil)
	req4.PostForm = url.Values{}
	req4.PostForm.Add("grant_type", "authorization_code")

	req5 := httptest.NewRequest("POST", "/oauth/token", nil)
	req5.PostForm = url.Values{}
	req5.PostForm.Add("grant_type", "invalid")

	req6 := httptest.NewRequest("POST", "/oauth/token", bytes.NewBufferString(";;;"))
	req6.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w http.ResponseWriter
		r *http.Request
	}{
		{
			name: "success",
			w:    httptest.NewRecorder(),
			r:    req1,
		},
		{
			name: "session_not_found",
			w:    httptest.NewRecorder(),
			r:    req2,
		},
		{
			name: "invalid_client_id",
			w:    httptest.NewRecorder(),
			r:    req3,
		},
		{
			name: "invalid_code",
			w:    httptest.NewRecorder(),
			r:    req4,
		},
		{
			name: "invalid_grant_type",
			w:    httptest.NewRecorder(),
			r:    req5,
		},
		{
			name: "invalid_request",
			w:    httptest.NewRecorder(),
			r:    req6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.ConfigCtxKey, cu.IM{
				"NT_AUTH_CLIENT_ID": "test",
			})
			ctx = context.WithValue(ctx, md.SessionServiceCtxKey, ses)
			OAuthToken(tt.w, tt.r.WithContext(ctx))
		})
	}
}

func TestOAuthRegistration(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w http.ResponseWriter
		r *http.Request
	}{
		{
			name: "success",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest("GET", "/oauth-registration", nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OAuthRegistration(tt.w, tt.r)
		})
	}
}

func TestOAuthCallback(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w http.ResponseWriter
		r *http.Request
	}{
		{
			name: "success",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest("GET", "/oauth-callback?code=test&state=SES012345", nil),
		},
		{
			name: "error",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest("GET", "/oauth-callback?code=test&state=SES012346", nil),
		},
	}
	ses := &api.SessionService{
		Config: api.SessionConfig{
			Method: md.SessionMethodMemory,
		},
		Conn: &td.TestDriver{Config: cu.IM{}},
	}
	ses.SaveSession("SES012345", cu.IM{})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.ConfigCtxKey, cu.IM{})
			ctx = context.WithValue(ctx, md.SessionServiceCtxKey, ses)
			OAuthCallback(tt.w, tt.r.WithContext(ctx))
		})
	}
}

func TestOAuthLogin(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/oauth-login", nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.ConfigCtxKey, cu.IM{})
			ctx = context.WithValue(ctx, md.SessionServiceCtxKey, &api.SessionService{})
			OAuthLogin(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestOAuthValidate(t *testing.T) {
	req1 := httptest.NewRequest("POST", "/oauth/validate", nil)
	req1.PostForm = url.Values{}
	req1.PostForm.Add("username", "test")
	req1.PostForm.Add("password", "123456")
	req1.PostForm.Add("database", "test")
	req1.PostForm.Add("session_id", "SES012345")

	req2 := httptest.NewRequest("POST", "/oauth/token", bytes.NewBufferString(";;;"))
	req2.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	ses1 := &api.SessionService{
		Config: api.SessionConfig{
			Method: md.SessionMethodMemory,
		},
		Conn: &td.TestDriver{Config: cu.IM{}},
	}
	ses1.SaveSession("SES012345", cu.IM{})

	ses2 := &api.SessionService{
		Config: api.SessionConfig{
			Method: md.SessionMethodMemory,
		},
		Conn: &td.TestDriver{Config: cu.IM{}},
	}
	ses2.SaveSession("SES012345", cu.SM{
		"redirect_uri": "https://example.com",
		"state":        "SES012345",
	})

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w   http.ResponseWriter
		r   *http.Request
		ds  *api.DataStore
		ses *api.SessionService
	}{
		{
			name: "success",
			w:    httptest.NewRecorder(),
			r:    req1,
			ses:  ses1,
			ds: &api.DataStore{
				Config: cu.IM{
					"NT_API_KEY": "test",
				},
				Db: &td.TestDriver{
					Config: cu.IM{
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
							return []cu.IM{{"id": 1, "name": "test", "value": "test"}}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ComparePasswordAndHash: func(password string, hash string) (err error) {
					return nil
				},
				ConvertToType: func(data interface{}, result any) (err error) {
					return nil
				},
				CreateLoginToken: func(params cu.SM, config cu.IM) (result string, err error) {
					return "test", nil
				},
			},
		},
		{
			name: "callback_url",
			w:    httptest.NewRecorder(),
			r:    req1,
			ses:  ses2,
			ds: &api.DataStore{
				Config: cu.IM{
					"NT_API_KEY": "test",
				},
				Db: &td.TestDriver{
					Config: cu.IM{
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
							return []cu.IM{{"id": 1, "name": "test", "value": "test"}}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ComparePasswordAndHash: func(password string, hash string) (err error) {
					return nil
				},
				ConvertToType: func(data interface{}, result any) (err error) {
					return nil
				},
				CreateLoginToken: func(params cu.SM, config cu.IM) (result string, err error) {
					return "test", nil
				},
			},
		},
		{
			name: "login_error",
			w:    httptest.NewRecorder(),
			r:    req1,
			ses:  ses2,
			ds: &api.DataStore{
				Config: cu.IM{
					"NT_API_KEY": "test",
				},
				Db: &td.TestDriver{
					Config: cu.IM{
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
							return []cu.IM{{"id": 1, "name": "test", "value": "test"}}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ComparePasswordAndHash: func(password string, hash string) (err error) {
					return nil
				},
				ConvertToType: func(data interface{}, result any) (err error) {
					return nil
				},
				CreateLoginToken: func(params cu.SM, config cu.IM) (result string, err error) {
					return "", errors.New("login error")
				},
			},
		},
		{
			name: "auth_error",
			w:    httptest.NewRecorder(),
			r:    req1,
			ses:  ses2,
			ds: &api.DataStore{
				Config: cu.IM{
					"NT_API_KEY": "test",
				},
				Db: &td.TestDriver{
					Config: cu.IM{},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ComparePasswordAndHash: func(password string, hash string) (err error) {
					return nil
				},
				ConvertToType: func(data interface{}, result any) (err error) {
					return nil
				},
				CreateLoginToken: func(params cu.SM, config cu.IM) (result string, err error) {
					return "test", nil
				},
			},
		},
		{
			name: "session_error",
			w:    httptest.NewRecorder(),
			r:    req1,
			ses:  &api.SessionService{},
			ds:   &api.DataStore{},
		},
		{
			name: "invalid_parameters",
			w:    httptest.NewRecorder(),
			r:    req2,
			ses:  &api.SessionService{},
			ds:   &api.DataStore{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.SessionServiceCtxKey, tt.ses)
			ctx = context.WithValue(ctx, md.DataStoreCtxKey, tt.ds)
			OAuthValidate(tt.w, tt.r.WithContext(ctx))
		})
	}
}
