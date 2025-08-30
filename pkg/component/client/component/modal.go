package component

import (
	"fmt"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
)

func modalInfoMessage(labels cu.SM, data cu.IM) (form ct.Form) {
	return ct.Form{
		Title: cu.ToString(data["title"], labels["inputbox_info"]),
		Icon:  cu.ToString(data["icon"], ct.IconInfoCircle),
		Modal: true,
		BodyRows: []ct.Row{
			{
				Columns: []ct.RowColumn{
					{Label: cu.ToString(data["info_label"], ""),
						Value: ct.Field{
							Type: ct.FieldTypeLabel,
							Value: cu.IM{
								"value": cu.ToString(data["info_message"], ""),
								"style": cu.SM{
									"font-weight": "normal",
									"font-style":  "italic",
								},
							},
						}},
				},
				Full:         false,
				FieldCol:     false,
				BorderTop:    false,
				BorderBottom: false,
			},
		},
		FooterRows: []ct.Row{
			{
				Columns: []ct.RowColumn{
					{Value: ct.Field{
						Type:  ct.FieldTypeLabel,
						Value: cu.IM{},
					}},
					{Value: ct.Field{
						Type: ct.FieldTypeButton,
						Value: cu.IM{
							"name":         ct.FormEventOK,
							"type":         ct.ButtonTypeSubmit,
							"button_style": ct.ButtonStylePrimary,
							"icon":         ct.IconCheck,
							"label":        cu.ToString(data["submit_label"], labels["inputbox_ok"]),
							"auto_focus":   true,
							"selected":     true,
						},
					}},
					{Value: ct.Field{
						Type:  ct.FieldTypeLabel,
						Value: cu.IM{},
					}},
				},
				Full:         true,
				FieldCol:     false,
				BorderTop:    false,
				BorderBottom: false,
			},
		},
	}
}

func modalWarningMessage(labels cu.SM, data cu.IM) (form ct.Form) {
	return ct.Form{
		Title: cu.ToString(data["title"], labels["inputbox_warning"]),
		Icon:  cu.ToString(data["icon"], ct.IconExclamationTriangle),
		Modal: true,
		BodyRows: []ct.Row{
			{
				Columns: []ct.RowColumn{
					{Label: cu.ToString(data["warning_label"], ""),
						Value: ct.Field{
							Type: ct.FieldTypeLabel,
							Value: cu.IM{
								"value": cu.ToString(data["warning_message"], ""),
								"style": cu.SM{
									"font-weight": "normal",
									"font-style":  "italic",
								},
							},
						}},
				},
				Full:         false,
				FieldCol:     false,
				BorderTop:    false,
				BorderBottom: false,
			},
		},
		FooterRows: []ct.Row{
			{
				Columns: []ct.RowColumn{
					{Value: ct.Field{
						Type: ct.FieldTypeButton,
						Value: cu.IM{
							"name":         ct.FormEventOK,
							"type":         ct.ButtonTypeSubmit,
							"button_style": ct.ButtonStylePrimary,
							"icon":         ct.IconCheck,
							"label":        cu.ToString(data["submit_label"], labels["inputbox_ok"]),
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
							"icon":         ct.IconTimes,
							"label":        cu.ToString(data["cancel_label"], labels["inputbox_cancel"]),
						},
					}},
				},
				Full:         true,
				FieldCol:     false,
				BorderTop:    false,
				BorderBottom: false,
			},
		},
	}
}

