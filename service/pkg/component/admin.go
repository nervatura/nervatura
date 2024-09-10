package component

import (
	"fmt"
	"strings"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	ut "github.com/nervatura/nervatura/service/pkg/utils"
)

const (
	AdminEventModule        = "module"
	AdminEventTheme         = "theme"
	AdminEventChange        = "change"
	AdminEventCreate        = "create"
	AdminEventLogin         = "login"
	AdminEventPassword      = "password"
	AdminEventReportInstall = "report_install"
	AdminEventReportDelete  = "report_delete"
	AdminEventLocalesUndo   = "undo"
	AdminEventLocalesSave   = "save"
	AdminEventLocalesError  = "error"
)

var adminDefaultLabel cu.SM = cu.SM{
	"admin_title":                   ut.GetMessage("admin_title"),
	"admin_login":                   ut.GetMessage("admin_login"),
	"admin_database":                ut.GetMessage("admin_database"),
	"admin_client":                  ut.GetMessage("admin_client"),
	"admin_locales":                 ut.GetMessage("admin_locales"),
	"admin_help":                    ut.GetMessage("admin_help"),
	"admin_api_key":                 ut.GetMessage("admin_api_key"),
	"admin_alias":                   ut.GetMessage("admin_alias"),
	"admin_demo":                    ut.GetMessage("admin_demo"),
	"admin_true":                    ut.GetMessage("admin_true"),
	"admin_false":                   ut.GetMessage("admin_false"),
	"admin_create":                  ut.GetMessage("admin_create"),
	"admin_create_result_state":     ut.GetMessage("admin_create_result_state"),
	"admin_create_result_stamp":     ut.GetMessage("admin_create_result_stamp"),
	"admin_create_result_message":   ut.GetMessage("admin_create_result_message"),
	"admin_create_result_section":   ut.GetMessage("admin_create_result_section"),
	"admin_create_result_datatype":  ut.GetMessage("admin_create_result_datatype"),
	"admin_username":                ut.GetMessage("admin_username"),
	"admin_password":                ut.GetMessage("admin_password"),
	"admin_report":                  ut.GetMessage("admin_report"),
	"admin_configuration":           ut.GetMessage("admin_configuration"),
	"admin_logout":                  ut.GetMessage("admin_logout"),
	"admin_confirm":                 ut.GetMessage("admin_confirm"),
	"admin_password_change":         ut.GetMessage("admin_password_change"),
	"admin_report_list_reportkey":   ut.GetMessage("admin_report_list_reportkey"),
	"admin_report_list_installed":   ut.GetMessage("admin_report_list_installed"),
	"admin_report_list_repname":     ut.GetMessage("admin_report_list_repname"),
	"admin_report_list_description": ut.GetMessage("admin_report_list_description"),
	"admin_report_list_reptype":     ut.GetMessage("admin_report_list_reptype"),
	"admin_report_list_filename":    ut.GetMessage("admin_report_list_filename"),
	"admin_report_list_label":       ut.GetMessage("admin_report_list_label"),
	"admin_env_list_key":            ut.GetMessage("admin_env_list_key"),
	"admin_env_list_value":          ut.GetMessage("admin_env_list_value"),
	"locales_title":                 ut.GetMessage("admin_locales_title"),
	"locales_missing":               ut.GetMessage("locales_missing"),
	"locales_update":                ut.GetMessage("locales_update"),
	"locales_undo":                  ut.GetMessage("locales_undo"),
	"locales_add":                   ut.GetMessage("locales_add"),
	"locales_filter":                ut.GetMessage("locales_filter"),
	"locales_lcode":                 ut.GetMessage("locales_lcode"),
	"locales_lname":                 ut.GetMessage("locales_lname"),
	"locales_existing_lang":         ut.GetMessage("error_existing_lang"),
	"locales_tag":                   ut.GetMessage("locales_tag"),
	"locales_key":                   ut.GetMessage("locales_key"),
	"locales_value":                 ut.GetMessage("locales_value"),
}

var adminIcoMap map[string][]string = map[string][]string{
	ct.ThemeDark: {ct.ThemeLight, "Sun"}, ct.ThemeLight: {ct.ThemeDark, "Moon"},
}

