package mcp

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
)

func Test_promptHandler(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		req     *mcp.GetPromptRequest
		wantErr bool
	}{
		{
			name: "success",
			req: &mcp.GetPromptRequest{
				Params: &mcp.GetPromptParams{Name: "test_prompt",
					Arguments: map[string]string{
						"name": "John",
						"age":  "30",
					}},
			},
			wantErr: false,
		},
		{
			name: "missing required argument",
			req: &mcp.GetPromptRequest{
				Params: &mcp.GetPromptParams{Name: "test_prompt",
					Arguments: map[string]string{
						"age": "30",
					}},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.ConfigCtxKey, cu.IM{
				"prompts": map[string]PromptData{
					"test_prompt": {
						Name:        "test_prompt",
						Title:       "Test Prompt",
						Description: "This is a test prompt.",
						Arguments: []*mcp.PromptArgument{
							{
								Name:        "name",
								Title:       "Name",
								Description: "Please enter the name of the person to greet.",
								Required:    true,
							},
						},
						PromptMessages: []*mcp.PromptMessage{
							{
								Content: &mcp.TextContent{Text: "Hello, {{name}}! You are {{age}} years old."},
								Role:    "user",
							},
							{
								Content: &mcp.EmbeddedResource{Resource: &mcp.ResourceContents{
									URI:      "template:ntr_customer_en.json",
									MIMEType: "application/json",
								}},
								Role: "user",
							},
							{
								Content: &mcp.ResourceLink{
									URI:         "http://localhost:5000/public/images/logo.svg",
									MIMEType:    "image/svg+xml",
									Name:        "test_link",
									Title:       "Test Link",
									Description: "This is a test link.",
								},
								Role: "user",
							},
						},
					},
				},
				"resources": map[string]mcp.Resource{
					"ntr_customer_en": {
						URI:      "template:ntr_customer_en.json",
						MIMEType: "application/json",
					},
				},
			})
			_, gotErr := promptHandler(ctx, tt.req)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("promptHandler() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("promptHandler() succeeded unexpectedly")
			}
		})
	}
}
