package mcp

/*
func init() {
	toolDataMap["nervatura_gemini_query"] = ToolData{
		Tool: mcp.Tool{
			Name:        "nervatura_gemini_query",
			Title:       "Gemini question",
			Description: "Send a question to Gemini and get the response.",
		},
		ModelSchema: GeminiSchema(),
		ConnectHandler: func(server *mcp.Server, tool *mcp.Tool) {
			mcp.AddTool(server, tool, geminiQueryHandler)
		},
		Scopes: []string{"public"},
	}
}

func GeminiSchema() (ms *ModelSchema) {
	return &ModelSchema{
		Name: "gemini",
		QueryInputSchema: func(scope string) (schema *jsonschema.Schema) {
			return &jsonschema.Schema{
				Type: "object",
				Properties: map[string]*jsonschema.Schema{
					"question": {Type: "string", Description: "The question to send to Gemini."},
				},
				Required: []string{"question"},
			}
		},
		QueryOutputSchema: func(scope string) (schema *jsonschema.Schema) {
			return &jsonschema.Schema{
				Type: "object",
				Properties: map[string]*jsonschema.Schema{
					"response": {Type: "string", Description: "The response from Gemini."},
				},
			}
		},
		Examples: map[string][]any{
			"question": {`What is the capital of France?`},
		},
	}
}

func geminiQueryHandler(ctx context.Context, req *mcp.CallToolRequest, inputData cu.IM) (result *mcp.CallToolResult, response any, err error) {
	config := ctx.Value(md.ConfigCtxKey).(cu.IM)
	model := cu.ToString(config["NT_GEMINI_MODEL"], "gemini-2.5-flash")
	apiKey := cu.ToString(config["NT_GEMINI_API_KEY"], "")

	if apiKey == "" {
		return result, response, errors.New("NT_GEMINI_API_KEY is not set")
	}
	question := cu.ToString(inputData["question"], "")
	if question == "" {
		return result, response, errors.New("question is required")
	}

	var client *genai.Client
	if client, err = genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: apiKey, Backend: genai.BackendGeminiAPI,
	}); err != nil {
		return result, response, err
	}

	var res *genai.GenerateContentResponse
	if res, err = client.Models.GenerateContent(
		ctx, model, genai.Text(question),
		&genai.GenerateContentConfig{},
	); err == nil {
		response = cu.IM{
			"response": res.Text(),
		}
	}

	return result, response, err
}
*/
