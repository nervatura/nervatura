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
	td "github.com/nervatura/nervatura/v6/test/driver"
)

func TestGService_PlaceUpdate(t *testing.T) {
	type fields struct {
		Config                 cu.IM
		AppLog                 *slog.Logger
		DecodeAuthToken        func(tokenString string) (cu.IM, error)
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx  context.Context
		req  *pb.Place
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
				req:  &pb.Place{Id: 1, Code: "1234", PlaceName: "1234", PlaceType: pb.PlaceType_PLACE_BANK, CurrencyCode: "USD"},
				user: md.Auth{UserGroup: md.UserGroupAdmin},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					Db: &td.TestDriver{
						Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1, "code": "1234", "place_name": "1234", "place_type": "bank", "place_object": `{"id": 1, "code": "1234"}`}}, nil
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
					ConvertToByte: func(v any) ([]byte, error) {
						return []byte(`{"id": 1, "code": "1234"}`), nil
					},
					ConvertToType: func(v any, t any) error {
						return ut.ConvertToType(v, t)
					},
					ConvertFromByte: func(data []byte, result interface{}) error {
						return cu.ConvertFromByte(data, result)
					},
				},
			},
			wantErr: false,
		},
		{
			name: "unprocessable entity",
			args: args{
				ctx:  context.Background(),
				req:  &pb.Place{Id: 1, Code: "1234", PlaceName: "Place", PlaceType: pb.PlaceType_PLACE_BANK, CurrencyCode: "USD"},
				user: md.Auth{UserGroup: md.UserGroupAdmin},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
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
					ReadAll: func(r io.Reader) ([]byte, error) {
						return []byte(`{"id": 1, "code": "1234"}`), nil
					},
					ConvertFromReader: func(r io.Reader, v any) error {
						return cu.ConvertFromByte([]byte(`{"id": 1, "code": "1234"}`), v)
					},
					ConvertToByte: func(v any) ([]byte, error) {
						return []byte(`{"id": 1, "code": "1234"}`), nil
					},
					ConvertToType: func(v any, t any) error {
						return ut.ConvertToType(v, t)
					},
					ConvertFromByte: func(data []byte, result interface{}) error {
						return cu.ConvertFromByte(data, result)
					},
				},
			},
			wantErr: true,
		},
		{
			name: "error update",
			args: args{
				ctx:  context.Background(),
				req:  &pb.Place{Id: 0, Code: "1234", PlaceName: "Place", PlaceType: pb.PlaceType_PLACE_BANK, CurrencyCode: "USD"},
				user: md.Auth{UserGroup: md.UserGroupAdmin},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					Db: &td.TestDriver{
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
					ConvertToByte: func(v any) ([]byte, error) {
						return []byte(`{"id": 1, "code": "1234"}`), nil
					},
					ConvertToType: func(v any, t any) error {
						return ut.ConvertToType(v, t)
					},
					ConvertFromByte: func(data []byte, result interface{}) error {
						return cu.ConvertFromByte(data, result)
					},
				},
			},
			wantErr: true,
		},
		{
			name: "get entity error",
			args: args{
				ctx:  context.Background(),
				req:  &pb.Place{Id: 1, Code: "1234", PlaceName: "Place", PlaceType: pb.PlaceType_PLACE_BANK, CurrencyCode: "USD"},
				user: md.Auth{UserGroup: md.UserGroupAdmin},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					Db: &td.TestDriver{
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
					ConvertToByte: func(v any) ([]byte, error) {
						return []byte(`{"id": 1, "code": "1234"}`), nil
					},
					ConvertToType: func(v any, t any) error {
						return ut.ConvertToType(v, t)
					},
					ConvertFromByte: func(data []byte, result interface{}) error {
						return cu.ConvertFromByte(data, result)
					},
				},
			},
			wantErr: true,
		},
		{
			name: "place name is required",
			args: args{
				ctx:  context.Background(),
				req:  &pb.Place{Id: 1, Code: "1234", PlaceType: pb.PlaceType_PLACE_BANK, CurrencyCode: "USD"},
				user: md.Auth{UserGroup: md.UserGroupAdmin},
			},
			wantErr: true,
		},
		{
			name: "guest user",
			args: args{
				ctx:  context.Background(),
				req:  &pb.Place{Id: 1, Code: "1234", PlaceName: "Place", PlaceType: pb.PlaceType_PLACE_BANK, CurrencyCode: "USD"},
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
			_, err := s.PlaceUpdate(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GService.PlaceUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGService_PlaceQuery(t *testing.T) {
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
				req: &pb.RequestQuery{Limit: 10, Offset: 0, Filters: []*pb.RequestQueryFilter{{Field: "place_name", Value: "Place"}}},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					Db: &td.TestDriver{
						Config: cu.IM{"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "1234", "place_name": "1234", "place_type": "bank", "place_object": `{"id": 1, "code": "1234"}`}}, nil
						}},
					},
					ReadAll: func(r io.Reader) ([]byte, error) {
						return []byte(`{"id": 1, "code": "1234"}`), nil
					},
					ConvertFromReader: func(r io.Reader, v any) error {
						return cu.ConvertFromByte([]byte(`{"id": 1, "code": "1234"}`), v)
					},
					ConvertToByte: func(v any) ([]byte, error) {
						return []byte(`{"id": 1, "code": "1234"}`), nil
					},
					ConvertToType: func(v any, t any) error {
						return ut.ConvertToType(v, t)
					},
					ConvertFromByte: func(data []byte, result interface{}) error {
						return cu.ConvertFromByte(data, result)
					},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid filter field",
			args: args{
				ctx: context.Background(),
				req: &pb.RequestQuery{Limit: 10, Offset: 0, Filters: []*pb.RequestQueryFilter{{Field: "invalid", Value: "Place"}}},
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
			_, err := s.PlaceQuery(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GService.PlaceQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
