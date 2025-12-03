package api

import (
	"errors"
	"io"
	"log/slog"
	"net/http"
	"path"
	"testing"

	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	st "github.com/nervatura/nervatura/v6/pkg/static"
)

func TestDataStore_ReportInstall(t *testing.T) {
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
	}
	type args struct {
		reportKey string
		reportDir string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantConfigID int64
		wantErr      bool
	}{
		{
			name: "success",
			fields: fields{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{}, nil
						},
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
					},
				},
				ReadFile: func(name string) ([]byte, error) {
					return []byte(`{"id": 1, "name": "test"}`), nil
				},
				ConvertFromByte: func(data []byte, result interface{}) error {
					return cu.ConvertFromByte([]byte(`{"meta": {"report_key": "test", "report_name": "test", "report_type": "test", "file_type": "FILE_CSV"}}`), result)
				},
				ConvertToByte: func(data interface{}) ([]byte, error) {
					return []byte(`{"meta": {"report_key": "test", "report_name": "test", "report_type": "test", "file_type": "FILE_CSV"}}`), nil
				},
			},
			args: args{
				reportKey: "ntr_customer_en",
				reportDir: "",
			},
			wantConfigID: 1,
			wantErr:      false,
		},
		{
			name: "exists template",
			fields: fields{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1}}, nil
						},
					},
				},
				ReadFile: func(name string) ([]byte, error) {
					return []byte(`{"id": 1, "name": "test"}`), nil
				},
				ConvertFromByte: func(data []byte, result interface{}) error {
					return cu.ConvertFromByte([]byte(`{"meta": {"report_key": "test", "report_name": "test", "report_type": "test", "file_type": "FILE_CSV"}}`), result)
				},
			},
			args: args{
				reportKey: "ntr_customer_en",
				reportDir: "",
			},
			wantConfigID: 0,
			wantErr:      true,
		},

		{
			name: "invalid template",
			fields: fields{
				Db: &md.TestDriver{
					Config: cu.IM{},
				},
				ReadFile: func(name string) ([]byte, error) {
					return []byte(`{"id": 1, "name": "test"}`), nil
				},
				ConvertFromByte: func(data []byte, result interface{}) error {
					return cu.ConvertFromByte([]byte(`{"meta": {"report_key": "test"}}`), result)
				},
			},
			args: args{
				reportKey: "ntr_customer_en",
				reportDir: "",
			},
			wantConfigID: 0,
			wantErr:      true,
		},
		{
			name: "convert error",
			fields: fields{
				Db: &md.TestDriver{
					Config: cu.IM{},
				},
				ReadFile: func(name string) ([]byte, error) {
					return []byte(`{"id": 1, "name": "test"}`), nil
				},
				ConvertFromByte: func(data []byte, result interface{}) error {
					return errors.New("test")
				},
			},
			args: args{
				reportKey: "ntr_customer_en",
				reportDir: "",
			},
			wantConfigID: 0,
			wantErr:      true,
		},
		{
			name: "missing template",
			fields: fields{
				Db: &md.TestDriver{
					Config: cu.IM{},
				},
				ReadFile: func(name string) ([]byte, error) {
					return nil, errors.New("test")
				},
				ConvertFromByte: func(data []byte, result interface{}) error {
					return errors.New("test")
				},
			},
			args: args{
				reportKey: "ntr_customer_en",
				reportDir: "path/to/template",
			},
			wantConfigID: 0,
			wantErr:      true,
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
			}
			gotConfigID, err := ds.ReportInstall(tt.args.reportKey, tt.args.reportDir)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataStore.ReportInstall() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotConfigID != tt.wantConfigID {
				t.Errorf("DataStore.ReportInstall() = %v, want %v", gotConfigID, tt.wantConfigID)
			}
		})
	}
}

