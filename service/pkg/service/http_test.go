//go:build http || all
// +build http all

package service

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	db "github.com/nervatura/nervatura/service/pkg/database"
	nt "github.com/nervatura/nervatura/service/pkg/nervatura"
)

func TestHTTPService_ClientConfig(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
		GetParam      func(req *http.Request, name string) string
		GetTokenKeys  func() map[string]map[string]string
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		result int
	}{
		{
			name: "empty",
			fields: fields{
				Config: nt.IM{
					"version":          testData.version,
					"NT_CLIENT_CONFIG": "",
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/config", nil),
			},
			result: http.StatusOK,
		},
		{
			name: "config_file",
			fields: fields{
				Config: nt.IM{
					"version":          testData.version,
					"NT_CLIENT_CONFIG": "../../data/test_client_config.json",
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/config", nil),
			},
			result: http.StatusOK,
		},
		{
			name: "config_file_error",
			fields: fields{
				Config: nt.IM{
					"version":          testData.version,
					"NT_CLIENT_CONFIG": "data/config.json",
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/config", nil),
			},
			result: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &HTTPService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
				GetParam:      tt.fields.GetParam,
				GetTokenKeys:  tt.fields.GetTokenKeys,
			}
			srv.ClientConfig(tt.args.w, tt.args.r)
			if status := tt.args.w.(*httptest.ResponseRecorder).Code; status != tt.result {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.result)
			}
		})
	}
}

func TestHTTPService_UserLogin(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
		GetParam      func(req *http.Request, name string) string
		GetTokenKeys  func() map[string]map[string]string
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		result int
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
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/login",
					bytes.NewBuffer([]byte(testData.encodeData(nt.IM{
						"username": "admin",
						"database": "test",
					})))),
			},
			result: http.StatusOK,
		},
		{
			name: "login_missing_database",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/login",
					bytes.NewBuffer([]byte(testData.encodeData(nt.IM{
						"username": "admin",
					})))),
			},
			result: http.StatusBadRequest,
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
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/login",
					bytes.NewBuffer([]byte(testData.encodeData(nt.IM{
						"username": "admin",
						"password": "1a1a1a",
						"database": "test"})))),
			},
			result: http.StatusBadRequest,
		},
		{
			name: "login_body_error",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/login",
					bytes.NewBuffer([]byte("{error"))),
			},
			result: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &HTTPService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
				GetParam:      tt.fields.GetParam,
				GetTokenKeys:  tt.fields.GetTokenKeys,
			}
			srv.UserLogin(tt.args.w, tt.args.r)
			if status := tt.args.w.(*httptest.ResponseRecorder).Code; status != tt.result {
				t.Errorf("handler returned wrong status code(%s): got %v want %v",
					tt.name, status, tt.result)
			}
		})
	}
}

