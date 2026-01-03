package component

import (
	"fmt"
	"html/template"
	"slices"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	st "github.com/nervatura/nervatura/v6/pkg/static"
)

type QueueEditor struct{}

func (e *QueueEditor) Frame(labels cu.SM, data cu.IM) (title, icon string) {
	return cu.ToString(data["editor_title"], labels["queue_title"]),
		cu.ToString(data["editor_icon"], ct.IconPrint)
}

func (e *QueueEditor) SideBar(labels cu.SM, data cu.IM) (items []ct.SideBarItem) {
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
			Name:  "queue_search",
			Value: "queue_search",
			Label: labels["queue_search"],
			Icon:  ct.IconSearch,
		},
		&ct.SideBarSeparator{},
		&ct.SideBarElementLink{
			SideBarElement: ct.SideBarElement{
				Name:  "editor_help",
				Value: "editor_help",
				Label: labels["editor_help"],
				Icon:  ct.IconQuestionCircle,
			},
			Href:       st.DocsClientPath + "#queue",
			LinkTarget: "_blank",
		},
	}
}

func (e *QueueEditor) View(labels cu.SM, data cu.IM) (views []ct.EditorView) {
	printQueue := cu.ToIMA(data["print_queue"], []cu.IM{})
	return []ct.EditorView{
		{
			Key:   "office_queue",
			Label: labels["queue_filters"],
			Icon:  ct.IconFilter,
		},
		{
			Key:   "items",
			Label: labels["queue_items"],
			Icon:  ct.IconPrint,
			Badge: cu.ToString(int64(len(printQueue)), "0"),
		},
	}
}

func (e *QueueEditor) Row(view string, labels cu.SM, data cu.IM) (rows []ct.Row) {
	if !slices.Contains([]string{"office_queue"}, view) {
		return []ct.Row{}
	}

	filters := cu.ToIM(data["filters"], cu.IM{})

	reportTypeOpt := func() (opt []ct.SelectOption) {
		opt = []ct.SelectOption{}
		for _, rtype := range []string{"CUSTOMER", "EMPLOYEE", "PRODUCT", "PROJECT", "TOOL", "TRANS"} {
			opt = append(opt, ct.SelectOption{
				Value: rtype, Text: rtype,
			})
		}
		return opt
	}

	return []ct.Row{
		{Columns: []ct.RowColumn{
			{Label: labels["queue_report_type"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "report_type",
				},
				Type: ct.FieldTypeSelect, Value: cu.IM{
					"name":    "report_type",
					"options": reportTypeOpt(),
					"is_null": true,
					"value":   cu.ToString(filters["report_type"], ""),
				},
			}},
			{Label: labels["queue_start_date"],
				Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "start_date",
					},
					Type: ct.FieldTypeDate, Value: cu.IM{
						"name":    "start_date",
						"is_null": true,
						"value":   cu.ToString(filters["start_date"], ""),
					},
				}},
			{Label: labels["queue_end_date"],
				Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "end_date",
					},
					Type: ct.FieldTypeDate, Value: cu.IM{
						"name":    "end_date",
						"is_null": true,
						"value":   cu.ToString(filters["end_date"], ""),
					},
				}},
			{Label: labels["queue_ref_code"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "ref_code",
				},
				Type: ct.FieldTypeString, Value: cu.IM{
					"name":  "ref_code",
					"value": cu.ToString(filters["ref_code"], ""),
				},
			}},
		}, Full: true, BorderBottom: true},
	}
}

func (s *QueueEditor) CustomTemplateCell(sessionID string, label, inline string) func(row cu.IM, col ct.TableColumn, value any, rowIndex int64) template.HTML {
	return func(row cu.IM, col ct.TableColumn, value any, rowIndex int64) template.HTML {
		lnk := ct.Link{
			LinkStyle: ct.LinkStyleDefault,
			Label:     label,
			//Icon:       ct.IconEdit,
			Href:       fmt.Sprintf(st.ClientPath+"/session/export/report/modal/%s?output=pdf&inline=%s&queue=%s", sessionID, inline, cu.ToString(row["code"], "")),
			LinkTarget: "_blank",
		}
		res, _ := lnk.Render()
		return res
	}
}

func (e *QueueEditor) Table(view string, labels cu.SM, data cu.IM) []ct.Table {
	if !slices.Contains([]string{"items"}, view) {
		return []ct.Table{}
	}

	items := cu.ToIMA(data["print_queue"], []cu.IM{})
	sessionID := cu.ToString(data["session_id"], "")
	return []ct.Table{
		{
			Fields: []ct.TableField{
				{Name: "preview", Label: labels["queue_pdf"],
					Column: &ct.TableColumn{Id: "preview",
						Header: labels["queue_pdf"],
						Cell:   e.CustomTemplateCell(sessionID, labels["queue_pdf"], "true")}},
				{Name: "download", Label: labels["queue_download"],
					Column: &ct.TableColumn{Id: "download",
						Header: labels["queue_download"],
						Cell:   e.CustomTemplateCell(sessionID, labels["queue_download"], "false")}},
				{Name: "delete_lbl", Label: labels["queue_delete"], FieldType: ct.TableFieldTypeLink},
				{Name: "ref_type", Label: labels["queue_report_type"]},
				{Name: "ref_code", Label: labels["queue_ref_code"]},
				{Name: "qty", Label: labels["queue_copies"], FieldType: ct.TableFieldTypeNumber},
				{Name: "orientation", Label: labels["queue_orientation"]},
				{Name: "paper_size", Label: labels["queue_size"]},
				{Name: "time_stamp", Label: labels["queue_time_stamp"], FieldType: ct.TableFieldTypeDate},
			},
			Rows:              items,
			Pagination:        ct.PaginationTypeTop,
			PageSize:          5,
			HidePaginatonSize: true,
			TableFilter:       true,
			FilterPlaceholder: labels["placeholder_filter"],
		},
	}
}

func (e *QueueEditor) Form(formKey string, labels cu.SM, data cu.IM) (form ct.Form) {
	return ct.Form{}
}
