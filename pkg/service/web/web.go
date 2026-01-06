package web

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/csv"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	cpu "github.com/nervatura/nervatura/v6/pkg/client/web"
	cls "github.com/nervatura/nervatura/v6/pkg/client/web/service"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
	st "github.com/nervatura/nervatura/v6/pkg/static"
	"golang.org/x/oauth2"
)

func ClientAuth(w http.ResponseWriter, r *http.Request) {
	var err error
	var html template.HTML

	cs := r.Context().Value(md.ClientServiceCtxKey).(*cls.ClientService)
	lang := cu.ToString(r.URL.Query().Get("lang"), st.DefaultLang)
	theme := cu.ToString(r.URL.Query().Get("theme"), st.DefaultTheme)

	sessionID := ut.GetSessionID()
	client := cs.GetClient(r.Host, sessionID, st.ClientPath+"/auth/event", lang, theme)
	ccApp := cpu.Application(sessionID, client)

	if html, err = ccApp.Render(); err == nil {
		cs.Session.SaveSession(client.Ticket.SessionID, client)
	}
	cpu.RespondMessage(w, html, err)
}

func ClientAuthEvent(w http.ResponseWriter, r *http.Request) {
	var te ct.TriggerEvent
	var err error
	var html template.HTML
	var evt ct.ResponseEvent
	var client *ct.Client

	cs := r.Context().Value(md.ClientServiceCtxKey).(*cls.ClientService)
	sessionID := r.Header.Get("X-Session-Token")
	te, err = cs.TriggerEvent(r)

	if err == nil {
		if client, err = cs.LoadSession(sessionID); err == nil {
			evt = client.OnRequest(te)
		}

		for key, value := range evt.Header {
			w.Header().Set(key, value)
		}
		if evt.Trigger != nil {
			html, err = evt.Trigger.Render()
		} else {
			err = errors.New(http.StatusText(http.StatusInternalServerError))
		}
	}

	if err != nil {
		w.Header().Set(ct.HeaderRetarget, "#toast-msg")
		w.Header().Set(ct.HeaderReswap, "innerHTML")
		html, _ = (&ct.Toast{
			Type: ct.ToastTypeError, Value: err.Error(),
		}).Render()
	} else {
		cs.Session.SaveSession(sessionID, client)
	}
	cpu.RespondMessage(w, html, nil)
}

