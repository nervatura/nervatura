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

func Test_mapCreateHandler(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		req       *mcp.CallToolRequest
		inputData md.ConfigMap
		wantErr   bool
		ds        *api.DataStore
	}{
		{
			name: "success",
			req:  &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_map_create"}},
			inputData: md.ConfigMap{
				FieldName:   "Test Field",
				FieldType:   md.FieldTypeString,
				Description: "Test Description",
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
			name: "missing_name",
			req:  &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_map_create"}},
			inputData: md.ConfigMap{
				FieldType:   md.FieldTypeString,
				Description: "Test Description",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			_, _, gotErr := mapCreateHandler(ctx, tt.req, tt.inputData)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("mapCreateHandler() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("mapCreateHandler() succeeded unexpectedly")
			}
		})
	}
}

func Test_mapUpdate(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		req       *mcp.CallToolRequest
		inputData cu.IM
		wantErr   bool
		ds        *api.DataStore
	}{
		{
			name: "success",
			req:  &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_map_update"}},
			inputData: cu.IM{
				"code":        "CUS123456",
				"field_name":  "Test Field",
				"field_type":  md.FieldTypeString,
				"description": "Test Description",
			},
			wantErr: false,
			ds: &api.DataStore{
				Db: &td.TestDriver{
					Config: cu.IM{
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "CUS123456", "data": `{"field_name": "Test Field", "field_type": "FIELD_TYPE_STRING", "description": "Test Description"}`}}, nil
						},
						"GetDataField": func(data any, JSONName string) (fieldName string, fieldValue interface{}) {
							return JSONName, "value"
						},
					},
				},
				ConvertFromByte: func(data []byte, v any) error {
					return cu.ConvertFromByte(data, v)
				},
				ConvertToByte: func(v any) ([]byte, error) {
					return cu.ConvertToByte(v)
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
		},
		{
			name: "invalid code",
			req:  &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_map_update"}},
			inputData: cu.IM{
				"code":        "CUS123456",
				"field_name":  "Test Field",
				"field_type":  md.FieldTypeString,
				"description": "Test Description",
			},
			wantErr: true,
			ds: &api.DataStore{
				Db: &td.TestDriver{
					Config: cu.IM{
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{}, nil
						},
						"GetDataField": func(data any, JSONName string) (fieldName string, fieldValue interface{}) {
							return JSONName, "value"
						},
					},
				},
				ConvertFromByte: func(data []byte, v any) error {
					return cu.ConvertFromByte(data, v)
				},
				ConvertToByte: func(v any) ([]byte, error) {
					return cu.ConvertToByte(v)
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
		},
		{
			name: "missing code",
			req:  &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_map_update"}},
			inputData: cu.IM{
				"field_name":  "Test Field",
				"field_type":  md.FieldTypeString,
				"description": "Test Description",
			},
			wantErr: true,
			ds: &api.DataStore{
				Db: &td.TestDriver{
					Config: cu.IM{
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{}, nil
						},
						"GetDataField": func(data any, JSONName string) (fieldName string, fieldValue interface{}) {
							return JSONName, "value"
						},
					},
				},
				ConvertFromByte: func(data []byte, v any) error {
					return cu.ConvertFromByte(data, v)
				},
				ConvertToByte: func(v any) ([]byte, error) {
					return cu.ConvertToByte(v)
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			_, _, gotErr := mapUpdate(ctx, tt.req, tt.inputData)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("mapUpdate() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("mapUpdate() succeeded unexpectedly")
			}
		})
	}
}
