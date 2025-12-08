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
	toolDataMap["nervatura_contact_create"] = ToolData{
		Tool: mcp.Tool{
			Name:        "nervatura_contact_create",
			Title:       "Contact Data Create",
			Description: "Create a new %s contact.",
		},
		Extend:            true,
		ModelExtendSchema: ContactSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, extendCreate)
		},
		Scopes: []string{"customer", "project", "place"},
	}
	toolDataMap["nervatura_contact_update"] = ToolData{
		Tool: mcp.Tool{
			Name:        "nervatura_contact_update",
			Title:       "Contact Data Update",
			Description: "Update a %s contact.",
		},
		Extend:            true,
		ModelExtendSchema: ContactSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, extendUpdate)
		},
		Scopes: []string{"customer", "project", "place"},
	}
	toolDataMap["nervatura_contact_query"] = ToolData{
		Tool: mcp.Tool{
			Name:        "nervatura_contact_query",
			Title:       "Contact Query",
			Description: "Query %s contacts by parameters. The result is all contacts that match the filter criteria.",
		},
		Extend:            true,
		ModelExtendSchema: ContactSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, extendQuery)
		},
		Scopes: []string{"customer", "project", "place"},
	}
	toolDataMap["nervatura_contact_delete"] = ToolData{
		Tool:              createExtendDeleteTool("nervatura_contact_delete", "contact"),
		Extend:            true,
		ModelExtendSchema: ContactSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, extendDelete)
		},
		Scopes: []string{"customer", "project", "place"},
	}
}

type contactCreate struct {
	Code string `json:"code" jsonschema:"Database independent unique external key. Valid customer, project or place code."`
	md.Contact
}

type contactUpdate struct {
	Code  string `json:"code" jsonschema:"Database independent unique external key. Valid customer, project or place code."`
	Index int64  `json:"index" jsonschema:"Index of the contact in the list."`
	md.Contact
}

type contactCustomer struct {
	Code         string `json:"code,omitempty" jsonschema:"Database independent unique external key."`
	CustomerName string `json:"customer_name,omitempty" jsonschema:"Full name of the customer."`
	md.Contact
}

type contactProject struct {
	Code        string `json:"code,omitempty" jsonschema:"Database independent unique external key."`
	ProjectName string `json:"project_name,omitempty" jsonschema:"Full name of the project."`
	md.Contact
}

type contactPlace struct {
	Code      string `json:"code,omitempty" jsonschema:"Database independent unique external key."`
	PlaceName string `json:"place_name,omitempty" jsonschema:"Full name of the place."`
	md.Contact
}

type contactParameter struct {
	Model   string `json:"model" jsonschema:"Model. Enum values. Required."`
	Surname string `json:"surname,omitempty" jsonschema:"Surname."`
	Email   string `json:"email,omitempty" jsonschema:"Email."`
	Tag     string `json:"tag,omitempty" jsonschema:"Tag."`
	Limit   int64  `json:"limit,omitempty" jsonschema:"Limit."`
	Offset  int64  `json:"offset,omitempty" jsonschema:"Offset."`
}

func ContactSchema() (ms *ModelExtendSchema) {
	return &ModelExtendSchema{
		Model: "contact",
		ViewName: cu.SM{
			"customer": "customer_contacts",
			"project":  "project_contacts",
			"place":    "place_contacts",
		},
		ModelFromCode: func(code string) (model, field string, err error) {
			prefix := cu.ToString(code[:3], "XXX")
			switch prefix {
			case "CUS":
				return "customer", "contacts", nil
			case "PRJ":
				return "project", "contacts", nil
			case "PLA":
				return "place", "contacts", nil
			}
			return "", "", fmt.Errorf("invalid code: %s", code)
		},
		CreateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[contactCreate](nil); err == nil {
				schema.Description = fmt.Sprintf("Create a new %s contact.", scope)
				schema.Properties["tags"].Default = []byte(`[]`)
				schema.Properties["contact_map"].Default = []byte(`{}`)
				schema.Required = []string{"code"}
			}
			return schema
		},
		UpdateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[contactUpdate](nil); err == nil {
				schema.Description = fmt.Sprintf("Update a %s contact.", scope)
				schema.Properties["contact_map"].Default = []byte(`{}`)
				schema.Required = []string{"code", "index"}
			}
			return schema
		},
		QueryInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[contactParameter](nil); err == nil {
				schema.Description = "Query contacts by parameters. The result is all contacts that match the filter criteria."
				schema.Properties["model"].Type = "string"
				schema.Properties["model"].Enum = []any{"customer", "place", "project"}
				if slices.Contains([]string{"customer", "place", "project"}, scope) {
					schema.Properties["model"].Enum = []any{scope}
					schema.Properties["model"].Default = []byte(`"` + scope + `"`)
				}
			}
			return schema
		},
		QueryOutputSchema: func(scope string) *jsonschema.Schema {
			schemaList := []*jsonschema.Schema{}
			if schema, err := jsonschema.For[contactCustomer](nil); err == nil {
				schema.Description = "Customer contacts."
				schema.AdditionalProperties = &jsonschema.Schema{}
				schema.Required = []string{}
				schemaList = append(schemaList, schema)
			}
			if schema, err := jsonschema.For[contactProject](nil); err == nil {
				schema.Description = "Project contacts."
				schema.AdditionalProperties = &jsonschema.Schema{}
				schema.Required = []string{}
				schemaList = append(schemaList, schema)
			}
			if schema, err := jsonschema.For[contactPlace](nil); err == nil {
				schema.Description = "Place contacts."
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
			var contact []md.Contact = []md.Contact{}
			err = cu.ConvertToType(data, &contact)
			return contact, err
		},
		LoadList: func(model string, rows []cu.IM) (items any, err error) {
			switch model {
			case "customer":
				var customers []contactCustomer = []contactCustomer{}
				err = cu.ConvertToType(rows, &customers)
				return customers, err
			case "project":
				var projects []contactProject = []contactProject{}
				err = cu.ConvertToType(rows, &projects)
				return projects, err
			case "place":
				var places []contactPlace = []contactPlace{}
				err = cu.ConvertToType(rows, &places)
				return places, err
			default:
				return rows, err
			}
		},
	}
}