func TestHTTPService_TokenLogin(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
		GetParam      func(req *http.Request, name string) string
		GetTokenKeys  func() map[string]map[string]string
	}
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		token   string
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
				r: httptest.NewRequest("POST", "/", nil),
			},
			token: testData.adminToken,
		},
		{
			name: "token_missing",
			args: args{
				r: httptest.NewRequest("POST", "/", nil),
			},
			token:   "",
			wantErr: true,
		},
		{
			name: "token_wrong",
			args: args{
				r: httptest.NewRequest("POST", "/", nil),
			},
			token:   "010101010101",
			wantErr: true,
		},
		{
			name: "token_expired",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				r: httptest.NewRequest("POST", "/", nil),
			},
			token:   "eyJhbGciOiJIUzI1NiIsImtpZCI6IjhiZDBlMDI1MDk0ODJmNThjZWZkM2MwZWNkNDFmZjBlIiwidHlwIjoiSldUIn0.eyJ1c2VybmFtZSI6ImFkbWluIiwiZGF0YWJhc2UiOiJkZW1vIiwiZXhwIjoxNjI4NTM1ODQzLCJpc3MiOiJuZXJ2YXR1cmEifQ.ErrxPYNENXc7tvi7PVzE8z6qe8QEEtnEgdcPEqzAvos",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &HTTPService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
				GetParam:      tt.fields.GetParam,
				GetTokenKeys:  tt.fields.GetTokenKeys,
			}
			tt.args.r.Header.Set("Authorization", "Bearer "+tt.token)
			_, err := srv.TokenLogin(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPService.TokenLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestHTTPService_UserPassword(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
		GetParam      func(req *http.Request, name string) string
		GetTokenKeys  func() map[string]map[string]string
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		token  string
		result int
	}{
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
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/password",
					bytes.NewBuffer([]byte(testData.encodeData(nt.IM{
						"username": "guest",
						"password": "123",
						"confirm":  "123"})))),
			},
			token:  testData.adminToken,
			result: http.StatusNoContent,
		},
		{
			name: "custnumber_ok",
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
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/password",
					bytes.NewBuffer([]byte(testData.encodeData(nt.IM{
						"custnumber": "DMCUST/00001",
						"password":   "123",
						"confirm":    "123"})))),
			},
			token:  testData.adminToken,
			result: http.StatusNoContent,
		},
		{
			name: "customer_token_error",
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
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/password",
					bytes.NewBuffer([]byte(testData.encodeData(nt.IM{
						"password": "123",
						"confirm":  "123"})))),
			},
			token:  testData.customerToken,
			result: http.StatusUnauthorized,
		},
		{
			name: "empty_error",
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
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/password",
					bytes.NewBuffer([]byte(testData.encodeData(nt.IM{
						"password": "",
						"confirm":  ""})))),
			},
			token:  testData.adminToken,
			result: http.StatusBadRequest,
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
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/password",
					bytes.NewBuffer([]byte(testData.encodeData(nt.IM{
						"username": "demo",
						"password": "123",
						"confirm":  "123"})))),
			},
			token:  testData.userToken,
			result: http.StatusUnauthorized,
		},
		{
			name: "custnumber_scope_error",
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
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/password",
					bytes.NewBuffer([]byte(testData.encodeData(nt.IM{
						"custnumber": "DMCUST/00001",
						"password":   "123",
						"confirm":    "123"})))),
			},
			token:  testData.userToken,
			result: http.StatusUnauthorized,
		},
		{
			name: "pw_body_error",
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
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/password",
					bytes.NewBuffer([]byte("{error"))),
			},
			token:  testData.adminToken,
			result: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &HTTPService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
				GetParam:      tt.fields.GetParam,
				GetTokenKeys:  tt.fields.GetTokenKeys,
			}
			if tt.name != "pw_unauthorized" {
				tt.args.r.Header.Set("Authorization", "Bearer "+tt.token)
				ctx, err := srv.TokenLogin(tt.args.r)
				if err != nil {
					t.Fatal(err)
				}
				tt.args.r = tt.args.r.WithContext(ctx)
			}
			srv.UserPassword(tt.args.w, tt.args.r)
			if status := tt.args.w.(*httptest.ResponseRecorder).Code; status != tt.result {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.result)
			}
		})
	}
}

func TestHTTPService_TokenValidate(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
		GetParam      func(req *http.Request, name string) string
		GetTokenKeys  func() map[string]map[string]string
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		result int
	}{
		{
			name: "validate_ok",
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
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/refresh", nil),
			},
			result: http.StatusOK,
		},
		{
			name: "validate_unauthorized",
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
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/refresh", nil),
			},
			result: http.StatusUnauthorized,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &HTTPService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
				GetParam:      tt.fields.GetParam,
				GetTokenKeys:  tt.fields.GetTokenKeys,
			}
			if tt.name != "validate_unauthorized" {
				tt.args.r.Header.Set("Authorization", "Bearer "+testData.adminToken)
				ctx, err := srv.TokenLogin(tt.args.r)
				if err != nil {
					t.Fatal(err)
				}
				tt.args.r = tt.args.r.WithContext(ctx)
			}
			srv.TokenValidate(tt.args.w, tt.args.r)
			if status := tt.args.w.(*httptest.ResponseRecorder).Code; status != tt.result {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.result)
			}
		})
	}
}

