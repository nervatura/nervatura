package service

import (
	"slices"
	"time"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	api "github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

type RateService struct {
	cls *ClientService
}

func NewRateService(cls *ClientService) *RateService {
	return &RateService{
		cls: cls,
	}
}

func (s *RateService) Data(evt ct.ResponseEvent, params cu.IM) (data cu.IM, err error) {
	client := evt.Trigger.(*ct.Client)
	ds := s.cls.getDataStore(client.Ticket.Database)
	user := cu.ToIM(client.Ticket.User, cu.IM{})

	data = cu.IM{
		"rate": cu.IM{
			"rate_type": md.RateType(0),
			"rate_date": time.Now().Format(time.RFC3339),
			"rate_meta": cu.IM{
				"tags": []string{},
			},
			"rate_map": cu.IM{},
		},
		"currencies":   cu.IM{},
		"places":       cu.IM{},
		"user":         user,
		"dirty":        false,
		"editor_icon":  ct.IconGlobe,
		"editor_title": "",
	}

	var rows []cu.IM = []cu.IM{}
	if cu.ToString(params["rate_id"], "") != "" || cu.ToString(params["rate_code"], "") != "" {
		if rows, err = ds.StoreDataQuery(md.Query{
			Fields: []string{"*"}, From: "rate",
			Filters: []md.Filter{
				{Field: "deleted", Comp: "==", Value: false},
				{Field: "id", Comp: "==", Value: cu.ToInteger(params["rate_id"], 0)},
				{Or: true, Field: "code", Comp: "==", Value: cu.ToString(params["rate_code"], "")},
			},
		}, false); err != nil {
			return data, err
		}
		if len(rows) > 0 {
			data["rate"] = rows[0]
			data["editor_title"] = cu.ToString(rows[0]["code"], "")
		}
	}
	var rate cu.IM = cu.ToIM(data["rate"], cu.IM{})

	if rows, err = ds.StoreDataQuery(md.Query{
		Fields: []string{"*"},
		From:   "currency",
		Filters: []md.Filter{
			{Field: "deleted", Comp: "==", Value: false},
		},
	}, false); err != nil {
		return data, err
	}
	data["currencies"] = rows
	if len(rows) > 0 {
		rate["currency_code"] = cu.ToString(rows[0]["code"], "")
	}

	if rows, err = ds.StoreDataQuery(md.Query{
		Fields: []string{"*"},
		From:   "place",
		Filters: []md.Filter{
			{Field: "deleted", Comp: "==", Value: false},
			{Field: "place_type", Comp: "==", Value: md.PlaceTypeBank.String()},
		},
		OrderBy: []string{"code"},
	}, false); err != nil {
		return data, err
	}
	data["places"] = rows

	return data, err
}

func (s *RateService) Response(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	switch evt.Name {
	case ct.FormEventOK:
		return s.formNext(evt)

	case ct.ClientEventSideMenu:
		return s.sideMenu(evt)

	default:
		return s.editorField(evt)
	}
}

func (s *RateService) update(ds *api.DataStore, data cu.IM) (editor cu.IM, err error) {
	var rate md.Rate = md.Rate{}
	ut.ConvertToType(data["rate"], &rate)
	values := cu.IM{
		"rate_type":     rate.RateType.String(),
		"rate_date":     rate.RateDate,
		"currency_code": rate.CurrencyCode,
		"place_code":    nil,
	}
	if rate.Code != "" {
		values["code"] = rate.Code
	}
	if rate.PlaceCode != "" {
		values["place_code"] = rate.PlaceCode
	}

	ut.ConvertByteToIMData(rate.RateMeta, values, "rate_meta")
	ut.ConvertByteToIMData(rate.RateMap, values, "rate_map")

	var rateID int64
	newRate := (rate.Id == 0)
	update := md.Update{Values: values, Model: "rate"}
	if !newRate {
		update.IDKey = rate.Id
	}
	if rateID, err = ds.StoreDataUpdate(update); err == nil && newRate {
		var rates []cu.IM = []cu.IM{}
		if rates, err = ds.StoreDataGet(cu.IM{"id": rateID, "model": "rate"}, true); err == nil {
			data["rate"] = rates[0]
			data["editor_title"] = cu.ToString(cu.ToIM(rates[0], cu.IM{})["code"], "")
		}
	}
	return data, err
}

func (s *RateService) delete(ds *api.DataStore, rateID int64) (err error) {
	return ds.DataDelete("rate", rateID, "")
}

func (s *RateService) formNext(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := s.cls.getDataStore(client.Ticket.Database)
	rate := cu.ToIM(stateData["rate"], cu.IM{})
	rateMeta := cu.ToIM(rate["rate_meta"], cu.IM{})

	resultUpdate := func(dirty bool) (re ct.ResponseEvent, err error) {
		rate["rate_meta"] = rateMeta
		stateData["rate"] = rate
		if dirty {
			stateData["dirty"] = dirty
		}
		client.SetEditor("rate", cu.ToString(stateData["view"], ""), stateData)
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
			if err = s.delete(ds, cu.ToInteger(rate["id"], 0)); err != nil {
				return evt, err
			}
			client.ResetEditor()
			return evt, err
		},

		"editor_add_tag": func() (re ct.ResponseEvent, err error) {
			tag := cu.ToString(frmValue["value"], "")
			if tag != "" {
				tags := ut.ToStringArray(rateMeta["tags"])
				if !slices.Contains(tags, tag) {
					tags = append(tags, tag)
					rateMeta["tags"] = tags
					return resultUpdate(true)
				}
			}
			return evt, nil
		},

		"bookmark_add": func() (re ct.ResponseEvent, err error) {
			label := cu.ToString(frmValue["value"], "")
			bookmark := md.Bookmark{
				BookmarkType: md.BookmarkTypeEditor,
				Label:        label,
				Key:          "rate",
				Code:         cu.ToString(rate["code"], ""),
				Filters:      []any{},
				Columns:      map[string]bool{},
				TimeStamp:    time.Now().Format(time.RFC3339),
			}
			return s.cls.addBookmark(evt, bookmark), nil
		},
	}

	if fn, ok := nextMap[cu.ToString(frmData["next"], "")]; ok {
		return fn()
	}
	return evt, err
}

