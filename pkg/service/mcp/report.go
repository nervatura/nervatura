package mcp

import (
	"context"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

type reportParameter struct {
	ReportKey  string `json:"report_key" jsonschema:"Report key."`
	ReportType string `json:"report_type" jsonschema:"Report type."`
	ReportName string `json:"report_name" jsonschema:"Report name."`
	Label      string `json:"label" jsonschema:"Label."`
}

func ReportSchema() (ms *ModelSchema) {
	return &ModelSchema{
		Name:             "report",
		Prefix:           "",
		ResultType:       func() any { return md.ConfigReport{} },
		ResultListType:   func() any { return []md.ConfigReport{} },
		InputSchema:      func() (*jsonschema.Schema, error) { return jsonschema.For[md.ConfigReport](nil) },
		ParameterSchema:  func() (*jsonschema.Schema, error) { return jsonschema.For[reportParameter](nil) },
		ResultSchema:     func() (*jsonschema.Schema, error) { return jsonschema.For[md.ConfigReport](nil) },
		ResultListSchema: func() (*jsonschema.Schema, error) { return jsonschema.For[[]md.ConfigReport](nil) },
		SchemaModify: func(schemaType SchemaType, schema *jsonschema.Schema) {
			schema.Description = "Report templates"
			schema.Properties["report_key"].ReadOnly = true
			schema.Properties["report_type"].Enum = []any{"REPORT", "CUSTOMER", "EMPLOYEE", "PRODUCT", "PROJECT", "TOOL", "TRANS"}
			schema.Properties["report_type"].Default = []byte(`"REPORT"`)
			schema.Properties["trans_type"].Enum = []any{"INVOICE", "RECEIPT", "ORDER", "OFFER", "WORKSHEET", "RENT", "DELIVERY", "INVENTORY", "WAYBILL", "PRODUCTION", "FORMULA", "BANK", "CASH"}
			schema.Properties["direction"].Enum = []any{"OUT", "IN", "TRANSFER"}
			schema.Properties["file_type"].Enum = []any{md.FileTypePDF.String(), md.FileTypeCSV.String()}
			schema.Properties["file_type"].Default = []byte(`"` + md.FileTypePDF.String() + `"`)

		},
		Examples: map[string][]any{
			"report_key":  {`ntr_invoice_en`},
			"report_type": {`TRANS`},
			"trans_type":  {`INVOICE`},
			"direction":   {`OUT`},
			"file_type":   {`FILE_TYPE_PDF`},
		},
		PrimaryFields: []string{"report_key", "report_type", "trans_type", "direction", "file_type", "template"},
		Required:      []string{"report_key", "report_type", "file_type"},
	}
}

var reportDataCodeTool = mcp.Tool{
	Name:        "nervatura_report_data_code",
	Title:       "Report Data Code",
	Description: "Get report data by code.",
	InputSchema: &jsonschema.Schema{
		Type: "object",
		Properties: map[string]*jsonschema.Schema{
			"report_key": {Type: "string",
				Description: "The unique key of the report template. Example: ntr_invoice_en",
				Examples:    []any{`ntr_invoice_en`, `ntr_customer_en`}},
			"code": {Type: "string", MinLength: ut.AnyPointer(12),
				Description: "The unique key of the result model data. Example: CUS1731101982N123",
				Examples:    []any{`CUS1731101982N123`, `PRD1731101982N123`}},
			"output": {Type: "string", Enum: []any{"base64", "xml", "pdf"},
				Description: "The output format",
				Examples:    []any{`base64`, `xml`, `pdf`},
				Default:     []byte(`"base64"`)},
			"orientation": {Type: "string", Enum: []any{"p", "l"},
				Description: "The orientation of the report",
				Examples:    []any{`p`, `l`},
				Default:     []byte(`"p"`)},
			"size": {Type: "string", Enum: []any{"a3", "a4", "a5", "letter", "legal"},
				Description: "The size of the report",
				Examples:    []any{`a3`, `a4`, `a5`, `letter`, `legal`},
				Default:     []byte(`"a4"`)},
		},
		Required: []string{"report_key", "code"},
	},
	OutputSchema: &jsonschema.Schema{
		Type: "object",
		Properties: map[string]*jsonschema.Schema{
			"content_type": {Type: "string",
				Description: "The content type of the report",
				Examples:    []any{`application/pdf`, `application/xml`, `text/csv`},
				Default:     []byte(`"application/pdf"`),
			},
			"data": {Type: "string",
				Description: "The report data in the selected output format",
				Examples:    []any{`iVBORw0KGgoAAAANSUhEUgAA...`},
			},
		},
	},
}

func reportDataCode(ctx context.Context, req *mcp.CallToolRequest, parameters cu.IM) (result *mcp.CallToolResult, response any, err error) {
	extra := req.GetExtra()
	ds := extra.TokenInfo.Extra["ds"].(*api.DataStore)

	var report cu.IM
	if report, err = ds.GetReport(parameters); err == nil {
		response = cu.IM{
			"content_type": report["content_type"],
			"data":         report["template"],
		}
	}

	return result, response, err
}
