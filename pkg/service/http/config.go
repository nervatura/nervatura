package http

import (
	"errors"
	"net/http"
	"slices"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
)

var configMap = map[md.ConfigType]any{
	md.ConfigTypeMap:        &md.ConfigMap{},
	md.ConfigTypeShortcut:   &md.ConfigShortcut{},
	md.ConfigTypeMessage:    &md.ConfigMessage{},
	md.ConfigTypePattern:    &md.ConfigPattern{},
	md.ConfigTypeReport:     &md.ConfigReport{},
	md.ConfigTypePrintQueue: &md.ConfigPrintQueue{},
	md.ConfigTypeData:       &cu.IM{},
}

// ConfigPost - create new config
func ConfigPost(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	user := r.Context().Value(md.AuthUserCtxKey).(md.Auth)

	// convert request body to struct and schema validation
	var data md.Config = md.Config{
		ConfigType: md.ConfigType(md.ConfigTypeMap),
	}
	err := ds.ConvertFromReader(r.Body, &data)
	if err != nil {
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
	}

	if user.UserGroup != md.UserGroupAdmin &&
		!slices.Contains([]string{"CONFIG_PRINT_QUEUE", "CONFIG_PATTERN"}, data.ConfigType.String()) {
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, errors.New(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}

	// prepare values for database update
	values := cu.IM{
		"config_type": data.ConfigType.String(),
	}
	if data.Code != "" {
		values["code"] = data.Code
	}
	configDataByte, err := ds.ConvertToByte(data.Data)
	if err == nil {
		// schema validation
		err = ds.ConvertFromByte(configDataByte, configMap[data.ConfigType])
		if err == nil {
			values["data"] = string(configDataByte[:])
		}
	}
	if err != nil {
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
	}

	// database insert
	var rows []cu.IM
	var result cu.IM
	var configID int64
	if configID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "config"}); err == nil {
		if rows, err = ds.StoreDataGet(cu.IM{"id": configID, "model": "config"}, true); err == nil {
			result = rows[0]
		}
	}
	RespondMessage(w, http.StatusCreated, result, http.StatusUnprocessableEntity, err)
}

// ConfigPut - update config
func ConfigPut(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	user := r.Context().Value(md.AuthUserCtxKey).(md.Auth)

	configID := cu.ToInteger(r.PathValue("id_code"), 0)
	configCode := cu.ToString(r.PathValue("id_code"), "")

	var config md.Config
	var inputFields, metaFields []string
	var data cu.IM
	var err error

	if data, inputFields, metaFields, err = ds.GetBodyData("config", r.Body, &config); err == nil {
		if cu.ToString(data["config_type"], "") == "" {
			err = errors.New("config_type is required")
			RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
			return
		}
		if user.UserGroup != md.UserGroupAdmin &&
			!slices.Contains([]string{"CONFIG_PRINT_QUEUE", "CONFIG_PATTERN"}, cu.ToString(data["config_type"], "")) {
			RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, errors.New(http.StatusText(http.StatusMethodNotAllowed)))
			return
		}

		if slices.Contains(inputFields, "data") {
			var configDataByte []byte
			configDataByte, err = ds.ConvertToByte(config.Data)
			if err == nil {
				// schema validation
				err = ds.ConvertFromByte(configDataByte, configMap[config.ConfigType])
				if err == nil {
					_, err = ds.UpdateData(md.UpdateDataOptions{
						Model: "config", IDKey: configID, Code: configCode,
						Data: config, Fields: inputFields, MetaFields: metaFields,
					})
				}
			}
		}
	}
	RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// ConfigDelete - delete config
func ConfigDelete(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	user := r.Context().Value(md.AuthUserCtxKey).(md.Auth)
	configID := cu.ToInteger(r.PathValue("id_code"), 0)
	configCode := cu.ToString(r.PathValue("id_code"), "")

	var configs []cu.IM
	var err error
	if configs, err = ds.GetDataByID("config", configID, configCode, true); err == nil {
		config := configs[0]
		if user.UserGroup != md.UserGroupAdmin &&
			!slices.Contains([]string{"CONFIG_PRINT_QUEUE", "CONFIG_PATTERN"}, cu.ToString(config["config_type"], "")) {
			RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, errors.New(http.StatusText(http.StatusMethodNotAllowed)))
			return
		}
		_, err = ds.StoreDataUpdate(md.Update{Model: "config", IDKey: cu.ToInteger(config["id"], 0)})
	}

	RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// ConfigQuery - get configs
func ConfigQuery(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)

	var params cu.IM = cu.IM{
		"model":  "config",
		"limit":  cu.ToInteger(r.URL.Query().Get("limit"), 0),
		"offset": cu.ToInteger(r.URL.Query().Get("offset"), 0),
	}
	if r.URL.Query().Get("config_type") != "" {
		params["config_type"] = cu.ToString(r.URL.Query().Get("config_type"), "")
	}
	response, err := ds.StoreDataGet(params, false)
	RespondMessage(w, http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// ConfigGet - get config by id or code
func ConfigGet(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	configID := cu.ToInteger(r.PathValue("id_code"), 0)
	configCode := cu.ToString(r.PathValue("id_code"), "")
	var err error
	var configs []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if configs, err = ds.GetDataByID("config", configID, configCode, true); err == nil {
		response = configs[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	RespondMessage(w, http.StatusOK, response, errStatus, err)
}
