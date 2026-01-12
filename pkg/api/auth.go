package api

import (
	"errors"
	"fmt"
	"net/http"

	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
	st "github.com/nervatura/nervatura/v6/pkg/static"
)

/*
TokenRefresh - create/refresh a JWT token
*/
func (ds *DataStore) TokenRefresh(user md.Auth) (token string, err error) {
	if token, err = ds.CreateLoginToken(cu.SM{"code": user.Code, "user_name": user.UserName, "scope": user.UserGroup.String(), "alias": ds.Alias}, ds.Config); err != nil {
		err = ds.SetError(err, errors.New(http.StatusText(http.StatusInternalServerError)))
	}
	return token, err
}

func (ds *DataStore) AuthUser(uid, username string) (user md.Auth, err error) {
	query := md.Query{
		Fields: []string{"*"}, From: "auth",
		Filters: []md.Filter{
			{Field: "deleted", Comp: "==", Value: false},
			{Field: "disabled", Comp: "==", Value: false},
		},
		Filter: fmt.Sprintf("and (code='%s' or user_name='%s')", uid, username),
	}
	var rows []cu.IM
	if rows, err = ds.StoreDataQuery(query, true); err == nil {
		err = ds.ConvertData(rows[0], &user)
	}
	return user, err
}

func (ds *DataStore) UserLogin(username, password string, createToken bool) (token string, err error) {

	result := func(user md.Auth) (string, error) {
		// create a token
		if createToken {
			return ds.TokenRefresh(user)
		}
		return token, nil
	}

	// check username or email
	var user md.Auth
	if user, err = ds.AuthUser("", username); err != nil {
		return token, err
	}

	// check password
	var pwhash string
	refnumber := ut.GetHash(cu.ToString(user.Code, ""), "sha256")
	query := md.Query{
		Fields: []string{"value"}, From: "usref",
		Filters: []md.Filter{
			{Field: "refnumber", Comp: "==", Value: refnumber},
		},
	}
	if rows, err := ds.StoreDataQuery(query, false); (err == nil) && (len(rows) > 0) {
		pwhash = cu.ToString(rows[0]["value"], "")
	}

	if pwhash == "" {
		if user.Code == st.DefaultConfig["connection"]["default_admin"] && user.UserGroup == md.UserGroupAdmin {
			if cu.ToString(ds.Config["NT_API_KEY"], "") != password {
				return token, errors.New(ut.GetMessage("default_admin_password"))
			}
			return result(user)
		}
		return token, errors.New(ut.GetMessage("missing_password"))
	}

	if err := ds.ComparePasswordAndHash(password, pwhash); err != nil {
		return token, errors.New(ut.GetMessage("wrong_password"))
	}
	return result(user)
}

func (ds *DataStore) TokenLogin(token string) (user md.Auth, err error) {
	if token == "" {
		return user, errors.New("token is empty")
	}
	tokenKeys := ds.Config["tokenKeys"].([]cu.SM)
	claim, err := ds.ParseToken(token, tokenKeys, ds.Config)
	if err != nil {
		return user, err
	}
	return ds.AuthUser(cu.ToString(claim["user_code"], ""), cu.ToString(claim["user_name"], ""))
}

func (ds *DataStore) UserPassword(userCode, password, confirm string) (err error) {
	if userCode == "" {
		return errors.New("missing user code")
	}
	if password == "" {
		return errors.New(ut.GetMessage("empty_password"))
	}
	if password != confirm {
		return errors.New(ut.GetMessage("verify_password"))
	}

	var pwhash string
	pwhash, err = ds.CreatePasswordHash(password)
	if err != nil {
		return err
	}
	update := md.Update{Model: "usref", Values: cu.IM{"value": pwhash}}

	// check password
	refnumber := ut.GetHash(userCode, "sha256")
	query := md.Query{
		Fields: []string{"id"}, From: "usref",
		Filters: []md.Filter{
			{Field: "refnumber", Comp: "==", Value: refnumber},
		},
	}
	var rows []cu.IM
	if rows, err = ds.StoreDataQuery(query, false); err != nil {
		return err
	}
	if len(rows) > 0 {
		update.IDKey = cu.ToInteger(rows[0]["id"], 0)
	} else {
		update.Values["refnumber"] = refnumber
	}
	_, err = ds.StoreDataUpdate(update)
	return err
}
