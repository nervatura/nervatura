//go:build http || all
// +build http all

package service

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	nt "github.com/nervatura/nervatura/service/pkg/nervatura"
	ut "github.com/nervatura/nervatura/service/pkg/utils"
)

// HTTPService implements the Nervatura API service
type HTTPService struct {
	Config        map[string]interface{}
	GetNervaStore func(database string) *nt.NervaStore
	GetParam      func(req *http.Request, name string) string
	GetTokenKeys  func() map[string]map[string]string
}

const contentKey = "Content-Type"

// respondMessage write json response format
func (srv *HTTPService) respondMessage(w http.ResponseWriter, code int, payload interface{}, errCode int, err error) {
	var response []byte
	var jerr error
	if err != nil || payload != nil {
		w.Header().Set(contentKey, "application/json")
		if err != nil {
			w.WriteHeader(errCode)
			response, jerr = ut.ConvertToByte(nt.SM{"code": strconv.Itoa(errCode), "message": err.Error()})
		} else {
			w.WriteHeader(code)
			response, jerr = ut.ConvertToByte(payload)
		}
		if jerr == nil {
			w.Write(response)
		}
	} else {
		w.WriteHeader(code)
	}
}

func (srv *HTTPService) ClientConfig(w http.ResponseWriter, r *http.Request) {
	config := nt.IM{}
	if srv.Config["NT_CLIENT_CONFIG"] != "" {
		if content, err := os.ReadFile(srv.Config["NT_CLIENT_CONFIG"].(string)); err == nil {
			_ = ut.ConvertFromByte(content, &config)
		}
	}
	srv.respondMessage(w, http.StatusOK, config, http.StatusBadRequest, nil)
}

func (srv *HTTPService) TokenLogin(r *http.Request) (ctx context.Context, err error) {
	tokenStr := ""
	bearer := r.Header.Get("Authorization")
	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
		tokenStr = bearer[7:]
	}
	if tokenStr == "" {
		return ctx, errors.New(ut.GetMessage("error_unauthorized"))
	}
	claim, err := ut.TokenDecode(tokenStr)
	if err != nil {
		return ctx, err
	}
	database := ut.ToString(claim["database"], "")
	nstore := srv.GetNervaStore(database)
	err = (&nt.API{NStore: nstore}).TokenLogin(nt.IM{"token": tokenStr, "keys": srv.GetTokenKeys()})
	if err != nil {
		return ctx, err
	}
	ctx = context.WithValue(r.Context(), NstoreCtxKey, nstore)
	return ctx, nil
}

func (srv *HTTPService) UserLogin(w http.ResponseWriter, r *http.Request) {
	data := nt.IM{}
	err := ut.ConvertFromReader(r.Body, &data)
	if err != nil {
		srv.respondMessage(w, 0, nil, http.StatusBadRequest, err)
		return
	}
	if _, found := data["database"]; !found {
		srv.respondMessage(w, 0, nil, http.StatusBadRequest, errors.New(ut.GetMessage("missing_database")))
		return
	}
	nstore := srv.GetNervaStore(data["database"].(string))
	token, engine, err := (&nt.API{NStore: nstore}).UserLogin(data)
	srv.respondMessage(w, http.StatusOK, nt.SM{"token": token, "engine": engine, "version": srv.Config["version"].(string)}, http.StatusBadRequest, err)
}

func (srv *HTTPService) getStore(ctx context.Context, admin bool) (nstore *nt.NervaStore, err error) {
	if ctx.Value(NstoreCtxKey) == nil {
		return nil, errors.New(ut.GetMessage("error_unauthorized"))
	}
	nstore = ctx.Value(NstoreCtxKey).(*nt.NervaStore)
	if nstore.User == nil {
		return nil, errors.New(ut.GetMessage("error_unauthorized"))
	}
	if admin && nstore.User.Scope != "admin" {
		return nil, errors.New(ut.GetMessage("error_unauthorized"))
	}
	return nstore, nil
}

func (srv *HTTPService) UserPassword(w http.ResponseWriter, r *http.Request) {
	data := nt.IM{}
	err := ut.ConvertFromReader(r.Body, &data)
	if err != nil {
		srv.respondMessage(w, 0, nil, http.StatusBadRequest, err)
		return
	}
	username := ut.ToString(data["username"], "")
	custnumber := ut.ToString(data["custnumber"], "")
	nstore, err := srv.getStore(r.Context(), false)
	if err != nil || ((username != nstore.User.Username && nstore.User.Scope != "admin") || (custnumber != "" && nstore.User.Scope != "admin")) {
		srv.respondMessage(w, 0, nil, http.StatusUnauthorized, errors.New(ut.GetMessage("error_unauthorized")))
		return
	}
	if username == "" && custnumber == "" {
		data["username"] = nstore.User.Username
	}
	err = (&nt.API{NStore: nstore}).UserPassword(data)
	srv.respondMessage(w, http.StatusNoContent, nil, http.StatusBadRequest, err)
}

