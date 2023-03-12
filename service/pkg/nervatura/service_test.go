package nervatura

import (
	"errors"
	"testing"

	smtpmock "github.com/mocktools/go-smtp-mock/v2"
)

func TestNervaStore_GetService(t *testing.T) {
	type fields struct {
		ds       DataDriver
		User     *User
		Customer IM
		models   IM
		config   IM
	}
	type args struct {
		key     string
		options IM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "nextNumber",
			fields: fields{
				ds: &testDriver{Config: IM{}},
			},
			args: args{
				key: "nextNumber",
				options: IM{
					"numberkey": "custnumber",
					"step":      false,
				},
			},
			wantErr: false,
		},
		{
			name: "getPriceValue",
			fields: fields{
				ds: &testDriver{Config: IM{}},
			},
			args: args{
				key: "getPriceValue",
				options: IM{
					"curr":        "EUR",
					"product_id":  2,
					"customer_id": 2,
				},
			},
			wantErr: false,
		},
		{
			name: "sendEmail",
			fields: fields{
				ds: &testDriver{Config: IM{}},
				config: IM{
					"NT_SMTP_HOST": "localhost",
					"NT_SMTP_PORT": 0,
				},
			},
			args: args{
				key: "sendEmail",
				options: IM{
					"provider": "smtp",
					"email": map[string]interface{}{
						"from": "info@nervatura.com", "name": "Nervatura",
						"recipients": []interface{}{
							map[string]interface{}{"email": "sample@company.com"}},
						"subject": "Demo Invoice",
						"text":    "Email sending with attached invoice",
						"attachments": []interface{}{
							map[string]interface{}{
								"reportkey": "ntr_invoice_en",
								"nervatype": "trans",
								"refnumber": "DMINV/00001"}},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "unknown_method",
			fields: fields{
				ds: &testDriver{Config: IM{}},
			},
			args: args{
				key:     "unknown",
				options: IM{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nstore := &NervaStore{
				ds:       tt.fields.ds,
				User:     tt.fields.User,
				Customer: tt.fields.Customer,
				models:   tt.fields.models,
				config:   tt.fields.config,
			}
			_, err := nstore.GetService(tt.args.key, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("NervaStore.GetService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNervaStore_nextNumber(t *testing.T) {
	type fields struct {
		ds       DataDriver
		User     *User
		Customer IM
		models   IM
		config   IM
	}
	type args struct {
		options IM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "step_true",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": int64(1), "curvalue": int64(0), "len": int64(5),
								"value": "value", "isyear": 1, "sep": "sep"},
						}, nil
					},
				}},
			},
			args: args{
				options: IM{
					"numberkey": "custnumber",
					"step":      true,
					"trans":     "trans",
				},
			},
			wantErr: false,
		},
		{
			name: "missing_numberkey",
			fields: fields{
				ds: &testDriver{Config: IM{}},
			},
			args: args{
				options: IM{
					"step":  true,
					"trans": "trans",
				},
			},
			wantErr: true,
		},
		{
			name: "not_connect",
			fields: fields{
				ds: &testDriver{Config: IM{
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
							Connected: false,
							Engine:    "test",
						}
					},
				}},
			},
			args: args{
				options: IM{
					"numberkey": "custnumber",
					"step":      true,
					"trans":     "trans",
				},
			},
			wantErr: true,
		},
		{
			name: "Query_error",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return nil, errors.New("error")
					},
					"RollbackTransaction": func() error {
						return errors.New("error")
					},
				}},
			},
			args: args{
				options: IM{
					"numberkey": "custnumber",
					"step":      true,
				},
			},
			wantErr: true,
		},
		{
			name: "Query_error_2",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						if queries[0].From == "fieldvalue" {
							return nil, errors.New("error")
						}
						return []IM{
							{"id": int64(1), "curvalue": int64(0), "len": int64(5),
								"value": "value", "isyear": 1, "sep": "sep"},
						}, nil
					},
				}},
			},
			args: args{
				options: IM{
					"numberkey": "custnumber",
					"step":      true,
				},
			},
			wantErr: true,
		},
		{
			name: "BeginTransaction_error",
			fields: fields{
				ds: &testDriver{Config: IM{
					"BeginTransaction": func() (interface{}, error) {
						return nil, errors.New("error")
					},
				}},
			},
			args: args{
				options: IM{
					"numberkey": "custnumber",
					"step":      true,
				},
			},
			wantErr: true,
		},
		{
			name: "Update_error_1",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": int64(1), "curvalue": int64(0), "len": int64(5),
								"value": "value", "isyear": 1, "sep": "sep"},
						}, nil
					},
					"Update": func(data Update) (int64, error) {
						return 0, errors.New("error")
					},
				}},
			},
			args: args{
				options: IM{
					"numberkey": "custnumber",
					"step":      true,
					"trans":     "trans",
				},
			},
			wantErr: true,
		},
		{
			name: "Update_error_2",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Update": func(data Update) (int64, error) {
						return 0, errors.New("error")
					},
				}},
			},
			args: args{
				options: IM{
					"numberkey": "custnumber",
					"step":      true,
					"trans":     "trans",
				},
			},
			wantErr: true,
		},
		{
			name: "insert_key",
			fields: fields{
				ds: &testDriver{Config: IM{}},
			},
			args: args{
				options: IM{
					"numberkey":  "custnumber",
					"step":       true,
					"trans":      "trans",
					"insert_key": false,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nstore := &NervaStore{
				ds:       tt.fields.ds,
				User:     tt.fields.User,
				Customer: tt.fields.Customer,
				models:   tt.fields.models,
				config:   tt.fields.config,
			}
			_, err := nstore.nextNumber(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("NervaStore.nextNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNervaStore_getPriceValue(t *testing.T) {
	type fields struct {
		ds       DataDriver
		User     *User
		Customer IM
		models   IM
		config   IM
	}
	type args struct {
		options IM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"discount": float64(10)},
						}, nil
					},
					"QueryKey": func(options IM) ([]IM, error) {
						if options["qkey"] == "custprice" {
							return []IM{
								{"mp": float64(50)},
							}, nil
						}
						if options["qkey"] == "grouprice" {
							return []IM{
								{"mp": float64(20)},
							}, nil
						}
						return []IM{
							{"mp": float64(100)},
						}, nil
					},
				}},
			},
			args: args{
				options: IM{
					"curr":        "EUR",
					"product_id":  2,
					"customer_id": 2,
				},
			},
			wantErr: false,
		},
		{
			name: "missing_curr",
			fields: fields{
				ds: &testDriver{Config: IM{}},
			},
			args: args{
				options: IM{
					"product_id":  2,
					"customer_id": 2,
				},
			},
			wantErr: true,
		},
		{
			name: "missing_product_id",
			fields: fields{
				ds: &testDriver{Config: IM{}},
			},
			args: args{
				options: IM{
					"curr":        "EUR",
					"customer_id": 2,
				},
			},
			wantErr: true,
		},
		{
			name: "QueryKey_error_listprice",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"discount": float64(10)},
						}, nil
					},
					"QueryKey": func(options IM) ([]IM, error) {
						if options["qkey"] == "listprice" {
							return nil, errors.New("error")
						}
						return []IM{
							{"mp": float64(100)},
						}, nil
					},
				}},
			},
			args: args{
				options: IM{
					"curr":        "EUR",
					"product_id":  2,
					"customer_id": 2,
				},
			},
			wantErr: true,
		},
		{
			name: "QueryKey_error_custprice",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"discount": float64(10)},
						}, nil
					},
					"QueryKey": func(options IM) ([]IM, error) {
						if options["qkey"] == "custprice" {
							return nil, errors.New("error")
						}
						return []IM{
							{"mp": float64(100)},
						}, nil
					},
				}},
			},
			args: args{
				options: IM{
					"curr":        "EUR",
					"product_id":  2,
					"customer_id": 2,
				},
			},
			wantErr: true,
		},
		{
			name: "QueryKey_error_grouprice",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"discount": float64(10)},
						}, nil
					},
					"QueryKey": func(options IM) ([]IM, error) {
						if options["qkey"] == "grouprice" {
							return nil, errors.New("error")
						}
						return []IM{
							{"mp": float64(100)},
						}, nil
					},
				}},
			},
			args: args{
				options: IM{
					"curr":        "EUR",
					"product_id":  2,
					"customer_id": 2,
				},
			},
			wantErr: true,
		},
		{
			name: "Query_error",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return nil, errors.New("error")
					},
					"QueryKey": func(options IM) ([]IM, error) {
						if options["qkey"] == "custprice" {
							return []IM{
								{"mp": float64(50)},
							}, nil
						}
						if options["qkey"] == "grouprice" {
							return []IM{
								{"mp": float64(20)},
							}, nil
						}
						return []IM{
							{"mp": float64(100)},
						}, nil
					},
				}},
			},
			args: args{
				options: IM{
					"curr":        "EUR",
					"product_id":  2,
					"customer_id": 2,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nstore := &NervaStore{
				ds:       tt.fields.ds,
				User:     tt.fields.User,
				Customer: tt.fields.Customer,
				models:   tt.fields.models,
				config:   tt.fields.config,
			}
			_, err := nstore.getPriceValue(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("NervaStore.getPriceValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNervaStore_createEmail(t *testing.T) {
	type fields struct {
		ds       DataDriver
		User     *User
		Customer IM
		models   IM
		config   IM
	}
	type args struct {
		from     string
		emailTo  []string
		emailOpt IM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "html_ok",
			fields: fields{
				ds: &testDriver{Config: IM{}},
			},
			args: args{
				from:    "info@nervatura.com",
				emailTo: []string{"sample@company.com"},
				emailOpt: IM{
					"from": "info@nervatura.com", "name": "Nervatura",
					"recipients": []interface{}{
						map[string]interface{}{"email": "sample@company.com"}},
					"subject": "Demo Invoice",
					"html":    "Email sending with attached invoice",
				},
			},
			wantErr: false,
		},
		{
			name: "text_attachments_error",
			fields: fields{
				ds: &testDriver{Config: IM{}},
			},
			args: args{
				from:    "info@nervatura.com",
				emailTo: []string{"sample@company.com"},
				emailOpt: IM{
					"from": "info@nervatura.com", "name": "Nervatura",
					"recipients": []interface{}{
						map[string]interface{}{"email": "sample@company.com"}},
					"subject": "Demo Invoice",
					"text":    "Email sending with attached invoice",
					"attachments": []interface{}{
						map[string]interface{}{
							"reportkey": "ntr_invoice_en",
							"nervatype": "trans",
							"refnumber": "DMINV/00001",
							"filename":  "filename.pdf",
							"report_id": int64(123),
							"ref_id":    int64(123),
						}},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nstore := &NervaStore{
				ds:       tt.fields.ds,
				User:     tt.fields.User,
				Customer: tt.fields.Customer,
				models:   tt.fields.models,
				config:   tt.fields.config,
			}
			_, err := nstore.createEmail(tt.args.from, tt.args.emailTo, tt.args.emailOpt)
			if (err != nil) != tt.wantErr {
				t.Errorf("NervaStore.createEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNervaStore_sendEmail(t *testing.T) {
	type fields struct {
		ds       DataDriver
		User     *User
		Customer IM
		models   IM
		config   IM
	}
	type args struct {
		options IM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "send_ok",
			fields: fields{
				ds: &testDriver{Config: IM{}},
				config: IM{
					"NT_SMTP_HOST":            "localhost",
					"NT_SMTP_PORT":            2525,
					"NT_SMTP_CONN":            "net",
					"NT_SMTP_AUTH":            "none",
					"NT_SMTP_TLS_MIN_VERSION": 769,
				},
			},
			args: args{
				options: IM{
					"provider": "smtp",
					"email": map[string]interface{}{
						"from": "info@nervatura.com", "name": "Nervatura",
						"recipients": []interface{}{
							map[string]interface{}{"email": "sample@company.com"}},
						"subject":     "Demo Invoice",
						"text":        "Email sending with attached invoice",
						"attachments": []interface{}{},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "missing_email",
			args: args{
				options: IM{},
			},
			wantErr: true,
		},
		{
			name: "invalid_provider",
			args: args{
				options: IM{
					"provider": "xxxx",
					"email": map[string]interface{}{
						"from": "info@nervatura.com", "name": "Nervatura",
						"recipients": []interface{}{
							map[string]interface{}{"email": "sample@company.com"}},
						"subject":     "Demo Invoice",
						"text":        "Email sending with attached invoice",
						"attachments": []interface{}{},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "auth_err",
			fields: fields{
				ds: &testDriver{Config: IM{}},
				config: IM{
					"NT_SMTP_HOST":            "localhost",
					"NT_SMTP_PORT":            2525,
					"NT_SMTP_CONN":            "net",
					"NT_SMTP_AUTH":            "auth",
					"NT_SMTP_TLS_MIN_VERSION": 769,
				},
			},
			args: args{
				options: IM{
					"provider": "smtp",
					"email": map[string]interface{}{
						"from": "info@nervatura.com", "name": "Nervatura",
						"recipients": []interface{}{
							map[string]interface{}{"email": "sample@company.com"}},
						"subject":     "Demo Invoice",
						"text":        "Email sending with attached invoice",
						"attachments": []interface{}{},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "missing_recipients",
			fields: fields{
				ds: &testDriver{Config: IM{}},
				config: IM{
					"NT_SMTP_HOST":            "localhost",
					"NT_SMTP_PORT":            2525,
					"NT_SMTP_CONN":            "net",
					"NT_SMTP_AUTH":            "none",
					"NT_SMTP_TLS_MIN_VERSION": 769,
				},
			},
			args: args{
				options: IM{
					"provider": "smtp",
					"email": map[string]interface{}{
						"from": "info@nervatura.com", "name": "Nervatura",
						"subject":     "Demo Invoice",
						"text":        "Email sending with attached invoice",
						"attachments": []interface{}{},
					},
				},
			},
			wantErr: true,
		},
	}
	server := smtpmock.New(smtpmock.ConfigurationAttr{
		HostAddress:       "localhost",
		PortNumber:        2525,
		LogToStdout:       false,
		LogServerActivity: false,
	})
	server.Start()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nstore := &NervaStore{
				ds:       tt.fields.ds,
				User:     tt.fields.User,
				Customer: tt.fields.Customer,
				models:   tt.fields.models,
				config:   tt.fields.config,
			}
			_, err := nstore.sendEmail(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("NervaStore.sendEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
	server.Stop()
}
