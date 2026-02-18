package service

import (
	"fmt"
	"slices"
	"strings"
	"time"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	api "github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

type ToolService struct {
	cls *ClientService
}

func NewToolService(cls *ClientService) *ToolService {
	return &ToolService{
		cls: cls,
	}
}

func (s *ToolService) Data(evt ct.ResponseEvent, params cu.IM) (data cu.IM, err error) {
	client := evt.Trigger.(*ct.Client)
	ds := s.cls.getDataStore(client.Ticket.Database)
	user := cu.ToIM(client.Ticket.User, cu.IM{})

	data = cu.IM{
		"tool":          cu.IM{},
		"config_map":    cu.IM{},
		"config_data":   cu.IM{},
		"config_report": cu.IM{},
		"user":          user,
		"dirty":         false,
		"editor_icon":   ct.IconWrench,
		"editor_title":  "",
	}

	if cu.ToString(params["tool_id"], "") != "" || cu.ToString(params["tool_code"], "") != "" {
		var tools []cu.IM = []cu.IM{}
		if tools, err = ds.StoreDataQuery(md.Query{
			Fields: []string{"*"}, From: "tool",
			Filters: []md.Filter{
				{Field: "deleted", Comp: "==", Value: false},
				{BlockStart: true, Field: "id", Comp: "==", Value: cu.ToInteger(params["tool_id"], 0)},
				{Or: true, BlockEnd: true, Field: "code", Comp: "==", Value: cu.ToString(params["tool_code"], "")},
			},
		}, false); err != nil {
			return data, err
		}
		if len(tools) > 0 {
			data["tool"] = tools[0]
			data["editor_title"] = cu.ToString(tools[0]["code"], "")
		}
	}

	var rows []cu.IM = []cu.IM{}
	if rows, err = ds.StoreDataQuery(md.Query{
		Fields: []string{"*"}, From: "config_map",
	}, false); err != nil {
		return data, err
	}
	data["config_map"] = rows

	if rows, err = ds.StoreDataQuery(md.Query{
		Fields: []string{"*"}, From: "config_data",
	}, false); err != nil {
		return data, err
	}
	data["config_data"] = rows

	if rows, err = ds.StoreDataQuery(md.Query{
		Fields: []string{"id", "report_key", "report_name"}, From: "config_report",
		Filters: []md.Filter{
			{Field: "report_type", Comp: "==", Value: "TOOL"},
		},
	}, false); err != nil {
		return data, err
	}
	data["config_report"] = rows

	return data, err
}

func (s *ToolService) update(ds *api.DataStore, data cu.IM) (editor cu.IM, err error) {
	var tool md.Tool = md.Tool{}
	ds.ConvertData(data["tool"], &tool)
	values := cu.IM{
		"description":  tool.Description,
		"product_code": tool.ProductCode,
	}
	if tool.Code != "" {
		values["code"] = tool.Code
	}

	ut.ConvertByteToIMData(tool.Events, values, "events")
	ut.ConvertByteToIMData(tool.ToolMeta, values, "tool_meta")
	ut.ConvertByteToIMData(tool.ToolMap, values, "tool_map")

	var toolID int64
	newTool := (tool.Id == 0)
	update := md.Update{Values: values, Model: "tool"}
	if !newTool {
		update.IDKey = tool.Id
	}
	if toolID, err = ds.StoreDataUpdate(update); err == nil && newTool {
		var tools []cu.IM = []cu.IM{}
		if tools, err = ds.StoreDataGet(cu.IM{"id": toolID, "model": "tool"}, true); err == nil {
			data["tool"] = tools[0]
			data["editor_title"] = cu.ToString(cu.ToIM(tools[0], cu.IM{})["code"], "")
		}
	}
	return data, err
}

func (s *ToolService) delete(ds *api.DataStore, toolID int64) (err error) {
	return ds.DataDelete("tool", toolID, "")
}

func (s *ToolService) formNext(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := s.cls.getDataStore(client.Ticket.Database)
	tool := cu.ToIM(stateData["tool"], cu.IM{})
	toolMeta := cu.ToIM(tool["tool_meta"], cu.IM{})
	toolMap := cu.ToIM(tool["tool_map"], cu.IM{})

	resultUpdate := func(dirty bool) (re ct.ResponseEvent, err error) {
		tool["tool_meta"] = toolMeta
		tool["tool_map"] = toolMap
		stateData["tool"] = tool
		if dirty {
			stateData["dirty"] = dirty
		}
		client.SetEditor("tool", cu.ToString(stateData["view"], ""), stateData)
		return evt, err
	}

	frmValues := cu.ToIM(evt.Value, cu.IM{})
	frmData := cu.ToIM(frmValues["data"], cu.IM{})
	frmValue := cu.ToIM(frmValues["value"], cu.IM{})

	nextMap := map[string]func() (re ct.ResponseEvent, err error){
		"editor_cancel": func() (re ct.ResponseEvent, err error) {
			client.ResetEditor()
			return evt, err
		},

		"editor_delete": func() (re ct.ResponseEvent, err error) {
			if err = s.delete(ds, cu.ToInteger(tool["id"], 0)); err != nil {
				return evt, err
			}
			client.ResetEditor()
			return evt, err
		},

		"product": func() (re ct.ResponseEvent, err error) {
			params := cu.ToIM(stateData["params"], cu.IM{})
			return s.cls.setEditor(evt, "product", params), nil
		},

		"editor_add_tag": func() (re ct.ResponseEvent, err error) {
			tag := cu.ToString(frmValue["value"], "")
			if tag != "" {
				tags := ut.ToStringArray(toolMeta["tags"])
				if !slices.Contains(tags, tag) {
					tags = append(tags, tag)
					toolMeta["tags"] = tags
					return resultUpdate(true)
				}
			}
			return evt, nil
		},

		"form_add_tag": func() (re ct.ResponseEvent, err error) {
			tag := cu.ToString(frmValue["value"], "")
			frmKey := cu.ToString(frmData["frm_key"], "")
			frmIndex := cu.ToInteger(frmData["frm_index"], 0)
			row := cu.ToIM(frmData["row"], cu.IM{})
			metaName := ut.MetaName(row, "_meta")
			rowField := cu.ToString(frmData["row_field"], "")
			if tag != "" {
				tags := ut.ToStringArray(row[rowField])
				if metaName != "" {
					tags = ut.ToStringArray(cu.ToIM(row[metaName], cu.IM{})[rowField])
				}
				if !slices.Contains(tags, tag) {
					tags = append(tags, tag)
					if metaName != "" {
						cu.ToIM(row[metaName], cu.IM{})[rowField] = tags
					} else {
						row[rowField] = tags
					}
					client.SetForm(frmKey, row, frmIndex, false)
					return evt, nil
				}
			}
			return evt, nil
		},

		"bookmark_add": func() (re ct.ResponseEvent, err error) {
			label := cu.ToString(frmValue["value"], "")
			bookmark := md.Bookmark{
				BookmarkType: md.BookmarkTypeEditor,
				Label:        label,
				Key:          "tool",
				Code:         cu.ToString(tool["code"], ""),
				Filters:      []any{},
				Columns:      map[string]bool{},
				TimeStamp:    time.Now().Format(time.RFC3339),
			}
			return s.cls.addBookmark(evt, bookmark), nil
		},

		"editor_map_value": func() (re ct.ResponseEvent, err error) {
			code := cu.ToString(frmValue["value"], "")
			model := cu.ToString(frmData["model"], "")
			mapField := cu.ToString(frmData["map_field"], "")
			if _, err := ds.GetDataByID(model, 0, code, true); err != nil {
				frmData["label"] = fmt.Sprintf("%s: %s (%s)", client.Msg("invalid_code"), code, model)
				frmData["default_value"] = code
				frmData["invalid"] = true
				client.SetForm("input_string", frmData, 0, true)
				return evt, nil
			}
			toolMap[mapField] = code
			stateData["map_field"] = ""
			return resultUpdate(true)
		},
	}

	if fn, ok := nextMap[cu.ToString(frmData["next"], "")]; ok {
		return fn()
	}
	return evt, err
}

func (s *ToolService) formEvent(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	tool := cu.ToIM(stateData["tool"], cu.IM{})

	frmValues := cu.ToIM(evt.Value, cu.IM{})
	frmValue := cu.ToIM(frmValues["value"], cu.IM{})
	frmData := cu.ToIM(frmValues["data"], cu.IM{})
	form := cu.ToIM(frmData["form"], cu.IM{})
	frmIndex := cu.ToInteger(form["index"], 0)
	frmKey := cu.ToString(form["key"], "")
	frmBaseValues := cu.ToIM(form["data"], cu.IM{})
	frmEvent := cu.ToString(frmValues["event"], "")
	rows := cu.ToIMA(tool[frmKey], []cu.IM{})
	if srows, found := stateData[frmKey]; found && (len(rows) == 0) {
		rows = cu.ToIMA(srows, []cu.IM{})
	}
	delete := (cu.ToString(frmValue["form_delete"], "") != "")

	resultUpdate := func() (re ct.ResponseEvent, err error) {
		if _, found := tool[frmKey]; found {
			tool[frmKey] = rows
		} else {
			stateData[frmKey] = rows
		}
		stateData["dirty"] = true
		return evt, err
	}

	eventMap := map[string]func() (re ct.ResponseEvent, err error){
		ct.FormEventOK: func() (re ct.ResponseEvent, err error) {
			//rowMeta := cu.ToIM(rows[frmIndex][ut.MetaName(rows[frmIndex], "_meta")], cu.IM{})
			//rowMap := cu.ToIM(rows[frmIndex][ut.MetaName(rows[frmIndex], "_map")], cu.IM{})
			customValues := map[string]func(value any){
				"base_tags": func(value any) {
					rows[frmIndex]["tags"] = value
				},
			}
			return s.cls.editorFormOK(evt, rows, customValues)
		},

		ct.FormEventCancel: func() (re ct.ResponseEvent, err error) {
			if delete {
				if _, found := stateData[frmKey]; found {
					deleteRows := cu.ToIMA(stateData[frmKey+"_delete"], []cu.IM{})
					deleteRows = append(deleteRows, rows[frmIndex])
					stateData[frmKey+"_delete"] = deleteRows
				}
				rows = append(rows[:frmIndex], rows[frmIndex+1:]...)
				return resultUpdate()
			}
			return evt, nil
		},

		ct.FormEventChange: func() (re ct.ResponseEvent, err error) {
			fieldName := cu.ToString(frmValues["name"], "")
			switch fieldName {
			case "tags":
				return s.cls.editorFormTags(cu.IM{"row_field": fieldName}, evt)
			default:
				frmBaseValues[fieldName] = frmValues["value"]
				cu.ToSM(evt.Header, cu.SM{})[ct.HeaderReswap] = ct.SwapNone
			}
			return evt, nil
		},
	}

	if len(rows) > 0 && frmIndex < int64(len(rows)) {
		if fn, ok := eventMap[frmEvent]; ok {
			return fn()
		}
	}

	return evt, err
}

func (s *ToolService) sideMenu(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	tool := cu.ToIM(stateData["tool"], cu.IM{})
	ds := s.cls.getDataStore(client.Ticket.Database)

	menuMap := map[string]func() (re ct.ResponseEvent, err error){
		"editor_save": func() (re ct.ResponseEvent, err error) {
			if stateData, err = s.update(ds, stateData); err != nil {
				return evt, err
			}
			stateData["dirty"] = false
			client.SetEditor("tool", cu.ToString(stateData["view"], ""), stateData)
			return evt, err
		},

		"editor_delete": func() (re ct.ResponseEvent, err error) {
			modal := cu.IM{
				"warning_label":   client.Msg("inputbox_delete"),
				"warning_message": "",
				"next":            "editor_delete",
			}
			client.SetForm("warning", modal, 0, true)
			return evt, err
		},

		"editor_cancel": func() (re ct.ResponseEvent, err error) {
			if cu.ToBoolean(stateData["dirty"], false) {
				modal := cu.IM{
					"warning_label":   client.Msg("inputbox_dirty"),
					"warning_message": client.Msg("inputbox_drop"),
					"next":            "editor_cancel",
				}
				client.SetForm("warning", modal, 0, true)
			} else {
				client.ResetEditor()
			}
			return evt, err
		},

		"editor_new": func() (re ct.ResponseEvent, err error) {
			return s.cls.setEditor(evt, "tool",
				cu.IM{
					"session_id": client.Ticket.SessionID,
				}), nil
		},

		"editor_report": func() (re ct.ResponseEvent, err error) {
			return s.cls.showReportSelector(evt, "TOOL", cu.ToString(tool["code"], ""))
		},

		"editor_bookmark": func() (re ct.ResponseEvent, err error) {
			modal := cu.IM{
				"title":         client.Msg("bookmark_new"),
				"icon":          ct.IconStar,
				"label":         client.Msg("bookmark_enter"),
				"placeholder":   "",
				"field_name":    "value",
				"default_value": "",
				"required":      false,
				"next":          "bookmark_add",
			}
			client.SetForm("input_string", modal, 0, true)
			return evt, nil
		},
	}

	if fn, ok := menuMap[cu.ToString(evt.Value, "")]; ok {
		return fn()
	}

	return evt, err
}

func (s *ToolService) editorField(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := s.cls.getDataStore(client.Ticket.Database)
	tool := cu.ToIM(stateData["tool"], cu.IM{})
	toolMeta := cu.ToIM(tool["tool_meta"], cu.IM{})
	toolMap := cu.ToIM(tool["tool_map"], cu.IM{})

	resultUpdate := func(params cu.IM) (re ct.ResponseEvent, err error) {
		tool["tool_meta"] = toolMeta
		tool["tool_map"] = toolMap
		stateData["tool"] = tool
		if cu.ToBoolean(params["dirty"], false) {
			stateData["dirty"] = true
		}
		client.SetEditor("tool", cu.ToString(stateData["view"], ""), stateData)
		return evt, err
	}

	values := cu.ToIM(evt.Value, cu.IM{})
	fieldName := cu.ToString(values["name"], "")

	value := cu.ToString(values["value"], "")

	fieldMap := map[string]func() (re ct.ResponseEvent, err error){
		ct.TableEventRowSelected: func() (re ct.ResponseEvent, err error) {
			valueData := cu.ToIM(values["value"], cu.IM{})
			client.SetForm(cu.ToString(stateData["view"], ""), cu.ToIM(valueData["row"], cu.IM{}), cu.ToInteger(valueData["index"], 0), false)
			return evt, nil
		},

		ct.TableEventAddItem: func() (re ct.ResponseEvent, err error) {
			view := cu.ToString(stateData["view"], "")
			typeMap := map[string]func() cu.IM{
				"events": func() cu.IM {
					var event cu.IM
					ds.ConvertData(md.Event{
						Tags:     []string{},
						EventMap: cu.IM{},
					}, &event)
					return event
				},
			}
			if slices.Contains([]string{"events"}, view) {
				getBase := func() (base cu.IM) {
					if _, found := tool[view]; found {
						return tool
					}
					return stateData
				}
				base := getBase()
				rows := cu.ToIMA(base[view], []cu.IM{})
				rows = append(rows, typeMap[view]())
				base[view] = rows
				client.SetForm(view, cu.MergeIM(typeMap[view](), cu.IM{}), cu.ToInteger(len(rows)-1, 0), false)
				return evt, nil
			}
			return s.cls.addMapField(evt, toolMap, resultUpdate)
		},

		ct.TableEventFormDelete: func() (re ct.ResponseEvent, err error) {
			valueData := cu.ToIM(values["value"], cu.IM{})
			row := cu.ToIM(valueData["row"], cu.IM{})
			fieldName := cu.ToString(row["field_name"], "")
			delete(toolMap, fieldName)
			return resultUpdate(cu.IM{"dirty": true})
		},

		ct.TableEventFormUpdate: func() (re ct.ResponseEvent, err error) {
			return s.cls.updateMapField(evt, toolMap, resultUpdate)
		},

		ct.TableEventFormChange: func() (re ct.ResponseEvent, err error) {
			return evt, nil
		},

		ct.TableEventFormCancel: func() (re ct.ResponseEvent, err error) {
			return evt, nil
		},

		"map_field": func() (re ct.ResponseEvent, err error) {
			stateData["map_field"] = value
			return resultUpdate(cu.IM{"dirty": false})
		},

		"queue": func() (re ct.ResponseEvent, err error) {
			modal := cu.ToIM(client.Data["modal"], cu.IM{})
			modalData := cu.ToIM(modal["data"], cu.IM{})
			if err = s.cls.insertPrintQueue(ds, modalData); err == nil {
				return s.cls.evtMsg(evt.Name, evt.TriggerName, client.Msg("report_add_queue"), ct.ToastTypeSuccess, 5), nil
			}
			return evt, err
		},

		"tags": func() (re ct.ResponseEvent, err error) {
			return s.cls.editorTags(evt, toolMeta, resultUpdate)
		},

		"description": func() (re ct.ResponseEvent, err error) {
			tool[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"product_code": func() (re ct.ResponseEvent, err error) {
			return s.cls.editorCodeSelector(evt, "tool", strings.Split(fieldName, "_")[0], tool, resultUpdate)
		},

		"notes": func() (re ct.ResponseEvent, err error) {
			toolMeta[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"serial_number": func() (re ct.ResponseEvent, err error) {
			toolMeta[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"inactive": func() (re ct.ResponseEvent, err error) {
			toolMeta[fieldName] = cu.ToBoolean(value, false)
			return resultUpdate(cu.IM{"dirty": true})
		},
	}

	if fn, ok := fieldMap[fieldName]; ok {
		return fn()
	}
	if slices.Contains([]string{"orientation", "template", "paper_size", "copy"}, fieldName) {
		modal := cu.ToIM(client.Data["modal"], cu.IM{})
		modalData := cu.ToIM(modal["data"], cu.IM{})
		modalData[fieldName] = value
		client.SetForm("report", modalData, 0, true)
		return evt, nil
	}
	return evt, nil
}

func (s *ToolService) Response(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	switch evt.Name {
	case ct.FormEventOK:
		return s.formNext(evt)

	case ct.ClientEventForm:
		return s.formEvent(evt)

	case ct.ClientEventSideMenu:
		return s.sideMenu(evt)

	default:
		return s.editorField(evt)
	}
}
