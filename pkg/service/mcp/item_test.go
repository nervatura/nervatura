package mcp

import (
	"bytes"
	"context"
	"errors"
	"log/slog"
	"strings"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/auth"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
)

func Test_itemCreateHandler(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		req       *mcp.CallToolRequest
		inputData itemCreate
		wantErr   bool
		ds        *api.DataStore
	}{
		{
			name: "success",
			req: &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{
				Name: "nervatura_invoice_create"},
				Extra: &mcp.RequestExtra{
					TokenInfo: &auth.TokenInfo{
						Extra: cu.IM{
							"user": md.Auth{
								Code: "123456",
							},
						},
					},
				}},
			inputData: itemCreate{
				TransCode:   "INV1731101982N123",
				ProductCode: "PRD1731101982N123",
				TaxCode:     "VAT20",
				itemMeta: itemMeta{
					InputType:   "DEFAULT_PRICE",
					Qty:         1,
					Amount:      100,
					Discount:    0,
					Description: "Test Item",
					Deposit:     false,
					ActionPrice: false,
				},
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
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
		},
		{
			name: "calc_error",
			req: &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{
				Name: "nervatura_invoice_create"},
				Extra: &mcp.RequestExtra{
					TokenInfo: &auth.TokenInfo{
						Extra: cu.IM{
							"user": md.Auth{
								Code: "123456",
							},
						},
					},
				}},
			inputData: itemCreate{
				TransCode:   "INV1731101982N123",
				ProductCode: "PRD1731101982N123",
				TaxCode:     "VAT20",
				itemMeta: itemMeta{
					InputType:   "DEFAULT_PRICE",
					Qty:         1,
					Amount:      100,
					Discount:    0,
					Description: "Test Item",
					Deposit:     false,
					ActionPrice: false,
				},
			},
			wantErr: true,
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							if strings.Contains(queries[0].From, "currency_view") {
								return nil, errors.New("error")
							}
							return []cu.IM{{"id": 1, "code": "CUS123456"}}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
		},
		{
			name: "price_error",
			req: &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{
				Name: "nervatura_invoice_create"},
				Extra: &mcp.RequestExtra{
					TokenInfo: &auth.TokenInfo{
						Extra: cu.IM{
							"user": md.Auth{
								Code: "123456",
							},
						},
					},
				}},
			inputData: itemCreate{
				TransCode:   "INV1731101982N123",
				ProductCode: "PRD1731101982N123",
				TaxCode:     "VAT20",
				itemMeta: itemMeta{
					InputType:   "DEFAULT_PRICE",
					Qty:         1,
					Amount:      100,
					Discount:    0,
					Description: "Test Item",
					Deposit:     false,
					ActionPrice: false,
				},
			},
			wantErr: true,
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							if strings.Contains(queries[0].From, "price_view") {
								return nil, errors.New("error")
							}
							return []cu.IM{{"id": 1, "code": "CUS123456"}}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
		},
		{
			name: "trans_error",
			req: &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{
				Name: "nervatura_invoice_create"},
				Extra: &mcp.RequestExtra{
					TokenInfo: &auth.TokenInfo{
						Extra: cu.IM{
							"user": md.Auth{
								Code: "123456",
							},
						},
					},
				}},
			inputData: itemCreate{
				TransCode:   "INV1731101982N123",
				ProductCode: "PRD1731101982N123",
				TaxCode:     "VAT20",
				itemMeta: itemMeta{
					InputType:   "DEFAULT_PRICE",
					Qty:         1,
					Amount:      100,
					Discount:    0,
					Description: "Test Item",
					Deposit:     false,
					ActionPrice: false,
				},
			},
			wantErr: true,
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							if strings.Contains(queries[0].From, "trans") {
								return nil, errors.New("error")
							}
							return []cu.IM{{"id": 1, "code": "CUS123456"}}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
		},
		{
			name: "product_error",
			req: &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{
				Name: "nervatura_invoice_create"},
				Extra: &mcp.RequestExtra{
					TokenInfo: &auth.TokenInfo{
						Extra: cu.IM{
							"user": md.Auth{
								Code: "123456",
							},
						},
					},
				}},
			inputData: itemCreate{
				TransCode:   "INV1731101982N123",
				ProductCode: "PRD1731101982N123",
				TaxCode:     "VAT20",
				itemMeta: itemMeta{
					InputType:   "DEFAULT_PRICE",
					Qty:         1,
					Amount:      100,
					Discount:    0,
					Description: "Test Item",
					Deposit:     false,
					ActionPrice: false,
				},
			},
			wantErr: true,
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							if strings.Contains(queries[0].From, "product_view") {
								return nil, errors.New("error")
							}
							return []cu.IM{{"id": 1, "code": "CUS123456"}}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
		},
		{
			name: "required_error",
			req: &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{
				Name: "nervatura_invoice_create"},
				Extra: &mcp.RequestExtra{
					TokenInfo: &auth.TokenInfo{
						Extra: cu.IM{
							"user": md.Auth{
								Code: "123456",
							},
						},
					},
				}},
			inputData: itemCreate{
				ProductCode: "PRD1731101982N123",
				TaxCode:     "VAT20",
				itemMeta: itemMeta{
					InputType:   "DEFAULT_PRICE",
					Qty:         1,
					Amount:      100,
					Discount:    0,
					Description: "Test Item",
					Deposit:     false,
					ActionPrice: false,
				},
			},
			wantErr: true,
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							if strings.Contains(queries[0].From, "product_view") {
								return nil, errors.New("error")
							}
							return []cu.IM{{"id": 1, "code": "CUS123456"}}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			_, _, gotErr := itemCreateHandler(ctx, tt.req, tt.inputData)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("itemCreateHandler() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("itemCreateHandler() succeeded unexpectedly")
			}
		})
	}
}

