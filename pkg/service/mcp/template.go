package mcp

import (
	"context"
	"fmt"
	"net/url"
	"path"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	st "github.com/nervatura/nervatura/v6/pkg/static"
)

var ntrCustomerEnResource = mcp.Resource{
	Name:        "ntr_customer_en",
	Title:       "Customer Sheet",
	Description: "Customer Information Sheet in English",
	MIMEType:    "application/json",
	URI:         "template:ntr_customer_en.json",
}

func templateResource(ctx context.Context, req *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
	u, err := url.Parse(req.Params.URI)
	if err != nil {
		return nil, err
	}
	if u.Scheme != "template" {
		return nil, fmt.Errorf("wrong scheme: %q", u.Scheme)
	}
	key := u.Opaque
	file, err := st.Report.ReadFile(path.Join("template", key))
	if err != nil {
		return nil, err
	}
	return &mcp.ReadResourceResult{
		Contents: []*mcp.ResourceContents{
			{URI: req.Params.URI, MIMEType: "application/json", Text: string(file)},
		},
	}, nil
}
