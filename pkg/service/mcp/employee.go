package mcp

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

func init() {
	toolDataMap["nervatura_employee_create"] = McpTool{
		Tool: mcp.Tool{
			Name:        "nervatura_employee_create",
			Title:       "Create a new employee",
			Description: "Create a new employee. Related tools: event.",
			Meta: mcp.Meta{
				"scopes": []string{"employee"},
			},
		},
		ModelSchema: EmployeeSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, employeeCreateHandler)
		},
	}
	toolDataMap["nervatura_employee_query"] = McpTool{
		Tool: mcp.Tool{
			Name:        "nervatura_employee_query",
			Title:       "Query employees by parameters",
			Description: "Query employees by parameters. The result is all employees that match the filter criteria.",
			Meta: mcp.Meta{
				"scopes": []string{"employee"},
			},
		},
		ModelSchema: EmployeeSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelQuery)
		},
	}
	toolDataMap["nervatura_employee_update"] = McpTool{
		Tool: mcp.Tool{
			Name:        "nervatura_employee_update",
			Title:       "Update a employee by code",
			Description: "Update a employee by code. When modifying, only the specified values change. Related tools: event.",
			Meta: mcp.Meta{
				"scopes": []string{"employee"},
			},
		},
		ModelSchema: EmployeeSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, employeeUpdateHandler)
		},
	}
	toolDataMap["nervatura_employee_delete"] = McpTool{
		Tool: createDeleteTool("nervatura_employee_delete", "employee",
			mcp.Meta{"scopes": []string{"employee"}}),
		ModelSchema: EmployeeSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelDelete)
		},
	}
}

type employeeContact struct {
	FirstName string `json:"first_name,omitempty" jsonschema:"First name."`
	Surname   string `json:"surname,omitempty" jsonschema:"Surname."`
	Status    string `json:"status,omitempty" jsonschema:"Status."`
	Phone     string `json:"phone,omitempty" jsonschema:"Phone."`
	Mobile    string `json:"mobile,omitempty" jsonschema:"Mobile."`
	Email     string `json:"email,omitempty" jsonschema:"Email."`
}

type employeeAddress struct {
	Country string `json:"country,omitempty" jsonschema:"Country."`
	State   string `json:"state,omitempty" jsonschema:"State."`
	ZipCode string `json:"zip_code,omitempty" jsonschema:"Zip code."`
	City    string `json:"city,omitempty" jsonschema:"City."`
	Street  string `json:"street,omitempty" jsonschema:"Street."`
}

type employeeCreate struct {
	employeeContact
	employeeAddress
	md.EmployeeMeta
	EmployeeMap cu.IM `json:"employee_map,omitempty" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
}

type employeeUpdate struct {
	Code string `json:"code" jsonschema:"Database independent unique key. Required when updating an existing employee."`
	employeeContact
	employeeAddress
	md.EmployeeMeta
	EmployeeMap cu.IM `json:"employee_map,omitempty" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
}

type employeeParameter struct {
	Code   string `json:"code,omitempty" jsonschema:"Database independent unique key."`
	Tag    string `json:"tag,omitempty" jsonschema:"Tag."`
	Limit  int64  `json:"limit,omitempty" jsonschema:"Limit."`
	Offset int64  `json:"offset,omitempty" jsonschema:"Offset."`
}

func EmployeeSchema() (ms *ModelSchema) {
	return &ModelSchema{
		Name:   "employee",
		Prefix: "EMP",
		CreateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[employeeCreate](nil); err == nil {
				schema.Properties["start_date"].Type = "string"
				schema.Properties["start_date"].Format = "date"
				schema.Properties["start_date"].Default = []byte(`"` + time.Now().Format("2006-01-02") + `"`)
				schema.Properties["end_date"].Type = "string"
				schema.Properties["end_date"].Format = "date"
				schema.Properties["employee_map"].Default = []byte(`{}`)
				schema.Properties["tags"].Default = []byte(`[]`)
				schema.Required = []string{"first_name", "surname", "start_date"}
			}
			return schema
		},
		UpdateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[employeeUpdate](nil); err == nil {
				schema.Properties["start_date"].Type = "string"
				schema.Properties["start_date"].Format = "date"
				schema.Properties["end_date"].Type = "string"
				schema.Properties["end_date"].Format = "date"
				schema.Properties["employee_map"].Default = []byte(`{}`)
				schema.Required = []string{"code"}
			}
			return schema
		},
		QueryInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			schema, _ = jsonschema.For[employeeParameter](nil)
			return schema
		},
		QueryOutputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[md.Employee](nil); err == nil {
				schema.Description = "Employee data"
				schema.AdditionalProperties = &jsonschema.Schema{}
			}
			return &jsonschema.Schema{
				Type:  "object",
				Items: schema,
			}
		},
		LoadData: employeeLoadData,
		LoadList: func(rows []cu.IM) (items any, err error) {
			var employees []md.Employee = []md.Employee{}
			err = cu.ConvertToType(rows, &employees)
			return employees, err
		},
		PrimaryFields: []string{"id", "code", "employee_meta", "employee_map"},
		Required:      []string{"first_name", "surname", "start_date"},
	}
}

