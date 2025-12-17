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
	toolDataMap["nervatura_invoice_create"] = ToolData{
		Tool: mcp.Tool{
			Name:        "nervatura_invoice_create",
			Title:       "Create a new invoice",
			Description: "Create a new invoice. Related tools: item.",
			Meta: mcp.Meta{
				"scopes": []string{"invoice"},
			},
		},
		ModelSchema: InvoiceSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, invoiceCreateHandler)
		},
	}
	toolDataMap["nervatura_invoice_query"] = ToolData{
		Tool: mcp.Tool{
			Name:        "nervatura_invoice_query",
			Title:       "Query invoices by parameters",
			Description: "Query invoices by parameters. The result is all invoices that match the filter criteria.",
			Meta: mcp.Meta{
				"scopes": []string{"invoice"},
			},
		},
		ModelSchema: InvoiceSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelQuery)
		},
	}
	toolDataMap["nervatura_invoice_update"] = ToolData{
		Tool: mcp.Tool{
			Name:        "nervatura_invoice_update",
			Title:       "Update a invoice by code",
			Description: "Update a invoice by code. When modifying, only the specified values change. Related tools: item.",
			Meta: mcp.Meta{
				"scopes": []string{"invoice"},
			},
		},
		ModelSchema: InvoiceSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelUpdate)
		},
	}
	toolDataMap["nervatura_invoice_delete"] = ToolData{
		Tool:        createDeleteTool("nervatura_invoice_delete", "invoice", mcp.Meta{"scopes": []string{"invoice"}}),
		ModelSchema: InvoiceSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelDelete)
		},
	}
}

type invoiceCreate struct {
	TransDate    string `json:"trans_date" jsonschema:"Transaction date. Required when creating a new invoice. Example: 2025-01-01"`
	Direction    string `json:"direction" jsonschema:"Transaction direction. Enum values. Required when creating a new invoice. Example: DIRECTION_OUT"`
	TransCode    string `json:"trans_code" jsonschema:"Other transaction (invoice, receipt, offer, order, worksheet, rent etc.) reference. Optional. Example: ORD1731101982N123"`
	CustomerCode string `json:"customer_code" jsonschema:"Customer reference. Required when creating a new invoice. Example: CUS1731101982N123"`
	EmployeeCode string `json:"employee_code" jsonschema:"Employee reference. Optional. Example: EMP1731101982N123"`
	ProjectCode  string `json:"project_code" jsonschema:"Project reference. Optional. Example: PRJ1731101982N123"`
	CurrencyCode string `json:"currency_code" jsonschema:"Currency iso code. Required when creating a new invoice. Example: USD"`
	invoiceMeta
	TransMap cu.IM `json:"trans_map,omitempty" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
}

type invoiceUpdate struct {
	Code         string `json:"code" jsonschema:"Database independent unique key. Required when updating an existing invoice."`
	TransDate    string `json:"trans_date,omitempty" jsonschema:"Transaction date."`
	TransCode    string `json:"trans_code,omitempty" jsonschema:"Other transaction (invoice, receipt, offer, order, worksheet, rent etc.) reference. Optional. Example: ORD1731101982N123"`
	CustomerCode string `json:"customer_code,omitempty" jsonschema:"Customer reference. Example: CUS1731101982N123"`
	EmployeeCode string `json:"employee_code,omitempty" jsonschema:"Employee reference. Example: EMP1731101982N123"`
	ProjectCode  string `json:"project_code,omitempty" jsonschema:"Project reference. Example: PRJ1731101982N123"`
	CurrencyCode string `json:"currency_code,omitempty" jsonschema:"Currency iso code. Example: USD"`
	invoiceMeta
	TransMap cu.IM `json:"trans_map,omitempty" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
}

