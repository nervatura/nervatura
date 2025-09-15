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

func (cls *ClientService) employeeData(ds *api.DataStore, user, params cu.IM) (data cu.IM, err error) {
	data = cu.IM{
		"employee": cu.IM{
			"employee_meta": cu.IM{
				"start_date": "",
				"end_date":   "",
				"tags":       []string{},
			},
			"employee_map": cu.IM{},
			"address": cu.IM{
				"tags":        []string{},
				"address_map": cu.IM{},
			},
			"contact": cu.IM{
				"tags":        []string{},
				"contact_map": cu.IM{},
			},
			"events":     []cu.IM{},
			"time_stamp": md.TimeDateTime{Time: time.Now()},
		},
		"config_map":    cu.IM{},
		"config_data":   cu.IM{},
		"config_report": cu.IM{},
		"user":          user,
		"dirty":         false,
		"editor_icon":   ct.IconUser,
		"editor_title":  "",
	}

	if cu.ToString(params["employee_id"], "") != "" || cu.ToString(params["employee_code"], "") != "" {
		var employees []cu.IM = []cu.IM{}
		if employees, err = ds.StoreDataQuery(md.Query{
			Fields: []string{"*"}, From: "employee",
			Filters: []md.Filter{
				{Field: "deleted", Comp: "==", Value: false},
				{Field: "id", Comp: "==", Value: cu.ToInteger(params["employee_id"], 0)},
				{Or: true, Field: "code", Comp: "==", Value: cu.ToString(params["employee_code"], "")},
			},
		}, false); err != nil {
			return data, err
		}
		if len(employees) > 0 {
			data["employee"] = employees[0]
			data["editor_title"] = cu.ToString(employees[0]["code"], "")
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
			{Field: "report_type", Comp: "==", Value: "EMPLOYEE"},
		},
	}, false); err != nil {
		return data, err
	}
	data["config_report"] = rows

	return data, err
}

func (cls *ClientService) employeeUpdate(ds *api.DataStore, data cu.IM) (editor cu.IM, err error) {
	var employee md.Employee = md.Employee{}
	ut.ConvertToType(data["employee"], &employee)
	values := cu.IM{}
	if employee.Code != "" {
		values["code"] = employee.Code
	}

	ut.ConvertByteToIMData(employee.Address, values, "address")
	ut.ConvertByteToIMData(employee.Contact, values, "contact")
	ut.ConvertByteToIMData(employee.Events, values, "events")
	ut.ConvertByteToIMData(employee.EmployeeMeta, values, "employee_meta")
	ut.ConvertByteToIMData(employee.EmployeeMap, values, "employee_map")

	var employeeID int64
	newEmployee := (employee.Id == 0)
	update := md.Update{Values: values, Model: "employee"}
	if !newEmployee {
		update.IDKey = employee.Id
	}
	if employeeID, err = ds.StoreDataUpdate(update); err == nil && newEmployee {
		var employees []cu.IM = []cu.IM{}
		if employees, err = ds.StoreDataGet(cu.IM{"id": employeeID, "model": "employee"}, true); err == nil {
			data["employee"] = employees[0]
			data["editor_title"] = cu.ToString(cu.ToIM(employees[0], cu.IM{})["code"], "")
		}
	}
	return data, err
}

func (cls *ClientService) employeeDelete(ds *api.DataStore, employeeID int64) (err error) {
	return ds.DataDelete("employee", employeeID, "")
}

func (cls *ClientService) employeeResponseFormNext(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := cls.getDataStore(client.Ticket.Database)
	employee := cu.ToIM(stateData["employee"], cu.IM{})
	employeeMeta := cu.ToIM(employee["employee_meta"], cu.IM{})
	employeeMap := cu.ToIM(employee["employee_map"], cu.IM{})

	resultUpdate := func(dirty bool) (re ct.ResponseEvent, err error) {
		employee["employee_meta"] = employeeMeta
		employee["employee_map"] = employeeMap
		stateData["employee"] = employee
		if dirty {
			stateData["dirty"] = dirty
		}
		client.SetEditor("employee", cu.ToString(stateData["view"], ""), stateData)
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
			if err = cls.employeeDelete(ds, cu.ToInteger(employee["id"], 0)); err != nil {
				return evt, err
			}
			client.ResetEditor()
			return evt, err
		},

		"editor_add_tag": func() (re ct.ResponseEvent, err error) {
			tag := cu.ToString(frmValue["value"], "")
			if tag != "" {
				tags := ut.ToStringArray(employeeMeta["tags"])
				if !slices.Contains(tags, tag) {
					tags = append(tags, tag)
					employeeMeta["tags"] = tags
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
				Key:          "employee",
				Code:         cu.ToString(employee["code"], ""),
				Filters:      []any{},
				Columns:      map[string]bool{},
				TimeStamp:    md.TimeDateTime{Time: time.Now()},
			}
			return cls.addBookmark(evt, bookmark), nil
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
			employeeMap[mapField] = code
			stateData["map_field"] = ""
			return resultUpdate(true)
		},
	}

	if fn, ok := nextMap[cu.ToString(frmData["next"], "")]; ok {
		return fn()
	}
	return evt, err
}

