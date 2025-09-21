package service

import (
	"slices"
	"strings"

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
		"currency":      []cu.IM{},
		"tax":           []cu.IM{},
		"auth":          []cu.IM{},
		"dirty":         false,
		"user":          user,
		"editor_icon":   ct.IconCog,
		"editor_title":  "",
	}

	var rows []cu.IM = []cu.IM{}
	if rows, err = ds.StoreDataQuery(md.Query{
		Fields: []string{"*"}, From: "config_data", OrderBy: []string{"config_code", "config_key"},
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

	if rows, err = ds.StoreDataQuery(md.Query{
		Fields: []string{"*"}, From: "currency",
	}, false); err != nil {
		return data, err
	}
	data["currency"] = rows

	if rows, err = ds.StoreDataQuery(md.Query{
		Fields: []string{"*"}, From: "tax",
	}, false); err != nil {
		return data, err
	}
	data["tax"] = rows

	if rows, err = ds.StoreDataQuery(md.Query{
		Fields: []string{"*"}, From: "auth",
	}, false); err != nil {
		return data, err
	}
	data["auth"] = rows

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

func (cls *ClientService) currencyUpdate(ds *api.DataStore, data cu.IM) (editor cu.IM, err error) {
	var currencyData md.Currency = md.Currency{
		CurrencyMeta: md.CurrencyMeta{
			Tags: []string{},
		},
		CurrencyMap: cu.IM{},
	}
	ut.ConvertToType(data, &currencyData)
	values := cu.IM{}
	if currencyData.Code != "" {
		values["code"] = currencyData.Code
	}

	ut.ConvertByteToIMData(currencyData.CurrencyMeta, values, "currency_meta")
	ut.ConvertByteToIMData(currencyData.CurrencyMap, values, "currency_map")

	var currencyID int64
	newConfig := (currencyData.Id == 0)
	update := md.Update{Values: values, Model: "currency"}
	if !newConfig {
		update.IDKey = currencyData.Id
	}
	if currencyID, err = ds.StoreDataUpdate(update); err == nil && newConfig {
		var currencies []cu.IM = []cu.IM{}
		if currencies, err = ds.StoreDataGet(cu.IM{"id": currencyID, "model": "currency"}, true); err == nil {
			data = currencies[0]
		}
	}
	return data, err
}

func (cls *ClientService) currencyAdd(evt ct.ResponseEvent, code string) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := cls.getDataStore(client.Ticket.Database)

	errorModal := func(msg string) {
		modal := cu.IM{
			"title":      client.Msg("currency_new"),
			"info_label": msg,
			"icon":       ct.IconExclamationTriangle,
		}
		client.SetForm("info", modal, 0, true)
	}

	code = strings.ToUpper(code)
	if len(code) != 3 {
		errorModal(client.Msg("currency_invalid"))
		return evt, nil
	}

	if _, err = ds.StoreDataGet(cu.IM{"code": code, "model": "currency"}, true); err == nil {
		errorModal(client.Msg("currency_exists"))
		return evt, nil
	}

	var currencyData cu.IM
	if currencyData, err = cls.currencyUpdate(ds, cu.IM{"code": code}); err == nil {
		currencies := cu.ToIMA(stateData["currency"], []cu.IM{})
		currencies = append(currencies, currencyData)
		stateData["currency"] = currencies
	}
	if err != nil {
		errorModal(err.Error())
		return evt, nil
	}
	return evt, nil
}

func (cls *ClientService) taxUpdate(ds *api.DataStore, data cu.IM) (editor cu.IM, err error) {
	var taxData md.Tax = md.Tax{
		TaxMeta: md.TaxMeta{
			Tags: []string{},
		},
		TaxMap: cu.IM{},
	}
	ut.ConvertToType(data, &taxData)
	values := cu.IM{}
	if taxData.Code != "" {
		values["code"] = taxData.Code
	}

	ut.ConvertByteToIMData(taxData.TaxMeta, values, "tax_meta")
	ut.ConvertByteToIMData(taxData.TaxMap, values, "tax_map")

	var taxID int64
	newConfig := (taxData.Id == 0)
	update := md.Update{Values: values, Model: "tax"}
	if !newConfig {
		update.IDKey = taxData.Id
	}
	if taxID, err = ds.StoreDataUpdate(update); err == nil && newConfig {
		var taxes []cu.IM = []cu.IM{}
		if taxes, err = ds.StoreDataGet(cu.IM{"id": taxID, "model": "tax"}, true); err == nil {
			data = taxes[0]
		}
	}
	return data, err
}

func (cls *ClientService) taxAdd(evt ct.ResponseEvent, code string) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := cls.getDataStore(client.Ticket.Database)

	errorModal := func(msg string) {
		modal := cu.IM{
			"title":      client.Msg("tax_new"),
			"info_label": msg,
			"icon":       ct.IconExclamationTriangle,
		}
		client.SetForm("info", modal, 0, true)
	}

	code = strings.ToUpper(code)
	if _, err = ds.StoreDataGet(cu.IM{"code": code, "model": "tax"}, true); err == nil {
		errorModal(client.Msg("tax_exists"))
		return evt, nil
	}

	var taxData cu.IM
	if taxData, err = cls.taxUpdate(ds, cu.IM{"code": code}); err == nil {
		taxes := cu.ToIMA(stateData["tax"], []cu.IM{})
		taxes = append(taxes, taxData)
		stateData["tax"] = taxes
	}
	if err != nil {
		errorModal(err.Error())
		return evt, nil
	}
	return evt, nil
}