func modalInputString(labels cu.SM, data cu.IM) (form ct.Form) {
	return ct.Form{
		Title: cu.ToString(data["title"], labels["inputbox_string"]),
		Icon:  cu.ToString(data["icon"], ct.IconEdit),
		Modal: true,
		BodyRows: []ct.Row{
			{
				Columns: []ct.RowColumn{
					{Label: cu.ToString(data["label"], ""),
						Value: ct.Field{
							Type: ct.FieldTypeString,
							Value: cu.IM{
								"name":        cu.ToString(data["field_name"], "value"),
								"value":       cu.ToString(data["default_value"], ""),
								"label":       cu.ToString(data["label"], ""),
								"placeholder": cu.ToString(data["placeholder"], ""),
								"required":    cu.ToBoolean(data["required"], false),
								"invalid":     cu.ToBoolean(data["invalid"], false),
								"auto_focus":  true,
							},
						}},
				},
				Full:         true,
				FieldCol:     false,
				BorderTop:    false,
				BorderBottom: false,
			},
		},
		FooterRows: []ct.Row{
			{
				Columns: []ct.RowColumn{
					{Value: ct.Field{
						Type: ct.FieldTypeButton,
						Value: cu.IM{
							"name":         ct.FormEventOK,
							"type":         ct.ButtonTypeSubmit,
							"button_style": ct.ButtonStylePrimary,
							"icon":         ct.IconCheck,
							"label":        cu.ToString(data["submit_label"], labels["inputbox_ok"]),
						},
					}},
					{Value: ct.Field{
						Type: ct.FieldTypeButton,
						Value: cu.IM{
							"name":         ct.FormEventCancel,
							"type":         ct.ButtonTypeSubmit,
							"button_style": ct.ButtonStyleDefault,
							"icon":         ct.IconTimes,
							"label":        cu.ToString(data["cancel_label"], labels["inputbox_cancel"]),
						},
					}},
				},
				Full:         true,
				FieldCol:     false,
				BorderTop:    false,
				BorderBottom: false,
			},
		},
	}
}

