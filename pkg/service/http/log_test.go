package http

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
)

func TestLogQuery(t *testing.T) {
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
				r: httptest.NewRequest("GET", "/log?limit=10&offset=0&log_type=INFO&ref_type=USER&ref_code=123456&tag=test", nil),
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
					return []byte(`{"id": 1}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"id": 1}`), v)
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
			LogQuery(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestLogGet(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	req := httptest.NewRequest("GET", "/log/123456", nil)
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
					return []byte(`{"id": 1}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"id": 1}`), v)
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
					return []byte(`{"id": 1}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"id": 1}`), v)
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
			LogGet(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}
