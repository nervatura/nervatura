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

func TestCLIService_EmployeeInsert(t *testing.T) {
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
							return []byte(`{"code": "123456"}`), nil
						},
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte([]byte(`{"code": "123456"}`), v)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(v)
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
								"Update": func(data md.Update) (int64, error) {
									return 1, nil
								},
							},
						},
						AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
						ReadAll: func(r io.Reader) ([]byte, error) {
							return []byte(`{}`), nil
						},
						ConvertFromByte: func(data []byte, v any) error {
							return errors.New("error")
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(v)
						},
					},
				},
				requestData: `{}`,
			},
			want: `{"code":422,"data":"error"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := &CLIService{
				Config: tt.fields.Config,
			}
			if got := cli.EmployeeInsert(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.EmployeeInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_EmployeeUpdate(t *testing.T) {
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
							return cu.ConvertFromByte(data, v)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(v)
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
			if got := cli.EmployeeUpdate(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.EmployeeUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_EmployeeDelete(t *testing.T) {
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
							return cu.ConvertFromByte(data, v)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(v)
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
			if got := cli.EmployeeDelete(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.EmployeeDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_EmployeeQuery(t *testing.T) {
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
							return []byte(`{}`), nil
						},
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte(data, v)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(v)
						},
					},
					"tag": "test",
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
			if got := cli.EmployeeQuery(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.EmployeeQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_EmployeeGet(t *testing.T) {
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
							return []byte(`{}`), nil
						},
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte(data, v)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(v)
						},
					},
					"id": 1,
				},
				requestData: `{}`,
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
							Config: cu.IM{},
						},
					},
					"id": 1,
				},
				requestData: `{}`,
			},
			want: `{"code":404,"data":"Not Found"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := &CLIService{
				Config: tt.fields.Config,
			}
			if got := cli.EmployeeGet(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.EmployeeGet() = %v, want %v", got, tt.want)
			}
		})
	}
}