func modalReport(labels cu.SM, data cu.IM) (form ct.Form) {
	configData := cu.ToIMA(data["config_data"], []cu.IM{})
	configReport := cu.ToIMA(data["config_report"], []cu.IM{})
	fromReport := func() (options []ct.SelectOption) {
		options = []ct.SelectOption{}
		for _, config := range configReport {
			options = append(options, ct.SelectOption{
				Value: cu.ToString(config["report_key"], ""),
				Text:  cu.ToString(config["report_name"], ""),
			})
		}
		return options
	}
	orientations := fromConfig("orientation", configData)
	sizes := fromConfig("paper_size", configData)
	templates := fromReport()
	return ct.Form{
		Title: cu.ToString(data["title"], labels["editor_report"]),
		Icon:  cu.ToString(data["icon"], ct.IconEdit),
		Modal: true,
		BodyRows: []ct.Row{
			{
				Columns: []ct.RowColumn{
					{Label: labels["report_template"],
						Value: ct.Field{
							Type: ct.FieldTypeSelect,
							Value: cu.IM{
								"name":    "template",
								"options": templates,
								"value":   cu.ToString(data["template"], ""),
								"is_null": false,
							},
							FormTrigger: true,
						}},
				},
				Full:         false,
				FieldCol:     false,
				BorderTop:    false,
				BorderBottom: false,
			},
			{
				Columns: []ct.RowColumn{
					{Label: labels["report_orientation"],
						Value: ct.Field{
							Type: ct.FieldTypeSelect,
							Value: cu.IM{
								"name":    "orientation",
								"options": orientations,
								"value":   cu.ToString(data["orientation"], ""),
								"is_null": false,
							},
							FormTrigger: true,
						}},
					{Label: labels["report_size"],
						Value: ct.Field{
							Type: ct.FieldTypeSelect,
							Value: cu.IM{
								"name":    "paper_size",
								"options": sizes,
								"value":   cu.ToString(data["paper_size"], ""),
								"is_null": false,
							},
							FormTrigger: true,
						}},
					{Label: labels["report_copy"],
						Value: ct.Field{
							Type: ct.FieldTypeInteger,
							Value: cu.IM{
								"name":  "copy",
								"value": cu.ToInteger(data["copy"], 1),
							},
							FormTrigger: true,
						}},
				},
				Full:         false,
				FieldCol:     false,
				BorderTop:    false,
				BorderBottom: false,
			},
		},
		FooterRows: []ct.Row{
			{
				Columns: []ct.RowColumn{
					{Value: ct.Field{
						Type: ct.FieldTypeUrlLink,
						Value: cu.IM{
							"name":        "url_pdf",
							"link_style":  ct.LinkStylePrimary,
							"label":       labels["report_print"],
							"icon":        ct.IconPrint,
							"href":        cu.ToString(data["url_pdf"], ""),
							"link_target": "_blank",
							"disabled":    (len(templates) == 0),
						},
					}},
					{Value: ct.Field{
						Type: ct.FieldTypeUrlLink,
						Value: cu.IM{
							"name":        "url_export",
							"link_style":  ct.LinkStylePrimary,
							"label":       labels["report_export"],
							"icon":        ct.IconDownload,
							"href":        cu.ToString(data["url_export"], ""),
							"download":    fmt.Sprintf("%s.pdf", cu.ToString(data["title"], labels["editor_report"])),
							"link_target": "_blank",
							"disabled":    (len(templates) == 0),
						},
					}},
				},
				Full:         true,
				FieldCol:     false,
				BorderTop:    false,
				BorderBottom: false,
			},
			{
				Columns: []ct.RowColumn{
					/*
						{Value: ct.Field{
							Type: ct.FieldTypeUrlLink,
							Value: cu.IM{
								"name":        "url_xml",
								"link_style":  ct.LinkStylePrimary,
								"label":       labels["report_xml"],
								"icon":        ct.IconCode,
								"href":        cu.ToString(data["url_xml"], ""),
								"download":    fmt.Sprintf("%s.xml", cu.ToString(data["title"], labels["editor_report"])),
								"link_target": "_blank",
								"disabled":    (len(templates) == 0),
							},
						}},
					*/
					{Value: ct.Field{
						Type: ct.FieldTypeButton,
						Value: cu.IM{
							"name":         "queue",
							"type":         ct.ButtonTypeButton,
							"button_style": ct.ButtonStylePrimary,
							"icon":         ct.IconDatabase,
							"label":        labels["report_queue"],
							"disabled":     (len(templates) == 0),
						},
						FormTrigger: true,
					}},
					{Value: ct.Field{
						Type: ct.FieldTypeButton,
						Value: cu.IM{
							"name":         ct.FormEventCancel,
							"type":         ct.ButtonTypeSubmit,
							"button_style": ct.ButtonStyleDefault,
							"icon":         ct.IconTimes,
							"label":        cu.ToString(data["cancel_label"], labels["inputbox_cancel"]),
						},
					}},
				},
				Full:         true,
				FieldCol:     false,
				BorderTop:    false,
				BorderBottom: false,
			},
		},
	}
}

func modalSelector(labels cu.SM, data cu.IM) (form ct.Form) {
	return ct.Form{
		Title: cu.ToString(data["title"], labels["inputbox_string"]),
		Icon:  cu.ToString(data["icon"], ct.IconSearch),
		Modal: true,
		BodyRows: []ct.Row{
			{
				Columns: []ct.RowColumn{
					{Value: ct.Field{
						Type: ct.FieldTypeList,
						Value: cu.IM{
							"name":                cu.ToString(data["next"], "selector"),
							"rows":                cu.ToIMA(data["rows"], []cu.IM{}),
							"pagination":          ct.PaginationTypeBottom,
							"page_size":           5,
							"hide_paginaton_size": true,
							"edit_item":           cu.ToBoolean(data["edit_item"], true),
							"edit_icon":           cu.ToString(data["edit_icon"], ct.IconEdit),
							"list_filter":         cu.ToBoolean(data["list_filter"], true),
							"delete_item":         cu.ToBoolean(data["delete_item"], false),
						},
						FormTrigger: true,
					}},
				},
				Full:         true,
				FieldCol:     false,
				BorderTop:    false,
				BorderBottom: false,
			},
		},
		FooterRows: []ct.Row{},
	}
}
