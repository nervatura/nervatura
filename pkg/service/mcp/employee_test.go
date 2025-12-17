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

func Test_employeeCreateHandler(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		req       *mcp.CallToolRequest
		inputData employeeCreate
		wantErr   bool
		ds        *api.DataStore
	}{
		{
			name: "success",
			req:  &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_employee_create"}},
			inputData: employeeCreate{
				employeeContact: employeeContact{
					FirstName: "Test Employee",
					Surname:   "Test Employee",
				},
				employeeAddress: employeeAddress{
					Country: "Test Country",
					State:   "Test State",
					ZipCode: "Test Zip Code",
					City:    "Test City",
					Street:  "Test Street",
				},
				EmployeeMeta: md.EmployeeMeta{
					StartDate: "2025-01-01",
				},
				EmployeeMap: cu.IM{
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
			name: "missing_code",
			req:  &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_employee_create"}},
			inputData: employeeCreate{
				employeeContact: employeeContact{
					FirstName: "Test Employee",
					Surname:   "Test Employee",
				},
				employeeAddress: employeeAddress{
					Country: "Test Country",
					State:   "Test State",
					ZipCode: "Test Zip Code",
					City:    "Test City",
					Street:  "Test Street",
				},
				EmployeeMeta: md.EmployeeMeta{
					Tags: []string{"test"},
				},
				EmployeeMap: cu.IM{
					"test": "test",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			_, _, gotErr := employeeCreateHandler(ctx, tt.req, tt.inputData)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("employeeCreateHandler() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("employeeCreateHandler() succeeded unexpectedly")
			}
		})
	}
}

func Test_employeeUpdateHandler(t *testing.T) {
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
			req:  &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_employee_update"}},
			inputData: cu.IM{
				"code":    "CUS123456",
				"surname": "Test Employee",
				"city":    "Test City",
				"notes":   "Test Notes",
				"employee_map": cu.IM{
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
							return []cu.IM{{"id": 1, "code": "CUS123456",
								"address":       cu.IM{"country": "Test Country", "state": "Test State", "zip_code": "Test Zip Code", "city": "Test City", "street": "Test Street"},
								"contact":       cu.IM{"first_name": "Test Employee", "surname": "Test Employee"},
								"employee_meta": cu.IM{"start_date": "2025-01-01", "tags": []string{"test"}, "notes": "Test Notes"},
								"employee_map":  cu.IM{"test": "test"}},
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
					return []byte(`{"code": "CUS123456"}`), nil
				},
			},
		},
		{
			name: "missing_employee",
			req:  &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_employee_update"}},
			inputData: cu.IM{
				"code":    "CUS123456",
				"surname": "Test Employee",
				"city":    "Test City",
				"notes":   "Test Notes",
				"employee_map": cu.IM{
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
					return []byte(`{"code": "CUS123456"}`), nil
				},
			},
		},
		{
			name: "missing_code",
			req:  &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_employee_update"}},
			inputData: cu.IM{
				"surname": "Test Employee",
				"city":    "Test City",
				"notes":   "Test Notes",
				"employee_map": cu.IM{
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
					return []byte(`{"code": "CUS123456"}`), nil
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			_, _, gotErr := employeeUpdateHandler(ctx, tt.req, tt.inputData)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("employeeUpdateHandler() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("employeeUpdateHandler() succeeded unexpectedly")
			}
		})
	}
}

func Test_employeeLoadData(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		data    any
		wantErr bool
	}{
		{
			name: "success",
			data: cu.IM{
				"code": "CUS123456",
				"address": cu.IM{
					"country": "Test Country",
				},
				"contact": cu.IM{
					"first_name": "Test Employee",
					"surname":    "Test Employee",
				},
				"employee_meta": cu.IM{
					"start_date": "2025-01-01",
					"tags":       []string{"test"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, gotErr := employeeLoadData(tt.data)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("employeeLoadData() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("employeeLoadData() succeeded unexpectedly")
			}
		})
	}
}