type invoiceMeta struct {
	DueTime       string              `json:"due_time,omitempty" jsonschema:"Due date. Required when creating a new invoice. Example: 2025-01-01"`
	RefNumber     string              `json:"ref_number,omitempty" jsonschema:"Ref number. Example: REF1731101982N123"`
	PaidType      string              `json:"paid_type,omitempty" jsonschema:"Paid type. Enum values. Example: PAID_TYPE_CASH"`
	Paid          bool                `json:"paid,omitempty" jsonschema:"Paid invoice"`
	Rate          float64             `json:"rate,omitempty" jsonschema:"Rate. Example: 1.0"`
	Closed        bool                `json:"closed,omitempty" jsonschema:"Closed invoice"`
	Status        string              `json:"status,omitempty" jsonschema:"Status. Enum values. Example: TRANS_STATUS_NORMAL"`
	TransState    string              `json:"trans_state,omitempty" jsonschema:"Trans state. Enum values. Example: TRANS_STATE_OK"`
	Notes         string              `json:"notes" jsonschema:"Notes"`
	InternalNotes string              `json:"internal_notes,omitempty" jsonschema:"Internal notes"`
	ReportNotes   string              `json:"report_notes,omitempty" jsonschema:"Report notes."`
	Invoice       md.TransMetaInvoice `json:"invoice,omitempty"`
	Tags          []string            `json:"tags,omitempty" jsonschema:"Tags. Example: [TAG1, TAG2]"`
}

