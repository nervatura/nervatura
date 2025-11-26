package mcp

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

var queryCodeTool = mcp.Tool{
	Name:        "nervatura_query_code",
	Title:       "Query model data by code",
	Description: "Query data by model code. It returns all the data related to the model as a result.",
	InputSchema: &jsonschema.Schema{
		Type: "object",
		Properties: map[string]*jsonschema.Schema{
			"code": {Type: "string", MinLength: ut.AnyPointer(12),
				Description: "The unique key of the result model data. Example: CUS1731101982N123",
				Examples:    []any{`CUS1731101982N123`, `PRD1731101982N123`}},
		},
	},
	OutputSchema: &jsonschema.Schema{
		Type:  "object",
		OneOf: makeModelSchemaList(SchemaTypeResult),
	},
}

var queryParametersTool = mcp.Tool{
	Name:        "nervatura_query_parameters",
	Title:       "Query model data by parameters",
	Description: "Query model data by parameters. The result is all data that matches the filter criteria.",
	InputSchema: &jsonschema.Schema{
		Type: "object",
		Properties: map[string]*jsonschema.Schema{
			"model": {Type: "string",
				Description: "Please select the data type.",
				Enum:        []any{`customer`, `product`},
			},
			"parameters": {Type: "array", Items: &jsonschema.Schema{Type: "string"},
				Description: "Parameters to filter the data. Example: customer_type=CUSTOMER_COMPANY",
				Examples:    []any{`customer_type=CUSTOMER_COMPANY`, `customer_name=John Doe`}},
		},
		Required: []string{"model"},
	},
	OutputSchema: &jsonschema.Schema{
		Type: "object",
		Items: &jsonschema.Schema{
			OneOf: makeModelSchemaList(SchemaTypeResultList),
		},
	},
}

var queryElicitingTool = mcp.Tool{
	Name:        "nervatura_query_eliciting",
	Title:       "Query model data by parameters",
	Description: "Query model data by parameters. After selecting the model, the server returns the possible conditions that can be specified for the query. The result is all data that matches the filter criteria.",
	InputSchema: &jsonschema.Schema{
		Type: "object",
		Properties: map[string]*jsonschema.Schema{
			"model": {Type: "string",
				Description: "Please select the data type.",
				Enum:        getModelEnum(SchemaTypeResult),
			},
		},
		Required: []string{"model"},
	},
	OutputSchema: &jsonschema.Schema{
		Type: "object",
		Items: &jsonschema.Schema{
			OneOf: makeModelSchemaList(SchemaTypeResultList),
		},
	},
}

func queryCode(ctx context.Context, req *mcp.CallToolRequest, inputData map[string]any) (result *mcp.CallToolResult, response any, err error) {
	extra := req.GetExtra()
	ds := extra.TokenInfo.Extra["ds"].(*api.DataStore)

	code := cu.ToString(inputData["code"], "XXX")
	prefix := code[:3]

	var sm *ModelSchema
	if sm, err = getModelSchemaByPrefix(prefix); err != nil {
		return nil, nil, err
	}
	response = sm.ResultType()

	var rows []cu.IM
	if rows, err = ds.StoreDataGet(
		cu.IM{"fields": []string{sm.Name + "_object"}, "code": code, "model": sm.Name + "_view"}, true); err == nil {
		if dataJson, found := rows[0][sm.Name+"_object"].(string); found {
			err = ds.ConvertFromByte([]byte(dataJson), &response)
		}
	}
	return &mcp.CallToolResult{
		StructuredContent: response,
		IsError:           err != nil,
	}, nil, err
}

func queryEliciting(ctx context.Context, req *mcp.CallToolRequest, inputData cu.IM) (result *mcp.CallToolResult, response any, err error) {
	extra := req.GetExtra()
	ds := extra.TokenInfo.Extra["ds"].(*api.DataStore)
	model := cu.ToString(inputData["model"], "")

	var sm *ModelSchema
	var found bool
	if sm, found = getSchemaMap()[model]; !found {
		return nil, nil, fmt.Errorf("invalid model: %s", model)
	}

	var res *mcp.ElicitResult
	if res, err = req.Session.Elicit(ctx, &mcp.ElicitParams{
		Message:         fmt.Sprintf("List %s(s) parameters", sm.Name),
		RequestedSchema: makeModelSchema(model, SchemaTypeParameter),
	}); err != nil {
		return nil, nil, fmt.Errorf("eliciting failed: %v", err)
	}

	var params cu.IM = cu.IM{
		"fields": []string{model + "_object"},
		"model":  model + "_view",
		"limit":  cu.ToInteger(res.Content["limit"], 0),
		"offset": cu.ToInteger(res.Content["offset"], 0),
	}
	for key, value := range res.Content {
		if !slices.Contains([]string{"limit", "offset"}, key) {
			params[key] = cu.ToString(value, "")
		}
	}

	response = []any{}
	var rows []cu.IM
	if rows, err = ds.StoreDataGet(params, true); err == nil {
		for _, row := range rows {
			var result any = sm.ResultListType()
			if dataJson, found := row[model+"_object"].(string); found {
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

func queryParameters(ctx context.Context, req *mcp.CallToolRequest, inputData cu.IM) (result *mcp.CallToolResult, response any, err error) {
	extra := req.GetExtra()
	ds := extra.TokenInfo.Extra["ds"].(*api.DataStore)
	model := cu.ToString(inputData["model"], "")
	parameters := ut.ToStringArray(inputData["parameters"])

	var sm *ModelSchema
	var found bool
	if sm, found = getSchemaMap()[model]; !found {
		return nil, nil, fmt.Errorf("invalid model: %s", model)
	}

	var prm cu.IM = cu.IM{}
	for _, parameter := range parameters {
		keyValue := strings.Split(strings.Trim(parameter, " "), "=")
		if len(keyValue) == 2 {
			prm[keyValue[0]] = keyValue[1]
		}
	}

	var params cu.IM = cu.IM{
		"fields": []string{model + "_object"},
		"model":  model + "_view",
		"limit":  cu.ToInteger(prm["limit"], 0),
		"offset": cu.ToInteger(prm["offset"], 0),
	}
	for key, value := range prm {
		if !slices.Contains([]string{"limit", "offset"}, key) {
			params[key] = cu.ToString(value, "")
		}
	}

	response = []any{}
	var rows []cu.IM
	if rows, err = ds.StoreDataGet(params, true); err == nil {
		for _, row := range rows {
			var result any = sm.ResultListType()
			if dataJson, found := row[model+"_object"].(string); found {
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
