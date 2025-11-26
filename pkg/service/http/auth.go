package http

import (
	"errors"
	"net/http"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

// AuthPost - create new auth
func AuthPost(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	user := r.Context().Value(md.AuthUserCtxKey).(md.Auth)
	if user.UserGroup != md.UserGroupAdmin {
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, errors.New(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}

	// convert request body to struct and schema validation
	var data md.Auth = md.Auth{
		UserGroup: md.UserGroup(md.UserGroupUser),
		AuthMeta: md.AuthMeta{
			Tags:      []string{},
			Bookmarks: []md.Bookmark{},
		},
		AuthMap: cu.IM{},
	}
	err := ds.ConvertFromReader(r.Body, &data)
	if err != nil {
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
	}

	if data.UserName == "" {
		err = errors.New("auth user_name and user_group are required")
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
	}

	// prepare values for database update
	values := cu.IM{
		"user_name":  data.UserName,
		"user_group": data.UserGroup.String(),
	}
	if data.Code != "" {
		values["code"] = data.Code
	}

	ut.ConvertByteToIMData(data.AuthMeta, values, "auth_meta")
	ut.ConvertByteToIMData(data.AuthMap, values, "auth_map")

	// database insert
	var rows []cu.IM
	var result cu.IM
	var authID int64
	if authID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "auth"}); err == nil {
		if rows, err = ds.StoreDataGet(cu.IM{"id": authID, "model": "auth"}, true); err == nil {
			result = rows[0]
		}
	}
	RespondMessage(w, http.StatusCreated, result, http.StatusUnprocessableEntity, err)
}

// AuthPut - update auth
func AuthPut(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	user := r.Context().Value(md.AuthUserCtxKey).(md.Auth)

	authID := cu.ToInteger(r.PathValue("id_code"), 0)
	authCode := cu.ToString(r.PathValue("id_code"), "")
	if user.UserGroup != md.UserGroupAdmin && (authID != user.Id || (authID == 0 && authCode != user.Code)) {
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, errors.New(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}

	var auth md.Auth
	var inputFields, metaFields []string
	var err error
	if _, inputFields, metaFields, err = ds.GetBodyData("auth", r.Body, &auth); err == nil {
		_, err = ds.UpdateData(md.UpdateDataOptions{
			Model: "auth", IDKey: authID, Code: authCode,
			Data: auth, Meta: auth.AuthMeta, Fields: inputFields, MetaFields: metaFields,
		})
	}
	RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// AuthQuery - get auths
func AuthQuery(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	user := r.Context().Value(md.AuthUserCtxKey).(md.Auth)

	if user.UserGroup != md.UserGroupAdmin {
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, errors.New(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}

	var params cu.IM = cu.IM{
		"model":  "auth",
		"limit":  cu.ToInteger(r.URL.Query().Get("limit"), 0),
		"offset": cu.ToInteger(r.URL.Query().Get("offset"), 0),
	}
	if r.URL.Query().Get("user_group") != "" {
		params["user_group"] = cu.ToString(r.URL.Query().Get("user_group"), "")
	}
	if r.URL.Query().Get("tag") != "" {
		params["tag"] = cu.ToString(r.URL.Query().Get("tag"), "")
	}
	response, err := ds.StoreDataGet(params, false)
	RespondMessage(w, http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// AuthGet - get auth by id or code
func AuthGet(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	user := r.Context().Value(md.AuthUserCtxKey).(md.Auth)

	authID := cu.ToInteger(r.PathValue("id_code"), 0)
	authCode := cu.ToString(r.PathValue("id_code"), "")
	if user.UserGroup != md.UserGroupAdmin || (authID == 0 && authCode == "") {
		authID = user.Id
		authCode = user.Code
	}

	var err error
	var auths []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if auths, err = ds.GetDataByID("auth", authID, authCode, true); err == nil {
		response = auths[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	RespondMessage(w, http.StatusOK, response, errStatus, err)
}

// AuthLogin - Login by username and password and get access token
func AuthLogin(w http.ResponseWriter, r *http.Request) {
	opt := r.Context().Value(md.AuthOptionsCtxKey).(md.AuthOptions)

	var loginData cu.IM
	if err := opt.ConvertFromReader(r.Body, &loginData); err != nil {
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
	}

	ds := api.NewDataStore(opt.Config, cu.ToString(loginData["database"], ""), opt.AppLog)
	token, err := ds.UserLogin(cu.ToString(loginData["user_name"], ""), cu.ToString(loginData["password"], ""), true)
	result := cu.IM{"token": token, "version": cu.ToString(opt.Config["version"], "")}
	RespondMessage(w, http.StatusOK, result, http.StatusUnprocessableEntity, err)
}

func AuthPassword(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	user := r.Context().Value(md.AuthUserCtxKey).(md.Auth)

	var passwordData cu.IM
	if err := ds.ConvertFromReader(r.Body, &passwordData); err != nil {
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
	}
	err := ds.UserPassword(user.Code, cu.ToString(passwordData["password"], ""), cu.ToString(passwordData["confirm"], ""))
	RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// AuthReset - Reset a user password and result a new password
// If user is not admin, the user id and code are set to the current token user id or code
func AuthReset(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	user := r.Context().Value(md.AuthUserCtxKey).(md.Auth)

	authID := cu.ToInteger(r.PathValue("id_code"), 0)
	authCode := cu.ToString(r.PathValue("id_code"), "")
	if user.UserGroup != md.UserGroupAdmin && (authID != user.Id || (authID == 0 && authCode != user.Code)) {
		authID = user.Id
		authCode = user.Code
	}

	var auths []cu.IM
	var err error
	password := cu.RandString(16)
	if auths, err = ds.GetDataByID("auth", authID, authCode, true); err == nil {
		err = ds.UserPassword(cu.ToString(auths[0]["code"], ""), password, password)
	}
	RespondMessage(w, http.StatusOK, cu.IM{"password": password}, http.StatusUnprocessableEntity, err)
}

// AuthToken - Refresh access token
func AuthToken(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	user := r.Context().Value(md.AuthUserCtxKey).(md.Auth)
	token, err := ds.TokenRefresh(user)
	result := cu.IM{"token": token, "version": cu.ToString(ds.Config["version"], "")}
	RespondMessage(w, http.StatusOK, result, http.StatusUnprocessableEntity, err)
}
