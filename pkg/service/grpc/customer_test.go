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

func TestGService_CustomerUpdate(t *testing.T) {
	type fields struct {
		Config                 cu.IM
		AppLog                 *slog.Logger
		DecodeAuthToken        func(tokenString string) (cu.IM, error)
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx  context.Context
		req  *pb.Customer
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
				ctx: context.Background(),
				req: &pb.Customer{Id: 0, Code: "123456", CustomerName: "test"},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					Db: &td.TestDriver{
						Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1, "customer_object": `{"id": 1, "code": "123456"}`}}, nil
							},
							"Update": func(data md.Update) (int64, error) {
								return 1, nil
							},
						},
					},
					ReadAll: func(r io.Reader) ([]byte, error) {
						return []byte(`{""code": "123456"}`), nil
					},
					ConvertFromReader: func(r io.Reader, v any) error {
						return cu.ConvertFromByte([]byte(`{""code": "123456"}`), v)
					},
					ConvertToByte: func(v any) ([]byte, error) {
						return []byte(`{""code": "123456"}`), nil
					},
					ConvertToType: func(v any, t any) error {
						return ut.ConvertToType(v, t)
					},
					ConvertFromByte: func(data []byte, result interface{}) error {
						return cu.ConvertFromByte(data, result)
					},
				},
				user: md.Auth{UserGroup: md.UserGroupAdmin},
			},
			wantErr: false,
		},
		{
			name: "customer name is required",
			args: args{
				ctx: context.Background(),
				req: &pb.Customer{Id: 0, Code: ""},
				ds:  &api.DataStore{},
			},
			wantErr: true,
		},
		{
			name: "unprocessable get entity",
			args: args{
				ctx: context.Background(),
				req: &pb.Customer{Id: 0, Code: "1234", CustomerName: "test"},
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
						return nil, errors.New("error")
					},
					ConvertFromReader: func(r io.Reader, v any) error {
						return cu.ConvertFromByte([]byte(`{""code": "123456"}`), v)
					},
					ConvertToByte: func(v any) ([]byte, error) {
						return nil, errors.New("error")
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
				ctx: context.Background(),
				req: &pb.Customer{Id: 0, Code: "1234", CustomerName: "test"},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					Db: &td.TestDriver{
						Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{}, nil
							},
							"Update": func(data md.Update) (int64, error) {
								return 0, errors.New("error")
							},
						},
					},
					ReadAll: func(r io.Reader) ([]byte, error) {
						return []byte(`{""code": "123456"}`), nil
					},
					ConvertFromReader: func(r io.Reader, v any) error {
						return cu.ConvertFromByte([]byte(`{""code": "123456"}`), v)
					},
					ConvertToByte: func(v any) ([]byte, error) {
						return nil, errors.New("error")
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
			name: "error get entity",
			args: args{
				ctx: context.Background(),
				req: &pb.Customer{Id: 0, Code: "1234", CustomerName: "test"},
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
						return []byte(`{""code": "123456"}`), nil
					},
					ConvertFromReader: func(r io.Reader, v any) error {
						return cu.ConvertFromByte([]byte(`{""code": "123456"}`), v)
					},
					ConvertToByte: func(v any) ([]byte, error) {
						return nil, errors.New("error")
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
			name: "guest user",
			args: args{
				ctx:  context.Background(),
				req:  &pb.Customer{Id: 0, Code: "1234", CustomerName: "test"},
				ds:   &api.DataStore{},
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
			_, err := s.CustomerUpdate(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GService.CustomerUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGService_CustomerQuery(t *testing.T) {
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
					{Field: "customer_type", Value: "1"},
					{Field: "customer_name", Value: "test"},
					{Field: "tag", Value: "tag1"},
				}},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					Db: &td.TestDriver{
						Config: cu.IM{"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "customer_object": `{"id": 1, "code": "123456"}`}}, nil
						}},
					},
					ReadAll: func(r io.Reader) ([]byte, error) {
						return []byte(`{""code": "123456"}`), nil
					},
					ConvertFromReader: func(r io.Reader, v any) error {
						return cu.ConvertFromByte([]byte(`{""code": "123456"}`), v)
					},
					ConvertToByte: func(v any) ([]byte, error) {
						return nil, errors.New("error")
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
				req: &pb.RequestQuery{Limit: 10, Offset: 0, Filters: []*pb.RequestQueryFilter{{Field: "invalid", Value: "test"}}},
				ds:  &api.DataStore{},
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
			_, err := s.CustomerQuery(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GService.CustomerQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
