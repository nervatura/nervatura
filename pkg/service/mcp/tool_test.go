package mcp

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log/slog"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
)

func Test_modelQuery(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		req        *mcp.CallToolRequest
		parameters cu.IM
		wantErr    bool
		ds         *api.DataStore
	}{
		{
			name: "invalid tool",
			req:  &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "invalid"}},
			parameters: cu.IM{
				"limit":  10,
				"offset": 0,
			},
			wantErr: true,
		},
		{
			name: "customer_query",
			req:  &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_customer_query"}},
			parameters: cu.IM{
				"limit":         10,
				"offset":        0,
				"customer_type": "CUSTOMER_COMPANY",
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1}}, nil
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			wantErr: false,
		},
		{
			name: "product_query",
			req:  &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_product_query"}},
			parameters: cu.IM{
				"limit":        10,
				"offset":       0,
				"product_type": "PRODUCT_ITEM",
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1}}, nil
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			_, _, gotErr := modelQuery(ctx, tt.req, tt.parameters)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("modelQuery() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("modelQuery() succeeded unexpectedly")
			}

		})
	}
}

func Test_modelUpdate(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		req       *mcp.CallToolRequest
		inputData cu.IM
		wantErr   bool
		ds        *api.DataStore
	}{
		{
			name: "missing code",
			req:  &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_customer_update"}},
			inputData: cu.IM{
				"code": "",
			},
			wantErr: true,
		},
		{
			name: "invalid tool",
			req:  &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "invalid"}},
			inputData: cu.IM{
				"code": "123456",
			},
			wantErr: true,
		},
		{
			name: "success",
			req:  &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_customer_update"}},
			inputData: cu.IM{
				"code": "123456",
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "123456"}}, nil
						},
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
					},
				},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
				ReadAll: func(r io.Reader) ([]byte, error) {
					return []byte(`{"id": 1, "code": "123456"}`), nil
				},
				ConvertToByte: func(v any) ([]byte, error) {
					return cu.ConvertToByte(v)
				},
				ConvertFromByte: func(data []byte, v any) error {
					return cu.ConvertFromByte(data, v)
				},
				GetDataField: func(data any, JSONName string) (fieldName string, fieldValue interface{}) {
					return JSONName, "value"
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			_, _, gotErr := modelUpdate(ctx, tt.req, tt.inputData)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("modelUpdate() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("modelUpdate() succeeded unexpectedly")
			}
		})
	}
}

