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

type CustomerService struct {
	cls *ClientService
}

func NewCustomerService(cls *ClientService) *CustomerService {
	return &CustomerService{
		cls: cls,
	}
}

func (s *CustomerService) Data(evt ct.ResponseEvent, params cu.IM) (data cu.IM, err error) {
	client := evt.Trigger.(*ct.Client)
	ds := s.cls.getDataStore(client.Ticket.Database)
	user := cu.ToIM(client.Ticket.User, cu.IM{})

	data = cu.IM{
		"customer": cu.IM{
			"customer_type": md.CustomerType(0),
			"customer_meta": cu.IM{
				"tags": []string{},
			},
			"customer_map": cu.IM{},
			"addresses":    []cu.IM{},
			"contacts":     []cu.IM{},
			"events":       []cu.IM{},
		},
		"config_map":    cu.IM{},
		"config_data":   cu.IM{},
		"config_report": cu.IM{},
		"user":          user,
		"dirty":         false,
		"editor_icon":   ct.IconUser,
		"editor_title":  "",
	}

	if cu.ToString(params["customer_id"], "") != "" || cu.ToString(params["customer_code"], "") != "" {
		var customers []cu.IM = []cu.IM{}
		if customers, err = ds.StoreDataQuery(md.Query{
			Fields: []string{"*"}, From: "customer",
			Filters: []md.Filter{
				{Field: "deleted", Comp: "==", Value: false},
				{Field: "id", Comp: "==", Value: cu.ToInteger(params["customer_id"], 0)},
				{Or: true, Field: "code", Comp: "==", Value: cu.ToString(params["customer_code"], "")},
			},
		}, false); err != nil {
			return data, err
		}
		if len(customers) > 0 {
			data["customer"] = customers[0]
			data["editor_title"] = cu.ToString(customers[0]["code"], "")
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
			{Field: "report_type", Comp: "==", Value: "CUSTOMER"},
		},
	}, false); err != nil {
		return data, err
	}
	data["config_report"] = rows

	return data, err
}

