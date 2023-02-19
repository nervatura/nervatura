//go:build http || all
// +build http all

package service

import (
	"errors"
	"net/http"
	"os"
	"path"
	"sort"
	"strings"
	"text/template"

	nt "github.com/nervatura/nervatura/service/pkg/nervatura"
	ut "github.com/nervatura/nervatura/service/pkg/utils"
)

// LocalesService implements the Nervatura Client translation helper tool
type LocalesService struct {
	ClientMsg    string
	ConfigFile   string
	Version      string
	GetParam     func(req *http.Request, name string) string
	theme        string
	templates    *template.Template
	templateName string
	langs        nt.SL
	tagKeys      nt.SL
	tagValues    map[string]nt.SL
	filter       string
	edited       bool
	deflang      nt.IM
	locfile      nt.IM
}

func (loc *LocalesService) loadData() error {
	loc.langs = nt.SL{"client"}
	loc.tagKeys = make(nt.SL, 0)
	loc.tagValues = make(map[string]nt.SL)
	loc.edited = false
	loc.locfile = nt.IM{
		"locales": make(nt.IM),
	}

	var jsonMessages, _ = ut.Static.ReadFile(loc.ClientMsg)
	if err := ut.ConvertFromByte(jsonMessages, &loc.deflang); err != nil {
		return err
	}
	for rowKey := range loc.deflang {
		tag := strings.Split(rowKey, "_")[0]
		if !ut.Contains(loc.tagKeys, tag) && !ut.Contains(nt.SL{"en", "key"}, tag) {
			loc.tagKeys = append(loc.tagKeys, tag)
			loc.tagValues[tag] = make(nt.SL, 0)
		}
		loc.tagValues[tag] = append(loc.tagValues[tag], rowKey)
		sort.Strings(loc.tagValues[tag])
	}
	sort.Strings(loc.tagKeys)

	if content, err := os.ReadFile(loc.ConfigFile); err == nil {
		config := nt.IM{}
		if err = ut.ConvertFromByte(content, &config); err == nil {
			if locales, valid := config["locales"].(nt.IM); valid {
				loc.locfile["locales"] = locales
				for langKey := range locales {
					loc.langs = append(loc.langs, langKey)
					sort.Strings(loc.langs)
				}
			}
		}
	}
	return nil
}

// LoadLocales -Load all json files from data dir
func (loc *LocalesService) LoadLocales() error {
	loc.templates, _ = template.ParseFS(ut.Static, path.Join("static", "views", "*.html"))
	loc.templateName = "locales"
	loc.theme = "light"
	return loc.loadData()
}

func (loc *LocalesService) sendMsg(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	response, _ := ut.ConvertToByte(nt.IM{"code": code, "message": message})
	w.Write(response)
}

