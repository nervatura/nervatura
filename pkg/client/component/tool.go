package component

import (
	"slices"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
	st "github.com/nervatura/nervatura/v6/static"
)

type ToolEditor struct{}

func (e *ToolEditor) Frame(labels cu.SM, data cu.IM) (title, icon string) {
	return cu.ToString(data["editor_title"], labels["tool_title"]),
		cu.ToString(data["editor_icon"], ct.IconWrench)
}

func (e *ToolEditor) SideBar(labels cu.SM, data cu.IM) (items []ct.SideBarItem) {
	var tool cu.IM = cu.ToIM(data["tool"], cu.IM{"tool_meta": cu.IM{}})
	user := cu.ToIM(data["user"], cu.IM{})

	readonly := (cu.ToString(user["user_group"], "") == md.UserGroupGuest.String())
	dirty := cu.ToBoolean(data["dirty"], false)
	newInput := (cu.ToInteger(tool["id"], 0) == 0)
	updateLabel := labels["editor_save"]
	if newInput {
		updateLabel = labels["editor_create"]
	}
	updateDisabled := func() (disabled bool) {
		return (cu.ToString(tool["description"], "") == "") || (cu.ToString(tool["product_code"], "") == "") || readonly
	}

	smState := func() *ct.SideBarStatic {
		if cu.ToBoolean(tool["inactive"], false) {
			return &ct.SideBarStatic{
				Icon: ct.IconLock, Label: labels["state_closed"], Color: "brown",
			}
		}
		if newInput {
			return &ct.SideBarStatic{
				Icon: ct.IconPlus, Label: labels["state_new"], Color: "yellow",
			}
		}
		return &ct.SideBarStatic{
			Icon: ct.IconEdit, Label: labels["state_edit"], Color: "green",
		}
	}

	return []ct.SideBarItem{
		&ct.SideBarSeparator{},
		&ct.SideBarElement{
			Name:    "editor_cancel",
			Value:   "editor_cancel",
			Label:   labels["browser_title"],
			Icon:    ct.IconReply,
			NotFull: true,
		},
		&ct.SideBarSeparator{},
		smState(),
		&ct.SideBarSeparator{},
		&ct.SideBarElement{
			Name:     "editor_save",
			Value:    "editor_save",
			Label:    updateLabel,
			Icon:     ct.IconUpload,
			Selected: dirty,
			Disabled: updateDisabled(),
		},
		&ct.SideBarElement{
			Name:     "editor_delete",
			Value:    "editor_delete",
			Label:    labels["editor_delete"],
			Icon:     ct.IconTimes,
			Disabled: newInput || readonly,
		},
		&ct.SideBarSeparator{},
		&ct.SideBarElement{
			Name:     "editor_new",
			Value:    "editor_new",
			Label:    labels["tool_new"],
			Icon:     ct.IconUser,
			Disabled: newInput || dirty || readonly,
		},
		&ct.SideBarSeparator{},
		&ct.SideBarElement{
			Name:     "editor_report",
			Value:    "editor_report",
			Label:    labels["editor_report"],
			Icon:     ct.IconChartBar,
			Disabled: newInput || dirty,
		},
		&ct.SideBarSeparator{},
		&ct.SideBarElement{
			Name:     "editor_bookmark",
			Value:    "editor_bookmark",
			Label:    labels["editor_bookmark"],
			Icon:     ct.IconStar,
			Disabled: newInput,
		},
		&ct.SideBarElementLink{
			SideBarElement: ct.SideBarElement{
				Name:  "editor_help",
				Value: "editor_help",
				Label: labels["editor_help"],
				Icon:  ct.IconQuestionCircle,
			},
			Href:       st.DocsClientPath + "#tool",
			LinkTarget: "_blank",
		},
	}
}

func (e *ToolEditor) View(labels cu.SM, data cu.IM) (views []ct.EditorView) {
	var tool cu.IM = cu.ToIM(data["tool"], cu.IM{})
	toolMap := cu.ToIM(tool["tool_map"], cu.IM{})
	event := cu.ToIMA(tool["events"], []cu.IM{})
	newInput := (cu.ToInteger(tool["id"], 0) == 0)

	if newInput {
		return []ct.EditorView{
			{
				Key:   "tool",
				Label: labels["tool_view"],
				Icon:  ct.IconWrench,
			},
		}
	}
	return []ct.EditorView{
		{
			Key:   "tool",
			Label: labels["tool_view"],
			Icon:  ct.IconWrench,
		},
		{
			Key:   "maps",
			Label: labels["map_view"],
			Icon:  ct.IconDatabase,
			Badge: cu.ToString(int64(len(toolMap)), "0"),
		},
		{
			Key:   "events",
			Label: labels["event_view"],
			Icon:  ct.IconCalendar,
			Badge: cu.ToString(int64(len(event)), "0"),
		},
	}
}

