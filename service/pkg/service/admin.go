//go:build http || all
// +build http all

package service

import (
	"errors"
	"net/http"
	"os"
	"sort"
	"strings"
	"text/template"

	fm "github.com/nervatura/component/component/atom"
	bc "github.com/nervatura/component/component/base"
	pg "github.com/nervatura/component/component/page"
	cp "github.com/nervatura/nervatura/service/pkg/component"
	db "github.com/nervatura/nervatura/service/pkg/database"
	nt "github.com/nervatura/nervatura/service/pkg/nervatura"
	ut "github.com/nervatura/nervatura/service/pkg/utils"
)

const taskPage = `<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0, viewport-fit=cover" />
		<title>{{ .title }}</title>
		<link rel="icon" type="image/svg+xml" href="/style/static/favicon.svg">
		<link rel="stylesheet" href="/style/index.css" />
		<link rel="stylesheet" href="/css/admin.css" />
	</head>
	<body><div class="admin row mobile" theme="dark" style="margin:auto;">
	{{if .env_result}}
		<div class="container section">
		{{range .env_result}}
			<div class="row full border-top" >
				<div class="cell mobile small">
					<div class="cell padding-normal bold" style="white-space:nowrap;vertical-align: top;" >{{ .envkey }}</div>
				<div class="cell mobile padding-normal" >
				{{if .envvalue}}<span style="color:rgb(var(--functional-green));white-space:wrap;">{{ .envvalue }}</span>
				{{else}}<span style="color:rgb(var(--functional-red));">X</span>{{end}}
				</div>
			</div>
		</div>
		{{end}}
	</div>
	{{end}}
	</div></body>
</html>`

// AdminService implements the Nervatura Admin GUI
type AdminService struct {
	Config        map[string]interface{}
	GetNervaStore func(database string) *nt.NervaStore
	GetParam      func(req *http.Request, name string) string
	GetTokenKeys  func() map[string]map[string]string
	GetTaskSecKey func() string
	Session       nt.SessionService
}

func (adm *AdminService) envList() []nt.IM {
	envResult := make([]nt.IM, 0)
	keys := make([]string, 0)
	configs := nt.IM{}
	for key, value := range adm.Config {
		keys = append(keys, key)
		configs[key] = value
	}
	for _, env := range os.Environ() {
		if strings.HasPrefix(env, "NT_ALIAS_") {
			keys = append(keys, strings.Split(env, "=")[0])
			configs[strings.Split(env, "=")[0]] = strings.Split(env, "=")[1]
		}
	}

	sort.Strings(keys)
	for _, key := range keys {
		envResult = append(envResult, nt.IM{"envkey": strings.ToUpper(key), "envvalue": ut.ToString(configs[key], "")})
	}
	return envResult
}

func (adm *AdminService) createDatabase(data nt.IM) (result []nt.IM, err error) {
	if adm.Config["NT_API_KEY"] != data["api_key"] {
		return result, errors.New(ut.GetMessage("invalid_api_key"))
	}
	log, err := (&nt.API{NStore: nt.New(&db.SQLDriver{Config: adm.Config}, adm.Config)}).DatabaseCreate(nt.IM{
		"database": data["alias"], "demo": data["demo"],
	})
	if err == nil {
		result = ut.SMAtoIMA(log)
	}
	return result, err
}

func (adm *AdminService) userLogin(data nt.IM) (token string, nstore *nt.NervaStore, err error) {
	nstore = adm.GetNervaStore(ut.ToString(data["database"], ""))
	token, _, err = (&nt.API{NStore: nstore}).UserLogin(data)
	if err == nil {
		if nstore.User.Scope != "admin" {
			return token, nstore, errors.New(ut.GetMessage("admin_rights"))
		}
	}
	return token, nstore, err
}

