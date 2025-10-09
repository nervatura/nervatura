package service

import (
	"errors"
	"log/slog"
	"reflect"
	"testing"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	api "github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	"golang.org/x/oauth2"
)

func TestClientService_searchFilter(t *testing.T) {
	type fields struct {
		Config       cu.IM
		AuthConfigs  map[string]*oauth2.Config
		AppLog       *slog.Logger
		Session      *api.SessionService
		NewDataStore func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore
	}
	type args struct {
		view         string
		filter       ct.BrowserFilter
		queryFilters []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "customer_inactive",
			args: args{
				view: "customer",
				filter: ct.BrowserFilter{
					Field: "inactive",
					Comp:  "==",
					Value: true,
				},
				queryFilters: []string{},
			},
			want: []string{"and (inactive = true)"},
		},
		{
			name: "customer_code",
			args: args{
				view: "customer",
				filter: ct.BrowserFilter{
					Or:    true,
					Field: "code",
					Comp:  "==",
					Value: "123",
				},
				queryFilters: []string{},
			},
			want: []string{"or (code like '%123%')"},
		},
		{
			name: "customer_id",
			args: args{
				view: "customer",
				filter: ct.BrowserFilter{
					Field: "id",
					Comp:  "==",
					Value: "456",
				},
				queryFilters: []string{},
			},
			want: []string{"and (id = 456)"},
		},
		{
			name: "customer_name",
			args: args{
				view: "customer",
				filter: ct.BrowserFilter{
					Field: "name",
					Comp:  "==",
					Value: "John Doe",
				},
				queryFilters: []string{},
			},
			want: []string{},
		},
		{
			name: "customer_simple",
			args: args{
				view: "customer_simple",
				filter: ct.BrowserFilter{
					Field: "code",
					Comp:  "==",
					Value: "123",
				},
				queryFilters: []string{},
			},
			want: []string{"and (code like '%123%')"},
		},
		{
			name: "customer_addresses",
			args: args{
				view: "customer_addresses",
				filter: ct.BrowserFilter{
					Field: "city",
					Comp:  "==",
					Value: "123",
				},
				queryFilters: []string{},
			},
			want: []string{"and (city like '%123%')"},
		},
		{
			name: "customer_contacts",
			args: args{
				view: "customer_contacts",
				filter: ct.BrowserFilter{
					Field: "name",
					Comp:  "==",
					Value: "John Doe",
				},
				queryFilters: []string{},
			},
			want: []string{"and (name like '%John Doe%')"},
		},
		{
			name: "customer_events-start_time",
			args: args{
				view: "customer_events",
				filter: ct.BrowserFilter{
					Field: "start_time",
					Comp:  "==",
					Value: "2021-01-01",
				},
				queryFilters: []string{},
			},
			want: []string{"and (start_time = '2021-01-01')"},
		},
		{
			name: "customer_events-default",
			args: args{
				view: "customer_events",
				filter: ct.BrowserFilter{
					Field: "name",
					Comp:  "==",
					Value: "John Doe",
				},
				queryFilters: []string{},
			},
			want: []string{"and (name like '%John Doe%')"},
		},
		{
			name: "customer_map-default",
			args: args{
				view: "customer_map",
				filter: ct.BrowserFilter{
					Field: "name",
					Comp:  "==",
					Value: "John Doe",
				},
				queryFilters: []string{},
			},
			want: []string{"and (name like '%John Doe%')"},
		},
		{
			name: "product_simple",
			args: args{
				view: "product_simple",
				filter: ct.BrowserFilter{
					Field: "code",
					Comp:  "==",
					Value: "123",
				},
				queryFilters: []string{},
			},
			want: []string{"and (code like '%123%')"},
		},
		{
			name: "product_inactive",
			args: args{
				view: "product",
				filter: ct.BrowserFilter{
					Field: "inactive",
					Comp:  "==",
					Value: true,
				},
				queryFilters: []string{},
			},
			want: []string{"and (inactive = true)"},
		},
		{
			name: "product_code",
			args: args{
				view: "product",
				filter: ct.BrowserFilter{
					Field: "code",
					Comp:  "==",
					Value: "123",
				},
				queryFilters: []string{},
			},
			want: []string{"and (code like '%123%')"},
		},
		{
			name: "product_id",
			args: args{
				view: "product",
				filter: ct.BrowserFilter{
					Field: "id",
					Comp:  "==",
					Value: "456",
				},
				queryFilters: []string{},
			},
			want: []string{"and (id = 456)"},
		},
		{
			name: "product_default",
			args: args{
				view: "product",
				filter: ct.BrowserFilter{
					Field: "default",
				},
				queryFilters: []string{},
			},
			want: []string{},
		},
		{
			name: "product_map",
			args: args{
				view: "product_map",
				filter: ct.BrowserFilter{
					Field: "demo_number",
					Comp:  "==",
					Value: "123",
				},
				queryFilters: []string{},
			},
			want: []string{"and (demo_number like '%123%')"},
		},
		{
			name: "product_events",
			args: args{
				view: "product_events",
				filter: ct.BrowserFilter{
					Field: "start_time",
					Comp:  "==",
					Value: "2021-01-01",
				},
				queryFilters: []string{},
			},
			want: []string{"and (start_time = '2021-01-01')"},
		},
		{
			name: "product_events_default",
			args: args{
				view: "product_events",
				filter: ct.BrowserFilter{
					Field: "name",
					Comp:  "==",
					Value: "John Doe",
				},
				queryFilters: []string{},
			},
			want: []string{"and (name like '%John Doe%')"},
		},
		{
			name: "product_prices_valid_from",
			args: args{
				view: "product_prices",
				filter: ct.BrowserFilter{
					Field: "valid_from",
					Comp:  "==",
					Value: "2021-01-01",
				},
				queryFilters: []string{},
			},
			want: []string{"and (valid_from = '2021-01-01')"},
		},
		{
			name: "product_prices_id",
			args: args{
				view: "product_prices",
				filter: ct.BrowserFilter{
					Field: "id",
					Comp:  "==",
					Value: "456",
				},
				queryFilters: []string{},
			},
			want: []string{"and (id = 456)"},
		},
		{
			name: "product_default",
			args: args{
				view: "product_prices",
				filter: ct.BrowserFilter{
					Field: "default",
					Comp:  "==",
					Value: "123",
				},
				queryFilters: []string{},
			},
			want: []string{"and (default like '%123%')"},
		},
		{
			name: "tool_simple",
			args: args{
				view: "tool_simple",
				filter: ct.BrowserFilter{
					Field: "code",
					Comp:  "==",
					Value: "123",
				},
				queryFilters: []string{},
			},
			want: []string{"and (code like '%123%')"},
		},
		{
			name: "tool_inactive",
			args: args{
				view: "tool",
				filter: ct.BrowserFilter{
					Field: "inactive",
					Comp:  "==",
					Value: true,
				},
				queryFilters: []string{},
			},
			want: []string{"and (inactive = true)"},
		},
		{
			name: "tool_code",
			args: args{
				view: "tool",
				filter: ct.BrowserFilter{
					Field: "code",
					Comp:  "==",
					Value: "123",
				},
				queryFilters: []string{},
			},
			want: []string{"and (code like '%123%')"},
		},
		{
			name: "tool_id",
			args: args{
				view: "tool",
				filter: ct.BrowserFilter{
					Field: "id",
					Comp:  "==",
					Value: "456",
				},
				queryFilters: []string{},
			},
			want: []string{"and (id = 456)"},
		},
		{
			name: "tool_default",
			args: args{
				view: "tool",
				filter: ct.BrowserFilter{
					Field: "default",
					Comp:  "==",
					Value: "123",
				},
				queryFilters: []string{},
			},
			want: []string{},
		},
		{
			name: "tool_map",
			args: args{
				view: "tool_map",
				filter: ct.BrowserFilter{
					Field: "demo_number",
					Comp:  "==",
					Value: "123",
				},
				queryFilters: []string{},
			},
			want: []string{"and (demo_number like '%123%')"},
		},
		{
			name: "tool_events",
			args: args{
				view: "tool_events",
				filter: ct.BrowserFilter{
					Field: "start_time",
					Comp:  "==",
					Value: "2021-01-01",
				},
				queryFilters: []string{},
			},
			want: []string{"and (start_time = '2021-01-01')"},
		},
		{
			name: "tool_events_default",
			args: args{
				view: "tool_events",
				filter: ct.BrowserFilter{
					Field: "name",
					Comp:  "==",
					Value: "John Doe",
				},
				queryFilters: []string{},
			},
			want: []string{"and (name like '%John Doe%')"},
		},
		{
			name: "project_simple",
			args: args{
				view: "project_simple",
				filter: ct.BrowserFilter{
					Field: "code",
					Comp:  "==",
					Value: "123",
				},
				queryFilters: []string{},
			},
			want: []string{"and (code like '%123%')"},
		},
		{
			name: "project_inactive",
			args: args{
				view: "project",
				filter: ct.BrowserFilter{
					Field: "inactive",
					Comp:  "==",
					Value: true,
				},
				queryFilters: []string{},
			},
			want: []string{"and (inactive = true)"},
		},
		{
			name: "project_start_date",
			args: args{
				view: "project",
				filter: ct.BrowserFilter{
					Field: "start_date",
					Comp:  "==",
					Value: "2021-01-01",
				},
				queryFilters: []string{},
			},
			want: []string{"and (start_date = '2021-01-01')"},
		},
		{
			name: "project_code",
			args: args{
				view: "project",
				filter: ct.BrowserFilter{
					Field: "code",
					Comp:  "==",
					Value: "123",
				},
				queryFilters: []string{},
			},
			want: []string{"and (code like '%123%')"},
		},
		{
			name: "project_id",
			args: args{
				view: "project",
				filter: ct.BrowserFilter{
					Field: "id",
					Comp:  "==",
					Value: "456",
				},
				queryFilters: []string{},
			},
			want: []string{"and (id = 456)"},
		},
		{
			name: "project_default",
			args: args{
				view: "project",
				filter: ct.BrowserFilter{
					Field: "default",
				},
				queryFilters: []string{},
			},
			want: []string{},
		},
		{
			name: "project_map",
			args: args{
				view: "project_map",
				filter: ct.BrowserFilter{
					Field: "demo_number",
					Comp:  "==",
					Value: "123",
				},
				queryFilters: []string{},
			},
			want: []string{"and (demo_number like '%123%')"},
		},
		{
			name: "project_addresses",
			args: args{
				view: "project_addresses",
				filter: ct.BrowserFilter{
					Field: "city",
					Comp:  "==",
					Value: "123",
				},
				queryFilters: []string{},
			},
			want: []string{"and (city like '%123%')"},
		},
		{
			name: "project_contacts",
			args: args{
				view: "project_contacts",
				filter: ct.BrowserFilter{
					Field: "name",
					Comp:  "==",
					Value: "John Doe",
				},
				queryFilters: []string{},
			},
			want: []string{"and (name like '%John Doe%')"},
		},
		{
			name: "project_events_start_time",
			args: args{
				view: "project_events",
				filter: ct.BrowserFilter{
					Field: "start_time",
					Comp:  "==",
					Value: "2021-01-01",
				},
				queryFilters: []string{},
			},
			want: []string{"and (start_time = '2021-01-01')"},
		},
		{
			name: "project_events_default",
			args: args{
				view: "project_events",
				filter: ct.BrowserFilter{
					Field: "name",
					Comp:  "==",
					Value: "John Doe",
				},
				queryFilters: []string{},
			},
			want: []string{"and (name like '%John Doe%')"},
		},
		{
			name: "employee_simple",
			args: args{
				view: "employee_simple",
				filter: ct.BrowserFilter{
					Field: "code",
					Comp:  "==",
					Value: "123",
				},
				queryFilters: []string{},
			},
			want: []string{"and (code like '%123%')"},
		},
		{
			name: "employee_inactive",
			args: args{
				view: "employee",
				filter: ct.BrowserFilter{
					Field: "inactive",
					Comp:  "==",
					Value: true,
				},
				queryFilters: []string{},
			},
			want: []string{"and (inactive = true)"},
		},
		{
			name: "employee_start_date",
			args: args{
				view: "employee",
				filter: ct.BrowserFilter{
					Field: "start_date",
					Comp:  "==",
					Value: "2021-01-01",
				},
				queryFilters: []string{},
			},
			want: []string{"and (start_date = '2021-01-01')"},
		},
		{
			name: "employee_code",
			args: args{
				view: "employee",
				filter: ct.BrowserFilter{
					Field: "code",
					Comp:  "==",
					Value: "123",
				},
				queryFilters: []string{},
			},
			want: []string{"and (code like '%123%')"},
		},
		{
			name: "employee_id",
			args: args{
				view: "employee",
				filter: ct.BrowserFilter{
					Field: "id",
					Comp:  "==",
					Value: "456",
				},
				queryFilters: []string{},
			},
			want: []string{"and (id = 456)"},
		},
		{
			name: "employee_default",
			args: args{
				view: "employee",
				filter: ct.BrowserFilter{
					Field: "default",
				},
				queryFilters: []string{},
			},
			want: []string{},
		},
		{
			name: "employee_map",
			args: args{
				view: "employee_map",
				filter: ct.BrowserFilter{
					Field: "demo_number",
					Comp:  "==",
					Value: "123",
				},
				queryFilters: []string{},
			},
			want: []string{"and (demo_number like '%123%')"},
		},
		{
			name: "employee_events",
			args: args{
				view: "employee_events",
				filter: ct.BrowserFilter{
					Field: "start_time",
					Comp:  "==",
					Value: "2021-01-01",
				},
				queryFilters: []string{},
			},
			want: []string{"and (start_time = '2021-01-01')"},
		},
		{
			name: "employee_events_default",
			args: args{
				view: "employee_events",
				filter: ct.BrowserFilter{
					Field: "name",
					Comp:  "==",
					Value: "John Doe",
				},
				queryFilters: []string{},
			},
			want: []string{"and (name like '%John Doe%')"},
		},
		{
			name: "place_simple",
			args: args{
				view: "place_simple",
				filter: ct.BrowserFilter{
					Field: "code",
					Comp:  "==",
					Value: "123",
				},
			},
			want: []string{"and (code like '%123%')"},
		},
		{
			name: "place_inactive",
			args: args{
				view: "place",
				filter: ct.BrowserFilter{
					Field: "inactive",
					Comp:  "==",
					Value: true,
				},
				queryFilters: []string{},
			},
			want: []string{"and (inactive = true)"},
		},
		{
			name: "place_code",
			args: args{
				view: "place",
				filter: ct.BrowserFilter{
					Field: "code",
					Comp:  "==",
					Value: "123",
				},
			},
			want: []string{"and (code like '%123%')"},
		},
		{
			name: "place_id",
			args: args{
				view: "place",
				filter: ct.BrowserFilter{
					Field: "id",
					Comp:  "==",
					Value: "456",
				},
				queryFilters: []string{},
			},
			want: []string{"and (id = 456)"},
		},
		{
			name: "place_default",
			args: args{
				view: "place",
				filter: ct.BrowserFilter{
					Field: "default",
				},
				queryFilters: []string{},
			},
			want: []string{},
		},
		{
			name: "place_map",
			args: args{
				view: "place_map",
				filter: ct.BrowserFilter{
					Field: "demo_number",
					Comp:  "==",
					Value: "123",
				},
				queryFilters: []string{},
			},
			want: []string{"and (demo_number like '%123%')"},
		},
		{
			name: "place_contacts",
			args: args{
				view: "place_contacts",
				filter: ct.BrowserFilter{
					Field: "name",
					Comp:  "==",
					Value: "John Doe",
				},
				queryFilters: []string{},
			},
			want: []string{"and (name like '%John Doe%')"},
		},
		{
			name: "transitem_simple",
			args: args{
				view: "transitem_simple",
				filter: ct.BrowserFilter{
					Field: "code",
					Comp:  "==",
					Value: "123",
				},
				queryFilters: []string{},
			},
			want: []string{"and (code like '%123%')"},
		},
		{
			name: "invoice_simple",
			args: args{
				view: "invoice_simple",
				filter: ct.BrowserFilter{
					Field: "code",
					Comp:  "==",
					Value: "123",
				},
				queryFilters: []string{},
			},
			want: []string{"and (code like '%123%')"},
		},
		{
			name: "transitem_closed",
			args: args{
				view: "transitem",
				filter: ct.BrowserFilter{
					Field: "closed",
					Comp:  "==",
					Value: true,
				},
				queryFilters: []string{},
			},
			want: []string{"and (closed = true)"},
		},
		{
			name: "transitem_trans_date",
			args: args{
				view: "transitem",
				filter: ct.BrowserFilter{
					Field: "trans_date",
					Comp:  "==",
					Value: "2021-01-01",
				},
				queryFilters: []string{},
			},
			want: []string{"and (trans_date = '2021-01-01')"},
		},
		{
			name: "transitem_id",
			args: args{
				view: "transitem",
				filter: ct.BrowserFilter{
					Field: "id",
					Comp:  "==",
					Value: "456",
				},
				queryFilters: []string{},
			},
			want: []string{"and (id = 456)"},
		},
		{
			name: "transitem_default",
			args: args{
				view: "transitem",
				filter: ct.BrowserFilter{
					Field: "default",
				},
				queryFilters: []string{},
			},
			want: []string{"and (default  '%%')"},
		},
		{
			name: "transitem_map",
			args: args{
				view: "transitem_map",
				filter: ct.BrowserFilter{
					Field: "demo_number",
					Comp:  "==",
					Value: "123",
				},
				queryFilters: []string{},
			},
			want: []string{"and (demo_number like '%123%')"},
		},
		{
			name: "transitem_item",
			args: args{
				view: "transitem_item",
				filter: ct.BrowserFilter{
					Field: "qty",
					Comp:  "==",
					Value: "123",
				},
				queryFilters: []string{},
			},
			want: []string{"and (qty = 123)"},
		},
		{
			name: "transitem_item_deposit",
			args: args{
				view: "transitem_item",
				filter: ct.BrowserFilter{
					Field: "deposit",
					Comp:  "==",
					Value: true,
				},
				queryFilters: []string{},
			},
			want: []string{"and (deposit = true)"},
		},
		{
			name: "transitem_item_default",
			args: args{
				view: "transitem_item",
				filter: ct.BrowserFilter{
					Field: "default",
				},
				queryFilters: []string{},
			},
			want: []string{"and (default  '%%')"},
		},
		{
			name: "transpayment_simple",
			args: args{
				view: "transpayment_simple",
				filter: ct.BrowserFilter{
					Field: "default",
				},
				queryFilters: []string{},
			},
			want: []string{"and (default  '%%')"},
		},
		{
			name: "transpayment_closed",
			args: args{
				view: "transpayment",
				filter: ct.BrowserFilter{
					Field: "closed",
					Comp:  "==",
					Value: true,
				},
			},
			want: []string{"and (closed = true)"},
		},
		{
			name: "transpayment_trans_date",
			args: args{
				view: "transpayment",
				filter: ct.BrowserFilter{
					Field: "trans_date",
					Comp:  "==",
					Value: "2021-01-01",
				},
			},
			want: []string{"and (trans_date = '2021-01-01')"},
		},
		{
			name: "transpayment_id",
			args: args{
				view: "transpayment",
				filter: ct.BrowserFilter{
					Field: "id",
					Comp:  "==",
					Value: "456",
				},
			},
			want: []string{"and (id = 456)"},
		},
		{
			name: "transpayment_default",
			args: args{
				view: "transpayment",
				filter: ct.BrowserFilter{
					Field: "default",
				},
				queryFilters: []string{},
			},
			want: []string{"and (default  '%%')"},
		},
		{
			name: "transpayment_map",
			args: args{
				view: "transpayment_map",
				filter: ct.BrowserFilter{
					Field: "demo_number",
					Comp:  "==",
					Value: "123",
				},
				queryFilters: []string{},
			},
			want: []string{"and (demo_number like '%123%')"},
		},
		{
			name: "transpayment_invoice",
			args: args{
				view: "transpayment_invoice",
				filter: ct.BrowserFilter{
					Field: "invoice_amount",
					Comp:  "==",
					Value: "123",
				},
				queryFilters: []string{},
			},
			want: []string{"and (invoice_amount = 123)"},
		},
		{
			name: "transpayment_invoice_id",
			args: args{
				view: "transpayment_invoice",
				filter: ct.BrowserFilter{
					Field: "id",
					Comp:  "==",
					Value: "456",
				},
			},
			want: []string{"and (id = 456)"},
		},
		{
			name: "transpayment_invoice_trans_date",
			args: args{
				view: "transpayment_invoice",
				filter: ct.BrowserFilter{
					Field: "trans_date",
					Comp:  "==",
					Value: "2021-01-01",
				},
				queryFilters: []string{},
			},
			want: []string{"and (trans_date = '2021-01-01')"},
		},
		{
			name: "transpayment_invoice_currency_code",
			args: args{
				view: "transpayment_invoice",
				filter: ct.BrowserFilter{
					Field: "currency_code",
					Comp:  "==",
					Value: "USD",
				},
			},
			want: []string{"and (currency_code = 'USD')"},
		},
		{
			name: "transpayment_invoice_default",
			args: args{
				view: "transpayment_invoice",
				filter: ct.BrowserFilter{
					Field: "default",
					Comp:  "==",
					Value: "abc",
				},
				queryFilters: []string{},
			},
			want: []string{"and (default like '%abc%')"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClientService{
				Config:       tt.fields.Config,
				AuthConfigs:  tt.fields.AuthConfigs,
				AppLog:       tt.fields.AppLog,
				Session:      tt.fields.Session,
				NewDataStore: tt.fields.NewDataStore,
			}
			if got := cls.searchFilter(tt.args.view, tt.args.filter, tt.args.queryFilters); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientService.searchFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientService_searchData(t *testing.T) {
	type fields struct {
		Config       cu.IM
		AuthConfigs  map[string]*oauth2.Config
		AppLog       *slog.Logger
		Session      *api.SessionService
		NewDataStore func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore
	}
	type args struct {
		ds      *api.DataStore
		view    string
		query   md.Query
		filters []ct.BrowserFilter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "customer",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
			},
			args: args{
				ds: &api.DataStore{
					Db: &md.TestDriver{Config: cu.IM{
						"Query": func(queries []md.Query) ([]cu.IM, error) {
							return []cu.IM{{"id": 1, "name": "test"}}, errors.New("error")
						},
					}},
					Config: cu.IM{},
					AppLog: slog.Default(),
					ConvertToType: func(data interface{}, result any) (err error) {
						return nil
					},
				},
				view: "customer",
				query: md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
					Filter: "field like 'abc'",
				},
				filters: []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClientService{
				Config:       tt.fields.Config,
				AuthConfigs:  tt.fields.AuthConfigs,
				AppLog:       tt.fields.AppLog,
				Session:      tt.fields.Session,
				NewDataStore: tt.fields.NewDataStore,
			}
			_, err := cls.searchData(tt.args.ds, tt.args.view, tt.args.query, tt.args.filters)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientService.searchData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClientService_searchResponse(t *testing.T) {
	type fields struct {
		Config       cu.IM
		AuthConfigs  map[string]*oauth2.Config
		AppLog       *slog.Logger
		Session      *api.SessionService
		NewDataStore func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore
	}
	type args struct {
		evt ct.ResponseEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "side_menu_group_open",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:     &md.TestDriver{Config: cu.IM{}},
						Config: config,
						AppLog: appLog,
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.ClientEventSideMenu,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"search": cu.IM{
									"view":       "customer",
									"side_group": "group_customer",
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: "group_product",
				},
			},
		},
		{
			name: "side_menu_group_close",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:     &md.TestDriver{Config: cu.IM{}},
						Config: config,
						AppLog: appLog,
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.ClientEventSideMenu,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"search": cu.IM{
									"view":       "customer",
									"side_group": "group_value_1",
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: "group_value_1",
				},
			},
		},
		{
			name: "side_menu_default",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db:     &md.TestDriver{Config: cu.IM{}},
						Config: config,
						AppLog: appLog,
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Name: ct.ClientEventSideMenu,
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"search": cu.IM{
									"view": "customer",
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{},
						},
					},
					Value: "value_2",
				},
			},
		},
		{
			name: "bookmark_add",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Update": func(data md.Update) (int64, error) {
								return 1, nil
							},
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1}}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
						ConvertToType: func(data interface{}, result any) (err error) {
							return nil
						},
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"search": cu.IM{
									"view":     "customer",
									"customer": cu.IM{},
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{
								"id": 1,
							},
						},
						Lang: "en",
						ClientLabels: func(lang string) cu.SM {
							return cu.SM{}
						},
					},
					Name:  ct.FormEventOK,
					Value: cu.IM{"data": cu.IM{"next": "bookmark_add"}, "value": cu.IM{"value": "label"}},
				},
			},
			wantErr: false,
		},
		{
			name: "bookmark_new",
			fields: fields{
				Config: cu.IM{},
				AppLog: slog.Default(),
				NewDataStore: func(config cu.IM, alias string, appLog *slog.Logger) *api.DataStore {
					return &api.DataStore{
						Db: &md.TestDriver{Config: cu.IM{
							"Update": func(data md.Update) (int64, error) {
								return 1, nil
							},
							"Query": func(queries []md.Query) ([]cu.IM, error) {
								return []cu.IM{{"id": 1}}, nil
							},
						}},
						Config: config,
						AppLog: appLog,
						ConvertToType: func(data interface{}, result any) (err error) {
							return nil
						},
					}
				},
			},
			args: args{
				evt: ct.ResponseEvent{
					Trigger: &ct.Client{
						BaseComponent: ct.BaseComponent{
							Data: cu.IM{
								"search": cu.IM{
									"view":     "customer",
									"customer": cu.IM{},
								},
							},
						},
						Ticket: ct.Ticket{
							User: cu.IM{
								"id": 1,
							},
						},
						Lang: "en",
						ClientLabels: func(lang string) cu.SM {
							return cu.SM{}
						},
					},
					Name:  ct.BrowserEventBookmark,
					Value: cu.IM{"data": cu.IM{"next": "bookmark_add"}, "value": cu.IM{"value": "label"}},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClientService{
				Config:       tt.fields.Config,
				AuthConfigs:  tt.fields.AuthConfigs,
				AppLog:       tt.fields.AppLog,
				Session:      tt.fields.Session,
				NewDataStore: tt.fields.NewDataStore,
			}
			_, err := cls.searchResponse(tt.args.evt)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientService.searchResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
