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
	toolDataMap["nervatura_customer_create"] = ToolData{
		Tool: mcp.Tool{
			Name:        "nervatura_customer_create",
			Title:       "Create a new customer",
			Description: "Create a new customer. Related tools: contact, address, event.",
			Meta: mcp.Meta{
				"scopes": []string{"customer"},
			},
		},
		ModelSchema: CustomerSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, customerCreateHandler)
		},
	}
	toolDataMap["nervatura_customer_query"] = ToolData{
		Tool: mcp.Tool{
			Name:        "nervatura_customer_query",
			Title:       "Query customers by parameters",
			Description: "Query customers by parameters. The result is all customers that match the filter criteria.",
			Meta: mcp.Meta{
				"scopes": []string{"customer"},
			},
		},
		ModelSchema: CustomerSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelQuery)
		},
	}
	toolDataMap["nervatura_customer_update"] = ToolData{
		Tool: mcp.Tool{
			Name:        "nervatura_customer_update",
			Title:       "Update a customer by code",
			Description: "Update a customer by code. When modifying, only the specified values change. Related tools: contact, address, event.",
			Meta: mcp.Meta{
				"scopes": []string{"customer"},
			},
		},
		ModelSchema: CustomerSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelUpdate)
		},
	}
	toolDataMap["nervatura_customer_delete"] = ToolData{
		Tool: createDeleteTool("nervatura_customer_delete", "customer",
			mcp.Meta{"scopes": []string{"customer"}}),
		ModelSchema: CustomerSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelDelete)
		},
	}
}

type customerCreate struct {
	CustomerType string `json:"customer_type" jsonschema:"Customer type. Enum values. Required when creating a new customer."`
	CustomerName string `json:"customer_name" jsonschema:"Full name of the customer. Required when creating a new customer."`
	md.CustomerMeta
	CustomerMap cu.IM `json:"customer_map,omitempty" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
}

type customerUpdate struct {
	Code         string `json:"code" jsonschema:"Database independent unique key. Required when updating an existing customer."`
	CustomerType string `json:"customer_type,omitempty" jsonschema:"Customer type. Enum values."`
	CustomerName string `json:"customer_name,omitempty" jsonschema:"Full name of the customer."`
	md.CustomerMeta
	CustomerMap cu.IM `json:"customer_map,omitempty" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
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
				schema.Properties["customer_type"].Type = "string"
				schema.Properties["customer_type"].Enum = ut.ToAnyArray(md.CustomerType(0).Keys())
				schema.Properties["customer_type"].Default = []byte(`"` + md.CustomerTypeCompany.String() + `"`)
				schema.Properties["customer_map"].Default = []byte(`{}`)
				schema.Properties["tags"].Default = []byte(`[]`)
				schema.Required = []string{"customer_type", "customer_name"}
			}
			return schema
		},
		UpdateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[customerUpdate](nil); err == nil {
				schema.Properties["customer_type"].Type = "string"
				schema.Properties["customer_type"].Enum = ut.ToAnyArray(md.CustomerType(0).Keys())
				schema.Properties["customer_map"].Default = []byte(`{}`)
				schema.Required = []string{"code"}
			}
			return schema
		},
		QueryInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			schema, _ = jsonschema.For[customerParameter](nil)
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
			err = cu.ConvertToType(data, &customer)
			return customer, customer.CustomerMeta, err
		},
		LoadList: func(rows []cu.IM) (items any, err error) {
			var customers []md.Customer = []md.Customer{}
			err = cu.ConvertToType(rows, &customers)
			return customers, err
		},
		PrimaryFields: []string{"id", "code", "customer_type", "customer_name", "customer_map"},
		Required:      []string{"customer_name", "customer_type"},
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

	ut.ConvertByteToIMData([]md.Contact{}, values, "contacts")
	ut.ConvertByteToIMData([]md.Address{}, values, "addresses")
	ut.ConvertByteToIMData([]md.Event{}, values, "events")
	ut.ConvertByteToIMData(inputData.CustomerMeta, values, "customer_meta")
	ut.ConvertByteToIMData(inputData.CustomerMap, values, "customer_map")

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
