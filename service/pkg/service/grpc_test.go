//go:build grpc || all
// +build grpc all

package service

import (
	"context"
	"reflect"
	"testing"

	db "github.com/nervatura/nervatura/service/pkg/database"
	nt "github.com/nervatura/nervatura/service/pkg/nervatura"
	pb "github.com/nervatura/nervatura/service/pkg/proto"
)

func TestRPCService_itemMap(t *testing.T) {
	type fields struct {
		Config                 map[string]interface{}
		GetNervaStore          func(database string) *nt.NervaStore
		GetTokenKeys           func() map[string]map[string]string
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		key  string
		data nt.IM
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *pb.ResponseGet_Value
	}{
		{
			name: "address",
			args: args{
				key:  "address",
				data: nt.IM{},
			},
			want: &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Address{Address: &pb.Address{
					Metadata: []*pb.MetaData{},
				}}},
		},
		{
			name: "barcode",
			args: args{
				key:  "barcode",
				data: nt.IM{},
			},
			want: &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Barcode{Barcode: &pb.Barcode{}}},
		},
		{
			name: "contact",
			args: args{
				key:  "contact",
				data: nt.IM{},
			},
			want: &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Contact{Contact: &pb.Contact{
					Metadata: []*pb.MetaData{},
				}}},
		},
		{
			name: "currency",
			args: args{
				key:  "currency",
				data: nt.IM{},
			},
			want: &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Currency{Currency: &pb.Currency{
					Metadata: []*pb.MetaData{},
				}}},
		},
		{
			name: "customer",
			args: args{
				key:  "customer",
				data: nt.IM{},
			},
			want: &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Customer{Customer: &pb.Customer{
					Metadata: []*pb.MetaData{},
				}}},
		},
		{
			name: "deffield",
			args: args{
				key:  "deffield",
				data: nt.IM{},
			},
			want: &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Deffield{Deffield: &pb.Deffield{}}},
		},
		{
			name: "employee",
			args: args{
				key:  "employee",
				data: nt.IM{},
			},
			want: &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Employee{Employee: &pb.Employee{
					Metadata: []*pb.MetaData{},
				}}},
		},
		{
			name: "event",
			args: args{
				key:  "event",
				data: nt.IM{},
			},
			want: &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Event{Event: &pb.Event{
					Metadata: []*pb.MetaData{},
				}}},
		},
		{
			name: "fieldvalue",
			args: args{
				key:  "fieldvalue",
				data: nt.IM{},
			},
			want: &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Fieldvalue{Fieldvalue: &pb.Fieldvalue{}}},
		},
		{
			name: "groups",
			args: args{
				key:  "groups",
				data: nt.IM{},
			},
			want: &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Groups{Groups: &pb.Groups{}}},
		},
		{
			name: "item",
			args: args{
				key:  "item",
				data: nt.IM{},
			},
			want: &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Item{Item: &pb.Item{
					Metadata: []*pb.MetaData{},
				}}},
		},
		{
			name: "link",
			args: args{
				key:  "link",
				data: nt.IM{},
			},
			want: &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Link{Link: &pb.Link{
					Metadata: []*pb.MetaData{},
				}}},
		},
		{
			name: "log",
			args: args{
				key:  "log",
				data: nt.IM{},
			},
			want: &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Log{Log: &pb.Log{
					Metadata: []*pb.MetaData{},
				}}},
		},
		{
			name: "movement",
			args: args{
				key:  "movement",
				data: nt.IM{},
			},
			want: &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Movement{Movement: &pb.Movement{
					Metadata: []*pb.MetaData{},
				}}},
		},
		{
			name: "numberdef",
			args: args{
				key:  "numberdef",
				data: nt.IM{},
			},
			want: &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Numberdef{Numberdef: &pb.Numberdef{}}},
		},
		{
			name: "pattern",
			args: args{
				key:  "pattern",
				data: nt.IM{},
			},
			want: &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Pattern{Pattern: &pb.Pattern{}}},
		},
		{
			name: "payment",
			args: args{
				key:  "payment",
				data: nt.IM{},
			},
			want: &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Payment{Payment: &pb.Payment{
					Metadata: []*pb.MetaData{},
				}}},
		},
		{
			name: "place",
			args: args{
				key:  "place",
				data: nt.IM{},
			},
			want: &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Place{Place: &pb.Place{
					Metadata: []*pb.MetaData{},
				}}},
		},
		{
			name: "price",
			args: args{
				key:  "price",
				data: nt.IM{},
			},
			want: &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Price{Price: &pb.Price{
					Metadata: []*pb.MetaData{},
				}}},
		},
		{
			name: "product",
			args: args{
				key:  "product",
				data: nt.IM{},
			},
			want: &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Product{Product: &pb.Product{
					Metadata: []*pb.MetaData{},
				}}},
		},
		{
			name: "project",
			args: args{
				key:  "project",
				data: nt.IM{},
			},
			want: &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Project{Project: &pb.Project{
					Metadata: []*pb.MetaData{},
				}}},
		},
		{
			name: "rate",
			args: args{
				key:  "rate",
				data: nt.IM{},
			},
			want: &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Rate{Rate: &pb.Rate{
					Metadata: []*pb.MetaData{},
				}}},
		},
		{
			name: "tax",
			args: args{
				key:  "tax",
				data: nt.IM{},
			},
			want: &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Tax{Tax: &pb.Tax{
					Metadata: []*pb.MetaData{},
				}}},
		},
		{
			name: "tool",
			args: args{
				key:  "tool",
				data: nt.IM{},
			},
			want: &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Tool{Tool: &pb.Tool{
					Metadata: []*pb.MetaData{},
				}}},
		},
		{
			name: "trans",
			args: args{
				key:  "trans",
				data: nt.IM{},
			},
			want: &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Trans{Trans: &pb.Trans{
					Metadata: []*pb.MetaData{},
				}}},
		},
		{
			name: "ui_audit",
			args: args{
				key:  "ui_audit",
				data: nt.IM{},
			},
			want: &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_UiAudit{UiAudit: &pb.UiAudit{}}},
		},
		{
			name: "ui_menu",
			args: args{
				key:  "ui_menu",
				data: nt.IM{},
			},
			want: &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_UiMenu{UiMenu: &pb.UiMenu{}}},
		},
		{
			name: "ui_menufields",
			args: args{
				key:  "ui_menufields",
				data: nt.IM{},
			},
			want: &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_UiMenufields{UiMenufields: &pb.UiMenufields{}}},
		},
		{
			name: "ui_message",
			args: args{
				key:  "ui_message",
				data: nt.IM{},
			},
			want: &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_UiMessage{UiMessage: &pb.UiMessage{}}},
		},
		{
			name: "ui_printqueue",
			args: args{
				key:  "ui_printqueue",
				data: nt.IM{},
			},
			want: &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_UiPrintqueue{UiPrintqueue: &pb.UiPrintqueue{}}},
		},
		{
			name: "ui_report",
			args: args{
				key:  "ui_report",
				data: nt.IM{},
			},
			want: &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_UiReport{UiReport: &pb.UiReport{}}},
		},
		{
			name: "ui_userconfig",
			args: args{
				key:  "ui_userconfig",
				data: nt.IM{},
			},
			want: &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_UiUserconfig{UiUserconfig: &pb.UiUserconfig{}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &RPCService{
				Config:                 tt.fields.Config,
				GetNervaStore:          tt.fields.GetNervaStore,
				GetTokenKeys:           tt.fields.GetTokenKeys,
				UnimplementedAPIServer: tt.fields.UnimplementedAPIServer,
			}
			if got := srv.itemMap(tt.args.key, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RPCService.itemMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getValue(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want *pb.Value
	}{
		{
			name: "nil_value",
			args: args{
				value: nil,
			},
			want: &pb.Value{Value: &pb.Value_Text{Text: "null"}},
		},
		{
			name: "bool_value",
			args: args{
				value: true,
			},
			want: &pb.Value{Value: &pb.Value_Boolean{Boolean: true}},
		},
		{
			name: "int_value",
			args: args{
				value: int64(123),
			},
			want: &pb.Value{Value: &pb.Value_Number{Number: float64(123)}},
		},
		{
			name: "float_value",
			args: args{
				value: float64(123),
			},
			want: &pb.Value{Value: &pb.Value_Number{Number: float64(123)}},
		},
		{
			name: "string_value",
			args: args{
				value: "value",
			},
			want: &pb.Value{Value: &pb.Value_Text{Text: "value"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getValue(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getIValue(t *testing.T) {
	type args struct {
		value *pb.Value
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Boolean",
			args: args{
				value: &pb.Value{Value: &pb.Value_Boolean{Boolean: true}},
			},
			want: true,
		},
		{
			name: "Number",
			args: args{
				value: &pb.Value{Value: &pb.Value_Number{Number: 123}},
			},
			want: float64(123),
		},
		{
			name: "empty_Text",
			args: args{
				value: &pb.Value{Value: &pb.Value_Text{Text: ""}},
			},
			want: nil,
		},
		{
			name: "Text",
			args: args{
				value: &pb.Value{Value: &pb.Value_Text{Text: "value"}},
			},
			want: "value",
		},
		{
			name: "invalid",
			args: args{
				value: &pb.Value{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIValue(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRPCService_TokenAuth(t *testing.T) {
	type fields struct {
		Config                 map[string]interface{}
		GetNervaStore          func(database string) *nt.NervaStore
		GetTokenKeys           func() map[string]map[string]string
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		authorization []string
		parent        context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "token_ok",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				authorization: []string{"Bearer " + testData.adminToken},
				parent:        context.Background(),
			},
			wantErr: false,
		},
		{
			name: "token_missing",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				authorization: []string{},
				parent:        context.Background(),
			},
			wantErr: true,
		},
		{
			name: "token_empty",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				authorization: []string{"010101010101"},
				parent:        context.Background(),
			},
			wantErr: true,
		},
		{
			name: "token_wrong",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				authorization: []string{""},
				parent:        context.Background(),
			},
			wantErr: true,
		},
		{
			name: "token_expired",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				authorization: []string{"eyJhbGciOiJIUzI1NiIsImtpZCI6IjhiZDBlMDI1MDk0ODJmNThjZWZkM2MwZWNkNDFmZjBlIiwidHlwIjoiSldUIn0.eyJ1c2VybmFtZSI6ImFkbWluIiwiZGF0YWJhc2UiOiJkZW1vIiwiZXhwIjoxNjI4NTM1ODQzLCJpc3MiOiJuZXJ2YXR1cmEifQ.ErrxPYNENXc7tvi7PVzE8z6qe8QEEtnEgdcPEqzAvos"},
				parent:        context.Background(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &RPCService{
				Config:                 tt.fields.Config,
				GetNervaStore:          tt.fields.GetNervaStore,
				GetTokenKeys:           tt.fields.GetTokenKeys,
				UnimplementedAPIServer: tt.fields.UnimplementedAPIServer,
			}
			_, err := srv.TokenAuth(tt.args.authorization, tt.args.parent)
			if (err != nil) != tt.wantErr {
				t.Errorf("RPCService.TokenAuth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRPCService_ApiKeyAuth(t *testing.T) {
	type fields struct {
		Config                 map[string]interface{}
		GetNervaStore          func(database string) *nt.NervaStore
		GetTokenKeys           func() map[string]map[string]string
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		authorization []string
		parent        context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "key_ok",
			fields: fields{
				Config: nt.IM{
					"NT_API_KEY": testData.apiKey,
				},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{})
				},
			},
			args: args{
				authorization: []string{testData.apiKey},
				parent:        context.Background(),
			},
			wantErr: false,
		},
		{
			name: "missing_key",
			fields: fields{
				Config: nt.IM{
					"NT_API_KEY": testData.apiKey,
				},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{})
				},
			},
			args: args{
				authorization: []string{},
				parent:        context.Background(),
			},
			wantErr: true,
		},
		{
			name: "empty_key",
			fields: fields{
				Config: nt.IM{
					"NT_API_KEY": testData.apiKey,
				},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{})
				},
			},
			args: args{
				authorization: []string{""},
				parent:        context.Background(),
			},
			wantErr: true,
		},
		{
			name: "invalid_key",
			fields: fields{
				Config: nt.IM{
					"NT_API_KEY": testData.apiKey,
				},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{})
				},
			},
			args: args{
				authorization: []string{"INVALID_KEY"},
				parent:        context.Background(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &RPCService{
				Config:                 tt.fields.Config,
				GetNervaStore:          tt.fields.GetNervaStore,
				GetTokenKeys:           tt.fields.GetTokenKeys,
				UnimplementedAPIServer: tt.fields.UnimplementedAPIServer,
			}
			_, err := srv.ApiKeyAuth(tt.args.authorization, tt.args.parent)
			if (err != nil) != tt.wantErr {
				t.Errorf("RPCService.ApiKeyAuth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRPCService_UserLogin(t *testing.T) {
	type fields struct {
		Config                 map[string]interface{}
		GetNervaStore          func(database string) *nt.NervaStore
		GetTokenKeys           func() map[string]map[string]string
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx context.Context
		req *pb.RequestUserLogin
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *pb.ResponseUserLogin
		wantErr bool
	}{
		{
			name: "login_ok",
			fields: fields{
				Config: nt.IM{
					"version": testData.version,
				},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE": testData.hashTable,
					})
				},
			},
			args: args{
				req: &pb.RequestUserLogin{
					Username: "admin",
					Database: "test",
				},
			},
			wantErr: false,
		},
		{
			name: "login_missing_database",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{})
				},
			},
			args: args{
				req: &pb.RequestUserLogin{
					Username: "admin",
				},
			},
			wantErr: true,
		},
		{
			name: "login_wrong_password",
			fields: fields{
				Config: nt.IM{
					"version": testData.version,
				},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE": testData.hashTable,
					})
				},
			},
			args: args{
				req: &pb.RequestUserLogin{
					Username: "admin",
					Password: "1a1a1a",
					Database: "test",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &RPCService{
				Config:                 tt.fields.Config,
				GetNervaStore:          tt.fields.GetNervaStore,
				GetTokenKeys:           tt.fields.GetTokenKeys,
				UnimplementedAPIServer: tt.fields.UnimplementedAPIServer,
			}
			_, err := srv.UserLogin(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("RPCService.UserLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRPCService_UserPassword(t *testing.T) {
	type fields struct {
		Config                 map[string]interface{}
		GetNervaStore          func(database string) *nt.NervaStore
		GetTokenKeys           func() map[string]map[string]string
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx context.Context
		req *pb.RequestUserPassword
	}
	config := nt.IM{
		"version":              testData.version,
		"NT_HASHTABLE":         testData.hashTable,
		"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "user_ok",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.adminToken)),
				req: &pb.RequestUserPassword{
					Username: "guest",
					Password: "123",
					Confirm:  "123",
				},
			},
			wantErr: false,
		},
		{
			name: "custnumber_ok",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.adminToken)),
				req: &pb.RequestUserPassword{
					Custnumber: "DMCUST/00001",
					Password:   "123",
					Confirm:    "123",
				},
			},
			wantErr: false,
		},
		{
			name: "customer_token_error",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.customerToken)),
				req: &pb.RequestUserPassword{
					Password: "123",
					Confirm:  "123",
				},
			},
			wantErr: true,
		},
		{
			name: "empty_error",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.adminToken)),
				req: &pb.RequestUserPassword{
					Password: "",
					Confirm:  "",
				},
			},
			wantErr: true,
		},
		{
			name: "user_scope_error",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.userToken)),
				req: &pb.RequestUserPassword{
					Username: "demo",
					Password: "",
					Confirm:  "",
				},
			},
			wantErr: true,
		},
		{
			name: "custnumber_scope_error",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.userToken)),
				req: &pb.RequestUserPassword{
					Custnumber: "DMCUST/00001",
					Password:   "",
					Confirm:    "",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &RPCService{
				Config:                 tt.fields.Config,
				GetNervaStore:          tt.fields.GetNervaStore,
				GetTokenKeys:           tt.fields.GetTokenKeys,
				UnimplementedAPIServer: tt.fields.UnimplementedAPIServer,
			}
			_, err := srv.UserPassword(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("RPCService.UserPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRPCService_TokenDecode(t *testing.T) {
	type fields struct {
		Config                 map[string]interface{}
		GetNervaStore          func(database string) *nt.NervaStore
		GetTokenKeys           func() map[string]map[string]string
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx context.Context
		req *pb.RequestTokenDecode
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "admin_token",
			args: args{
				req: &pb.RequestTokenDecode{
					Value: testData.adminToken,
				},
			},
			wantErr: false,
		},
		{
			name: "error_token",
			args: args{
				req: &pb.RequestTokenDecode{
					Value: "ASND1233kjkjdjk2222jkjkd",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &RPCService{
				Config:                 tt.fields.Config,
				GetNervaStore:          tt.fields.GetNervaStore,
				GetTokenKeys:           tt.fields.GetTokenKeys,
				UnimplementedAPIServer: tt.fields.UnimplementedAPIServer,
			}
			_, err := srv.TokenDecode(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("RPCService.TokenDecode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRPCService_TokenLogin(t *testing.T) {
	type fields struct {
		Config                 map[string]interface{}
		GetNervaStore          func(database string) *nt.NervaStore
		GetTokenKeys           func() map[string]map[string]string
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx context.Context
		req *pb.RequestEmpty
	}
	config := nt.IM{
		"version":              testData.version,
		"NT_HASHTABLE":         testData.hashTable,
		"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.ResponseTokenLogin
		wantErr bool
	}{
		{
			name: "user_token",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.adminToken)),
			},
			wantErr: false,
		},
		{
			name: "customer_token",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.customerToken)),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &RPCService{
				Config:                 tt.fields.Config,
				GetNervaStore:          tt.fields.GetNervaStore,
				GetTokenKeys:           tt.fields.GetTokenKeys,
				UnimplementedAPIServer: tt.fields.UnimplementedAPIServer,
			}
			_, err := srv.TokenLogin(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("RPCService.TokenLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRPCService_TokenRefresh(t *testing.T) {
	type fields struct {
		Config                 map[string]interface{}
		GetNervaStore          func(database string) *nt.NervaStore
		GetTokenKeys           func() map[string]map[string]string
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx context.Context
		req *pb.RequestEmpty
	}
	config := nt.IM{
		"version":              testData.version,
		"NT_HASHTABLE":         testData.hashTable,
		"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "refresh_ok",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.adminToken)),
			},
			wantErr: false,
		},
		{
			name: "refresh_unauthorized",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &RPCService{
				Config:                 tt.fields.Config,
				GetNervaStore:          tt.fields.GetNervaStore,
				GetTokenKeys:           tt.fields.GetTokenKeys,
				UnimplementedAPIServer: tt.fields.UnimplementedAPIServer,
			}
			_, err := srv.TokenRefresh(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("RPCService.TokenRefresh() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRPCService_Get(t *testing.T) {
	type fields struct {
		Config                 map[string]interface{}
		GetNervaStore          func(database string) *nt.NervaStore
		GetTokenKeys           func() map[string]map[string]string
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx context.Context
		req *pb.RequestGet
	}
	config := nt.IM{
		"version":              testData.version,
		"NT_HASHTABLE":         testData.hashTable,
		"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "get_ids_ok",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.adminToken)),
				req: &pb.RequestGet{
					Nervatype: pb.DataType_trans,
					Metadata:  true,
					Ids:       []int64{4},
				},
			},
			wantErr: false,
		},
		{
			name: "get_filter_ok",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.adminToken)),
				req: &pb.RequestGet{
					Nervatype: pb.DataType_customer,
					Metadata:  true,
					Filter:    []string{"custnumber;==;DMCUST/00001"},
				},
			},
			wantErr: false,
		},
		{
			name: "get_scope_error",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.customerToken)),
				req: &pb.RequestGet{
					Nervatype: pb.DataType_trans,
					Metadata:  true,
					Ids:       []int64{4},
				},
			},
			wantErr: true,
		},
		{
			name: "get_error",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.adminToken)),
				req: &pb.RequestGet{
					Nervatype: pb.DataType_customer,
					Metadata:  true,
					Filter:    []string{"credit;==;DMCUST/00001"},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &RPCService{
				Config:                 tt.fields.Config,
				GetNervaStore:          tt.fields.GetNervaStore,
				GetTokenKeys:           tt.fields.GetTokenKeys,
				UnimplementedAPIServer: tt.fields.UnimplementedAPIServer,
			}
			_, err := srv.Get(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("RPCService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRPCService_Update(t *testing.T) {
	type fields struct {
		Config                 map[string]interface{}
		GetNervaStore          func(database string) *nt.NervaStore
		GetTokenKeys           func() map[string]map[string]string
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx context.Context
		req *pb.RequestUpdate
	}
	config := nt.IM{
		"version":              testData.version,
		"NT_HASHTABLE":         testData.hashTable,
		"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "update_ok",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.adminToken)),
				req: &pb.RequestUpdate{
					Nervatype: pb.DataType_address,
					Items: []*pb.RequestUpdate_Item{
						{
							Values: map[string]*pb.Value{
								"zipcode":           {Value: &pb.Value_Text{Text: "12345"}},
								"city":              {Value: &pb.Value_Text{Text: "BigCity"}},
								"notes":             {Value: &pb.Value_Text{Text: "Create a new item by IDs"}},
								"address_metadata1": {Value: &pb.Value_Text{Text: "value1"}},
								"address_metadata2": {Value: &pb.Value_Text{Text: "value2~note2"}},
							},
							Keys: map[string]*pb.Value{
								"nervatype": {Value: &pb.Value_Text{Text: "customer"}},
								"ref_id":    {Value: &pb.Value_Text{Text: "customer/DMCUST/00001"}},
							}},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "update_scope_err",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.customerToken)),
				req: &pb.RequestUpdate{
					Nervatype: pb.DataType_address,
					Items: []*pb.RequestUpdate_Item{
						{
							Values: map[string]*pb.Value{
								"zipcode":           {Value: &pb.Value_Text{Text: "12345"}},
								"city":              {Value: &pb.Value_Text{Text: "BigCity"}},
								"notes":             {Value: &pb.Value_Text{Text: "Create a new item by IDs"}},
								"address_metadata1": {Value: &pb.Value_Text{Text: "value1"}},
								"address_metadata2": {Value: &pb.Value_Text{Text: "value2~note2"}},
							},
							Keys: map[string]*pb.Value{
								"nervatype": {Value: &pb.Value_Text{Text: "customer"}},
								"ref_id":    {Value: &pb.Value_Text{Text: "customer/DMCUST/00001"}},
							}},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "update_error",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.adminToken)),
				req: &pb.RequestUpdate{
					Nervatype: pb.DataType_address,
					Items: []*pb.RequestUpdate_Item{
						{
							Values: map[string]*pb.Value{
								"zipcode": {Value: &pb.Value_Text{Text: "12345"}},
								"city":    {Value: &pb.Value_Text{Text: "BigCity"}},
							},
							Keys: map[string]*pb.Value{
								"nervatype": {Value: &pb.Value_Text{Text: "customer"}},
								"ref_id":    {Value: &pb.Value_Text{Text: "customer/00001"}},
							}},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &RPCService{
				Config:                 tt.fields.Config,
				GetNervaStore:          tt.fields.GetNervaStore,
				GetTokenKeys:           tt.fields.GetTokenKeys,
				UnimplementedAPIServer: tt.fields.UnimplementedAPIServer,
			}
			if _, err := srv.Update(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("RPCService.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRPCService_Delete(t *testing.T) {
	type fields struct {
		Config                 map[string]interface{}
		GetNervaStore          func(database string) *nt.NervaStore
		GetTokenKeys           func() map[string]map[string]string
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx context.Context
		req *pb.RequestDelete
	}
	config := nt.IM{
		"version":              testData.version,
		"NT_HASHTABLE":         testData.hashTable,
		"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "delete_ok",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.adminToken)),
				req: &pb.RequestDelete{
					Nervatype: pb.DataType_address, Id: 55,
				},
			},
			wantErr: false,
		},
		{
			name: "delete_scope_error",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.customerToken)),
				req: &pb.RequestDelete{
					Nervatype: pb.DataType_address, Id: 55,
				},
			},
			wantErr: true,
		},
		{
			name: "delete_error",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.adminToken)),
				req: &pb.RequestDelete{
					Nervatype: pb.DataType_address, Key: "abrakadabra",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &RPCService{
				Config:                 tt.fields.Config,
				GetNervaStore:          tt.fields.GetNervaStore,
				GetTokenKeys:           tt.fields.GetTokenKeys,
				UnimplementedAPIServer: tt.fields.UnimplementedAPIServer,
			}
			if _, err := srv.Delete(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("RPCService.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRPCService_View(t *testing.T) {
	type fields struct {
		Config                 map[string]interface{}
		GetNervaStore          func(database string) *nt.NervaStore
		GetTokenKeys           func() map[string]map[string]string
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx context.Context
		req *pb.RequestView
	}
	config := nt.IM{
		"version":              testData.version,
		"NT_HASHTABLE":         testData.hashTable,
		"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "view_ok",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.adminToken)),
				req: &pb.RequestView{
					Options: []*pb.RequestView_Query{
						{
							Key:    "customers",
							Text:   "select c.id, ct.groupvalue as custtype, c.custnumber, c.custname from customer c inner join groups ct on c.custtype = ct.id where c.deleted = 0 and c.custnumber <> 'HOME'",
							Values: []*pb.Value{},
						},
						{
							Key: "invoices",
							Text: `select t.id, t.transnumber, tt.groupvalue as transtype, td.groupvalue as direction, t.transdate, c.custname, t.curr, items.amount
										from trans t inner join groups tt on t.transtype = tt.id
										inner join groups td on t.direction = td.id
										inner join customer c on t.customer_id = c.id
										inner join (
											select trans_id, sum(amount) amount
											from item where deleted = 0 group by trans_id) items on t.id = items.trans_id
										where t.deleted = 0 and tt.groupvalue = 'invoice'`,
							Values: []*pb.Value{},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "view_scope_error",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.customerToken)),
				req: &pb.RequestView{
					Options: []*pb.RequestView_Query{
						{
							Key:    "customers",
							Text:   "select c.id, ct.groupvalue as custtype, c.custnumber, c.custname from customer c inner join groups ct on c.custtype = ct.id where c.deleted = 0 and c.custnumber <> 'HOME'",
							Values: []*pb.Value{},
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "view_error",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.adminToken)),
				req: &pb.RequestView{
					Options: []*pb.RequestView_Query{
						{
							Key:    "customers",
							Text:   "select c.id, ct.groupvalue as custtype, c.custnumber, ",
							Values: []*pb.Value{{Value: &pb.Value_Text{Text: "param"}}},
						},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &RPCService{
				Config:                 tt.fields.Config,
				GetNervaStore:          tt.fields.GetNervaStore,
				GetTokenKeys:           tt.fields.GetTokenKeys,
				UnimplementedAPIServer: tt.fields.UnimplementedAPIServer,
			}
			if _, err := srv.View(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("RPCService.View() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRPCService_Function(t *testing.T) {
	type fields struct {
		Config                 map[string]interface{}
		GetNervaStore          func(database string) *nt.NervaStore
		GetTokenKeys           func() map[string]map[string]string
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx context.Context
		req *pb.RequestFunction
	}
	config := nt.IM{
		"version":              testData.version,
		"NT_HASHTABLE":         testData.hashTable,
		"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "function_ok",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.adminToken)),
				req: &pb.RequestFunction{
					Key: "nextNumber",
					Values: map[string]*pb.Value{
						"numberkey": {Value: &pb.Value_Text{Text: "custnumber"}},
						"step":      {Value: &pb.Value_Boolean{Boolean: false}},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "function_scope_error",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.customerToken)),
				req: &pb.RequestFunction{
					Key: "nextNumber",
					Values: map[string]*pb.Value{
						"numberkey": {Value: &pb.Value_Text{Text: "custnumber"}},
						"step":      {Value: &pb.Value_Boolean{Boolean: false}},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "function_error",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.adminToken)),
				req: &pb.RequestFunction{
					Key:    "number",
					Values: map[string]*pb.Value{},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &RPCService{
				Config:                 tt.fields.Config,
				GetNervaStore:          tt.fields.GetNervaStore,
				GetTokenKeys:           tt.fields.GetTokenKeys,
				UnimplementedAPIServer: tt.fields.UnimplementedAPIServer,
			}
			if _, err := srv.Function(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("RPCService.Function() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRPCService_ReportList(t *testing.T) {
	type fields struct {
		Config                 map[string]interface{}
		GetNervaStore          func(database string) *nt.NervaStore
		GetTokenKeys           func() map[string]map[string]string
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx context.Context
		req *pb.RequestReportList
	}
	config := nt.IM{
		"version":              testData.version,
		"NT_HASHTABLE":         testData.hashTable,
		"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
		"NT_REPORT_DIR":        "",
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "list_ok",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.adminToken)),
				req: &pb.RequestReportList{
					Label: "invoice",
				},
			},
			wantErr: false,
		},
		{
			name: "list_scope_error",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.userToken)),
				req: &pb.RequestReportList{
					Label: "invoice",
				},
			},
			wantErr: true,
		},
		{
			name: "list_error",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, nt.IM{
						"version":              testData.version,
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
						"NT_REPORT_DIR":        "bad",
					}), testData.adminToken)),
				req: &pb.RequestReportList{
					Label: "invoice",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &RPCService{
				Config:                 tt.fields.Config,
				GetNervaStore:          tt.fields.GetNervaStore,
				GetTokenKeys:           tt.fields.GetTokenKeys,
				UnimplementedAPIServer: tt.fields.UnimplementedAPIServer,
			}
			if _, err := srv.ReportList(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("RPCService.ReportList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRPCService_ReportDelete(t *testing.T) {
	type fields struct {
		Config                 map[string]interface{}
		GetNervaStore          func(database string) *nt.NervaStore
		GetTokenKeys           func() map[string]map[string]string
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx context.Context
		req *pb.RequestReportDelete
	}
	config := nt.IM{
		"version":              testData.version,
		"NT_HASHTABLE":         testData.hashTable,
		"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
		"NT_REPORT_DIR":        "",
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "delete_ok",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.adminToken)),
				req: &pb.RequestReportDelete{
					Reportkey: "ntr_cash_in_en",
				},
			},
			wantErr: false,
		},
		{
			name: "delete_error",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.adminToken)),
				req: &pb.RequestReportDelete{
					Reportkey: "ntr_cash_in_en",
				},
			},
			wantErr: true,
		},
		{
			name: "delete_scope_error",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.userToken)),
				req: &pb.RequestReportDelete{
					Reportkey: "ntr_cash_in_en",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &RPCService{
				Config:                 tt.fields.Config,
				GetNervaStore:          tt.fields.GetNervaStore,
				GetTokenKeys:           tt.fields.GetTokenKeys,
				UnimplementedAPIServer: tt.fields.UnimplementedAPIServer,
			}
			if _, err := srv.ReportDelete(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("RPCService.ReportDelete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRPCService_ReportInstall(t *testing.T) {
	type fields struct {
		Config                 map[string]interface{}
		GetNervaStore          func(database string) *nt.NervaStore
		GetTokenKeys           func() map[string]map[string]string
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx context.Context
		req *pb.RequestReportInstall
	}
	config := nt.IM{
		"version":              testData.version,
		"NT_HASHTABLE":         testData.hashTable,
		"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
		"NT_REPORT_DIR":        "",
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "in_ok",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.adminToken)),
				req: &pb.RequestReportInstall{
					Reportkey: "ntr_cash_in_en",
				},
			},
			wantErr: false,
		},
		{
			name: "in_error",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.adminToken)),
				req: &pb.RequestReportInstall{
					Reportkey: "ntr_cash_in_en",
				},
			},
			wantErr: true,
		},
		{
			name: "in_scope_error",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.userToken)),
				req: &pb.RequestReportInstall{
					Reportkey: "ntr_cash_in_en",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &RPCService{
				Config:                 tt.fields.Config,
				GetNervaStore:          tt.fields.GetNervaStore,
				GetTokenKeys:           tt.fields.GetTokenKeys,
				UnimplementedAPIServer: tt.fields.UnimplementedAPIServer,
			}
			if _, err := srv.ReportInstall(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("RPCService.ReportInstall() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRPCService_Report(t *testing.T) {
	type fields struct {
		Config                 map[string]interface{}
		GetNervaStore          func(database string) *nt.NervaStore
		GetTokenKeys           func() map[string]map[string]string
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx context.Context
		req *pb.RequestReport
	}
	config := nt.IM{
		"version":              testData.version,
		"NT_HASHTABLE":         testData.hashTable,
		"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
		"NT_REPORT_DIR":        "",
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "report_ok",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.adminToken)),
				req: &pb.RequestReport{
					Reportkey:   "ntr_invoice_en",
					Orientation: pb.ReportOrientation_portrait,
					Size:        pb.ReportSize_a4,
					Type:        pb.ReportType_report_trans,
					Refnumber:   "DMINV/00001",
				},
			},
			wantErr: false,
		},
		{
			name: "report_error",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(testData.driver, config), testData.adminToken)),
				req: &pb.RequestReport{
					Reportkey:   "ntr_invoice",
					Orientation: pb.ReportOrientation_portrait,
					Size:        pb.ReportSize_a4,
					Type:        pb.ReportType_report_trans,
					Refnumber:   "DMINV/00001",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &RPCService{
				Config:                 tt.fields.Config,
				GetNervaStore:          tt.fields.GetNervaStore,
				GetTokenKeys:           tt.fields.GetTokenKeys,
				UnimplementedAPIServer: tt.fields.UnimplementedAPIServer,
			}
			if _, err := srv.Report(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("RPCService.Report() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRPCService_DatabaseCreate(t *testing.T) {
	type fields struct {
		Config                 map[string]interface{}
		GetNervaStore          func(database string) *nt.NervaStore
		GetTokenKeys           func() map[string]map[string]string
		UnimplementedAPIServer pb.UnimplementedAPIServer
	}
	type args struct {
		ctx context.Context
		req *pb.RequestDatabaseCreate
	}
	config := nt.IM{
		"version":       testData.version,
		"NT_HASHTABLE":  testData.hashTable,
		"NT_API_KEY":    testData.apiKey,
		"NT_ALIAS_TEST": testData.testDatabase,
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "database_ok",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(&db.SQLDriver{
						Config: nt.IM{"version": testData.version}}, config), testData.adminToken)),
				req: &pb.RequestDatabaseCreate{
					Alias: "test", Demo: true,
				},
			},
			wantErr: false,
		},
		{
			name: "database_error",
			args: args{
				ctx: context.WithValue(context.Background(), NstoreCtxKey,
					testData.getApi(nt.New(&db.SQLDriver{
						Config: nt.IM{"version": testData.version}}, config), testData.adminToken)),
				req: &pb.RequestDatabaseCreate{
					Alias: "kalevala", Demo: true,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &RPCService{
				Config:                 tt.fields.Config,
				GetNervaStore:          tt.fields.GetNervaStore,
				GetTokenKeys:           tt.fields.GetTokenKeys,
				UnimplementedAPIServer: tt.fields.UnimplementedAPIServer,
			}
			if _, err := srv.DatabaseCreate(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("RPCService.DatabaseCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
