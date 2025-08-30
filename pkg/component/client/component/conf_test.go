package component

import (
	"reflect"
	"testing"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
)

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
						"id": 1,
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
			name: "invalid",
			args: args{
				mKey: "invalid",
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

func TestClientMenu(t *testing.T) {
	type args struct {
		labels cu.SM
		config cu.IM
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "default",
			args: args{
				labels: cu.SM{},
				config: cu.IM{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ClientMenu(tt.args.labels, tt.args.config)
		})
	}
}

func TestClientLogin(t *testing.T) {
	type args struct {
		labels cu.SM
		config cu.IM
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "default",
			args: args{
				labels: cu.SM{},
				config: cu.IM{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ClientLogin(tt.args.labels, tt.args.config)
		})
	}
}

func TestClientSideBar(t *testing.T) {
	type args struct {
		moduleKey string
		labels    cu.SM
		data      cu.IM
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "search",
			args: args{
				moduleKey: "search",
				labels:    cu.SM{},
				data:      cu.IM{},
			},
		},
		{
			name: "browser",
			args: args{
				moduleKey: "browser",
				labels:    cu.SM{},
				data:      cu.IM{},
			},
		},
		{
			name: "customer_new",
			args: args{
				moduleKey: "customer",
				labels:    cu.SM{},
				data: cu.IM{
					"customer": cu.IM{
						"id": 0,
					},
				},
			},
		},
		{
			name: "customer_edit",
			args: args{
				moduleKey: "customer",
				labels:    cu.SM{},
				data: cu.IM{
					"customer": cu.IM{
						"id": 1,
					},
				},
			},
		},
		{
			name: "customer_inactive",
			args: args{
				moduleKey: "customer",
				labels:    cu.SM{},
				data: cu.IM{
					"customer": cu.IM{
						"id":       1,
						"inactive": true,
					},
				},
			},
		},
		{
			name: "product_new",
			args: args{
				moduleKey: "product",
				labels:    cu.SM{},
				data: cu.IM{
					"product": cu.IM{
						"id": 0,
					},
				},
			},
		},
		{
			name: "product_edit",
			args: args{
				moduleKey: "product",
				labels:    cu.SM{},
				data: cu.IM{
					"product": cu.IM{
						"id": 1,
					},
				},
			},
		},
		{
			name: "product_inactive",
			args: args{
				moduleKey: "product",
				labels:    cu.SM{},
				data: cu.IM{
					"product": cu.IM{
						"id":       1,
						"inactive": true,
					},
				},
			},
		},
		{
			name: "tool_new",
			args: args{
				moduleKey: "tool",
				labels:    cu.SM{},
				data: cu.IM{
					"tool": cu.IM{
						"id": 0,
					},
				},
			},
		},
		{
			name: "tool_edit",
			args: args{
				moduleKey: "tool",
				labels:    cu.SM{},
				data: cu.IM{
					"tool": cu.IM{
						"id": 1,
					},
				},
			},
		},
		{
			name: "product_inactive",
			args: args{
				moduleKey: "tool",
				labels:    cu.SM{},
				data: cu.IM{
					"tool": cu.IM{
						"id":       1,
						"inactive": true,
					},
				},
			},
		},
		{
			name: "setting",
			args: args{
				moduleKey: "setting",
				labels:    cu.SM{},
				data:      cu.IM{},
			},
		},
		{
			name: "invalid",
			args: args{
				moduleKey: "invalid",
			},
		},
		{
			name: "project_new",
			args: args{
				moduleKey: "project",
				labels:    cu.SM{},
				data: cu.IM{
					"project": cu.IM{
						"id": 0,
					},
				},
			},
		},
		{
			name: "project_edit",
			args: args{
				moduleKey: "project",
				labels:    cu.SM{},
				data: cu.IM{
					"project": cu.IM{
						"id": 1,
					},
				},
			},
		},
		{
			name: "project_inactive",
			args: args{
				moduleKey: "project",
				labels:    cu.SM{},
				data: cu.IM{
					"project": cu.IM{
						"id":       1,
						"inactive": true,
					},
				},
			},
		},
		{
			name: "employee_new",
			args: args{
				moduleKey: "employee",
				labels:    cu.SM{},
				data: cu.IM{
					"employee": cu.IM{
						"id": 0,
					},
				},
			},
		},
		{
			name: "employee_edit",
			args: args{
				moduleKey: "employee",
				labels:    cu.SM{},
				data: cu.IM{
					"employee": cu.IM{
						"id": 1,
					},
				},
			},
		},
		{
			name: "employee_inactive",
			args: args{
				moduleKey: "employee",
				labels:    cu.SM{},
				data: cu.IM{
					"employee": cu.IM{
						"id":       1,
						"inactive": true,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ClientSideBar(tt.args.moduleKey, tt.args.labels, tt.args.data)
		})
	}
}

func TestClientEditor(t *testing.T) {
	type args struct {
		editorKey  string
		viewName   string
		labels     cu.SM
		editorData cu.IM
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "search",
			args: args{
				editorKey:  "search",
				viewName:   "search",
				labels:     cu.SM{},
				editorData: cu.IM{},
			},
		},
		{
			name: "customer",
			args: args{
				editorKey:  "customer",
				viewName:   "customer",
				labels:     cu.SM{},
				editorData: cu.IM{},
			},
		},
		{
			name: "product",
			args: args{
				editorKey:  "product",
				viewName:   "product",
				labels:     cu.SM{},
				editorData: cu.IM{},
			},
		},
		{
			name: "tool",
			args: args{
				editorKey:  "tool",
				viewName:   "tool",
				labels:     cu.SM{},
				editorData: cu.IM{},
			},
		},
		{
			name: "project",
			args: args{
				editorKey:  "project",
				viewName:   "project",
				labels:     cu.SM{},
				editorData: cu.IM{},
			},
		},
		{
			name: "setting",
			args: args{
				editorKey:  "setting",
				viewName:   "setting",
				labels:     cu.SM{},
				editorData: cu.IM{},
			},
		},
		{
			name: "employee",
			args: args{
				editorKey:  "employee",
				viewName:   "employee",
				labels:     cu.SM{},
				editorData: cu.IM{},
			},
		},
		{
			name: "invalid",
			args: args{
				editorKey: "invalid",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ClientEditor(tt.args.editorKey, tt.args.viewName, tt.args.labels, tt.args.editorData)
		})
	}
}

func TestClientForm(t *testing.T) {
	type args struct {
		editorKey string
		formKey   string
		labels    cu.SM
		data      cu.IM
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "customer_addresses",
			args: args{
				editorKey: "customer",
				formKey:   "addresses",
				labels:    cu.SM{},
				data:      cu.IM{},
			},
		},
		{
			name: "customer_contacts",
			args: args{
				editorKey: "customer",
				formKey:   "contacts",
				labels:    cu.SM{},
				data:      cu.IM{},
			},
		},
		{
			name: "customer_events",
			args: args{
				editorKey: "customer",
				formKey:   "events",
				labels:    cu.SM{},
				data:      cu.IM{},
			},
		},
		{
			name: "customer_customer",
			args: args{
				editorKey: "customer",
				formKey:   "customer",
				labels:    cu.SM{},
				data:      cu.IM{},
			},
		},
		{
			name: "product_events",
			args: args{
				editorKey: "product",
				formKey:   "events",
				labels:    cu.SM{},
				data:      cu.IM{},
			},
		},
		{
			name: "product_prices",
			args: args{
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
		},
		{
			name: "product_product",
			args: args{
				editorKey: "product",
				formKey:   "product",
				labels:    cu.SM{},
				data:      cu.IM{},
			},
		},
		{
			name: "tool_events",
			args: args{
				editorKey: "tool",
				formKey:   "events",
				labels:    cu.SM{},
				data:      cu.IM{},
			},
		},
		{
			name: "tool_product",
			args: args{
				editorKey: "tool",
				formKey:   "tool",
				labels:    cu.SM{},
				data:      cu.IM{},
			},
		},
		{
			name: "setting",
			args: args{
				editorKey: "setting",
				formKey:   "setting",
				labels:    cu.SM{},
				data:      cu.IM{},
			},
		},
		{
			name: "project_addresses",
			args: args{
				editorKey: "project",
				formKey:   "addresses",
				labels:    cu.SM{},
				data:      cu.IM{},
			},
		},
		{
			name: "project_contacts",
			args: args{
				editorKey: "project",
				formKey:   "contacts",
				labels:    cu.SM{},
				data:      cu.IM{},
			},
		},
		{
			name: "project_events",
			args: args{
				editorKey: "project",
				formKey:   "events",
				labels:    cu.SM{},
				data:      cu.IM{},
			},
		},
		{
			name: "project_project",
			args: args{
				editorKey: "project",
				formKey:   "project",
				labels:    cu.SM{},
				data:      cu.IM{},
			},
		},
		{
			name: "employee_events",
			args: args{
				editorKey: "employee",
				formKey:   "events",
				labels:    cu.SM{},
				data:      cu.IM{},
			},
		},
		{
			name: "employee_employee",
			args: args{
				editorKey: "employee",
				formKey:   "employee",
				labels:    cu.SM{},
				data:      cu.IM{},
			},
		},
		{
			name: "invalid",
			args: args{
				editorKey: "invalid",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ClientForm(tt.args.editorKey, tt.args.formKey, tt.args.labels, tt.args.data)
		})
	}
}

func TestClientModalForm(t *testing.T) {
	type args struct {
		formKey string
		labels  cu.SM
		data    cu.IM
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "info",
			args: args{
				formKey: "info",
				labels:  cu.SM{},
				data:    cu.IM{},
			},
		},
		{
			name: "warning",
			args: args{
				formKey: "warning",
				labels:  cu.SM{},
				data:    cu.IM{},
			},
		},
		{
			name: "input_string",
			args: args{
				formKey: "input_string",
				labels:  cu.SM{},
				data:    cu.IM{},
			},
		},
		{
			name: "report",
			args: args{
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
		},
		{
			name: "selector",
			args: args{
				formKey: "selector",
				labels:  cu.SM{},
				data:    cu.IM{},
			},
		},
		{
			name: "missing",
			args: args{
				formKey: "missing",
				labels:  cu.SM{},
				data:    cu.IM{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ClientModalForm(tt.args.formKey, tt.args.labels, tt.args.data)
		})
	}
}

func TestClientBrowser(t *testing.T) {
	type args struct {
		viewName   string
		labels     cu.SM
		searchData cu.IM
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "customer",
			args: args{
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ClientBrowser(tt.args.viewName, tt.args.labels, tt.args.searchData)
		})
	}
}

func TestClientSearch(t *testing.T) {
	type args struct {
		viewName   string
		labels     cu.SM
		searchData cu.IM
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "customer",
			args: args{
				viewName:   "customer",
				labels:     cu.SM{},
				searchData: cu.IM{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ClientSearch(tt.args.viewName, tt.args.labels, tt.args.searchData)
		})
	}
}

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
