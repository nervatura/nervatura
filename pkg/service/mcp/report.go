package mcp

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
	st "github.com/nervatura/nervatura/v6/pkg/static"
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
		Scopes: []string{"customer", "product", "invoice"},
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
					"url_result": {Type: "boolean",
						Description: "If true, the result will be a URL to the report result. If false, the result will be the report data in the selected output format",
						Default:     []byte(`true`)},
					"inline": {Type: "boolean",
						Description: "If true, the result will be displayed in the browser or downloaded as a file",
						Default:     []byte(`true`)},
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
						Description: "The report data in the selected output format or a URL to the report result",
						Examples:    []any{`iVBORw0KGgoAAAANSUhEUgAA...`},
					},
				},
			}
		},
	}
}

func getDefaultReportKey(ctx context.Context, code string) (reportKey string, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)

	var ms *ModelSchema
	if ms, err = getModelSchemaByPrefix(code[:3]); err != nil {
		return reportKey, err
	}
	values := cu.IM{
		"fields": []string{"report_key"}, "model": "config_report", "report_type": strings.ToUpper(ms.Name)}
	var rows []cu.IM
	if ms.Name == "trans" {
		if rows, err = ds.StoreDataGet(cu.IM{"model": "trans", "code": code}, true); err != nil {
			return reportKey, err
		}
		values["trans_type"] = strings.TrimPrefix(cu.ToString(rows[0]["trans_type"], ""), "TRANS_")
		values["direction"] = cu.ToString(rows[0]["direction"], "")
	}
	if rows, err = ds.StoreDataGet(values, true); err != nil {
		return reportKey, errors.New("missing report key for " + code)
	}
	reportKey = cu.ToString(rows[0]["report_key"], "")
	return reportKey, nil
}

func reportQueryHandler(ctx context.Context, req *mcp.CallToolRequest, parameters cu.IM) (result *mcp.CallToolResult, response any, err error) {
	tokenInfo := req.Extra.TokenInfo

	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	session := ctx.Value(md.SessionServiceCtxKey).(*api.SessionService)
	config := ctx.Value(md.ConfigCtxKey).(cu.IM)
	code := cu.ToString(parameters["code"], "XXX")
	urlResult := cu.ToBoolean(parameters["url_result"], false)

	if cu.ToString(parameters["report_key"], "") == "" {
		if parameters["report_key"], err = getDefaultReportKey(ctx, code); err != nil {
			return nil, nil, err
		}
	}

	var report cu.IM
	if urlResult {
		sessionID := req.Session.ID()
		output := cu.ToString(parameters["output"], "pdf")
		if output == "base64" {
			output = "pdf"
		}
		baseURL := cu.ToString(config["NT_PUBLIC_HOST"], "")
		response = cu.IM{
			"content_type": "application/pdf",
			"data":         fmt.Sprintf(baseURL+st.ClientPath+"/session/export/report/modal/%s?output=%s&inline=%s", sessionID, output, cu.ToString(parameters["inline"], "true")),
		}
		clientData := ct.Client{
			Ticket: ct.Ticket{
				SessionID: sessionID,
				User:      cu.IM{},
				Expiry:    tokenInfo.Expiration,
				Database:  cu.ToString(tokenInfo.Extra["alias"], ""),
			},
			BaseComponent: ct.BaseComponent{
				Data: cu.IM{
					"modal": cu.IM{
						"data": cu.IM{
							"code":        code,
							"template":    cu.ToString(parameters["report_key"], ""),
							"orientation": cu.ToString(parameters["orientation"], ""),
							"paper_size":  cu.ToString(parameters["size"], ""),
						},
					},
				},
			},
		}
		session.SaveSession(sessionID, &clientData)
		return result, response, nil
	}
	if report, err = ds.GetReport(parameters); err == nil {
		response = cu.IM{
			"content_type": report["content_type"],
			"data":         report["template"],
		}
	}

	return result, response, err
}