func (cls *ClientService) authUpdate(ds *api.DataStore, data cu.IM) (editor cu.IM, err error) {
	var authData md.Auth = md.Auth{
		UserGroup: md.UserGroup(md.UserGroupUser),
		AuthMeta: md.AuthMeta{
			Tags:      []string{},
			Filter:    []md.AuthFilter{},
			Bookmarks: []md.Bookmark{},
		},
		AuthMap: cu.IM{},
	}
	ut.ConvertToType(data, &authData)
	values := cu.IM{}
	if authData.UserName != "admin" {
		values = cu.IM{
			"user_name":  authData.UserName,
			"user_group": authData.UserGroup.String(),
			"disabled":   authData.Disabled,
		}
	}

	ut.ConvertByteToIMData(authData.AuthMeta, values, "auth_meta")
	ut.ConvertByteToIMData(authData.AuthMap, values, "auth_map")

	var authID int64
	newAuth := (authData.Id == 0)
	update := md.Update{Values: values, Model: "auth"}
	if !newAuth {
		update.IDKey = authData.Id
	}
	if authID, err = ds.StoreDataUpdate(update); err == nil && newAuth {
		var auths []cu.IM = []cu.IM{}
		if auths, err = ds.StoreDataGet(cu.IM{"id": authID, "model": "auth"}, true); err == nil {
			data = auths[0]
		}
	}
	return data, err
}

func (cls *ClientService) authAdd(evt ct.ResponseEvent, userName string) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := cls.getDataStore(client.Ticket.Database)

	errorModal := func(msg string) {
		modal := cu.IM{
			"title":      client.Msg("auth_new"),
			"info_label": msg,
			"icon":       ct.IconExclamationTriangle,
		}
		client.SetForm("info", modal, 0, true)
	}

	userName = strings.ReplaceAll(strings.ToLower(userName), " ", "")
	if _, err = ds.StoreDataGet(cu.IM{"user_name": userName, "model": "auth"}, true); err == nil {
		errorModal(client.Msg("auth_exists"))
		return evt, nil
	}

	var authData cu.IM
	if authData, err = cls.authUpdate(ds,
		cu.IM{"user_name": userName, "user_group": md.UserGroupUser.String()}); err == nil {
		auths := cu.ToIMA(stateData["auth"], []cu.IM{})
		auths = append(auths, authData)
		stateData["auth"] = auths
	}
	if err != nil {
		errorModal(err.Error())
		return evt, nil
	}
	return evt, nil
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

