package component

import (
	"net/url"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

type ShortcutEditor struct{}

func (e *ShortcutEditor) Frame(labels cu.SM, data cu.IM) (title, icon string) {
	return cu.ToString(data["editor_title"], labels["shortcut_title"]),
		cu.ToString(data["editor_icon"], ct.IconShare)
}

func (e *ShortcutEditor) SideBar(labels cu.SM, data cu.IM) (items []ct.SideBarItem) {
	shortcut := cu.ToIM(data["shortcut"], cu.IM{})
	result := cu.ToString(data["result"], "")
	params := cu.ToIM(data["params"], cu.IM{})
	urlLink := (cu.ToString(shortcut["method"], "") == md.ShortcutMethodGET.String()) && (cu.ToString(shortcut["address"], "") != "")

	items = []ct.SideBarItem{
		&ct.SideBarSeparator{},
		&ct.SideBarElement{
			Name:    "editor_cancel",
			Value:   "editor_cancel",
			Label:   labels["browser_title"],
			Icon:    ct.IconReply,
			NotFull: true,
		},
		&ct.SideBarSeparator{},
		&ct.SideBarElement{
			Name:     "shortcut_list",
			Value:    "shortcut_list",
			Label:    labels["shortcut_title"],
			Icon:     ct.IconListUl,
			Disabled: len(shortcut) == 0,
		},
		&ct.SideBarSeparator{},
		&ct.SideBarElement{
			Name:     "shortcut_reset",
			Value:    "shortcut_reset",
			Label:    labels["shortcut_reset"],
			Icon:     ct.IconUndo,
			Disabled: result == "",
		},
		&ct.SideBarSeparator{},
	}

	if urlLink {
		urlParams := url.Values{}
		for key, pvalue := range params {
			urlParams.Set(key, cu.ToString(pvalue, ""))
		}
		items = append(items, &ct.SideBarElementLink{
			SideBarElement: ct.SideBarElement{
				Name:     "shortcut_call",
				Value:    "shortcut_call",
				Label:    labels["shortcut_call"],
				Icon:     ct.IconShare,
				Disabled: len(shortcut) == 0,
			},
			Href: cu.ToString(shortcut["address"], "") + "?" + urlParams.Encode(),
			//Download:   fmt.Sprintf("declaration_%s.pdf", cu.ToString(submission["id"], "")),
			LinkTarget: "_blank",
		})
	} else {
		items = append(items, &ct.SideBarElement{
			Name:     "shortcut_call",
			Value:    "shortcut_call",
			Label:    labels["shortcut_call"],
			Icon:     ct.IconShare,
			Disabled: len(shortcut) == 0,
		})
	}
	return items
}

func (e *ShortcutEditor) View(labels cu.SM, data cu.IM) (views []ct.EditorView) {
	shortcut := cu.ToIM(data["shortcut"], cu.IM{})
	return []ct.EditorView{
		{
			Key:   "office_shortcut",
			Label: cu.ToString(shortcut["description"], labels["shortcut_title"]),
			Icon:  ct.IconShare,
		},
	}
}

func (e *ShortcutEditor) fieldTypeMap(field, params cu.IM) ct.Field {
	fieldName := cu.ToString(field["field_name"], "")
	switch cu.ToString(field["field_type"], "") {
	case md.ShortcutFieldBool.String():
		return ct.Field{
			Type: ct.FieldTypeBool,
			Value: cu.IM{
				"name":  fieldName,
				"value": cu.ToBoolean(params[fieldName], false),
			},
		}
	case md.ShortcutFieldInteger.String():
		return ct.Field{
			Type: ct.FieldTypeInteger,
			Value: cu.IM{
				"name":  fieldName,
				"value": cu.ToInteger(params[fieldName], 0),
			},
		}
	case md.ShortcutFieldNumber.String():
		return ct.Field{
			Type: ct.FieldTypeNumber,
			Value: cu.IM{
				"name":  fieldName,
				"value": cu.ToFloat(params[fieldName], 0),
			},
		}
	case md.ShortcutFieldDate.String():
		return ct.Field{
			Type: ct.FieldTypeDate,
			Value: cu.IM{
				"name":  fieldName,
				"value": cu.ToString(params[fieldName], ""),
			},
		}
	default:
		return ct.Field{
			Type: ct.FieldTypeString,
			Value: cu.IM{
				"name":  fieldName,
				"value": cu.ToString(params[fieldName], ""),
			},
		}
	}
}

func (e *ShortcutEditor) Row(view string, labels cu.SM, data cu.IM) (rows []ct.Row) {
	shortcut := cu.ToIM(data["shortcut"], cu.IM{})
	result := cu.ToString(data["result"], "")
	params := cu.ToIM(data["params"], cu.IM{})
	if len(shortcut) > 0 && result != "" {
		return []ct.Row{
			{Columns: []ct.RowColumn{
				{Label: labels["shortcut_result"],
					Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Id: "shortcut_result",
						},
						Type: ct.FieldTypeText,
						Value: cu.IM{
							"name":     "shortcut_result",
							"value":    result,
							"readonly": true,
							"rows":     10,
						},
					}},
			}, Full: true, BorderBottom: true},
		}
	}

	if len(shortcut) > 0 {
		rows = []ct.Row{}
		fields := cu.ToIMA(shortcut["fields"], []cu.IM{})
		ut.SortIMData(fields, "order")
		row := ct.Row{Columns: []ct.RowColumn{}, Full: true, BorderBottom: true}
		for _, field := range fields {
			col := ct.RowColumn{Label: cu.ToString(field["description"], ""), Value: e.fieldTypeMap(field, params)}
			if len(row.Columns) > 1 {
				rows = append(rows, row)
				row = ct.Row{Columns: []ct.RowColumn{}, Full: true, BorderBottom: true}
			}
			row.Columns = append(row.Columns, col)
		}
		if len(row.Columns) > 0 {
			rows = append(rows, row)
		}
		return rows
	}

	configValues := cu.ToIMA(data["config_values"], []cu.IM{})
	mapRows := []cu.IM{}
	for _, row := range configValues {
		configMap := cu.ToIM(row["data"], cu.IM{})
		mapRows = append(mapRows, cu.MergeIM(row,
			cu.IM{"lslabel": cu.ToString(configMap["shortcut_key"], ""), "lsvalue": cu.ToString(configMap["description"], "")}))
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
					"add_item":            false,
					"edit_item":           true,
					"edit_icon":           ct.IconShare,
					"delete_item":         false,
				},
			}},
		}, Full: true, BorderBottom: true},
	}
}

func (e *ShortcutEditor) Table(view string, labels cu.SM, data cu.IM) []ct.Table {
	return []ct.Table{}
}

func (e *ShortcutEditor) Form(formKey string, labels cu.SM, data cu.IM) (form ct.Form) {
	return ct.Form{}
}
