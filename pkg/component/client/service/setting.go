package service

import (
	"slices"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	api "github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
	st "github.com/nervatura/nervatura/v6/pkg/static"
)

func (cls *ClientService) settingData(ds *api.DataStore, user, _ cu.IM) (data cu.IM, err error) {
	userConfig := cu.ToIM(user["auth_map"], cu.IM{})
	data = cu.IM{
		"setting": cu.IM{
			"id":          cu.ToInteger(user["id"], 0),
			"lang":        cu.ToString(userConfig["lang"], st.DefaultLang),
			"theme":       cu.ToString(userConfig["theme"], st.DefaultTheme),
			"export_sep":  cu.ToString(userConfig["export_sep"], st.DefaultExportSep),
			"page_size":   cu.ToString(userConfig["page_size"], st.DefaultPaperSize),
			"orientation": cu.ToString(userConfig["orientation"], st.DefaultOrientation),
			"pagination":  cu.ToString(userConfig["pagination"], st.DefaultPagination),
			"history":     cu.ToInteger(userConfig["history"], st.DefaultHistory),
			"password":    "",
			"confirm":     "",
		},
		"config_data":   []cu.IM{},
		"config_map":    []cu.IM{},
		"config_values": []cu.IM{},
		"dirty":         false,
		"editor_icon":   ct.IconCog,
		"editor_title":  "",
	}

	var rows []cu.IM = []cu.IM{}
	if rows, err = ds.StoreDataQuery(md.Query{
		Fields: []string{"*"}, From: "config_data",
	}, false); err != nil {
		return data, err
	}
	data["config_data"] = rows

	if rows, err = ds.StoreDataQuery(md.Query{
		Fields: []string{"*"}, From: "config_map", OrderBy: []string{"field_name"},
	}, false); err != nil {
		return data, err
	}
	data["config_map"] = rows

	if rows, err = ds.StoreDataQuery(md.Query{
		Fields: []string{"*"}, From: "config",
		Filters: []md.Filter{
			{Field: "config_type", Comp: "!=", Value: "CONFIG_REPORT"},
		},
	}, false); err != nil {
		return data, err
	}
	data["config_values"] = rows

	return data, nil
}

func (cls *ClientService) settingUpdate(ds *api.DataStore, user, data cu.IM) (err error) {
	values := cu.IM{}
	config, err := ds.ConvertToByte(data)
	if err == nil {
		values["auth_map"] = string(config[:])
		_, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "auth", IDKey: cu.ToInteger(user["id"], 0)})
	}

	return err
}

func (cls *ClientService) configUpdate(ds *api.DataStore, data cu.IM) (editor cu.IM, err error) {
	var configData md.Config = md.Config{}
	ut.ConvertToType(data, &configData)
	values := cu.IM{
		"config_type": configData.ConfigType.String(),
	}
	if configData.Code != "" {
		values["code"] = configData.Code
	}

	ut.ConvertByteToIMData(configData.Data, values, "data")

	var configID int64
	newConfig := (configData.Id == 0)
	update := md.Update{Values: values, Model: "config"}
	if !newConfig {
		update.IDKey = configData.Id
	}
	if configID, err = ds.StoreDataUpdate(update); err == nil && newConfig {
		var configs []cu.IM = []cu.IM{}
		if configs, err = ds.StoreDataGet(cu.IM{"id": configID, "model": "config"}, true); err == nil {
			data = configs[0]
		}
	}
	return data, err
}

func (cls *ClientService) settingPassword(ds *api.DataStore, user, data cu.IM) (err error) {
	return ds.UserPassword(
		cu.ToString(user["code"], ""), cu.ToString(data["password"], ""), cu.ToString(data["confirm"], ""),
	)
}

func (cls *ClientService) settingResponseSideMenu(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	stateData["view"] = cu.ToString(evt.Value, "")

	/*
		menuMap := map[string]func() (re ct.ResponseEvent, err error){
			"config_map": func() (re ct.ResponseEvent, err error) {
				stateData["view"] = "config_map"
				return evt, err
			},
		}

		if fn, ok := menuMap[cu.ToString(evt.Value, "")]; ok {
			return fn()
		}
	*/

	return evt, err
}

