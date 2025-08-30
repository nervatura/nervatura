package cli

import (
	"bytes"
	"errors"
	"io"
	"net/http"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

// AuthInsert - create new auth
func (cli *CLIService) AuthInsert(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	// convert request body to struct and schema validation
	var data md.Auth = md.Auth{
		UserGroup: md.UserGroup(md.UserGroupUser),
		AuthMeta: md.AuthMeta{
			Tags:      []string{},
			Bookmarks: []md.Bookmark{},
		},
		AuthMap: cu.IM{},
	}
	err := ds.ConvertFromByte([]byte(requestData), &data)
	if err != nil {
		return cli.respondString(http.StatusUnprocessableEntity, nil, http.StatusUnprocessableEntity, err)
	}

	if data.UserName == "" {
		err = errors.New("auth user_name and user_group are required")
		return cli.respondString(http.StatusUnprocessableEntity, nil, http.StatusUnprocessableEntity, err)
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
	return cli.respondString(http.StatusCreated, result, http.StatusUnprocessableEntity, err)
}

// AuthUpdate - update auth
func (cli *CLIService) AuthUpdate(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	authID := cu.ToInteger(options["id"], 0)
	authCode := cu.ToString(options["code"], "")

	var auth md.Auth
	var inputFields, metaFields []string
	var err error
	reader := io.NopCloser(bytes.NewReader([]byte(requestData)))
	if _, inputFields, metaFields, err = ds.GetBodyData("auth", reader, &auth); err == nil {
		err = ds.UpdateData(md.UpdateDataOptions{
			Model: "auth", IDKey: authID, Code: authCode,
			Data: auth, Meta: auth.AuthMeta, Fields: inputFields, MetaFields: metaFields,
		})
	}
	return cli.respondString(http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// AuthQuery - query auths
func (cli *CLIService) AuthQuery(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	var params cu.IM = cu.IM{"model": "auth"}
	if userGroup, found := options["user_group"].(string); found {
		params["user_group"] = userGroup
	}
	if tag, found := options["tag"].(string); found {
		params["tag"] = tag
	}
	response, err := ds.StoreDataGet(params, false)
	return cli.respondString(http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// AuthGet - get auth
func (cli *CLIService) AuthGet(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	authID := cu.ToInteger(options["id"], 0)
	authCode := cu.ToString(options["code"], "")

	var err error
	var auths []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if auths, err = ds.GetDataByID("auth", authID, authCode, true); err == nil {
		response = auths[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	return cli.respondString(http.StatusOK, response, errStatus, err)
}
