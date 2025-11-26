package service

import (
	"slices"
	"time"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

type ShippingService struct {
	cls *ClientService
}

func NewShippingService(cls *ClientService) *ShippingService {
	return &ShippingService{
		cls: cls,
	}
}

func (s *ShippingService) Data(evt ct.ResponseEvent, params cu.IM) (data cu.IM, err error) {
	client := evt.Trigger.(*ct.Client)
	ds := s.cls.getDataStore(client.Ticket.Database)
	user := cu.ToIM(client.Ticket.User, cu.IM{})

	data = cu.IM{
		"shipping":     cu.IM{},
		"items":        []cu.IM{},
		"movements":    []cu.IM{},
		"places":       []cu.IM{},
		"create_items": []cu.IM{},
		"user":         user,
		"dirty":        false,
		"editor_icon":  ct.IconTruck,
		"editor_title": "",
	}

	var rows []cu.IM = []cu.IM{}
	var shipping cu.IM = cu.IM{}
	if cu.ToString(params["trans_id"], "") != "" || cu.ToString(params["trans_code"], "") != "" {
		if rows, err = ds.StoreDataQuery(md.Query{
			Fields: []string{"t.id", "t.code", "t.trans_type", "t.direction", "t.trans_date",
				"t.customer_code", "c.customer_name"},
			From: "trans t inner join customer c on t.customer_code = c.code",
			Filters: []md.Filter{
				{Field: "t.deleted", Comp: "==", Value: false},
				{Field: "c.deleted", Comp: "==", Value: false},
				{Field: "t.id", Comp: "==", Value: cu.ToInteger(params["trans_id"], 0)},
				{Or: true, Field: "t.code", Comp: "==", Value: cu.ToString(params["trans_code"], "")},
			},
		}, false); err != nil {
			return data, err
		}
		if len(rows) > 0 {
			data["shipping"] = cu.MergeIM(rows[0], cu.IM{
				"shipping_time": time.Now().Format("2006-01-02T00:00:00"),
				"place_code":    "",
			})
			data["editor_title"] = cu.ToString(rows[0]["code"], "")
		}
		shipping = cu.ToIM(data["shipping"], cu.IM{})

		if rows, err = ds.StoreDataQuery(md.Query{
			Fields: []string{"*"},
			From:   "item_shipping",
			Filters: []md.Filter{
				{Field: "trans_code", Comp: "==", Value: cu.ToString(shipping["code"], "")},
			},
			OrderBy: []string{"code"},
		}, false); err != nil {
			return data, err
		}
		items := []cu.IM{}
		direction := cu.ToString(shipping["direction"], "")
		for _, row := range rows {
			item := cu.MergeIM(row,
				cu.IM{
					"batch_no": "", "qty": 0, "difference": 0, "disabled": false,
					"turnover": cu.ToFloat(row["movement_qty"], 0)})
			if direction == md.DirectionOut.String() {
				item["turnover"] = -cu.ToFloat(row["movement_qty"], 0)
			}
			item["difference"] = cu.ToFloat(item["item_qty"], 0) - cu.ToFloat(item["turnover"], 0)
			if cu.ToFloat(item["difference"], 0) == 0 {
				item["disabled"] = true
			}
			items = append(items, item)
		}
		data["items"] = items

		if rows, err = ds.StoreDataQuery(md.Query{
			Fields: []string{
				"i.id", "i.code", "i.description", "i.qty", "mv.trans_code",
				"mv.shipping_time as shipping_date", "mv.qty as shipping_qty", "pr.product_name", "pr.unit",
				"pr.code as product_code",
			},
			From: `item_view i inner join movement_view mv on mv.item_code = i.code inner join product_view pr on pr.code = mv.product_code`,
			Filters: []md.Filter{
				{Field: "i.trans_code", Comp: "==", Value: cu.ToString(shipping["code"], "")},
			},
			OrderBy: []string{"mv.trans_code", "mv.id"},
		}, false); err != nil {
			return data, err
		}
		data["movements"] = rows
	}

	if rows, err = ds.StoreDataQuery(md.Query{
		Fields: []string{"*"},
		From:   "place",
		Filters: []md.Filter{
			{Field: "deleted", Comp: "==", Value: false},
			{Field: "place_type", Comp: "==", Value: md.PlaceTypeWarehouse.String()},
		},
		OrderBy: []string{"code"},
	}, false); err != nil {
		return data, err
	}
	data["places"] = rows
	if len(rows) > 0 {
		shipping["place_code"] = cu.ToString(rows[0]["code"], "")
	}

	return data, nil
}

func (s *ShippingService) Response(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	switch evt.Name {
	case ct.FormEventOK:
		return s.formNext(evt)

	//case ct.ClientEventForm:
	//	return evt, err

	case ct.ClientEventSideMenu:
		return s.sideMenu(evt)

	default:
		return s.editorField(evt)
	}
}

func (s *ShippingService) createDelivery(evt ct.ResponseEvent, movements []md.Movement) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := s.cls.getDataStore(client.Ticket.Database)

	shipping := cu.ToIM(stateData["shipping"], cu.IM{})
	shippingTime, _ := cu.StringToDateTime(cu.ToString(shipping["shipping_time"], ""))
	user := cu.ToIM(stateData["user"], cu.IM{})

	values := cu.IM{
		"trans_type": md.TransTypeDelivery.String(),
		"direction":  shipping["direction"],
		"trans_date": shippingTime.Format(time.DateOnly),
		"auth_code":  user["code"],
	}
	meta := md.TransMeta{
		PaidType:   md.PaidTypeTransfer,
		Status:     md.TransStatusNormal,
		TransState: md.TransStateOK,
		Worksheet:  md.TransMetaWorksheet{},
		Rent:       md.TransMetaRent{},
		Invoice:    md.TransMetaInvoice{},
		Tags:       []string{},
	}
	ut.ConvertByteToIMData(meta, values, "trans_meta")
	ut.ConvertByteToIMData(cu.IM{}, values, "trans_map")

	var transID int64
	var transCode string
	update := md.Update{Values: values, Model: "trans"}
	if transID, err = ds.StoreDataUpdate(update); err == nil {
		var rows []cu.IM = []cu.IM{}
		if rows, err = ds.StoreDataGet(cu.IM{"id": transID, "model": "trans"}, true); err == nil {
			transCode = cu.ToString(cu.ToIM(rows[0], cu.IM{})["code"], "")
		}
	}

	if err != nil {
		return s.cls.errorModal(evt, client.Msg("shipping_create"), client.Msg(err.Error()))
	}

	for _, mv := range movements {
		values := cu.IM{
			"movement_type": mv.MovementType.String(),
			"shipping_time": mv.ShippingTime.Format(time.RFC3339),
			"trans_code":    transCode,
			"product_code":  mv.ProductCode,
			"place_code":    shipping["place_code"],
			"item_code":     mv.ItemCode,
		}
		ut.ConvertByteToIMData(mv.MovementMeta, values, "movement_meta")
		ut.ConvertByteToIMData(mv.MovementMap, values, "movement_map")
		if _, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "movement"}); err != nil {
			return s.cls.errorModal(evt, client.Msg("shipping_create"), client.Msg(err.Error()))
		}
	}

	return s.cls.setEditor(evt, "shipping", cu.IM{
		"editor_view": "items",
		"trans_code":  shipping["code"],
		"session_id":  client.Ticket.SessionID,
	}), err
}

