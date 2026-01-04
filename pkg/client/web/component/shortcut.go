package component

import (
	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
)

type ShortcutEditor struct{}

func (e *ShortcutEditor) Frame(labels cu.SM, data cu.IM) (title, icon string) {
	return cu.ToString(data["editor_title"], labels["shortcut_title"]),
		cu.ToString(data["editor_icon"], ct.IconShare)
}

func (e *ShortcutEditor) SideBar(labels cu.SM, data cu.IM) (items []ct.SideBarItem) {
	shortcut := cu.ToIM(data["shortcut"], cu.IM{})
	shortcutData := cu.ToIM(shortcut["data"], cu.IM{})

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
		&ct.SideBarElement{
			Name:     "shortcut_recall",
			Value:    "shortcut_recall",
			Label:    labels["shortcut_recall"],
			Icon:     ct.IconUndo,
			Disabled: len(shortcutData) == 0,
		},
		&ct.SideBarSeparator{},
		&ct.SideBarElement{
			Name:     "shortcut_reset",
			Value:    "shortcut_reset",
			Label:    labels["shortcut_reset"],
			Icon:     ct.IconListUl,
			Disabled: len(shortcutData) == 0,
		},
	}
}

func (e *ShortcutEditor) View(labels cu.SM, data cu.IM) (views []ct.EditorView) {
	return []ct.EditorView{
		{
			Key:   "office_shortcut",
			Label: labels["office_shortcut_title"],
			Icon:  ct.IconShare,
		},
	}
}

func (e *ShortcutEditor) Row(view string, labels cu.SM, data cu.IM) (rows []ct.Row) {
	shortcut := cu.ToIM(data["shortcut"], cu.IM{})
	result := cu.ToString(shortcut["result"], "")
	if result != "" {
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
