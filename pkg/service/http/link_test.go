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

func TestLinkPost(t *testing.T) {
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
				r: httptest.NewRequest("POST", "/link", nil),
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
					return []byte(`{"code": "123456", "link_type_1": "LINK_ITEM", "link_code_1": "123456", "link_type_2": "LINK_ITEM", "link_code_2": "123456"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"code": "123456", "link_type_1": "LINK_ITEM", "link_code_1": "123456", "link_type_2": "LINK_ITEM", "link_code_2": "123456"}`), v)
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
			name: "unprocessable entity",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/link", nil),
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ReadAll: func(r io.Reader) ([]byte, error) {
					return []byte(`{"code": "123456", "link_type_1": "LINK_ITEM", "link_code_1": "123456", "link_type_2": "LINK_ITEM", "link_code_2": "123456"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return errors.New("unprocessable entity")
				},
			},
		},
		{
			name: "bad request",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/link", nil),
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ReadAll: func(r io.Reader) ([]byte, error) {
					return []byte(`{"code": "123456", "link_type_1": "LINK_ITEM", "link_type_2": "LINK_ITEM"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"code": "123456", "link_type_1": "LINK_ITEM", "link_type_2": "LINK_ITEM"}`), v)
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
			LinkPost(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestLinkPut(t *testing.T) {
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
				r: httptest.NewRequest("PUT", "/link/123456", nil),
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
					return []byte(`{"code": "123456", "link_type_1": "LINK_ITEM", "link_code_1": "123456", "link_type_2": "LINK_ITEM", "link_code_2": "123456"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"code": "123456", "link_type_1": "LINK_ITEM", "link_type_2": "LINK_ITEM"}`), v)
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
			LinkPut(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestLinkDelete(t *testing.T) {
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
				r: httptest.NewRequest("DELETE", "/link/123456", nil),
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
					return []byte(`{"code": "123456", "link_type_1": "LINK_ITEM", "link_type_2": "LINK_ITEM"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"code": "123456", "link_type_1": "LINK_ITEM", "link_type_2": "LINK_ITEM"}`), v)
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
			LinkDelete(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestLinkQuery(t *testing.T) {
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
				r: httptest.NewRequest("GET", "/link?limit=10&offset=0&link_type_1=LINK_ITEM&link_code_1=123456&link_type_2=LINK_ITEM&link_code_2=123456&tag=tag1", nil),
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
					return []byte(`{"code": "123456", "link_type_1": "LINK_ITEM", "link_type_2": "LINK_ITEM"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"code": "123456", "link_type_1": "LINK_ITEM", "link_type_2": "LINK_ITEM"}`), v)
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
			LinkQuery(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestLinkGet(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	req := httptest.NewRequest("GET", "/link/123456", nil)
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
					return []byte(`{"code": "123456", "link_type_1": "LINK_ITEM", "link_type_2": "LINK_ITEM"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"code": "123456", "link_type_1": "LINK_ITEM", "link_type_2": "LINK_ITEM"}`), v)
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
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ReadAll: func(r io.Reader) ([]byte, error) {
					return []byte(`{"code": "123456", "link_type_1": "LINK_ITEM", "link_type_2": "LINK_ITEM"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"code": "123456", "link_type_1": "LINK_ITEM", "link_type_2": "LINK_ITEM"}`), v)
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
			LinkGet(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}
