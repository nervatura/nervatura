package component

import (
	"reflect"
	"testing"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
)

func TestDefaultMapValue(t *testing.T) {
	type args struct {
		ftype string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "number",
			args: args{
				ftype: "FIELD_NUMBER",
			},
			want: 0,
		},
		{
			name: "invalid",
			args: args{
				ftype: "invalid",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefaultMapValue(tt.args.ftype); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultMapValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapTableRows(t *testing.T) {
	type args struct {
		mapData   cu.IM
		configMap []cu.IM
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "number",
			args: args{
				mapData: cu.IM{
					"demo_number": 123,
				},
				configMap: []cu.IM{
					{
						"field_name":  "demo_number",
						"field_type":  "FIELD_NUMBER",
						"description": "Demo Number",
						"filter":      []interface{}{"FILTER_CUSTOMER"},
						"tags":        []interface{}{},
					},
				},
			},
		},
		{
			name: "enum",
			args: args{
				mapData: cu.IM{
					"demo_enum": "enum_value",
				},
				configMap: []cu.IM{
					{
						"field_name":  "demo_enum",
						"field_type":  "FIELD_ENUM",
						"description": "Demo Enum",
						"filter":      []interface{}{"FILTER_CUSTOMER"},
						"tags":        []interface{}{"value1", "value2"},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapTableRows(tt.args.mapData, tt.args.configMap)
		})
	}
}

func TestClientComponent_Menu(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		labels cu.SM
		config cu.IM
	}{
		{
			name:   "default",
			labels: cu.SM{},
			config: cu.IM{
				"login_disabled": false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := NewClientComponent()
			cc.Menu(tt.labels, tt.config)
		})
	}
}

func TestClientComponent_Login(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		labels cu.SM
		config cu.IM
	}{
		{
			name:   "default",
			labels: cu.SM{},
			config: cu.IM{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := NewClientComponent()
			cc.Login(tt.labels, tt.config)
		})
	}
}

func TestClientComponent_SideBar(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		moduleKey string
		labels    cu.SM
		data      cu.IM
	}{
		{
			name:      "search",
			moduleKey: "search",
			labels:    cu.SM{},
			data: cu.IM{
				"config": cu.IM{
					"ticket": ct.Ticket{
						SessionID: "123",
					},
				},
			},
		},
		{
			name:      "browser",
			moduleKey: "browser",
			labels:    cu.SM{},
			data:      cu.IM{},
		},
		{
			name:      "customer_new",
			moduleKey: "customer",
			labels:    cu.SM{},
			data: cu.IM{
				"customer": cu.IM{
					"id": 0,
				},
			},
		},
		{
			name:      "customer_edit",
			moduleKey: "customer",
			labels:    cu.SM{},
			data: cu.IM{
				"customer": cu.IM{
					"id": 1,
				},
			},
		},
		{
			name:      "customer_inactive",
			moduleKey: "customer",
			labels:    cu.SM{},
			data: cu.IM{
				"customer": cu.IM{
					"id":       1,
					"inactive": true,
				},
			},
		},
		{
			name:      "product_new",
			moduleKey: "product",
			labels:    cu.SM{},
			data: cu.IM{
				"product": cu.IM{
					"id": 0,
				},
			},
		},
		{
			name:      "product_edit",
			moduleKey: "product",
			labels:    cu.SM{},
			data: cu.IM{
				"product": cu.IM{
					"id": 1,
				},
			},
		},
		{
			name:      "product_inactive",
			moduleKey: "product",
			labels:    cu.SM{},
			data: cu.IM{
				"product": cu.IM{
					"id":       1,
					"inactive": true,
				},
			},
		},
		{
			name:      "tool_new",
			moduleKey: "tool",
			labels:    cu.SM{},
			data: cu.IM{
				"tool": cu.IM{
					"id": 0,
				},
			},
		},
		{
			name:      "tool_edit",
			moduleKey: "tool",
			labels:    cu.SM{},
			data: cu.IM{
				"tool": cu.IM{
					"id": 1,
				},
			},
		},
		{
			name:      "product_inactive",
			moduleKey: "tool",
			labels:    cu.SM{},
			data: cu.IM{
				"tool": cu.IM{
					"id":       1,
					"inactive": true,
				},
			},
		},
		{
			name:      "setting",
			moduleKey: "setting",
			labels:    cu.SM{},
			data: cu.IM{
				"user": cu.IM{
					"user_group": md.UserGroupAdmin.String(),
				},
			},
		},
		{
			name:      "invalid",
			moduleKey: "invalid",
		},
		{
			name:      "project_new",
			moduleKey: "project",
			labels:    cu.SM{},
			data: cu.IM{
				"project": cu.IM{
					"id": 0,
				},
			},
		},
		{
			name:      "project_edit",
			moduleKey: "project",
			labels:    cu.SM{},
			data: cu.IM{
				"project": cu.IM{
					"id": 1,
				},
			},
		},
		{
			name:      "project_inactive",
			moduleKey: "project",
			labels:    cu.SM{},
			data: cu.IM{
				"project": cu.IM{
					"id":       1,
					"inactive": true,
				},
			},
		},
		{
			name:      "employee_new",
			moduleKey: "employee",
			labels:    cu.SM{},
			data: cu.IM{
				"employee": cu.IM{
					"id": 0,
				},
			},
		},
		{
			name:      "employee_edit",
			moduleKey: "employee",
			labels:    cu.SM{},
			data: cu.IM{
				"employee": cu.IM{
					"id": 1,
				},
			},
		},
		{
			name:      "employee_inactive",
			moduleKey: "employee",
			labels:    cu.SM{},
			data: cu.IM{
				"employee": cu.IM{
					"id":       1,
					"inactive": true,
				},
			},
		},
		{
			name:      "place_new",
			moduleKey: "place",
			labels:    cu.SM{},
			data: cu.IM{
				"place": cu.IM{
					"id": 0,
				},
			},
		},
		{
			name:      "place_edit",
			moduleKey: "place",
			labels:    cu.SM{},
			data: cu.IM{
				"place": cu.IM{
					"id": 1,
				},
			},
		},
		{
			name:      "place_inactive",
			moduleKey: "place",
			labels:    cu.SM{},
			data: cu.IM{
				"place": cu.IM{
					"id":       1,
					"inactive": true,
				},
			},
		},
		{
			name:      "trans_new",
			moduleKey: "trans",
			labels:    cu.SM{},
			data: cu.IM{
				"trans": cu.IM{
					"id": 0,
				},
			},
		},
		{
			name:      "trans_edit",
			moduleKey: "trans",
			labels:    cu.SM{},
			data: cu.IM{
				"trans": cu.IM{
					"id":         1,
					"trans_type": md.TransTypeInvoice.String(),
					"direction":  md.DirectionOut.String(),
					"trans_meta": cu.IM{
						"status": md.TransStatusNormal.String(),
					},
				},
			},
		},
		{
			name:      "trans_cash",
			moduleKey: "trans",
			labels:    cu.SM{},
			data: cu.IM{
				"trans": cu.IM{
					"id":         1,
					"trans_type": md.TransTypeCash.String(),
					"direction":  md.DirectionOut.String(),
					"trans_meta": cu.IM{
						"status": md.TransStatusNormal.String(),
					},
				},
			},
		},
		{
			name:      "trans_inventory",
			moduleKey: "trans",
			labels:    cu.SM{},
			data: cu.IM{
				"trans": cu.IM{
					"id":         1,
					"trans_type": md.TransTypeInventory.String(),
					"direction":  md.DirectionTransfer.String(),
					"trans_meta": cu.IM{
						"status": md.TransStatusNormal.String(),
					},
				},
			},
		},
		{
			name:      "trans_production",
			moduleKey: "trans",
			labels:    cu.SM{},
			data: cu.IM{
				"trans": cu.IM{
					"id":         1,
					"trans_type": md.TransTypeProduction.String(),
					"direction":  md.DirectionTransfer.String(),
					"trans_meta": cu.IM{
						"status": md.TransStatusNormal.String(),
					},
				},
			},
		},
		{
			name:      "trans_waybill",
			moduleKey: "trans",
			labels:    cu.SM{},
			data: cu.IM{
				"trans": cu.IM{
					"id":         1,
					"trans_type": md.TransTypeWaybill.String(),
					"direction":  md.DirectionTransfer.String(),
					"trans_meta": cu.IM{
						"status": md.TransStatusNormal.String(),
					},
				},
			},
		},
		{
			name:      "trans_formula",
			moduleKey: "trans",
			labels:    cu.SM{},
			data: cu.IM{
				"trans": cu.IM{
					"id":         1,
					"trans_type": md.TransTypeFormula.String(),
					"direction":  md.DirectionTransfer.String(),
					"trans_meta": cu.IM{
						"status": md.TransStatusNormal.String(),
					},
				},
			},
		},
		{
			name:      "trans_cash_deleted",
			moduleKey: "trans",
			labels:    cu.SM{},
			data: cu.IM{
				"trans": cu.IM{
					"id":         1,
					"trans_type": md.TransTypeCash.String(),
					"direction":  md.DirectionOut.String(),
					"trans_meta": cu.IM{
						"status": md.TransStatusDeleted.String(),
					},
				},
			},
		},
		{
			name:      "trans_closed",
			moduleKey: "trans",
			labels:    cu.SM{},
			data: cu.IM{
				"trans": cu.IM{
					"id":         1,
					"trans_type": md.TransTypeOrder.String(),
					"direction":  md.DirectionOut.String(),
					"trans_meta": cu.IM{
						"status": md.TransStatusNormal.String(),
						"closed": true,
					},
				},
			},
		},
		{
			name:      "trans_deleted",
			moduleKey: "trans",
			labels:    cu.SM{},
			data: cu.IM{
				"trans": cu.IM{
					"id":         1,
					"trans_type": md.TransTypeInvoice.String(),
					"direction":  md.DirectionOut.String(),
					"trans_meta": cu.IM{
						"status": md.TransStatusDeleted.String(),
					},
				},
			},
		},
		{
			name:      "shipping",
			moduleKey: "shipping",
			labels:    cu.SM{},
			data: cu.IM{
				"shipping": cu.IM{
					"id": 1,
				},
				"items": []cu.IM{
					{
						"id":  1,
						"qty": 100,
					},
				},
				"dirty": true,
			},
		},
		{
			name:      "rate_new",
			moduleKey: "rate",
			labels:    cu.SM{},
			data: cu.IM{
				"rate": cu.IM{
					"id": 0,
				},
			},
		},
		{
			name:      "rate_edit",
			moduleKey: "rate",
			labels:    cu.SM{},
			data: cu.IM{
				"rate": cu.IM{
					"id": 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := NewClientComponent()
			cc.SideBar(tt.moduleKey, tt.labels, tt.data)
		})
	}
}

func TestClientComponent_Search(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		viewName   string
		labels     cu.SM
		searchData cu.IM
	}{
		{
			name:     "customer",
			viewName: "customer",
			labels:   cu.SM{},
			searchData: cu.IM{
				"config": cu.IM{
					"ticket": ct.Ticket{
						SessionID: "123",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := NewClientComponent()
			cc.Search(tt.viewName, tt.labels, tt.searchData)
		})
	}
}

func TestClientComponent_Browser(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		viewName   string
		labels     cu.SM
		searchData cu.IM
	}{
		{
			name:     "customer",
			viewName: "customer",
			labels:   cu.SM{},
			searchData: cu.IM{
				"config": cu.IM{
					"ticket": ct.Ticket{
						SessionID: "123",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := NewClientComponent()
			cc.Browser(tt.viewName, tt.labels, tt.searchData)
		})
	}
}

func TestClientComponent_Editor(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		editorKey  string
		viewName   string
		labels     cu.SM
		editorData cu.IM
	}{
		{
			name:       "search",
			editorKey:  "search",
			viewName:   "search",
			labels:     cu.SM{},
			editorData: cu.IM{},
		},
		{
			name:       "customer",
			editorKey:  "customer",
			viewName:   "customer",
			labels:     cu.SM{},
			editorData: cu.IM{},
		},
		{
			name:       "product",
			editorKey:  "product",
			viewName:   "product",
			labels:     cu.SM{},
			editorData: cu.IM{},
		},
		{
			name:       "tool",
			editorKey:  "tool",
			viewName:   "tool",
			labels:     cu.SM{},
			editorData: cu.IM{},
		},
		{
			name:       "project",
			editorKey:  "project",
			viewName:   "project",
			labels:     cu.SM{},
			editorData: cu.IM{},
		},
		{
			name:       "setting",
			editorKey:  "setting",
			viewName:   "setting",
			labels:     cu.SM{},
			editorData: cu.IM{},
		},
		{
			name:       "employee",
			editorKey:  "employee",
			viewName:   "employee",
			labels:     cu.SM{},
			editorData: cu.IM{},
		},
		{
			name:       "place",
			editorKey:  "place",
			viewName:   "place",
			labels:     cu.SM{},
			editorData: cu.IM{},
		},
		{
			name:       "trans",
			editorKey:  "trans",
			viewName:   "trans",
			labels:     cu.SM{},
			editorData: cu.IM{},
		},
		{
			name:       "shipping",
			editorKey:  "shipping",
			viewName:   "shipping",
			labels:     cu.SM{},
			editorData: cu.IM{},
		},
		{
			name:      "rate",
			editorKey: "rate",
			viewName:  "rate",
			labels:    cu.SM{},
			editorData: cu.IM{
				"currencies": []cu.IM{
					{
						"code": "USD",
					},
				},
				"places": []cu.IM{
					{
						"code": "123",
					},
				},
			},
		},
		{
			name:      "invalid",
			editorKey: "invalid",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := NewClientComponent()
			cc.Editor(tt.editorKey, tt.viewName, tt.labels, tt.editorData)
		})
	}
}

func TestClientComponent_Form(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		editorKey string
		formKey   string
		labels    cu.SM
		data      cu.IM
	}{
		{
			name:      "customer_addresses",
			editorKey: "customer",
			formKey:   "addresses",
			labels:    cu.SM{},
			data:      cu.IM{},
		},
		{
			name:      "customer_contacts",
			editorKey: "customer",
			formKey:   "contacts",
			labels:    cu.SM{},
			data:      cu.IM{},
		},
		{
			name:      "customer_events",
			editorKey: "customer",
			formKey:   "events",
			labels:    cu.SM{},
			data:      cu.IM{},
		},
		{
			name:      "customer_customer",
			editorKey: "customer",
			formKey:   "customer",
			labels:    cu.SM{},
			data:      cu.IM{},
		},
		{
			name:      "product_events",
			editorKey: "product",
			formKey:   "events",
			labels:    cu.SM{},
			data:      cu.IM{},
		},
		{
			name:      "product_prices",
			editorKey: "product",
			formKey:   "prices",
			labels:    cu.SM{},
			data: cu.IM{
				"currencies": []cu.IM{
					{
						"currency_code": "USD",
					},
				},
			},
		},
		{
			name:      "product_components",
			editorKey: "product",
			formKey:   "components",
			labels:    cu.SM{},
			data: cu.IM{
				"components": []cu.IM{
					{
						"id": 1,
					},
				},
				"product_selector": []cu.IM{
					{
						"id": 1,
					},
				},
			},
		},
		{
			name:      "product_product",
			editorKey: "product",
			formKey:   "product",
			labels:    cu.SM{},
			data:      cu.IM{},
		},
		{
			name:      "tool_events",
			editorKey: "tool",
			formKey:   "events",
			labels:    cu.SM{},
			data:      cu.IM{},
		},
		{
			name:      "tool_product",
			editorKey: "tool",
			formKey:   "tool",
			labels:    cu.SM{},
			data:      cu.IM{},
		},
		{
			name:      "project_addresses",
			editorKey: "project",
			formKey:   "addresses",
			labels:    cu.SM{},
			data:      cu.IM{},
		},
		{
			name:      "project_contacts",
			editorKey: "project",
			formKey:   "contacts",
			labels:    cu.SM{},
			data:      cu.IM{},
		},
		{
			name:      "project_events",
			editorKey: "project",
			formKey:   "events",
			labels:    cu.SM{},
			data:      cu.IM{},
		},
		{
			name:      "project_project",
			editorKey: "project",
			formKey:   "project",
			labels:    cu.SM{},
			data:      cu.IM{},
		},
		{
			name:      "employee_events",
			editorKey: "employee",
			formKey:   "events",
			labels:    cu.SM{},
			data:      cu.IM{},
		},
		{
			name:      "employee_employee",
			editorKey: "employee",
			formKey:   "employee",
			labels:    cu.SM{},
			data:      cu.IM{},
		},
		{
			name:      "place_contacts",
			editorKey: "place",
			formKey:   "contacts",
			labels:    cu.SM{},
			data:      cu.IM{},
		},
		{
			name:      "place_events",
			editorKey: "place",
			formKey:   "events",
			labels:    cu.SM{},
			data:      cu.IM{},
		},
		{
			name:      "place_place",
			editorKey: "place",
			formKey:   "place",
			labels:    cu.SM{},
			data:      cu.IM{},
		},
		{
			name:      "trans_items",
			editorKey: "trans",
			formKey:   "items",
			labels:    cu.SM{},
			data: cu.IM{
				"tax_codes": []cu.IM{
					{
						"code": "123",
						"tax_meta": cu.IM{
							"description": "123",
							"rate_value":  123,
						},
					},
				},
			},
		},
		{
			name:      "trans_payments",
			editorKey: "trans",
			formKey:   "payments",
			labels:    cu.SM{},
			data: cu.IM{
				"payments": []cu.IM{
					{
						"id": 1,
						"payment_meta": cu.IM{
							"amount": 100,
						},
					},
					{
						"id": 2,
						"payment_meta": cu.IM{
							"amount": -100,
						},
					},
				},
			},
		},
		{
			name:      "trans_payment_link",
			editorKey: "trans",
			formKey:   "payment_link",
			labels:    cu.SM{},
			data: cu.IM{
				"payment_link": []cu.IM{
					{
						"id": 1,
						"link_meta": cu.IM{
							"amount": 100,
						},
					},
					{
						"id": 2,
						"link_meta": cu.IM{
							"amount": -100,
						},
					},
				},
			},
		},
		{
			name:      "trans_trans",
			editorKey: "trans",
			formKey:   "trans",
			labels:    cu.SM{},
			data:      cu.IM{},
		},
		{
			name:      "setting_config_map",
			editorKey: "setting",
			formKey:   "config_map",
			labels:    cu.SM{},
			data: cu.IM{
				"code": "123",
				"data": cu.IM{
					"field_name":  "123",
					"field_type":  "123",
					"description": "123",
				},
			},
		},
		{
			name:      "setting_auth",
			editorKey: "setting",
			formKey:   "auth",
			labels:    cu.SM{},
			data: cu.IM{
				"code": "123",
				"data": cu.IM{
					"user_group": md.UserGroupAdmin.String(),
					"user_name":  "admin",
					"disabled":   false,
					"auth_meta": cu.IM{
						"tags": []string{},
					},
					"auth_map": cu.IM{},
				},
			},
		},
		{
			name:      "setting_setting",
			editorKey: "setting",
			formKey:   "setting",
			labels:    cu.SM{},
			data:      cu.IM{},
		},
		{
			name:      "setting_shortcut",
			editorKey: "setting",
			formKey:   "shortcut",
			labels:    cu.SM{},
			data: cu.IM{
				"data": cu.IM{
					"shortcut_key": "shortcut_key",
					"description":  "description",
					"modul":        "modul",
					"method":       "method",
					"func_name":    "func_name",
					"address":      "address",
					"fields": []cu.IM{
						{
							"field_name":  "field_name",
							"field_type":  "FIELD_NUMBER",
							"description": "Demo Number",
						},
					},
				},
			},
		},
		{
			name:      "trans_movements",
			editorKey: "trans",
			formKey:   "movements",
			labels:    cu.SM{},
			data: cu.IM{
				"movements": []cu.IM{
					{
						"id": 1,
						"movement_meta": cu.IM{
							"qty": 100,
						},
					},
					{
						"id": 2,
						"movement_meta": cu.IM{
							"qty": -100,
						},
					},
				},
				"places": []cu.IM{
					{
						"code":       "1",
						"place_type": md.PlaceTypeWarehouse.String(),
					},
				},
			},
		},
		{
			name:      "shipping",
			editorKey: "shipping",
			formKey:   "shipping",
			labels:    cu.SM{},
			data: cu.IM{
				"shipping": cu.IM{
					"id": 1,
				},
			},
		},
		{
			name:      "rate",
			editorKey: "rate",
			formKey:   "rate",
			labels:    cu.SM{},
			data: cu.IM{
				"rate": cu.IM{
					"id": 1,
				},
			},
		},
		{
			name:      "invalid",
			editorKey: "invalid",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := NewClientComponent()
			cc.Form(tt.editorKey, tt.formKey, tt.labels, tt.data)
		})
	}
}

