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

func TestGService_TransUpdate(t *testing.T) {
	type fields struct {
		Config                 cu.IM
		AppLog                 *slog.Logger
		DecodeAuthToken        func(tokenString string) (cu.IM, error)
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx  context.Context
		req  *pb.Trans
		ds   *api.DataStore
		user md.Auth
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantPbTrans *pb.Trans
		wantErr     bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: &pb.Trans{Id: 1, Code: "1234", TransType: pb.TransType_TRANS_INVOICE, Direction: pb.Direction_DIRECTION_OUT, TransDate: "2024-01-01",
					TransCode: "T01234", CustomerCode: "C01234", EmployeeCode: "E01133", ProjectCode: "P01123", PlaceCode: "PL011222",
					CurrencyCode: "USD", AuthCode: "A01234"},
				ds: &api.DataStore{
					Db: &td.TestDriver{
						Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1, "code": "1234", "trans_object": `{"id": 1, "code": "1234"}`}}, nil
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
				user: md.Auth{UserGroup: md.UserGroupAdmin},
			},
			wantErr: false,
		},
		{
			name: "unprocessable entity",
			args: args{
				ctx: context.Background(),
				req: &pb.Trans{Id: 1, Code: "1234", TransType: pb.TransType_TRANS_INVOICE, Direction: pb.Direction_DIRECTION_OUT, TransDate: "2024-01-01",
					TransCode: "T01234", CustomerCode: "C01234", EmployeeCode: "E01133", ProjectCode: "P01123", PlaceCode: "PL011222",
					CurrencyCode: "USD", AuthCode: "A01234"},
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
						return cu.ConvertToByte(v)
					},
					ConvertToType: func(data interface{}, result any) error {
						return ut.ConvertToType(data, result)
					},
					ConvertFromByte: func(data []byte, v any) error {
						return cu.ConvertFromByte(data, v)
					},
				},
				user: md.Auth{UserGroup: md.UserGroupAdmin},
			},
			wantErr: true,
		},
		{
			name: "error update",
			args: args{
				ctx: context.Background(),
				req: &pb.Trans{Id: 0, Code: "1234", TransType: pb.TransType_TRANS_INVOICE, Direction: pb.Direction_DIRECTION_OUT, TransDate: "2024-01-01",
					TransCode: "T01234", CustomerCode: "C01234", EmployeeCode: "E01133", ProjectCode: "P01123", PlaceCode: "PL011222",
					CurrencyCode: "USD", AuthCode: "A01234"},
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
				ctx: context.Background(),
				req: &pb.Trans{Id: 0, Code: "1234", TransType: pb.TransType_TRANS_INVOICE, Direction: pb.Direction_DIRECTION_OUT, TransDate: "2024-01-01",
					TransCode: "T01234", CustomerCode: "C01234", EmployeeCode: "E01133", ProjectCode: "P01123", PlaceCode: "PL011222",
					CurrencyCode: "USD", AuthCode: "A01234"},
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
			name: "customer code and currency code is required",
			args: args{
				ctx: context.Background(),
				req: &pb.Trans{Id: 0, Code: "1234", TransType: pb.TransType_TRANS_INVOICE, Direction: pb.Direction_DIRECTION_OUT, TransDate: "2024-01-01",
					TransCode: "T01234", EmployeeCode: "E01133", ProjectCode: "P01123", PlaceCode: "PL011222",
					AuthCode: "A01234"},
				user: md.Auth{UserGroup: md.UserGroupAdmin},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				},
			},
			wantErr: true,
		},
		{
			name: "trans date is required",
			args: args{
				ctx: context.Background(),
				req: &pb.Trans{Id: 0, Code: "1234", TransType: pb.TransType_TRANS_INVOICE, Direction: pb.Direction_DIRECTION_OUT,
					TransCode: "T01234", CustomerCode: "C01234", EmployeeCode: "E01133", ProjectCode: "P01123", PlaceCode: "PL011222",
					CurrencyCode: "USD", AuthCode: "A01234"},
				user: md.Auth{UserGroup: md.UserGroupAdmin},
			},
			wantErr: true,
		},
		{
			name: "guest user",
			args: args{
				ctx: context.Background(),
				req: &pb.Trans{Id: 0, Code: "1234", TransType: pb.TransType_TRANS_INVOICE, Direction: pb.Direction_DIRECTION_OUT, TransDate: "2024-01-01",
					TransCode: "T01234", CustomerCode: "C01234", EmployeeCode: "E01133", ProjectCode: "P01123", PlaceCode: "PL011222",
					CurrencyCode: "USD", AuthCode: "A01234"},
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
			_, err := s.TransUpdate(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GService.TransUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGService_TransQuery(t *testing.T) {
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
					{Field: "trans_type", Value: "invoice"},
					{Field: "direction", Value: "out"},
					{Field: "trans_date", Value: "2024-01-01"},
					{Field: "tag", Value: "tag1"},
				}},
				ds: &api.DataStore{
					AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					Db: &td.TestDriver{
						Config: cu.IM{
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1, "code": "1234", "trans_object": `{"id": 1, "code": "1234"}`}}, nil
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
					{Field: "invalid", Value: "invalid"},
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
			_, err := s.TransQuery(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GService.TransQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