func ClientSession(w http.ResponseWriter, r *http.Request) {
	var err error
	var html template.HTML
	var loginClient *ct.Client

	cs := r.Context().Value(md.ClientServiceCtxKey).(*cls.ClientService)
	loginID := r.PathValue("session_id")
	if loginClient, err = cs.LoadSession(loginID); err != nil || !loginClient.Ticket.Valid() {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	cs.Session.DeleteSession(loginID)

	loginClient.Ticket.SessionID = ut.GetSessionID()
	client := cs.GetClient(
		r.Host,
		loginClient.Ticket.SessionID, st.ClientPath+"/session/event",
		cu.ToString(loginClient.Lang, st.DefaultLang),
		cu.ToString(loginClient.Theme, st.DefaultTheme))
	client.Ticket = loginClient.Ticket
	client.Data = loginClient.Data

	if _, found := loginClient.Data["request_data"]; found {
		client.SetSearch(st.DefaultSearchView, cu.IM{}, true)
		client.LoginDisabled = true
		//client.HideMenuExit = true
		//if editorData, err := cls.requestUpdate(loginClient.Module, loginClient.Data); err == nil && editorData != nil {
		//	client.Data["editor"] = editorData
		//client.SetSearch("view",cu.IM{},false)
		//client.SetEditor("key", "view", cu.IM{})
		//}
	}

	ccApp := cpu.Application(client.Ticket.SessionID, client)

	if html, err = ccApp.Render(); err == nil {
		cs.Session.SaveSession(client.Ticket.SessionID, client)
	}
	cpu.RespondMessage(w, html, err)
}

func ClientSessionEvent(w http.ResponseWriter, r *http.Request) {
	var te ct.TriggerEvent
	var err error
	var html template.HTML
	var evt ct.ResponseEvent
	var client *ct.Client

	cs := r.Context().Value(md.ClientServiceCtxKey).(*cls.ClientService)
	sessionID := r.Header.Get("X-Session-Token")
	te, err = cs.TriggerEvent(r)

	if err == nil {
		if client, err = cs.LoadSession(sessionID); err == nil && client.Ticket.Valid() {
			evt = client.OnRequest(te)
		} else {
			evt = cs.EvtRedirect(ct.LoginEventAuth, evt.Name, "/")
		}

		for key, value := range evt.Header {
			w.Header().Set(key, value)
		}

		err = errors.New(http.StatusText(http.StatusInternalServerError))
		if evt.Trigger != nil {
			html, err = evt.Trigger.Render()
		}
	}

	if err != nil {
		w.Header().Set(ct.HeaderRetarget, "#toast-msg")
		w.Header().Set(ct.HeaderReswap, "innerHTML")
		html, _ = (&ct.Toast{
			Type: ct.ToastTypeError, Value: err.Error(),
		}).Render()
	} else {
		cs.Session.SaveSession(sessionID, client)
	}
	cpu.RespondMessage(w, html, nil)
}

/*
	curl -X 'POST' 'http://localhost:5000/client/api/session' \
		-H 'X-API-KEY: TEST_API_KEY' \
		-d '{"database": "demo", "username": "admin", "request_id": "RQTEST0001", "module": "search", "lang": "en", "theme": "light", "data": { "code": "code", "name": "name" }}'
*/
func ClientSessionCreate(w http.ResponseWriter, r *http.Request) {
	var err error
	var sessionReq cpu.SessionRequest

	err = cu.ConvertFromReader(r.Body, &sessionReq)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var user cu.IM
	cs := r.Context().Value(md.ClientServiceCtxKey).(*cls.ClientService)
	database := cu.ToString(sessionReq.Database, cu.ToString(cs.Config["NT_DEFAULT_ALIAS"], ""))
	if user, err = cs.AuthUser(database, sessionReq.Username); err != nil {
		errMsg := fmt.Sprintf("%s: %s", ut.GetMessage("en", "unknown_user"), sessionReq.Username)
		http.Error(w, errMsg, http.StatusUnauthorized)
		return
	}

	sessionID := ut.GetSessionID()
	userConfig := cu.ToIM(user["auth_map"], cu.IM{})
	lang := cu.ToString(cu.ToString(userConfig["lang"], sessionReq.Lang), st.DefaultLang)
	theme := cu.ToString(cu.ToString(userConfig["theme"], sessionReq.Theme), st.DefaultTheme)
	url := fmt.Sprintf(st.ClientPath+"/session/%s", sessionID)

	client := cs.GetClient(r.Host, sessionID, st.ClientPath+"/auth/event", lang, theme)
	client.Ticket = ct.Ticket{
		SessionID:  sessionID,
		User:       user,
		Expiry:     time.Now().Add(time.Duration(cu.ToFloat(cs.Config["NT_SESSION_EXP"], 1)) * time.Hour),
		AuthMethod: "provider",
		Database:   database,
		Host:       r.Host,
	}

	switch sessionReq.Module {
	default:
		client.Data["request_state"] = "search"
		client.Data["request_module"] = st.DefaultSearchView
	}

	client.Data["request_user"] = user
	client.Data["request_data"] = sessionReq.Data
	client.Data["request_id"] = sessionReq.RequestID

	cs.Session.SaveSession(sessionID, client)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(url))
}

