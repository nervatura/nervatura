package mcp

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
)

func Test_resourceHandler(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		req     *mcp.ReadResourceRequest
		wantErr bool
		config  cu.IM
	}{
		{
			name:    "success",
			req:     &mcp.ReadResourceRequest{Params: &mcp.ReadResourceParams{URI: "template:ntr_customer_en.json"}},
			wantErr: false,
			config: cu.IM{
				"resources": map[string]ResourceData{
					"ntr_customer_en": {
						Resource: mcp.Resource{URI: "template:ntr_customer_en.json"},
					},
				},
			},
		},
		{
			name:    "missing resources",
			req:     &mcp.ReadResourceRequest{Params: &mcp.ReadResourceParams{URI: "template:ntr_customer_en.json"}},
			wantErr: true,
			config:  cu.IM{},
		},
		{
			name:    "missing resource",
			req:     &mcp.ReadResourceRequest{Params: &mcp.ReadResourceParams{URI: "template:ntr_customer_en.json"}},
			wantErr: true,
			config: cu.IM{
				"resources": map[string]ResourceData{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.ConfigCtxKey, tt.config)
			_, gotErr := resourceHandler(ctx, tt.req)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("resourceHandler() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("resourceHandler() succeeded unexpectedly")
			}

		})
	}
}
