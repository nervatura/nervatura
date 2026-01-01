package component

import (
	"fmt"
	"slices"
	"strings"
	"time"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
	st "github.com/nervatura/nervatura/v6/pkg/static"
)

// EditorInterface is an interface that defines the methods for a Nervatura Client UI editor components.
type EditorInterface interface {
	Frame(labels cu.SM, data cu.IM) (title, icon string)
	SideBar(labels cu.SM, data cu.IM) (items []ct.SideBarItem)
	View(labels cu.SM, data cu.IM) (views []ct.EditorView)
	Row(view string, labels cu.SM, data cu.IM) (rows []ct.Row)
	Table(view string, labels cu.SM, data cu.IM) []ct.Table
	Form(formKey string, labels cu.SM, data cu.IM) (form ct.Form)
}

// SearchInterface is an interface that defines the methods for a Nervatura Client UI search components.
type SearchInterface interface {
	SideBar(labels cu.SM, data cu.IM) (items []ct.SideBarItem)
	SideGroups(labels cu.SM) []md.SideGroup
	View(view string, labels cu.SM) md.SearchView
	Query(key string, params cu.IM) md.Query
	Filter(view string, filter ct.BrowserFilter, queryFilters []string) []string
}

var editorMap = map[string]EditorInterface{
	"customer": &CustomerEditor{},
	"employee": &EmployeeEditor{},
	"place":    &PlaceEditor{},
	"product":  &ProductEditor{},
	"project":  &ProjectEditor{},
	"tool":     &ToolEditor{},
	"setting":  &SettingEditor{},
	"trans":    &TransEditor{},
	"shipping": &ShippingEditor{},
	"rate":     &RateEditor{},
}

var modalMap = map[string]func(labels cu.SM, data cu.IM) ct.Form{
	"info": func(labels cu.SM, data cu.IM) ct.Form {
		return modalInfoMessage(labels, data)
	},
	"warning": func(labels cu.SM, data cu.IM) ct.Form {
		return modalWarningMessage(labels, data)
	},
	"input_string": func(labels cu.SM, data cu.IM) ct.Form {
		return modalInputString(labels, data)
	},
	"select": func(labels cu.SM, data cu.IM) ct.Form {
		return modalSelect(labels, data)
	},
	"report": func(labels cu.SM, data cu.IM) ct.Form {
		return modalReport(labels, data)
	},
	"selector": func(labels cu.SM, data cu.IM) ct.Form {
		return modalSelector(labels, data)
	},
	"config_field": func(labels cu.SM, data cu.IM) ct.Form {
		return modalConfigField(labels, data)
	},
	"trans_create": func(labels cu.SM, data cu.IM) ct.Form {
		return modalTransCreate(labels, data)
	},
}

type ClientComponent struct {
	languages     []string
	helpURL       string
	clientHelpURL string
	exportURL     string
	SearchConfig  *SearchConfig
	editorMap     map[string]EditorInterface
	modalMap      map[string]func(labels cu.SM, data cu.IM) ct.Form
}

func NewClientComponent() *ClientComponent {
	return &ClientComponent{
		languages:     st.ClientLang,
		helpURL:       st.DocsPath,
		clientHelpURL: st.DocsClientPath + "#search",
		exportURL:     st.ClientPath + "/session/export/browser/%s",
		SearchConfig:  &SearchConfig{},
		editorMap:     editorMap,
		modalMap:      modalMap,
	}
}

func (cc *ClientComponent) Labels(lang string) cu.SM {
	return ut.GetLangMessages(lang)
}

func (cc *ClientComponent) Menu(labels cu.SM, config cu.IM) ct.MenuBar {
	//theme := cu.ToString(config["theme"], "light")
	helpURL := cc.clientHelpURL //+ "/browser",
	hideExit := cu.ToBoolean(config["login_disabled"], false)
	mnu := ct.MenuBar{
		Items: []ct.MenuBarItem{
			//{Value: "theme", Label: labels["theme_"+ct.ClientIcoMap[theme][0]], Icon: ct.ClientIcoMap[theme][1]},
			{Value: "search", Label: labels["mnu_search"], Icon: ct.IconSearch},
			{Value: "bookmark", Label: labels["mnu_bookmark"], Icon: ct.IconStar},
			{Value: "setting", Label: labels["mnu_setting"], Icon: ct.IconCog},
			{Value: "help", Label: labels["mnu_help"], Icon: ct.IconQuestionCircle, ItemURL: helpURL},
		},
		LabelMenu: labels["mnu_menu"],
		LabelHide: labels["mnu_hide"],
		SideBar:   true,
	}
	if !hideExit {
		mnu.Items = append(mnu.Items, ct.MenuBarItem{
			Value: "logout", Label: labels["mnu_logout"], Icon: ct.IconExit,
		})
	}
	return mnu
}