type Admin struct {
	ct.BaseComponent
	Version    string                            `json:"version"`
	Theme      string                            `json:"theme"`
	Module     string                            `json:"module"`
	View       string                            `json:"view"`
	Token      string                            `json:"token"`
	HelpURL    string                            `json:"help_url"`
	ClientURL  string                            `json:"client_url"`
	LocalesURL string                            `json:"locales_url"`
	Labels     cu.SM                             `json:"labels"`
	TokenLogin func(database, token string) bool `json:"-"` // Token validation
}

func (adm *Admin) Properties() cu.IM {
	return cu.MergeIM(
		adm.BaseComponent.Properties(),
		cu.IM{
			"version":     adm.Version,
			"theme":       adm.Theme,
			"module":      adm.Module,
			"view":        adm.View,
			"token":       adm.Token,
			"help_url":    adm.HelpURL,
			"client_url":  adm.ClientURL,
			"locales_url": adm.LocalesURL,
			"labels":      adm.Labels,
		})
}

func (adm *Admin) GetProperty(propName string) interface{} {
	return adm.Properties()[propName]
}

func (adm *Admin) Validation(propName string, propValue interface{}) interface{} {
	pm := map[string]func() interface{}{
		"theme": func() interface{} {
			return adm.CheckEnumValue(cu.ToString(propValue, ""), ct.ThemeLight, ct.Theme)
		},
		"labels": func() interface{} {
			value := cu.SetSMValue(adm.Labels, "", "")
			if smap, valid := propValue.(cu.SM); valid {
				value = cu.MergeSM(value, smap)
			}
			if len(value) == 0 {
				value = adminDefaultLabel
			}
			return value
		},
		"target": func() interface{} {
			adm.SetProperty("id", adm.Id)
			value := cu.ToString(propValue, adm.Id)
			if value != "this" && !strings.HasPrefix(value, "#") {
				value = "#" + value
			}
			return value
		},
	}
	if _, found := pm[propName]; found {
		return pm[propName]()
	}
	if adm.BaseComponent.GetProperty(propName) != nil {
		return adm.BaseComponent.Validation(propName, propValue)
	}
	return propValue
}

func (adm *Admin) SetProperty(propName string, propValue interface{}) interface{} {
	pm := map[string]func() interface{}{
		"version": func() interface{} {
			adm.Version = cu.ToString(propValue, "1.0.0")
			return adm.Version
		},
		"theme": func() interface{} {
			adm.Theme = adm.Validation(propName, propValue).(string)
			return adm.Theme
		},
		"module": func() interface{} {
			adm.Module = cu.ToString(propValue, "database")
			return adm.Module
		},
		"view": func() interface{} {
			adm.View = cu.ToString(propValue, "password")
			return adm.View
		},
		"token": func() interface{} {
			adm.Token = cu.ToString(propValue, "")
			return adm.Token
		},
		"help_url": func() interface{} {
			adm.HelpURL = cu.ToString(propValue, "")
			return adm.HelpURL
		},
		"client_url": func() interface{} {
			adm.ClientURL = cu.ToString(propValue, "")
			return adm.ClientURL
		},
		"locales_url": func() interface{} {
			adm.LocalesURL = cu.ToString(propValue, "")
			return adm.LocalesURL
		},
		"labels": func() interface{} {
			adm.Labels = adm.Validation(propName, propValue).(cu.SM)
			return adm.Labels
		},
		"target": func() interface{} {
			adm.Target = adm.Validation(propName, propValue).(string)
			return adm.Target
		},
	}
	if _, found := pm[propName]; found {
		return adm.SetRequestValue(propName, pm[propName](), []string{})
	}
	if adm.BaseComponent.GetProperty(propName) != nil {
		return adm.BaseComponent.SetProperty(propName, propValue)
	}
	return propValue
}