func (cls *ClientService) settingResponseFormNextTags(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)

	frmValues := cu.ToIM(evt.Value, cu.IM{})
	frmData := cu.ToIM(frmValues["data"], cu.IM{})
	frmValue := cu.ToIM(frmValues["value"], cu.IM{})

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
}

func (cls *ClientService) settingResponseFormNext(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := cls.getDataStore(client.Ticket.Database)
	//setting := cu.ToIM(stateData["setting"], cu.IM{})
	configValues := cu.ToIMA(stateData["config_values"], []cu.IM{})
	currencies := cu.ToIMA(stateData["currency"], []cu.IM{})
	taxes := cu.ToIMA(stateData["tax"], []cu.IM{})
	auth := cu.ToIMA(stateData["auth"], []cu.IM{})

	frmValues := cu.ToIM(evt.Value, cu.IM{})
	frmData := cu.ToIM(frmValues["data"], cu.IM{})
	frmValue := cu.ToIM(frmValues["value"], cu.IM{})

	nextMap := map[string]func() (re ct.ResponseEvent, err error){
		"form_add_tag": func() (re ct.ResponseEvent, err error) {
			return cls.settingResponseFormNextTags(evt)
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

		"auth_delete": func() (re ct.ResponseEvent, err error) {
			if idx := slices.IndexFunc(auth, func(c cu.IM) bool {
				return cu.ToString(c["code"], "") == cu.ToString(frmData["code"], "")
			}); idx > int(-1) {
				if _, err = ds.StoreDataUpdate(md.Update{Model: "auth", IDKey: cu.ToInteger(auth[idx]["id"], 0)}); err == nil {
					auth = append(auth[:idx], auth[idx+1:]...)
					stateData["auth"] = auth
				}
			}
			return evt, err
		},

		"password_reset": func() (re ct.ResponseEvent, err error) {
			if err = cls.settingPassword(ds,
				cu.IM{"code": cu.ToString(frmData["code"], "")},
				cu.IM{"password": cu.ToString(frmValue["value"], ""), "confirm": cu.ToString(frmValue["value"], "")}); err == nil {
				modal := cu.IM{
					"info_label":   client.Msg("auth_password_reset_ok"),
					"info_message": "",
					"icon":         ct.IconCheck,
				}
				client.SetForm("info", modal, 0, true)
			}
			return evt, err
		},

		"currency_delete": func() (re ct.ResponseEvent, err error) {
			if idx := slices.IndexFunc(currencies, func(c cu.IM) bool {
				return cu.ToString(c["code"], "") == cu.ToString(frmData["code"], "")
			}); idx > int(-1) {
				if _, err = ds.StoreDataUpdate(md.Update{Model: "currency", IDKey: cu.ToInteger(currencies[idx]["id"], 0)}); err == nil {
					currencies = append(currencies[:idx], currencies[idx+1:]...)
					stateData["currency"] = currencies
				}
			}
			if err != nil {
				modal := cu.IM{
					"info_label":   client.Msg("inputbox_delete_error"),
					"info_message": err.Error(),
					"icon":         ct.IconExclamationTriangle,
				}
				client.SetForm("info", modal, 0, true)
			}
			return evt, nil
		},

		"currency_add": func() (re ct.ResponseEvent, err error) {
			return cls.currencyAdd(evt, cu.ToString(frmValue["value"], ""))
		},

		"tax_delete": func() (re ct.ResponseEvent, err error) {
			if idx := slices.IndexFunc(taxes, func(c cu.IM) bool {
				return cu.ToString(c["code"], "") == cu.ToString(frmData["code"], "")
			}); idx > int(-1) {
				if _, err = ds.StoreDataUpdate(md.Update{Model: "tax", IDKey: cu.ToInteger(taxes[idx]["id"], 0)}); err == nil {
					taxes = append(taxes[:idx], taxes[idx+1:]...)
					stateData["tax"] = taxes
				}
			}
			if err != nil {
				modal := cu.IM{
					"info_label":   client.Msg("inputbox_delete_error"),
					"info_message": err.Error(),
					"icon":         ct.IconExclamationTriangle,
				}
				client.SetForm("info", modal, 0, true)
			}
			return evt, nil
		},

		"tax_add": func() (re ct.ResponseEvent, err error) {
			return cls.taxAdd(evt, cu.ToString(frmValue["value"], ""))
		},

		"auth_add": func() (re ct.ResponseEvent, err error) {
			return cls.authAdd(evt, cu.ToString(frmValue["value"], ""))
		},
	}

	if fn, ok := nextMap[cu.ToString(frmData["next"], "")]; ok {
		return fn()
	}
	return evt, err
}

func (cls *ClientService) settingResponseFormEventChange(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()

	frmValues := cu.ToIM(evt.Value, cu.IM{})
	frmData := cu.ToIM(frmValues["data"], cu.IM{})
	form := cu.ToIM(frmData["form"], cu.IM{})
	frmIndex := cu.ToInteger(form["index"], 0)
	configValue := cu.ToIM(form["data"], cu.IM{})
	configMeta := cu.ToIM(configValue["data"], cu.IM{})
	authMeta := cu.ToIM(configValue["auth_meta"], cu.IM{})

	fieldName := cu.ToString(frmValues["name"], "")
	switch cu.ToString(form["key"], "") {
	case "config_map":
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
		}

	case "auth":
		switch fieldName {
		case "tags":
			return cls.editorFormTags(cu.IM{"row_field": fieldName, "meta_name": "auth_meta"}, evt)
		case "filter":
			opt := []ct.SelectOption{}
			ft := md.AuthFilter(0)
			for _, ftKey := range ft.Keys() {
				opt = append(opt, ct.SelectOption{
					Value: ftKey, Text: ftKey,
				})
			}
			return cls.editorFormTags(cu.IM{"row_field": fieldName, "meta_name": "auth_meta",
				"options": opt, "value": opt[0].Value, "is_null": false, "form_key": "select",
				"icon": ct.IconFilter, "title": client.Msg("inputbox_new_filter"),
				"label": client.Msg("inputbox_enter_filter")}, evt)
		case "user_name", "disabled":
			configValue[fieldName] = frmValues["value"]
			client.SetForm(cu.ToString(stateData["view"], ""), configValue, frmIndex, false)
		case "user_group":
			configValue[fieldName] = frmValues["value"]
			if cu.ToString(frmValues["value"], "") == md.UserGroupAdmin.String() {
				authMeta["filter"] = []md.AuthFilter{}
			}
			client.SetForm(cu.ToString(stateData["view"], ""), configValue, frmIndex, false)
		case "description":
			authMeta[fieldName] = frmValues["value"]
			client.SetForm(cu.ToString(stateData["view"], ""), configValue, frmIndex, false)
		case "password":
			modal := cu.IM{
				"title":         client.Msg("auth_password"),
				"icon":          ct.IconKey,
				"label":         client.Msg("auth_password_enter"),
				"placeholder":   "",
				"field_name":    "value",
				"default_value": "",
				"required":      false,
				"next":          "password_reset",
				"code":          cu.ToString(configValue["code"], ""),
			}
			client.SetForm("input_string", modal, 0, true)
		}
	}

	return evt, err
}

