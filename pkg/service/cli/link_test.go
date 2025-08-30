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

func TestCLIService_LinkInsert(t *testing.T) {
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
			name: "success",
			fields: fields{
				Config: cu.IM{},
			},
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
							return cu.ConvertFromByte([]byte(`{"code": "test", "link_type1": "LINK_TRANS", "link_type2": "LINK_TRANS", "link_code_1": "test", "link_code_2": "test", "link_meta": {"tags": []}, "link_map": {}}`), v)
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
							},
						},
						AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
						ReadAll: func(r io.Reader) ([]byte, error) {
							return []byte(`{}`), nil
						},
						ConvertFromByte: func(data []byte, v any) error {
							return errors.New("error")
						},
					},
				},
				requestData: `{}`,
			},
			want: `{"code":422,"data":"error"}`,
		},
		{
			name:   "required fields error",
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
							return []byte(`{}`), nil
						},
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte([]byte(`{"code": "test", "link_code_2": "test", "link_meta": {"tags": []}, "link_map": {}}`), v)
						},
					},
				},
				requestData: `{}`,
			},
			want: `{"code":422,"data":"link link_code_1 and link_code_2 are required"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := &CLIService{
				Config: tt.fields.Config,
			}
			if got := cli.LinkInsert(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.LinkInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_LinkUpdate(t *testing.T) {
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
			name: "success",
			fields: fields{
				Config: cu.IM{},
			},
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
							return []byte(`{"code": "test", "link_type_1": "LINK_TRANS", "link_type_2": "LINK_TRANS", "link_code_1": "test", "link_code_2": "test", "link_meta": {"tags": []}, "link_map": {}}`), nil
						},
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte([]byte(`{"code": "test", "link_type_1": "LINK_TRANS", "link_type_2": "LINK_TRANS", "link_code_1": "test", "link_code_2": "test", "link_meta": {"tags": []}, "link_map": {}}`), v)
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
			if got := cli.LinkUpdate(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.LinkUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_LinkDelete(t *testing.T) {
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
			name: "success",
			fields: fields{
				Config: cu.IM{},
			},
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
			if got := cli.LinkDelete(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.LinkDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_LinkQuery(t *testing.T) {
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
			name: "success",
			fields: fields{
				Config: cu.IM{},
			},
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
					},
					"link_type_1": "LINK_TRANS",
					"link_code_1": "test",
					"link_type_2": "LINK_TRANS",
					"link_code_2": "test",
					"tag":         "test",
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
			if got := cli.LinkQuery(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.LinkQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_LinkGet(t *testing.T) {
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
			name: "success",
			fields: fields{
				Config: cu.IM{},
			},
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
					},
					"id": 1,
				},
			},
			want: `{"code":200,"data":{"id":1}}`,
		},
		{
			name: "not found",
			fields: fields{
				Config: cu.IM{},
			},
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
			if got := cli.LinkGet(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.LinkGet() = %v, want %v", got, tt.want)
			}
		})
	}
}