func (srv *HTTPService) TokenValidate(w http.ResponseWriter, r *http.Request) {
	nstore, err := srv.getStore(r.Context(), false)
	if err != nil {
		srv.respondMessage(w, 0, nil, http.StatusUnauthorized, err)
		return
	}
	conn := nstore.Connection()
	results := nt.SM{
		"username": nstore.User.Username,
		"database": conn.Alias,
		"engine":   conn.Engine,
		"version":  srv.Config["version"].(string),
	}
	srv.respondMessage(w, http.StatusOK, results, http.StatusBadRequest, err)
}

func (srv *HTTPService) TokenRefresh(w http.ResponseWriter, r *http.Request) {
	if r.Context().Value(NstoreCtxKey) == nil {
		srv.respondMessage(w, 0, nil, http.StatusUnauthorized, errors.New(ut.GetMessage("error_unauthorized")))
		return
	}
	tokenStr, err := (&nt.API{NStore: r.Context().Value(NstoreCtxKey).(*nt.NervaStore)}).TokenRefresh()
	srv.respondMessage(w, http.StatusOK, nt.SM{"token": tokenStr}, http.StatusBadRequest, err)
}

func (srv *HTTPService) GetFilter(w http.ResponseWriter, r *http.Request) {
	nstore, err := srv.getStore(r.Context(), false)
	if err != nil {
		srv.respondMessage(w, 0, nil, http.StatusUnauthorized, err)
		return
	}

	params := nt.IM{"nervatype": srv.GetParam(r, "nervatype"),
		"metadata": r.URL.Query().Get("metadata")}
	query := strings.Split(r.URL.RawQuery, "&")
	for index := 0; index < len(query); index++ {
		if strings.HasPrefix(query[index], "filter=") {
			filter, err := url.QueryUnescape(query[index][7:])
			if err == nil {
				params["filter"] = filter
			}
		}
	}
	results, err := (&nt.API{NStore: nstore}).Get(params)
	srv.respondMessage(w, http.StatusOK, results, http.StatusBadRequest, err)
}

func (srv *HTTPService) GetIds(w http.ResponseWriter, r *http.Request) {
	nstore, err := srv.getStore(r.Context(), false)
	if err != nil {
		srv.respondMessage(w, 0, nil, http.StatusUnauthorized, err)
		return
	}

	params := nt.IM{"nervatype": srv.GetParam(r, "nervatype"),
		"metadata": r.URL.Query().Get("metadata"), "ids": srv.GetParam(r, "IDs")}
	results, err := (&nt.API{NStore: nstore}).Get(params)
	srv.respondMessage(w, http.StatusOK, results, http.StatusBadRequest, err)
}

func (srv *HTTPService) View(w http.ResponseWriter, r *http.Request) {
	nstore, err := srv.getStore(r.Context(), false)
	if err != nil {
		srv.respondMessage(w, 0, nil, http.StatusUnauthorized, err)
		return
	}

	data := make([]nt.IM, 0)
	err = ut.ConvertFromReader(r.Body, &data)
	if err != nil {
		srv.respondMessage(w, 0, nil, http.StatusBadRequest, err)
		return
	}
	results, err := (&nt.API{NStore: nstore}).View(data)

	srv.respondMessage(w, http.StatusOK, results, http.StatusBadRequest, err)
}

func (srv *HTTPService) Function(w http.ResponseWriter, r *http.Request) {
	nstore, err := srv.getStore(r.Context(), false)
	if err != nil {
		srv.respondMessage(w, 0, nil, http.StatusUnauthorized, err)
		return
	}

	data := nt.IM{}
	err = ut.ConvertFromReader(r.Body, &data)
	if err != nil {
		srv.respondMessage(w, 0, nil, http.StatusBadRequest, err)
		return
	}
	results, err := (&nt.API{NStore: nstore}).Function(data)
	srv.respondMessage(w, http.StatusOK, results, http.StatusBadRequest, err)
}

func (srv *HTTPService) Update(w http.ResponseWriter, r *http.Request) {
	nstore, err := srv.getStore(r.Context(), false)
	if err != nil {
		srv.respondMessage(w, 0, nil, http.StatusUnauthorized, err)
		return
	}

	data := make([]nt.IM, 0)
	err = ut.ConvertFromReader(r.Body, &data)
	if err != nil {
		srv.respondMessage(w, 0, nil, http.StatusBadRequest, err)
		return
	}
	results, err := (&nt.API{NStore: nstore}).Update(srv.GetParam(r, "nervatype"), data)
	srv.respondMessage(w, http.StatusOK, results, http.StatusBadRequest, err)
}

