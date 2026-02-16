package mcp

import (
	"context"
	"errors"
	"strings"
	"unicode"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
)

// maxPromptArgLen limits argument length to mitigate prompt injection and DoS.
const maxPromptArgLen = 4096

type PromptData struct {
	Name              string                `json:"name,omitempty" jsonschema:"Prompt name."`
	Title             string                `json:"title,omitempty" jsonschema:"Prompt title."`
	Description       string                `json:"description,omitempty" jsonschema:"Prompt description."`
	Arguments         []*mcp.PromptArgument `json:"arguments,omitempty" jsonschema:"Prompt arguments."`
	Meta              mcp.Meta              `json:"_meta,omitempty" jsonschema:"Prompt meta."`
	PromptDescription string                `json:"prompt_description,omitempty" jsonschema:"Result description."`
	PromptMessages    []*mcp.PromptMessage  `json:"prompt_messages,omitempty" jsonschema:"Prompt messages."`
}

// sanitizePromptArg restricts argument values to mitigate prompt injection.
// Applies length limit, replaces newlines with spaces, and strips control characters.
func sanitizePromptArg(value string) string {
	if len(value) > maxPromptArgLen {
		value = value[:maxPromptArgLen]
	}
	var b strings.Builder
	b.Grow(len(value))
	for _, r := range value {
		switch {
		case r == '\n' || r == '\r':
			b.WriteRune(' ')
		case unicode.IsControl(r):
			// skip other control chars
			continue
		default:
			b.WriteRune(r)
		}
	}
	return b.String()
}

func promptHandler(ctx context.Context, req *mcp.GetPromptRequest) (result *mcp.GetPromptResult, err error) {
	config := ctx.Value(md.ConfigCtxKey).(cu.IM)
	texReplace := func(text string) string {
		for key, value := range req.Params.Arguments {
			text = strings.ReplaceAll(text, "{{"+key+"}}", sanitizePromptArg(value))
		}
		return text
	}
	checkRequired := func(arguments []*mcp.PromptArgument) error {
		for _, argument := range arguments {
			if argument.Required {
				if value, ok := req.Params.Arguments[argument.Name]; !ok || value == "" {
					return errors.New("required argument is missing: " + argument.Name)
				}
			}
		}
		return nil
	}
	if prompts, ok := config["prompts"].(map[string]PromptData); ok {
		if prompt, ok := prompts[req.Params.Name]; ok {
			result = &mcp.GetPromptResult{
				Description: prompt.PromptDescription,
				Messages:    []*mcp.PromptMessage{},
			}
			if err = checkRequired(prompt.Arguments); err != nil {
				return result, err
			}
			for _, message := range prompt.PromptMessages {
				switch message.Content.(type) {
				case *mcp.TextContent:
					result.Messages = append(result.Messages, &mcp.PromptMessage{
						Role:    message.Role,
						Content: &mcp.TextContent{Text: texReplace(message.Content.(*mcp.TextContent).Text)},
					})
				case *mcp.EmbeddedResource:
					var content *mcp.ResourceContents
					var uri string = message.Content.(*mcp.EmbeddedResource).Resource.URI
					if content, err = getResourceContent(ctx, uri); err == nil {
						result.Messages = append(result.Messages, &mcp.PromptMessage{
							Role:    message.Role,
							Content: &mcp.EmbeddedResource{Resource: content},
						})
					}
				default:
					result.Messages = append(result.Messages, message)
				}

			}
		}
	}
	return result, err
}
