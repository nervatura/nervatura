package mcp

import (
	"context"
	"errors"
	"net/url"
	"path"
	"slices"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	st "github.com/nervatura/nervatura/v6/pkg/static"
)

type ResourceData struct {
	mcp.Resource
	Scopes []string
}

func getResource(config cu.IM, uri string) (resource *ResourceData, err error) {
	resources, ok := config["resources"].(map[string]ResourceData)
	if !ok {
		return nil, errors.New("resources not found")
	}
	for _, resource := range resources {
		if resource.URI == uri {
			return &resource, nil
		}
	}
	return nil, errors.New("resources not found")
}

func getResourceContent(ctx context.Context, uri string) (result *mcp.ResourceContents, err error) {
	config := ctx.Value(md.ConfigCtxKey).(cu.IM)
	var resource *ResourceData
	if resource, err = getResource(config, uri); err != nil {
		return nil, err
	}
	var content *mcp.ResourceContents = &mcp.ResourceContents{
		URI:      resource.URI,
		MIMEType: resource.MIMEType,
	}

	var u *url.URL
	if u, err = url.Parse(resource.URI); err == nil {
		if slices.Contains([]string{"template", "mcp"}, u.Scheme) {
			var file []byte
			if file, err = st.Static.ReadFile(path.Join(u.Scheme, u.Opaque)); err == nil {
				content.Text = string(file)
			}
		}
	}
	return content, err
}

func resourceHandler(ctx context.Context, req *mcp.ReadResourceRequest) (result *mcp.ReadResourceResult, err error) {
	var content *mcp.ResourceContents
	content, err = getResourceContent(ctx, req.Params.URI)
	result = &mcp.ReadResourceResult{
		Contents: []*mcp.ResourceContents{content},
	}
	return result, err
}
