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
	toolDataMap["nervatura_project_create"] = McpTool{
		Tool: mcp.Tool{
			Name:        "nervatura_project_create",
			Title:       "Create a new project",
			Description: "Create a new project. Related tools: contact, address, event.",
			Meta: mcp.Meta{
				"scopes": []string{"project"},
			},
		},
		ModelSchema: ProjectSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, projectCreateHandler)
		},
	}
	toolDataMap["nervatura_project_query"] = McpTool{
		Tool: mcp.Tool{
			Name:        "nervatura_project_query",
			Title:       "Query projects by parameters",
			Description: "Query projects by parameters. The result is all projects that match the filter criteria.",
			Meta: mcp.Meta{
				"scopes": []string{"project"},
			},
		},
		ModelSchema: ProjectSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelQuery)
		},
	}
	toolDataMap["nervatura_project_update"] = McpTool{
		Tool: mcp.Tool{
			Name:        "nervatura_project_update",
			Title:       "Update a project by code",
			Description: "Update a project by code. When modifying, only the specified values change. Related tools: contact, address, event.",
			Meta: mcp.Meta{
				"scopes": []string{"project"},
			},
		},
		ModelSchema: ProjectSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelUpdate)
		},
	}
	toolDataMap["nervatura_project_delete"] = McpTool{
		Tool: createDeleteTool("nervatura_project_delete", "project",
			mcp.Meta{"scopes": []string{"project"}}),
		ModelSchema: ProjectSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelDelete)
		},
	}
}

type projectCreate struct {
	ProjectName  string `json:"project_name" jsonschema:"Full name of the project. Required when creating a new project."`
	CustomerCode string `json:"customer_code" jsonschema:"Optional customer reference. Example: CUS1731101982N123"`
	md.ProjectMeta
	ProjectMap cu.IM `json:"project_map,omitempty" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
}

type projectUpdate struct {
	Code         string `json:"code" jsonschema:"Database independent unique key. Required when updating an existing project."`
	ProjectName  string `json:"project_name,omitempty" jsonschema:"Full name of the project."`
	CustomerCode string `json:"customer_code,omitempty" jsonschema:"Optional customer reference. Example: CUS1731101982N123"`
	md.ProjectMeta
	ProjectMap cu.IM `json:"project_map,omitempty" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
}

type projectParameter struct {
	Code         string `json:"code,omitempty" jsonschema:"Database independent unique key."`
	ProjectName  string `json:"like_project_name,omitempty" jsonschema:"Full name of the project. It is not case sensitive and partial values ​​can be specified."`
	CustomerCode string `json:"customer_code,omitempty" jsonschema:"Optional customer reference. Example: CUS1731101982N123"`
	Tag          string `json:"tag,omitempty" jsonschema:"Tag."`
	Limit        int64  `json:"limit,omitempty" jsonschema:"Limit."`
	Offset       int64  `json:"offset,omitempty" jsonschema:"Offset."`
}

func ProjectSchema() (ms *ModelSchema) {
	return &ModelSchema{
		Name:   "project",
		Prefix: "PRJ",
		CreateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[projectCreate](nil); err == nil {
				schema.Properties["project_map"].Default = []byte(`{}`)
				schema.Properties["tags"].Default = []byte(`[]`)
				schema.Required = []string{"project_name"}
			}
			return schema
		},
		UpdateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[projectUpdate](nil); err == nil {
				schema.Properties["project_map"].Default = []byte(`{}`)
				schema.Required = []string{"code"}
			}
			return schema
		},
		QueryInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			schema, _ = jsonschema.For[projectParameter](nil)
			return schema
		},
		QueryOutputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[md.Project](nil); err == nil {
				schema.Description = "Project data"
				schema.AdditionalProperties = &jsonschema.Schema{}
			}
			return &jsonschema.Schema{
				Type:  "object",
				Items: schema,
			}
		},
		LoadData: projectLoadData,
		LoadList: func(rows []cu.IM) (items any, err error) {
			var projects []md.Project = []md.Project{}
			err = cu.ConvertToType(rows, &projects)
			return projects, err
		},
		PrimaryFields: []string{"id", "code", "project_name", "customer_code", "project_map"},
		Required:      []string{"project_name"},
	}
}

func projectLoadData(data any) (modelData, metaData any, err error) {
	var project md.Project = md.Project{
		Addresses: []md.Address{},
		Contacts:  []md.Contact{},
		Events:    []md.Event{},
		ProjectMeta: md.ProjectMeta{
			Tags: []string{},
		},
		ProjectMap: cu.IM{},
	}
	err = cu.ConvertToType(data, &project)
	return project, project.ProjectMeta, err
}

func projectCreateHandler(ctx context.Context, req *mcp.CallToolRequest, inputData projectCreate) (result *mcp.CallToolResult, response UpdateResponseData, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)

	if inputData.ProjectName == "" {
		return result, UpdateResponseData{}, errors.New("project name is required")
	}

	values := cu.IM{
		"project_name": inputData.ProjectName,
	}
	if inputData.CustomerCode != "" {
		values["customer_code"] = inputData.CustomerCode
	}
	if inputData.ProjectMeta.Tags == nil {
		inputData.ProjectMeta.Tags = []string{}
	}

	ut.ConvertByteToIMData([]md.Contact{}, values, "contacts")
	ut.ConvertByteToIMData([]md.Address{}, values, "addresses")
	ut.ConvertByteToIMData([]md.Event{}, values, "events")
	ut.ConvertByteToIMData(inputData.ProjectMeta, values, "project_meta")
	ut.ConvertByteToIMData(inputData.ProjectMap, values, "project_map")

	var rows []cu.IM
	var projectID int64
	var code string
	if projectID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "project"}); err == nil {
		if rows, err = ds.StoreDataGet(cu.IM{"id": projectID, "model": "project"}, true); err == nil {
			code = cu.ToString(rows[0]["code"], "")
		}
	}
	response = UpdateResponseData{
		Model: "project",
		Code:  code,
		ID:    projectID,
	}

	return result, response, err
}
