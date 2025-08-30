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

func TestPlacePost(t *testing.T) {
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
				r: httptest.NewRequest("POST", "/place", nil),
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
					return []byte(`{"code": "123456", "place_name": "123456", "place_type": "PLACE_WAREHOUSE", "currency_code": "USD"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"code": "123456", "place_name": "123456", "place_type": "PLACE_WAREHOUSE", "currency_code": "USD"}`), v)
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
			name: "place name is required",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/place", nil),
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ReadAll: func(r io.Reader) ([]byte, error) {
					return []byte(`{"code": "123456", "place_type": "PLACE_WAREHOUSE", "currency_code": "USD"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"code": "123456", "place_type": "PLACE_WAREHOUSE", "currency_code": "USD"}`), v)
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
			name: "unprocessable entity",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/place", nil),
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ReadAll: func(r io.Reader) ([]byte, error) {
					return []byte(`{"code": "123456", "place_name": "123456", "place_type": "PLACE_WAREHOUSE", "currency_code": "USD"}`), nil
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
			PlacePost(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestPlacePut(t *testing.T) {
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
				r: httptest.NewRequest("PUT", "/place/1", nil),
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
					return []byte(`{"code": "123456", "place_name": "123456", "place_type": "PLACE_WAREHOUSE", "currency_code": "USD"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"code": "123456", "place_name": "123456", "place_type": "PLACE_WAREHOUSE", "currency_code": "USD"}`), v)
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
			PlacePut(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestPlaceDelete(t *testing.T) {
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
				r: httptest.NewRequest("DELETE", "/place/1", nil),
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
					return []byte(`{"code": "123456", "place_name": "123456", "place_type": "PLACE_WAREHOUSE", "currency_code": "USD"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"code": "123456", "place_name": "123456", "place_type": "PLACE_WAREHOUSE", "currency_code": "USD"}`), v)
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
			PlaceDelete(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestPlaceQuery(t *testing.T) {
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
				r: httptest.NewRequest("GET", "/place?limit=10&offset=0&place_type=PLACE_WAREHOUSE&place_name=123456&tag=123456", nil),
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
					return []byte(`{"code": "123456", "place_name": "123456", "place_type": "PLACE_WAREHOUSE", "currency_code": "USD"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"code": "123456", "place_name": "123456", "place_type": "PLACE_WAREHOUSE", "currency_code": "USD"}`), v)
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
			PlaceQuery(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestPlaceGet(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	req := httptest.NewRequest("GET", "/place/1", nil)
	req.SetPathValue("id_code", "1")
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
					return []byte(`{"code": "123456", "place_name": "123456", "place_type": "PLACE_WAREHOUSE", "currency_code": "USD"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"code": "123456", "place_name": "123456", "place_type": "PLACE_WAREHOUSE", "currency_code": "USD"}`), v)
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
					return []byte(`{"code": "123456", "place_name": "123456", "place_type": "PLACE_WAREHOUSE", "currency_code": "USD"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"code": "123456", "place_name": "123456", "place_type": "PLACE_WAREHOUSE", "currency_code": "USD"}`), v)
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
			PlaceGet(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}
