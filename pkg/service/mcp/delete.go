package mcp

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

func createDeleteTool(name, model string, meta mcp.Meta) (tool mcp.Tool) {
	return mcp.Tool{
		Name:        name,
		Title:       fmt.Sprintf("Delete %s data by code", model),
		Description: fmt.Sprintf("Delete data by %s model code. It returns the success of the operation or an error message.", model),
		Meta:        meta,
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"code": {Type: "string", MinLength: ut.AnyPointer(12),
					Description: "The unique key of the result model data. Example: CUS1731101982N123",
					Examples:    []any{`CUS1731101982N123`, `PRD1731101982N123`}},
			},
			Required: []string{"code"},
		},
		OutputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"success": {Type: "boolean", Description: "The success of the operation."},
			},
		},
	}
}

func createExtendDeleteTool(name, model string, meta mcp.Meta) (tool mcp.Tool) {
	return mcp.Tool{
		Name:        name,
		Title:       fmt.Sprintf("Delete %s data by code", model),
		Description: fmt.Sprintf("Delete data by %s model code. It returns the success of the operation or an error message.", model),
		Meta:        meta,
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"code": {Type: "string", MinLength: ut.AnyPointer(12),
					Description: "The unique key of the result model data. Example: CUS1731101982N123",
					Examples:    []any{`CUS1731101982N123`, `PRD1731101982N123`}},
				"index": {Type: "integer", Description: "The index of the data to delete."},
			},
			Required: []string{"code", "index"},
		},
		OutputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"success": {Type: "boolean", Description: "The success of the operation."},
			},
		},
	}
}

func modelDelete(ctx context.Context, req *mcp.CallToolRequest, inputData map[string]any) (result *mcp.CallToolResult, response any, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)

	code := cu.ToString(inputData["code"], "???")
	prefix := code[:3]

	var sm *ModelSchema
	if sm, err = getModelSchemaByPrefix(prefix); err != nil {
		return nil, nil, err
	}

	var res *mcp.ElicitResult
	if res, err = req.Session.Elicit(ctx, &mcp.ElicitParams{
		RequestedSchema: &jsonschema.Schema{Type: "object",
			Description: "Confirm the deletion of the data.",
			Properties: map[string]*jsonschema.Schema{
				"confirm": {Type: "string",
					Description: "Confirm the deletion of the data.",
					Enum:        []any{`YES`, `NO`}},
			},
			Required: []string{"confirm"},
		},
	}); err != nil {
		return nil, nil, fmt.Errorf("eliciting failed: %v", err)
	}
	confirm := (cu.ToString(res.Content["confirm"], "NO") == "YES")

	if confirm {
		err = ds.DataDelete(sm.Name, 0, code)
	}

	return &mcp.CallToolResult{
		StructuredContent: cu.IM{"success": err == nil && confirm},
		IsError:           err != nil,
	}, nil, err
}

func extendDelete(ctx context.Context, req *mcp.CallToolRequest, inputData map[string]any) (result *mcp.CallToolResult, response any, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)

	code := cu.ToString(inputData["code"], "???")
	index := int(cu.ToInteger(inputData["index"], 0))

	var mt McpTool
	var found bool
	if mt, found = toolDataMap[req.Params.Name]; !found || !mt.Extend {
		return nil, nil, fmt.Errorf("invalid tool: %s", req.Params.Name)
	}
	var ms *ModelExtendSchema = mt.ModelExtendSchema

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

	var res *mcp.ElicitResult
	if res, err = req.Session.Elicit(ctx, &mcp.ElicitParams{
		RequestedSchema: &jsonschema.Schema{Type: "object",
			Description: "Confirm the deletion of the data.",
			Properties: map[string]*jsonschema.Schema{
				"confirm": {Type: "string",
					Description: "Confirm the deletion of the data.",
					Enum:        []any{`YES`, `NO`}},
			},
			Required: []string{"confirm"},
		},
	}); err != nil {
		return nil, nil, fmt.Errorf("eliciting failed: %v", err)
	}
	confirm := (cu.ToString(res.Content["confirm"], "NO") == "YES")

	if confirm {
		fieldValues = append(fieldValues[:index], fieldValues[index+1:]...)

		var modelData any
		if modelData, err = ms.LoadData(fieldValues); err == nil {
			values := cu.IM{}
			ut.ConvertByteToIMData(modelData, values, fieldName)
			_, err = ds.StoreDataUpdate(md.Update{Values: values, Model: baseModel, IDKey: updateID})
		}
	}

	return &mcp.CallToolResult{
		StructuredContent: cu.IM{"success": err == nil && confirm},
		IsError:           err != nil,
	}, nil, err
}
