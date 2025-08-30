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

func TestGService_LogGet(t *testing.T) {
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
				req: &pb.RequestGet{Id: 1, Code: "1234"},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					Db: &md.TestDriver{
						Config: cu.IM{"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "1234", "log_object": `{"id": 1, "code": "1234"}`}}, nil
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
		},
		{
			name: "unprocessable entity",
			args: args{
				ctx: context.Background(),
				req: &pb.RequestGet{Id: 1, Code: "1234"},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					Db: &md.TestDriver{
						Config: cu.IM{"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "1234"}}, nil
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
			_, err := s.LogGet(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GService.LogGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGService_LogQuery(t *testing.T) {
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
				req: &pb.RequestQuery{Limit: 10, Offset: 0, Filters: []*pb.RequestQueryFilter{{Field: "log_type", Value: "1"}}},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					Db: &md.TestDriver{
						Config: cu.IM{"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "1234", "log_object": `{"id": 1, "code": "1234"}`}}, nil
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
		},
		{
			name: "invalid filter field",
			args: args{
				ctx: context.Background(),
				req: &pb.RequestQuery{Limit: 10, Offset: 0, Filters: []*pb.RequestQueryFilter{{Field: "invalid", Value: "1"}}},
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
			_, err := s.LogQuery(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GService.LogQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