func (adm *Admin) OnRequest(te ct.TriggerEvent) (re ct.ResponseEvent) {
	if cc, found := adm.RequestMap[te.Id]; found {
		return cc.OnRequest(te)
	}
	re = ct.ResponseEvent{
		Trigger: &ct.Toast{
			Type:  ct.ToastTypeError,
			Value: fmt.Sprintf("Invalid parameter: %s", te.Id),
		},
		TriggerName: te.Name,
		Name:        te.Name,
		Header: cu.SM{
			ct.HeaderRetarget: "#toast-msg",
			ct.HeaderReswap:   "innerHTML",
		},
	}
	return re
}

func (adm *Admin) response(evt ct.ResponseEvent) (re ct.ResponseEvent) {
	admEvt := ct.ResponseEvent{
		Trigger: adm, TriggerName: adm.Name, Value: evt.Value,
	}
	switch evt.TriggerName {

	case "create_result", "env_list":
		return evt

	case "locales":
		switch evt.Name {
		case LocalesEventUndo:
			admEvt.Trigger = evt.Trigger
			admEvt.Name = AdminEventLocalesUndo
		case LocalesEventSave:
			admEvt.Trigger = evt.Trigger
			admEvt.Name = AdminEventLocalesSave
		case LocalesEventError:
			admEvt.Trigger = evt.Trigger
			admEvt.Name = AdminEventLocalesError
		default:
			return evt
		}

	case "report_list":
		if evt.Name == ct.TableEventCurrentPage {
			adm.SetProperty("data", cu.MergeIM(adm.Data, cu.IM{"report_list_current_page": evt.Value}))
		}
		return evt

	case "api_key", "alias", "demo", "username", "password", "database", "confirm", "report_key":
		admEvt.Name = AdminEventChange
		adm.SetProperty("data", cu.IM{evt.TriggerName: admEvt.Value})

	case "theme":
		admEvt.Name = AdminEventTheme
		adm.SetProperty("theme", adminIcoMap[adm.Theme][0])

	case "main_menu":
		admEvt.Name = AdminEventModule
		adm.SetProperty("module", admEvt.Value)
		if admEvt.Value == "help" && adm.HelpURL != "" {
			admEvt.Header = cu.MergeSM(admEvt.Header,
				cu.SM{ct.HeaderRedirect: adm.HelpURL})
		}
		if admEvt.Value == "client" && adm.ClientURL != "" {
			admEvt.Header = cu.MergeSM(admEvt.Header,
				cu.SM{ct.HeaderRedirect: adm.ClientURL})
		}

	case "view_menu":
		admEvt.Name = AdminEventChange
		adm.SetProperty("view", admEvt.Value)
		if admEvt.Value == "logout" {
			adm.SetProperty("token", "")
		}

	case "create":
		admEvt.Name = AdminEventCreate

	case "login":
		admEvt.Name = AdminEventLogin

	case "report_install", "report_delete":
		reportkey := cu.ToString(evt.Trigger.GetProperty("data").(cu.IM)["reportkey"], "")
		admEvt.Name = evt.TriggerName
		admEvt.Value = reportkey

	case "password_change":
		admEvt.Name = AdminEventPassword

	default:
	}
	if adm.OnResponse != nil {
		return adm.OnResponse(admEvt)
	}
	return admEvt
}

