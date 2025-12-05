package mcp

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

func modelQuery(ctx context.Context, req *mcp.CallToolRequest, parameters cu.IM) (result *mcp.CallToolResult, response any, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)

	model := strings.TrimSuffix(strings.TrimPrefix(req.Params.Name, "nervatura_"), "_query")
	var ms *ModelSchema
	var found bool
	if ms, found = getSchemaMap()[model]; !found {
		return nil, nil, fmt.Errorf("invalid model: %s", model)
	}

	var params cu.IM = cu.IM{
		"fields": []string{"*"},
		"model":  ms.Name,
		"limit":  cu.ToInteger(parameters["limit"], 0),
		"offset": cu.ToInteger(parameters["offset"], 0),
	}
	for key, value := range parameters {
		if !slices.Contains([]string{"limit", "offset"}, key) {
			params[key] = cu.ToString(value, "")
		}
	}

	content := cu.IM{"items": []any{}}
	var rows []cu.IM
	if rows, err = ds.StoreDataGet(params, true); err == nil {
		content["items"], err = ms.LoadList(rows)
	}

	return &mcp.CallToolResult{
		StructuredContent: content,
		IsError:           err != nil,
	}, nil, err
}

func modelUpdate(ctx context.Context, req *mcp.CallToolRequest, inputData cu.IM) (result *mcp.CallToolResult, response UpdateResponseData, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)

	var updateID int64
	var inputFields, metaFields []string
	var modelData, metaData any

	code := cu.ToString(inputData["code"], "")
	if code == "" {
		return nil, UpdateResponseData{}, fmt.Errorf("code is required")
	}
	model := strings.TrimSuffix(strings.TrimPrefix(req.Params.Name, "nervatura_"), "_update")
	var ms *ModelSchema
	var found bool
	if ms, found = getSchemaMap()[model]; !found {
		return nil, UpdateResponseData{}, fmt.Errorf("invalid model: %s", model)
	}

	if modelData, metaData, inputFields, metaFields, err = getSchemaData(inputData, ms, getParamsMeta(req)); err == nil {
		updateID, err = ds.UpdateData(md.UpdateDataOptions{
			Model: ms.Name, IDKey: 0, Code: code,
			Data: modelData, Meta: metaData, Fields: inputFields, MetaFields: metaFields,
		})
	}
	response = UpdateResponseData{
		Model: ms.Name,
		Code:  code,
		ID:    updateID,
	}

	return result, response, err
}

func extendQuery(ctx context.Context, req *mcp.CallToolRequest, inputData cu.IM) (result *mcp.CallToolResult, response any, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	baseModel := cu.ToString(inputData["model"], "")

	model := strings.TrimSuffix(strings.TrimPrefix(req.Params.Name, "nervatura_"), "_query")
	var ms *ModelExtendSchema
	var found bool
	if ms, found = getExtendSchemaMap()[model]; !found {
		return nil, nil, fmt.Errorf("invalid model: %s", model)
	}

	var params cu.IM = cu.IM{
		"fields": []string{"*"},
		"model":  ms.ViewName[baseModel],
		"limit":  cu.ToInteger(inputData["limit"], 0),
		"offset": cu.ToInteger(inputData["offset"], 0),
	}
	for key, value := range inputData {
		if !slices.Contains([]string{"model", "limit", "offset"}, key) {
			params[key] = cu.ToString(value, "")
		}
	}

	content := cu.IM{"items": []any{}}
	var rows []cu.IM
	if rows, err = ds.StoreDataGet(params, false); err == nil {
		content["items"], err = ms.LoadList(baseModel, rows)
	}

	return &mcp.CallToolResult{
		StructuredContent: content,
		IsError:           err != nil,
	}, nil, err
}

