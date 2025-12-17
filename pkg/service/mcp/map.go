package mcp

import (
	"context"
	"errors"
	"fmt"
	"maps"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

func init() {
	toolDataMap["nervatura_map_create"] = ToolData{
		Tool: mcp.Tool{
			Name:        "nervatura_map_create",
			Title:       "Create a new custom field definition",
			Description: "Create a new custom field definition. Example: customer_map, product_map, etc.",
			Meta: mcp.Meta{
				"scopes": []string{"setting"},
			},
		},
		ModelSchema: MapSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, mapCreateHandler)
		},
	}
	toolDataMap["nervatura_map_query"] = ToolData{
		Tool: mcp.Tool{
			Name:        "nervatura_map_query",
			Title:       "Query map configs by parameters",
			Description: "Query map configs by parameters. The result is all map configs that match the filter criteria.",
			Meta: mcp.Meta{
				"scopes": []string{"setting"},
			},
		},
		ModelSchema: MapSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelQuery)
		},
	}
	toolDataMap["nervatura_map_update"] = ToolData{
		Tool: mcp.Tool{
			Name:        "nervatura_map_update",
			Title:       "Update a map config by code",
			Description: "Update a map config by code. When modifying, only the specified values change.",
			Meta: mcp.Meta{
				"scopes": []string{"setting"},
			},
		},
		ModelSchema: MapSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, mapUpdate)
		},
	}
	toolDataMap["nervatura_map_delete"] = ToolData{
		Tool: createDeleteTool("nervatura_map_delete", "config",
			mcp.Meta{"scopes": []string{"setting"}}),
		ModelSchema: MapSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelDelete)
		},
	}
}

type updateMap struct {
	Code string `json:"code" jsonschema:"Database independent unique key. Required when updating an existing map config."`
	md.ConfigMap
}

type mapParameter struct {
	Code      string `json:"code,omitempty" jsonschema:"Database independent unique key."`
	FieldType string `json:"field_type,omitempty" jsonschema:"Field type. Enum values."`
	Tag       string `json:"tag,omitempty" jsonschema:"Tag."`
	Limit     int64  `json:"limit,omitempty" jsonschema:"Limit."`
	Offset    int64  `json:"offset,omitempty" jsonschema:"Offset."`
}

func MapSchema() (ms *ModelSchema) {
	return &ModelSchema{
		Name:       "config",
		Prefix:     "CNF",
		CustomFrom: "config_map",
		CreateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[md.ConfigMap](nil); err == nil {
				schema.Properties["field_type"].Type = "string"
				schema.Properties["field_type"].Enum = ut.ToAnyArray(md.FieldType(0).Keys())
				schema.Properties["field_type"].Default = []byte(`"` + md.FieldTypeString.String() + `"`)
				schema.Properties["filter"].Type = "string"
				schema.Properties["filter"].Enum = ut.ToAnyArray(md.MapFilter(0).Keys())
				schema.Properties["tags"].Default = []byte(`[]`)
				schema.Required = []string{"field_name", "field_type", "description"}
			}
			return schema
		},
		UpdateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[updateMap](nil); err == nil {
				schema.Properties["field_type"].Type = "string"
				schema.Properties["field_type"].Enum = ut.ToAnyArray(md.FieldType(0).Keys())
				schema.Properties["filter"].Type = "string"
				schema.Properties["filter"].Enum = ut.ToAnyArray(md.MapFilter(0).Keys())
				schema.Required = []string{"code"}
			}
			return schema
		},
		QueryInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[mapParameter](nil); err == nil {
				schema.Properties["field_type"].Type = "string"
				schema.Properties["field_type"].Enum = ut.ToAnyArray(md.FieldType(0).Keys())
			}
			return schema
		},
		QueryOutputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[md.Config](nil); err == nil {
				schema.Description = "Map configs"
				schema.AdditionalProperties = &jsonschema.Schema{}
			}
			return &jsonschema.Schema{
				Type:  "object",
				Items: schema,
			}
		},
		LoadList: func(rows []cu.IM) (items any, err error) {
			var configs []updateMap = []updateMap{}
			err = cu.ConvertToType(rows, &configs)
			return configs, err
		},
	}
}

func mapCreateHandler(ctx context.Context, req *mcp.CallToolRequest, inputData md.ConfigMap) (result *mcp.CallToolResult, response UpdateResponseData, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)

	if inputData.FieldName == "" || inputData.Description == "" {
		return result, UpdateResponseData{}, errors.New("field name and description are required")
	}

	values := cu.IM{
		"config_type": md.ConfigTypeMap.String(),
	}
	if inputData.Tags == nil {
		inputData.Tags = []string{}
	}
	if inputData.Filter == nil {
		inputData.Filter = []md.MapFilter{}
	}

	ut.ConvertByteToIMData(inputData, values, "data")

	var rows []cu.IM
	var configID int64
	var code string
	if configID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "config"}); err == nil {
		if rows, err = ds.StoreDataGet(cu.IM{"id": configID, "model": "config"}, true); err == nil {
			code = cu.ToString(rows[0]["code"], "")
		}
	}
	response = UpdateResponseData{
		Model: "config",
		Code:  code,
		ID:    configID,
	}

	return result, response, err
}

func mapUpdate(ctx context.Context, req *mcp.CallToolRequest, inputData cu.IM) (result *mcp.CallToolResult, response UpdateResponseData, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)

	code := cu.ToString(inputData["code"], "")
	if code == "" {
		return nil, UpdateResponseData{}, fmt.Errorf("code is required")
	}

	var rows []cu.IM
	if rows, err = ds.StoreDataGet(cu.IM{
		"fields": []string{"*"}, "model": "config", "code": code}, true); err != nil {
		return nil, UpdateResponseData{}, errors.New("invalid code: " + code)
	}
	updateID := cu.ToInteger(rows[0]["id"], 0)
	fieldValues := cu.ToIM(rows[0]["data"], cu.IM{})
	maps.Copy(fieldValues, inputData)

	var cnf md.ConfigMap = md.ConfigMap{
		Tags:   []string{},
		Filter: []md.MapFilter{},
	}
	if err = cu.ConvertToType(fieldValues, &cnf); err == nil {
		values := cu.IM{}
		ut.ConvertByteToIMData(cnf, values, "data")
		_, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "config", IDKey: updateID})
	}

	return result, UpdateResponseData{
		Model: "config",
		Code:  code,
		ID:    updateID,
	}, err
}
