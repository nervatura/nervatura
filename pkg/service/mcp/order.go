package mcp

import (
	"context"
	"errors"
	"time"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

func init() {
	toolDataMap["nervatura_order_create"] = ToolData{
		Tool: mcp.Tool{
			Name:        "nervatura_order_create",
			Title:       "Create a new order",
			Description: "Create a new order. Related tools: item.",
			Meta: mcp.Meta{
				"scopes": []string{"order"},
			},
		},
		ModelSchema: OrderSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, orderCreateHandler)
		},
	}
	toolDataMap["nervatura_order_query"] = ToolData{
		Tool: mcp.Tool{
			Name:        "nervatura_order_query",
			Title:       "Query orders by parameters",
			Description: "Query orders by parameters. The result is all orders that match the filter criteria.",
			Meta: mcp.Meta{
				"scopes": []string{"order"},
			},
		},
		ModelSchema: OrderSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelQuery)
		},
	}
	toolDataMap["nervatura_order_update"] = ToolData{
		Tool: mcp.Tool{
			Name:        "nervatura_order_update",
			Title:       "Update a order by code",
			Description: "Update a order by code. When modifying, only the specified values change. Related tools: item.",
			Meta: mcp.Meta{
				"scopes": []string{"order"},
			},
		},
		ModelSchema: OrderSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelUpdate)
		},
	}
	toolDataMap["nervatura_order_delete"] = ToolData{
		Tool:        createDeleteTool("nervatura_order_delete", "order", mcp.Meta{"scopes": []string{"order"}}),
		ModelSchema: OrderSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelDelete)
		},
	}
}

type orderCreate struct {
	TransDate    string `json:"trans_date" jsonschema:"Order date. Required when creating a new order. Example: 2025-01-01"`
	Direction    string `json:"direction" jsonschema:"Transaction direction. Enum values. Required when creating a new order. Example: DIRECTION_OUT"`
	TransCode    string `json:"trans_code" jsonschema:"Other transaction (invoice, receipt, offer, order, worksheet, rent etc.) reference. Optional. Example: ORD1731101982N123"`
	CustomerCode string `json:"customer_code" jsonschema:"Customer reference. Required when creating a new order. Example: CUS1731101982N123"`
	EmployeeCode string `json:"employee_code" jsonschema:"Employee reference. Optional. Example: EMP1731101982N123"`
	ProjectCode  string `json:"project_code" jsonschema:"Project reference. Optional. Example: PRJ1731101982N123"`
	CurrencyCode string `json:"currency_code" jsonschema:"Currency iso code. Required when creating a new order. Example: USD"`
	orderMeta
	TransMap cu.IM `json:"trans_map,omitempty" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
}

type orderUpdate struct {
	Code         string `json:"code" jsonschema:"Database independent unique key. Required when updating an existing order."`
	TransDate    string `json:"trans_date,omitempty" jsonschema:"Order date."`
	TransCode    string `json:"trans_code,omitempty" jsonschema:"Other transaction (invoice, receipt, offer, order, worksheet, rent etc.) reference. Optional. Example: ORD1731101982N123"`
	CustomerCode string `json:"customer_code,omitempty" jsonschema:"Customer reference. Example: CUS1731101982N123"`
	EmployeeCode string `json:"employee_code,omitempty" jsonschema:"Employee reference. Example: EMP1731101982N123"`
	ProjectCode  string `json:"project_code,omitempty" jsonschema:"Project reference. Example: PRJ1731101982N123"`
	CurrencyCode string `json:"currency_code,omitempty" jsonschema:"Currency iso code. Example: USD"`
	orderMeta
	TransMap cu.IM `json:"trans_map,omitempty" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
}

type orderMeta struct {
	DueTime       string   `json:"due_time,omitempty" jsonschema:"Delivery date. Required when creating a new order. Example: 2025-01-01"`
	RefNumber     string   `json:"ref_number,omitempty" jsonschema:"Ref number. Example: REF1731101982N123"`
	PaidType      string   `json:"paid_type,omitempty" jsonschema:"Paid type. Enum values. Example: PAID_TYPE_CASH"`
	Paid          bool     `json:"paid,omitempty" jsonschema:"Released"`
	Rate          float64  `json:"rate,omitempty" jsonschema:"Payment days"`
	Closed        bool     `json:"closed,omitempty" jsonschema:"Closed order"`
	TransState    string   `json:"trans_state,omitempty" jsonschema:"Trans state. Enum values. Example: TRANS_STATE_OK"`
	Notes         string   `json:"notes" jsonschema:"Notes"`
	InternalNotes string   `json:"internal_notes,omitempty" jsonschema:"Internal notes"`
	ReportNotes   string   `json:"report_notes,omitempty" jsonschema:"Report notes."`
	Tags          []string `json:"tags,omitempty" jsonschema:"Tags. Example: [TAG1, TAG2]"`
}

