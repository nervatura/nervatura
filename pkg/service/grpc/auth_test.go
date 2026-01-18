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
	pb "github.com/nervatura/nervatura/v6/protos/go"
	td "github.com/nervatura/nervatura/v6/test/driver"
)

func TestGService_AuthLogin(t *testing.T) {
	type fields struct {
		Config                 cu.IM
		AppLog                 *slog.Logger
		DecodeAuthToken        func(tokenString string) (cu.IM, error)
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx context.Context
		req *pb.RequestAuthLogin
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
				Config: cu.IM{"version": "1.0.0"},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args: args{
				ctx: context.Background(),
				req: &pb.RequestAuthLogin{UserName: "test", Password: "test"},
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
			_, err := s.AuthLogin(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GService.AuthLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGService_AuthUpdate(t *testing.T) {
	type fields struct {
		Config                 cu.IM
		AppLog                 *slog.Logger
		DecodeAuthToken        func(tokenString string) (cu.IM, error)
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx  context.Context
		req  *pb.Auth
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
				Config: cu.IM{"version": "1.0.0"},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args: args{
				ctx: context.Background(),
				req: &pb.Auth{Id: 1, Code: "1234", UserName: "test"},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					Db: &td.TestDriver{
						Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1, "code": "1234", "auth_object": `{"id": 1, "code": "1234"}`}}, nil
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
				user: md.Auth{
					UserGroup: md.UserGroupUser,
					Id:        1,
					Code:      "1234",
				},
			},
			wantErr: false,
		},
		{
			name: "unprocessable entity",
			fields: fields{
				Config: cu.IM{"version": "1.0.0"},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args: args{
				ctx: context.Background(),
				req: &pb.Auth{Id: 1, Code: "1234", UserName: "test"},
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
				user: md.Auth{
					UserGroup: md.UserGroupAdmin,
					Id:        1,
					Code:      "1234",
				},
			},
			wantErr: true,
		},
		{
			name: "error update",
			fields: fields{
				Config: cu.IM{"version": "1.0.0"},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args: args{
				ctx: context.Background(),
				req: &pb.Auth{Id: 0, Code: "1234", UserName: "test"},
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
						return cu.ConvertToByte(v)
					},
					ConvertToType: func(data interface{}, result any) error {
						return ut.ConvertToType(data, result)
					},
					ConvertFromByte: func(data []byte, v any) error {
						return cu.ConvertFromByte(data, v)
					},
				},
				user: md.Auth{
					UserGroup: md.UserGroupAdmin,
					Id:        1,
					Code:      "1234",
				},
			},
			wantErr: true,
		},
		{
			name: "user not allowed",
			fields: fields{
				Config: cu.IM{"version": "1.0.0"},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args: args{
				ctx: context.Background(),
				req: &pb.Auth{Id: 2, Code: "2222", UserName: "test"},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					Db: &td.TestDriver{
						Config: cu.IM{
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
				user: md.Auth{
					UserGroup: md.UserGroupUser,
					Id:        1,
					Code:      "1234",
				},
			},
			wantErr: true,
		},
		{
			name: "error get data by id",
			fields: fields{
				Config: cu.IM{"version": "1.0.0"},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args: args{
				ctx: context.Background(),
				req: &pb.Auth{Id: 0, Code: "1234", UserName: "test"},
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
			name: "user name is required",
			fields: fields{
				Config: cu.IM{"version": "1.0.0"},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args: args{
				ctx: context.Background(),
				req: &pb.Auth{Id: 0, Code: "1234"},
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
			_, err := s.AuthUpdate(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GService.AuthUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGService_AuthPassword(t *testing.T) {
	type fields struct {
		Config                 cu.IM
		AppLog                 *slog.Logger
		DecodeAuthToken        func(tokenString string) (cu.IM, error)
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx  context.Context
		req  *pb.RequestPasswordChange
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
				Config: cu.IM{"version": "1.0.0"},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args: args{
				ctx: context.Background(),
				req: &pb.RequestPasswordChange{Password: "test", Confirm: "test"},
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
					ConvertToByte: func(v any) ([]byte, error) {
						return cu.ConvertToByte(v)
					},
					ConvertToType: func(data interface{}, result any) error {
						return ut.ConvertToType(data, result)
					},
					ConvertFromByte: func(data []byte, v any) error {
						return cu.ConvertFromByte(data, v)
					},
					CreatePasswordHash: func(password string) (hash string, err error) {
						return password, nil
					},
				},
				user: md.Auth{
					UserGroup: md.UserGroupAdmin,
					Id:        1,
					Code:      "1234",
				},
			},
			wantErr: false,
		},
		{
			name: "password error",
			fields: fields{
				Config: cu.IM{"version": "1.0.0"},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args: args{
				ctx: context.Background(),
				req: &pb.RequestPasswordChange{Password: "test", Confirm: "test"},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					Db: &td.TestDriver{
						Config: cu.IM{
							"Update": func(data md.Update) (int64, error) {
								return 1, nil
							},
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1, "code": "1234"}}, nil
							},
						},
					},
					CreatePasswordHash: func(password string) (hash string, err error) {
						return "", errors.New("error")
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
			_, err := s.AuthPassword(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GService.AuthPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGService_AuthPasswordReset(t *testing.T) {
	type fields struct {
		Config                 cu.IM
		AppLog                 *slog.Logger
		DecodeAuthToken        func(tokenString string) (cu.IM, error)
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx  context.Context
		req  *pb.RequestGet
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
				Config: cu.IM{"version": "1.0.0"},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args: args{
				ctx: context.Background(),
				req: &pb.RequestGet{Id: 1, Code: "1234"},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					Db: &td.TestDriver{
						Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1, "code": "1234", "auth_object": `{"id": 1, "code": "1234"}`}}, nil
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
					CreatePasswordHash: func(password string) (hash string, err error) {
						return password, nil
					},
				},
				user: md.Auth{
					UserGroup: md.UserGroupAdmin,
					Id:        1,
					Code:      "1234",
				},
			},
			wantErr: false,
		},
		{
			name: "password error",
			fields: fields{
				Config: cu.IM{"version": "1.0.0"},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args: args{
				ctx: context.Background(),
				req: &pb.RequestGet{Id: 1, Code: "1234"},
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
					CreatePasswordHash: func(password string) (hash string, err error) {
						return "", errors.New("error")
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
			name: "user not allowed",
			fields: fields{
				Config: cu.IM{"version": "1.0.0"},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args: args{
				ctx: context.Background(),
				req: &pb.RequestGet{Id: 2, Code: "1234"},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					Db: &td.TestDriver{
						Config: cu.IM{"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 2, "code": "1234", "auth_object": `{"id": 2, "code": "1234"}`}}, nil
						}},
					},
					ReadAll: func(r io.Reader) ([]byte, error) {
						return []byte(`{"id": 2, "code": "1234"}`), nil
					},
					ConvertFromReader: func(r io.Reader, v any) error {
						return cu.ConvertFromByte([]byte(`{"id": 2, "code": "1234"}`), v)
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
				user: md.Auth{
					UserGroup: md.UserGroupUser,
					Id:        1,
					Code:      "2222",
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
			ctx := context.WithValue(tt.args.ctx, md.DataStoreCtxKey, tt.args.ds)
			ctx = context.WithValue(ctx, md.AuthUserCtxKey, tt.args.user)
			_, err := s.AuthPasswordReset(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GService.AuthPasswordReset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
