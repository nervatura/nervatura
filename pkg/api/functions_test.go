package api

import (
	"bytes"
	"errors"
	"io"
	"log/slog"
	"net"
	"net/http"
	"testing"

	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
)

func TestDataStore_ProductPrice(t *testing.T) {
	type fields struct {
		Db                     DataDriver
		Alias                  string
		Config                 cu.IM
		AppLog                 *slog.Logger
		ReadAll                func(r io.Reader) ([]byte, error)
		ConvertToByte          func(data interface{}) ([]byte, error)
		ConvertFromByte        func(data []byte, result interface{}) error
		ConvertFromReader      func(data io.Reader, result interface{}) error
		ConvertToType          func(data interface{}, result any) (err error)
		GetDataField           func(data any, JSONName string) (fieldName string, fieldValue interface{})
		NewRequest             func(method string, url string, body io.Reader) (*http.Request, error)
		RequestDo              func(req *http.Request) (*http.Response, error)
		CreateLoginToken       func(params cu.SM, config cu.IM) (result string, err error)
		ParseToken             func(token string, keyMap []cu.SM, config cu.IM) (cu.IM, error)
		CreatePasswordHash     func(password string) (hash string, err error)
		ComparePasswordAndHash func(password string, hash string) (err error)
		ReadFile               func(name string) ([]byte, error)
		NewSmtpClient          func(conn net.Conn, host string) (md.SmtpClient, error)
	}
	type args struct {
		options cu.IM
	}
	cn := 0
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							cn++
							if cn == 1 {
								return []cu.IM{{"mp": 100}}, nil
							}
							return []cu.IM{{"mp": 50}}, nil
						},
					},
				},
			},
			args: args{
				options: cu.IM{
					"product_code":  "123",
					"currency_code": "USD",
					"qty":           1,
					"customer_code": "123",
					"tag":           "test",
				},
			},
			wantErr: false,
		},
		{
			name: "query error 1",
			fields: fields{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return nil, errors.New("test")
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args: args{
				options: cu.IM{
					"product_code":  "123",
					"currency_code": "USD",
					"qty":           1,
					"customer_code": "123",
					"tag":           "test",
				},
			},
			wantErr: true,
		},
		{
			name: "query error 2",
			fields: fields{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							cn++
							if cn == 2 {
								return nil, errors.New("test")
							}
							return []cu.IM{{"mp": 50}}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args: args{
				options: cu.IM{
					"product_code":  "123",
					"currency_code": "USD",
					"qty":           1,
					"customer_code": "123",
					"tag":           "test",
				},
			},
			wantErr: true,
		},
		{
			name: "query error 3",
			fields: fields{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							cn++
							if cn == 3 {
								return nil, errors.New("test")
							}
							return []cu.IM{{"mp": 50}}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			args: args{
				options: cu.IM{
					"product_code":  "123",
					"currency_code": "USD",
					"qty":           1,
					"customer_code": "123",
					"tag":           "test",
				},
			},
			wantErr: true,
		},
		{
			name: "missing product code",
			fields: fields{
				Db: &md.TestDriver{
					Config: cu.IM{},
				},
			},
			args: args{
				options: cu.IM{
					"currency_code": "USD",
					"qty":           1,
					"customer_code": "123",
					"tag":           "test",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid price type",
			fields: fields{
				Db: &md.TestDriver{
					Config: cu.IM{},
				},
			},
			args: args{
				options: cu.IM{
					"product_code":  "123",
					"currency_code": "USD",
					"qty":           1,
					"customer_code": "123",
					"tag":           "test",
					"price_type":    "invalid",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &DataStore{
				Db:                     tt.fields.Db,
				Alias:                  tt.fields.Alias,
				Config:                 tt.fields.Config,
				AppLog:                 tt.fields.AppLog,
				ReadAll:                tt.fields.ReadAll,
				ConvertToByte:          tt.fields.ConvertToByte,
				ConvertFromByte:        tt.fields.ConvertFromByte,
				ConvertFromReader:      tt.fields.ConvertFromReader,
				ConvertToType:          tt.fields.ConvertToType,
				GetDataField:           tt.fields.GetDataField,
				NewRequest:             tt.fields.NewRequest,
				RequestDo:              tt.fields.RequestDo,
				CreateLoginToken:       tt.fields.CreateLoginToken,
				ParseToken:             tt.fields.ParseToken,
				CreatePasswordHash:     tt.fields.CreatePasswordHash,
				ComparePasswordAndHash: tt.fields.ComparePasswordAndHash,
				ReadFile:               tt.fields.ReadFile,
				NewSmtpClient:          tt.fields.NewSmtpClient,
			}
			cn = 0
			_, err := ds.ProductPrice(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataStore.ProductPrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDataStore_Function(t *testing.T) {
	type fields struct {
		Db                     DataDriver
		Alias                  string
		Config                 cu.IM
		AppLog                 *slog.Logger
		ReadAll                func(r io.Reader) ([]byte, error)
		ConvertToByte          func(data interface{}) ([]byte, error)
		ConvertFromByte        func(data []byte, result interface{}) error
		ConvertFromReader      func(data io.Reader, result interface{}) error
		ConvertToType          func(data interface{}, result any) (err error)
		GetDataField           func(data any, JSONName string) (fieldName string, fieldValue interface{})
		NewRequest             func(method string, url string, body io.Reader) (*http.Request, error)
		RequestDo              func(req *http.Request) (*http.Response, error)
		CreateLoginToken       func(params cu.SM, config cu.IM) (result string, err error)
		ParseToken             func(token string, keyMap []cu.SM, config cu.IM) (cu.IM, error)
		CreatePasswordHash     func(password string) (hash string, err error)
		ComparePasswordAndHash func(password string, hash string) (err error)
		ReadFile               func(name string) ([]byte, error)
		NewSmtpClient          func(conn net.Conn, host string) (md.SmtpClient, error)
	}
	type args struct {
		functionName string
		options      cu.IM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "test",
			fields: fields{},
			args: args{
				functionName: "test",
				options:      cu.IM{},
			},
			wantErr: false,
		},
		{
			name:   "missing function name",
			fields: fields{},
			args: args{
				options: cu.IM{},
			},
			wantErr: true,
		},
		{
			name:   "report_install",
			fields: fields{},
			args: args{
				functionName: "report_install",
				options:      cu.IM{},
			},
			wantErr: true,
		},
		{
			name: "report_list",
			fields: fields{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "report_key": "ntr_customer_en", "report_name": "test", "report_type": "test", "file_type": "FILE_CSV"}}, nil
						},
					},
				},
				ReadFile: func(name string) ([]byte, error) {
					return []byte(`{"meta": {"report_key": "ntr_customer_en", "report_name": "test", "report_type": "test", "file_type": "FILE_CSV"}}`), nil
				},
				ConvertFromByte: func(data []byte, result interface{}) error {
					return cu.ConvertFromByte([]byte(`{"meta": {"report_key": "ntr_customer_en", "report_name": "test", "report_type": "test", "file_type": "FILE_CSV"}}`), result)
				},
			},
			args: args{
				functionName: "report_list",
				options: cu.IM{
					"report_dir": "test",
					"label":      "test",
				},
			},
			wantErr: true,
		},
		{
			name:   "report_get",
			fields: fields{},
			args: args{
				functionName: "report_get",
				options:      cu.IM{},
			},
			wantErr: true,
		},
		{
			name:   "email_send",
			fields: fields{},
			args: args{
				functionName: "email_send",
				options:      cu.IM{},
			},
			wantErr: true,
		},
		{
			name:   "product_price",
			fields: fields{},
			args: args{
				functionName: "product_price",
				options:      cu.IM{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &DataStore{
				Db:                     tt.fields.Db,
				Alias:                  tt.fields.Alias,
				Config:                 tt.fields.Config,
				AppLog:                 tt.fields.AppLog,
				ReadAll:                tt.fields.ReadAll,
				ConvertToByte:          tt.fields.ConvertToByte,
				ConvertFromByte:        tt.fields.ConvertFromByte,
				ConvertFromReader:      tt.fields.ConvertFromReader,
				ConvertToType:          tt.fields.ConvertToType,
				GetDataField:           tt.fields.GetDataField,
				NewRequest:             tt.fields.NewRequest,
				RequestDo:              tt.fields.RequestDo,
				CreateLoginToken:       tt.fields.CreateLoginToken,
				ParseToken:             tt.fields.ParseToken,
				CreatePasswordHash:     tt.fields.CreatePasswordHash,
				ComparePasswordAndHash: tt.fields.ComparePasswordAndHash,
				ReadFile:               tt.fields.ReadFile,
				NewSmtpClient:          tt.fields.NewSmtpClient,
			}
			_, err := ds.Function(tt.args.functionName, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataStore.Function() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
