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

func Test_worksheetLoadData(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		data    any
		wantErr bool
	}{
		{
			name: "success",
			data: cu.IM{
				"code": "WOR1731101982N123",
				"worksheet": cu.IM{
					"distance": 100,
					"repair":   10,
					"total":    110,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, gotErr := worksheetLoadData(tt.data)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("worksheetLoadData() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("worksheetLoadData() succeeded unexpectedly")
			}
		})
	}
}

func Test_worksheetCreateHandler(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		req       *mcp.CallToolRequest
		inputData worksheetCreate
		wantErr   bool
		ds        *api.DataStore
	}{
		{
			name: "success",
			req: &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{
				Name: "nervatura_worksheet_create"},
				Extra: &mcp.RequestExtra{
					TokenInfo: &auth.TokenInfo{
						Extra: cu.IM{
							"user": md.Auth{
								Code: "123456",
							},
						},
					},
				}},
			inputData: worksheetCreate{
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
							return []cu.IM{{"id": 1, "code": "WOR1731101982N123"}}, nil
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
				Name: "nervatura_worksheet_create"},
				Extra: &mcp.RequestExtra{
					TokenInfo: &auth.TokenInfo{
						Extra: cu.IM{
							"user": md.Auth{
								Code: "123456",
							},
						},
					},
				}},
			inputData: worksheetCreate{
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
			_, _, gotErr := worksheetCreateHandler(ctx, tt.req, tt.inputData)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("worksheetCreateHandler() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("worksheetCreateHandler() succeeded unexpectedly")
			}
		})
	}
}

func Test_worksheetValidate(t *testing.T) {
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
				"code": "WOR1731101982N123",
			},
			wantErr: true,
			ds: &api.DataStore{
				Db: &td.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "WOR1731101982N123", "trans_meta": cu.IM{"closed": true}}}, nil
						},
					},
				},
			},
		},
		{
			name: "success",
			input: cu.IM{
				"code": "WOR1731101982N123",
			},
			wantErr: false,
			ds: &api.DataStore{
				Db: &td.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "WOR1731101982N123", "trans_meta": cu.IM{}}}, nil
						},
					},
				},
			},
		},
		{
			name: "missing_code",
			input: cu.IM{
				"code": "WOR1731101982N123",
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
			_, gotErr := worksheetValidate(ctx, tt.input)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("worksheetValidate() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("worksheetValidate() succeeded unexpectedly")
			}
		})
	}
}
