package mcp

import (
	"context"
	"log/slog"
	"path"
	"testing"

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
			req:  &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_report_query"}},
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
			name: "invalid report",
			req:  &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_report_query"}},
			parameters: cu.IM{
				"code": "XXX1731101982N123", "report_key": "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
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