type orderQuery struct {
	Id           int64     `json:"id,omitempty" jsonschema:"Database primary key."`
	Code         string    `json:"code,omitempty" jsonschema:"Database independent unique key."`
	TransDate    string    `json:"trans_date,omitempty" jsonschema:"Order date."`
	Direction    string    `json:"direction,omitempty" jsonschema:"Transaction direction. Enum values. Example: DIRECTION_OUT"`
	TransCode    string    `json:"trans_code,omitempty" jsonschema:"Other transaction (invoice, receipt, offer, order, worksheet, rent etc.) reference. Optional. Example: ORD1731101982N123"`
	CustomerCode string    `json:"customer_code,omitempty" jsonschema:"Customer reference. Example: CUS1731101982N123"`
	EmployeeCode string    `json:"employee_code,omitempty" jsonschema:"Employee reference. Example: EMP1731101982N123"`
	ProjectCode  string    `json:"project_code,omitempty" jsonschema:"Project reference. Example: PRJ1731101982N123"`
	CurrencyCode string    `json:"currency_code,omitempty" jsonschema:"Currency iso code. Example: USD"`
	Amount       float64   `json:"amount,omitempty" jsonschema:"Total amount."`
	TransMeta    orderMeta `json:"trans_meta,omitempty" jsonschema:"Trans metadata."`
	TransMap     cu.IM     `json:"trans_map,omitempty" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
}

type orderParameter struct {
	Code         string `json:"code,omitempty" jsonschema:"Database independent unique key."`
	TransType    string `json:"trans_type,omitempty" jsonschema:"Transaction type. Enum values."`
	Direction    string `json:"direction,omitempty" jsonschema:"Transaction direction. Enum values."`
	TransDate    string `json:"trans_date,omitempty" jsonschema:"Order date."`
	CustomerCode string `json:"customer_code,omitempty" jsonschema:"Customer reference. Example: CUS1731101982N123"`
	Tag          string `json:"tag,omitempty" jsonschema:"Tag."`
	Limit        int64  `json:"limit,omitempty" jsonschema:"Limit."`
	Offset       int64  `json:"offset,omitempty" jsonschema:"Offset."`
}

func OrderSchema() (ms *ModelSchema) {
	return &ModelSchema{
		Name:       "trans",
		CustomFrom: "trans t left join(select trans_code as invoice_code, sum(amount) as amount from item_view group by trans_code) i on t.code = i.invoice_code",
		CustomParameters: func(params cu.IM) cu.IM {
			params["filter"] = "trans_type = '" + md.TransTypeOrder.String() + "'"
			return params
		},
		Prefix: "ORD",
		CreateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[orderCreate](nil); err == nil {
				schema.Properties["direction"].Type = "string"
				schema.Properties["direction"].Enum = []any{md.DirectionOut.String(), md.DirectionIn.String()}
				schema.Properties["direction"].Default = []byte(`"` + md.DirectionOut.String() + `"`)
				schema.Properties["trans_date"].Type = "string"
				schema.Properties["trans_date"].Format = "date"
				schema.Properties["trans_date"].Default = []byte(`"` + time.Now().Format("2006-01-02") + `"`)
				schema.Properties["trans_map"].Default = []byte(`{}`)
				schema.Properties["due_time"].Type = "string"
				schema.Properties["due_time"].Format = "date"
				schema.Properties["due_time"].Default = []byte(`"` + time.Now().Format("2006-01-02") + `"`)
				schema.Properties["paid_type"].Type = "string"
				schema.Properties["paid_type"].Enum = ut.ToAnyArray(md.PaidType(0).Keys())
				schema.Properties["paid_type"].Default = []byte(`"` + md.PaidTypeCard.String() + `"`)
				schema.Properties["trans_state"].Type = "string"
				schema.Properties["trans_state"].Enum = ut.ToAnyArray(md.TransState(0).Keys())
				schema.Properties["trans_state"].Default = []byte(`"` + md.TransStateOK.String() + `"`)
				schema.Properties["tags"].Default = []byte(`[]`)
				schema.Required = []string{"trans_date", "direction", "customer_code", "currency_code", "due_time", "paid_type"}
			}
			return schema
		},
		UpdateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[orderUpdate](nil); err == nil {
				schema.Properties["trans_date"].Type = "string"
				schema.Properties["trans_date"].Format = "date"
				schema.Properties["trans_map"].Default = []byte(`{}`)
				schema.Properties["due_time"].Type = "string"
				schema.Properties["due_time"].Format = "date"
				schema.Properties["paid_type"].Type = "string"
				schema.Properties["paid_type"].Enum = ut.ToAnyArray(md.PaidType(0).Keys())
				schema.Properties["trans_state"].Type = "string"
				schema.Properties["trans_state"].Enum = ut.ToAnyArray(md.TransState(0).Keys())
				schema.Required = []string{"code"}
			}
			return schema
		},
		QueryInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			schema, _ = jsonschema.For[orderParameter](nil)
			return schema
		},
		QueryOutputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[orderQuery](nil); err == nil {
				schema.Description = "Order data"
				schema.AdditionalProperties = &jsonschema.Schema{}
			}
			return &jsonschema.Schema{
				Type:  "object",
				Items: schema,
			}
		},
		LoadData: orderLoadData,
		LoadList: func(rows []cu.IM) (items any, err error) {
			var orders []orderQuery = []orderQuery{}
			err = cu.ConvertToType(rows, &orders)
			return orders, err
		},
		Validate:      orderValidate,
		PrimaryFields: []string{"id", "code", "trans_type", "trans_date", "direction", "trans_code", "customer_code", "employee_code", "project_code", "currency_code", "trans_meta", "trans_map"},
		Required:      []string{"trans_date", "direction", "customer_code", "currency_code", "due_time", "paid_type"},
	}
}

func orderLoadData(data any) (modelData, metaData any, err error) {
	var order md.Trans = md.Trans{
		TransType: md.TransTypeOrder,
		Direction: md.DirectionOut,
		TransMeta: md.TransMeta{
			Status:     md.TransStatusNormal,
			TransState: md.TransStateOK,
			Tags:       []string{},
		},
		TransMap: cu.IM{},
	}
	err = cu.ConvertToType(data, &order)
	return order, order.TransMeta, err
}

func orderCreateHandler(ctx context.Context, req *mcp.CallToolRequest, inputData orderCreate) (result *mcp.CallToolResult, response UpdateResponseData, err error) {
	tokenInfo := req.Extra.TokenInfo
	user := tokenInfo.Extra["user"].(md.Auth)
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)

	if inputData.TransDate == "" || inputData.Direction == "" || inputData.CustomerCode == "" || inputData.CurrencyCode == "" {
		return result, UpdateResponseData{}, errors.New("trans date, direction, customer code and currency code are required")
	}

	values := cu.IM{
		"trans_type": md.TransTypeOrder.String(),
		"direction":  inputData.Direction,
		"trans_date": inputData.TransDate,
		"auth_code":  user.Code,
	}

	// Optional fields
	optionalFields := map[string]string{
		"customer_code": inputData.CustomerCode,
		"employee_code": inputData.EmployeeCode,
		"project_code":  inputData.ProjectCode,
		"trans_code":    inputData.TransCode,
		"currency_code": inputData.CurrencyCode,
	}

	for key, value := range optionalFields {
		if value != "" {
			values[key] = value
		}
	}

	transMeta := md.TransMeta{
		DueTime:    md.TimeDateTime{Time: time.Now()},
		Status:     md.TransStatus(md.TransStatusNormal),
		TransState: md.TransState(md.TransStateOK),
		Worksheet:  md.TransMetaWorksheet{},
		Rent:       md.TransMetaRent{},
		Invoice:    md.TransMetaInvoice{},
		Tags:       []string{},
	}
	cu.ConvertToType(inputData.orderMeta, &transMeta)

	ut.ConvertByteToIMData(transMeta, values, "trans_meta")
	ut.ConvertByteToIMData(inputData.TransMap, values, "trans_map")

	var rows []cu.IM
	var transID int64
	var code string
	if transID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "trans"}); err == nil {
		if rows, err = ds.StoreDataGet(cu.IM{"id": transID, "model": "trans"}, true); err == nil {
			code = cu.ToString(rows[0]["code"], "")
		}
	}
	response = UpdateResponseData{
		Model: "trans",
		Code:  code,
		ID:    transID,
	}

	return result, response, err
}

func orderValidate(ctx context.Context, input cu.IM) (data cu.IM, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)

	code := cu.ToString(input["code"], "")
	var rows []cu.IM
	if rows, err = ds.StoreDataGet(cu.IM{"model": "trans", "code": code}, true); err != nil {
		return data, err
	}
	transMeta := cu.ToIM(rows[0]["trans_meta"], cu.IM{})
	if cu.ToBoolean(transMeta["closed"], false) {
		return data, errors.New("order is not updatable because it is closed")
	}
	return input, nil
}
