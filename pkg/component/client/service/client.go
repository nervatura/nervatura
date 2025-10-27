package service

import (
	"errors"
	"fmt"
	"log/slog"
	"slices"
	"strings"
	"time"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	api "github.com/nervatura/nervatura/v6/pkg/api"
	cpu "github.com/nervatura/nervatura/v6/pkg/component"
	cp "github.com/nervatura/nervatura/v6/pkg/component/client/component"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
	st "github.com/nervatura/nervatura/v6/pkg/static"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// ClientService implements the Nervatura Client GUI Service.
type ClientService struct {
	Config       cu.IM
	AuthConfigs  map[string]*oauth2.Config
	AppLog       *slog.Logger
	Session      *api.SessionService
	NewDataStore func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore
	Modules      map[string]func(cls *ClientService) ServiceModule
	UI           *cp.ClientComponent
}

// ServiceModule is an interface that defines the methods for a module in the ClientService.
type ServiceModule interface {
	// Data is the method that returns the data for the module.
	Data(evt ct.ResponseEvent, params cu.IM) (data cu.IM, err error)
	// Response is the method that returns the response for the module.
	Response(evt ct.ResponseEvent) (re ct.ResponseEvent, err error)
}

var moduleMap = map[string]func(cls *ClientService) ServiceModule{
	"search": func(cls *ClientService) ServiceModule {
		return NewSearchService(cls)
	},
	"customer": func(cls *ClientService) ServiceModule {
		return NewCustomerService(cls)
	},
	"employee": func(cls *ClientService) ServiceModule {
		return NewEmployeeService(cls)
	},
	"place": func(cls *ClientService) ServiceModule {
		return NewPlaceService(cls)
	},
	"product": func(cls *ClientService) ServiceModule {
		return NewProductService(cls)
	},
	"project": func(cls *ClientService) ServiceModule {
		return NewProjectService(cls)
	},
	"tool": func(cls *ClientService) ServiceModule {
		return NewToolService(cls)
	},
	"trans": func(cls *ClientService) ServiceModule {
		return NewTransService(cls)
	},
}

func NewClientService(config cu.IM, appLog *slog.Logger) *ClientService {
	cls := &ClientService{
		Config:       config,
		AppLog:       appLog,
		Session:      api.NewSession(config, "", cu.ToString(config["NT_SESSION_ALIAS"], "")),
		NewDataStore: api.NewDataStore,
		Modules:      moduleMap,
		UI:           cp.NewClientComponent(),
	}
	cls.AuthConfigs = cls.getAuthConfig()
	return cls
}

func (cls *ClientService) getDataStore(alias string) *api.DataStore {
	return cls.NewDataStore(cls.Config, alias, cls.AppLog)
}

func (cls *ClientService) getAuthConfig() (authMap map[string]*oauth2.Config) {
	cfMap := map[string]*oauth2.Config{
		"NT_GOOGLE_CLIENT": {
			ClientID:     cu.ToString(cls.Config["NT_GOOGLE_CLIENT_ID"], ""),
			ClientSecret: cu.ToString(cls.Config["NT_GOOGLE_CLIENT_SECRET"], ""),
			RedirectURL:  st.AuthRedirectURL,
			Scopes:       []string{"email"},
			Endpoint:     google.Endpoint,
		},
	}
	authMap = map[string]*oauth2.Config{}
	for _, key := range []string{"NT_GOOGLE_CLIENT", "NT_FACEBOOK_CLIENT", "NT_GITHUB_CLIENT", "NT_MICROSOFT_CLIENT"} {
		if cu.ToString(cls.Config[key+"_ID"], "") != "" && cu.ToString(cls.Config[key+"_SECRET"], "") != "" {
			authMap[key] = cfMap[key]
		}
	}
	return authMap
}

func (cls *ClientService) GetClient(host, sessionID, eventURL, lang, theme string) (client *ct.Client) {
	authButtons := func() (authBtn []ct.LoginAuthButton) {
		btnValue := cu.SM{
			"NT_GOOGLE_CLIENT":    "google,Google,Google",
			"NT_FACEBOOK_CLIENT":  "facebook,Facebook,Facebook",
			"NT_GITHUB_CLIENT":    "github,Github,Github",
			"NT_MICROSOFT_CLIENT": "microsoft,Microsoft,Microsoft",
		}
		authBtn = []ct.LoginAuthButton{}
		for key := range cls.AuthConfigs {
			authBtn = append(authBtn, ct.LoginAuthButton{
				Id:    strings.Split(btnValue[key], ",")[0],
				Label: strings.Split(btnValue[key], ",")[1],
				Icon:  strings.Split(btnValue[key], ",")[2],
			})
		}
		return authBtn
	}
	cli := &ct.Client{
		BaseComponent: ct.BaseComponent{
			Id:           sessionID + "_client",
			EventURL:     eventURL,
			OnResponse:   cls.MainResponse,
			RequestValue: map[string]cu.IM{},
			RequestMap:   map[string]ct.ClientComponent{},
			Data:         cu.IM{},
		},
		Ticket: ct.Ticket{
			Host:       host,
			AuthMethod: "password",
			Database:   cu.ToString(cls.Config["NT_DEFAULT_ALIAS"], ""),
			SessionID:  sessionID,
		},
		LoginURL: st.ClientPath + "/auth/",
		//LoginDisabled:   false,
		LoginButtons: authButtons(),
		Version:      cu.ToString(cls.Config["version"], ""),
		//HideSideBar:     false,
		//HideMenu:        false,
		Lang:            lang,
		Theme:           theme,
		ClientLabels:    cls.UI.Labels,
		ClientMenu:      cls.UI.Menu,
		ClientSideBar:   cls.UI.SideBar,
		ClientLogin:     cls.UI.Login,
		ClientSearch:    cls.UI.Search,
		ClientBrowser:   cls.UI.Browser,
		ClientEditor:    cls.UI.Editor,
		ClientModalForm: cls.UI.ModalForm,
		ClientForm:      cls.UI.Form,
	}
	//cli.SetConfigValue("hide_exit", false)
	return cli
}

func (cls *ClientService) LoadSession(sessionID string) (client *ct.Client, err error) {
	var memClient interface{}
	if memClient, err = cls.Session.LoadSession(sessionID, &client); err == nil {
		if memClient, found := memClient.(*ct.Client); found {
			client = memClient
		} else {
			client.OnResponse = cls.MainResponse

			client.ClientLabels = cls.UI.Labels
			client.ClientMenu = cls.UI.Menu
			client.ClientSideBar = cls.UI.SideBar
			client.ClientLogin = cls.UI.Login
			client.ClientSearch = cls.UI.Search
			client.ClientBrowser = cls.UI.Browser
			client.ClientEditor = cls.UI.Editor
			client.ClientModalForm = cls.UI.ModalForm
			client.ClientForm = cls.UI.Form
			_, err = client.Render()
		}
	}
	return client, err
}

func (cls *ClientService) AuthUser(database, username string) (user cu.IM, err error) {
	ds := cls.getDataStore(database)
	if authUser, err := ds.AuthUser("", username); err == nil {
		user = cu.IM{
			"id": authUser.Id, "code": authUser.Code,
			"user_name": authUser.UserName, "user_group": authUser.UserGroup.String(),
			"tags": authUser.AuthMeta.Tags, "auth_map": authUser.AuthMap,
			"bookmarks": authUser.AuthMeta.Bookmarks, "auth_filter": authUser.AuthMeta.Filter,
		}
	}
	return user, err
}

func (cls *ClientService) userLogin(database, username, password string) (user cu.IM, err error) {
	ds := cls.getDataStore(database)

	if user, err = cls.AuthUser(database, username); err == nil {
		_, err = ds.UserLogin(username, password, false)
	}
	return user, err
}

func (cls *ClientService) codeName(ds *api.DataStore, code, model string) (name string) {
	if code != "" {
		if rows, err := ds.StoreDataQuery(md.Query{
			Fields: []string{"*"}, From: model,
			Filters: []md.Filter{
				{Field: "code", Comp: "==", Value: code},
			},
		}, true); err == nil {
			name = cu.ToString(rows[0][model+"_name"], "")
		}
	}
	return name
}

/*
func (cls *ClientService) moduleData(evt ct.ResponseEvent, ds *api.DataStore, mKey string, user cu.IM, params cu.IM) (cu.IM, error) {
	modelMap := map[string]func() (data cu.IM, err error){
		/*
						"customer": func() (cu.IM, error) {
							//return cls.customerData(ds, user, params)
							return NewCustomerService(cls).Data(evt, params)
						},

					"product": func() (cu.IM, error) {
						return cls.productData(ds, user, params)
					},
				"tool": func() (cu.IM, error) {
					return cls.toolData(ds, user, params)
				},
					"project": func() (cu.IM, error) {
						return cls.projectData(ds, user, params)
					},

						"employee": func() (cu.IM, error) {
							return cls.employeeData(ds, user, params)
						},
			"setting": func() (cu.IM, error) {
				return cls.settingData(ds, user, params)
			},
					"trans": func() (cu.IM, error) {
						//return cls.transData(ds, user, params)
						return NewTransService(cls).Data(evt, params)
					},
				"place": func() (cu.IM, error) {
					return cls.placeData(ds, user, params)
				},

	}
	if sm, found := cls.Modules[mKey]; found {
		return sm(cls).Data(evt, params)
	}
	if rm, found := modelMap[mKey]; found {
		return rm()
	}
	return cu.IM{}, nil
}
*/

/*
func (cls *ClientService) moduleResponse(mKey string, evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	modelMap := map[string]func() (re ct.ResponseEvent, err error){

			"search": func() (re ct.ResponseEvent, err error) {
				return cls.searchResponse(evt)
			},

							"customer": func() (re ct.ResponseEvent, err error) {
								//return cls.customerResponse(evt)
								return NewCustomerService(cls).Response(evt)
							},

						"product": func() (re ct.ResponseEvent, err error) {
							return cls.productResponse(evt)
						},
					"tool": func() (re ct.ResponseEvent, err error) {
						return cls.toolResponse(evt)
					},

						"project": func() (re ct.ResponseEvent, err error) {
							return cls.projectResponse(evt)
						},

							"employee": func() (re ct.ResponseEvent, err error) {
								return cls.employeeResponse(evt)
							},
				"setting": func() (re ct.ResponseEvent, err error) {
					return cls.settingResponse(evt)
				},
						"trans": func() (re ct.ResponseEvent, err error) {
							return NewTransService(cls).Response(evt)
						},
					"place": func() (re ct.ResponseEvent, err error) {
						return cls.placeResponse(evt)
					},

	}
	if sm, found := cls.Modules[mKey]; found {
		return sm(cls).Response(evt)
	}
	if rm, found := modelMap[mKey]; found {
		return rm()
	}
	return evt, nil
}
*/

func (cls *ClientService) evtMsg(name, triggerName, value, toastType string, timeout int64) ct.ResponseEvent {
	return ct.ResponseEvent{
		Trigger: &ct.Toast{
			Type:    toastType,
			Value:   value,
			Timeout: timeout,
		},
		TriggerName: triggerName,
		Name:        name,
		Header: cu.SM{
			ct.HeaderRetarget: "#toast-msg",
			ct.HeaderReswap:   "innerHTML",
		},
	}
}

func (cls *ClientService) editorCodeSelector(evt ct.ResponseEvent, editor, codeType string, editorData cu.IM,
	resultUpdate func(params cu.IM) (re ct.ResponseEvent, err error)) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	//ds := cls.getDataStore(client.Ticket.Database)

	values := cu.ToIM(evt.Value, cu.IM{})
	value := cu.ToString(values["value"], "")
	event := cu.ToString(values["event"], "")
	if fe, found := values["form_event"]; found {
		event = cu.ToString(fe, "")
	}

	if event == ct.SelectorEventSearch {
		sConf := cls.UI.SearchConfig.View(codeType+"_simple", client.Labels())
		filters := client.ToFilters(cu.ToString(value, ""), sConf.Filters)
		var resultData cu.IM
		if resultData, err = cls.Modules["search"](cls).Data(evt, cu.IM{
			"view":    codeType + "_simple",
			"query":   cls.UI.SearchConfig.Query(codeType+"_simple", cu.IM{"editor": editor}),
			"filters": filters,
		}); err != nil {
			return evt, err
		}
		/*
			if stateData[codeType+"_selector"], err = cls.searchData(
				ds, codeType+"_simple", searchQuery[codeType+"_simple"](editor), filters); err != nil {
				return evt, err
			}
		*/
		stateData[codeType+"_selector"] = cu.ToIMA(resultData["result"], []cu.IM{})
		return resultUpdate(cu.IM{"event": event, "dirty": false})
	}
	if slices.Contains([]string{"transitem", "transmovement", "transpayment", "invoice"}, codeType) {
		codeType = "trans"
	}
	if event == ct.SelectorEventSelected {
		valueData := cu.ToIM(values["value"], cu.IM{})
		selectedRow := cu.ToIM(valueData["row"], cu.IM{})
		editorData[codeType+"_code"] = selectedRow["code"]
		if cu.ToString(selectedRow[codeType+"_name"], "") != "" {
			stateData[codeType+"_name"] = selectedRow[codeType+"_name"]
		}
		return resultUpdate(cu.IM{"event": event, "dirty": true, "values": selectedRow})
	}
	if event == ct.SelectorEventDelete {
		editorData[codeType+"_code"] = nil
		return resultUpdate(cu.IM{"event": event, "dirty": true})
	}
	if event == ct.SelectorEventLink {
		selectValue := values["value"].(ct.SelectOption)
		stateData["params"] = cu.IM{
			codeType + "_code": cu.ToString(selectValue.Value, ""),
			"session_id":       client.Ticket.SessionID}
		if cu.ToBoolean(stateData["dirty"], false) {
			modal := cu.IM{
				"warning_label":   client.Msg("inputbox_dirty"),
				"warning_message": client.Msg("inputbox_drop"),
				"next":            codeType,
			}
			client.SetForm("warning", modal, 0, true)
			return evt, nil
		} else {
			return cls.setEditor(evt, codeType, stateData["params"].(cu.IM)), nil
		}
	}
	return evt, err
}

