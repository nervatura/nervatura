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

func TestCLIService_ProjectInsert(t *testing.T) {
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
							return cu.ConvertFromByte([]byte(`{"code": "123456", "customer_code": "123456", "customer_type": "CUSTOMER_PRIVATE", "project_name": "John Doe"}`), v)
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
			name:   "unprocessable entity",
			fields: fields{},
			args: args{
				options: cu.IM{
					"ds": &api.DataStore{
						Config: cu.IM{},
						Db: &md.TestDriver{
							Config: cu.IM{
								"Query": func(queries []md.Query) ([]cu.IM, error) {
									return []cu.IM{}, nil
								},
								"Update": func(data md.Update) (int64, error) {
									return 0, errors.New("error")
								},
							},
						},
						AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
						ConvertFromByte: func(data []byte, v any) error {
							return errors.New("unprocessable entity")
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(cu.IM{"id": 1})
						},
					},
				},
				requestData: `{}`,
			},
			want: `{"code":422,"data":"unprocessable entity"}`,
		},
		{
			name:   "project name is required",
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
							return cu.ConvertFromByte([]byte(`{"code": "123456", "customer_code": "123456", "customer_type": "CUSTOMER_PRIVATE"}`), v)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(cu.IM{"id": 1})
						},
					},
				},
				requestData: `{}`,
			},
			want: `{"code":422,"data":"project name is required"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := &CLIService{
				Config: tt.fields.Config,
			}
			if got := cli.ProjectInsert(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.ProjectInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_ProjectUpdate(t *testing.T) {
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
						ReadAll: func(r io.Reader) ([]byte, error) {
							return []byte(`{"code": "123456", "customer_code": "123456", "project_name": "John Doe"}`), nil
						},
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte([]byte(`{"code": "123456", "customer_code": "123456", "project_name": "John Doe"}`), v)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(cu.IM{"id": 1})
						},
						GetDataField: func(data any, JSONName string) (fieldName string, fieldValue interface{}) {
							return "code", "123456"
						},
					},
					"id": 1,
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
			if got := cli.ProjectUpdate(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.ProjectUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_ProjectDelete(t *testing.T) {
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
							return []byte(`{"code": "123456", "customer_code": "123456", "project_name": "John Doe"}`), nil
						},
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte([]byte(`{"code": "123456", "customer_code": "123456", "project_name": "John Doe"}`), v)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(cu.IM{"id": 1})
						},
					},
					"id": 1,
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
			if got := cli.ProjectDelete(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.ProjectDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_ProjectQuery(t *testing.T) {
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
						ReadAll: func(r io.Reader) ([]byte, error) {
							return []byte(`[{"id": 1}]`), nil
						},
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte([]byte(`[{"id": 1}]`), v)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(cu.IM{"id": 1})
						},
					},
					"project_name":  "John Doe",
					"customer_code": "123456",
					"tag":           "tag1",
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
			if got := cli.ProjectQuery(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.ProjectQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_ProjectGet(t *testing.T) {
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
						ReadAll: func(r io.Reader) ([]byte, error) {
							return []byte(`[{"id": 1}]`), nil
						},
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte([]byte(`[{"id": 1}]`), v)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(cu.IM{"id": 1})
						},
					},
					"id": 1,
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
						Db: &md.TestDriver{
							Config: cu.IM{
								"Query": func(queries []md.Query) ([]cu.IM, error) {
									return []cu.IM{}, nil
								},
							},
						},
						AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
						ReadAll: func(r io.Reader) ([]byte, error) {
							return []byte(`[{"id": 1}]`), nil
						},
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte([]byte(`[{"id": 1}]`), v)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(cu.IM{"id": 1})
						},
					},
					"id": 1,
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
			if got := cli.ProjectGet(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.ProjectGet() = %v, want %v", got, tt.want)
			}
		})
	}
}