func (adm *AdminService) reportInstall(token, database, reportkey string) (nstore *nt.NervaStore, err error) {
	nstore = adm.GetNervaStore(database)
	err = (&nt.API{NStore: nstore}).TokenLogin(nt.IM{"token": token, "keys": adm.GetTokenKeys()})
	if err == nil {
		_, err = (&nt.API{NStore: nstore}).ReportInstall(nt.IM{"reportkey": reportkey})
	}
	return nstore, err
}

func (adm *AdminService) reportDelete(token, database, reportkey string) (nstore *nt.NervaStore, err error) {
	nstore = adm.GetNervaStore(database)
	err = (&nt.API{NStore: nstore}).TokenLogin(nt.IM{"token": token, "keys": adm.GetTokenKeys()})
	if err == nil {
		err = (&nt.API{NStore: nstore}).ReportDelete(nt.IM{"reportkey": reportkey})
	}
	return nstore, err
}

func (adm *AdminService) userPassword(token, database string, data nt.IM) (err error) {
	nstore := adm.GetNervaStore(database)
	err = (&nt.API{NStore: nstore}).TokenLogin(nt.IM{"token": token, "keys": adm.GetTokenKeys()})
	if err == nil {
		err = (&nt.API{NStore: nstore}).UserPassword(data)
	}
	return err
}

func (adm *AdminService) appResponse(evt bc.ResponseEvent) (re bc.ResponseEvent) {
	errMsg := func(value, toastType string) (re bc.ResponseEvent) {
		return bc.ResponseEvent{
			Trigger: &fm.Toast{
				Type:    toastType,
				Value:   value,
				Timeout: 0,
			},
			TriggerName: evt.TriggerName,
			Name:        evt.Name,
			Header: bc.SM{
				bc.HeaderRetarget: "#toast-msg",
				bc.HeaderReswap:   "innerHTML",
			},
		}
	}
	switch evt.Name {
	case cp.AdminEventCreate:
		data := evt.Trigger.GetProperty("data").(nt.IM)
		result, err := adm.createDatabase(data)
		if err != nil {
			return errMsg(err.Error(), fm.ToastTypeError)
		}
		evt.Trigger.SetProperty("data", bc.MergeIM(data, nt.IM{"create_result": result}))
		evt.Trigger.SetProperty("token", "")

	case cp.AdminEventLogin:
		data := evt.Trigger.GetProperty("data").(nt.IM)
		token, nstore, err := adm.userLogin(data)
		if err != nil {
			return errMsg(err.Error(), fm.ToastTypeError)
		}
		evt.Trigger.SetProperty("token", token)
		evt.Trigger.SetProperty("view", "password")
		evt.Trigger.SetProperty("data", bc.MergeIM(data, nt.IM{"env_list": adm.envList()}))
		reportList, _ := (&nt.API{NStore: nstore}).ReportList(data)
		evt.Trigger.SetProperty("data", bc.MergeIM(data, nt.IM{"report_list": reportList}))

	case cp.AdminEventReportInstall, cp.AdminEventReportDelete:
		data := evt.Trigger.GetProperty("data").(nt.IM)
		database := ut.ToString(data["database"], "")
		token := ut.ToString(evt.Trigger.GetProperty("token"), "")
		reportkey := ut.ToString(evt.Value, "")
		var nstore *nt.NervaStore
		var err error
		if evt.Name == cp.AdminEventReportDelete {
			nstore, err = adm.reportDelete(token, database, reportkey)
		} else {
			nstore, err = adm.reportInstall(token, database, reportkey)
		}
		if err != nil {
			return errMsg(err.Error(), fm.ToastTypeError)
		}
		reportList, _ := (&nt.API{NStore: nstore}).ReportList(data)
		evt.Trigger.SetProperty("data", bc.MergeIM(data, nt.IM{"report_list": reportList}))

	case cp.AdminEventPassword:
		data := evt.Trigger.GetProperty("data").(nt.IM)
		database := ut.ToString(data["database"], "")
		token := ut.ToString(evt.Trigger.GetProperty("token"), "")
		err := adm.userPassword(token, database, data)
		if err != nil {
			return errMsg(err.Error(), fm.ToastTypeError)
		}
		return errMsg(ut.GetMessage("password_change"), fm.ToastTypeSuccess)

	}
	return evt
}

