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

func TestGService_TokenAuth(t *testing.T) {
	type fields struct {
		Config                 cu.IM
		AppLog                 *slog.Logger
		DecodeAuthToken        func(tokenString string) (cu.IM, error)
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		authorization []string
		parent        context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test TokenAuth",
			fields: fields{
				Config: cu.IM{
					"tokenKeys": []cu.SM{
						{"type": "private", "value": "SECRET_KEY"},
					},
					"db": &td.TestDriver{
						Config: cu.IM{
							"Connection": func() struct {
								Alias     string
								Connected bool
								Engine    string
							} {
								return struct {
									Alias     string
									Connected bool
									Engine    string
								}{
									Alias:     "test",
									Connected: true,
									Engine:    "sqlite",
								}
							},
							"QuerySQL": func(sqlString string) ([]cu.IM, error) {
								return []cu.IM{{"id": 1, "name": "test"}}, nil
							},
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1}}, nil
							},
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				DecodeAuthToken: func(tokenString string) (cu.IM, error) {
					return cu.IM{
						"alias": "test",
					}, nil
				},
				UnimplementedAPIServer: pb.UnimplementedAPIServer{},
			},
			args:    args{authorization: []string{"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJJRDAxMjMiLCJlbWFpbCI6InVzZXJAbWFpbC5jb20iLCJpc3MiOiJlb25leCJ9.1TKb3k2xYslwhPDQD50AuSsFqRlIxUB7kErrQqaHVIg"}, parent: context.Background()},
			wantErr: false,
		},
		{
			name: "TokenLogin error",
			fields: fields{
				Config: cu.IM{
					"tokenKeys": []cu.SM{
						{"type": "private", "value": "SECRET_KEY"},
					},
					"db": &td.TestDriver{
						Config: cu.IM{},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				DecodeAuthToken: func(tokenString string) (cu.IM, error) {
					return cu.IM{}, nil
				},
				UnimplementedAPIServer: pb.UnimplementedAPIServer{},
			},
			args:    args{authorization: []string{"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"}, parent: context.Background()},
			wantErr: true,
		},
		{
			name: "invalid token",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				DecodeAuthToken: func(tokenString string) (cu.IM, error) {
					return cu.IM{}, errors.New("invalid token")
				},
				UnimplementedAPIServer: pb.UnimplementedAPIServer{},
			},
			args:    args{authorization: []string{"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"}, parent: context.Background()},
			wantErr: true,
		},
		{
			name: "no token 1",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				DecodeAuthToken: func(tokenString string) (cu.IM, error) {
					return cu.IM{}, nil
				},
				UnimplementedAPIServer: pb.UnimplementedAPIServer{},
			},
			args:    args{authorization: []string{"Bearer "}, parent: context.Background()},
			wantErr: true,
		},
		{
			name: "no token 2",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				DecodeAuthToken: func(tokenString string) (cu.IM, error) {
					return cu.IM{}, nil
				},
				UnimplementedAPIServer: pb.UnimplementedAPIServer{},
			},
			args:    args{authorization: []string{}, parent: context.Background()},
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
			_, err := s.TokenAuth(tt.args.authorization, tt.args.parent)
			if (err != nil) != tt.wantErr {
				t.Errorf("GService.TokenAuth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGService_ApiKeyAuth(t *testing.T) {
	type fields struct {
		Config                 cu.IM
		AppLog                 *slog.Logger
		DecodeAuthToken        func(tokenString string) (cu.IM, error)
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		authorization []string
		parent        context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test ApiKeyAuth",
			fields: fields{
				Config: cu.IM{"NT_API_KEY": "test"},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args:    args{authorization: []string{"test"}, parent: context.Background()},
			wantErr: false,
		},
		{
			name: "invalid api key",
			fields: fields{
				Config: cu.IM{"NT_API_KEY": "test"},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args:    args{authorization: []string{"invalid"}, parent: context.Background()},
			wantErr: true,
		},
		{
			name: "no api key",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args:    args{authorization: []string{}, parent: context.Background()},
			wantErr: true,
		},
		{
			name: "empty api key",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args:    args{authorization: []string{" "}, parent: context.Background()},
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
			_, err := s.ApiKeyAuth(tt.args.authorization, tt.args.parent)
			if (err != nil) != tt.wantErr {
				t.Errorf("GService.ApiKeyAuth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGService_Delete(t *testing.T) {
	type fields struct {
		Config                 cu.IM
		AppLog                 *slog.Logger
		DecodeAuthToken        func(tokenString string) (cu.IM, error)
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx  context.Context
		req  *pb.RequestDelete
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
			name: "default",
			args: args{
				ctx:  context.Background(),
				req:  &pb.RequestDelete{Model: pb.Model_CUSTOMER, Id: 1, Code: "CUS0000000000N1"},
				user: md.Auth{UserGroup: md.UserGroupAdmin},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					Db: &td.TestDriver{
						Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1, "code": "CUS0000000000N1"}}, nil
							},
							"Update": func(data md.Update) (int64, error) {
								return 1, errors.New("error")
							},
						},
					},
					ReadAll: func(r io.Reader) ([]byte, error) {
						return []byte(`{"id": 1, "code": "CUS0000000000N1"}`), nil
					},
					ConvertFromReader: func(r io.Reader, v any) error {
						return cu.ConvertFromByte([]byte(`{"id": 1, "code": "CUS0000000000N1"}`), v)
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
			name: "config ok",
			args: args{
				ctx:  context.Background(),
				req:  &pb.RequestDelete{Model: pb.Model_CONFIG, Id: 1, Code: "CONFIG_PRINT_QUEUE"},
				user: md.Auth{UserGroup: md.UserGroupAdmin},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					Db: &td.TestDriver{
						Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1, "code": "CONFIG_PRINT_QUEUE"}}, nil
							},
							"Update": func(data md.Update) (int64, error) {
								return 1, nil
							},
						},
					},
					ReadAll: func(r io.Reader) ([]byte, error) {
						return []byte(`{"id": 1, "code": "CONFIG_PRINT_QUEUE"}`), nil
					},
					ConvertFromReader: func(r io.Reader, v any) error {
						return cu.ConvertFromByte([]byte(`{"id": 1, "code": "CONFIG_PRINT_QUEUE"}`), v)
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
			name: "config admin error",
			args: args{
				ctx:  context.Background(),
				req:  &pb.RequestDelete{Model: pb.Model_CONFIG, Id: 1, Code: "CONFIG_MAP"},
				user: md.Auth{UserGroup: md.UserGroupUser},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					Db: &td.TestDriver{
						Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1, "code": "CONFIG_MAP"}}, nil
							},
							"Update": func(data md.Update) (int64, error) {
								return 1, nil
							},
						},
					},
					ReadAll: func(r io.Reader) ([]byte, error) {
						return []byte(`{"id": 1, "code": "CONFIG_MAP"}`), nil
					},
					ConvertFromReader: func(r io.Reader, v any) error {
						return cu.ConvertFromByte([]byte(`{"id": 1, "code": "CONFIG_MAP"}`), v)
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
			name: "config get error",
			args: args{
				ctx:  context.Background(),
				req:  &pb.RequestDelete{Model: pb.Model_CONFIG, Id: 1, Code: "CONFIG_PRINT_QUEUE"},
				user: md.Auth{UserGroup: md.UserGroupUser},
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
						return []byte(`{"id": 1, "code": "CONFIG_MAP"}`), nil
					},
					ConvertFromReader: func(r io.Reader, v any) error {
						return cu.ConvertFromByte([]byte(`{"id": 1, "code": "CONFIG_MAP"}`), v)
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
			name: "auth ok",
			args: args{
				ctx:  context.Background(),
				req:  &pb.RequestDelete{Model: pb.Model_AUTH, Id: 2, Code: "USR0000000000N2"},
				user: md.Auth{UserGroup: md.UserGroupAdmin},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					Db: &td.TestDriver{
						Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 2, "code": "USR0000000000N2", "auth_object": `{"id": 2, "code": "USR0000000000N2"}`}}, nil
							},
							"Update": func(data md.Update) (int64, error) {
								return 1, nil
							},
						},
					},
					ReadAll: func(r io.Reader) ([]byte, error) {
						return []byte(`{"id": 2, "code": "USR0000000000N2", "auth_object": {"id": 2, "code": "USR0000000000N2"}}`), nil
					},
					ConvertFromReader: func(r io.Reader, v any) error {
						return cu.ConvertFromByte([]byte(`{"id": 2, "code": "USR0000000000N2", "auth_object": {"id": 2, "code": "USR0000000000N2"}}`), v)
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
			name: "auth get error",
			args: args{
				ctx:  context.Background(),
				req:  &pb.RequestDelete{Model: pb.Model_AUTH, Id: 2, Code: "USR0000000000N2"},
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
						return []byte(`{"id": 2, "code": "USR0000000000N2", "auth_object": {"id": 2, "code": "USR0000000000N2"}}`), nil
					},
					ConvertFromReader: func(r io.Reader, v any) error {
						return cu.ConvertFromByte([]byte(`{"id": 2, "code": "USR0000000000N2", "auth_object": {"id": 2, "code": "USR0000000000N2"}}`), v)
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
			name: "auth protected admin",
			args: args{
				ctx:  context.Background(),
				req:  &pb.RequestDelete{Model: pb.Model_AUTH, Id: 1, Code: "USR0000000000N1"},
				user: md.Auth{UserGroup: md.UserGroupAdmin},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					Db: &td.TestDriver{
						Config: cu.IM{},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "guest user",
			args: args{
				ctx:  context.Background(),
				req:  &pb.RequestDelete{Model: pb.Model_AUTH, Id: 2, Code: "USR0000000000N2"},
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
			_, err := s.Delete(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GService.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGService_Function(t *testing.T) {
	type fields struct {
		Config                 cu.IM
		AppLog                 *slog.Logger
		DecodeAuthToken        func(tokenString string) (cu.IM, error)
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx context.Context
		req *pb.RequestFunction
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
				req: &pb.RequestFunction{Function: "test", Args: cu.SM{"key": "value"}},
				ds: &api.DataStore{
					Db: &td.TestDriver{
						Config: cu.IM{},
					},
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
			_, err := s.Function(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GService.Function() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGService_Database(t *testing.T) {
	type fields struct {
		Config                 cu.IM
		AppLog                 *slog.Logger
		DecodeAuthToken        func(tokenString string) (cu.IM, error)
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx context.Context
		req *pb.RequestDatabase
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
				req: &pb.RequestDatabase{Alias: "test", Demo: false},
				ds: &api.DataStore{
					Db: &td.TestDriver{
						Config: cu.IM{},
					},
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					ReadAll: func(r io.Reader) ([]byte, error) {
						return []byte(`{"alias": "test", "demo": false}`), nil
					},
					ConvertFromReader: func(r io.Reader, v any) error {
						return cu.ConvertFromByte([]byte(`{"alias": "test", "demo": false}`), v)
					},
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
			_, err := s.Database(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GService.Database() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGService_View(t *testing.T) {
	type fields struct {
		Config                 cu.IM
		AppLog                 *slog.Logger
		DecodeAuthToken        func(tokenString string) (cu.IM, error)
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx context.Context
		req *pb.RequestView
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
				req: &pb.RequestView{Name: pb.ViewName_VIEW_CURRENCY_VIEW, Filter: "", OrderBy: []string{"id"}, Limit: 10, Offset: 0},
				ds: &api.DataStore{
					Db: &td.TestDriver{
						Config: cu.IM{},
					},
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					ReadAll: func(r io.Reader) ([]byte, error) {
						return []byte(`{"id": 1, "code": "USD"}`), nil
					},
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
			_, err := s.View(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GService.View() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
