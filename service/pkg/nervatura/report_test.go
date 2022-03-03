package nervatura

import (
	"errors"
	"path"
	"testing"

	ut "github.com/nervatura/nervatura/service/pkg/utils"
)

func TestNervaStore_getReportRefValues(t *testing.T) {
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
			name: "info",
			fields: fields{
				ds: &testDriver{Config: IM{}},
			},
			args: args{
				options: IM{
					"nervatype": "nervatype",
					"refnumber": "refnumber",
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
			_, _, err := nstore.getReportRefValues(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("NervaStore.getReportRefValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNervaStore_getReportDefault(t *testing.T) {
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
			name: "QueryKey_error",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						if options["qkey"] == "default_report" {
							return nil, errors.New("error")
						}
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
			wantErr: true,
		},
		{
			name: "not_exist",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						if options["qkey"] == "default_report" {
							return []IM{}, nil
						}
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
			_, err := nstore.getReportDefault(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("NervaStore.getReportDefault() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNervaStore_getReportHead(t *testing.T) {
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
			name: "report",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						if options["qkey"] == "default_report" {
							return []IM{
								{"id": int64(1), "reportkey": "sample",
									"nervatype": int64(1), "transtype": int64(1), "direction": int64(1),
									"repname": "Sample json", "description": "Sample json", "label": "sample",
									"filetype": int64(1), "report": "", "reptype": "pdf",
								},
							}, nil
						}
						return []IM{}, nil
					},
				}},
			},
			args: args{
				options: IM{
					"reportkey": "sample",
					"nervatype": "trans",
				},
			},
			wantErr: false,
		},
		{
			name: "not_exist",
			fields: fields{
				ds: &testDriver{Config: IM{}},
			},
			args: args{
				options: IM{
					"report_id": int64(123),
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
			_, err := nstore.getReportHead(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("NervaStore.getReportHead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNervaStore_getReportDataWhere(t *testing.T) {
	type fields struct {
		ds       DataDriver
		User     *User
		Customer IM
		models   IM
		config   IM
	}
	type args struct {
		reportTemplate IM
		filters        IM
		sources        []SM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "fields",
			fields: fields{},
			args: args{
				reportTemplate: IM{
					"fields": IM{
						"date1": IM{
							"fieldtype": "date",
							"wheretype": "where",
						},
						"string1": IM{
							"fieldtype": "string",
							"wheretype": "where",
							"sql":       "string1",
						},
						"date2": IM{
							"fieldtype": "date",
							"wheretype": "where",
							"dataset":   "dsdate",
						},
						"string2": IM{
							"fieldtype": "string",
							"wheretype": "in",
						},
						"string3": IM{
							"fieldtype": "string",
							"wheretype": "in",
							"sql":       "string1",
						},
					},
				},
				filters: IM{
					"date1":   "2021-12-24",
					"string1": "value",
					"date2":   "2021-12-24",
					"string2": "value",
					"string3": "value",
				},
				sources: []SM{
					{
						"dataset": "dsdate",
						"sqlstr":  "select * from table",
					},
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
			_, err := nstore.getReportDataWhere(tt.args.reportTemplate, tt.args.filters, tt.args.sources)
			if (err != nil) != tt.wantErr {
				t.Errorf("NervaStore.getReportDataWhere() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNervaStore_getReportData(t *testing.T) {
	type fields struct {
		ds       DataDriver
		User     *User
		Customer IM
		models   IM
		config   IM
	}
	type args struct {
		reportTemplate IM
		filters        IM
		sources        []SM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "getReportDataWhere_error",
			fields: fields{},
			args: args{
				reportTemplate: IM{
					"data": IM{},
				},
				filters: IM{
					"@error": "value",
				},
				sources: []SM{},
			},
			wantErr: true,
		},
		{
			name: "nodata",
			fields: fields{
				ds: &testDriver{Config: IM{}},
			},
			args: args{
				reportTemplate: IM{
					"meta": IM{
						"rapname": "repname",
					},
					"data": IM{},
					"fields": IM{
						"date1": IM{
							"fieldtype": "date",
							"wheretype": "where",
						},
						"date2": IM{
							"fieldtype": "date",
							"wheretype": "where",
							"dataset":   "dsdate",
						},
					},
				},
				filters: IM{
					"date1": "2021-12-24",
					"date2": "2021-12-24",
				},
				sources: []SM{
					{
						"dataset": "dsdate",
						"sqlstr":  "select * from table",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "ds_nodata",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QuerySQL": func(sqlString string) ([]IM, error) {
						if sqlString == "ds" {
							return []IM{}, nil
						}
						return []IM{
							{"id": int64(1)},
						}, nil
					},
				}},
			},
			args: args{
				reportTemplate: IM{
					"meta": IM{
						"rapname": "repname",
					},
					"data":   IM{},
					"fields": IM{},
				},
				filters: IM{},
				sources: []SM{
					{
						"dataset": "dsdate",
						"sqlstr":  "select * from table",
					},
					{
						"dataset": "ds",
						"sqlstr":  "ds",
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
			_, err := nstore.getReportData(tt.args.reportTemplate, tt.args.filters, tt.args.sources)
			if (err != nil) != tt.wantErr {
				t.Errorf("NervaStore.getReportData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNervaStore_getReportCSV(t *testing.T) {
	type fields struct {
		ds       DataDriver
		User     *User
		Customer IM
		models   IM
		config   IM
	}
	type args struct {
		reportTemplate IM
		datarows       IM
		base64Encoding bool
	}
	csv_json, _ := ut.Report.ReadFile(path.Join("static", "templates", "csv_custpos_en.json"))
	reportTemplate := IM{}
	_ = ut.ConvertFromByte([]byte(csv_json), &reportTemplate)
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "base64",
			fields: fields{},
			args: args{
				reportTemplate: reportTemplate,
				datarows:       make(map[string]interface{}),
				base64Encoding: true,
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
			_, err := nstore.getReportCSV(tt.args.reportTemplate, tt.args.datarows, tt.args.base64Encoding)
			if (err != nil) != tt.wantErr {
				t.Errorf("NervaStore.getReportCSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNervaStore_getReportPDF(t *testing.T) {
	type fields struct {
		ds       DataDriver
		User     *User
		Customer IM
		models   IM
		config   IM
	}
	type args struct {
		options      IM
		datarows     IM
		jsonTemplate string
	}
	sample_json, _ := ut.Report.ReadFile(path.Join("static", "templates", "sample.json"))
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "base64",
			fields: fields{},
			args: args{
				options: IM{
					"output": "base64",
				},
				datarows: IM{
					"labels": IM{
						"label": "value",
					},
				},
				jsonTemplate: string(sample_json),
			},
			wantErr: false,
		},
		{
			name:   "xml",
			fields: fields{},
			args: args{
				options: IM{
					"output": "xml",
				},
				datarows:     IM{},
				jsonTemplate: string(sample_json),
			},
			wantErr: false,
		},
		{
			name:   "LoadJSONDefinition_error",
			fields: fields{},
			args: args{
				options: IM{
					"output": "pdf",
				},
				datarows:     IM{},
				jsonTemplate: "json",
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
			_, err := nstore.getReportPDF(tt.args.options, tt.args.datarows, tt.args.jsonTemplate)
			if (err != nil) != tt.wantErr {
				t.Errorf("NervaStore.getReportPDF() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNervaStore_getReport(t *testing.T) {
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
	pdf_json, _ := ut.Report.ReadFile(path.Join("static", "templates", "ntr_invoice_en.json"))
	csv_json, _ := ut.Report.ReadFile(path.Join("static", "templates", "csv_custpos_en.json"))
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "pdf_report",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QuerySQL": func(sqlString string) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
				}},
			},
			args: args{
				options: IM{
					"filters": IM{},
					"report": IM{
						"reportkey": "sample",
						"report":    string(pdf_json),
						"ref_id":    int64(1),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "data_output",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QuerySQL": func(sqlString string) ([]IM, error) {
						return []IM{
							{"id": int64(1)},
						}, nil
					},
				}},
			},
			args: args{
				options: IM{
					"output":  "data",
					"filters": IM{},
					"report": IM{
						"reportkey": "sample",
						"report":    string(pdf_json),
						"ref_id":    int64(1),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "csv_report",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QuerySQL": func(sqlString string) ([]IM, error) {
						return []IM{
							{"id": int64(1), "customer": "customer"},
						}, nil
					},
				}},
			},
			args: args{
				options: IM{
					"report": IM{
						"reptype":   "csv",
						"reportkey": "sample",
						"report":    string(csv_json),
						"ref_id":    int64(1),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "getReportHead_error",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QueryKey": func(options IM) ([]IM, error) {
						return nil, errors.New("error")
					},
				}},
			},
			args: args{
				options: IM{
					"reportkey": "sample",
				},
			},
			wantErr: true,
		},
		{
			name: "missing_reportkey",
			fields: fields{
				ds: &testDriver{Config: IM{}},
			},
			args: args{
				options: IM{},
			},
			wantErr: true,
		},
		{
			name: "getReportData_error",
			fields: fields{
				ds: &testDriver{Config: IM{
					"QuerySQL": func(sqlString string) ([]IM, error) {
						return nil, errors.New("error")
					},
				}},
			},
			args: args{
				options: IM{
					"filters": IM{},
					"report": IM{
						"reportkey": "sample",
						"report":    string(pdf_json),
						"ref_id":    int64(1),
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
			_, err := nstore.getReport(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("NervaStore.getReport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