func (adm *AdminService) tokenLogin(database, token string) bool {
	if (database != "") && (token != "") {
		nstore := adm.GetNervaStore(database)
		err := (&nt.API{NStore: nstore}).TokenLogin(nt.IM{"token": token, "keys": adm.GetTokenKeys()})
		if err == nil {
			return true
		}
	}
	return false
}

func (adm *AdminService) Home(w http.ResponseWriter, r *http.Request) {
	sessionID := ut.RandString(30)
	admin := &cp.Admin{
		BaseComponent: bc.BaseComponent{
			Id:           bc.GetComponentID(),
			EventURL:     "/admin/event",
			OnResponse:   adm.appResponse,
			RequestValue: map[string]bc.IM{},
			RequestMap:   map[string]bc.ClientComponent{},
			Data: bc.IM{
				"alias": "demo",
				"demo":  "true",
			},
		},
		Module:     "database",
		HelpURL:    ut.ToString(adm.Config["NT_DOCS_URL"], ""),
		ClientURL:  "/client",
		LocalesURL: "/locales",
		Version:    ut.ToString(adm.Config["version"], ""),
		TokenLogin: adm.tokenLogin,
	}
	ccApp := &pg.Application{
		Title:  ut.GetMessage("view_admin"),
		Header: bc.SM{"X-Session-Token": sessionID},
		/*
			Script: []string{
				"/js/htmx.min.js",
				"/js/remove-me.js",
			},
		*/
		HeadLink: []pg.HeadLink{
			{Rel: "icon", Href: "/style/static/favicon.svg", Type: "image/svg+xml"},
			{Rel: "stylesheet", Href: "/style/index.css"},
			{Rel: "stylesheet", Href: "/css/admin.css"},
		},
		MainComponent: admin,
	}
	var err error
	var res string
	res, err = ccApp.Render()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	adm.Session.SaveSession(sessionID, admin)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(res))
}

func (adm *AdminService) AppEvent(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	sessionID := r.Header.Get("X-Session-Token")
	te := bc.TriggerEvent{
		Id:     r.Header.Get("HX-Trigger"),
		Name:   r.Header.Get("HX-Trigger-Name"),
		Target: r.Header.Get("HX-Target"),
		Values: r.Form,
	}
	var err error
	var evt bc.ResponseEvent
	var admin *cp.Admin
	var memAdmin interface{}
	memAdmin, err = adm.Session.LoadSession(sessionID, &admin)
	if err == nil {
		if memAdmin, found := memAdmin.(*cp.Admin); found {
			admin = memAdmin
			evt = admin.OnRequest(te)
		} else {
			admin.OnResponse = adm.appResponse
			admin.TokenLogin = adm.tokenLogin
			_, err := admin.Render()
			if err == nil {
				evt = admin.OnRequest(te)
			}
		}
	}

	for key, value := range evt.Header {
		w.Header().Set(key, value)
	}
	var res string
	if evt.Trigger != nil {
		res, err = evt.Trigger.Render()
	}
	if err != nil {
		res, _ = (&fm.Toast{
			Type: fm.ToastTypeError, Value: err.Error(),
		}).Render()
	}
	adm.Session.SaveSession(sessionID, admin)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(res))
}

func (adm *AdminService) Task(w http.ResponseWriter, r *http.Request) {
	taskName := adm.GetParam(r, "taskName")
	secKey := adm.GetParam(r, "secKey")
	if secKey != adm.GetTaskSecKey() {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(ut.GetMessage("error_unauthorized")))
		return
	}

	data := nt.IM{}
	if taskName == "config" {
		data["title"] = ut.GetMessage("view_configuration")
		data["env_result"] = adm.envList()
	}
	tmp, _ := template.New("task").Parse(taskPage)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmp.ExecuteTemplate(w, "task", data); err != nil {
		http.Error(w, ut.GetMessage("error_internal"), http.StatusInternalServerError)
	}
}