func Test_itemLoadData(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		data    any
		wantErr bool
	}{
		{
			name: "success",
			data: cu.IM{
				"code":         "ITM1731101982N123",
				"trans_code":   "INV1731101982N123",
				"product_code": "PRD1731101982N123",
				"tax_code":     "VAT20",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, gotErr := itemLoadData(tt.data)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("itemLoadData() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("itemLoadData() succeeded unexpectedly")
			}
		})
	}
}

func Test_calcItemPrice(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		inputData cu.IM
		wantErr   bool
		ds        *api.DataStore
	}{
		{
			name: "FX_PRICE",
			inputData: cu.IM{
				"code":         "ITM1731101982N123",
				"trans_code":   "INV1731101982N123",
				"product_code": "PRD1731101982N123",
				"tax_code":     "VAT20",
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "CUS123456"}}, nil
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "NET_AMOUNT",
			inputData: cu.IM{
				"code":       "ITM1731101982N123",
				"input_type": "NET_AMOUNT",
				"amount":     100,
				"qty":        1,
				"discount":   0,
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "CUS123456"}}, nil
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "AMOUNT",
			inputData: cu.IM{
				"trans_code": "INV1731101982N123",
				"tax_code":   "VAT20",
				"input_type": "AMOUNT",
				"amount":     120,
				"qty":        1,
				"discount":   10,
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "CUS123456"}}, nil
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "tax_error",
			inputData: cu.IM{
				"trans_code": "INV1731101982N123",
				"tax_code":   "VAT20",
				"input_type": "AMOUNT",
				"amount":     120,
				"qty":        1,
				"discount":   10,
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							if strings.Contains(queries[0].From, "tax_view") {
								return nil, errors.New("error")
							}
							return []cu.IM{{"id": 1, "code": "CUS123456"}}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			wantErr: true,
		},
		{
			name: "currency_error",
			inputData: cu.IM{
				"trans_code": "INV1731101982N123",
				"tax_code":   "VAT20",
				"input_type": "AMOUNT",
				"amount":     120,
				"qty":        1,
				"discount":   10,
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							if strings.Contains(queries[0].From, "currency_view") {
								return nil, errors.New("error")
							}
							return []cu.IM{{"id": 1, "code": "CUS123456"}}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			wantErr: true,
		},
		{
			name: "item_error",
			inputData: cu.IM{
				"code":       "ITM1731101982N123",
				"trans_code": "INV1731101982N123",
				"tax_code":   "VAT20",
				"input_type": "AMOUNT",
				"amount":     120,
				"qty":        1,
				"discount":   10,
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return nil, errors.New("error")
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			wantErr: true,
		},
		{
			name: "required_error",
			inputData: cu.IM{
				"tax_code":   "VAT20",
				"input_type": "AMOUNT",
				"amount":     120,
				"qty":        1,
				"discount":   10,
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return nil, errors.New("error")
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			_, gotErr := calcItemPrice(ctx, tt.inputData)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("calcItemPrice() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("calcItemPrice() succeeded unexpectedly")
			}
		})
	}
}

func Test_itemValidate(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		input   cu.IM
		wantErr bool
		ds      *api.DataStore
	}{
		{
			name: "success",
			input: cu.IM{
				"code":         "ITM1731101982N123",
				"trans_code":   "INV1731101982N123",
				"product_code": "PRD1731101982N123",
				"tax_code":     "VAT20",
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "CUS123456"}}, nil
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "simple",
			input: cu.IM{
				"code":  "ITM1731101982N123",
				"notes": "Test Item",
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "CUS123456"}}, nil
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "calc_error",
			input: cu.IM{
				"code":         "ITM1731101982N123",
				"trans_code":   "INV1731101982N123",
				"product_code": "PRD1731101982N123",
				"tax_code":     "VAT20",
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							if strings.Contains(queries[0].From, "currency_view") {
								return nil, errors.New("error")
							}
							return []cu.IM{{"id": 1, "code": "CUS123456"}}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			wantErr: true,
		},
		{
			name: "deleted_error",
			input: cu.IM{
				"code":  "ITM1731101982N123",
				"notes": "Test Item",
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "CUS123456", "status": md.TransStatusDeleted.String()}}, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			wantErr: true,
		},
		{
			name: "trans_error",
			input: cu.IM{
				"code":  "ITM1731101982N123",
				"notes": "Test Item",
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return nil, errors.New("error")
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			_, gotErr := itemValidate(ctx, tt.input)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("itemValidate() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("itemValidate() succeeded unexpectedly")
			}
		})
	}
}