func TestHTTPService_TokenRefresh(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
		GetParam      func(req *http.Request, name string) string
		GetTokenKeys  func() map[string]map[string]string
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		result int
	}{
		{
			name: "refresh_ok",
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
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/refresh", nil),
			},
			result: http.StatusOK,
		},
		{
			name: "refresh_unauthorized",
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
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/refresh", nil),
			},
			result: http.StatusUnauthorized,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &HTTPService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
				GetParam:      tt.fields.GetParam,
				GetTokenKeys:  tt.fields.GetTokenKeys,
			}
			if tt.name != "refresh_unauthorized" {
				tt.args.r.Header.Set("Authorization", "Bearer "+testData.adminToken)
				ctx, err := srv.TokenLogin(tt.args.r)
				if err != nil {
					t.Fatal(err)
				}
				tt.args.r = tt.args.r.WithContext(ctx)
			}
			srv.TokenRefresh(tt.args.w, tt.args.r)
			if status := tt.args.w.(*httptest.ResponseRecorder).Code; status != tt.result {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.result)
			}
		})
	}
}

func TestHTTPService_GetFilter(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
		GetParam      func(req *http.Request, name string) string
		GetTokenKeys  func() map[string]map[string]string
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		token  string
		result int
	}{
		{
			name: "filter_ok",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetParam: func(req *http.Request, name string) string {
					return "customer"
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/customer?metadata=true&filter=custnumber;==;DMCUST/00001", nil),
			},
			token:  testData.adminToken,
			result: http.StatusOK,
		},
		{
			name: "filter_error",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetParam: func(req *http.Request, name string) string {
					return "customer"
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/customer?metadata=false&filter=credit;==;DMCUST/00001", nil),
			},
			token:  testData.adminToken,
			result: http.StatusBadRequest,
		},
		{
			name: "filter_scope_error",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetParam: func(req *http.Request, name string) string {
					return "customer"
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/customer?metadata=false&filter=credit;==;DMCUST/00001", nil),
			},
			token:  testData.customerToken,
			result: http.StatusUnauthorized,
		},
		{
			name: "filter_unauthorized",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetParam: func(req *http.Request, name string) string {
					return "customer"
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/customer?metadata=false", nil),
			},
			token:  "",
			result: http.StatusUnauthorized,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &HTTPService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
				GetParam:      tt.fields.GetParam,
				GetTokenKeys:  tt.fields.GetTokenKeys,
			}
			if tt.token != "" {
				tt.args.r.Header.Set("Authorization", "Bearer "+tt.token)
				ctx, err := srv.TokenLogin(tt.args.r)
				if err != nil {
					t.Fatal(err)
				}
				tt.args.r = tt.args.r.WithContext(ctx)
			}
			srv.GetFilter(tt.args.w, tt.args.r)
			if status := tt.args.w.(*httptest.ResponseRecorder).Code; status != tt.result {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.result)
			}
		})
	}
}

func TestHTTPService_GetIds(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
		GetParam      func(req *http.Request, name string) string
		GetTokenKeys  func() map[string]map[string]string
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		result int
	}{
		{
			name: "filter_ok",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetParam: func(req *http.Request, name string) string {
					if name == "nervatype" {
						return "customer"
					}
					return "2,4"
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/customer?metadata=true", nil),
			},
			result: http.StatusOK,
		},
		{
			name: "filter_error",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetParam: func(req *http.Request, name string) string {
					if name == "nervatype" {
						return "customer"
					}
					return ""
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/customer?metadata=true", nil),
			},
			result: http.StatusBadRequest,
		},
		{
			name: "filter_unauthorized",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetParam: func(req *http.Request, name string) string {
					if name == "nervatype" {
						return "customer"
					}
					return ""
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/customer?metadata=true", nil),
			},
			result: http.StatusUnauthorized,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &HTTPService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
				GetParam:      tt.fields.GetParam,
				GetTokenKeys:  tt.fields.GetTokenKeys,
			}
			if tt.name != "filter_unauthorized" {
				tt.args.r.Header.Set("Authorization", "Bearer "+testData.adminToken)
				ctx, err := srv.TokenLogin(tt.args.r)
				if err != nil {
					t.Fatal(err)
				}
				tt.args.r = tt.args.r.WithContext(ctx)
			}
			srv.GetIds(tt.args.w, tt.args.r)
			if status := tt.args.w.(*httptest.ResponseRecorder).Code; status != tt.result {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.result)
			}
		})
	}
}

