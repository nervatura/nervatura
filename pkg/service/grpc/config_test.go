package grpc

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"testing"

	cu "github.com/nervatura/component/pkg/util"
	api "github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	pb "github.com/nervatura/nervatura/v6/pkg/service/grpc/proto"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

func TestGService_ConfigUpdate(t *testing.T) {
	type fields struct {
		Config                 cu.IM
		AppLog                 *slog.Logger
		DecodeAuthToken        func(tokenString string) (cu.IM, error)
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx  context.Context
		req  *pb.Config
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
			fields: fields{
				Config:                 cu.IM{},
				AppLog:                 slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				DecodeAuthToken:        nil,
				UnimplementedAPIServer: pb.UnimplementedAPIServer{},
			},
			args: args{
				ctx: context.Background(),
				req: &pb.Config{
					Id:         0,
					Code:       "123456",
					ConfigType: pb.ConfigType_CONFIG_PRINT_QUEUE,
					Data:       &pb.Config_Map{Map: &pb.ConfigMap{}},
				},
				ds: &api.DataStore{
					Config: cu.IM{},
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
						return []byte(`{"code": "123456", "config_type": "CONFIG_PRINT_QUEUE", "data": {"key": "value"}}`), nil
					},
					ConvertFromReader: func(r io.Reader, v any) error {
						return cu.ConvertFromByte([]byte(`{"code": "123456", "config_type": "CONFIG_PRINT_QUEUE", "data": {"key": "value"}}`), v)
					},
					ConvertToByte: func(v any) ([]byte, error) {
						return []byte(`{"code": "123456", "config_type": "CONFIG_PRINT_QUEUE"}`), nil
					},
					ConvertFromByte: func(b []byte, v any) error {
						return cu.ConvertFromByte(b, v)
					},
					ConvertToType: func(v any, t any) error {
						return ut.ConvertToType(v, t)
					},
				},
				user: md.Auth{
					UserGroup: md.UserGroupAdmin,
				},
			},
			wantErr: false,
		},
		{
			name: "data is required",
			args: args{
				req: &pb.Config{
					Id:         0,
					Code:       "123456",
					ConfigType: pb.ConfigType_CONFIG_PRINT_QUEUE,
				},
				ds: &api.DataStore{
					Config: cu.IM{},
					Db: &md.TestDriver{
						Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 0}}, nil
							},
						},
					},
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					ReadAll: func(r io.Reader) ([]byte, error) {
						return []byte(`{"code": "123456", "config_type": "CONFIG_PRINT_QUEUE", "data": {"key": "value"}}`), nil
					},
					ConvertFromReader: func(r io.Reader, v any) error {
						return cu.ConvertFromByte([]byte(`{"code": "123456", "config_type": "CONFIG_PRINT_QUEUE", "data": {"key": "value"}}`), v)
					},
					ConvertToByte: func(v any) ([]byte, error) {
						return []byte(`{"code": "123456", "config_type": "CONFIG_PRINT_QUEUE"}`), nil
					},
					ConvertFromByte: func(b []byte, v any) error {
						return cu.ConvertFromByte(b, v)
					},
					ConvertToType: func(v any, t any) error {
						return ut.ConvertToType(v, t)
					},
				},
				user: md.Auth{
					UserGroup: md.UserGroupAdmin,
				},
			},
			wantErr: true,
		},
		{
			name: "not found",
			args: args{
				req: &pb.Config{
					Id:         1,
					Code:       "123456",
					ConfigType: pb.ConfigType_CONFIG_PRINT_QUEUE,
				},
				ds: &api.DataStore{
					Config: cu.IM{},
					Db: &md.TestDriver{
						Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{}, nil
							},
						},
					},
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					ReadAll: func(r io.Reader) ([]byte, error) {
						return []byte(`{"code": "123456", "config_type": "CONFIG_PRINT_QUEUE", "data": {"key": "value"}}`), nil
					},
					ConvertFromReader: func(r io.Reader, v any) error {
						return cu.ConvertFromByte([]byte(`{"code": "123456", "config_type": "CONFIG_PRINT_QUEUE", "data": {"key": "value"}}`), v)
					},
					ConvertToByte: func(v any) ([]byte, error) {
						return []byte(`{"code": "123456", "config_type": "CONFIG_PRINT_QUEUE"}`), nil
					},
					ConvertFromByte: func(b []byte, v any) error {
						return cu.ConvertFromByte(b, v)
					},
					ConvertToType: func(v any, t any) error {
						return ut.ConvertToType(v, t)
					},
				},
			},
			wantErr: true,
		},
		{
			name: "not allowed guest group",
			args: args{
				req: &pb.Config{
					Id:         0,
					Code:       "123456",
					ConfigType: pb.ConfigType_CONFIG_PRINT_QUEUE,
				},
				user: md.Auth{
					UserGroup: md.UserGroupGuest,
				},
			},
			wantErr: true,
		},
		{
			name: "not allowed user group",
			args: args{
				req: &pb.Config{
					Id:         0,
					Code:       "123456",
					ConfigType: pb.ConfigType_CONFIG_MAP,
				},
				user: md.Auth{
					UserGroup: md.UserGroupUser,
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
			_, err := s.ConfigUpdate(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GService.ConfigUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGService_ConfigGet(t *testing.T) {
	type fields struct {
		Config                 cu.IM
		AppLog                 *slog.Logger
		DecodeAuthToken        func(tokenString string) (cu.IM, error)
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx context.Context
		req *pb.RequestGet
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
				req: &pb.RequestGet{Id: 1},
				ds: &api.DataStore{
					Config: cu.IM{},
					Db: &md.TestDriver{
						Config: cu.IM{"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1}}, nil
						}},
					},
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					ReadAll: func(r io.Reader) ([]byte, error) {
						return []byte(`{"id": 1, "code": "123456", "config_type": "CONFIG_PRINT_QUEUE", "data": {"key": "value"}}`), nil
					},
					ConvertFromReader: func(r io.Reader, v any) error {
						return cu.ConvertFromByte([]byte(`{"id": 1, "code": "123456", "config_type": "CONFIG_PRINT_QUEUE", "data": {"key": "value"}}`), v)
					},
					ConvertToByte: func(v any) ([]byte, error) {
						return []byte(`{"id": 1, "code": "123456", "config_type": "CONFIG_PRINT_QUEUE"}`), nil
					},
					ConvertFromByte: func(b []byte, v any) error {
						return cu.ConvertFromByte(b, v)
					},
					ConvertToType: func(v any, t any) error {
						return ut.ConvertToType(v, t)
					},
				},
			},
			wantErr: false,
		},
		{
			name: "not found",
			args: args{
				ctx: context.Background(),
				req: &pb.RequestGet{Code: "123456"},
				ds: &api.DataStore{
					Config: cu.IM{},
					Db: &md.TestDriver{
						Config: cu.IM{"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{}, nil
						}},
					},
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					ReadAll: func(r io.Reader) ([]byte, error) {
						return []byte(`{"id": 1, "code": "123456", "config_type": "CONFIG_PRINT_QUEUE", "data": {"key": "value"}}`), nil
					},
					ConvertFromReader: func(r io.Reader, v any) error {
						return cu.ConvertFromByte([]byte(`{"id": 1, "code": "123456", "config_type": "CONFIG_PRINT_QUEUE", "data": {"key": "value"}}`), v)
					},
					ConvertToByte: func(v any) ([]byte, error) {
						return []byte(`{"id": 1, "code": "123456", "config_type": "CONFIG_PRINT_QUEUE"}`), nil
					},
					ConvertFromByte: func(b []byte, v any) error {
						return cu.ConvertFromByte(b, v)
					},
					ConvertToType: func(v any, t any) error {
						return ut.ConvertToType(v, t)
					},
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
			_, err := s.ConfigGet(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GService.ConfigGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGService_ConfigQuery(t *testing.T) {
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
				req: &pb.RequestQuery{Limit: 10, Offset: 0, Filters: []*pb.RequestQueryFilter{{Field: "config_type", Value: "CONFIG_PRINT_QUEUE"}}},
				ds: &api.DataStore{
					Config: cu.IM{},
					Db: &md.TestDriver{
						Config: cu.IM{"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1}}, nil
						}},
					},
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					ReadAll: func(r io.Reader) ([]byte, error) {
						return []byte(`[{"id": 1, "code": "123456", "config_type": "CONFIG_PRINT_QUEUE", "data": {"key": "value"}}]`), nil
					},
					ConvertFromReader: func(r io.Reader, v any) error {
						return cu.ConvertFromByte([]byte(`[{"id": 1, "code": "123456", "config_type": "CONFIG_PRINT_QUEUE", "data": {"key": "value"}}]`), v)
					},
					ConvertToByte: func(v any) ([]byte, error) {
						return []byte(`[{"id": 1, "code": "123456", "config_type": "CONFIG_PRINT_QUEUE"}]`), nil
					},
					ConvertFromByte: func(b []byte, v any) error {
						return cu.ConvertFromByte(b, v)
					},
					ConvertToType: func(v any, t any) error {
						return ut.ConvertToType(v, t)
					},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid filter field",
			args: args{
				ctx: context.Background(),
				req: &pb.RequestQuery{Limit: 10, Offset: 0, Filters: []*pb.RequestQueryFilter{{Field: "invalid", Value: "CONFIG_PRINT_QUEUE"}}},
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
			_, err := s.ConfigQuery(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GService.ConfigQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
