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
	toolDataMap["nervatura_rate_create"] = McpTool{
		Tool: mcp.Tool{
			Name:        "nervatura_rate_create",
			Title:       "Create a new rate",
			Description: "Create a new rate. Related tools: place(bank rate), currency.",
			Meta: mcp.Meta{
				"scopes": []string{"setting"},
			},
		},
		ModelSchema: RateSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, rateCreateHandler)
		},
	}
	toolDataMap["nervatura_rate_query"] = McpTool{
		Tool: mcp.Tool{
			Name:        "nervatura_rate_query",
			Title:       "Query rates by parameters",
			Description: "Query rates by parameters. The result is all rates that match the filter criteria.",
			Meta: mcp.Meta{
				"scopes": []string{"setting"},
			},
		},
		ModelSchema: RateSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelQuery)
		},
	}
	toolDataMap["nervatura_rate_update"] = McpTool{
		Tool: mcp.Tool{
			Name:        "nervatura_rate_update",
			Title:       "Update a rate by code",
			Description: "Update a rate by code. When modifying, only the specified values change. Related tools: place(bank rate), currency.",
			Meta: mcp.Meta{
				"scopes": []string{"setting"},
			},
		},
		ModelSchema: RateSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelUpdate)
		},
	}
	toolDataMap["nervatura_rate_delete"] = McpTool{
		Tool:        createDeleteTool("nervatura_rate_delete", "rate", mcp.Meta{"scopes": []string{"setting"}}),
		ModelSchema: RateSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelDelete)
		},
	}
}

type rateCreate struct {
	RateType     string `json:"rate_type" jsonschema:"Rate type. Enum values. Required when creating a new rate."`
	RateDate     string `json:"rate_date" jsonschema:"Rate date. ISO 8601 date. Required when creating a new rate. Example: 2025-01-01"`
	PlaceCode    string `json:"place_code" jsonschema:"Optional reference to place.code (bank rate). Example: PLA1731101982N123"`
	CurrencyCode string `json:"currency_code" jsonschema:"Currency code. Required when creating a new rate. Example: EUR"`
	md.RateMeta
	RateMap cu.IM `json:"rate_map,omitempty" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
}

type rateUpdate struct {
	Code         string `json:"code" jsonschema:"Database independent unique key. Required when updating an existing rate."`
	RateType     string `json:"rate_type" jsonschema:"Rate type. Enum values."`
	RateDate     string `json:"rate_date" jsonschema:"Rate date. ISO 8601 date. Example: 2025-01-01"`
	PlaceCode    string `json:"place_code" jsonschema:"Optional reference to place.code (bank rate). Example: PLA1731101982N123"`
	CurrencyCode string `json:"currency_code" jsonschema:"Currency code. Example: EUR"`
	md.RateMeta
	RateMap cu.IM `json:"rate_map,omitempty" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
}

type rateParameter struct {
	Code         string `json:"code,omitempty" jsonschema:"Database independent unique key."`
	RateType     string `json:"rate_type,omitempty" jsonschema:"Rate type. Enum values."`
	RateDate     string `json:"rate_date,omitempty" jsonschema:"Rate date. ISO 8601 date."`
	PlaceCode    string `json:"place_code,omitempty" jsonschema:"Optional reference to place.code (bank rate). Example: PLA1731101982N123"`
	CurrencyCode string `json:"currency_code,omitempty" jsonschema:"Currency code."`
	Tag          string `json:"tag,omitempty" jsonschema:"Tag."`
	Limit        int64  `json:"limit,omitempty" jsonschema:"Limit."`
	Offset       int64  `json:"offset,omitempty" jsonschema:"Offset."`
}

func RateSchema() (ms *ModelSchema) {
	return &ModelSchema{
		Name:   "rate",
		Prefix: "RAT",
		CreateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[rateCreate](nil); err == nil {
				schema.Properties["rate_type"].Type = "string"
				schema.Properties["rate_type"].Enum = ut.ToAnyArray(md.RateType(0).Keys())
				schema.Properties["rate_type"].Default = []byte(`"` + md.RateTypeRate.String() + `"`)
				schema.Properties["rate_map"].Default = []byte(`{}`)
				schema.Properties["tags"].Default = []byte(`[]`)
				schema.Properties["rate_date"].Format = "date"
				schema.Required = []string{"rate_type", "rate_date", "currency_code"}
			}
			return schema
		},
		UpdateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[rateUpdate](nil); err == nil {
				schema.Properties["rate_type"].Type = "string"
				schema.Properties["rate_type"].Enum = ut.ToAnyArray(md.RateType(0).Keys())
				schema.Properties["rate_map"].Default = []byte(`{}`)
				schema.Properties["rate_date"].Format = "date"
				schema.Required = []string{"code"}
			}
			return schema
		},
		QueryInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			schema, _ = jsonschema.For[rateParameter](nil)
			return schema
		},
		QueryOutputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[md.Rate](nil); err == nil {
				schema.Description = "Rate data"
				schema.AdditionalProperties = &jsonschema.Schema{}
			}
			return &jsonschema.Schema{
				Type:  "object",
				Items: schema,
			}
		},
		LoadData: rateLoadData,
		LoadList: func(rows []cu.IM) (items any, err error) {
			var rates []md.Rate = []md.Rate{}
			err = cu.ConvertToType(rows, &rates)
			return rates, err
		},
		PrimaryFields: []string{"id", "code", "rate_type", "rate_date", "place_code", "currency_code", "rate_meta", "rate_map"},
		Required:      []string{"rate_type", "rate_date", "currency_code"},
	}
}

func rateLoadData(data any) (modelData, metaData any, err error) {
	var rate md.Rate = md.Rate{
		RateType: md.RateTypeRate,
		RateDate: time.Now().Format(time.RFC3339),
		RateMeta: md.RateMeta{
			Tags: []string{},
		},
		RateMap: cu.IM{},
	}
	err = cu.ConvertToType(data, &rate)
	return rate, rate.RateMeta, err
}

func rateCreateHandler(ctx context.Context, req *mcp.CallToolRequest, inputData rateCreate) (result *mcp.CallToolResult, response UpdateResponseData, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)

	if inputData.RateDate == "" || inputData.CurrencyCode == "" {
		return result, UpdateResponseData{}, errors.New("rate date and currency code are required")
	}

	values := cu.IM{
		"rate_date":     inputData.RateDate,
		"rate_type":     inputData.RateType,
		"currency_code": inputData.CurrencyCode,
	}
	if inputData.PlaceCode != "" {
		values["place_code"] = inputData.PlaceCode
	}
	if inputData.RateMeta.Tags == nil {
		inputData.RateMeta.Tags = []string{}
	}

	ut.ConvertByteToIMData(inputData.RateMeta, values, "rate_meta")
	ut.ConvertByteToIMData(inputData.RateMap, values, "rate_map")

	var rows []cu.IM
	var rateID int64
	var code string
	if rateID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "rate"}); err == nil {
		if rows, err = ds.StoreDataGet(cu.IM{"id": rateID, "model": "rate"}, true); err == nil {
			code = cu.ToString(rows[0]["code"], "")
		}
	}
	response = UpdateResponseData{
		Model: "rate",
		Code:  code,
		ID:    rateID,
	}

	return result, response, err
}
