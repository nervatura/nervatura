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
	toolDataMap["nervatura_event_create"] = ToolData{
		Tool: mcp.Tool{
			Name:        "nervatura_event_create",
			Title:       "Event Data Create",
			Description: "Create a new %s event.",
		},
		Extend:            true,
		ModelExtendSchema: EventSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, extendCreate)
		},
		Scopes: []string{"customer", "employee", "product", "project", "tool"},
	}
	toolDataMap["nervatura_event_update"] = ToolData{
		Tool: mcp.Tool{
			Name:        "nervatura_event_update",
			Title:       "Event Data Update",
			Description: "Update a %s event.",
		},
		Extend:            true,
		ModelExtendSchema: EventSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, extendUpdate)
		},
		Scopes: []string{"customer", "employee", "product", "project", "tool"},
	}
	toolDataMap["nervatura_event_query"] = ToolData{
		Tool: mcp.Tool{
			Name:        "nervatura_event_query",
			Title:       "Event Query",
			Description: "Query %s events by parameters. The result is all events that match the filter criteria.",
		},
		Extend:            true,
		ModelExtendSchema: EventSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, extendQuery)
		},
		Scopes: []string{"customer", "employee", "product", "project", "tool"},
	}
	toolDataMap["nervatura_event_delete"] = ToolData{
		Tool:              createExtendDeleteTool("nervatura_event_delete", "event"),
		Extend:            true,
		ModelExtendSchema: EventSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, extendDelete)
		},
		Scopes: []string{"customer", "employee", "product", "project", "tool"},
	}
}

type eventCreate struct {
	Code string `json:"code" jsonschema:"Database independent unique external key. Valid customer, project or place code."`
	md.Event
}

type eventUpdate struct {
	Code  string `json:"code" jsonschema:"Database independent unique external key. Valid customer, project or place code."`
	Index int64  `json:"index" jsonschema:"Index of the event in the list."`
	md.Event
}

type eventCustomer struct {
	Code         string `json:"code,omitempty" jsonschema:"Database independent unique external key."`
	CustomerName string `json:"customer_name,omitempty" jsonschema:"Full name of the customer."`
	md.Event
}

type eventEmployee struct {
	Code      string `json:"code,omitempty" jsonschema:"Database independent unique external key."`
	FirstName string `json:"first_name,omitempty" jsonschema:"First name of the employee."`
	Surname   string `json:"surname,omitempty" jsonschema:"Surname of the employee."`
	md.Event
}

type eventProduct struct {
	Code        string `json:"code,omitempty" jsonschema:"Database independent unique external key."`
	ProductName string `json:"product_name,omitempty" jsonschema:"Product name."`
	md.Event
}

type eventProject struct {
	Code        string `json:"code,omitempty" jsonschema:"Database independent unique external key."`
	ProjectName string `json:"project_name,omitempty" jsonschema:"Full name of the project."`
	md.Event
}

type eventTool struct {
	Code            string `json:"code,omitempty" jsonschema:"Database independent unique external key."`
	ToolDescription string `json:"tool_description,omitempty" jsonschema:"Tool description."`
	md.Event
}

type eventParameter struct {
	Code    string `json:"code,omitempty" jsonschema:"Database independent unique external key. Example: CUS1731101982N123"`
	Model   string `json:"model" jsonschema:"Model. Enum values. Required."`
	Subject string `json:"subject,omitempty" jsonschema:"Subject."`
	Place   string `json:"place,omitempty" jsonschema:"Place."`
	Tag     string `json:"tag,omitempty" jsonschema:"Tag."`
	Limit   int64  `json:"limit,omitempty" jsonschema:"Limit."`
	Offset  int64  `json:"offset,omitempty" jsonschema:"Offset."`
}

