package component

import (
	"reflect"
	"testing"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
)

func TestTestLocale(t *testing.T) {
	for _, tt := range TestLocale(&ct.BaseComponent{EventURL: "/demo"}) {
		t.Run(tt.Label, func(t *testing.T) {
			tt.Component.Render()
		})
	}
	testLocaleResponse(ct.ResponseEvent{Name: LocalesEventError, Trigger: &Locale{}})
	testLocaleResponse(ct.ResponseEvent{Name: LocalesEventSave, Trigger: &Locale{}})
	testLocaleResponse(ct.ResponseEvent{Name: LocalesEventUndo, Trigger: &Locale{}})
}

func TestLocale_Validation(t *testing.T) {
	type fields struct {
		BaseComponent ct.BaseComponent
		Locales       []ct.SelectOption
		TagKeys       []ct.SelectOption
		FilterValue   string
		Dirty         bool
		AddItem       bool
		Labels        cu.SM
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
		{
			name: "locales",
			args: args{
				propName:  "locales",
				propValue: []ct.SelectOption{},
			},
			want: []ct.SelectOption{{Value: "default", Text: "default"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loc := &Locale{
				BaseComponent: tt.fields.BaseComponent,
				Locales:       tt.fields.Locales,
				TagKeys:       tt.fields.TagKeys,
				FilterValue:   tt.fields.FilterValue,
				Dirty:         tt.fields.Dirty,
				AddItem:       tt.fields.AddItem,
				Labels:        tt.fields.Labels,
			}
			if got := loc.Validation(tt.args.propName, tt.args.propValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Locale.Validation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocale_SetProperty(t *testing.T) {
	type fields struct {
		BaseComponent ct.BaseComponent
		Locales       []ct.SelectOption
		TagKeys       []ct.SelectOption
		FilterValue   string
		Dirty         bool
		AddItem       bool
		Labels        cu.SM
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
			loc := &Locale{
				BaseComponent: tt.fields.BaseComponent,
				Locales:       tt.fields.Locales,
				TagKeys:       tt.fields.TagKeys,
				FilterValue:   tt.fields.FilterValue,
				Dirty:         tt.fields.Dirty,
				AddItem:       tt.fields.AddItem,
				Labels:        tt.fields.Labels,
			}
			if got := loc.SetProperty(tt.args.propName, tt.args.propValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Locale.SetProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocale_response(t *testing.T) {
	type fields struct {
		BaseComponent ct.BaseComponent
		Locales       []ct.SelectOption
		TagKeys       []ct.SelectOption
		FilterValue   string
		Dirty         bool
		AddItem       bool
		Labels        cu.SM
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
			name: "values",
			args: args{
				evt: ct.ResponseEvent{
					TriggerName: "values",
				},
			},
		},
		{
			name: "tag_keys",
			args: args{
				evt: ct.ResponseEvent{
					TriggerName: "tag_keys",
				},
			},
		},
		{
			name: "locales",
			fields: fields{
				TagKeys: []ct.SelectOption{{Value: "address", Text: "address"}},
			},
			args: args{
				evt: ct.ResponseEvent{
					TriggerName: "locales",
				},
			},
		},
		{
			name: "undo",
			args: args{
				evt: ct.ResponseEvent{
					TriggerName: "undo",
				},
			},
		},
		{
			name: "update",
			args: args{
				evt: ct.ResponseEvent{
					TriggerName: "update",
				},
			},
		},
		{
			name: "add_locales_missing",
			fields: fields{
				BaseComponent: ct.BaseComponent{
					Data: cu.IM{
						"locfile": cu.IM{"locales": cu.IM{}},
					},
					OnResponse: func(evt ct.ResponseEvent) (re ct.ResponseEvent) {
						evt.Trigger = &ct.BaseComponent{}
						return evt
					},
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					TriggerName: "add",
				},
			},
		},
		{
			name: "add_existing_lang",
			fields: fields{
				BaseComponent: ct.BaseComponent{
					Data: cu.IM{
						"lang_key": "en",
						"locfile":  cu.IM{"locales": cu.IM{}},
					},
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					TriggerName: "add",
				},
			},
		},
		{
			name: "add",
			fields: fields{
				BaseComponent: ct.BaseComponent{
					Data: cu.IM{
						"lang_key":  "de",
						"lang_name": "de",
						"locfile":   cu.IM{"locales": cu.IM{}},
					},
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					TriggerName: "add",
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
		{
			name: "tag_cell",
			args: args{
				evt: ct.ResponseEvent{
					TriggerName: "tag_cell",
				},
			},
		},
		{
			name: "value_cell",
			fields: fields{
				BaseComponent: ct.BaseComponent{
					Data: cu.IM{
						"locales": "de",
						"locfile": cu.IM{"locales": cu.IM{"de": cu.IM{}}},
					},
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					TriggerName: "value_cell",
					Trigger: &ct.Input{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"key": "abc",
							},
						},
					},
				},
			},
		},
		{
			name: "lang_key",
			args: args{
				evt: ct.ResponseEvent{
					TriggerName: "lang_key",
				},
			},
		},
		{
			name: "lang_name",
			args: args{
				evt: ct.ResponseEvent{
					TriggerName: "lang_name",
				},
			},
		},
		{
			name: "filter",
			args: args{
				evt: ct.ResponseEvent{
					TriggerName: "filter",
				},
			},
		},
		{
			name: "add_item",
			args: args{
				evt: ct.ResponseEvent{
					TriggerName: "add_item",
				},
			},
		},
		{
			name: "default",
			args: args{
				evt: ct.ResponseEvent{
					TriggerName: "default",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loc := &Locale{
				BaseComponent: tt.fields.BaseComponent,
				Locales:       tt.fields.Locales,
				TagKeys:       tt.fields.TagKeys,
				FilterValue:   tt.fields.FilterValue,
				Dirty:         tt.fields.Dirty,
				AddItem:       tt.fields.AddItem,
				Labels:        tt.fields.Labels,
			}
			loc.response(tt.args.evt)
		})
	}
}

func TestLocale_getComponent(t *testing.T) {
	type fields struct {
		BaseComponent ct.BaseComponent
		Locales       []ct.SelectOption
		TagKeys       []ct.SelectOption
		FilterValue   string
		Dirty         bool
		AddItem       bool
		Labels        cu.SM
	}
	type args struct {
		name string
		data cu.IM
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "tag_keys",
			fields: fields{
				BaseComponent: ct.BaseComponent{
					Data: cu.IM{
						"locales": "de",
					},
				},
			},
			args: args{
				name: "tag_keys",
				data: cu.IM{},
			},
			wantErr: false,
		},
		{
			name: "missing",
			args: args{
				name: "missing",
				data: cu.IM{},
			},
			wantErr: false,
		},
		{
			name: "update",
			args: args{
				name: "update",
				data: cu.IM{},
			},
			wantErr: false,
		},
		{
			name: "undo",
			args: args{
				name: "undo",
				data: cu.IM{},
			},
			wantErr: false,
		},
		{
			name: "add_item",
			fields: fields{
				AddItem: true,
			},
			args: args{
				name: "add_item",
				data: cu.IM{},
			},
			wantErr: false,
		},
		{
			name: "add",
			args: args{
				name: "add",
				data: cu.IM{},
			},
			wantErr: false,
		},
		{
			name: "lang_key",
			fields: fields{
				BaseComponent: ct.BaseComponent{
					Data: cu.IM{
						"lang_key": "key",
					},
				},
			},
			args: args{
				name: "lang_key",
				data: cu.IM{},
			},
			wantErr: false,
		},
		{
			name: "lang_name",
			fields: fields{
				BaseComponent: ct.BaseComponent{
					Data: cu.IM{
						"lang_name": "name",
					},
				},
			},
			args: args{
				name: "lang_name",
				data: cu.IM{},
			},
			wantErr: false,
		},
		{
			name: "tag_cell",
			args: args{
				name: "tag_cell",
				data: cu.IM{},
			},
			wantErr: false,
		},
		{
			name: "value_cell",
			args: args{
				name: "value_cell",
				data: cu.IM{},
			},
			wantErr: false,
		},
		{
			name: "values",
			fields: fields{
				BaseComponent: ct.BaseComponent{
					Data: cu.IM{
						"locales":    "de",
						"tag_keys":   "tag",
						"locfile":    cu.IM{"locales": cu.IM{"de": cu.IM{"tag_key1": "value1"}}},
						"deflang":    cu.IM{"tag_key1": "value1", "tag_key2": "value2"},
						"tag_values": map[string][]string{"tag": {"tag_key1", "tag_key2"}},
					},
				},
			},
			args: args{
				name: "values",
				data: cu.IM{},
			},
			wantErr: false,
		},
		{
			name: "values_client",
			fields: fields{
				BaseComponent: ct.BaseComponent{
					Data: cu.IM{
						"locales":    "client",
						"tag_keys":   "tag",
						"locfile":    cu.IM{"locales": cu.IM{"de": cu.IM{"tag_key1": "value1"}}},
						"deflang":    cu.IM{"tag_key1": "value1", "tag_key2": "value2"},
						"tag_values": map[string][]string{"tag": {"tag_key1", "tag_key2"}},
					},
				},
			},
			args: args{
				name: "values",
				data: cu.IM{},
			},
			wantErr: false,
		},
		{
			name: "values_missing",
			fields: fields{
				BaseComponent: ct.BaseComponent{
					Data: cu.IM{
						"locales":    "de",
						"tag_keys":   "missing",
						"locfile":    cu.IM{"locales": cu.IM{"de": cu.IM{"tag_key1": "value1"}}},
						"deflang":    cu.IM{"tag_key1": "value1", "tag_key2": "value2"},
						"tag_values": map[string][]string{"tag": {"tag_key1", "tag_key2"}},
					},
				},
			},
			args: args{
				name: "values",
				data: cu.IM{},
			},
			wantErr: false,
		},
		{
			name: "values_filter",
			fields: fields{
				BaseComponent: ct.BaseComponent{
					Data: cu.IM{
						"locales":    "de",
						"tag_keys":   "tag",
						"locfile":    cu.IM{"locales": cu.IM{"de": cu.IM{"tag_key1": "value1", "tag_key2": "key1"}}},
						"deflang":    cu.IM{"tag_key1": "value1", "tag_key2": "value2"},
						"tag_values": map[string][]string{"tag": {"tag_key1", "tag_key2"}},
					},
				},
				FilterValue: "key1",
			},
			args: args{
				name: "values",
				data: cu.IM{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loc := &Locale{
				BaseComponent: tt.fields.BaseComponent,
				Locales:       tt.fields.Locales,
				TagKeys:       tt.fields.TagKeys,
				FilterValue:   tt.fields.FilterValue,
				Dirty:         tt.fields.Dirty,
				AddItem:       tt.fields.AddItem,
				Labels:        tt.fields.Labels,
			}
			_, err := loc.getComponent(tt.args.name, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Locale.getComponent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