func (e *ToolEditor) Row(view string, labels cu.SM, data cu.IM) (rows []ct.Row) {
	if !slices.Contains([]string{"tool", "maps"}, view) {
		return []ct.Row{}
	}

	var tool md.Tool = md.Tool{}
	ut.ConvertToType(data["tool"], &tool)

	configMap := cu.ToIMA(data["config_map"], []cu.IM{})
	selectedField := cu.ToString(data["map_field"], "")
	productSelectorRows := cu.ToIMA(data["product_selector"], []cu.IM{})

	var productSelectorFields []ct.TableField = []ct.TableField{
		{Name: "code", Label: labels["product_code"]},
		{Name: "product_name", Label: labels["product_name"]},
		{Name: "product_type", Label: labels["product_type"]},
		{Name: "tag_lst", Label: labels["product_tags"]},
	}

	mapFieldOpt := func() (opt []ct.SelectOption) {
		opt = []ct.SelectOption{}
		for _, field := range configMap {
			filter := ut.ToStringArray(field["filter"])
			if slices.Contains(filter, "FILTER_TOOL") || len(filter) == 0 {
				if _, ok := tool.ToolMap[cu.ToString(field["field_name"], "")]; !ok {
					opt = append(opt, ct.SelectOption{
						Value: cu.ToString(field["field_name"], ""), Text: cu.ToString(field["description"], ""),
					})
				}
			}
		}
		return opt
	}

	if view == "maps" {
		return []ct.Row{
			{Columns: []ct.RowColumn{
				{Label: labels["map_fields"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "map_field_" + cu.ToString(tool.Id, ""),
					},
					Type: ct.FieldTypeSelect, Value: cu.IM{
						"name":    "map_field",
						"options": mapFieldOpt(),
						"is_null": true,
						"value":   selectedField,
					},
				}},
			}, Full: false, FieldCol: true},
		}
	}

	return []ct.Row{
		{Columns: []ct.RowColumn{
			{Label: labels["tool_serial_number"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "serial_number_" + cu.ToString(tool.Id, ""),
				},
				Type: ct.FieldTypeString, Value: cu.IM{
					"name":  "serial_number",
					"value": tool.ToolMeta.SerialNumber,
				},
			}},
			{Label: labels["tool_description"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "description_" + cu.ToString(tool.Id, ""),
				},
				Type: ct.FieldTypeString, Value: cu.IM{
					"name":        "description",
					"invalid":     (tool.Description == ""),
					"placeholder": labels["mandatory_data"],
					"value":       tool.Description,
				},
			}},
		}, Full: true, BorderBottom: true},
		{Columns: []ct.RowColumn{
			{Label: labels["tool_code"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "code_" + cu.ToString(tool.Id, ""),
				},
				Type: ct.FieldTypeString, Value: cu.IM{
					"name":     "code",
					"value":    tool.Code,
					"disabled": true,
				},
			}},
			{
				Label: labels["product_code"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "product_code_" + cu.ToString(tool.Id, ""),
					},
					Type: ct.FieldTypeSelector, Value: cu.IM{
						"name":  "product_code",
						"title": labels["view_product"],
						"value": ct.SelectOption{
							Value: tool.ProductCode,
							Text:  tool.ProductCode,
						},
						"fields":  productSelectorFields,
						"rows":    productSelectorRows,
						"link":    true,
						"is_null": false,
					},
					FormTrigger: true,
				},
			},
			{Label: labels["tool_inactive"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "inactive_" + cu.ToString(tool.Id, ""),
				},
				Type: ct.FieldTypeBool, Value: cu.IM{
					"name":  "inactive",
					"value": cu.ToBoolean(tool.ToolMeta.Inactive, false),
				},
			}},
		}, Full: true, BorderBottom: true},
		{Columns: []ct.RowColumn{
			{Label: labels["tool_notes"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "notes_" + cu.ToString(tool.Id, ""),
				},
				Type: ct.FieldTypeText, Value: cu.IM{
					"name":  "notes",
					"value": tool.ToolMeta.Notes,
					"rows":  4,
				},
			}},
			{
				Label: labels["tool_tags"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "tags_" + cu.ToString(tool.Id, ""),
					},
					Type: ct.FieldTypeList, Value: cu.IM{
						"name":                "tags",
						"rows":                ut.ToTagList(tool.ToolMeta.Tags),
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
				},
			},
		}, Full: true, BorderBottom: true},
	}
}

