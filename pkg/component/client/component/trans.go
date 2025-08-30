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

func transSideBar(labels cu.SM, data cu.IM) (items []ct.SideBarItem) {
	var trans cu.IM = cu.ToIM(data["trans"], cu.IM{"trans_meta": cu.IM{}})

	dirty := cu.ToBoolean(data["dirty"], false)
	newInput := (cu.ToInteger(trans["id"], 0) == 0)
	updateLabel := labels["editor_save"]
	if newInput {
		updateLabel = labels["editor_create"]
	}
	updateDisabled := func() (disabled bool) {
		return (cu.ToString(trans["trans_name"], "") == "")
	}

	smState := func() *ct.SideBarStatic {
		if cu.ToBoolean(trans["inactive"], false) {
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
			Disabled: newInput,
		},
		&ct.SideBarSeparator{},
		&ct.SideBarElement{
			Name:     "editor_new",
			Value:    "editor_new",
			Label:    labels["transitem_new"],
			Icon:     ct.IconFileText,
			Disabled: newInput || dirty,
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
			Href:       st.DocsClientPath, //+ "/trans",
			LinkTarget: "_blank",
		},
	}
}

func transEditorView(labels cu.SM, data cu.IM) (views []ct.EditorView) {
	var trans cu.IM = cu.ToIM(data["trans"], cu.IM{})
	transMap := cu.ToIM(trans["trans_map"], cu.IM{})
	items := cu.ToIMA(data["items"], []cu.IM{})
	newInput := (cu.ToInteger(trans["id"], 0) == 0)

	if newInput {
		return []ct.EditorView{
			{
				Key:   "trans",
				Label: labels[strings.ToLower(cu.ToString(trans["trans_type"], ""))],
				Icon:  ct.IconFileText,
			},
		}
	}
	return []ct.EditorView{
		{
			Key:   "trans",
			Label: labels[strings.ToLower(cu.ToString(trans["trans_type"], ""))],
			Icon:  ct.IconFileText,
		},
		{
			Key:   "maps",
			Label: labels["map_view"],
			Icon:  ct.IconFileText,
			Badge: cu.ToString(int64(len(transMap)), "0"),
		},
		{
			Key:   "items",
			Label: labels["items_view"],
			Icon:  ct.IconListOl,
			Badge: cu.ToString(int64(len(items)), "0"),
		},
	}
}

