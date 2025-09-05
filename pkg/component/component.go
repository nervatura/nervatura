package component

import (
	"html/template"
	"io"
	"net/http"
	"strings"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
	st "github.com/nervatura/nervatura/v6/pkg/static"
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
		"title":       ut.GetMessage("en", "app_title"),
		"error_title": title,
		"error_msg":   msg,
		"login":       ut.GetMessage("en", "title_login"),
	}
	html, err = cu.TemplateBuilder("error", st.ErrorTemplate, map[string]any{}, pageData)
	RespondMessage(w, html, err)
}

func Application(sessionID string, mainComponent ct.ClientComponent) (ccApp *ct.Application) {
	return &ct.Application{
		Title: ut.GetMessage("en", "app_title"),
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

func TriggerEvent(r *http.Request) (te ct.TriggerEvent, err error) {
	te = ct.TriggerEvent{
		Id:     r.Header.Get("HX-Trigger"),
		Name:   r.Header.Get("HX-Trigger-Name"),
		Target: r.Header.Get("HX-Target"),
	}

	switch strings.Split(r.Header.Get("Content-Type"), ";")[0] {
	case "multipart/form-data":
		// File upload handling
		//var file multipart.File
		//var handler *multipart.FileHeader
		//var dst *os.File
		// Parse request body as multipart form data with 32MB max memory
		if err = r.ParseMultipartForm(32 << 20); err == nil {
			// Get file from Form
			_, _, err = r.FormFile("file")
			/*
				if file, _, err = r.FormFile("file"); err == nil {
					// Create file locally
					if dst, err = os.Create(handler.Filename); err == nil {
						// Copy the uploaded file data to the newly created file on the filesystem
						_, err = io.Copy(dst, file)
					}
					defer dst.Close()
				}
				defer file.Close()
			*/
		}
	case "application/x-www-form-urlencoded":
		if err = r.ParseForm(); err == nil {
			te.Values = r.Form
		}
	default:
		// text/plain, application/json
		te.Data, err = io.ReadAll(r.Body)
	}
	return te, err
}

func EvtRedirect(name, triggerName, url string) ct.ResponseEvent {
	return ct.ResponseEvent{
		Trigger:     &ct.BaseComponent{},
		TriggerName: name,
		Name:        triggerName,
		Header: cu.SM{
			ct.HeaderReswap:   ct.SwapNone,
			ct.HeaderRedirect: url,
		},
	}
}
