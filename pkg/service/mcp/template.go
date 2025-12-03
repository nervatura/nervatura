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

func loadTemplate(uri string) (result *mcp.ResourceContents, err error) {
	var u *url.URL
	if u, err = url.Parse(uri); err == nil {
		if u.Scheme != "template" {
			return nil, fmt.Errorf("wrong scheme: %q", u.Scheme)
		}
		var file []byte
		if file, err = st.Report.ReadFile(path.Join("template", u.Opaque)); err == nil {
			result = &mcp.ResourceContents{
				URI:      uri,
				MIMEType: "application/json",
				Text:     string(file),
			}
		}
	}
	return result, err
}

func templateResource(ctx context.Context, req *mcp.ReadResourceRequest) (result *mcp.ReadResourceResult, err error) {
	result = &mcp.ReadResourceResult{
		Contents: []*mcp.ResourceContents{},
	}
	var content *mcp.ResourceContents
	if content, err = loadTemplate(req.Params.URI); err == nil {
		result.Contents = append(result.Contents, content)
	}
	return result, err
}
