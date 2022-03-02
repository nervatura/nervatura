package nervatura

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	type args struct {
		driver DataDriver
		config IM
	}
	tests := []struct {
		name       string
		args       args
		wantNstore *NervaStore
	}{
		{
			name: "ok",
			args: args{
				driver: &testDriver{Config: IM{}},
				config: IM{},
			},
			wantNstore: New(&testDriver{Config: IM{}}, IM{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNstore := New(tt.args.driver, tt.args.config); !reflect.DeepEqual(gotNstore, tt.wantNstore) {
				t.Errorf("New() = %v, want %v", gotNstore, tt.wantNstore)
			}
		})
	}
}

func TestNervaStore_connected(t *testing.T) {
	type fields struct {
		ds       DataDriver
		User     *User
		Customer IM
		models   IM
		config   IM
	}
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				ds: &testDriver{Config: IM{}},
			},
			want:    true,
			wantErr: false,
		},
		{
			name:    "error",
			want:    false,
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
			got, err := nstore.connected()
			if (err != nil) != tt.wantErr {
				t.Errorf("NervaStore.connected() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NervaStore.connected() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validFieldValue(t *testing.T) {
	type args struct {
		nervatype string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid",
			args: args{
				nervatype: "trans",
			},
			want: true,
		},
		{
			name: "groups",
			args: args{
				nervatype: "groups",
			},
			want: false,
		},
		{
			name: "ui_report",
			args: args{
				nervatype: "ui_report",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validFieldValue(tt.args.nervatype); got != tt.want {
				t.Errorf("validFieldValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNervaStore_getTableKey(t *testing.T) {
	type fields struct {
		ds       DataDriver
		User     *User
		Customer IM
		models   IM
		config   IM
	}
	type args struct {
		nervatype string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "trans",
			fields: fields{
				models: DataModel()["model"].(IM),
			},
			args: args{
				nervatype: "trans",
			},
			want: "transnumber",
		},
		{
			name: "address",
			fields: fields{
				models: DataModel()["model"].(IM),
			},
			args: args{
				nervatype: "address",
			},
			want: "",
		},
		{
			name: "missing",
			fields: fields{
				models: DataModel()["model"].(IM),
			},
			args: args{
				nervatype: "missing",
			},
			want: "",
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
			if got := nstore.getTableKey(tt.args.nervatype); got != tt.want {
				t.Errorf("NervaStore.getTableKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkFieldvalueBool(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "nil_false",
			args: args{
				value: nil,
			},
			want: "false",
		},
		{
			name: "true",
			args: args{
				value: true,
			},
			want: "true",
		},
		{
			name: "false",
			args: args{
				value: 0,
			},
			want: "false",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkFieldvalueBool(tt.args.value)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkFieldvalueBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkFieldvalueDate(t *testing.T) {
	type args struct {
		value     interface{}
		fieldname string
		fieldtype string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "nil",
			args: args{
				value:     nil,
				fieldname: "fieldname",
				fieldtype: "fieldtype",
			},
			wantErr: true,
		},
		{
			name: "time",
			args: args{
				value:     time.Now(),
				fieldname: "fieldname",
				fieldtype: "fieldtype",
			},
			wantErr: false,
		},
		{
			name: "string_error",
			args: args{
				value:     "datestring",
				fieldname: "fieldname",
				fieldtype: "fieldtype",
			},
			wantErr: true,
		},
		{
			name: "string_valid",
			args: args{
				value:     "2021-01-01",
				fieldname: "fieldname",
				fieldtype: "fieldtype",
			},
			wantErr: false,
		},
		{
			name: "invalid",
			args: args{
				value:     1234,
				fieldname: "fieldname",
				fieldtype: "fieldtype",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := checkFieldvalueDate(tt.args.value, tt.args.fieldname, tt.args.fieldtype)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkFieldvalueDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_checkFieldvalueTime(t *testing.T) {
	type args struct {
		value     interface{}
		fieldname string
		fieldtype string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "nil",
			args: args{
				value:     nil,
				fieldname: "fieldname",
				fieldtype: "fieldtype",
			},
			wantErr: true,
		},
		{
			name: "time",
			args: args{
				value:     time.Now(),
				fieldname: "fieldname",
				fieldtype: "fieldtype",
			},
			wantErr: false,
		},
		{
			name: "string_error",
			args: args{
				value:     "datestring",
				fieldname: "fieldname",
				fieldtype: "fieldtype",
			},
			wantErr: true,
		},
		{
			name: "string_valid",
			args: args{
				value:     "22:54",
				fieldname: "fieldname",
				fieldtype: "fieldtype",
			},
			wantErr: false,
		},
		{
			name: "invalid",
			args: args{
				value:     1234,
				fieldname: "fieldname",
				fieldtype: "fieldtype",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := checkFieldvalueTime(tt.args.value, tt.args.fieldname, tt.args.fieldtype)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkFieldvalueTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNervaStore_checkFieldvalueNervatype(t *testing.T) {
	type fields struct {
		ds       DataDriver
		User     *User
		Customer IM
		models   IM
		config   IM
	}
	type args struct {
		value     interface{}
		fieldname string
		fieldtype string
		trans     interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "trans_id_ok",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{{"id": 1}}, nil
					},
				}},
			},
			args: args{
				fieldname: "fieldname",
				fieldtype: "transitem",
				value:     int64(1),
			},
			wantErr: false,
		},
		{
			name: "trans_id_string_nf",
			fields: fields{
				ds: &testDriver{Config: IM{}},
			},
			args: args{
				fieldname: "fieldname",
				fieldtype: "transitem",
				value:     "1",
			},
			wantErr: true,
		},
		{
			name: "customer_string_error",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return nil, errors.New("error")
					},
				}},
			},
			args: args{
				fieldname: "fieldname",
				fieldtype: "customer",
				value:     "custnumber",
			},
			wantErr: true,
		},
		{
			name: "customer_invalid_value",
			fields: fields{
				ds: &testDriver{Config: IM{}},
			},
			args: args{
				fieldname: "fieldname",
				fieldtype: "customer",
				value:     true,
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
			_, err := nstore.checkFieldvalueNervatype(tt.args.value, tt.args.fieldname, tt.args.fieldtype, tt.args.trans)
			if (err != nil) != tt.wantErr {
				t.Errorf("NervaStore.checkFieldvalueNervatype() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNervaStore_checkFieldvalue(t *testing.T) {
	type fields struct {
		ds       DataDriver
		User     *User
		Customer IM
		models   IM
		config   IM
	}
	type args struct {
		fieldname string
		value     interface{}
		trans     interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "boolField",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"fieldtype": "bool"},
						}, nil
					},
				}},
			},
			args: args{
				fieldname: "boolField",
				value:     true,
			},
			wantErr: false,
		},
		{
			name: "intField",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"fieldtype": "integer"},
						}, nil
					},
				}},
			},
			args: args{
				fieldname: "intField",
				value:     int64(123),
			},
			wantErr: false,
		},
		{
			name: "floatField",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"fieldtype": "float"},
						}, nil
					},
				}},
			},
			args: args{
				fieldname: "floatField",
				value:     float64(123),
			},
			wantErr: false,
		},
		{
			name: "dateField",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"fieldtype": "date"},
						}, nil
					},
				}},
			},
			args: args{
				fieldname: "dateField",
				value:     "2021-12-12",
			},
			wantErr: false,
		},
		{
			name: "timeField",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"fieldtype": "time"},
						}, nil
					},
				}},
			},
			args: args{
				fieldname: "timeField",
				value:     "16:10",
			},
			wantErr: false,
		},
		{
			name: "stringField",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"fieldtype": "string"},
						}, nil
					},
				}},
			},
			args: args{
				fieldname: "stringField",
				value:     "string",
			},
			wantErr: false,
		},
		{
			name: "product",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"fieldtype": "product"},
						}, nil
					},
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{{"id": 1}}, nil
					},
				}},
			},
			args: args{
				fieldname: "product",
				value:     "product",
			},
			wantErr: false,
		},
		{
			name: "invalid",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"fieldtype": "invalid"},
						}, nil
					},
				}},
			},
			args: args{
				fieldname: "invalid",
				value:     "invalid",
			},
			wantErr: true,
		},
		{
			name: "QueryKey_error",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return nil, errors.New("error")
					},
				}},
			},
			args: args{
				fieldname: "error",
				value:     "error",
			},
			wantErr: true,
		},
		{
			name: "missing",
			fields: fields{
				ds: &testDriver{Config: IM{}},
			},
			args: args{
				fieldname: "missing",
				value:     "missing",
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
			_, err := nstore.checkFieldvalue(tt.args.fieldname, tt.args.value, tt.args.trans)
			if (err != nil) != tt.wantErr {
				t.Errorf("NervaStore.checkFieldvalue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNervaStore_insertLog(t *testing.T) {
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
			name: "log_ok",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": int64(1), "fieldname": "logstate", "value": true},
							{"id": int64(2), "fieldname": "nervatype", "value": true},
							{"id": int64(3), "fieldname": "log_logstate", "value": true},
						}, nil
					},
				}},
				User: &User{Id: int64(1)},
			},
			args: args{
				options: IM{
					"logstate":  "logstate",
					"nervatype": "nervatype",
					"trans":     "trans",
					"ref_id":    123,
				},
			},
			wantErr: false,
		},
		{
			name: "log_nervatype_logstate",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": int64(4), "fieldname": "log_nervatype_logstate", "value": true},
						}, nil
					},
				}},
				User: &User{Id: int64(1)},
			},
			args: args{
				options: IM{
					"logstate":  "logstate",
					"nervatype": "nervatype",
					"trans":     "trans",
					"ref_id":    123,
				},
			},
			wantErr: false,
		},
		{
			name: "missing_nervatype",
			fields: fields{
				ds:   &testDriver{Config: IM{}},
				User: &User{Id: int64(1)},
			},
			args: args{
				options: IM{
					"logstate": "logstate",
					"trans":    "trans",
					"ref_id":   123,
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
				User: &User{Id: int64(1)},
			},
			args: args{
				options: IM{
					"logstate":  "logstate",
					"nervatype": "nervatype",
					"trans":     "trans",
					"ref_id":    123,
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
				}},
				User: &User{Id: int64(1)},
			},
			args: args{
				options: IM{
					"logstate":  "logstate",
					"nervatype": "nervatype",
					"trans":     "trans",
					"ref_id":    123,
				},
			},
			wantErr: true,
		},
		{
			name: "Update_error",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": int64(1), "fieldname": "logstate", "value": true},
							{"id": int64(2), "fieldname": "nervatype", "value": true},
							{"id": int64(3), "fieldname": "log_logstate", "value": true},
						}, nil
					},
					"Update": func(data Update) (int64, error) {
						return 0, errors.New("error")
					},
				}},
				User: &User{Id: int64(1)},
			},
			args: args{
				options: IM{
					"logstate":  "logstate",
					"nervatype": "nervatype",
					"trans":     "trans",
					"ref_id":    123,
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
			if err := nstore.insertLog(tt.args.options); (err != nil) != tt.wantErr {
				t.Errorf("NervaStore.insertLog() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNervaStore_updateValidate(t *testing.T) {
	type fields struct {
		ds       DataDriver
		User     *User
		Customer IM
		models   IM
		config   IM
	}
	type args struct {
		nervatype   string
		checkValues IM
		trans       interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "integer_float_string_bool",
			fields: fields{
				ds:     &testDriver{Config: IM{}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				nervatype: "customer",
				checkValues: IM{
					"values": IM{
						"id":          0,
						"terms":       123,
						"creditlimit": 123,
						"custnumber":  nil,
						"custname":    nil,
						"taxnumber":   nil,
						"inactive":    int64(1),
						"notax":       0,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "string_length",
			fields: fields{
				ds:     &testDriver{Config: IM{}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				nervatype: "ui_message",
				checkValues: IM{
					"values": IM{
						"lang": "lang",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "date_datetime_curr",
			fields: fields{
				ds:     &testDriver{Config: IM{}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				nervatype: "trans",
				checkValues: IM{
					"values": IM{
						"duedate":   nil,
						"transdate": "2021-01-01",
						"crdate":    time.Now(),
						"curr":      "",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "datetime",
			fields: fields{
				ds:     &testDriver{Config: IM{}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				nervatype: "event",
				checkValues: IM{
					"values": IM{
						"fromdate": "2021-01-01",
						"todate":   time.Now(),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "date_error",
			fields: fields{
				ds:     &testDriver{Config: IM{}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				nervatype: "employee",
				checkValues: IM{
					"values": IM{
						"startdate": nil,
						"enddate":   "20210101",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "date_error",
			fields: fields{
				ds:     &testDriver{Config: IM{}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				nervatype: "employee",
				checkValues: IM{
					"values": IM{
						"enddate": 123,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "datetime_error",
			fields: fields{
				ds:     &testDriver{Config: IM{}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				nervatype: "event",
				checkValues: IM{
					"values": IM{
						"fromdate": "20210101",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "datetime_error",
			fields: fields{
				ds:     &testDriver{Config: IM{}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				nervatype: "event",
				checkValues: IM{
					"values": IM{
						"fromdate": 123,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "notnull_error",
			fields: fields{
				ds:     &testDriver{Config: IM{}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				nervatype: "ui_menu",
				checkValues: IM{
					"values": IM{
						"method": nil,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "min_error",
			fields: fields{
				ds:     &testDriver{Config: IM{}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				nervatype: "item",
				checkValues: IM{
					"values": IM{
						"discount": int64(-1),
					},
				},
			},
			wantErr: true,
		},
		{
			name: "max_error",
			fields: fields{
				ds:     &testDriver{Config: IM{}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				nervatype: "item",
				checkValues: IM{
					"values": IM{
						"discount": int64(10000),
					},
				},
			},
			wantErr: true,
		},
		{
			name: "curr_error",
			fields: fields{
				ds:     &testDriver{Config: IM{}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				nervatype: "trans",
				checkValues: IM{
					"values": IM{
						"curr": int64(122),
					},
				},
			},
			wantErr: true,
		},
		{
			name: "groups_type",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": int64(1), "groupname": "transtype", "groupvalue": "invoice"},
							{"id": int64(2), "groupname": "curr", "groupvalue": "EUR"},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				nervatype: "trans",
				checkValues: IM{
					"values": IM{
						"transtype": int64(1),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "groups_type_limit",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": int64(1), "groupname": "nervatype", "groupvalue": "employee"},
							{"id": int64(2), "groupname": "nervatype", "groupvalue": "customer"},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				nervatype: "address",
				checkValues: IM{
					"values": IM{
						"nervatype": 2,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Query_error",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return nil, errors.New("error")
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				nervatype: "address",
				checkValues: IM{
					"values": IM{
						"nervatype": 2,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "groups_nil",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				nervatype: "log",
				checkValues: IM{
					"values": IM{
						"nervatype": nil,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "groups_invalid",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				nervatype: "log",
				checkValues: IM{
					"values": IM{
						"nervatype": "nervatype",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "groups_invalid",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				nervatype: "log",
				checkValues: IM{
					"values": IM{
						"nervatype": 123,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "groups_invalid",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": int64(1), "groupname": "nervatype", "groupvalue": "employee"},
							{"id": int64(2), "groupname": "transtype", "groupvalue": "invoice"},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				nervatype: "address",
				checkValues: IM{
					"values": IM{
						"nervatype": 2,
					},
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
			_, err := nstore.updateValidate(tt.args.nervatype, tt.args.checkValues, tt.args.trans)
			if (err != nil) != tt.wantErr {
				t.Errorf("NervaStore.updateValidate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNervaStore_UpdateData(t *testing.T) {
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
			name: "id_ok",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
					"QueryKey": func(options IM) ([]IM, error) {
						if options["qkey"] == "update_deffields" {
							return []IM{
								{"fieldname": "custfield", "fieldvalue_id": int64(123), "fieldtype": "string"},
							}, nil
						}
						return []IM{}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype":  "customer",
					"insert_row": true,
					"values": IM{
						"id":       float64(1),
						"custname": "custname", "custfield": "fieldvalue",
					},
					"trans": "",
				},
			},
			wantErr: false,
		},
		{
			name: "new_deffields",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
					"QueryKey": func(options IM) ([]IM, error) {
						if options["qkey"] == "update_deffields" {
							return []IM{
								{"fieldname": "custfield_string", "fieldvalue_id": int64(1), "fieldtype": "string", "addnew": 1, "visible": 1},
								{"fieldname": "custfield_bool", "fieldvalue_id": int64(2), "fieldtype": "bool", "addnew": 1, "visible": 1},
								{"fieldname": "custfield_integer", "fieldvalue_id": int64(3), "fieldtype": "integer", "addnew": 1, "visible": 1},
								{"fieldname": "fieldtype_string", "fieldvalue_id": int64(4)},
								{"fieldname": "nervatype_id", "fieldvalue_id": int64(5)},
							}, nil
						}
						return []IM{}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype":    "customer",
					"insert_row":   true,
					"insert_field": true,
					"values": IM{
						"custname": "custname", "custfield_new": "newvalue~1", "custfield_new~0": "newvalue",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "new_fieldvalue",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
					"QueryKey": func(options IM) ([]IM, error) {
						if options["qkey"] == "update_deffields" {
							return []IM{
								{"fieldname": "fieldname", "fieldtype": "string"},
							}, nil
						}
						return []IM{}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype":    "fieldvalue",
					"insert_row":   true,
					"insert_field": true,
					"values": IM{
						"fieldname": "fieldname",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "missing_nervatype",
			fields: fields{
				ds:     &testDriver{Config: IM{}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{},
			},
			wantErr: true,
		},
		{
			name: "invalid_nervatype",
			fields: fields{
				ds:     &testDriver{Config: IM{}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "kalevala",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid_value",
			fields: fields{
				ds:     &testDriver{Config: IM{}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "fieldvalue",
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
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype":  "customer",
					"insert_row": true,
					"values": IM{
						"id":       float64(1),
						"custname": "custname", "custfield": "fieldvalue",
					},
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
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype":  "customer",
					"insert_row": true,
					"values": IM{
						"id":       float64(1),
						"custname": "custname", "custfield": "fieldvalue",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "invalid_id",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{}, nil
					},
					"QueryKey": func(options IM) ([]IM, error) {
						if options["qkey"] == "update_deffields" {
							return []IM{
								{"fieldname": "custfield", "fieldvalue_id": int64(123), "fieldtype": "string"},
							}, nil
						}
						return []IM{}, nil
					},
					"RollbackTransaction": func() error {
						return errors.New("error")
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype":  "customer",
					"insert_row": true,
					"values": IM{
						"id":       float64(1),
						"custname": "custname", "custfield": "fieldvalue",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "disabled_update",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
					"QueryKey": func(options IM) ([]IM, error) {
						if options["qkey"] == "update_deffields" {
							return []IM{
								{"fieldname": "custfield", "fieldvalue_id": int64(123), "fieldtype": "string"},
							}, nil
						}
						return []IM{}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype":  "customer",
					"update_row": false,
					"values": IM{
						"id":       float64(1),
						"custname": "custname", "custfield": "fieldvalue",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "disabled_insert",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
					"QueryKey": func(options IM) ([]IM, error) {
						if options["qkey"] == "update_deffields" {
							return []IM{
								{"fieldname": "custfield_string", "fieldvalue_id": int64(1), "fieldtype": "string", "addnew": 1, "visible": 1},
								{"fieldname": "custfield_bool", "fieldvalue_id": int64(2), "fieldtype": "bool", "addnew": 1, "visible": 1},
								{"fieldname": "custfield_integer", "fieldvalue_id": int64(3), "fieldtype": "integer", "addnew": 1, "visible": 1},
								{"fieldname": "fieldtype_string", "fieldvalue_id": int64(4)},
								{"fieldname": "nervatype_id", "fieldvalue_id": int64(5)},
							}, nil
						}
						return []IM{}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "customer",
					"values": IM{
						"custname": "custname", "custfield_new": "newvalue~1", "custfield_new~0": "newvalue",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "unknown_fieldname",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
					"QueryKey": func(options IM) ([]IM, error) {
						if options["qkey"] == "update_deffields" {
							return []IM{
								{"fieldname": "custfield", "fieldvalue_id": int64(123), "fieldtype": "string"},
							}, nil
						}
						return []IM{}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "groups",
					"values": IM{
						"id":       float64(1),
						"custname": "custname", "custfield": "fieldvalue",
					},
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
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "customer",
					"values": IM{
						"id":       float64(1),
						"custname": "custname", "custfield": "fieldvalue",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "QueryKey_error",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{}, errors.New("error")
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "customer",
					"values": IM{
						"id":       float64(1),
						"custname": "custname", "custfield": "fieldvalue",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "updateValidate_error",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						if queries[0].From == "groups g" {
							return nil, errors.New("error")
						}
						return []IM{
							{"id": int64(1)},
						}, nil
					},
					"QueryKey": func(options IM) ([]IM, error) {
						if options["qkey"] == "update_deffields" {
							return []IM{
								{"fieldname": "custfield", "fieldvalue_id": int64(123), "fieldtype": "string"},
							}, nil
						}
						return []IM{}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "customer",
					"values": IM{
						"id":       float64(1),
						"custname": "custname", "custfield": "fieldvalue",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "fieldvalue_missing_fieldname",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
					"QueryKey": func(options IM) ([]IM, error) {
						if options["qkey"] == "update_deffields" {
							return []IM{
								{"fieldname": "fieldname", "fieldtype": "string"},
							}, nil
						}
						return []IM{}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype":    "fieldvalue",
					"insert_row":   true,
					"insert_field": true,
					"values": IM{
						"value": "value",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "validFieldValue_QueryKey_error",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{}, errors.New("error")
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype":    "customer",
					"insert_row":   true,
					"insert_field": true,
					"values": IM{
						"custname": "custname", "custfield_new": "newvalue~1", "custfield_new~0": "newvalue",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "fieldvalue_checkFieldvalue_error",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{}, errors.New("error")
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype":    "fieldvalue",
					"insert_row":   true,
					"insert_field": true,
					"values": IM{
						"fieldname": "fieldname",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Update_error",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
					"QueryKey": func(options IM) ([]IM, error) {
						if options["qkey"] == "update_deffields" {
							return []IM{
								{"fieldname": "custfield", "fieldvalue_id": int64(123), "fieldtype": "string"},
							}, nil
						}
						return []IM{}, nil
					},
					"Update": func(Update) (int64, error) {
						return 0, errors.New("error")
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype":  "customer",
					"insert_row": true,
					"values": IM{
						"id":       float64(1),
						"custname": "custname", "custfield": "fieldvalue",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "missing_insert_field",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
					"QueryKey": func(options IM) ([]IM, error) {
						if options["qkey"] == "update_deffields" {
							return []IM{
								{"fieldname": "custfield_string", "fieldvalue_id": int64(1), "fieldtype": "string", "addnew": 1, "visible": 1},
								{"fieldname": "custfield_bool", "fieldvalue_id": int64(2), "fieldtype": "bool", "addnew": 1, "visible": 1},
								{"fieldname": "custfield_integer", "fieldvalue_id": int64(3), "fieldtype": "integer", "addnew": 1, "visible": 1},
								{"fieldname": "fieldtype_string", "fieldvalue_id": int64(4)},
								{"fieldname": "nervatype_id", "fieldvalue_id": int64(5)},
							}, nil
						}
						return []IM{}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype":  "customer",
					"insert_row": true,
					"values": IM{
						"custname": "custname", "custfield_new": "newvalue~1", "custfield_new~0": "newvalue",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "deffield_Update_error",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
					"QueryKey": func(options IM) ([]IM, error) {
						if options["qkey"] == "update_deffields" {
							return []IM{
								{"fieldname": "custfield_string", "fieldvalue_id": int64(1), "fieldtype": "string", "addnew": 1, "visible": 1},
								{"fieldname": "custfield_bool", "fieldvalue_id": int64(2), "fieldtype": "bool", "addnew": 1, "visible": 1},
								{"fieldname": "custfield_integer", "fieldvalue_id": int64(3), "fieldtype": "integer", "addnew": 1, "visible": 1},
								{"fieldname": "fieldtype_string", "fieldvalue_id": int64(4)},
								{"fieldname": "nervatype_id", "fieldvalue_id": int64(5)},
							}, nil
						}
						return []IM{}, nil
					},
					"Update": func(options Update) (int64, error) {
						if options.Model == "deffield" {
							return 0, errors.New("error")
						}
						return 0, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype":    "customer",
					"insert_row":   true,
					"insert_field": true,
					"values": IM{
						"custname": "custname", "custfield_new": "newvalue~1", "custfield_new~0": "newvalue",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "checkFieldvalue_error",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
					"QueryKey": func(options IM) ([]IM, error) {
						if options["qkey"] == "update_deffields" && options["fieldname"] == "custfield" {
							return nil, errors.New("error")
						}
						if options["qkey"] == "update_deffields" {
							return []IM{
								{"fieldname": "custfield", "fieldvalue_id": int64(123), "fieldtype": "string"},
							}, nil
						}
						return []IM{}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype":  "customer",
					"insert_row": true,
					"values": IM{
						"id":       float64(1),
						"custname": "custname", "custfield": "fieldvalue",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "fieldvalue_Update_error",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
					"QueryKey": func(options IM) ([]IM, error) {
						if options["qkey"] == "update_deffields" {
							return []IM{
								{"fieldname": "custfield", "fieldvalue_id": int64(123), "fieldtype": "string"},
							}, nil
						}
						return []IM{}, nil
					},
					"Update": func(options Update) (int64, error) {
						if options.Model == "fieldvalue" {
							return 0, errors.New("error")
						}
						return 0, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype":  "customer",
					"insert_row": true,
					"values": IM{
						"id":       float64(1),
						"custname": "custname", "custfield": "fieldvalue",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "insertLog_error",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						if queries[0].From == "groups" {
							return nil, errors.New("error")
						}
						return []IM{
							{"id": int64(1)},
						}, nil
					},
					"QueryKey": func(options IM) ([]IM, error) {
						if options["qkey"] == "update_deffields" {
							return []IM{
								{"fieldname": "custfield", "fieldvalue_id": int64(123), "fieldtype": "string"},
							}, nil
						}
						return []IM{}, nil
					},
				}},
				models: DataModel()["model"].(IM),
				User:   &User{Id: int64(1)},
			},
			args: args{
				options: IM{
					"nervatype":  "customer",
					"insert_row": true,
					"values": IM{
						"id":       float64(1),
						"custname": "custname", "custfield": "fieldvalue",
					},
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
			_, err := nstore.UpdateData(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("NervaStore.UpdateData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNervaStore_GetInfofromRefnumber(t *testing.T) {
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
			name: "tool",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "tool",
					"refnumber": "ABC-123",
				},
			},
			wantErr: false,
		},
		{
			name: "currency",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "currency",
					"refnumber": "EUR",
				},
			},
			wantErr: false,
		},
		{
			name: "address",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "address",
					"refnumber": "customer/DMCUST/00001~1",
				},
			},
			wantErr: false,
		},
		{
			name: "groups",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "groups",
					"refnumber": "barcodetype~QR",
				},
			},
			wantErr: false,
		},
		{
			name: "fieldvalue",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1), "ref_nervatype": "customer", "custtype": "company"},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "fieldvalue",
					"refnumber": "DMCUST/00001~~sample_customer_date~1",
				},
			},
			wantErr: false,
		},
		{
			name: "setting",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "fieldvalue",
					"refnumber": "default_unit",
				},
			},
			wantErr: false,
		},
		{
			name: "item",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1), "transtype": "invoice", "direction": "out",
								"digit": 2, "qty": 1, "discount": 0, "tax_id": 1, "rate": 1},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "item",
					"refnumber": "DMINV/00001~1",
				},
			},
			wantErr: false,
		},
		{
			name: "movement",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1), "transtype": "delivery", "direction": "out", "movetype": "inventory"},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "movement",
					"refnumber": "DMCORR/00001~1",
				},
			},
			wantErr: false,
		},
		{
			name: "payment",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1), "transtype": "bank", "direction": "transfer"},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "payment",
					"refnumber": "DMPMT/00001~1",
				},
			},
			wantErr: false,
		},
		{
			name: "price_5",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "price",
					"refnumber": "DMPROD/00001~price~2020-04-05~EUR~0",
				},
			},
			wantErr: false,
		},
		{
			name: "price_4",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "price",
					"refnumber": "DMPROD/00001~2020-04-05~EUR~0",
				},
			},
			wantErr: false,
		},
		{
			name: "link",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1), "transtype": "invoice", "direction": "out",
								"digit": 2, "qty": 1, "discount": 0, "tax_id": 1, "rate": 1,
								"movetype": "inventory"},
							{"id": int64(1), "transtype": "invoice", "direction": "out",
								"digit": 2, "qty": 1, "discount": 0, "tax_id": 1, "rate": 1,
								"movetype": "inventory"},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "link",
					"refnumber": "movement~DMDEL/00001~2~~item~DMORD/00001~2",
				},
			},
			wantErr: false,
		},
		{
			name: "rate",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "rate",
					"refnumber": "rate~2020-04-05~EUR~bank",
				},
			},
			wantErr: false,
		},
		{
			name: "log",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "log",
					"refnumber": "admin~2019-09-03T17:46:00",
				},
			},
			wantErr: false,
		},
		{
			name: "audit",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "ui_audit",
					"refnumber": "user~trans~invoice",
				},
			},
			wantErr: false,
		},
		{
			name: "menufields",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "ui_menufields",
					"refnumber": "nextNumber~step",
				},
			},
			wantErr: false,
		},
		{
			name: "customer",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1),
								"zipcode": "zipcode", "city": "city", "street": "street",
								"custtype": "own", "custname": "custname", "taxnumber": "taxnumber", "terms": 0},
							{"id": int64(1),
								"zipcode": "zipcode", "city": "city", "street": "street",
								"custtype": "company", "custname": "custname", "taxnumber": "taxnumber", "terms": 0},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype":  "customer",
					"refnumber":  "DMCUST/00001",
					"extra_info": true,
				},
			},
			wantErr: false,
		},
		{
			name: "product",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1), "description": "description", "unit": "unit", "tax_id": 1, "rate": 1},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "product",
					"refnumber": "DMPROD/00001",
				},
			},
			wantErr: false,
		},
		{
			name: "place",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1), "placetype": "bank"},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "place",
					"refnumber": "bank",
				},
			},
			wantErr: false,
		},
		{
			name: "event",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1), "ref_nervatype": "customer"},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "event",
					"refnumber": "DMEVT/00001",
				},
			},
			wantErr: false,
		},
		{
			name: "tax",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1), "rate": 1},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "tax",
					"refnumber": "15%",
				},
			},
			wantErr: false,
		},
		{
			name: "trans",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1), "transtype": "invoice", "direction": "out", "digit": 2},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "trans",
					"refnumber": "DMINV/00001",
				},
			},
			wantErr: false,
		},
		{
			name: "missing_nervatype",
			fields: fields{
				ds:     &testDriver{Config: IM{}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"refnumber": "DMINV/00001",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid_nervatype",
			fields: fields{
				ds:     &testDriver{Config: IM{}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "kalevala",
					"refnumber": "DMINV/00001",
				},
			},
			wantErr: true,
		},
		{
			name: "missing_fieldname",
			fields: fields{
				ds:     &testDriver{Config: IM{}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "customer",
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
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "trans",
					"refnumber": "DMINV/00001",
				},
			},
			wantErr: true,
		},
		{
			name: "address_invalid_refnumber_1",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "address",
					"refnumber": "customer/DMCUST/00001~",
				},
			},
			wantErr: true,
		},
		{
			name: "address_invalid_refnumber_2",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "address",
					"refnumber": "customer/DMCUST/00001~0",
				},
			},
			wantErr: true,
		},
		{
			name: "address_invalid_refnumber_3",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "address",
					"refnumber": "groups/DMCUST/00001~1",
				},
			},
			wantErr: true,
		},
		{
			name: "address_invalid_refnumber_4",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "address",
					"refnumber": "customer",
				},
			},
			wantErr: true,
		},
		{
			name: "groups_invalid_refnumber",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "groups",
					"refnumber": "barcodetype",
				},
			},
			wantErr: true,
		},
		{
			name: "fieldvalue_invalid_refnumber_1",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1), "ref_nervatype": "customer", "custtype": "company"},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "fieldvalue",
					"refnumber": "DMCUST/00001~~sample_customer_date~",
				},
			},
			wantErr: true,
		},
		{
			name: "fieldvalue_invalid_refnumber_2",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1), "ref_nervatype": "customer", "custtype": "company"},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "fieldvalue",
					"refnumber": "DMCUST/00001~~sample_customer_date~0",
				},
			},
			wantErr: true,
		},
		{
			name: "fieldvalue_invalid_refnumber_3",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						if options["nervatype"] == "deffield" {
							return nil, errors.New("error")
						}
						return []IM{
							{"id": int64(1), "ref_nervatype": "customer", "custtype": "company"},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "fieldvalue",
					"refnumber": "DMCUST/00001~~sample_customer_date~1",
				},
			},
			wantErr: true,
		},
		{
			name: "fieldvalue_invalid_refnumber_4",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						if options["nervatype"] == "customer" {
							return nil, errors.New("error")
						}
						return []IM{
							{"id": int64(1), "ref_nervatype": "customer", "custtype": "company"},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "fieldvalue",
					"refnumber": "DMCUST/00001~~sample_customer_date~1",
				},
			},
			wantErr: true,
		},
		{
			name: "item_invalid_refnumber_1",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1), "transtype": "invoice", "direction": "out",
								"digit": 2, "qty": 1, "discount": 0, "tax_id": 1, "rate": 1},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "item",
					"refnumber": "DMINV/00001~",
				},
			},
			wantErr: true,
		},
		{
			name: "item_invalid_refnumber_2",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1), "transtype": "invoice", "direction": "out",
								"digit": 2, "qty": 1, "discount": 0, "tax_id": 1, "rate": 1},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "item",
					"refnumber": "DMINV/00001~0",
				},
			},
			wantErr: true,
		},
		{
			name: "price_invalid_refnumber_1",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "price",
					"refnumber": "DMPROD/00001~price~2020-04-05~EUR~",
				},
			},
			wantErr: true,
		},
		{
			name: "price_invalid_refnumber_2",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "price",
					"refnumber": "DMPROD/00001~price",
				},
			},
			wantErr: true,
		},
		{
			name: "link_invalid_refnumber_1",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1), "transtype": "invoice", "direction": "out",
								"digit": 2, "qty": 1, "discount": 0, "tax_id": 1, "rate": 1,
								"movetype": "inventory"},
							{"id": int64(1), "transtype": "invoice", "direction": "out",
								"digit": 2, "qty": 1, "discount": 0, "tax_id": 1, "rate": 1,
								"movetype": "inventory"},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "link",
					"refnumber": "movement~DMDEL/00001~2",
				},
			},
			wantErr: true,
		},
		{
			name: "link_invalid_refnumber_2",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						if options["nervatype"] == "movement" {
							return nil, errors.New("error")
						}
						return []IM{
							{"id": int64(1), "transtype": "invoice", "direction": "out",
								"digit": 2, "qty": 1, "discount": 0, "tax_id": 1, "rate": 1,
								"movetype": "inventory"},
							{"id": int64(1), "transtype": "invoice", "direction": "out",
								"digit": 2, "qty": 1, "discount": 0, "tax_id": 1, "rate": 1,
								"movetype": "inventory"},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "link",
					"refnumber": "movement~DMDEL/00001~2~~item~DMORD/00001~2",
				},
			},
			wantErr: true,
		},
		{
			name: "link_invalid_refnumber_3",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						if options["nervatype"] == "item" {
							return nil, errors.New("error")
						}
						return []IM{
							{"id": int64(1), "transtype": "invoice", "direction": "out",
								"digit": 2, "qty": 1, "discount": 0, "tax_id": 1, "rate": 1,
								"movetype": "inventory"},
							{"id": int64(1), "transtype": "invoice", "direction": "out",
								"digit": 2, "qty": 1, "discount": 0, "tax_id": 1, "rate": 1,
								"movetype": "inventory"},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "link",
					"refnumber": "movement~DMDEL/00001~2~~item~DMORD/00001~2",
				},
			},
			wantErr: true,
		},
		{
			name: "rate_invalid_refnumber",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "rate",
					"refnumber": "rate~2020-04-05",
				},
			},
			wantErr: true,
		},
		{
			name: "log_invalid_refnumber",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "log",
					"refnumber": "admin",
				},
			},
			wantErr: true,
		},
		{
			name: "audit_invalid_refnumber_1",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "ui_audit",
					"refnumber": "user~groups~invoice",
				},
			},
			wantErr: true,
		},
		{
			name: "audit_invalid_refnumber_2",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "ui_audit",
					"refnumber": "user",
				},
			},
			wantErr: true,
		},
		{
			name: "menufields_invalid_refnumber",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "ui_menufields",
					"refnumber": "nextNumber",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid_refnumber_1",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "ui_userconfig",
					"refnumber": "nextNumber",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid_refnumber_2",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "tool",
					"refnumber": "ABC-123",
				},
			},
			wantErr: true,
		},
		{
			name: "customer_invalid_refnumber",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1),
								"zipcode": "zipcode", "city": "city", "street": "street",
								"custtype": "own", "custname": "custname", "taxnumber": "taxnumber", "terms": 0},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype":  "customer",
					"refnumber":  "DMCUST/00001",
					"extra_info": true,
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
			_, err := nstore.GetInfofromRefnumber(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("NervaStore.GetInfofromRefnumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNervaStore_DeleteData(t *testing.T) {
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
			name: "logical_delete_ok",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1), "custtype": "company", "custname": "custname", "count": int64(0)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "customer",
					"refnumber": "DMCUST/00001",
				},
			},
			wantErr: false,
		},
		{
			name: "delete_ok",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1), "custtype": "company", "custname": "custname", "count": int64(0)},
						}, nil
					},
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{{"value": true}}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "customer",
					"refnumber": "DMCUST/00001",
					"trans":     "trans",
				},
			},
			wantErr: false,
		},
		{
			name: "missing_nervatype",
			fields: fields{
				ds:     &testDriver{Config: IM{}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"refnumber": "DMCUST/00001",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid_nervatype",
			fields: fields{
				ds:     &testDriver{Config: IM{}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "kalevala",
					"refnumber": "DMCUST/00001",
				},
			},
			wantErr: true,
		},
		{
			name: "missing_fieldname",
			fields: fields{
				ds:     &testDriver{Config: IM{}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "customer",
				},
			},
			wantErr: true,
		},
		{
			name: "missing_id",
			fields: fields{
				ds:     &testDriver{Config: IM{}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "customer",
					"refnumber": "DMCUST/00001",
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
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "customer",
					"ref_id":    int64(1),
				},
			},
			wantErr: true,
		},
		{
			name: "QueryKey_error",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return nil, errors.New("error")
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "customer",
					"ref_id":    int64(1),
				},
			},
			wantErr: true,
		},
		{
			name: "integrity_error_1",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"count": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "customer",
					"ref_id":    int64(1),
				},
			},
			wantErr: true,
		},
		{
			name: "numberdef_integrity_error",
			fields: fields{
				ds:     &testDriver{Config: IM{}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "numberdef",
					"ref_id":    int64(1),
				},
			},
			wantErr: true,
		},
		{
			name: "QueryKey_error",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						if options["qkey"] == "delete_deffields" {
							return nil, errors.New("error")
						}
						return []IM{
							{"id": int64(1), "custtype": "company", "custname": "custname", "count": int64(0)},
						}, nil
					},
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{{"value": true}}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "customer",
					"refnumber": "DMCUST/00001",
				},
			},
			wantErr: true,
		},
		{
			name: "BeginTransaction_error",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1), "custtype": "company", "custname": "custname", "count": int64(0)},
						}, nil
					},
					"BeginTransaction": func() (interface{}, error) {
						return nil, errors.New("error")
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "customer",
					"refnumber": "DMCUST/00001",
				},
			},
			wantErr: true,
		},
		{
			name: "Update_error",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1), "custtype": "company", "custname": "custname", "count": int64(0)},
						}, nil
					},
					"Update": func(data Update) (int64, error) {
						return 0, errors.New("error")
					},
					"RollbackTransaction": func() error {
						return errors.New("error")
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "customer",
					"refnumber": "DMCUST/00001",
				},
			},
			wantErr: true,
		},
		{
			name: "Query_error",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1), "custtype": "company", "custname": "custname", "count": int64(0)},
						}, nil
					},
					"Query": func(queries []Query) ([]IM, error) {
						return nil, errors.New("error")
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "customer",
					"refnumber": "DMCUST/00001",
				},
			},
			wantErr: true,
		},
		{
			name: "Update_error_2",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1), "custtype": "company", "custname": "custname", "count": int64(0)},
						}, nil
					},
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{{"value": true}}, nil
					},
					"Update": func(data Update) (int64, error) {
						if data.Model == "fieldvalue" {
							return 0, errors.New("error")
						}
						return 0, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "customer",
					"refnumber": "DMCUST/00001",
				},
			},
			wantErr: true,
		},
		{
			name: "insertLog_error",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"id": int64(1), "custtype": "company", "custname": "custname", "count": int64(0)},
						}, nil
					},
					"Query": func(queries []Query) ([]IM, error) {
						if queries[0].From == "groups" {
							return nil, errors.New("error")
						}
						return []IM{{"value": true}}, nil
					},
				}},
				models: DataModel()["model"].(IM),
				User:   &User{Id: int64(1)},
			},
			args: args{
				options: IM{
					"nervatype": "customer",
					"refnumber": "DMCUST/00001",
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
			if err := nstore.DeleteData(tt.args.options); (err != nil) != tt.wantErr {
				t.Errorf("NervaStore.DeleteData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNervaStore_GetRefnumber(t *testing.T) {
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
			name: "address",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"head_nervatype": "customer", "city": "city", "custnumber": "custnumber",
								"ref_id": int64(1), "nervatype": int64(1), "count": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "address", "ref_id": 6, "retfield": "city",
					"use_deleted": false,
				},
			},
			wantErr: false,
		},
		{
			name: "fieldvalue",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"head_nervatype": "customer", "value": "value", "custnumber": "custnumber", "fieldname": "fieldname",
								"ref_id": int64(1), "nervatype": int64(1), "count": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "fieldvalue", "ref_id": 6, "retfield": "fieldname",
					"use_deleted": false,
				},
			},
			wantErr: false,
		},
		{
			name: "setting",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"head_nervatype": "setting", "value": "value", "custnumber": "custnumber", "fieldname": "fieldname",
								"ref_id": nil, "nervatype": int64(1), "count": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "setting", "ref_id": 6, "retfield": "value",
					"use_deleted": false,
				},
			},
			wantErr: false,
		},
		{
			name: "groups",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"groupname": "groupname", "groupvalue": "groupvalue"},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "groups", "ref_id": 6, "retfield": "",
					"use_deleted": false,
				},
			},
			wantErr: false,
		},
		{
			name: "item",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"transnumber": "transnumber", "trans_id": int64(1), "count": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "item", "ref_id": 6, "retfield": "",
					"use_deleted": false,
				},
			},
			wantErr: false,
		},
		{
			name: "price",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"discount": int64(0), "qty": float64(1), "partnumber": "partnumber", "validfrom": "2021-01-01", "curr": "curr"},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "price", "ref_id": 6, "retfield": "",
					"use_deleted": false,
				},
			},
			wantErr: false,
		},
		{
			name: "link",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"custnumber": "custnumber", "nervatype1": "customer", "ref_id_1": int64(1), "nervatype2": "customer", "ref_id_2": int64(2)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "link", "ref_id": 6, "retfield": "",
					"use_deleted": false,
				},
			},
			wantErr: false,
		},
		{
			name: "rate",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"ratedate": "2021-12-01", "rate_type": "rate", "planumber": "planumber"},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "rate", "ref_id": 6, "retfield": "",
					"use_deleted": false,
				},
			},
			wantErr: false,
		},
		{
			name: "log",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"crdate": "2021-12-01", "empnumber": "empnumber"},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "log", "ref_id": 6, "retfield": "",
					"use_deleted": false,
				},
			},
			wantErr: false,
		},
		{
			name: "missing_nervatype",
			fields: fields{
				ds:     &testDriver{Config: IM{}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"ref_id": 6, "retfield": "",
					"use_deleted": false,
				},
			},
			wantErr: true,
		},
		{
			name: "invalid_nervatype",
			fields: fields{
				ds:     &testDriver{Config: IM{}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "nervatype",
					"ref_id":    6, "retfield": "",
					"use_deleted": false,
				},
			},
			wantErr: true,
		},
		{
			name: "missing_ref_id",
			fields: fields{
				ds:     &testDriver{Config: IM{}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "customer",
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
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "address", "ref_id": 6, "retfield": "city",
					"use_deleted": false,
				},
			},
			wantErr: true,
		},
		{
			name: "QueryKey_error",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return nil, errors.New("error")
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "address", "ref_id": 6, "retfield": "city",
					"use_deleted": false,
				},
			},
			wantErr: true,
		},
		{
			name: "invalid_refnumber",
			fields: fields{
				ds:     &testDriver{Config: IM{}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "address", "ref_id": 6, "retfield": "city",
					"use_deleted": false,
				},
			},
			wantErr: true,
		},
		{
			name: "address_QueryKey_error",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						if options["refId"] == "1" {
							return nil, errors.New("error")
						}
						return []IM{
							{"head_nervatype": "customer", "city": "city", "custnumber": "custnumber",
								"ref_id": int64(1), "nervatype": int64(1), "count": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "address", "ref_id": 6, "retfield": "city",
					"use_deleted": false,
				},
			},
			wantErr: true,
		},
		{
			name: "address_invalid_refnumber",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						if options["refId"] == "1" {
							return []IM{}, nil
						}
						return []IM{
							{"head_nervatype": "customer", "city": "city", "custnumber": "custnumber",
								"ref_id": int64(1), "nervatype": int64(1), "count": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "address", "ref_id": 6, "retfield": "city",
					"use_deleted": false,
				},
			},
			wantErr: true,
		},
		{
			name: "address_GetRefnumber_error",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						if options["nervatype"] == "customer" {
							return nil, errors.New("error")
						}
						return []IM{
							{"head_nervatype": "customer", "city": "city", "custnumber": "custnumber",
								"ref_id": int64(1), "nervatype": int64(1), "count": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "address", "ref_id": 6, "retfield": "city",
					"use_deleted": false,
				},
			},
			wantErr: true,
		},
		{
			name: "fieldvalue_QueryKey_error",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						if options["refId"] == "1" {
							return nil, errors.New("error")
						}
						return []IM{
							{"head_nervatype": "customer", "value": "value", "custnumber": "custnumber", "fieldname": "fieldname",
								"ref_id": int64(1), "nervatype": int64(1), "count": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "fieldvalue", "ref_id": 6, "retfield": "fieldname",
					"use_deleted": false,
				},
			},
			wantErr: true,
		},
		{
			name: "fieldvalue_invalid_refnumber",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						if options["refId"] == "1" {
							return []IM{}, nil
						}
						return []IM{
							{"head_nervatype": "customer", "value": "value", "custnumber": "custnumber", "fieldname": "fieldname",
								"ref_id": int64(1), "nervatype": int64(1), "count": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "fieldvalue", "ref_id": 6, "retfield": "fieldname",
					"use_deleted": false,
				},
			},
			wantErr: true,
		},
		{
			name: "fieldvalue_GetRefnumber_error",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						if options["nervatype"] == "customer" {
							return nil, errors.New("error")
						}
						return []IM{
							{"head_nervatype": "customer", "value": "value", "custnumber": "custnumber", "fieldname": "fieldname",
								"ref_id": int64(1), "nervatype": int64(1), "count": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "fieldvalue", "ref_id": 6, "retfield": "fieldname",
					"use_deleted": false,
				},
			},
			wantErr: true,
		},
		{
			name: "item_QueryKey_error",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						if options["refId"] == "1" {
							return nil, errors.New("error")
						}
						return []IM{
							{"transnumber": "transnumber", "trans_id": int64(1), "count": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "item", "ref_id": 6, "retfield": "",
					"use_deleted": false,
				},
			},
			wantErr: true,
		},
		{
			name: "item_invalid_refnumber",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						if options["refId"] == "1" {
							return []IM{}, nil
						}
						return []IM{
							{"transnumber": "transnumber", "trans_id": int64(1), "count": int64(1)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "item", "ref_id": 6, "retfield": "",
					"use_deleted": false,
				},
			},
			wantErr: true,
		},
		{
			name: "link_GetRefnumber_error_1",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						if options["nervatype"] == "customer" {
							return nil, errors.New("error")
						}
						return []IM{
							{"custnumber": "custnumber", "nervatype1": "customer", "ref_id_1": int64(1), "nervatype2": "employee", "ref_id_2": int64(2)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "link", "ref_id": 6, "retfield": "",
					"use_deleted": false,
				},
			},
			wantErr: true,
		},
		{
			name: "link_GetRefnumber_error_2",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						if options["nervatype"] == "employee" {
							return nil, errors.New("error")
						}
						return []IM{
							{"custnumber": "custnumber", "nervatype1": "customer", "ref_id_1": int64(1), "nervatype2": "employee", "ref_id_2": int64(2)},
						}, nil
					},
				}},
				models: DataModel()["model"].(IM),
			},
			args: args{
				options: IM{
					"nervatype": "link", "ref_id": 6, "retfield": "",
					"use_deleted": false,
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
			_, err := nstore.GetRefnumber(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("NervaStore.GetRefnumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNervaStore_GetDataAudit(t *testing.T) {
	type fields struct {
		ds       DataDriver
		User     *User
		Customer IM
		models   IM
		config   IM
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "audit_transfilter",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"transfilter": "transfilter"},
						}, nil
					},
				}},
				User: &User{Id: int64(1)},
			},
			wantErr: false,
		},
		{
			name: "audit_all",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{}, nil
					},
				}},
				User: &User{Id: int64(1)},
			},
			wantErr: false,
		},
		{
			name: "invalid_login",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{}, nil
					},
				}},
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
				User: &User{Id: int64(1)},
			},
			wantErr: true,
		},
		{
			name: "QueryKey_error",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return nil, errors.New("error")
					},
				}},
				User: &User{Id: int64(1)},
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
			_, err := nstore.GetDataAudit()
			if (err != nil) != tt.wantErr {
				t.Errorf("NervaStore.GetDataAudit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNervaStore_GetObjectAudit(t *testing.T) {
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
			name: "audit_list_ok",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"inputfilter": "disabled"},
						}, nil
					},
				}},
				User: &User{Id: int64(1)},
			},
			args: args{
				options: IM{
					"nervatype":    []string{"customer", "product"},
					"nervatype_id": []int64{1, 2, 3},
					"transtype":    []string{"invoice", "order"},
					"transtype_id": []int64{1, 2, 3},
				},
			},
			wantErr: false,
		},
		{
			name: "audit_ok",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"inputfilter": "readonly"},
						}, nil
					},
				}},
				User: &User{Id: int64(1)},
			},
			args: args{
				options: IM{
					"nervatype":    "customer",
					"nervatype_id": int64(1),
					"transtype":    "invoice",
					"transtype_id": int64(1),
				},
			},
			wantErr: false,
		},
		{
			name:   "invalid_login",
			fields: fields{},
			args: args{
				options: IM{
					"nervatype":    "customer",
					"nervatype_id": int64(1),
					"transtype":    "invoice",
					"transtype_id": int64(1),
				},
			},
			wantErr: true,
		},
		{
			name: "missing_nervatype",
			fields: fields{
				User: &User{Id: int64(1)},
			},
			args: args{
				options: IM{},
			},
			wantErr: true,
		},
		{
			name: "sql_all",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"inputfilter": "readonly"},
						}, nil
					},
				}},
				User: &User{Id: int64(1)},
			},
			args: args{
				options: IM{
					"nervatype": "sql",
				},
			},
			wantErr: false,
		},
		{
			name: "not_connect",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"inputfilter": "readonly"},
						}, nil
					},
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
				User: &User{Id: int64(1)},
			},
			args: args{
				options: IM{
					"nervatype":    "customer",
					"nervatype_id": int64(1),
					"transtype":    "invoice",
					"transtype_id": int64(1),
				},
			},
			wantErr: true,
		},
		{
			name: "QueryKey_error",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return nil, errors.New("error")
					},
				}},
				User: &User{Id: int64(1)},
			},
			args: args{
				options: IM{
					"nervatype":    "customer",
					"nervatype_id": int64(1),
					"transtype":    "invoice",
					"transtype_id": int64(1),
				},
			},
			wantErr: true,
		},
		{
			name: "audit_all_ok",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{}, nil
					},
				}},
				User: &User{Id: int64(1)},
			},
			args: args{
				options: IM{
					"nervatype":    "customer",
					"nervatype_id": int64(1),
					"transtype":    "invoice",
					"transtype_id": int64(1),
				},
			},
			wantErr: false,
		},
		{
			name: "audit_update_ok",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return []IM{
							{"inputfilter": "update"},
						}, nil
					},
				}},
				User: &User{Id: int64(1)},
			},
			args: args{
				options: IM{
					"nervatype":    "customer",
					"nervatype_id": int64(1),
					"transtype":    "invoice",
					"transtype_id": int64(1),
				},
			},
			wantErr: false,
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
			_, err := nstore.GetObjectAudit(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("NervaStore.GetObjectAudit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNervaStore_GetGroups(t *testing.T) {
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
			name: "groups_list_ok",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": int64(1), "groupname": "groupname", "groupvalue": "groupvalue"},
						}, nil
					},
				}},
			},
			args: args{
				options: IM{
					"groupname": []string{"transtype", "usergroup"},
				},
			},
			wantErr: false,
		},
		{
			name: "groups_ok",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": int64(1), "groupname": "groupname", "groupvalue": "groupvalue"},
						}, nil
					},
				}},
			},
			args: args{
				options: IM{
					"groupname": "transtype",
				},
			},
			wantErr: false,
		},
		{
			name: "not_connect",
			fields: fields{
				ds: &testDriver{Config: IM{
					"Query": func(queries []Query) ([]IM, error) {
						return []IM{
							{"id": int64(1), "groupname": "groupname", "groupvalue": "groupvalue"},
						}, nil
					},
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
					"groupname": "transtype",
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
				}},
			},
			args: args{
				options: IM{
					"groupname": "transtype",
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
			_, err := nstore.GetGroups(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("NervaStore.GetGroups() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
