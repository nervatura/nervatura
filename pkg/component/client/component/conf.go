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

func ClientMenu(labels cu.SM, config cu.IM) ct.MenuBar {
	//theme := cu.ToString(config["theme"], "light")
	helpURL := st.DocsClientPath //+ "/browser",
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

func getLocales(clientLang []string) (locales []ct.SelectOption) {
	locales = []ct.SelectOption{}
	for _, loc := range clientLang {
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

func ClientLogin(labels cu.SM, config cu.IM) ct.Login {
	theme := cu.ToString(config["theme"], "light")
	version := cu.ToString(config["version"], "1.0.0")
	lang := cu.ToString(config["lang"], "en")
	login := ct.Login{
		Locales:      getLocales(st.ClientLang),
		AuthButtons:  []ct.LoginAuthButton{},
		Version:      version,
		Theme:        theme,
		Labels:       labels,
		Lang:         lang,
		HideDatabase: false,
		HidePassword: false,
		ShowHelp:     true,
		HelpURL:      st.DocsPath,
	}
	return login
}

func ClientSideBar(moduleKey string, labels cu.SM, data cu.IM) ct.SideBar {
	itemMap := map[string]func() []ct.SideBarItem{
		"search": func() []ct.SideBarItem {
			return searchSideBar(labels, data)
		},
		"browser": func() []ct.SideBarItem {
			return searchSideBar(labels, data)
		},
		"customer": func() []ct.SideBarItem {
			return customerSideBar(labels, data)
		},
		"product": func() []ct.SideBarItem {
			return productSideBar(labels, data)
		},
		"tool": func() []ct.SideBarItem {
			return toolSideBar(labels, data)
		},
		"project": func() []ct.SideBarItem {
			return projectSideBar(labels, data)
		},
		"employee": func() []ct.SideBarItem {
			return employeeSideBar(labels, data)
		},
		"trans": func() []ct.SideBarItem {
			return transSideBar(labels, data)
		},
		"setting": func() []ct.SideBarItem {
			return settingSideBar(labels, data)
		},
		"place": func() []ct.SideBarItem {
			return placeSideBar(labels, data)
		},
	}
	sb := ct.SideBar{
		Items: []ct.SideBarItem{},
	}
	if item, found := itemMap[moduleKey]; found {
		sb.Items = item()
	}
	return sb
}

func ClientEditor(editorKey, viewName string, labels cu.SM, editorData cu.IM) ct.Editor {
	edi := ct.Editor{
		Title:  cu.ToString(editorData["editor_title"], labels[editorKey+"_title"]),
		Icon:   cu.ToString(editorData["editor_icon"], ""),
		Views:  moduleEditorView(editorKey, labels, editorData),
		Rows:   moduleEditorRow(editorKey, viewName, labels, editorData),
		Tables: moduleEditorTable(editorKey, viewName, labels, editorData),
	}
	edi.SetProperty("view", viewName)
	return edi
}

func moduleEditorView(editorKey string, labels cu.SM, data cu.IM) []ct.EditorView {
	viewMap := map[string]func() []ct.EditorView{
		"customer": func() []ct.EditorView {
			return customerEditorView(labels, data)
		},
		"product": func() []ct.EditorView {
			return productEditorView(labels, data)
		},
		"tool": func() []ct.EditorView {
			return toolEditorView(labels, data)
		},
		"project": func() []ct.EditorView {
			return projectEditorView(labels, data)
		},
		"employee": func() []ct.EditorView {
			return employeeEditorView(labels, data)
		},
		"trans": func() []ct.EditorView {
			return transEditorView(labels, data)
		},
		"setting": func() []ct.EditorView {
			return settingEditorView(labels, data)
		},
		"place": func() []ct.EditorView {
			return placeEditorView(labels, data)
		},
	}
	if view, found := viewMap[editorKey]; found {
		return view()
	}
	return []ct.EditorView{}
}

func moduleEditorRow(editorKey, viewName string, labels cu.SM, data cu.IM) []ct.Row {
	rowMap := map[string]func() []ct.Row{
		"customer": func() []ct.Row {
			return customerRow(viewName, labels, data)
		},
		"product": func() []ct.Row {
			return productRow(viewName, labels, data)
		},
		"tool": func() []ct.Row {
			return toolRow(viewName, labels, data)
		},
		"project": func() []ct.Row {
			return projectRow(viewName, labels, data)
		},
		"employee": func() []ct.Row {
			return employeeRow(viewName, labels, data)
		},
		"trans": func() []ct.Row {
			return transRow(viewName, labels, data)
		},
		"setting": func() []ct.Row {
			return settingRow(viewName, labels, data)
		},
		"place": func() []ct.Row {
			return placeRow(viewName, labels, data)
		},
	}
	if row, found := rowMap[editorKey]; found {
		return row()
	}
	return []ct.Row{}
}

func moduleEditorTable(editorKey, viewName string, labels cu.SM, data cu.IM) []ct.Table {
	tblMap := map[string]func() []ct.Table{
		"customer": func() []ct.Table {
			return customerTable(viewName, labels, data)
		},
		"product": func() []ct.Table {
			return productTable(viewName, labels, data)
		},
		"tool": func() []ct.Table {
			return toolTable(viewName, labels, data)
		},
		"project": func() []ct.Table {
			return projectTable(viewName, labels, data)
		},
		"employee": func() []ct.Table {
			return employeeTable(viewName, labels, data)
		},
		"trans": func() []ct.Table {
			return transTable(viewName, labels, data)
		},
		"setting": func() []ct.Table {
			return settingTable(viewName, labels, data)
		},
		"place": func() []ct.Table {
			return placeTable(viewName, labels, data)
		},
	}
	if tbl, found := tblMap[editorKey]; found {
		return tbl()
	}
	return []ct.Table{}
}

func ClientForm(editorKey, formKey string, labels cu.SM, data cu.IM) (form ct.Form) {
	frmMap := map[string]func() ct.Form{
		"customer": func() ct.Form {
			return customerForm(formKey, labels, data)
		},
		"product": func() ct.Form {
			return productForm(formKey, labels, data)
		},
		"tool": func() ct.Form {
			return toolForm(formKey, labels, data)
		},
		"project": func() ct.Form {
			return projectForm(formKey, labels, data)
		},
		"employee": func() ct.Form {
			return employeeForm(formKey, labels, data)
		},
		"trans": func() ct.Form {
			return transForm(formKey, labels, data)
		},
		"setting": func() ct.Form {
			return settingForm(formKey, labels, data)
		},
		"place": func() ct.Form {
			return placeForm(formKey, labels, data)
		},
	}
	if frm, found := frmMap[editorKey]; found {
		return frm()
	}
	return ct.Form{}
}

func ClientModalForm(formKey string, labels cu.SM, data cu.IM) (form ct.Form) {
	mMap := map[string]func() ct.Form{
		"info": func() ct.Form {
			return modalInfoMessage(labels, data)
		},
		"warning": func() ct.Form {
			return modalWarningMessage(labels, data)
		},
		"input_string": func() ct.Form {
			return modalInputString(labels, data)
		},
		"select": func() ct.Form {
			return modalSelect(labels, data)
		},
		"report": func() ct.Form {
			return modalReport(labels, data)
		},
		"selector": func() ct.Form {
			return modalSelector(labels, data)
		},
		"config_field": func() ct.Form {
			return modalConfigField(labels, data)
		},
		"trans_create": func() ct.Form {
			return modalTransCreate(labels, data)
		},
	}
	if frm, found := mMap[formKey]; found {
		return frm()
	}
	return ct.Form{}
}

func ClientBrowser(viewName string, labels cu.SM, searchData cu.IM) ct.Browser {
	sConf := SearchViewConfig(viewName, labels)
	rows := cu.ToIMA(searchData["rows"], []cu.IM{})
	var sessionID string
	config := cu.ToIM(searchData["config"], cu.IM{})
	userConfig := cu.ToIM(searchData["user_config"], cu.IM{})
	if ticket, found := config["ticket"].(ct.Ticket); found {
		sessionID = ticket.SessionID
	}
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
		ExportURL:   fmt.Sprintf(st.ClientPath+"/session/export/browser/%s", sessionID),
		HelpURL:     st.DocsClientPath, //+ "/browser",
		Download:    fmt.Sprintf("%s.csv", viewName),
		ReadOnly:    sConf.ReadOnly,
	}
	bro.SetProperty("visible_columns", sConf.VisibleColumns)
	bro.SetProperty("hide_filters", sConf.HideFilters)
	bro.SetProperty("filters", sConf.Filters)
	return bro
}

func ClientSearch(viewName string, labels cu.SM, searchData cu.IM) ct.Search {
	sConf := SearchViewConfig(viewName, labels)
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
