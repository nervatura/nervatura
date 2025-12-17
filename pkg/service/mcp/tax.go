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

func init() {
	toolDataMap["nervatura_tax_create"] = ToolData{
		Tool: mcp.Tool{
			Name:        "nervatura_tax_create",
			Title:       "Create a new tax",
			Description: "Create a new tax.",
			Meta: mcp.Meta{
				"scopes": []string{"setting"},
			},
		},
		ModelSchema: TaxSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, taxCreateHandler)
		},
	}
	toolDataMap["nervatura_tax_query"] = ToolData{
		Tool: mcp.Tool{
			Name:        "nervatura_tax_query",
			Title:       "Query taxes by parameters",
			Description: "Query taxes by parameters. The result is all taxes that match the filter criteria.",
			Meta: mcp.Meta{
				"scopes": []string{"setting"},
			},
		},
		ModelSchema: TaxSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelQuery)
		},
	}
	toolDataMap["nervatura_tax_update"] = ToolData{
		Tool: mcp.Tool{
			Name:        "nervatura_tax_update",
			Title:       "Update a tax by code",
			Description: "Update a tax by code. When modifying, only the specified values change.",
			Meta: mcp.Meta{
				"scopes": []string{"setting"},
			},
		},
		ModelSchema: TaxSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelUpdate)
		},
	}
	toolDataMap["nervatura_tax_delete"] = ToolData{
		Tool: mcp.Tool{
			Name:        "nervatura_tax_delete",
			Title:       "Delete a tax by code",
			Description: "Delete a tax by code.",
			Meta: mcp.Meta{
				"scopes": []string{"setting"},
			},
			InputSchema: &jsonschema.Schema{
				Type: "object",
				Properties: map[string]*jsonschema.Schema{
					"code": {Type: "string", Description: "The unique key of the deleted data. Example: VAT25"},
				},
				Required: []string{"code"},
			},
			OutputSchema: &jsonschema.Schema{
				Type: "object",
				Properties: map[string]*jsonschema.Schema{
					"success": {Type: "boolean", Description: "The success of the operation."},
				},
			},
		},
		ModelSchema: TaxSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, taxDelete)
		},
	}
}

type taxUpdate struct {
	Code string `json:"code" jsonschema:"Database independent unique key. Required when updating an existing tax. Example: VAT25"`
	md.TaxMeta
	TaxMap cu.IM `json:"tax_map,omitempty" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
}

type taxParameter struct {
	Code   string `json:"code,omitempty" jsonschema:"Database independent unique key."`
	Tag    string `json:"tag,omitempty" jsonschema:"Tag."`
	Limit  int64  `json:"limit,omitempty" jsonschema:"Limit."`
	Offset int64  `json:"offset,omitempty" jsonschema:"Offset."`
}

func TaxSchema() (ms *ModelSchema) {
	return &ModelSchema{
		Name:   "tax",
		Prefix: "",
		CreateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[taxUpdate](nil); err == nil {
				schema.Properties["rate_value"].Minimum = ut.AnyPointer(0.0)
				schema.Properties["rate_value"].Maximum = ut.AnyPointer(1.0)
				schema.Properties["rate_value"].Default = []byte(`0.0`)
				schema.Properties["tax_map"].Default = []byte(`{}`)
				schema.Properties["tags"].Default = []byte(`[]`)
				schema.Required = []string{"code", "description", "rate_value"}
			}
			return schema
		},
		UpdateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[taxUpdate](nil); err == nil {
				schema.Properties["rate_value"].Minimum = ut.AnyPointer(0.0)
				schema.Properties["rate_value"].Maximum = ut.AnyPointer(1.0)
				schema.Properties["tax_map"].Default = []byte(`{}`)
				schema.Required = []string{"code"}
			}
			return schema
		},
		QueryInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			schema, _ = jsonschema.For[taxParameter](nil)
			return schema
		},
		QueryOutputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[md.Tax](nil); err == nil {
				schema.Description = "Taxes"
				schema.AdditionalProperties = &jsonschema.Schema{}
			}
			return &jsonschema.Schema{
				Type:  "object",
				Items: schema,
			}
		},
		LoadData: taxLoadData,
		LoadList: func(rows []cu.IM) (items any, err error) {
			var taxes []md.Tax = []md.Tax{}
			err = cu.ConvertToType(rows, &taxes)
			return taxes, err
		},
		PrimaryFields: []string{"id", "code", "tax_meta", "tax_map"},
		Required:      []string{"code", "description"},
	}
}

func taxLoadData(data any) (modelData, metaData any, err error) {
	var tax md.Tax = md.Tax{
		TaxMeta: md.TaxMeta{
			Tags: []string{},
		},
		TaxMap: cu.IM{},
	}
	err = cu.ConvertToType(data, &tax)
	return tax, tax.TaxMeta, err
}

func taxCreateHandler(ctx context.Context, req *mcp.CallToolRequest, inputData taxUpdate) (result *mcp.CallToolResult, response UpdateResponseData, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)

	if inputData.Code == "" {
		return result, UpdateResponseData{}, errors.New("code is required")
	}

	values := cu.IM{
		"code": inputData.Code,
	}
	if inputData.TaxMeta.Tags == nil {
		inputData.TaxMeta.Tags = []string{}
	}

	ut.ConvertByteToIMData(inputData.TaxMeta, values, "tax_meta")
	ut.ConvertByteToIMData(inputData.TaxMap, values, "tax_map")

	var taxID int64
	var code string
	taxID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "tax"})
	response = UpdateResponseData{
		Model: "tax",
		Code:  code,
		ID:    taxID,
	}

	return result, response, err
}

func taxDelete(ctx context.Context, req *mcp.CallToolRequest, inputData map[string]any) (result *mcp.CallToolResult, response any, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)

	code := cu.ToString(inputData["code"], "")
	if code == "" {
		return nil, nil, errors.New("code is required")
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
		err = ds.DataDelete("tax", 0, code)
	}

	return &mcp.CallToolResult{
		StructuredContent: cu.IM{"success": err == nil && confirm},
		IsError:           err != nil,
	}, nil, err
}
