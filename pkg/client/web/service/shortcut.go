package service

import (
	"encoding/json"
	"net/url"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
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
		"shortcut": cu.IM{
			"data":   cu.IM{},
			"result": "",
		},
		"config_values": cu.IM{},
		"user":          user,
		"editor_icon":   ct.IconShare,
		"editor_title":  client.Msg("office_shortcut_title"),
	}

	var rows []cu.IM = []cu.IM{}
	if rows, err = ds.StoreDataQuery(md.Query{
		Fields: []string{"*"}, From: "config",
		Filters: []md.Filter{
			{Field: "config_type", Comp: "==", Value: md.ConfigTypeShortcut.String()},
		},
	}, false); err != nil {
		return data, err
	}
	data["config_values"] = rows

	return data, err
}

func (s *ShortcutService) Response(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	switch evt.Name {
	case ct.FormEventOK:
		return s.formNext(evt)

	case ct.ClientEventSideMenu:
		return s.sideMenu(evt)

	case ct.FormEventChange:
		return s.formEventChange(evt)

	default:
		return s.editorField(evt)
	}
}

func (s *ShortcutService) sideMenu(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()

	menuMap := map[string]func() (re ct.ResponseEvent, err error){
		"editor_cancel": func() (re ct.ResponseEvent, err error) {
			client.ResetEditor()
			return evt, err
		},
		"shortcut_recall": func() (re ct.ResponseEvent, err error) {
			shortcut := cu.ToIM(stateData["shortcut"], cu.IM{})
			shortcutData := cu.ToIM(shortcut["data"], cu.IM{})
			client.SetForm("shortcut", shortcutData, 0, true)
			return evt, err
		},
		"shortcut_reset": func() (re ct.ResponseEvent, err error) {
			stateData["shortcut"] = cu.IM{
				"data":   cu.IM{},
				"result": "",
			}
			return evt, err
		},
	}

	if fn, ok := menuMap[cu.ToString(evt.Value, "")]; ok {
		return fn()
	}

	return evt, err
}

func (s *ShortcutService) formEventChange(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)

	frmValues := cu.ToIM(evt.Value, cu.IM{})
	frmData := cu.ToIM(frmValues["data"], cu.IM{})

	fieldName := cu.ToString(frmValues["name"], "")
	value := cu.ToString(frmValues["value"], "")
	params := cu.ToIM(frmData["params"], cu.IM{})
	params[fieldName] = value
	frmData["params"] = params
	urlParams := url.Values{}
	for key, pvalue := range params {
		urlParams.Set(key, cu.ToString(pvalue, ""))
	}
	rowData := cu.ToIM(frmData["shortcut"], cu.IM{})
	frmData["url"] = cu.ToString(rowData["address"], "") + "?" + urlParams.Encode()

	client.SetForm("shortcut", frmData, 0, true)

	return evt, err
}

func (s *ShortcutService) formNext(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := s.cls.getDataStore(client.Ticket.Database)

	frmValues := cu.ToIM(evt.Value, cu.IM{})
	frmData := cu.ToIM(frmValues["data"], cu.IM{})

	shortcut := cu.ToIM(frmData["shortcut"], cu.IM{})
	params := cu.ToIM(frmValues["value"], cu.IM{})
	frmData["params"] = params

	stateData["shortcut"] = cu.IM{
		"data":   frmData,
		"result": "",
	}

	url := cu.ToString(frmData["url"], "")
	if url == "" {
		var response any
		if response, err = ds.Function(cu.ToString(shortcut["func_name"], ""), params); err == nil {
			var jsonStr []byte
			if jsonStr, err = s.formatJson(response, "", "  "); err == nil {
				stateData["shortcut"] = cu.MergeIM(cu.ToIM(stateData["shortcut"], cu.IM{}), cu.IM{
					"result": string(jsonStr),
				})
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
		stateData["shortcut"] = cu.MergeIM(cu.ToIM(stateData["shortcut"], cu.IM{}), cu.IM{
			"result": string(result),
		})
	}
	return evt, err
}

func (s *ShortcutService) editorField(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)

	values := cu.ToIM(evt.Value, cu.IM{})
	fieldName := cu.ToString(values["name"], "")
	value := cu.ToIM(values["value"], cu.IM{})
	row := cu.ToIM(value["row"], cu.IM{})
	rowData := cu.ToIM(row["data"], cu.IM{})

	switch fieldName {
	case "shortcut":
		modal := cu.IM{
			"title":    cu.ToString(rowData["description"], ""),
			"icon":     ct.IconShare,
			"next":     "call_shortcut",
			"shortcut": rowData,
			"url":      cu.ToString(rowData["address"], ""),
			"params":   cu.IM{},
		}
		client.SetForm("shortcut", modal, 0, true)
	}
	return evt, err
}