// OAuth callback handler handle the redirect request from the OAuth provider.
// It read the code query parameter and exchange it to get the access token.
func ClientAuthCallback(w http.ResponseWriter, r *http.Request) {
	var err error
	var token *oauth2.Token
	var appLogin *ct.Client

	code := r.URL.Query().Get("code")
	loginID := r.URL.Query().Get("state")
	cs := r.Context().Value(md.ClientServiceCtxKey).(*cls.ClientService)

	errMsg := ut.GetMessage("en", "failed_authentication")
	errTitle := ut.GetMessage("en", "error_authentication")

	if appLogin, err = cs.LoadSession(loginID); err != nil {
		cpu.ErrorPage(w, errTitle, errMsg)
		return
	}

	verifier := cu.ToString(appLogin.Data["verifier"], "")
	token, err = cs.AuthConfig.Exchange(context.Background(), code, oauth2.VerifierOption(verifier))
	if err != nil {
		cpu.ErrorPage(w, errTitle, errMsg)
		return
	}

	var email string
	idToken := cu.ToString(token.Extra("id_token"), token.AccessToken)
	if idToken != "" && len(strings.Split(idToken, ".")) > 1 {
		var uDec []byte
		if uDec, err = base64.StdEncoding.WithPadding(-1).DecodeString(strings.Split(idToken, ".")[1]); err == nil {
			var data cu.IM
			if err = cu.ConvertFromByte(uDec, &data); err == nil {
				email = cu.ToString(data["email"], cu.ToString(data["preferred_username"], ""))
			}
		}
	}
	if err != nil || email == "" {
		cpu.ErrorPage(w, errTitle, errMsg)
		return
	}

	var user cu.IM
	if user, err = cs.AuthUser(appLogin.Ticket.Database, email); err != nil {
		errMsg = fmt.Sprintf("%s: %s", ut.GetMessage("en", "unknown_user"), email)
		cpu.ErrorPage(w, errTitle, errMsg)
		return
	}

	appLogin.Ticket.AuthMethod = "oauth_login"
	appLogin.Ticket.User = user
	appLogin.Ticket.Expiry = time.Now().Add(time.Duration(cu.ToFloat(cs.Config["NT_SESSION_EXP"], 1)) * time.Hour)
	userConfig := cu.ToIM(user["auth_map"], cu.IM{})
	appLogin.Lang = cu.ToString(userConfig["lang"], st.DefaultLang)
	appLogin.Theme = cu.ToString(userConfig["theme"], st.DefaultTheme)
	appLogin.SetSearch(st.DefaultSearchView, cu.IM{
		"user_config": userConfig,
		"auth_filter": user["auth_filter"],
		"user_group":  user["user_group"],
	}, true)

	cs.Session.SaveSession(loginID, appLogin)

	url := fmt.Sprintf(st.ClientPath+"/session/%s", loginID)
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func ClientExportBrowser(w http.ResponseWriter, r *http.Request) {
	sessionID := r.PathValue("session_id")
	cs := r.Context().Value(md.ClientServiceCtxKey).(*cls.ClientService)

	var client *ct.Client
	var err error
	if client, err = cs.LoadSession(sessionID); err != nil ||
		!client.Ticket.Valid() && !client.LoginDisabled {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	_, stateKey, stateData := client.GetStateData()

	labels := client.CustomFunctions.Labels(client.Lang)
	sConf := cs.UI.SearchConfig.View(stateKey, labels, client.Ticket.SessionID)
	browserFields := sConf.Fields
	visibleColumns := client.GetSearchVisibleColumns(ut.ToBoolMap(sConf.VisibleColumns, map[string]bool{}))
	fileName := fmt.Sprintf("%s.csv", stateKey)

	visibleField := func(fieldname string) bool {
		if visible, found := visibleColumns[fieldname]; found {
			return visible
		}
		return false
	}

	var sRows [][]string = make([][]string, 0)

	//labels
	sRow := make([]string, 0)
	for _, field := range browserFields {
		if visibleField(field.Name) {
			sRow = append(sRow, field.Label)
		}
	}
	sRows = append(sRows, sRow)

	rows := cu.ToIMA(stateData["rows"], []cu.IM{})
	// table data
	for _, tRow := range rows {
		sRow = make([]string, 0)
		for _, field := range browserFields {
			if visibleField(field.Name) {
				sRow = append(sRow, cu.ToString(tRow[field.Name], ""))
			}
		}
		sRows = append(sRows, sRow)
	}

	var csvData []byte
	var b bytes.Buffer
	writr := csv.NewWriter(&b)
	userConfig := cu.ToIM(client.Ticket.User["auth_map"], cu.IM{})
	writr.Comma = rune(cu.ToString(userConfig["export_sep"], st.DefaultExportSep)[0])
	if err := writr.WriteAll(sRows); err == nil {
		csvData = b.Bytes()
	}

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename="+fileName)
	w.Write(csvData)
	w.WriteHeader(http.StatusOK)
}

func ClientExportReport(w http.ResponseWriter, r *http.Request) {
	sessionID := r.PathValue("session_id")
	output := cu.ToString(r.URL.Query().Get("output"), "pdf")
	queueCode := cu.ToString(r.URL.Query().Get("queue"), "")
	export := cu.ToBoolean(r.URL.Query().Get("export"), false)

	disposition := "attachment"
	if cu.ToBoolean(r.URL.Query().Get("inline"), false) {
		disposition = "inline"
	}
	cs := r.Context().Value(md.ClientServiceCtxKey).(*cls.ClientService)

	inputType := func() string {
		if queueCode != "" {
			return "queue"
		}
		if export {
			return "export"
		}
		return "report"
	}

	var client *ct.Client
	var err error
	if client, err = cs.LoadSession(sessionID); err != nil ||
		!client.Ticket.Valid() && !client.LoginDisabled {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	options := cu.IM{}
	ds := cs.NewDataStore(cs.Config, client.Ticket.Database, cs.AppLog)

	switch inputType() {
	case "queue":
		var rows []cu.IM = []cu.IM{}
		if rows, err = ds.StoreDataQuery(md.Query{
			Fields: []string{"*"}, From: "config_print_queue",
			Filters: []md.Filter{
				{Field: "code", Comp: "==", Value: queueCode},
			},
		}, true); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		options = cu.IM{
			"report_key":  rows[0]["report_code"],
			"orientation": rows[0]["orientation"],
			"size":        rows[0]["paper_size"],
			"code":        rows[0]["ref_code"],
			"output":      output,
			"filters":     cu.IM{},
			"queue_id":    rows[0]["id"],
		}

	case "export":
		editor := cu.ToIM(client.Data["editor"], cu.IM{})
		shortcut := cu.ToIM(editor["shortcut"], cu.IM{})
		params := cu.ToIM(editor["params"], cu.IM{})

		options = cu.IM{
			"report_key": cu.ToString(shortcut["report_key"], ""),
			"output":     output,
			"filters":    params,
		}

	default:
		modal := cu.ToIM(client.Data["modal"], cu.IM{})
		modalData := cu.ToIM(modal["data"], cu.IM{})
		options = cu.IM{
			"report_key":  modalData["template"],
			"orientation": modalData["orientation"],
			"size":        modalData["paper_size"],
			"code":        modalData["code"],
			"output":      output,
			"filters":     cu.IM{},
		}
	}

	results, err := ds.GetReport(options)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fileName := fmt.Sprintf("%s.%s", sessionID, output)
	w.Header().Set("Content-Type", cu.ToString(results["content_type"], ""))
	w.Header().Set("Content-Disposition", disposition+";filename="+fileName)
	if cu.ToString(results["content_type"], "") == "application/pdf" {
		w.Write(results["template"].([]uint8))
		return
	}
	w.Write([]byte(results["template"].(string)))
}

func ClientTemplateEditor(w http.ResponseWriter, r *http.Request) {
	sessionID := r.PathValue("session_id")
	code := r.PathValue("code")
	cs := r.Context().Value(md.ClientServiceCtxKey).(*cls.ClientService)

	var client *ct.Client
	var err error
	if client, err = cs.LoadSession(sessionID); err != nil ||
		!client.Ticket.Valid() && !client.LoginDisabled {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	ds := cs.NewDataStore(cs.Config, client.Ticket.Database, cs.AppLog)
	var rows []cu.IM = []cu.IM{}
	if rows, err = ds.StoreDataQuery(md.Query{
		Fields: []string{"*"}, From: "config",
		Filters: []md.Filter{
			{Field: "code", Comp: "==", Value: code},
		},
	}, false); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var token string
	user := client.Ticket.User
	if token, err = ds.CreateLoginToken(
		cu.SM{"code": cu.ToString(user["code"], ""), "user_name": cu.ToString(user["user_name"], ""), "scope": cu.ToString(user["user_group"], ""), "alias": ds.Alias}, ds.Config); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := cu.IM{
		"token":  token,
		"report": rows[0],
	}

	w.Header().Set("Content-Type", "application/json")
	if response, err := cu.ConvertToByte(result); err == nil {
		w.Write(response)
	}

}
