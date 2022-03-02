package service

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	"text/template"

	db "github.com/nervatura/nervatura-service/pkg/database"
	nt "github.com/nervatura/nervatura-service/pkg/nervatura"
)

func TestAdminService_LoadTemplates(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
		templates     *template.Template
		GetTokenKeys  func() map[string]map[string]string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "load_templates",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adm := &AdminService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
				templates:     tt.fields.templates,
				GetTokenKeys:  tt.fields.GetTokenKeys,
			}
			adm.LoadTemplates()
		})
	}
}

func TestAdminService_render(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
		templates     *template.Template
		GetTokenKeys  func() map[string]map[string]string
	}
	type args struct {
		w        http.ResponseWriter
		template string
		data     interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "missing_template",
			fields: fields{
				Config: nt.IM{
					"version": "test",
				},
				templates: testData.templates(),
			},
			args: args{
				w:        httptest.NewRecorder(),
				template: "missing",
				data:     nt.IM{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adm := &AdminService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
				templates:     tt.fields.templates,
				GetTokenKeys:  tt.fields.GetTokenKeys,
			}
			adm.render(tt.args.w, tt.args.template, tt.args.data)
		})
	}
}

func TestAdminService_Home(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
		templates     *template.Template
		GetTokenKeys  func() map[string]map[string]string
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
			name: "login",
			fields: fields{
				Config: nt.IM{
					"version": "test",
				},
				templates: testData.templates(),
			},
			args: args{
				r: testData.formReq(url.Values{}),
				w: httptest.NewRecorder(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adm := &AdminService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
				templates:     tt.fields.templates,
				GetTokenKeys:  tt.fields.GetTokenKeys,
			}
			adm.Home(tt.args.w, tt.args.r)
		})
	}
}

func TestAdminService_Login(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
		templates     *template.Template
		GetTokenKeys  func() map[string]map[string]string
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
			name: "login_ok",
			fields: fields{
				Config: nt.IM{
					"version": testData.version,
				},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE": testData.hashTable,
					})
				},
				templates: testData.templates(),
			},
			args: args{
				r: testData.formReq(url.Values{
					"username": []string{"admin"},
					"database": []string{"test"},
				}),
				w: httptest.NewRecorder(),
			},
		},
		{
			name: "missing_database",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE": testData.hashTable,
					})
				},
				templates: testData.templates(),
			},
			args: args{
				r: testData.formReq(url.Values{
					"username": []string{"admin"},
				}),
				w: httptest.NewRecorder(),
			},
		},
		{
			name: "login_error",
			fields: fields{
				Config: nt.IM{
					"version": testData.version,
				},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE": testData.hashTable,
					})
				},
				templates: testData.templates(),
			},
			args: args{
				r: testData.formReq(url.Values{
					"username": []string{"admin"},
					"password": []string{"12345678"},
					"database": []string{"test"},
				}),
				w: httptest.NewRecorder(),
			},
		},
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
				templates: testData.templates(),
			},
			args: args{
				r: testData.formReq(url.Values{
					"username": []string{"user"},
					"database": []string{"test"},
				}),
				w: httptest.NewRecorder(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adm := &AdminService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
				templates:     tt.fields.templates,
				GetTokenKeys:  tt.fields.GetTokenKeys,
			}
			adm.Login(tt.args.w, tt.args.r)
		})
	}
}

func TestAdminService_Menu(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
		templates     *template.Template
		GetTokenKeys  func() map[string]map[string]string
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
			name: "database",
			fields: fields{
				Config: nt.IM{
					"version": "test",
				},
				templates: testData.templates(),
			},
			args: args{
				r: testData.formReq(url.Values{
					"menu": []string{"database"},
				}),
				w: httptest.NewRecorder(),
			},
		},
		{
			name: "theme_dark",
			fields: fields{
				Config: nt.IM{
					"version": "test",
				},
				templates: testData.templates(),
			},
			args: args{
				r: testData.formReq(url.Values{
					"menu": []string{"theme"},
				}),
				w: httptest.NewRecorder(),
			},
		},
		{
			name: "theme_light",
			fields: fields{
				Config: nt.IM{
					"version": "test",
				},
				templates: testData.templates(),
			},
			args: args{
				r: testData.formReq(url.Values{
					"menu":  []string{"theme"},
					"theme": []string{"dark"},
				}),
				w: httptest.NewRecorder(),
			},
		},
		{
			name: "logout",
			fields: fields{
				Config: nt.IM{
					"version": "test",
				},
				templates: testData.templates(),
			},
			args: args{
				r: testData.formReq(url.Values{
					"menu": []string{"logout"},
				}),
				w: httptest.NewRecorder(),
			},
		},
		{
			name: "login",
			fields: fields{
				Config: nt.IM{
					"version": "test",
				},
				templates: testData.templates(),
			},
			args: args{
				r: testData.formReq(url.Values{
					"menu": []string{"login"},
				}),
				w: httptest.NewRecorder(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adm := &AdminService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
				templates:     tt.fields.templates,
				GetTokenKeys:  tt.fields.GetTokenKeys,
			}
			adm.Menu(tt.args.w, tt.args.r)
		})
	}
}