func (s *CustomerService) Response(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
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

func (s *CustomerService) update(ds *api.DataStore, data cu.IM) (editor cu.IM, err error) {
	var customer md.Customer = md.Customer{}
	ut.ConvertToType(data["customer"], &customer)
	values := cu.IM{
		"customer_type": customer.CustomerType.String(),
		"customer_name": customer.CustomerName,
	}
	if customer.Code != "" {
		values["code"] = customer.Code
	}

	ut.ConvertByteToIMData(customer.Contacts, values, "contacts")
	ut.ConvertByteToIMData(customer.Addresses, values, "addresses")
	ut.ConvertByteToIMData(customer.Events, values, "events")
	ut.ConvertByteToIMData(customer.CustomerMeta, values, "customer_meta")
	ut.ConvertByteToIMData(customer.CustomerMap, values, "customer_map")

	var customerID int64
	newCustomer := (customer.Id == 0)
	update := md.Update{Values: values, Model: "customer"}
	if !newCustomer {
		update.IDKey = customer.Id
	}
	if customerID, err = ds.StoreDataUpdate(update); err == nil && newCustomer {
		var customers []cu.IM = []cu.IM{}
		if customers, err = ds.StoreDataGet(cu.IM{"id": customerID, "model": "customer"}, true); err == nil {
			data["customer"] = customers[0]
			data["editor_title"] = cu.ToString(cu.ToIM(customers[0], cu.IM{})["code"], "")
		}
	}
	return data, err
}

func (s *CustomerService) delete(ds *api.DataStore, customerID int64) (err error) {
	return ds.DataDelete("customer", customerID, "")
}

func (s *CustomerService) formNext(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := s.cls.getDataStore(client.Ticket.Database)
	customer := cu.ToIM(stateData["customer"], cu.IM{})
	customerMeta := cu.ToIM(customer["customer_meta"], cu.IM{})
	customerMap := cu.ToIM(customer["customer_map"], cu.IM{})

	resultUpdate := func(dirty bool) (re ct.ResponseEvent, err error) {
		customer["customer_meta"] = customerMeta
		customer["customer_map"] = customerMap
		stateData["customer"] = customer
		if dirty {
			stateData["dirty"] = dirty
		}
		client.SetEditor("customer", cu.ToString(stateData["view"], ""), stateData)
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
			if err = s.delete(ds, cu.ToInteger(customer["id"], 0)); err != nil {
				return evt, err
			}
			client.ResetEditor()
			return evt, err
		},

		"editor_add_tag": func() (re ct.ResponseEvent, err error) {
			tag := cu.ToString(frmValue["value"], "")
			if tag != "" {
				tags := ut.ToStringArray(customerMeta["tags"])
				if !slices.Contains(tags, tag) {
					tags = append(tags, tag)
					customerMeta["tags"] = tags
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
				Key:          "customer",
				Code:         cu.ToString(customer["code"], ""),
				Filters:      []any{},
				Columns:      map[string]bool{},
				TimeStamp:    md.TimeDateTime{Time: time.Now()},
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
			customerMap[mapField] = code
			stateData["map_field"] = ""
			return resultUpdate(true)
		},
	}

	if fn, ok := nextMap[cu.ToString(frmData["next"], "")]; ok {
		return fn()
	}
	return evt, err
}

func (s *CustomerService) formEvent(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	customer := cu.ToIM(stateData["customer"], cu.IM{})

	frmValues := cu.ToIM(evt.Value, cu.IM{})
	frmValue := cu.ToIM(frmValues["value"], cu.IM{})
	frmData := cu.ToIM(frmValues["data"], cu.IM{})
	form := cu.ToIM(frmData["form"], cu.IM{})
	frmIndex := cu.ToInteger(form["index"], 0)
	frmKey := cu.ToString(form["key"], "")
	frmBaseValues := cu.ToIM(form["data"], cu.IM{})
	frmEvent := cu.ToString(frmValues["event"], "")
	rows := cu.ToIMA(customer[frmKey], []cu.IM{})
	if srows, found := stateData[frmKey]; found && (len(rows) == 0) {
		rows = cu.ToIMA(srows, []cu.IM{})
	}
	delete := (cu.ToString(frmValue["form_delete"], "") != "")

	resultUpdate := func() (re ct.ResponseEvent, err error) {
		if _, found := customer[frmKey]; found {
			customer[frmKey] = rows
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

func (s *CustomerService) sideMenu(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	customer := cu.ToIM(stateData["customer"], cu.IM{})
	ds := s.cls.getDataStore(client.Ticket.Database)

	menuMap := map[string]func() (re ct.ResponseEvent, err error){
		"editor_save": func() (re ct.ResponseEvent, err error) {
			if stateData, err = s.update(ds, stateData); err != nil {
				return evt, err
			}
			stateData["dirty"] = false
			client.SetEditor("customer", cu.ToString(stateData["view"], ""), stateData)
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
			return s.cls.setEditor(evt, "customer",
				cu.IM{
					"session_id": client.Ticket.SessionID,
				}), nil
		},

		"editor_report": func() (re ct.ResponseEvent, err error) {
			return s.cls.showReportSelector(evt, "CUSTOMER", cu.ToString(customer["code"], ""))
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

func (s *CustomerService) editorField(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := s.cls.getDataStore(client.Ticket.Database)
	customer := cu.ToIM(stateData["customer"], cu.IM{})
	customerMeta := cu.ToIM(customer["customer_meta"], cu.IM{})
	customerMap := cu.ToIM(customer["customer_map"], cu.IM{})

	resultUpdate := func(params cu.IM) (re ct.ResponseEvent, err error) {
		customer["customer_meta"] = customerMeta
		customer["customer_map"] = customerMap
		stateData["customer"] = customer
		if cu.ToBoolean(params["dirty"], false) {
			stateData["dirty"] = true
		}
		client.SetEditor("customer", cu.ToString(stateData["view"], ""), stateData)
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
				"addresses": func() cu.IM {
					var address cu.IM
					ut.ConvertToType(md.Address{
						Tags:       []string{},
						AddressMap: cu.IM{},
					}, &address)
					return address
				},
				"contacts": func() cu.IM {
					var contact cu.IM
					ut.ConvertToType(md.Contact{
						Tags:       []string{},
						ContactMap: cu.IM{},
					}, &contact)
					return contact
				},
				"events": func() cu.IM {
					var event cu.IM
					ut.ConvertToType(md.Event{
						Tags:     []string{},
						EventMap: cu.IM{},
					}, &event)
					return event
				},
			}
			if slices.Contains([]string{"addresses", "contacts", "events"}, view) {
				getBase := func() (base cu.IM) {
					if _, found := customer[view]; found {
						return customer
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
			return s.cls.addMapField(evt, customerMap, resultUpdate)
		},

		ct.TableEventFormDelete: func() (re ct.ResponseEvent, err error) {
			valueData := cu.ToIM(values["value"], cu.IM{})
			row := cu.ToIM(valueData["row"], cu.IM{})
			fieldName := cu.ToString(row["field_name"], "")
			delete(customerMap, fieldName)
			return resultUpdate(cu.IM{"dirty": true})
		},

		ct.TableEventFormUpdate: func() (re ct.ResponseEvent, err error) {
			return s.cls.updateMapField(evt, customerMap, resultUpdate)
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
			return s.cls.editorTags(evt, customerMeta, resultUpdate)
		},

		"customer_name": func() (re ct.ResponseEvent, err error) {
			customer[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"customer_type": func() (re ct.ResponseEvent, err error) {
			customer[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"notes": func() (re ct.ResponseEvent, err error) {
			customerMeta[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"terms": func() (re ct.ResponseEvent, err error) {
			customerMeta[fieldName] = cu.ToInteger(value, 0)
			return resultUpdate(cu.IM{"dirty": true})
		},

		"credit_limit": func() (re ct.ResponseEvent, err error) {
			customerMeta[fieldName] = cu.ToFloat(value, 0)
			return resultUpdate(cu.IM{"dirty": true})
		},

		"discount": func() (re ct.ResponseEvent, err error) {
			customerMeta[fieldName] = cu.ToFloat(value, 0)
			return resultUpdate(cu.IM{"dirty": true})
		},

		"inactive": func() (re ct.ResponseEvent, err error) {
			customerMeta[fieldName] = cu.ToBoolean(value, false)
			return resultUpdate(cu.IM{"dirty": true})
		},

		"tax_free": func() (re ct.ResponseEvent, err error) {
			customerMeta[fieldName] = cu.ToBoolean(value, false)
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