func (cc *ClientComponent) getLocales() (locales []ct.SelectOption) {
	locales = []ct.SelectOption{}
	for _, loc := range cc.languages {
		lang := strings.Split(loc, ",")
		if len(lang) > 1 {
			locales = append(locales, ct.SelectOption{
				Value: lang[0],
				Text:  lang[1],
			})
		}
	}
	return locales
}

func (cc *ClientComponent) Login(labels cu.SM, config cu.IM) ct.Login {
	theme := cu.ToString(config["theme"], "light")
	version := cu.ToString(config["version"], "1.0.0")
	lang := cu.ToString(config["lang"], "en")
	login := ct.Login{
		Locales:      cc.getLocales(),
		AuthButtons:  []ct.LoginAuthButton{},
		Version:      version,
		Theme:        theme,
		Labels:       labels,
		Lang:         lang,
		HideDatabase: false,
		HidePassword: false,
		ShowHelp:     true,
		HelpURL:      cc.helpURL,
	}
	return login
}

func (cc *ClientComponent) SideBar(moduleKey string, labels cu.SM, data cu.IM) ct.SideBar {
	sb := ct.SideBar{
		Items: []ct.SideBarItem{},
	}
	if mb, found := cc.editorMap[moduleKey]; found {
		sb.Items = mb.SideBar(labels, data)
	}
	if slices.Contains([]string{"search", "browser"}, moduleKey) {
		sb.Items = cc.SearchConfig.SideBar(labels, data)
	}
	return sb
}

func (cc *ClientComponent) Search(viewName string, labels cu.SM, searchData cu.IM) ct.Search {
	var sessionID string
	config := cu.ToIM(searchData["config"], cu.IM{})
	if ticket, found := config["ticket"].(ct.Ticket); found {
		sessionID = ticket.SessionID
	}

	sConf := cc.SearchConfig.View(viewName, labels, sessionID)
	rows := cu.ToIMA(searchData["rows"], []cu.IM{})
	search := ct.Search{
		Fields:            sConf.Fields,
		Title:             sConf.Title,
		FilterPlaceholder: sConf.FilterPlaceholder,
		AutoFocus:         true,
		Full:              true,
		Rows:              rows,
	}
	return search
}

func (cc *ClientComponent) Browser(viewName string, labels cu.SM, searchData cu.IM) ct.Browser {
	var sessionID string
	config := cu.ToIM(searchData["config"], cu.IM{})
	if ticket, found := config["ticket"].(ct.Ticket); found {
		sessionID = ticket.SessionID
	}
	sConf := cc.SearchConfig.View(viewName, labels, sessionID)
	rows := cu.ToIMA(searchData["rows"], []cu.IM{})
	userConfig := cu.ToIM(searchData["user_config"], cu.IM{})

	bro := ct.Browser{
		Table: ct.Table{
			TableFilter:       true,
			HidePaginatonSize: false,
			Fields:            sConf.Fields,
			Rows:              rows,
			LabelAdd:          sConf.LabelAdd,
			AddItem:           (sConf.LabelAdd != ""),
			PageSize:          cu.ToInteger(userConfig["pagination"], cu.ToInteger(st.DefaultPagination, 10)),
			RowKey:            "id",
		},
		Title:       sConf.Title,
		View:        viewName,
		ExportLimit: 65000,
		Labels:      labels,
		ExportURL:   fmt.Sprintf(cc.exportURL, sessionID),
		HelpURL:     cc.clientHelpURL, //+ "/browser",
		Download:    fmt.Sprintf("%s.csv", viewName),
		ReadOnly:    sConf.ReadOnly,
	}
	bro.SetProperty("visible_columns", sConf.VisibleColumns)
	bro.SetProperty("hide_filters", sConf.HideFilters)
	bro.SetProperty("filters", sConf.Filters)
	return bro
}

func (cc *ClientComponent) Editor(editorKey, viewName string, labels cu.SM, editorData cu.IM) ct.Editor {
	edi := ct.Editor{
		Title:  cu.ToString(editorData["editor_title"], labels[editorKey+"_title"]),
		Icon:   cu.ToString(editorData["editor_icon"], ""),
		Views:  []ct.EditorView{},
		Rows:   []ct.Row{},
		Tables: []ct.Table{},
	}
	if em, found := cc.editorMap[editorKey]; found {
		edi.Title, edi.Icon = em.Frame(labels, editorData)
		edi.Views = em.View(labels, editorData)
		edi.Rows = em.Row(viewName, labels, editorData)
		edi.Tables = em.Table(viewName, labels, editorData)
	}
	edi.SetProperty("view", viewName)
	return edi
}

