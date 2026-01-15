package mcp

import (
	"bytes"
	"context"
	"errors"
	"log/slog"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	td "github.com/nervatura/nervatura/v6/test/driver"
)

func Test_currencyCreateHandler(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		req       *mcp.CallToolRequest
		inputData currencyUpdate
		wantErr   bool
		ds        *api.DataStore
	}{
		{
			name: "success",
			req:  &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_currency_create"}},
			inputData: currencyUpdate{
				Code: "EUR",
				CurrencyMeta: md.CurrencyMeta{
					Description: "Euro",
				},
				CurrencyMap: cu.IM{
					"symbol": "€",
				},
			},
			wantErr: false,
			ds: &api.DataStore{
				Db: &td.TestDriver{
					Config: cu.IM{
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "EUR"}}, nil
						},
						"GetDataField": func(data any, JSONName string) (fieldName string, fieldValue interface{}) {
							return JSONName, "value"
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
		},
		{
			name: "missing_code",
			req:  &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_customer_create"}},
			inputData: currencyUpdate{
				CurrencyMeta: md.CurrencyMeta{
					Description: "Euro",
				},
				CurrencyMap: cu.IM{
					"symbol": "€",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			_, _, gotErr := currencyCreateHandler(ctx, tt.req, tt.inputData)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("currencyCreateHandler() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("currencyCreateHandler() succeeded unexpectedly")
			}
		})
	}
}

func Test_currencyDelete(t *testing.T) {
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
				"code": "EUR",
			},
			wantErr:   false,
			resultErr: nil,
			ds: &api.DataStore{
				Db: &td.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "EUR"}}, nil
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
				"code": "EUR",
			},
			wantErr:   true,
			resultErr: errors.New("eliciting failed: context canceled"),
		},
		{
			name: "code not found",
			inputData: map[string]any{
				"code": "",
			},
			wantErr:   true,
			resultErr: errors.New("code is required and must be 3 characters long"),
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
			_, _, gotErr := currencyDelete(ctx, &mcp.CallToolRequest{
				Params:  &mcp.CallToolParamsRaw{Name: "nervatura_currency_delete"},
				Session: serverSession}, tt.inputData)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("currencyDelete() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("currencyDelete() succeeded unexpectedly")
			}
		})
	}
}

func Test_currencyLoadData(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		data    any
		wantErr bool
	}{
		{
			name: "success",
			data: cu.IM{
				"code":        "EUR",
				"description": "Euro",
				"symbol":      "€",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, gotErr := currencyLoadData(tt.data)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("currencyLoadData() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("currencyLoadData() succeeded unexpectedly")
			}
		})
	}
}
