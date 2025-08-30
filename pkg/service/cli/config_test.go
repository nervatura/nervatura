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
)

func TestCLIService_ConfigInsert(t *testing.T) {
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
						Db:     &md.TestDriver{},
						AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
						ConvertFromByte: func(data []byte, v any) error {
							return errors.New("test error")
						},
					},
				},
				requestData: `{}`,
			},
			want: `{"code":422,"data":"test error"}`,
		},
		{
			name:   "config type unprocessable entity",
			fields: fields{},
			args: args{
				options: cu.IM{
					"ds": &api.DataStore{
						Config: cu.IM{},
						Db:     &md.TestDriver{},
						AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
						ConvertFromByte: func(data []byte, v any) error {
							strData := string(data)
							if strData == `{}` {
								return cu.ConvertFromByte([]byte(`{"config_type": "CONFIG_MAP", "code": "123456"}`), v)
							}
							return errors.New("test error")
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(cu.IM{"id": "error"})
						},
					},
				},
				requestData: `{}`,
			},
			want: `{"code":422,"data":"test error"}`,
		},
		{
			name:   "success",
			fields: fields{},
			args: args{
				options: cu.IM{
					"ds": &api.DataStore{
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
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte([]byte(`{"config_type": "CONFIG_MAP", "code": "123456"}`), v)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(cu.IM{"id": 1})
						},
					},
				},
				requestData: `{}`,
			},
			want: `{"code":201,"data":{"id":1}}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := &CLIService{
				Config: tt.fields.Config,
			}
			if got := cli.ConfigInsert(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.ConfigInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_ConfigUpdate(t *testing.T) {
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
							return []byte(`{"config_type": "CONFIG_MAP", "code": "123456"}`), nil
						},
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte([]byte(`{"config_type": "CONFIG_MAP", "code": "123456", "data": {}}`), v)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(cu.IM{"id": 1})
						},
						GetDataField: func(data any, JSONName string) (fieldName string, fieldValue interface{}) {
							return "code", "123456"
						},
					},
					"id":   1,
					"code": "123456",
				},
				requestData: ``,
			},
			want: `{"code":204,"data":"No Content"}`,
		},
		{
			name:   "missing config_type",
			fields: fields{},
			args: args{
				options: cu.IM{
					"ds": &api.DataStore{
						Config: cu.IM{},
						Db:     &md.TestDriver{},
						AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
						ReadAll: func(r io.Reader) ([]byte, error) {
							return []byte(`{"code": "123456", "data": {}}`), nil
						},
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte([]byte(`{"code": "123456", "data": {}}`), v)
						},
						GetDataField: func(data any, JSONName string) (fieldName string, fieldValue interface{}) {
							return "code", "123456"
						},
					},
					"id":   1,
					"code": "123456",
				},
				requestData: `{}`,
			},
			want: `{"code":422,"data":"config_type is required"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := &CLIService{
				Config: tt.fields.Config,
			}
			if got := cli.ConfigUpdate(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.ConfigUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_ConfigDelete(t *testing.T) {
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
					},
					"id":   1,
					"code": "123456",
				},
				requestData: `{}`,
			},
			want: `{"code":204,"data":"No Content"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := &CLIService{
				Config: tt.fields.Config,
			}
			if got := cli.ConfigDelete(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.ConfigDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_ConfigQuery(t *testing.T) {
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
						Db: &md.TestDriver{
							Config: cu.IM{
								"Query": func(queries []md.Query) ([]cu.IM, error) {
									return []cu.IM{{"id": 1}}, nil
								},
							},
						},
						AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					},
					"config_type": "CONFIG_MAP",
				},
				requestData: `{}`,
			},
			want: `{"code":200,"data":[{"id":1}]}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := &CLIService{
				Config: tt.fields.Config,
			}
			if got := cli.ConfigQuery(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.ConfigQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_ConfigGet(t *testing.T) {
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
						Db: &md.TestDriver{
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
						Db: &md.TestDriver{
							Config: cu.IM{
								"Query": func(queries []md.Query) ([]cu.IM, error) {
									return []cu.IM{}, nil
								},
							},
						},
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
			if got := cli.ConfigGet(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.ConfigGet() = %v, want %v", got, tt.want)
			}
		})
	}
}
