package service

import (
	"fmt"
	"slices"
	"time"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	api "github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

type PlaceService struct {
	cls *ClientService
}

func NewPlaceService(cls *ClientService) *PlaceService {
	return &PlaceService{
		cls: cls,
	}
}

func (s *PlaceService) Data(evt ct.ResponseEvent, params cu.IM) (data cu.IM, err error) {
	client := evt.Trigger.(*ct.Client)
	ds := s.cls.getDataStore(client.Ticket.Database)
	user := cu.ToIM(client.Ticket.User, cu.IM{})

	data = cu.IM{
		"place": cu.IM{
			"place_type": md.PlaceType(0),
			"place_meta": cu.IM{
				"tags": []string{},
			},
			"place_map": cu.IM{},
			"address":   cu.IM{},
			"contacts":  []cu.IM{},
			"events":    []cu.IM{},
		},
		"config_map":   cu.IM{},
		"config_data":  cu.IM{},
		"currencies":   cu.IM{},
		"user":         user,
		"dirty":        false,
		"editor_icon":  ct.IconUser,
		"editor_title": "",
	}

	if cu.ToString(params["place_id"], "") != "" || cu.ToString(params["place_code"], "") != "" {
		var places []cu.IM = []cu.IM{}
		if places, err = ds.StoreDataQuery(md.Query{
			Fields: []string{"*"}, From: "place",
			Filters: []md.Filter{
				{Field: "deleted", Comp: "==", Value: false},
				{BlockStart: true, Field: "id", Comp: "==", Value: cu.ToInteger(params["place_id"], 0)},
				{Or: true, BlockEnd: true, Field: "code", Comp: "==", Value: cu.ToString(params["place_code"], "")},
			},
		}, false); err != nil {
			return data, err
		}
		if len(places) > 0 {
			data["place"] = places[0]
			data["editor_title"] = cu.ToString(places[0]["code"], "")
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
		Fields: []string{"code", "description", "digit"}, From: "currency_view",
		OrderBy: []string{"code"},
	}, false); err != nil {
		return data, err
	}
	data["currencies"] = rows

	return data, err
}

func (s *PlaceService) update(ds *api.DataStore, data cu.IM) (editor cu.IM, err error) {
	var place md.Place = md.Place{}
	ds.ConvertData(data["place"], &place)
	values := cu.IM{
		"place_type": place.PlaceType.String(),
		"place_name": place.PlaceName,
	}
	if place.Code != "" {
		values["code"] = place.Code
	}

	ut.ConvertByteToIMData(place.Contacts, values, "contacts")
	ut.ConvertByteToIMData(place.Address, values, "address")
	ut.ConvertByteToIMData(place.Events, values, "events")
	ut.ConvertByteToIMData(place.PlaceMeta, values, "place_meta")
	ut.ConvertByteToIMData(place.PlaceMap, values, "place_map")

	var placeID int64
	newPlace := (place.Id == 0)
	update := md.Update{Values: values, Model: "place"}
	if !newPlace {
		update.IDKey = place.Id
	}
	if placeID, err = ds.StoreDataUpdate(update); err == nil && newPlace {
		var places []cu.IM = []cu.IM{}
		if places, err = ds.StoreDataGet(cu.IM{"id": placeID, "model": "place"}, true); err == nil {
			data["place"] = places[0]
			data["editor_title"] = cu.ToString(cu.ToIM(places[0], cu.IM{})["code"], "")
		}
	}
	return data, err
}

func (s *PlaceService) delete(ds *api.DataStore, placeID int64) (err error) {
	return ds.DataDelete("place", placeID, "")
}

func (s *PlaceService) formNext(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := s.cls.getDataStore(client.Ticket.Database)
	place := cu.ToIM(stateData["place"], cu.IM{})
	placeMeta := cu.ToIM(place["place_meta"], cu.IM{})
	placeMap := cu.ToIM(place["place_map"], cu.IM{})

	resultUpdate := func(dirty bool) (re ct.ResponseEvent, err error) {
		place["place_meta"] = placeMeta
		place["place_map"] = placeMap
		stateData["place"] = place
		if dirty {
			stateData["dirty"] = dirty
		}
		client.SetEditor("place", cu.ToString(stateData["view"], ""), stateData)
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
			if err = s.delete(ds, cu.ToInteger(place["id"], 0)); err != nil {
				return evt, err
			}
			client.ResetEditor()
			return evt, err
		},

		"editor_add_tag": func() (re ct.ResponseEvent, err error) {
			tag := cu.ToString(frmValue["value"], "")
			if tag != "" {
				tags := ut.ToStringArray(placeMeta["tags"])
				if !slices.Contains(tags, tag) {
					tags = append(tags, tag)
					placeMeta["tags"] = tags
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
				Key:          "place",
				Code:         cu.ToString(place["code"], ""),
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
			placeMap[mapField] = code
			stateData["map_field"] = ""
			return resultUpdate(true)
		},
	}

	if fn, ok := nextMap[cu.ToString(frmData["next"], "")]; ok {
		return fn()
	}
	return evt, err
}

func (s *PlaceService) formEvent(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	place := cu.ToIM(stateData["place"], cu.IM{})

	frmValues := cu.ToIM(evt.Value, cu.IM{})
	frmValue := cu.ToIM(frmValues["value"], cu.IM{})
	frmData := cu.ToIM(frmValues["data"], cu.IM{})
	form := cu.ToIM(frmData["form"], cu.IM{})
	frmIndex := cu.ToInteger(form["index"], 0)
	frmKey := cu.ToString(form["key"], "")
	frmBaseValues := cu.ToIM(form["data"], cu.IM{})
	frmEvent := cu.ToString(frmValues["event"], "")
	rows := cu.ToIMA(place[frmKey], []cu.IM{})
	if srows, found := stateData[frmKey]; found && (len(rows) == 0) {
		rows = cu.ToIMA(srows, []cu.IM{})
	}
	delete := (cu.ToString(frmValue["form_delete"], "") != "")

	resultUpdate := func() (re ct.ResponseEvent, err error) {
		if _, found := place[frmKey]; found {
			place[frmKey] = rows
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

func (s *PlaceService) sideMenu(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := s.cls.getDataStore(client.Ticket.Database)

	menuMap := map[string]func() (re ct.ResponseEvent, err error){
		"editor_save": func() (re ct.ResponseEvent, err error) {
			if stateData, err = s.update(ds, stateData); err != nil {
				return evt, err
			}
			stateData["dirty"] = false
			client.SetEditor("place", cu.ToString(stateData["view"], ""), stateData)
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
			return s.cls.setEditor(evt, "place",
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

func (s *PlaceService) editorField(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := s.cls.getDataStore(client.Ticket.Database)

	place := cu.ToIM(stateData["place"], cu.IM{})
	placeMeta := cu.ToIM(place["place_meta"], cu.IM{})
	placeMap := cu.ToIM(place["place_map"], cu.IM{})
	address := cu.ToIM(place["address"], cu.IM{})

	resultUpdate := func(params cu.IM) (re ct.ResponseEvent, err error) {
		place["place_meta"] = placeMeta
		place["place_map"] = placeMap
		place["address"] = address
		stateData["place"] = place
		if cu.ToBoolean(params["dirty"], false) {
			stateData["dirty"] = true
		}
		client.SetEditor("place", cu.ToString(stateData["view"], ""), stateData)
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
				"contacts": func() cu.IM {
					var contact cu.IM
					ds.ConvertData(md.Contact{
						Tags:       []string{},
						ContactMap: cu.IM{},
					}, &contact)
					return contact
				},
				"events": func() cu.IM {
					var event cu.IM
					ds.ConvertData(md.Event{
						Tags:     []string{},
						EventMap: cu.IM{},
					}, &event)
					return event
				},
			}
			if slices.Contains([]string{"contacts", "events"}, view) {
				getBase := func() (base cu.IM) {
					if _, found := place[view]; found {
						return place
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
			return s.cls.addMapField(evt, placeMap, resultUpdate)
		},

		ct.TableEventFormDelete: func() (re ct.ResponseEvent, err error) {
			valueData := cu.ToIM(values["value"], cu.IM{})
			row := cu.ToIM(valueData["row"], cu.IM{})
			fieldName := cu.ToString(row["field_name"], "")
			delete(placeMap, fieldName)
			return resultUpdate(cu.IM{"dirty": true})
		},

		ct.TableEventFormUpdate: func() (re ct.ResponseEvent, err error) {
			return s.cls.updateMapField(evt, placeMap, resultUpdate)
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

		"tags": func() (re ct.ResponseEvent, err error) {
			return s.cls.editorTags(evt, placeMeta, resultUpdate)
		},

		"place_name": func() (re ct.ResponseEvent, err error) {
			place[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"place_type": func() (re ct.ResponseEvent, err error) {
			place[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"notes": func() (re ct.ResponseEvent, err error) {
			placeMeta[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"inactive": func() (re ct.ResponseEvent, err error) {
			placeMeta[fieldName] = cu.ToBoolean(value, false)
			return resultUpdate(cu.IM{"dirty": true})
		},

		"currency_code": func() (re ct.ResponseEvent, err error) {
			place[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"country": func() (re ct.ResponseEvent, err error) {
			address[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"state": func() (re ct.ResponseEvent, err error) {
			address[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"zip_code": func() (re ct.ResponseEvent, err error) {
			address[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"city": func() (re ct.ResponseEvent, err error) {
			address[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"street": func() (re ct.ResponseEvent, err error) {
			address[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},
	}

	if fn, ok := fieldMap[fieldName]; ok {
		return fn()
	}
	return evt, nil
}

func (s *PlaceService) Response(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
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
