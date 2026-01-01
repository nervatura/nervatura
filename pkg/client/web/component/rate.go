package component

import (
	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
	st "github.com/nervatura/nervatura/v6/pkg/static"
)

type RateEditor struct{}

func (e *RateEditor) Frame(labels cu.SM, data cu.IM) (title, icon string) {
	return cu.ToString(data["editor_title"], labels["rate_title"]),
		cu.ToString(data["editor_icon"], ct.IconGlobe)
}

func (e *RateEditor) SideBar(labels cu.SM, data cu.IM) (items []ct.SideBarItem) {
	var rate cu.IM = cu.ToIM(data["rate"], cu.IM{"rate_meta": cu.IM{}})
	user := cu.ToIM(data["user"], cu.IM{})

	dirty := cu.ToBoolean(data["dirty"], false)
	readonly := (cu.ToString(user["user_group"], "") == md.UserGroupGuest.String())
	newInput := (cu.ToInteger(rate["id"], 0) == 0)
	updateLabel := labels["editor_save"]
	if newInput {
		updateLabel = labels["editor_create"]
	}
	updateDisabled := func() (disabled bool) {
		return (cu.ToString(rate["currency_code"], "") == "") || readonly
	}

	smState := func() *ct.SideBarStatic {
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
			Label:    labels["rate_new"],
			Icon:     ct.IconUser,
			Disabled: newInput || dirty || readonly,
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
			Href:       st.DocsClientPath + "#rate",
			LinkTarget: "_blank",
		},
	}
}

func (e *RateEditor) View(labels cu.SM, data cu.IM) (views []ct.EditorView) {
	return []ct.EditorView{
		{
			Key:   "rate",
			Label: labels["rate_view"],
			Icon:  ct.IconGlobe,
		},
	}
}

func (e *RateEditor) Row(view string, labels cu.SM, data cu.IM) (rows []ct.Row) {
	var rate md.Rate = md.Rate{}
	ut.ConvertToType(data["rate"], &rate)

	currencies := cu.ToIMA(data["currencies"], []cu.IM{})
	places := cu.ToIMA(data["places"], []cu.IM{})

	rateTypeOpt := func() (opt []ct.SelectOption) {
		opt = []ct.SelectOption{}
		for _, rtype := range md.RateType(0).Keys() {
			opt = append(opt, ct.SelectOption{
				Value: rtype, Text: rtype,
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

	placeOpt := func() (opt []ct.SelectOption) {
		opt = []ct.SelectOption{}
		for _, place := range places {
			opt = append(opt, ct.SelectOption{
				Value: cu.ToString(place["code"], ""), Text: cu.ToString(place["place_name"], ""),
			})
		}
		return opt
	}

	return []ct.Row{
		{Columns: []ct.RowColumn{
			{Label: labels["rate_code"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "code_" + cu.ToString(rate.Id, ""),
				},
				Type: ct.FieldTypeString, Value: cu.IM{
					"name":     "code",
					"value":    rate.Code,
					"disabled": true,
				},
			}},
			{Label: labels["rate_type"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "rate_type_" + cu.ToString(rate.Id, ""),
				},
				Type: ct.FieldTypeSelect, Value: cu.IM{
					"name":    "rate_type",
					"options": rateTypeOpt(),
					"is_null": false,
					"value":   rate.RateType.String(),
				},
			}},
			{
				Label: labels["rate_account"],
				Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "place_code_" + cu.ToString(rate.Id, ""),
					},
					Type: ct.FieldTypeSelect, Value: cu.IM{
						"name":    "place_code",
						"options": placeOpt(),
						"is_null": true,
						"value":   rate.PlaceCode,
					},
				}},
		}, Full: true, BorderBottom: true},
		{Columns: []ct.RowColumn{
			{Label: labels["rate_date"],
				Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "rate_date_" + cu.ToString(rate.Id, ""),
					},
					Type: ct.FieldTypeDate, Value: cu.IM{
						"name":    "rate_date",
						"is_null": false,
						"value":   rate.RateDate,
					},
				}},
			{Label: labels["currency_code"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "currency_code_" + cu.ToString(rate.Id, ""),
				},
				Type: ct.FieldTypeSelect, Value: cu.IM{
					"name":    "currency_code",
					"options": currencyOpt(),
					"is_null": false,
					"value":   rate.CurrencyCode,
				},
			}},
			{
				Label: labels["rate_value"],
				Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "rate_value_" + cu.ToString(rate.Id, ""),
					},
					Type: ct.FieldTypeNumber, Value: cu.IM{
						"name":  "rate_value",
						"value": rate.RateMeta.RateValue,
					},
				}},
		}, Full: true, BorderBottom: true},
		{Columns: []ct.RowColumn{
			{Label: labels["rate_notes"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "notes_" + cu.ToString(rate.Id, ""),
				},
				Type: ct.FieldTypeText, Value: cu.IM{
					"name":  "notes",
					"value": rate.RateMeta.Notes,
					"rows":  4,
				},
			}},
			{
				Label: labels["rate_tags"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "tags_" + cu.ToString(rate.Id, ""),
					},
					Type: ct.FieldTypeList, Value: cu.IM{
						"name":                "tags",
						"rows":                ut.ToTagList(rate.RateMeta.Tags),
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

func (e *RateEditor) Table(view string, labels cu.SM, data cu.IM) []ct.Table {
	return []ct.Table{}
}

func (e *RateEditor) Form(formKey string, labels cu.SM, data cu.IM) (form ct.Form) {
	return ct.Form{}
}
