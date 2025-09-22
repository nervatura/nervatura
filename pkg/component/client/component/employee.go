package component

import (
	"slices"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
	st "github.com/nervatura/nervatura/v6/pkg/static"
)

func employeeSideBar(labels cu.SM, data cu.IM) (items []ct.SideBarItem) {
	var employee cu.IM = cu.ToIM(data["employee"], cu.IM{"employee_meta": cu.IM{}})
	user := cu.ToIM(data["user"], cu.IM{})

	dirty := cu.ToBoolean(data["dirty"], false)
	readonly := (cu.ToString(user["user_group"], "") == md.UserGroupGuest.String())
	newInput := (cu.ToInteger(employee["id"], 0) == 0)
	updateLabel := labels["editor_save"]
	if newInput {
		updateLabel = labels["editor_create"]
	}

	smState := func() *ct.SideBarStatic {
		if cu.ToBoolean(employee["inactive"], false) {
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
			Disabled: readonly,
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
			Label:    labels["employee_new"],
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
			Name:  "editor_bookmark",
			Value: "editor_bookmark",
			Label: labels["editor_bookmark"],
			Icon:  ct.IconStar,
		},
		&ct.SideBarElementLink{
			SideBarElement: ct.SideBarElement{
				Name:  "editor_help",
				Value: "editor_help",
				Label: labels["editor_help"],
				Icon:  ct.IconQuestionCircle,
			},
			Href:       st.DocsClientPath, //+ "/employee",
			LinkTarget: "_blank",
		},
	}
}

func employeeEditorView(labels cu.SM, data cu.IM) (views []ct.EditorView) {
	var employee cu.IM = cu.ToIM(data["employee"], cu.IM{})
	employeeMap := cu.ToIM(employee["employee_map"], cu.IM{})
	event := cu.ToIMA(employee["events"], []cu.IM{})
	newInput := (cu.ToInteger(employee["id"], 0) == 0)

	if newInput {
		return []ct.EditorView{
			{
				Key:   "employee",
				Label: labels["employee_view"],
				Icon:  ct.IconClock,
			},
		}
	}
	return []ct.EditorView{
		{
			Key:   "employee",
			Label: labels["employee_view"],
			Icon:  ct.IconClock,
		},
		{
			Key:   "contact",
			Label: labels["contact_view"],
			Icon:  ct.IconMobile,
		},
		{
			Key:   "address",
			Label: labels["address_view"],
			Icon:  ct.IconHome,
		},
		{
			Key:   "maps",
			Label: labels["map_view"],
			Icon:  ct.IconClock,
			Badge: cu.ToString(int64(len(employeeMap)), "0"),
		},
		{
			Key:   "events",
			Label: labels["event_view"],
			Icon:  ct.IconCalendar,
			Badge: cu.ToString(int64(len(event)), "0"),
		},
	}
}

