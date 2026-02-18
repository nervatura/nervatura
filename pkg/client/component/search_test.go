package component_test

import (
	"html/template"
	"reflect"
	"testing"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/client/component"
	md "github.com/nervatura/nervatura/v6/pkg/model"
)

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
			name: "product_components",
			key:  "product_components",
			params: cu.IM{
				"view": "product_components",
				"query": md.Query{
					Filters: []md.Filter{},
				},
				"filters": []ct.BrowserFilter{},
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
		{
			name: "place_events",
			key:  "place_events",
			params: cu.IM{
				"view": "place_events",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "office_template_editor",
			key:  "office_template_editor",
			params: cu.IM{
				"view": "office_template_editor",
				"query": md.Query{
					Filters: []md.Filter{
						{Field: "id", Comp: "==", Value: 1},
					},
				},
				"filters": []ct.BrowserFilter{{Field: "id", Comp: "==", Value: 1}},
			},
		},
		{
			name: "office_rate",
			key:  "office_rate",
			params: cu.IM{
				"view": "office_rate",
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

func TestSearchConfig_CustomTemplateCell(t *testing.T) {
	type args struct {
		sessionID string
	}
	tests := []struct {
		name string
		s    *component.SearchConfig
		args args
		want func(row cu.IM, col ct.TableColumn, value any, rowIndex int64) template.HTML
	}{
		{
			name: "custom_template_cell",
			s:    &component.SearchConfig{},
			args: args{
				sessionID: "123",
			},
			want: func(row cu.IM, col ct.TableColumn, value any, rowIndex int64) template.HTML {
				return template.HTML("test")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &component.SearchConfig{}
			got := s.CustomTemplateCell(tt.args.sessionID)
			got(cu.IM{"code": "123"}, ct.TableColumn{}, "123", 0, nil)
		})
	}
}

func TestSearchConfig_Filter(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		view    string
		filter  ct.BrowserFilter
		filters []md.Filter
		want    []md.Filter
	}{
		{
			name: "date-filter",
			view: "customer_events",
			filter: ct.BrowserFilter{
				Field: "end_time",
				Comp:  "==",
				Value: "2026-01-01",
			},
			filters: []md.Filter{},
			want: []md.Filter{
				{Field: "start_time", Comp: "=", Value: "2026-01-01"},
			},
		},
		{
			name: "number-filter",
			view: "transitem",
			filter: ct.BrowserFilter{
				Field: "amount",
				Comp:  "==",
				Value: 100,
			},
			filters: []md.Filter{},
			want: []md.Filter{
				{Field: "amount", Comp: "=", Value: float64(100)},
			},
		},
		{
			name: "integer-filter",
			view: "customer",
			filter: ct.BrowserFilter{
				Field: "discount",
				Comp:  "==",
				Value: 100,
			},
			filters: []md.Filter{},
			want: []md.Filter{
				{Field: "discount", Comp: "=", Value: int64(100)},
			},
		},
		{
			name: "string-filter",
			view: "transitem",
			filter: ct.BrowserFilter{
				Field: "customer_name",
				Comp:  "==",
				Value: "John Doe",
			},
			filters: []md.Filter{},
			want: []md.Filter{
				{Field: "CAST(customer_name as CHAR(255))", Comp: "like", Value: "%John Doe%"},
			},
		},
		{
			name: "boolean-filter",
			view: "transitem",
			filter: ct.BrowserFilter{
				Field: "paid",
				Comp:  "==",
				Value: true,
			},
			filters: []md.Filter{},
			want: []md.Filter{
				{Field: "paid", Comp: "=", Value: true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var s component.SearchConfig
			got := s.Filter(tt.view, tt.filter, tt.filters)
			// TODO: update the condition below to compare got with tt.want.
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
