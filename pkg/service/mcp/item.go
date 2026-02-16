package mcp

import (
	"context"
	"errors"
	"math"
	"slices"
	"strings"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

func init() {
	toolDataMap["nervatura_item_create"] = McpTool{
		Tool: mcp.Tool{
			Name:        "nervatura_item_create",
			Title:       "Create a new document item",
			Description: "Create a new document item. Related tools: offer, order, invoice, receipt, worksheet, rental.",
			Meta: mcp.Meta{
				"scopes": []string{"offer", "order", "invoice", "worksheet", "rent"},
			},
		},
		ModelSchema: ItemSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, itemCreateHandler)
		},
	}
	toolDataMap["nervatura_item_query"] = McpTool{
		Tool: mcp.Tool{
			Name:        "nervatura_item_query",
			Title:       "Query items by parameters",
			Description: "Query items by parameters. The result is all items that match the filter criteria.",
			Meta: mcp.Meta{
				"scopes": []string{"offer", "order", "invoice", "worksheet", "rent"},
			},
		},
		ModelSchema: ItemSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelQuery)
		},
	}
	toolDataMap["nervatura_item_update"] = McpTool{
		Tool: mcp.Tool{
			Name:        "nervatura_item_update",
			Title:       "Update a document item by code",
			Description: "Update a document item by code. When modifying, only the specified values change. Related tools: offer, order, invoice, receipt, worksheet, rental.",
			Meta: mcp.Meta{
				"scopes": []string{"offer", "order", "invoice", "worksheet", "rent"},
			},
		},
		ModelSchema: ItemSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelUpdate)
		},
	}
	toolDataMap["nervatura_item_delete"] = McpTool{
		Tool:        createDeleteTool("nervatura_item_delete", "item", mcp.Meta{"scopes": []string{"offer", "order", "invoice", "worksheet", "rent"}}),
		ModelSchema: ItemSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelDelete)
		},
	}
}

type itemMeta struct {
	Unit        string  `json:"unit,omitempty" jsonschema:"Unit. If empty, the product unit will be used. Example: piece"`
	Qty         float64 `json:"qty,omitempty" jsonschema:"Quantity."`
	InputType   string  `json:"input_type,omitempty" jsonschema:"Input type. Enum values."`
	Amount      float64 `json:"amount,omitempty" jsonschema:"Input amount."`
	Discount    float64 `json:"discount,omitempty" jsonschema:"Discount."`
	Description string  `json:"description,omitempty" jsonschema:"Description."`
	Deposit     bool    `json:"deposit,omitempty" jsonschema:"Deposit."`
	OwnStock    float64 `json:"own_stock,omitempty" jsonschema:"Own stock."`
	ActionPrice bool    `json:"action_price,omitempty" jsonschema:"Action price."`
	// Additional tags for the item
	Tags []string `json:"tags,omitempty" jsonschema:"Additional tags for the item. The value is an array of strings. Example: [TAG1, TAG2]"`
}

type itemCreate struct {
	TransCode   string `json:"trans_code" jsonschema:"Transaction code. Required when creating a new document item. Example: INV1731101982N123"`
	ProductCode string `json:"product_code" jsonschema:"Product code. Required when creating a new document item. Example: PRD1731101982N123"`
	TaxCode     string `json:"tax_code,omitempty" jsonschema:"Tax code. If empty, the product tax code will be used. Example: VAT20"`
	itemMeta
	ItemMap cu.IM `json:"item_map,omitempty" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
}

type itemUpdate struct {
	Code        string `json:"code" jsonschema:"Database independent unique key. Required when updating an existing document item."`
	TransCode   string `json:"trans_code" jsonschema:"Offer, order, invoice, receipt, worksheet, rental reference code. Example: INV1731101982N123"`
	ProductCode string `json:"product_code" jsonschema:"Product code. Example: PRD1731101982N123"`
	TaxCode     string `json:"tax_code" jsonschema:"Tax code. Example: VAT20"`
	itemMeta
	ItemMap cu.IM `json:"item_map,omitempty" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
}

