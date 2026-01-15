package http

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"path"
	"testing"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	st "github.com/nervatura/nervatura/v6/pkg/static"
	td "github.com/nervatura/nervatura/v6/test/driver"
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
	req := httptest.NewRequest("POST", st.ApiPath, nil)
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
						"db": &td.TestDriver{
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
						"db": &td.TestDriver{
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
						"db": &td.TestDriver{
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
				Db: &td.TestDriver{
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
				Db: &td.TestDriver{
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
	pdf_json, _ := st.Report.ReadFile(path.Join("template", "ntr_customer_en.json"))
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
				Db: &td.TestDriver{
					Config: cu.IM{},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ConvertFromReader: func(r io.Reader, v any) error {
					return nil
				},
			},
		},
		{
			name: "pdf response",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/function", bytes.NewBufferString(`{"name": "report_get", "values": {"report_id": 1}}`)),
			},
			ds: &api.DataStore{
				Db: &td.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "data": cu.IM{"file_type": "FILE_PDF"}}}, nil
						},
						"QuerySQL": func(sqlString string) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "data": cu.IM{"file_type": "FILE_PDF"}}}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ConvertFromReader: func(r io.Reader, v any) error {
					options := cu.IM{
						"name": "report_get",
						"values": cu.IM{
							"report_key":  "ntr_customer_en",
							"orientation": "portrait",
							"size":        "a4",
							"code":        "test",
							"template":    string(pdf_json),
							"filters":     cu.IM{},
						},
					}
					return cu.ConvertToType(options, v)
				},
				ConvertFromByte: func(data []byte, result interface{}) error {
					return cu.ConvertFromByte(data, result)
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
				Db: &td.TestDriver{
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
				Db: &td.TestDriver{
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

func TestHealth(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w http.ResponseWriter
		r *http.Request
	}{
		{
			name: "success",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest("GET", "/health", nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Health(tt.w, tt.r)
		})
	}
}