// createView - Create a result view data
func (loc *LocalesService) createView(lang, tag string) nt.IM {
	tags := loc.tagKeys
	if lang != "client" {
		tags = append(tags, "missing")
	}
	if !ut.Contains(loc.langs, lang) {
		lang = "client"
	}
	if !ut.Contains(tags, tag) {
		tag = tags[0]
	}
	toLangValue := func(key string, value interface{}) []map[string]string {
		values := make([]map[string]string, 0)
		if v, valid := value.([]interface{}); valid {
			for i := 0; i < len(v); i++ {
				values = append(values, map[string]string{"key": key, "index": ut.ToString(i, ""), "value": ut.ToString(v[i], "")})
			}
		} else {
			values = append(values, map[string]string{"key": key, "index": "0", "value": ut.ToString(value, "")})
		}
		return values
	}
	toValues := func(key string, value interface{}) nt.SL {
		values := make(nt.SL, 0)
		if v, valid := value.([]interface{}); valid {
			for i := 0; i < len(v); i++ {
				values = append(values, ut.ToString(v[i], ""))
			}
		} else {
			values = append(values, ut.ToString(value, ""))
		}
		return values
	}
	find := func(a nt.SL, x string) bool {
		for _, n := range a {
			if strings.Contains(n, x) {
				return true
			}
		}
		return false
	}
	view := nt.IM{
		"theme": loc.theme, "filter": loc.filter, "dirty": ut.ToString(loc.edited, "false"),
		"lang": lang, "tag": tag, "version": loc.Version,
		"langs": loc.langs, "tags": tags,
		"view_admin": ut.GetMessage("view_admin"),
		"labels": map[string]string{
			"title":              ut.GetMessage("view_locales_title"),
			"title_missing":      ut.GetMessage("title_missing"),
			"title_update":       ut.GetMessage("title_update"),
			"title_undo":         ut.GetMessage("title_undo"),
			"title_theme":        ut.GetMessage("title_theme"),
			"title_add":          ut.GetMessage("title_add"),
			"placeholder_filter": ut.GetMessage("placeholder_filter"),
			"placeholder_lcode":  ut.GetMessage("placeholder_lcode"),
			"placeholder_lname":  ut.GetMessage("placeholder_lname"),
			"version":            ut.GetMessage("view_version"),
		},
		"data": make([]nt.IM, 0),
	}
	locales := loc.locfile["locales"].(nt.IM)
	if loc.filter != "" {
		filter := make(nt.SL, 0)
		for rowKey := range loc.deflang {
			if strings.Contains(rowKey, loc.filter) ||
				find(toValues(rowKey, loc.deflang[rowKey]), loc.filter) {
				filter = append(filter, rowKey)
			} else if lang != "client" {
				if find(toValues(rowKey, locales[lang].(nt.IM)[rowKey]), loc.filter) {
					filter = append(filter, rowKey)
				}
			}
		}
		sort.Strings(filter)
		for i := 0; i < len(filter); i++ {
			row := nt.IM{
				"tag":    strings.Split(filter[i], "_")[0],
				"key":    filter[i],
				"client": toLangValue(filter[i], loc.deflang[filter[i]]),
			}
			if lang != "client" {
				row["lang_values"] = toLangValue(filter[i], locales[lang].(nt.IM)[filter[i]])
			}
			view["data"] = append(view["data"].([]nt.IM), row)
		}
	} else if tag == "missing" {
		missing := make(nt.SL, 0)
		for rowKey := range loc.deflang {
			if _, found := locales[lang].(nt.IM)[rowKey]; !found {
				missing = append(missing, rowKey)
			}
		}
		sort.Strings(missing)
		for i := 0; i < len(missing); i++ {
			view["data"] = append(view["data"].([]nt.IM), nt.IM{
				"tag":         strings.Split(missing[i], "_")[0],
				"key":         missing[i],
				"client":      toLangValue(missing[i], loc.deflang[missing[i]]),
				"lang_values": toLangValue(missing[i], ""),
			})
		}
	} else {
		keys := loc.tagValues[tag]
		for i := 0; i < len(keys); i++ {
			row := nt.IM{
				"tag":    strings.Split(keys[i], "_")[0],
				"key":    keys[i],
				"client": toLangValue(keys[i], loc.deflang[keys[i]]),
			}
			if lang != "client" {
				row["lang_values"] = toLangValue(keys[i], locales[lang].(nt.IM)[keys[i]])
			}
			view["data"] = append(view["data"].([]nt.IM), row)
		}
	}
	return view
}

// Render - template rendering
func (loc *LocalesService) Render(w http.ResponseWriter, r *http.Request) {
	view := loc.createView(loc.GetParam(r, "lang"), loc.GetParam(r, "tag"))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := loc.templates.ExecuteTemplate(w, loc.templateName, view); err != nil {
		http.Error(w, ut.GetMessage("error_internal"), http.StatusInternalServerError)
	}
}

// SetTheme - light/dark form theme
func (loc *LocalesService) SetTheme(w http.ResponseWriter, r *http.Request) {
	theme := nt.SM{
		"light": "dark", "dark": "light",
	}
	loc.theme = theme[loc.theme]
	loc.sendMsg(w, 200, "OK")
}

