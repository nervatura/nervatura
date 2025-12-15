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
)

func Test_offerLoadData(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		data    any
		wantErr bool
	}{
		{
			name: "success",
			data: cu.IM{
				"trans_type": md.TransTypeOffer.String(),
				"direction":  md.DirectionOut.String(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, gotErr := offerLoadData(tt.data)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("offerLoadData() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("offerLoadData() succeeded unexpectedly")
			}
		})
	}
}

func Test_offerCreateHandler(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		req       *mcp.CallToolRequest
		inputData offerCreate
		wantErr   bool
		ds        *api.DataStore
	}{
		{
			name: "success",
			req: &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{
				Name: "nervatura_offer_create"},
				Extra: &mcp.RequestExtra{
					TokenInfo: &auth.TokenInfo{
						Extra: cu.IM{
							"user": md.Auth{
								Code: "123456",
							},
						},
					},
				}},
			inputData: offerCreate{
				TransDate:    "2025-01-01",
				Direction:    "DIRECTION_OUT",
				CustomerCode: "CUS123456",
				CurrencyCode: "USD",
			},
			wantErr: false,
			ds: &api.DataStore{
				Db: &md.TestDriver{
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
				Name: "nervatura_offer_create"},
				Extra: &mcp.RequestExtra{
					TokenInfo: &auth.TokenInfo{
						Extra: cu.IM{
							"user": md.Auth{
								Code: "123456",
							},
						},
					},
				}},
			inputData: offerCreate{
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
			_, _, gotErr := offerCreateHandler(ctx, tt.req, tt.inputData)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("offerCreateHandler() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("offerCreateHandler() succeeded unexpectedly")
			}
		})
	}
}

func Test_offerValidate(t *testing.T) {
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
				"code": "OFF1731101982N123",
			},
			wantErr: true,
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "OFF1731101982N123", "trans_meta": cu.IM{"closed": true}}}, nil
						},
					},
				},
			},
		},
		{
			name: "success",
			input: cu.IM{
				"code": "OFF1731101982N123",
			},
			wantErr: false,
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "OFF1731101982N123", "trans_meta": cu.IM{}}}, nil
						},
					},
				},
			},
		},
		{
			name: "missing_code",
			input: cu.IM{
				"code": "OFF1731101982N123",
			},
			wantErr: true,
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			_, gotErr := offerValidate(ctx, tt.input)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("offerValidate() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("offerValidate() succeeded unexpectedly")
			}
		})
	}
}
