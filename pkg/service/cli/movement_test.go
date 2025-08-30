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

func TestCLIService_MovementInsert(t *testing.T) {
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
								"Update": func(data md.Update) (int64, error) {
									return 1, nil
								},
							},
						},
						AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte([]byte(`{"code": "123456", "tool_code": "123456", "place_code": "123456", "movement_type":"MOVEMENT_INVENTORY","shipping_time":"2024-01-01","trans_code":"123456", "product_code":"123456"}`), v)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(cu.IM{"id": 1})
						},
					},
				},
				requestData: `{"movement_type":"MOVEMENT_INVENTORY","shipping_time":"2024-01-01","trans_code":"123456"}`,
			},
			want: `{"code":201,"data":{"id":1}}`,
		},
		{
			name: "unprocessable entity",
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
									return 0, errors.New("error")
								},
							},
						},
						AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
						ConvertFromByte: func(data []byte, v any) error {
							return errors.New("error")
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(cu.IM{"id": 1})
						},
					},
				},
			},
			want: `{"code":422,"data":"error"}`,
		},
		{
			name: "required fields shipping_time and trans_code",
			fields: fields{
				Config: cu.IM{},
			},
			args: args{
				options: cu.IM{
					"ds": &api.DataStore{
						Config: cu.IM{},
						Db: &md.TestDriver{
							Config: cu.IM{},
						},
						AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte([]byte(`{"movement_type":"MOVEMENT_INVENTORY"}`), v)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(cu.IM{"id": 1})
						},
					},
				},
				requestData: `{"movement_type":"MOVEMENT_INVENTORY"}`,
			},
			want: `{"code":422,"data":"movement shipping_time and trans_code are required"}`,
		},
		{
			name: "required fields product_code or tool_code",
			fields: fields{
				Config: cu.IM{},
			},
			args: args{
				options: cu.IM{
					"ds": &api.DataStore{
						Config: cu.IM{},
						Db: &md.TestDriver{
							Config: cu.IM{},
						},
						AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte([]byte(`{"movement_type":"MOVEMENT_INVENTORY","shipping_time":"2024-01-01","trans_code":"123456"}`), v)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(cu.IM{"id": 1})
						},
					},
				},
				requestData: `{"movement_type":"MOVEMENT_INVENTORY","shipping_time":"2024-01-01","trans_code":"123456"}`,
			},
			want: `{"code":422,"data":"movement product_code or tool_code are required"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := &CLIService{
				Config: tt.fields.Config,
			}
			if got := cli.MovementInsert(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.MovementInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_MovementUpdate(t *testing.T) {
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
						ReadAll: func(reader io.Reader) ([]byte, error) {
							return cu.ConvertToByte([]byte(`{"movement_type":"MOVEMENT_INVENTORY","shipping_time":"2024-01-01","trans_code":"123456"}`))
						},
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte([]byte(`{"movement_type":"MOVEMENT_INVENTORY","shipping_time":"2024-01-01","trans_code":"123456"}`), v)
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
				requestData: `{"movement_type":"MOVEMENT_INVENTORY","shipping_time":"2024-01-01","trans_code":"123456"}`,
			},
			want: `{"code":204,"data":"No Content"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := &CLIService{
				Config: tt.fields.Config,
			}
			if got := cli.MovementUpdate(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.MovementUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_MovementDelete(t *testing.T) {
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
						ReadAll: func(reader io.Reader) ([]byte, error) {
							return cu.ConvertToByte([]byte(`{"movement_type":"MOVEMENT_INVENTORY","shipping_time":"2024-01-01","trans_code":"123456"}`))
						},
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte([]byte(`{"movement_type":"MOVEMENT_INVENTORY","shipping_time":"2024-01-01","trans_code":"123456"}`), v)
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
			if got := cli.MovementDelete(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.MovementDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_MovementQuery(t *testing.T) {
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
						AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
						ReadAll: func(reader io.Reader) ([]byte, error) {
							return cu.ConvertToByte([]byte(`{"movement_type":"MOVEMENT_INVENTORY","shipping_time":"2024-01-01","trans_code":"123456"}`))
						},
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte([]byte(`{"movement_type":"MOVEMENT_INVENTORY","shipping_time":"2024-01-01","trans_code":"123456"}`), v)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(cu.IM{"id": 1})
						},
					},
					"trans_code":    "123456",
					"movement_type": "MOVEMENT_INVENTORY",
					"tag":           "test",
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
			if got := cli.MovementQuery(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.MovementQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_MovementGet(t *testing.T) {
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
						AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
						ReadAll: func(reader io.Reader) ([]byte, error) {
							return cu.ConvertToByte([]byte(`{"movement_type":"MOVEMENT_INVENTORY","shipping_time":"2024-01-01","trans_code":"123456"}`))
						},
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte([]byte(`{"movement_type":"MOVEMENT_INVENTORY","shipping_time":"2024-01-01","trans_code":"123456"}`), v)
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
						AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
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
			if got := cli.MovementGet(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.MovementGet() = %v, want %v", got, tt.want)
			}
		})
	}
}