func (cls *ClientService) showReportSelector(evt ct.ResponseEvent, refType, code string) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	userConfig := cu.ToIM(client.Ticket.User["auth_map"], cu.IM{})

	configReport := cu.ToIMA(stateData["config_report"], []cu.IM{})
	modal := cu.IM{
		"title":         code,
		"icon":          ct.IconChartBar,
		"config_data":   stateData["config_data"],
		"config_report": configReport,
		"code":          code,
		"ref_type":      refType,
		"orientation":   cu.ToString(userConfig["orientation"], st.DefaultOrientation),
		"paper_size":    cu.ToString(userConfig["paper_size"], st.DefaultPaperSize),
		"copy":          1,
		"auth_code":     cu.ToString(client.Ticket.User["code"], ""),
		"url_pdf":       fmt.Sprintf(st.ClientPath+"/session/export/report/modal/%s?output=pdf&inline=true", client.Ticket.SessionID),
		"url_export":    fmt.Sprintf(st.ClientPath+"/session/export/report/modal/%s?output=pdf", client.Ticket.SessionID),
		"url_xml":       fmt.Sprintf(st.ClientPath+"/session/export/report/modal/%s?output=xml", client.Ticket.SessionID),
		"next":          "report_queue",
	}
	if len(configReport) > 0 {
		modal["template"] = configReport[0]["report_key"]
	}
	client.SetForm("report", modal, 0, true)
	return evt, err
}

