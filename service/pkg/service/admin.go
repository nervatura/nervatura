package service

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"sort"
	"strings"
	"text/template"

	db "github.com/nervatura/nervatura-service/pkg/database"
	nt "github.com/nervatura/nervatura-service/pkg/nervatura"
	ut "github.com/nervatura/nervatura-service/pkg/utils"
)

// AdminService implements the Nervatura Admin GUI
type AdminService struct {
	Config        map[string]interface{}
	GetNervaStore func(database string) *nt.NervaStore
	templates     *template.Template
	GetTokenKeys  func() map[string]map[string]string
}

func (adm *AdminService) LoadTemplates() {
	adm.templates, _ = template.ParseFS(ut.Static, path.Join("static", "views", "*.html"))
}

// template rendering
func (adm *AdminService) render(w http.ResponseWriter, template string, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := adm.templates.ExecuteTemplate(w, template, data); err != nil {
		http.Error(w, ut.GetMessage("error_internal"), http.StatusInternalServerError)
	}
}

func (adm *AdminService) parseData(r *http.Request) nt.IM {
	return nt.IM{
		"version": adm.Config["version"], "view_version": ut.GetMessage("view_version"),
		"theme": r.PostFormValue("theme"), "menu": r.PostFormValue("menu"),
		"token":    r.PostFormValue("token"),
		"database": r.PostFormValue("database"), "demo": r.PostFormValue("demo"),
		"apikey":   r.PostFormValue("apikey"),
		"username": r.PostFormValue("username"), "password": r.PostFormValue("password"),
		"confirm": r.PostFormValue("confirm"), "reportkey": r.PostFormValue("reportkey"),
		"errors": nt.SM{}, "result": []nt.SM{},
		"view_admin":    ut.GetMessage("view_admin"),
		"view_database": ut.GetMessage("view_database"), "view_client": ut.GetMessage("view_client"),
		"view_docs": ut.GetMessage("view_docs"), "view_theme": ut.GetMessage("view_theme"),
		"view_login": ut.GetMessage("view_login"), "view_submit": ut.GetMessage("view_submit"),
		"view_username": ut.GetMessage("view_username"), "view_password": ut.GetMessage("view_password"),
		"view_create": ut.GetMessage("view_create"), "view_api_key": ut.GetMessage("view_api_key"),
		"view_alias": ut.GetMessage("view_alias"), "view_demo": ut.GetMessage("view_demo"),
		"view_logout": ut.GetMessage("view_logout"), "view_token": ut.GetMessage("view_token"),
		"view_refresh": ut.GetMessage("view_refresh"), "view_password_change": ut.GetMessage("view_password_change"),
		"view_confirm": ut.GetMessage("view_confirm"), "view_report": ut.GetMessage("view_report"),
		"view_list": ut.GetMessage("view_list"), "view_report_key": ut.GetMessage("view_report_key"),
		"view_install": ut.GetMessage("view_install"), "view_installed": ut.GetMessage("view_installed"),
		"view_name": ut.GetMessage("view_name"), "view_description": ut.GetMessage("view_description"),
		"view_type": ut.GetMessage("view_type"), "view_filename": ut.GetMessage("view_filename"),
		"view_label": ut.GetMessage("view_label"), "view_delete": ut.GetMessage("view_delete"),
		"view_configuration": ut.GetMessage("view_configuration"), "env_result": []nt.SM{},
		"view_envkey": ut.GetMessage("view_envkey"), "view_envvalue": ut.GetMessage("view_envvalue"),
	}
}

func (adm *AdminService) Home(w http.ResponseWriter, r *http.Request) {
	data := adm.parseData(r)
	adm.render(w, "login", data)
}

func (adm *AdminService) Login(w http.ResponseWriter, r *http.Request) {
	data := adm.parseData(r)
	if data["database"] == "" {
		data["errors"].(nt.SM)["login"] = ut.GetMessage("missing_database")
		adm.render(w, "login", data)
		return
	}
	nstore := adm.GetNervaStore(data["database"].(string))
	token, _, err := (&nt.API{NStore: nstore}).UserLogin(data)
	if err != nil {
		data["errors"].(nt.SM)["login"] = err.Error()
		adm.render(w, "login", data)
		return
	}
	if nstore.User.Scope != "admin" {
		data["errors"].(nt.SM)["login"] = ut.GetMessage("admin_rights")
		adm.render(w, "login", data)
		return
	}
	data["token"] = token
	data["password"] = ""
	adm.render(w, "admin", data)
}

