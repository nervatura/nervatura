package component

import (
	"slices"
	"strings"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
	st "github.com/nervatura/nervatura/v6/pkg/static"
)

type SettingEditor struct{}

var settingIconMap = map[string]string{
	"setting":     ct.IconKeyboard,
	"config_map":  ct.IconDatabase,
	"config_data": ct.IconCog,
	"currency":    ct.IconDollar,
	"tax":         ct.IconTicket,
	"auth":        ct.IconUserLock,
	"shortcut":    ct.IconShare,
	"template":    ct.IconBook,
}

func (e *SettingEditor) Frame(labels cu.SM, data cu.IM) (title, icon string) {
	return cu.ToString(data["editor_title"], labels["setting_title"]),
		cu.ToString(data["editor_icon"], ct.IconCog)
}

func (e *SettingEditor) SideBar(labels cu.SM, data cu.IM) (items []ct.SideBarItem) {
	viewName := cu.ToString(data["view"], "")
	user := cu.ToIM(data["user"], cu.IM{})

	sideElement := func(name string, selected bool) *ct.SideBarElement {
		return &ct.SideBarElement{
			Name:     name,
			Value:    name,
			Label:    " " + labels["setting_"+name],
			Icon:     settingIconMap[name],
			Selected: selected,
		}
	}
	sb := []ct.SideBarItem{
		&ct.SideBarSeparator{},
	}
	sb = append(sb, sideElement("setting", viewName == "setting"))
	if cu.ToString(user["user_group"], "") == md.UserGroupAdmin.String() {
		sb = append(sb, sideElement("config_data", viewName == "config_data"))
		sb = append(sb, sideElement("config_map", viewName == "config_map"))
		sb = append(sb, sideElement("auth", viewName == "auth"))
		sb = append(sb, sideElement("currency", viewName == "currency"))
		sb = append(sb, sideElement("tax", viewName == "tax"))
		sb = append(sb, sideElement("shortcut", viewName == "shortcut"))
		sb = append(sb, sideElement("template", viewName == "template"))
	}

	sb = append(sb,
		&ct.SideBarSeparator{},
		&ct.SideBarElementLink{
			SideBarElement: ct.SideBarElement{
				Name:  "editor_help",
				Value: "editor_help",
				Label: labels["editor_help"],
				Icon:  ct.IconQuestionCircle,
			},
			Href:       st.DocsClientPath + "#settings",
			LinkTarget: "_blank",
		})
	return sb
}

func (e *SettingEditor) View(labels cu.SM, data cu.IM) (views []ct.EditorView) {
	viewName := cu.ToString(data["view"], "")
	return []ct.EditorView{
		{
			Key:   viewName,
			Label: labels["setting_"+viewName],
			Icon:  settingIconMap[viewName],
		},
	}
}

