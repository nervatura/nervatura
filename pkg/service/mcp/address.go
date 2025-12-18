package mcp

import (
	"fmt"
	"slices"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
)

func init() {
	toolDataMap["nervatura_address_create"] = McpTool{
		Tool: mcp.Tool{
			Name:        "nervatura_address_create",
			Title:       "Create a new address",
			Description: "Create a new %s address.",
			Meta: mcp.Meta{
				"scopes": []string{"customer", "project"},
			},
		},
		Extend:            true,
		ModelExtendSchema: AddressSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, extendCreate)
		},
	}
	toolDataMap["nervatura_address_update"] = McpTool{
		Tool: mcp.Tool{
			Name:        "nervatura_address_update",
			Title:       "Update a address by code",
			Description: "Update a %s address.",
			Meta: mcp.Meta{
				"scopes": []string{"customer", "project"},
			},
		},
		Extend:            true,
		ModelExtendSchema: AddressSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, extendUpdate)
		},
	}
	toolDataMap["nervatura_address_query"] = McpTool{
		Tool: mcp.Tool{
			Name:        "nervatura_address_query",
			Title:       "Query addresses by parameters",
			Description: "Query %s addresses by parameters. The result is all addresses that match the filter criteria.",
			Meta: mcp.Meta{
				"scopes": []string{"customer", "project"},
			},
		},
		Extend:            true,
		ModelExtendSchema: AddressSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, extendQuery)
		},
	}
	toolDataMap["nervatura_address_delete"] = McpTool{
		Tool:              createExtendDeleteTool("nervatura_address_delete", "address", mcp.Meta{"scopes": []string{"customer", "project"}}),
		Extend:            true,
		ModelExtendSchema: AddressSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, extendDelete)
		},
	}
}

type addressCreate struct {
	Code string `json:"code" jsonschema:"Database independent unique external key. Valid customer, project or place code."`
	md.Address
}

type addressUpdate struct {
	Code  string `json:"code" jsonschema:"Database independent unique external key. Valid customer, project or place code."`
	Index int64  `json:"index" jsonschema:"Index of the address in the list."`
	md.Address
}

type addressCustomer struct {
	Code         string `json:"code,omitempty" jsonschema:"Database independent unique external key."`
	CustomerName string `json:"customer_name,omitempty" jsonschema:"Full name of the customer."`
	md.Address
}

type addressProject struct {
	Code        string `json:"code,omitempty" jsonschema:"Database independent unique external key."`
	ProjectName string `json:"project_name,omitempty" jsonschema:"Full name of the project."`
	md.Address
}

type addressParameter struct {
	Code    string `json:"code,omitempty" jsonschema:"Database independent unique external key. Example: CUS1731101982N123"`
	Model   string `json:"model" jsonschema:"Model. Enum values. Required."`
	City    string `json:"city,omitempty" jsonschema:"City."`
	State   string `json:"state,omitempty" jsonschema:"State."`
	ZipCode string `json:"zip_code,omitempty" jsonschema:"Zip code."`
	Street  string `json:"street,omitempty" jsonschema:"Street."`
	Tag     string `json:"tag,omitempty" jsonschema:"Tag."`
	Limit   int64  `json:"limit,omitempty" jsonschema:"Limit."`
	Offset  int64  `json:"offset,omitempty" jsonschema:"Offset."`
}

func AddressSchema() (ms *ModelExtendSchema) {
	return &ModelExtendSchema{
		Model: "address",
		ViewName: cu.SM{
			"customer": "customer_addresses",
			"project":  "project_addresses",
		},
		ModelFromCode: func(code string) (model, field string, err error) {
			prefix := cu.ToString(code[:3], "XXX")
			switch prefix {
			case "CUS":
				return "customer", "addresses", nil
			case "PRJ":
				return "project", "addresses", nil
			}
			return "", "", fmt.Errorf("invalid code: %s", code)
		},
		CreateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[addressCreate](nil); err == nil {
				schema.Description = fmt.Sprintf("Create a new %s address.", scope)
				schema.Properties["tags"].Default = []byte(`[]`)
				schema.Properties["address_map"].Default = []byte(`{}`)
				schema.Required = []string{"code"}
			}
			return schema
		},
		UpdateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[addressUpdate](nil); err == nil {
				schema.Description = fmt.Sprintf("Update a %s address.", scope)
				schema.Properties["address_map"].Default = []byte(`{}`)
				schema.Required = []string{"code", "index"}
			}
			return schema
		},
		QueryInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[addressParameter](nil); err == nil {
				schema.Description = "Query addresses by parameters. The result is all addresses that match the filter criteria."
				schema.Properties["model"].Type = "string"
				schema.Properties["model"].Enum = []any{"customer", "project"}
				if slices.Contains([]string{"customer", "project"}, scope) {
					schema.Properties["model"].Enum = []any{scope}
					schema.Properties["model"].Default = []byte(`"` + scope + `"`)
				}
			}
			return schema
		},
		QueryOutputSchema: func(scope string) *jsonschema.Schema {
			schemaList := []*jsonschema.Schema{}
			if schema, err := jsonschema.For[addressCustomer](nil); err == nil {
				schema.Description = "Customer addresses."
				schema.AdditionalProperties = &jsonschema.Schema{}
				schema.Required = []string{}
				schemaList = append(schemaList, schema)
			}
			if schema, err := jsonschema.For[addressProject](nil); err == nil {
				schema.Description = "Project addresses."
				schema.AdditionalProperties = &jsonschema.Schema{}
				schema.Required = []string{}
				schemaList = append(schemaList, schema)
			}
			return &jsonschema.Schema{
				Type:  "object",
				AnyOf: schemaList,
			}
		},
		LoadData: func(data any) (modelData any, err error) {
			var address []md.Address = []md.Address{}
			err = cu.ConvertToType(data, &address)
			return address, err
		},
		LoadList: func(model string, rows []cu.IM) (items any, err error) {
			switch model {
			case "customer":
				var customers []addressCustomer = []addressCustomer{}
				err = cu.ConvertToType(rows, &customers)
				return customers, err
			case "project":
				var projects []addressProject = []addressProject{}
				err = cu.ConvertToType(rows, &projects)
				return projects, err
			default:
				return rows, err
			}
		},
	}
}
