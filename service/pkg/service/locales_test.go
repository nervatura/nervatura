//go:build http || all
// +build http all

package service

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"

	nt "github.com/nervatura/nervatura/service/pkg/nervatura"
)

func TestLocalesService_loadData(t *testing.T) {
	type fields struct {
		ClientMsg  string
		ConfigFile string
		Version    string
		GetParam   func(req *http.Request, name string) string
		theme      string
		templates  *template.Template
		langs      nt.SL
		tagKeys    nt.SL
		tagValues  map[string]nt.SL
		filter     string
		edited     bool
		deflang    nt.IM
		locfile    nt.IM
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				Version:    "test",
				ConfigFile: "../../data/test_client_config.json",
				ClientMsg:  "static/locales/client.json",
			},
			wantErr: false,
		},
		{
			name: "error 1.",
			fields: fields{
				Version:    "test",
				ConfigFile: "../../data/test_client_config.json",
				ClientMsg:  "",
			},
			wantErr: true,
		},
		{
			name: "missing client.json",
			fields: fields{
				Version:    "test",
				ConfigFile: "",
				ClientMsg:  "static/locales/client.json",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loc := &LocalesService{
				ClientMsg:  tt.fields.ClientMsg,
				ConfigFile: tt.fields.ConfigFile,
				Version:    tt.fields.Version,
				GetParam:   tt.fields.GetParam,
				theme:      tt.fields.theme,
				templates:  tt.fields.templates,
				langs:      tt.fields.langs,
				tagKeys:    tt.fields.tagKeys,
				tagValues:  tt.fields.tagValues,
				filter:     tt.fields.filter,
				edited:     tt.fields.edited,
				deflang:    tt.fields.deflang,
				locfile:    tt.fields.locfile,
			}
			if err := loc.loadData(); (err != nil) != tt.wantErr {
				t.Errorf("LocalesService.loadData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLocalesService_LoadLocales(t *testing.T) {
	type fields struct {
		ClientMsg  string
		ConfigFile string
		Version    string
		GetParam   func(req *http.Request, name string) string
		theme      string
		templates  *template.Template
		langs      nt.SL
		tagKeys    nt.SL
		tagValues  map[string]nt.SL
		filter     string
		edited     bool
		deflang    nt.IM
		locfile    nt.IM
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				Version:    "test",
				ConfigFile: "../../data/test_client_config.json",
				ClientMsg:  "static/locales/client.json",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loc := &LocalesService{
				ClientMsg:  tt.fields.ClientMsg,
				ConfigFile: tt.fields.ConfigFile,
				Version:    tt.fields.Version,
				GetParam:   tt.fields.GetParam,
				theme:      tt.fields.theme,
				templates:  tt.fields.templates,
				langs:      tt.fields.langs,
				tagKeys:    tt.fields.tagKeys,
				tagValues:  tt.fields.tagValues,
				filter:     tt.fields.filter,
				edited:     tt.fields.edited,
				deflang:    tt.fields.deflang,
				locfile:    tt.fields.locfile,
			}
			if err := loc.LoadLocales(); (err != nil) != tt.wantErr {
				t.Errorf("LocalesService.LoadLocales() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLocalesService_sendMsg(t *testing.T) {
	type fields struct {
		ClientMsg  string
		ConfigFile string
		Version    string
		GetParam   func(req *http.Request, name string) string
		theme      string
		templates  *template.Template
		langs      nt.SL
		tagKeys    nt.SL
		tagValues  map[string]nt.SL
		filter     string
		edited     bool
		deflang    nt.IM
		locfile    nt.IM
	}
	type args struct {
		w       http.ResponseWriter
		code    int
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "ok",
			args: args{
				w:       httptest.NewRecorder(),
				code:    200,
				message: "OK",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loc := &LocalesService{
				ClientMsg:  tt.fields.ClientMsg,
				ConfigFile: tt.fields.ConfigFile,
				Version:    tt.fields.Version,
				GetParam:   tt.fields.GetParam,
				theme:      tt.fields.theme,
				templates:  tt.fields.templates,
				langs:      tt.fields.langs,
				tagKeys:    tt.fields.tagKeys,
				tagValues:  tt.fields.tagValues,
				filter:     tt.fields.filter,
				edited:     tt.fields.edited,
				deflang:    tt.fields.deflang,
				locfile:    tt.fields.locfile,
			}
			loc.sendMsg(tt.args.w, tt.args.code, tt.args.message)
		})
	}
}

func TestLocalesService_createView(t *testing.T) {
	type fields struct {
		ClientMsg  string
		ConfigFile string
		Version    string
		GetParam   func(req *http.Request, name string) string
		theme      string
		templates  *template.Template
		langs      nt.SL
		tagKeys    nt.SL
		tagValues  map[string]nt.SL
		filter     string
		edited     bool
		deflang    nt.IM
		locfile    nt.IM
	}
	type args struct {
		lang string
		tag  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "tag view",
			args: args{
				lang: "ts",
				tag:  "login",
			},
			fields: fields{
				theme:   "dark",
				filter:  "",
				edited:  false,
				Version: "test",
				langs:   nt.SL{"ts", "client"},
				tagKeys: nt.SL{"login"},
				tagValues: map[string]nt.SL{
					"login": {"login_username", "login_password"},
				},
				locfile: nt.IM{
					"locales": nt.IM{
						"ts": nt.IM{
							"login_username": "Username",
						},
					},
				},
				deflang: nt.IM{
					"login_password": "Password",
				},
			},
		},
		{
			name: "missing view",
			args: args{
				lang: "ts",
				tag:  "missing",
			},
			fields: fields{
				theme:   "dark",
				filter:  "",
				edited:  false,
				Version: "test",
				langs:   nt.SL{"ts", "client"},
				tagKeys: nt.SL{"login"},
				tagValues: map[string]nt.SL{
					"login": {"login_username", "login_password"},
				},
				locfile: nt.IM{
					"locales": nt.IM{
						"ts": nt.IM{
							"login_username": "Username",
						},
					},
				},
				deflang: nt.IM{
					"login_username": "Username",
					"login_password": "Password",
				},
			},
		},
		{
			name: "filter view 1",
			args: args{
				lang: "tx",
				tag:  "",
			},
			fields: fields{
				theme:   "dark",
				filter:  "login",
				edited:  false,
				Version: "test",
				langs:   nt.SL{"ts", "client"},
				tagKeys: nt.SL{"login"},
				tagValues: map[string]nt.SL{
					"login": {"login_username", "login_password"},
				},
				locfile: nt.IM{
					"locales": nt.IM{
						"ts": nt.IM{
							"login_username": "Username",
						},
					},
				},
				deflang: nt.IM{
					"login_username": "Username",
					"login_password": "Password",
					"menu_side":      "Menu",
					"menu_test":      "login",
				},
			},
		},
		{
			name: "filter view 2",
			args: args{
				lang: "ts",
				tag:  "",
			},
			fields: fields{
				theme:   "dark",
				filter:  "login",
				edited:  false,
				Version: "test",
				langs:   nt.SL{"ts", "client"},
				tagKeys: nt.SL{"login"},
				tagValues: map[string]nt.SL{
					"login": {"login_username", "login_password"},
				},
				locfile: nt.IM{
					"locales": nt.IM{
						"ts": nt.IM{
							"login_username": "Username",
							"menu_side":      "login",
						},
					},
				},
				deflang: nt.IM{
					"login_username": "Username",
					"login_password": "Password",
					"menu_side":      "Menu",
					"menu_test":      "login",
					"menu_lst":       []interface{}{"apple", "moon", "login"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loc := &LocalesService{
				ClientMsg:  tt.fields.ClientMsg,
				ConfigFile: tt.fields.ConfigFile,
				Version:    tt.fields.Version,
				GetParam:   tt.fields.GetParam,
				theme:      tt.fields.theme,
				templates:  tt.fields.templates,
				langs:      tt.fields.langs,
				tagKeys:    tt.fields.tagKeys,
				tagValues:  tt.fields.tagValues,
				filter:     tt.fields.filter,
				edited:     tt.fields.edited,
				deflang:    tt.fields.deflang,
				locfile:    tt.fields.locfile,
			}
			if got := loc.createView(tt.args.lang, tt.args.tag); len(got) != 11 {
				t.Errorf("LocalesService.createView() = %v", got)
			}
		})
	}
}

func TestLocalesService_Render(t *testing.T) {
	type fields struct {
		ClientMsg    string
		ConfigFile   string
		Version      string
		GetParam     func(req *http.Request, name string) string
		theme        string
		templates    *template.Template
		templateName string
		langs        nt.SL
		tagKeys      nt.SL
		tagValues    map[string]nt.SL
		filter       string
		edited       bool
		deflang      nt.IM
		locfile      nt.IM
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
				theme:   "dark",
				filter:  "",
				edited:  false,
				Version: "test",
				langs:   nt.SL{"ts", "client"},
				tagKeys: nt.SL{"login"},
				tagValues: map[string]nt.SL{
					"login": {"login_username", "login_password"},
				},
				locfile: nt.IM{
					"locales": nt.IM{
						"ts": nt.IM{
							"login_username": "Username",
						},
					},
				},
				deflang: nt.IM{
					"login_password": "Password",
				},
				GetParam: func(req *http.Request, name string) string {
					if name == "lang" {
						return "ts"
					}
					return "login"
				},
				templates:    testData.templates(),
				templateName: "locales",
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/", nil),
			},
		},
		{
			name: "missing template",
			fields: fields{
				theme:   "dark",
				filter:  "",
				edited:  false,
				Version: "test",
				langs:   nt.SL{"ts", "client"},
				tagKeys: nt.SL{"login"},
				tagValues: map[string]nt.SL{
					"login": {"login_username", "login_password"},
				},
				locfile: nt.IM{
					"locales": nt.IM{
						"ts": nt.IM{
							"login_username": "Username",
						},
					},
				},
				deflang: nt.IM{
					"login_password": "Password",
				},
				GetParam: func(req *http.Request, name string) string {
					if name == "lang" {
						return "ts"
					}
					return "login"
				},
				templates:    testData.templates(),
				templateName: "missing",
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/", nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loc := &LocalesService{
				ClientMsg:  tt.fields.ClientMsg,
				ConfigFile: tt.fields.ConfigFile,
				Version:    tt.fields.Version,
				GetParam:   tt.fields.GetParam,
				theme:      tt.fields.theme,
				templates:  tt.fields.templates,
				langs:      tt.fields.langs,
				tagKeys:    tt.fields.tagKeys,
				tagValues:  tt.fields.tagValues,
				filter:     tt.fields.filter,
				edited:     tt.fields.edited,
				deflang:    tt.fields.deflang,
				locfile:    tt.fields.locfile,
			}
			loc.Render(tt.args.w, tt.args.r)
		})
	}
}

func TestLocalesService_SetTheme(t *testing.T) {
	type fields struct {
		ClientMsg    string
		ConfigFile   string
		Version      string
		GetParam     func(req *http.Request, name string) string
		theme        string
		templates    *template.Template
		templateName string
		langs        nt.SL
		tagKeys      nt.SL
		tagValues    map[string]nt.SL
		filter       string
		edited       bool
		deflang      nt.IM
		locfile      nt.IM
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
				theme: "light",
			},
			args: args{
				w: httptest.NewRecorder(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loc := &LocalesService{
				ClientMsg:    tt.fields.ClientMsg,
				ConfigFile:   tt.fields.ConfigFile,
				Version:      tt.fields.Version,
				GetParam:     tt.fields.GetParam,
				theme:        tt.fields.theme,
				templates:    tt.fields.templates,
				templateName: tt.fields.templateName,
				langs:        tt.fields.langs,
				tagKeys:      tt.fields.tagKeys,
				tagValues:    tt.fields.tagValues,
				filter:       tt.fields.filter,
				edited:       tt.fields.edited,
				deflang:      tt.fields.deflang,
				locfile:      tt.fields.locfile,
			}
			loc.SetTheme(tt.args.w, tt.args.r)
		})
	}
}

func TestLocalesService_SetFilter(t *testing.T) {
	type fields struct {
		ClientMsg    string
		ConfigFile   string
		Version      string
		GetParam     func(req *http.Request, name string) string
		theme        string
		templates    *template.Template
		templateName string
		langs        nt.SL
		tagKeys      nt.SL
		tagValues    map[string]nt.SL
		filter       string
		edited       bool
		deflang      nt.IM
		locfile      nt.IM
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
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/filter", bytes.NewBuffer([]byte(testData.encodeData(nt.IM{
					"filter_value": "login",
				})))),
			},
		},
		{
			name: "error",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/filter", bytes.NewBuffer([]byte(testData.encodeData([]nt.IM{})))),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loc := &LocalesService{
				ClientMsg:    tt.fields.ClientMsg,
				ConfigFile:   tt.fields.ConfigFile,
				Version:      tt.fields.Version,
				GetParam:     tt.fields.GetParam,
				theme:        tt.fields.theme,
				templates:    tt.fields.templates,
				templateName: tt.fields.templateName,
				langs:        tt.fields.langs,
				tagKeys:      tt.fields.tagKeys,
				tagValues:    tt.fields.tagValues,
				filter:       tt.fields.filter,
				edited:       tt.fields.edited,
				deflang:      tt.fields.deflang,
				locfile:      tt.fields.locfile,
			}
			loc.SetFilter(tt.args.w, tt.args.r)
		})
	}
}