func TestHTTPService_View(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
		GetParam      func(req *http.Request, name string) string
		GetTokenKeys  func() map[string]map[string]string
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		result int
	}{
		{
			name: "view_ok",
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
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/view", bytes.NewBuffer([]byte(testData.encodeData([]nt.IM{
					{
						"key":    "customers",
						"text":   "select c.id, ct.groupvalue as custtype, c.custnumber, c.custname from customer c inner join groups ct on c.custtype = ct.id where c.deleted = 0 and c.custnumber <> 'HOME'",
						"values": []interface{}{},
					},
					{
						"key":    "invoices",
						"text":   "select t.id, t.transnumber, tt.groupvalue as transtype, td.groupvalue as direction, t.transdate, c.custname, t.curr, items.amount from trans t inner join groups tt on t.transtype = tt.id inner join groups td on t.direction = td.id inner join customer c on t.customer_id = c.id inner join ( select trans_id, sum(amount) amount from item where deleted = 0 group by trans_id) items on t.id = items.trans_id where t.deleted = 0 and tt.groupvalue = 'invoice'",
						"values": []interface{}{},
					},
				})))),
			},
			result: http.StatusOK,
		},
		{
			name: "view_error",
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
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/view", bytes.NewBuffer([]byte(testData.encodeData([]nt.IM{
					{
						"key":    "customers",
						"text":   "select c.id, ct.groupvalue as custtype",
						"values": []interface{}{},
					},
				})))),
			},
			result: http.StatusBadRequest,
		},
		{
			name: "view_body_error",
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
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/view", bytes.NewBuffer([]byte("{error"))),
			},
			result: http.StatusBadRequest,
		},
		{
			name: "view_unauthorized",
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
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/view", bytes.NewBuffer([]byte(""))),
			},
			result: http.StatusUnauthorized,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &HTTPService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
				GetParam:      tt.fields.GetParam,
				GetTokenKeys:  tt.fields.GetTokenKeys,
			}
			if tt.name != "view_unauthorized" {
				tt.args.r.Header.Set("Authorization", "Bearer "+testData.adminToken)
				ctx, err := srv.TokenLogin(tt.args.r)
				if err != nil {
					t.Fatal(err)
				}
				tt.args.r = tt.args.r.WithContext(ctx)
			}
			srv.View(tt.args.w, tt.args.r)
			if status := tt.args.w.(*httptest.ResponseRecorder).Code; status != tt.result {
				t.Errorf("handler returned wrong status code(%s): got %v want %v",
					tt.name, status, tt.result)
			}
		})
	}
}

func TestHTTPService_Function(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
		GetParam      func(req *http.Request, name string) string
		GetTokenKeys  func() map[string]map[string]string
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		result int
	}{
		{
			name: "function_ok",
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
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/function", bytes.NewBuffer([]byte(testData.encodeData(nt.IM{
					"key": "nextNumber",
					"values": nt.IM{
						"numberkey": "custnumber",
						"step":      false,
					},
				})))),
			},
			result: http.StatusOK,
		},
		{
			name: "function_error",
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
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/function", bytes.NewBuffer([]byte(testData.encodeData(nt.IM{
					"key":    "number",
					"values": nt.IM{},
				})))),
			},
			result: http.StatusBadRequest,
		},
		{
			name: "function_body_error",
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
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/function", bytes.NewBuffer([]byte("{error"))),
			},
			result: http.StatusBadRequest,
		},
		{
			name: "function_unauthorized",
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
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/function", bytes.NewBuffer([]byte(""))),
			},
			result: http.StatusUnauthorized,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &HTTPService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
				GetParam:      tt.fields.GetParam,
				GetTokenKeys:  tt.fields.GetTokenKeys,
			}
			if tt.name != "function_unauthorized" {
				tt.args.r.Header.Set("Authorization", "Bearer "+testData.adminToken)
				ctx, err := srv.TokenLogin(tt.args.r)
				if err != nil {
					t.Fatal(err)
				}
				tt.args.r = tt.args.r.WithContext(ctx)
			}
			srv.Function(tt.args.w, tt.args.r)
			if status := tt.args.w.(*httptest.ResponseRecorder).Code; status != tt.result {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.result)
			}
		})
	}
}

