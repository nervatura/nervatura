package cli

import (
	"bytes"
	"errors"
	"io"
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

// ConfigInsert - create new config
func (cli *CLIService) ConfigInsert(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	// convert request body to struct and schema validation
	var data md.Config = md.Config{
		ConfigType: md.ConfigType(md.ConfigTypeMap),
	}
	err := ds.ConvertFromByte([]byte(requestData), &data)
	if err != nil {
		return cli.respondString(0, nil, http.StatusUnprocessableEntity, err)
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
		return cli.respondString(0, nil, http.StatusUnprocessableEntity, err)
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
	return cli.respondString(http.StatusCreated, result, http.StatusUnprocessableEntity, err)
}

// ConfigUpdate - update config
func (cli *CLIService) ConfigUpdate(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	configID := cu.ToInteger(options["id"], 0)
	configCode := cu.ToString(options["code"], "")

	var config md.Config
	var inputFields, metaFields []string
	var data cu.IM
	var err error

	reader := io.NopCloser(bytes.NewReader([]byte(requestData)))
	if data, inputFields, metaFields, err = ds.GetBodyData("config", reader, &config); err == nil {
		if cu.ToString(data["config_type"], "") == "" {
			err = errors.New("config_type is required")
			return cli.respondString(0, nil, http.StatusUnprocessableEntity, err)
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
	return cli.respondString(http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// ConfigDelete - delete config
func (cli *CLIService) ConfigDelete(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	configID := cu.ToInteger(options["id"], 0)
	configCode := cu.ToString(options["code"], "")

	var configs []cu.IM
	var err error
	if configs, err = ds.GetDataByID("config", configID, configCode, true); err == nil {
		config := configs[0]
		_, err = ds.StoreDataUpdate(md.Update{Model: "config", IDKey: cu.ToInteger(config["id"], 0)})
	}

	return cli.respondString(http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// ConfigQuery - get configs
func (cli *CLIService) ConfigQuery(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	var params cu.IM = cu.IM{"model": "config"}
	if configType, found := options["config_type"].(string); found {
		params["config_type"] = configType
	}
	response, err := ds.StoreDataGet(params, false)
	return cli.respondString(http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// ConfigGet - get config by id or code
func (cli *CLIService) ConfigGet(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	configID := cu.ToInteger(options["id"], 0)
	configCode := cu.ToString(options["code"], "")
	var err error
	var configs []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if configs, err = ds.GetDataByID("config", configID, configCode, true); err == nil {
		response = configs[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	return cli.respondString(http.StatusOK, response, errStatus, err)
}
