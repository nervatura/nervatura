//go:build http || all
// +build http all

package service

import (
	"errors"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	bc "github.com/nervatura/component/component/base"
	cp "github.com/nervatura/nervatura/service/pkg/component"
	nt "github.com/nervatura/nervatura/service/pkg/nervatura"
	ut "github.com/nervatura/nervatura/service/pkg/utils"
)

func TestAdminService_appResponse(t *testing.T) {
	type fields struct {
		Config          map[string]interface{}
		GetNervaStore   func(database string) *nt.NervaStore
		GetParam        func(req *http.Request, name string) string
		GetTokenKeys    func() map[string]map[string]string
		GetTaskSecKey   func() string
		ReadFile        func(name string) ([]byte, error)
		ConvertFromByte func(data []byte, result interface{}) error
		CreateFile      func(name string) (*os.File, error)
		ConvertToWriter func(out io.Writer, data interface{}) error
		Session         nt.SessionService
	}
	type args struct {
		evt bc.ResponseEvent
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "save",
			fields: fields{
				Config: nt.IM{
					"version": testData.version,
				},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
				CreateFile: func(name string) (*os.File, error) {
					return os.NewFile(1234, "file"), nil
				},
				ConvertToWriter: func(out io.Writer, data interface{}) error {
					return nil
				},
			},
			args: args{
				evt: bc.ResponseEvent{
					Trigger: &cp.Admin{
						BaseComponent: bc.BaseComponent{
							Data: nt.IM{
								"locfile": nt.IM{"locales": nt.IM{}},
							},
						},
						Token: testData.adminToken,
					},
					Name: cp.AdminEventLocalesSave,
				},
			},
		},
		{
			name: "save_create_err",
			fields: fields{
				Config: nt.IM{
					"version": testData.version,
				},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
				CreateFile: func(name string) (*os.File, error) {
					return os.NewFile(0, "file"), errors.New("error")
				},
			},
			args: args{
				evt: bc.ResponseEvent{
					Trigger: &cp.Admin{
						BaseComponent: bc.BaseComponent{
							Data: nt.IM{
								"locfile": nt.IM{"locales": nt.IM{}},
							},
						},
						Token: testData.adminToken,
					},
					Name: cp.AdminEventLocalesSave,
				},
			},
		},
		{
			name: "save_convert_err",
			fields: fields{
				Config: nt.IM{
					"version": testData.version,
				},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
				CreateFile: func(name string) (*os.File, error) {
					return nil, nil
				},
				ConvertToWriter: func(out io.Writer, data interface{}) error {
					return errors.New("error")
				},
			},
			args: args{
				evt: bc.ResponseEvent{
					Trigger: &cp.Admin{
						BaseComponent: bc.BaseComponent{
							Data: nt.IM{
								"locfile": nt.IM{"locales": nt.IM{}},
							},
						},
						Token: testData.adminToken,
					},
					Name: cp.AdminEventLocalesSave,
				},
			},
		},
		{
			name: "error",
			fields: fields{
				Config: nt.IM{
					"version": testData.version,
				},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				evt: bc.ResponseEvent{
					Trigger: &cp.Admin{
						BaseComponent: bc.BaseComponent{
							Data: nt.IM{
								"locfile": nt.IM{"locales": nt.IM{}},
							},
						},
						Token: testData.adminToken,
					},
					Name: cp.AdminEventLocalesError,
				},
			},
		},
		{
			name: "create",
			fields: fields{
				Config: nt.IM{
					"version":       testData.version,
					"NT_HASHTABLE":  testData.hashTable,
					"NT_API_KEY":    testData.apiKey,
					"NT_ALIAS_TEST": testData.testDatabase,
				},
			},
			args: args{
				evt: bc.ResponseEvent{
					Trigger: &bc.BaseComponent{
						Data: nt.IM{
							"api_key": testData.apiKey,
							"alias":   "test",
							"demo":    true,
						},
					},
					Name: cp.AdminEventCreate,
				},
			},
		},
		{
			name: "create_error",
			fields: fields{
				Config: nt.IM{
					"version":       testData.version,
					"NT_HASHTABLE":  testData.hashTable,
					"NT_API_KEY":    testData.apiKey,
					"NT_ALIAS_TEST": testData.testDatabase,
				},
			},
			args: args{
				evt: bc.ResponseEvent{
					Trigger: &bc.BaseComponent{
						Data: nt.IM{
							"api_key": testData.apiKey,
							"alias":   "",
							"demo":    "false",
						},
					},
					Name: cp.AdminEventCreate,
				},
			},
		},
		{
			name: "login",
			fields: fields{
				Config: nt.IM{
					"version": testData.version,
				},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE": testData.hashTable,
					})
				},
			},
			args: args{
				evt: bc.ResponseEvent{
					Trigger: &bc.BaseComponent{
						Data: nt.IM{
							"username": "admin",
							"database": testData.testDatabase,
						},
					},
					Name: cp.AdminEventLogin,
				},
			},
		},
		{
			name: "login_err",
			fields: fields{
				Config: nt.IM{
					"version": testData.version,
				},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE": testData.hashTable,
					})
				},
			},
			args: args{
				evt: bc.ResponseEvent{
					Trigger: &bc.BaseComponent{
						Data: nt.IM{
							"username": "user",
							"database": testData.testDatabase,
						},
					},
					Name: cp.AdminEventLogin,
				},
			},
		},
		{
			name: "report_install",
			fields: fields{
				Config: nt.IM{
					"version": testData.version,
				},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				evt: bc.ResponseEvent{
					Trigger: &cp.Admin{
						BaseComponent: bc.BaseComponent{
							Data: nt.IM{
								"database": testData.testDatabase,
							},
						},
						Token: testData.adminToken,
					},
					Name:  cp.AdminEventReportInstall,
					Value: "missing",
				},
			},
		},
		{
			name: "report_delete",
			fields: fields{
				Config: nt.IM{
					"version": testData.version,
				},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				evt: bc.ResponseEvent{
					Trigger: &cp.Admin{
						BaseComponent: bc.BaseComponent{
							Data: nt.IM{
								"database": testData.testDatabase,
							},
						},
						Token: testData.adminToken,
					},
					Name:  cp.AdminEventReportDelete,
					Value: "ntr_cash_out_en",
				},
			},
		},
		{
			name: "password",
			fields: fields{
				Config: nt.IM{
					"version": testData.version,
				},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				evt: bc.ResponseEvent{
					Trigger: &cp.Admin{
						BaseComponent: bc.BaseComponent{
							Data: nt.IM{
								"username": "demo",
								"password": "123",
								"confirm":  "123",
							},
						},
						Token: testData.adminToken,
					},
					Name: cp.AdminEventPassword,
				},
			},
		},
		{
			name: "password_err",
			fields: fields{
				Config: nt.IM{
					"version": testData.version,
				},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				evt: bc.ResponseEvent{
					Trigger: &cp.Admin{
						BaseComponent: bc.BaseComponent{
							Data: nt.IM{
								"username": "name",
								"password": "123",
								"confirm":  "123",
							},
						},
						Token: testData.adminToken,
					},
					Name: cp.AdminEventPassword,
				},
			},
		},
		{
			name: "undo",
			fields: fields{
				Config: nt.IM{
					"version": testData.version,
				},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
				ReadFile: func(name string) ([]byte, error) {
					return []byte{}, nil
				},
				ConvertFromByte: func(data []byte, result interface{}) error {
					bt, _ := ut.ConvertToByte(&nt.IM{"locales": nt.IM{}})
					ut.ConvertFromByte(bt, result)
					return nil
				},
			},
			args: args{
				evt: bc.ResponseEvent{
					Trigger: &cp.Admin{
						BaseComponent: bc.BaseComponent{
							Data: nt.IM{
								"username": "demo",
								"password": "123",
								"confirm":  "123",
							},
						},
						Token: testData.adminToken,
					},
					Name: cp.AdminEventLocalesUndo,
				},
			},
		},
		{
			name: "undo_err",
			fields: fields{
				Config: nt.IM{
					"version": testData.version,
				},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
				ReadFile: func(name string) ([]byte, error) {
					return []byte{}, nil
				},
				ConvertFromByte: func(data []byte, result interface{}) error {
					return errors.New("error")
				},
			},
			args: args{
				evt: bc.ResponseEvent{
					Trigger: &cp.Admin{
						BaseComponent: bc.BaseComponent{
							Data: nt.IM{
								"username": "demo",
								"password": "123",
								"confirm":  "123",
							},
						},
						Token: testData.adminToken,
					},
					Name: cp.AdminEventLocalesUndo,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adm := &AdminService{
				Config:          tt.fields.Config,
				GetNervaStore:   tt.fields.GetNervaStore,
				GetParam:        tt.fields.GetParam,
				GetTokenKeys:    tt.fields.GetTokenKeys,
				GetTaskSecKey:   tt.fields.GetTaskSecKey,
				ReadFile:        tt.fields.ReadFile,
				ConvertFromByte: tt.fields.ConvertFromByte,
				CreateFile:      tt.fields.CreateFile,
				ConvertToWriter: tt.fields.ConvertToWriter,
				Session:         tt.fields.Session,
			}
			adm.appResponse(tt.args.evt)
		})
	}
}