func (adm *Admin) getComponent(name string, data cu.IM) (res string, err error) {
	ccLbl := func() *ct.Label {
		return &ct.Label{
			Value: adm.msg(name),
		}
	}
	ccSel := func(options []ct.SelectOption) *ct.Select {
		return &ct.Select{
			BaseComponent: ct.BaseComponent{
				Id: adm.Id + "_" + name, Name: name,
				EventURL:     adm.EventURL,
				Target:       adm.Target,
				OnResponse:   adm.response,
				RequestValue: adm.RequestValue,
				RequestMap:   adm.RequestMap,
			},
			Label:   adm.msg("admin_" + name),
			IsNull:  false,
			Value:   cu.ToString(adm.Data[name], ""),
			Options: options,
		}
	}
	ccInp := func(itype string) *ct.Input {
		return &ct.Input{
			BaseComponent: ct.BaseComponent{
				Id: adm.Id + "_" + name, Name: name,
				EventURL:     adm.EventURL,
				Target:       adm.Target,
				Swap:         ct.SwapOuterHTML,
				OnResponse:   adm.response,
				RequestValue: adm.RequestValue,
				RequestMap:   adm.RequestMap,
			},
			Type:  itype,
			Label: adm.msg("admin_" + name),
			Value: cu.ToString(adm.Data[name], ""),
			Full:  true,
		}
	}
	ccBtn := func(btnType, label string, disabled bool) *ct.Button {
		return &ct.Button{
			BaseComponent: ct.BaseComponent{
				Id: adm.Id + "_" + name, Name: name,
				EventURL:     adm.EventURL,
				Target:       adm.Target,
				OnResponse:   adm.response,
				RequestValue: adm.RequestValue,
				RequestMap:   adm.RequestMap,
			},
			ButtonStyle: btnType,
			Label:       label,
			Disabled:    disabled,
		}
	}
	ccMenu := func(items []ct.MenuBarItem, value, class string) *ct.MenuBar {
		return &ct.MenuBar{
			BaseComponent: ct.BaseComponent{
				Id: adm.Id + "_" + name, Name: name,
				EventURL:     adm.EventURL,
				Target:       adm.Target,
				OnResponse:   adm.response,
				RequestValue: adm.RequestValue,
				RequestMap:   adm.RequestMap,
				Class:        []string{class},
			},
			Items:   items,
			Value:   value,
			SideBar: false,
		}
	}
	ccTbl := func(rowKey string, rows []cu.IM, fields []ct.TableField, currentPage int64) *ct.Table {
		tbl := &ct.Table{
			BaseComponent: ct.BaseComponent{
				Id: adm.Id + "_" + name, Name: name,
				EventURL:     adm.EventURL,
				OnResponse:   adm.response,
				RequestValue: adm.RequestValue,
				RequestMap:   adm.RequestMap,
			},
			Rows:        rows,
			Fields:      fields,
			Pagination:  ct.PaginationTypeTop,
			PageSize:    10,
			RowKey:      rowKey,
			TableFilter: true,
			AddItem:     false,
		}
		if currentPage > 0 {
			tbl.CurrentPage = currentPage
		}
		return tbl
	}
	ccMap := map[string]func() ct.ClientComponent{
		"main_menu": func() ct.ClientComponent {
			return ccMenu(
				[]ct.MenuBarItem{
					{Value: "database", Label: adm.msg("admin_database"), Icon: "Database"},
					{Value: "login", Label: adm.msg("admin_login"), Icon: "Edit"},
					{Value: "client", Label: adm.msg("admin_client"), Icon: "Globe"},
					{Value: "locales", Label: adm.msg("admin_locales"), Icon: "User"},
					{Value: "help", Label: adm.msg("admin_help"), Icon: "QuestionCircle"},
				},
				cu.ToString(adm.GetProperty("module"), ""), "border-top")
		},
		"view_menu": func() ct.ClientComponent {
			return ccMenu(
				[]ct.MenuBarItem{
					{Value: "password", Label: adm.msg("admin_password"), Icon: "Key"},
					{Value: "report", Label: adm.msg("admin_report"), Icon: "ChartBar"},
					{Value: "configuration", Label: adm.msg("admin_configuration"), Icon: "Cog"},
					{Value: "logout", Label: adm.msg("admin_logout"), Icon: "Exit"},
				},
				cu.ToString(adm.GetProperty("view"), ""), "border-top")
		},
		"admin_api_key": func() ct.ClientComponent {
			return ccLbl()
		},
		"api_key": func() ct.ClientComponent {
			return ccInp(ct.InputTypeString)
		},
		"admin_alias": func() ct.ClientComponent {
			return ccLbl()
		},
		"alias": func() ct.ClientComponent {
			return ccInp(ct.InputTypeString)
		},
		"admin_demo": func() ct.ClientComponent {
			return ccLbl()
		},
		"demo": func() ct.ClientComponent {
			return ccSel([]ct.SelectOption{
				{Value: "true", Text: adm.msg("admin_true")},
				{Value: "false", Text: adm.msg("admin_false")},
			})
		},
		"create": func() ct.ClientComponent {
			disabled := (cu.ToString(adm.Data["alias"], "") == "") || (cu.ToString(adm.Data["api_key"], "") == "")
			return ccBtn(ct.ButtonStylePrimary, adm.msg("admin_"+name), disabled)
		},
		"theme": func() ct.ClientComponent {
			themeBtn := ccBtn(ct.ButtonStylePrimary, "", false)
			themeBtn.Style = cu.SM{"padding": "4px"}
			themeBtn.LabelComponent = &ct.Icon{Value: adminIcoMap[adm.Theme][1], Width: 18, Height: 18}
			return themeBtn
		},
		"create_result": func() ct.ClientComponent {
			fields := []ct.TableField{
				{Column: &ct.TableColumn{
					Id:        "state",
					Header:    adm.msg("admin_create_result_state"),
					CellStyle: cu.SM{"text-align": "center"},
					Cell: func(row cu.IM, col ct.TableColumn, value interface{}) string {
						icoKey := "InfoCircle"
						color := "orange"
						if value == "err" {
							icoKey = "ExclamationTriangle"
							color = "red"
						}
						res, _ := (&ct.Icon{Value: icoKey, Color: color}).Render()
						return fmt.Sprintf(
							`<span class="cell-label">%s</span>%s`, col.Header, res)
					}}},
				{Name: "stamp", FieldType: ct.TableFieldTypeTime, Label: adm.msg("admin_create_result_stamp")},
				{Name: "section", FieldType: ct.TableFieldTypeString, Label: adm.msg("admin_create_result_section")},
				{Name: "datatype", FieldType: ct.TableFieldTypeString, Label: adm.msg("admin_create_result_datatype")},
				{Column: &ct.TableColumn{
					Id:     "message",
					Header: adm.msg("admin_create_result_message"),
					Cell: func(row cu.IM, col ct.TableColumn, value interface{}) string {
						style := ""
						if row["state"] == "err" {
							style = `style="color:red;"`
						}
						return fmt.Sprintf(
							`<span class="cell-label">%s</span>
							<span %s >%s</span>`, col.Header, style, cu.ToString(value, ""))
					}}},
			}
			rows := cu.ToIMA(adm.Data["create_result"], []cu.IM{})
			return ccTbl("stamp", rows, fields, 0)
		},
		"admin_username": func() ct.ClientComponent {
			return ccLbl()
		},
		"username": func() ct.ClientComponent {
			return ccInp(ct.InputTypeString)
		},
		"admin_password": func() ct.ClientComponent {
			return ccLbl()
		},
		"password": func() ct.ClientComponent {
			return ccInp(ct.InputTypePassword)
		},
		"admin_confirm": func() ct.ClientComponent {
			return ccLbl()
		},
		"confirm": func() ct.ClientComponent {
			return ccInp(ct.InputTypePassword)
		},
		"admin_database": func() ct.ClientComponent {
			return ccLbl()
		},
		"database": func() ct.ClientComponent {
			return ccInp(ct.InputTypeString)
		},
		"login": func() ct.ClientComponent {
			disabled := (cu.ToString(adm.Data["username"], "") == "") || (cu.ToString(adm.Data["database"], "") == "")
			return ccBtn(ct.ButtonStylePrimary, adm.msg("admin_"+name), disabled)
		},
		"password_change": func() ct.ClientComponent {
			disabled := (cu.ToString(adm.Data["username"], "") == "") || (cu.ToString(adm.Data["password"], "") == "") || (cu.ToString(adm.Data["confirm"], "") == "")
			return ccBtn(ct.ButtonStylePrimary, adm.msg("admin_"+name), disabled)
		},
		"install_ico": func() ct.ClientComponent {
			return &ct.Icon{
				BaseComponent: ct.BaseComponent{
					Id: adm.Id + "_" + cu.ToString(data["reportkey"], ""), Name: cu.ToString(data["event"], ""),
					EventURL:     adm.EventURL,
					Target:       adm.Target,
					Indicator:    ct.IndicatorSpinner,
					OnResponse:   adm.response,
					RequestValue: adm.RequestValue,
					RequestMap:   adm.RequestMap,
					Data:         data,
				},
				Color: cu.ToString(data["color"], ""),
				Value: cu.ToString(data["ico_key"], ""),
				Width: 20, Height: 20,
			}
		},
		"report_list": func() ct.ClientComponent {
			fields := []ct.TableField{
				{Column: &ct.TableColumn{
					Id:        "installed",
					Header:    adm.msg("admin_report_list_installed"),
					CellStyle: cu.SM{"text-align": "center"},
					Cell: func(row cu.IM, col ct.TableColumn, value interface{}) string {
						idata := cu.IM{
							"ico_key": "Plus", "color": "green",
							"event": AdminEventReportInstall,
						}
						if cu.ToBoolean(value, false) {
							idata = cu.IM{
								"ico_key": "Times", "color": "red",
								"event": AdminEventReportDelete,
							}
						}
						idata["reportkey"] = row["reportkey"]
						res, _ := adm.getComponent("install_ico", idata)
						return fmt.Sprintf(
							`<span class="cell-label">%s</span>%s`, col.Header, res)
					}}},
				{Name: "reportkey", FieldType: ct.TableFieldTypeString, Label: adm.msg("admin_report_list_reportkey")},
				{Name: "repname", FieldType: ct.TableFieldTypeString, Label: adm.msg("admin_report_list_repname")},
				{Name: "description", FieldType: ct.TableFieldTypeString, Label: adm.msg("admin_report_list_description")},
				{Name: "reptype", FieldType: ct.TableFieldTypeString, Label: adm.msg("admin_report_list_reptype")},
				{Name: "filename", FieldType: ct.TableFieldTypeString, Label: adm.msg("admin_report_list_filename")},
				{Name: "label", FieldType: ct.TableFieldTypeString, Label: adm.msg("admin_report_list_label")},
			}
			rows := cu.ToIMA(adm.Data["report_list"], []cu.IM{})
			currentPage := int64(0)
			if current, valid := adm.Data["report_list_current_page"].(int64); valid {
				currentPage = current
			}
			return ccTbl("reportkey", rows, fields, currentPage)
		},
		"env_list": func() ct.ClientComponent {
			fields := []ct.TableField{
				{Name: "envkey", FieldType: ct.TableFieldTypeString, Label: adm.msg("admin_env_list_key")},
				{Name: "envvalue", FieldType: ct.TableFieldTypeString, Label: adm.msg("admin_env_list_value")},
			}
			rows := cu.ToIMA(adm.Data["env_list"], []cu.IM{})
			return ccTbl("envkey", rows, fields, 0)
		},
		"locales": func() ct.ClientComponent {
			locales := adm.Data["locales"].(cu.IM)
			return &Locale{
				BaseComponent: ct.BaseComponent{
					Id: adm.Id + "_" + name, Name: name,
					EventURL:     adm.EventURL,
					OnResponse:   adm.response,
					RequestValue: adm.RequestValue,
					RequestMap:   adm.RequestMap,
					Data: cu.IM{
						"deflang":    locales["deflang"],
						"locales":    locales["locale"],
						"tag_keys":   locales["tag_key"],
						"tag_values": locales["tag_values"],
						"locfile":    locales["locfile"],
					},
				},
				Locales: locales["locales"].([]ct.SelectOption),
				TagKeys: locales["tag_keys"].([]ct.SelectOption),
				Labels:  adminDefaultLabel,
			}
		},
	}
	cc := ccMap[name]()
	res, err = cc.Render()
	return res, err
}

