package service

import (
	"strings"
	"testing"

	db "github.com/nervatura/nervatura/service/pkg/database"
	nt "github.com/nervatura/nervatura/service/pkg/nervatura"
)

func TestCLIService_TokenDecode(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
	}
	type args struct {
		token string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "admin_token",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
			},
			args: args{
				token: testData.adminToken,
			},
			want: `{"database":"test","exp":`,
		},
		{
			name: "error_token",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
			},
			args: args{
				token: "ASND1233kjkjdjk2222jkjkd",
			},
			want: `{"code":400,"message":`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &CLIService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
			}
			if got := srv.TokenDecode(tt.args.token); !strings.HasPrefix(got, tt.want) {
				t.Errorf("CLIService.TokenDecode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_TokenLogin(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
	}
	type args struct {
		token     string
		tokenKeys map[string]map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "admin_token",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
			},
			args: args{
				token:     testData.adminToken,
				tokenKeys: make(map[string]map[string]string),
			},
			want: `{"id":1,"username":"admin","empnumber":"admin","usergroup":102,"scope":"admin"}`,
		},
		{
			name: "unauthorized_token",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
			},
			args: args{
				token:     "ASND1233kjkjdjk2222jkjkd",
				tokenKeys: make(map[string]map[string]string),
			},
			want: `{"code":401,"message":"Unauthorized"}`,
		},
		{
			name: "invalid_token",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
			},
			args: args{
				token:     "eyJhbGciOiJIUzI1NiIsImtpZCI6IjhiZDBlMDI1MDk0ODJmNThjZWZkM2MwZWNkNDFmZjBlIiwidHlwIjoiSldUIn0.eyJ1c2VybmFtZSI6ImFkbWluIiwiZGF0YWJhc2UiOiJkZW1vIiwiZXhwIjoxNjI4NTM1ODQzLCJpc3MiOiJuZXJ2YXR1cmEifQ.ErrxPYNENXc7tvi7PVzE8z6qe8QEEtnEgdcPEqzAvos",
				tokenKeys: make(map[string]map[string]string),
			},
			want: `{"code":401,"message":"signature is invalid"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &CLIService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
			}
			if _, got := srv.TokenLogin(tt.args.token, tt.args.tokenKeys); !strings.HasPrefix(got, tt.want) {
				t.Errorf("CLIService.TokenLogin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_UserLogin(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
	}
	type args struct {
		options nt.IM
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "login_ok",
			fields: fields{
				Config: nt.IM{
					"version": testData.version,
				},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
			},
			args: args{
				options: nt.IM{
					"username": "admin",
					"database": "test"},
			},
			want: `{"engine":"`,
		},
		{
			name: "login_missing_database",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
			},
			args: args{
				options: nt.IM{
					"username": "admin"},
			},
			want: `{"code":400,"message":"Missing database"}`,
		},
		{
			name: "login_wrong_password",
			fields: fields{
				Config: nt.IM{
					"version": testData.version,
				},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
			},
			args: args{
				options: nt.IM{
					"username": "admin",
					"password": "123",
					"database": "test"},
			},
			want: `{"code":400,"message":"Incorrect password"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &CLIService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
			}
			if got := srv.UserLogin(tt.args.options); !strings.HasPrefix(got, tt.want) {
				t.Errorf("CLIService.UserLogin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_UserPassword(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
	}
	type args struct {
		api     *nt.API
		options nt.IM
	}
	config := nt.IM{
		"version":              testData.version,
		"NT_HASHTABLE":         testData.hashTable,
		"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "admin_ok",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
			},
			args: args{
				api: testData.getApi(nt.New(testData.driver, config), testData.adminToken),
				options: nt.IM{
					"username": "guest",
					"password": "123",
					"confirm":  "123"},
			},
			want: `{"code":204,"message":"OK"}`,
		},
		{
			name: "user_ok",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
			},
			args: args{
				api: testData.getApi(nt.New(testData.driver, config), testData.userToken),
				options: nt.IM{
					"password": "123",
					"confirm":  "123"},
			},
			want: `{"code":204,"message":"OK"}`,
		},
		{
			name: "user_scope_error",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
			},
			args: args{
				api: testData.getApi(nt.New(testData.driver, config), testData.userToken),
				options: nt.IM{
					"username": "guest",
					"password": "123",
					"confirm":  "123"},
			},
			want: `{"code":401,"message":"Unauthorized"}`,
		},
		{
			name: "customer_scope_error",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
			},
			args: args{
				api: testData.getApi(nt.New(testData.driver, config), testData.customerToken),
				options: nt.IM{
					"password": "123",
					"confirm":  "123"},
			},
			want: `{"code":401,"message":"Unauthorized"}`,
		},
		{
			name: "user_empty_error",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
			},
			args: args{
				api: testData.getApi(nt.New(testData.driver, config), testData.userToken),
				options: nt.IM{
					"password": "",
					"confirm":  ""},
			},
			want: `{"code":400,"message":"The new password can not be empty!"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &CLIService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
			}
			if got := srv.UserPassword(tt.args.api, tt.args.options); got != tt.want {
				t.Errorf("CLIService.UserPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_TokenRefresh(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
	}
	type args struct {
		api *nt.API
	}
	config := nt.IM{
		"version":              testData.version,
		"NT_HASHTABLE":         testData.hashTable,
		"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "refresh_ok",
			fields: fields{
				Config: nt.IM{},
			},
			args: args{
				api: testData.getApi(nt.New(testData.driver, config), testData.userToken),
			},
			want: `{"token":`,
		},
		{
			name: "api_nil",
			want: `{"code":401,"message":"Unauthorized"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &CLIService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
			}
			if got := srv.TokenRefresh(tt.args.api); !strings.HasPrefix(got, tt.want) {
				t.Errorf("CLIService.TokenRefresh() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_Get(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
	}
	type args struct {
		api     *nt.API
		options nt.IM
	}
	config := nt.IM{
		"version":              testData.version,
		"NT_HASHTABLE":         testData.hashTable,
		"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "get_ok",
			fields: fields{
				Config: config,
			},
			args: args{
				api: testData.getApi(nt.New(testData.driver, config), testData.adminToken),
				options: nt.IM{
					"nervatype": "customer",
					"metadata":  true,
					"filter":    "custnumber;==;DMCUST/00001"},
			},
			want: `[{"account":null,"creditlimit":1000000,"custname":"First Customer Co.",`,
		},
		{
			name: "get_error",
			fields: fields{
				Config: config,
			},
			args: args{
				api: testData.getApi(nt.New(testData.driver, config), testData.adminToken),
				options: nt.IM{
					"nervatype": "kalevala",
					"metadata":  true,
					"filter":    "custnumber;==;DMCUST/00001"},
			},
			want: `{"code":400,"message":"Invalid nervatype value: kalevala"}`,
		},
		{
			name: "get_scope_error",
			fields: fields{
				Config: config,
			},
			args: args{
				api: testData.getApi(nt.New(testData.driver, config), testData.customerToken),
				options: nt.IM{
					"nervatype": "customer",
					"metadata":  true,
					"filter":    "custnumber;==;DMCUST/00001"},
			},
			want: `{"code":401,"message":"Unauthorized"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &CLIService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
			}
			if got := srv.Get(tt.args.api, tt.args.options); !strings.HasPrefix(got, tt.want) {
				t.Errorf("CLIService.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_View(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
	}
	type args struct {
		api  *nt.API
		data []nt.IM
	}
	config := nt.IM{
		"version":              testData.version,
		"NT_HASHTABLE":         testData.hashTable,
		"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "view_ok",
			fields: fields{
				Config: config,
			},
			args: args{
				api: testData.getApi(nt.New(testData.driver, config), testData.adminToken),
				data: []nt.IM{
					{
						"key":    "customers",
						"text":   "select c.id, ct.groupvalue as custtype, c.custnumber, c.custname from customer c inner join groups ct on c.custtype = ct.id where c.deleted = 0 and c.custnumber <> 'HOME'",
						"values": []interface{}{},
					},
					{
						"key":    "invoices",
						"text":   "select t.id, t.transnumber, tt.groupvalue as transtype, td.groupvalue as direction, t.transdate, c.custname, t.curr, items.amount from trans t inner join groups tt on t.transtype = tt.id inner join groups td on t.direction = td.id inner join customer c on t.customer_id = c.id inner join ( select trans_id, sum(amount) amount from item where deleted = 0 group by trans_id) items on t.id = items.trans_id where t.deleted = 0 and tt.groupvalue = 'invoice'",
						"values": []interface{}{},
					}},
			},
			want: `{"customers":[{"custname":"First Customer Co.","custnumber":"DMCUST/00001",`,
		},
		{
			name: "view_error",
			fields: fields{
				Config: config,
			},
			args: args{
				api: testData.getApi(nt.New(testData.driver, config), testData.adminToken),
				data: []nt.IM{
					{
						"key":    "customers",
						"text":   "select c.id, ct.groupvalue as custtype",
						"values": []interface{}{},
					}},
			},
			want: `{"code":400,"message":`,
		},
		{
			name: "view_scope_error",
			fields: fields{
				Config: config,
			},
			args: args{
				api: testData.getApi(nt.New(testData.driver, config), testData.customerToken),
				data: []nt.IM{
					{
						"key":    "customers",
						"text":   "select c.id, ct.groupvalue as custtype, c.custnumber, c.custname from customer c inner join groups ct on c.custtype = ct.id where c.deleted = 0 and c.custnumber <> 'HOME'",
						"values": []interface{}{},
					},
					{
						"key":    "invoices",
						"text":   "select t.id, t.transnumber, tt.groupvalue as transtype, td.groupvalue as direction, t.transdate, c.custname, t.curr, items.amount from trans t inner join groups tt on t.transtype = tt.id inner join groups td on t.direction = td.id inner join customer c on t.customer_id = c.id inner join ( select trans_id, sum(amount) amount from item where deleted = 0 group by trans_id) items on t.id = items.trans_id where t.deleted = 0 and tt.groupvalue = 'invoice'",
						"values": []interface{}{},
					}},
			},
			want: `{"code":401,"message":"Unauthorized"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &CLIService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
			}
			if got := srv.View(tt.args.api, tt.args.data); !strings.HasPrefix(got, tt.want) {
				t.Errorf("CLIService.View() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_Function(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
	}
	type args struct {
		api     *nt.API
		options nt.IM
	}
	config := nt.IM{
		"version":              testData.version,
		"NT_HASHTABLE":         testData.hashTable,
		"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "func_ok",
			fields: fields{
				Config: config,
			},
			args: args{
				api: testData.getApi(nt.New(testData.driver, config), testData.adminToken),
				options: nt.IM{
					"key": "nextNumber",
					"values": nt.IM{
						"numberkey": "custnumber",
						"step":      false,
					}},
			},
			want: `"CUS/00001"`,
		},
		{
			name: "func_error",
			fields: fields{
				Config: config,
			},
			args: args{
				api: testData.getApi(nt.New(testData.driver, config), testData.adminToken),
				options: nt.IM{
					"key":    "number",
					"values": nt.IM{}},
			},
			want: `{"code":400,"message":"Unknown method: number"}`,
		},
		{
			name: "func_scope_error",
			fields: fields{
				Config: config,
			},
			args: args{
				api: testData.getApi(nt.New(testData.driver, config), testData.customerToken),
				options: nt.IM{
					"key": "nextNumber",
					"values": nt.IM{
						"numberkey": "custnumber",
						"step":      false,
					}},
			},
			want: `{"code":401,"message":"Unauthorized"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &CLIService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
			}
			if got := srv.Function(tt.args.api, tt.args.options); !strings.HasPrefix(got, tt.want) {
				t.Errorf("CLIService.Function() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_Update(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
	}
	type args struct {
		api       *nt.API
		nervatype string
		data      []nt.IM
	}
	config := nt.IM{
		"version":              testData.version,
		"NT_HASHTABLE":         testData.hashTable,
		"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "update_ok",
			fields: fields{
				Config: config,
			},
			args: args{
				api:       testData.getApi(nt.New(testData.driver, config), testData.adminToken),
				nervatype: "address",
				data: []nt.IM{
					{
						"keys": nt.IM{
							"id": "customer/DMCUST/00001~1"},
						"zipcode": "54321",
						"city":    "BigCity",
						"notes":   "Update an item by Keys"}},
			},
			want: `[3]`,
		},
		{
			name: "update_error",
			fields: fields{
				Config: config,
			},
			args: args{
				api:       testData.getApi(nt.New(testData.driver, config), testData.adminToken),
				nervatype: "customer",
				data: []nt.IM{
					{
						"keys": nt.IM{
							"nervatype": "customer",
							"ref_id":    "customer/00001"},
						"zipcode": "12345",
						"city":    "BigCity",
					}},
			},
			want: `{"code":400,"message":"Invalid`,
		},
		{
			name: "update_scope_error",
			fields: fields{
				Config: config,
			},
			args: args{
				api:       testData.getApi(nt.New(testData.driver, config), testData.customerToken),
				nervatype: "address",
				data: []nt.IM{
					{
						"keys": nt.IM{
							"id": "customer/DMCUST/00001~1"},
						"zipcode": "54321",
						"city":    "BigCity",
						"notes":   "Update an item by Keys"}},
			},
			want: `{"code":401,"message":"Unauthorized"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &CLIService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
			}
			if got := srv.Update(tt.args.api, tt.args.nervatype, tt.args.data); !strings.HasPrefix(got, tt.want) {
				t.Errorf("CLIService.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_Delete(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
	}
	type args struct {
		api     *nt.API
		options nt.IM
	}
	config := nt.IM{
		"version":              testData.version,
		"NT_HASHTABLE":         testData.hashTable,
		"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "delete_ok",
			fields: fields{
				Config: config,
			},
			args: args{
				api: testData.getApi(nt.New(testData.driver, config), testData.adminToken),
				options: nt.IM{
					"nervatype": "address",
					"ref_id":    55},
			},
			want: `{"code":204,"message":"OK"}`,
		},
		{
			name: "delete_error",
			fields: fields{
				Config: config,
			},
			args: args{
				api: testData.getApi(nt.New(testData.driver, config), testData.adminToken),
				options: nt.IM{
					"nervatype": "cust",
					"ref_id":    55},
			},
			want: `{"code":400,"message":"Invalid nervatype value: cust"}`,
		},
		{
			name: "delete_scope_error",
			fields: fields{
				Config: config,
			},
			args: args{
				api: testData.getApi(nt.New(testData.driver, config), testData.customerToken),
				options: nt.IM{
					"nervatype": "address",
					"ref_id":    55},
			},
			want: `{"code":401,"message":"Unauthorized"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &CLIService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
			}
			if got := srv.Delete(tt.args.api, tt.args.options); !strings.HasPrefix(got, tt.want) {
				t.Errorf("CLIService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_ReportList(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
	}
	type args struct {
		api     *nt.API
		options nt.IM
	}
	config := nt.IM{
		"version":              testData.version,
		"NT_HASHTABLE":         testData.hashTable,
		"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
		"NT_REPORT_DIR":        "",
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "list_ok",
			fields: fields{
				Config: config,
			},
			args: args{
				api: testData.getApi(nt.New(testData.driver, config), testData.adminToken),
				options: nt.IM{
					"label": "invoice"},
			},
			want: `[{"description":"Customer invoice","filename":"ntr_invoice_en.json",`,
		},
		{
			name: "list_scope_error",
			fields: fields{
				Config: config,
			},
			args: args{
				api: testData.getApi(nt.New(testData.driver, config), testData.userToken),
				options: nt.IM{
					"label": "invoice"},
			},
			want: `{"code":401,"message":"Unauthorized"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &CLIService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
			}
			if got := srv.ReportList(tt.args.api, tt.args.options); !strings.HasPrefix(got, tt.want) {
				t.Errorf("CLIService.ReportList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_ReportDelete(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
	}
	type args struct {
		api     *nt.API
		options nt.IM
	}
	config := nt.IM{
		"version":              testData.version,
		"NT_HASHTABLE":         testData.hashTable,
		"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
		"NT_REPORT_DIR":        "",
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "delete_ok",
			fields: fields{
				Config: config,
			},
			args: args{
				api: testData.getApi(nt.New(testData.driver, config), testData.adminToken),
				options: nt.IM{
					"reportkey": "ntr_order_in_en"},
			},
			want: `{"code":204,"message":"OK"}`,
		},
		{
			name: "delete_error",
			fields: fields{
				Config: config,
			},
			args: args{
				api: testData.getApi(nt.New(testData.driver, config), testData.adminToken),
				options: nt.IM{
					"reportkey": "ntr_order_in_en"},
			},
			want: `{"code":400,"message":"Missing reportkey: ntr_order_in_en"}`,
		},
		{
			name: "delete_scope_error",
			fields: fields{
				Config: config,
			},
			args: args{
				api: testData.getApi(nt.New(testData.driver, config), testData.userToken),
				options: nt.IM{
					"reportkey": "ntr_order_in_en"},
			},
			want: `{"code":401,"message":"Unauthorized"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &CLIService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
			}
			if got := srv.ReportDelete(tt.args.api, tt.args.options); !strings.HasPrefix(got, tt.want) {
				t.Errorf("CLIService.ReportDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_ReportInstall(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
	}
	type args struct {
		api     *nt.API
		options nt.IM
	}
	config := nt.IM{
		"version":              testData.version,
		"NT_HASHTABLE":         testData.hashTable,
		"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
		"NT_REPORT_DIR":        "",
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "report_ok",
			fields: fields{
				Config: config,
			},
			args: args{
				api: testData.getApi(nt.New(testData.driver, config), testData.adminToken),
				options: nt.IM{
					"reportkey": "ntr_order_in_en"},
			},
			want: `3`,
		},
		{
			name: "report_error",
			fields: fields{
				Config: config,
			},
			args: args{
				api: testData.getApi(nt.New(testData.driver, config), testData.adminToken),
				options: nt.IM{
					"reportkey": "ntr_order_en"},
			},
			want: `{"code":400,"message":`,
		},
		{
			name: "report_scope_error",
			fields: fields{
				Config: config,
			},
			args: args{
				api: testData.getApi(nt.New(testData.driver, config), testData.userToken),
				options: nt.IM{
					"reportkey": "ntr_order_in_en"},
			},
			want: `{"code":401,"message":"Unauthorized"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &CLIService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
			}
			if got := srv.ReportInstall(tt.args.api, tt.args.options); !strings.HasPrefix(got, tt.want) {
				t.Errorf("CLIService.ReportInstall() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_Report(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
	}
	type args struct {
		api     *nt.API
		options nt.IM
	}
	config := nt.IM{
		"version":              testData.version,
		"NT_HASHTABLE":         testData.hashTable,
		"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
		"NT_REPORT_DIR":        "",
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "report_ok",
			fields: fields{
				Config: config,
			},
			args: args{
				api: testData.getApi(nt.New(testData.driver, config), testData.adminToken),
				options: nt.IM{
					"reportkey":   "ntr_invoice_en",
					"orientation": "portrait",
					"size":        "a4",
					"nervatype":   "trans",
					"refnumber":   "DMINV/00001"},
			},
			want: `{"data":null,"filetype":"base64","template":"data:application/pdf;filename=Report.pdf;base64,`,
		},
		{
			name: "report_error",
			fields: fields{
				Config: config,
			},
			args: args{
				api: testData.getApi(nt.New(testData.driver, config), testData.adminToken),
				options: nt.IM{
					"reportkey":   "invoice",
					"orientation": "portrait",
					"size":        "a4",
					"nervatype":   "trans",
					"refnumber":   "DMINV/00001"},
			},
			want: `{"code":400,"message":"does not exist"}`,
		},
		{
			name: "api_nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &CLIService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
			}
			if got := srv.Report(tt.args.api, tt.args.options); !strings.HasPrefix(got, tt.want) {
				t.Errorf("CLIService.Report() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_DatabaseCreate(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
	}
	type args struct {
		apiKey  string
		options nt.IM
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "create_ok",
			fields: fields{
				Config: nt.IM{
					"NT_API_KEY": testData.apiKey,
				},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(&db.SQLDriver{
						Config: nt.IM{"version": testData.version}}, nt.IM{
						"version":       testData.version,
						"NT_HASHTABLE":  testData.hashTable,
						"NT_API_KEY":    testData.apiKey,
						"NT_ALIAS_TEST": testData.testDatabase,
					})
				},
			},
			args: args{
				apiKey: testData.apiKey,
				options: nt.IM{
					"database": "test",
					"demo":     true},
			},
			want: `[{"database":"test","message":"Start process",`,
		},
		{
			name: "create_error",
			fields: fields{
				Config: nt.IM{
					"NT_API_KEY": testData.apiKey,
				},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(&db.SQLDriver{
						Config: nt.IM{"version": testData.version}}, nt.IM{
						"version":       testData.version,
						"NT_HASHTABLE":  testData.hashTable,
						"NT_API_KEY":    testData.apiKey,
						"NT_ALIAS_TEST": testData.testDatabase,
					})
				},
			},
			args: args{
				apiKey: testData.apiKey,
				options: nt.IM{
					"database": "kalevala",
					"demo":     true},
			},
			want: `{"code":400,"message":"Could not connect to the database"}`,
		},
		{
			name: "database_unauthorized",
			fields: fields{
				Config: nt.IM{
					"NT_API_KEY": testData.apiKey,
				},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(&db.SQLDriver{
						Config: nt.IM{"version": testData.version}}, nt.IM{
						"version":       testData.version,
						"NT_HASHTABLE":  testData.hashTable,
						"NT_API_KEY":    testData.apiKey,
						"NT_ALIAS_TEST": testData.testDatabase,
					})
				},
			},
			args: args{
				apiKey: "API_KEY",
				options: nt.IM{
					"database": "test",
					"demo":     true},
			},
			want: `{"code":401,"message":"Unauthorized"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &CLIService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
			}
			if got := srv.DatabaseCreate(tt.args.apiKey, tt.args.options); !strings.HasPrefix(got, tt.want) {
				t.Errorf("CLIService.DatabaseCreate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCLIService_checkUser(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
	}
	type args struct {
		api   *nt.API
		admin bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "api_nil",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &CLIService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
			}
			if err := srv.checkUser(tt.args.api, tt.args.admin); (err != nil) != tt.wantErr {
				t.Errorf("CLIService.checkUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