func (s *ShippingService) formNext(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	shipping := cu.ToIM(stateData["shipping"], cu.IM{})
	items := cu.ToIMA(stateData["items"], []cu.IM{})

	frmValues := cu.ToIM(evt.Value, cu.IM{})
	frmData := cu.ToIM(frmValues["data"], cu.IM{})

	shippingTime, _ := cu.StringToDateTime(cu.ToString(shipping["shipping_time"], ""))
	direction := cu.ToString(shipping["direction"], "")
	mvRow := func(item cu.IM, qty float64) (mv md.Movement) {
		mv = md.Movement{
			MovementType: md.MovementTypeInventory,
			ShippingTime: md.TimeDateTime{Time: shippingTime},
			ProductCode:  cu.ToString(item["product_code"], ""),
			PlaceCode:    cu.ToString(shipping["place_code"], ""),
			ItemCode:     cu.ToString(item["code"], ""),
			MovementMeta: md.MovementMeta{
				Qty:   qty,
				Notes: cu.ToString(item["batch_no"], ""),
				Tags:  []string{},
			},
			MovementMap: cu.IM{},
		}
		if direction == md.DirectionOut.String() {
			mv.MovementMeta.Qty = -mv.MovementMeta.Qty
		}
		return mv
	}

	nextMap := map[string]func() (re ct.ResponseEvent, err error){
		"editor_cancel": func() (re ct.ResponseEvent, err error) {
			evt = s.cls.setEditor(evt, "trans",
				cu.IM{
					"editor_view": "transitem_shipping",
					"trans_code":  cu.ToString(shipping["code"], ""),
					"session_id":  client.Ticket.SessionID,
				})
			return evt, err
		},
		"create_all": func() (re ct.ResponseEvent, err error) {
			mv := []md.Movement{}
			for _, item := range items {
				if cu.ToFloat(item["difference"], 0) > 0 {
					mv = append(mv, mvRow(item, cu.ToFloat(item["difference"], 0)))
				}
			}
			return s.createDelivery(evt, mv)
		},
		"create_item": func() (re ct.ResponseEvent, err error) {
			mv := []md.Movement{}
			for _, item := range items {
				if cu.ToFloat(item["qty"], 0) > 0 {
					mv = append(mv, mvRow(item, cu.ToFloat(item["qty"], 0)))
				}
			}
			return s.createDelivery(evt, mv)
		},
	}

	if fn, ok := nextMap[cu.ToString(frmData["next"], "")]; ok {
		return fn()
	}
	return evt, err
}