func TestAdminService_Admin(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
		templates     *template.Template
		GetTokenKeys  func() map[string]map[string]string
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	os.Setenv("NT_ALIAS_KALEVALA", "sqlite5://file:../data/empty.db")
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "refresh",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
				templates: testData.templates(),
			},
			args: args{
				r: testData.formReq(url.Values{
					"cmd":   []string{"refresh"},
					"token": []string{testData.adminToken},
				}),
				w: httptest.NewRecorder(),
			},
		},
		{
			name: "password",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
				templates: testData.templates(),
			},
			args: args{
				r: testData.formReq(url.Values{
					"cmd":      []string{"password"},
					"token":    []string{testData.adminToken},
					"username": []string{"demo"},
					"password": []string{"123"},
					"confirm":  []string{"123"},
				}),
				w: httptest.NewRecorder(),
			},
		},
		{
			name: "report_list",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
						"NT_REPORT_DIR":        "",
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
				templates: testData.templates(),
			},
			args: args{
				r: testData.formReq(url.Values{
					"cmd":   []string{"list"},
					"token": []string{testData.adminToken},
				}),
				w: httptest.NewRecorder(),
			},
		},
		{
			name: "report_delete",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
						"NT_REPORT_DIR":        "",
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
				templates: testData.templates(),
			},
			args: args{
				r: testData.formReq(url.Values{
					"cmd":       []string{"delete"},
					"token":     []string{testData.adminToken},
					"reportkey": []string{"ntr_invoice_en"},
				}),
				w: httptest.NewRecorder(),
			},
		},
		{
			name: "report_install",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
						"NT_REPORT_DIR":        "",
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
				templates: testData.templates(),
			},
			args: args{
				r: testData.formReq(url.Values{
					"cmd":       []string{"install"},
					"token":     []string{testData.adminToken},
					"reportkey": []string{"ntr_invoice_en"},
				}),
				w: httptest.NewRecorder(),
			},
		},
		{
			name: "env_list",
			fields: fields{
				Config: nt.IM{
					"NT_ALIAS_TEST": testData.testDatabase,
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
				templates: testData.templates(),
			},
			args: args{
				r: testData.formReq(url.Values{
					"cmd":   []string{"env_list"},
					"token": []string{testData.adminToken},
				}),
				w: httptest.NewRecorder(),
			},
		},
		{
			name: "error",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
						"NT_REPORT_DIR":        "",
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
				templates: testData.templates(),
			},
			args: args{
				r: testData.formReq(url.Values{
					"cmd":       []string{"install"},
					"token":     []string{testData.adminToken},
					"reportkey": []string{"ntr_invoice_en"},
				}),
				w: httptest.NewRecorder(),
			},
		},
		{
			name: "admin_rights",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
				templates: testData.templates(),
			},
			args: args{
				r: testData.formReq(url.Values{
					"cmd":   []string{"refresh"},
					"token": []string{testData.customerToken},
				}),
				w: httptest.NewRecorder(),
			},
		},
		{
			name: "unauthorized",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
				templates: testData.templates(),
			},
			args: args{
				r: testData.formReq(url.Values{
					"cmd":   []string{"refresh"},
					"token": []string{"ERRORORTOKENE"},
				}),
				w: httptest.NewRecorder(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adm := &AdminService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
				templates:     tt.fields.templates,
				GetTokenKeys:  tt.fields.GetTokenKeys,
			}
			adm.Admin(tt.args.w, tt.args.r)
		})
	}
}

func TestAdminService_Database(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
		templates     *template.Template
		GetTokenKeys  func() map[string]map[string]string
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
			name: "invalid_api_key",
			fields: fields{
				Config: nt.IM{
					"NT_API_KEY": testData.apiKey,
				},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(&db.SQLDriver{
						Config: nt.IM{"version": testData.version}}, nt.IM{
						"version":       testData.version,
						"NT_HASHTABLE":  testData.hashTable,
						"NT_API_KEY":    testData.apiKey,
						"NT_ALIAS_TEST": testData.testDatabase,
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
				templates: testData.templates(),
			},
			args: args{
				r: testData.formReq(url.Values{
					"database": []string{"test"},
					"demo":     []string{"true"},
				}),
				w: httptest.NewRecorder(),
			},
		},
		{
			name: "create",
			fields: fields{
				Config: nt.IM{
					"NT_API_KEY": testData.apiKey,
				},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(&db.SQLDriver{
						Config: nt.IM{"version": testData.version}}, nt.IM{
						"version":       testData.version,
						"NT_HASHTABLE":  testData.hashTable,
						"NT_API_KEY":    testData.apiKey,
						"NT_ALIAS_TEST": testData.testDatabase,
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
				templates: testData.templates(),
			},
			args: args{
				r: testData.formReq(url.Values{
					"database": []string{"test"},
					"demo":     []string{"true"},
					"apikey":   []string{testData.apiKey},
				}),
				w: httptest.NewRecorder(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adm := &AdminService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
				templates:     tt.fields.templates,
				GetTokenKeys:  tt.fields.GetTokenKeys,
			}
			adm.Database(tt.args.w, tt.args.r)
		})
	}
}