func (adm *AdminService) Menu(w http.ResponseWriter, r *http.Request) {
	data := adm.parseData(r)
	switch data["menu"] {
	case "database":
		adm.render(w, "database", data)
	case "theme":
		if data["theme"] == "" || data["theme"] == "light" {
			data["theme"] = "dark"
		} else {
			data["theme"] = "light"
		}
		adm.render(w, r.PostFormValue("pageID"), data)
	case "logout":
		data["token"] = ""
		adm.render(w, "login", data)
	default:
		adm.render(w, "login", data)
	}
}

func (adm *AdminService) Admin(w http.ResponseWriter, r *http.Request) {
	data := adm.parseData(r)
	unauthorized := func(errMsg string) {
		data["errors"].(nt.SM)["login"] = errMsg
		adm.render(w, "login", data)
	}

	nstore := adm.GetNervaStore(data["database"].(string))
	err := (&nt.API{NStore: nstore}).TokenLogin(nt.IM{"token": data["token"].(string), "keys": adm.GetTokenKeys()})
	if err != nil {
		unauthorized(err.Error())
		return
	}
	if nstore.User == nil || nstore.User.Scope != "admin" {
		unauthorized(ut.GetMessage("admin_rights"))
		return
	}
	switch r.PostFormValue("cmd") {
	case "refresh":
		data["token"], err = (&nt.API{NStore: nstore}).TokenRefresh()
	case "password":
		err = (&nt.API{NStore: nstore}).UserPassword(data)
		if err == nil {
			data["success"] = ut.GetMessage("password_change")
			data["password"] = ""
			data["confirm"] = ""
		}
	case "list":
		data["result"], err = (&nt.API{NStore: nstore}).ReportList(data)
	case "install":
		var id int64
		id, err = (&nt.API{NStore: nstore}).ReportInstall(nt.IM{"reportkey": ut.ToString(data["reportkey"], "")})
		if err == nil {
			data["success"] = fmt.Sprintf(ut.GetMessage("result_id"), id)
			data["reportkey"] = ""
		}
	case "delete":
		err = (&nt.API{NStore: nstore}).ReportDelete(nt.IM{"reportkey": ut.ToString(data["reportkey"], "")})
		if err == nil {
			data["success"] = ut.GetMessage("successful_delete")
			data["reportkey"] = ""
		}
	case "env_list":
		envResult := make([]nt.SM, 0)
		keys := make([]string, 0)
		configs := nt.IM{}
		for key, value := range adm.Config {
			keys = append(keys, key)
			configs[key] = value
		}
		for _, env := range os.Environ() {
			if strings.HasPrefix(env, "NT_ALIAS_") {
				configs[strings.Split(env, "=")[0]] = strings.Split(env, "=")[1]
			}
		}

		sort.Strings(keys)
		for _, key := range keys {
			envResult = append(envResult, nt.SM{"envkey": strings.ToUpper(key), "envvalue": ut.ToString(adm.Config[key], "")})
		}
		data["env_result"] = envResult
	}
	if err != nil {
		data["errors"].(nt.SM)["admin"] = err.Error()
	}
	adm.render(w, "admin", data)
}

func (adm *AdminService) Database(w http.ResponseWriter, r *http.Request) {
	data := adm.parseData(r)
	if adm.Config["NT_API_KEY"] != data["apikey"] {
		data["errors"].(nt.SM)["database"] = ut.GetMessage("invalid_api_key")
		adm.render(w, "database", data)
		return
	}
	data["result"], _ = (&nt.API{NStore: nt.New(&db.SQLDriver{Config: adm.Config}, adm.Config)}).DatabaseCreate(data)
	adm.render(w, "database", data)
}
