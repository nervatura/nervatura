package component

import (
	"fmt"
	"html/template"
	"sort"
	"strings"

	ct "github.com/nervatura/component/pkg/component"
	ut "github.com/nervatura/component/pkg/util"
)

// [Locale] constants
const (
	ComponentTypeLocale = "locale"

	LocalesEventChange = "change"
	LocalesEventUndo   = "undo"
	LocalesEventSave   = "save"
	LocalesEventError  = "error"
)

var localeDefaultLabel ut.SM = ut.SM{
	"locale_title":         "Translation helper tool",
	"locale_missing":       "Missing values",
	"locale_update":        "Write changes to the json files",
	"locale_undo":          "Discard all changes",
	"locale_add":           "Add a new language",
	"locale_filter":        "Filter rows",
	"locale_lcode":         "Lang.code (e.g. de)",
	"locale_lname":         "Lang.name(e.g.Deutsche)",
	"locale_existing_lang": "Existing language code",
	"locale_tag":           "Tag",
	"locale_key":           "Fieldname",
	"locale_value":         "Value",
}

/*
Creates an translation helper tool

Example component with the following main features:
  - server-side state management
  - [Input], [Select], [Label], [Button] and [Table] components
  - customized table cells: link field and input control
  - data filtering
  - state-bound control visibility
  - dynamic data sources for the [Select] and [Table] controls

For example:

	&Locale{
	  BaseComponent: ct.BaseComponent{
	    Id:           "id_locale_default",
	    EventURL:     "/event",
	    OnResponse:   func(evt ct.ResponseEvent) (re ct.ResponseEvent) {
	      // return parent_component response
	      return evt
	    },
	    RequestValue: parent_component.GetProperty("request_value").(map[string]ut.IM),
	    RequestMap:   parent_component.GetProperty("request_map").(map[string]ct.ClientComponent),
	    Data: ut.IM{
	      "deflang": ut.IM{
	        "key": "en", "en": "English", "address_view": "Address Data",
	      },
	      "locales":  "default",
	      "tag_keys": "address",
	      "tag_values": map[string][]string{
	        "address": { "address_view", "address_country" },
	        "login": { "login_username", "login_password", "login_database" },
	      },
	      "locfile": ut.IM{
	        "locales": ut.IM{
	          "de": ut.IM{ "de": "Deutsche", "key": "de", "login_database": "Datenbank" }
	        },
	      },
	    },
	  },
	  Locales: []ct.SelectOption{
	    {Value: "default", Text: "Default"}, {Value: "de", Text: "Deutsch"},
	  },
	  TagKeys: []ct.SelectOption{
	    {Value: "address", Text: "address"}, {Value: "login", Text: "login"},
	  },
	}
*/
type Locale struct {
	ct.BaseComponent
	// The languages that can be selected from the data source
	Locales []ct.SelectOption `json:"locales"`
	// The groups of localization texts
	TagKeys []ct.SelectOption `json:"tag_keys"`
	// The filter condition of the displayed data
	FilterValue string `json:"filter_value"`
	// Data changed from user input
	Dirty bool `json:"dirty"`
	// Show/hide Add a new language section
	AddItem bool `json:"add_item"`
	// The texts of the labels of the controls
	Labels ut.SM `json:"labels"`
}

/*
Returns all properties of the [Locale]
*/
func (loc *Locale) Properties() ut.IM {
	return ut.MergeIM(
		loc.BaseComponent.Properties(),
		ut.IM{
			"locales":      loc.Locales,
			"tag_keys":     loc.TagKeys,
			"filter_value": loc.FilterValue,
			"dirty":        loc.Dirty,
			"add_item":     loc.AddItem,
			"labels":       loc.Labels,
		})
}

/*
Returns the value of the property of the [Locale] with the specified name.
*/
func (loc *Locale) GetProperty(propName string) interface{} {
	return loc.Properties()[propName]
}

