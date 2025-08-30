package grpc

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log/slog"
	"testing"

	cu "github.com/nervatura/component/pkg/util"
	api "github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	pb "github.com/nervatura/nervatura/v6/pkg/service/grpc/proto"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

func TestGService_RateUpdate(t *testing.T) {
	type fields struct {
		Config                 cu.IM
		AppLog                 *slog.Logger
		DecodeAuthToken        func(tokenString string) (cu.IM, error)
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx  context.Context
		req  *pb.Rate
		ds   *api.DataStore
		user md.Auth
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx:  context.Background(),
				req:  &pb.Rate{Id: 1, Code: "1234", RateType: pb.RateType_RATE_AVERAGE, RateDate: "2024-01-01", PlaceCode: "1234", CurrencyCode: "USD"},
				user: md.Auth{UserGroup: md.UserGroupAdmin},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					Db: &md.TestDriver{
						Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1, "code": "1234", "rate_object": `{"id": 1, "code": "1234"}`}}, nil
							},
							"Update": func(data md.Update) (int64, error) {
								return 1, nil
							},
						},
					},
					ReadAll: func(r io.Reader) ([]byte, error) {
						return []byte(`{"id": 1, "code": "1234"}`), nil
					},
					ConvertFromReader: func(r io.Reader, v any) error {
						return cu.ConvertFromByte([]byte(`{"id": 1, "code": "1234"}`), v)
					},
					ConvertFromByte: func(data []byte, v any) error {
						return cu.ConvertFromByte(data, v)
					},
					ConvertToByte: func(v any) ([]byte, error) {
						return cu.ConvertToByte(v)
					},
					ConvertToType: func(data interface{}, result any) error {
						return ut.ConvertToType(data, result)
					},
				},
			},
			wantErr: false,
		},
		{
			name: "unprocessable entity",
			args: args{
				ctx:  context.Background(),
				req:  &pb.Rate{Id: 1, Code: "1234", RateType: pb.RateType_RATE_AVERAGE, RateDate: "2024-01-01", PlaceCode: "1234", CurrencyCode: "USD"},
				user: md.Auth{UserGroup: md.UserGroupAdmin},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
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
					ReadAll: func(r io.Reader) ([]byte, error) {
						return []byte(`{"id": 1, "code": "1234"}`), nil
					},
					ConvertFromReader: func(r io.Reader, v any) error {
						return cu.ConvertFromByte([]byte(`{"id": 1, "code": "1234"}`), v)
					},
					ConvertFromByte: func(data []byte, v any) error {
						return cu.ConvertFromByte(data, v)
					},
					ConvertToByte: func(v any) ([]byte, error) {
						return cu.ConvertToByte(v)
					},
					ConvertToType: func(data interface{}, result any) error {
						return ut.ConvertToType(data, result)
					},
				},
			},
			wantErr: true,
		},
		{
			name: "error update",
			args: args{
				ctx:  context.Background(),
				req:  &pb.Rate{Id: 0, Code: "1234", RateType: pb.RateType_RATE_AVERAGE, RateDate: "2024-01-01", PlaceCode: "1234", CurrencyCode: "USD"},
				user: md.Auth{UserGroup: md.UserGroupAdmin},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					Db: &md.TestDriver{
						Config: cu.IM{
							"Update": func(data md.Update) (int64, error) {
								return 0, errors.New("error")
							},
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{}, nil
							},
						},
					},
					ReadAll: func(r io.Reader) ([]byte, error) {
						return []byte(`{"id": 1, "code": "1234"}`), nil
					},
					ConvertFromReader: func(r io.Reader, v any) error {
						return cu.ConvertFromByte([]byte(`{"id": 1, "code": "1234"}`), v)
					},
					ConvertFromByte: func(data []byte, v any) error {
						return cu.ConvertFromByte(data, v)
					},
					ConvertToByte: func(v any) ([]byte, error) {
						return cu.ConvertToByte(v)
					},
					ConvertToType: func(data interface{}, result any) error {
						return ut.ConvertToType(data, result)
					},
				},
			},
			wantErr: true,
		},
		{
			name: "error get data by id",
			args: args{
				ctx:  context.Background(),
				req:  &pb.Rate{Id: 0, Code: "1234", RateType: pb.RateType_RATE_AVERAGE, RateDate: "2024-01-01", PlaceCode: "1234", CurrencyCode: "USD"},
				user: md.Auth{UserGroup: md.UserGroupAdmin},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					Db: &md.TestDriver{
						Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{}, errors.New("error")
							},
						},
					},
					ReadAll: func(r io.Reader) ([]byte, error) {
						return []byte(`{"id": 1, "code": "1234"}`), nil
					},
					ConvertFromReader: func(r io.Reader, v any) error {
						return cu.ConvertFromByte([]byte(`{"id": 1, "code": "1234"}`), v)
					},
					ConvertFromByte: func(data []byte, v any) error {
						return cu.ConvertFromByte(data, v)
					},
					ConvertToByte: func(v any) ([]byte, error) {
						return cu.ConvertToByte(v)
					},
					ConvertToType: func(data interface{}, result any) error {
						return ut.ConvertToType(data, result)
					},
				},
			},
			wantErr: true,
		},
		{
			name: "rate date, place code and currency code are required",
			args: args{
				ctx:  context.Background(),
				req:  &pb.Rate{Id: 0, Code: "1234", RateType: pb.RateType_RATE_AVERAGE, RateDate: "", PlaceCode: "1234", CurrencyCode: "USD"},
				user: md.Auth{UserGroup: md.UserGroupAdmin},
			},
			wantErr: true,
		},
		{
			name: "guest user",
			args: args{
				ctx:  context.Background(),
				req:  &pb.Rate{Id: 0, Code: "1234", RateType: pb.RateType_RATE_AVERAGE, RateDate: "2024-01-01", PlaceCode: "1234", CurrencyCode: "USD"},
				user: md.Auth{UserGroup: md.UserGroupGuest},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &GService{
				Config:                 tt.fields.Config,
				AppLog:                 tt.fields.AppLog,
				DecodeAuthToken:        tt.fields.DecodeAuthToken,
				UnimplementedAPIServer: tt.fields.UnimplementedAPIServer,
			}
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.args.ds)
			ctx = context.WithValue(ctx, md.AuthUserCtxKey, tt.args.user)
			_, err := s.RateUpdate(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GService.RateUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGService_RateQuery(t *testing.T) {
	type fields struct {
		Config                 cu.IM
		AppLog                 *slog.Logger
		DecodeAuthToken        func(tokenString string) (cu.IM, error)
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx context.Context
		req *pb.RequestQuery
		ds  *api.DataStore
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: &pb.RequestQuery{Limit: 10, Offset: 0, Filters: []*pb.RequestQueryFilter{
					{Field: "rate_type", Value: "average"},
					{Field: "currency_code", Value: "USD"},
					{Field: "tag", Value: "test"},
				}},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					Db: &md.TestDriver{
						Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1, "code": "1234", "rate_object": `{"id": 1, "code": "1234"}`}}, nil
							},
						},
					},
					ReadAll: func(r io.Reader) ([]byte, error) {
						return []byte(`{"id": 1, "code": "1234"}`), nil
					},
					ConvertFromReader: func(r io.Reader, v any) error {
						return cu.ConvertFromByte([]byte(`{"id": 1, "code": "1234"}`), v)
					},
					ConvertFromByte: func(data []byte, v any) error {
						return cu.ConvertFromByte(data, v)
					},
					ConvertToByte: func(v any) ([]byte, error) {
						return cu.ConvertToByte(v)
					},
					ConvertToType: func(data interface{}, result any) error {
						return ut.ConvertToType(data, result)
					},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid filter field",
			args: args{
				ctx: context.Background(),
				req: &pb.RequestQuery{Limit: 10, Offset: 0, Filters: []*pb.RequestQueryFilter{{Field: "invalid", Value: "test"}}},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &GService{
				Config:                 tt.fields.Config,
				AppLog:                 tt.fields.AppLog,
				DecodeAuthToken:        tt.fields.DecodeAuthToken,
				UnimplementedAPIServer: tt.fields.UnimplementedAPIServer,
			}
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.args.ds)
			_, err := s.RateQuery(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GService.RateQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
