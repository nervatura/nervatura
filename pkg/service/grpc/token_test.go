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

func TestGService_TokenLogin(t *testing.T) {
	type fields struct {
		Config                 cu.IM
		AppLog                 *slog.Logger
		DecodeAuthToken        func(tokenString string) (cu.IM, error)
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx  context.Context
		req  *pb.RequestEmpty
		ds   *api.DataStore
		auth md.Auth
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
				req: &pb.RequestEmpty{},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					Db: &md.TestDriver{
						Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1, "code": "1234", "auth_object": `{"id": 1, "code": "1234"}`}}, nil
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
				auth: md.Auth{
					UserGroup: md.UserGroupAdmin,
					Id:        1,
					Code:      "1234",
				},
			},
			wantErr: false,
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
			ctx = context.WithValue(ctx, md.AuthUserCtxKey, tt.args.auth)
			_, err := s.TokenLogin(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GService.TokenLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGService_TokenRefresh(t *testing.T) {
	type fields struct {
		Config                 cu.IM
		AppLog                 *slog.Logger
		DecodeAuthToken        func(tokenString string) (cu.IM, error)
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx  context.Context
		req  *pb.RequestEmpty
		ds   *api.DataStore
		auth md.Auth
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
				req: &pb.RequestEmpty{},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					Db: &md.TestDriver{
						Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1, "code": "1234", "auth_object": `{"id": 1, "code": "1234"}`}}, nil
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
					CreateLoginToken: func(params cu.SM, config cu.IM) (result string, err error) {
						return "new_token", nil
					},
					Config: cu.IM{},
				},
				auth: md.Auth{UserGroup: md.UserGroupAdmin},
			},
			wantErr: false,
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
			ctx = context.WithValue(ctx, md.AuthUserCtxKey, tt.args.auth)
			_, err := s.TokenRefresh(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GService.TokenRefresh() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGService_TokenDecode(t *testing.T) {
	type fields struct {
		Config                 cu.IM
		AppLog                 *slog.Logger
		DecodeAuthToken        func(tokenString string) (cu.IM, error)
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx context.Context
		req *pb.RequestTokenDecode
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
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				DecodeAuthToken: func(tokenString string) (cu.IM, error) {
					return cu.IM{"sub": "test_sub", "user_name": "test_user_name", "alias": "test_alias", "exp": "test_exp", "iss": "test_iss"}, nil
				},
			},
			args: args{
				ctx: context.Background(),
				req: &pb.RequestTokenDecode{Token: "test_token"},
			},
			wantErr: false,
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
			_, err := s.TokenDecode(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GService.TokenDecode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