func TestHTTPService_Update(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
		GetParam      func(req *http.Request, name string) string
		GetTokenKeys  func() map[string]map[string]string
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		result int
	}{
		{
			name: "update_ok",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetParam: func(req *http.Request, name string) string {
					return "address"
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/update", bytes.NewBuffer([]byte(testData.encodeData([]nt.IM{
					{
						"keys": nt.IM{
							"nervatype": "customer",
							"ref_id":    "customer/DMCUST/00001"},
						"zipcode":           "12345",
						"city":              "BigCity",
						"notes":             "Create a new item by Keys",
						"address_metadata1": "value1",
						"address_metadata2": "value2~note2"},
				})))),
			},
			result: http.StatusOK,
		},
		{
			name: "update_error",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetParam: func(req *http.Request, name string) string {
					return "address"
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/update", bytes.NewBuffer([]byte(testData.encodeData([]nt.IM{
					{
						"keys": nt.IM{
							"nervatype": "customer",
							"ref_id":    "customer/00001"},
						"zipcode": "12345",
						"city":    "BigCity"},
				})))),
			},
			result: http.StatusBadRequest,
		},
		{
			name: "update_body_error",
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
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/update", bytes.NewBuffer([]byte("{error"))),
			},
			result: http.StatusBadRequest,
		},
		{
			name: "update_unauthorized",
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
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/update", bytes.NewBuffer([]byte(""))),
			},
			result: http.StatusUnauthorized,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &HTTPService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
				GetParam:      tt.fields.GetParam,
				GetTokenKeys:  tt.fields.GetTokenKeys,
			}
			if tt.name != "update_unauthorized" {
				tt.args.r.Header.Set("Authorization", "Bearer "+testData.adminToken)
				ctx, err := srv.TokenLogin(tt.args.r)
				if err != nil {
					t.Fatal(err)
				}
				tt.args.r = tt.args.r.WithContext(ctx)
			}
			srv.Update(tt.args.w, tt.args.r)
			if status := tt.args.w.(*httptest.ResponseRecorder).Code; status != tt.result {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.result)
			}
		})
	}
}

func TestHTTPService_Delete(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
		GetParam      func(req *http.Request, name string) string
		GetTokenKeys  func() map[string]map[string]string
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		result int
	}{
		{
			name: "delete_ok",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetParam: func(req *http.Request, name string) string {
					return "address"
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("DELETE", "/address?id=55", nil),
			},
			result: http.StatusNoContent,
		},
		{
			name: "delete_error",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
					})
				},
				GetParam: func(req *http.Request, name string) string {
					return "kalevala"
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("DELETE", "/customer?id=44", nil),
			},
			result: http.StatusBadRequest,
		},
		{
			name: "delete_unauthorized",
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
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("DELETE", "/customer?id=44", nil),
			},
			result: http.StatusUnauthorized,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &HTTPService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
				GetParam:      tt.fields.GetParam,
				GetTokenKeys:  tt.fields.GetTokenKeys,
			}
			if tt.name != "delete_unauthorized" {
				tt.args.r.Header.Set("Authorization", "Bearer "+testData.adminToken)
				ctx, err := srv.TokenLogin(tt.args.r)
				if err != nil {
					t.Fatal(err)
				}
				tt.args.r = tt.args.r.WithContext(ctx)
			}
			srv.Delete(tt.args.w, tt.args.r)
			if status := tt.args.w.(*httptest.ResponseRecorder).Code; status != tt.result {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.result)
			}
		})
	}
}

func TestHTTPService_ReportList(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
		GetParam      func(req *http.Request, name string) string
		GetTokenKeys  func() map[string]map[string]string
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		token  string
		result int
	}{
		{
			name: "list_ok",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
						"NT_REPORT_DIR":        "",
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/report/list?label=invoice", nil),
			},
			token:  testData.adminToken,
			result: http.StatusOK,
		},
		{
			name: "list_scope_error",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
						"NT_REPORT_DIR":        "",
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/report/list?label=invoice", nil),
			},
			token:  testData.userToken,
			result: http.StatusUnauthorized,
		},
		{
			name: "list_error",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
						"NT_REPORT_DIR":        "bad",
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/report/list?label=invoice", nil),
			},
			token:  testData.adminToken,
			result: http.StatusBadRequest,
		},
		{
			name: "list_unauthorized",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
						"NT_REPORT_DIR":        "",
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/report/list?label=invoice", nil),
			},
			token:  "",
			result: http.StatusUnauthorized,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &HTTPService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
				GetParam:      tt.fields.GetParam,
				GetTokenKeys:  tt.fields.GetTokenKeys,
			}
			if tt.name != "list_unauthorized" {
				tt.args.r.Header.Set("Authorization", "Bearer "+tt.token)
				ctx, err := srv.TokenLogin(tt.args.r)
				if err != nil {
					t.Fatal(err)
				}
				tt.args.r = tt.args.r.WithContext(ctx)
			}
			srv.ReportList(tt.args.w, tt.args.r)
			if status := tt.args.w.(*httptest.ResponseRecorder).Code; status != tt.result {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.result)
			}
		})
	}
}