func (e *SettingEditor) Row(view string, labels cu.SM, data cu.IM) (rows []ct.Row) {
	var setting cu.IM = cu.ToIM(data["setting"], cu.IM{})
	configData := cu.ToIMA(data["config_data"], []cu.IM{})
	configValues := cu.ToIMA(data["config_values"], []cu.IM{})
	auth := cu.ToIMA(data["auth"], []cu.IM{})

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
		return ct.SelectOptionRangeValidation(data["locales"], []ct.SelectOption{{Value: "en", Text: "English"}})
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
		"config_map": func() []ct.Row {
			mapRows := []cu.IM{}
			for _, row := range configValues {
				if row["config_type"] == md.ConfigTypeMap.String() {
					configMap := cu.ToIM(row["data"], cu.IM{})
					mapRows = append(mapRows, cu.MergeIM(row,
						cu.IM{"lslabel": cu.ToString(configMap["field_name"], ""), "lsvalue": cu.ToString(configMap["description"], "")}))
				}
			}
			return []ct.Row{
				{Columns: []ct.RowColumn{
					{Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Id: "config_map_list",
						},
						Type: ct.FieldTypeList,
						Value: cu.IM{
							"name":                "config_map",
							"rows":                mapRows,
							"pagination":          ct.PaginationTypeTop,
							"page_size":           10,
							"hide_paginaton_size": true,
							"list_filter":         true,
							"filter_placeholder":  labels["placeholder_filter"],
							"add_item":            true,
							"edit_item":           true,
							"delete_item":         true,
						},
					}},
				}, Full: true, BorderBottom: false},
			}
		},
		"shortcut": func() []ct.Row {
			mapRows := []cu.IM{}
			for _, row := range configValues {
				if row["config_type"] == md.ConfigTypeShortcut.String() {
					configMap := cu.ToIM(row["data"], cu.IM{})
					mapRows = append(mapRows, cu.MergeIM(row,
						cu.IM{"lslabel": cu.ToString(configMap["shortcut_key"], ""), "lsvalue": cu.ToString(configMap["description"], "")}))
				}
			}
			return []ct.Row{
				{Columns: []ct.RowColumn{
					{Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Id: "shortcut_list",
						},
						Type: ct.FieldTypeList,
						Value: cu.IM{
							"name":                "shortcut",
							"rows":                mapRows,
							"pagination":          ct.PaginationTypeTop,
							"page_size":           10,
							"hide_paginaton_size": true,
							"list_filter":         true,
							"filter_placeholder":  labels["placeholder_filter"],
							"add_item":            true,
							"edit_item":           true,
							"delete_item":         true,
						},
					}},
				}, Full: true, BorderBottom: true},
			}
		},
		"auth": func() []ct.Row {
			userRows := []cu.IM{}
			for _, row := range auth {
				userRows = append(userRows, cu.MergeIM(row,
					cu.IM{"lslabel": cu.ToString(row["user_name"], ""),
						"lsvalue": labels[cu.ToString(row["user_group"], "")]}))
			}
			return []ct.Row{
				{Columns: []ct.RowColumn{
					{Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Id: "auth_list",
						},
						Type: ct.FieldTypeList,
						Value: cu.IM{
							"name":                "auth",
							"rows":                userRows,
							"pagination":          ct.PaginationTypeTop,
							"page_size":           10,
							"hide_paginaton_size": true,
							"list_filter":         true,
							"filter_placeholder":  labels["placeholder_filter"],
							"add_item":            true,
							"edit_item":           true,
							"delete_item":         false,
						},
					}},
				}, Full: true, BorderBottom: false},
			}
		},
	}

	if rows, found := viewMap[view]; found {
		return rows()
	}
	return []ct.Row{}
}