/*
It checks the value given to the property of the [Locale] and always returns a valid value
*/
func (loc *Locale) Validation(propName string, propValue interface{}) interface{} {
	pm := map[string]func() interface{}{
		"labels": func() interface{} {
			value := ut.ToSM(loc.Labels, ut.SM{})
			if smap, valid := propValue.(ut.SM); valid {
				value = ut.MergeSM(value, smap)
			}
			if len(value) == 0 {
				value = localeDefaultLabel
			}
			return value
		},
		"locales": func() interface{} {
			value := loc.Locales
			if lang, valid := propValue.([]ct.SelectOption); valid && len(lang) > 0 {
				value = lang
			}
			if len(value) == 0 {
				lang := ut.ToString(loc.Data["locales"], "default")
				value = []ct.SelectOption{{Value: lang, Text: lang}}
			}
			return value
		},
		"tag_keys": func() interface{} {
			value := loc.TagKeys
			if key, valid := propValue.([]ct.SelectOption); valid && len(key) > 0 {
				value = key
			}
			return value
		},
		"target": func() interface{} {
			loc.SetProperty("id", loc.Id)
			value := ut.ToString(propValue, loc.Id)
			if value != "this" && !strings.HasPrefix(value, "#") {
				value = "#" + value
			}
			return value
		},
	}
	if _, found := pm[propName]; found {
		return pm[propName]()
	}
	if loc.BaseComponent.GetProperty(propName) != nil {
		return loc.BaseComponent.Validation(propName, propValue)
	}
	return propValue
}

/*
Setting a property of the [Locale] value safely. Checks the entered value.
In case of an invalid value, the default value will be set.
*/
func (loc *Locale) SetProperty(propName string, propValue interface{}) interface{} {
	pm := map[string]func() interface{}{
		"locales": func() interface{} {
			loc.Locales = loc.Validation(propName, propValue).([]ct.SelectOption)
			return loc.Locales
		},
		"tag_keys": func() interface{} {
			loc.TagKeys = loc.Validation(propName, propValue).([]ct.SelectOption)
			return loc.TagKeys
		},
		"filter_value": func() interface{} {
			loc.FilterValue = ut.ToString(propValue, "")
			return loc.FilterValue
		},
		"dirty": func() interface{} {
			loc.Dirty = ut.ToBoolean(propValue, false)
			return loc.Dirty
		},
		"add_item": func() interface{} {
			loc.AddItem = ut.ToBoolean(propValue, false)
			return loc.AddItem
		},
		"labels": func() interface{} {
			loc.Labels = loc.Validation(propName, propValue).(ut.SM)
			return loc.Labels
		},
		"target": func() interface{} {
			loc.Target = loc.Validation(propName, propValue).(string)
			return loc.Target
		},
	}
	if _, found := pm[propName]; found {
		return loc.SetRequestValue(propName, pm[propName](), []string{})
	}
	if loc.BaseComponent.GetProperty(propName) != nil {
		return loc.BaseComponent.SetProperty(propName, propValue)
	}
	return propValue
}