func (cls *ClientService) addMapField(evt ct.ResponseEvent, editorMap cu.IM,
	resultUpdate func(params cu.IM) (re ct.ResponseEvent, err error)) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()

	mapField := cu.ToString(stateData["map_field"], "")
	configMap := cu.ToIMA(stateData["config_map"], []cu.IM{})
	var defaultValue any = ""
	if idx := slices.IndexFunc(configMap, func(c cu.IM) bool {
		return cu.ToString(c["field_name"], "") == mapField
	}); idx > int(-1) {
		fieldType := cu.ToString(configMap[idx]["field_type"], "")
		fieldDesc := cu.ToString(configMap[idx]["description"], "")
		if slices.Contains([]string{
			"FIELD_CUSTOMER", "FIELD_EMPLOYEE", "FIELD_PLACE", "FIELD_PRODUCT", "FIELD_PROJECT", "FIELD_TOOL",
			"FIELD_TRANS_ITEM", "FIELD_TRANS_MOVEMENT", "FIELD_TRANS_PAYMENT"}, fieldType) {
			model := strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(fieldType, "FIELD_", ""), "TRANS_", ""))
			modal := cu.IM{
				"title":         client.Msg("map_new"),
				"icon":          ct.IconUser,
				"label":         fieldDesc,
				"placeholder":   "",
				"field_name":    "value",
				"default_value": "",
				"required":      false,
				"next":          "editor_map_value",
				"model":         model,
				"map_field":     mapField,
			}
			client.SetForm("input_string", modal, 0, true)
			return evt, nil
		}
		defaultValue = cp.DefaultMapValue(fieldType)
		if fieldType == md.FieldTypeEnum.String() {
			defaultValue = cu.ToString(ut.ToStringArray(configMap[idx]["tags"])[0], "")
		}
	}
	editorMap[mapField] = defaultValue
	stateData["map_field"] = ""
	return resultUpdate(cu.IM{"dirty": true})
}

