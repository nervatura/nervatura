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
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
	pb "github.com/nervatura/nervatura/v6/proto"
	td "github.com/nervatura/nervatura/v6/test/driver"
)

func TestGService_TaxUpdate(t *testing.T) {
	type fields struct {
		Config                 cu.IM
		AppLog                 *slog.Logger
		DecodeAuthToken        func(tokenString string) (cu.IM, error)
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx  context.Context
		req  *pb.Tax
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
				req:  &pb.Tax{Id: 1, Code: "1234"},
				user: md.Auth{UserGroup: md.UserGroupAdmin},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					Db: &td.TestDriver{
						Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1, "code": "1234", "tax_object": `{"id": 1, "code": "1234"}`}}, nil
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
					ConvertFromByte: func(data []byte, result interface{}) error {
						return cu.ConvertFromByte(data, result)
					},
					ConvertToByte: func(v any) ([]byte, error) {
						return cu.ConvertToByte(v)
					},
					ConvertToType: func(v any, result any) error {
						return ut.ConvertToType(v, result)
					},
				},
			},
			wantErr: false,
		},
		{
			name: "unprocessable entity",
			args: args{
				ctx:  context.Background(),
				req:  &pb.Tax{Id: 0, Code: "TEST"},
				user: md.Auth{UserGroup: md.UserGroupAdmin},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					Db: &td.TestDriver{
						Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1, "code": "1234"}}, nil
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
					ConvertFromByte: func(data []byte, result interface{}) error {
						return cu.ConvertFromByte(data, result)
					},
					ConvertToByte: func(v any) ([]byte, error) {
						return cu.ConvertToByte(v)
					},
					ConvertToType: func(v any, result any) error {
						return ut.ConvertToType(v, result)
					},
				},
			},
			wantErr: true,
		},
		{
			name: "error get data by id",
			args: args{
				ctx:  context.Background(),
				req:  &pb.Tax{Id: 0, Code: "1234"},
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
					ConvertFromByte: func(data []byte, result interface{}) error {
						return cu.ConvertFromByte(data, result)
					},
					ConvertToByte: func(v any) ([]byte, error) {
						return cu.ConvertToByte(v)
					},
					ConvertToType: func(v any, result any) error {
						return ut.ConvertToType(v, result)
					},
				},
			},
			wantErr: true,
		},
		{
			name: "tax code is required",
			args: args{
				ctx:  context.Background(),
				req:  &pb.Tax{Id: 1, Code: ""},
				user: md.Auth{UserGroup: md.UserGroupAdmin},
			},
			wantErr: true,
		},
		{
			name: "guest user",
			args: args{
				ctx:  context.Background(),
				req:  &pb.Tax{Id: 1, Code: "1234"},
				user: md.Auth{UserGroup: md.UserGroupGuest},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				},
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
			_, err := s.TaxUpdate(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GService.TaxUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGService_TaxQuery(t *testing.T) {
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
				req: &pb.RequestQuery{Limit: 10, Offset: 0, Filters: []*pb.RequestQueryFilter{{Field: "tag", Value: "TEST"}}},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					Db: &td.TestDriver{
						Config: cu.IM{"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "1234", "tax_object": `{"id": 1, "code": "1234"}`}}, nil
						}},
					},
					ReadAll: func(r io.Reader) ([]byte, error) {
						return []byte(`{"id": 1, "code": "1234"}`), nil
					},
					ConvertFromReader: func(r io.Reader, v any) error {
						return cu.ConvertFromByte([]byte(`{"id": 1, "code": "1234"}`), v)
					},
					ConvertFromByte: func(data []byte, result interface{}) error {
						return cu.ConvertFromByte(data, result)
					},
					ConvertToByte: func(v any) ([]byte, error) {
						return cu.ConvertToByte(v)
					},
					ConvertToType: func(v any, result any) error {
						return ut.ConvertToType(v, result)
					},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid filter field",
			args: args{
				ctx: context.Background(),
				req: &pb.RequestQuery{Limit: 10, Offset: 0, Filters: []*pb.RequestQueryFilter{{Field: "invalid", Value: "TEST"}}},
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
			_, err := s.TaxQuery(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GService.TaxQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