func (e *SettingEditor) Table(view string, labels cu.SM, data cu.IM) []ct.Table {
	if !slices.Contains([]string{"config_data", "currency", "tax", "template"}, view) {
		return []ct.Table{}
	}

	tblMap := map[string]func() []ct.Table{
		"config_data": func() []ct.Table {
			configData := cu.ToIMA(data["config_data"], []cu.IM{})
			return []ct.Table{
				{
					BaseComponent: ct.BaseComponent{
						Name: "config_data",
					},
					Fields: []ct.TableField{
						{Name: "config_code", Label: labels["setting_config_code"], ReadOnly: true},
						{Name: "config_key", Label: labels["setting_config_key"], ReadOnly: true},
						{Name: "config_value", Label: labels["setting_config_value"], Required: true},
					},
					//RowKey:             "id",
					Rows:               configData,
					Pagination:         ct.PaginationTypeTop,
					PageSize:           10,
					HidePaginatonSize:  true,
					RowSelected:        true,
					TableFilter:        true,
					FilterPlaceholder:  labels["placeholder_filter"],
					Editable:           true,
					EditDeleteDisabled: true,
					Unsortable:         true,
				},
			}
		},
		"currency": func() []ct.Table {
			itemRows := func() []cu.IM {
				rows := []cu.IM{}
				currencies := cu.ToIMA(data["currency"], []cu.IM{})
				for _, currency := range currencies {
					currencyMeta := cu.ToIM(currency["currency_meta"], cu.IM{})
					rows = append(rows, cu.IM{
						"id":          currency["id"],
						"code":        currency["code"],
						"description": currencyMeta["description"],
						"digit":       cu.ToInteger(currencyMeta["digit"], 0),
						"cash_round":  cu.ToInteger(currencyMeta["cash_round"], 0),
					})
				}
				return rows
			}
			return []ct.Table{
				{
					Fields: []ct.TableField{
						{Name: "code", Label: labels["currency_code"], ReadOnly: true},
						{Name: "description", Label: labels["currency_description"]},
						{Name: "digit", Label: labels["currency_digit"], FieldType: ct.TableFieldTypeInteger},
						{Name: "cash_round", Label: labels["currency_cash_round"], FieldType: ct.TableFieldTypeInteger},
					},
					Rows:              itemRows(),
					Pagination:        ct.PaginationTypeTop,
					PageSize:          10,
					HidePaginatonSize: true,
					RowSelected:       true,
					TableFilter:       true,
					FilterPlaceholder: labels["placeholder_filter"],
					AddItem:           true,
					LabelAdd:          labels["currency_new"],
					Editable:          true,
					Unsortable:        true,
				},
			}
		},
		"tax": func() []ct.Table {
			itemRows := func() []cu.IM {
				rows := []cu.IM{}
				taxes := cu.ToIMA(data["tax"], []cu.IM{})
				for _, tax := range taxes {
					taxMeta := cu.ToIM(tax["tax_meta"], cu.IM{})
					rows = append(rows, cu.IM{
						"id":          tax["id"],
						"code":        tax["code"],
						"description": taxMeta["description"],
						"rate_value":  cu.ToFloat(taxMeta["rate_value"], 0),
					})
				}
				return rows
			}
			return []ct.Table{
				{
					Fields: []ct.TableField{
						{Name: "code", Label: labels["tax_code"], ReadOnly: true},
						{Name: "description", Label: labels["tax_description"]},
						{Name: "rate_value", Label: labels["tax_rate_value"], FieldType: ct.TableFieldTypeNumber},
					},
					Rows:              itemRows(),
					Pagination:        ct.PaginationTypeTop,
					PageSize:          10,
					HidePaginatonSize: true,
					RowSelected:       true,
					TableFilter:       true,
					FilterPlaceholder: labels["placeholder_filter"],
					AddItem:           true,
					LabelAdd:          labels["tax_new"],
					Editable:          true,
					Unsortable:        true,
				},
			}
		},
		"template": func() []ct.Table {
			templates := cu.ToIMA(data["template"], []cu.IM{})
			return []ct.Table{
				{
					Fields: []ct.TableField{
						{Name: "installed", Label: labels["template_installed"],
							FieldType: ct.TableFieldTypeBool, TextAlign: ct.TextAlignCenter, ReadOnly: true},
						{Name: "report_key", Label: labels["template_report_key"], ReadOnly: true},
						{Name: "report_name", Label: labels["template_report_name"], ReadOnly: true},
						{Name: "label", Label: labels["template_label"], ReadOnly: true},
						//{Name: "file_type", Label: labels["template_file_type"], ReadOnly: true},
						{Name: "description", Label: labels["template_description"], ReadOnly: true},
						//{Name: "file_name", Label: labels["template_file_name"], ReadOnly: true},
					},
					Rows:              templates,
					RowKey:            "report_key",
					Pagination:        ct.PaginationTypeTop,
					PageSize:          10,
					HidePaginatonSize: true,
					RowSelected:       true,
					TableFilter:       true,
					FilterPlaceholder: labels["placeholder_filter"],
					Editable:          true,
					Unsortable:        true,
				},
			}
		},
	}

	return tblMap[view]()
}

