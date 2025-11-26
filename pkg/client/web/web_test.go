package web_test

import (
	"errors"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	ct "github.com/nervatura/component/pkg/component"
	"github.com/nervatura/nervatura/v6/pkg/client/web"
)

func TestRespondMessage(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w    http.ResponseWriter
		html template.HTML
		err  error
	}{
		{
			name: "ok",
			w:    httptest.NewRecorder(),
			html: template.HTML("test"),
			err:  nil,
		},
		{
			name: "err",
			w:    httptest.NewRecorder(),
			html: template.HTML("test"),
			err:  errors.New("test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			web.RespondMessage(tt.w, tt.html, tt.err)
		})
	}
}

func TestErrorPage(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w     http.ResponseWriter
		title string
		msg   string
	}{
		{
			name:  "ok",
			w:     httptest.NewRecorder(),
			title: "test",
			msg:   "test",
		},
		{
			name:  "err",
			w:     httptest.NewRecorder(),
			title: "test",
			msg:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			web.ErrorPage(tt.w, tt.title, tt.msg)
		})
	}
}

func TestApplication(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		sessionID     string
		mainComponent ct.ClientComponent
	}{
		{
			name:          "ok",
			sessionID:     "test",
			mainComponent: &ct.Button{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			web.Application(tt.sessionID, tt.mainComponent)

		})
	}
}
