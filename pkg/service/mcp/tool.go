package mcp

import (
	"context"
	"errors"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

func init() {
	toolDataMap["nervatura_tool_create"] = McpTool{
		Tool: mcp.Tool{
			Name:        "nervatura_tool_create",
			Title:       "Create a new tool",
			Description: "Create a new tool. Related tools: event, waybill.",
			Meta: mcp.Meta{
				"scopes": []string{"tool"},
			},
		},
		ModelSchema: ToolSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, toolCreateHandler)
		},
	}
	toolDataMap["nervatura_tool_query"] = McpTool{
		Tool: mcp.Tool{
			Name:        "nervatura_tool_query",
			Title:       "Query tools by parameters",
			Description: "Query tools by parameters. The result is all tools that match the filter criteria.",
			Meta: mcp.Meta{
				"scopes": []string{"tool"},
			},
		},
		ModelSchema: ToolSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelQuery)
		},
	}
	toolDataMap["nervatura_tool_update"] = McpTool{
		Tool: mcp.Tool{
			Name:        "nervatura_tool_update",
			Title:       "Update a tool by code",
			Description: "Update a tool by code. When modifying, only the specified values change. Related tools: event, waybill.",
			Meta: mcp.Meta{
				"scopes": []string{"tool"},
			},
		},
		ModelSchema: ToolSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelUpdate)
		},
	}
	toolDataMap["nervatura_tool_delete"] = McpTool{
		Tool: createDeleteTool("nervatura_tool_delete", "tool",
			mcp.Meta{"scopes": []string{"tool"}}),
		ModelSchema: ToolSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelDelete)
		},
	}
}

type toolCreate struct {
	Description string `json:"description" jsonschema:"Tool name or description. Required when creating a new tool."`
	ProductCode string `json:"product_code" jsonschema:"Reference to product.code. Example: PRD1731101982N12. Required when creating a new tool."`
	md.ToolMeta
	ToolMap cu.IM `json:"tool_map,omitempty" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
}

type toolUpdate struct {
	Code        string `json:"code" jsonschema:"Database independent unique key. Required when updating an existing tool."`
	Description string `json:"description,omitempty" jsonschema:"Tool name or description."`
	ProductCode string `json:"product_code,omitempty" jsonschema:"Reference to product.code. Example: PRD1731101982N12"`
	md.ToolMeta
	ToolMap cu.IM `json:"tool_map,omitempty" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
}

type toolParameter struct {
	Code        string `json:"code,omitempty" jsonschema:"Database independent unique key."`
	Description string `json:"description,omitempty" jsonschema:"Tool name or description."`
	ProductCode string `json:"product_code,omitempty" jsonschema:"Reference to product.code. Example: PRD1731101982N12"`
	Tag         string `json:"tag,omitempty" jsonschema:"Tag."`
	Limit       int64  `json:"limit,omitempty" jsonschema:"Limit."`
	Offset      int64  `json:"offset,omitempty" jsonschema:"Offset."`
}

func ToolSchema() (ms *ModelSchema) {
	return &ModelSchema{
		Name:   "tool",
		Prefix: "SER",
		CreateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[toolCreate](nil); err == nil {
				schema.Properties["tool_map"].Default = []byte(`{}`)
				schema.Properties["tags"].Default = []byte(`[]`)
				schema.Required = []string{"description", "product_code"}
			}
			return schema
		},
		UpdateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[toolUpdate](nil); err == nil {
				schema.Properties["tool_map"].Default = []byte(`{}`)
				schema.Required = []string{"code"}
			}
			return schema
		},
		QueryInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			schema, _ = jsonschema.For[toolParameter](nil)
			return schema
		},
		QueryOutputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[md.Tool](nil); err == nil {
				schema.Description = "Tool data"
				schema.AdditionalProperties = &jsonschema.Schema{}
			}
			return &jsonschema.Schema{
				Type:  "object",
				Items: schema,
			}
		},
		LoadData: toolLoadData,
		LoadList: func(rows []cu.IM) (items any, err error) {
			var tools []md.Tool = []md.Tool{}
			err = cu.ConvertToType(rows, &tools)
			return tools, err
		},
		PrimaryFields: []string{"id", "code", "description", "product_code", "tool_map"},
		Required:      []string{"description", "product_code"},
	}
}

func toolLoadData(data any) (modelData, metaData any, err error) {
	var tool md.Tool = md.Tool{
		Events: []md.Event{},
		ToolMeta: md.ToolMeta{
			Tags: []string{},
		},
		ToolMap: cu.IM{},
	}
	err = cu.ConvertToType(data, &tool)
	return tool, tool.ToolMeta, err
}

func toolCreateHandler(ctx context.Context, req *mcp.CallToolRequest, inputData toolCreate) (result *mcp.CallToolResult, response UpdateResponseData, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)

	if inputData.Description == "" || inputData.ProductCode == "" {
		return result, UpdateResponseData{}, errors.New("description and product code are required")
	}

	values := cu.IM{
		"description":  inputData.Description,
		"product_code": inputData.ProductCode,
	}
	if inputData.ToolMeta.Tags == nil {
		inputData.ToolMeta.Tags = []string{}
	}

	ut.ConvertByteToIMData([]md.Event{}, values, "events")
	ut.ConvertByteToIMData(inputData.ToolMeta, values, "tool_meta")
	ut.ConvertByteToIMData(inputData.ToolMap, values, "tool_map")

	var rows []cu.IM
	var toolID int64
	var code string
	if toolID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "tool"}); err == nil {
		if rows, err = ds.StoreDataGet(cu.IM{"id": toolID, "model": "tool"}, true); err == nil {
			code = cu.ToString(rows[0]["code"], "")
		}
	}
	response = UpdateResponseData{
		Model: "tool",
		Code:  code,
		ID:    toolID,
	}

	return result, response, err
}
