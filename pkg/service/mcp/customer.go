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

type customerCreate struct {
	CustomerType string `json:"customer_type" jsonschema:"Customer type. Enum values. Required when creating a new customer."`
	CustomerName string `json:"customer_name" jsonschema:"Full name of the customer. Required when creating a new customer."`
	md.CustomerMeta
}

type customerUpdate struct {
	Code         string `json:"code" jsonschema:"Database independent unique key. Required when updating an existing customer."`
	CustomerType string `json:"customer_type,omitempty" jsonschema:"Customer type. Enum values."`
	CustomerName string `json:"customer_name,omitempty" jsonschema:"Full name of the customer."`
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
		CreateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[customerCreate](nil); err == nil {
				schema.Description = "Create a new customer."
				schema.Properties["customer_type"].Type = "string"
				schema.Properties["customer_type"].Enum = []any{md.CustomerTypeCompany.String(), md.CustomerTypePrivate.String(), md.CustomerTypeOther.String(),
					md.CustomerTypeOwn.String()}
				schema.Properties["customer_type"].Default = []byte(`"` + md.CustomerTypeCompany.String() + `"`)
				schema.Required = []string{"customer_type", "customer_name"}
			}
			return schema
		},
		UpdateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[customerUpdate](nil); err == nil {
				schema.Description = "Update an existing customer."
				schema.Properties["customer_type"].Type = "string"
				schema.Properties["customer_type"].Enum = []any{md.CustomerTypeCompany.String(), md.CustomerTypePrivate.String(), md.CustomerTypeOther.String(),
					md.CustomerTypeOwn.String()}
				schema.Required = []string{"code"}
			}
			return schema
		},
		QueryInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[customerParameter](nil); err == nil {
				schema.Description = "Query customers by parameters. The result is all customers that match the filter criteria."
			}
			return schema
		},
		QueryOutputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[md.Customer](nil); err == nil {
				schema.Description = "Customer data"
				schema.AdditionalProperties = &jsonschema.Schema{}
			}
			return &jsonschema.Schema{
				Type:  "object",
				Items: schema,
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
		LoadList: func(rows []cu.IM) (items any, err error) {
			var customers []md.Customer = []md.Customer{}
			err = cu.ConvertToType(rows, &customers)
			return customers, err
		},
		Examples: map[string][]any{
			"id":            {12345},
			"code":          {`CUS1731101982N123`},
			"customer_type": {md.CustomerTypeCompany.String()},
			"customer_name": {`First Customer LTD`},
		},
		PrimaryFields: []string{"id", "code", "customer_type", "customer_name"},
		Required:      []string{"customer_name", "customer_type"},
	}
}

func customerQueryTool(scope string) (tool *mcp.Tool) {
	return &mcp.Tool{
		Name:         "nervatura_customer_query",
		Title:        "Customer Data Query",
		Description:  "Query customers by parameters. The result is all customers that match the filter criteria.",
		InputSchema:  getSchemaMap()["customer"].QueryInputSchema(scope),
		OutputSchema: getSchemaMap()["customer"].QueryOutputSchema(scope),
	}
}

func customerUpdateTool(scope string) (tool *mcp.Tool) {
	return &mcp.Tool{
		Name:        "nervatura_customer_update",
		Title:       "Customer Data Update",
		Description: "Update a customer by code. When modifying, only the specified values change. Related tools: contact, address, event.",
		InputSchema: getSchemaMap()["customer"].UpdateInputSchema(scope),
	}
}

func customerCreateTool(scope string) (tool *mcp.Tool) {
	return &mcp.Tool{
		Name:        "nervatura_customer_create",
		Title:       "Customer Data Create",
		Description: "Create a new customer. Related tools: contact, address, event.",
		InputSchema: getSchemaMap()["customer"].CreateInputSchema(scope),
	}
}

func customerCreateHandler(ctx context.Context, req *mcp.CallToolRequest, inputData customerCreate) (result *mcp.CallToolResult, response UpdateResponseData, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)

	if inputData.CustomerName == "" {
		return result, UpdateResponseData{}, errors.New("customer name is required")
	}

	values := cu.IM{
		"customer_type": inputData.CustomerType,
		"customer_name": inputData.CustomerName,
	}
	if inputData.CustomerMeta.Tags == nil {
		inputData.CustomerMeta.Tags = []string{}
	}

	customerMap := getParamsMeta(req)

	ut.ConvertByteToIMData([]md.Contact{}, values, "contacts")
	ut.ConvertByteToIMData([]md.Address{}, values, "addresses")
	ut.ConvertByteToIMData([]md.Event{}, values, "events")
	ut.ConvertByteToIMData(inputData.CustomerMeta, values, "customer_meta")
	ut.ConvertByteToIMData(customerMap, values, "customer_map")

	var rows []cu.IM
	var customerID int64
	var code string
	if customerID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "customer"}); err == nil {
		if rows, err = ds.StoreDataGet(cu.IM{"id": customerID, "model": "customer"}, true); err == nil {
			code = cu.ToString(rows[0]["code"], "")
		}
	}
	response = UpdateResponseData{
		Model: "customer",
		Code:  code,
		ID:    customerID,
	}

	return result, response, err
}