func (cls *ClientService) employeeResponseFormEvent(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	employee := cu.ToIM(stateData["employee"], cu.IM{})

	frmValues := cu.ToIM(evt.Value, cu.IM{})
	frmValue := cu.ToIM(frmValues["value"], cu.IM{})
	frmData := cu.ToIM(frmValues["data"], cu.IM{})
	form := cu.ToIM(frmData["form"], cu.IM{})
	frmIndex := cu.ToInteger(form["index"], 0)
	frmKey := cu.ToString(form["key"], "")
	frmBaseValues := cu.ToIM(form["data"], cu.IM{})
	frmEvent := cu.ToString(frmValues["event"], "")
	rows := cu.ToIMA(employee[frmKey], []cu.IM{})
	if srows, found := stateData[frmKey]; found && (len(rows) == 0) {
		rows = cu.ToIMA(srows, []cu.IM{})
	}
	delete := (cu.ToString(frmValue["form_delete"], "") != "")

	resultUpdate := func() (re ct.ResponseEvent, err error) {
		if _, found := employee[frmKey]; found {
			employee[frmKey] = rows
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
			return cls.editorFormOK(evt, rows, customValues)
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
				return cls.editorFormTags(cu.IM{"row_field": fieldName}, evt)
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

func (cls *ClientService) employeeResponseSideMenu(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	employee := cu.ToIM(stateData["employee"], cu.IM{})
	ds := cls.getDataStore(client.Ticket.Database)

	menuMap := map[string]func() (re ct.ResponseEvent, err error){
		"editor_save": func() (re ct.ResponseEvent, err error) {
			if stateData, err = cls.employeeUpdate(ds, stateData); err != nil {
				return evt, err
			}
			stateData["dirty"] = false
			client.SetEditor("employee", cu.ToString(stateData["view"], ""), stateData)
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
			return cls.setEditor(evt, "employee",
				cu.IM{
					"session_id": client.Ticket.SessionID,
				}), nil
		},

		"editor_report": func() (re ct.ResponseEvent, err error) {
			return cls.showReportSelector(evt, "EMPLOYEE", cu.ToString(employee["code"], ""))
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

func (cls *ClientService) employeeResponseEditorField(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := cls.getDataStore(client.Ticket.Database)
	employee := cu.ToIM(stateData["employee"], cu.IM{})
	employeeMeta := cu.ToIM(employee["employee_meta"], cu.IM{})
	employeeMap := cu.ToIM(employee["employee_map"], cu.IM{})
	address := cu.ToIM(employee["address"], cu.IM{})
	contact := cu.ToIM(employee["contact"], cu.IM{})

	resultUpdate := func(params cu.IM) (re ct.ResponseEvent, err error) {
		employee["employee_meta"] = employeeMeta
		employee["employee_map"] = employeeMap
		employee["address"] = address
		employee["contact"] = contact
		stateData["employee"] = employee
		if cu.ToBoolean(params["dirty"], false) {
			stateData["dirty"] = true
		}
		client.SetEditor("employee", cu.ToString(stateData["view"], ""), stateData)
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
					ut.ConvertToType(md.Event{
						Tags:     []string{},
						EventMap: cu.IM{},
					}, &event)
					return event
				},
			}
			if slices.Contains([]string{"events"}, view) {
				getBase := func() (base cu.IM) {
					if _, found := employee[view]; found {
						return employee
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
			return cls.addMapField(evt, employeeMap, resultUpdate)
		},

		ct.TableEventFormDelete: func() (re ct.ResponseEvent, err error) {
			valueData := cu.ToIM(values["value"], cu.IM{})
			row := cu.ToIM(valueData["row"], cu.IM{})
			fieldName := cu.ToString(row["field_name"], "")
			delete(employeeMap, fieldName)
			return resultUpdate(cu.IM{"dirty": true})
		},

		ct.TableEventFormUpdate: func() (re ct.ResponseEvent, err error) {
			return cls.updateMapField(evt, employeeMap, resultUpdate)
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
			if err = cls.insertPrintQueue(ds, modalData); err == nil {
				return cls.evtMsg(evt.Name, evt.TriggerName, client.Msg("report_add_queue"), ct.ToastTypeSuccess, 5), nil
			}
			return evt, err
		},

		"tags": func() (re ct.ResponseEvent, err error) {
			return cls.editorTags(evt, employeeMeta, resultUpdate)
		},

		"start_date": func() (re ct.ResponseEvent, err error) {
			employeeMeta[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"end_date": func() (re ct.ResponseEvent, err error) {
			employeeMeta[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"notes": func() (re ct.ResponseEvent, err error) {
			employeeMeta[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"inactive": func() (re ct.ResponseEvent, err error) {
			employeeMeta[fieldName] = cu.ToBoolean(value, false)
			return resultUpdate(cu.IM{"dirty": true})
		},

		"first_name": func() (re ct.ResponseEvent, err error) {
			contact[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"surname": func() (re ct.ResponseEvent, err error) {
			contact[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"status": func() (re ct.ResponseEvent, err error) {
			contact[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"email": func() (re ct.ResponseEvent, err error) {
			contact[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"phone": func() (re ct.ResponseEvent, err error) {
			contact[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"mobile": func() (re ct.ResponseEvent, err error) {
			contact[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"contact_notes": func() (re ct.ResponseEvent, err error) {
			contact["notes"] = value
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

		"address_notes": func() (re ct.ResponseEvent, err error) {
			address["notes"] = value
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

func (cls *ClientService) employeeResponse(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	switch evt.Name {
	case ct.FormEventOK:
		return cls.employeeResponseFormNext(evt)

	case ct.ClientEventForm:
		return cls.employeeResponseFormEvent(evt)

	case ct.ClientEventSideMenu:
		return cls.employeeResponseSideMenu(evt)

	default:
		return cls.employeeResponseEditorField(evt)
	}
}
