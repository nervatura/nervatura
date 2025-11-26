package mcp

import (
	"github.com/google/jsonschema-go/jsonschema"
	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
)

type productInput struct {
	Code        string `json:"code,omitempty" jsonschema:"Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation."`
	ProductType string `json:"product_type,omitempty" jsonschema:"Product type. Enum values."`
	ProductName string `json:"product_name,omitempty" jsonschema:"Product name. Required when creating a new product."`
	TaxCode     string `json:"tax_code,omitempty" jsonschema:"Tax code."`
	md.ProductMeta
	ProductMap cu.IM `json:"product_map,omitempty" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
}

type productParameter struct {
	Code        string `json:"code,omitempty" jsonschema:"Database independent unique external key."`
	ProductType string `json:"product_type,omitempty" jsonschema:"Product type. Enum values."`
	ProductName string `json:"product_name,omitempty" jsonschema:"Product name."`
	Tag         string `json:"tag,omitempty" jsonschema:"Tag."`
	Limit       int64  `json:"limit,omitempty" jsonschema:"Limit."`
	Offset      int64  `json:"offset,omitempty" jsonschema:"Offset."`
}

func ProductSchema() (ms *ModelSchema) {
	return &ModelSchema{
		Name:             "product",
		Prefix:           "PRD",
		ResultType:       func() any { return md.Product{} },
		ResultListType:   func() any { return []md.Product{} },
		InputSchema:      func() (*jsonschema.Schema, error) { return jsonschema.For[productInput](nil) },
		ParameterSchema:  func() (*jsonschema.Schema, error) { return jsonschema.For[productParameter](nil) },
		ResultSchema:     func() (*jsonschema.Schema, error) { return jsonschema.For[md.Product](nil) },
		ResultListSchema: func() (*jsonschema.Schema, error) { return jsonschema.For[[]md.Product](nil) },
		SchemaModify: func(schemaType SchemaType, schema *jsonschema.Schema) {
			switch schemaType {
			case SchemaTypeResult, SchemaTypeResultList:
				schema.Description = "Product data"
				schema.Properties["id"].ReadOnly = true
				schema.Properties["code"].ReadOnly = true
				schema.Properties["product_type"].Type = "string"
				schema.Properties["product_type"].Enum = []any{md.ProductTypeItem.String(), md.ProductTypeService.String(), md.ProductTypeVirtual.String()}
				schema.Properties["product_type"].Default = []byte(`"` + md.ProductTypeItem.String() + `"`)
				schema.Properties["product_map"].Default = []byte(`{}`)
				schema.Properties["product_map"].AdditionalProperties = &jsonschema.Schema{}
				schema.Properties["product_meta"].Default = []byte(`{}`)
				schema.Properties["product_meta"].Properties["barcode_type"].Type = "string"
				schema.Properties["product_meta"].Properties["barcode_type"].Enum = []any{
					md.BarcodeTypeCode128.String(), md.BarcodeTypeCode39.String(), md.BarcodeTypeEan13.String(), md.BarcodeTypeEan8.String(),
					md.BarcodeTypeQRCode.String()}
				schema.Properties["time_stamp"].Type = "string"
				schema.Properties["time_stamp"].ReadOnly = true
				schema.Properties["events"].Items.Properties["start_time"].Type = "string"
				schema.Properties["events"].Items.Properties["end_time"].Type = "string"
			case SchemaTypeInput, SchemaTypeParameter:
				schema.Properties["product_type"].Enum = []any{md.ProductTypeItem.String(), md.ProductTypeService.String(), md.ProductTypeVirtual.String()}
			}
		},
		Examples: map[string][]any{
			"id":           {12345},
			"code":         {`PRD1731101982N123`},
			"product_type": {md.ProductTypeItem.String()},
			"product_name": {`First Product`},
			"tag":          {`First Tag`},
			"tax_code":     {`TAX1234567890`},
		},
		PrimaryFields: []string{"id", "code", "product_type", "product_name", "tax_code"},
		Required:      []string{"product_name"},
	}
}