func (cls *ClientService) settingResponseFormEvent(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := cls.getDataStore(client.Ticket.Database)
	configValues := cu.ToIMA(stateData["config_values"], []cu.IM{})
	authValues := cu.ToIMA(stateData["auth"], []cu.IM{})

	frmValues := cu.ToIM(evt.Value, cu.IM{})
	frmValue := cu.ToIM(frmValues["value"], cu.IM{})
	frmData := cu.ToIM(frmValues["data"], cu.IM{})
	form := cu.ToIM(frmData["form"], cu.IM{})
	frmEvent := cu.ToString(frmValues["event"], "")
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

			case "auth":
				if idx := slices.IndexFunc(authValues, func(c cu.IM) bool {
					return cu.ToString(c["code"], "") == cu.ToString(configValue["code"], "")
				}); idx > int(-1) {
					if configValue, err = cls.authUpdate(ds, configValue); err == nil {
						authValues[idx] = configValue
						stateData["auth"] = authValues
					}
				}
			}
			return evt, err
		},

		ct.FormEventCancel: func() (re ct.ResponseEvent, err error) {
			switch cu.ToString(form["key"], "") {
			case "config_map":
				if delete {
					configValue := cu.ToIM(form["data"], cu.IM{})
					modal := cu.IM{
						"warning_label":   client.Msg("inputbox_delete"),
						"warning_message": "",
						"next":            "config_delete",
						"code":            cu.ToString(configValue["code"], ""),
					}
					client.SetForm("warning", modal, 0, true)
				}
				if cu.ToString(configValue["code"], "") == "" {
					if idx := slices.IndexFunc(configValues, func(c cu.IM) bool {
						return cu.ToString(c["code"], "") == ""
					}); idx > int(-1) {
						configValues = append(configValues[:idx], configValues[idx+1:]...)
						stateData["config_values"] = configValues
					}
				}
			case "auth":
				if delete {
					authValue := cu.ToIM(form["data"], cu.IM{})
					modal := cu.IM{
						"warning_label":   client.Msg("inputbox_delete"),
						"warning_message": "",
						"next":            "auth_delete",
						"code":            cu.ToString(authValue["code"], ""),
					}
					client.SetForm("warning", modal, 0, true)
				}
			}
			return evt, err
		},

		ct.FormEventChange: func() (re ct.ResponseEvent, err error) {
			return cls.settingResponseFormEventChange(evt)
		},
	}

	if fn, ok := eventMap[frmEvent]; ok {
		return fn()
	}

	return evt, err

}

