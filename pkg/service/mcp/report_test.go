package mcp

import (
	"context"
	"log/slog"
	"path"
	"testing"
	"time"

	"github.com/modelcontextprotocol/go-sdk/auth"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	st "github.com/nervatura/nervatura/v6/pkg/static"
)

func Test_reportQueryHandler(t *testing.T) {
	pdf_json, _ := st.Report.ReadFile(path.Join("template", "ntr_customer_en.json"))
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		req        *mcp.CallToolRequest
		parameters cu.IM
		wantErr    bool
		ds         *api.DataStore
	}{
		{
			name: "success",
			req: &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_report_query"},
				Extra: &mcp.RequestExtra{
					TokenInfo: &auth.TokenInfo{Expiration: time.Now().Add(time.Duration(1) * time.Hour)},
				}},
			parameters: cu.IM{
				"code": "CUS1731101982N123", "report_key": "",
			},
			wantErr: false,
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "report_key": "ntr_customer_en", "data": cu.IM{"file_type": "FILE_PDF", "template": string(pdf_json)}}}, nil
						},
						"QuerySQL": func(sqlString string) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "report_key": "ntr_customer_en", "data": cu.IM{"file_type": "FILE_PDF"}}}, nil
						},
					},
				},
				Config:          cu.IM{},
				AppLog:          slog.Default(),
				ConvertFromByte: cu.ConvertFromByte,
				ConvertToByte: func(v any) ([]byte, error) {
					return nil, nil
				},
				ConvertToType: func(data interface{}, result any) (err error) {
					return nil
				},
			},
		},
		{
			name: "url_result",
			req: &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_report_query"},
				Extra: &mcp.RequestExtra{
					TokenInfo: &auth.TokenInfo{Expiration: time.Now().Add(time.Duration(1) * time.Hour)},
				},
				Session: &mcp.ServerSession{}},
			parameters: cu.IM{
				"code": "CUS1731101982N123", "report_key": "", "url_result": true, "output": "base64",
			},
			wantErr: false,
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "report_key": "ntr_customer_en", "data": cu.IM{"file_type": "FILE_PDF", "template": string(pdf_json)}}}, nil
						},
						"QuerySQL": func(sqlString string) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "report_key": "ntr_customer_en", "data": cu.IM{"file_type": "FILE_PDF"}}}, nil
						},
					},
				},
				Config:          cu.IM{},
				AppLog:          slog.Default(),
				ConvertFromByte: cu.ConvertFromByte,
				ConvertToByte: func(v any) ([]byte, error) {
					return nil, nil
				},
				ConvertToType: func(data interface{}, result any) (err error) {
					return nil
				},
			},
		},
		{
			name: "invalid report",
			req: &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_report_query"},
				Extra: &mcp.RequestExtra{
					TokenInfo: &auth.TokenInfo{Expiration: time.Now().Add(time.Duration(1) * time.Hour)},
				}},
			parameters: cu.IM{
				"code": "XXX1731101982N123", "report_key": "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			ctx = context.WithValue(ctx, md.SessionServiceCtxKey, &api.SessionService{
				Config: api.SessionConfig{
					Method: md.SessionMethodMemory,
				},
				Conn: &md.TestDriver{Config: cu.IM{}},
			})
			ctx = context.WithValue(ctx, md.ConfigCtxKey, cu.IM{})
			_, _, gotErr := reportQueryHandler(ctx, tt.req, tt.parameters)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("reportQueryHandler() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("reportQueryHandler() succeeded unexpectedly")
			}
		})
	}
}

func Test_getDefaultReportKey(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		code    string
		wantErr bool
		ds      *api.DataStore
	}{
		{
			name:    "success",
			code:    "INV1731101982N123",
			wantErr: false,
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "report_key": "ntr_invoice_en", "trans_type": "TRANS_INVOICE", "direction": "DIRECTION_OUT"}}, nil
						},
					},
				},
			},
		},
		{
			name:    "missing default report key",
			code:    "INV1731101982N123",
			wantErr: true,
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							if queries[0].From == "config_report" {
								return []cu.IM{}, nil
							}
							return []cu.IM{{"id": 1, "report_key": "ntr_invoice_en", "trans_type": "TRANS_INVOICE", "direction": "DIRECTION_OUT"}}, nil
						},
					},
				},
			},
		},
		{
			name:    "missing code",
			code:    "INV1731101982N123",
			wantErr: true,
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{},
				},
			},
		},
		{
			name:    "invalid model",
			code:    "XXX1731101982N123",
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
			_, gotErr := getDefaultReportKey(ctx, tt.code)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("getDefaultReportKey() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("getDefaultReportKey() succeeded unexpectedly")
			}
		})
	}
}