func TestLocalesService_updateData(t *testing.T) {
	type fields struct {
		ClientMsg    string
		ConfigFile   string
		Version      string
		GetParam     func(req *http.Request, name string) string
		theme        string
		templates    *template.Template
		templateName string
		langs        nt.SL
		tagKeys      nt.SL
		tagValues    map[string]nt.SL
		filter       string
		edited       bool
		deflang      nt.IM
		locfile      nt.IM
	}
	type args struct {
		lang  string
		key   string
		value string
		index int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "simple OK",
			args: args{
				lang:  "ts",
				key:   "menu_side",
				value: "test",
				index: 0,
			},
			fields: fields{
				edited: false,
				locfile: nt.IM{
					"locales": nt.IM{
						"ts": nt.IM{
							"login_username": "Username",
							"menu_side":      "login",
						},
					},
				},
				deflang: nt.IM{
					"login_username": "Username",
					"login_password": "Password",
					"menu_side":      "Menu",
					"menu_test":      "login",
					"menu_lst":       []interface{}{"apple", "moon", "login"},
				},
			},
			wantErr: false,
		},
		{
			name: "arr. OK",
			args: args{
				lang:  "ts",
				key:   "menu_lst",
				value: "baba",
				index: 1,
			},
			fields: fields{
				edited: false,
				locfile: nt.IM{
					"locales": nt.IM{
						"ts": nt.IM{
							"login_username": "Username",
							"menu_side":      "login",
						},
					},
				},
				deflang: nt.IM{
					"login_username": "Username",
					"login_password": "Password",
					"menu_side":      "Menu",
					"menu_test":      "login",
					"menu_lst":       []interface{}{"apple", "moon", "login"},
				},
			},
			wantErr: false,
		},
		{
			name: "arr. error",
			args: args{
				lang:  "ts",
				key:   "menu_lst",
				value: "baba",
				index: 6,
			},
			fields: fields{
				edited: false,
				locfile: nt.IM{
					"locales": nt.IM{
						"ts": nt.IM{
							"login_username": "Username",
							"menu_side":      "login",
						},
					},
				},
				deflang: nt.IM{
					"login_username": "Username",
					"login_password": "Password",
					"menu_side":      "Menu",
					"menu_test":      "login",
					"menu_lst":       []interface{}{"apple", "moon", "login"},
				},
			},
			wantErr: true,
		},
		{
			name: "missing",
			args: args{
				lang:  "ts",
				key:   "missing",
				value: "baba",
				index: 0,
			},
			fields: fields{
				edited: false,
				locfile: nt.IM{
					"locales": nt.IM{
						"ts": nt.IM{
							"login_username": "Username",
							"menu_side":      "login",
						},
					},
				},
				deflang: nt.IM{
					"login_username": "Username",
					"login_password": "Password",
					"menu_side":      "Menu",
					"menu_test":      "login",
					"menu_lst":       []interface{}{"apple", "moon", "login"},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loc := &LocalesService{
				ClientMsg:    tt.fields.ClientMsg,
				ConfigFile:   tt.fields.ConfigFile,
				Version:      tt.fields.Version,
				GetParam:     tt.fields.GetParam,
				theme:        tt.fields.theme,
				templates:    tt.fields.templates,
				templateName: tt.fields.templateName,
				langs:        tt.fields.langs,
				tagKeys:      tt.fields.tagKeys,
				tagValues:    tt.fields.tagValues,
				filter:       tt.fields.filter,
				edited:       tt.fields.edited,
				deflang:      tt.fields.deflang,
				locfile:      tt.fields.locfile,
			}
			if err := loc.updateData(tt.args.lang, tt.args.key, tt.args.value, tt.args.index); (err != nil) != tt.wantErr {
				t.Errorf("LocalesService.updateData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLocalesService_UpdateRow(t *testing.T) {
	type fields struct {
		ClientMsg    string
		ConfigFile   string
		Version      string
		GetParam     func(req *http.Request, name string) string
		theme        string
		templates    *template.Template
		templateName string
		langs        nt.SL
		tagKeys      nt.SL
		tagValues    map[string]nt.SL
		filter       string
		edited       bool
		deflang      nt.IM
		locfile      nt.IM
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
				edited: false,
				locfile: nt.IM{
					"locales": nt.IM{
						"ts": nt.IM{
							"login_username": "Username",
							"menu_side":      "login",
						},
					},
				},
				deflang: nt.IM{
					"login_username": "Username",
					"login_password": "Password",
					"menu_side":      "Menu",
					"menu_test":      "login",
					"menu_lst":       []interface{}{"apple", "moon", "login"},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/filter", bytes.NewBuffer([]byte(testData.encodeData(nt.IM{
					"lang": "ts", "key": "menu_side", "value": "test", "index": 0,
				})))),
			},
		},
		{
			name: "update error",
			fields: fields{
				edited: false,
				locfile: nt.IM{
					"locales": nt.IM{
						"ts": nt.IM{
							"login_username": "Username",
							"menu_side":      "login",
						},
					},
				},
				deflang: nt.IM{
					"login_username": "Username",
					"login_password": "Password",
					"menu_side":      "Menu",
					"menu_test":      "login",
					"menu_lst":       []interface{}{"apple", "moon", "login"},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/filter", bytes.NewBuffer([]byte(testData.encodeData(nt.IM{
					"lang": "ts", "key": "menu_lst", "value": "test", "index": 9,
				})))),
			},
		},
		{
			name: "convert error",
			fields: fields{
				edited: false,
				locfile: nt.IM{
					"locales": nt.IM{
						"ts": nt.IM{
							"login_username": "Username",
							"menu_side":      "login",
						},
					},
				},
				deflang: nt.IM{
					"login_username": "Username",
					"login_password": "Password",
					"menu_side":      "Menu",
					"menu_test":      "login",
					"menu_lst":       []interface{}{"apple", "moon", "login"},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/filter", bytes.NewBuffer([]byte(testData.encodeData([]nt.IM{})))),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loc := &LocalesService{
				ClientMsg:    tt.fields.ClientMsg,
				ConfigFile:   tt.fields.ConfigFile,
				Version:      tt.fields.Version,
				GetParam:     tt.fields.GetParam,
				theme:        tt.fields.theme,
				templates:    tt.fields.templates,
				templateName: tt.fields.templateName,
				langs:        tt.fields.langs,
				tagKeys:      tt.fields.tagKeys,
				tagValues:    tt.fields.tagValues,
				filter:       tt.fields.filter,
				edited:       tt.fields.edited,
				deflang:      tt.fields.deflang,
				locfile:      tt.fields.locfile,
			}
			loc.UpdateRow(tt.args.w, tt.args.r)
		})
	}
}