func (loc *Locale) response(evt ct.ResponseEvent) (re ct.ResponseEvent) {
	locEvt := ct.ResponseEvent{
		Trigger: loc, TriggerName: loc.Name, Value: evt.Value,
	}
	switch evt.TriggerName {
	case "values":
		return evt

	case "tag_keys", "locales", "undo", "update", "add", "missing", "tag_cell", "value_cell",
		"lang_key", "lang_name", "filter", "add_item":
		evtMap := map[string]func(){
			"tag_keys": func() {
				locEvt.Name = LocalesEventChange
				loc.SetProperty("data", ut.IM{evt.TriggerName: locEvt.Value})
				loc.SetProperty("filter_value", "")
			},
			"locales": func() {
				locEvt.Name = LocalesEventChange
				loc.SetProperty("data", ut.IM{evt.TriggerName: locEvt.Value})
				loc.SetProperty("filter_value", "")
				loc.SetProperty("data", ut.IM{"tag_keys": loc.TagKeys[0].Value})
			},
			"undo": func() {
				locEvt.Name = LocalesEventUndo
			},
			"update": func() {
				locEvt.Name = LocalesEventSave
			},
			"add": func() {
				lang_key := ut.ToString(loc.GetProperty("data").(ut.IM)["lang_key"], "")
				lang_name := ut.ToString(loc.GetProperty("data").(ut.IM)["lang_name"], "")
				locales := loc.Data["locfile"].(ut.IM)["locales"].(ut.IM)
				if _, found := locales[lang_key].(ut.IM); found || lang_key == "en" {
					locEvt.Name = LocalesEventError
					locEvt.Value = loc.msg("locale_existing_lang")
				} else if lang_key == "" || lang_name == "" {
					locEvt.Name = LocalesEventError
					locEvt.Value = loc.msg("locale_missing")
				} else {
					locEvt.Name = LocalesEventChange
					locales[lang_key] = ut.IM{
						"key":    lang_key,
						lang_key: lang_name,
					}
					langs := append(loc.Locales, ct.SelectOption{Value: lang_key, Text: lang_key})
					loc.SetProperty("locales", langs)
					loc.SetProperty("add_item", false)
					loc.SetProperty("data", ut.IM{"lang_key": ""})
					loc.SetProperty("data", ut.IM{"lang_name": ""})
				}
			},
			"missing": func() {
				locEvt.Name = LocalesEventChange
				loc.SetProperty("data", ut.IM{"tag_keys": "missing"})
				loc.SetProperty("filter_value", "")
			},
			"tag_cell": func() {
				locEvt.Name = LocalesEventChange
				loc.SetProperty("data", ut.IM{"tag_keys": locEvt.Value})
				loc.SetProperty("filter_value", "")
			},
			"value_cell": func() {
				locEvt.Name = LocalesEventChange
				key := ut.ToString(evt.Trigger.GetProperty("data").(ut.IM)["key"], "")
				locales := loc.Data["locfile"].(ut.IM)["locales"].(ut.IM)
				lang := ut.ToString(loc.Data["locales"], "")
				if langValues, found := locales[lang].(ut.IM); found {
					langValues[key] = locEvt.Value
				}
				loc.SetProperty("dirty", true)
			},
			"lang_key": func() {
				locEvt.Name = LocalesEventChange
				loc.SetProperty("data", ut.IM{evt.TriggerName: locEvt.Value})
				loc.SetProperty("dirty", true)
			},
			"lang_name": func() {
				locEvt.Name = LocalesEventChange
				loc.SetProperty("data", ut.IM{evt.TriggerName: locEvt.Value})
				loc.SetProperty("dirty", true)
			},
			"filter": func() {
				locEvt.Name = LocalesEventChange
				loc.SetProperty("filter_value", locEvt.Value)
			},
			"add_item": func() {
				loc.SetProperty("add_item", !loc.AddItem)
			},
		}
		evtMap[evt.TriggerName]()

	default:
	}
	if loc.OnResponse != nil {
		return loc.OnResponse(locEvt)
	}
	return locEvt
}