func TestClientComponent_Modal(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		formKey string
		labels  cu.SM
		data    cu.IM
	}{
		{
			name:    "info",
			formKey: "info",
			labels:  cu.SM{},
			data:    cu.IM{},
		},
		{
			name:    "warning",
			formKey: "warning",
			labels:  cu.SM{},
			data:    cu.IM{},
		},
		{
			name:    "input_string",
			formKey: "input_string",
			labels:  cu.SM{},
			data:    cu.IM{},
		},
		{
			name:    "report",
			formKey: "report",
			labels:  cu.SM{},
			data: cu.IM{
				"config_data": []cu.IM{
					{
						"config_code":  "orientation",
						"config_value": "landscape",
					},
				},
				"config_report": []cu.IM{
					{
						"report_key":  "report_key",
						"report_name": "report_name",
					},
				},
			},
		},
		{
			name:    "selector",
			formKey: "selector",
			labels:  cu.SM{},
			data:    cu.IM{},
		},
		{
			name:    "select",
			formKey: "select",
			labels:  cu.SM{},
			data: cu.IM{
				"info_label":   "label",
				"info_message": "message",
			},
		},
		{
			name:    "config_field",
			formKey: "config_field",
			labels:  cu.SM{},
			data: cu.IM{
				"field_name":  "field_name",
				"description": "description",
				"field_type":  "field_type",
				"order":       "order",
			},
		},
		{
			name:    "trans_create",
			formKey: "trans_create",
			labels:  cu.SM{},
			data: cu.IM{
				"trans_type":        md.TransTypeOrder.String(),
				"direction":         md.DirectionOut.String(),
				"status":            md.TransStatusNormal.String(),
				"show_delivery":     true,
				"trans_types":       []string{md.TransTypeOrder.String(), md.TransTypeInvoice.String(), md.TransTypeReceipt.String()},
				"create_trans_type": md.TransTypeInvoice.String(),
				"next":              "trans_create",
			},
		},
		{
			name:    "missing",
			formKey: "missing",
			labels:  cu.SM{},
			data:    cu.IM{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := NewClientComponent()
			cc.Modal(tt.formKey, tt.labels, tt.data)
		})
	}
}

