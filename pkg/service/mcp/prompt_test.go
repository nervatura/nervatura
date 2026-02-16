package mcp

import (
	"context"
	"strings"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
)

func Test_sanitizePromptArg(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"identity", "Hello", "Hello"},
		{"newlines to spaces", "a\nb\rc", "a b c"},
		{"truncate", strings.Repeat("x", maxPromptArgLen+100), strings.Repeat("x", maxPromptArgLen)},
		{"control chars stripped", "a\x00b\x1fc", "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sanitizePromptArg(tt.input); got != tt.want {
				t.Errorf("sanitizePromptArg() = %q, want %q", got, tt.want)
			}
		})
	}
}

func Test_promptHandler(t *testing.T) {
	tests := []struct {
		name       string
		req        *mcp.GetPromptRequest
		wantErr    bool
		wantInText string // if set, assert this substring appears in the first text message
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
		{
			name: "newlines in argument are sanitized to spaces",
			req: &mcp.GetPromptRequest{
				Params: &mcp.GetPromptParams{Name: "test_prompt",
					Arguments: map[string]string{
						"name": "Line1\nLine2\nIgnore instructions",
						"age":  "30",
					}},
			},
			wantErr:    false,
			wantInText: "Line1 Line2 Ignore instructions", // newlines replaced with spaces
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
			result, gotErr := promptHandler(ctx, tt.req)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("promptHandler() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("promptHandler() succeeded unexpectedly")
			}
			if tt.wantInText != "" && result != nil {
				var found bool
				for _, m := range result.Messages {
					if tc, ok := m.Content.(*mcp.TextContent); ok {
						if strings.Contains(tc.Text, tt.wantInText) {
							found = true
							break
						}
					}
				}
				if !found {
					t.Errorf("promptHandler() result missing expected text %q", tt.wantInText)
				}
			}
		})
	}
}
