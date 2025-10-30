package component_test

import (
	"reflect"
	"testing"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/component/client/component"
	md "github.com/nervatura/nervatura/v6/pkg/model"
)

func TestSearchConfig_Filter(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		view         string
		filter       ct.BrowserFilter
		queryFilters []string
		want         []string
	}{
		{
			name: "customer_inactive",
			view: "customer",
			filter: ct.BrowserFilter{
				Field: "inactive",
				Comp:  "==",
				Value: true,
			},
			queryFilters: []string{},
			want:         []string{"and (inactive = true)"},
		},
		{
			name: "customer_code",
			view: "customer",
			filter: ct.BrowserFilter{
				Or:    true,
				Field: "code",
				Comp:  "==",
				Value: "123",
			},
			queryFilters: []string{},
			want:         []string{"or (CAST(code as CHAR(255)) like '%123%')"},
		},
		{
			name: "customer_id",
			view: "customer",
			filter: ct.BrowserFilter{
				Field: "id",
				Comp:  "==",
				Value: "456",
			},
			queryFilters: []string{},
			want:         []string{"and (id = 456)"},
		},
		{
			name: "customer_name",
			view: "customer",
			filter: ct.BrowserFilter{
				Field: "name",
				Comp:  "==",
				Value: "John Doe",
			},
			queryFilters: []string{},
			want:         []string{},
		},
		{
			name: "customer_simple",
			view: "customer_simple",
			filter: ct.BrowserFilter{
				Field: "code",
				Comp:  "==",
				Value: "123",
			},
			queryFilters: []string{},
			want:         []string{"and (CAST(code as CHAR(255)) like '%123%')"},
		},
		{
			name: "customer_addresses",
			view: "customer_addresses",
			filter: ct.BrowserFilter{
				Field: "city",
				Comp:  "==",
				Value: "123",
			},
			queryFilters: []string{},
			want:         []string{"and (city like '%123%')"},
		},
		{
			name: "customer_contacts",
			view: "customer_contacts",
			filter: ct.BrowserFilter{
				Field: "name",
				Comp:  "==",
				Value: "John Doe",
			},
			queryFilters: []string{},
			want:         []string{"and (name like '%John Doe%')"},
		},
		{
			name: "customer_events-start_time",
			view: "customer_events",
			filter: ct.BrowserFilter{
				Field: "start_time",
				Comp:  "==",
				Value: "2021-01-01",
			},
			queryFilters: []string{},
			want:         []string{"and (start_time = '2021-01-01')"},
		},
		{
			name: "customer_events-default",
			view: "customer_events",
			filter: ct.BrowserFilter{
				Field: "name",
				Comp:  "==",
				Value: "John Doe",
			},
			queryFilters: []string{},
			want:         []string{"and (CAST(name as CHAR(255)) like '%John Doe%')"},
		},
		{
			name: "customer_map-default",
			view: "customer_map",
			filter: ct.BrowserFilter{
				Field: "name",
				Comp:  "==",
				Value: "John Doe",
			},
			queryFilters: []string{},
			want:         []string{"and (name like '%John Doe%')"},
		},
		{
			name: "product_simple",
			view: "product_simple",
			filter: ct.BrowserFilter{
				Field: "code",
				Comp:  "==",
				Value: "123",
			},
			queryFilters: []string{},
			want:         []string{"and (CAST(code as CHAR(255)) like '%123%')"},
		},
		{
			name: "product_inactive",
			view: "product",
			filter: ct.BrowserFilter{
				Field: "inactive",
				Comp:  "==",
				Value: true,
			},
			queryFilters: []string{},
			want:         []string{"and (inactive = true)"},
		},
		{
			name: "product_code",
			view: "product",
			filter: ct.BrowserFilter{
				Field: "code",
				Comp:  "==",
				Value: "123",
			},
			queryFilters: []string{},
			want:         []string{"and (CAST(code as CHAR(255)) like '%123%')"},
		},
		{
			name: "product_id",
			view: "product",
			filter: ct.BrowserFilter{
				Field: "id",
				Comp:  "==",
				Value: "456",
			},
			queryFilters: []string{},
			want:         []string{"and (id = 456)"},
		},
		{
			name: "product_default",
			view: "product",
			filter: ct.BrowserFilter{
				Field: "default",
			},
			queryFilters: []string{},
			want:         []string{},
		},
		{
			name: "product_map",
			view: "product_map",
			filter: ct.BrowserFilter{
				Field: "demo_number",
				Comp:  "==",
				Value: "123",
			},
			queryFilters: []string{},
			want:         []string{"and (demo_number like '%123%')"},
		},
		{
			name: "product_events",
			view: "product_events",
			filter: ct.BrowserFilter{
				Field: "start_time",
				Comp:  "==",
				Value: "2021-01-01",
			},
			queryFilters: []string{},
			want:         []string{"and (start_time = '2021-01-01')"},
		},
		{
			name: "product_events_default",
			view: "product_events",
			filter: ct.BrowserFilter{
				Field: "name",
				Comp:  "==",
				Value: "John Doe",
			},
			queryFilters: []string{},
			want:         []string{"and (CAST(name as CHAR(255)) like '%John Doe%')"},
		},
		{
			name: "product_prices_valid_from",
			view: "product_prices",
			filter: ct.BrowserFilter{
				Field: "valid_from",
				Comp:  "==",
				Value: "2021-01-01",
			},
			queryFilters: []string{},
			want:         []string{"and (valid_from = '2021-01-01')"},
		},
		{
			name: "product_prices_id",
			view: "product_prices",
			filter: ct.BrowserFilter{
				Field: "id",
				Comp:  "==",
				Value: "456",
			},
			queryFilters: []string{},
			want:         []string{"and (id = 456)"},
		},
		{
			name: "product_default",
			view: "product_prices",
			filter: ct.BrowserFilter{
				Field: "default",
				Comp:  "==",
				Value: "123",
			},
			queryFilters: []string{},
			want:         []string{"and (CAST(default as CHAR(255)) like '%123%')"},
		},
		{
			name: "tool_simple",
			view: "tool_simple",
			filter: ct.BrowserFilter{
				Field: "code",
				Comp:  "==",
				Value: "123",
			},
			queryFilters: []string{},
			want:         []string{"and (CAST(code as CHAR(255)) like '%123%')"},
		},
		{
			name: "tool_inactive",
			view: "tool",
			filter: ct.BrowserFilter{
				Field: "inactive",
				Comp:  "==",
				Value: true,
			},
			queryFilters: []string{},
			want:         []string{"and (inactive = true)"},
		},
		{
			name: "tool_code",
			view: "tool",
			filter: ct.BrowserFilter{
				Field: "code",
				Comp:  "==",
				Value: "123",
			},
			queryFilters: []string{},
			want:         []string{"and (CAST(code as CHAR(255)) like '%123%')"},
		},
		{
			name: "tool_id",
			view: "tool",
			filter: ct.BrowserFilter{
				Field: "id",
				Comp:  "==",
				Value: "456",
			},
			queryFilters: []string{},
			want:         []string{"and (id = 456)"},
		},
		{
			name: "tool_default",
			view: "tool",
			filter: ct.BrowserFilter{
				Field: "default",
				Comp:  "==",
				Value: "123",
			},
			queryFilters: []string{},
			want:         []string{},
		},
		{
			name: "tool_map",
			view: "tool_map",
			filter: ct.BrowserFilter{
				Field: "demo_number",
				Comp:  "==",
				Value: "123",
			},
			queryFilters: []string{},
			want:         []string{"and (demo_number like '%123%')"},
		},
		{
			name: "tool_events",
			view: "tool_events",
			filter: ct.BrowserFilter{
				Field: "start_time",
				Comp:  "==",
				Value: "2021-01-01",
			},
			queryFilters: []string{},
			want:         []string{"and (start_time = '2021-01-01')"},
		},
		{
			name: "tool_events_default",
			view: "tool_events",
			filter: ct.BrowserFilter{
				Field: "name",
				Comp:  "==",
				Value: "John Doe",
			},
			queryFilters: []string{},
			want:         []string{"and (CAST(name as CHAR(255)) like '%John Doe%')"},
		},
		{
			name: "project_simple",
			view: "project_simple",
			filter: ct.BrowserFilter{
				Field: "code",
				Comp:  "==",
				Value: "123",
			},
			queryFilters: []string{},
			want:         []string{"and (CAST(code as CHAR(255)) like '%123%')"},
		},
		{
			name: "project_inactive",
			view: "project",
			filter: ct.BrowserFilter{
				Field: "inactive",
				Comp:  "==",
				Value: true,
			},
			queryFilters: []string{},
			want:         []string{"and (inactive = true)"},
		},
		{
			name: "project_start_date",
			view: "project",
			filter: ct.BrowserFilter{
				Field: "start_date",
				Comp:  "==",
				Value: "2021-01-01",
			},
			queryFilters: []string{},
			want:         []string{"and (start_date = '2021-01-01')"},
		},
		{
			name: "project_code",
			view: "project",
			filter: ct.BrowserFilter{
				Field: "code",
				Comp:  "==",
				Value: "123",
			},
			queryFilters: []string{},
			want:         []string{"and (CAST(code as CHAR(255)) like '%123%')"},
		},
		{
			name: "project_id",
			view: "project",
			filter: ct.BrowserFilter{
				Field: "id",
				Comp:  "==",
				Value: "456",
			},
			queryFilters: []string{},
			want:         []string{"and (id = 456)"},
		},
		{
			name: "project_default",
			view: "project",
			filter: ct.BrowserFilter{
				Field: "default",
			},
			queryFilters: []string{},
			want:         []string{},
		},
		{
			name: "project_map",
			view: "project_map",
			filter: ct.BrowserFilter{
				Field: "demo_number",
				Comp:  "==",
				Value: "123",
			},
			queryFilters: []string{},
			want:         []string{"and (demo_number like '%123%')"},
		},
		{
			name: "project_addresses",
			view: "project_addresses",
			filter: ct.BrowserFilter{
				Field: "city",
				Comp:  "==",
				Value: "123",
			},
			queryFilters: []string{},
			want:         []string{"and (city like '%123%')"},
		},
		{
			name: "project_contacts",
			view: "project_contacts",
			filter: ct.BrowserFilter{
				Field: "name",
				Comp:  "==",
				Value: "John Doe",
			},
			queryFilters: []string{},
			want:         []string{"and (name like '%John Doe%')"},
		},
		{
			name: "project_events_start_time",
			view: "project_events",
			filter: ct.BrowserFilter{
				Field: "start_time",
				Comp:  "==",
				Value: "2021-01-01",
			},
			queryFilters: []string{},
			want:         []string{"and (start_time = '2021-01-01')"},
		},
		{
			name: "project_events_default",
			view: "project_events",
			filter: ct.BrowserFilter{
				Field: "name",
				Comp:  "==",
				Value: "John Doe",
			},
			queryFilters: []string{},
			want:         []string{"and (CAST(name as CHAR(255)) like '%John Doe%')"},
		},
		{
			name: "employee_simple",
			view: "employee_simple",
			filter: ct.BrowserFilter{
				Field: "code",
				Comp:  "==",
				Value: "123",
			},
			queryFilters: []string{},
			want:         []string{"and (CAST(code as CHAR(255)) like '%123%')"},
		},
		{
			name: "employee_inactive",
			view: "employee",
			filter: ct.BrowserFilter{
				Field: "inactive",
				Comp:  "==",
				Value: true,
			},
			queryFilters: []string{},
			want:         []string{"and (inactive = true)"},
		},
		{
			name: "employee_start_date",
			view: "employee",
			filter: ct.BrowserFilter{
				Field: "start_date",
				Comp:  "==",
				Value: "2021-01-01",
			},
			queryFilters: []string{},
			want:         []string{"and (start_date = '2021-01-01')"},
		},
		{
			name: "employee_code",
			view: "employee",
			filter: ct.BrowserFilter{
				Field: "code",
				Comp:  "==",
				Value: "123",
			},
			queryFilters: []string{},
			want:         []string{"and (CAST(code as CHAR(255)) like '%123%')"},
		},
		{
			name: "employee_id",
			view: "employee",
			filter: ct.BrowserFilter{
				Field: "id",
				Comp:  "==",
				Value: "456",
			},
			queryFilters: []string{},
			want:         []string{"and (id = 456)"},
		},
		{
			name: "employee_default",
			view: "employee",
			filter: ct.BrowserFilter{
				Field: "default",
			},
			queryFilters: []string{},
			want:         []string{},
		},
		{
			name: "employee_map",
			view: "employee_map",
			filter: ct.BrowserFilter{
				Field: "demo_number",
				Comp:  "==",
				Value: "123",
			},
			queryFilters: []string{},
			want:         []string{"and (demo_number like '%123%')"},
		},
		{
			name: "employee_events",
			view: "employee_events",
			filter: ct.BrowserFilter{
				Field: "start_time",
				Comp:  "==",
				Value: "2021-01-01",
			},
			queryFilters: []string{},
			want:         []string{"and (start_time = '2021-01-01')"},
		},
		{
			name: "employee_events_default",
			view: "employee_events",
			filter: ct.BrowserFilter{
				Field: "name",
				Comp:  "==",
				Value: "John Doe",
			},
			queryFilters: []string{},
			want:         []string{"and (CAST(name as CHAR(255)) like '%John Doe%')"},
		},
		{
			name: "place_simple",
			view: "place_simple",
			filter: ct.BrowserFilter{
				Field: "code",
				Comp:  "==",
				Value: "123",
			},
			want: []string{"and (CAST(code as CHAR(255)) like '%123%')"},
		},
		{
			name: "place_inactive",
			view: "place",
			filter: ct.BrowserFilter{
				Field: "inactive",
				Comp:  "==",
				Value: true,
			},
			queryFilters: []string{},
			want:         []string{"and (inactive = true)"},
		},
		{
			name: "place_code",
			view: "place",
			filter: ct.BrowserFilter{
				Field: "code",
				Comp:  "==",
				Value: "123",
			},
			want: []string{"and (code like '%123%')"},
		},
		{
			name: "place_id",
			view: "place",
			filter: ct.BrowserFilter{
				Field: "id",
				Comp:  "==",
				Value: "456",
			},
			queryFilters: []string{},
			want:         []string{"and (id = 456)"},
		},
		{
			name: "place_default",
			view: "place",
			filter: ct.BrowserFilter{
				Field: "default",
			},
			queryFilters: []string{},
			want:         []string{},
		},
		{
			name: "place_map",
			view: "place_map",
			filter: ct.BrowserFilter{
				Field: "demo_number",
				Comp:  "==",
				Value: "123",
			},
			queryFilters: []string{},
			want:         []string{"and (demo_number like '%123%')"},
		},
		{
			name: "place_contacts",
			view: "place_contacts",
			filter: ct.BrowserFilter{
				Field: "name",
				Comp:  "==",
				Value: "John Doe",
			},
			queryFilters: []string{},
			want:         []string{"and (name like '%John Doe%')"},
		},
		{
			name: "transitem_simple",
			view: "transitem_simple",
			filter: ct.BrowserFilter{
				Field: "code",
				Comp:  "==",
				Value: "123",
			},
			queryFilters: []string{},
			want:         []string{"and (CAST(code as CHAR(255)) like '%123%')"},
		},
		{
			name: "invoice_simple",
			view: "invoice_simple",
			filter: ct.BrowserFilter{
				Field: "code",
				Comp:  "==",
				Value: "123",
			},
			queryFilters: []string{},
			want:         []string{"and (CAST(code as CHAR(255)) like '%123%')"},
		},
		{
			name: "transitem_closed",
			view: "transitem",
			filter: ct.BrowserFilter{
				Field: "closed",
				Comp:  "==",
				Value: true,
			},
			queryFilters: []string{},
			want:         []string{"and (closed = true)"},
		},
		{
			name: "transitem_trans_date",
			view: "transitem",
			filter: ct.BrowserFilter{
				Field: "trans_date",
				Comp:  "==",
				Value: "2021-01-01",
			},
			queryFilters: []string{},
			want:         []string{"and (trans_date = '2021-01-01')"},
		},
		{
			name: "transitem_id",
			view: "transitem",
			filter: ct.BrowserFilter{
				Field: "id",
				Comp:  "==",
				Value: "456",
			},
			queryFilters: []string{},
			want:         []string{"and (id = 456)"},
		},
		{
			name: "transitem_customer_name",
			view: "transitem",
			filter: ct.BrowserFilter{
				Field: "customer_name",
				Comp:  "==",
				Value: "John Doe",
			},
			queryFilters: []string{},
			want:         []string{"and (CAST(customer_name as CHAR(255)) like '%John Doe%')"},
		},
		{
			name: "transitem_default",
			view: "transitem",
			filter: ct.BrowserFilter{
				Field: "default",
			},
			queryFilters: []string{},
			want:         []string{"and (CAST(t.default as CHAR(255))  '%%')"},
		},
		{
			name: "transitem_map",
			view: "transitem_map",
			filter: ct.BrowserFilter{
				Field: "demo_number",
				Comp:  "==",
				Value: "123",
			},
			queryFilters: []string{},
			want:         []string{"and (demo_number like '%123%')"},
		},
		{
			name: "transitem_item",
			view: "transitem_item",
			filter: ct.BrowserFilter{
				Field: "qty",
				Comp:  "==",
				Value: "123",
			},
			queryFilters: []string{},
			want:         []string{"and (qty = 123)"},
		},
		{
			name: "transitem_item_deposit",
			view: "transitem_item",
			filter: ct.BrowserFilter{
				Field: "deposit",
				Comp:  "==",
				Value: true,
			},
			queryFilters: []string{},
			want:         []string{"and (deposit = true)"},
		},
		{
			name: "transitem_item_default",
			view: "transitem_item",
			filter: ct.BrowserFilter{
				Field: "default",
			},
			queryFilters: []string{},
			want:         []string{"and (CAST(default as CHAR(255))  '%%')"},
		},
		{
			name: "transpayment_simple",
			view: "transpayment_simple",
			filter: ct.BrowserFilter{
				Field: "default",
			},
			queryFilters: []string{},
			want:         []string{"and (CAST(default as CHAR(255))  '%%')"},
		},
		{
			name: "transpayment_closed",
			view: "transpayment",
			filter: ct.BrowserFilter{
				Field: "closed",
				Comp:  "==",
				Value: true,
			},
			want: []string{"and (closed = true)"},
		},
		{
			name: "transpayment_trans_date",
			view: "transpayment",
			filter: ct.BrowserFilter{
				Field: "trans_date",
				Comp:  "==",
				Value: "2021-01-01",
			},
			want: []string{"and (trans_date = '2021-01-01')"},
		},
		{
			name: "transpayment_id",
			view: "transpayment",
			filter: ct.BrowserFilter{
				Field: "id",
				Comp:  "==",
				Value: "456",
			},
			want: []string{"and (id = 456)"},
		},
		{
			name: "transpayment_code",
			view: "transpayment",
			filter: ct.BrowserFilter{
				Field: "code",
				Comp:  "==",
				Value: "123",
			},
			want: []string{"and (CAST(t.code as CHAR(255)) like '%123%')"},
		},
		{
			name: "transpayment_default",
			view: "transpayment",
			filter: ct.BrowserFilter{
				Field: "default",
			},
			queryFilters: []string{},
			want:         []string{"and (CAST(default as CHAR(255))  '%%')"},
		},
		{
			name: "transpayment_map",
			view: "transpayment_map",
			filter: ct.BrowserFilter{
				Field: "demo_number",
				Comp:  "==",
				Value: "123",
			},
			queryFilters: []string{},
			want:         []string{"and (demo_number like '%123%')"},
		},
		{
			name: "transpayment_invoice",
			view: "transpayment_invoice",
			filter: ct.BrowserFilter{
				Field: "invoice_amount",
				Comp:  "==",
				Value: "123",
			},
			queryFilters: []string{},
			want:         []string{"and (invoice_amount = 123)"},
		},
		{
			name: "transpayment_invoice_id",
			view: "transpayment_invoice",
			filter: ct.BrowserFilter{
				Field: "id",
				Comp:  "==",
				Value: "456",
			},
			want: []string{"and (id = 456)"},
		},
		{
			name: "transpayment_invoice_trans_date",
			view: "transpayment_invoice",
			filter: ct.BrowserFilter{
				Field: "trans_date",
				Comp:  "==",
				Value: "2021-01-01",
			},
			queryFilters: []string{},
			want:         []string{"and (trans_date = '2021-01-01')"},
		},
		{
			name: "transpayment_invoice_currency_code",
			view: "transpayment_invoice",
			filter: ct.BrowserFilter{
				Field: "currency_code",
				Comp:  "==",
				Value: "USD",
			},
			want: []string{"and (CAST(currency_code as CHAR(255)) like '%USD%')"},
		},
		{
			name: "transpayment_invoice_default",
			view: "transpayment_invoice",
			filter: ct.BrowserFilter{
				Field: "default",
				Comp:  "==",
				Value: "abc",
			},
			queryFilters: []string{},
			want:         []string{"and (CAST(default as CHAR(255)) like '%abc%')"},
		},
		{
			name: "transmovement_simple",
			view: "transmovement_simple",
			filter: ct.BrowserFilter{
				Field: "code",
				Comp:  "==",
				Value: "123",
			},
			queryFilters: []string{},
			want:         []string{"and (CAST(code as CHAR(255)) like '%123%')"},
		},
		{
			name: "transmovement_stock_posdate",
			view: "transmovement_stock",
			filter: ct.BrowserFilter{
				Field: "posdate",
				Comp:  "==",
				Value: "2021-01-01",
			},
			queryFilters: []string{},
			want:         []string{"and (posdate = '2021-01-01')"},
		},
		{
			name: "transmovement_stock_id",
			view: "transmovement_stock",
			filter: ct.BrowserFilter{
				Field: "id",
				Comp:  "==",
				Value: "456",
			},
			queryFilters: []string{},
			want:         []string{"and (id = 456)"},
		},
		{
			name: "transmovement_stock_qty",
			view: "transmovement_stock",
			filter: ct.BrowserFilter{
				Field: "qty",
				Comp:  "==",
				Value: "123",
			},
			queryFilters: []string{},
			want:         []string{"and (qty = 123)"},
		},
		{
			name: "transmovement_stock_default",
			view: "transmovement_stock",
			filter: ct.BrowserFilter{
				Field: "default",
				Comp:  "==",
				Value: "abc",
			},
			queryFilters: []string{},
			want:         []string{"and (CAST(default as CHAR(255)) like '%abc%')"},
		},
		{
			name: "transmovement_id",
			view: "transmovement",
			filter: ct.BrowserFilter{
				Field: "id",
				Comp:  "==",
				Value: "456",
			},
			queryFilters: []string{},
			want:         []string{"and (id = 456)"},
		},
		{
			name: "transmovement_shipping_date",
			view: "transmovement",
			filter: ct.BrowserFilter{
				Field: "shipping_date",
				Comp:  "==",
				Value: "2021-01-01",
			},
			queryFilters: []string{},
			want:         []string{"and (shipping_date = '2021-01-01')"},
		},
		{
			name: "transmovement_default",
			view: "transmovement",
			filter: ct.BrowserFilter{
				Field: "default",
				Comp:  "==",
				Value: "abc",
			},
			queryFilters: []string{},
			want:         []string{"and (CAST(default as CHAR(255)) like '%abc%')"},
		},
		{
			name: "transmovement_waybill_id",
			view: "transmovement_waybill",
			filter: ct.BrowserFilter{
				Field: "id",
				Comp:  "==",
				Value: "456",
			},
			queryFilters: []string{},
			want:         []string{"and (id = 456)"},
		},
		{
			name: "transmovement_waybill_shipping_time",
			view: "transmovement_waybill",
			filter: ct.BrowserFilter{
				Field: "shipping_time",
				Comp:  "==",
				Value: "2021-01-01",
			},
			queryFilters: []string{},
			want:         []string{"and (shipping_time = '2021-01-01')"},
		},
		{
			name: "transmovement_waybill_default",
			view: "transmovement_waybill",
			filter: ct.BrowserFilter{
				Field: "default",
				Comp:  "==",
				Value: "abc",
			},
			queryFilters: []string{},
			want:         []string{"and (CAST(default as CHAR(255)) like '%abc%')"},
		},
		{
			name: "transmovement_formula_id",
			view: "transmovement_formula",
			filter: ct.BrowserFilter{
				Field: "id",
				Comp:  "==",
				Value: "456",
			},
			queryFilters: []string{},
			want:         []string{"and (id = 456)"},
		},
		{
			name: "transmovement_formula_shared",
			view: "transmovement_formula",
			filter: ct.BrowserFilter{
				Field: "shared",
				Comp:  "==",
				Value: true,
			},
			queryFilters: []string{},
			want:         []string{"and (shared = true)"},
		},
		{
			name: "transmovement_formula_default",
			view: "transmovement_formula",
			filter: ct.BrowserFilter{
				Field: "default",
				Comp:  "==",
				Value: "abc",
			},
			queryFilters: []string{},
			want:         []string{"and (CAST(default as CHAR(255)) like '%abc%')"},
		},
		{
			name: "transmovement_map",
			view: "transmovement_map",
			filter: ct.BrowserFilter{
				Field: "demo_number",
				Comp:  "==",
				Value: "123",
			},
			queryFilters: []string{},
			want:         []string{"and (demo_number like '%123%')"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var s component.SearchConfig
			got := s.Filter(tt.view, tt.filter, tt.queryFilters)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("component.SearchFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchConfig_Query(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		key    string
		params cu.IM
	}{
		{
			name: "transitem_simple",
			key:  "transitem_simple",
			params: cu.IM{
				"view": "transitem_simple",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "invoice_simple",
			key:  "invoice_simple",
			params: cu.IM{
				"view": "invoice_simple",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "transitem",
			key:  "transitem",
			params: cu.IM{
				"view": "transitem",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "transitem_map",
			key:  "transitem_map",
			params: cu.IM{
				"view": "transitem_map",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "transitem_item",
			key:  "transitem_item",
			params: cu.IM{
				"view": "transitem_item",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "transpayment_simple",
			key:  "transpayment_simple",
			params: cu.IM{
				"view": "transpayment_simple",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "transpayment",
			key:  "transpayment",
			params: cu.IM{
				"view": "transpayment",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "transpayment_map",
			key:  "transpayment_map",
			params: cu.IM{
				"view": "transpayment_map",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "transpayment_invoice",
			key:  "transpayment_invoice",
			params: cu.IM{
				"view": "transpayment_invoice",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
			},
		},
		{
			name: "transmovement_simple",
			key:  "transmovement_simple",
			params: cu.IM{
				"view": "transmovement_simple",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "transmovement_stock",
			key:  "transmovement_stock",
			params: cu.IM{
				"view": "transmovement_stock",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "transmovement",
			key:  "transmovement",
			params: cu.IM{
				"view": "transmovement",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "transmovement_waybill",
			key:  "transmovement_waybill",
			params: cu.IM{
				"view": "transmovement_waybill",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "transmovement_formula",
			key:  "transmovement_formula",
			params: cu.IM{
				"view": "transmovement_formula",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "transmovement_map",
			key:  "transmovement_map",
			params: cu.IM{
				"view": "transmovement_map",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "customer_simple",
			key:  "customer_simple",
			params: cu.IM{
				"view": "customer_simple",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "customer",
			key:  "customer",
			params: cu.IM{
				"view": "customer",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "customer_map",
			key:  "customer_map",
			params: cu.IM{
				"view": "customer_map",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "customer_addresses",
			key:  "customer_addresses",
			params: cu.IM{
				"view": "customer_addresses",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "customer_contacts",
			key:  "customer_contacts",
			params: cu.IM{
				"view": "customer_contacts",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "customer_events",
			key:  "customer_events",
			params: cu.IM{
				"view": "customer_events",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "product_simple",
			key:  "product_simple",
			params: cu.IM{
				"view": "product_simple",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "product",
			key:  "product",
			params: cu.IM{
				"view": "product",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "product_map",
			key:  "product_map",
			params: cu.IM{
				"view": "product_map",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "product_events",
			key:  "product_events",
			params: cu.IM{
				"view": "product_events",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "product_prices",
			key:  "product_prices",
			params: cu.IM{
				"view": "product_prices",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "tool_simple",
			key:  "tool_simple",
			params: cu.IM{
				"view": "tool_simple",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "tool",
			key:  "tool",
			params: cu.IM{
				"view": "tool",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "tool_map",
			key:  "tool_map",
			params: cu.IM{
				"view": "tool_map",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "tool_events",
			key:  "tool_events",
			params: cu.IM{
				"view": "tool_events",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "project_simple",
			key:  "project_simple",
			params: cu.IM{
				"view": "project_simple",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "project",
			key:  "project",
			params: cu.IM{
				"view": "project",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "project_map",
			key:  "project_map",
			params: cu.IM{
				"view": "project_map",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "project_addresses",
			key:  "project_addresses",
			params: cu.IM{
				"view": "project_addresses",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "project_contacts",
			key:  "project_contacts",
			params: cu.IM{
				"view": "project_contacts",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "project_events",
			key:  "project_events",
			params: cu.IM{
				"view": "project_events",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "employee_simple",
			key:  "employee_simple",
			params: cu.IM{
				"view": "employee_simple",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "employee",
			key:  "employee",
			params: cu.IM{
				"view": "employee",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "employee_map",
			key:  "employee_map",
			params: cu.IM{
				"view": "employee_map",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "employee_events",
			key:  "employee_events",
			params: cu.IM{
				"view": "employee_events",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "place_simple",
			key:  "place_simple",
			params: cu.IM{
				"view": "place_simple",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "place",
			key:  "place",
			params: cu.IM{
				"view": "place",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "place_map",
			key:  "place_map",
			params: cu.IM{
				"view": "place_map",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "place_contacts",
			key:  "place_contacts",
			params: cu.IM{
				"view": "place_contacts",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var s component.SearchConfig
			s.Query(tt.key, tt.params)
		})
	}
}
