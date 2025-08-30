package http

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
)

func TestApiKeyAuth(t *testing.T) {
	type args struct {
		opt md.AuthOptions
	}
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("X-API-KEY", "test")
	tests := []struct {
		name        string
		args        args
		wantErrCode int
	}{
		{
			name: "valid api key",
			args: args{
				opt: md.AuthOptions{
					Request: req,
					Config: cu.IM{
						"NT_API_KEY": "test", "NT_DEFAULT_ALIAS": "test", "NT_DEFAULT_ADMIN": "test",
						"tokenKeys": []cu.SM{{"alias": "alias"}, {"user_code": "user_code"}, {"user_name": "user_name"}},
					},
				},
			},
			wantErrCode: 0,
		},
		{
			name: "invalid api key",
			args: args{
				opt: md.AuthOptions{
					Request: req,
					Config: cu.IM{
						"NT_API_KEY": "invalid", "NT_DEFAULT_ALIAS": "test", "NT_DEFAULT_ADMIN": "test",
						"tokenKeys": []cu.SM{{"alias": "alias"}, {"user_code": "user_code"}, {"user_name": "user_name"}},
					},
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				},
			},
			wantErrCode: http.StatusUnauthorized,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotErrCode := ApiKeyAuth(tt.args.opt)
			if gotErrCode != tt.wantErrCode {
				t.Errorf("ApiKeyAuth() gotErrCode = %v, want %v", gotErrCode, tt.wantErrCode)
			}
		})
	}
}

func TestTokenAuth(t *testing.T) {
	type args struct {
		opt md.AuthOptions
	}
	req := httptest.NewRequest("POST", "/", nil)
	req.Header.Set("Authorization", "Bearer test")
	tests := []struct {
		name        string
		args        args
		wantErrCode int
	}{
		{
			name: "valid token",
			args: args{
				opt: md.AuthOptions{
					Request: req,
					Config: cu.IM{
						"tokenKeys": []cu.SM{{"alias": "alias"}, {"user_code": "user_code"}, {"user_name": "user_name"}},
						"db": &md.TestDriver{
							Config: cu.IM{
								"Query": func(queries []md.Query) ([]cu.IM, error) {
									return []cu.IM{{"id": 1}}, nil
								},
							},
						},
					},
					ParseToken: func(tokenString string, keyMap []cu.SM, config cu.IM) (cu.IM, error) {
						return cu.IM{"alias": "test", "user_code": "123456", "user_name": "test"}, nil
					},
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				},
			},
			wantErrCode: 0,
		},
		{
			name: "missing token",
			args: args{
				opt: md.AuthOptions{
					Request: httptest.NewRequest("GET", "/", nil),
					Config: cu.IM{
						"tokenKeys": []cu.SM{{"alias": "alias"}, {"user_code": "user_code"}, {"user_name": "user_name"}},
					},
				},
			},
			wantErrCode: http.StatusUnauthorized,
		},
		{
			name: "invalid token",
			args: args{
				opt: md.AuthOptions{
					Request: req,
					Config: cu.IM{
						"tokenKeys": []cu.SM{{"alias": "alias"}, {"user_code": "user_code"}, {"user_name": "user_name"}},
					},
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					ParseToken: func(tokenString string, keyMap []cu.SM, config cu.IM) (cu.IM, error) {
						return cu.IM{}, errors.New("invalid token")
					},
				},
			},
			wantErrCode: http.StatusUnauthorized,
		},
		{
			name: "missing user_code",
			args: args{
				opt: md.AuthOptions{
					Request: req,
					Config: cu.IM{
						"tokenKeys": []cu.SM{{"alias": "alias"}, {"user_name": "user_name"}},
					},
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					ParseToken: func(tokenString string, keyMap []cu.SM, config cu.IM) (cu.IM, error) {
						return cu.IM{"alias": "test"}, nil
					},
				},
			},
			wantErrCode: http.StatusUnauthorized,
		},
		{
			name: "error authenticating user",
			args: args{
				opt: md.AuthOptions{
					Request: req,
					Config: cu.IM{
						"tokenKeys": []cu.SM{{"alias": "alias"}, {"user_code": "user_code"}},
						"db": &md.TestDriver{
							Config: cu.IM{"Query": func(queries []md.Query) ([]cu.IM, error) {
								return nil, errors.New("error querying")
							}},
						},
					},
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					ParseToken: func(tokenString string, keyMap []cu.SM, config cu.IM) (cu.IM, error) {
						return cu.IM{"alias": "test", "user_code": "123456", "user_name": "test"}, nil
					},
				},
			},
			wantErrCode: http.StatusUnauthorized,
		},
		{
			name: "user group is guest",
			args: args{
				opt: md.AuthOptions{
					Request: req,
					Config: cu.IM{
						"tokenKeys": []cu.SM{{"alias": "alias"}, {"user_code": "user_code"}},
						"db": &md.TestDriver{
							Config: cu.IM{"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"user_group": md.UserGroupGuest}}, nil
							}},
						},
					},
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					ParseToken: func(tokenString string, keyMap []cu.SM, config cu.IM) (cu.IM, error) {
						return cu.IM{"alias": "test", "user_code": "123456", "user_name": "test"}, nil
					},
				},
			},
			wantErrCode: http.StatusMethodNotAllowed,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotErrCode := TokenAuth(tt.args.opt)
			if gotErrCode != tt.wantErrCode {
				t.Errorf("TokenAuth() gotErrCode = %v, want %v", gotErrCode, tt.wantErrCode)
			}
		})
	}
}

func TestDatabase(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		ds   *api.DataStore
	}{
		{
			name: "success",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/database", nil),
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ReadAll: func(r io.Reader) ([]byte, error) {
					return []byte(`{"alias": "test", "demo": false}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return nil
				},
			},
		},
		{
			name: "error reading body",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/database", nil),
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ReadAll: func(r io.Reader) ([]byte, error) {
					return nil, errors.New("error reading body")
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return errors.New("error converting from reader")
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			Database(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestFunction(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		ds   *api.DataStore
	}{
		{
			name: "success",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/function", nil),
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ConvertFromReader: func(r io.Reader, v any) error {
					return nil
				},
			},
		},
		{
			name: "error converting from reader",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/function", nil),
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ConvertFromReader: func(r io.Reader, v any) error {
					return errors.New("error converting from reader")
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			Function(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestView(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		ds   *api.DataStore
	}{
		{
			name: "success",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/view", nil),
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ConvertFromReader: func(r io.Reader, v any) error {
					return nil
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			View(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestRespondMessage(t *testing.T) {
	type args struct {
		w       http.ResponseWriter
		code    int
		payload interface{}
		errCode int
		err     error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RespondMessage(tt.args.w, tt.args.code, tt.args.payload, tt.args.errCode, tt.args.err)
		})
	}
}
