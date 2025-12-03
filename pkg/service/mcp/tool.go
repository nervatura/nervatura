package mcp

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
)

func modelQuery(ctx context.Context, req *mcp.CallToolRequest, parameters cu.IM) (result *mcp.CallToolResult, response any, err error) {
	extra := req.GetExtra()
	ds := extra.TokenInfo.Extra["ds"].(*api.DataStore)

	model := strings.TrimSuffix(strings.TrimPrefix(req.Params.Name, "nervatura_"), "_query")
	var ms *ModelSchema
	var found bool
	if ms, found = getSchemaMap()[model]; !found {
		return nil, nil, fmt.Errorf("invalid model: %s", model)
	}

	var params cu.IM = cu.IM{
		"fields": []string{ms.Name + "_object"},
		"model":  ms.Name + "_view",
		"limit":  cu.ToInteger(parameters["limit"], 0),
		"offset": cu.ToInteger(parameters["offset"], 0),
	}
	for key, value := range parameters {
		if !slices.Contains([]string{"limit", "offset"}, key) {
			params[key] = cu.ToString(value, "")
		}
	}

	response = []any{}
	var rows []cu.IM
	if rows, err = ds.StoreDataGet(params, true); err == nil {
		for _, row := range rows {
			var result any = ms.ResultType()
			if dataJson, found := row[ms.Name+"_object"].(string); found {
				if err = ds.ConvertFromByte([]byte(dataJson), &result); err != nil {
					return nil, nil, err
				}
				response = append(response.([]any), result)
			}
		}
	}

	return &mcp.CallToolResult{
		StructuredContent: cu.IM{"items": response},
		IsError:           err != nil,
	}, nil, err
}

func modelUpdate(ctx context.Context, req *mcp.CallToolRequest, inputData cu.IM) (result *mcp.CallToolResult, response UpdateResponseData, err error) {
	extra := req.GetExtra()
	ds := extra.TokenInfo.Extra["ds"].(*api.DataStore)

	var updateID int64
	var inputFields, metaFields []string
	var modelData, metaData any

	code := cu.ToString(inputData["code"], "")
	model := strings.TrimSuffix(strings.TrimPrefix(req.Params.Name, "nervatura_"), "_update")
	var ms *ModelSchema
	var found bool
	if ms, found = getSchemaMap()[model]; !found {
		return nil, UpdateResponseData{}, fmt.Errorf("invalid model: %s", model)
	}

	if modelData, metaData, inputFields, metaFields, err = getSchemaData(inputData, ms); err == nil {
		if code != "" {
			updateID, err = ds.UpdateData(md.UpdateDataOptions{
				Model: ms.Name, IDKey: 0, Code: code,
				Data: modelData, Meta: metaData, Fields: inputFields, MetaFields: metaFields,
			})
		} else {
			// prepare values for database update
			values := ms.InsertValues(modelData)
			updateID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: ms.Name})
		}
	}
	if err == nil {
		response = UpdateResponseData{
			Model: ms.Name,
			Code:  code,
			ID:    updateID,
		}

		if code == "" {
			var rows []cu.IM
			if rows, err = ds.StoreDataGet(cu.IM{"fields": []string{"code"}, "id": updateID, "model": ms.Name}, true); err == nil {
				response.Code = cu.ToString(rows[0]["code"], "")
			}
		}

	}

	return result, response, err
}