func Test_moduleEditorView(t *testing.T) {
	type args struct {
		mKey   string
		labels cu.SM
		data   cu.IM
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "search",
			args: args{
				mKey:   "search",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "customer_new",
			args: args{
				mKey:   "customer",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "customer_edit",
			args: args{
				mKey:   "customer",
				labels: cu.SM{},
				data: cu.IM{
					"customer": cu.IM{
						"id": 1,
					},
				},
			},
		},
		{
			name: "product_new",
			args: args{
				mKey:   "product",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "product_edit",
			args: args{
				mKey:   "product",
				labels: cu.SM{},
				data: cu.IM{
					"product": cu.IM{
						"id":           1,
						"product_type": md.ProductTypeVirtual.String(),
						"components": []cu.IM{
							{
								"id": 1,
							},
						},
					},
				},
			},
		},
		{
			name: "tool_new",
			args: args{
				mKey:   "tool",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "tool_edit",
			args: args{
				mKey:   "tool",
				labels: cu.SM{},
				data: cu.IM{
					"tool": cu.IM{
						"id": 1,
					},
				},
			},
		},
		{
			name: "project_new",
			args: args{
				mKey:   "project",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "project_edit",
			args: args{
				mKey:   "project",
				labels: cu.SM{},
				data: cu.IM{
					"project": cu.IM{
						"id": 1,
					},
				},
			},
		},
		{
			name: "employee_new",
			args: args{
				mKey:   "employee",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "employee_edit",
			args: args{
				mKey:   "employee",
				labels: cu.SM{},
				data: cu.IM{
					"employee": cu.IM{
						"id": 1,
					},
				},
			},
		},
		{
			name: "place_new",
			args: args{
				mKey:   "place",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "place_edit",
			args: args{
				mKey:   "place",
				labels: cu.SM{},
				data: cu.IM{
					"place": cu.IM{
						"id": 1,
					},
				},
			},
		},
		{
			name: "trans_new",
			args: args{
				mKey:   "trans",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "trans_edit",
			args: args{
				mKey:   "trans",
				labels: cu.SM{},
				data: cu.IM{
					"trans": cu.IM{
						"id":         1,
						"trans_type": md.TransTypeOrder.String(),
					},
				},
			},
		},
		{
			name: "trans_bank",
			args: args{
				mKey:   "trans",
				labels: cu.SM{},
				data: cu.IM{
					"trans": cu.IM{
						"id":         1,
						"trans_type": md.TransTypeBank.String(),
					},
				},
			},
		},
		{
			name: "trans_movement",
			args: args{
				mKey:   "trans",
				labels: cu.SM{},
				data: cu.IM{
					"trans": cu.IM{
						"id":         1,
						"trans_type": md.TransTypeDelivery.String(),
					},
				},
			},
		},
		{
			name: "setting",
			args: args{
				mKey:   "setting",
				labels: cu.SM{},
				data: cu.IM{
					"setting": cu.IM{
						"id": 1,
					},
				},
			},
		},
		{
			name: "shipping",
			args: args{
				mKey:   "shipping",
				labels: cu.SM{},
				data: cu.IM{
					"shipping": cu.IM{
						"id": 1,
					},
					"items": []cu.IM{
						{
							"id":         1,
							"difference": 100,
						},
						{
							"id":         1,
							"difference": 0,
						},
					},
					"dirty": false,
				},
			},
		},
		{
			name: "invalid",
			args: args{
				mKey:   "invalid",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moduleEditorView(tt.args.mKey, tt.args.labels, tt.args.data)
		})
	}
}

func Test_moduleEditorRow(t *testing.T) {
	type args struct {
		mKey   string
		view   string
		labels cu.SM
		data   cu.IM
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "invalid",
			args: args{
				mKey:   "trans",
				view:   "invalid",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "invalid_trans_type",
			args: args{
				mKey:   "trans",
				view:   "trans",
				labels: cu.SM{},
				data: cu.IM{
					"trans": cu.IM{
						"trans_type": "INVALID",
					},
				},
			},
		},
		{
			name: "search",
			args: args{
				mKey:   "search",
				view:   "customer",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "customer",
			args: args{
				mKey:   "customer",
				view:   "customer",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "customer_maps",
			args: args{
				mKey:   "customer",
				view:   "maps",
				labels: cu.SM{},
				data: cu.IM{
					"config_map": []cu.IM{
						{
							"field_name":  "demo_number",
							"field_type":  "FIELD_NUMBER",
							"description": "Demo Number",
						},
					},
				},
			},
		},
		{
			name: "customer_contacts",
			args: args{
				mKey:   "customer",
				view:   "contacts",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "setting_invalid",
			args: args{
				mKey:   "setting",
				view:   "invalid",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "setting_config_map",
			args: args{
				mKey:   "setting",
				view:   "config_map",
				labels: cu.SM{},
				data: cu.IM{
					"config_values": []cu.IM{
						{
							"code":        "123",
							"config_type": "CONFIG_MAP",
							"data": cu.IM{
								"field_name":  "demo_number",
								"field_type":  "FIELD_NUMBER",
								"description": "Demo Number",
							},
						},
					},
				},
			},
		},
		{
			name: "setting_shortcut",
			args: args{
				mKey:   "setting",
				view:   "shortcut",
				labels: cu.SM{},
				data: cu.IM{
					"config_values": []cu.IM{
						{
							"code":        "123",
							"config_type": "CONFIG_SHORTCUT",
							"data": cu.IM{
								"shortcut_key": "shortcut_key",
								"description":  "description",
								"modul":        "modul",
								"method":       "method",
								"func_name":    "func_name",
								"address":      "address",
								"fields": []cu.IM{
									{
										"field_name":  "field_name",
										"field_type":  "FIELD_NUMBER",
										"description": "Demo Number",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "setting_auth",
			args: args{
				mKey:   "setting",
				view:   "auth",
				labels: cu.SM{},
				data: cu.IM{
					"auth": []cu.IM{
						{
							"code":       "123",
							"user_group": md.UserGroupAdmin.String(),
							"user_name":  "admin",
							"disabled":   false,
							"auth_meta": cu.IM{
								"tags": []string{},
							},
							"auth_map": cu.IM{},
						},
					},
				},
			},
		},
		{
			name: "product",
			args: args{
				mKey:   "product",
				view:   "product",
				labels: cu.SM{},
				data: cu.IM{
					"tax_codes": []cu.IM{
						{
							"tax_code": "123",
						},
					},
				},
			},
		},
		{
			name: "product_maps",
			args: args{
				mKey:   "product",
				view:   "maps",
				labels: cu.SM{},
				data: cu.IM{
					"config_map": []cu.IM{
						{
							"field_name":  "demo_number",
							"field_type":  "FIELD_NUMBER",
							"description": "Demo Number",
						},
					},
				},
			},
		},
		{
			name: "product_events",
			args: args{
				mKey:   "product",
				view:   "events",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "tool",
			args: args{
				mKey:   "tool",
				view:   "tool",
				labels: cu.SM{},
				data: cu.IM{
					"tool": cu.IM{
						"id": 1,
					},
				},
			},
		},
		{
			name: "tool_maps",
			args: args{
				mKey:   "tool",
				view:   "maps",
				labels: cu.SM{},
				data: cu.IM{
					"config_map": []cu.IM{
						{
							"field_name":  "demo_number",
							"field_type":  "FIELD_NUMBER",
							"description": "Demo Number",
						},
					},
				},
			},
		},
		{
			name: "tool_events",
			args: args{
				mKey:   "tool",
				view:   "events",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "project",
			args: args{
				mKey:   "project",
				view:   "project",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "project_maps",
			args: args{
				mKey:   "project",
				view:   "maps",
				labels: cu.SM{},
				data: cu.IM{
					"config_map": []cu.IM{
						{
							"field_name":  "demo_number",
							"field_type":  "FIELD_NUMBER",
							"description": "Demo Number",
						},
					},
				},
			},
		},
		{
			name: "project_contacts",
			args: args{
				mKey:   "project",
				view:   "contacts",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "employee",
			args: args{
				mKey:   "employee",
				view:   "employee",
				labels: cu.SM{},
				data: cu.IM{
					"employee": cu.IM{
						"id": 1,
					},
				},
			},
		},
		{
			name: "employee_address",
			args: args{
				mKey:   "employee",
				view:   "address",
				labels: cu.SM{},
				data: cu.IM{
					"employee": cu.IM{
						"id": 1,
					},
				},
			},
		},
		{
			name: "employee_contact",
			args: args{
				mKey:   "employee",
				view:   "contact",
				labels: cu.SM{},
				data: cu.IM{
					"employee": cu.IM{
						"id": 1,
					},
				},
			},
		},
		{
			name: "employee_maps",
			args: args{
				mKey:   "employee",
				view:   "maps",
				labels: cu.SM{},
				data: cu.IM{
					"config_map": []cu.IM{
						{
							"field_name":  "demo_number",
							"field_type":  "FIELD_NUMBER",
							"description": "Demo Number",
						},
					},
				},
			},
		},
		{
			name: "employee_events",
			args: args{
				mKey:   "employee",
				view:   "events",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "place",
			args: args{
				mKey:   "place",
				view:   "place",
				labels: cu.SM{},
				data: cu.IM{
					"place": cu.IM{
						"id":         1,
						"place_type": "PLACE_BANK",
					},
					"currencies": []cu.IM{
						{
							"code": "USD",
						},
					},
				},
			},
		},
		{
			name: "place_maps",
			args: args{
				mKey:   "place",
				view:   "maps",
				labels: cu.SM{},
				data: cu.IM{
					"config_map": []cu.IM{
						{
							"field_name":  "demo_number",
							"field_type":  "FIELD_NUMBER",
							"description": "Demo Number",
						},
					},
				},
			},
		},
		{
			name: "place_contacts",
			args: args{
				mKey:   "place",
				view:   "contacts",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "place_events",
			args: args{
				mKey:   "place",
				view:   "events",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "trans_order",
			args: args{
				mKey:   "trans",
				view:   "trans",
				labels: cu.SM{},
				data: cu.IM{
					"trans": cu.IM{
						"id":         1,
						"trans_type": "TRANS_ORDER",
					},
					"currencies": []cu.IM{
						{
							"code": "USD",
						},
					},
				},
			},
		},
		{
			name: "trans_receipt",
			args: args{
				mKey:   "trans",
				view:   "trans",
				labels: cu.SM{},
				data: cu.IM{
					"trans": cu.IM{
						"id":         1,
						"trans_type": "TRANS_RECEIPT",
					},
					"currencies": []cu.IM{
						{
							"code": "USD",
						},
					},
				},
			},
		},
		{
			name: "trans_worksheet",
			args: args{
				mKey:   "trans",
				view:   "trans",
				labels: cu.SM{},
				data: cu.IM{
					"trans": cu.IM{
						"id":         1,
						"trans_type": "TRANS_WORKSHEET",
						"trans_meta": cu.IM{
							"status": "STATUS_AMENDMENT",
						},
					},
					"currencies": []cu.IM{
						{
							"code": "USD",
						},
					},
				},
			},
		},
		{
			name: "trans_rent",
			args: args{
				mKey:   "trans",
				view:   "trans",
				labels: cu.SM{},
				data: cu.IM{
					"trans": cu.IM{
						"id":         1,
						"trans_type": "TRANS_RENT",
					},
					"currencies": []cu.IM{
						{
							"code": "USD",
						},
					},
				},
			},
		},
		{
			name: "trans_delivery",
			args: args{
				mKey:   "trans",
				view:   "trans",
				labels: cu.SM{},
				data: cu.IM{
					"trans": cu.IM{
						"id":         1,
						"trans_type": "TRANS_DELIVERY",
						"direction":  md.DirectionOut.String(),
						"trans_meta": cu.IM{
							"status": md.TransStatusNormal.String(),
						},
					},
					"currencies": []cu.IM{
						{
							"code": "USD",
						},
					},
					"movement_inventory": []cu.IM{
						{
							"id":             1,
							"ref_trans_code": "TR123",
						},
					},
					"places": []cu.IM{
						{
							"code":       "1",
							"place_type": md.PlaceTypeWarehouse.String(),
						},
					},
					"movements": []cu.IM{
						{
							"id":            1,
							"place_code":    "1",
							"movement_type": "MOVEMENT_TYPE_HEAD",
							"movement_meta": cu.IM{
								"qty": 1,
							},
						},
					},
				},
			},
		},
		{
			name: "trans_delivery",
			args: args{
				mKey:   "trans",
				view:   "trans",
				labels: cu.SM{},
				data: cu.IM{
					"trans": cu.IM{
						"id":         1,
						"trans_type": "TRANS_DELIVERY",
						"direction":  md.DirectionTransfer.String(),
						"trans_meta": cu.IM{
							"status": md.TransStatusNormal.String(),
						},
					},
					"currencies": []cu.IM{
						{
							"code": "USD",
						},
					},
					"places": []cu.IM{
						{
							"code":       "1",
							"place_type": md.PlaceTypeWarehouse.String(),
						},
					},
					"movements": []cu.IM{
						{
							"id":            1,
							"place_code":    "1",
							"movement_type": "MOVEMENT_TYPE_HEAD",
							"movement_meta": cu.IM{
								"qty": 1,
							},
						},
						{
							"id":            2,
							"movement_type": "MOVEMENT_TYPE_ITEM",
							"movement_code": "MV123",
							"movement_meta": cu.IM{
								"qty": 1,
							},
						},
					},
					"movement_inventory": []cu.IM{
						{
							"id":             1,
							"ref_trans_code": "TR123",
						},
					},
				},
			},
		},
		{
			name: "trans_inventory",
			args: args{
				mKey:   "trans",
				view:   "trans",
				labels: cu.SM{},
				data: cu.IM{
					"trans": cu.IM{
						"id":         1,
						"trans_type": "TRANS_INVENTORY",
						"direction":  md.DirectionTransfer.String(),
						"trans_meta": cu.IM{
							"status": md.TransStatusNormal.String(),
						},
					},
					"currencies": []cu.IM{
						{
							"code": "USD",
						},
					},
					"places": []cu.IM{
						{
							"code":       "1",
							"place_type": md.PlaceTypeWarehouse.String(),
						},
					},
					"movements": []cu.IM{
						{
							"id":            1,
							"place_code":    "1",
							"movement_type": "MOVEMENT_TYPE_HEAD",
							"movement_meta": cu.IM{
								"qty": 1,
							},
						},
						{
							"id":            2,
							"movement_type": "MOVEMENT_TYPE_ITEM",
							"movement_code": "MV123",
							"movement_meta": cu.IM{
								"qty": 1,
							},
						},
					},
				},
			},
		},
		{
			name: "trans_waybill",
			args: args{
				mKey:   "trans",
				view:   "trans",
				labels: cu.SM{},
				data: cu.IM{
					"trans": cu.IM{
						"id":         1,
						"trans_type": "TRANS_WAYBILL",
						"direction":  md.DirectionIn.String(),
						"trans_meta": cu.IM{
							"status": md.TransStatusNormal.String(),
						},
					},
					"currencies": []cu.IM{
						{
							"code": "USD",
						},
					},
					"places": []cu.IM{
						{
							"code":       "1",
							"place_type": md.PlaceTypeWarehouse.String(),
						},
					},
					"movements": []cu.IM{
						{
							"id":            1,
							"place_code":    "1",
							"movement_type": "MOVEMENT_TYPE_HEAD",
							"movement_meta": cu.IM{
								"qty": 1,
							},
						},
						{
							"id":            2,
							"movement_type": "MOVEMENT_TYPE_ITEM",
							"movement_code": "MV123",
							"movement_meta": cu.IM{
								"qty": 1,
							},
						},
					},
				},
			},
		},
		{
			name: "trans_production",
			args: args{
				mKey:   "trans",
				view:   "trans",
				labels: cu.SM{},
				data: cu.IM{
					"trans": cu.IM{
						"id":         1,
						"trans_type": "TRANS_PRODUCTION",
						"direction":  md.DirectionTransfer.String(),
						"trans_meta": cu.IM{
							"status": md.TransStatusNormal.String(),
						},
					},
					"places": []cu.IM{
						{
							"code":       "1",
							"place_type": md.PlaceTypeWarehouse.String(),
						},
					},
					"movements": []cu.IM{
						{
							"id":            1,
							"place_code":    "1",
							"movement_type": "MOVEMENT_TYPE_HEAD",
							"movement_meta": cu.IM{
								"qty": 1,
							},
						},
						{
							"id":            2,
							"movement_type": "MOVEMENT_TYPE_ITEM",
							"movement_code": "MV123",
							"movement_meta": cu.IM{
								"qty":    1,
								"shared": true,
							},
						},
					},
				},
			},
		},
		{
			name: "trans_production2",
			args: args{
				mKey:   "trans",
				view:   "trans",
				labels: cu.SM{},
				data: cu.IM{
					"trans": cu.IM{
						"id":         1,
						"trans_type": "TRANS_PRODUCTION",
						"direction":  md.DirectionTransfer.String(),
						"trans_meta": cu.IM{
							"status": md.TransStatusNormal.String(),
						},
					},
					"places": []cu.IM{
						{
							"code":       "1",
							"place_type": md.PlaceTypeWarehouse.String(),
						},
					},
					"movements": []cu.IM{
						{
							"id":            1,
							"place_code":    "1",
							"movement_type": "MOVEMENT_HEAD",
							"movement_meta": cu.IM{
								"qty": 1,
							},
						},
						{
							"id":            2,
							"movement_type": "MOVEMENT_ITEM",
							"movement_code": "MV123",
							"movement_meta": cu.IM{
								"qty": 1,
							},
						},
					},
				},
			},
		},
		{
			name: "trans_formula",
			args: args{
				mKey:   "trans",
				view:   "trans",
				labels: cu.SM{},
				data: cu.IM{
					"trans": cu.IM{
						"id":         1,
						"trans_type": "TRANS_FORMULA",
						"direction":  md.DirectionTransfer.String(),
						"trans_meta": cu.IM{
							"status": md.TransStatusNormal.String(),
						},
					},
					"places": []cu.IM{
						{
							"code":       "1",
							"place_type": md.PlaceTypeWarehouse.String(),
						},
					},
					"movements": []cu.IM{
						{
							"id":            1,
							"place_code":    "1",
							"movement_type": "MOVEMENT_HEAD",
							"movement_meta": cu.IM{
								"qty": 1,
							},
						},
						{
							"id":            2,
							"movement_type": "MOVEMENT_PLAN",
							"movement_code": "MV123",
							"movement_meta": cu.IM{
								"qty": 1,
							},
						},
					},
				},
			},
		},
		{
			name: "trans_formula2",
			args: args{
				mKey:   "trans",
				view:   "trans",
				labels: cu.SM{},
				data: cu.IM{
					"trans": cu.IM{
						"id":         1,
						"trans_type": "TRANS_FORMULA",
						"direction":  md.DirectionTransfer.String(),
						"trans_meta": cu.IM{
							"status": md.TransStatusNormal.String(),
						},
					},
					"places": []cu.IM{
						{
							"code":       "1",
							"place_type": md.PlaceTypeWarehouse.String(),
						},
					},
					"movements": []cu.IM{
						{
							"id":            2,
							"movement_type": "MOVEMENT_PLAN",
							"movement_code": "MV123",
							"movement_meta": cu.IM{
								"qty": 1,
							},
						},
					},
				},
			},
		},
		{
			name: "trans_maps",
			args: args{
				mKey:   "trans",
				view:   "maps",
				labels: cu.SM{},
				data: cu.IM{
					"config_map": []cu.IM{
						{
							"field_name":  "demo_number",
							"field_type":  "FIELD_NUMBER",
							"description": "Demo Number",
						},
					},
				},
			},
		},
		{
			name: "trans_items",
			args: args{
				mKey:   "trans",
				view:   "items",
				labels: cu.SM{},
				data: cu.IM{
					"items": []cu.IM{
						{
							"product_code": "123",
							"tax_code":     "123",
						},
					},
				},
			},
		},
		{
			name: "trans_payments",
			args: args{
				mKey:   "trans",
				view:   "payments",
				labels: cu.SM{},
				data: cu.IM{
					"trans": cu.IM{
						"id": 1,
					},
					"currencies": []cu.IM{
						{
							"code": "USD",
						},
					},
					"payments": []cu.IM{
						{
							"id": 1,
							"payment_meta": cu.IM{
								"amount": 100,
							},
						},
						{
							"id": 2,
							"payment_meta": cu.IM{
								"amount": -100,
							},
						},
					},
				},
			},
		},
		{
			name: "trans_bank",
			args: args{
				mKey:   "trans",
				view:   "trans",
				labels: cu.SM{},
				data: cu.IM{
					"trans": cu.IM{
						"id":         1,
						"trans_type": md.TransTypeBank.String(),
					},
					"currencies": []cu.IM{
						{
							"code": "USD",
						},
					},
					"places": []cu.IM{
						{
							"code":       "1",
							"place_type": md.PlaceTypeBank.String(),
						},
					},
					"payments": []cu.IM{
						{
							"id": 1,
							"payment_meta": cu.IM{
								"amount": 100,
							},
						},
					},
				},
			},
		},
		{
			name: "trans_cash",
			args: args{
				mKey:   "trans",
				view:   "trans",
				labels: cu.SM{},
				data: cu.IM{
					"trans": cu.IM{
						"id":         1,
						"trans_type": md.TransTypeCash.String(),
					},
					"currencies": []cu.IM{
						{
							"code": "USD",
						},
					},
					"places": []cu.IM{
						{
							"code":       "1",
							"place_type": md.PlaceTypeCash.String(),
						},
					},
					"payments": []cu.IM{
						{
							"id": 1,
							"payment_meta": cu.IM{
								"amount": 100,
							},
						},
					},
				},
			},
		},
		{
			name: "shipping",
			args: args{
				mKey:   "shipping",
				view:   "shipping",
				labels: cu.SM{},
				data: cu.IM{
					"shipping": cu.IM{
						"id": 1,
					},
					"items": []cu.IM{
						{
							"id":         1,
							"difference": 100,
						},
					},
					"places": []cu.IM{
						{
							"code":       "1",
							"place_name": "Place 1",
						},
					},
				},
			},
		},
		{
			name: "shipping",
			args: args{
				mKey:   "shipping",
				view:   "items",
				labels: cu.SM{},
				data: cu.IM{
					"shipping": cu.IM{
						"id": 1,
					},
					"items": []cu.IM{
						{
							"id":         1,
							"difference": 100,
						},
					},
					"places": []cu.IM{
						{
							"code":       "1",
							"place_name": "Place 1",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moduleEditorRow(tt.args.mKey, tt.args.view, tt.args.labels, tt.args.data)
		})
	}
}

func Test_moduleEditorTable(t *testing.T) {
	type args struct {
		mKey   string
		view   string
		labels cu.SM
		data   cu.IM
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "search",
			args: args{
				mKey:   "search",
				view:   "customer",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "customer",
			args: args{
				mKey:   "customer",
				view:   "customer",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "customer_events",
			args: args{
				mKey:   "customer",
				view:   "events",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "customer_addresses",
			args: args{
				mKey:   "customer",
				view:   "addresses",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "customer_contacts",
			args: args{
				mKey:   "customer",
				view:   "contacts",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "customer_maps",
			args: args{
				mKey:   "customer",
				view:   "maps",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "tool",
			args: args{
				mKey:   "tool",
				view:   "tool",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "tool_events",
			args: args{
				mKey:   "tool",
				view:   "events",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "tool_maps",
			args: args{
				mKey:   "tool",
				view:   "maps",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "product",
			args: args{
				mKey:   "product",
				view:   "product",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "product_events",
			args: args{
				mKey:   "product",
				view:   "events",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "product_prices",
			args: args{
				mKey:   "product",
				view:   "prices",
				labels: cu.SM{},
				data: cu.IM{
					"prices": []cu.IM{
						{
							"price": "123",
						},
					},
				},
			},
		},
		{
			name: "product_components",
			args: args{
				mKey:   "product",
				view:   "components",
				labels: cu.SM{},
				data: cu.IM{
					"components": []cu.IM{
						{
							"id": 1,
							"link_meta": cu.IM{
								"qty":   1,
								"notes": "123",
							},
						},
					},
				},
			},
		},
		{
			name: "product_maps",
			args: args{
				mKey:   "product",
				view:   "maps",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "project",
			args: args{
				mKey:   "project",
				view:   "project",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "project_events",
			args: args{
				mKey:   "project",
				view:   "events",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "project_addresses",
			args: args{
				mKey:   "project",
				view:   "addresses",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "project_contacts",
			args: args{
				mKey:   "project",
				view:   "contacts",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "project_maps",
			args: args{
				mKey:   "project",
				view:   "maps",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "employee",
			args: args{
				mKey:   "employee",
				view:   "employee",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "employee_events",
			args: args{
				mKey:   "employee",
				view:   "events",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "employee_address",
			args: args{
				mKey:   "employee",
				view:   "address",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "employee_contact",
			args: args{
				mKey:   "employee",
				view:   "contact",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "employee_maps",
			args: args{
				mKey:   "employee",
				view:   "maps",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "place",
			args: args{
				mKey:   "place",
				view:   "place",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "place_contacts",
			args: args{
				mKey:   "place",
				view:   "contacts",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "place_events",
			args: args{
				mKey:   "place",
				view:   "events",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "place_maps",
			args: args{
				mKey:   "place",
				view:   "maps",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "trans",
			args: args{
				mKey:   "trans",
				view:   "trans",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "trans_items",
			args: args{
				mKey:   "trans",
				view:   "items",
				labels: cu.SM{},
				data: cu.IM{
					"items": []cu.IM{
						{
							"trans_code":  "123",
							"trans_date":  "2021-01-01",
							"description": "Demo Number",
							"unit":        "1",
							"qty":         "1",
							"amount":      "1",
							"deposit":     "1",
						},
					},
				},
			},
		},
		{
			name: "trans_transitem_invoice",
			args: args{
				mKey:   "trans",
				view:   "transitem_invoice",
				labels: cu.SM{},
				data: cu.IM{
					"transitem_invoice": []cu.IM{
						{
							"trans_code":  "123",
							"trans_date":  "2021-01-01",
							"description": "Demo Number",
							"unit":        "1",
							"qty":         "1",
							"amount":      "1",
							"deposit":     "1",
						},
					},
				},
			},
		},
		{
			name: "trans_payments",
			args: args{
				mKey:   "trans",
				view:   "payments",
				labels: cu.SM{},
				data: cu.IM{
					"payments": []cu.IM{
						{
							"id": 1,
							"payment_meta": cu.IM{
								"amount": 100,
							},
						},
						{
							"id": 2,
							"payment_meta": cu.IM{
								"amount": -100,
							},
						},
					},
				},
			},
		},
		{
			name: "trans_payment_link",
			args: args{
				mKey:   "trans",
				view:   "payment_link",
				labels: cu.SM{},
				data: cu.IM{
					"trans": cu.IM{
						"id":         1,
						"trans_type": md.TransTypeInvoice.String(),
					},
					"payment_link": []cu.IM{
						{
							"id": 1,
							"link_meta": cu.IM{
								"amount": 100,
							},
						},
						{
							"id": 2,
							"link_meta": cu.IM{
								"amount": -100,
							},
						},
					},
				},
			},
		},
		{
			name: "trans_maps",
			args: args{
				mKey:   "trans",
				view:   "maps",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "setting",
			args: args{
				mKey:   "setting",
				view:   "setting",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "setting_config_data",
			args: args{
				mKey:   "setting",
				view:   "config_data",
				labels: cu.SM{},
				data: cu.IM{
					"config_data": []cu.IM{
						{
							"config_code":  "123",
							"config_key":   "123",
							"config_value": "123",
						},
					},
				},
			},
		},
		{
			name: "setting_currency",
			args: args{
				mKey:   "setting",
				view:   "currency",
				labels: cu.SM{},
				data: cu.IM{
					"currency": []cu.IM{
						{
							"code": "123",
							"currency_meta": cu.IM{
								"description": "123",
								"digit":       123,
								"cash_round":  123,
							},
						},
					},
				},
			},
		},
		{
			name: "setting_tax",
			args: args{
				mKey:   "setting",
				view:   "tax",
				labels: cu.SM{},
				data: cu.IM{
					"tax": []cu.IM{
						{
							"code": "123",
							"tax_meta": cu.IM{
								"description": "123",
								"rate_value":  123,
							},
						},
					},
				},
			},
		},
		{
			name: "setting_template",
			args: args{
				mKey:   "setting",
				view:   "template",
				labels: cu.SM{},
				data: cu.IM{
					"template": []cu.IM{
						{
							"code":        "123",
							"report_key":  "123",
							"report_name": "123",
							"label":       "123",
							"description": "123",
							"installed":   true,
							"id":          123,
						},
					},
				},
			},
		},
		{
			name: "trans_movements_delivery",
			args: args{
				mKey:   "trans",
				view:   "movements",
				labels: cu.SM{},
				data: cu.IM{
					"trans": cu.IM{
						"id":         1,
						"trans_type": md.TransTypeDelivery.String(),
						"direction":  md.DirectionOut.String(),
						"trans_meta": cu.IM{
							"status": md.TransStatusNormal.String(),
						},
					},
					"movements": []cu.IM{
						{
							"id": 1,
							"movement_meta": cu.IM{
								"qty": 100,
							},
						},
						{
							"id": 2,
							"movement_meta": cu.IM{
								"qty": -100,
							},
						},
					},
					"movement_inventory": []cu.IM{
						{
							"id":             1,
							"ref_trans_code": "TR123",
						},
					},
					"places": []cu.IM{
						{
							"code":       "1",
							"place_type": md.PlaceTypeWarehouse.String(),
						},
					},
				},
			},
		},
		{
			name: "trans_movements_inventory",
			args: args{
				mKey:   "trans",
				view:   "movements",
				labels: cu.SM{},
				data: cu.IM{
					"trans": cu.IM{
						"id":         1,
						"trans_type": md.TransTypeInventory.String(),
						"direction":  md.DirectionTransfer.String(),
						"trans_meta": cu.IM{
							"status": md.TransStatusNormal.String(),
						},
					},
					"movements": []cu.IM{
						{
							"id": 1,
							"movement_meta": cu.IM{
								"qty": 100,
							},
						},
						{
							"id": 2,
							"movement_meta": cu.IM{
								"qty": -100,
							},
						},
					},
					"movement_inventory": []cu.IM{
						{
							"id":             1,
							"ref_trans_code": "TR123",
						},
					},
					"places": []cu.IM{
						{
							"code":       "1",
							"place_type": md.PlaceTypeWarehouse.String(),
						},
					},
				},
			},
		},
		{
			name: "trans_movements_production",
			args: args{
				mKey:   "trans",
				view:   "movements",
				labels: cu.SM{},
				data: cu.IM{
					"trans": cu.IM{
						"id":         1,
						"trans_type": md.TransTypeProduction.String(),
						"direction":  md.DirectionTransfer.String(),
						"trans_meta": cu.IM{
							"status": md.TransStatusNormal.String(),
						},
					},
					"movements": []cu.IM{
						{
							"id": 1,
							"movement_meta": cu.IM{
								"qty": 100,
							},
						},
						{
							"id": 2,
							"movement_meta": cu.IM{
								"qty": -100,
							},
						},
					},
					"movement_inventory": []cu.IM{
						{
							"id":             1,
							"ref_trans_code": "TR123",
						},
					},
					"places": []cu.IM{
						{
							"code":       "1",
							"place_type": md.PlaceTypeWarehouse.String(),
						},
					},
				},
			},
		},
		{
			name: "trans_movements_waybill",
			args: args{
				mKey:   "trans",
				view:   "movements",
				labels: cu.SM{},
				data: cu.IM{
					"trans": cu.IM{
						"id":         1,
						"trans_type": md.TransTypeWaybill.String(),
						"direction":  md.DirectionTransfer.String(),
						"trans_meta": cu.IM{
							"status": md.TransStatusNormal.String(),
						},
					},
					"movements": []cu.IM{
						{
							"id": 1,
							"movement_meta": cu.IM{
								"qty": 100,
							},
						},
						{
							"id": 2,
							"movement_meta": cu.IM{
								"qty": -100,
							},
						},
					},
					"movement_waybill": []cu.IM{
						{
							"id":             1,
							"ref_trans_code": "TR123",
						},
					},
					"places": []cu.IM{
						{
							"code":       "1",
							"place_type": md.PlaceTypeWarehouse.String(),
						},
					},
				},
			},
		},
		{
			name: "trans_movements_formula",
			args: args{
				mKey:   "trans",
				view:   "movements",
				labels: cu.SM{},
				data: cu.IM{
					"trans": cu.IM{
						"id":         1,
						"trans_type": md.TransTypeFormula.String(),
						"direction":  md.DirectionTransfer.String(),
						"trans_meta": cu.IM{
							"status": md.TransStatusNormal.String(),
						},
					},
					"movements": []cu.IM{
						{
							"id": 1,
							"movement_meta": cu.IM{
								"qty": 100,
							},
						},
						{
							"id": 2,
							"movement_meta": cu.IM{
								"qty": -100,
							},
						},
					},
					"movement_formula": []cu.IM{
						{
							"id":             1,
							"ref_trans_code": "TR123",
						},
					},
					"places": []cu.IM{
						{
							"code":       "1",
							"place_type": md.PlaceTypeWarehouse.String(),
						},
					},
				},
			},
		},
		{
			name: "trans_transitem_shipping",
			args: args{
				mKey:   "trans",
				view:   "transitem_shipping",
				labels: cu.SM{},
				data: cu.IM{
					"trans": cu.IM{
						"id":         1,
						"trans_type": md.TransTypeInvoice.String(),
					},
					"movements": []cu.IM{
						{
							"id": 1,
							"movement_meta": cu.IM{
								"qty": 100,
							},
						},
					},
					"places": []cu.IM{
						{
							"code":       "1",
							"place_type": md.PlaceTypeWarehouse.String(),
						},
					},
				},
			},
		},
		{
			name: "trans_tool_movement",
			args: args{
				mKey:   "trans",
				view:   "tool_movement",
				labels: cu.SM{},
				data: cu.IM{
					"trans": cu.IM{
						"id":         1,
						"trans_type": md.TransTypeInvoice.String(),
					},
					"movements": []cu.IM{
						{
							"id": 1,
							"movement_meta": cu.IM{
								"qty": 100,
							},
						},
					},
				},
			},
		},
		{
			name: "shipping_movements",
			args: args{
				mKey:   "shipping",
				view:   "movements",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "shipping_items",
			args: args{
				mKey:   "shipping",
				view:   "items",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "shipping_shipping",
			args: args{
				mKey:   "shipping",
				view:   "shipping",
				labels: cu.SM{},
				data:   cu.IM{},
			},
		},
		{
			name: "invalid",
			args: args{
				mKey: "invalid",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moduleEditorTable(tt.args.mKey, tt.args.view, tt.args.labels, tt.args.data)
		})
	}
}

func TestClientComponent_Labels(t *testing.T) {
	type fields struct {
		languages     []string
		helpURL       string
		clientHelpURL string
		exportURL     string
		SearchConfig  *SearchConfig
		editorMap     map[string]EditorInterface
		modalMap      map[string]func(labels cu.SM, data cu.IM) ct.Form
	}
	type args struct {
		lang string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "en",
			fields: fields{
				languages: []string{"en"},
			},
			args: args{
				lang: "en",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := &ClientComponent{
				languages:     tt.fields.languages,
				helpURL:       tt.fields.helpURL,
				clientHelpURL: tt.fields.clientHelpURL,
				exportURL:     tt.fields.exportURL,
				SearchConfig:  tt.fields.SearchConfig,
				editorMap:     tt.fields.editorMap,
				modalMap:      tt.fields.modalMap,
			}
			cc.Labels(tt.args.lang)
		})
	}
}
