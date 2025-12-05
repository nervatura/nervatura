package mcp

import (
	"fmt"
	"slices"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
)

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
	baseContact
}

type contactProject struct {
	Code        string `json:"code,omitempty" jsonschema:"Database independent unique external key."`
	ProjectName string `json:"project_name,omitempty" jsonschema:"Full name of the project."`
	baseContact
}

type contactPlace struct {
	Code      string `json:"code,omitempty" jsonschema:"Database independent unique external key."`
	PlaceName string `json:"place_name,omitempty" jsonschema:"Full name of the place."`
	baseContact
}

type baseContact struct {
	FirstName  string   `json:"first_name,omitempty" jsonschema:"First name."`
	Surname    string   `json:"surname,omitempty" jsonschema:"Surname."`
	Status     string   `json:"status,omitempty" jsonschema:"Status."`
	Phone      string   `json:"phone,omitempty" jsonschema:"Phone."`
	Mobile     string   `json:"mobile,omitempty" jsonschema:"Mobile."`
	Email      string   `json:"email,omitempty" jsonschema:"Email."`
	Notes      string   `json:"notes,omitempty" jsonschema:"Notes."`
	Tags       []string `json:"tags,omitempty" jsonschema:"Additional tags for the contact. The value is an array of strings."`
	ContactMap cu.IM    `json:"contact_map,omitempty" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
}

type contactParameter struct {
	Model   string `json:"model" jsonschema:"Model. Enum values. Required."`
	Surname string `json:"surname,omitempty" jsonschema:"Surname."`
	Email   string `json:"email,omitempty" jsonschema:"Email."`
	Tag     string `json:"tag,omitempty" jsonschema:"Tag."`
	Limit   int64  `json:"limit,omitempty" jsonschema:"Limit."`
	Offset  int64  `json:"offset,omitempty" jsonschema:"Offset."`
}

func contactCreateTool(scope string) (tool *mcp.Tool) {
	if scope == "all" || scope == "public" {
		scope = ""
	}
	return &mcp.Tool{
		Name:        "nervatura_contact_create",
		Title:       "Contact Data Create",
		Description: fmt.Sprintf("Create a new %s contact.", scope),
		InputSchema: getExtendSchemaMap()["contact"].CreateInputSchema(scope),
	}
}

func contactUpdateTool(scope string) (tool *mcp.Tool) {
	if scope == "all" || scope == "public" {
		scope = ""
	}
	return &mcp.Tool{
		Name:        "nervatura_contact_update",
		Title:       "Contact Data Update",
		Description: fmt.Sprintf("Update a %s contact.", scope),
		InputSchema: getExtendSchemaMap()["contact"].UpdateInputSchema(scope),
	}
}

func contactQueryTool(scope string) (tool *mcp.Tool) {
	return &mcp.Tool{
		Name:         "nervatura_contact_query",
		Title:        "Contact Query",
		Description:  "Query contacts by parameters. The result is all contacts that match the filter criteria.",
		InputSchema:  getExtendSchemaMap()["contact"].QueryInputSchema(scope),
		OutputSchema: getExtendSchemaMap()["contact"].QueryOutputSchema(scope),
	}
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
				delete(schema.Properties, "contact_map")
				schema.Properties["tags"].Default = []byte(`[]`)
				schema.Required = []string{"code"}
			}
			return schema
		},
		UpdateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[contactUpdate](nil); err == nil {
				schema.Description = fmt.Sprintf("Update a %s contact.", scope)
				delete(schema.Properties, "contact_map")
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
				schemaList = append(schemaList, schema)
			}
			if schema, err := jsonschema.For[contactProject](nil); err == nil {
				schema.Description = "Project contacts."
				schema.AdditionalProperties = &jsonschema.Schema{}
				schemaList = append(schemaList, schema)
			}
			if schema, err := jsonschema.For[contactPlace](nil); err == nil {
				schema.Description = "Place contacts."
				schema.AdditionalProperties = &jsonschema.Schema{}
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
