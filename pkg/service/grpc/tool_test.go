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

func TestGService_ToolUpdate(t *testing.T) {
	type fields struct {
		Config                 cu.IM
		AppLog                 *slog.Logger
		DecodeAuthToken        func(tokenString string) (cu.IM, error)
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx  context.Context
		req  *pb.Tool
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
				req:  &pb.Tool{Id: 1, Code: "1234", ProductCode: "1234", Description: "1234"},
				user: md.Auth{UserGroup: md.UserGroupAdmin},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					Db: &md.TestDriver{
						Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1, "code": "1234", "tool_object": `{"id": 1, "code": "1234"}`}}, nil
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
						return cu.ConvertToByte(v)
					},
					ConvertToType: func(data interface{}, result any) error {
						return ut.ConvertToType(data, result)
					},
					ConvertFromByte: func(data []byte, v any) error {
						return cu.ConvertFromByte(data, v)
					},
				},
			},
			wantErr: false,
		},
		{
			name: "unprocessable entity",
			args: args{
				ctx:  context.Background(),
				req:  &pb.Tool{Id: 1, Code: "1234", ProductCode: "P1234", Description: "1234"},
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
					ConvertToByte: func(v any) ([]byte, error) {
						return cu.ConvertToByte(v)
					},
					ConvertToType: func(data interface{}, result any) error {
						return ut.ConvertToType(data, result)
					},
					ConvertFromByte: func(data []byte, v any) error {
						return cu.ConvertFromByte(data, v)
					},
				},
			},
			wantErr: true,
		},
		{
			name: "error update",
			args: args{
				ctx:  context.Background(),
				req:  &pb.Tool{Id: 0, Code: "1234", ProductCode: "P1234", Description: "1234"},
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
					ConvertToByte: func(v any) ([]byte, error) {
						return cu.ConvertToByte(v)
					},
					ConvertToType: func(data interface{}, result any) error {
						return ut.ConvertToType(data, result)
					},
					ConvertFromByte: func(data []byte, v any) error {
						return cu.ConvertFromByte(data, v)
					},
				},
			},
			wantErr: true,
		},
		{
			name: "error get data by id",
			args: args{
				ctx:  context.Background(),
				req:  &pb.Tool{Id: 0, Code: "1234", ProductCode: "P1234", Description: "1234"},
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
					ConvertToByte: func(v any) ([]byte, error) {
						return cu.ConvertToByte(v)
					},
					ConvertToType: func(data interface{}, result any) error {
						return ut.ConvertToType(data, result)
					},
					ConvertFromByte: func(data []byte, v any) error {
						return cu.ConvertFromByte(data, v)
					},
				},
			},
			wantErr: true,
		},
		{
			name: "product code is required",
			args: args{
				ctx:  context.Background(),
				req:  &pb.Tool{Id: 1, Code: "1234", Description: "1234"},
				user: md.Auth{UserGroup: md.UserGroupAdmin},
			},
			wantErr: true,
		},
		{
			name: "guest user",
			args: args{
				ctx:  context.Background(),
				req:  &pb.Tool{Id: 1, Code: "1234", ProductCode: "P1234", Description: "1234"},
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
			ctx := context.WithValue(tt.args.ctx, md.DataStoreCtxKey, tt.args.ds)
			ctx = context.WithValue(ctx, md.AuthUserCtxKey, tt.args.user)
			_, err := s.ToolUpdate(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GService.ToolUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGService_ToolQuery(t *testing.T) {
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
					{Field: "product_code", Value: "P1234"},
				}},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					Db: &md.TestDriver{
						Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1, "code": "1234", "tool_object": `{"id": 1, "code": "1234"}`}}, nil
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
						return cu.ConvertToByte(v)
					},
					ConvertToType: func(data interface{}, result any) error {
						return ut.ConvertToType(data, result)
					},
					ConvertFromByte: func(data []byte, v any) error {
						return cu.ConvertFromByte(data, v)
					},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid filter field",
			args: args{
				ctx: context.Background(),
				req: &pb.RequestQuery{Limit: 10, Offset: 0, Filters: []*pb.RequestQueryFilter{
					{Field: "invalid", Value: "P1234"},
				}},
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
			ctx := context.WithValue(tt.args.ctx, md.DataStoreCtxKey, tt.args.ds)
			_, err := s.ToolQuery(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GService.ToolQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