func (cls *ClientService) updateMapField(evt ct.ResponseEvent, editorMap cu.IM,
	resultUpdate func(params cu.IM) (re ct.ResponseEvent, err error)) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	ds := cls.getDataStore(client.Ticket.Database)

	values := cu.ToIM(evt.Value, cu.IM{})
	valueData := cu.ToIM(values["value"], cu.IM{})
	row := cu.ToIM(valueData["row"], cu.IM{})
	fieldName := cu.ToString(row["field_name"], "")
	fieldType := cu.ToString(row["field_type"], "")
	if slices.Contains([]string{
		"FIELD_CUSTOMER", "FIELD_EMPLOYEE", "FIELD_PLACE", "FIELD_PRODUCT", "FIELD_PROJECT", "FIELD_TOOL",
		"FIELD_TRANS_ITEM", "FIELD_TRANS_MOVEMENT", "FIELD_TRANS_PAYMENT"}, fieldType) {
		model := strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(fieldType, "FIELD_", ""), "TRANS_", ""))
		if _, err := ds.GetDataByID(model, 0, cu.ToString(row["value"], ""), true); err != nil {
			errMsg := fmt.Sprintf("%s: %s (%s)", client.Msg("invalid_code"), cu.ToString(row["value"], ""), model)
			return evt, errors.New(errMsg)
		}
	}
	editorMap[fieldName] = row["value"]
	return resultUpdate(cu.IM{"dirty": true})
}

func (cls *ClientService) editorFormOK(evt ct.ResponseEvent, rows []cu.IM, customValues map[string]func(value any)) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()

	frmValues := cu.ToIM(evt.Value, cu.IM{})
	frmValue := cu.ToIM(frmValues["value"], cu.IM{})
	frmData := cu.ToIM(frmValues["data"], cu.IM{})
	form := cu.ToIM(frmData["form"], cu.IM{})
	frmIndex := cu.ToInteger(form["index"], 0)
	frmBaseValues := cu.ToIM(form["data"], cu.IM{})

	for field := range frmValue {
		if fn, ok := customValues["frm_"+field]; ok {
			fn(frmValue[field])
		} else if _, found := rows[frmIndex][field]; found {
			rows[frmIndex][field] = frmValue[field]
		}
	}
	for field := range frmBaseValues {
		if fn, ok := customValues["base_"+field]; ok {
			fn(frmBaseValues[field])
		}
	}
	stateData["dirty"] = true
	return evt, err
}

func (cls *ClientService) editorFormTags(params cu.IM, evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()

	frmValues := cu.ToIM(evt.Value, cu.IM{})
	frmValue := cu.ToIM(frmValues["value"], cu.IM{})
	frmData := cu.ToIM(frmValues["data"], cu.IM{})
	form := cu.ToIM(frmData["form"], cu.IM{})
	frmIndex := cu.ToInteger(form["index"], 0)
	frmKey := cu.ToString(form["key"], "")
	frmBaseValues := cu.ToIM(form["data"], cu.IM{})
	formEvent := cu.ToString(frmValues["form_event"], "")
	rowField := cu.ToString(params["row_field"], "tags")

	if formEvent == ct.ListEventAddItem {
		modal := cu.IM{
			"title":         cu.ToString(params["title"], client.Msg("inputbox_new_tag")),
			"icon":          cu.ToString(params["icon"], ct.IconTag),
			"label":         cu.ToString(params["label"], client.Msg("inputbox_enter_tag")),
			"placeholder":   "",
			"field_name":    "value",
			"default_value": "",
			"required":      false,
			"next":          "form_add_tag",
			"frm_key":       frmKey,
			"frm_index":     frmIndex,
			"row":           frmBaseValues,
			"row_field":     rowField,
			"meta_name":     params["meta_name"],
			"options":       params["options"],
			"value":         params["value"],
			"is_null":       params["is_null"],
		}
		client.SetForm(cu.ToString(params["form_key"], "input_string"), modal, 0, true)
		return evt, nil
	}
	if formEvent == ct.ListEventDelete {
		metaName := ut.MetaName(frmBaseValues, cu.ToString(params["meta_name"], "_meta"))
		tags := ut.ToStringArray(frmBaseValues[rowField])
		if metaName != "" {
			tags = ut.ToStringArray(cu.ToIM(frmBaseValues[metaName], cu.IM{})[rowField])
		}
		row := cu.ToIM(frmValue["row"], cu.IM{})
		if idx := slices.Index(tags, cu.ToString(row["tag"], "")); idx > int(-1) {
			tags = append(tags[:idx], tags[idx+1:]...)
			if metaName != "" {
				cu.ToIM(frmBaseValues[metaName], cu.IM{})[rowField] = tags
			} else {
				frmBaseValues[rowField] = tags
			}
		}
		stateData["dirty"] = true
		client.SetForm(cu.ToString(stateData["view"], ""), frmBaseValues, frmIndex, false)
		return evt, nil
	}
	return evt, nil
}

