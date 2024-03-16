package component

import (
	"fmt"
	"strings"

	fm "github.com/nervatura/component/component/atom"
	bc "github.com/nervatura/component/component/base"
	mc "github.com/nervatura/component/component/molecule"
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

var adminDefaultLabel bc.SM = bc.SM{
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
}

var adminIcoMap map[string][]string = map[string][]string{
	bc.ThemeDark: {bc.ThemeLight, "Sun"}, bc.ThemeLight: {bc.ThemeDark, "Moon"},
}

type Admin struct {
	bc.BaseComponent
	Version    string                            `json:"version"`
	Theme      string                            `json:"theme"`
	Module     string                            `json:"module"`
	View       string                            `json:"view"`
	Token      string                            `json:"token"`
	HelpURL    string                            `json:"help_url"`
	ClientURL  string                            `json:"client_url"`
	LocalesURL string                            `json:"locales_url"`
	Labels     bc.SM                             `json:"labels"`
	TokenLogin func(database, token string) bool `json:"-"` // Token validation
}

func (adm *Admin) Properties() bc.IM {
	return bc.MergeIM(
		adm.BaseComponent.Properties(),
		bc.IM{
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
			return adm.CheckEnumValue(bc.ToString(propValue, ""), bc.ThemeLight, bc.Theme)
		},
		"labels": func() interface{} {
			value := bc.SetSMValue(adm.Labels, "", "")
			if smap, valid := propValue.(bc.SM); valid {
				value = bc.MergeSM(value, smap)
			}
			if len(value) == 0 {
				value = adminDefaultLabel
			}
			return value
		},
		"target": func() interface{} {
			adm.SetProperty("id", adm.Id)
			value := bc.ToString(propValue, adm.Id)
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
			adm.Version = bc.ToString(propValue, "1.0.0")
			return adm.Version
		},
		"theme": func() interface{} {
			adm.Theme = adm.Validation(propName, propValue).(string)
			return adm.Theme
		},
		"module": func() interface{} {
			adm.Module = bc.ToString(propValue, "database")
			return adm.Module
		},
		"view": func() interface{} {
			adm.View = bc.ToString(propValue, "password")
			return adm.View
		},
		"token": func() interface{} {
			adm.Token = bc.ToString(propValue, "")
			return adm.Token
		},
		"help_url": func() interface{} {
			adm.HelpURL = bc.ToString(propValue, "")
			return adm.HelpURL
		},
		"client_url": func() interface{} {
			adm.ClientURL = bc.ToString(propValue, "")
			return adm.ClientURL
		},
		"locales_url": func() interface{} {
			adm.LocalesURL = bc.ToString(propValue, "")
			return adm.LocalesURL
		},
		"labels": func() interface{} {
			adm.Labels = adm.Validation(propName, propValue).(bc.SM)
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

func (adm *Admin) OnRequest(te bc.TriggerEvent) (re bc.ResponseEvent) {
	if cc, found := adm.RequestMap[te.Id]; found {
		return cc.OnRequest(te)
	}
	re = bc.ResponseEvent{
		Trigger: &fm.Toast{
			Type:  fm.ToastTypeError,
			Value: fmt.Sprintf("Invalid parameter: %s", te.Id),
		},
		TriggerName: te.Name,
		Name:        te.Name,
		Header: bc.SM{
			bc.HeaderRetarget: "#toast-msg",
			bc.HeaderReswap:   "innerHTML",
		},
	}
	return re
}

func (adm *Admin) response(evt bc.ResponseEvent) (re bc.ResponseEvent) {
	admEvt := bc.ResponseEvent{
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
		if evt.Name == mc.TableEventCurrentPage {
			adm.SetProperty("data", bc.MergeIM(adm.Data, bc.IM{"report_list_current_page": evt.Value}))
		}
		return evt

	case "api_key", "alias", "demo", "username", "password", "database", "confirm", "report_key":
		admEvt.Name = AdminEventChange
		adm.SetProperty("data", bc.IM{evt.TriggerName: admEvt.Value})

	case "theme":
		admEvt.Name = AdminEventTheme
		adm.SetProperty("theme", adminIcoMap[adm.Theme][0])

	case "main_menu":
		admEvt.Name = AdminEventModule
		adm.SetProperty("module", admEvt.Value)
		if admEvt.Value == "help" && adm.HelpURL != "" {
			admEvt.Header = bc.MergeSM(admEvt.Header,
				bc.SM{bc.HeaderRedirect: adm.HelpURL})
		}
		if admEvt.Value == "client" && adm.ClientURL != "" {
			admEvt.Header = bc.MergeSM(admEvt.Header,
				bc.SM{bc.HeaderRedirect: adm.ClientURL})
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
		reportkey := bc.ToString(evt.Trigger.GetProperty("data").(bc.IM)["reportkey"], "")
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

func (adm *Admin) getComponent(name string, data bc.IM) (res string, err error) {
	ccLbl := func() *fm.Label {
		return &fm.Label{
			Value: adm.msg(name),
		}
	}
	ccSel := func(options []fm.SelectOption) *fm.Select {
		return &fm.Select{
			BaseComponent: bc.BaseComponent{
				Id: adm.Id + "_" + name, Name: name,
				EventURL:     adm.EventURL,
				Target:       adm.Target,
				OnResponse:   adm.response,
				RequestValue: adm.RequestValue,
				RequestMap:   adm.RequestMap,
			},
			Label:   adm.msg("admin_" + name),
			IsNull:  false,
			Value:   bc.ToString(adm.Data[name], ""),
			Options: options,
		}
	}
	ccInp := func(itype string) *fm.Input {
		return &fm.Input{
			BaseComponent: bc.BaseComponent{
				Id: adm.Id + "_" + name, Name: name,
				EventURL:     adm.EventURL,
				Target:       adm.Target,
				Swap:         bc.SwapOuterHTML,
				OnResponse:   adm.response,
				RequestValue: adm.RequestValue,
				RequestMap:   adm.RequestMap,
			},
			Type:  itype,
			Label: adm.msg("admin_" + name),
			Value: bc.ToString(adm.Data[name], ""),
			Full:  true,
		}
	}
	ccBtn := func(btnType, label string, disabled bool) *fm.Button {
		return &fm.Button{
			BaseComponent: bc.BaseComponent{
				Id: adm.Id + "_" + name, Name: name,
				EventURL:     adm.EventURL,
				Target:       adm.Target,
				OnResponse:   adm.response,
				RequestValue: adm.RequestValue,
				RequestMap:   adm.RequestMap,
			},
			Type:     btnType,
			Label:    label,
			Disabled: disabled,
		}
	}
	ccMenu := func(items []mc.MenuBarItem, value, class string) *mc.MenuBar {
		return &mc.MenuBar{
			BaseComponent: bc.BaseComponent{
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
	ccTbl := func(rowKey string, rows []bc.IM, fields []mc.TableField, currentPage int64) *mc.Table {
		tbl := &mc.Table{
			BaseComponent: bc.BaseComponent{
				Id: adm.Id + "_" + name, Name: name,
				EventURL:     adm.EventURL,
				OnResponse:   adm.response,
				RequestValue: adm.RequestValue,
				RequestMap:   adm.RequestMap,
			},
			Rows:        rows,
			Fields:      fields,
			Pagination:  mc.PaginationTypeTop,
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
	ccMap := map[string]func() bc.ClientComponent{
		"main_menu": func() bc.ClientComponent {
			return ccMenu(
				[]mc.MenuBarItem{
					{Value: "database", Label: adm.msg("admin_database"), Icon: "Database"},
					{Value: "login", Label: adm.msg("admin_login"), Icon: "Edit"},
					{Value: "client", Label: adm.msg("admin_client"), Icon: "Globe"},
					{Value: "locales", Label: adm.msg("admin_locales"), Icon: "User"},
					{Value: "help", Label: adm.msg("admin_help"), Icon: "QuestionCircle"},
				},
				bc.ToString(adm.GetProperty("module"), ""), "border-top")
		},
		"view_menu": func() bc.ClientComponent {
			return ccMenu(
				[]mc.MenuBarItem{
					{Value: "password", Label: adm.msg("admin_password"), Icon: "Key"},
					{Value: "report", Label: adm.msg("admin_report"), Icon: "ChartBar"},
					{Value: "configuration", Label: adm.msg("admin_configuration"), Icon: "Cog"},
					{Value: "logout", Label: adm.msg("admin_logout"), Icon: "Exit"},
				},
				bc.ToString(adm.GetProperty("view"), ""), "border-top")
		},
		"admin_api_key": func() bc.ClientComponent {
			return ccLbl()
		},
		"api_key": func() bc.ClientComponent {
			return ccInp(fm.InputTypeText)
		},
		"admin_alias": func() bc.ClientComponent {
			return ccLbl()
		},
		"alias": func() bc.ClientComponent {
			return ccInp(fm.InputTypeText)
		},
		"admin_demo": func() bc.ClientComponent {
			return ccLbl()
		},
		"demo": func() bc.ClientComponent {
			return ccSel([]fm.SelectOption{
				{Value: "true", Text: adm.msg("admin_true")},
				{Value: "false", Text: adm.msg("admin_false")},
			})
		},
		"create": func() bc.ClientComponent {
			disabled := (bc.ToString(adm.Data["alias"], "") == "") || (bc.ToString(adm.Data["api_key"], "") == "")
			return ccBtn(fm.ButtonTypePrimary, adm.msg("admin_"+name), disabled)
		},
		"theme": func() bc.ClientComponent {
			themeBtn := ccBtn(fm.ButtonTypePrimary, "", false)
			themeBtn.Style = bc.SM{"padding": "4px"}
			themeBtn.LabelComponent = &fm.Icon{Value: adminIcoMap[adm.Theme][1], Width: 18, Height: 18}
			return themeBtn
		},
		"create_result": func() bc.ClientComponent {
			fields := []mc.TableField{
				{Column: &mc.TableColumn{
					Id:        "state",
					Header:    adm.msg("admin_create_result_state"),
					CellStyle: bc.SM{"text-align": "center"},
					Cell: func(row bc.IM, col mc.TableColumn, value interface{}) string {
						icoKey := "InfoCircle"
						color := "orange"
						if value == "err" {
							icoKey = "ExclamationTriangle"
							color = "red"
						}
						res, _ := (&fm.Icon{Value: icoKey, Color: color}).Render()
						return fmt.Sprintf(
							`<span class="cell-label">%s</span>%s`, col.Header, res)
					}}},
				{Name: "stamp", FieldType: mc.TableFieldTypeTime, Label: adm.msg("admin_create_result_stamp")},
				{Name: "section", FieldType: mc.TableFieldTypeString, Label: adm.msg("admin_create_result_section")},
				{Name: "datatype", FieldType: mc.TableFieldTypeString, Label: adm.msg("admin_create_result_datatype")},
				{Column: &mc.TableColumn{
					Id:     "message",
					Header: adm.msg("admin_create_result_message"),
					Cell: func(row bc.IM, col mc.TableColumn, value interface{}) string {
						style := ""
						if row["state"] == "err" {
							style = `style="color:red;"`
						}
						return fmt.Sprintf(
							`<span class="cell-label">%s</span>
							<span %s >%s</span>`, col.Header, style, bc.ToString(value, ""))
					}}},
			}
			rows := bc.ToIMA(adm.Data["create_result"], []bc.IM{})
			return ccTbl("stamp", rows, fields, 0)
		},
		"admin_username": func() bc.ClientComponent {
			return ccLbl()
		},
		"username": func() bc.ClientComponent {
			return ccInp(fm.InputTypeText)
		},
		"admin_password": func() bc.ClientComponent {
			return ccLbl()
		},
		"password": func() bc.ClientComponent {
			return ccInp(fm.InputTypePassword)
		},
		"admin_confirm": func() bc.ClientComponent {
			return ccLbl()
		},
		"confirm": func() bc.ClientComponent {
			return ccInp(fm.InputTypePassword)
		},
		"admin_database": func() bc.ClientComponent {
			return ccLbl()
		},
		"database": func() bc.ClientComponent {
			return ccInp(fm.InputTypeText)
		},
		"login": func() bc.ClientComponent {
			disabled := (bc.ToString(adm.Data["username"], "") == "") || (bc.ToString(adm.Data["database"], "") == "")
			return ccBtn(fm.ButtonTypePrimary, adm.msg("admin_"+name), disabled)
		},
		"password_change": func() bc.ClientComponent {
			disabled := (bc.ToString(adm.Data["username"], "") == "") || (bc.ToString(adm.Data["password"], "") == "") || (bc.ToString(adm.Data["confirm"], "") == "")
			return ccBtn(fm.ButtonTypePrimary, adm.msg("admin_"+name), disabled)
		},
		"install_ico": func() bc.ClientComponent {
			return &fm.Icon{
				BaseComponent: bc.BaseComponent{
					Id: adm.Id + "_" + bc.ToString(data["reportkey"], ""), Name: bc.ToString(data["event"], ""),
					EventURL:     adm.EventURL,
					Target:       adm.Target,
					Indicator:    bc.IndicatorSpinner,
					OnResponse:   adm.response,
					RequestValue: adm.RequestValue,
					RequestMap:   adm.RequestMap,
					Data:         data,
				},
				Color: bc.ToString(data["color"], ""),
				Value: bc.ToString(data["ico_key"], ""),
				Width: 20, Height: 20,
			}
		},
		"report_list": func() bc.ClientComponent {
			fields := []mc.TableField{
				{Column: &mc.TableColumn{
					Id:        "installed",
					Header:    adm.msg("admin_report_list_installed"),
					CellStyle: bc.SM{"text-align": "center"},
					Cell: func(row bc.IM, col mc.TableColumn, value interface{}) string {
						idata := bc.IM{
							"ico_key": "Plus", "color": "green",
							"event": AdminEventReportInstall,
						}
						if bc.ToBoolean(value, false) {
							idata = bc.IM{
								"ico_key": "Times", "color": "red",
								"event": AdminEventReportDelete,
							}
						}
						idata["reportkey"] = row["reportkey"]
						res, _ := adm.getComponent("install_ico", idata)
						return fmt.Sprintf(
							`<span class="cell-label">%s</span>%s`, col.Header, res)
					}}},
				{Name: "reportkey", FieldType: mc.TableFieldTypeString, Label: adm.msg("admin_report_list_reportkey")},
				{Name: "repname", FieldType: mc.TableFieldTypeString, Label: adm.msg("admin_report_list_repname")},
				{Name: "description", FieldType: mc.TableFieldTypeString, Label: adm.msg("admin_report_list_description")},
				{Name: "reptype", FieldType: mc.TableFieldTypeString, Label: adm.msg("admin_report_list_reptype")},
				{Name: "filename", FieldType: mc.TableFieldTypeString, Label: adm.msg("admin_report_list_filename")},
				{Name: "label", FieldType: mc.TableFieldTypeString, Label: adm.msg("admin_report_list_label")},
			}
			rows := bc.ToIMA(adm.Data["report_list"], []bc.IM{})
			currentPage := int64(0)
			if current, valid := adm.Data["report_list_current_page"].(int64); valid {
				currentPage = current
			}
			return ccTbl("reportkey", rows, fields, currentPage)
		},
		"env_list": func() bc.ClientComponent {
			fields := []mc.TableField{
				{Name: "envkey", FieldType: mc.TableFieldTypeString, Label: adm.msg("admin_env_list_key")},
				{Name: "envvalue", FieldType: mc.TableFieldTypeString, Label: adm.msg("admin_env_list_value")},
			}
			rows := bc.ToIMA(adm.Data["env_list"], []bc.IM{})
			return ccTbl("envkey", rows, fields, 0)
		},
		"locales": func() bc.ClientComponent {
			locales := adm.Data["locales"].(bc.IM)
			return &Locales{
				BaseComponent: bc.BaseComponent{
					Id: adm.Id + "_" + name, Name: name,
					EventURL:     adm.EventURL,
					OnResponse:   adm.response,
					RequestValue: adm.RequestValue,
					RequestMap:   adm.RequestMap,
					Data: bc.IM{
						"deflang":    locales["deflang"],
						"locales":    locales["locale"],
						"tag_keys":   locales["tag_key"],
						"tag_values": locales["tag_values"],
						"locfile":    locales["locfile"],
					},
				},
				Locales: locales["locales"].([]fm.SelectOption),
				TagKeys: locales["tag_keys"].([]fm.SelectOption),
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
			return adm.getComponent(name, bc.IM{})
		},
		"showResult": func() bool {
			return len(bc.ToIMA(adm.Data["create_result"], []bc.IM{})) > 0
		},
		"tokenLogin": func() bool {
			if adm.TokenLogin != nil {
				database := bc.ToString(adm.Data["database"], "")
				return adm.TokenLogin(database, adm.Token)
			}
			return false
		},
	}
	tpl := `<div id="{{ .Id }}" theme="{{ .Theme }}" class="admin row mobile {{ customClass }}"
	{{ if styleMap }} style="{{ range $key, $value := .Style }}{{ $key }}:{{ $value }};{{ end }}"{{ end }}>
	<div class="row title">
	<div class="cell">
	<div class="cell title-cell" ><span>{{ msg "admin_title" }}</span></div>
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

	return bc.TemplateBuilder("admin", tpl, funcMap, adm)
}
