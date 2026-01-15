package cli

import (
	"bytes"
	"errors"
	"io"
	"log/slog"
	"testing"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	td "github.com/nervatura/nervatura/v6/test/driver"
)

func TestCLIService_AuthInsert(t *testing.T) {
	type fields struct {
		Config map[string]interface{}
	}
	type args struct {
		options     cu.IM
		requestData string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "unprocessable entity",
			fields: fields{},
			args: args{
				options: cu.IM{
					"ds": &api.DataStore{
						Config: cu.IM{},
						Db:     &td.TestDriver{},
						AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
						ConvertFromByte: func(data []byte, v any) error {
							return errors.New("test error")
						},
					},
				},
				requestData: `{"user_name": "test", "user_group": "user"}`,
			},
			want: `{"code":422,"data":"test error"}`,
		},
		{
			name:   "missing user name",
			fields: fields{},
			args: args{
				options: cu.IM{
					"ds": &api.DataStore{
						Config: cu.IM{},
						Db:     &td.TestDriver{},
						AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte([]byte(`{"user_group": "GROUP_USER"}`), v)
						},
					},
				},
				requestData: `{"user_group": "GROUP_USER"}`,
			},
			want: `{"code":422,"data":"auth user_name and user_group are required"}`,
		},
		{
			name:   "success",
			fields: fields{},
			args: args{
				options: cu.IM{
					"ds": &api.DataStore{
						Config: cu.IM{},
						Db: &td.TestDriver{
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
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte([]byte(`{"user_name": "test", "user_group": "GROUP_USER", "code": "123456"}`), v)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(cu.IM{"id": 1})
						},
					},
				},
				requestData: `{"user_name": "test", "user_group": "GROUP_USER"}`,
			},
			want: `{"code":201,"data":{"id":1}}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := &CLIService{
				Config: tt.fields.Config,
			}
			if got := cli.AuthInsert(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.AuthInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_AuthUpdate(t *testing.T) {
	type fields struct {
		Config map[string]interface{}
	}
	type args struct {
		options     cu.IM
		requestData string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "success",
			fields: fields{},
			args: args{
				options: cu.IM{
					"ds": &api.DataStore{
						Config: cu.IM{},
						Db: &td.TestDriver{
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
							return []byte(`{"user_name": "test", "user_group": "GROUP_USER"}`), nil
						},
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte([]byte(`{"user_name": "test", "user_group": "GROUP_USER"}`), v)
						},
						GetDataField: func(data any, JSONName string) (fieldName string, fieldValue interface{}) {
							return "user_name", "test"
						},
					},
					"id":   1,
					"code": "123456",
				},
				requestData: `{"user_name": "test", "user_group": "GROUP_USER"}`,
			},
			want: `{"code":204,"data":"No Content"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := &CLIService{
				Config: tt.fields.Config,
			}
			if got := cli.AuthUpdate(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.AuthUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_AuthQuery(t *testing.T) {
	type fields struct {
		Config map[string]interface{}
	}
	type args struct {
		options     cu.IM
		requestData string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "success",
			fields: fields{},
			args: args{
				options: cu.IM{
					"ds": &api.DataStore{
						Config: cu.IM{},
						Db: &td.TestDriver{
							Config: cu.IM{
								"Query": func(queries []md.Query) ([]cu.IM, error) {
									return []cu.IM{{"id": 1}}, nil
								},
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
							},
						},
					},
					"user_group": "GROUP_USER",
					"tag":        "test",
				},
			},
			want: `{"code":200,"data":[{"id":1}]}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := &CLIService{
				Config: tt.fields.Config,
			}
			if got := cli.AuthQuery(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.AuthQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_AuthGet(t *testing.T) {
	type fields struct {
		Config map[string]interface{}
	}
	type args struct {
		options     cu.IM
		requestData string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "success",
			fields: fields{},
			args: args{
				options: cu.IM{
					"ds": &api.DataStore{
						Config: cu.IM{},
						Db: &td.TestDriver{
							Config: cu.IM{
								"Query": func(queries []md.Query) ([]cu.IM, error) {
									return []cu.IM{{"id": 1}}, nil
								},
							},
						},
						AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					},
					"id":   1,
					"code": "123456",
				},
			},
			want: `{"code":200,"data":{"id":1}}`,
		},
		{
			name:   "not found",
			fields: fields{},
			args: args{
				options: cu.IM{
					"ds": &api.DataStore{
						Config: cu.IM{},
						Db:     &td.TestDriver{},
						AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					},
					"id":   1,
					"code": "123456",
				},
			},
			want: `{"code":404,"data":"Not Found"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := &CLIService{
				Config: tt.fields.Config,
			}
			if got := cli.AuthGet(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.AuthGet() = %v, want %v", got, tt.want)
			}
		})
	}
}