type invoiceQuery struct {
	Id           int64       `json:"id,omitempty" jsonschema:"Database primary key."`
	Code         string      `json:"code,omitempty" jsonschema:"Database independent unique key."`
	TransDate    string      `json:"trans_date,omitempty" jsonschema:"Transaction date."`
	Direction    string      `json:"direction,omitempty" jsonschema:"Transaction direction. Enum values. Example: DIRECTION_OUT"`
	TransCode    string      `json:"trans_code,omitempty" jsonschema:"Other transaction (invoice, receipt, offer, order, worksheet, rent etc.) reference. Optional. Example: ORD1731101982N123"`
	CustomerCode string      `json:"customer_code,omitempty" jsonschema:"Customer reference. Example: CUS1731101982N123"`
	EmployeeCode string      `json:"employee_code,omitempty" jsonschema:"Employee reference. Example: EMP1731101982N123"`
	ProjectCode  string      `json:"project_code,omitempty" jsonschema:"Project reference. Example: PRJ1731101982N123"`
	CurrencyCode string      `json:"currency_code,omitempty" jsonschema:"Currency iso code. Example: USD"`
	Amount       float64     `json:"amount,omitempty" jsonschema:"Total amount."`
	TransMeta    invoiceMeta `json:"trans_meta,omitempty" jsonschema:"Trans metadata."`
	TransMap     cu.IM       `json:"trans_map,omitempty" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
}

type invoiceParameter struct {
	Code         string `json:"code,omitempty" jsonschema:"Database independent unique key."`
	TransType    string `json:"trans_type,omitempty" jsonschema:"Transaction type. Enum values."`
	Direction    string `json:"direction,omitempty" jsonschema:"Transaction direction. Enum values."`
	TransDate    string `json:"trans_date,omitempty" jsonschema:"Transaction date."`
	CustomerCode string `json:"customer_code,omitempty" jsonschema:"Customer reference. Example: CUS1731101982N123"`
	Tag          string `json:"tag,omitempty" jsonschema:"Tag."`
	Limit        int64  `json:"limit,omitempty" jsonschema:"Limit."`
	Offset       int64  `json:"offset,omitempty" jsonschema:"Offset."`
}

func InvoiceSchema() (ms *ModelSchema) {
	return &ModelSchema{
		Name:       "trans",
		CustomFrom: "trans t left join(select trans_code as invoice_code, sum(amount) as amount from item_view group by trans_code) i on t.code = i.invoice_code",
		CustomParameters: func(params cu.IM) cu.IM {
			params["filter"] = "trans_type = '" + md.TransTypeInvoice.String() + "'"
			return params
		},
		Prefix: "INV",
		CreateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[invoiceCreate](nil); err == nil {
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
				delete(schema.Properties, "status")
				delete(schema.Properties, "invoice")
				schema.Properties["tags"].Default = []byte(`[]`)
				schema.Required = []string{"trans_date", "direction", "customer_code", "currency_code", "due_time", "paid_type"}
			}
			return schema
		},
		UpdateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[invoiceUpdate](nil); err == nil {
				schema.Properties["trans_date"].Type = "string"
				schema.Properties["trans_date"].Format = "date"
				schema.Properties["trans_map"].Default = []byte(`{}`)
				schema.Properties["due_time"].Type = "string"
				schema.Properties["due_time"].Format = "date"
				schema.Properties["paid_type"].Type = "string"
				schema.Properties["paid_type"].Enum = ut.ToAnyArray(md.PaidType(0).Keys())
				schema.Properties["trans_state"].Type = "string"
				schema.Properties["trans_state"].Enum = ut.ToAnyArray(md.TransState(0).Keys())
				delete(schema.Properties, "status")
				delete(schema.Properties, "invoice")
				schema.Required = []string{"code"}
			}
			return schema
		},
		QueryInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			schema, _ = jsonschema.For[invoiceParameter](nil)
			return schema
		},
		QueryOutputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[invoiceQuery](nil); err == nil {
				schema.Description = "Invoice data"
				schema.AdditionalProperties = &jsonschema.Schema{}
			}
			return &jsonschema.Schema{
				Type:  "object",
				Items: schema,
			}
		},
		LoadData: invoiceLoadData,
		LoadList: func(rows []cu.IM) (items any, err error) {
			var invoices []invoiceQuery = []invoiceQuery{}
			err = cu.ConvertToType(rows, &invoices)
			return invoices, err
		},
		Validate:      invoiceValidate,
		PrimaryFields: []string{"id", "code", "trans_type", "trans_date", "direction", "trans_code", "customer_code", "employee_code", "project_code", "currency_code", "trans_meta", "trans_map"},
		Required:      []string{"trans_date", "direction", "customer_code", "currency_code", "due_time", "paid_type"},
	}
}

func invoiceLoadData(data any) (modelData, metaData any, err error) {
	var invoice md.Trans = md.Trans{
		TransType: md.TransTypeInvoice,
		Direction: md.DirectionOut,
		TransMeta: md.TransMeta{
			Status:     md.TransStatusNormal,
			TransState: md.TransStateOK,
			Invoice:    md.TransMetaInvoice{},
			Tags:       []string{},
		},
		TransMap: cu.IM{},
	}
	err = cu.ConvertToType(data, &invoice)
	return invoice, invoice.TransMeta, err
}

func invoiceCreateHandler(ctx context.Context, req *mcp.CallToolRequest, inputData invoiceCreate) (result *mcp.CallToolResult, response UpdateResponseData, err error) {
	tokenInfo := req.Extra.TokenInfo
	user := tokenInfo.Extra["user"].(md.Auth)
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)

	if inputData.TransDate == "" || inputData.Direction == "" || inputData.CustomerCode == "" || inputData.CurrencyCode == "" {
		return result, UpdateResponseData{}, errors.New("trans date, direction, customer code and currency code are required")
	}

	values := cu.IM{
		"trans_type": md.TransTypeInvoice.String(),
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
	cu.ConvertToType(inputData.invoiceMeta, &transMeta)

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

func invoiceValidate(ctx context.Context, input cu.IM) (data cu.IM, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)

	code := cu.ToString(input["code"], "")
	var rows []cu.IM
	if rows, err = ds.StoreDataGet(cu.IM{"model": "trans", "code": code}, true); err != nil {
		return data, err
	}
	transMeta := cu.ToIM(rows[0]["trans_meta"], cu.IM{})
	if cu.ToString(transMeta["status"], "") == md.TransStatusDeleted.String() || cu.ToBoolean(transMeta["closed"], false) {
		return data, errors.New("invoice is not updatable because it is deleted or closed")
	}
	return input, nil
}
