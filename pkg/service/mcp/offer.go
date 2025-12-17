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
	toolDataMap["nervatura_offer_create"] = McpTool{
		Tool: mcp.Tool{
			Name:        "nervatura_offer_create",
			Title:       "Create a new offer",
			Description: "Create a new offer. Related tools: item.",
			Meta: mcp.Meta{
				"scopes": []string{"offer"},
			},
		},
		ModelSchema: OfferSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, offerCreateHandler)
		},
	}
	toolDataMap["nervatura_offer_query"] = McpTool{
		Tool: mcp.Tool{
			Name:        "nervatura_offer_query",
			Title:       "Query offers by parameters",
			Description: "Query offers by parameters. The result is all offers that match the filter criteria.",
			Meta: mcp.Meta{
				"scopes": []string{"offer"},
			},
		},
		ModelSchema: OfferSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelQuery)
		},
	}
	toolDataMap["nervatura_offer_update"] = McpTool{
		Tool: mcp.Tool{
			Name:        "nervatura_offer_update",
			Title:       "Update a offer by code",
			Description: "Update a offer by code. When modifying, only the specified values change. Related tools: item.",
			Meta: mcp.Meta{
				"scopes": []string{"offer"},
			},
		},
		ModelSchema: OfferSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelUpdate)
		},
	}
	toolDataMap["nervatura_offer_delete"] = McpTool{
		Tool:        createDeleteTool("nervatura_offer_delete", "offer", mcp.Meta{"scopes": []string{"offer"}}),
		ModelSchema: OfferSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelDelete)
		},
	}
}

type offerCreate struct {
	TransDate    string `json:"trans_date" jsonschema:"Offer date. Required when creating a new offer. Example: 2025-01-01"`
	Direction    string `json:"direction" jsonschema:"Transaction direction. Enum values. Required when creating a new offer. Example: DIRECTION_OUT"`
	TransCode    string `json:"trans_code" jsonschema:"Other transaction (invoice, receipt, offer, order, worksheet, rent etc.) reference. Optional. Example: ORD1731101982N123"`
	CustomerCode string `json:"customer_code" jsonschema:"Customer reference. Required when creating a new offer. Example: CUS1731101982N123"`
	EmployeeCode string `json:"employee_code" jsonschema:"Employee reference. Optional. Example: EMP1731101982N123"`
	ProjectCode  string `json:"project_code" jsonschema:"Project reference. Optional. Example: PRJ1731101982N123"`
	CurrencyCode string `json:"currency_code" jsonschema:"Currency iso code. Required when creating a new offer. Example: USD"`
	offerMeta
	TransMap cu.IM `json:"trans_map,omitempty" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
}

type offerUpdate struct {
	Code         string `json:"code" jsonschema:"Database independent unique key. Required when updating an existing offer."`
	TransDate    string `json:"trans_date,omitempty" jsonschema:"Offer date."`
	TransCode    string `json:"trans_code,omitempty" jsonschema:"Other transaction (invoice, receipt, offer, order, worksheet, rent etc.) reference. Optional. Example: ORD1731101982N123"`
	CustomerCode string `json:"customer_code,omitempty" jsonschema:"Customer reference. Example: CUS1731101982N123"`
	EmployeeCode string `json:"employee_code,omitempty" jsonschema:"Employee reference. Example: EMP1731101982N123"`
	ProjectCode  string `json:"project_code,omitempty" jsonschema:"Project reference. Example: PRJ1731101982N123"`
	CurrencyCode string `json:"currency_code,omitempty" jsonschema:"Currency iso code. Example: USD"`
	offerMeta
	TransMap cu.IM `json:"trans_map,omitempty" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
}

type offerMeta struct {
	DueTime       string   `json:"due_time,omitempty" jsonschema:"Validity date. Required when creating a new offer. Example: 2025-01-01"`
	RefNumber     string   `json:"ref_number,omitempty" jsonschema:"Ref number. Example: REF1731101982N123"`
	PaidType      string   `json:"paid_type,omitempty" jsonschema:"Paid type. Enum values. Example: PAID_TYPE_CASH"`
	Paid          bool     `json:"paid,omitempty" jsonschema:"Released"`
	Rate          float64  `json:"rate,omitempty" jsonschema:"Payment days"`
	Closed        bool     `json:"closed,omitempty" jsonschema:"Closed offer"`
	TransState    string   `json:"trans_state,omitempty" jsonschema:"Trans state. Enum values. Example: TRANS_STATE_OK"`
	Notes         string   `json:"notes" jsonschema:"Notes"`
	InternalNotes string   `json:"internal_notes,omitempty" jsonschema:"Internal notes"`
	ReportNotes   string   `json:"report_notes,omitempty" jsonschema:"Report notes."`
	Tags          []string `json:"tags,omitempty" jsonschema:"Tags. Example: [TAG1, TAG2]"`
}