func (cls *ClientService) settingResponseFormNext(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := cls.getDataStore(client.Ticket.Database)
	//setting := cu.ToIM(stateData["setting"], cu.IM{})
	configValues := cu.ToIMA(stateData["config_values"], []cu.IM{})

	frmValues := cu.ToIM(evt.Value, cu.IM{})
	frmData := cu.ToIM(frmValues["data"], cu.IM{})
	frmValue := cu.ToIM(frmValues["value"], cu.IM{})

	nextMap := map[string]func() (re ct.ResponseEvent, err error){
		"form_add_tag": func() (re ct.ResponseEvent, err error) {
			tag := cu.ToString(frmValue["value"], "")
			frmKey := cu.ToString(frmData["frm_key"], "")
			frmIndex := cu.ToInteger(frmData["frm_index"], 0)
			row := cu.ToIM(frmData["row"], cu.IM{})
			rowField := cu.ToString(frmData["row_field"], "")
			metaName := ut.MetaName(row, cu.ToString(frmData["meta_name"], "_meta"))
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

		"config_delete": func() (re ct.ResponseEvent, err error) {
			if idx := slices.IndexFunc(configValues, func(c cu.IM) bool {
				return cu.ToString(c["code"], "") == cu.ToString(frmData["code"], "")
			}); idx > int(-1) {
				if _, err = ds.StoreDataUpdate(md.Update{Model: "config", IDKey: cu.ToInteger(configValues[idx]["id"], 0)}); err == nil {
					configValues = append(configValues[:idx], configValues[idx+1:]...)
					stateData["config_values"] = configValues
				}
			}
			return evt, err
		},
	}

	if fn, ok := nextMap[cu.ToString(frmData["next"], "")]; ok {
		return fn()
	}
	return evt, err
}

func (cls *ClientService) settingResponseFormEvent(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := cls.getDataStore(client.Ticket.Database)
	configValues := cu.ToIMA(stateData["config_values"], []cu.IM{})

	frmValues := cu.ToIM(evt.Value, cu.IM{})
	frmValue := cu.ToIM(frmValues["value"], cu.IM{})
	frmData := cu.ToIM(frmValues["data"], cu.IM{})
	form := cu.ToIM(frmData["form"], cu.IM{})
	frmIndex := cu.ToInteger(form["index"], 0)
	//frmKey := cu.ToString(form["key"], "")
	//frmBaseValues := cu.ToIM(form["data"], cu.IM{})
	frmEvent := cu.ToString(frmValues["event"], "")
	//rows := cu.ToIMA(stateData[frmKey], []cu.IM{})
	configValue := cu.ToIM(form["data"], cu.IM{})
	configMeta := cu.ToIM(configValue["data"], cu.IM{})
	delete := (cu.ToString(frmValue["form_delete"], "") != "")

	eventMap := map[string]func() (re ct.ResponseEvent, err error){
		ct.FormEventOK: func() (re ct.ResponseEvent, err error) {
			switch cu.ToString(form["key"], "") {
			case "config_map":
				configMeta["field_name"] = cu.ToString(frmValue["field_name"], "")
				configMeta["field_type"] = cu.ToString(frmValue["field_type"], "")
				configMeta["description"] = cu.ToString(frmValue["description"], "")
				if idx := slices.IndexFunc(configValues, func(c cu.IM) bool {
					return cu.ToString(c["code"], "") == cu.ToString(configValue["code"], "")
				}); idx > int(-1) {
					if configValue, err = cls.configUpdate(ds, configValue); err == nil {
						configValues[idx] = configValue
						stateData["config_values"] = configValues
					}
				}
			}
			return evt, err
		},

		ct.FormEventCancel: func() (re ct.ResponseEvent, err error) {
			if delete {
				switch cu.ToString(form["key"], "") {
				case "config_map":
					configValue := cu.ToIM(form["data"], cu.IM{})
					modal := cu.IM{
						"warning_label":   client.Msg("inputbox_delete"),
						"warning_message": "",
						"next":            "config_delete",
						"code":            cu.ToString(configValue["code"], ""),
					}
					client.SetForm("warning", modal, 0, true)
				}
			}
			if cu.ToString(configValue["code"], "") == "" {
				if idx := slices.IndexFunc(configValues, func(c cu.IM) bool {
					return cu.ToString(c["code"], "") == ""
				}); idx > int(-1) {
					configValues = append(configValues[:idx], configValues[idx+1:]...)
					stateData["config_values"] = configValues
				}
			}
			return evt, err
		},

		ct.FormEventChange: func() (re ct.ResponseEvent, err error) {
			fieldName := cu.ToString(frmValues["name"], "")
			switch fieldName {
			case "tags":
				return cls.editorFormTags(cu.IM{"row_field": fieldName, "meta_name": "data"}, evt)
			case "filter":
				opt := []ct.SelectOption{}
				ft := md.MapFilter(0)
				for _, ftKey := range ft.Keys() {
					opt = append(opt, ct.SelectOption{
						Value: ftKey, Text: ftKey,
					})
				}
				return cls.editorFormTags(cu.IM{"row_field": fieldName, "meta_name": "data",
					"options": opt, "value": opt[0].Value, "is_null": false, "form_key": "select",
					"icon": ct.IconFilter, "title": client.Msg("inputbox_new_filter"),
					"label": client.Msg("inputbox_enter_filter")}, evt)

			case "field_name", "field_type", "description":
				configMeta[fieldName] = frmValues["value"]
				client.SetForm(cu.ToString(stateData["view"], ""), configValue, frmIndex, false)
				//cu.ToSM(evt.Header, cu.SM{})[ct.HeaderReswap] = ct.SwapNone

			default:
				//frmBaseValues[fieldName] = frmValues["value"]
				//cu.ToSM(evt.Header, cu.SM{})[ct.HeaderReswap] = ct.SwapNone
			}
			return evt, nil
		},
	}

	if fn, ok := eventMap[frmEvent]; ok {
		return fn()
	}

	return evt, err

}

func (cls *ClientService) settingResponseEditorField(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, stateKey, stateData := client.GetStateData()
	ds := cls.getDataStore(client.Ticket.Database)
	user := client.Ticket.User
	setting := cu.ToIM(stateData["setting"], cu.IM{})
	configValues := cu.ToIMA(stateData["config_values"], []cu.IM{})

	values := cu.ToIM(evt.Value, cu.IM{})
	fieldName := cu.ToString(values["name"], "")
	value := cu.ToString(values["value"], "")

	configUpdate := func() (re ct.ResponseEvent, err error) {
		setting[fieldName] = value
		userConfig := cu.ToIM(user["auth_map"], cu.IM{})
		userConfig[fieldName] = value
		if err = cls.settingUpdate(ds, user, userConfig); err != nil {
			return evt, err
		}
		client.Ticket.User["auth_map"] = userConfig
		client.SetProperty(fieldName, value)
		stateData["setting"] = setting
		client.SetEditor(stateKey, cu.ToString(stateData["view"], ""), stateData)
		return evt, err
	}

	resultUpdate := func(dirty bool) (re ct.ResponseEvent, err error) {
		setting[fieldName] = value
		stateData["setting"] = setting
		if dirty {
			stateData["dirty"] = dirty
		}
		client.SetEditor("setting", cu.ToString(stateData["view"], ""), stateData)
		return evt, err
	}

	fieldMap := map[string]func() (re ct.ResponseEvent, err error){
		"theme": func() (re ct.ResponseEvent, err error) {
			return configUpdate()
		},
		"lang": func() (re ct.ResponseEvent, err error) {
			return configUpdate()
		},
		"page_size": func() (re ct.ResponseEvent, err error) {
			return configUpdate()
		},
		"orientation": func() (re ct.ResponseEvent, err error) {
			return configUpdate()
		},
		"pagination": func() (re ct.ResponseEvent, err error) {
			return configUpdate()
		},
		"history": func() (re ct.ResponseEvent, err error) {
			return configUpdate()
		},
		"export_sep": func() (re ct.ResponseEvent, err error) {
			return configUpdate()
		},
		"password": func() (re ct.ResponseEvent, err error) {
			return resultUpdate(true)
		},
		"confirm": func() (re ct.ResponseEvent, err error) {
			return resultUpdate(true)
		},
		"change_password": func() (re ct.ResponseEvent, err error) {
			if err = cls.settingPassword(ds, user, setting); err != nil {
				return evt, err
			}
			setting["password"] = ""
			setting["confirm"] = ""
			stateData["setting"] = setting
			stateData["dirty"] = false
			client.SetEditor(stateKey, cu.ToString(stateData["view"], ""), stateData)
			return cls.evtMsg(evt.Name, evt.TriggerName, client.Msg("setting_password_ok"), ct.ToastTypeSuccess, 5), nil
		},
		"config_map": func() (re ct.ResponseEvent, err error) {
			event := cu.ToString(cu.ToIM(evt.Value, cu.IM{})["event"], "")
			evValue := cu.ToIM(cu.ToIM(evt.Value, cu.IM{})["value"], cu.IM{})
			row := cu.ToIM(evValue["row"], cu.IM{})
			index := cu.ToInteger(evValue["index"], 0)
			switch event {
			case ct.ListEventEditItem:
				client.SetForm("config_map", row, index, false)

			case ct.ListEventDelete:
				modal := cu.IM{
					"warning_label":   client.Msg("inputbox_delete"),
					"warning_message": "",
					"next":            "config_delete",
					"code":            cu.ToString(row["code"], ""),
				}
				client.SetForm("warning", modal, 0, true)

			case ct.ListEventAddItem:
				configValue := cu.IM{
					"id":          0,
					"code":        "",
					"config_type": md.ConfigTypeMap.String(),
					"data": cu.IM{
						"field_name":  "",
						"field_type":  md.FieldTypeString.String(),
						"description": "",
						"tags":        []string{},
						"filter":      []md.MapFilter{},
					},
				}
				configValues = append(configValues, configValue)
				stateData["config_values"] = configValues
				client.SetForm("config_map", configValue, int64(len(configValues)-1), false)
			}
			return evt, err
		},
	}

	if fn, ok := fieldMap[fieldName]; ok {
		return fn()
	}
	return evt, err
}

func (cls *ClientService) settingResponse(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	switch evt.Name {
	case ct.FormEventOK:
		return cls.settingResponseFormNext(evt)

	case ct.ClientEventForm:
		return cls.settingResponseFormEvent(evt)

	case ct.ClientEventSideMenu:
		return cls.settingResponseSideMenu(evt)

	default:
		return cls.settingResponseEditorField(evt)
	}
}