func Test_extendQuery(t *testing.T) {
	type args struct {
		//ctx       context.Context
		req       *mcp.CallToolRequest
		inputData cu.IM
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		ds      *api.DataStore
	}{
		{
			name: "invalid tool",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "invalid"}},
				inputData: cu.IM{"model": "customer", "limit": 10, "offset": 0},
			},
			wantErr: true,
		},
		{
			name: "contact_customer",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_contact_query"}},
				inputData: cu.IM{"model": "customer", "limit": 10, "offset": 0, "surname": "Doe"},
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "123456"}}, nil
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			wantErr: false,
		},
		{
			name: "contact_project",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_contact_query"}},
				inputData: cu.IM{"model": "project", "limit": 10, "offset": 0, "surname": "Doe"},
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "123456"}}, nil
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			wantErr: false,
		},
		{
			name: "contact_place",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_contact_query"}},
				inputData: cu.IM{"model": "place", "limit": 10, "offset": 0, "surname": "Doe"},
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "123456"}}, nil
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			wantErr: false,
		},
		{
			name: "contact_invalid_model",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_contact_query"}},
				inputData: cu.IM{"model": "invalid", "limit": 10, "offset": 0, "surname": "Doe"},
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "123456"}}, nil
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			wantErr: false,
		},
		{
			name: "address_customer",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_address_query"}},
				inputData: cu.IM{"model": "customer", "limit": 10, "offset": 0, "city": "New York"},
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "123456"}}, nil
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			wantErr: false,
		},
		{
			name: "address_project",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_address_query"}},
				inputData: cu.IM{"model": "project", "limit": 10, "offset": 0, "city": "New York"},
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "123456"}}, nil
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			wantErr: false,
		},
		{
			name: "address_invalid_model",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_address_query"}},
				inputData: cu.IM{"model": "invalid", "limit": 10, "offset": 0, "city": "New York"},
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "123456"}}, nil
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			wantErr: false,
		},
		{
			name: "event_customer",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_event_query"}},
				inputData: cu.IM{"model": "customer", "limit": 10, "offset": 0, "subject": "Meeting"},
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "123456"}}, nil
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			wantErr: false,
		},
		{
			name: "event_project",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_event_query"}},
				inputData: cu.IM{"model": "project", "limit": 10, "offset": 0, "subject": "Meeting"},
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "123456"}}, nil
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			wantErr: false,
		},
		{
			name: "event_employee",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_event_query"}},
				inputData: cu.IM{"model": "employee", "limit": 10, "offset": 0, "subject": "Meeting"},
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "123456"}}, nil
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			wantErr: false,
		},
		{
			name: "event_product",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_event_query"}},
				inputData: cu.IM{"model": "product", "limit": 10, "offset": 0, "subject": "Meeting"},
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "123456"}}, nil
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			wantErr: false,
		},
		{
			name: "event_tool",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_event_query"}},
				inputData: cu.IM{"model": "tool", "limit": 10, "offset": 0, "subject": "Meeting"},
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "123456"}}, nil
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			wantErr: false,
		},
		{
			name: "event_invalid_model",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_event_query"}},
				inputData: cu.IM{"model": "invalid", "limit": 10, "offset": 0, "subject": "Meeting"},
			},
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "123456"}}, nil
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			_, _, err := extendQuery(ctx, tt.args.req, tt.args.inputData)
			if (err != nil) != tt.wantErr {
				t.Errorf("extendQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_extendUpdate(t *testing.T) {
	type args struct {
		//ctx       context.Context
		req       *mcp.CallToolRequest
		inputData cu.IM
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		ds      *api.DataStore
	}{
		{
			name: "missing code",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_contact_update"}},
				inputData: cu.IM{"code": ""},
			},
			wantErr: true,
		},
		{
			name: "invalid tool",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "invalid"}},
				inputData: cu.IM{"code": "123456"},
			},
			wantErr: true,
		},
		{
			name: "contact_customer",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_contact_update"}},
				inputData: cu.IM{"code": "CUS123456", "index": 0, "surname": "Doe", "contact_map": cu.IM{"fiels": "value"}},
			},
			wantErr: false,
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "CUS123456", "contacts": []cu.IM{{"surname": "Doe"}}}}, nil
						},
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
		},
		{
			name: "contact_project",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_contact_update"}},
				inputData: cu.IM{"code": "PRJ123456", "index": 0, "surname": "Doe"},
			},
			wantErr: true,
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{}, nil
						},
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
		},
		{
			name: "contact_place",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_contact_update"}},
				inputData: cu.IM{"code": "PLA123456", "index": 0, "surname": "Doe"},
			},
			wantErr: true,
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "CUS123456", "contacts": []cu.IM{{"surname": "Doe"}}}}, nil
						},
						"Update": func(data md.Update) (int64, error) {
							return 1, errors.New("error")
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
		},
		{
			name: "contact_invalid",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_contact_update"}},
				inputData: cu.IM{"code": "XXX123456", "index": 0, "surname": "Doe"},
			},
			wantErr: true,
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "CUS123456", "contacts": []cu.IM{{"surname": "Doe"}}}}, nil
						},
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
		},
		{
			name: "address_customer",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_address_update"}},
				inputData: cu.IM{"code": "CUS123456", "index": 2, "city": "New York"},
			},
			wantErr: true,
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "CUS123456", "addresses": []cu.IM{{"city": "New York"}}}}, nil
						},
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
		},
		{
			name: "address_project",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_address_update"}},
				inputData: cu.IM{"code": "PRJ123456", "index": 0, "city": "New York", "missing": 12345},
			},
			wantErr: false,
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "CUS123456", "addresses": []cu.IM{{"city": "New York"}}}}, nil
						},
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
		},
		{
			name: "address_missing",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_address_update"}},
				inputData: cu.IM{"code": "XXX123456", "index": 0, "city": "New York", "missing": 12345},
			},
			wantErr: true,
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "CUS123456", "addresses": []cu.IM{{"city": "New York"}}}}, nil
						},
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
		},
		{
			name: "event_customer",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_event_update"}},
				inputData: cu.IM{"code": "CUS123456", "index": 0, "subject": "Meeting"},
			},
			wantErr: false,
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "CUS123456", "events": []cu.IM{{"subject": "Meeting"}}}}, nil
						},
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
		},
		{
			name: "event_project",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_event_update"}},
				inputData: cu.IM{"code": "PRJ123456", "index": 0, "subject": "Meeting"},
			},
			wantErr: false,
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "CUS123456", "events": []cu.IM{{"subject": "Meeting"}}}}, nil
						},
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
		},
		{
			name: "event_employee",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_event_update"}},
				inputData: cu.IM{"code": "EMP123456", "index": 0, "subject": "Meeting"},
			},
			wantErr: false,
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "CUS123456", "events": []cu.IM{{"subject": "Meeting"}}}}, nil
						},
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
		},
		{
			name: "event_tool",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_event_update"}},
				inputData: cu.IM{"code": "SER123456", "index": 0, "subject": "Meeting"},
			},
			wantErr: false,
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "CUS123456", "events": []cu.IM{{"subject": "Meeting"}}}}, nil
						},
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
		},
		{
			name: "event_product",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_event_update"}},
				inputData: cu.IM{"code": "PRD123456", "index": 0, "subject": "Meeting"},
			},
			wantErr: false,
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "CUS123456", "events": []cu.IM{{"subject": "Meeting"}}}}, nil
						},
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
		},
		{
			name: "event_missing",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_event_update"}},
				inputData: cu.IM{"code": "XXX123456", "index": 0, "subject": "Meeting"},
			},
			wantErr: true,
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "CUS123456", "events": []cu.IM{{"subject": "Meeting"}}}}, nil
						},
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			_, _, err := extendUpdate(ctx, tt.args.req, tt.args.inputData)
			if (err != nil) != tt.wantErr {
				t.Errorf("extendUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_extendCreate(t *testing.T) {
	type args struct {
		//ctx       context.Context
		req       *mcp.CallToolRequest
		inputData cu.IM
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		ds      *api.DataStore
	}{
		{
			name: "missing code",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_contact_update"}},
				inputData: cu.IM{"code": ""},
			},
			wantErr: true,
		},
		{
			name: "invalid tool",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "invalid"}},
				inputData: cu.IM{"code": "123456"},
			},
			wantErr: true,
		},
		{
			name: "contact_customer",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_contact_update"}},
				inputData: cu.IM{"code": "CUS123456", "index": 0, "surname": "Doe", "contact_map": cu.IM{"fiels": "value"}},
			},
			wantErr: false,
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "CUS123456", "contacts": []cu.IM{{"surname": "Doe"}}}}, nil
						},
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
		},
		{
			name: "event_product",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_event_update"}},
				inputData: cu.IM{"code": "PRD123456", "index": 0, "subject": "Meeting", "event_map": cu.IM{"fiels": "value"}},
			},
			wantErr: false,
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "PRD123456", "events": []cu.IM{{"subject": "Meeting"}}}}, nil
						},
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
		},
		{
			name: "invalid_code",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_contact_update"}},
				inputData: cu.IM{"code": "XXX123456", "index": 0, "surname": "Doe", "contact_map": cu.IM{"fiels": "value"}},
			},
			wantErr: true,
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "code": "CUS123456", "contacts": []cu.IM{{"surname": "Doe"}}}}, nil
						},
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
		},
		{
			name: "missing_code",
			args: args{
				req:       &mcp.CallToolRequest{Params: &mcp.CallToolParamsRaw{Name: "nervatura_contact_update"}},
				inputData: cu.IM{"code": "CUS123456", "index": 0, "surname": "Doe", "contact_map": cu.IM{"fiels": "value"}},
			},
			wantErr: true,
			ds: &api.DataStore{
				Db: &md.TestDriver{
					Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{}, nil
						},
						"Update": func(data md.Update) (int64, error) {
							return 1, nil
						},
					},
				},
				Config: cu.IM{},
				AppLog: slog.New(slog.NewTextHandler(bytes.NewBufferString(""), nil)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), md.DataStoreCtxKey, tt.ds)
			_, _, err := extendCreate(ctx, tt.args.req, tt.args.inputData)
			if (err != nil) != tt.wantErr {
				t.Errorf("extendCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
