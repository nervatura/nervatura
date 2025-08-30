package service

import (
	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	api "github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
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
		"config_data":  cu.IM{},
		"dirty":        false,
		"editor_icon":  ct.IconCog,
		"editor_title": "",
	}

	var rows []cu.IM = []cu.IM{}
	if rows, err = ds.StoreDataQuery(md.Query{
		Fields: []string{"*"}, From: "config_data",
	}, false); err != nil {
		return data, err
	}
	data["config_data"] = rows

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

func (cls *ClientService) settingPassword(ds *api.DataStore, user, data cu.IM) (err error) {
	return ds.UserPassword(
		cu.ToString(user["code"], ""), cu.ToString(data["password"], ""), cu.ToString(data["confirm"], ""),
	)
}

func (cls *ClientService) settingResponseSideMenu(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	/*
		//client := evt.Trigger.(*ct.Client)
		//_, _, stateData := client.GetStateData()

		menuMap := map[string]func() (re ct.ResponseEvent, err error){}

		if fn, ok := menuMap[cu.ToString(evt.Value, "")]; ok {
			return fn()
		}
	*/

	return evt, err
}

func (cls *ClientService) settingResponseFormNext(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	/*
		//client := evt.Trigger.(*ct.Client)
		//_, _, stateData := client.GetStateData()
		//ds := cls.getDataStore(client.Ticket.Database)
		//setting := cu.ToIM(stateData["setting"], cu.IM{})

		frmValues := cu.ToIM(evt.Value, cu.IM{})
		frmData := cu.ToIM(frmValues["data"], cu.IM{})
		//frmValue := cu.ToIM(frmValues["value"], cu.IM{})

		nextMap := map[string]func() (re ct.ResponseEvent, err error){}

		if fn, ok := nextMap[cu.ToString(frmData["next"], "")]; ok {
			return fn()
		}
	*/
	return evt, err
}

func (cls *ClientService) settingResponseFormEvent(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	/*
		client := evt.Trigger.(*ct.Client)
		_, _, stateData := client.GetStateData()
		setting := cu.ToIM(stateData["setting"], cu.IM{})

		frmValues := cu.ToIM(evt.Value, cu.IM{})
		//frmValue := cu.ToIM(frmValues["value"], cu.IM{})
		frmData := cu.ToIM(frmValues["data"], cu.IM{})
		form := cu.ToIM(frmData["form"], cu.IM{})
		frmIndex := cu.ToInteger(form["index"], 0)
		frmKey := cu.ToString(form["key"], "")
		//frmBaseValues := cu.ToIM(form["data"], cu.IM{})
		frmEvent := cu.ToString(frmValues["event"], "")
		rows := cu.ToIMA(setting[frmKey], []cu.IM{})
		//delete := (cu.ToString(frmValue["form_delete"], "") != "")

		eventMap := map[string]func() (re ct.ResponseEvent, err error){}

		if len(rows) > 0 && frmIndex < int64(len(rows)) {
			if fn, ok := eventMap[frmEvent]; ok {
				return fn()
			}
		}
	*/

	return evt, err

}

func (cls *ClientService) settingResponseEditorField(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, stateKey, stateData := client.GetStateData()
	ds := cls.getDataStore(client.Ticket.Database)
	user := client.Ticket.User
	setting := cu.ToIM(stateData["setting"], cu.IM{})

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