type itemParameter struct {
	Code        string `json:"code,omitempty" jsonschema:"Database independent unique key."`
	TransType   string `json:"trans_type,omitempty" jsonschema:"Transaction type. Enum values."`
	TransCode   string `json:"trans_code,omitempty" jsonschema:"Offer, order, invoice, receipt, worksheet, rental reference code. Example: INV1731101982N123"`
	ProductCode string `json:"product_code,omitempty" jsonschema:"Product code. Example: PRD1731101982N123"`
	TaxCode     string `json:"tax_code,omitempty" jsonschema:"Tax code. Example: VAT20"`
	Tag         string `json:"tag,omitempty" jsonschema:"Tag."`
	Limit       int64  `json:"limit,omitempty" jsonschema:"Limit."`
	Offset      int64  `json:"offset,omitempty" jsonschema:"Offset."`
}

func ItemSchema() (ms *ModelSchema) {
	return &ModelSchema{
		Name:   "item",
		Prefix: "ITM",
		CustomParameters: func(params cu.IM) cu.IM {
			query := md.Query{
				Fields: []string{"code"},
				From:   "trans",
				Filters: []md.Filter{
					{Field: "deleted", Comp: "==", Value: false},
				},
			}
			if _, found := params["trans_type"]; found {
				query.Filters = append(query.Filters, md.Filter{Field: "trans_type", Comp: "==", Value: cu.ToString(params["trans_type"], "")})
				delete(params, "trans_type")
			}
			params["in_trans_code"] = query
			return params
		},
		CreateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[itemCreate](nil); err == nil {
				schema.Properties["input_type"].Type = "string"
				schema.Properties["input_type"].Enum = []any{"DEFAULT_PRICE", "FX_PRICE", "NET_AMOUNT", "AMOUNT"}
				schema.Properties["input_type"].Default = []byte(`"DEFAULT_PRICE"`)
				schema.Properties["input_type"].Description = "Input type. Enum values. If DEFAULT_PRICE, the valid default price set in the database will be set."
				schema.Properties["amount"].Description = "Input amount. If input type is DEFAULT_PRICE, this field is ignored."
				schema.Properties["discount"].Description = "Discount. If input type is DEFAULT_PRICE, this field is ignored."
				schema.Properties["description"].Description = "Description. If empty, the product description will be used."
				schema.Properties["item_map"].Default = []byte(`{}`)
				schema.Properties["tags"].Default = []byte(`[]`)
				schema.Properties["qty"].Default = []byte(`1`)
				schema.Required = []string{"trans_code", "product_code", "input_type", "qty"}
			}
			return schema
		},
		UpdateInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[itemUpdate](nil); err == nil {
				schema.Properties["input_type"].Type = "string"
				schema.Properties["input_type"].Enum = []any{"FX_PRICE", "NET_AMOUNT", "AMOUNT"}
				schema.Required = []string{"code"}
			}
			return schema
		},
		QueryInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[itemParameter](nil); err == nil {
				schema.Description = "Query items by parameters. The result is all items that match the filter criteria."
				schema.Properties["trans_type"].Type = "string"
				schema.Properties["trans_type"].Enum = []any{
					md.TransTypeOrder.String(), md.TransTypeOffer.String(), md.TransTypeInvoice.String(), md.TransTypeReceipt.String(),
					md.TransTypeWorksheet.String(), md.TransTypeRent.String(),
				}
				if slices.Contains([]string{"invoice", "receipt", "order", "offer", "worksheet", "rent"}, scope) {
					schema.Properties["trans_type"].Enum = []any{"TRANS_" + strings.ToUpper(scope)}
					schema.Required = []string{"trans_type"}
					schema.Properties["trans_type"].Default = []byte(`"` + "TRANS_" + strings.ToUpper(scope) + `"`)
				}
			}
			return schema
		},
		QueryOutputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[md.Item](nil); err == nil {
				schema.Description = "Item data"
				schema.AdditionalProperties = &jsonschema.Schema{}
			}
			return &jsonschema.Schema{
				Type:  "object",
				Items: schema,
			}
		},
		LoadData: itemLoadData,
		LoadList: func(rows []cu.IM) (items any, err error) {
			var values []md.Item = []md.Item{}
			err = cu.ConvertToType(rows, &values)
			return values, err
		},
		Validate:      itemValidate,
		PrimaryFields: []string{"id", "code", "trans_code", "product_code", "tax_code", "item_meta", "item_map"},
		Required:      []string{"trans_code", "product_code", "input_type"},
	}
}

