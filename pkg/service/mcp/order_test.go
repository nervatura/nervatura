package mcp

import (
	"bytes"
	"context"
	"log/slog"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/auth"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	td "github.com/nervatura/nervatura/v6/test/driver"
)

func Test_orderLoadData(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		data    any
		wantErr bool
	}{
		{
			name: "success",
			data: cu.IM{
				"trans_type": md.TransTypeOrder.String(),
				"direction":  md.DirectionOut.String(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, gotErr := orderLoadData(tt.data)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("orderLoadData() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("orderLoadData() succeeded unexpectedly")
			}
		})
	}
}

func Test_orderCreateHandler(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		req       *mcp.CallToolRequest
		inputData orderCreate
		wantErr   bool
		ds        *api.DataStore
	}{
		{
			name: "success",
			req: &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{
				Name: "nervatura_order_create"},
				Extra: &mcp.RequestExtra{
					TokenInfo: &auth.TokenInfo{
						Extra: cu.IM{
							"user": md.Auth{
								Code: "123456",
							},
						},
					},
				}},
			inputData: orderCreate{
				TransDate:    "2025-01-01",
				Direction:    "DIRECTION_OUT",
				CustomerCode: "CUS123456",
				CurrencyCode: "USD",
			},
			wantErr: false,
			ds: &api.DataStore{
				Db: &td.TestDriver{
					Config: cu.IM{
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "CUS123456"}}, nil
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
			name: "missing_trans_date",
			req: &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{
				Name: "nervatura_order_create"},
				Extra: &mcp.RequestExtra{
					TokenInfo: &auth.TokenInfo{
						Extra: cu.IM{
							"user": md.Auth{
								Code: "123456",
							},
						},
					},
				}},
			inputData: orderCreate{
				Direction:    "DIRECTION_OUT",
				CustomerCode: "CUS123456",
				CurrencyCode: "USD",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			_, _, gotErr := orderCreateHandler(ctx, tt.req, tt.inputData)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("orderCreateHandler() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("orderCreateHandler() succeeded unexpectedly")
			}
		})
	}
}

func Test_orderValidate(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		input   cu.IM
		wantErr bool
		ds      *api.DataStore
	}{
		{
			name: "deleted",
			input: cu.IM{
				"code": "ORD1731101982N123",
			},
			wantErr: true,
			ds: &api.DataStore{
				Db: &td.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "ORD1731101982N123", "trans_meta": cu.IM{"closed": true}}}, nil
						},
					},
				},
			},
		},
		{
			name: "success",
			input: cu.IM{
				"code": "ORD1731101982N123",
			},
			wantErr: false,
			ds: &api.DataStore{
				Db: &td.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "ORD1731101982N123", "trans_meta": cu.IM{}}}, nil
						},
					},
				},
			},
		},
		{
			name: "missing_code",
			input: cu.IM{
				"code": "ORD1731101982N123",
			},
			wantErr: true,
			ds: &api.DataStore{
				Db: &td.TestDriver{
					Config: cu.IM{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			_, gotErr := orderValidate(ctx, tt.input)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("orderValidate() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("orderValidate() succeeded unexpectedly")
			}
		})
	}
}
