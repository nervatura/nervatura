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
	td "github.com/nervatura/nervatura/v6/test/driver"
)

func TestConfigPost(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		ds   *api.DataStore
		user md.Auth
	}{
		{
			name: "success",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/config", nil),
			},
			ds: &api.DataStore{
				Config: cu.IM{},
				Db: &td.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1}}, nil
						},
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"code": "123456", "user_name": "test", "user_group": "GROUP_ADMIN"}`), v)
				},
				ConvertToByte: func(v any) ([]byte, error) {
					return cu.ConvertToByte(cu.IM{})
				},
				ConvertFromByte: func(b []byte, v any) error {
					return cu.ConvertFromByte(b, v)
				},
			},
			user: md.Auth{
				UserGroup: md.UserGroupAdmin,
			},
		},
		{
			name: "unprocessable entity",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/config", nil),
			},
			ds: &api.DataStore{
				Config: cu.IM{},
				Db: &td.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ConvertFromReader: func(r io.Reader, v any) error {
					return errors.New("unprocessable entity")
				},
				ConvertToByte: func(v any) ([]byte, error) {
					return cu.ConvertToByte(cu.IM{})
				},
				ConvertFromByte: func(b []byte, v any) error {
					return cu.ConvertFromByte(b, v)
				},
			},
			user: md.Auth{
				UserGroup: md.UserGroupAdmin,
			},
		},
		{
			name: "print queue not allowed",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/config", nil),
			},
			ds: &api.DataStore{
				Config: cu.IM{},
				Db:     &td.TestDriver{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"code": "123456", "user_name": "test", "user_group": "GROUP_USER", "config_type": "CONFIG_MAP"}`), v)
				},
				ConvertToByte: func(v any) ([]byte, error) {
					return cu.ConvertToByte(cu.IM{})
				},
				ConvertFromByte: func(b []byte, v any) error {
					return cu.ConvertFromByte(b, v)
				},
			},
			user: md.Auth{
				UserGroup: md.UserGroupUser,
			},
		},
		{
			name: "schema validation error",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/config", nil),
			},
			ds: &api.DataStore{
				Config: cu.IM{},
				Db:     &td.TestDriver{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"code": "123456", "user_name": "test", "user_group": "GROUP_ADMIN", "config_type": "CONFIG_MAP"}`), v)
				},
				ConvertToByte: func(v any) ([]byte, error) {
					return nil, errors.New("schema validation error")
				},
				ConvertFromByte: func(b []byte, v any) error {
					return cu.ConvertFromByte(b, v)
				},
			},
			user: md.Auth{
				UserGroup: md.UserGroupAdmin,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			ctx = context.WithValue(ctx, md.AuthUserCtxKey, tt.user)
			ConfigPost(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestConfigPut(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		ds   *api.DataStore
		user md.Auth
	}{
		{
			name: "success",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("PUT", "/config/123456", nil),
			},
			ds: &api.DataStore{
				Config: cu.IM{},
				Db: &td.TestDriver{
					Config: cu.IM{
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1}}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ReadAll: func(r io.Reader) ([]byte, error) {
					return []byte(`{"code": "123456", "user_group": "GROUP_ADMIN", "config_type": "CONFIG_PRINT_QUEUE", "data": {}}`), nil
				},
				ConvertFromByte: func(data []byte, v any) error {
					return cu.ConvertFromByte(data, v)
				},
				ConvertToByte: func(v any) ([]byte, error) {
					return cu.ConvertToByte(cu.IM{})
				},
			},
			user: md.Auth{
				UserGroup: md.UserGroupAdmin,
			},
		},
		{
			name: "config type not found",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("PUT", "/config/123456", nil),
			},
			ds: &api.DataStore{
				Config: cu.IM{},
				Db:     &td.TestDriver{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ReadAll: func(r io.Reader) ([]byte, error) {
					return []byte(`{"code": "123456", "user_group": "GROUP_ADMIN"}`), nil
				},
				ConvertFromByte: func(data []byte, v any) error {
					return cu.ConvertFromByte(data, v)
				},
				ConvertToByte: func(v any) ([]byte, error) {
					return cu.ConvertToByte(cu.IM{})
				},
			},
			user: md.Auth{
				UserGroup: md.UserGroupAdmin,
			},
		},
		{
			name: "print queue not allowed",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("PUT", "/config/123456", nil),
			},
			ds: &api.DataStore{
				Config: cu.IM{},
				Db:     &td.TestDriver{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ReadAll: func(r io.Reader) ([]byte, error) {
					return []byte(`{"code": "123456", "user_group": "GROUP_ADMIN", "config_type": "CONFIG_MAP"}`), nil
				},
				ConvertFromByte: func(data []byte, v any) error {
					return cu.ConvertFromByte(data, v)
				},
				ConvertToByte: func(v any) ([]byte, error) {
					return cu.ConvertToByte(cu.IM{})
				},
			},
			user: md.Auth{
				UserGroup: md.UserGroupUser,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			ctx = context.WithValue(ctx, md.AuthUserCtxKey, tt.user)
			ConfigPut(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestConfigDelete(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	req := httptest.NewRequest("DELETE", "/config/123456", nil)
	req.SetPathValue("id_code", "123456")
	tests := []struct {
		name string
		args args
		ds   *api.DataStore
		user md.Auth
	}{
		{
			name: "success",
			args: args{
				w: httptest.NewRecorder(),
				r: req,
			},
			ds: &api.DataStore{
				Config: cu.IM{},
				Db: &td.TestDriver{
					Config: cu.IM{
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1}}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ReadAll: func(r io.Reader) ([]byte, error) {
					return []byte(`{"code": "123456", "user_group": "GROUP_ADMIN"}`), nil
				},
				ConvertFromByte: func(data []byte, v any) error {
					return cu.ConvertFromByte(data, v)
				},
				ConvertToByte: func(v any) ([]byte, error) {
					return cu.ConvertToByte(cu.IM{})
				},
			},
			user: md.Auth{
				UserGroup: md.UserGroupAdmin,
			},
		},
		{
			name: "not allowed",
			args: args{
				w: httptest.NewRecorder(),
				r: req,
			},
			user: md.Auth{
				UserGroup: md.UserGroupUser,
			},
			ds: &api.DataStore{
				Config: cu.IM{},
				Db: &td.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1}}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ReadAll: func(r io.Reader) ([]byte, error) {
					return []byte(`{"code": "123456", "user_group": "GROUP_ADMIN", "config_type": "CONFIG_MAP"}`), nil
				},
				ConvertFromByte: func(data []byte, v any) error {
					return cu.ConvertFromByte(data, v)
				},
				ConvertToByte: func(v any) ([]byte, error) {
					return cu.ConvertToByte(cu.IM{})
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			ctx = context.WithValue(ctx, md.AuthUserCtxKey, tt.user)
			ConfigDelete(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestConfigQuery(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		ds   *api.DataStore
		user md.Auth
	}{
		{
			name: "success",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/config?limit=10&offset=0&config_type=CONFIG_MAP", nil),
			},
			ds: &api.DataStore{
				Config: cu.IM{},
				Db: &td.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1}}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ReadAll: func(r io.Reader) ([]byte, error) {
					return []byte(`{"code": "123456", "user_group": "GROUP_ADMIN"}`), nil
				},
				ConvertFromByte: func(data []byte, v any) error {
					return cu.ConvertFromByte(data, v)
				},
				ConvertToByte: func(v any) ([]byte, error) {
					return cu.ConvertToByte(cu.IM{})
				},
			},
			user: md.Auth{
				UserGroup: md.UserGroupAdmin,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			ctx = context.WithValue(ctx, md.AuthUserCtxKey, tt.user)
			ConfigQuery(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestConfigGet(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	req := httptest.NewRequest("GET", "/config/123456", nil)
	req.SetPathValue("id_code", "123456")
	tests := []struct {
		name string
		args args
		ds   *api.DataStore
		user md.Auth
	}{
		{
			name: "success",
			args: args{
				w: httptest.NewRecorder(),
				r: req,
			},
			ds: &api.DataStore{
				Config: cu.IM{},
				Db: &td.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1}}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ReadAll: func(r io.Reader) ([]byte, error) {
					return []byte(`{"code": "123456", "user_group": "GROUP_ADMIN"}`), nil
				},
				ConvertFromByte: func(data []byte, v any) error {
					return cu.ConvertFromByte(data, v)
				},
				ConvertToByte: func(v any) ([]byte, error) {
					return cu.ConvertToByte(cu.IM{})
				},
			},
		},
		{
			name: "not found",
			args: args{
				w: httptest.NewRecorder(),
				r: req,
			},
			ds: &api.DataStore{
				Config: cu.IM{},
				Db: &td.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ReadAll: func(r io.Reader) ([]byte, error) {
					return []byte(`{"code": "123456", "user_group": "GROUP_ADMIN"}`), nil
				},
				ConvertFromByte: func(data []byte, v any) error {
					return cu.ConvertFromByte(data, v)
				},
				ConvertToByte: func(v any) ([]byte, error) {
					return cu.ConvertToByte(cu.IM{})
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			ctx = context.WithValue(ctx, md.AuthUserCtxKey, tt.user)
			ConfigGet(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}