func (cls *ClientService) editorTags(evt ct.ResponseEvent, editorMeta cu.IM,
	resultUpdate func(params cu.IM) (re ct.ResponseEvent, err error)) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	values := cu.ToIM(evt.Value, cu.IM{})
	event := cu.ToString(values["event"], "")

	if event == ct.ListEventAddItem {
		modal := cu.IM{
			"title":         client.Msg("inputbox_new_tag"),
			"icon":          ct.IconTag,
			"label":         client.Msg("inputbox_enter_tag"),
			"placeholder":   "",
			"field_name":    "value",
			"default_value": "",
			"required":      false,
			"next":          "editor_add_tag",
		}
		client.SetForm("input_string", modal, 0, true)
		return evt, nil
	}
	if event == ct.ListEventDelete {
		tags := ut.ToStringArray(editorMeta["tags"])
		valueData := cu.ToIM(values["value"], cu.IM{})
		row := cu.ToIM(valueData["row"], cu.IM{})
		if idx := slices.Index(tags, cu.ToString(row["tag"], "")); idx > int(-1) {
			tags = append(tags[:idx], tags[idx+1:]...)
			editorMeta["tags"] = tags
		}
	}
	return resultUpdate(cu.IM{"event": event, "dirty": true})
}

func (cls *ClientService) addBookmark(evt ct.ResponseEvent, bm md.Bookmark) ct.ResponseEvent {
	client := evt.Trigger.(*ct.Client)
	ds := cls.getDataStore(client.Ticket.Database)
	var err error
	var authUser md.Auth
	if authUser, err = ds.AuthUser(cu.ToString(client.Ticket.User["code"], ""), ""); err == nil {
		authUser.AuthMeta.Bookmarks = append(authUser.AuthMeta.Bookmarks, bm)
		var bookmarks []cu.IM
		if err = ut.ConvertToType(authUser.AuthMeta.Bookmarks, &bookmarks); err == nil {
			client.Ticket.User["bookmarks"] = bookmarks
		}
		values := cu.IM{}
		ut.ConvertByteToIMData(authUser.AuthMeta, values, "auth_meta")
		update := md.Update{Values: values, Model: "auth", IDKey: authUser.Id}
		_, err = ds.StoreDataUpdate(update)
	}
	if err != nil {
		return cls.evtMsg(evt.Name, evt.TriggerName, err.Error(), ct.ToastTypeError, 0)
	}
	return evt
}

func (cls *ClientService) deleteBookmark(evt ct.ResponseEvent) ct.ResponseEvent {
	client := evt.Trigger.(*ct.Client)
	ds := cls.getDataStore(client.Ticket.Database)
	bookmarks := cu.ToIMA(client.Ticket.User["bookmarks"], []cu.IM{})

	data := cu.ToIM(evt.Value, cu.IM{})
	value := cu.ToIM(data["value"], cu.IM{})
	idx := cu.ToInteger(value["index"], 0)

	var err error
	var authUser md.Auth
	if authUser, err = ds.AuthUser(cu.ToString(client.Ticket.User["code"], ""), ""); err == nil {
		if idx < int64(len(authUser.AuthMeta.Bookmarks)) {
			authUser.AuthMeta.Bookmarks = append(authUser.AuthMeta.Bookmarks[:idx], authUser.AuthMeta.Bookmarks[idx+1:]...)
			bookmarks = append(bookmarks[:idx], bookmarks[idx+1:]...)
		}
		client.Ticket.User["bookmarks"] = bookmarks
		values := cu.IM{}
		ut.ConvertByteToIMData(authUser.AuthMeta, values, "auth_meta")
		update := md.Update{Values: values, Model: "auth", IDKey: authUser.Id}
		_, err = ds.StoreDataUpdate(update)
	}
	if err != nil {
		return cls.evtMsg(evt.Name, evt.TriggerName, err.Error(), ct.ToastTypeError, 0)
	}
	return cls.mainResponseBookmark(evt)
}

func (cls *ClientService) setBookmark(evt ct.ResponseEvent) ct.ResponseEvent {
	client := evt.Trigger.(*ct.Client)
	bookmarks := cu.ToIMA(client.Ticket.User["bookmarks"], []cu.IM{})

	evtData := cu.ToIM(evt.Value, cu.IM{})
	value := cu.ToIM(evtData["value"], cu.IM{})
	valueRow := cu.ToIM(value["row"], cu.IM{})
	idx := cu.ToInteger(valueRow["index"], cu.ToInteger(value["index"], 0))
	bookmark := bookmarks[idx]
	client.CloseModal()

	if cu.ToString(bookmark["bookmark_type"], "") == md.BookmarkTypeEditor.String() {
		return cls.setEditor(evt, cu.ToString(bookmark["key"], ""),
			cu.IM{
				"editor_view": cu.ToString(bookmark["key"], ""),
				cu.ToString(bookmark["key"], "") + "_code": cu.ToString(bookmark["code"], ""),
				"session_id": client.Ticket.SessionID,
			})
	}
	userConfig := cu.ToIM(client.Ticket.User["auth_map"], cu.IM{})
	client.SetSearch(cu.ToString(bookmark["key"], ""), cu.IM{
		"session_id":  client.Ticket.SessionID,
		"user_config": userConfig,
		"auth_filter": client.Ticket.User["auth_filter"],
		"user_group":  client.Ticket.User["user_group"],
		cu.ToString(bookmark["key"], ""): cu.IM{
			"filters":         bookmark["filters"],
			"visible_columns": bookmark["columns"],
		},
	}, false)
	return evt
}

func (cls *ClientService) setEditor(evt ct.ResponseEvent, module string, params cu.IM) ct.ResponseEvent {
	client := evt.Trigger.(*ct.Client)
	var moData cu.IM = cu.IM{}
	var err error
	if sm, found := cls.Modules[module]; found {
		if moData, err = sm(cls).Data(evt, params); err != nil {
			return cls.evtMsg(evt.Name, evt.TriggerName, err.Error(), ct.ToastTypeError, 0)
		}
	}
	client.SetEditor(module, cu.ToString(params["editor_view"], module), moData)
	return evt
}

