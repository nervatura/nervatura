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

func TestCLIService_TransInsert(t *testing.T) {
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
						ReadAll: func(reader io.Reader) ([]byte, error) {
							return []byte(`{"code": "123456", "trans_code": "123456", "place_code": "123456", "project_code": "123456", "employee_code": "123456", "customer_code": "123456", "currency_code": "USD", "trans_date": "2024-01-01", "trans_type": "TRANS_INVOICE", "direction": "DIRECTION_OUT", "trans_meta": { "due_time": "2024-01-01", "status": "STATUS_NORMAL", "trans_state": "STATE_OK", "worksheet": { "worksheet_type": "WORK_ORDER", "worksheet_id": 1 } } }`), nil
						},
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte([]byte(`{"code": "123456", "trans_code": "123456", "place_code": "123456", "project_code": "123456", "employee_code": "123456", "customer_code": "123456", "currency_code": "USD", "trans_date": "2024-01-01", "trans_type": "TRANS_INVOICE", "direction": "DIRECTION_OUT", "trans_meta": { "due_time": "2024-01-01", "status": "STATUS_NORMAL", "trans_state": "STATE_OK", "worksheet": { "worksheet_type": "WORK_ORDER", "worksheet_id": 1 } } }`), v)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(cu.IM{"id": 1})
						},
					},
				},
			},
			want: `{"code":201,"data":{"id":1}}`,
		},
		{
			name:   "trans date is required",
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
						ReadAll: func(reader io.Reader) ([]byte, error) {
							return []byte(`{"trans_type": "TRANS_INVOICE", "direction": "DIRECTION_OUT"}`), nil
						},
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte(data, v)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(v)
						},
					},
				},
				requestData: `{"trans_type": "TRANS_INVOICE", "direction": "DIRECTION_OUT"}`,
			},
			want: `{"code":422,"data":"trans date is required"}`,
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
									return []cu.IM{{"id": 1}}, nil
								},
								"Update": func(data md.Update) (int64, error) {
									return 1, nil
								},
							},
						},
						AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
						ReadAll: func(reader io.Reader) ([]byte, error) {
							return []byte(`{"trans_type": "TRANS_INVOICE", "direction": "DIRECTION_OUT"}`), nil
						},
						ConvertFromByte: func(data []byte, v any) error {
							return errors.New("unprocessable entity")
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(v)
						},
					},
				},
				requestData: `{"trans_type": "TRANS_INVOICE", "direction": "DIRECTION_OUT"}`,
			},
			want: `{"code":422,"data":"unprocessable entity"}`,
		},
		{
			name:   "customer code required",
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
						ReadAll: func(reader io.Reader) ([]byte, error) {
							return []byte(`{"code": "123456", "trans_code": "123456", "place_code": "123456", "project_code": "123456", "employee_code": "123456",  "currency_code": "USD", "trans_date": "2024-01-01", "trans_type": "TRANS_INVOICE", "direction": "DIRECTION_OUT", "trans_meta": { "due_time": "2024-01-01", "status": "STATUS_NORMAL", "trans_state": "STATE_OK", "worksheet": { "worksheet_type": "WORK_ORDER", "worksheet_id": 1 } } }`), nil
						},
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte([]byte(`{"code": "123456", "trans_code": "123456", "place_code": "123456", "project_code": "123456", "employee_code": "123456",  "currency_code": "USD", "trans_date": "2024-01-01", "trans_type": "TRANS_INVOICE", "direction": "DIRECTION_OUT", "trans_meta": { "due_time": "2024-01-01", "status": "STATUS_NORMAL", "trans_state": "STATE_OK", "worksheet": { "worksheet_type": "WORK_ORDER", "worksheet_id": 1 } } }`), v)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(v)
						},
					},
				},
				requestData: `{"trans_type": "TRANS_INVOICE", "direction": "DIRECTION_OUT"}`,
			},
			want: `{"code":422,"data":"invoice, receipt, offer, order, worksheet and rent must have customer code and currency code"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := &CLIService{
				Config: tt.fields.Config,
			}
			if got := cli.TransInsert(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.TransInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_TransUpdate(t *testing.T) {
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
						ReadAll: func(reader io.Reader) ([]byte, error) {
							return []byte(`{"trans_type": "TRANS_INVOICE", "direction": "DIRECTION_OUT"}`), nil
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
					"id":   1,
					"code": "123456",
				},
				requestData: `{"trans_type": "TRANS_INVOICE", "direction": "DIRECTION_OUT"}`,
			},
			want: `{"code":204,"data":"No Content"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := &CLIService{
				Config: tt.fields.Config,
			}
			if got := cli.TransUpdate(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.TransUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_TransDelete(t *testing.T) {
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
						ReadAll: func(reader io.Reader) ([]byte, error) {
							return []byte(`{"trans_type": "TRANS_INVOICE", "direction": "DIRECTION_OUT"}`), nil
						},
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte(data, v)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(v)
						},
					},
					"id":   1,
					"code": "123456",
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
			if got := cli.TransDelete(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.TransDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_TransQuery(t *testing.T) {
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
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte(data, v)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(v)
						},
					},
					"trans_type": "TRANS_INVOICE",
					"direction":  "DIRECTION_OUT",
					"trans_date": "2024-01-01",
					"tag":        "tag1",
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
			if got := cli.TransQuery(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.TransQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_TransGet(t *testing.T) {
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
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte(data, v)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(v)
						},
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
						Db: &md.TestDriver{
							Config: cu.IM{
								"Query": func(queries []md.Query) ([]cu.IM, error) {
									return []cu.IM{}, nil
								},
							},
						},
						AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
						ConvertFromByte: func(data []byte, v any) error {
							return cu.ConvertFromByte(data, v)
						},
						ConvertToByte: func(v any) ([]byte, error) {
							return cu.ConvertToByte(v)
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
			if got := cli.TransGet(tt.args.options, tt.args.requestData); got != tt.want {
				t.Errorf("CLIService.TransGet() = %v, want %v", got, tt.want)
			}
		})
	}
}