func itemLoadData(data any) (modelData, metaData any, err error) {
	var item md.Item = md.Item{
		ItemMeta: md.ItemMeta{
			Tags: []string{},
		},
		ItemMap: cu.IM{},
	}
	err = cu.ConvertToType(data, &item)
	return item, item.ItemMeta, err
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func calcItemPrice(ctx context.Context, inputData cu.IM) (data cu.IM, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)

	code := cu.ToString(inputData["code"], "")
	transCode := cu.ToString(inputData["trans_code"], "")
	taxCode := cu.ToString(inputData["tax_code"], "")
	inputType := cu.ToString(inputData["input_type"], "FX_PRICE")
	inputAmount := cu.ToFloat(inputData["amount"], 0)
	qty := cu.ToFloat(inputData["qty"], 0)
	discount := cu.ToFloat(inputData["discount"], 0)
	rate := float64(0)
	digit := uint(0)

	if code == "" && (transCode == "" || taxCode == "") {
		return data, errors.New("trans code and tax code are required")
	}

	var rows []cu.IM
	if code != "" {
		if rows, err = ds.StoreDataGet(cu.IM{
			"fields": []string{"i.trans_code", "i.tax_code", "i.product_code", "tx.rate_value", "cv.digit", "i.fx_price", "i.qty", "i.discount"},
			"model": `item_view i inner join tax_view tx on i.tax_code = tx.code inner join trans t on i.trans_code = t.code 
			inner join currency_view cv on t.currency_code = cv.code`,
			"i.code": code}, true); err != nil {
			return data, err
		}
		digit = uint(cu.ToInteger(rows[0]["digit"], 0))
		inputAmount = cu.ToFloat(inputAmount, cu.ToFloat(rows[0]["fx_price"], 0))
		qty = cu.ToFloat(qty, cu.ToFloat(rows[0]["qty"], 0))
		discount = cu.ToFloat(discount, cu.ToFloat(rows[0]["discount"], 0))
		if taxCode == "" {
			rate = cu.ToFloat(rows[0]["rate_value"], 0)
		}
	}
	if code == "" {
		if rows, err = ds.StoreDataGet(cu.IM{
			"fields": []string{"cv.digit"},
			"model":  `trans t inner join currency_view cv on t.currency_code = cv.code`,
			"t.code": transCode}, true); err != nil {
			return data, err
		}
		digit = uint(cu.ToInteger(rows[0]["digit"], 0))
	}
	if taxCode != "" {
		if rows, err = ds.StoreDataGet(cu.IM{"fields": []string{"rate_value"}, "model": `tax_view`, "code": taxCode}, true); err != nil {
			return data, err
		}
		rate = cu.ToFloat(rows[0]["rate_value"], 0)
	}

	var netAmount, vatAmount, amount, fxPrice float64
	typeMap := map[string]func(){
		"FX_PRICE": func() {
			fxPrice = inputAmount
			netAmount = roundFloat(fxPrice*(1-discount/100)*qty, digit)
			vatAmount = roundFloat(fxPrice*(1-discount/100)*qty*rate, digit)
			amount = roundFloat(netAmount+vatAmount, digit)
		},
		"NET_AMOUNT": func() {
			netAmount = inputAmount
			if qty != 0 {
				fxPrice = roundFloat(netAmount/(1-discount/100)/qty, digit)
				vatAmount = roundFloat(netAmount*rate, digit)
			}
			amount = roundFloat(netAmount+vatAmount, digit)
		},
		"AMOUNT": func() {
			amount = inputAmount
			if qty != 0 {
				netAmount = roundFloat(amount/(1+rate), digit)
				vatAmount = roundFloat(amount-netAmount, digit)
				fxPrice = roundFloat(netAmount/(1-discount/100)/qty, digit)
			}
		},
	}
	if fn, ok := typeMap[inputType]; ok {
		fn()
	}

	return cu.IM{
		"net_amount": netAmount, "vat_amount": vatAmount, "amount": amount, "fx_price": fxPrice,
	}, err
}

