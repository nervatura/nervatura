package mcp

import (
	"context"
	"errors"
	"time"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

func init() {
	toolDataMap["nervatura_price_create"] = ToolData{
		Tool: mcp.Tool{
			Name:        "nervatura_price_create",
			Title:       "Create a new price",
			Description: "Create a new price. Related tools: product, currency, customer.",
			Meta: mcp.Meta{
				"scopes": []string{"product"},
			},
		},
		ModelSchema: PriceSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, priceCreateHandler)
		},
	}
	toolDataMap["nervatura_price_query"] = ToolData{
		Tool: mcp.Tool{
			Name:        "nervatura_price_query",
			Title:       "Query prices by parameters",
			Description: "Query prices by parameters. The result is all prices that match the filter criteria.",
			Meta: mcp.Meta{
				"scopes": []string{"product"},
			},
		},
		ModelSchema: PriceSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelQuery)
		},
	}
	toolDataMap["nervatura_price_update"] = ToolData{
		Tool: mcp.Tool{
			Name:        "nervatura_price_update",
			Title:       "Update a price by code",
			Description: "Update a price by code. When modifying, only the specified values change. Related tools: product, currency, customer.",
			Meta: mcp.Meta{
				"scopes": []string{"product"},
			},
		},
		ModelSchema: PriceSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelUpdate)
		},
	}
	toolDataMap["nervatura_price_delete"] = ToolData{
		Tool:        createDeleteTool("nervatura_price_delete", "price", mcp.Meta{"scopes": []string{"product"}}),
		ModelSchema: PriceSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelDelete)
		},
	}
}

type priceCreate struct {
	PriceType    string  `json:"price_type" jsonschema:"Price type. Enum values. Required when creating a new price."`
	ValidFrom    string  `json:"valid_from" jsonschema:"Valid from date. ISO 8601 date. Required when creating a new price. Example: 2025-01-01"`
	ValidTo      string  `json:"valid_to" jsonschema:"Valid validity date. ISO 8601 date. Optional. Example: 2025-01-01"`
	ProductCode  string  `json:"product_code" jsonschema:"Product code. Required when creating a new price. Example: PRD1731101982N123"`
	CurrencyCode string  `json:"currency_code" jsonschema:"Currency code. Required when creating a new price. Example: EUR"`
	CustomerCode string  `json:"customer_code" jsonschema:"Customer code if the price is for a specific customer. Optional. Example: CUS1731101982N123"`
	Qty          float64 `json:"qty" jsonschema:"Quantity. Price ranges can also be specified, thus different price can be set for a smaller and bigger quantity of the same product. The quantity should be used as the lower threshold, ie. this should be the minimum quantity for the price set."`
	md.PriceMeta
	PriceMap cu.IM `json:"price_map,omitempty" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
}

type priceUpdate struct {
	Code         string  `json:"code" jsonschema:"Database independent unique key. Required when updating an existing price."`
	PriceType    string  `json:"price_type" jsonschema:"Price type. Enum values."`
	ValidFrom    string  `json:"valid_from" jsonschema:"Valid from date. ISO 8601 date. Example: 2025-01-01"`
	ValidTo      string  `json:"valid_to" jsonschema:"Valid validity date. ISO 8601 date. Example: 2025-01-01"`
	ProductCode  string  `json:"product_code" jsonschema:"Product code. Example: PRD1731101982N123"`
	CurrencyCode string  `json:"currency_code" jsonschema:"Currency code. Example: EUR"`
	CustomerCode string  `json:"customer_code" jsonschema:"Customer code if the price is for a specific customer. Example: CUS1731101982N123"`
	Qty          float64 `json:"qty" jsonschema:"Quantity. Price ranges can also be specified, thus different price can be set for a smaller and bigger quantity of the same product. The quantity should be used as the lower threshold, ie. this should be the minimum quantity for the price set."`
	md.PriceMeta
	PriceMap cu.IM `json:"price_map,omitempty" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
}

type priceParameter struct {
	Code         string `json:"code,omitempty" jsonschema:"Database independent unique key."`
	PriceType    string `json:"price_type,omitempty" jsonschema:"Price type. Enum values."`
	ValidFrom    string `json:"valid_from,omitempty" jsonschema:"Valid from date. ISO 8601 date."`
	ValidTo      string `json:"valid_to,omitempty" jsonschema:"Valid validity date. ISO 8601 date."`
	ProductCode  string `json:"product_code,omitempty" jsonschema:"Product code."`
	CurrencyCode string `json:"currency_code,omitempty" jsonschema:"Currency code."`
	CustomerCode string `json:"customer_code,omitempty" jsonschema:"Customer code if the price is for a specific customer."`
	Tag          string `json:"tag,omitempty" jsonschema:"Tag."`
	Limit        int64  `json:"limit,omitempty" jsonschema:"Limit."`
	Offset       int64  `json:"offset,omitempty" jsonschema:"Offset."`
}

