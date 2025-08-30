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

func TestPricePost(t *testing.T) {
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
				r: httptest.NewRequest("POST", "/price", nil),
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
					return []byte(`{"code": "123456", "price_type": "PRICE_CUSTOMER", "currency_code": "USD", "valid_from": "2024-01-01", "qty": 1, "product_code": "123456", "valid_to": "2024-01-01", "customer_code": "123456"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"code": "123456", "price_type": "PRICE_CUSTOMER", "currency_code": "USD", "valid_from": "2024-01-01", "qty": 1, "product_code": "123456", "valid_to": "2024-01-01", "customer_code": "123456"}`), v)
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
			name: "product code is required",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/price", nil),
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ReadAll: func(r io.Reader) ([]byte, error) {
					return []byte(`{"code": "123456", "price_type": "PRICE_CUSTOMER", "currency_code": "USD", "valid_from": "2024-01-01", "qty": 1}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"code": "123456", "price_type": "PRICE_CUSTOMER", "currency_code": "USD", "valid_from": "2024-01-01", "qty": 1}`), v)
				},
			},
		},
		{
			name: "unprocessable entity",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/price", nil),
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ReadAll: func(r io.Reader) ([]byte, error) {
					return []byte(`{"code": "123456", "price_type": "PRICE_CUSTOMER", "currency_code": "USD", "valid_from": "2024-01-01", "qty": 1}`), nil
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
			PricePost(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestPricePut(t *testing.T) {
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
				r: httptest.NewRequest("PUT", "/price/123456", nil),
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
					return []byte(`{"code": "123456", "price_type": "PRICE_CUSTOMER", "currency_code": "USD", "valid_from": "2024-01-01", "qty": 1, "product_code": "123456", "valid_to": "2024-01-01"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"code": "123456", "price_type": "PRICE_CUSTOMER", "currency_code": "USD", "valid_from": "2024-01-01", "qty": 1, "product_code": "123456", "valid_to": "2024-01-01"}`), v)
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
			PricePut(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestPriceDelete(t *testing.T) {
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
				r: httptest.NewRequest("DELETE", "/price/123456", nil),
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
					return []byte(`{"code": "123456", "price_type": "PRICE_CUSTOMER", "currency_code": "USD", "valid_from": "2024-01-01", "qty": 1, "product_code": "123456", "valid_to": "2024-01-01"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"code": "123456", "price_type": "PRICE_CUSTOMER", "currency_code": "USD", "valid_from": "2024-01-01", "qty": 1, "product_code": "123456", "valid_to": "2024-01-01"}`), v)
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
			PriceDelete(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestPriceQuery(t *testing.T) {
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
				r: httptest.NewRequest("GET", "/price?limit=10&offset=0&price_type=PRICE_CUSTOMER&product_code=123456&currency_code=USD&tag=tag1", nil),
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
					return []byte(`{"code": "123456", "price_type": "PRICE_CUSTOMER", "currency_code": "USD", "valid_from": "2024-01-01", "qty": 1, "product_code": "123456", "valid_to": "2024-01-01"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"code": "123456", "price_type": "PRICE_CUSTOMER", "currency_code": "USD", "valid_from": "2024-01-01", "qty": 1, "product_code": "123456", "valid_to": "2024-01-01"}`), v)
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
			PriceQuery(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}

func TestPriceGet(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	req := httptest.NewRequest("GET", "/price/123456", nil)
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
					return []byte(`{"code": "123456", "price_type": "PRICE_CUSTOMER", "currency_code": "USD", "valid_from": "2024-01-01", "qty": 1, "product_code": "123456", "valid_to": "2024-01-01"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"code": "123456", "price_type": "PRICE_CUSTOMER", "currency_code": "USD", "valid_from": "2024-01-01", "qty": 1, "product_code": "123456", "valid_to": "2024-01-01"}`), v)
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
					return []byte(`{"code": "123456", "price_type": "PRICE_CUSTOMER", "currency_code": "USD", "valid_from": "2024-01-01", "qty": 1, "product_code": "123456", "valid_to": "2024-01-01"}`), nil
				},
				ConvertFromReader: func(r io.Reader, v any) error {
					return cu.ConvertFromByte([]byte(`{"code": "123456", "price_type": "PRICE_CUSTOMER", "currency_code": "USD", "valid_from": "2024-01-01", "qty": 1, "product_code": "123456", "valid_to": "2024-01-01"}`), v)
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
			PriceGet(tt.args.w, tt.args.r.WithContext(ctx))
		})
	}
}