func (loc *Locale) getComponent(name string, data ut.IM) (html template.HTML, err error) {
	ccSel := func(options []ct.SelectOption) *ct.Select {
		sel := &ct.Select{
			BaseComponent: ct.BaseComponent{
				Id: loc.Id + "_" + name, Name: name,
				EventURL:     loc.EventURL,
				Target:       loc.Target,
				OnResponse:   loc.response,
				RequestValue: loc.RequestValue,
				RequestMap:   loc.RequestMap,
			},
			IsNull:  false,
			Options: options,
		}
		sel.SetProperty("value", loc.Data[name])
		return sel
	}
	ccBtn := func(icoKey, label, bstyle string) *ct.Button {
		return &ct.Button{
			BaseComponent: ct.BaseComponent{
				Id: loc.Id + "_" + name, Name: name,
				Style:        ut.SM{"padding": "8px"},
				EventURL:     loc.EventURL,
				Target:       loc.Id,
				OnResponse:   loc.response,
				RequestValue: loc.RequestValue,
				RequestMap:   loc.RequestMap,
			},
			ButtonStyle:    bstyle,
			Label:          loc.msg(label),
			LabelComponent: &ct.Icon{Value: icoKey, Width: 18, Height: 18},
		}
	}
	ccInp := func(label, placeholder, value string) *ct.Input {
		inp := &ct.Input{
			BaseComponent: ct.BaseComponent{
				Id:           loc.Id + "_" + name + "_" + ut.ToString(data["key"], ""),
				Name:         name,
				EventURL:     loc.EventURL,
				Target:       loc.Target,
				OnResponse:   loc.response,
				RequestValue: loc.RequestValue,
				RequestMap:   loc.RequestMap,
				Data:         data,
			},
			Type:        ct.InputTypeString,
			Label:       label,
			Placeholder: placeholder,
			Full:        true,
		}
		inp.SetProperty("value", value)
		return inp
	}
	ccTbl := func(rowKey string, rows []ut.IM, fields []ct.TableField) *ct.Table {
		tbl := &ct.Table{
			BaseComponent: ct.BaseComponent{
				Id: loc.Id + "_" + name, Name: name,
				EventURL:     loc.EventURL,
				OnResponse:   loc.response,
				RequestValue: loc.RequestValue,
				RequestMap:   loc.RequestMap,
			},
			Rows:        rows,
			Fields:      fields,
			Pagination:  ct.PaginationTypeTop,
			PageSize:    10,
			RowKey:      rowKey,
			TableFilter: false,
			AddItem:     false,
		}
		return tbl
	}
	ccMap := map[string]func() ct.ClientComponent{
		"locales": func() ct.ClientComponent {
			return ccSel(loc.Locales)
		},
		"tag_keys": func() ct.ClientComponent {
			keys := loc.TagKeys
			if ut.ToString(loc.Data["locales"], "") != "default" {
				keys = append(keys, ct.SelectOption{Value: "missing", Text: "missing"})
			}
			return ccSel(keys)
		},
		"missing": func() ct.ClientComponent {
			return ccBtn("QuestionCircle", "locale_missing", ct.ButtonStyleDefault)
		},
		"filter": func() ct.ClientComponent {
			return ccInp(loc.msg("locale_filter"), loc.msg("locale_filter"), loc.FilterValue)
		},
		"update": func() ct.ClientComponent {
			return ccBtn("Check", "locale_update", ct.ButtonStylePrimary)
		},
		"undo": func() ct.ClientComponent {
			return ccBtn("Undo", "locale_undo", ct.ButtonStylePrimary)
		},
		"add_item": func() ct.ClientComponent {
			icon := "Plus"
			if loc.AddItem {
				icon = "ArrowUp"
			}
			return ccBtn(icon, "locale_add", ct.ButtonStyleDefault)
		},
		"add": func() ct.ClientComponent {
			return ccBtn("Plus", "locale_add", ct.ButtonStyleDefault)
		},
		"lang_key": func() ct.ClientComponent {
			lang_key := ut.ToString(loc.GetProperty("data").(ut.IM)["lang_key"], "")
			inp := ccInp(loc.msg("locale_lcode"), loc.msg("locale_lcode"), lang_key)
			inp.Full = false
			inp.MaxLength = 5
			inp.Style = ut.SM{"text-transform": "lowercase"}
			return inp
		},
		"lang_name": func() ct.ClientComponent {
			lang_name := ut.ToString(loc.GetProperty("data").(ut.IM)["lang_name"], "")
			return ccInp(loc.msg("locale_lname"), loc.msg("locale_lname"), lang_name)
		},
		"tag_cell": func() ct.ClientComponent {
			return &ct.Label{
				BaseComponent: ct.BaseComponent{
					Id:           loc.Id + "_" + ut.ToString(data["key"], ""),
					Name:         name,
					EventURL:     loc.EventURL,
					Target:       loc.Target,
					Data:         data,
					OnResponse:   loc.response,
					RequestValue: loc.RequestValue,
					RequestMap:   loc.RequestMap,
				},
				Value: ut.ToString(data["tag"], ""),
			}
		},
		"value_cell": func() ct.ClientComponent {
			return ccInp(ut.ToString(data["default"], ""), ut.ToString(data["default"], ""), ut.ToString(data["value"], ""))
		},
		"values": func() ct.ClientComponent {
			toValue := func(lang, key string, locales ut.IM) string {
				if langValues, found := locales[lang].(ut.IM); found {
					return ut.ToString(langValues[key], "")
				}
				return ""
			}

			lang := ut.ToString(loc.Data["locales"], "")
			tag := ut.ToString(loc.Data["tag_keys"], "")
			rows := []ut.IM{}
			deflang := loc.Data["deflang"].(ut.IM)
			locales := loc.Data["locfile"].(ut.IM)["locales"].(ut.IM)
			fields := []ct.TableField{
				{Column: &ct.TableColumn{
					Id:     "tag",
					Header: loc.msg("locale_tag"),
					Cell: func(row ut.IM, col ct.TableColumn, value interface{}, rowIndex int64) template.HTML {
						linkLabel := fmt.Sprintf(
							`<span class="cell-label">%s</span>`, value)
						var link template.HTML
						link, _ = loc.getComponent("tag_cell", row)
						return template.HTML(linkLabel + string(link))
					}}},
				{Name: "key", FieldType: ct.TableFieldTypeString, Label: loc.msg("locale_key")},
			}
			if lang == "default" {
				fields = append(fields, ct.TableField{
					Name: "default", FieldType: ct.TableFieldTypeString, Label: loc.msg("locale_value")})
			} else {
				fields = append(fields,
					ct.TableField{Column: &ct.TableColumn{
						Id:     "value",
						Header: loc.msg("locale_value"),
						Cell: func(row ut.IM, col ct.TableColumn, value interface{}, rowIndex int64) template.HTML {
							input, _ := loc.getComponent("value_cell", row)
							return input
						}}},
				)
			}
			var keys []string
			if loc.FilterValue != "" {
				keys = make([]string, 0)
				for rowKey := range deflang {
					if strings.Contains(rowKey, loc.FilterValue) ||
						strings.Contains(ut.ToString(deflang[rowKey], ""), loc.FilterValue) {
						keys = append(keys, rowKey)
					} else if lang != "default" {
						if strings.Contains(toValue(lang, rowKey, locales), loc.FilterValue) {
							keys = append(keys, rowKey)
						}
					}
				}
			} else if tag == "missing" {
				keys = make([]string, 0)
				for rowKey := range deflang {
					if _, found := locales[lang].(ut.IM)[rowKey]; !found {
						keys = append(keys, rowKey)
					}
				}
			} else {
				keys = loc.Data["tag_values"].(map[string][]string)[tag]
			}
			sort.Strings(keys)
			for _, key := range keys {
				if len(strings.Split(key, "_")) > 1 {
					rows = append(rows, ut.IM{
						"tag":     strings.Split(key, "_")[0],
						"key":     key,
						"default": ut.ToString(deflang[key], ""),
						"value":   toValue(lang, key, locales),
					})
				}
			}
			return ccTbl("key", rows, fields)
		},
	}
	cc := ccMap[name]()
	html, err = cc.Render()
	return html, err
}

