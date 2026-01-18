package mcp

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

func init() {
	toolDataMap["nervatura_place_create"] = McpTool{
		Tool: mcp.Tool{
			Name:        "nervatura_place_create",
			Title:       "Create a new place (warehouse, bank, cash, etc.)",
			Description: "Create a new place (warehouse, bank, cash, etc.). Related tools: contact, event.",
			Meta: mcp.Meta{
				"scopes": []string{"place"},
			},
		},
		ModelSchema: PlaceSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, placeCreateHandler)
		},
	}
	toolDataMap["nervatura_place_query"] = McpTool{
		Tool: mcp.Tool{
			Name:        "nervatura_place_query",
			Title:       "Query places by parameters",
			Description: "Query places by parameters. The result is all places that match the filter criteria.",
			Meta: mcp.Meta{
				"scopes": []string{"place"},
			},
		},
		ModelSchema: PlaceSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelQuery)
		},
	}
	toolDataMap["nervatura_place_update"] = McpTool{
		Tool: mcp.Tool{
			Name:        "nervatura_place_update",
			Title:       "Update a place by code",
			Description: "Update a place by code. When modifying, only the specified values change. Related tools: contact, event.",
			Meta: mcp.Meta{
				"scopes": []string{"place"},
			},
		},
		ModelSchema: PlaceSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, placeUpdateHandler)
		},
	}
	toolDataMap["nervatura_place_delete"] = McpTool{
		Tool: createDeleteTool("nervatura_place_delete", "place",
			mcp.Meta{"scopes": []string{"place"}}),
		ModelSchema: PlaceSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelDelete)
		},
	}
}

type placeAddress struct {
	Country string `json:"country,omitempty" jsonschema:"Country."`
	State   string `json:"state,omitempty" jsonschema:"State."`
	ZipCode string `json:"zip_code,omitempty" jsonschema:"Zip code."`
	City    string `json:"city,omitempty" jsonschema:"City."`
	Street  string `json:"street,omitempty" jsonschema:"Street."`
}

type placeCreate struct {
	PlaceType    md.PlaceType `json:"place_type" jsonschema:"Place type. Enum values. Example: PLACE_WAREHOUSE"`
	PlaceName    string       `json:"place_name" jsonschema:"Place name. Example: Main warehouse"`
	CurrencyCode string       `json:"currency_code" jsonschema:"Optional reference to currency.code (PLACE_BANK, PLACE_CASH). Example: EUR"`
	placeAddress
	md.PlaceMeta
	PlaceMap cu.IM `json:"place_map,omitempty" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
}

type placeUpdate struct {
	Code         string       `json:"code" jsonschema:"Database independent unique key. Required when updating an existing place."`
	PlaceType    md.PlaceType `json:"place_type" jsonschema:"Place type. Enum values. Example: PLACE_WAREHOUSE"`
	PlaceName    string       `json:"place_name" jsonschema:"Place name. Example: Main warehouse"`
	CurrencyCode string       `json:"currency_code" jsonschema:"Optional reference to currency.code (PLACE_BANK, PLACE_CASH). Example: EUR"`
	placeAddress
	md.PlaceMeta
	PlaceMap cu.IM `json:"place_map,omitempty" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
}

type placeParameter struct {
	Code         string       `json:"code,omitempty" jsonschema:"Database independent unique key."`
	PlaceType    md.PlaceType `json:"place_type,omitempty" jsonschema:"Place type. Enum values. Example: PLACE_WAREHOUSE"`
	PlaceName    string       `json:"like_place_name,omitempty" jsonschema:"Place name. It is not case sensitive and partial values ​​can be specified. Example: Main warehouse"`
	CurrencyCode string       `json:"currency_code,omitempty" jsonschema:"Optional reference to currency.code (PLACE_BANK, PLACE_CASH). Example: EUR"`
	Tag          string       `json:"tag,omitempty" jsonschema:"Tag."`
	Limit        int64        `json:"limit,omitempty" jsonschema:"Limit."`
	Offset       int64        `json:"offset,omitempty" jsonschema:"Offset."`
}

