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
	toolDataMap["nervatura_currency_create"] = McpTool{
		Tool: mcp.Tool{
			Name:        "nervatura_currency_create",
			Title:       "Create a new currency",
			Description: "Create a new currency.",
			Meta: mcp.Meta{
				"scopes": []string{"setting"},
			},
		},
		ModelSchema: CurrencySchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, currencyCreateHandler)
		},
	}
	toolDataMap["nervatura_currency_query"] = McpTool{
		Tool: mcp.Tool{
			Name:        "nervatura_currency_query",
			Title:       "Query currencies by parameters",
			Description: "Query currencies by parameters. The result is all currencies that match the filter criteria.",
			Meta: mcp.Meta{
				"scopes": []string{"setting"},
			},
		},
		ModelSchema: CurrencySchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelQuery)
		},
	}
	toolDataMap["nervatura_currency_update"] = McpTool{
		Tool: mcp.Tool{
			Name:        "nervatura_currency_update",
			Title:       "Update a currency by code",
			Description: "Update a currency by code. When modifying, only the specified values change.",
			Meta: mcp.Meta{
				"scopes": []string{"setting"},
			},
		},
		ModelSchema: CurrencySchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelUpdate)
		},
	}
	toolDataMap["nervatura_currency_delete"] = McpTool{
		Tool: mcp.Tool{
			Name:        "nervatura_currency_delete",
			Title:       "Delete a currency by code",
			Description: "Delete a currency by code.",
			Meta: mcp.Meta{
				"scopes": []string{"setting"},
			},
			InputSchema: &jsonschema.Schema{
				Type: "object",
				Properties: map[string]*jsonschema.Schema{
					"code": {Type: "string", MinLength: ut.AnyPointer(3), MaxLength: ut.AnyPointer(3), Pattern: `^[A-Z]{3}$`},
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
		ModelSchema: CurrencySchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, currencyDelete)
		},
	}
}

type currencyUpdate struct {
	Code string `json:"code" jsonschema:"The ISO 4217 code of the currency. The value is always mandatory. Example: EUR"`
	md.CurrencyMeta
	CurrencyMap cu.IM `json:"currency_map,omitempty" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
}

type currencyParameter struct {
	Code   string `json:"code,omitempty" jsonschema:"Database independent unique key."`
	Tag    string `json:"tag,omitempty" jsonschema:"Tag."`
	Limit  int64  `json:"limit,omitempty" jsonschema:"Limit."`
	Offset int64  `json:"offset,omitempty" jsonschema:"Offset."`
}

func CurrencySchema() (ms *ModelSchema) {
	return &ModelSchema{
		Name:   "currency",
		Prefix: "",
		CreateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[currencyUpdate](nil); err == nil {
				schema.Properties["code"].MinLength = ut.AnyPointer(3)
				schema.Properties["code"].MaxLength = ut.AnyPointer(3)
				schema.Properties["code"].Pattern = `^[A-Z]{3}$`
				schema.Properties["currency_map"].Default = []byte(`{}`)
				schema.Properties["tags"].Default = []byte(`[]`)
				schema.Required = []string{"code", "description"}
			}
			return schema
		},
		UpdateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[currencyUpdate](nil); err == nil {
				schema.Properties["code"].MinLength = ut.AnyPointer(3)
				schema.Properties["code"].MaxLength = ut.AnyPointer(3)
				schema.Properties["code"].Pattern = `^[A-Z]{3}$`
				schema.Properties["currency_map"].Default = []byte(`{}`)
				schema.Required = []string{"code"}
			}
			return schema
		},
		QueryInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			schema, _ = jsonschema.For[currencyParameter](nil)
			return schema
		},
		QueryOutputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[md.Currency](nil); err == nil {
				schema.Description = "Currencies"
				schema.AdditionalProperties = &jsonschema.Schema{}
			}
			return &jsonschema.Schema{
				Type:  "object",
				Items: schema,
			}
		},
		LoadData: currencyLoadData,
		LoadList: func(rows []cu.IM) (items any, err error) {
			var currencies []md.Currency = []md.Currency{}
			err = cu.ConvertToType(rows, &currencies)
			return currencies, err
		},
		PrimaryFields: []string{"id", "code", "currency_meta", "currency_map"},
		Required:      []string{"code", "description"},
	}
}

func currencyLoadData(data any) (modelData, metaData any, err error) {
	var currency md.Currency = md.Currency{
		CurrencyMeta: md.CurrencyMeta{
			Tags: []string{},
		},
		CurrencyMap: cu.IM{},
	}
	err = cu.ConvertToType(data, &currency)
	return currency, currency.CurrencyMeta, err
}

func currencyCreateHandler(ctx context.Context, req *mcp.CallToolRequest, inputData currencyUpdate) (result *mcp.CallToolResult, response UpdateResponseData, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)

	if inputData.Code == "" || len(inputData.Code) != 3 {
		return result, UpdateResponseData{}, errors.New("code is required and must be 3 characters long")
	}

	values := cu.IM{
		"code": inputData.Code,
	}
	if inputData.CurrencyMeta.Tags == nil {
		inputData.CurrencyMeta.Tags = []string{}
	}

	ut.ConvertByteToIMData(inputData.CurrencyMeta, values, "currency_meta")
	ut.ConvertByteToIMData(inputData.CurrencyMap, values, "currency_map")

	var currencyID int64
	var code string
	currencyID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "currency"})
	response = UpdateResponseData{
		Model: "currency",
		Code:  code,
		ID:    currencyID,
	}

	return result, response, err
}

func currencyDelete(ctx context.Context, req *mcp.CallToolRequest, inputData map[string]any) (result *mcp.CallToolResult, response any, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)

	code := cu.ToString(inputData["code"], "")
	if code == "" || len(code) != 3 {
		return nil, nil, errors.New("code is required and must be 3 characters long")
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
		err = ds.DataDelete("currency", 0, code)
	}

	return &mcp.CallToolResult{
		StructuredContent: cu.IM{"success": err == nil && confirm},
		IsError:           err != nil,
	}, nil, err
}