type offerQuery struct {
	Id           int64     `json:"id,omitempty" jsonschema:"Database primary key."`
	Code         string    `json:"code,omitempty" jsonschema:"Database independent unique key."`
	TransDate    string    `json:"trans_date,omitempty" jsonschema:"Offer date."`
	Direction    string    `json:"direction,omitempty" jsonschema:"Transaction direction. Enum values. Example: DIRECTION_OUT"`
	TransCode    string    `json:"trans_code,omitempty" jsonschema:"Other transaction (invoice, receipt, offer, order, worksheet, rent etc.) reference. Optional. Example: ORD1731101982N123"`
	CustomerCode string    `json:"customer_code,omitempty" jsonschema:"Customer reference. Example: CUS1731101982N123"`
	EmployeeCode string    `json:"employee_code,omitempty" jsonschema:"Employee reference. Example: EMP1731101982N123"`
	ProjectCode  string    `json:"project_code,omitempty" jsonschema:"Project reference. Example: PRJ1731101982N123"`
	CurrencyCode string    `json:"currency_code,omitempty" jsonschema:"Currency iso code. Example: USD"`
	Amount       float64   `json:"amount,omitempty" jsonschema:"Total amount."`
	TransMeta    offerMeta `json:"trans_meta,omitempty" jsonschema:"Trans metadata."`
	TransMap     cu.IM     `json:"trans_map,omitempty" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
}

type offerParameter struct {
	Code         string `json:"code,omitempty" jsonschema:"Database independent unique key."`
	TransType    string `json:"trans_type,omitempty" jsonschema:"Transaction type. Enum values."`
	Direction    string `json:"direction,omitempty" jsonschema:"Transaction direction. Enum values."`
	TransDate    string `json:"trans_date,omitempty" jsonschema:"Offer date."`
	CustomerCode string `json:"customer_code,omitempty" jsonschema:"Customer reference. Example: CUS1731101982N123"`
	Tag          string `json:"tag,omitempty" jsonschema:"Tag."`
	Limit        int64  `json:"limit,omitempty" jsonschema:"Limit."`
	Offset       int64  `json:"offset,omitempty" jsonschema:"Offset."`
}

func OfferSchema() (ms *ModelSchema) {
	return &ModelSchema{
		Name:       "trans",
		CustomFrom: "trans t left join(select trans_code as invoice_code, sum(amount) as amount from item_view group by trans_code) i on t.code = i.invoice_code",
		CustomParameters: func(params cu.IM) cu.IM {
			params["filter"] = "trans_type = '" + md.TransTypeOffer.String() + "'"
			return params
		},
		Prefix: "OFF",
		CreateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[offerCreate](nil); err == nil {
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
			if schema, err = jsonschema.For[offerUpdate](nil); err == nil {
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
			schema, _ = jsonschema.For[offerParameter](nil)
			return schema
		},
		QueryOutputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[offerQuery](nil); err == nil {
				schema.Description = "Offer data"
				schema.AdditionalProperties = &jsonschema.Schema{}
			}
			return &jsonschema.Schema{
				Type:  "object",
				Items: schema,
			}
		},
		LoadData: offerLoadData,
		LoadList: func(rows []cu.IM) (items any, err error) {
			var offers []offerQuery = []offerQuery{}
			err = cu.ConvertToType(rows, &offers)
			return offers, err
		},
		Validate:      offerValidate,
		PrimaryFields: []string{"id", "code", "trans_type", "trans_date", "direction", "trans_code", "customer_code", "employee_code", "project_code", "currency_code", "trans_meta", "trans_map"},
		Required:      []string{"trans_date", "direction", "customer_code", "currency_code", "due_time", "paid_type"},
	}
}

func offerLoadData(data any) (modelData, metaData any, err error) {
	var offer md.Trans = md.Trans{
		TransType: md.TransTypeOffer,
		Direction: md.DirectionOut,
		TransMeta: md.TransMeta{
			Status:     md.TransStatusNormal,
			TransState: md.TransStateOK,
			Tags:       []string{},
		},
		TransMap: cu.IM{},
	}
	err = cu.ConvertToType(data, &offer)
	return offer, offer.TransMeta, err
}

func offerCreateHandler(ctx context.Context, req *mcp.CallToolRequest, inputData offerCreate) (result *mcp.CallToolResult, response UpdateResponseData, err error) {
	tokenInfo := req.Extra.TokenInfo
	user := tokenInfo.Extra["user"].(md.Auth)
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)

	if inputData.TransDate == "" || inputData.Direction == "" || inputData.CustomerCode == "" || inputData.CurrencyCode == "" {
		return result, UpdateResponseData{}, errors.New("trans date, direction, customer code and currency code are required")
	}

	values := cu.IM{
		"trans_type": md.TransTypeOffer.String(),
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
	cu.ConvertToType(inputData.offerMeta, &transMeta)

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

func offerValidate(ctx context.Context, input cu.IM) (data cu.IM, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)

	code := cu.ToString(input["code"], "")
	var rows []cu.IM
	if rows, err = ds.StoreDataGet(cu.IM{"model": "trans", "code": code}, true); err != nil {
		return data, err
	}
	transMeta := cu.ToIM(rows[0]["trans_meta"], cu.IM{})
	if cu.ToBoolean(transMeta["closed"], false) {
		return data, errors.New("offer is not updatable because it is closed")
	}
	return input, nil
}