func (e *SettingEditor) Form(formKey string, labels cu.SM, data cu.IM) (form ct.Form) {
	formData := cu.ToIM(data, cu.IM{})
	footerRows := func(updateDisabled, deleteDisabled bool) []ct.Row {
		rows := []ct.Row{
			{
				Columns: []ct.RowColumn{
					{Value: ct.Field{
						Type: ct.FieldTypeButton,
						Value: cu.IM{
							"name":         ct.FormEventOK,
							"type":         ct.ButtonTypeSubmit,
							"button_style": ct.ButtonStylePrimary,
							"icon":         ct.IconCheck,
							"label":        labels["editor_save"],
							//"auto_focus":   true,
							"selected": true,
							"disabled": updateDisabled,
						},
					}},
					{Value: ct.Field{
						Type: ct.FieldTypeButton,
						Value: cu.IM{
							"name":         ct.FormEventCancel,
							"type":         ct.ButtonTypeSubmit,
							"button_style": ct.ButtonStyleDefault,
							"icon":         ct.IconReply,
							"label":        labels["editor_back"],
						},
					}},
					{Value: ct.Field{
						Type:  ct.FieldTypeLabel,
						Value: cu.IM{},
					}},
					{Value: ct.Field{
						Type: ct.FieldTypeButton,
						Value: cu.IM{
							"name":         "form_delete",
							"type":         ct.ButtonTypeSubmit,
							"button_style": ct.ButtonStyleBorder,
							"icon":         ct.IconTimes,
							"label":        labels["editor_delete"],
							"style":        cu.SM{"color": "red", "fill": "red"},
							"disabled":     deleteDisabled,
						},
					}},
				},
				Full:         true,
				FieldCol:     false,
				BorderTop:    false,
				BorderBottom: false,
			},
		}
		return rows
	}
	frmMap := map[string]func() ct.Form{
		"config_map": func() ct.Form {
			var mp md.ConfigMap = md.ConfigMap{}
			cfData := cu.ToIM(formData["data"], cu.IM{})
			ut.ConvertToType(cfData, &mp)
			fieldTypeOpt := func() (opt []ct.SelectOption) {
				opt = []ct.SelectOption{}
				ft := md.FieldType(0)
				for _, ftKey := range ft.Keys() {
					opt = append(opt, ct.SelectOption{
						Value: ftKey, Text: ftKey,
					})
				}
				return opt
			}
			return ct.Form{
				Title: labels["setting_config_map"],
				Icon:  ct.IconDatabase,
				BodyRows: []ct.Row{
					{Columns: []ct.RowColumn{
						{Label: labels["setting_config_code"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "code",
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":     "code",
								"value":    cu.ToString(formData["code"], ""),
								"disabled": true,
							},
							//FormTrigger: true,
						}},
						{Label: labels["setting_map_field_type"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "field_type",
							},
							Type: ct.FieldTypeSelect, Value: cu.IM{
								"name":    "field_type",
								"options": fieldTypeOpt(),
								"is_null": false,
								"value":   mp.FieldType.String(),
							},
							FormTrigger: true,
						}},
						{Label: labels["setting_map_field_name"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "field_name",
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":        "field_name",
								"invalid":     (mp.FieldName == ""),
								"placeholder": labels["mandatory_data"],
								"value":       mp.FieldName,
							},
							FormTrigger: true,
						}},
					}, Full: true, BorderBottom: true},
					{Columns: []ct.RowColumn{
						{Label: labels["setting_map_description"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "description",
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":        "description",
								"invalid":     (mp.Description == ""),
								"placeholder": labels["mandatory_data"],
								"value":       mp.Description,
							},
							FormTrigger: true,
						}},
					}, Full: true, BorderBottom: true, FieldCol: true},
					{Columns: []ct.RowColumn{
						{
							Label: labels["setting_map_tags"], Value: ct.Field{
								BaseComponent: ct.BaseComponent{
									Name: "tags",
								},
								Type: ct.FieldTypeList, Value: cu.IM{
									"name":                "tags",
									"rows":                ut.ToTagList(mp.Tags),
									"label_value":         "tag",
									"pagination":          ct.PaginationTypeBottom,
									"page_size":           5,
									"hide_paginaton_size": true,
									"list_filter":         true,
									"filter_placeholder":  labels["placeholder_filter"],
									"add_item":            true,
									"add_icon":            ct.IconTag,
									"edit_item":           false,
									"delete_item":         true,
									"indicator":           ct.IndicatorSpinner,
								},
								FormTrigger: true,
							},
						},
						{
							Label: labels["setting_map_filter"], Value: ct.Field{
								BaseComponent: ct.BaseComponent{
									Name: "filter",
								},
								Type: ct.FieldTypeList, Value: cu.IM{
									"name":                "filter",
									"rows":                ut.ToTagList(ut.ToStringArray(cfData["filter"])),
									"label_value":         "tag",
									"pagination":          ct.PaginationTypeBottom,
									"page_size":           5,
									"hide_paginaton_size": true,
									"list_filter":         true,
									"filter_placeholder":  labels["placeholder_filter"],
									"add_item":            true,
									"add_icon":            ct.IconFilter,
									"edit_item":           false,
									"delete_item":         true,
									"indicator":           ct.IndicatorSpinner,
								},
								FormTrigger: true,
							},
						},
					}, Full: true, BorderBottom: true},
				},
				FooterRows: footerRows(((mp.FieldName == "") || (mp.Description == "")), (cu.ToString(formData["code"], "") == "")),
			}
		},
		"shortcut": func() ct.Form {
			var mp md.ConfigShortcut = md.ConfigShortcut{}
			cfData := cu.ToIM(formData["data"], cu.IM{})
			ut.ConvertToType(cfData, &mp)
			methodOpt := func() (opt []ct.SelectOption) {
				opt = []ct.SelectOption{}
				mm := md.ShortcutMethod(0)
				for _, mmKey := range mm.Keys() {
					opt = append(opt, ct.SelectOption{
						Value: mmKey, Text: mmKey,
					})
				}
				return opt
			}
			fieldRows := []cu.IM{}
			sfields := cu.ToIMA(cfData["fields"], []cu.IM{})
			for _, row := range sfields {
				fieldRows = append(fieldRows, cu.MergeIM(row,
					cu.IM{
						"lslabel": cu.ToString(row["field_name"], "") + " - " + cu.ToString(row["field_type"], ""),
						"lsvalue": cu.ToString(row["description"], "")}))
			}
			return ct.Form{
				Title: labels["setting_shortcut"],
				Icon:  ct.IconShare,
				BodyRows: []ct.Row{
					{Columns: []ct.RowColumn{
						{Label: labels["setting_config_code"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "code",
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":     "code",
								"value":    cu.ToString(formData["code"], ""),
								"disabled": true,
							},
							//FormTrigger: true,
						}},
						{Label: labels["setting_shortcut_method"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "method",
							},
							Type: ct.FieldTypeSelect, Value: cu.IM{
								"name":    "method",
								"options": methodOpt(),
								"is_null": false,
								"value":   mp.Method.String(),
							},
							FormTrigger: true,
						}},
						{Label: labels["setting_shortcut_shortcut_key"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "shortcut_key",
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":        "shortcut_key",
								"invalid":     (mp.ShortcutKey == ""),
								"placeholder": labels["mandatory_data"],
								"value":       mp.ShortcutKey,
							},
							FormTrigger: true,
						}},
					}, Full: true, BorderBottom: true},
					{Columns: []ct.RowColumn{
						{Label: labels["setting_shortcut_modul"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "modul",
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":  "modul",
								"value": mp.Modul,
							},
							FormTrigger: true,
						}},
						{Label: labels["setting_shortcut_icon"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "icon",
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":  "icon",
								"value": mp.Icon,
							},
							FormTrigger: true,
						}},
						{Label: labels["setting_shortcut_func_name"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "func_name",
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":  "func_name",
								"value": mp.Funcname,
							},
							FormTrigger: true,
						}},
					}, Full: true, BorderBottom: true},
					{Columns: []ct.RowColumn{
						{Label: labels["setting_shortcut_description"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "description",
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":        "description",
								"invalid":     (mp.Description == ""),
								"placeholder": labels["mandatory_data"],
								"value":       mp.Description,
							},
							FormTrigger: true,
						}},
						{Label: labels["setting_shortcut_address"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "address",
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":  "address",
								"value": mp.Address,
							},
							FormTrigger: true,
						}},
					}, Full: true, BorderBottom: true},
					{Columns: []ct.RowColumn{
						{
							Label: labels["setting_shortcut_fields"], Value: ct.Field{
								BaseComponent: ct.BaseComponent{
									Name: "fields",
								},
								Type: ct.FieldTypeList, Value: cu.IM{
									"name":                "fields",
									"rows":                fieldRows,
									"pagination":          ct.PaginationTypeBottom,
									"page_size":           5,
									"hide_paginaton_size": true,
									"list_filter":         true,
									"filter_placeholder":  labels["placeholder_filter"],
									"add_item":            true,
									"add_icon":            ct.IconKeyboard,
									"edit_item":           true,
									"delete_item":         true,
									"indicator":           ct.IndicatorSpinner,
								},
								FormTrigger: true,
							},
						},
					}, Full: true, BorderBottom: true},
				},
				FooterRows: footerRows(((mp.ShortcutKey == "") || (mp.Description == "")), (cu.ToString(formData["code"], "") == "")),
			}
		},
		"auth": func() ct.Form {
			var auth md.Auth = md.Auth{}
			authMeta := cu.ToIM(formData["auth_meta"], cu.IM{})
			ut.ConvertToType(formData, &auth)
			userGroupOpt := func() (opt []ct.SelectOption) {
				opt = []ct.SelectOption{}
				ug := md.UserGroup(0)
				for _, ugKey := range ug.Keys() {
					opt = append(opt, ct.SelectOption{
						Value: ugKey, Text: ugKey,
					})
				}
				return opt
			}
			return ct.Form{
				Title: labels["setting_auth"],
				Icon:  ct.IconUserLock,
				BodyRows: []ct.Row{
					{Columns: []ct.RowColumn{
						{Label: labels["auth_code"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "code",
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":     "code",
								"value":    cu.ToString(formData["code"], ""),
								"disabled": true,
							},
							//FormTrigger: true,
						}},
						{Label: labels["auth_user_group"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "user_group",
							},
							Type: ct.FieldTypeSelect, Value: cu.IM{
								"name":     "user_group",
								"options":  userGroupOpt(),
								"is_null":  false,
								"value":    auth.UserGroup.String(),
								"disabled": (auth.UserName == "admin"),
							},
							FormTrigger: true,
						}},
						{Label: labels["auth_user_name"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "user_name",
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":        "user_name",
								"invalid":     (auth.UserName == ""),
								"placeholder": labels["mandatory_data"],
								"value":       auth.UserName,
								"disabled":    true,
							},
							FormTrigger: true,
						}},
						{Label: labels["auth_disabled"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "disabled",
							},
							Type: ct.FieldTypeBool, Value: cu.IM{
								"name":     "disabled",
								"value":    auth.Disabled,
								"disabled": (auth.UserName == "admin"),
							},
							FormTrigger: true,
						}},
					}, Full: true, BorderBottom: true},
					{Columns: []ct.RowColumn{
						{Label: labels["auth_description"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "description",
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":  "description",
								"value": auth.AuthMeta.Description,
							},
							FormTrigger: true,
						}},
						{Label: labels["auth_password"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "password",
							},
							Type: ct.FieldTypeButton, Value: cu.IM{
								"name":         "password",
								"type":         ct.ButtonTypeSubmit,
								"button_style": ct.ButtonStyleBorder,
								"icon":         ct.IconKey,
								"label":        labels["auth_password_reset"],
							},
							FormTrigger: true,
						}},
					}, Full: true, BorderBottom: true},
					{Columns: []ct.RowColumn{
						{
							Label: labels["auth_tags"], Value: ct.Field{
								BaseComponent: ct.BaseComponent{
									Name: "tags",
								},
								Type: ct.FieldTypeList, Value: cu.IM{
									"name":                "tags",
									"rows":                ut.ToTagList(auth.AuthMeta.Tags),
									"label_value":         "tag",
									"pagination":          ct.PaginationTypeBottom,
									"page_size":           5,
									"hide_paginaton_size": true,
									"list_filter":         true,
									"filter_placeholder":  labels["placeholder_filter"],
									"add_item":            true,
									"add_icon":            ct.IconTag,
									"edit_item":           false,
									"delete_item":         true,
									"indicator":           ct.IndicatorSpinner,
								},
								FormTrigger: true,
							},
						},
						{
							Label: labels["setting_map_filter"], Value: ct.Field{
								BaseComponent: ct.BaseComponent{
									Name: "filter",
								},
								Type: ct.FieldTypeList, Value: cu.IM{
									"name":                "filter",
									"rows":                ut.ToTagList(ut.ToStringArray(authMeta["filter"])),
									"label_value":         "tag",
									"pagination":          ct.PaginationTypeBottom,
									"page_size":           5,
									"hide_paginaton_size": true,
									"list_filter":         auth.UserGroup != md.UserGroupAdmin,
									"filter_placeholder":  labels["placeholder_filter"],
									"add_item":            auth.UserGroup != md.UserGroupAdmin,
									"add_icon":            ct.IconFilter,
									"edit_item":           false,
									"delete_item":         true,
									"indicator":           ct.IndicatorSpinner,
								},
								FormTrigger: true,
							},
						},
					}, Full: true, BorderBottom: true},
				},
				FooterRows: footerRows((auth.UserName == ""), (auth.UserName == "admin") || (auth.UserName == "")),
			}
		},
	}

	if frm, found := frmMap[formKey]; found {
		return frm()
	}
	return ct.Form{}
}
