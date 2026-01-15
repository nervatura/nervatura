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

func Test_rentLoadData(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		data    any
		wantErr bool
	}{
		{
			name: "success",
			data: cu.IM{
				"code": "REN1731101982N123",
				"rent": cu.IM{
					"distance": 100,
					"repair":   10,
					"total":    110,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, gotErr := rentLoadData(tt.data)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("rentLoadData() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("rentLoadData() succeeded unexpectedly")
			}
		})
	}
}

func Test_rentCreateHandler(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		req       *mcp.CallToolRequest
		inputData rentCreate
		wantErr   bool
		ds        *api.DataStore
	}{
		{
			name: "success",
			req: &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{
				Name: "nervatura_rent_create"},
				Extra: &mcp.RequestExtra{
					TokenInfo: &auth.TokenInfo{
						Extra: cu.IM{
							"user": md.Auth{
								Code: "123456",
							},
						},
					},
				}},
			inputData: rentCreate{
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
							return []cu.IM{{"id": 1, "code": "REN1731101982N123"}}, nil
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
				Name: "nervatura_rent_create"},
				Extra: &mcp.RequestExtra{
					TokenInfo: &auth.TokenInfo{
						Extra: cu.IM{
							"user": md.Auth{
								Code: "123456",
							},
						},
					},
				}},
			inputData: rentCreate{
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
			_, _, gotErr := rentCreateHandler(ctx, tt.req, tt.inputData)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("rentCreateHandler() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("rentCreateHandler() succeeded unexpectedly")
			}
		})
	}
}

func Test_rentValidate(t *testing.T) {
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
				"code": "REN1731101982N123",
			},
			wantErr: true,
			ds: &api.DataStore{
				Db: &td.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "REN1731101982N123", "trans_meta": cu.IM{"closed": true}}}, nil
						},
					},
				},
			},
		},
		{
			name: "success",
			input: cu.IM{
				"code": "REN1731101982N123",
			},
			wantErr: false,
			ds: &api.DataStore{
				Db: &td.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "REN1731101982N123", "trans_meta": cu.IM{}}}, nil
						},
					},
				},
			},
		},
		{
			name: "missing_code",
			input: cu.IM{
				"code": "REN1731101982N123",
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
			_, gotErr := rentValidate(ctx, tt.input)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("rentValidate() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("rentValidate() succeeded unexpectedly")
			}
		})
	}
}
