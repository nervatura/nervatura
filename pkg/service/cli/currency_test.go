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

func TestCLIService_CurrencyInsert(t *testing.T) {
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
							return cu.ConvertFromByte([]byte(`{"currency_type": "CURRENCY_MAP", "code": "123456"}`), v)
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
		{
			name:   "code is required",
			fields: fields{},
			args: args{
				options: cu.IM{
					"ds": &api.DataStore{
						Db: &td.TestDriver{
							Config: cu.IM{},
						},
						AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte([]byte(`{}`), v)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(cu.IM{})
						},
					},
				},
				requestData: `{}`,
			},
			want: `{"code":422,"data":"code is required"}`,
		},
		{
			name:   "unprocessable entity",
			fields: fields{},
			args: args{
				options: cu.IM{
					"ds": &api.DataStore{
						Db: &td.TestDriver{
							Config: cu.IM{},
						},
						AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
						ConvertFromByte: func(data []byte, v any) error {
							return errors.New("unprocessable entity")
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(cu.IM{})
						},
					},
				},
				requestData: `{}`,
			},
			want: `{"code":422,"data":"unprocessable entity"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := &CLIService{
				Config: tt.fields.Config,
			}
			if got := cli.CurrencyInsert(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.CurrencyInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_CurrencyUpdate(t *testing.T) {
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
							return []byte(`{"code": "EUR"}`), nil
						},
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte([]byte(`{"code": "EUR"}`), v)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(cu.IM{"id": 1})
						},
						GetDataField: func(data any, JSONName string) (fieldName string, fieldValue interface{}) {
							return "code", "EUR"
						},
					},
					"id":   1,
					"code": "EUR",
				},
				requestData: ``,
			},
			want: `{"code":204,"data":"No Content"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := &CLIService{
				Config: tt.fields.Config,
			}
			if got := cli.CurrencyUpdate(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.CurrencyUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_CurrencyDelete(t *testing.T) {
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
							return []byte(`{"code": "EUR"}`), nil
						},
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte([]byte(`{"code": "EUR"}`), v)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(cu.IM{"id": 1})
						},
					},
					"id":   1,
					"code": "EUR",
				},
				requestData: ``,
			},
			want: `{"code":204,"data":"No Content"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := &CLIService{
				Config: tt.fields.Config,
			}
			if got := cli.CurrencyDelete(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.CurrencyDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_CurrencyQuery(t *testing.T) {
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
						Db: &td.TestDriver{
							Config: cu.IM{
								"Query": func(queries []md.Query) ([]cu.IM, error) {
									return []cu.IM{{"id": 1}}, nil
								},
							},
						},
						AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
					},
					"tag":  "tag",
					"id":   1,
					"code": "EUR",
				},
				requestData: ``,
			},
			want: `{"code":200,"data":[{"id":1}]}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := &CLIService{
				Config: tt.fields.Config,
			}
			if got := cli.CurrencyQuery(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.CurrencyQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_CurrencyGet(t *testing.T) {
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
						Db: &td.TestDriver{
							Config: cu.IM{
								"Query": func(queries []md.Query) ([]cu.IM, error) {
									return []cu.IM{{"id": 1}}, nil
								},
							},
						},
					},
					"id":   1,
					"code": "EUR",
				},
				requestData: ``,
			},
			want: `{"code":200,"data":{"id":1}}`,
		},
		{
			name:   "not found",
			fields: fields{},
			args: args{
				options: cu.IM{
					"ds": &api.DataStore{
						Db: &td.TestDriver{
							Config: cu.IM{},
						},
					},
					"id":   1,
					"code": "EUR",
				},
				requestData: ``,
			},
			want: `{"code":404,"data":"Not Found"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := &CLIService{
				Config: tt.fields.Config,
			}
			if got := cli.CurrencyGet(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.CurrencyGet() = %v, want %v", got, tt.want)
			}
		})
	}
}