func TestHTTPService_ReportDelete(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
		GetParam      func(req *http.Request, name string) string
		GetTokenKeys  func() map[string]map[string]string
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		token  string
		result int
	}{
		{
			name: "delete_ok",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
						"NT_REPORT_DIR":        "",
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("DELETE", "/delete?reportkey=ntr_invoice_en", nil),
			},
			token:  testData.adminToken,
			result: http.StatusNoContent,
		},
		{
			name: "delete_error",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
						"NT_REPORT_DIR":        "",
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("DELETE", "/delete?reportkey=ntr_invoice_en", nil),
			},
			token:  testData.adminToken,
			result: http.StatusBadRequest,
		},
		{
			name: "delete_scope_error",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
						"NT_REPORT_DIR":        "",
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("DELETE", "/delete?reportkey=ntr_invoice_en", nil),
			},
			token:  testData.userToken,
			result: http.StatusUnauthorized,
		},
		{
			name: "delete_unauthorized",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
						"NT_REPORT_DIR":        "",
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("DELETE", "/delete?reportkey=ntr_invoice_en", nil),
			},
			token:  "",
			result: http.StatusUnauthorized,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &HTTPService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
				GetParam:      tt.fields.GetParam,
				GetTokenKeys:  tt.fields.GetTokenKeys,
			}
			if tt.name != "delete_unauthorized" {
				tt.args.r.Header.Set("Authorization", "Bearer "+tt.token)
				ctx, err := srv.TokenLogin(tt.args.r)
				if err != nil {
					t.Fatal(err)
				}
				tt.args.r = tt.args.r.WithContext(ctx)
			}
			srv.ReportDelete(tt.args.w, tt.args.r)
			if status := tt.args.w.(*httptest.ResponseRecorder).Code; status != tt.result {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.result)
			}
		})
	}
}

func TestHTTPService_ReportInstall(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
		GetParam      func(req *http.Request, name string) string
		GetTokenKeys  func() map[string]map[string]string
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		token  string
		result int
	}{
		{
			name: "in_ok",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
						"NT_REPORT_DIR":        "",
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/install?reportkey=ntr_invoice_en", nil),
			},
			token:  testData.adminToken,
			result: http.StatusOK,
		},
		{
			name: "in_error",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
						"NT_REPORT_DIR":        "",
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/install?reportkey=ntr_invoice_en", nil),
			},
			token:  testData.adminToken,
			result: http.StatusBadRequest,
		},
		{
			name: "in_scope_error",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
						"NT_REPORT_DIR":        "",
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/install?reportkey=ntr_invoice_en", nil),
			},
			token:  testData.userToken,
			result: http.StatusUnauthorized,
		},
		{
			name: "in_unauthorized",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
						"NT_REPORT_DIR":        "",
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/install?reportkey=ntr_invoice_en", nil),
			},
			token:  "",
			result: http.StatusUnauthorized,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &HTTPService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
				GetParam:      tt.fields.GetParam,
				GetTokenKeys:  tt.fields.GetTokenKeys,
			}
			if tt.name != "in_unauthorized" {
				tt.args.r.Header.Set("Authorization", "Bearer "+tt.token)
				ctx, err := srv.TokenLogin(tt.args.r)
				if err != nil {
					t.Fatal(err)
				}
				tt.args.r = tt.args.r.WithContext(ctx)
			}
			srv.ReportInstall(tt.args.w, tt.args.r)
			if status := tt.args.w.(*httptest.ResponseRecorder).Code; status != tt.result {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.result)
			}
		})
	}
}