func employeeLoadData(data any) (modelData, metaData any, err error) {
	var employee md.Employee = md.Employee{
		Address: md.Address{
			Tags:       []string{},
			AddressMap: cu.IM{},
		},
		Contact: md.Contact{
			Tags:       []string{},
			ContactMap: cu.IM{},
		},
		Events: []md.Event{},
		EmployeeMeta: md.EmployeeMeta{
			Tags: []string{},
		},
		EmployeeMap: cu.IM{},
	}
	err = cu.ConvertToType(data, &employee)
	return employee, employee.EmployeeMeta, err
}

func employeeCreateHandler(ctx context.Context, req *mcp.CallToolRequest, inputData employeeCreate) (result *mcp.CallToolResult, response UpdateResponseData, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)

	if inputData.FirstName == "" || inputData.Surname == "" || inputData.StartDate == "" {
		return result, UpdateResponseData{}, errors.New("first name, surname and start date are required")
	}

	values := cu.IM{}
	if inputData.EmployeeMeta.Tags == nil {
		inputData.EmployeeMeta.Tags = []string{}
	}

	var contact md.Contact = md.Contact{
		Tags:       []string{},
		ContactMap: cu.IM{},
	}
	if err = cu.ConvertToType(inputData.employeeContact, &contact); err == nil {
		ut.ConvertByteToIMData(contact, values, "contact")
	}

	var address md.Address = md.Address{
		Tags:       []string{},
		AddressMap: cu.IM{},
	}
	if err = cu.ConvertToType(inputData.employeeAddress, &address); err == nil {
		ut.ConvertByteToIMData(address, values, "address")
	}

	ut.ConvertByteToIMData([]md.Event{}, values, "events")
	ut.ConvertByteToIMData(inputData.EmployeeMeta, values, "employee_meta")
	ut.ConvertByteToIMData(inputData.EmployeeMap, values, "employee_map")

	var rows []cu.IM
	var employeeID int64
	var code string
	if employeeID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "employee"}); err == nil {
		if rows, err = ds.StoreDataGet(cu.IM{"id": employeeID, "model": "employee"}, true); err == nil {
			code = cu.ToString(rows[0]["code"], "")
		}
	}
	response = UpdateResponseData{
		Model: "employee",
		Code:  code,
		ID:    employeeID,
	}

	return result, response, err
}

func employeeUpdateHandler(ctx context.Context, req *mcp.CallToolRequest, inputData cu.IM) (result *mcp.CallToolResult, response UpdateResponseData, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)

	code := cu.ToString(inputData["code"], "")
	if code == "" {
		return nil, UpdateResponseData{}, fmt.Errorf("code is required")
	}

	var rows []cu.IM
	if rows, err = ds.StoreDataGet(cu.IM{
		"fields": []string{"*"}, "model": "employee", "code": code}, true); err != nil {
		return nil, UpdateResponseData{}, errors.New("invalid code: " + code)
	}
	updateID := cu.ToInteger(rows[0]["id"], 0)

	setValue := func(data cu.IM, excludeFields []string) (values cu.IM, dirty bool) {
		for key, value := range inputData {
			if _, found := data[key]; found && !slices.Contains(excludeFields, key) {
				data[key] = value
				dirty = true
			}
		}
		return data, dirty
	}

	values := cu.IM{}
	if addressValues, dirty := setValue(cu.ToIM(rows[0]["address"], cu.IM{}), []string{"tags", "notes"}); dirty {
		var address md.Address = md.Address{Tags: []string{}, AddressMap: cu.IM{}}
		if err = cu.ConvertToType(addressValues, &address); err == nil {
			ut.ConvertByteToIMData(address, values, "address")
		}
	}
	if contactValues, dirty := setValue(cu.ToIM(rows[0]["contact"], cu.IM{}), []string{"tags", "notes"}); dirty {
		var contact md.Contact = md.Contact{Tags: []string{}, ContactMap: cu.IM{}}
		if err = cu.ConvertToType(contactValues, &contact); err == nil {
			ut.ConvertByteToIMData(contact, values, "contact")
		}
	}
	if employeeMeta, dirty := setValue(cu.ToIM(rows[0]["employee_meta"], cu.IM{}), []string{}); dirty {
		var meta md.EmployeeMeta = md.EmployeeMeta{Tags: []string{}}
		if err = cu.ConvertToType(employeeMeta, &meta); err == nil {
			ut.ConvertByteToIMData(meta, values, "employee_meta")
		}
	}

	mapValues := cu.ToIM(inputData["employee_map"], cu.IM{})
	if len(mapValues) > 0 {
		ut.ConvertByteToIMData(cu.MergeIM(cu.ToIM(rows[0]["employee_map"], cu.IM{}), mapValues), values, "employee_map")
	}

	_, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "employee", IDKey: updateID})

	return result, UpdateResponseData{
		Model: "employee",
		Code:  code,
		ID:    updateID,
	}, err
}
