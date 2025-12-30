package http

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"time"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	st "github.com/nervatura/nervatura/v6/pkg/static"
)

func RespondMessage(w http.ResponseWriter, code int, payload interface{}, errCode int, err error) {
	var response []byte
	var jerr error
	if err != nil || payload != nil {
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(errCode)
			response, jerr = cu.ConvertToByte(cu.SM{"code": strconv.Itoa(errCode), "message": err.Error()})
		} else {
			w.WriteHeader(code)
			response, jerr = cu.ConvertToByte(payload)
		}
		if jerr == nil {
			w.Write(response)
		}
	} else {
		w.WriteHeader(code)
	}
}

func ApiKeyAuth(opt md.AuthOptions) (ctx context.Context, errCode int) {
	apiKey := opt.Request.Header.Get("X-API-KEY")
	if (apiKey != "") && !slices.Contains(strings.Split(cu.ToString(opt.Config["NT_API_KEY"], ""), ","), apiKey) {
		opt.AppLog.Error("API key is not valid")
		return ctx, http.StatusUnauthorized
	}
	alias := cu.ToString(opt.Config["NT_DEFAULT_ALIAS"], "")
	userCode := cu.ToString(opt.Config["NT_DEFAULT_ADMIN"], "")
	ds := api.NewDataStore(opt.Config, alias, opt.AppLog)

	var user md.Auth = md.Auth{
		UserGroup: md.UserGroupAdmin,
		Code:      userCode,
		UserName:  "admin",
	}

	ctx = context.WithValue(opt.Request.Context(), md.DataStoreCtxKey, ds)
	ctx = context.WithValue(ctx, md.AuthUserCtxKey, user)
	return ctx, 0
}

func TokenAuth(opt md.AuthOptions) (ctx context.Context, errCode int) {
	bearerToken := opt.Request.Header.Get("Authorization")

	if len(bearerToken) > 7 && strings.ToUpper(bearerToken[0:6]) == "BEARER" {
		bearerToken = bearerToken[7:]
	}
	if bearerToken == "" {
		return ctx, http.StatusUnauthorized
	}
	var tokenData cu.IM
	var err error
	if tokenData, err = opt.ParseToken(bearerToken, opt.Config["tokenKeys"].([]cu.SM), opt.Config); err != nil {
		opt.AppLog.Error(fmt.Sprintf("Error parsing token: %v", err))
		return ctx, http.StatusUnauthorized
	}
	alias := cu.ToString(tokenData["alias"], "")
	userCode := cu.ToString(tokenData["user_code"], "")
	userName := cu.ToString(tokenData["user_name"], cu.ToString(tokenData["email"], ""))

	if userCode == "" && userName == "" {
		opt.AppLog.Error("Missing API key or bearer token")
		return ctx, http.StatusUnauthorized
	}

	ds := api.NewDataStore(opt.Config, alias, opt.AppLog)
	var user md.Auth
	if user, err = ds.AuthUser(userCode, userName); err != nil {
		opt.AppLog.Error(fmt.Sprintf("Error authenticating user: %v", err))
		return ctx, http.StatusUnauthorized
	}
	if opt.Request.Method == http.MethodPost || opt.Request.Method == http.MethodPut || opt.Request.Method == http.MethodDelete {
		if user.UserGroup == md.UserGroupGuest && strings.Contains(opt.Request.URL.Path, st.ApiPath) {
			return ctx, http.StatusMethodNotAllowed
		}
	}

	ctx = context.WithValue(opt.Request.Context(), md.DataStoreCtxKey, ds)
	ctx = context.WithValue(ctx, md.AuthUserCtxKey, user)

	return ctx, errCode
}

func Database(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	var option cu.IM
	if err := ds.ConvertFromReader(r.Body, &option); err != nil {
		RespondMessage(w, http.StatusUnprocessableEntity, nil, http.StatusUnprocessableEntity, err)
		return
	}

	response := api.CreateDatabase(option, ds.Config)
	RespondMessage(w, http.StatusOK, response, http.StatusUnprocessableEntity, nil)
}

func Function(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	var option cu.IM
	if err := ds.ConvertFromReader(r.Body, &option); err != nil {
		RespondMessage(w, http.StatusUnprocessableEntity, nil, http.StatusUnprocessableEntity, err)
		return
	}

	prms := cu.ToIM(option["values"], cu.IM{})
	response, err := ds.Function(cu.ToString(option["name"], ""), prms)
	if err == nil && cu.ToString(cu.ToIM(response, cu.IM{})["content_type"], "") == "application/pdf" {
		if pdf, found := cu.ToIM(response, cu.IM{})["template"].([]uint8); found {
			w.Header().Set("Content-Type", "application/pdf")
			w.Header().Set("Content-Disposition", "attachment; filename=report.pdf")
			w.Write(pdf)
			return
		}
	}
	RespondMessage(w, http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

func View(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)

	var view md.View
	var response []cu.IM
	var err error
	if err = ds.ConvertFromReader(r.Body, &view); err == nil {
		query := md.Query{
			Fields:  []string{"*"},
			From:    strings.ToLower(strings.TrimPrefix(view.Name.String(), "VIEW_")),
			Filter:  view.Filter,
			OrderBy: view.OrderBy,
			Limit:   view.Limit,
			Offset:  view.Offset,
		}
		response, err = ds.StoreDataQuery(query, false)
	}

	RespondMessage(w, http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

func Health(w http.ResponseWriter, r *http.Request) {
	response := cu.IM{
		"status": "ok",
		"time":   time.Now().Format(time.RFC3339),
	}
	RespondMessage(w, http.StatusOK, response, http.StatusInternalServerError, nil)
}