func (cls *ClientService) settingResponseEditorFieldTable(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := cls.getDataStore(client.Ticket.Database)
	view := cu.ToString(stateData["view"], "")
	configValues := cu.ToIMA(stateData["config_values"], []cu.IM{})
	currencies := cu.ToIMA(stateData["currency"], []cu.IM{})
	taxes := cu.ToIMA(stateData["tax"], []cu.IM{})

	values := cu.ToIM(evt.Value, cu.IM{})
	fieldName := cu.ToString(values["name"], "")

	fieldMap := map[string]func() (re ct.ResponseEvent, err error){
		ct.TableEventRowSelected: func() (re ct.ResponseEvent, err error) {
			valueData := cu.ToIM(values["value"], cu.IM{})
			client.SetForm(cu.ToString(stateData["view"], ""),
				cu.ToIM(valueData["row"], cu.IM{}),
				cu.ToInteger(valueData["index"], 0), false)
			return evt, nil
		},

		ct.TableEventFormUpdate: func() (re ct.ResponseEvent, err error) {
			values := cu.ToIM(evt.Value, cu.IM{})
			valueData := cu.ToIM(values["value"], cu.IM{})
			row := cu.ToIM(valueData["row"], cu.IM{})
			switch view {
			case "config_data":
				configCode := cu.ToString(row["config_code"], "")
				fieldName := cu.ToString(row["config_key"], "")
				if idx := slices.IndexFunc(configValues, func(c cu.IM) bool {
					return cu.ToString(c["code"], "") == configCode
				}); idx > int(-1) {
					configData := cu.ToIM(configValues[idx]["data"], cu.IM{})
					configData[fieldName] = row["config_value"]
					configValues[idx]["data"] = configData
					if configValues[idx], err = cls.configUpdate(ds, configValues[idx]); err == nil {
						stateData["config_values"] = configValues
					}
				}

			case "currency":
				currencyCode := cu.ToString(row["code"], "")
				if idx := slices.IndexFunc(currencies, func(c cu.IM) bool {
					return cu.ToString(c["code"], "") == currencyCode
				}); idx > int(-1) {
					currencyMeta := cu.ToIM(currencies[idx]["currency_meta"], cu.IM{})
					currencyMeta["description"] = row["description"]
					currencyMeta["digit"] = cu.ToInteger(row["digit"], 0)
					currencyMeta["cash_round"] = cu.ToInteger(row["cash_round"], 0)
					currencies[idx]["currency_meta"] = currencyMeta
					if currencies[idx], err = cls.currencyUpdate(ds, currencies[idx]); err == nil {
						stateData["currency"] = currencies
					}
				}

			case "tax":
				taxCode := cu.ToString(row["code"], "")
				if idx := slices.IndexFunc(taxes, func(c cu.IM) bool {
					return cu.ToString(c["code"], "") == taxCode
				}); idx > int(-1) {
					taxMeta := cu.ToIM(taxes[idx]["tax_meta"], cu.IM{})
					taxMeta["description"] = row["description"]
					taxMeta["rate_value"] = cu.ToFloat(row["rate_value"], 0)
					taxes[idx]["tax_meta"] = taxMeta
					if taxes[idx], err = cls.taxUpdate(ds, taxes[idx]); err == nil {
						stateData["tax"] = taxes
					}
				}
			}
			return evt, err
		},

		ct.TableEventFormDelete: func() (re ct.ResponseEvent, err error) {
			values := cu.ToIM(evt.Value, cu.IM{})
			valueData := cu.ToIM(values["value"], cu.IM{})
			row := cu.ToIM(valueData["row"], cu.IM{})
			switch view {
			case "currency":
				modal := cu.IM{
					"warning_label":   client.Msg("inputbox_delete"),
					"warning_message": "",
					"next":            "currency_delete",
					"code":            cu.ToString(row["code"], ""),
				}
				client.SetForm("warning", modal, 0, true)

			case "tax":
				modal := cu.IM{
					"warning_label":   client.Msg("inputbox_delete"),
					"warning_message": "",
					"next":            "tax_delete",
					"code":            cu.ToString(row["code"], ""),
				}
				client.SetForm("warning", modal, 0, true)
			}
			return evt, err
		},

		ct.TableEventAddItem: func() (re ct.ResponseEvent, err error) {
			switch view {
			case "currency":
				modal := cu.IM{
					"title":         client.Msg("currency_new"),
					"icon":          ct.IconDollar,
					"label":         client.Msg("currency_enter"),
					"placeholder":   "",
					"field_name":    "value",
					"default_value": "",
					"required":      false,
					"next":          "currency_add",
				}
				client.SetForm("input_string", modal, 0, true)

			case "tax":
				modal := cu.IM{
					"title":         client.Msg("tax_new"),
					"icon":          ct.IconTicket,
					"label":         client.Msg("tax_enter"),
					"placeholder":   "",
					"field_name":    "value",
					"default_value": "",
					"required":      false,
					"next":          "tax_add",
				}
				client.SetForm("input_string", modal, 0, true)
			}
			return evt, err
		},
	}

	return fieldMap[fieldName]()
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

		"auth": func() (re ct.ResponseEvent, err error) {
			event := cu.ToString(cu.ToIM(evt.Value, cu.IM{})["event"], "")
			evValue := cu.ToIM(cu.ToIM(evt.Value, cu.IM{})["value"], cu.IM{})
			row := cu.ToIM(evValue["row"], cu.IM{})
			index := cu.ToInteger(evValue["index"], 0)
			switch event {
			case ct.ListEventEditItem:
				client.SetForm("auth", row, index, false)

			case ct.ListEventAddItem:
				modal := cu.IM{
					"title":         client.Msg("auth_new"),
					"icon":          ct.IconUserLock,
					"label":         client.Msg("auth_enter"),
					"placeholder":   "",
					"field_name":    "value",
					"default_value": "",
					"required":      false,
					"next":          "auth_add",
				}
				client.SetForm("input_string", modal, 0, true)
			}
			return evt, err
		},

		ct.TableEventRowSelected: func() (re ct.ResponseEvent, err error) {
			return cls.settingResponseEditorFieldTable(evt)
		},

		ct.TableEventFormUpdate: func() (re ct.ResponseEvent, err error) {
			return cls.settingResponseEditorFieldTable(evt)
		},

		ct.TableEventFormDelete: func() (re ct.ResponseEvent, err error) {
			return cls.settingResponseEditorFieldTable(evt)
		},

		ct.TableEventAddItem: func() (re ct.ResponseEvent, err error) {
			return cls.settingResponseEditorFieldTable(evt)
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