func (cls *ClientService) searchEvent(evt ct.ResponseEvent) (re ct.ResponseEvent) {
	var err error
	client := evt.Trigger.(*ct.Client)
	_, stateKey, stateData := client.GetStateData()
	//ds := cls.getDataStore(client.Ticket.Database)
	//var result []cu.IM
	var filter string = cu.ToString(evt.Value, "")
	sConf := cls.UI.SearchConfig.View(stateKey, client.Labels())
	var resultData cu.IM
	if resultData, err = cls.Modules["search"](cls).Data(evt, cu.IM{
		"view":    stateKey,
		"query":   cls.UI.SearchConfig.Query(stateKey, cu.IM{"editor": ""}),
		"filters": client.GetSearchFilters(filter, sConf.Filters),
	}); err != nil {
		return cls.evtMsg(evt.Name, evt.TriggerName, err.Error(), ct.ToastTypeError, 0)
	}
	stateData["rows"] = cu.ToIMA(resultData["result"], []cu.IM{})
	/*
		if result, err = cls.searchData(ds, stateKey, searchQuery[stateKey](""), client.GetSearchFilters(filter, sConf.Filters)); err != nil {
			return cls.evtMsg(evt.Name, evt.TriggerName, err.Error(), ct.ToastTypeError, 0)
		}
	*/
	//stateData["rows"] = result
	stateData["filter_value"] = filter
	client.SetSearch(stateKey, stateData, cu.ToBoolean(stateData["simple"], false))
	return evt
}

func (cls *ClientService) insertPrintQueue(ds *api.DataStore, data cu.IM) (err error) {
	values := cu.IM{
		"config_type": md.ConfigTypePrintQueue.String(),
	}
	if configDataByte, err := ds.ConvertToByte(md.ConfigPrintQueue{
		RefType:     cu.ToString(data["ref_type"], ""),
		RefCode:     cu.ToString(data["code"], ""),
		Qty:         cu.ToInteger(data["copy"], 0),
		ReportCode:  cu.ToString(data["template"], ""),
		Orientation: cu.ToString(data["orientation"], ""),
		PaperSize:   cu.ToString(data["paper_size"], ""),
		AuthCode:    cu.ToString(data["auth_code"], ""),
	}); err == nil {
		values["data"] = string(configDataByte[:])
	}
	_, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "config"})
	return err
}

func (cls *ClientService) mainResponseLogin(evt ct.ResponseEvent) (re ct.ResponseEvent) {
	var err error
	client := evt.Trigger.(*ct.Client)
	evtData := cu.ToIM(evt.Value, cu.IM{})
	username := cu.ToString(evtData["username"], "")
	password := cu.ToString(evtData["password"], "")
	database := cu.ToString(evtData["database"], cu.ToString(cls.Config["NT_DEFAULT_ALIAS"], ""))
	var user cu.IM
	if user, err = cls.userLogin(database, username, password); err != nil {
		return cls.evtMsg(evt.Name, evt.TriggerName, err.Error(), ct.ToastTypeError, 0)
	}
	client.Ticket.Database = database
	client.Ticket.User = user
	client.Ticket.Expiry = time.Now().Add(time.Duration(cu.ToFloat(cls.Config["NT_SESSION_EXP"], 1)) * time.Hour)
	userConfig := cu.ToIM(user["auth_map"], cu.IM{})
	client.Lang = cu.ToString(userConfig["lang"], st.DefaultLang)
	client.Theme = cu.ToString(userConfig["theme"], st.DefaultTheme)
	client.SetSearch(st.DefaultSearchView, cu.IM{
		"user_config": userConfig,
		"auth_filter": user["auth_filter"],
		"user_group":  user["user_group"],
	}, true)

	url := fmt.Sprintf(st.ClientPath+"/session/%s", client.Ticket.SessionID)
	return cpu.EvtRedirect(ct.LoginEventLogin, evt.Name, url)
}

func (cls *ClientService) mainResponseAuth(evt ct.ResponseEvent) (re ct.ResponseEvent) {
	client := evt.Trigger.(*ct.Client)
	authConfig := cu.ToString(evt.Value, "")
	if config, found := cls.AuthConfigs[cu.ToString(evt.Value, "")]; found {
		config.RedirectURL = fmt.Sprintf(config.RedirectURL, client.Ticket.Host)
		verifier := oauth2.GenerateVerifier()
		url := config.AuthCodeURL(client.Ticket.SessionID, oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(verifier))
		client.Data["verifier"] = verifier
		client.Data["auth_config"] = authConfig
		return cpu.EvtRedirect(ct.LoginEventAuth, evt.Name, url)
	}
	return cls.evtMsg(evt.Name, evt.TriggerName, "Invalid oAuth provider", ct.ToastTypeError, 0)
}

func (cls *ClientService) mainResponseLink(evt ct.ResponseEvent, evtData cu.IM) (re ct.ResponseEvent) {
	client := evt.Trigger.(*ct.Client)
	row := cu.ToIM(evtData["row"], cu.IM{})
	fieldName := cu.ToString(evtData["fieldname"], "")
	fieldType := cu.ToString(row["field_type"], "")
	module := strings.Split(strings.TrimPrefix(fieldName, "ref_"), "_")[0]
	rowId := cu.ToString(row[module+"_id"], "")
	if fieldName == "value" && fieldType == "FIELD_URL" {
		return cpu.EvtRedirect(evt.Name, evt.TriggerName, cu.ToString(row["value"], ""))
	}
	if fieldName == "value" && slices.Contains([]string{
		"FIELD_CUSTOMER", "FIELD_EMPLOYEE", "FIELD_PLACE", "FIELD_PRODUCT", "FIELD_PROJECT",
		"FIELD_TOOL", "FIELD_TRANS_ITEM", "FIELD_TRANS_MOVEMENT", "FIELD_TRANS_PAYMENT"}, fieldType) {
		module = strings.ToLower(strings.TrimPrefix(fieldType, "FIELD_"))
		return cls.setEditor(evt, module, cu.IM{
			"editor_view":    module,
			module + "_code": cu.ToString(row["value"], ""),
			"session_id":     client.Ticket.SessionID,
		})
	}
	if strings.HasSuffix(fieldName, "_code") {
		return cls.setEditor(evt, module, cu.IM{
			"editor_view":                         module,
			strings.TrimPrefix(fieldName, "ref_"): cu.ToString(row[fieldName], ""),
			"session_id":                          client.Ticket.SessionID,
		})
	}
	if fieldName == "code" {
		module = cu.ToString(row["editor"], "")
		rowId = cu.ToString(row["id"], "")
	}
	return cls.setEditor(evt, module, cu.IM{
		"editor_view":  cu.ToString(row["editor_view"], ""),
		module + "_id": rowId,
		"session_id":   client.Ticket.SessionID,
	})
}