func (loc *Locale) msg(labelID string) string {
	if label, found := loc.Labels[labelID]; found {
		return label
	}
	return labelID
}

/*
Based on the values, it will generate the html code of the [Locale] or return with an error message.
*/
func (loc *Locale) Render() (html template.HTML, err error) {
	loc.InitProps(loc)

	funcMap := map[string]any{
		"styleMap": func() bool {
			return len(loc.Style) > 0
		},
		"customClass": func() string {
			return strings.Join(loc.Class, " ")
		},
		"localeComponent": func(name string) (html template.HTML, err error) {
			return loc.getComponent(name, ut.IM{})
		},
		"lang": func() string {
			return ut.ToString(loc.Data["locales"], "")
		},
	}
	tpl := `<div id="{{ .Id }}" name="{{ .Name }}" class="row full {{ customClass }}"
	{{ if styleMap }} style="{{ range $key, $value := .Style }}{{ $key }}:{{ $value }};{{ end }}"{{ end }}
	><div class="row full section container-small">
	<div class="row full" >
	<div class="cell mobile" >
	<div class="cell mobile">
	<div class="cell padding-small" >{{ localeComponent "locales" }}</div>
	<div class="cell padding-small" >{{ localeComponent "tag_keys" }}</div>
	{{ if ne lang "default" }}<div class="cell padding-small" >{{ localeComponent "missing" }}</div>{{ end }}
	</div>
	<div class="cell padding-small mobile" >{{ localeComponent "filter" }}</div>
	</div>
	<div class="cell mobile" >
	<div class="right">
	{{ if eq .Dirty true }}
	<div class="cell padding-small" >{{ localeComponent "update" }}</div>
	<div class="cell padding-small" >{{ localeComponent "undo" }}</div>
	{{ end }}
	<div class="cell padding-small" >{{ localeComponent "add_item" }}</div>
	</div>
	</div>
	</div>
	{{ if eq .AddItem true }}
	<div class="row" >
	<div class="cell padding-small mobile" >{{ localeComponent "lang_key" }}</div>
	<div class="cell padding-small mobile" >{{ localeComponent "lang_name" }}</div>
	<div class="cell padding-small" >{{ localeComponent "add" }}</div>
	</div>
	{{ end }}
	</div>
	<div class="container section-small" >
	{{ localeComponent "values" }}
	</div>
	</div>`

	return ut.TemplateBuilder("locales", tpl, funcMap, loc)
}

func localeLocfile() ut.IM {
	return ut.IM{
		"locales": ut.IM{
			"de": ut.IM{
				"de":             "Deutsche",
				"key":            "de",
				"login_database": "Datenbank",
				"login_password": "Passwort",
				"login_username": "Nutzername",
			},
			"jp": ut.IM{
				"jp":             "日本語",
				"key":            "jp",
				"login_database": "データベース",
				"login_password": "パスワード",
				"login_username": "ユーザー名",
			},
		},
	}
}