func TestLocalesService_UndoAll(t *testing.T) {
	type fields struct {
		ClientMsg    string
		ConfigFile   string
		Version      string
		GetParam     func(req *http.Request, name string) string
		theme        string
		templates    *template.Template
		templateName string
		langs        nt.SL
		tagKeys      nt.SL
		tagValues    map[string]nt.SL
		filter       string
		edited       bool
		deflang      nt.IM
		locfile      nt.IM
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
				Version:    "test",
				ConfigFile: "../../data/test_client_config.json",
				ClientMsg:  "static/locales/client.json",
			},
			args: args{
				w: httptest.NewRecorder(),
			},
		},
		{
			name: "error",
			fields: fields{
				Version:    "test",
				ConfigFile: "../../data/test_client_config.json",
				ClientMsg:  "",
			},
			args: args{
				w: httptest.NewRecorder(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loc := &LocalesService{
				ClientMsg:    tt.fields.ClientMsg,
				ConfigFile:   tt.fields.ConfigFile,
				Version:      tt.fields.Version,
				GetParam:     tt.fields.GetParam,
				theme:        tt.fields.theme,
				templates:    tt.fields.templates,
				templateName: tt.fields.templateName,
				langs:        tt.fields.langs,
				tagKeys:      tt.fields.tagKeys,
				tagValues:    tt.fields.tagValues,
				filter:       tt.fields.filter,
				edited:       tt.fields.edited,
				deflang:      tt.fields.deflang,
				locfile:      tt.fields.locfile,
			}
			loc.UndoAll(tt.args.w, tt.args.r)
		})
	}
}