func EventSchema() (ms *ModelExtendSchema) {
	return &ModelExtendSchema{
		Model: "event",
		ViewName: cu.SM{
			"customer": "customer_events",
			"employee": "employee_events",
			"product":  "product_events",
			"project":  "project_events",
			"tool":     "tool_events",
		},
		ModelFromCode: func(code string) (model, field string, err error) {
			prefix := cu.ToString(code[:3], "XXX")
			switch prefix {
			case "CUS":
				return "customer", "events", nil
			case "EMP":
				return "employee", "events", nil
			case "PRD":
				return "product", "events", nil
			case "PRJ":
				return "project", "events", nil
			case "SER":
				return "tool", "events", nil
			}
			return "", "", fmt.Errorf("invalid code: %s", code)
		},
		CreateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[eventCreate](nil); err == nil {
				schema.Description = fmt.Sprintf("Create a new %s event.", scope)
				schema.Properties["tags"].Default = []byte(`[]`)
				schema.Properties["event_map"].Default = []byte(`{}`)
				schema.Required = []string{"code"}
			}
			return schema
		},
		UpdateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[eventUpdate](nil); err == nil {
				schema.Description = fmt.Sprintf("Update a %s event.", scope)
				schema.Properties["event_map"].Default = []byte(`{}`)
				schema.Required = []string{"code", "index"}
			}
			return schema
		},
		QueryInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[eventParameter](nil); err == nil {
				schema.Description = "Query events by parameters. The result is all events that match the filter criteria."
				schema.Properties["model"].Type = "string"
				schema.Properties["model"].Enum = []any{"customer", "employee", "product", "project", "tool"}
				if slices.Contains([]string{"customer", "project", "employee", "product", "tool"}, scope) {
					schema.Properties["model"].Enum = []any{scope}
					schema.Properties["model"].Default = []byte(`"` + scope + `"`)
				}
			}
			return schema
		},
		QueryOutputSchema: func(scope string) *jsonschema.Schema {
			schemaList := []*jsonschema.Schema{}
			if schema, err := jsonschema.For[eventCustomer](nil); err == nil {
				schema.Description = "Customer events."
				schema.AdditionalProperties = &jsonschema.Schema{}
				schema.Required = []string{}
				schemaList = append(schemaList, schema)
			}
			if schema, err := jsonschema.For[eventEmployee](nil); err == nil {
				schema.Description = "Employee events."
				schema.AdditionalProperties = &jsonschema.Schema{}
				schema.Required = []string{}
				schemaList = append(schemaList, schema)
			}
			if schema, err := jsonschema.For[eventProduct](nil); err == nil {
				schema.Description = "Product events."
				schema.AdditionalProperties = &jsonschema.Schema{}
				schema.Required = []string{}
				schemaList = append(schemaList, schema)
			}
			if schema, err := jsonschema.For[eventTool](nil); err == nil {
				schema.Description = "Tool events."
				schema.AdditionalProperties = &jsonschema.Schema{}
				schema.Required = []string{}
				schemaList = append(schemaList, schema)
			}
			if schema, err := jsonschema.For[eventProject](nil); err == nil {
				schema.Description = "Project events."
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
			var event []md.Event = []md.Event{}
			err = cu.ConvertToType(data, &event)
			return event, err
		},
		LoadList: func(model string, rows []cu.IM) (items any, err error) {
			switch model {
			case "customer":
				var customers []eventCustomer = []eventCustomer{}
				err = cu.ConvertToType(rows, &customers)
				return customers, err
			case "project":
				var projects []eventProject = []eventProject{}
				err = cu.ConvertToType(rows, &projects)
				return projects, err
			case "employee":
				var employees []eventEmployee = []eventEmployee{}
				err = cu.ConvertToType(rows, &employees)
				return employees, err
			case "product":
				var products []eventProduct = []eventProduct{}
				err = cu.ConvertToType(rows, &products)
				return products, err
			case "tool":
				var tools []eventTool = []eventTool{}
				err = cu.ConvertToType(rows, &tools)
				return tools, err
			default:
				return rows, err
			}
		},
	}
}
