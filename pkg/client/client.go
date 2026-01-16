package client

import (
	"html/template"
	"net/http"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
	st "github.com/nervatura/nervatura/v6/static"
)

type SessionRequest struct {
	Database  string `json:"database"`
	Username  string `json:"username"`
	Lang      string `json:"lang"`
	Module    string `json:"module"`
	Theme     string `json:"theme"`
	RequestID string `json:"request_id"`
	Data      cu.IM  `json:"data"`
}

func RespondMessage(w http.ResponseWriter, html template.HTML, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}

func ErrorPage(w http.ResponseWriter, title, msg string) {
	var err error
	var html template.HTML
	msg = cu.ToString(msg, http.StatusText(http.StatusUnauthorized))
	pageData := cu.IM{
		"title":       ut.GetMessage("app_title"),
		"error_title": title,
		"error_msg":   msg,
		"login":       ut.GetMessage("title_login"),
	}
	html, err = cu.TemplateBuilder("error", st.ErrorTemplate, map[string]any{}, pageData)
	RespondMessage(w, html, err)
}

func Application(sessionID string, mainComponent ct.ClientComponent) (ccApp *ct.Application) {
	return &ct.Application{
		Title: ut.GetMessage("app_title"),
		Header: cu.SM{
			"X-Session-Token": sessionID,
		},
		Script: []string{
			"/static/js/htmx.min.js",
			//"/static/js/htmx.v1.min.js",
			//"https://unpkg.com/htmx.org@latest",
			//"https://unpkg.com/htmx.org@2.0.3",
			"/static/js/remove-me.js",
		},
		HeadLink: []ct.HeadLink{
			{Rel: "icon", Href: "/static/favicon.svg", Type: "image/svg+xml"},
			{Rel: "stylesheet", Href: "/public/css/app.css"},
			{Rel: "stylesheet", Href: "/static/css/index.css"},
		},
		MainComponent: mainComponent,
	}
}
