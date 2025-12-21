package mcp

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

func init() {
	toolDataMap["nervatura_email_send"] = McpTool{
		Tool: mcp.Tool{
			Name:        "nervatura_email_send",
			Title:       "Send an email with attachments",
			Description: "Send an email with attachments to the specified email addresses. The result is the email sent successfully.",
			Meta: mcp.Meta{
				"scopes": []string{"customer", "product", "employee", "project", "tool", "offer", "order", "invoice", "worksheet", "rent"},
			},
		},
		ModelSchema: EmailSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, emailSendHandler)
		},
	}
}

func EmailSchema() (ms *ModelSchema) {
	return &ModelSchema{
		Name: "email",
		QueryInputSchema: func(scope string) (schema *jsonschema.Schema) {
			return &jsonschema.Schema{
				Type: "object",
				Properties: map[string]*jsonschema.Schema{
					"recipients": {Type: "array", Items: &jsonschema.Schema{Type: "string"},
						Description: "The email addresses of the recipients. Required. Example: [\"sample@company.com\"]",
						Examples:    []any{`["sample@company.com"]`},
					},
					"subject": {Type: "string",
						Description: "The subject of the email. Required.",
						Examples:    []any{`Email subject`},
					},
					"text": {Type: "string",
						Description: "The text of the email. Optional.",
						Examples:    []any{`Email text`},
					},
					"html": {Type: "string",
						Description: "The HTML of the email. Optional.",
						Examples:    []any{`Email HTML`},
					},
					"code": {Type: "string", MinLength: ut.AnyPointer(12),
						Description: "The unique model key of the attached report. Optional. Example: CUS1731101982N123",
						Examples:    []any{`CUS1731101982N123`, `PRD1731101982N123`}},
					"report_key": {Type: "string",
						Description: fmt.Sprintf("The unique key of the attached %s report template. If not specified, the default report template will be used. Example: ntr_invoice_en", scope),
						Examples:    []any{`ntr_invoice_en`, `ntr_customer_en`}},
					"orientation": {Type: "string", Enum: []any{"p", "l"},
						Description: "The orientation of the report",
						Examples:    []any{`p`, `l`},
						Default:     []byte(`"p"`)},
					"size": {Type: "string", Enum: []any{"a3", "a4", "a5", "letter", "legal"},
						Description: "The size of the report",
						Examples:    []any{`a3`, `a4`, `a5`, `letter`, `legal`},
						Default:     []byte(`"a4"`)},
				},
				Required: []string{"recipients", "subject"},
			}
		},
		QueryOutputSchema: func(scope string) (schema *jsonschema.Schema) {
			return &jsonschema.Schema{
				Type: "object",
				Properties: map[string]*jsonschema.Schema{
					"result": {Type: "string",
						Description: "The result of the email sending",
						Examples:    []any{`OK`, `ERROR`},
						Default:     []byte(`"OK"`),
					},
					"message": {Type: "string",
						Description: "The message of the email sending",
						Examples:    []any{`Email sent successfully`, `Email sending failed`},
						Default:     []byte(`"Email sent successfully"`),
					},
				},
			}
		},
	}
}

func emailSendHandler(ctx context.Context, req *mcp.CallToolRequest, parameters cu.IM) (result *mcp.CallToolResult, response any, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)

	subject := cu.ToString(parameters["subject"], "")
	if subject == "" {
		return nil, nil, errors.New("subject is required")
	}
	code := cu.ToString(parameters["code"], "XXX")
	reportKey := cu.ToString(parameters["report_key"], "")

	if reportKey == "" && code != "XXX" {
		if reportKey, err = getDefaultReportKey(ctx, code); err != nil {
			return nil, nil, err
		}
	}

	recipients := []cu.IM{}
	for _, recipient := range ut.ToStringArray(parameters["recipients"]) {
		recipients = append(recipients, cu.IM{"email": recipient})
	}
	attachments := []cu.IM{}
	if code != "XXX" && reportKey != "" {
		attachments = append(attachments, cu.IM{
			"report_key":  reportKey,
			"code":        code,
			"filename":    code + ".pdf",
			"orientation": cu.ToString(parameters["orientation"], "p"),
			"size":        cu.ToString(parameters["size"], "a4"),
		})
	}
	options := cu.IM{
		"email": cu.IM{
			"recipients":  recipients,
			"subject":     subject,
			"text":        cu.ToString(parameters["text"], ""),
			"html":        cu.ToString(parameters["html"], ""),
			"attachments": attachments,
		},
	}

	response = cu.IM{"result": "OK", "message": "Email sent successfully"}
	if _, err = ds.SendEmail(options); err != nil {
		response = cu.IM{"result": "ERROR", "message": "Email sending failed"}
	}

	return result, response, err
}
