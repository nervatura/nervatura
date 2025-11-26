package mcp

import (
	"context"
	"fmt"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

var deleteCodeTool = mcp.Tool{
	Name:        "nervatura_delete_code",
	Title:       "Delete model data by code",
	Description: "Delete data by model code. It returns the success of the operation or an error message.",
	InputSchema: &jsonschema.Schema{
		Type: "object",
		Properties: map[string]*jsonschema.Schema{
			"code": {Type: "string", MinLength: ut.AnyPointer(12),
				Description: "The unique key of the result model data. Example: CUS1731101982N123",
				Examples:    []any{`CUS1731101982N123`, `PRD1731101982N123`}},
		},
		Required: []string{"code"},
	},
	OutputSchema: &jsonschema.Schema{
		Type: "object",
		Properties: map[string]*jsonschema.Schema{
			"success": {Type: "boolean", Description: "The success of the operation."},
		},
	},
}

func deleteCode(ctx context.Context, req *mcp.CallToolRequest, inputData map[string]any) (result *mcp.CallToolResult, response any, err error) {
	extra := req.GetExtra()
	ds := extra.TokenInfo.Extra["ds"].(*api.DataStore)

	code := cu.ToString(inputData["code"], "???")
	prefix := code[:3]

	var sm *ModelSchema
	if sm, err = getModelSchemaByPrefix(prefix); err != nil {
		return nil, nil, err
	}

	var res *mcp.ElicitResult
	if res, err = req.Session.Elicit(ctx, &mcp.ElicitParams{
		RequestedSchema: &jsonschema.Schema{Type: "object",
			Description: "Confirm the deletion of the data.",
			Properties: map[string]*jsonschema.Schema{
				"confirm": {Type: "string",
					Description: "Confirm the deletion of the data.",
					Enum:        []any{`YES`, `NO`}},
			},
			Required: []string{"confirm"},
		},
	}); err != nil {
		return nil, nil, fmt.Errorf("eliciting failed: %v", err)
	}
	confirm := (cu.ToString(res.Content["confirm"], "NO") == "YES")

	if confirm {
		err = ds.DataDelete(sm.Name, 0, code)
	}

	return &mcp.CallToolResult{
		StructuredContent: cu.IM{"success": err == nil && confirm},
		IsError:           err != nil,
	}, nil, err
}