func (s *RateService) sideMenu(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := s.cls.getDataStore(client.Ticket.Database)

	menuMap := map[string]func() (re ct.ResponseEvent, err error){
		"editor_save": func() (re ct.ResponseEvent, err error) {
			if stateData, err = s.update(ds, stateData); err != nil {
				return evt, err
			}
			stateData["dirty"] = false
			client.SetEditor("rate", cu.ToString(stateData["view"], ""), stateData)
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
			return s.cls.setEditor(evt, "rate",
				cu.IM{
					"session_id": client.Ticket.SessionID,
				}), nil
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

func (s *RateService) editorField(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	rate := cu.ToIM(stateData["rate"], cu.IM{})
	rateMeta := cu.ToIM(rate["rate_meta"], cu.IM{})

	resultUpdate := func(params cu.IM) (re ct.ResponseEvent, err error) {
		rate["rate_meta"] = rateMeta
		stateData["rate"] = rate
		if cu.ToBoolean(params["dirty"], false) {
			stateData["dirty"] = true
		}
		client.SetEditor("rate", cu.ToString(stateData["view"], ""), stateData)
		return evt, err
	}

	values := cu.ToIM(evt.Value, cu.IM{})
	fieldName := cu.ToString(values["name"], "")
	value := cu.ToString(values["value"], "")

	fieldMap := map[string]func() (re ct.ResponseEvent, err error){
		"tags": func() (re ct.ResponseEvent, err error) {
			return s.cls.editorTags(evt, rateMeta, resultUpdate)
		},

		"rate_date": func() (re ct.ResponseEvent, err error) {
			rate[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"rate_type": func() (re ct.ResponseEvent, err error) {
			rate[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"notes": func() (re ct.ResponseEvent, err error) {
			rateMeta[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"rate_value": func() (re ct.ResponseEvent, err error) {
			rateMeta[fieldName] = cu.ToFloat(value, 0)
			return resultUpdate(cu.IM{"dirty": true})
		},

		"place_code": func() (re ct.ResponseEvent, err error) {
			rate[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"currency_code": func() (re ct.ResponseEvent, err error) {
			rate[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},
	}

	if fn, ok := fieldMap[fieldName]; ok {
		return fn()
	}
	return evt, nil
}
