package mcp

import (
	"bytes"
	"context"
	"log/slog"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	td "github.com/nervatura/nervatura/v6/test/driver"
)

func Test_projectLoadData(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		data    any
		wantErr bool
	}{
		{
			name: "success",
			data: cu.IM{
				"project_name":  "Test Project",
				"customer_code": "CUS1731101982N123",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, gotErr := projectLoadData(tt.data)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("projectLoadData() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("projectLoadData() succeeded unexpectedly")
			}
		})
	}
}

func Test_projectCreateHandler(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		req       *mcp.CallToolRequest
		inputData projectCreate
		wantErr   bool
		ds        *api.DataStore
	}{
		{
			name: "success",
			req:  &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_project_create"}},
			inputData: projectCreate{
				ProjectName:  "Test Project",
				CustomerCode: "CUS1731101982N123",
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
			name:      "missing_name",
			req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_project_create"}},
			inputData: projectCreate{},
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			_, _, gotErr := projectCreateHandler(ctx, tt.req, tt.inputData)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("projectCreateHandler() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("projectCreateHandler() succeeded unexpectedly")
			}
		})
	}
}