// SetFilter - input form filter
func (loc *LocalesService) SetFilter(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}
	err := ut.ConvertFromReader(r.Body, &data)
	if err != nil {
		loc.sendMsg(w, http.StatusBadRequest, err.Error())
		return
	}
	loc.filter = ut.ToString(data["filter_value"], "")
	loc.sendMsg(w, 200, "OK")
}

// updateData - Write changes to the language map
func (loc *LocalesService) updateData(lang, key, value string, index int64) error {
	locales := loc.locfile["locales"].(nt.IM)
	if _, found := locales[lang]; found {
		if ivalue, found := loc.deflang[key]; found {
			if values, valid := ivalue.([]interface{}); valid {
				if len(values) > int(index) {
					if _, init := locales[lang].(map[string]interface{})[key].([]interface{}); !init {
						locales[lang].(map[string]interface{})[key] = make([]interface{}, len(values))
					}
					locales[lang].(map[string]interface{})[key].([]interface{})[index] = value
				} else {
					return errors.New(ut.GetMessage("update_error"))
				}
			} else {
				locales[lang].(map[string]interface{})[key] = value
			}
			loc.edited = true
			return nil
		}
	}
	return errors.New(ut.GetMessage("update_error"))
}

// UpdateRow - input form value update
func (loc *LocalesService) UpdateRow(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}
	err := ut.ConvertFromReader(r.Body, &data)
	if err != nil {
		loc.sendMsg(w, http.StatusBadRequest, err.Error())
		return
	}
	err = loc.updateData(ut.ToString(data["lang"], ""), ut.ToString(data["key"], ""),
		ut.ToString(data["value"], ""), ut.ToInteger(data["index"], 0))
	if err != nil {
		loc.sendMsg(w, 400, err.Error())
		return
	}
	loc.sendMsg(w, 200, "OK")
}

// UndoAll - Discard all changes
func (loc *LocalesService) UndoAll(w http.ResponseWriter, r *http.Request) {
	err := loc.loadData()
	if err == nil {
		loc.sendMsg(w, 200, "OK")
	} else {
		loc.sendMsg(w, 400, err.Error())
	}
}

func (loc *LocalesService) saveFile() error {
	if loc.edited {
		fw, err := os.Create(loc.ConfigFile)
		if err != nil {
			return err
		}
		defer fw.Close()
		return ut.ConvertToWriter(fw, loc.locfile)
	}
	return nil
}

// CreateFile - Write changes to the json files
func (loc *LocalesService) CreateFile(w http.ResponseWriter, r *http.Request) {
	err := loc.saveFile()
	if err == nil {
		err = loc.loadData()
	}
	if err == nil {
		loc.sendMsg(w, 200, "OK")
	} else {
		loc.sendMsg(w, 400, err.Error())
	}
}

// AddLang - Add a new language
func (loc *LocalesService) AddLang(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}
	err := ut.ConvertFromReader(r.Body, &data)
	if err != nil {
		loc.sendMsg(w, http.StatusBadRequest, err.Error())
		return
	}
	lang_key := ut.ToString(data["lang_key"], "")
	lang_name := ut.ToString(data["lang_name"], "")
	if lang_key == "" || lang_name == "" {
		loc.sendMsg(w, http.StatusBadRequest, errors.New(ut.GetMessage("title_missing")).Error())
		return
	}
	if ut.Contains(loc.langs, lang_key) || lang_key == "en" {
		loc.sendMsg(w, http.StatusBadRequest, errors.New(ut.GetMessage("error_existing_lang")).Error())
		return
	}

	loc.langs = append(loc.langs, lang_key)
	sort.Strings(loc.langs)
	loc.locfile["locales"].(nt.IM)[lang_key] = nt.IM{
		"key":    lang_key,
		lang_key: lang_name,
	}
	loc.edited = true
	loc.sendMsg(w, 200, "OK")
}
