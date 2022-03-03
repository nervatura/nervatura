package service

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path"
	"strings"
	"testing"
	"text/template"

	db "github.com/nervatura/nervatura/service/pkg/database"
	nt "github.com/nervatura/nervatura/service/pkg/nervatura"
	ut "github.com/nervatura/nervatura/service/pkg/utils"
)

var testData = struct {
	version, hashTable, tokenKey, apiKey, testDatabase,
	adminToken, userToken, customerToken string
	encodeData func(data interface{}) string
	getApi     func(nstore *nt.NervaStore, token string) *nt.API
	formReq    func(params url.Values) *http.Request
	templates  func() *template.Template
	driver     nt.DataDriver
}{
	version:      "test",
	hashTable:    "ref17890714",
	tokenKey:     "TEST_TOKEN_KEY",
	apiKey:       "TEST_API_KEY",
	testDatabase: "sqlite://file::memory:?cache=shared",
	//testDatabase: "postgres://postgres:admin@172.19.0.1:5432/nervatura?sslmode=disable",
	//testDatabase: "mysql://root:admin@tcp(localhost:3306)/nervatura",
	encodeData: func(data interface{}) string {
		jdata, _ := json.Marshal(data)
		return string(jdata)
	},
	getApi: func(nstore *nt.NervaStore, token string) *nt.API {
		api := &nt.API{NStore: nstore}
		api.TokenLogin(nt.IM{"token": token, "keys": make(map[string]map[string]string)})
		return api
	},
	formReq: func(params url.Values) *http.Request {
		req := httptest.NewRequest("POST", "/admin", strings.NewReader(params.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		return req
	},
	templates: func() *template.Template {
		tmpl := new(template.Template)
		tmpl.ParseFS(ut.Static, path.Join("static", "views", "*.html"))
		return tmpl
	},
}

func init() {
	testData.driver = &db.SQLDriver{
		Config: nt.IM{"version": testData.version}}
	api := &nt.API{NStore: nt.New(testData.driver, nt.IM{
		"NT_ALIAS_TEST":        testData.testDatabase,
		"NT_HASHTABLE":         testData.hashTable,
		"NT_TOKEN_PRIVATE_KEY": testData.tokenKey,
	})}
	_, err := api.DatabaseCreate(nt.IM{"database": "test", "demo": true})
	if err != nil {
		return
	}
	testData.adminToken, _ = ut.CreateToken("admin", "test", nt.IM{"NT_TOKEN_PRIVATE_KEY": testData.tokenKey})
	testData.userToken, _ = ut.CreateToken("user", "test", nt.IM{"NT_TOKEN_PRIVATE_KEY": testData.tokenKey})
	testData.customerToken, _ = ut.CreateToken("DMCUST/00002", "test", nt.IM{"NT_TOKEN_PRIVATE_KEY": testData.tokenKey})
}

func TestContextKey_String(t *testing.T) {
	type fields struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "nstore",
			fields: fields{name: "nstore"},
			want:   "nstore",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &ContextKey{
				name: tt.fields.name,
			}
			if got := k.String(); got != tt.want {
				t.Errorf("ContextKey.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