func (e *ToolEditor) Table(view string, labels cu.SM, data cu.IM) []ct.Table {
	if !slices.Contains([]string{"maps", "events"}, view) {
		return []ct.Table{}
	}

	var tool cu.IM = cu.ToIM(data["tool"], cu.IM{})
	newInput := (cu.ToInteger(tool["id"], 0) == 0)
	tblMap := map[string]func() []ct.Table{
		"maps": func() []ct.Table {
			configMap := cu.ToIMA(data["config_map"], []cu.IM{})
			toolMap := cu.ToIM(tool["tool_map"], cu.IM{})
			return []ct.Table{
				{
					Fields: []ct.TableField{
						{Name: "description", Label: labels["map_description"], ReadOnly: true},
						{Name: "value", Label: labels["map_value"], FieldType: ct.TableFieldTypeMeta, Required: true},
					},
					Rows:              mapTableRows(toolMap, configMap),
					Pagination:        ct.PaginationTypeTop,
					PageSize:          5,
					HidePaginatonSize: true,
					RowSelected:       true,
					TableFilter:       true,
					FilterPlaceholder: labels["placeholder_filter"],
					AddItem:           !newInput && (cu.ToString(data["map_field"], "") != ""),
					LabelAdd:          labels["map_new"],
					Editable:          true,
					Unsortable:        true,
				},
			}
		},
		"events": func() []ct.Table {
			event := cu.ToIMA(tool["events"], []cu.IM{})
			return []ct.Table{
				{
					Fields: []ct.TableField{
						{Name: "subject", Label: labels["event_subject"]},
						{Name: "start_time", Label: labels["event_start_time"], FieldType: ct.TableFieldTypeDateTime},
						{Name: "end_time", Label: labels["event_end_time"], FieldType: ct.TableFieldTypeDateTime},
						{Name: "place", Label: labels["event_place"]},
						{Name: "description", Label: labels["event_description"]},
						//{Name: "tag_lst", Label: labels["event_tags"]},
					},
					Rows:              event,
					Pagination:        ct.PaginationTypeTop,
					PageSize:          5,
					HidePaginatonSize: true,
					RowSelected:       true,
					TableFilter:       true,
					FilterPlaceholder: labels["placeholder_filter"],
					AddItem:           !newInput,
					LabelAdd:          labels["event_new"],
				},
			}
		},
	}
	return tblMap[view]()
}

func (e *ToolEditor) Form(formKey string, labels cu.SM, data cu.IM) (form ct.Form) {
	formData := cu.ToIM(data, cu.IM{})
	footerRows := []ct.Row{
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
						"auto_focus":   true,
						"selected":     true,
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
					},
				}},
			},
			Full:         true,
			FieldCol:     false,
			BorderTop:    false,
			BorderBottom: false,
		},
	}
	frmMap := map[string]func() ct.Form{
		"events": func() ct.Form {
			var event md.Event = md.Event{}
			ut.ConvertToType(formData, &event)
			return ct.Form{
				Title: labels["event_view"],
				Icon:  ct.IconCalendar,
				BodyRows: []ct.Row{
					{Columns: []ct.RowColumn{
						{Label: labels["event_subject"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "subject",
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":  "subject",
								"value": event.Subject,
							},
							FormTrigger: true,
						}},
					}, Full: true, BorderBottom: true, FieldCol: true},
					{Columns: []ct.RowColumn{
						{Label: labels["event_start_time"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "start_time",
							},
							Type: ct.FieldTypeDateTime, Value: cu.IM{
								"name":    "start_time",
								"value":   event.StartTime,
								"is_null": false,
							},
							FormTrigger: true,
						}},
						{Label: labels["event_end_time"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "end_time",
							},
							Type: ct.FieldTypeDateTime, Value: cu.IM{
								"name":    "end_time",
								"is_null": true,
								"value":   event.EndTime,
							},
							FormTrigger: true,
						}},
						{Label: labels["event_place"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "place",
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":  "place",
								"value": event.Place,
							},
							FormTrigger: true,
						}},
					}, Full: true, BorderBottom: true},
					{Columns: []ct.RowColumn{
						{Label: labels["event_description"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "description",
							},
							Type: ct.FieldTypeText, Value: cu.IM{
								"name":  "description",
								"value": event.Description,
								"rows":  4,
							},
							FormTrigger: true,
						}},
						{
							Label: labels["event_tags"], Value: ct.Field{
								BaseComponent: ct.BaseComponent{
									Name: "tags",
								},
								Type: ct.FieldTypeList, Value: cu.IM{
									"name":                "tags",
									"rows":                ut.ToTagList(event.Tags),
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
					}, Full: true, BorderBottom: true},
				},
				FooterRows: footerRows,
			}
		},
	}

	if frm, found := frmMap[formKey]; found {
		return frm()
	}
	return ct.Form{}
}