func PlaceSchema() (ms *ModelSchema) {
	return &ModelSchema{
		Name:   "place",
		Prefix: "PLA",
		CreateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[placeCreate](nil); err == nil {
				schema.Properties["place_type"].Type = "string"
				schema.Properties["place_type"].Enum = ut.ToAnyArray(md.PlaceType(0).Keys())
				schema.Properties["place_type"].Default = []byte(`"` + md.PlaceTypeWarehouse.String() + `"`)
				schema.Properties["place_map"].Default = []byte(`{}`)
				schema.Properties["tags"].Default = []byte(`[]`)
				schema.Required = []string{"place_type", "place_name"}
			}
			return schema
		},
		UpdateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[placeUpdate](nil); err == nil {
				schema.Properties["place_type"].Type = "string"
				schema.Properties["place_type"].Enum = ut.ToAnyArray(md.PlaceType(0).Keys())
				schema.Properties["place_map"].Default = []byte(`{}`)
				schema.Required = []string{"code"}
			}
			return schema
		},
		QueryInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			schema, _ = jsonschema.For[placeParameter](nil)
			return schema
		},
		QueryOutputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[md.Place](nil); err == nil {
				schema.Description = "Place data"
				schema.AdditionalProperties = &jsonschema.Schema{}
			}
			return &jsonschema.Schema{
				Type:  "object",
				Items: schema,
			}
		},
		LoadData: placeLoadData,
		LoadList: func(rows []cu.IM) (items any, err error) {
			var places []md.Place = []md.Place{}
			err = cu.ConvertToType(rows, &places)
			return places, err
		},
		PrimaryFields: []string{"id", "code", "place_type", "place_name", "currency_code", "place_meta", "place_map"},
		Required:      []string{"place_type", "place_name"},
	}
}

func placeLoadData(data any) (modelData, metaData any, err error) {
	var place md.Place = md.Place{
		PlaceType: md.PlaceTypeWarehouse,
		Address: md.Address{
			Tags:       []string{},
			AddressMap: cu.IM{},
		},
		Contacts: []md.Contact{},
		Events:   []md.Event{},
		PlaceMeta: md.PlaceMeta{
			Tags: []string{},
		},
		PlaceMap: cu.IM{},
	}
	err = cu.ConvertToType(data, &place)
	return place, place.PlaceMeta, err
}

func placeCreateHandler(ctx context.Context, req *mcp.CallToolRequest, inputData placeCreate) (result *mcp.CallToolResult, response UpdateResponseData, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)

	if inputData.PlaceName == "" {
		return result, UpdateResponseData{}, errors.New("place name is required")
	}

	values := cu.IM{
		"place_type": inputData.PlaceType.String(),
		"place_name": inputData.PlaceName,
	}
	if inputData.CurrencyCode != "" {
		values["currency_code"] = inputData.CurrencyCode
	}
	if inputData.PlaceMeta.Tags == nil {
		inputData.PlaceMeta.Tags = []string{}
	}

	var address md.Address = md.Address{
		Tags:       []string{},
		AddressMap: cu.IM{},
	}
	if err = cu.ConvertToType(inputData.placeAddress, &address); err == nil {
		ut.ConvertByteToIMData(address, values, "address")
	}

	ut.ConvertByteToIMData([]md.Contact{}, values, "contacts")
	ut.ConvertByteToIMData([]md.Event{}, values, "events")
	ut.ConvertByteToIMData(inputData.PlaceMeta, values, "place_meta")
	ut.ConvertByteToIMData(inputData.PlaceMap, values, "place_map")

	var rows []cu.IM
	var placeID int64
	var code string
	if placeID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "place"}); err == nil {
		if rows, err = ds.StoreDataGet(cu.IM{"id": placeID, "model": "place"}, true); err == nil {
			code = cu.ToString(rows[0]["code"], "")
		}
	}
	response = UpdateResponseData{
		Model: "place",
		Code:  code,
		ID:    placeID,
	}

	return result, response, err
}

func placeUpdateHandler(ctx context.Context, req *mcp.CallToolRequest, inputData cu.IM) (result *mcp.CallToolResult, response UpdateResponseData, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)

	code := cu.ToString(inputData["code"], "")
	if code == "" {
		return nil, UpdateResponseData{}, fmt.Errorf("code is required")
	}

	var rows []cu.IM
	if rows, err = ds.StoreDataGet(cu.IM{
		"fields": []string{"*"}, "model": "place", "code": code}, true); err != nil {
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
	if placeMeta, dirty := setValue(cu.ToIM(rows[0]["place_meta"], cu.IM{}), []string{}); dirty {
		var meta md.PlaceMeta = md.PlaceMeta{Tags: []string{}}
		if err = cu.ConvertToType(placeMeta, &meta); err == nil {
			ut.ConvertByteToIMData(meta, values, "place_meta")
		}
	}

	mapValues := cu.ToIM(inputData["place_map"], cu.IM{})
	if len(mapValues) > 0 {
		ut.ConvertByteToIMData(cu.MergeIM(cu.ToIM(rows[0]["place_map"], cu.IM{}), mapValues), values, "place_map")
	}

	_, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "place", IDKey: updateID})

	return result, UpdateResponseData{
		Model: "place",
		Code:  code,
		ID:    updateID,
	}, err
}
