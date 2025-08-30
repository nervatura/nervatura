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

func TestProjectPost(t *testing.T) {
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
				r: httptest.NewRequest("POST", "/project", nil),
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
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
				ReadAll: func(r io.Reader) ([]byte, error) {
					return []byte(`{"code": "123456", "project_name": "Project 1", "customer_code": "CUST_123456"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"code": "123456", "project_name": "Project 1", "customer_code": "CUST_123456"}`), v)
				},
				ConvertToByte: func(v any) ([]byte, error) {
					return cu.ConvertToByte(cu.IM{})
				},
				ConvertFromByte: func(data []byte, v any) error {
					return cu.ConvertFromByte(data, v)
				},
			},
		},
		{
			name: "project name is required",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/project", nil),
			},
			ds: &api.DataStore{
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
					return []byte(`{"code": "123456", "customer_code": "CUST_123456"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"code": "123456", "customer_code": "CUST_123456"}`), v)
				},
			},
		},
		{
			name: "unprocessable entity",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/project", nil),
			},
			ds: &api.DataStore{
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
					return []byte(`{"code": "123456", "customer_code": "CUST_123456"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return errors.New("unprocessable entity")
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			ProjectPost(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestProjectPut(t *testing.T) {
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
				r: httptest.NewRequest("PUT", "/project/123456", nil),
			},
			ds: &api.DataStore{
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
					return []byte(`{"code": "123456", "project_name": "Project 1", "customer_code": "CUST_123456"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"code": "123456", "project_name": "Project 1", "customer_code": "CUST_123456"}`), v)
				},
				ConvertToByte: func(v any) ([]byte, error) {
					return cu.ConvertToByte(cu.IM{})
				},
				ConvertFromByte: func(data []byte, v any) error {
					return cu.ConvertFromByte(data, v)
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			ProjectPut(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestProjectDelete(t *testing.T) {
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
				r: httptest.NewRequest("DELETE", "/project/123456", nil),
			},
			ds: &api.DataStore{
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
					return []byte(`{"code": "123456", "project_name": "Project 1", "customer_code": "CUST_123456"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"code": "123456", "project_name": "Project 1", "customer_code": "CUST_123456"}`), v)
				},
				ConvertToByte: func(v any) ([]byte, error) {
					return cu.ConvertToByte(cu.IM{})
				},
				ConvertFromByte: func(data []byte, v any) error {
					return cu.ConvertFromByte(data, v)
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			ProjectDelete(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestProjectQuery(t *testing.T) {
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
				r: httptest.NewRequest("GET", "/project?limit=10&offset=0&project_name=Project1&customer_code=CUST_123456&tag=TAG_123456", nil),
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1}}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ReadAll: func(r io.Reader) ([]byte, error) {
					return []byte(`[{"id": 1, "project_name": "Project 1", "customer_code": "CUST_123456"}]`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`[{"id": 1, "project_name": "Project 1", "customer_code": "CUST_123456"}]`), v)
				},
				ConvertToByte: func(v any) ([]byte, error) {
					return cu.ConvertToByte(cu.IM{})
				},
				ConvertFromByte: func(data []byte, v any) error {
					return cu.ConvertFromByte(data, v)
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			ProjectQuery(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestProjectGet(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	req := httptest.NewRequest("GET", "/project/123456", nil)
	req.SetPathValue("id_code", "123456")
	tests := []struct {
		name string
		args args
		ds   *api.DataStore
	}{
		{
			name: "success",
			args: args{
				w: httptest.NewRecorder(),
				r: req,
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1}}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ReadAll: func(r io.Reader) ([]byte, error) {
					return []byte(`{"id": 1, "project_name": "Project 1", "customer_code": "CUST_123456"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"id": 1, "project_name": "Project 1", "customer_code": "CUST_123456"}`), v)
				},
				ConvertToByte: func(v any) ([]byte, error) {
					return cu.ConvertToByte(cu.IM{})
				},
				ConvertFromByte: func(data []byte, v any) error {
					return cu.ConvertFromByte(data, v)
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
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ReadAll: func(r io.Reader) ([]byte, error) {
					return []byte(`{"id": 1, "project_name": "Project 1", "customer_code": "CUST_123456"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"id": 1, "project_name": "Project 1", "customer_code": "CUST_123456"}`), v)
				},
				ConvertToByte: func(v any) ([]byte, error) {
					return cu.ConvertToByte(cu.IM{})
				},
				ConvertFromByte: func(data []byte, v any) error {
					return cu.ConvertFromByte(data, v)
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			ProjectGet(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}