func transRow(view string, labels cu.SM, data cu.IM) (rows []ct.Row) {
	if !slices.Contains([]string{"trans", "maps", "items"}, view) {
		return []ct.Row{}
	}

	var trans md.Trans = md.Trans{}
	ut.ConvertToType(data["trans"], &trans)

	configMap := cu.ToIMA(data["config_map"], []cu.IM{})
	currencies := cu.ToIMA(data["currencies"], []cu.IM{})
	items := cu.ToIMA(data["items"], []cu.IM{})
	selectedField := cu.ToString(data["map_field"], "")
	customerSelectorRows := cu.ToIMA(data["customer_selector"], []cu.IM{})
	customerName := cu.ToString(data["customer_name"], "")
	transitemSelectorRows := cu.ToIMA(data["transitem_selector"], []cu.IM{})
	employeeSelectorRows := cu.ToIMA(data["employee_selector"], []cu.IM{})
	projectSelectorRows := cu.ToIMA(data["project_selector"], []cu.IM{})

	mapFieldOpt := func() (opt []ct.SelectOption) {
		opt = []ct.SelectOption{}
		for _, field := range configMap {
			filter := ut.ToStringArray(field["filter"])
			if slices.Contains(filter, "FILTER_TRANS") || len(filter) == 0 {
				if _, ok := trans.TransMap[cu.ToString(field["field_name"], "")]; !ok {
					opt = append(opt, ct.SelectOption{
						Value: cu.ToString(field["field_name"], ""), Text: cu.ToString(field["description"], ""),
					})
				}
			}
		}
		return opt
	}

	transStateOpt := func() (opt []ct.SelectOption) {
		opt = []ct.SelectOption{}
		for _, state := range []md.TransState{
			md.TransStateOK, md.TransStateNew, md.TransStateBack,
		} {
			opt = append(opt, ct.SelectOption{
				Value: state.String(), Text: state.String(),
			})
		}
		return opt
	}

	currencyOpt := func() (opt []ct.SelectOption) {
		opt = []ct.SelectOption{}
		for _, currency := range currencies {
			opt = append(opt, ct.SelectOption{
				Value: cu.ToString(currency["code"], ""), Text: cu.ToString(currency["code"], ""),
			})
		}
		return opt
	}

	directionOpt := func() (opt []ct.SelectOption) {
		opt = []ct.SelectOption{}
		for _, direction := range []md.Direction{
			md.DirectionOut, md.DirectionIn, md.DirectionTransfer,
		} {
			opt = append(opt, ct.SelectOption{
				Value: direction.String(), Text: direction.String(),
			})
		}
		return opt
	}

	paidTypeOpt := func() (opt []ct.SelectOption) {
		opt = []ct.SelectOption{}
		for _, paidType := range []md.PaidType{
			md.PaidTypeOnline, md.PaidTypeCard, md.PaidTypeTransfer, md.PaidTypeCash, md.PaidTypeOther,
		} {
			opt = append(opt, ct.SelectOption{
				Value: paidType.String(), Text: paidType.String(),
			})
		}
		return opt
	}

	var customerSelectorFields []ct.TableField = []ct.TableField{
		{Name: "code", Label: labels["customer_code"]},
		{Name: "customer_name", Label: labels["customer_name"]},
		{Name: "tax_number", Label: labels["customer_tax_number"]},
	}

	var transitemSelectorFields []ct.TableField = []ct.TableField{
		{Name: "code", Label: labels["trans_code"]},
		{Name: "trans_date", Label: labels["trans_date"]},
		{Name: "trans_type", Label: labels["trans_type"]},
		{Name: "direction", Label: labels["trans_direction"]},
	}

	var employeeSelectorFields []ct.TableField = []ct.TableField{
		{Name: "code", Label: labels["employee_code"]},
		{Name: "first_name", Label: labels["contact_first_name"]},
		{Name: "surname", Label: labels["contact_surname"]},
		{Name: "status", Label: labels["contact_status"]},
	}

	var projectSelectorFields []ct.TableField = []ct.TableField{
		{Name: "code", Label: labels["project_code"]},
		{Name: "project_name", Label: labels["project_name"]},
		{Name: "customer_code", Label: labels["customer_code"]},
	}

	typeLabel := func(key string) string {
		return cu.ToString(labels[key+"_"+strings.ToLower(strings.Split(trans.TransType.String(), "_")[1])], labels[key])
	}

	itemTotal := func() (netAmount, vatAmount, amount float64) {
		for _, item := range items {
			netAmount += cu.ToFloat(cu.ToIM(item["item_meta"], cu.IM{})["net_amount"], 0)
			vatAmount += cu.ToFloat(cu.ToIM(item["item_meta"], cu.IM{})["vat_amount"], 0)
			amount += cu.ToFloat(cu.ToIM(item["item_meta"], cu.IM{})["amount"], 0)
		}
		return netAmount, vatAmount, amount
	}

	if view == "maps" {
		return []ct.Row{
			{Columns: []ct.RowColumn{
				{Label: labels["map_fields"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "map_field_" + cu.ToString(trans.Id, ""),
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

	if view == "items" {
		netAmount, vatAmount, amount := itemTotal()
		return []ct.Row{
			{Columns: []ct.RowColumn{
				{
					Label: labels["item_net_amount"],
					Value: ct.Field{
						Type: ct.FieldTypeNumber, Value: cu.IM{
							"name":     "net_amount",
							"value":    netAmount,
							"disabled": true,
							"style": cu.SM{
								"opacity": "1",
							},
						},
					}},
				{
					Label: labels["item_vat_amount"],
					Value: ct.Field{
						Type: ct.FieldTypeNumber, Value: cu.IM{
							"name":     "vat_amount",
							"value":    vatAmount,
							"disabled": true,
							"style": cu.SM{
								"opacity": "1",
							},
						},
					}},
				{
					Label: labels["item_amount"],
					Value: ct.Field{
						Type: ct.FieldTypeNumber, Value: cu.IM{
							"name":     "amount",
							"value":    amount,
							"disabled": true,
							"style": cu.SM{
								"opacity": "1",
							},
						},
					}},
			}, Full: false},
		}
	}

	if slices.Contains([]string{
		md.TransTypeInvoice.String(), md.TransTypeReceipt.String(), md.TransTypeOrder.String(),
		md.TransTypeOffer.String(), md.TransTypeWorksheet.String(), md.TransTypeRent.String()},
		trans.TransType.String(),
	) {
		rows = []ct.Row{{Columns: []ct.RowColumn{
			{Label: labels["trans_code"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "code_" + cu.ToString(trans.Id, ""),
				},
				Type: ct.FieldTypeString, Value: cu.IM{
					"name":     "code",
					"value":    trans.Code,
					"disabled": true,
				},
			}},
			{Label: labels["trans_ref_number"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "ref_number_" + cu.ToString(trans.Id, ""),
				},
				Type: ct.FieldTypeString, Value: cu.IM{
					"name":  "ref_number",
					"value": trans.TransMeta.RefNumber,
				},
			}},
			{Label: labels["trans_state"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "trans_state_" + cu.ToString(trans.Id, ""),
				},
				Type: ct.FieldTypeSelect, Value: cu.IM{
					"name":    "trans_state",
					"options": transStateOpt(),
					"is_null": false,
					"value":   trans.TransMeta.TransState.String(),
				},
			}},
		}, Full: true, BorderBottom: true},
			{Columns: []ct.RowColumn{
				{Label: labels["trans_direction"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "trans_direction_" + cu.ToString(trans.Id, ""),
					},
					Type: ct.FieldTypeSelect, Value: cu.IM{
						"name":     "direction",
						"options":  directionOpt(),
						"is_null":  false,
						"value":    trans.Direction.String(),
						"disabled": (trans.Id > 0),
					},
				}},
				{Label: labels["trans_time_stamp"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "time_stamp_" + cu.ToString(trans.Id, ""),
					},
					Type: ct.FieldTypeDate, Value: cu.IM{
						"name":     "time_stamp",
						"is_null":  false,
						"value":    trans.TimeStamp.String(),
						"disabled": true,
					},
				}},
				{Label: typeLabel("trans_date"),
					Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "trans_date_" + cu.ToString(trans.Id, ""),
						},
						Type: ct.FieldTypeDate, Value: cu.IM{
							"name":    "trans_date",
							"is_null": false,
							"value":   trans.TransDate.String(),
						},
					}},
				{Label: typeLabel("trans_due_time"),
					Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "due_time_" + cu.ToString(trans.Id, ""),
						},
						Type: ct.FieldTypeDate, Value: cu.IM{
							"name":    "due_time",
							"is_null": false,
							"value":   trans.TransMeta.DueTime.String(),
						},
					}},
			}, Full: true, BorderBottom: true}}
		if trans.TransType == md.TransTypeReceipt {
			rows = append(rows, ct.Row{Columns: []ct.RowColumn{
				{Label: labels["trans_paid_type"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "paid_type_" + cu.ToString(trans.Id, ""),
					},
					Type: ct.FieldTypeSelect, Value: cu.IM{
						"name":    "paid_type",
						"options": paidTypeOpt(),
						"is_null": false,
						"value":   trans.TransMeta.PaidType.String(),
					},
				}},
				{
					Label: labels["trans_code"], Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "trans_code_" + cu.ToString(trans.Id, ""),
						},
						Type: ct.FieldTypeSelector, Value: cu.IM{
							"name":  "transitem_code",
							"title": labels["transitem_view"],
							"value": ct.SelectOption{
								Value: trans.TransCode,
								Text:  trans.TransCode,
							},
							"fields":  transitemSelectorFields,
							"rows":    transitemSelectorRows,
							"link":    true,
							"is_null": true,
						},
						FormTrigger: true,
					},
				},
			},
				Full: true, BorderBottom: true},
				ct.Row{Columns: []ct.RowColumn{
					{
						Label: labels["employee_code"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "employee_code_" + cu.ToString(trans.Id, ""),
							},
							Type: ct.FieldTypeSelector, Value: cu.IM{
								"name":  "employee_code",
								"title": labels["employee_view"],
								"value": ct.SelectOption{
									Value: trans.EmployeeCode,
									Text:  trans.EmployeeCode,
								},
								"fields":  employeeSelectorFields,
								"rows":    employeeSelectorRows,
								"link":    true,
								"is_null": true,
							},
							FormTrigger: true,
						},
					},
					{
						Label: labels["project_code"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "project_code_" + cu.ToString(trans.Id, ""),
							},
							Type: ct.FieldTypeSelector, Value: cu.IM{
								"name":  "project_code",
								"title": labels["project_view"],
								"value": ct.SelectOption{
									Value: trans.ProjectCode,
									Text:  trans.ProjectCode,
								},
								"fields":  projectSelectorFields,
								"rows":    projectSelectorRows,
								"link":    true,
								"is_null": true,
							},
							FormTrigger: true,
						},
					},
				},
					Full: true, BorderBottom: true})
		} else {
			rows = append(rows, ct.Row{Columns: []ct.RowColumn{
				{
					Label: labels["customer_name"], Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "customer_code_" + cu.ToString(trans.Id, ""),
						},
						Type: ct.FieldTypeSelector, Value: cu.IM{
							"name":  "customer_code",
							"title": labels["view_customer"],
							"value": ct.SelectOption{
								Value: trans.CustomerCode,
								Text:  customerName,
							},
							"fields":  customerSelectorFields,
							"rows":    customerSelectorRows,
							"link":    true,
							"is_null": true,
						},
						FormTrigger: true,
					},
				},
				{
					Label: labels["trans_code"], Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "trans_code_" + cu.ToString(trans.Id, ""),
						},
						Type: ct.FieldTypeSelector, Value: cu.IM{
							"name":  "transitem_code",
							"title": labels["transitem_view"],
							"value": ct.SelectOption{
								Value: trans.TransCode,
								Text:  trans.TransCode,
							},
							"fields":  transitemSelectorFields,
							"rows":    transitemSelectorRows,
							"link":    true,
							"is_null": true,
						},
						FormTrigger: true,
					},
				},
			},
				Full: true, BorderBottom: true},
				ct.Row{Columns: []ct.RowColumn{
					{
						Label: labels["employee_code"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "employee_code_" + cu.ToString(trans.Id, ""),
							},
							Type: ct.FieldTypeSelector, Value: cu.IM{
								"name":  "employee_code",
								"title": labels["employee_view"],
								"value": ct.SelectOption{
									Value: trans.EmployeeCode,
									Text:  trans.EmployeeCode,
								},
								"fields":  employeeSelectorFields,
								"rows":    employeeSelectorRows,
								"link":    true,
								"is_null": true,
							},
							FormTrigger: true,
						},
					},
					{
						Label: labels["project_code"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "project_code_" + cu.ToString(trans.Id, ""),
							},
							Type: ct.FieldTypeSelector, Value: cu.IM{
								"name":  "project_code",
								"title": labels["project_view"],
								"value": ct.SelectOption{
									Value: trans.ProjectCode,
									Text:  trans.ProjectCode,
								},
								"fields":  projectSelectorFields,
								"rows":    projectSelectorRows,
								"link":    true,
								"is_null": true,
							},
							FormTrigger: true,
						},
					},
					{Label: labels["trans_paid_type"], Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "paid_type_" + cu.ToString(trans.Id, ""),
						},
						Type: ct.FieldTypeSelect, Value: cu.IM{
							"name":    "paid_type",
							"options": paidTypeOpt(),
							"is_null": false,
							"value":   trans.TransMeta.PaidType.String(),
						},
					}},
				},
					Full: true, BorderBottom: true})
		}
		if trans.TransType == md.TransTypeWorksheet {
			rows = append(rows, ct.Row{
				Columns: []ct.RowColumn{{
					Label: labels["trans_worksheet_distance"],
					Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "worksheet_distance_" + cu.ToString(trans.Id, ""),
						},
						Type: ct.FieldTypeNumber, Value: cu.IM{
							"name":  "worksheet_distance",
							"value": trans.TransMeta.Worksheet.Distance,
						},
					}},
					{
						Label: labels["trans_worksheet_repair"],
						Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "worksheet_repair_" + cu.ToString(trans.Id, ""),
							},
							Type: ct.FieldTypeNumber, Value: cu.IM{
								"name":  "worksheet_repair",
								"value": trans.TransMeta.Worksheet.Repair,
							},
						}},
					{
						Label: labels["trans_worksheet_total"],
						Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "worksheet_total_" + cu.ToString(trans.Id, ""),
							},
							Type: ct.FieldTypeNumber, Value: cu.IM{
								"name":  "worksheet_total",
								"value": trans.TransMeta.Worksheet.Total,
							},
						}},
				}, Full: true, BorderBottom: true,
			},
				ct.Row{
					Columns: []ct.RowColumn{
						{
							Label: labels["trans_worksheet_notes"],
							Value: ct.Field{
								BaseComponent: ct.BaseComponent{
									Name: "worksheet_notes_" + cu.ToString(trans.Id, ""),
								},
								Type: ct.FieldTypeString, Value: cu.IM{
									"name":  "worksheet_notes",
									"value": trans.TransMeta.Worksheet.Notes,
								},
							}},
					}, Full: true, BorderBottom: true,
				})
		}
		if trans.TransType == md.TransTypeRent {
			rows = append(rows, ct.Row{
				Columns: []ct.RowColumn{{
					Label: labels["trans_rent_holiday"],
					Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "rent_holiday_" + cu.ToString(trans.Id, ""),
						},
						Type: ct.FieldTypeNumber, Value: cu.IM{
							"name":  "rent_holiday",
							"value": trans.TransMeta.Rent.Holiday,
						},
					}},
					{
						Label: labels["trans_rent_bad_tool"],
						Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "rent_bad_tool_" + cu.ToString(trans.Id, ""),
							},
							Type: ct.FieldTypeNumber, Value: cu.IM{
								"name":  "rent_bad_tool",
								"value": trans.TransMeta.Rent.BadTool,
							},
						}},
					{
						Label: labels["trans_rent_other"],
						Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "rent_other_" + cu.ToString(trans.Id, ""),
							},
							Type: ct.FieldTypeNumber, Value: cu.IM{
								"name":  "rent_other",
								"value": trans.TransMeta.Rent.Other,
							},
						}},
				}, Full: true, BorderBottom: true,
			},
				ct.Row{
					Columns: []ct.RowColumn{
						{
							Label: labels["trans_rent_notes"],
							Value: ct.Field{
								BaseComponent: ct.BaseComponent{
									Name: "rent_notes_" + cu.ToString(trans.Id, ""),
								},
								Type: ct.FieldTypeString, Value: cu.IM{
									"name":  "rent_notes",
									"value": trans.TransMeta.Rent.Notes,
								},
							}},
					}, Full: true, BorderBottom: true,
				})
		}
		rows = append(rows,
			ct.Row{Columns: []ct.RowColumn{
				{Label: labels["currency_code"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "currency_code_" + cu.ToString(trans.Id, ""),
					},
					Type: ct.FieldTypeSelect, Value: cu.IM{
						"name":    "currency_code",
						"options": currencyOpt(),
						"is_null": false,
						"value":   trans.CurrencyCode,
					},
				}},
				{Label: typeLabel("trans_rate"), Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "rate_" + cu.ToString(trans.Id, ""),
					},
					Type: ct.FieldTypeNumber, Value: cu.IM{
						"name":  "rate",
						"value": trans.TransMeta.Rate,
					},
				}},
				{Label: typeLabel("trans_paid"), Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "paid_" + cu.ToString(trans.Id, ""),
					},
					Type: ct.FieldTypeBool, Value: cu.IM{
						"name":  "paid",
						"value": cu.ToBoolean(trans.TransMeta.Paid, false),
					},
				}},
				{Label: labels["trans_closed"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "closed_" + cu.ToString(trans.Id, ""),
					},
					Type: ct.FieldTypeBool, Value: cu.IM{
						"name":  "closed",
						"value": cu.ToBoolean(trans.TransMeta.Closed, false),
					},
				}},
			}, Full: true, BorderBottom: true},
			ct.Row{Columns: []ct.RowColumn{
				{Label: labels["trans_notes"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "notes_" + cu.ToString(trans.Id, ""),
					},
					Type: ct.FieldTypeText, Value: cu.IM{
						"name":  "notes",
						"value": trans.TransMeta.Notes,
						"rows":  4,
					},
				}},
				{
					Label: labels["trans_tags"], Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "tags_" + cu.ToString(trans.Id, ""),
						},
						Type: ct.FieldTypeList, Value: cu.IM{
							"name":                "tags",
							"rows":                ut.ToTagList(trans.TransMeta.Tags),
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
			ct.Row{Columns: []ct.RowColumn{
				{Label: labels["trans_internal_notes"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "internal_notes_" + cu.ToString(trans.Id, ""),
					},
					Type: ct.FieldTypeText, Value: cu.IM{
						"name":  "internal_notes",
						"value": trans.TransMeta.InternalNotes,
						"rows":  4,
					},
				}},
			}, Full: true, BorderBottom: true})
		return rows
	}

	return rows
}

