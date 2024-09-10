package component

import (
	"reflect"
	"testing"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
)

func TestAdmin_Render(t *testing.T) {
	type fields struct {
		BaseComponent ct.BaseComponent
		Version       string
		Theme         string
		Module        string
		View          string
		Token         string
		HelpURL       string
		ClientURL     string
		LocalesURL    string
		Labels        cu.SM
		TokenLogin    func(database, token string) bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Default",
			fields: fields{
				BaseComponent: ct.BaseComponent{
					Id:       cu.GetComponentID(),
					EventURL: "/event",
					OnResponse: func(evt ct.ResponseEvent) (re ct.ResponseEvent) {
						return re
					},
					Data: cu.IM{
						"alias": "demo",
					},
				},
				Module:     "database",
				HelpURL:    "https://nervatura.github.io/nervatura/",
				ClientURL:  "/client",
				LocalesURL: "/locales",
			},
			wantErr: false,
		},
		{
			name: "Database result",
			fields: fields{
				BaseComponent: ct.BaseComponent{
					Id:       cu.GetComponentID(),
					EventURL: "/event",
					Data: cu.IM{
						"api_key": "DEMO_API_KEY",
						"alias":   "demo",
						"create_result": []cu.IM{
							{"database": "demo", "message": "Start process", "stamp": "2023-12-22 17:03:26", "state": "create"},
							{"message": "The existing table is dropped...", "stamp": "2023-12-22 17:03:26", "state": "err"},
							{"message": "Creating the tables...", "stamp": "2023-12-22 17:03:26", "state": "create"},
						},
					},
				},
				Module: "database",
			},
			wantErr: false,
		},
		{
			name: "Login",
			fields: fields{
				BaseComponent: ct.BaseComponent{
					Id:       cu.GetComponentID(),
					EventURL: "/event",
					Data: cu.IM{
						"username": "admin",
						"database": "demo",
					},
				},
				Module: "login",
			},
			wantErr: false,
		},
		{
			name: "Password change",
			fields: fields{
				BaseComponent: ct.BaseComponent{
					Id:       cu.GetComponentID(),
					EventURL: "/event",
					Data: cu.IM{
						"username": "admin",
						"database": "demo",
					},
				},
				Module:     "login",
				Token:      "TOKEN0123456789",
				TokenLogin: func(database, token string) bool { return (token != "") },
				View:       "password",
			},
			wantErr: false,
		},
		{
			name: "Report",
			fields: fields{
				BaseComponent: ct.BaseComponent{
					Id:       cu.GetComponentID(),
					EventURL: "/event",
					Data: cu.IM{
						"username": "admin",
						"database": "demo",
						"report_list": []cu.IM{
							{"description": "Accounts Payable and Receivable.", "filename": "csv_custpos_en.json", "installed": true,
								"label": "", "repname": "Payments Due List - CSV output", "reportkey": "csv_custpos_en", "reptype": "csv"},
							{"description": "Recoverable and payable VAT summary grouped by currency.", "filename": "csv_vat_en.json",
								"installed": false, "label": "", "repname": "VAT Summary - CSV output.", "reportkey": "csv_vat_en", "reptype": "csv"},
						},
						"report_list_current_page": int64(2),
					},
				},
				Module:     "login",
				Token:      "TOKEN0123456789",
				TokenLogin: func(database, token string) bool { return (token != "") },
				View:       "report",
			},
			wantErr: false,
		},
		{
			name: "Configuration",
			fields: fields{
				BaseComponent: ct.BaseComponent{
					Id:       cu.GetComponentID(),
					EventURL: "/event",
					Data: cu.IM{
						"username": "admin",
						"database": "demo",
						"env_list": []cu.IM{
							{"envkey": "NT_ALIAS_DEMO", "envvalue": "sqlite://file:data/demo.db?cache=shared&mode=rwc"},
							{"envkey": "NT_CLIENT_CONFIG", "envvalue": "data/client_config_loc.json"},
						},
					},
				},
				Module:     "login",
				Token:      "TOKEN0123456789",
				TokenLogin: func(database, token string) bool { return (token != "") },
				View:       "configuration",
			},
			wantErr: false,
		},
		{
			name: "Locales",
			fields: fields{
				BaseComponent: ct.BaseComponent{
					Id:       cu.GetComponentID(),
					EventURL: "/event",
					Data: cu.IM{
						"locales": cu.IM{
							"locale":     "de",
							"tag_key":    "tag",
							"locfile":    cu.IM{"locales": cu.IM{"de": cu.IM{"tag_key1": "value1"}}},
							"deflang":    cu.IM{"tag_key1": "value1", "tag_key2": "value2"},
							"tag_values": map[string][]string{"tag": {"tag_key1", "tag_key2"}},
							"locales":    []ct.SelectOption{},
							"tag_keys":   []ct.SelectOption{},
						},
					},
				},
				Module: "locales",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adm := &Admin{
				BaseComponent: tt.fields.BaseComponent,
				Version:       tt.fields.Version,
				Theme:         tt.fields.Theme,
				Module:        tt.fields.Module,
				View:          tt.fields.View,
				Token:         tt.fields.Token,
				HelpURL:       tt.fields.HelpURL,
				ClientURL:     tt.fields.ClientURL,
				LocalesURL:    tt.fields.LocalesURL,
				Labels:        tt.fields.Labels,
				TokenLogin:    tt.fields.TokenLogin,
			}
			_, err := adm.Render()
			if (err != nil) != tt.wantErr {
				t.Errorf("Admin.Render() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestAdmin_Validation(t *testing.T) {
	type fields struct {
		BaseComponent ct.BaseComponent
		Version       string
		Theme         string
		Module        string
		View          string
		Token         string
		HelpURL       string
		ClientURL     string
		Labels        cu.SM
		TokenLogin    func(darabase, token string) bool
	}
	type args struct {
		propName  string
		propValue interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{
			name: "base",
			args: args{
				propName:  "id",
				propValue: "BTNID",
			},
			want: "BTNID",
		},
		{
			name: "invalid",
			args: args{
				propName:  "invalid",
				propValue: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adm := &Admin{
				BaseComponent: tt.fields.BaseComponent,
				Version:       tt.fields.Version,
				Theme:         tt.fields.Theme,
				Module:        tt.fields.Module,
				View:          tt.fields.View,
				Token:         tt.fields.Token,
				HelpURL:       tt.fields.HelpURL,
				ClientURL:     tt.fields.ClientURL,
				Labels:        tt.fields.Labels,
				TokenLogin:    tt.fields.TokenLogin,
			}
			if got := adm.Validation(tt.args.propName, tt.args.propValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Admin.Validation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdmin_SetProperty(t *testing.T) {
	type fields struct {
		BaseComponent ct.BaseComponent
		Version       string
		Theme         string
		Module        string
		View          string
		Token         string
		HelpURL       string
		ClientURL     string
		Labels        cu.SM
		TokenLogin    func(darabase, token string) bool
	}
	type args struct {
		propName  string
		propValue interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{
			name: "base",
			args: args{
				propName:  "id",
				propValue: "BTNID",
			},
			want: "BTNID",
		},
		{
			name: "invalid",
			args: args{
				propName:  "invalid",
				propValue: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adm := &Admin{
				BaseComponent: tt.fields.BaseComponent,
				Version:       tt.fields.Version,
				Theme:         tt.fields.Theme,
				Module:        tt.fields.Module,
				View:          tt.fields.View,
				Token:         tt.fields.Token,
				HelpURL:       tt.fields.HelpURL,
				ClientURL:     tt.fields.ClientURL,
				Labels:        tt.fields.Labels,
				TokenLogin:    tt.fields.TokenLogin,
			}
			if got := adm.SetProperty(tt.args.propName, tt.args.propValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Admin.SetProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdmin_msg(t *testing.T) {
	type fields struct {
		BaseComponent ct.BaseComponent
		Version       string
		Theme         string
		Module        string
		View          string
		Token         string
		HelpURL       string
		ClientURL     string
		Labels        cu.SM
		TokenLogin    func(darabase, token string) bool
	}
	type args struct {
		labelID string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "missing",
			args: args{
				labelID: "missing",
			},
			want: "missing",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adm := &Admin{
				BaseComponent: tt.fields.BaseComponent,
				Version:       tt.fields.Version,
				Theme:         tt.fields.Theme,
				Module:        tt.fields.Module,
				View:          tt.fields.View,
				Token:         tt.fields.Token,
				HelpURL:       tt.fields.HelpURL,
				ClientURL:     tt.fields.ClientURL,
				Labels:        tt.fields.Labels,
				TokenLogin:    tt.fields.TokenLogin,
			}
			if got := adm.msg(tt.args.labelID); got != tt.want {
				t.Errorf("Admin.msg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdmin_response(t *testing.T) {
	type fields struct {
		BaseComponent ct.BaseComponent
		Version       string
		Theme         string
		Module        string
		View          string
		Token         string
		HelpURL       string
		ClientURL     string
		LocalesURL    string
		Labels        cu.SM
		TokenLogin    func(darabase, token string) bool
	}
	type args struct {
		evt ct.ResponseEvent
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "api_key",
			args: args{
				evt: ct.ResponseEvent{
					TriggerName: "api_key",
				},
			},
		},
		{
			name: "create_result",
			args: args{
				evt: ct.ResponseEvent{
					TriggerName: "create_result",
				},
			},
		},
		{
			name: "theme",
			fields: fields{
				BaseComponent: ct.BaseComponent{
					OnResponse: func(evt ct.ResponseEvent) (re ct.ResponseEvent) {
						evt.Trigger = &Admin{}
						return evt
					},
				},
				Theme: ct.ThemeDark,
			},
			args: args{
				evt: ct.ResponseEvent{
					TriggerName: "theme",
				},
			},
		},
		{
			name: "main_menu_help",
			fields: fields{
				HelpURL: "/help",
			},
			args: args{
				evt: ct.ResponseEvent{
					TriggerName: "main_menu",
					Value:       "help",
				},
			},
		},
		{
			name: "main_menu_client",
			fields: fields{
				ClientURL: "/client",
			},
			args: args{
				evt: ct.ResponseEvent{
					TriggerName: "main_menu",
					Value:       "client",
				},
			},
		},
		{
			name: "main_menu_locales",
			fields: fields{
				LocalesURL: "/locales",
			},
			args: args{
				evt: ct.ResponseEvent{
					TriggerName: "main_menu",
					Value:       "locales",
				},
			},
		},
		{
			name: "view_menu",
			args: args{
				evt: ct.ResponseEvent{
					TriggerName: "view_menu",
					Value:       "logout",
				},
			},
		},
		{
			name: "create",
			args: args{
				evt: ct.ResponseEvent{
					TriggerName: "create",
				},
			},
		},
		{
			name: "login",
			args: args{
				evt: ct.ResponseEvent{
					TriggerName: "login",
				},
			},
		},
		{
			name: "report_install",
			args: args{
				evt: ct.ResponseEvent{
					TriggerName: "report_install",
					Trigger: &ct.BaseComponent{
						Data: make(map[string]interface{}),
					},
				},
			},
		},
		{
			name: "report_delete",
			args: args{
				evt: ct.ResponseEvent{
					TriggerName: "report_delete",
					Trigger: &ct.BaseComponent{
						Data: make(map[string]interface{}),
					},
				},
			},
		},
		{
			name: "report_list",
			args: args{
				evt: ct.ResponseEvent{
					Name:        ct.TableEventCurrentPage,
					TriggerName: "report_list",
					Trigger: &ct.BaseComponent{
						Data: make(map[string]interface{}),
					},
				},
			},
		},
		{
			name: "password_change",
			args: args{
				evt: ct.ResponseEvent{
					TriggerName: "password_change",
				},
			},
		},
		{
			name: "locales_undo",
			args: args{
				evt: ct.ResponseEvent{
					TriggerName: "locales",
					Name:        LocalesEventUndo,
				},
			},
		},
		{
			name: "locales_save",
			args: args{
				evt: ct.ResponseEvent{
					TriggerName: "locales",
					Name:        AdminEventLocalesSave,
				},
			},
		},
		{
			name: "locales_error",
			args: args{
				evt: ct.ResponseEvent{
					TriggerName: "locales",
					Name:        AdminEventLocalesError,
				},
			},
		},
		{
			name: "locales",
			args: args{
				evt: ct.ResponseEvent{
					TriggerName: "locales",
					Name:        "",
				},
			},
		},
		{
			name: "missing",
			args: args{
				evt: ct.ResponseEvent{
					TriggerName: "missing",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adm := &Admin{
				BaseComponent: tt.fields.BaseComponent,
				Version:       tt.fields.Version,
				Theme:         tt.fields.Theme,
				Module:        tt.fields.Module,
				View:          tt.fields.View,
				Token:         tt.fields.Token,
				HelpURL:       tt.fields.HelpURL,
				ClientURL:     tt.fields.ClientURL,
				LocalesURL:    tt.fields.LocalesURL,
				Labels:        tt.fields.Labels,
				TokenLogin:    tt.fields.TokenLogin,
			}
			adm.response(tt.args.evt)
		})
	}
}

func TestAdmin_OnRequest(t *testing.T) {
	type fields struct {
		BaseComponent ct.BaseComponent
		Version       string
		Theme         string
		Module        string
		View          string
		Token         string
		HelpURL       string
		ClientURL     string
		LocalesURL    string
		Labels        cu.SM
		TokenLogin    func(database, token string) bool
	}
	type args struct {
		te ct.TriggerEvent
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "invalid",
			args: args{
				te: ct.TriggerEvent{},
			},
		},
		{
			name: "valid",
			fields: fields{
				BaseComponent: ct.BaseComponent{
					RequestMap: map[string]ct.ClientComponent{
						"ID12345": &Admin{},
					},
				},
			},
			args: args{
				te: ct.TriggerEvent{
					Id: "ID12345",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adm := &Admin{
				BaseComponent: tt.fields.BaseComponent,
				Version:       tt.fields.Version,
				Theme:         tt.fields.Theme,
				Module:        tt.fields.Module,
				View:          tt.fields.View,
				Token:         tt.fields.Token,
				HelpURL:       tt.fields.HelpURL,
				ClientURL:     tt.fields.ClientURL,
				LocalesURL:    tt.fields.LocalesURL,
				Labels:        tt.fields.Labels,
				TokenLogin:    tt.fields.TokenLogin,
			}
			adm.OnRequest(tt.args.te)
		})
	}
}
