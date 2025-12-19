package mcp

import (
	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"

	cu "github.com/nervatura/component/pkg/util"
)

func init() {
	toolDataMap["nervatura_stock_query"] = McpTool{
		Tool: mcp.Tool{
			Name:        "nervatura_stock_query",
			Title:       "Query inventory by parameters",
			Description: "Query inventory by parameters. The result is all inventory that match the filter criteria.",
			Meta: mcp.Meta{
				"scopes": []string{"stock"},
			},
		},
		ModelSchema: StockSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, modelQuery)
		},
	}
}

type stockData struct {
	PlaceCode   string  `json:"place_code" jsonschema:"Place code."`
	PlaceName   string  `json:"place_name" jsonschema:"Place name."`
	ProductCode string  `json:"product_code" jsonschema:"Product code."`
	ProductName string  `json:"product_name" jsonschema:"Product name."`
	Unit        string  `json:"unit" jsonschema:"Unit."`
	BatchNo     string  `json:"batch_no" jsonschema:"Batch number."`
	Qty         float64 `json:"qty" jsonschema:"Quantity."`
	Posdate     string  `json:"posdate" jsonschema:"Inventory date."`
}

type stockParameter struct {
	PlaceCode   string  `json:"place_code,omitempty" jsonschema:"Place code."`
	PlaceName   string  `json:"place_name,omitempty" jsonschema:"Place name."`
	ProductCode string  `json:"product_code,omitempty" jsonschema:"Product code."`
	ProductName string  `json:"product_name,omitempty" jsonschema:"Product name."`
	Qty         float64 `json:"qty,omitempty" jsonschema:"Quantity."`
	Tag         string  `json:"tag,omitempty" jsonschema:"Tag."`
	Limit       int64   `json:"limit,omitempty" jsonschema:"Limit."`
	Offset      int64   `json:"offset,omitempty" jsonschema:"Offset."`
}

func StockSchema() (ms *ModelSchema) {
	return &ModelSchema{
		Name:       "stock",
		Prefix:     "",
		CustomFrom: "movement_stock",
		QueryInputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			schema, _ = jsonschema.For[stockParameter](nil)
			return schema
		},
		QueryOutputSchema: func(scope string) (schema *jsonschema.Schema) {
			schema = &jsonschema.Schema{}
			var err error
			if schema, err = jsonschema.For[stockData](nil); err == nil {
				schema.Description = "Inventory"
				schema.AdditionalProperties = &jsonschema.Schema{}
			}
			return &jsonschema.Schema{
				Type:  "object",
				Items: schema,
			}
		},
		LoadList: func(rows []cu.IM) (items any, err error) {
			var stocks []stockData = []stockData{}
			err = cu.ConvertToType(rows, &stocks)
			return stocks, err
		},
	}
}