func TestAdminService_Task(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
		GetParam      func(req *http.Request, name string) string
		GetTokenKeys  func() map[string]map[string]string
		GetTaskSecKey func() string
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "unauthorized",
			fields: fields{
				GetTaskSecKey: func() string {
					return "SEC01234"
				},
				GetParam: func(req *http.Request, name string) string {
					return "config"
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/admin/task/config/seckey", nil),
			},
		},
		{
			name: "config",
			fields: fields{
				GetTaskSecKey: func() string {
					return "config"
				},
				GetParam: func(req *http.Request, name string) string {
					return "config"
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/admin/task/config/seckey", nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adm := &AdminService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
				GetParam:      tt.fields.GetParam,
				GetTokenKeys:  tt.fields.GetTokenKeys,
				GetTaskSecKey: tt.fields.GetTaskSecKey,
			}
			adm.Task(tt.args.w, tt.args.r)
		})
	}
}

func TestAdminService_envList(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
		GetParam      func(req *http.Request, name string) string
		GetTokenKeys  func() map[string]map[string]string
		GetTaskSecKey func() string
		Session       nt.SessionService
	}
	os.Setenv("NT_ALIAS_KALEVALA", "sqlite5://file:../data/empty.db")
	tests := []struct {
		name   string
		fields fields
		want   []nt.IM
	}{
		{
			name: "ok",
			fields: fields{
				Config: nt.IM{
					"NT_API_KEY": "EXAMPLE_API_KEY",
				},
			},
			want: []nt.IM{
				{"envkey": "NT_ALIAS_KALEVALA", "envvalue": "sqlite5://file:../data/empty.db"},
				{"envkey": "NT_API_KEY", "envvalue": "EXAMPLE_API_KEY"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adm := &AdminService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
				GetParam:      tt.fields.GetParam,
				GetTokenKeys:  tt.fields.GetTokenKeys,
				GetTaskSecKey: tt.fields.GetTaskSecKey,
				Session:       tt.fields.Session,
			}
			if got := adm.envList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AdminService.envList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdminService_userLogin(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
		GetParam      func(req *http.Request, name string) string
		GetTokenKeys  func() map[string]map[string]string
		GetTaskSecKey func() string
		Session       nt.SessionService
	}
	type args struct {
		data nt.IM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "scope_error",
			fields: fields{
				Config: nt.IM{
					"version": testData.version,
				},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE": testData.hashTable,
					})
				},
			},
			args: args{
				data: nt.IM{
					"username": "user",
					"database": testData.testDatabase,
				},
			},
			wantErr: true,
		},
		{
			name: "admin",
			fields: fields{
				Config: nt.IM{
					"version": testData.version,
				},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE": testData.hashTable,
					})
				},
			},
			args: args{
				data: nt.IM{
					"username": "admin",
					"database": testData.testDatabase,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adm := &AdminService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
				GetParam:      tt.fields.GetParam,
				GetTokenKeys:  tt.fields.GetTokenKeys,
				GetTaskSecKey: tt.fields.GetTaskSecKey,
				Session:       tt.fields.Session,
			}
			_, _, err := adm.userLogin(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("AdminService.userLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestAdminService_createDatabase(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
		GetParam      func(req *http.Request, name string) string
		GetTokenKeys  func() map[string]map[string]string
		GetTaskSecKey func() string
		Session       nt.SessionService
	}
	type args struct {
		data nt.IM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "invalid_api_key",
			fields: fields{
				Config: nt.IM{
					"NT_API_KEY": testData.apiKey,
				},
			},
			args: args{
				data: nt.IM{
					"api_key": "",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adm := &AdminService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
				GetParam:      tt.fields.GetParam,
				GetTokenKeys:  tt.fields.GetTokenKeys,
				GetTaskSecKey: tt.fields.GetTaskSecKey,
				Session:       tt.fields.Session,
			}
			_, err := adm.createDatabase(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("AdminService.createDatabase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestAdminService_reportInstall(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
		GetParam      func(req *http.Request, name string) string
		GetTokenKeys  func() map[string]map[string]string
		GetTaskSecKey func() string
		Session       nt.SessionService
	}
	type args struct {
		token     string
		database  string
		reportkey string
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
				Config: nt.IM{
					"version": testData.version,
				},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				token:     testData.adminToken,
				database:  testData.testDatabase,
				reportkey: "key",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adm := &AdminService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
				GetParam:      tt.fields.GetParam,
				GetTokenKeys:  tt.fields.GetTokenKeys,
				GetTaskSecKey: tt.fields.GetTaskSecKey,
				Session:       tt.fields.Session,
			}
			_, err := adm.reportInstall(tt.args.token, tt.args.database, tt.args.reportkey)
			if (err != nil) != tt.wantErr {
				t.Errorf("AdminService.reportInstall() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestAdminService_reportDelete(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
		GetParam      func(req *http.Request, name string) string
		GetTokenKeys  func() map[string]map[string]string
		GetTaskSecKey func() string
		Session       nt.SessionService
	}
	type args struct {
		token     string
		database  string
		reportkey string
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
				Config: nt.IM{
					"version": testData.version,
				},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				token:     testData.adminToken,
				database:  testData.testDatabase,
				reportkey: "key",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adm := &AdminService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
				GetParam:      tt.fields.GetParam,
				GetTokenKeys:  tt.fields.GetTokenKeys,
				GetTaskSecKey: tt.fields.GetTaskSecKey,
				Session:       tt.fields.Session,
			}
			_, err := adm.reportDelete(tt.args.token, tt.args.database, tt.args.reportkey)
			if (err != nil) != tt.wantErr {
				t.Errorf("AdminService.reportDelete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestAdminService_userPassword(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
		GetParam      func(req *http.Request, name string) string
		GetTokenKeys  func() map[string]map[string]string
		GetTaskSecKey func() string
		Session       nt.SessionService
	}
	type args struct {
		token    string
		database string
		data     nt.IM
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
				Config: nt.IM{
					"version": testData.version,
				},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				token:    testData.adminToken,
				database: testData.testDatabase,
				data: nt.IM{
					"username": "name",
					"password": "123",
					"confirm":  "123",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adm := &AdminService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
				GetParam:      tt.fields.GetParam,
				GetTokenKeys:  tt.fields.GetTokenKeys,
				GetTaskSecKey: tt.fields.GetTaskSecKey,
				Session:       tt.fields.Session,
			}
			if err := adm.userPassword(tt.args.token, tt.args.database, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("AdminService.userPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdminService_tokenLogin(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
		GetParam      func(req *http.Request, name string) string
		GetTokenKeys  func() map[string]map[string]string
		GetTaskSecKey func() string
		Session       nt.SessionService
	}
	type args struct {
		database string
		token    string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "true",
			fields: fields{
				Config: nt.IM{
					"version": testData.version,
				},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				database: testData.testDatabase,
				token:    testData.adminToken,
			},
			want: true,
		},
		{
			name: "false",
			fields: fields{
				Config: nt.IM{
					"version": testData.version,
				},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				database: testData.testDatabase,
				token:    "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adm := &AdminService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
				GetParam:      tt.fields.GetParam,
				GetTokenKeys:  tt.fields.GetTokenKeys,
				GetTaskSecKey: tt.fields.GetTaskSecKey,
				Session:       tt.fields.Session,
			}
			if got := adm.tokenLogin(tt.args.database, tt.args.token); got != tt.want {
				t.Errorf("AdminService.tokenLogin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdminService_Home(t *testing.T) {
	type fields struct {
		Config          map[string]interface{}
		GetNervaStore   func(database string) *nt.NervaStore
		GetParam        func(req *http.Request, name string) string
		GetTokenKeys    func() map[string]map[string]string
		GetTaskSecKey   func() string
		ReadFile        func(name string) ([]byte, error)
		ConvertFromByte func(data []byte, result interface{}) error
		Session         nt.SessionService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "ok",
			fields: fields{
				Config: nt.IM{},
				ReadFile: func(name string) ([]byte, error) {
					return []byte{}, nil
				},
				ConvertFromByte: func(data []byte, result interface{}) error {
					bt, _ := ut.ConvertToByte(&nt.IM{"locales": nt.IM{"de": nt.IM{"key": "value"}}})
					ut.ConvertFromByte(bt, result)
					return nil
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/admin", nil),
			},
		},
		{
			name: "locales_err",
			fields: fields{
				Config: nt.IM{},
				ReadFile: func(name string) ([]byte, error) {
					return []byte{}, nil
				},
				ConvertFromByte: func(data []byte, result interface{}) error {
					return errors.New("error")
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/admin", nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adm := &AdminService{
				Config:          tt.fields.Config,
				GetNervaStore:   tt.fields.GetNervaStore,
				GetParam:        tt.fields.GetParam,
				GetTokenKeys:    tt.fields.GetTokenKeys,
				GetTaskSecKey:   tt.fields.GetTaskSecKey,
				ReadFile:        tt.fields.ReadFile,
				ConvertFromByte: tt.fields.ConvertFromByte,
				Session:         tt.fields.Session,
			}
			adm.Home(tt.args.w, tt.args.r)
		})
	}
}

func TestAdminService_AppEvent(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
		GetParam      func(req *http.Request, name string) string
		GetTokenKeys  func() map[string]map[string]string
		GetTaskSecKey func() string
		Session       nt.SessionService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "mem",
			fields: fields{
				Config: nt.IM{},
				Session: nt.SessionService{
					MemSession: nt.IM{
						"SessionID": &cp.Admin{},
					},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/event", nil),
			},
		},
		{
			name: "file",
			fields: fields{
				Config: nt.IM{},
				Session: nt.SessionService{
					Config: nt.IM{
						"NT_SESSION_DIR": "dir",
					},
					ReadFile: func(name string) ([]byte, error) {
						return []byte{}, nil
					},
					ConvertFromByte: func(data []byte, result interface{}) error {
						bt, _ := ut.ConvertToByte(&cp.Admin{})
						ut.ConvertFromByte(bt, result)
						return nil
					},
					FileStat: func(name string) (fs.FileInfo, error) {
						return nil, nil
					},
					CreateFile: func(name string) (*os.File, error) {
						return nil, errors.New("error")
					},
					Method: "file",
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/event", nil),
			},
		},
		{
			name: "nodata",
			fields: fields{
				Config:  nt.IM{},
				Session: nt.SessionService{},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/event", nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adm := &AdminService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
				GetParam:      tt.fields.GetParam,
				GetTokenKeys:  tt.fields.GetTokenKeys,
				GetTaskSecKey: tt.fields.GetTaskSecKey,
				Session:       tt.fields.Session,
			}
			tt.args.r.Header.Set("X-Session-Token", "SessionID")
			adm.AppEvent(tt.args.w, tt.args.r)
		})
	}
}
