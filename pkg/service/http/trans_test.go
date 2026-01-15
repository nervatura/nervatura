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

func TestTransPost(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		ds   *api.DataStore
		user *md.Auth
	}{
		{
			name: "success",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/trans", nil),
			},
			ds: &api.DataStore{
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
					return []byte(`{"product_code": "123456", "description": "test", "code": "123456",
					"trans_type": "TRANS_INVOICE","direction": "DIRECTION_OUT","auth_code": "123456", 
					"trans_date": "2024-01-01", "customer_code": "123456", "currency_code": "123456", 
					"employee_code": "123456", "project_code": "123456", "place_code": "123456", "trans_code": "123456"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"product_code": "123456", "description": "test", "code": "123456",
					"trans_type": "TRANS_INVOICE","direction": "DIRECTION_OUT","auth_code": "123456", "trans_date": "2024-01-01", 
					"customer_code": "123456", "currency_code": "123456", "employee_code": "123456", "project_code": "123456", 
					"place_code": "123456", "trans_code": "123456"}`), v)
				},
				ConvertToByte: func(v any) ([]byte, error) {
					return cu.ConvertToByte(cu.IM{})
				},
				ConvertFromByte: func(data []byte, v any) error {
					return cu.ConvertFromByte(data, v)
				},
			},
			user: &md.Auth{
				Code:      "123456",
				UserGroup: md.UserGroup(md.UserGroupAdmin),
			},
		},
		{
			name: "customer code and currency code required",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/trans", nil),
			},
			ds: &api.DataStore{
				Db: &td.TestDriver{
					Config: cu.IM{},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ReadAll: func(r io.Reader) ([]byte, error) {
					return []byte(`{"product_code": "123456", "description": "test", "code": "123456",
					"trans_type": "TRANS_INVOICE","direction": "DIRECTION_OUT","auth_code": "123456", 
					"trans_date": "2024-01-01", "currency_code": "123456", 
					"employee_code": "123456", "project_code": "123456", "place_code": "123456", "trans_code": "123456"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"product_code": "123456", "description": "test", "code": "123456",
					"trans_type": "TRANS_INVOICE","direction": "DIRECTION_OUT","auth_code": "123456", 
					"trans_date": "2024-01-01", "currency_code": "123456", 
					"employee_code": "123456", "project_code": "123456", "place_code": "123456", "trans_code": "123456"}`), v)
				},
			},
			user: &md.Auth{
				Code:      "123456",
				UserGroup: md.UserGroup(md.UserGroupAdmin),
			},
		},
		{
			name: "trans date is required",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/trans", nil),
			},
			ds: &api.DataStore{
				Db: &td.TestDriver{
					Config: cu.IM{},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ReadAll: func(r io.Reader) ([]byte, error) {
					return []byte(`{"product_code": "123456", "description": "test", "code": "123456",
					"trans_type": "TRANS_INVOICE","direction": "DIRECTION_OUT","auth_code": "123456", 
					"customer_code": "123456", "currency_code": "123456", 
					"employee_code": "123456", "project_code": "123456", "place_code": "123456", "trans_code": "123456"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"product_code": "123456", "description": "test", "code": "123456",
					"trans_type": "TRANS_INVOICE","direction": "DIRECTION_OUT","auth_code": "123456", 
					"customer_code": "123456", "currency_code": "123456", 
					"employee_code": "123456", "project_code": "123456", "place_code": "123456", "trans_code": "123456"}`), v)
				},
			},
			user: &md.Auth{
				Code:      "123456",
				UserGroup: md.UserGroup(md.UserGroupAdmin),
			},
		},
		{
			name: "unprocessable entity",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/trans", nil),
			},
			ds: &api.DataStore{
				Db: &td.TestDriver{
					Config: cu.IM{},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ReadAll: func(r io.Reader) ([]byte, error) {
					return []byte(`{"product_code": "123456", "description": "test", "code": "123456",
					"trans_type": "TRANS_INVOICE","direction": "DIRECTION_OUT","auth_code": "123456", 
					"trans_date": "2024-01-01", "customer_code": "123456", "currency_code": "123456", 
					"employee_code": "123456", "project_code": "123456", "place_code": "123456", "trans_code": "123456"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return errors.New("unprocessable entity")
				},
			},
			user: &md.Auth{
				Code:      "123456",
				UserGroup: md.UserGroup(md.UserGroupAdmin),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			ctx = context.WithValue(ctx, md.AuthUserCtxKey, tt.user)
			TransPost(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestTransPut(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		ds   *api.DataStore
		user *md.Auth
	}{
		{
			name: "success",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("PUT", "/trans/123456", nil),
			},
			ds: &api.DataStore{
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
					return []byte(`{"product_code": "123456", "description": "test", "code": "123456"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"product_code": "123456", "description": "test", "code": "123456"}`), v)
				},
				ConvertToByte: func(v any) ([]byte, error) {
					return cu.ConvertToByte(cu.IM{})
				},
				ConvertFromByte: func(data []byte, v any) error {
					return cu.ConvertFromByte(data, v)
				},
			},
			user: &md.Auth{
				Code:      "123456",
				UserGroup: md.UserGroup(md.UserGroupAdmin),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			ctx = context.WithValue(ctx, md.AuthUserCtxKey, tt.user)
			TransPut(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestTransDelete(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		ds   *api.DataStore
		user *md.Auth
	}{
		{
			name: "success",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("DELETE", "/trans/123456", nil),
			},
			ds: &api.DataStore{
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
					return []byte(`{"product_code": "123456", "description": "test", "code": "123456"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"product_code": "123456", "description": "test", "code": "123456"}`), v)
				},
				ConvertToByte: func(v any) ([]byte, error) {
					return cu.ConvertToByte(cu.IM{})
				},
				ConvertFromByte: func(data []byte, v any) error {
					return cu.ConvertFromByte(data, v)
				},
			},
			user: &md.Auth{
				Code:      "123456",
				UserGroup: md.UserGroup(md.UserGroupAdmin),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			ctx = context.WithValue(ctx, md.AuthUserCtxKey, tt.user)
			TransDelete(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestTransQuery(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		ds   *api.DataStore
		user *md.Auth
	}{
		{
			name: "success",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/trans?limit=10&offset=0&trans_type=TRANS_INVOICE&direction=DIRECTION_OUT&trans_date=2024-01-01&tag=test", nil),
			},
			ds: &api.DataStore{
				Db: &td.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1}}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ReadAll: func(r io.Reader) ([]byte, error) {
					return []byte(`{"product_code": "123456", "description": "test", "code": "123456"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"product_code": "123456", "description": "test", "code": "123456"}`), v)
				},
				ConvertToByte: func(v any) ([]byte, error) {
					return cu.ConvertToByte(cu.IM{})
				},
				ConvertFromByte: func(data []byte, v any) error {
					return cu.ConvertFromByte(data, v)
				},
			},
			user: &md.Auth{
				Code:      "123456",
				UserGroup: md.UserGroup(md.UserGroupAdmin),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			ctx = context.WithValue(ctx, md.AuthUserCtxKey, tt.user)
			TransQuery(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestTransGet(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	req := httptest.NewRequest("GET", "/trans/123456", nil)
	req.SetPathValue("id_code", "123456")
	tests := []struct {
		name string
		args args
		ds   *api.DataStore
		user *md.Auth
	}{
		{
			name: "success",
			args: args{
				w: httptest.NewRecorder(),
				r: req,
			},
			ds: &api.DataStore{
				Db: &td.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1}}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ReadAll: func(r io.Reader) ([]byte, error) {
					return []byte(`{"product_code": "123456", "description": "test", "code": "123456"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"product_code": "123456", "description": "test", "code": "123456"}`), v)
				},
				ConvertToByte: func(v any) ([]byte, error) {
					return cu.ConvertToByte(cu.IM{})
				},
				ConvertFromByte: func(data []byte, v any) error {
					return cu.ConvertFromByte(data, v)
				},
			},
			user: &md.Auth{
				Code:      "123456",
				UserGroup: md.UserGroup(md.UserGroupAdmin),
			},
		},
		{
			name: "not found",
			args: args{
				w: httptest.NewRecorder(),
				r: req,
			},
			ds: &api.DataStore{
				Db: &td.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ReadAll: func(r io.Reader) ([]byte, error) {
					return []byte(`{"product_code": "123456", "description": "test", "code": "123456"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"product_code": "123456", "description": "test", "code": "123456"}`), v)
				},
				ConvertToByte: func(v any) ([]byte, error) {
					return cu.ConvertToByte(cu.IM{})
				},
				ConvertFromByte: func(data []byte, v any) error {
					return cu.ConvertFromByte(data, v)
				},
			},
			user: &md.Auth{
				Code:      "123456",
				UserGroup: md.UserGroup(md.UserGroupAdmin),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			ctx = context.WithValue(ctx, md.AuthUserCtxKey, tt.user)
			TransGet(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}