var testLocaleResponse func(evt ct.ResponseEvent) (re ct.ResponseEvent) = func(evt ct.ResponseEvent) (re ct.ResponseEvent) {
	switch evt.Name {
	case LocalesEventError:
		re = ct.ResponseEvent{
			Trigger: &ct.Toast{
				Type:  ct.ToastTypeError,
				Value: ut.ToString(evt.Value, ""),
			},
			TriggerName: evt.TriggerName,
			Name:        evt.Name,
			Header: ut.SM{
				ct.HeaderRetarget: "#toast-msg",
				ct.HeaderReswap:   ct.SwapInnerHTML,
			},
		}
		return re

	case LocalesEventSave:
		evt.Trigger.SetProperty("dirty", false)

	case LocalesEventUndo:
		evt.Trigger.SetProperty("data",
			ut.IM{"locfile": localeLocfile(),
				"locales": "default", "tag_keys": "address", "lang_key": "", "lang_name": "",
			})
		evt.Trigger.SetProperty("locales", []ct.SelectOption{
			{Value: "default", Text: "Default"},
			{Value: "de", Text: "Deutsch"},
			{Value: "jp", Text: "Japanese"},
		})
		evt.Trigger.SetProperty("filter_value", "")
		evt.Trigger.SetProperty("add_item", false)
		evt.Trigger.SetProperty("dirty", false)
	}
	return evt
}

func testLocaleData() ut.IM {
	return ut.IM{
		"deflang": ut.IM{
			"key":               "en",
			"en":                "English",
			"address_view":      "Address Data",
			"address_country":   "Country",
			"address_state":     "State",
			"address_zipcode":   "Zipcode",
			"address_city":      "City",
			"address_street":    "Street",
			"address_notes":     "Comment",
			"login_username":    "Username",
			"login_password":    "Password",
			"login_database":    "Database",
			"login_lang":        "Language",
			"login_login":       "Login",
			"login_server":      "Server URL",
			"login_engine_err":  "Invalid database type!",
			"login_version_err": "Invalid service version!",
		},
		"locales":  "default",
		"tag_keys": "address",
		"tag_values": map[string][]string{
			"address": {
				"address_view", "address_country", "address_state", "address_zipcode",
				"address_city", "address_street", "address_notes",
			},
			"login": {
				"login_username", "login_password", "login_database", "login_lang",
				"login_login", "login_server", "login_engine_err", "login_version_err",
			},
		},
		"locfile": localeLocfile(),
	}
}

// [Locale] test and demo data
func TestLocale(cc ct.ClientComponent) []ct.TestComponent {
	id := ut.ToString(cc.GetProperty("id"), "")
	eventURL := ut.ToString(cc.GetProperty("event_url"), "")
	requestValue := cc.GetProperty("request_value").(map[string]ut.IM)
	requestMap := cc.GetProperty("request_map").(map[string]ct.ClientComponent)
	return []ct.TestComponent{
		{
			Label:         "Default",
			ComponentType: ComponentTypeLocale,
			Component: &Locale{
				BaseComponent: ct.BaseComponent{
					Id:           id + "_locale_default",
					EventURL:     eventURL,
					OnResponse:   testLocaleResponse,
					RequestValue: requestValue,
					RequestMap:   requestMap,
					Data:         testLocaleData(),
				},
				Locales: []ct.SelectOption{
					{Value: "default", Text: "Default"},
					{Value: "de", Text: "Deutsch"},
					{Value: "jp", Text: "Japanese"},
				},
				TagKeys: []ct.SelectOption{
					{Value: "address", Text: "address"},
					{Value: "login", Text: "login"},
				},
			}},
		{
			Label:         "Input",
			ComponentType: ComponentTypeLocale,
			Component: &Locale{
				BaseComponent: ct.BaseComponent{
					Id:           id + "_locale_input",
					EventURL:     eventURL,
					OnResponse:   testLocaleResponse,
					RequestValue: requestValue,
					RequestMap:   requestMap,
					Data: ut.MergeIM(testLocaleData(), ut.IM{
						"locales":  "de",
						"tag_keys": "login",
					}),
				},
				Locales: []ct.SelectOption{
					{Value: "default", Text: "Default"},
					{Value: "de", Text: "Deutsch"},
					{Value: "jp", Text: "Japanese"},
				},
				TagKeys: []ct.SelectOption{
					{Value: "address", Text: "address"},
					{Value: "login", Text: "login"},
				},
			}},
	}
}