func transTable(view string, labels cu.SM, data cu.IM) []ct.Table {
	if !slices.Contains([]string{"maps", "items"}, view) {
		return []ct.Table{}
	}

	var trans cu.IM = cu.ToIM(data["trans"], cu.IM{})
	newInput := (cu.ToInteger(trans["id"], 0) == 0)
	tblMap := map[string]func() []ct.Table{
		"maps": func() []ct.Table {
			configMap := cu.ToIMA(data["config_map"], []cu.IM{})
			transMap := cu.ToIM(trans["trans_map"], cu.IM{})
			return []ct.Table{
				{
					Fields: []ct.TableField{
						{Name: "description", Label: labels["map_description"], ReadOnly: true},
						{Name: "value", Label: labels["map_value"], FieldType: ct.TableFieldTypeMeta, Required: true},
					},
					Rows:              mapTableRows(transMap, configMap),
					Pagination:        ct.PaginationTypeTop,
					PageSize:          5,
					HidePaginatonSize: true,
					RowSelected:       true,
					TableFilter:       true,
					FilterPlaceholder: labels["placeholder_filter"],
					AddItem:           !newInput && (cu.ToString(data["map_field"], "") != ""),
					LabelAdd:          labels["map_new"],
					Editable:          true,
				},
			}
		},
		"items": func() []ct.Table {
			itemRows := func() []cu.IM {
				rows := []cu.IM{}
				items := cu.ToIMA(data["items"], []cu.IM{})
				for _, item := range items {
					rows = append(rows, cu.IM{
						"product_code": item["product_code"],
						"tax_code":     item["tax_code"],
						"description":  cu.ToIM(item["item_meta"], cu.IM{})["description"],
						"unit":         cu.ToIM(item["item_meta"], cu.IM{})["unit"],
						"qty":          cu.ToFloat(cu.ToIM(item["item_meta"], cu.IM{})["qty"], 0),
						"amount":       cu.ToFloat(cu.ToIM(item["item_meta"], cu.IM{})["amount"], 0),
						"item_meta":    item["item_meta"],
					})
				}
				return rows
			}
			return []ct.Table{
				{
					Fields: []ct.TableField{
						{Name: "description", Label: labels["item_description"]},
						{Name: "unit", Label: labels["item_unit"]},
						{Name: "qty", Label: labels["item_qty"], FieldType: ct.TableFieldTypeNumber},
						{Name: "amount", Label: labels["item_amount"], FieldType: ct.TableFieldTypeNumber},
					},
					Rows:              itemRows(),
					Pagination:        ct.PaginationTypeTop,
					PageSize:          5,
					HidePaginatonSize: true,
					RowSelected:       true,
					TableFilter:       true,
					FilterPlaceholder: labels["placeholder_filter"],
					AddItem:           !newInput,
					LabelAdd:          labels["item_new"],
				},
			}
		},
	}
	return tblMap[view]()
}

