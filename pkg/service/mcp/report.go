package mcp

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

func init() {
	toolDataMap["nervatura_report_query"] = ToolData{
		Tool: mcp.Tool{
			Name:        "nervatura_report_query",
			Title:       "Get a PDF report or XML data by parameters",
			Description: "Get a %s PDF report or XML data by parameters. The result is the report data in the selected output format.",
		},
		ModelSchema: ReportSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, reportQueryHandler)
		},
		Scopes: []string{"customer"},
	}
}

func ReportSchema() (ms *ModelSchema) {
	return &ModelSchema{
		Name: "report",
		QueryInputSchema: func(scope string) (schema *jsonschema.Schema) {
			return &jsonschema.Schema{
				Type: "object",
				Properties: map[string]*jsonschema.Schema{
					"code": {Type: "string", MinLength: ut.AnyPointer(12),
						Description: "The unique key of the result model data. Required. Example: CUS1731101982N123",
						Examples:    []any{`CUS1731101982N123`, `PRD1731101982N123`}},
					"report_key": {Type: "string",
						Description: fmt.Sprintf("The unique key of the %s report template. If not specified, the default report template will be used. Example: ntr_invoice_en", scope),
						Examples:    []any{`ntr_invoice_en`, `ntr_customer_en`}},
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
				Required: []string{"code"},
			}
		},
		QueryOutputSchema: func(scope string) (schema *jsonschema.Schema) {
			return &jsonschema.Schema{
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
			}
		},
	}
}

func reportQueryHandler(ctx context.Context, req *mcp.CallToolRequest, parameters cu.IM) (result *mcp.CallToolResult, response any, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	code := cu.ToString(parameters["code"], "XXX")

	if cu.ToString(parameters["report_key"], "") == "" {
		var ms *ModelSchema
		if ms, err = getModelSchemaByPrefix(code[:3]); err != nil {
			return nil, nil, errors.New("invalid code: " + code)
		}
		var rows []cu.IM
		if rows, err = ds.StoreDataGet(cu.IM{
			"fields": []string{"report_key"}, "model": "config_report", "report_type": strings.ToUpper(ms.Name)}, true); err == nil {
			parameters["report_key"] = cu.ToString(rows[0]["report_key"], "")
		}
	}

	var report cu.IM
	if report, err = ds.GetReport(parameters); err == nil {
		response = cu.IM{
			"content_type": report["content_type"],
			"data":         report["template"],
		}
	}

	return result, response, err
}
