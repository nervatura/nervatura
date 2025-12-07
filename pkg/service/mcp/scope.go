package mcp

import cu "github.com/nervatura/component/pkg/util"

var ScopeInstruction cu.SM = cu.SM{
	"root":     "Nervatura is a business management framework. To use the tools, resources, and prompts found on MCP endpoints, first check out resources/README.md!",
	"all":      "This scope is used for access to all models. It allows you to query, create, update and delete data for all models.",
	"public":   "This scope is used for public access to the API. Additional endpoints: /mcp/customer, /mcp/all.",
	"customer": "This tools are used for access to the customer data. It allows you to query, create, update and delete customers. Related tools: contact, address, event.",
}
