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
	toolDataMap["nervatura_product_create"] = ToolData{
		Tool: mcp.Tool{
			Name:        "nervatura_product_create",
			Title:       "Create a new product",
			Description: "Create a new product. Related tools: price, event.",
			Meta: mcp.Meta{
				"scopes": []string{"product"},
			},
		},
		ModelSchema: ProductSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, productCreateHandler)
		},
	}
	toolDataMap["nervatura_product_query"] = ToolData{
		Tool: mcp.Tool{
			Name:        "nervatura_product_query",
			Title:       "Query products by parameters",
			Description: "Query products by parameters. The result is all products that match the filter criteria.",
			Meta: mcp.Meta{
				"scopes": []string{"product"},
			},
		},
		ModelSchema: ProductSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelQuery)
		},
	}
	toolDataMap["nervatura_product_update"] = ToolData{
		Tool: mcp.Tool{
			Name:        "nervatura_product_update",
			Title:       "Update a product by code",
			Description: "Update a product by code. When modifying, only the specified values change. Related tools: price, event.",
			Meta: mcp.Meta{
				"scopes": []string{"product"},
			},
		},
		ModelSchema: ProductSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelUpdate)
		},
	}
	toolDataMap["nervatura_product_delete"] = ToolData{
		Tool:        createDeleteTool("nervatura_product_delete", "product", mcp.Meta{"scopes": []string{"product"}}),
		ModelSchema: ProductSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelDelete)
		},
	}
}

type productCreate struct {
	ProductType string `json:"product_type" jsonschema:"Product type. Enum values. Required when creating a new product."`
	ProductName string `json:"product_name" jsonschema:"Full name of the product. Required when creating a new product."`
	TaxCode     string `json:"tax_code" jsonschema:"Tax code. Required when creating a new product. Example: VAT20"`
	md.ProductMeta
	ProductMap cu.IM `json:"product_map,omitempty" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
}

type productUpdate struct {
	Code        string `json:"code" jsonschema:"Database independent unique key. Required when updating an existing product."`
	ProductType string `json:"product_type,omitempty" jsonschema:"Product type. Enum values."`
	ProductName string `json:"product_name,omitempty" jsonschema:"Full name of the product."`
	TaxCode     string `json:"tax_code,omitempty" jsonschema:"Tax code. Example: VAT20"`
	md.ProductMeta
	ProductMap cu.IM `json:"product_map,omitempty" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
}

type productParameter struct {
	Code        string `json:"code,omitempty" jsonschema:"Database independent unique key."`
	ProductType string `json:"product_type,omitempty" jsonschema:"Product type. Enum values."`
	ProductName string `json:"product_name,omitempty" jsonschema:"Full name of the product."`
	Tag         string `json:"tag,omitempty" jsonschema:"Tag."`
	Limit       int64  `json:"limit,omitempty" jsonschema:"Limit."`
	Offset      int64  `json:"offset,omitempty" jsonschema:"Offset."`
}

func ProductSchema() (ms *ModelSchema) {
	return &ModelSchema{
		Name:   "product",
		Prefix: "PRD",
		CreateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[productCreate](nil); err == nil {
				schema.Properties["product_type"].Type = "string"
				schema.Properties["product_type"].Enum = ut.ToAnyArray(md.ProductType(0).Keys())
				schema.Properties["product_type"].Default = []byte(`"` + md.ProductTypeItem.String() + `"`)
				schema.Properties["product_map"].Default = []byte(`{}`)
				schema.Properties["tags"].Default = []byte(`[]`)
				schema.Required = []string{"product_type", "product_name", "tax_code"}
			}
			return schema
		},
		UpdateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[productUpdate](nil); err == nil {
				schema.Properties["product_type"].Type = "string"
				schema.Properties["product_type"].Enum = []any{md.ProductType(0).Keys()}
				schema.Properties["product_map"].Default = []byte(`{}`)
				schema.Required = []string{"code"}
			}
			return schema
		},
		QueryInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			schema, _ = jsonschema.For[productParameter](nil)
			return schema
		},
		QueryOutputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[md.Product](nil); err == nil {
				schema.Description = "Product data"
				schema.AdditionalProperties = &jsonschema.Schema{}
			}
			return &jsonschema.Schema{
				Type:  "object",
				Items: schema,
			}
		},
		LoadData: func(data any) (modelData, metaData any, err error) {
			var product md.Product = md.Product{
				ProductType: md.ProductTypeItem,
				Events:      []md.Event{},
				ProductMeta: md.ProductMeta{
					Tags: []string{},
				},
				ProductMap: cu.IM{},
			}
			err = cu.ConvertToType(data, &product)
			return product, product.ProductMeta, err
		},
		LoadList: func(rows []cu.IM) (items any, err error) {
			var products []md.Product = []md.Product{}
			err = cu.ConvertToType(rows, &products)
			return products, err
		},
		PrimaryFields: []string{"id", "code", "product_type", "product_name", "tax_code", "product_map"},
		Required:      []string{"product_name", "product_type", "tax_code"},
	}
}

func productCreateHandler(ctx context.Context, req *mcp.CallToolRequest, inputData productCreate) (result *mcp.CallToolResult, response UpdateResponseData, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)

	if inputData.ProductName == "" || inputData.TaxCode == "" {
		return result, UpdateResponseData{}, errors.New("product name and tax code are required")
	}

	values := cu.IM{
		"product_type": inputData.ProductType,
		"product_name": inputData.ProductName,
		"tax_code":     inputData.TaxCode,
	}
	if inputData.ProductMeta.Tags == nil {
		inputData.ProductMeta.Tags = []string{}
	}

	ut.ConvertByteToIMData([]md.Event{}, values, "events")
	ut.ConvertByteToIMData(inputData.ProductMeta, values, "product_meta")
	ut.ConvertByteToIMData(inputData.ProductMap, values, "product_map")

	var rows []cu.IM
	var productID int64
	var code string
	if productID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "product"}); err == nil {
		if rows, err = ds.StoreDataGet(cu.IM{"id": productID, "model": "product"}, true); err == nil {
			code = cu.ToString(rows[0]["code"], "")
		}
	}
	response = UpdateResponseData{
		Model: "product",
		Code:  code,
		ID:    productID,
	}

	return result, response, err
}