func (cls *ClientService) mainResponseBookmark(evt ct.ResponseEvent) (re ct.ResponseEvent) {
	client := evt.Trigger.(*ct.Client)
	var listRows []cu.IM = []cu.IM{}
	bookmarks := cu.ToIMA(client.Ticket.User["bookmarks"], []cu.IM{})
	for idx, bookmark := range bookmarks {
		listRows = append(listRows, cu.IM{
			"lslabel": cu.ToString(bookmark["label"], ""),
			"lsvalue": cu.ToString(bookmark["code"], "") +
				" - " + client.Msg(cu.ToString(bookmark["bookmark_type"], "")),
			"index": idx,
		})
	}
	modal := cu.IM{
		"title":       client.Msg("bookmark_title"),
		"icon":        ct.IconSearch,
		"edit_item":   true,
		"edit_icon":   ct.IconStar,
		"list_filter": true,
		"delete_item": true,
		"rows":        listRows,
		"next":        "bookmark",
	}
	client.SetForm("selector", modal, 0, true)
	return evt
}

func (cls *ClientService) mainResponseModule(evt ct.ResponseEvent) (re ct.ResponseEvent) {
	client := evt.Trigger.(*ct.Client)
	state, stateKey, stateData := client.GetStateData()

	value := cu.ToString(evt.Value, "")
	mdKey := stateKey
	if state == "search" || state == "browser" {
		mdKey = "search"
	}
	if value != mdKey {
		if cu.ToBoolean(stateData["dirty"], false) {
			modal := cu.IM{
				"name":            "warning",
				"warning_label":   client.Msg("inputbox_dirty"),
				"warning_message": client.Msg("inputbox_drop"),
				"next":            "set_" + value,
			}
			client.SetForm("warning", modal, 0, true)
			return evt
		}
		switch value {
		case "bookmark":
			return cls.mainResponseBookmark(evt)
		case "search":
			client.ResetEditor()
			return evt
		default:
			return cls.setEditor(evt, value,
				cu.IM{
					"editor_view": value,
					"session_id":  client.Ticket.SessionID,
				})
		}
	}
	return evt
}

func (cls *ClientService) mainResponseModuleEvent(evt ct.ResponseEvent, nextKey string) (re ct.ResponseEvent) {
	var err error
	client := evt.Trigger.(*ct.Client)
	state, stateKey, stateData := client.GetStateData()

	moduleKey := cu.ToString(nextKey, stateKey)
	if values, ok := evt.Value.(cu.IM); ok && cu.ToString(values["name"], "") == "bookmark" {
		if cu.ToString(values["event"], "") == "list_delete" {
			return cls.deleteBookmark(evt)
		}
		if cu.ToString(values["event"], "") == "list_filter_change" {
			return evt
		}
		return cls.setBookmark(evt)
	}
	if (state == "search" || state == "browser") && nextKey == "" {
		moduleKey = "search"
	}
	if state == "form" {
		moduleKey = cu.ToString(stateData["key"], "")
	}
	if sm, found := cls.Modules[moduleKey]; found {
		if evt, err = sm(cls).Response(evt); err != nil {
			return cls.evtMsg(evt.Name, evt.TriggerName, err.Error(), ct.ToastTypeError, 0)
		}
	}
	/*
		if evt, err = cls.moduleResponse(moduleKey, evt); err != nil {
			return cls.evtMsg(evt.Name, evt.TriggerName, err.Error(), ct.ToastTypeError, 0)
		}
	*/
	return evt
}