func PriceSchema() (ms *ModelSchema) {
	return &ModelSchema{
		Name:   "price",
		Prefix: "PRC",
		CreateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[priceCreate](nil); err == nil {
				schema.Properties["price_type"].Type = "string"
				schema.Properties["price_type"].Enum = ut.ToAnyArray(md.PriceType(0).Keys())
				schema.Properties["price_type"].Default = []byte(`"` + md.PriceTypeCustomer.String() + `"`)
				schema.Properties["price_map"].Default = []byte(`{}`)
				schema.Properties["tags"].Default = []byte(`[]`)
				schema.Properties["valid_from"].Format = "date"
				schema.Properties["valid_to"].Format = "date"
				schema.Required = []string{"price_type", "valid_from", "product_code", "currency_code"}
			}
			return schema
		},
		UpdateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[priceUpdate](nil); err == nil {
				schema.Properties["price_type"].Type = "string"
				schema.Properties["price_type"].Enum = []any{md.PriceTypeCustomer.String(), md.PriceTypeVendor.String()}
				schema.Properties["price_map"].Default = []byte(`{}`)
				schema.Properties["valid_from"].Format = "date"
				schema.Properties["valid_to"].Format = "date"
				schema.Required = []string{"code"}
			}
			return schema
		},
		QueryInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			schema, _ = jsonschema.For[priceParameter](nil)
			return schema
		},
		QueryOutputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[md.Price](nil); err == nil {
				schema.Description = "Price data"
				schema.AdditionalProperties = &jsonschema.Schema{}
			}
			return &jsonschema.Schema{
				Type:  "object",
				Items: schema,
			}
		},
		LoadData: func(data any) (modelData, metaData any, err error) {
			var price md.Price = md.Price{
				PriceType: md.PriceTypeCustomer,
				ValidFrom: md.TimeDate{Time: time.Now()},
				PriceMeta: md.PriceMeta{
					Tags: []string{},
				},
				PriceMap: cu.IM{},
			}
			err = cu.ConvertToType(data, &price)
			return price, price.PriceMeta, err
		},
		LoadList: func(rows []cu.IM) (items any, err error) {
			var prices []md.Price = []md.Price{}
			err = cu.ConvertToType(rows, &prices)
			return prices, err
		},
		PrimaryFields: []string{"id", "code", "price_type", "valid_from", "valid_to", "product_code", "currency_code", "customer_code", "qty", "price_map"},
		Required:      []string{"price_type", "valid_from", "product_code", "currency_code"},
	}
}

func priceCreateHandler(ctx context.Context, req *mcp.CallToolRequest, inputData priceCreate) (result *mcp.CallToolResult, response UpdateResponseData, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)

	if inputData.ValidFrom == "" || inputData.ProductCode == "" || inputData.CurrencyCode == "" {
		return result, UpdateResponseData{}, errors.New("valid from, product code and currency code are required")
	}

	values := cu.IM{
		"valid_from":    inputData.ValidFrom,
		"product_code":  inputData.ProductCode,
		"price_type":    inputData.PriceType,
		"currency_code": inputData.CurrencyCode,
		"qty":           inputData.Qty,
	}
	if inputData.CustomerCode != "" {
		values["customer_code"] = inputData.CustomerCode
	}
	if inputData.ValidTo != "" {
		values["valid_to"] = inputData.ValidTo
	}
	if inputData.PriceMeta.Tags == nil {
		inputData.PriceMeta.Tags = []string{}
	}

	ut.ConvertByteToIMData(inputData.PriceMeta, values, "price_meta")
	ut.ConvertByteToIMData(inputData.PriceMap, values, "price_map")

	var rows []cu.IM
	var priceID int64
	var code string
	if priceID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "price"}); err == nil {
		if rows, err = ds.StoreDataGet(cu.IM{"id": priceID, "model": "price"}, true); err == nil {
			code = cu.ToString(rows[0]["code"], "")
		}
	}
	response = UpdateResponseData{
		Model: "price",
		Code:  code,
		ID:    priceID,
	}

	return result, response, err
}