func TestHTTPService_Report(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
		GetParam      func(req *http.Request, name string) string
		GetTokenKeys  func() map[string]map[string]string
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		result int
	}{
		{
			name: "report_post_refnumber_pdf_ok",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
						"NT_REPORT_DIR":        "",
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/report", bytes.NewBuffer([]byte(testData.encodeData(nt.IM{
					"reportkey":   "ntr_invoice_en",
					"orientation": "portrait",
					"size":        "a4",
					"nervatype":   "trans",
					"refnumber":   "DMINV/00001",
				})))),
			},
			result: http.StatusOK,
		},
		{
			name: "report_get_report_id_pdf_error",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
						"NT_REPORT_DIR":        "",
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET",
					"/report?reportkey=ntr_invoice_en&orientation=portrait&size=a4&nervatype=trans&report_id=123456", nil),
			},
			result: http.StatusBadRequest,
		},
		{
			name: "report_get_xml_ok",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
						"NT_REPORT_DIR":        "",
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET",
					"/report?reportkey=ntr_invoice_en&output=xml&nervatype=trans&refnumber=DMINV/00001", nil),
			},
			result: http.StatusOK,
		},
		{
			name: "report_get_data_ok",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
						"NT_REPORT_DIR":        "",
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET",
					"/report?reportkey=ntr_invoice_en&output=data&nervatype=trans&refnumber=DMINV/00001", nil),
			},
			result: http.StatusOK,
		},
		{
			name: "report_get_csv_ok",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
						"NT_REPORT_DIR":        "",
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET",
					"/report?reportkey=csv_vat_en&filters[date_from]=2014-01-01&filters[date_to]=2024-01-01&filters[curr]=EUR", nil),
			},
			result: http.StatusBadRequest,
		},
		{
			name: "report_unauthorized",
			fields: fields{
				Config: nt.IM{},
				GetNervaStore: func(database string) *nt.NervaStore {
					return nt.New(testData.driver, nt.IM{
						"NT_HASHTABLE":         testData.hashTable,
						"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
						"NT_REPORT_DIR":        "",
					})
				},
				GetTokenKeys: func() map[string]map[string]string {
					return nil
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/report?reportkey=csv_vat_en", nil),
			},
			result: http.StatusUnauthorized,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &HTTPService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
				GetParam:      tt.fields.GetParam,
				GetTokenKeys:  tt.fields.GetTokenKeys,
			}
			if tt.name != "report_unauthorized" {
				tt.args.r.Header.Set("Authorization", "Bearer "+testData.adminToken)
				ctx, err := srv.TokenLogin(tt.args.r)
				if err != nil {
					t.Fatal(err)
				}
				tt.args.r = tt.args.r.WithContext(ctx)
			}
			srv.Report(tt.args.w, tt.args.r)
			if status := tt.args.w.(*httptest.ResponseRecorder).Code; status != tt.result {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.result)
			}
		})
	}
}

func TestHTTPService_DatabaseCreate(t *testing.T) {
	type fields struct {
		Config        map[string]interface{}
		GetNervaStore func(database string) *nt.NervaStore
		GetParam      func(req *http.Request, name string) string
		GetTokenKeys  func() map[string]map[string]string
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		result int
	}{
		{
			name: "database_ok",
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
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/database?alias=test&demo=true", nil),
			},
			result: http.StatusOK,
		},
		{
			name: "database_error",
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
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/database?alias=kalevala&demo=true", nil),
			},
			result: http.StatusBadRequest,
		},
		{
			name: "database_unauthorized",
			fields: fields{
				Config: nt.IM{},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/database?alias=test&demo=true", nil),
			},
			result: http.StatusUnauthorized,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &HTTPService{
				Config:        tt.fields.Config,
				GetNervaStore: tt.fields.GetNervaStore,
				GetParam:      tt.fields.GetParam,
				GetTokenKeys:  tt.fields.GetTokenKeys,
			}
			if tt.name != "database_unauthorized" {
				tt.args.r.Header.Set("X-Api-Key", testData.apiKey)
			}
			srv.DatabaseCreate(tt.args.w, tt.args.r)
			if status := tt.args.w.(*httptest.ResponseRecorder).Code; status != tt.result {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.result)
			}
		})
	}
}