func employeeRow(view string, labels cu.SM, data cu.IM) (rows []ct.Row) {
	if !slices.Contains([]string{"employee", "maps", "address", "contact"}, view) {
		return []ct.Row{}
	}

	var employee md.Employee = md.Employee{}
	ut.ConvertToType(data["employee"], &employee)

	var address md.Address = md.Address{}
	ut.ConvertToType(cu.ToIM(data["employee"], cu.IM{})["address"], &address)

	var contact md.Contact = md.Contact{}
	ut.ConvertToType(cu.ToIM(data["employee"], cu.IM{})["contact"], &contact)

	configMap := cu.ToIMA(data["config_map"], []cu.IM{})
	selectedField := cu.ToString(data["map_field"], "")

	mapFieldOpt := func() (opt []ct.SelectOption) {
		opt = []ct.SelectOption{}
		for _, field := range configMap {
			filter := ut.ToStringArray(field["filter"])
			if slices.Contains(filter, "FILTER_EMPLOYEE") || len(filter) == 0 {
				if _, ok := employee.EmployeeMap[cu.ToString(field["field_name"], "")]; !ok {
					opt = append(opt, ct.SelectOption{
						Value: cu.ToString(field["field_name"], ""), Text: cu.ToString(field["description"], ""),
					})
				}
			}
		}
		return opt
	}

	if view == "address" {
		return []ct.Row{
			{Columns: []ct.RowColumn{
				{Label: labels["address_country"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "country",
					},
					Type: ct.FieldTypeString, Value: cu.IM{
						"name":  "country",
						"value": address.Country,
					},
				}},
				{Label: labels["address_state"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "state",
					},
					Type: ct.FieldTypeString, Value: cu.IM{
						"name":  "state",
						"value": address.State,
					},
				}},
				{Label: labels["address_zip_code"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "zip_code",
					},
					Type: ct.FieldTypeString, Value: cu.IM{
						"name":  "zip_code",
						"value": address.ZipCode,
					},
				}},
			}, Full: true, BorderBottom: true},
			{Columns: []ct.RowColumn{
				{Label: labels["address_city"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "city",
					},
					Type: ct.FieldTypeString, Value: cu.IM{
						"name":  "city",
						"value": address.City,
					},
				}},
				{Label: labels["address_street"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "street",
					},
					Type: ct.FieldTypeString, Value: cu.IM{
						"name":  "street",
						"value": address.Street,
					},
				}},
			}, Full: true, BorderBottom: true},
			{Columns: []ct.RowColumn{
				{Label: labels["address_notes"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "address_notes",
					},
					Type: ct.FieldTypeText, Value: cu.IM{
						"name":  "address_notes",
						"value": address.Notes,
					},
				}},
			}, Full: true, BorderBottom: true},
		}
	}

	if view == "contact" {
		return []ct.Row{
			{Columns: []ct.RowColumn{
				{Label: labels["contact_first_name"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "first_name",
					},
					Type: ct.FieldTypeString, Value: cu.IM{
						"name":  "first_name",
						"value": contact.FirstName,
					},
				}},
				{Label: labels["contact_surname"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "surname",
					},
					Type: ct.FieldTypeString, Value: cu.IM{
						"name":  "surname",
						"value": contact.Surname,
					},
				}},
				{Label: labels["contact_status"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "status",
					},
					Type: ct.FieldTypeString, Value: cu.IM{
						"name":  "status",
						"value": contact.Status,
					},
				}},
			}, Full: true, BorderBottom: true},
			{Columns: []ct.RowColumn{
				{Label: labels["contact_phone"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "phone",
					},
					Type: ct.FieldTypeString, Value: cu.IM{
						"name":  "phone",
						"value": contact.Phone,
					},
				}},
				{Label: labels["contact_mobile"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "mobile",
					},
					Type: ct.FieldTypeString, Value: cu.IM{
						"name":  "mobile",
						"value": contact.Mobile,
					},
				}},
				{Label: labels["contact_email"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "email",
					},
					Type: ct.FieldTypeString, Value: cu.IM{
						"name":  "email",
						"value": contact.Email,
					},
				}},
			}, Full: true, BorderBottom: true},
			{Columns: []ct.RowColumn{
				{Label: labels["contact_notes"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "contact_notes",
					},
					Type: ct.FieldTypeText, Value: cu.IM{
						"name":  "contact_notes",
						"value": contact.Notes,
					},
				}},
			}, Full: true, BorderBottom: true},
		}
	}

	if view == "maps" {
		return []ct.Row{
			{Columns: []ct.RowColumn{
				{Label: labels["map_fields"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "map_field_" + cu.ToString(employee.Id, ""),
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
			{Label: labels["employee_code"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "code_" + cu.ToString(employee.Id, ""),
				},
				Type: ct.FieldTypeString, Value: cu.IM{
					"name":     "code",
					"value":    employee.Code,
					"disabled": true,
				},
			}},
			{Label: labels["employee_start_date"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "start_date",
				},
				Type: ct.FieldTypeDate, Value: cu.IM{
					"name":    "start_date",
					"is_null": true,
					"value":   employee.EmployeeMeta.StartDate,
				},
			}},
			{Label: labels["employee_end_date"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "end_date",
				},
				Type: ct.FieldTypeDate, Value: cu.IM{
					"name":    "end_date",
					"is_null": true,
					"value":   employee.EmployeeMeta.EndDate,
				},
			}},
			{Label: labels["employee_inactive"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "inactive_" + cu.ToString(employee.Id, ""),
				},
				Type: ct.FieldTypeBool, Value: cu.IM{
					"name":  "inactive",
					"value": cu.ToBoolean(employee.EmployeeMeta.Inactive, false),
				},
			}},
		}, Full: true, BorderBottom: true},
		{Columns: []ct.RowColumn{
			{Label: labels["employee_notes"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "notes_" + cu.ToString(employee.Id, ""),
				},
				Type: ct.FieldTypeText, Value: cu.IM{
					"name":  "notes",
					"value": employee.EmployeeMeta.Notes,
					"rows":  4,
				},
			}},
			{
				Label: labels["employee_tags"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "tags_" + cu.ToString(employee.Id, ""),
					},
					Type: ct.FieldTypeList, Value: cu.IM{
						"name":                "tags",
						"rows":                ut.ToTagList(employee.EmployeeMeta.Tags),
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

func employeeTable(view string, labels cu.SM, data cu.IM) []ct.Table {
	if !slices.Contains([]string{"maps", "events"}, view) {
		return []ct.Table{}
	}

	var employee cu.IM = cu.ToIM(data["employee"], cu.IM{})
	newInput := (cu.ToInteger(employee["id"], 0) == 0)
	tblMap := map[string]func() []ct.Table{
		"maps": func() []ct.Table {
			configMap := cu.ToIMA(data["config_map"], []cu.IM{})
			employeeMap := cu.ToIM(employee["employee_map"], cu.IM{})
			return []ct.Table{
				{
					Fields: []ct.TableField{
						{Name: "description", Label: labels["map_description"], ReadOnly: true},
						{Name: "value", Label: labels["map_value"], FieldType: ct.TableFieldTypeMeta, Required: true},
					},
					Rows:              mapTableRows(employeeMap, configMap),
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
			event := cu.ToIMA(employee["events"], []cu.IM{})
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

func employeeForm(formKey string, labels cu.SM, data cu.IM) (form ct.Form) {
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
								"value":   event.StartTime.String(),
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
								"value":   event.EndTime.String(),
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