func extendUpdate(ctx context.Context, req *mcp.CallToolRequest, inputData cu.IM) (result *mcp.CallToolResult, response UpdateResponseData, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)

	code := cu.ToString(inputData["code"], "")
	index := int(cu.ToInteger(inputData["index"], 0))
	if code == "" {
		return nil, UpdateResponseData{}, fmt.Errorf("code is required")
	}

	model := strings.TrimSuffix(strings.TrimPrefix(req.Params.Name, "nervatura_"), "_update")
	var ms *ModelExtendSchema
	var found bool
	if ms, found = getExtendSchemaMap()[model]; !found {
		return nil, UpdateResponseData{}, fmt.Errorf("invalid model: %s", model)
	}

	var baseModel, fieldName string
	if baseModel, fieldName, err = ms.ModelFromCode(code); err != nil {
		return nil, UpdateResponseData{}, fmt.Errorf("invalid code: %s", code)
	}

	var rows []cu.IM
	if rows, err = ds.StoreDataGet(cu.IM{
		"fields": []string{"*"}, "model": baseModel, "code": code}, true); err != nil {
		return nil, UpdateResponseData{}, errors.New("invalid code: " + code)
	}
	updateID := cu.ToInteger(rows[0]["id"], 0)
	fieldValues := cu.ToIMA(rows[0][fieldName], []cu.IM{})
	if len(fieldValues) < index+1 {
		return nil, UpdateResponseData{}, fmt.Errorf("index out of range: %d", index)
	}
	for field, value := range inputData {
		if _, found := fieldValues[index][field]; found {
			fieldValues[index][field] = value
		}
	}
	mapValues := getParamsMeta(req)
	if len(mapValues) > 0 {
		fieldValues[index][model+"_map"] = cu.MergeIM(cu.ToIM(fieldValues[index][model+"_map"], cu.IM{}), mapValues)
	}
	var modelData any
	if modelData, err = ms.LoadData(fieldValues); err != nil {
		return nil, UpdateResponseData{}, fmt.Errorf("invalid data: %s", err.Error())
	}
	values := cu.IM{}
	ut.ConvertByteToIMData(modelData, values, fieldName)
	_, err = ds.StoreDataUpdate(md.Update{Values: values, Model: baseModel, IDKey: updateID})

	return result, UpdateResponseData{
		Model: baseModel,
		Code:  code,
		ID:    updateID,
	}, err
}

func extendCreate(ctx context.Context, req *mcp.CallToolRequest, inputData cu.IM) (result *mcp.CallToolResult, response UpdateResponseData, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)

	code := cu.ToString(inputData["code"], "")
	if code == "" {
		return nil, UpdateResponseData{}, fmt.Errorf("code is required")
	}

	model := strings.TrimSuffix(strings.TrimPrefix(req.Params.Name, "nervatura_"), "_create")
	var ms *ModelExtendSchema
	var found bool
	if ms, found = getExtendSchemaMap()[model]; !found {
		return nil, UpdateResponseData{}, fmt.Errorf("invalid model: %s", model)
	}

	var baseModel, fieldName string
	if baseModel, fieldName, err = ms.ModelFromCode(code); err != nil {
		return nil, UpdateResponseData{}, fmt.Errorf("invalid code: %s", code)
	}

	var rows []cu.IM
	if rows, err = ds.StoreDataGet(cu.IM{
		"fields": []string{"*"}, "model": baseModel, "code": code}, true); err != nil {
		return nil, UpdateResponseData{}, errors.New("invalid code: " + code)
	}
	updateID := cu.ToInteger(rows[0]["id"], 0)
	fieldValues := cu.ToIMA(rows[0][fieldName], []cu.IM{})

	if _, found := inputData["tags"]; !found {
		inputData["tags"] = []string{}
	}
	inputData[model+"_map"] = getParamsMeta(req)
	fieldValues = append(fieldValues, inputData)

	var modelData any
	if modelData, err = ms.LoadData(fieldValues); err != nil {
		return nil, UpdateResponseData{}, fmt.Errorf("invalid data: %s", err.Error())
	}
	values := cu.IM{}
	ut.ConvertByteToIMData(modelData, values, fieldName)
	_, err = ds.StoreDataUpdate(md.Update{Values: values, Model: baseModel, IDKey: updateID})

	response = UpdateResponseData{
		Model: baseModel,
		Code:  code,
		ID:    updateID,
	}

	return result, response, err
}
