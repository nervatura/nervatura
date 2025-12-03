package mcp

import (
	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"

	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

type customerInput struct {
	Code         string `json:"code,omitempty" jsonschema:"Database independent unique key. If existing code is set, the customer is updated, otherwise a new customer is created."`
	CustomerType string `json:"customer_type,omitempty" jsonschema:"Customer type. Enum values."`
	CustomerName string `json:"customer_name,omitempty" jsonschema:"Full name of the customer. Required when creating a new customer."`
	md.CustomerMeta
}

type customerParameter struct {
	Code         string `json:"code,omitempty" jsonschema:"Database independent unique key."`
	CustomerType string `json:"customer_type,omitempty" jsonschema:"Customer type. Enum values."`
	CustomerName string `json:"customer_name,omitempty" jsonschema:"Full name of the customer."`
	Tag          string `json:"tag,omitempty" jsonschema:"Tag."`
	Limit        int64  `json:"limit,omitempty" jsonschema:"Limit."`
	Offset       int64  `json:"offset,omitempty" jsonschema:"Offset."`
}

func CustomerSchema() (ms *ModelSchema) {
	return &ModelSchema{
		Name:   "customer",
		Prefix: "CUS",
		ResultType: func() any {
			return md.Customer{
				CustomerType: md.CustomerTypeCompany,
				Addresses:    []md.Address{},
				Contacts:     []md.Contact{},
				Events:       []md.Event{},
				CustomerMeta: md.CustomerMeta{
					Tags: []string{},
				},
				CustomerMap: cu.IM{},
			}
		},
		ResultListType:   func() any { return []md.Customer{} },
		InputSchema:      func() (*jsonschema.Schema, error) { return jsonschema.For[customerInput](nil) },
		ParameterSchema:  func() (*jsonschema.Schema, error) { return jsonschema.For[customerParameter](nil) },
		ResultSchema:     func() (*jsonschema.Schema, error) { return jsonschema.For[md.Customer](nil) },
		ResultListSchema: func() (*jsonschema.Schema, error) { return jsonschema.For[[]md.Customer](nil) },
		SchemaModify: func(schemaType SchemaType, schema *jsonschema.Schema) {
			switch schemaType {
			case SchemaTypeResult, SchemaTypeResultList:
				schema.Description = "Customer data"
				schema.Properties["id"].ReadOnly = true
				schema.Properties["code"].ReadOnly = true
				schema.Properties["customer_type"].Type = "string"
				schema.Properties["customer_type"].Enum = []any{md.CustomerTypeCompany.String(), md.CustomerTypePrivate.String(), md.CustomerTypeOther.String(),
					md.CustomerTypeOwn.String()}
				schema.Properties["customer_type"].Default = []byte(`"` + md.CustomerTypeCompany.String() + `"`)
				schema.Properties["customer_map"].Default = []byte(`{}`)
				schema.Properties["customer_map"].AdditionalProperties = &jsonschema.Schema{}
				schema.Properties["customer_meta"].Required = []string{}
				schema.Properties["time_stamp"].Type = "string"
				schema.Properties["time_stamp"].ReadOnly = true
				schema.Properties["events"].Items.Properties["start_time"].Type = "string"
				schema.Properties["events"].Items.Properties["end_time"].Type = "string"
			case SchemaTypeInput, SchemaTypeParameter:
				schema.Properties["customer_type"].Enum = []any{md.CustomerTypeCompany.String(), md.CustomerTypePrivate.String(), md.CustomerTypeOther.String(),
					md.CustomerTypeOwn.String()}
				schema.Required = []string{}
			}
		},
		LoadData: func(data any) (modelData, metaData any, err error) {
			var customer md.Customer = md.Customer{
				CustomerType: md.CustomerTypeCompany,
				Addresses:    []md.Address{},
				Contacts:     []md.Contact{},
				Events:       []md.Event{},
				CustomerMeta: md.CustomerMeta{
					Tags: []string{},
				},
				CustomerMap: cu.IM{},
			}
			if err = cu.ConvertToType(data, &customer); err != nil {
				return customer, customer.CustomerMeta, err
			}
			return customer, customer.CustomerMeta, err
		},
		InsertValues: func(data any) (values cu.IM) {
			values = cu.IM{}
			if customer, ok := data.(md.Customer); ok {
				values["customer_type"] = customer.CustomerType.String()
				values["customer_name"] = customer.CustomerName
				if customer.CustomerMeta.Tags == nil {
					customer.CustomerMeta.Tags = []string{}
				}

				ut.ConvertByteToIMData(customer.Contacts, values, "contacts")
				ut.ConvertByteToIMData(customer.Addresses, values, "addresses")
				ut.ConvertByteToIMData(customer.Events, values, "events")
				ut.ConvertByteToIMData(customer.CustomerMeta, values, "customer_meta")
				ut.ConvertByteToIMData(customer.CustomerMap, values, "customer_map")
			}
			return values
		},
		Examples: map[string][]any{
			"id":            {12345},
			"code":          {`CUS1731101982N123`},
			"customer_type": {md.CustomerTypeCompany.String()},
			"customer_name": {`First Customer LTD`},
		},
		PrimaryFields: []string{"id", "code", "customer_type", "customer_name"},
		Required:      []string{"customer_name"},
	}
}

var customerQueryTool = mcp.Tool{
	Name:        "nervatura_customer_query",
	Title:       "Customer Data Query",
	Description: "Query customers by parameters. The result is all customers that match the filter criteria.",
	InputSchema: makeModelSchema("customer", SchemaTypeParameter),
	OutputSchema: &jsonschema.Schema{
		Type:  "object",
		Items: makeModelSchema("customer", SchemaTypeResultList),
	},
}

var customerUpdateTool = mcp.Tool{
	Name:        "nervatura_customer_update",
	Title:       "Customer Data Update",
	Description: "Update a customer by code or insert new customer data. When modifying, only the specified values change. Related tools: contact, address, event.",
	InputSchema: makeModelSchema("customer", SchemaTypeInput),
}