func (srv *HTTPService) Delete(w http.ResponseWriter, r *http.Request) {
	nstore, err := srv.getStore(r.Context(), false)
	if err != nil {
		srv.respondMessage(w, 0, nil, http.StatusUnauthorized, err)
		return
	}

	data := nt.IM{"nervatype": srv.GetParam(r, "nervatype"),
		"id": r.URL.Query().Get("id"), "key": r.URL.Query().Get("key")}
	err = (&nt.API{NStore: nstore}).Delete(data)
	srv.respondMessage(w, http.StatusNoContent, nil, http.StatusBadRequest, err)
}

func (srv *HTTPService) DatabaseCreate(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-Api-Key")
	if srv.Config["NT_API_KEY"] != apiKey {
		srv.respondMessage(w, 0, nil, http.StatusUnauthorized, errors.New(ut.GetMessage("error_unauthorized")))
		return
	}
	data := nt.IM{"database": r.URL.Query().Get("alias"), "demo": r.URL.Query().Get("demo")}
	nstore := srv.GetNervaStore(data["database"].(string))
	log, err := (&nt.API{NStore: nstore}).DatabaseCreate(data)
	srv.respondMessage(w, http.StatusOK, log, http.StatusBadRequest, err)
}

func (srv *HTTPService) ReportList(w http.ResponseWriter, r *http.Request) {
	nstore, err := srv.getStore(r.Context(), true)
	if err != nil {
		srv.respondMessage(w, 0, nil, http.StatusUnauthorized, err)
		return
	}

	params := nt.IM{"label": r.URL.Query().Get("label")}
	results, err := (&nt.API{NStore: nstore}).ReportList(params)
	srv.respondMessage(w, http.StatusOK, results, http.StatusBadRequest, err)
}

func (srv *HTTPService) ReportInstall(w http.ResponseWriter, r *http.Request) {
	nstore, err := srv.getStore(r.Context(), true)
	if err != nil {
		srv.respondMessage(w, 0, nil, http.StatusUnauthorized, err)
		return
	}

	params := nt.IM{"reportkey": r.URL.Query().Get("reportkey")}
	results, err := (&nt.API{NStore: nstore}).ReportInstall(params)
	srv.respondMessage(w, http.StatusOK, results, http.StatusBadRequest, err)
}

func (srv *HTTPService) ReportDelete(w http.ResponseWriter, r *http.Request) {
	nstore, err := srv.getStore(r.Context(), true)
	if err != nil {
		srv.respondMessage(w, 0, nil, http.StatusUnauthorized, err)
		return
	}

	params := nt.IM{"reportkey": r.URL.Query().Get("reportkey")}
	err = (&nt.API{NStore: nstore}).ReportDelete(params)
	srv.respondMessage(w, http.StatusNoContent, nil, http.StatusBadRequest, err)
}

func reportQueryFilters(values nt.SM) nt.IM {
	options := nt.IM{"filters": nt.IM{}}
	for key, value := range values {
		if strings.HasPrefix(key, "filters[") {
			fkey := key[8 : len(key)-1]
			options["filters"].(nt.IM)[fkey] = value
		} else {
			switch key {
			case "report_id":
				options["report_id"] = ut.ToInteger(value, 0)
			case "output":
				options["output"] = value
				if value == "data" {
					options["output"] = "tmp"
				}
			default:
				options[key] = value
			}
		}
	}
	return options
}

func reportFilters(values nt.SM, body io.ReadCloser) nt.IM {
	if len(values) > 0 {
		return reportQueryFilters(values)
	}
	data := nt.IM{}
	_ = ut.ConvertFromReader(body, &data)
	return data
}

func queryMap(uri string) nt.SM {
	// Request.URL.Query fixes
	values := make(nt.SM, 0)
	uri, _ = url.QueryUnescape(uri)
	query := strings.Split(uri, "?")
	if len(query) > 1 {
		svalues := strings.Split(query[1], "&")
		for i := 0; i < len(svalues); i++ {
			if len(strings.Split(svalues[i], "=")) > 1 {
				values[strings.Split(svalues[i], "=")[0]] = strings.Split(svalues[i], "=")[1]
			}
		}
	}
	return values
}

func (srv *HTTPService) Report(w http.ResponseWriter, r *http.Request) {
	nstore, err := srv.getStore(r.Context(), false)
	if err != nil {
		srv.respondMessage(w, 0, nil, http.StatusUnauthorized, err)
		return
	}

	options := reportFilters(queryMap(r.RequestURI), r.Body)

	results, err := (&nt.API{NStore: nstore}).Report(options)
	if err != nil {
		srv.respondMessage(w, 0, nil, http.StatusBadRequest, err)
		return
	}
	if options["output"] == "tmp" {
		srv.respondMessage(w, http.StatusOK, results, http.StatusBadRequest, err)
		return
	}
	filetype := nt.SM{
		"csv":    "text/csv",
		"xml":    "application/xml",
		"base64": "application/pdf",
		"pdf":    "application/pdf",
	}
	w.Header().Set(contentKey, filetype[results["filetype"].(string)])
	if results["filetype"] == "pdf" {
		w.Write(results["template"].([]uint8))
		return
	}
	w.Write([]byte(results["template"].(string)))
}
