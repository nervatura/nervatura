//go:build http || all
// +build http all

package service

import (
	"text/template"
)

// ClientService serve the Nervatura Client
type ClientService struct {
	Path      string
	Manifest  Manifest
	templates *template.Template
}

// Manifest reflects the structure of asset-manifest.json
type Manifest struct {
	Files       map[string]string `json:"files"`
	Entrypoints Entrypoints
}

type Entrypoints []string

/*
func (cli *ClientService) LoadManifest() (err error) {

	cli.templates, err = template.ParseFS(ut.Static, "static/views/client.html")
	if err != nil {
		return err
	}

	manifest, err := ut.Public.ReadFile("static/client/asset-manifest.json")
	if err == nil {
		err = ut.ConvertFromByte(manifest, &cli.Manifest)
	}

	return err
}

func (cli *ClientService) Render(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := cli.templates.ExecuteTemplate(w, "client.html", cli.Manifest); err != nil {
		http.Error(w, ut.GetMessage("error_internal"), http.StatusInternalServerError)
	}
}

func (e Entrypoints) Scripts() Entrypoints {
	var scripts Entrypoints

	for _, f := range e {
		if strings.HasSuffix(f, ".js") {
			scripts = append(scripts, f)
		}
	}

	return scripts
}

func (e Entrypoints) Styles() Entrypoints {
	var styles Entrypoints

	for _, f := range e {
		if strings.HasSuffix(f, ".css") {
			styles = append(styles, f)
		}
	}

	return styles
}
*/
