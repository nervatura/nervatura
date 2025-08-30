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

// ToolPost - create new tool
func (cli *CLIService) ToolInsert(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	// convert request body to struct and schema validation
	var data md.Tool = md.Tool{
		Events: []md.Event{},
		ToolMeta: md.ToolMeta{
			Tags: []string{},
		},
		ToolMap: cu.IM{},
	}
	err := ds.ConvertFromByte([]byte(requestData), &data)
	if err != nil {
		return cli.respondString(http.StatusUnprocessableEntity, nil, http.StatusUnprocessableEntity, err)
	}

	if data.ProductCode == "" || data.Description == "" {
		err = errors.New("product code and description are required")
		return cli.respondString(http.StatusUnprocessableEntity, nil, http.StatusUnprocessableEntity, err)
	}

	// prepare values for database update
	values := cu.IM{
		"product_code": data.ProductCode,
		"description":  data.Description,
	}
	if data.Code != "" {
		values["code"] = data.Code
	}

	ut.ConvertByteToIMData(data.Events, values, "events")
	ut.ConvertByteToIMData(data.ToolMeta, values, "tool_meta")
	ut.ConvertByteToIMData(data.ToolMap, values, "tool_map")

	// database insert
	var rows []cu.IM
	var result cu.IM
	var toolID int64
	if toolID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "tool"}); err == nil {
		if rows, err = ds.StoreDataGet(cu.IM{"id": toolID, "model": "tool"}, true); err == nil {
			result = rows[0]
		}
	}
	return cli.respondString(http.StatusCreated, result, http.StatusUnprocessableEntity, err)
}

// ToolPut - update tool
func (cli *CLIService) ToolUpdate(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	toolID := cu.ToInteger(options["id"], 0)
	toolCode := cu.ToString(options["code"], "")

	var tool md.Tool
	var inputFields, metaFields []string
	var err error
	reader := io.NopCloser(bytes.NewReader([]byte(requestData)))
	if _, inputFields, metaFields, err = ds.GetBodyData("tool", reader, &tool); err == nil {
		err = ds.UpdateData(md.UpdateDataOptions{
			Model: "tool", IDKey: toolID, Code: toolCode,
			Data: tool, Meta: tool.ToolMeta, Fields: inputFields, MetaFields: metaFields,
		})
	}
	return cli.respondString(http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

func (cli *CLIService) ToolDelete(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	toolID := cu.ToInteger(options["id"], 0)
	toolCode := cu.ToString(options["code"], "")
	err := ds.DataDelete("tool", toolID, toolCode)
	return cli.respondString(http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// ToolQuery - get tools
func (cli *CLIService) ToolQuery(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	var params cu.IM = cu.IM{"model": "tool"}
	for _, v := range []string{"product_code", "description", "tag"} {
		if options[v] != nil {
			params[v] = options[v]
		}
	}
	response, err := ds.StoreDataGet(params, false)
	return cli.respondString(http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// ToolGet - get tool by id or code
func (cli *CLIService) ToolGet(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	toolID := cu.ToInteger(options["id"], 0)
	toolCode := cu.ToString(options["code"], "")
	var err error
	var tools []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if tools, err = ds.GetDataByID("tool", toolID, toolCode, true); err == nil {
		response = tools[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	return cli.respondString(http.StatusOK, response, errStatus, err)
}