func TestDataStore_ReportList(t *testing.T) {
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
	}
	type args struct {
		reportDir string
		filter    string
	}
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
				reportDir: "",
				filter:    "",
			},
			wantErr: false,
		},
		{
			name: "report dir list",
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
				reportDir: "../../data",
				filter:    "",
			},
			wantErr: false,
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
			}
			_, err := ds.ReportList(tt.args.reportDir, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataStore.ReportList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDataStore_GetReport(t *testing.T) {
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
	}
	type args struct {
		options cu.IM
	}
	pdf_json, _ := st.Report.ReadFile(path.Join("template", "ntr_customer_en.json"))
	csv_json, _ := st.Report.ReadFile(path.Join("template", "csv_custpos_en.json"))
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success_pdf",
			fields: fields{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "data": cu.IM{"file_type": "FILE_PDF"}}}, nil
						},
						"QuerySQL": func(sqlString string) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "data": cu.IM{"file_type": "FILE_PDF"}}}, nil
						},
					},
				},
				ConvertFromByte: func(data []byte, result interface{}) error {
					return cu.ConvertFromByte(data, result)
				},
			},
			args: args{
				options: cu.IM{
					"report_key":  "ntr_customer_en",
					"orientation": "portrait",
					"size":        "a4",
					"code":        "test",
					"template":    string(pdf_json),
					"filters":     cu.IM{},
				},
			},
			wantErr: false,
		},
		{
			name: "data_output",
			fields: fields{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "data": cu.IM{"file_type": "FILE_PDF"}}}, nil
						},
						"QuerySQL": func(sqlString string) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "data": cu.IM{"file_type": "FILE_PDF"}}}, nil
						},
					},
				},
				ConvertFromByte: func(data []byte, result interface{}) error {
					return cu.ConvertFromByte(data, result)
				},
			},
			args: args{
				options: cu.IM{
					"report_key":  "ntr_customer_en",
					"orientation": "portrait",
					"size":        "a4",
					"code":        "test",
					"output":      "data",
					"template":    string(pdf_json),
					"filters":     cu.IM{},
				},
			},
			wantErr: false,
		},
		{
			name: "csv_output",
			fields: fields{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "data": cu.IM{"file_type": "FILE_CSV"}}}, nil
						},
						"QuerySQL": func(sqlString string) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "data": cu.IM{"file_type": "FILE_CSV"}, "invoice_no": "INV123456N1"}}, nil
						},
					},
				},
				ConvertFromByte: func(data []byte, result interface{}) error {
					return cu.ConvertFromByte(data, result)
				},
			},
			args: args{
				options: cu.IM{
					"report_key": "csv_custpos_en",
					"output":     "base64",
					"template":   string(csv_json),
					"filters": cu.IM{
						"posdate":  "2024-01-01",
						"curr":     "EUR",
						"customer": "Second%",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "data": cu.IM{"file_type": "FILE_PDF"}}}, nil
						},
						"QuerySQL": func(sqlString string) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "data": cu.IM{"file_type": "FILE_PDF"}}}, nil
						},
					},
				},
				ConvertFromByte: func(data []byte, result interface{}) error {
					return errors.New("test")
				},
			},
			args: args{
				options: cu.IM{
					"report_key":  "ntr_customer_en",
					"orientation": "portrait",
					"size":        "a4",
					"code":        "test",
					"template":    string(pdf_json),
					"filters":     cu.IM{},
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
			}
			_, err := ds.GetReport(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataStore.GetReport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDataStore_getReportData(t *testing.T) {
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
	}
	type args struct {
		reportTemplate cu.IM
		filters        cu.IM
		sources        []cu.SM
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
				reportTemplate: cu.IM{
					"data": cu.IM{},
				},
				filters: cu.IM{
					"@error": "value",
				},
				sources: []cu.SM{},
			},
			wantErr: true,
		},
		{
			name: "ds_nodata",
			fields: fields{
				Db: &md.TestDriver{
					Config: cu.IM{
						"QuerySQL": func(sqlString string) ([]cu.IM, error) {
							if sqlString == "ds" {
								return []cu.IM{}, nil
							}
							return []cu.IM{{"id": int64(1)}}, nil
						},
					}},
			},
			args: args{
				reportTemplate: cu.IM{
					"meta": cu.IM{
						"report_name": "repname",
					},
					"data":   cu.IM{},
					"fields": cu.IM{},
				},
				filters: cu.IM{},
				sources: []cu.SM{
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
		{
			name: "nodata",
			fields: fields{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "data": cu.IM{"file_type": "FILE_PDF"}}}, nil
						},
						"QuerySQL": func(sqlString string) ([]cu.IM, error) {
							return nil, errors.New("test")
						},
					},
				},
			},
			args: args{
				reportTemplate: cu.IM{
					"meta": cu.IM{
						"report_name": "repname",
						"report_type": "REPORT",
						"file_type":   "FILE_CSV",
					},
					"data": cu.IM{},
					"fields": cu.IM{
						"date1": cu.IM{
							"fieldtype": "date",
							"wheretype": "where",
						},
						"date2": cu.IM{
							"fieldtype": "date",
							"wheretype": "where",
							"dataset":   "dsdate",
						},
					},
				},
				filters: cu.IM{
					"date1": "2021-12-24",
					"date2": "2021-12-24",
				},
				sources: []cu.SM{
					{
						"dataset": "dsdate",
						"sqlstr":  "select * from table",
					},
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
			}
			_, err := ds.getReportData(tt.args.reportTemplate, tt.args.filters, tt.args.sources)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataStore.getReportData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
