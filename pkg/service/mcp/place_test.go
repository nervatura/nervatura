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
)

func Test_placeLoadData(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		data    any
		wantErr bool
	}{
		{
			name: "success",
			data: cu.IM{
				"place_type": md.PlaceTypeWarehouse.String(),
				"place_name": "Main warehouse",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, gotErr := placeLoadData(tt.data)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("placeLoadData() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("placeLoadData() succeeded unexpectedly")
			}
		})
	}
}

func Test_placeCreateHandler(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		req       *mcp.CallToolRequest
		inputData placeCreate
		wantErr   bool
		ds        *api.DataStore
	}{
		{
			name: "success",
			req:  &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_place_create"}},
			inputData: placeCreate{
				PlaceType:    md.PlaceTypeWarehouse,
				PlaceName:    "Main warehouse",
				CurrencyCode: "EUR",
				placeAddress: placeAddress{
					Country: "Test Country",
					State:   "Test State",
					ZipCode: "Test Zip Code",
					City:    "Test City",
					Street:  "Test Street",
				},
				PlaceMeta: md.PlaceMeta{},
				PlaceMap: cu.IM{
					"test": "test",
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
						"GetDataField": func(data any, JSONName string) (fieldName string, fieldValue interface{}) {
							return JSONName, "value"
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
		},
		{
			name: "missing_place_name",
			req:  &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_place_create"}},
			inputData: placeCreate{
				PlaceType:    md.PlaceTypeWarehouse,
				CurrencyCode: "EUR",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			_, _, gotErr := placeCreateHandler(ctx, tt.req, tt.inputData)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("placeCreateHandler() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("placeCreateHandler() succeeded unexpectedly")
			}
		})
	}
}

func Test_placeUpdateHandler(t *testing.T) {
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
			req:  &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_place_update"}},
			inputData: cu.IM{
				"code":  "CUS123456",
				"city":  "Test City",
				"notes": "Test Notes",
				"place_map": cu.IM{
					"test": "test",
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
							return []cu.IM{{"id": 1, "code": "PLA123456",
								"address":    cu.IM{"country": "Test Country", "state": "Test State", "zip_code": "Test Zip Code", "city": "Test City", "street": "Test Street"},
								"place_meta": cu.IM{"tags": []string{"test"}, "notes": "Test Notes"},
								"place_map":  cu.IM{"test": "test"}},
							}, nil
						},
						"GetDataField": func(data any, JSONName string) (fieldName string, fieldValue interface{}) {
							return JSONName, "value"
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ConvertToType: func(data interface{}, result any) (err error) {
					return nil
				},
				ConvertToByte: func(data any) ([]byte, error) {
					return []byte(`{"code": "PLA123456"}`), nil
				},
			},
		},
		{
			name: "missing_place",
			req:  &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_place_update"}},
			inputData: cu.IM{
				"code":    "CUS123456",
				"surname": "Test Employee",
				"city":    "Test City",
				"notes":   "Test Notes",
				"place_map": cu.IM{
					"test": "test",
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
							return []cu.IM{}, nil
						},
						"GetDataField": func(data any, JSONName string) (fieldName string, fieldValue interface{}) {
							return JSONName, "value"
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ConvertToType: func(data interface{}, result any) (err error) {
					return nil
				},
				ConvertToByte: func(data any) ([]byte, error) {
					return []byte(`{"code": "PLA123456"}`), nil
				},
			},
		},
		{
			name: "missing_code",
			req:  &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_place_update"}},
			inputData: cu.IM{
				"surname": "Test Employee",
				"city":    "Test City",
				"notes":   "Test Notes",
				"place_map": cu.IM{
					"test": "test",
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
							return []cu.IM{}, nil
						},
						"GetDataField": func(data any, JSONName string) (fieldName string, fieldValue interface{}) {
							return JSONName, "value"
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ConvertToType: func(data interface{}, result any) (err error) {
					return nil
				},
				ConvertToByte: func(data any) ([]byte, error) {
					return []byte(`{"code": "PLA123456"}`), nil
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			_, _, gotErr := placeUpdateHandler(ctx, tt.req, tt.inputData)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("placeUpdateHandler() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("placeUpdateHandler() succeeded unexpectedly")
			}
		})
	}
}