func TestLocalesService_saveFile(t *testing.T) {
	type fields struct {
		ClientMsg    string
		ConfigFile   string
		Version      string
		GetParam     func(req *http.Request, name string) string
		theme        string
		templates    *template.Template
		templateName string
		langs        nt.SL
		tagKeys      nt.SL
		tagValues    map[string]nt.SL
		filter       string
		edited       bool
		deflang      nt.IM
		locfile      nt.IM
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				edited:     true,
				ConfigFile: "../../data/test_client_config.json",
				locfile: nt.IM{
					"locales": nt.IM{
						"ts": nt.IM{
							"login_username": "Username",
							"menu_side":      "login",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "config file error",
			fields: fields{
				edited:     true,
				ConfigFile: "../data/test_client_config.json",
				locfile: nt.IM{
					"locales": nt.IM{
						"ts": nt.IM{
							"login_username": "Username",
							"menu_side":      "login",
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "skip",
			fields: fields{
				edited:     false,
				ConfigFile: "../../data/test_client_config.json",
				locfile: nt.IM{
					"locales": nt.IM{
						"ts": nt.IM{
							"login_username": "Username",
							"menu_side":      "login",
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loc := &LocalesService{
				ClientMsg:    tt.fields.ClientMsg,
				ConfigFile:   tt.fields.ConfigFile,
				Version:      tt.fields.Version,
				GetParam:     tt.fields.GetParam,
				theme:        tt.fields.theme,
				templates:    tt.fields.templates,
				templateName: tt.fields.templateName,
				langs:        tt.fields.langs,
				tagKeys:      tt.fields.tagKeys,
				tagValues:    tt.fields.tagValues,
				filter:       tt.fields.filter,
				edited:       tt.fields.edited,
				deflang:      tt.fields.deflang,
				locfile:      tt.fields.locfile,
			}
			if err := loc.saveFile(); (err != nil) != tt.wantErr {
				t.Errorf("LocalesService.saveFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLocalesService_CreateFile(t *testing.T) {
	type fields struct {
		ClientMsg    string
		ConfigFile   string
		Version      string
		GetParam     func(req *http.Request, name string) string
		theme        string
		templates    *template.Template
		templateName string
		langs        nt.SL
		tagKeys      nt.SL
		tagValues    map[string]nt.SL
		filter       string
		edited       bool
		deflang      nt.IM
		locfile      nt.IM
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
				edited:     true,
				Version:    "test",
				ConfigFile: "../../data/test_client_config.json",
				ClientMsg:  "static/locales/client.json",
				locfile: nt.IM{
					"locales": nt.IM{
						"ts": nt.IM{
							"login_username": "Username",
							"menu_side":      "login",
						},
					},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
			},
		},
		{
			name: "error",
			fields: fields{
				edited:     true,
				Version:    "test",
				ConfigFile: "../../data/test_client_config.json",
				ClientMsg:  "",
				locfile: nt.IM{
					"locales": nt.IM{
						"ts": nt.IM{
							"login_username": "Username",
							"menu_side":      "login",
						},
					},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loc := &LocalesService{
				ClientMsg:    tt.fields.ClientMsg,
				ConfigFile:   tt.fields.ConfigFile,
				Version:      tt.fields.Version,
				GetParam:     tt.fields.GetParam,
				theme:        tt.fields.theme,
				templates:    tt.fields.templates,
				templateName: tt.fields.templateName,
				langs:        tt.fields.langs,
				tagKeys:      tt.fields.tagKeys,
				tagValues:    tt.fields.tagValues,
				filter:       tt.fields.filter,
				edited:       tt.fields.edited,
				deflang:      tt.fields.deflang,
				locfile:      tt.fields.locfile,
			}
			loc.CreateFile(tt.args.w, tt.args.r)
		})
	}
}

func TestLocalesService_AddLang(t *testing.T) {
	type fields struct {
		ClientMsg    string
		ConfigFile   string
		Version      string
		GetParam     func(req *http.Request, name string) string
		theme        string
		templates    *template.Template
		templateName string
		langs        nt.SL
		tagKeys      nt.SL
		tagValues    map[string]nt.SL
		filter       string
		edited       bool
		deflang      nt.IM
		locfile      nt.IM
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
				edited: false,
				langs:  nt.SL{"client"},
				locfile: nt.IM{
					"locales": nt.IM{},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/add", bytes.NewBuffer([]byte(testData.encodeData(nt.IM{
					"lang_key": "new", "lang_name": "Lang",
				})))),
			},
		},
		{
			name: "missing value",
			fields: fields{
				edited: false,
				langs:  nt.SL{"client"},
				locfile: nt.IM{
					"locales": nt.IM{},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/add", bytes.NewBuffer([]byte(testData.encodeData(nt.IM{
					"lang_key": "", "lang_name": "Lang",
				})))),
			},
		},
		{
			name: "exists lang",
			fields: fields{
				edited: false,
				langs:  nt.SL{"client"},
				locfile: nt.IM{
					"locales": nt.IM{},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/add", bytes.NewBuffer([]byte(testData.encodeData(nt.IM{
					"lang_key": "en", "lang_name": "Lang",
				})))),
			},
		},
		{
			name: "convert error",
			fields: fields{
				edited: false,
				langs:  nt.SL{"client"},
				locfile: nt.IM{
					"locales": nt.IM{},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/add", bytes.NewBuffer([]byte(testData.encodeData([]nt.IM{})))),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loc := &LocalesService{
				ClientMsg:    tt.fields.ClientMsg,
				ConfigFile:   tt.fields.ConfigFile,
				Version:      tt.fields.Version,
				GetParam:     tt.fields.GetParam,
				theme:        tt.fields.theme,
				templates:    tt.fields.templates,
				templateName: tt.fields.templateName,
				langs:        tt.fields.langs,
				tagKeys:      tt.fields.tagKeys,
				tagValues:    tt.fields.tagValues,
				filter:       tt.fields.filter,
				edited:       tt.fields.edited,
				deflang:      tt.fields.deflang,
				locfile:      tt.fields.locfile,
			}
			loc.AddLang(tt.args.w, tt.args.r)
		})
	}
}