func (adm *Admin) msg(labelID string) string {
	if label, found := adm.Labels[labelID]; found {
		return label
	}
	return labelID
}

func (adm *Admin) Render() (res string, err error) {
	adm.InitProps(adm)

	funcMap := map[string]any{
		"msg": func(labelID string) string {
			return adm.msg(labelID)
		},
		"styleMap": func() bool {
			return len(adm.Style) > 0
		},
		"customClass": func() string {
			return strings.Join(adm.Class, " ")
		},
		"adminComponent": func(name string) (string, error) {
			return adm.getComponent(name, cu.IM{})
		},
		"showResult": func() bool {
			return len(cu.ToIMA(adm.Data["create_result"], []cu.IM{})) > 0
		},
		"tokenLogin": func() bool {
			if adm.TokenLogin != nil {
				database := cu.ToString(adm.Data["database"], "")
				return adm.TokenLogin(database, adm.Token)
			}
			return false
		},
	}
	tpl := `<div id="{{ .Id }}" theme="{{ .Theme }}" class="admin row mobile {{ customClass }}"
	{{ if styleMap }} style="{{ range $key, $value := .Style }}{{ $key }}:{{ $value }};{{ end }}"{{ end }}>
	<div class="row title">
	<div class="cell">
	<div class="cell container bold" ><span>{{ msg "admin_title" }}</span></div>
	<div class="cell">{{ adminComponent "theme" }}</div>
	</div>
	<div class="cell version-cell" ><span>{{ .Version }}</span></div>
	</div>
	{{ adminComponent "main_menu" }}
	{{ if eq .Module "database" }}
	<div class="row full section-small" >
	<div class="row full container-small section-small-bottom" >
	<div class="cell mobile">
	<div class="cell padding-small mobile" >
	<div class="section-small" >{{ adminComponent "admin_api_key" }}</div>
	{{ adminComponent "api_key" }}
	</div>
	<div class="cell padding-small mobile" >
	<div class="section-small" >{{ adminComponent "admin_alias" }}</div>
	{{ adminComponent "alias" }}
	</div>
	<div class="cell padding-small mobile" >
	<div class="section-small" >{{ adminComponent "admin_demo" }}</div>
	{{ adminComponent "demo" }}
	</div>
	</div></div>
	<div class="row full section-small" >
	<div class="cell container center" >
	{{ adminComponent "create" }}
	</div>
	</div>
	{{ if showResult }}
	<div class="container section-small" >
	{{ adminComponent "create_result" }}
	</div>
	{{ end }}
	</div>
	{{ end }}
	{{ if eq .Module "locales" }}
	{{ adminComponent "locales" }}
	{{ end }}
	{{ if eq .Module "login" }}
	{{ if eq tokenLogin false }}
	<div class="row full section-small" >
	<div class="row full container-small section-small-bottom" >
	<div class="cell padding-small mobile">
	<div class="section-small" >{{ adminComponent "admin_username" }}</div>
	{{ adminComponent "username" }}
	</div>
	<div class="cell padding-small mobile">
	<div class="section-small" >{{ adminComponent "admin_password" }}</div>
	{{ adminComponent "password" }}
	</div>
	<div class="cell padding-small mobile">
	<div class="section-small" >{{ adminComponent "admin_database" }}</div>
	{{ adminComponent "database" }}
	</div>
	</div></div>
	<div class="row full section-small-bottom" >
	<div class="cell container center section-small-bottom" >
	{{ adminComponent "login" }}
	</div>
	</div>
	{{ else }}
	{{ adminComponent "view_menu" }}
	{{ if eq .View "password" }}
	<div class="row full section-small" >
	<div class="row full container-small section-small-bottom" >
	<div class="cell padding-small mobile">
	<div class="section-small" >{{ adminComponent "admin_username" }}</div>
	{{ adminComponent "username" }}
	</div>
	<div class="cell padding-small mobile">
	<div class="section-small" >{{ adminComponent "admin_password" }}</div>
	{{ adminComponent "password" }}
	</div>
	<div class="cell padding-small mobile">
	<div class="section-small" >{{ adminComponent "admin_confirm" }}</div>
	{{ adminComponent "confirm" }}
	</div>
	</div></div>
	<div class="row full section-small-bottom" >
	<div class="cell container center section-small-bottom" >
	{{ adminComponent "password_change" }}
	</div>
	</div>
	{{ end }}
	{{ if eq .View "report" }}
	<div class="container section-small" >
	{{ adminComponent "report_list" }}
	</div>
	</div>
	{{ end }}
	{{ if eq .View "configuration" }}
	<div class="row full section-small" >
	<div class="container section-small" >
	{{ adminComponent "env_list" }}
	</div>
	</div>
	{{ end }}
	{{ end }}
	{{ end }}
	</div>`

	return cu.TemplateBuilder("admin", tpl, funcMap, adm)
}
