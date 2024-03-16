package component

import (
	"reflect"
	"testing"

	fm "github.com/nervatura/component/component/atom"
	bc "github.com/nervatura/component/component/base"
)

func TestLocales_Render(t *testing.T) {
	type fields struct {
		BaseComponent bc.BaseComponent
		Locales       []fm.SelectOption
		TagKeys       []fm.SelectOption
		FilterValue   string
		Dirty         bool
		AddItem       bool
		Labels        bc.SM
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "default",
			fields: fields{
				BaseComponent: bc.BaseComponent{
					Data: bc.IM{
						"deflang":    bc.IM{},
						"locales":    "client",
						"tag_keys":   "address",
						"tag_values": map[string][]string{},
						"locfile":    bc.IM{"locales": bc.IM{}},
					},
				},
				Locales: []fm.SelectOption{},
				TagKeys: []fm.SelectOption{{Value: "address", Text: "address"}},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loc := &Locales{
				BaseComponent: tt.fields.BaseComponent,
				Locales:       tt.fields.Locales,
				TagKeys:       tt.fields.TagKeys,
				FilterValue:   tt.fields.FilterValue,
				Dirty:         tt.fields.Dirty,
				AddItem:       tt.fields.AddItem,
				Labels:        tt.fields.Labels,
			}
			_, err := loc.Render()
			if (err != nil) != tt.wantErr {
				t.Errorf("Locales.Render() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestLocales_Validation(t *testing.T) {
	type fields struct {
		BaseComponent bc.BaseComponent
		Locales       []fm.SelectOption
		TagKeys       []fm.SelectOption
		FilterValue   string
		Dirty         bool
		AddItem       bool
		Labels        bc.SM
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
				propValue: []fm.SelectOption{{Value: "client", Text: "client"}},
			},
			want: []fm.SelectOption{{Value: "client", Text: "client"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loc := &Locales{
				BaseComponent: tt.fields.BaseComponent,
				Locales:       tt.fields.Locales,
				TagKeys:       tt.fields.TagKeys,
				FilterValue:   tt.fields.FilterValue,
				Dirty:         tt.fields.Dirty,
				AddItem:       tt.fields.AddItem,
				Labels:        tt.fields.Labels,
			}
			if got := loc.Validation(tt.args.propName, tt.args.propValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Locales.Validation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocales_SetProperty(t *testing.T) {
	type fields struct {
		BaseComponent bc.BaseComponent
		Locales       []fm.SelectOption
		TagKeys       []fm.SelectOption
		FilterValue   string
		Dirty         bool
		AddItem       bool
		Labels        bc.SM
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
			loc := &Locales{
				BaseComponent: tt.fields.BaseComponent,
				Locales:       tt.fields.Locales,
				TagKeys:       tt.fields.TagKeys,
				FilterValue:   tt.fields.FilterValue,
				Dirty:         tt.fields.Dirty,
				AddItem:       tt.fields.AddItem,
				Labels:        tt.fields.Labels,
			}
			if got := loc.SetProperty(tt.args.propName, tt.args.propValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Locales.SetProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocales_msg(t *testing.T) {
	type fields struct {
		BaseComponent bc.BaseComponent
		Locales       []fm.SelectOption
		TagKeys       []fm.SelectOption
		FilterValue   string
		Dirty         bool
		AddItem       bool
		Labels        bc.SM
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
			loc := &Locales{
				BaseComponent: tt.fields.BaseComponent,
				Locales:       tt.fields.Locales,
				TagKeys:       tt.fields.TagKeys,
				FilterValue:   tt.fields.FilterValue,
				Dirty:         tt.fields.Dirty,
				AddItem:       tt.fields.AddItem,
				Labels:        tt.fields.Labels,
			}
			if got := loc.msg(tt.args.labelID); got != tt.want {
				t.Errorf("Locales.msg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocales_response(t *testing.T) {
	type fields struct {
		BaseComponent bc.BaseComponent
		Locales       []fm.SelectOption
		TagKeys       []fm.SelectOption
		FilterValue   string
		Dirty         bool
		AddItem       bool
		Labels        bc.SM
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
			name: "values",
			args: args{
				evt: bc.ResponseEvent{
					TriggerName: "values",
				},
			},
		},
		{
			name: "tag_keys",
			args: args{
				evt: bc.ResponseEvent{
					TriggerName: "tag_keys",
				},
			},
		},
		{
			name: "locales",
			fields: fields{
				TagKeys: []fm.SelectOption{{Value: "address", Text: "address"}},
			},
			args: args{
				evt: bc.ResponseEvent{
					TriggerName: "locales",
				},
			},
		},
		{
			name: "undo",
			args: args{
				evt: bc.ResponseEvent{
					TriggerName: "undo",
				},
			},
		},
		{
			name: "update",
			args: args{
				evt: bc.ResponseEvent{
					TriggerName: "update",
				},
			},
		},
		{
			name: "add_locales_missing",
			fields: fields{
				BaseComponent: bc.BaseComponent{
					Data: bc.IM{
						"locfile": bc.IM{"locales": bc.IM{}},
					},
					OnResponse: func(evt bc.ResponseEvent) (re bc.ResponseEvent) {
						evt.Trigger = &Admin{}
						return evt
					},
				},
			},
			args: args{
				evt: bc.ResponseEvent{
					TriggerName: "add",
				},
			},
		},
		{
			name: "add_existing_lang",
			fields: fields{
				BaseComponent: bc.BaseComponent{
					Data: bc.IM{
						"lang_key": "en",
						"locfile":  bc.IM{"locales": bc.IM{}},
					},
				},
			},
			args: args{
				evt: bc.ResponseEvent{
					TriggerName: "add",
				},
			},
		},
		{
			name: "add",
			fields: fields{
				BaseComponent: bc.BaseComponent{
					Data: bc.IM{
						"lang_key":  "de",
						"lang_name": "de",
						"locfile":   bc.IM{"locales": bc.IM{}},
					},
				},
			},
			args: args{
				evt: bc.ResponseEvent{
					TriggerName: "add",
				},
			},
		},
		{
			name: "missing",
			args: args{
				evt: bc.ResponseEvent{
					TriggerName: "missing",
				},
			},
		},
		{
			name: "tag_cell",
			args: args{
				evt: bc.ResponseEvent{
					TriggerName: "tag_cell",
				},
			},
		},
		{
			name: "value_cell",
			fields: fields{
				BaseComponent: bc.BaseComponent{
					Data: bc.IM{
						"locales": "de",
						"locfile": bc.IM{"locales": bc.IM{"de": bc.IM{}}},
					},
				},
			},
			args: args{
				evt: bc.ResponseEvent{
					TriggerName: "value_cell",
					Trigger: &fm.Input{
						BaseComponent: bc.BaseComponent{
							Data: bc.IM{
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
				evt: bc.ResponseEvent{
					TriggerName: "lang_key",
				},
			},
		},
		{
			name: "lang_name",
			args: args{
				evt: bc.ResponseEvent{
					TriggerName: "lang_name",
				},
			},
		},
		{
			name: "filter",
			args: args{
				evt: bc.ResponseEvent{
					TriggerName: "filter",
				},
			},
		},
		{
			name: "add_item",
			args: args{
				evt: bc.ResponseEvent{
					TriggerName: "add_item",
				},
			},
		},
		{
			name: "default",
			args: args{
				evt: bc.ResponseEvent{
					TriggerName: "default",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loc := &Locales{
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

func TestLocales_getComponent(t *testing.T) {
	type fields struct {
		BaseComponent bc.BaseComponent
		Locales       []fm.SelectOption
		TagKeys       []fm.SelectOption
		FilterValue   string
		Dirty         bool
		AddItem       bool
		Labels        bc.SM
	}
	type args struct {
		name string
		data bc.IM
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
				BaseComponent: bc.BaseComponent{
					Data: bc.IM{
						"locales": "de",
					},
				},
			},
			args: args{
				name: "tag_keys",
				data: bc.IM{},
			},
			wantErr: false,
		},
		{
			name: "missing",
			args: args{
				name: "missing",
				data: bc.IM{},
			},
			wantErr: false,
		},
		{
			name: "update",
			args: args{
				name: "update",
				data: bc.IM{},
			},
			wantErr: false,
		},
		{
			name: "undo",
			args: args{
				name: "undo",
				data: bc.IM{},
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
				data: bc.IM{},
			},
			wantErr: false,
		},
		{
			name: "add",
			args: args{
				name: "add",
				data: bc.IM{},
			},
			wantErr: false,
		},
		{
			name: "lang_key",
			fields: fields{
				BaseComponent: bc.BaseComponent{
					Data: bc.IM{
						"lang_key": "key",
					},
				},
			},
			args: args{
				name: "lang_key",
				data: bc.IM{},
			},
			wantErr: false,
		},
		{
			name: "lang_name",
			fields: fields{
				BaseComponent: bc.BaseComponent{
					Data: bc.IM{
						"lang_name": "name",
					},
				},
			},
			args: args{
				name: "lang_name",
				data: bc.IM{},
			},
			wantErr: false,
		},
		{
			name: "tag_cell",
			args: args{
				name: "tag_cell",
				data: bc.IM{},
			},
			wantErr: false,
		},
		{
			name: "value_cell",
			args: args{
				name: "value_cell",
				data: bc.IM{},
			},
			wantErr: false,
		},
		{
			name: "values",
			fields: fields{
				BaseComponent: bc.BaseComponent{
					Data: bc.IM{
						"locales":    "de",
						"tag_keys":   "tag",
						"locfile":    bc.IM{"locales": bc.IM{"de": bc.IM{"tag_key1": "value1"}}},
						"deflang":    bc.IM{"tag_key1": "value1", "tag_key2": "value2"},
						"tag_values": map[string][]string{"tag": {"tag_key1", "tag_key2"}},
					},
				},
			},
			args: args{
				name: "values",
				data: bc.IM{},
			},
			wantErr: false,
		},
		{
			name: "values_client",
			fields: fields{
				BaseComponent: bc.BaseComponent{
					Data: bc.IM{
						"locales":    "client",
						"tag_keys":   "tag",
						"locfile":    bc.IM{"locales": bc.IM{"de": bc.IM{"tag_key1": "value1"}}},
						"deflang":    bc.IM{"tag_key1": "value1", "tag_key2": "value2"},
						"tag_values": map[string][]string{"tag": {"tag_key1", "tag_key2"}},
					},
				},
			},
			args: args{
				name: "values",
				data: bc.IM{},
			},
			wantErr: false,
		},
		{
			name: "values_missing",
			fields: fields{
				BaseComponent: bc.BaseComponent{
					Data: bc.IM{
						"locales":    "de",
						"tag_keys":   "missing",
						"locfile":    bc.IM{"locales": bc.IM{"de": bc.IM{"tag_key1": "value1"}}},
						"deflang":    bc.IM{"tag_key1": "value1", "tag_key2": "value2"},
						"tag_values": map[string][]string{"tag": {"tag_key1", "tag_key2"}},
					},
				},
			},
			args: args{
				name: "values",
				data: bc.IM{},
			},
			wantErr: false,
		},
		{
			name: "values_filter",
			fields: fields{
				BaseComponent: bc.BaseComponent{
					Data: bc.IM{
						"locales":    "de",
						"tag_keys":   "tag",
						"locfile":    bc.IM{"locales": bc.IM{"de": bc.IM{"tag_key1": "value1", "tag_key2": "key1"}}},
						"deflang":    bc.IM{"tag_key1": "value1", "tag_key2": "value2"},
						"tag_values": map[string][]string{"tag": {"tag_key1", "tag_key2"}},
					},
				},
				FilterValue: "key1",
			},
			args: args{
				name: "values",
				data: bc.IM{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loc := &Locales{
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
				t.Errorf("Locales.getComponent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