func transForm(formKey string, labels cu.SM, data cu.IM) (form ct.Form) {
	formData := cu.ToIM(data, cu.IM{})
	footerRows := func(disabled bool) []ct.Row {
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
							"disabled": disabled,
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
		return rows
	}
	frmMap := map[string]func() ct.Form{
		"items": func() ct.Form {
			var item md.Item = md.Item{}
			ut.ConvertToType(formData, &item)
			taxCodes := cu.ToIMA(data["tax_codes"], []cu.IM{})
			productSelectorRows := cu.ToIMA(data["product_selector"], []cu.IM{})
			taxCodeOpt := func() (opt []ct.SelectOption) {
				opt = []ct.SelectOption{}
				for _, taxCode := range taxCodes {
					opt = append(opt, ct.SelectOption{
						Value: cu.ToString(taxCode["code"], ""), Text: cu.ToString(taxCode["description"], ""),
					})
				}
				return opt
			}
			var productSelectorFields []ct.TableField = []ct.TableField{
				{Name: "code", Label: labels["product_code"]},
				{Name: "product_name", Label: labels["product_name"]},
				{Name: "product_type", Label: labels["product_type"]},
				{Name: "tag_lst", Label: labels["product_tags"]},
			}
			return ct.Form{
				Title: labels["item_view"],
				Icon:  ct.IconListOl,
				BodyRows: []ct.Row{
					{Columns: []ct.RowColumn{
						{
							Label: labels["product_code"], Value: ct.Field{
								BaseComponent: ct.BaseComponent{
									Name: "product_code_" + cu.ToString(item.Id, ""),
								},
								Type: ct.FieldTypeSelector, Value: cu.IM{
									"name":  "product_code",
									"title": labels["view_product"],
									"value": ct.SelectOption{
										Value: item.ProductCode,
										Text:  item.ProductCode,
									},
									"fields":  productSelectorFields,
									"rows":    productSelectorRows,
									"link":    true,
									"is_null": true,
								},
								FormTrigger: true,
							},
						},
					}, Full: true, BorderBottom: true, FieldCol: true},
					{Columns: []ct.RowColumn{
						{Label: labels["item_unit"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "unit",
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":  "unit",
								"value": item.ItemMeta.Unit,
							},
							FormTrigger: true,
						}},
						{Label: labels["item_own_stock"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "own_stock_" + cu.ToString(item.Id, ""),
							},
							Type: ct.FieldTypeNumber, Value: cu.IM{
								"name":  "own_stock",
								"value": item.ItemMeta.OwnStock,
							},
							FormTrigger: true,
						}},
						{Label: labels["item_deposit"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "deposit_" + cu.ToString(item.Id, ""),
							},
							Type: ct.FieldTypeBool, Value: cu.IM{
								"name":  "deposit",
								"value": cu.ToBoolean(item.ItemMeta.Deposit, false),
							},
							FormTrigger: true,
						}},
					}, Full: true, BorderBottom: true},
					{Columns: []ct.RowColumn{
						{Label: labels["item_qty"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "qty_" + cu.ToString(item.Id, ""),
							},
							Type: ct.FieldTypeNumber, Value: cu.IM{
								"name":  "qty",
								"value": item.ItemMeta.Qty,
							},
							FormTrigger: true,
						}},
						{Label: labels["item_discount"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "discount_" + cu.ToString(item.Id, ""),
							},
							Type: ct.FieldTypeInteger, Value: cu.IM{
								"name":      "discount",
								"set_max":   true,
								"max_value": 100,
								"set_min":   true,
								"min_value": 0,
								"value":     cu.ToInteger(item.ItemMeta.Discount, 0),
							},
							FormTrigger: true,
						}},
						{Label: labels["item_fx_price"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "fx_price_" + cu.ToString(item.Id, ""),
							},
							Type: ct.FieldTypeNumber, Value: cu.IM{
								"name":  "fx_price",
								"value": item.ItemMeta.FxPrice,
							},
							FormTrigger: true,
						}},
					}, Full: true, BorderBottom: true},
					{Columns: []ct.RowColumn{
						{Label: labels["item_net_amount"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "net_amount_" + cu.ToString(item.Id, ""),
							},
							Type: ct.FieldTypeNumber, Value: cu.IM{
								"name":  "net_amount",
								"value": item.ItemMeta.NetAmount,
							},
							FormTrigger: true,
						}},
						{Label: labels["tax_code"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "tax_code_" + cu.ToString(item.Id, ""),
							},
							Type: ct.FieldTypeSelect, Value: cu.IM{
								"name":    "tax_code",
								"options": taxCodeOpt(),
								"is_null": false,
								"value":   item.TaxCode,
							},
							FormTrigger: true,
						}},
						{Label: labels["item_amount"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "amount_" + cu.ToString(item.Id, ""),
							},
							Type: ct.FieldTypeNumber, Value: cu.IM{
								"name":  "amount",
								"value": item.ItemMeta.Amount,
							},
							FormTrigger: true,
						}},
					}, Full: true, BorderBottom: true},
					{Columns: []ct.RowColumn{
						{Label: labels["item_description"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "description",
							},
							Type: ct.FieldTypeText, Value: cu.IM{
								"name":  "description",
								"value": item.ItemMeta.Description,
								"rows":  3,
							},
							FormTrigger: true,
						}},
						{
							Label: labels["item_tags"], Value: ct.Field{
								BaseComponent: ct.BaseComponent{
									Name: "tags",
								},
								Type: ct.FieldTypeList, Value: cu.IM{
									"name":                "tags",
									"rows":                ut.ToTagList(item.ItemMeta.Tags),
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
					}, Full: true},
				},
				FooterRows: footerRows(item.ProductCode == ""),
			}
		},
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
				FooterRows: footerRows(false),
			}
		},
	}

	if frm, found := frmMap[formKey]; found {
		return frm()
	}
	return ct.Form{}
}