func (cc *ClientComponent) Form(editorKey, formKey string, labels cu.SM, data cu.IM) (form ct.Form) {
	if ef, found := cc.editorMap[editorKey]; found {
		return ef.Form(formKey, labels, data)
	}
	return ct.Form{}
}

func (cc *ClientComponent) Modal(formKey string, labels cu.SM, data cu.IM) (form ct.Form) {
	if frm, found := cc.modalMap[formKey]; found {
		return frm(labels, data)
	}
	return ct.Form{}
}

func moduleEditorView(editorKey string, labels cu.SM, data cu.IM) []ct.EditorView {
	if ev, found := editorMap[editorKey]; found {
		return ev.View(labels, data)
	}
	return []ct.EditorView{}
}

func moduleEditorRow(editorKey, viewName string, labels cu.SM, data cu.IM) []ct.Row {
	if er, found := editorMap[editorKey]; found {
		return er.Row(viewName, labels, data)
	}
	return []ct.Row{}
}

func moduleEditorTable(editorKey, viewName string, labels cu.SM, data cu.IM) []ct.Table {
	if et, found := editorMap[editorKey]; found {
		return et.Table(viewName, labels, data)
	}
	return []ct.Table{}
}

func DefaultMapValue(ftype string) interface{} {
	defvalue := cu.IM{
		md.FieldTypeNumber.String():   0,
		md.FieldTypeInteger.String():  0,
		md.FieldTypeDateTime.String(): time.Now().Format("2006-01-02T15:04"),
		md.FieldTypeDate.String():     time.Now().Format("2006-01-02"),
		md.FieldTypeBool.String():     false,
		md.FieldTypeURL.String():      "https://",
	}
	if value, found := defvalue[ftype]; found {
		return value
	}
	return ""
}

func mapTableRows(mapData cu.IM, configMap []cu.IM) (tableRows []cu.IM) {
	typeMap := func(fieldType string) string {
		tm := cu.SM{
			md.FieldTypeBool.String():          "bool",
			md.FieldTypeInteger.String():       "integer",
			md.FieldTypeNumber.String():        "number",
			md.FieldTypeDate.String():          "date",
			md.FieldTypeDateTime.String():      "datetime",
			md.FieldTypeCustomer.String():      "link",
			md.FieldTypeEmployee.String():      "link",
			md.FieldTypePlace.String():         "link",
			md.FieldTypeProduct.String():       "link",
			md.FieldTypeProject.String():       "link",
			md.FieldTypeTool.String():          "link",
			md.FieldTypeTransItem.String():     "link",
			md.FieldTypeTransMovement.String(): "link",
			md.FieldTypeTransPayment.String():  "link",
			md.FieldTypeURL.String():           "link",
		}
		if v, ok := tm[fieldType]; ok {
			return v
		}
		return "string"
	}
	tableRows = []cu.IM{}
	for field, value := range mapData {
		row := cu.IM{
			"id":         field,
			"field_name": field, "description": field, "value": value,
			"value_meta": "string", "field_type": "FIELD_STRING",
		}
		if idx := slices.IndexFunc(configMap, func(c cu.IM) bool {
			return cu.ToString(c["field_name"], "") == field
		}); idx > int(-1) {
			row["value_meta"] = typeMap(cu.ToString(configMap[idx]["field_type"], ""))
			row["description"] = cu.ToString(configMap[idx]["description"], "")
			row["field_type"] = cu.ToString(configMap[idx]["field_type"], "")
			if cu.ToString(configMap[idx]["field_type"], "") == "FIELD_ENUM" {
				valueOptions := []ct.SelectOption{}
				tags := ut.ToStringArray(configMap[idx]["tags"])
				for _, tag := range tags {
					valueOptions = append(valueOptions, ct.SelectOption{
						Value: tag, Text: tag,
					})
				}
				row["value_options"] = valueOptions
			}
		}
		tableRows = append(tableRows, row)
	}
	ut.SortIMData(tableRows, "field_name")
	return tableRows
}

func fromConfig(key string, configData []cu.IM) (options []ct.SelectOption) {
	options = []ct.SelectOption{}
	for _, config := range configData {
		if cu.ToString(config["config_code"], "") == key {
			options = append(options, ct.SelectOption{
				Value: cu.ToString(config["config_key"], ""),
				Text:  cu.ToString(config["config_value"], ""),
			})
		}
	}
	return options
}
