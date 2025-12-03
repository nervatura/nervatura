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

func TestAuthPost(t *testing.T) {
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
				r: httptest.NewRequest("POST", "/auth", nil),
			},
			ds: &api.DataStore{
				Config: cu.IM{},
				Db: &md.TestDriver{
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
				ConvertFromReader: func(data io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"code": "123456", "user_name": "test", "user_group": "GROUP_ADMIN"}`), v)
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
			name: "unauthorized",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/auth", nil),
			},
			ds: &api.DataStore{
				Config: cu.IM{},
				Db:     &md.TestDriver{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			user: md.Auth{
				UserGroup: md.UserGroupUser,
			},
		},
		{
			name: "unprocessable entity",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/auth", nil),
			},
			ds: &api.DataStore{
				Config: cu.IM{},
				Db:     &md.TestDriver{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ConvertFromReader: func(data io.Reader, v any) error {
					return errors.New("unprocessable entity")
				},
			},
			user: md.Auth{
				UserGroup: md.UserGroupAdmin,
			},
		},
		{
			name: "missing user name",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/auth", nil),
			},
			ds: &api.DataStore{
				Config: cu.IM{},
				Db:     &md.TestDriver{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ConvertFromReader: func(data io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"user_group": "GROUP_ADMIN"}`), v)
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
			AuthPost(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestAuthPut(t *testing.T) {
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
				r: httptest.NewRequest("PUT", "/auth/1", nil),
			},
			ds: &api.DataStore{
				Config: cu.IM{},
				Db: &md.TestDriver{
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
					return []byte(`{"user_group": "GROUP_ADMIN"}`), nil
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
				r: httptest.NewRequest("PUT", "/auth/1", nil),
			},
			ds: &api.DataStore{
				Config: cu.IM{},
				Db:     &md.TestDriver{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			user: md.Auth{
				Id:        1,
				Code:      "123456",
				UserGroup: md.UserGroupUser,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			ctx = context.WithValue(ctx, md.AuthUserCtxKey, tt.user)
			AuthPut(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestAuthQuery(t *testing.T) {
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
				r: httptest.NewRequest("GET", "/auth?limit=10&offset=0&user_group=GROUP_ADMIN&tag=test", nil),
			},
			ds: &api.DataStore{
				Config: cu.IM{},
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1}}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ReadAll: func(r io.Reader) ([]byte, error) {
					return []byte(`{"user_group": "GROUP_ADMIN"}`), nil
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
				r: httptest.NewRequest("GET", "/auth", nil),
			},
			ds: &api.DataStore{
				Config: cu.IM{},
				Db:     &md.TestDriver{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
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
			AuthQuery(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestAuthGet(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	req := httptest.NewRequest("GET", "/auth/1", nil)
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
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1}}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ReadAll: func(r io.Reader) ([]byte, error) {
					return []byte(`{"user_group": "GROUP_ADMIN"}`), nil
				},
				ConvertFromByte: func(data []byte, v any) error {
					return cu.ConvertFromByte(data, v)
				},
				ConvertToByte: func(v any) ([]byte, error) {
					return cu.ConvertToByte(cu.IM{})
				},
			},
			user: md.Auth{
				Id:        1,
				Code:      "111111",
				UserGroup: md.UserGroupUser,
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
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ReadAll: func(r io.Reader) ([]byte, error) {
					return []byte(`{"user_group": "GROUP_ADMIN"}`), nil
				},
				ConvertFromByte: func(data []byte, v any) error {
					return cu.ConvertFromByte(data, v)
				},
				ConvertToByte: func(v any) ([]byte, error) {
					return cu.ConvertToByte(cu.IM{})
				},
			},
			user: md.Auth{
				Id:        1,
				Code:      "111111",
				UserGroup: md.UserGroupAdmin,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			ctx = context.WithValue(ctx, md.AuthUserCtxKey, tt.user)
			AuthGet(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestAuthLogin(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		auth md.AuthOptions
	}{
		{
			name: "success",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/auth/login", nil),
			},
			auth: md.AuthOptions{
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ConvertFromReader: func(data io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"user_name": "test", "password": "123456"}`), v)
				},
			},
		},
		{
			name: "unprocessable entity",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/auth/login", nil),
			},
			auth: md.AuthOptions{
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ConvertFromReader: func(data io.Reader, v any) error {
					return errors.New("unprocessable entity")
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.AuthOptionsCtxKey, tt.auth)
			AuthLogin(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestAuthPassword(t *testing.T) {
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
				r: httptest.NewRequest("POST", "/auth/password", nil),
			},
			ds: &api.DataStore{
				Config: cu.IM{},
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1}}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ConvertFromReader: func(data io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"password": "123456", "confirm": "123456"}`), v)
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
			name: "unprocessable entity",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/auth/password", nil),
			},
			ds: &api.DataStore{
				Config: cu.IM{},
				Db:     &md.TestDriver{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ConvertFromReader: func(data io.Reader, v any) error {
					return errors.New("unprocessable entity")
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			ctx = context.WithValue(ctx, md.AuthUserCtxKey, tt.user)
			AuthPassword(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestAuthReset(t *testing.T) {
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
				r: httptest.NewRequest("POST", "/auth/reset", nil),
			},
			ds: &api.DataStore{
				Config: cu.IM{},
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1}}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ReadAll: func(r io.Reader) ([]byte, error) {
					return []byte(`{"user_group": "GROUP_ADMIN"}`), nil
				},
				ConvertFromByte: func(data []byte, v any) error {
					return cu.ConvertFromByte(data, v)
				},
				ConvertToByte: func(v any) ([]byte, error) {
					return cu.ConvertToByte(cu.IM{})
				},
			},
			user: md.Auth{
				Id:        1,
				Code:      "111111",
				UserGroup: md.UserGroupUser,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			ctx = context.WithValue(ctx, md.AuthUserCtxKey, tt.user)
			AuthReset(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestAuthToken(t *testing.T) {
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
				r: httptest.NewRequest("GET", "/auth/token", nil),
			},
			ds: &api.DataStore{
				Config: cu.IM{},
				Db:     &md.TestDriver{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				CreateLoginToken: func(params cu.SM, config cu.IM) (string, error) {
					return "token", nil
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
			AuthToken(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}
