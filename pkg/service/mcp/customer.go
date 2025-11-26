package mcp

import (
	"context"
	"slices"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

type customerInput struct {
	Code         string `json:"code,omitempty" jsonschema:"Database independent unique external key. If existing code is set, the customer is updated, otherwise a new customer is created."`
	CustomerType string `json:"customer_type,omitempty" jsonschema:"Customer type. Enum values."`
	CustomerName string `json:"customer_name,omitempty" jsonschema:"Full name of the customer. Required when creating a new customer."`
	md.CustomerMeta
}

type customerParameter struct {
	Code         string `json:"code,omitempty" jsonschema:"Database independent unique external key."`
	CustomerType string `json:"customer_type,omitempty" jsonschema:"Customer type. Enum values."`
	CustomerName string `json:"customer_name,omitempty" jsonschema:"Full name of the customer."`
	Tag          string `json:"tag,omitempty" jsonschema:"Tag."`
	Limit        int64  `json:"limit,omitempty" jsonschema:"Limit."`
	Offset       int64  `json:"offset,omitempty" jsonschema:"Offset."`
}

func CustomerSchema() (ms *ModelSchema) {
	return &ModelSchema{
		Name:             "customer",
		Prefix:           "CUS",
		ResultType:       func() any { return md.Customer{} },
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
				schema.Properties["customer_meta"].Default = []byte(`{}`)
				schema.Properties["time_stamp"].Type = "string"
				schema.Properties["time_stamp"].ReadOnly = true
				schema.Properties["events"].Items.Properties["start_time"].Type = "string"
				schema.Properties["events"].Items.Properties["end_time"].Type = "string"
			case SchemaTypeInput, SchemaTypeParameter:
				schema.Properties["customer_type"].Enum = []any{md.CustomerTypeCompany.String(), md.CustomerTypePrivate.String(), md.CustomerTypeOther.String(),
					md.CustomerTypeOwn.String()}
			}

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

func customerQuery(ctx context.Context, req *mcp.CallToolRequest, parameters cu.IM) (result *mcp.CallToolResult, response any, err error) {
	extra := req.GetExtra()
	ds := extra.TokenInfo.Extra["ds"].(*api.DataStore)

	var params cu.IM = cu.IM{
		"fields": []string{"customer_object"},
		"model":  "customer_view",
		"limit":  cu.ToInteger(parameters["limit"], 0),
		"offset": cu.ToInteger(parameters["offset"], 0),
	}
	for key, value := range parameters {
		if !slices.Contains([]string{"limit", "offset"}, key) {
			params[key] = cu.ToString(value, "")
		}
	}

	response = []md.Customer{}
	var rows []cu.IM
	if rows, err = ds.StoreDataGet(params, true); err == nil {
		for _, row := range rows {
			var result md.Customer
			if dataJson, found := row["customer_object"].(string); found {
				if err = ds.ConvertFromByte([]byte(dataJson), &result); err != nil {
					return nil, nil, err
				}
				response = append(response.([]md.Customer), result)
			}
		}
	}

	return &mcp.CallToolResult{
		StructuredContent: cu.IM{"items": response},
		IsError:           err != nil,
	}, nil, err
}

func customerUpdate(ctx context.Context, req *mcp.CallToolRequest, inputData customerInput) (result *mcp.CallToolResult, response UpdateResponseData, err error) {
	extra := req.GetExtra()
	ds := extra.TokenInfo.Extra["ds"].(*api.DataStore)

	var updateID int64
	var inputFields, metaFields []string
	var customer md.Customer = md.Customer{
		CustomerType: md.CustomerTypeCompany,
		Addresses:    []md.Address{},
		Contacts:     []md.Contact{},
		Events:       []md.Event{},
		CustomerMeta: md.CustomerMeta{Tags: []string{}},
		CustomerMap:  cu.IM{},
	}

	ms := CustomerSchema()
	if _, inputFields, metaFields, err = getSchemaData(req.Params.Arguments, ms, inputData, &customer); err == nil {
		if inputData.Code != "" {
			updateID, err = ds.UpdateData(md.UpdateDataOptions{
				Model: "customer", IDKey: 0, Code: inputData.Code,
				Data: customer, Meta: inputData.CustomerMeta, Fields: inputFields, MetaFields: metaFields,
			})
		} else {
			// prepare values for database update
			values := cu.IM{
				"customer_name": customer.CustomerName,
			}

			ut.ConvertByteToIMData(customer.Contacts, values, "contacts")
			ut.ConvertByteToIMData(customer.Addresses, values, "addresses")
			ut.ConvertByteToIMData(customer.Events, values, "events")
			ut.ConvertByteToIMData(inputData.CustomerMeta, values, "customer_meta")
			ut.ConvertByteToIMData(customer.CustomerMap, values, "customer_map")

			updateID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "customer"})
		}
	}
	if err == nil {
		response = UpdateResponseData{
			Model: ms.Name,
			Code:  inputData.Code,
			ID:    updateID,
		}

		if inputData.Code == "" {
			var rows []cu.IM
			if rows, err = ds.StoreDataGet(cu.IM{"fields": []string{"code"}, "id": updateID, "model": ms.Name}, true); err == nil {
				response.Code = cu.ToString(rows[0]["code"], "")
			}
		}

	}

	return result, response, err
}
