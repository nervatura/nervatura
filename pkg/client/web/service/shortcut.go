package service

import (
	"encoding/json"
	"fmt"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	st "github.com/nervatura/nervatura/v6/pkg/static"
)

type ShortcutService struct {
	cls        *ClientService
	formatJson func(v any, prefix string, indent string) ([]byte, error)
}

func NewShortcutService(cls *ClientService) *ShortcutService {
	return &ShortcutService{
		cls:        cls,
		formatJson: json.MarshalIndent,
	}
}

func (s *ShortcutService) Data(evt ct.ResponseEvent, params cu.IM) (data cu.IM, err error) {
	client := evt.Trigger.(*ct.Client)
	ds := s.cls.getDataStore(client.Ticket.Database)
	user := cu.ToIM(client.Ticket.User, cu.IM{})

	data = cu.IM{
		"shortcut":     cu.IM{},
		"params":       cu.IM{},
		"result":       "",
		"items":        []cu.IM{},
		"user":         user,
		"editor_icon":  ct.IconShare,
		"editor_title": client.Msg("office_shortcut_title"),
	}

	var items []cu.IM = []cu.IM{}
	var rows []cu.IM = []cu.IM{}

	if rows, err = ds.StoreDataQuery(md.Query{
		Fields: []string{"*"}, From: "config_shortcut",
	}, false); err != nil {
		return data, err
	}
	for _, row := range rows {
		items = append(items, cu.MergeIM(row,
			cu.IM{
				"lstype":  "shortcut",
				"lslabel": cu.ToString(row["shortcut_key"], ""),
				"lsvalue": cu.ToString(row["description"], ""),
			}))
	}

	if rows, err = ds.StoreDataQuery(md.Query{
		Fields: []string{"*"}, From: "config_report",
		Filters: []md.Filter{
			{Field: "report_type", Comp: "==", Value: "REPORT"},
			{Field: "file_type", Comp: "==", Value: md.FileTypeCSV.String()},
		},
	}, false); err != nil {
		return data, err
	}
	for _, row := range rows {
		items = append(items, cu.MergeIM(row,
			cu.IM{
				"lstype":  "report",
				"lslabel": cu.ToString(row["report_name"], ""),
				"lsvalue": cu.ToString(row["description"], ""),
			}))
	}
	data["items"] = items

	return data, err
}

func (s *ShortcutService) Response(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	switch evt.Name {

	case ct.ClientEventSideMenu:
		return s.sideMenu(evt)

	default:
		return s.editorField(evt)
	}
}

func (s *ShortcutService) callData(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := s.cls.getDataStore(client.Ticket.Database)

	shortcut := cu.ToIM(stateData["shortcut"], cu.IM{})
	params := cu.ToIM(stateData["params"], cu.IM{})

	url := cu.ToString(shortcut["address"], "")
	if url == "" {
		var response any
		if response, err = ds.Function(cu.ToString(shortcut["func_name"], ""), params); err == nil {
			var jsonStr []byte
			if jsonStr, err = s.formatJson(response, "", "  "); err == nil {
				stateData["result"] = string(jsonStr)
			}
		}
		return evt, err
	}

	var body []byte
	if len(params) > 0 {
		body, _ = cu.ConvertToByte(params)
	}
	var result []byte
	if result, err = ds.MakeRequest("POST", url, body, cu.SM{}); err == nil {
		stateData["result"] = string(result)
	}
	return evt, err
}

func (s *ShortcutService) sideMenu(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()

	menuMap := map[string]func() (re ct.ResponseEvent, err error){
		"editor_cancel": func() (re ct.ResponseEvent, err error) {
			client.ResetEditor()
			return evt, err
		},
		"shortcut_call": func() (re ct.ResponseEvent, err error) {
			return s.callData(evt)
		},
		"shortcut_list": func() (re ct.ResponseEvent, err error) {
			stateData["shortcut"] = cu.IM{}
			stateData["result"] = ""
			stateData["params"] = cu.IM{}
			return evt, err
		},
		"shortcut_reset": func() (re ct.ResponseEvent, err error) {
			stateData["result"] = ""
			return evt, err
		},
	}

	if fn, ok := menuMap[cu.ToString(evt.Value, "")]; ok {
		return fn()
	}

	return evt, err
}

func (s *ShortcutService) shortcutItem(row cu.IM, sessionID string) (item cu.IM) {
	if lstype := cu.ToString(row["lstype"], ""); lstype == "shortcut" {
		return row
	}
	template := cu.ToString(row["template"], "")
	var report cu.IM = cu.IM{}
	var fields []cu.IM = []cu.IM{}
	var err error
	if err = cu.ConvertFromByte([]byte(template), &report); err == nil {
		reportFields := cu.ToIM(report["fields"], cu.IM{})
		for key, values := range reportFields {
			values := cu.ToIM(values, cu.IM{})
			fields = append(fields, cu.IM{
				"field_name":  key,
				"description": cu.ToString(values["description"], ""),
				"field_type":  cu.ToString(values["fieldtype"], ""),
				"required":    (cu.ToString(values["wheretype"], "") == "in"),
				"order":       cu.ToInteger(values["orderby"], 0),
			})
		}
	}
	item = cu.MergeIM(row, cu.IM{
		"url":    fmt.Sprintf(st.ClientPath+"/session/export/report/%s?output=export", sessionID),
		"fields": fields,
	})
	return item
}

func (s *ShortcutService) editorField(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()

	values := cu.ToIM(evt.Value, cu.IM{})
	fieldName := cu.ToString(values["name"], "")

	switch fieldName {
	case "shortcut":
		value := cu.ToIM(values["value"], cu.IM{})
		row := cu.ToIM(value["row"], cu.IM{})
		stateData["shortcut"] = s.shortcutItem(row, client.Ticket.SessionID)
		stateData["result"] = ""
		stateData["params"] = cu.IM{}
	default:
		params := cu.ToIM(stateData["params"], cu.IM{})
		params[fieldName] = cu.ToString(values["value"], "")
		stateData["params"] = params
	}
	return evt, err
}