func (s *ShippingService) sideMenu(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	shipping := cu.ToIM(stateData["shipping"], cu.IM{})

	menuMap := map[string]func() (re ct.ResponseEvent, err error){
		"editor_cancel": func() (re ct.ResponseEvent, err error) {
			if cu.ToBoolean(stateData["dirty"], false) {
				modal := cu.IM{
					"warning_label":   client.Msg("inputbox_dirty"),
					"warning_message": client.Msg("inputbox_drop"),
					"next":            "editor_cancel",
				}
				client.SetForm("warning", modal, 0, true)
			} else {
				evt = s.cls.setEditor(evt, "trans",
					cu.IM{
						"editor_view": "transitem_shipping",
						"trans_code":  cu.ToString(shipping["code"], ""),
						"session_id":  client.Ticket.SessionID,
					})
			}
			return evt, err
		},

		"create_all": func() (re ct.ResponseEvent, err error) {
			modal := cu.IM{
				"title":           client.Msg("shipping_create"),
				"icon":            ct.IconTruck,
				"warning_label":   client.Msg("shipping_create_all"),
				"warning_message": client.Msg("inputbox_delete_info"),
				"next":            "create_all",
			}
			client.SetForm("warning", modal, 0, true)
			return evt, err
		},

		"create_item": func() (re ct.ResponseEvent, err error) {
			modal := cu.IM{
				"title":           client.Msg("shipping_create"),
				"icon":            ct.IconTruck,
				"warning_label":   client.Msg("shipping_create_item"),
				"warning_message": client.Msg("inputbox_delete_info"),
				"next":            "create_item",
			}
			client.SetForm("warning", modal, 0, true)
			return evt, err
		},
	}

	if fn, ok := menuMap[cu.ToString(evt.Value, "")]; ok {
		return fn()
	}

	return evt, err
}

func (s *ShippingService) editorField(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()

	resultUpdate := func(params cu.IM) {
		stateData[cu.ToString(params["key"], "")] = params["value"]
		if cu.ToBoolean(params["dirty"], false) {
			stateData["dirty"] = true
		}
		client.SetEditor("shipping", cu.ToString(stateData["view"], ""), stateData)
	}

	setShippingValue := func(fieldName string, value string) (re ct.ResponseEvent, err error) {
		shipping := cu.ToIM(stateData["shipping"], cu.IM{})
		shipping[fieldName] = value
		resultUpdate(cu.IM{"key": "shipping", "value": shipping, "dirty": false})
		return evt, err
	}

	values := cu.ToIM(evt.Value, cu.IM{})
	fieldName := cu.ToString(values["name"], "")
	value := cu.ToString(values["value"], "")
	valueData := cu.ToIM(values["value"], cu.IM{})
	row := cu.ToIM(valueData["row"], cu.IM{})

	fieldMap := map[string]func() (re ct.ResponseEvent, err error){
		ct.TableEventFormUpdate: func() (re ct.ResponseEvent, err error) {
			itemCode := cu.ToString(row["code"], "")
			items := cu.ToIMA(stateData["items"], []cu.IM{})
			if idx := slices.IndexFunc(items, func(c cu.IM) bool {
				return cu.ToString(c["code"], "") == itemCode
			}); idx > int(-1) {
				items[idx]["qty"] = cu.ToFloat(row["qty"], 0)
				items[idx]["batch_no"] = cu.ToString(row["batch_no"], "")
				resultUpdate(cu.IM{"key": "items", "value": items, "dirty": true})
			}
			return evt, err
		},
		ct.TableEventFormChange: func() (re ct.ResponseEvent, err error) {
			if trigger, ok := values["trigger"].(ct.ClientComponent); ok {
				rows := cu.ToIMA(trigger.GetProperty("rows"), []cu.IM{})
				filterIndex := cu.ToInteger(valueData["index"], 0)
				qty := cu.ToFloat(valueData["value"], 0)
				difference := cu.ToFloat(row["difference"], 0)
				if difference < qty {
					qty = difference
				}
				if qty < 0 {
					qty = 0
				}
				rows[filterIndex]["qty"] = qty
				trigger.SetProperty("rows", rows)
			}
			return evt, err
		},
		"place_code": func() (re ct.ResponseEvent, err error) {
			return setShippingValue(fieldName, value)
		},
		"shipping_time": func() (re ct.ResponseEvent, err error) {
			return setShippingValue(fieldName, value)
		},
	}
	if fn, ok := fieldMap[fieldName]; ok {
		return fn()
	}
	return evt, err
}
