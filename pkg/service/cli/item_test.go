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

func TestCLIService_ItemInsert(t *testing.T) {
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
								"Update": func(data md.Update) (int64, error) {
									return 1, nil
								},
							},
						},
						AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte([]byte(`{"trans_code": "test", "product_code": "test", "tax_code": "test", "code": "test"}`), v)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(cu.IM{"id": 1})
						},
					},
				},
				requestData: `{"trans_code": "test", "product_code": "test", "tax_code": "test"}`,
			},
			want: `{"code":201,"data":{"id":1}}`,
		},
		{
			name:   "trans_code is required",
			fields: fields{},
			args: args{
				options: cu.IM{
					"ds": &api.DataStore{
						Config: cu.IM{},
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte([]byte(`{"product_code": "test", "tax_code": "test"}`), v)
						},
					},
				},
				requestData: `{"product_code": "test", "tax_code": "test"}`,
			},
			want: `{"code":422,"data":"item trans_code, product_code and tax_code are required"}`,
		},
		{
			name:   "unprocessable entity",
			fields: fields{},
			args: args{
				options: cu.IM{
					"ds": &api.DataStore{
						Config: cu.IM{},
						ConvertFromByte: func(data []byte, v any) error {
							return errors.New("test")
						},
					},
				},
				requestData: `{"trans_code": "test", "product_code": "test", "tax_code": "test"}`,
			},
			want: `{"code":422,"data":"test"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := &CLIService{
				Config: tt.fields.Config,
			}
			if got := cli.ItemInsert(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.ItemInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_ItemUpdate(t *testing.T) {
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
							return []byte(`{}`), nil
						},
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte([]byte(`{"trans_code": "test", "product_code": "test", "tax_code": "test", "code": "test"}`), v)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(cu.IM{"id": 1})
						},
						GetDataField: func(data any, JSONName string) (fieldName string, fieldValue interface{}) {
							return "code", "123456"
						},
					},
					"id":   1,
					"code": "test",
				},
				requestData: `{"trans_code": "test", "product_code": "test", "tax_code": "test"}`,
			},
			want: `{"code":204,"data":"No Content"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := &CLIService{
				Config: tt.fields.Config,
			}
			if got := cli.ItemUpdate(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.ItemUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_ItemDelete(t *testing.T) {
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
							return []byte(`{}`), nil
						},
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte([]byte(`{"trans_code": "test", "product_code": "test", "tax_code": "test", "code": "test"}`), v)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(cu.IM{"id": 1})
						},
					},
					"id":   1,
					"code": "test",
				},
			},
			want: `{"code":204,"data":"No Content"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := &CLIService{
				Config: tt.fields.Config,
			}
			if got := cli.ItemDelete(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.ItemDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_ItemQuery(t *testing.T) {
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
						Db: &md.TestDriver{
							Config: cu.IM{
								"Query": func(queries []md.Query) ([]cu.IM, error) {
									return []cu.IM{{"id": 1}}, nil
								},
							},
						},
						AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
						ReadAll: func(r io.Reader) ([]byte, error) {
							return []byte(`{}`), nil
						},
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte([]byte(`{"trans_code": "test", "product_code": "test", "tax_code": "test", "code": "test"}`), v)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(cu.IM{"id": 1})
						},
					},
					"trans_code":   "test",
					"product_code": "test",
					"tag":          "test",
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
			if got := cli.ItemQuery(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.ItemQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_ItemGet(t *testing.T) {
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
						Db: &md.TestDriver{
							Config: cu.IM{
								"Query": func(queries []md.Query) ([]cu.IM, error) {
									return []cu.IM{{"id": 1}}, nil
								},
							},
						},
						AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
						ReadAll: func(r io.Reader) ([]byte, error) {
							return []byte(`{}`), nil
						},
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte([]byte(`{"trans_code": "test", "product_code": "test", "tax_code": "test", "code": "test"}`), v)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(cu.IM{"id": 1})
						},
					},
					"id":   1,
					"code": "test",
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
						AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
						ReadAll: func(r io.Reader) ([]byte, error) {
							return []byte(`{}`), nil
						},
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte([]byte(`{"trans_code": "test", "product_code": "test", "tax_code": "test", "code": "test"}`), v)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(cu.IM{"id": 1})
						},
					},
					"id":   1,
					"code": "test",
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
			if got := cli.ItemGet(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.ItemGet() = %v, want %v", got, tt.want)
			}
		})
	}
}