func itemCreateHandler(ctx context.Context, req *mcp.CallToolRequest, inputData itemCreate) (result *mcp.CallToolResult, response UpdateResponseData, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)

	if inputData.TransCode == "" || inputData.ProductCode == "" {
		return result, UpdateResponseData{}, errors.New("trans code and product code are required")
	}

	taxCode := inputData.TaxCode
	description := inputData.Description
	unit := inputData.Unit
	discount := inputData.Discount
	inputType := inputData.InputType
	inputPrice := inputData.Amount

	var rows []cu.IM
	if taxCode == "" || description == "" || unit == "" {
		if rows, err = ds.StoreDataGet(cu.IM{"model": `product_view`, "code": inputData.ProductCode}, true); err != nil {
			return result, UpdateResponseData{}, err
		}
		taxCode = cu.ToString(taxCode, cu.ToString(rows[0]["tax_code"], ""))
		description = cu.ToString(description, cu.ToString(rows[0]["product_name"], ""))
		unit = cu.ToString(unit, cu.ToString(rows[0]["unit"], ""))
	}

	if inputType == "DEFAULT_PRICE" {
		if rows, err = ds.StoreDataGet(cu.IM{"model": `trans`, "code": inputData.TransCode}, true); err != nil {
			return result, UpdateResponseData{}, err
		}
		customerCode := cu.ToString(rows[0]["customer_code"], "")
		currencyCode := cu.ToString(rows[0]["currency_code"], "")
		var priceData cu.IM
		if priceData, err = ds.ProductPrice(
			cu.IM{"product_code": inputData.ProductCode, "currency_code": currencyCode, "customer_code": customerCode, "qty": inputData.Qty},
		); err != nil {
			return result, UpdateResponseData{}, err
		}
		discount = cu.ToFloat(discount, cu.ToFloat(priceData["discount"], 0))
		inputPrice = cu.ToFloat(priceData["price"], 0)
		inputType = "FX_PRICE"
	}

	var prices cu.IM
	if prices, err = calcItemPrice(ctx, cu.IM{
		"code": "", "trans_code": inputData.TransCode, "tax_code": taxCode,
		"input_type": inputType, "amount": inputPrice, "qty": inputData.Qty, "discount": discount,
	}); err != nil {
		return result, UpdateResponseData{}, err
	}

	values := cu.IM{
		"trans_code":   inputData.TransCode,
		"product_code": inputData.ProductCode,
		"tax_code":     taxCode,
	}

	metaData := md.ItemMeta{
		Unit:        unit,
		Qty:         inputData.Qty,
		FxPrice:     cu.ToFloat(prices["fx_price"], 0),
		NetAmount:   cu.ToFloat(prices["net_amount"], 0),
		Discount:    discount,
		VatAmount:   cu.ToFloat(prices["vat_amount"], 0),
		Amount:      cu.ToFloat(prices["amount"], 0),
		Description: description,
		Deposit:     inputData.Deposit,
		OwnStock:    inputData.OwnStock,
		ActionPrice: inputData.ActionPrice,
		Tags:        inputData.Tags,
	}
	if metaData.Tags == nil {
		metaData.Tags = []string{}
	}

	ut.ConvertByteToIMData(metaData, values, "item_meta")
	ut.ConvertByteToIMData(inputData.ItemMap, values, "item_map")

	var itemID int64
	var code string
	if itemID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "item"}); err == nil {
		if rows, err = ds.StoreDataGet(cu.IM{"id": itemID, "model": "item"}, true); err == nil {
			code = cu.ToString(rows[0]["code"], "")
		}
	}
	response = UpdateResponseData{
		Model: "item",
		Code:  code,
		ID:    itemID,
	}

	return result, response, err
}

func itemValidate(ctx context.Context, input cu.IM) (data cu.IM, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	isPrice := func() bool {
		for key := range input {
			if slices.Contains([]string{"input_type", "qty", "discount", "amount", "tax_code"}, key) {
				return true
			}
		}
		return false
	}

	var rows []cu.IM
	if rows, err = ds.StoreDataQuery(md.Query{
		Fields: []string{"*"},
		From:   `trans_view`,
		Filters: []md.Filter{
			{Field: "code", Comp: "in", Value: md.Query{
				Fields:  []string{"trans_code"},
				From:    "item",
				Filters: []md.Filter{{Field: "code", Comp: "==", Value: input["code"]}},
			}},
		},
	}, true); err != nil {
		return input, err
	}
	if cu.ToString(rows[0]["status"], "") == md.TransStatusDeleted.String() || cu.ToBoolean(rows[0]["closed"], false) {
		return data, errors.New("item is not updatable because the transaction is deleted or closed")
	}

	if isPrice() {
		var prices cu.IM
		if prices, err = calcItemPrice(ctx, input); err != nil {
			return input, err
		}
		input = cu.MergeIM(input, prices)
		delete(input, "input_type")
	}
	return input, nil
}
