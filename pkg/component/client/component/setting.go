package component

import (
	"strings"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	st "github.com/nervatura/nervatura/v6/pkg/static"
)

var settingIconMap = map[string]string{
	"setting": ct.IconKeyboard,
}

func settingSideBar(labels cu.SM, data cu.IM) (items []ct.SideBarItem) {
	//sideGroup := cu.ToString(data["side_group"], "")
	viewName := cu.ToString(data["view"], "")

	sideElement := func(name string, selected bool) *ct.SideBarElement {
		return &ct.SideBarElement{
			Name:     name,
			Value:    name,
			Label:    labels["setting_"+name],
			Icon:     settingIconMap[name],
			Selected: selected,
		}
	}
	sb := []ct.SideBarItem{
		&ct.SideBarSeparator{},
	}
	sb = append(sb, sideElement("setting", viewName == "setting"))
	return sb
}

func settingEditorView(labels cu.SM, data cu.IM) (views []ct.EditorView) {
	viewName := cu.ToString(data["view"], "")
	return []ct.EditorView{
		{
			Key:   viewName,
			Label: labels["setting_"+viewName],
			Icon:  settingIconMap[viewName],
		},
	}
}

func settingRow(view string, labels cu.SM, data cu.IM) (rows []ct.Row) {
	var setting cu.IM = cu.ToIM(data["setting"], cu.IM{})
	configData := cu.ToIMA(data["config_data"], []cu.IM{})

	pageSizeOpt := fromConfig("paper_size", configData)
	orientationOpt := fromConfig("orientation", configData)
	themeOpt := func() (tmOpt []ct.SelectOption) {
		tmOpt = []ct.SelectOption{}
		for _, tm := range st.ClientTheme {
			theme := strings.Split(tm, ",")
			if len(theme) > 0 {
				tmOpt = append(tmOpt, ct.SelectOption{
					Value: theme[0],
					Text:  labels["theme_"+theme[0]],
				})
			}
		}
		return tmOpt
	}

	langOpt := func() (locales []ct.SelectOption) {
		locales = []ct.SelectOption{}
		for _, loc := range st.ClientLang {
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

	pgOpt := func() (options []ct.SelectOption) {
		options = []ct.SelectOption{}
		for _, size := range []string{"5", "10", "20", "50", "100"} {
			options = append(options, ct.SelectOption{
				Value: size,
				Text:  size,
			})
		}
		return options
	}

	viewMap := map[string]func() []ct.Row{
		"setting": func() []ct.Row {
			return []ct.Row{
				{Columns: []ct.RowColumn{
					{Label: labels["setting_theme"], Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "setting_theme",
						},
						Type: ct.FieldTypeSelect, Value: cu.IM{
							"name":    "theme",
							"options": themeOpt(),
							"is_null": false,
							"value":   cu.ToString(setting["theme"], st.DefaultTheme),
						},
					}},
					{Label: labels["setting_lang"], Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "setting_lang",
						},
						Type: ct.FieldTypeSelect, Value: cu.IM{
							"name":    "lang",
							"options": langOpt(),
							"is_null": false,
							"value":   cu.ToString(setting["lang"], st.DefaultLang),
						},
					}},
					{Label: labels["setting_pagination"], Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "setting_pagination",
						},
						Type: ct.FieldTypeSelect, Value: cu.IM{
							"name":    "pagination",
							"options": pgOpt(),
							"is_null": false,
							"value":   cu.ToString(setting["pagination"], st.DefaultPagination),
						},
					}},
				}, Full: true, BorderBottom: true},
				{Columns: []ct.RowColumn{
					{Label: labels["setting_page_size"], Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "setting_page_size",
						},
						Type: ct.FieldTypeSelect, Value: cu.IM{
							"name":    "page_size",
							"options": pageSizeOpt,
							"is_null": false,
							"value":   cu.ToString(setting["page_size"], st.DefaultPaperSize),
						},
					}},
					{Label: labels["setting_orientation"], Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "setting_orientation",
						},
						Type: ct.FieldTypeSelect, Value: cu.IM{
							"name":    "orientation",
							"options": orientationOpt,
							"is_null": false,
							"value":   cu.ToString(setting["orientation"], st.DefaultOrientation),
						},
					}},
					/*
						{Label: labels["setting_history"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "history",
							},
							Type: ct.FieldTypeInteger, Value: cu.IM{
								"name":  "history",
								"value": cu.ToInteger(setting["history"], st.DefaultHistory),
							},
						}},
					*/
					{Label: labels["setting_export_sep"], Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "export_sep",
						},
						Type: ct.FieldTypeString, Value: cu.IM{
							"name":  "export_sep",
							"value": cu.ToString(setting["export_sep"], st.DefaultExportSep),
							"size":  1,
						},
					}},
				}, Full: true, BorderBottom: true},
				{Columns: []ct.RowColumn{
					{Label: labels["setting_password"], Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "setting_password",
						},
						Type: ct.FieldTypePassword, Value: cu.IM{
							"name":        "password",
							"invalid":     (cu.ToString(setting["password"], "") == ""),
							"placeholder": labels["mandatory_data"],
							"value":       cu.ToString(setting["password"], ""),
						},
					}},
					{Label: labels["setting_confirm"], Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "setting_confirm",
						},
						Type: ct.FieldTypePassword, Value: cu.IM{
							"name":        "confirm",
							"invalid":     (cu.ToString(setting["confirm"], "") == ""),
							"placeholder": labels["mandatory_data"],
							"value":       cu.ToString(setting["confirm"], ""),
						},
					}},
					{Label: labels["setting_password_validation"],
						Value: ct.Field{
							Type: ct.FieldTypeButton,
							Value: cu.IM{
								"name":         "change_password",
								"type":         ct.ButtonTypeButton,
								"button_style": ct.ButtonStyleBorder,
								"icon":         ct.IconKey,
								"label":        labels["setting_change_password"],
								"disabled": (cu.ToString(setting["password"], "") == "") || (cu.ToString(setting["confirm"], "") == "") ||
									(cu.ToString(setting["password"], "") != cu.ToString(setting["confirm"], "")),
							},
						}},
				}, Full: true, BorderBottom: true},
			}
		},
	}

	if rows, found := viewMap[view]; found {
		return rows()
	}
	return []ct.Row{}
}

func settingTable(_ string, _ cu.SM, _ cu.IM) []ct.Table {
	return []ct.Table{}
}

func settingForm(_ string, _ cu.SM, _ cu.IM) (form ct.Form) {
	return ct.Form{}
}
