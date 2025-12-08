package mcp

import (
	"context"
	"errors"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
)

func Test_modelDelete(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		req       *mcp.CallToolRequest
		inputData map[string]any
		wantErr   bool
		ds        *api.DataStore
		resultErr error
	}{
		{
			name: "success",
			inputData: map[string]any{
				"code": "CUS123456",
			},
			wantErr:   false,
			resultErr: nil,
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "CUS123456"}}, nil
						},
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
					},
				},
			},
		},
		{
			name: "cancel",
			inputData: map[string]any{
				"code": "CUS123456",
			},
			wantErr:   true,
			resultErr: errors.New("eliciting failed: context canceled"),
		},
		{
			name: "code not found",
			inputData: map[string]any{
				"code": "XXX123456",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			clientTransport, serverTransport := mcp.NewInMemoryTransports()

			// Create server
			server := mcp.NewServer(&mcp.Implementation{Name: "config-server", Version: "v1.0.0"}, nil)
			serverSession, _ := server.Connect(ctx, serverTransport, nil)

			ctx = context.WithValue(ctx, md.DataStoreCtxKey, tt.ds)
			client := mcp.NewClient(&mcp.Implementation{Name: "config-client", Version: "v1.0.0"}, &mcp.ClientOptions{
				ElicitationHandler: func(ctx context.Context, request *mcp.ElicitRequest) (*mcp.ElicitResult, error) {
					return &mcp.ElicitResult{
						Action: "accept", Content: map[string]any{"confirm": "YES"},
					}, tt.resultErr
				},
			})

			client.Connect(ctx, clientTransport, nil)
			_, _, gotErr := modelDelete(ctx, &mcp.CallToolRequest{
				Params:  &mcp.CallToolParamsRaw{Name: "nervatura_customer_delete"},
				Session: serverSession}, tt.inputData)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("modelDelete() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("modelDelete() succeeded unexpectedly")
			}
		})
	}
}

func Test_extendDelete(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		req       *mcp.CallToolRequest
		inputData map[string]any
		wantErr   bool
		ds        *api.DataStore
		resultErr error
		toolName  string
	}{
		{
			name: "success",
			inputData: map[string]any{
				"code": "CUS123456",
			},
			wantErr:   false,
			resultErr: nil,
			toolName:  "nervatura_contact_delete",
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "CUS123456", "contacts": []cu.IM{{"id": 1, "surname": "Doe"}}}}, nil
						},
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
					},
				},
			},
		},
		{
			name: "cancel",
			inputData: map[string]any{
				"code": "CUS123456",
			},
			wantErr:   true,
			resultErr: errors.New("eliciting failed: context canceled"),
			toolName:  "nervatura_contact_delete",
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "CUS123456", "contacts": []cu.IM{{"id": 1, "surname": "Doe"}}}}, nil
						},
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
					},
				},
			},
		},
		{
			name: "index out of range",
			inputData: map[string]any{
				"code":  "CUS123456",
				"index": 1,
			},
			wantErr:   true,
			resultErr: nil,
			toolName:  "nervatura_contact_delete",
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "CUS123456", "contacts": []cu.IM{{"id": 1, "surname": "Doe"}}}}, nil
						},
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
					},
				},
			},
		},
		{
			name: "not found",
			inputData: map[string]any{
				"code":  "CUS123456",
				"index": 1,
			},
			wantErr:   true,
			resultErr: nil,
			toolName:  "nervatura_contact_delete",
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{}, nil
						},
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
					},
				},
			},
		},
		{
			name: "invalid code",
			inputData: map[string]any{
				"code":  "XXX123456",
				"index": 1,
			},
			wantErr:   true,
			resultErr: nil,
			toolName:  "nervatura_contact_delete",
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{}, nil
						},
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
					},
				},
			},
		},
		{
			name: "invalid tool",
			inputData: map[string]any{
				"code":  "XXX123456",
				"index": 1,
			},
			wantErr:   true,
			resultErr: nil,
			toolName:  "nervatura_customer_delete",
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{}, nil
						},
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			clientTransport, serverTransport := mcp.NewInMemoryTransports()

			// Create server
			server := mcp.NewServer(&mcp.Implementation{Name: "config-server", Version: "v1.0.0"}, nil)
			serverSession, _ := server.Connect(ctx, serverTransport, nil)

			ctx = context.WithValue(ctx, md.DataStoreCtxKey, tt.ds)
			client := mcp.NewClient(&mcp.Implementation{Name: "config-client", Version: "v1.0.0"}, &mcp.ClientOptions{
				ElicitationHandler: func(ctx context.Context, request *mcp.ElicitRequest) (*mcp.ElicitResult, error) {
					return &mcp.ElicitResult{
						Action: "accept", Content: map[string]any{"confirm": "YES"},
					}, tt.resultErr
				},
			})

			client.Connect(ctx, clientTransport, nil)
			_, _, gotErr := extendDelete(ctx, &mcp.CallToolRequest{
				Params:  &mcp.CallToolParamsRaw{Name: tt.toolName},
				Session: serverSession}, tt.inputData)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("extendDelete() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("extendDelete() succeeded unexpectedly")
			}
		})
	}
}