func (cls *ClientService) MainResponse(evt ct.ResponseEvent) (re ct.ResponseEvent) {
	client := evt.Trigger.(*ct.Client)
	_, stateKey, _ := client.GetStateData()

	/*
		moduleEvent := func(nextKey string) ct.ResponseEvent {
			moduleKey := cu.ToString(nextKey, stateKey)
			if values, ok := evt.Value.(cu.IM); ok && cu.ToString(values["name"], "") == "bookmark" {
				if cu.ToString(values["event"], "") == "list_delete" {
					return cls.deleteBookmark(evt)
				}
				if cu.ToString(values["event"], "") == "list_filter_change" {
					return evt
				}
				return cls.setBookmark(evt)
			}
			if (state == "search" || state == "browser") && nextKey == "" {
				moduleKey = "search"
			}
			if state == "form" {
				moduleKey = cu.ToString(stateData["key"], "")
			}
			if evt, err = cls.moduleResponse(moduleKey, evt); err != nil {
				return cls.evtMsg(evt.Name, evt.TriggerName, err.Error(), ct.ToastTypeError, 0)
			}
			return evt
		}
	*/

	reMap := map[string]func() ct.ResponseEvent{
		ct.BrowserEventSearch: func() ct.ResponseEvent {
			return cls.searchEvent(evt)
		},

		ct.SearchEventSearch: func() ct.ResponseEvent {
			return cls.searchEvent(evt)
		},

		ct.TableEventEditCell: func() ct.ResponseEvent {
			return cls.mainResponseLink(evt, cu.ToIM(evt.Value, cu.IM{}))
		},

		ct.TableEventAddItem: func() ct.ResponseEvent {
			if slices.Contains([]string{"transitem", "transmovement", "transpayment"}, stateKey) {
				return NewTransService(cls).createModal(evt,
					cu.IM{"state_key": stateKey, "next": stateKey + "_new"})
			}
			module := stateKey
			params := cu.IM{
				"session_id": client.Ticket.SessionID,
			}
			prmMap := map[string]cu.IM{
				"transmovement_formula": {
					"module":       "trans",
					"session_id":   client.Ticket.SessionID,
					"trans_type":   md.TransTypeFormula.String(),
					"direction":    md.DirectionTransfer.String(),
					"editor_title": client.Msg("trans_formula_new"),
					"editor_icon":  cp.TransTypeIcon("TRANS_FORMULA"),
				},
				"transmovement_waybill": {
					"module":       "trans",
					"session_id":   client.Ticket.SessionID,
					"trans_type":   md.TransTypeWaybill.String(),
					"direction":    md.DirectionOut.String(),
					"editor_title": client.Msg("trans_waybill_new"),
					"editor_icon":  cp.TransTypeIcon("TRANS_WAYBILL"),
				},
			}
			if prm, found := prmMap[module]; found {
				module = cu.ToString(prm["module"], "")
				params = prm
			}
			return cls.setEditor(evt, module, params)
		},

		ct.BrowserEventEditRow: func() ct.ResponseEvent {
			evtData := cu.ToIM(evt.Value, cu.IM{})
			resultMap := map[string]func() ct.ResponseEvent{
				stateKey: func() ct.ResponseEvent {
					return cls.setEditor(evt, cu.ToString(evtData["editor"], ""),
						cu.IM{
							"editor_view": cu.ToString(evtData["editor_view"], ""),
							cu.ToString(evtData["editor"], "") + "_id": cu.ToString(evtData["editor_id"], ""),
							"session_id": client.Ticket.SessionID,
						})
				},
			}
			return resultMap[stateKey]()
		},

		ct.SearchEventSelected: func() ct.ResponseEvent {
			evtData := cu.ToIM(evt.Value, cu.IM{})
			evtRow := cu.ToIM(evtData["row"], cu.IM{})
			resultMap := map[string]func() ct.ResponseEvent{
				stateKey: func() ct.ResponseEvent {
					return cls.setEditor(evt, cu.ToString(evtRow["editor"], ""),
						cu.IM{
							"editor_view": cu.ToString(evtRow["editor_view"], ""),
							cu.ToString(evtRow["editor"], "") + "_id": cu.ToString(evtRow["editor_id"], ""),
							"session_id": client.Ticket.SessionID,
						})
				},
			}
			return resultMap[stateKey]()
		},

		/*
			ct.EditorEventView: func() ct.ResponseEvent {
				//client.SetEditor(stateKey, cu.ToString(evt.Value, ""), stateData)
				//client.SetProperty("editor_view", cu.ToString(evt.Value, ""))
				//client.SetProperty("view", cu.ToString(evt.Value, ""))
				return evt
			},
		*/

		ct.ClientEventModule: func() ct.ResponseEvent {
			return cls.mainResponseModule(evt)
		},

		ct.FormEventOK: func() ct.ResponseEvent {
			frmValues := cu.ToIM(evt.Value, cu.IM{})
			frmData := cu.ToIM(frmValues["data"], cu.IM{})
			switch frmData["next"] {
			case "set_search":
				client.ResetEditor()
				return evt

			case "set_setting":
				return cls.setEditor(evt, "setting", cu.IM{
					"editor_view": "setting",
					"session_id":  client.Ticket.SessionID,
				})

			case "set_bookmark":
				return cls.mainResponseBookmark(evt)

			case "trans_new", "transitem_new", "transpayment_new", "transmovement_new":
				return cls.mainResponseModuleEvent(evt, "trans")

			default:
				return cls.mainResponseModuleEvent(evt, "")
			}
		},

		ct.FormEventChange: func() ct.ResponseEvent {
			values := cu.ToIM(evt.Value, cu.IM{})
			valueData := cu.ToIM(values["data"], cu.IM{})
			return cls.mainResponseModuleEvent(evt, cu.ToString(valueData["module"], ""))
		},

		ct.ClientEventForm: func() ct.ResponseEvent {
			return cls.mainResponseModuleEvent(evt, "")
		},

		ct.ClientEventSideMenu: func() ct.ResponseEvent {
			return cls.mainResponseModuleEvent(evt, "")
		},

		ct.BrowserEventBookmark: func() ct.ResponseEvent {
			return cls.mainResponseModuleEvent(evt, "")
		},

		ct.EditorEventField: func() ct.ResponseEvent {
			evtData := cu.ToIM(evt.Value, cu.IM{})
			fieldName := cu.ToString(evtData["name"], "")
			values := cu.ToIM(evtData["value"], cu.IM{})
			row := cu.ToIM(values["row"], cu.IM{})
			editor := cu.ToString(row["editor"], "")
			if fieldName == ct.TableEventEditCell && editor == "" {
				return cls.mainResponseLink(evt, cu.ToIM(evtData["value"], cu.IM{}))
			}
			return cls.mainResponseModuleEvent(evt, editor)
		},

		ct.LoginEventLogin: func() ct.ResponseEvent {
			return cls.mainResponseLogin(evt)
		},

		ct.LoginEventAuth: func() ct.ResponseEvent {
			return cls.mainResponseAuth(evt)
		},

		ct.ClientEventLogOut: func() ct.ResponseEvent {
			client.Ticket.User = nil
			client.Ticket.Expiry = time.Now()
			return cpu.EvtRedirect(ct.LoginEventAuth, evt.Name, client.LoginURL)
		},
	}

	if rm, found := reMap[evt.Name]; found {
		return rm()
	}
	return evt
}
