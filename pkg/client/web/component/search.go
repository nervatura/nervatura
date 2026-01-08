package component

import (
	"fmt"
	"html/template"
	"slices"
	"strings"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
	st "github.com/nervatura/nervatura/v6/pkg/static"
)

var compMap = cu.SM{
	"==": "=", "!=": "<>", "<": "<", "<=": "<=", ">": ">", ">=": ">=",
}

var compMapString = cu.SM{
	"==": "like", "!=": "not like",
}

var pre = func(or bool) string {
	if or {
		return "or"
	}
	return "and"
}

type SearchConfig struct{}

func (s *SearchConfig) SideBar(labels cu.SM, data cu.IM) (items []ct.SideBarItem) {
	sideGroup := cu.ToString(data["side_group"], "")
	authFilter := ut.ToStringArray(data["auth_filter"])
	userGroup := cu.ToString(data["user_group"], "")
	var sessionID string
	config := cu.ToIM(data["config"], cu.IM{})
	if ticket, found := config["ticket"].(ct.Ticket); found {
		sessionID = ticket.SessionID
	}

	sideElement := func(name string) *ct.SideBarElement {
		return &ct.SideBarElement{
			Name:     name,
			Value:    name,
			Label:    " " + s.View(name, labels, sessionID).Title,
			Icon:     s.View(name, labels, sessionID).Icon,
			Disabled: s.View(name, labels, sessionID).Disabled,
			//Selected: (cu.ToString(data["view"], "") == name),
		}
	}

	sideGroupElement := func(group md.SideGroup) *ct.SideBarGroup {
		sg := &ct.SideBarGroup{
			Name:  group.Name,
			Value: group.Name,
			Label: group.Label,
		}
		for _, name := range group.Views {
			se := sideElement(name)
			//se.Selected = (cu.ToString(data["view"], "") == name)
			sg.Items = append(sg.Items, *se)
		}
		return sg
	}

	selectedGroup := func(group md.SideGroup) bool {
		return (sideGroup == group.Name) ||
			((sideGroup == "") && slices.Contains(group.Views, cu.ToString(data["view"], "")))
	}

	visibleGroups := func(group md.SideGroup) bool {
		return userGroup == md.UserGroupAdmin.String() || len(authFilter) == 0 || slices.Contains(authFilter, group.AuthFilter)
	}

	sb := []ct.SideBarItem{
		&ct.SideBarSeparator{},
	}
	for _, group := range s.SideGroups(labels) {
		if visibleGroups(group) {
			groupElement := sideGroupElement(group)
			groupElement.Selected = selectedGroup(group)
			groupElement.Disabled = group.Disabled
			sb = append(sb, groupElement)
		}
	}

	return sb
}

func (s *SearchConfig) SideGroups(labels cu.SM) []md.SideGroup {
	return []md.SideGroup{
		{
			Name:  "group_transitem",
			Label: labels["transitem_title"],
			Views: []string{
				"transitem_simple", "transitem", "transitem_map", "transitem_item",
			},
			AuthFilter: md.AuthFilterTransItem.String(),
		},
		{
			Name:  "group_transpayment",
			Label: labels["transpayment_title"],
			Views: []string{
				"transpayment_simple", "transpayment", "transpayment_map", "transpayment_invoice",
			},
			AuthFilter: md.AuthFilterTransPayment.String(),
		},
		{
			Name:  "group_transmovement",
			Label: labels["transmovement_title"],
			Views: []string{
				"transmovement_simple", "transmovement_stock", "transmovement", "transmovement_waybill",
				"transmovement_formula", "transmovement_map",
			},
			AuthFilter: md.AuthFilterTransMovement.String(),
		},
		{
			Name:  "group_customer",
			Label: labels["customer_title"],
			Views: []string{
				"customer_simple", "customer", "customer_map", "customer_addresses", "customer_contacts", "customer_events",
			},
			AuthFilter: md.AuthFilterCustomer.String(),
		},
		{
			Name:  "group_product",
			Label: labels["product_title"],
			Views: []string{
				"product_simple", "product", "product_map", "product_events", "product_prices", "product_components",
			},
			AuthFilter: md.AuthFilterProduct.String(),
		},
		{
			Name:  "group_employee",
			Label: labels["employee_title"],
			Views: []string{
				"employee_simple", "employee", "employee_map", "employee_events",
			},
			AuthFilter: md.AuthFilterEmployee.String(),
		},
		{
			Name:  "group_tool",
			Label: labels["tool_title"],
			Views: []string{
				"tool_simple", "tool", "tool_map", "tool_events",
			},
			AuthFilter: md.AuthFilterTool.String(),
		},
		{
			Name:  "group_project",
			Label: labels["project_title"],
			Views: []string{
				"project_simple", "project", "project_map", "project_addresses", "project_contacts", "project_events",
			},
			AuthFilter: md.AuthFilterProject.String(),
		},
		{
			Name:  "group_place",
			Label: labels["place_title"],
			Views: []string{
				"place_simple", "place", "place_map", "place_contacts", "place_events",
			},
			AuthFilter: md.AuthFilterPlace.String(),
		},
		{
			Name:  "group_office",
			Label: labels["office_title"],
			Views: []string{
				"office_rate", "office_queue", "office_template_editor", "office_shortcut", "office_log",
			},
			AuthFilter: md.AuthFilterOffice.String(),
		},
	}
}

func (s *SearchConfig) View(view string, labels cu.SM, sessionID string) md.SearchView {
	viewMap := map[string]md.SearchView{
		"transitem_simple": {
			Title:    labels["quick_search"],
			Icon:     ct.IconBolt,
			Simple:   true,
			ReadOnly: false,
			LabelAdd: "",
			Fields: []ct.TableField{
				{Name: "code", Label: labels["trans_code"]},
				{Name: "trans_date", Label: labels["trans_date"]},
				{Name: "customer_name", Label: labels["customer_name"]},
				{Name: "currency_code", Label: labels["currency_code"]},
				{Name: "amount", Label: labels["item_amount"], FieldType: ct.TableFieldTypeNumber},
				{Name: "trans_type", Label: labels["trans_type"]},
				{Name: "direction", Label: labels["trans_direction"]},
			},
			VisibleColumns: cu.IM{},
			HideFilters:    cu.IM{},
			Filters: []any{
				cu.IM{"or": true, "field": "t.code", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "trans_type", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "direction", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "trans_date", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "customer_name", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "currency_code", "comp": "==", "value": ""},
			},
			FilterPlaceholder: strings.Join([]string{
				labels["trans_code"], labels["trans_date"],
				labels["customer_name"], labels["currency_code"], labels["trans_type"],
				labels["trans_direction"],
			}, ", "),
		},
		"invoice_simple": {
			Title:    labels["quick_search"],
			Icon:     ct.IconBolt,
			Simple:   true,
			ReadOnly: false,
			LabelAdd: "",
			Fields: []ct.TableField{
				{Name: "code", Label: labels["trans_code"]},
				{Name: "trans_date", Label: labels["trans_date"]},
				{Name: "customer_name", Label: labels["customer_name"]},
				{Name: "currency_code", Label: labels["currency_code"]},
				{Name: "amount", Label: labels["item_amount"], FieldType: ct.TableFieldTypeNumber},
				{Name: "trans_type", Label: labels["trans_type"]},
				{Name: "direction", Label: labels["trans_direction"]},
			},
			VisibleColumns: cu.IM{},
			HideFilters:    cu.IM{},
			Filters: []any{
				cu.IM{"or": true, "field": "t.code", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "trans_type", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "direction", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "trans_date", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "customer_name", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "currency_code", "comp": "==", "value": ""},
			},
			FilterPlaceholder: strings.Join([]string{
				labels["trans_code"], labels["trans_date"],
				labels["customer_name"], labels["currency_code"], labels["trans_type"],
				labels["trans_direction"],
			}, ", "),
		},
		"transitem": {
			Title:    labels["transitem_view"],
			Icon:     ct.IconFileText,
			Simple:   false,
			ReadOnly: false,
			LabelAdd: labels["transitem_new"],
			Fields: []ct.TableField{
				{Name: "code", Label: labels["transitem_code"]},
				{Name: "trans_type", Label: labels["trans_type"]},
				{Name: "direction", Label: labels["trans_direction"]},
				{Name: "trans_date", Label: labels["trans_date"]},
				{Name: "customer_code", Label: labels["customer_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "customer_name", Label: labels["customer_name"]},
				{Name: "currency_code", Label: labels["currency_code"]},
				{Name: "amount", Label: labels["item_amount"], FieldType: ct.TableFieldTypeNumber},
				{Name: "trans_code", Label: labels["trans_trans_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "employee_code", Label: labels["employee_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "project_code", Label: labels["project_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "place_code", Label: labels["place_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "due_time", Label: labels["trans_due_time"], FieldType: ct.TableFieldTypeDateTime},
				{Name: "ref_number", Label: labels["trans_ref_number"]},
				{Name: "paid_type", Label: labels["trans_paid_type"]},
				{Name: "rate", Label: labels["trans_rate"], FieldType: ct.TableFieldTypeNumber},
				{Name: "status", Label: labels["trans_status"]},
				{Name: "trans_state", Label: labels["trans_state"]},
				{Name: "notes", Label: labels["trans_notes"]},
				{Name: "internal_notes", Label: labels["trans_internal_notes"]},
				{Name: "tax_free", Label: labels["trans_tax_free"], FieldType: ct.TableFieldTypeBool},
				{Name: "paid", Label: labels["trans_paid"], FieldType: ct.TableFieldTypeBool},
				{Name: "closed", Label: labels["trans_closed"], FieldType: ct.TableFieldTypeBool},
				{Name: "tag_lst", Label: labels["trans_tags"]},
				{Name: "worksheet_distance", Label: labels["trans_worksheet_distance"], FieldType: ct.TableFieldTypeNumber},
				{Name: "worksheet_repair", Label: labels["trans_worksheet_repair"], FieldType: ct.TableFieldTypeNumber},
				{Name: "worksheet_total", Label: labels["trans_worksheet_total"], FieldType: ct.TableFieldTypeNumber},
				{Name: "worksheet_justification", Label: labels["trans_worksheet_justification"]},
				{Name: "rent_holiday", Label: labels["trans_rent_holiday"], FieldType: ct.TableFieldTypeNumber},
				{Name: "rent_bad_tool", Label: labels["trans_rent_bad_tool"], FieldType: ct.TableFieldTypeNumber},
				{Name: "rent_other", Label: labels["trans_rent_other"], FieldType: ct.TableFieldTypeNumber},
				{Name: "rent_justification", Label: labels["trans_rent_justification"]},
				{Name: "auth_code", Label: labels["auth_code"]},
			},
			VisibleColumns: cu.IM{
				"code": true, "trans_date": true, "customer_name": true, "currency_code": true, "amount": true,
				"trans_type": true, "direction": true,
			},
			HideFilters: cu.IM{},
			Filters:     []any{},
		},
		"transitem_map": {
			Title:       labels["map_view"],
			Icon:        ct.IconDatabase,
			Simple:      false,
			ReadOnly:    true,
			LabelAdd:    "",
			HideFilters: cu.IM{},
			Fields: []ct.TableField{
				{Name: "code", Label: labels["trans_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "trans_type", Label: labels["trans_type"]},
				{Name: "direction", Label: labels["trans_direction"]},
				{Name: "trans_date", Label: labels["trans_date"]},
				{Name: "description", Label: labels["map_description"]},
				{Name: "value", Label: labels["map_value"], FieldType: ct.TableFieldTypeMeta},
			},
			VisibleColumns: cu.IM{
				"code": true, "direction": true, "trans_date": true, "description": true, "value": true,
			},
			Filters: []any{},
		},
		"transitem_item": {
			Title:       labels["items_view"],
			Icon:        ct.IconFileText,
			Simple:      false,
			ReadOnly:    true,
			LabelAdd:    "",
			HideFilters: cu.IM{},
			Fields: []ct.TableField{
				{Name: "item_code", Label: labels["item_code"]},
				{Name: "trans_code", Label: labels["trans_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "trans_date", Label: labels["trans_date"], FieldType: ct.TableFieldTypeDate},
				{Name: "product_code", Label: labels["product_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "description", Label: labels["item_description"]},
				{Name: "tax_code", Label: labels["tax_code"]},
				{Name: "unit", Label: labels["item_unit"]},
				{Name: "qty", Label: labels["item_qty"], FieldType: ct.TableFieldTypeNumber},
				{Name: "currency_code", Label: labels["currency_code"]},
				{Name: "fx_price", Label: labels["item_fx_price"], FieldType: ct.TableFieldTypeNumber},
				{Name: "discount", Label: labels["item_discount"], FieldType: ct.TableFieldTypeNumber},
				{Name: "net_amount", Label: labels["item_net_amount"], FieldType: ct.TableFieldTypeNumber},
				{Name: "vat_amount", Label: labels["item_vat_amount"], FieldType: ct.TableFieldTypeNumber},
				{Name: "amount", Label: labels["item_amount"], FieldType: ct.TableFieldTypeNumber},
				{Name: "own_stock", Label: labels["item_own_stock"], FieldType: ct.TableFieldTypeNumber},
				{Name: "deposit", Label: labels["item_deposit"], FieldType: ct.TableFieldTypeBool},
				{Name: "action_price", Label: labels["item_action_price"], FieldType: ct.TableFieldTypeBool},
				{Name: "tag_lst", Label: labels["price_tags"]},
			},
			VisibleColumns: cu.IM{
				"trans_code": true, "trans_date": true, "description": true, "qty": true, "currency_code": true,
				"fx_price": true, "net_amount": true, "discount": true, "vat_amount": true,
				"amount": true, "deposit": true,
			},
			Filters: []any{},
		},
		"transpayment_simple": {
			Title:    labels["quick_search"],
			Icon:     ct.IconBolt,
			Simple:   true,
			ReadOnly: false,
			LabelAdd: "",
			Fields: []ct.TableField{
				{Name: "code", Label: labels["trans_code"]},
				{Name: "trans_date", Label: labels["trans_date"]},
				{Name: "place_name", Label: labels["place_name_payment"]},
				{Name: "currency_code", Label: labels["currency_code"]},
				{Name: "amount", Label: labels["item_amount"], FieldType: ct.TableFieldTypeNumber},
				{Name: "trans_type", Label: labels["trans_type"]},
				{Name: "direction", Label: labels["trans_direction"]},
			},
			VisibleColumns: cu.IM{},
			HideFilters:    cu.IM{},
			Filters: []any{
				cu.IM{"or": true, "field": "t.code", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "trans_type", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "direction", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "trans_date", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "place_name", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "p.currency_code", "comp": "==", "value": ""},
			},
			FilterPlaceholder: strings.Join([]string{
				labels["trans_code"], labels["trans_date"],
				labels["place_name_payment"], labels["currency_code"], labels["trans_type"],
				labels["trans_direction"],
			}, ", "),
		},
		"transpayment": {
			Title:    labels["transpayment_view"],
			Icon:     ct.IconMoney,
			Simple:   false,
			ReadOnly: true,
			LabelAdd: labels["transpayment_new"],
			Fields: []ct.TableField{
				{Name: "code", Label: labels["payment_code"]},
				{Name: "trans_code", Label: labels["trans_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "trans_type", Label: labels["trans_type"]},
				{Name: "direction", Label: labels["trans_direction"]},
				{Name: "status", Label: labels["trans_status"]},
				{Name: "ref_number", Label: labels["trans_ref_number"]},
				{Name: "trans_date", Label: labels["trans_date"]},
				{Name: "paid_date", Label: labels["payment_paid_date"], FieldType: ct.TableFieldTypeDate},
				{Name: "place_code", Label: labels["place_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "place_name", Label: labels["place_name_payment"]},
				{Name: "currency_code", Label: labels["currency_code"]},
				{Name: "amount", Label: labels["payment_amount"], FieldType: ct.TableFieldTypeNumber},
				{Name: "description", Label: labels["payment_notes"]},
				{Name: "employee_code", Label: labels["employee_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "tag_lst", Label: labels["trans_tags"]},
				{Name: "trans_state", Label: labels["trans_state"]},
				{Name: "notes", Label: labels["trans_notes"]},
				{Name: "internal_notes", Label: labels["trans_internal_notes"]},
				{Name: "closed", Label: labels["trans_closed"], FieldType: ct.TableFieldTypeBool},
				{Name: "auth_code", Label: labels["auth_code"]},
			},
			VisibleColumns: cu.IM{
				"trans_code": true, "ref_number": true,
				"paid_date": true, "place_name": true, "currency_code": true, "amount": true, "description": true,
			},
			HideFilters: cu.IM{},
			Filters:     []any{},
		},
		"transpayment_map": {
			Title:       labels["map_view"],
			Icon:        ct.IconDatabase,
			Simple:      false,
			ReadOnly:    true,
			LabelAdd:    "",
			HideFilters: cu.IM{},
			Fields: []ct.TableField{
				{Name: "code", Label: labels["trans_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "trans_type", Label: labels["trans_type"]},
				{Name: "direction", Label: labels["trans_direction"]},
				{Name: "trans_date", Label: labels["trans_date"]},
				{Name: "description", Label: labels["map_description"]},
				{Name: "value", Label: labels["map_value"], FieldType: ct.TableFieldTypeMeta},
			},
			VisibleColumns: cu.IM{
				"code": true, "direction": true, "trans_date": true, "description": true, "value": true,
			},
			Filters: []any{},
		},
		"transpayment_invoice": {
			Title:    labels["payment_link_view_bank"],
			Icon:     ct.IconFileText,
			Simple:   false,
			ReadOnly: true,
			LabelAdd: "",
			Fields: []ct.TableField{
				{Name: "code", Label: labels["payment_code"]},
				{Name: "trans_code", Label: labels["trans_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "trans_type", Label: labels["trans_type"]},
				{Name: "direction", Label: labels["trans_direction"]},
				{Name: "paid_date", Label: labels["payment_paid_date"]},
				{Name: "place_name", Label: labels["place_name_payment"]},
				{Name: "currency_code", Label: labels["currency_code"]},
				{Name: "paid_amount", Label: labels["payment_amount"], FieldType: ct.TableFieldTypeNumber},
				{Name: "paid_rate", Label: labels["payment_rate"], FieldType: ct.TableFieldTypeNumber},
				{Name: "ref_trans_code", Label: labels["invoice_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "invoice_curr", Label: labels["invoice_curr"]},
				{Name: "invoice_amount", Label: labels["invoice_amount"], FieldType: ct.TableFieldTypeNumber},
				{Name: "description", Label: labels["payment_notes"]},
			},
			VisibleColumns: cu.IM{
				"trans_code": true, "paid_date": true, "place_name": true, "currency_code": true,
				"paid_amount": true, "ref_trans_code": true, "description": true,
			},
			HideFilters: cu.IM{},
			Filters:     []any{},
		},
		"transmovement_simple": {
			Title:    labels["quick_search"],
			Icon:     ct.IconBolt,
			Simple:   true,
			ReadOnly: false,
			LabelAdd: "",
			Fields: []ct.TableField{
				{Name: "code", Label: labels["trans_code"]},
				{Name: "trans_date", Label: labels["trans_date"]},
				{Name: "trans_type", Label: labels["trans_type"]},
				{Name: "direction", Label: labels["trans_direction"]},
			},
			VisibleColumns: cu.IM{},
			HideFilters:    cu.IM{},
			Filters: []any{
				cu.IM{"or": true, "field": "t.code", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "trans_type", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "direction", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "trans_date", "comp": "==", "value": ""},
			},
			FilterPlaceholder: strings.Join([]string{
				labels["trans_code"], labels["trans_date"], labels["trans_type"], labels["trans_direction"],
			}, ", "),
		},
		"transmovement_stock": {
			Title:    labels["stock_view"],
			Icon:     ct.IconCalendar,
			Simple:   false,
			ReadOnly: true,
			LabelAdd: "",
			Fields: []ct.TableField{
				{Name: "place_code", Label: labels["place_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "place_name", Label: labels["place_name_movement"]},
				{Name: "product_code", Label: labels["product_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "product_name", Label: labels["product_name"]},
				{Name: "unit", Label: labels["product_unit"]},
				{Name: "batch_no", Label: labels["movement_batchnumber"]},
				{Name: "qty", Label: labels["stock_qty"], FieldType: ct.TableFieldTypeNumber},
				{Name: "posdate", Label: labels["stock_posdate"], FieldType: ct.TableFieldTypeDate},
			},
			VisibleColumns: cu.IM{
				"place_name": true, "product_code": true, "product_name": true, "batch_no": true,
				"product_group": true, "qty": true, "posdate": true,
			},
			HideFilters: cu.IM{},
			Filters:     []any{},
		},
		"transmovement": {
			Title:    labels["transmovement_view"],
			Icon:     ct.IconTruck,
			Simple:   false,
			ReadOnly: true,
			LabelAdd: labels["transmovement_new"],
			Fields: []ct.TableField{
				{Name: "code", Label: labels["movement_code"]},
				{Name: "trans_code", Label: labels["trans_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "trans_type", Label: labels["trans_type"]},
				{Name: "direction", Label: labels["trans_direction"]},
				{Name: "shipping_date", Label: labels["movement_shipping_date"]},
				{Name: "place_code", Label: labels["place_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "place_name", Label: labels["place_name_movement"]},
				{Name: "product_code", Label: labels["product_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "product_name", Label: labels["product_name"]},
				{Name: "unit", Label: labels["product_unit"]},
				{Name: "batch_no", Label: labels["movement_batchnumber"]},
				{Name: "qty", Label: labels["movement_qty"], FieldType: ct.TableFieldTypeNumber},
				{Name: "customer_code", Label: labels["customer_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "customer_name", Label: labels["customer_name"]},
				{Name: "ref_trans_code", Label: labels["trans_trans_code"], FieldType: ct.TableFieldTypeLink},
			},
			VisibleColumns: cu.IM{
				"trans_code": true, "shipping_date": true, "place_name": true,
				"product_name": true, "batch_no": true, "qty": true, "customer_name": true,
			},
			HideFilters: cu.IM{},
			Filters:     []any{},
		},
		"transmovement_waybill": {
			Title:    labels["waybill_view"],
			Icon:     ct.IconBriefcase,
			Simple:   false,
			ReadOnly: true,
			LabelAdd: labels["trans_waybill_new"],
			Fields: []ct.TableField{
				{Name: "code", Label: labels["movement_code"]},
				{Name: "trans_code", Label: labels["trans_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "direction", Label: labels["trans_direction"]},
				{Name: "shipping_time", Label: labels["movement_shipping_time"], FieldType: ct.TableFieldTypeDateTime},
				{Name: "tool_code", Label: labels["tool_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "serial_number", Label: labels["tool_serial_number"]},
				{Name: "description", Label: labels["tool_description"]},
				{Name: "mvnotes", Label: labels["movement_notes"]},
				{Name: "ref_trans_code", Label: labels["trans_trans_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "employee_code", Label: labels["employee_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "customer_code", Label: labels["customer_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "customer_name", Label: labels["customer_name"]},
				{Name: "trans_state", Label: labels["trans_state"]},
				{Name: "notes", Label: labels["trans_notes"]},
				{Name: "internal_notes", Label: labels["trans_internal_notes"]},
				{Name: "closed", Label: labels["trans_closed"], FieldType: ct.TableFieldTypeBool},
			},
			VisibleColumns: cu.IM{
				"trans_code": true, "direction": true, "shipping_time": true, "serial_number": true, "description": true,
				"ref_trans_code": true, "employee_code": true, "customer_name": true,
			},
			HideFilters: cu.IM{},
			Filters:     []any{},
		},
		"transmovement_formula": {
			Title:    labels["formula_view"],
			Icon:     ct.IconMagic,
			Simple:   false,
			ReadOnly: true,
			LabelAdd: labels["trans_formula_new"],
			Fields: []ct.TableField{
				{Name: "code", Label: labels["movement_code"]},
				{Name: "trans_code", Label: labels["trans_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "direction", Label: labels["trans_direction"]},
				{Name: "product_code", Label: labels["product_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "product_name", Label: labels["product_name"]},
				{Name: "unit", Label: labels["product_unit"]},
				{Name: "place_code", Label: labels["place_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "place_name", Label: labels["place_name_movement"]},
				{Name: "batch_no", Label: labels["movement_batchnumber"]},
				{Name: "qty", Label: labels["movement_qty"], FieldType: ct.TableFieldTypeNumber},
				{Name: "shared", Label: labels["movement_shared"], FieldType: ct.TableFieldTypeBool},
			},
			VisibleColumns: cu.IM{
				"trans_code": true, "direction": true, "product_name": true, "place_name": true, "batch_no": true, "qty": true, "shared": true,
			},
			HideFilters: cu.IM{},
			Filters:     []any{},
		},
		"transmovement_map": {
			Title:       labels["map_view"],
			Icon:        ct.IconDatabase,
			Simple:      false,
			ReadOnly:    true,
			LabelAdd:    "",
			HideFilters: cu.IM{},
			Fields: []ct.TableField{
				{Name: "code", Label: labels["trans_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "trans_type", Label: labels["trans_type"]},
				{Name: "direction", Label: labels["trans_direction"]},
				{Name: "trans_date", Label: labels["trans_date"]},
				{Name: "description", Label: labels["map_description"]},
				{Name: "value", Label: labels["map_value"], FieldType: ct.TableFieldTypeMeta},
			},
			VisibleColumns: cu.IM{
				"code": true, "direction": true, "trans_date": true, "description": true, "value": true,
			},
			Filters: []any{},
		},
		"customer_simple": {
			Title:    labels["quick_search"],
			Icon:     ct.IconBolt,
			Simple:   true,
			ReadOnly: false,
			LabelAdd: "",
			Fields: []ct.TableField{
				{Name: "code", Label: labels["customer_code"]},
				{Name: "customer_name", Label: labels["customer_name"]},
				{Name: "tax_number", Label: labels["customer_tax_number"]},
				{Name: "customer_type", Label: labels["customer_type"]},
			},
			VisibleColumns: cu.IM{},
			HideFilters:    cu.IM{},
			Filters: []any{
				cu.IM{"or": true, "field": "code", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "customer_name", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "tax_number", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "customer_type", "comp": "==", "value": ""},
			},
			FilterPlaceholder: strings.Join([]string{
				labels["customer_code"], labels["customer_name"],
				labels["customer_tax_number"], labels["customer_type"]}, ", "),
		},
		"customer": {
			Title:    labels["customer_view"],
			Icon:     ct.IconUser,
			Simple:   false,
			ReadOnly: false,
			LabelAdd: labels["customer_new"],
			Fields: []ct.TableField{
				{Name: "code", Label: labels["customer_code"]},
				{Name: "customer_name", Label: labels["customer_name"]},
				{Name: "tax_number", Label: labels["customer_tax_number"]},
				{Name: "customer_type", Label: labels["customer_type"]},
				{Name: "account", Label: labels["customer_account"]},
				{Name: "tax_free", Label: labels["customer_tax_free"], FieldType: ct.TableFieldTypeBool},
				{Name: "terms", Label: labels["customer_terms"], FieldType: ct.TableFieldTypeNumber},
				{Name: "credit_limit", Label: labels["customer_credit_limit"], FieldType: ct.TableFieldTypeNumber},
				{Name: "discount", Label: labels["customer_discount"], FieldType: ct.TableFieldTypeNumber},
				{Name: "notes", Label: labels["customer_notes"]},
				{Name: "tag_lst", Label: labels["customer_tags"]},
				{Name: "inactive", Label: labels["customer_inactive"], FieldType: ct.TableFieldTypeBool},
			},
			VisibleColumns: cu.IM{
				"code": true, "customer_name": true, "customer_type": true, "tax_number": true, "tag_lst": true,
			},
			HideFilters: cu.IM{},
			Filters:     []any{
				//cu.IM{"field": "customer_name", "comp": "==", "value": ""},
			},
		},
		"customer_map": {
			Title:       labels["map_view"],
			Icon:        ct.IconDatabase,
			Simple:      false,
			ReadOnly:    true,
			LabelAdd:    "",
			HideFilters: cu.IM{},
			Fields: []ct.TableField{
				{Name: "code", Label: labels["customer_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "customer_name", Label: labels["customer_name"], FieldType: ct.TableFieldTypeLink},
				{Name: "description", Label: labels["map_description"]},
				{Name: "value", Label: labels["map_value"], FieldType: ct.TableFieldTypeMeta},
			},
			VisibleColumns: cu.IM{
				"customer_name": true, "description": true, "value": true,
			},
			Filters: []any{},
		},
		"customer_addresses": {
			Title:       labels["address_view"],
			Icon:        ct.IconHome,
			Simple:      false,
			ReadOnly:    true,
			LabelAdd:    "",
			HideFilters: cu.IM{},
			Fields: []ct.TableField{
				{Name: "code", Label: labels["customer_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "customer_name", Label: labels["customer_name"], FieldType: ct.TableFieldTypeLink},
				{Name: "country", Label: labels["address_country"]},
				{Name: "state", Label: labels["address_state"]},
				{Name: "zip_code", Label: labels["address_zip_code"]},
				{Name: "city", Label: labels["address_city"]},
				{Name: "street", Label: labels["address_street"]},
				{Name: "notes", Label: labels["address_notes"]},
			},
			VisibleColumns: cu.IM{
				"customer_name": true, "zip_code": true, "city": true, "street": true,
			},
			Filters: []any{},
		},
		"customer_contacts": {
			Title:       labels["contact_view"],
			Icon:        ct.IconMobile,
			Simple:      false,
			ReadOnly:    true,
			LabelAdd:    "",
			HideFilters: cu.IM{},
			Fields: []ct.TableField{
				{Name: "code", Label: labels["customer_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "customer_name", Label: labels["customer_name"], FieldType: ct.TableFieldTypeLink},
				{Name: "first_name", Label: labels["contact_first_name"]},
				{Name: "surname", Label: labels["contact_surname"]},
				{Name: "status", Label: labels["contact_status"]},
				{Name: "phone", Label: labels["contact_phone"]},
				{Name: "mobile", Label: labels["contact_mobile"]},
				{Name: "email", Label: labels["contact_email"]},
				{Name: "notes", Label: labels["contact_notes"]},
			},
			VisibleColumns: cu.IM{
				"customer_name": true, "first_name": true, "surname": true, "phone": true, "email": true,
			},
			Filters: []any{},
		},
		"customer_events": {
			Title:       labels["event_view"],
			Icon:        ct.IconCalendar,
			Simple:      false,
			ReadOnly:    true,
			LabelAdd:    "",
			HideFilters: cu.IM{},
			Fields: []ct.TableField{
				{Name: "code", Label: labels["customer_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "customer_name", Label: labels["customer_name"], FieldType: ct.TableFieldTypeLink},
				{Name: "subject", Label: labels["event_subject"]},
				{Name: "start_time", Label: labels["event_start_time"], FieldType: ct.TableFieldTypeDateTime},
				{Name: "end_time", Label: labels["event_end_time"], FieldType: ct.TableFieldTypeDateTime},
				{Name: "place", Label: labels["event_place"]},
				{Name: "description", Label: labels["event_description"]},
				{Name: "tag_lst", Label: labels["event_tags"]},
			},
			VisibleColumns: cu.IM{
				"customer_name": true, "subject": true, "start_time": true, "end_time": true, "place": true, "tag_lst": true,
			},
			Filters: []any{},
		},
		"product": {
			Title:    labels["product_view"],
			Icon:     ct.IconShoppingCart,
			Simple:   false,
			ReadOnly: false,
			LabelAdd: labels["product_new"],
			Fields: []ct.TableField{
				{Name: "code", Label: labels["product_code"]},
				{Name: "product_name", Label: labels["product_name"]},
				{Name: "product_type", Label: labels["product_type"]},
				{Name: "tax_code", Label: labels["tax_code"]},
				{Name: "unit", Label: labels["product_unit"]},
				{Name: "barcode_type", Label: labels["product_barcode_type"]},
				{Name: "barcode", Label: labels["product_barcode"]},
				{Name: "barcode_qty", Label: labels["product_barcode_qty"], FieldType: ct.TableFieldTypeNumber},
				{Name: "notes", Label: labels["product_notes"]},
				{Name: "tag_lst", Label: labels["product_tags"]},
				{Name: "inactive", Label: labels["product_inactive"], FieldType: ct.TableFieldTypeBool},
			},
			VisibleColumns: cu.IM{
				"code": true, "product_name": true, "product_type": true, "unit": true, "tag_lst": true,
			},
			HideFilters: cu.IM{},
			Filters:     []any{},
		},
		"product_simple": {
			Title:    labels["quick_search"],
			Icon:     ct.IconBolt,
			Simple:   true,
			ReadOnly: false,
			LabelAdd: "",
			Fields: []ct.TableField{
				{Name: "code", Label: labels["product_code"]},
				{Name: "product_name", Label: labels["product_name"]},
				{Name: "product_type", Label: labels["product_type"]},
				{Name: "tag_lst", Label: labels["product_tags"]},
			},
			VisibleColumns: cu.IM{},
			HideFilters:    cu.IM{},
			Filters: []any{
				cu.IM{"or": true, "field": "code", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "product_name", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "product_type", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "tag_lst", "comp": "==", "value": ""},
			},
			FilterPlaceholder: strings.Join([]string{
				labels["product_code"], labels["product_name"],
				labels["product_type"], labels["product_tags"]}, ", "),
		},
		"product_map": {
			Title:       labels["map_view"],
			Icon:        ct.IconDatabase,
			Simple:      false,
			ReadOnly:    true,
			LabelAdd:    "",
			HideFilters: cu.IM{},
			Fields: []ct.TableField{
				{Name: "code", Label: labels["product_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "product_name", Label: labels["product_name"], FieldType: ct.TableFieldTypeLink},
				{Name: "description", Label: labels["map_description"]},
				{Name: "value", Label: labels["map_value"], FieldType: ct.TableFieldTypeMeta},
			},
			VisibleColumns: cu.IM{
				"product_name": true, "description": true, "value": true,
			},
			Filters: []any{},
		},
		"product_events": {
			Title:       labels["event_view"],
			Icon:        ct.IconCalendar,
			Simple:      false,
			ReadOnly:    true,
			LabelAdd:    "",
			HideFilters: cu.IM{},
			Fields: []ct.TableField{
				{Name: "code", Label: labels["product_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "product_name", Label: labels["product_name"], FieldType: ct.TableFieldTypeLink},
				{Name: "subject", Label: labels["event_subject"]},
				{Name: "start_time", Label: labels["event_start_time"], FieldType: ct.TableFieldTypeDateTime},
				{Name: "end_time", Label: labels["event_end_time"], FieldType: ct.TableFieldTypeDateTime},
				{Name: "place", Label: labels["event_place"]},
				{Name: "description", Label: labels["event_description"]},
				{Name: "tag_lst", Label: labels["event_tags"]},
			},
			VisibleColumns: cu.IM{
				"product_name": true, "subject": true, "start_time": true, "end_time": true, "place": true, "tag_lst": true,
			},
			Filters: []any{},
		},
		"product_prices": {
			Title:       labels["price_view"],
			Icon:        ct.IconMoney,
			Simple:      false,
			ReadOnly:    true,
			LabelAdd:    "",
			HideFilters: cu.IM{},
			Fields: []ct.TableField{
				{Name: "price_code", Label: labels["price_code"]},
				{Name: "code", Label: labels["product_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "product_name", Label: labels["product_name"], FieldType: ct.TableFieldTypeLink},
				{Name: "price_type", Label: labels["price_type"]},
				{Name: "valid_from", Label: labels["price_valid_from"], FieldType: ct.TableFieldTypeDate},
				{Name: "valid_to", Label: labels["price_valid_to"], FieldType: ct.TableFieldTypeDate},
				{Name: "customer_code", Label: labels["customer_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "currency_code", Label: labels["currency_code"]},
				{Name: "qty", Label: labels["price_qty"], FieldType: ct.TableFieldTypeNumber},
				{Name: "price_value", Label: labels["price_value"], FieldType: ct.TableFieldTypeNumber},
				{Name: "tag_lst", Label: labels["price_tags"]},
			},
			VisibleColumns: cu.IM{
				"product_name": true, "price_type": true, "valid_from": true, "valid_to": true,
				"customer_code": true, "currency_code": true, "qty": true, "price_value": true,
			},
			Filters: []any{},
		},
		"product_components": {
			Title:       labels["product_component_view"],
			Icon:        ct.IconFlask,
			Simple:      false,
			ReadOnly:    true,
			LabelAdd:    "",
			HideFilters: cu.IM{},
			Fields: []ct.TableField{
				{Name: "product_code", Label: labels["product_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "product_name", Label: labels["product_name"]},
				{Name: "ref_product_code", Label: labels["product_component_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "component_name", Label: labels["product_component_name"]},
				{Name: "component_unit", Label: labels["product_unit"]},
				{Name: "component_type", Label: labels["product_type"]},
				{Name: "qty", Label: labels["product_component_qty"], FieldType: ct.TableFieldTypeNumber},
				{Name: "notes", Label: labels["product_component_notes"]},
			},
			VisibleColumns: cu.IM{
				"product_code": true, "product_name": true, "ref_product_code": true, "component_name": true,
				"component_unit": true, "component_type": true, "qty": true,
			},
			Filters: []any{},
		},
		"tool_simple": {
			Title:    labels["quick_search"],
			Icon:     ct.IconBolt,
			Simple:   true,
			ReadOnly: false,
			LabelAdd: "",
			Fields: []ct.TableField{
				{Name: "code", Label: labels["tool_code"]},
				{Name: "description", Label: labels["tool_description"]},
				{Name: "product_code", Label: labels["product_code"]},
				{Name: "serial_number", Label: labels["tool_serial_number"]},
				{Name: "tag_lst", Label: labels["tool_tags"]},
			},
			VisibleColumns: cu.IM{},
			HideFilters:    cu.IM{},
			Filters: []any{
				cu.IM{"or": true, "field": "code", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "description", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "product_code", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "serial_number", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "tag_lst", "comp": "==", "value": ""},
			},
			FilterPlaceholder: strings.Join([]string{
				labels["tool_code"], labels["tool_description"],
				labels["product_code"], labels["tool_serial_number"],
				labels["tool_tags"]}, ", "),
		},
		"tool": {
			Title:    labels["tool_view"],
			Icon:     ct.IconWrench,
			Simple:   false,
			ReadOnly: false,
			LabelAdd: labels["tool_new"],
			Fields: []ct.TableField{
				{Name: "code", Label: labels["tool_code"]},
				{Name: "description", Label: labels["tool_description"]},
				{Name: "product_code", Label: labels["product_code"]},
				{Name: "serial_number", Label: labels["tool_serial_number"]},
				{Name: "notes", Label: labels["tool_notes"]},
				{Name: "tag_lst", Label: labels["tool_tags"]},
				{Name: "inactive", Label: labels["tool_inactive"], FieldType: ct.TableFieldTypeBool},
			},
			VisibleColumns: cu.IM{
				"description": true, "product_code": true, "serial_number": true, "tag_lst": true,
			},
			HideFilters: cu.IM{},
			Filters:     []any{},
		},
		"tool_map": {
			Title:       labels["map_view"],
			Icon:        ct.IconDatabase,
			Simple:      false,
			ReadOnly:    true,
			LabelAdd:    "",
			HideFilters: cu.IM{},
			Fields: []ct.TableField{
				{Name: "code", Label: labels["tool_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "tool_description", Label: labels["tool_description"], FieldType: ct.TableFieldTypeLink},
				{Name: "description", Label: labels["map_description"]},
				{Name: "value", Label: labels["map_value"], FieldType: ct.TableFieldTypeMeta},
			},
			VisibleColumns: cu.IM{
				"tool_description": true, "description": true, "value": true,
			},
			Filters: []any{},
		},
		"tool_events": {
			Title:       labels["event_view"],
			Icon:        ct.IconCalendar,
			Simple:      false,
			ReadOnly:    true,
			LabelAdd:    "",
			HideFilters: cu.IM{},
			Fields: []ct.TableField{
				{Name: "code", Label: labels["tool_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "tool_description", Label: labels["tool_description"], FieldType: ct.TableFieldTypeLink},
				{Name: "subject", Label: labels["event_subject"]},
				{Name: "start_time", Label: labels["event_start_time"], FieldType: ct.TableFieldTypeDateTime},
				{Name: "end_time", Label: labels["event_end_time"], FieldType: ct.TableFieldTypeDateTime},
				{Name: "place", Label: labels["event_place"]},
				{Name: "description", Label: labels["event_description"]},
				{Name: "tag_lst", Label: labels["event_tags"]},
			},
			VisibleColumns: cu.IM{
				"tool_description": true, "subject": true, "start_time": true, "end_time": true, "place": true, "tag_lst": true,
			},
			Filters: []any{},
		},
		"project_simple": {
			Title:    labels["quick_search"],
			Icon:     ct.IconBolt,
			Simple:   true,
			ReadOnly: false,
			LabelAdd: "",
			Fields: []ct.TableField{
				{Name: "code", Label: labels["project_code"]},
				{Name: "project_name", Label: labels["project_name"]},
				{Name: "customer_code", Label: labels["customer_code"]},
			},
			VisibleColumns: cu.IM{},
			HideFilters:    cu.IM{},
			Filters: []any{
				cu.IM{"or": true, "field": "code", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "project_name", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "customer_code", "comp": "==", "value": ""},
			},
			FilterPlaceholder: strings.Join([]string{
				labels["project_code"], labels["project_name"],
				labels["customer_code"]}, ", "),
		},
		"project": {
			Title:    labels["project_view"],
			Icon:     ct.IconClock,
			Simple:   false,
			ReadOnly: false,
			LabelAdd: labels["project_new"],
			Fields: []ct.TableField{
				{Name: "code", Label: labels["project_code"]},
				{Name: "project_name", Label: labels["project_name"]},
				{Name: "customer_code", Label: labels["customer_code"]},
				{Name: "start_date", Label: labels["project_start_date"], FieldType: ct.TableFieldTypeDate},
				{Name: "end_date", Label: labels["project_end_date"], FieldType: ct.TableFieldTypeDate},
				{Name: "notes", Label: labels["project_notes"]},
				{Name: "tag_lst", Label: labels["project_tags"]},
				{Name: "inactive", Label: labels["project_inactive"], FieldType: ct.TableFieldTypeBool},
			},
			VisibleColumns: cu.IM{
				"code": true, "project_name": true, "customer_code": true, "start_date": true, "end_date": true, "tag_lst": true,
			},
			HideFilters: cu.IM{},
			Filters:     []any{},
		},
		"project_map": {
			Title:       labels["map_view"],
			Icon:        ct.IconDatabase,
			Simple:      false,
			ReadOnly:    true,
			LabelAdd:    "",
			HideFilters: cu.IM{},
			Fields: []ct.TableField{
				{Name: "code", Label: labels["project_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "project_name", Label: labels["project_name"], FieldType: ct.TableFieldTypeLink},
				{Name: "description", Label: labels["map_description"]},
				{Name: "value", Label: labels["map_value"], FieldType: ct.TableFieldTypeMeta},
			},
			VisibleColumns: cu.IM{
				"project_name": true, "description": true, "value": true,
			},
			Filters: []any{},
		},
		"project_addresses": {
			Title:       labels["address_view"],
			Icon:        ct.IconHome,
			Simple:      false,
			ReadOnly:    true,
			LabelAdd:    "",
			HideFilters: cu.IM{},
			Fields: []ct.TableField{
				{Name: "code", Label: labels["project_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "project_name", Label: labels["project_name"], FieldType: ct.TableFieldTypeLink},
				{Name: "country", Label: labels["address_country"]},
				{Name: "state", Label: labels["address_state"]},
				{Name: "zip_code", Label: labels["address_zip_code"]},
				{Name: "city", Label: labels["address_city"]},
				{Name: "street", Label: labels["address_street"]},
				{Name: "notes", Label: labels["address_notes"]},
			},
			VisibleColumns: cu.IM{
				"project_name": true, "zip_code": true, "city": true, "street": true,
			},
			Filters: []any{},
		},
		"project_contacts": {
			Title:       labels["contact_view"],
			Icon:        ct.IconMobile,
			Simple:      false,
			ReadOnly:    true,
			LabelAdd:    "",
			HideFilters: cu.IM{},
			Fields: []ct.TableField{
				{Name: "code", Label: labels["project_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "project_name", Label: labels["project_name"], FieldType: ct.TableFieldTypeLink},
				{Name: "first_name", Label: labels["contact_first_name"]},
				{Name: "surname", Label: labels["contact_surname"]},
				{Name: "status", Label: labels["contact_status"]},
				{Name: "phone", Label: labels["contact_phone"]},
				{Name: "mobile", Label: labels["contact_mobile"]},
				{Name: "email", Label: labels["contact_email"]},
				{Name: "notes", Label: labels["contact_notes"]},
			},
			VisibleColumns: cu.IM{
				"project_name": true, "first_name": true, "surname": true, "phone": true, "email": true,
			},
			Filters: []any{},
		},
		"project_events": {
			Title:       labels["event_view"],
			Icon:        ct.IconCalendar,
			Simple:      false,
			ReadOnly:    true,
			LabelAdd:    "",
			HideFilters: cu.IM{},
			Fields: []ct.TableField{
				{Name: "code", Label: labels["project_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "project_name", Label: labels["project_name"], FieldType: ct.TableFieldTypeLink},
				{Name: "subject", Label: labels["event_subject"]},
				{Name: "start_time", Label: labels["event_start_time"], FieldType: ct.TableFieldTypeDateTime},
				{Name: "end_time", Label: labels["event_end_time"], FieldType: ct.TableFieldTypeDateTime},
				{Name: "place", Label: labels["event_place"]},
				{Name: "description", Label: labels["event_description"]},
				{Name: "tag_lst", Label: labels["event_tags"]},
			},
			VisibleColumns: cu.IM{
				"project_name": true, "subject": true, "start_time": true, "end_time": true, "place": true, "tag_lst": true,
			},
			Filters: []any{},
		},
		"employee_simple": {
			Title:    labels["quick_search"],
			Icon:     ct.IconBolt,
			Simple:   true,
			ReadOnly: false,
			LabelAdd: "",
			Fields: []ct.TableField{
				{Name: "code", Label: labels["employee_code"]},
				{Name: "first_name", Label: labels["contact_first_name"]},
				{Name: "surname", Label: labels["contact_surname"]},
				{Name: "status", Label: labels["contact_status"]},
				{Name: "email", Label: labels["contact_email"]},
				{Name: "tag_lst", Label: labels["employee_tags"]},
			},
			VisibleColumns: cu.IM{},
			HideFilters:    cu.IM{},
			Filters: []any{
				cu.IM{"or": true, "field": "code", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "first_name", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "surname", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "status", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "email", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "tag_lst", "comp": "==", "value": ""},
			},
			FilterPlaceholder: strings.Join([]string{
				labels["employee_code"], labels["contact_first_name"],
				labels["contact_surname"], labels["contact_status"],
				labels["contact_email"], labels["employee_tags"]}, ", "),
		},
		"employee": {
			Title:    labels["employee_view"],
			Icon:     ct.IconMale,
			Simple:   false,
			ReadOnly: false,
			LabelAdd: labels["employee_new"],
			Fields: []ct.TableField{
				{Name: "code", Label: labels["employee_code"]},
				{Name: "start_date", Label: labels["employee_start_date"], FieldType: ct.TableFieldTypeDate},
				{Name: "end_date", Label: labels["employee_end_date"], FieldType: ct.TableFieldTypeDate},
				{Name: "first_name", Label: labels["contact_first_name"]},
				{Name: "surname", Label: labels["contact_surname"]},
				{Name: "status", Label: labels["contact_status"]},
				{Name: "email", Label: labels["contact_email"]},
				{Name: "phone", Label: labels["contact_phone"]},
				{Name: "mobile", Label: labels["contact_mobile"]},
				{Name: "country", Label: labels["address_country"]},
				{Name: "state", Label: labels["address_state"]},
				{Name: "zip_code", Label: labels["address_zip_code"]},
				{Name: "city", Label: labels["address_city"]},
				{Name: "street", Label: labels["address_street"]},
				{Name: "notes", Label: labels["employee_notes"]},
				{Name: "tag_lst", Label: labels["employee_tags"]},
				{Name: "inactive", Label: labels["employee_inactive"], FieldType: ct.TableFieldTypeBool},
			},
			VisibleColumns: cu.IM{
				"code": true, "first_name": true, "surname": true, "status": true, "email": true, "phone": true, "mobile": true, "tag_lst": true,
			},
			HideFilters: cu.IM{},
			Filters:     []any{},
		},
		"employee_map": {
			Title:       labels["map_view"],
			Icon:        ct.IconDatabase,
			Simple:      false,
			ReadOnly:    true,
			LabelAdd:    "",
			HideFilters: cu.IM{},
			Fields: []ct.TableField{
				{Name: "code", Label: labels["employee_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "first_name", Label: labels["contact_first_name"]},
				{Name: "surname", Label: labels["contact_surname"]},
				{Name: "description", Label: labels["map_description"]},
				{Name: "value", Label: labels["map_value"], FieldType: ct.TableFieldTypeMeta},
			},
			VisibleColumns: cu.IM{
				"code": true, "first_name": true, "surname": true, "description": true, "value": true,
			},
			Filters: []any{},
		},
		"employee_events": {
			Title:       labels["event_view"],
			Icon:        ct.IconCalendar,
			Simple:      false,
			ReadOnly:    true,
			LabelAdd:    "",
			HideFilters: cu.IM{},
			Fields: []ct.TableField{
				{Name: "code", Label: labels["employee_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "first_name", Label: labels["contact_first_name"]},
				{Name: "surname", Label: labels["contact_surname"]},
				{Name: "subject", Label: labels["event_subject"]},
				{Name: "start_time", Label: labels["event_start_time"], FieldType: ct.TableFieldTypeDateTime},
				{Name: "end_time", Label: labels["event_end_time"], FieldType: ct.TableFieldTypeDateTime},
				{Name: "place", Label: labels["event_place"]},
				{Name: "description", Label: labels["event_description"]},
				{Name: "tag_lst", Label: labels["event_tags"]},
			},
			VisibleColumns: cu.IM{
				"code": true, "first_name": true, "surname": true, "subject": true, "start_time": true, "place": true, "tag_lst": true,
			},
			Filters: []any{},
		},
		"place_simple": {
			Title:    labels["quick_search"],
			Icon:     ct.IconBolt,
			Simple:   true,
			ReadOnly: false,
			LabelAdd: "",
			Fields: []ct.TableField{
				{Name: "code", Label: labels["place_code"]},
				{Name: "place_name", Label: labels["place_name"]},
				{Name: "place_type", Label: labels["place_type"]},
			},
			VisibleColumns: cu.IM{},
			HideFilters:    cu.IM{},
			Filters: []any{
				cu.IM{"or": true, "field": "code", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "place_name", "comp": "==", "value": ""},
				cu.IM{"or": true, "field": "place_type", "comp": "==", "value": ""},
			},
			FilterPlaceholder: strings.Join([]string{
				labels["place_code"], labels["place_name"],
				labels["place_type"]}, ", "),
		},
		"place": {
			Title:    labels["place_view"],
			Icon:     ct.IconHome,
			Simple:   false,
			ReadOnly: false,
			LabelAdd: labels["place_new"],
			Fields: []ct.TableField{
				{Name: "code", Label: labels["place_code"]},
				{Name: "place_name", Label: labels["place_name"]},
				{Name: "place_type", Label: labels["place_type"]},
				{Name: "currency_code", Label: labels["currency_code"]},
				{Name: "country", Label: labels["address_country"]},
				{Name: "state", Label: labels["address_state"]},
				{Name: "zip_code", Label: labels["address_zip_code"]},
				{Name: "city", Label: labels["address_city"]},
				{Name: "street", Label: labels["address_street"]},
				{Name: "notes", Label: labels["place_notes"]},
				{Name: "tag_lst", Label: labels["place_tags"]},
				{Name: "inactive", Label: labels["place_inactive"], FieldType: ct.TableFieldTypeBool},
			},
			VisibleColumns: cu.IM{
				"code": true, "place_name": true, "place_type": true, "currency_code": true, "tag_lst": true,
			},
			HideFilters: cu.IM{},
			Filters:     []any{
				//cu.IM{"field": "customer_name", "comp": "==", "value": ""},
			},
		},
		"place_map": {
			Title:       labels["map_view"],
			Icon:        ct.IconDatabase,
			Simple:      false,
			ReadOnly:    true,
			LabelAdd:    "",
			HideFilters: cu.IM{},
			Fields: []ct.TableField{
				{Name: "code", Label: labels["place_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "place_name", Label: labels["place_name"], FieldType: ct.TableFieldTypeLink},
				{Name: "description", Label: labels["map_description"]},
				{Name: "value", Label: labels["map_value"], FieldType: ct.TableFieldTypeMeta},
			},
			VisibleColumns: cu.IM{
				"place_name": true, "description": true, "value": true,
			},
			Filters: []any{},
		},
		"place_contacts": {
			Title:       labels["contact_view"],
			Icon:        ct.IconMobile,
			Simple:      false,
			ReadOnly:    true,
			LabelAdd:    "",
			HideFilters: cu.IM{},
			Fields: []ct.TableField{
				{Name: "code", Label: labels["place_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "place_name", Label: labels["place_name"], FieldType: ct.TableFieldTypeLink},
				{Name: "first_name", Label: labels["contact_first_name"]},
				{Name: "surname", Label: labels["contact_surname"]},
				{Name: "status", Label: labels["contact_status"]},
				{Name: "phone", Label: labels["contact_phone"]},
				{Name: "mobile", Label: labels["contact_mobile"]},
				{Name: "email", Label: labels["contact_email"]},
				{Name: "notes", Label: labels["contact_notes"]},
			},
			VisibleColumns: cu.IM{
				"place_name": true, "first_name": true, "surname": true, "phone": true, "email": true,
			},
			Filters: []any{},
		},
		"place_events": {
			Title:       labels["event_view"],
			Icon:        ct.IconCalendar,
			Simple:      false,
			ReadOnly:    true,
			LabelAdd:    "",
			HideFilters: cu.IM{},
			Fields: []ct.TableField{
				{Name: "code", Label: labels["place_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "place_name", Label: labels["place_name"], FieldType: ct.TableFieldTypeLink},
				{Name: "subject", Label: labels["event_subject"]},
				{Name: "start_time", Label: labels["event_start_time"], FieldType: ct.TableFieldTypeDateTime},
				{Name: "end_time", Label: labels["event_end_time"], FieldType: ct.TableFieldTypeDateTime},
				{Name: "place", Label: labels["event_place"]},
				{Name: "description", Label: labels["event_description"]},
				{Name: "tag_lst", Label: labels["event_tags"]},
			},
			VisibleColumns: cu.IM{
				"place_name": true, "subject": true, "start_time": true, "end_time": true, "place": true, "tag_lst": true,
			},
			Filters: []any{},
		},
		"office_queue": {
			Title: labels["office_queue_title"],
			Icon:  ct.IconPrint,
		},
		"office_template_editor": {
			Title:       labels["office_template_editor_title"],
			Icon:        ct.IconTextHeight,
			Simple:      false,
			ReadOnly:    true,
			LabelAdd:    "",
			HideFilters: cu.IM{},
			Fields: []ct.TableField{
				{Name: "code", Label: labels["template_code"],
					Column: &ct.TableColumn{Id: "code", Header: labels["template_code"],
						Cell: s.CustomTemplateCell(sessionID)}},
				{Name: "report_name", Label: labels["template_report_name"]},
				{Name: "report_type", Label: labels["template_report_type"]},
				{Name: "trans_type", Label: labels["template_trans_type"]},
				{Name: "direction", Label: labels["template_direction"]},
				{Name: "description", Label: labels["template_description"]},
				{Name: "label", Label: labels["template_label"]},
			},
			VisibleColumns: cu.IM{
				"code": true, "report_type": true, "trans_type": true, "direction": true, "report_name": true, "label": true,
			},
			Filters: []any{},
		},
		"office_rate": {
			Title:       labels["office_rate_title"],
			Icon:        ct.IconGlobe,
			Simple:      false,
			ReadOnly:    false,
			LabelAdd:    labels["rate_new"],
			HideFilters: cu.IM{},
			Fields: []ct.TableField{
				{Name: "code", Label: labels["rate_code"]},
				{Name: "rate_type", Label: labels["rate_type"]},
				{Name: "rate_date", Label: labels["rate_date"], FieldType: ct.TableFieldTypeDate},
				{Name: "place_code", Label: labels["place_code"]},
				{Name: "place_name", Label: labels["rate_account"]},
				{Name: "currency_code", Label: labels["currency_code"]},
				{Name: "rate_value", Label: labels["rate_value"], FieldType: ct.TableFieldTypeNumber},
				{Name: "tag_lst", Label: labels["rate_tags"]},
			},
			VisibleColumns: cu.IM{
				"code": true, "rate_type": true, "rate_date": true, "place_name": true, "currency_code": true, "rate_value": true, "tag_lst": true,
			},
			Filters: []any{},
		},
		"office_shortcut": {
			Title: labels["office_shortcut_title"],
			Icon:  ct.IconShare,
		},
		"office_log": {
			Title:    labels["office_log_title"],
			Icon:     ct.IconInfoCircle,
			Disabled: true,
		},
	}
	return viewMap[view]
}

func (s *SearchConfig) CustomTemplateCell(sessionID string) func(row cu.IM, col ct.TableColumn, value any, rowIndex int64) template.HTML {
	return func(row cu.IM, col ct.TableColumn, value any, rowIndex int64) template.HTML {
		lnk := ct.Link{
			LinkStyle: ct.LinkStyleDefault,
			Label:     cu.ToString(row["code"], ""),
			//Icon:       ct.IconEdit,
			Href:       fmt.Sprintf("/editor/?session=%s&code=%s", sessionID, cu.ToString(row["code"], "")),
			LinkTarget: "_blank",
		}
		res, _ := lnk.Render()
		return res
	}
}

func (s *SearchConfig) Query(key string, params cu.IM) (query md.Query) {
	qMap := map[string]func(editor string) md.Query{
		"transitem_simple": func(editor string) md.Query {
			return md.Query{
				Fields: []string{
					"t.*", "COALESCE(c.customer_name, '') as customer_name", "COALESCE(i.amount, 0) as amount",
					"'" + cu.ToString(editor, "trans") + "' as editor", "t.id as editor_id", "'trans' as editor_view"},
				From: `trans_view t left join customer c on t.customer_code = c.code  
			left join(select trans_code, sum(amount) as amount from item_view group by trans_code) i on t.code = i.trans_code`,
				Filter: fmt.Sprintf("trans_type in('%s','%s','%s','%s','%s','%s')",
					md.TransTypeInvoice.String(), md.TransTypeReceipt.String(), md.TransTypeOrder.String(),
					md.TransTypeOffer.String(), md.TransTypeWorksheet.String(), md.TransTypeRent.String()),
				OrderBy: []string{"t.id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"invoice_simple": func(editor string) md.Query {
			return md.Query{
				Fields: []string{
					"t.*", "COALESCE(c.customer_name, '') as customer_name", "COALESCE(i.amount, 0) as amount",
					"'" + cu.ToString(editor, "trans") + "' as editor", "t.id as editor_id", "'trans' as editor_view"},
				From: `trans_view t left join customer c on t.customer_code = c.code  
			left join(select trans_code, sum(amount) as amount from item_view group by trans_code) i on t.code = i.trans_code`,
				Filter: fmt.Sprintf("trans_type in('%s','%s')",
					md.TransTypeInvoice.String(), md.TransTypeReceipt.String(),
				),
				OrderBy: []string{"t.id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"transitem": func(editor string) md.Query {
			return md.Query{
				Fields: []string{
					"t.*", "COALESCE(c.customer_name, '') as customer_name", "COALESCE(i.amount, 0) as amount",
					"'" + cu.ToString(editor, "trans") + "' as editor", "t.id as editor_id", "'trans' as editor_view"},
				From: `trans_view t left join customer c on t.customer_code = c.code  
			left join(select trans_code, sum(amount) as amount from item_view group by trans_code) i on t.code = i.trans_code`,
				Filter: fmt.Sprintf("trans_type in('%s','%s','%s','%s','%s','%s')",
					md.TransTypeInvoice.String(), md.TransTypeReceipt.String(), md.TransTypeOrder.String(),
					md.TransTypeOffer.String(), md.TransTypeWorksheet.String(), md.TransTypeRent.String()),
				OrderBy: []string{"t.id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"transitem_map": func(editor string) md.Query {
			return md.Query{
				Fields: []string{"*", "map_key as field_name", "map_value as value",
					"id as trans_id", "'trans' as editor", "'maps' as editor_view"},
				From: "trans_map",
				Filter: fmt.Sprintf("trans_type in('%s','%s','%s','%s','%s','%s')",
					md.TransTypeInvoice.String(), md.TransTypeReceipt.String(), md.TransTypeOrder.String(),
					md.TransTypeOffer.String(), md.TransTypeWorksheet.String(), md.TransTypeRent.String()),
				OrderBy: []string{"id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"transitem_item": func(editor string) md.Query {
			return md.Query{
				Fields: []string{"i.*", "i.code as item_code", "t.trans_date", "t.currency_code",
					"t.id as trans_id", "'" + cu.ToString(editor, "trans") + "' as editor", "'items' as editor_view"},
				From:    "item_view i inner join trans t on i.trans_code = t.code",
				OrderBy: []string{"i.id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"transpayment_simple": func(editor string) md.Query {
			return md.Query{
				Fields: []string{
					"t.id", "t.code", "t.trans_date", "t.trans_type", "t.direction", "p.currency_code",
					"p.place_name as place_name", "COALESCE(pm.amount, 0) as amount",
					"'" + cu.ToString(editor, "trans") + "' as editor", "t.id as editor_id", "'trans' as editor_view"},
				From: `trans_view t inner join place p on t.place_code = p.code  
			left join(select trans_code, sum(amount) as amount from payment_view group by trans_code) pm on t.code = pm.trans_code`,
				Filter: fmt.Sprintf("trans_type in('%s','%s')",
					md.TransTypeBank.String(), md.TransTypeCash.String()),
				OrderBy: []string{"t.id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"transpayment": func(editor string) md.Query {
			return md.Query{
				Fields: []string{
					"pm.id", "pm.code", "pm.trans_code", "t.trans_type", "t.direction", "t.status", "t.ref_number", "t.trans_date", "pm.paid_date",
					"t.place_code", "p.place_name", "p.currency_code", "pm.amount", "pm.notes as description",
					"t.employee_code", "t.tag_lst", "t.trans_state", "t.notes", "t.internal_notes", "t.auth_code",
					"t.closed", "'" + cu.ToString(editor, "trans") + "' as editor", "t.id as editor_id", "'trans' as editor_view"},
				From: `trans_view t inner join place p on t.place_code = p.code  
			inner join payment_view pm on t.code = pm.trans_code`,
				Filter: fmt.Sprintf("trans_type in('%s','%s')",
					md.TransTypeBank.String(), md.TransTypeCash.String()),
				OrderBy: []string{"t.id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"transpayment_map": func(editor string) md.Query {
			return md.Query{
				Fields: []string{"*", "map_key as field_name", "map_value as value",
					"id as trans_id", "'" + cu.ToString(editor, "trans") + "' as editor", "'maps' as editor_view"},
				From: "trans_map",
				Filter: fmt.Sprintf("trans_type in('%s','%s')",
					md.TransTypeBank.String(), md.TransTypeCash.String()),
				OrderBy: []string{"id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"transpayment_invoice": func(editor string) md.Query {
			return md.Query{
				Fields:  []string{"*", "'" + cu.ToString(editor, "trans") + "' as editor", "id as editor_id", "'payment_link' as editor_view"},
				From:    `payment_invoice`,
				OrderBy: []string{"id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"transmovement_simple": func(editor string) md.Query {
			return md.Query{
				Fields: []string{
					"t.id", "t.code", "t.trans_date", "t.trans_type", "t.direction",
					"'" + cu.ToString(editor, "trans") + "' as editor", "t.id as editor_id", "'trans' as editor_view"},
				From: `trans_view t`,
				Filter: fmt.Sprintf("trans_type in('%s','%s','%s','%s','%s')",
					md.TransTypeDelivery.String(), md.TransTypeInventory.String(), md.TransTypeWaybill.String(),
					md.TransTypeProduction.String(), md.TransTypeFormula.String()),
				OrderBy: []string{"t.id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"transmovement_stock": func(editor string) md.Query {
			return md.Query{
				Fields:  []string{"*"},
				From:    `movement_stock`,
				OrderBy: []string{},
				Limit:   st.BrowserRowLimit,
			}
		},
		"transmovement": func(editor string) md.Query {
			return md.Query{
				Fields:  []string{"*"},
				From:    `movement_inventory`,
				OrderBy: []string{},
				Limit:   st.BrowserRowLimit,
			}
		},
		"transmovement_waybill": func(editor string) md.Query {
			return md.Query{
				Fields:  []string{"*"},
				From:    `movement_waybill`,
				OrderBy: []string{},
				Limit:   st.BrowserRowLimit,
			}
		},
		"transmovement_formula": func(editor string) md.Query {
			return md.Query{
				Fields:  []string{"*"},
				From:    `movement_formula`,
				OrderBy: []string{},
				Limit:   st.BrowserRowLimit,
			}
		},
		"transmovement_map": func(editor string) md.Query {
			return md.Query{
				Fields: []string{"*", "map_key as field_name", "map_value as value",
					"id as trans_id", "'" + cu.ToString(editor, "trans") + "' as editor", "'maps' as editor_view"},
				From: "trans_map",
				Filter: fmt.Sprintf("trans_type in('%s','%s','%s','%s','%s')",
					md.TransTypeDelivery.String(), md.TransTypeInventory.String(), md.TransTypeWaybill.String(),
					md.TransTypeProduction.String(), md.TransTypeFormula.String()),
				OrderBy: []string{"id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"customer_simple": func(editor string) md.Query {
			return md.Query{
				Fields: []string{
					"c.*", "'" + cu.ToString(editor, "customer") + "' as editor", "c.id as editor_id", "'customer' as editor_view"},
				From:    "customer_view c",
				OrderBy: []string{"c.id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"customer": func(editor string) md.Query {
			return md.Query{
				Fields: []string{
					"c.*", "'" + cu.ToString(editor, "customer") + "' as editor", "c.id as editor_id", "'customer' as editor_view"},
				From:    "customer_view c",
				OrderBy: []string{"c.id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"customer_map": func(editor string) md.Query {
			return md.Query{
				Fields: []string{"*", "map_key as field_name", "map_value as value",
					"id as customer_id", "'" + cu.ToString(editor, "customer") + "' as editor", "'maps' as editor_view"},
				From:    "customer_map",
				OrderBy: []string{"id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"customer_addresses": func(editor string) md.Query {
			return md.Query{
				Fields:  []string{"c.*", "c.id as customer_id", "'" + cu.ToString(editor, "customer") + "' as editor", "'addresses' as editor_view"},
				From:    "customer_addresses c",
				OrderBy: []string{"c.id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"customer_contacts": func(editor string) md.Query {
			return md.Query{
				Fields:  []string{"c.*", "c.id as customer_id", "'" + cu.ToString(editor, "customer") + "' as editor", "'contacts' as editor_view"},
				From:    "customer_contacts c",
				OrderBy: []string{"c.id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"customer_events": func(editor string) md.Query {
			return md.Query{
				Fields: []string{"c.*", "c.start_time as start_date", "c.end_time as end_date",
					"c.id as customer_id", "'" + cu.ToString(editor, "customer") + "' as editor", "'events' as editor_view"},
				From:    "customer_events c",
				OrderBy: []string{"c.id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"product_simple": func(editor string) md.Query {
			return md.Query{
				Fields: []string{
					"p.*", "'" + cu.ToString(editor, "product") + "' as editor", "p.id as editor_id", "'product' as editor_view"},
				From:    "product_view p",
				OrderBy: []string{"p.id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"product": func(editor string) md.Query {
			return md.Query{
				Fields: []string{
					"p.*", "'" + cu.ToString(editor, "product") + "' as editor", "p.id as editor_id", "'product' as editor_view"},
				From:    "product_view p",
				OrderBy: []string{"p.id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"product_map": func(editor string) md.Query {
			return md.Query{
				Fields: []string{"*", "map_key as field_name", "map_value as value",
					"id as product_id", "'" + cu.ToString(editor, "product") + "' as editor", "'maps' as editor_view"},
				From:    "product_map",
				OrderBy: []string{"id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"product_events": func(editor string) md.Query {
			return md.Query{
				Fields: []string{"p.*", "p.start_time as start_date", "p.end_time as end_date",
					"p.id as product_id", "'" + cu.ToString(editor, "product") + "' as editor", "'events' as editor_view"},
				From:    "product_events p",
				OrderBy: []string{"p.id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"product_prices": func(editor string) md.Query {
			return md.Query{
				Fields: []string{"pr.*", "p.code as code", "pr.code as price_code", "p.product_name",
					"p.id as product_id", "'" + cu.ToString(editor, "product") + "' as editor", "'prices' as editor_view"},
				From:    "price_view pr inner join product p on pr.product_code = p.code",
				OrderBy: []string{"p.id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"product_components": func(editor string) md.Query {
			return md.Query{
				Fields: []string{"*",
					"'" + cu.ToString(editor, "product") + "' as editor", "'components' as editor_view"},
				From:    "product_components",
				OrderBy: []string{"id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"tool_simple": func(editor string) md.Query {
			return md.Query{
				Fields: []string{
					"t.*", "'" + cu.ToString(editor, "tool") + "' as editor", "t.id as editor_id", "'tool' as editor_view"},
				From:    "tool_view t",
				OrderBy: []string{"t.id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"tool": func(editor string) md.Query {
			return md.Query{
				Fields: []string{
					"t.*", "'" + cu.ToString(editor, "tool") + "' as editor", "t.id as editor_id", "'tool' as editor_view"},
				From:    "tool_view t",
				OrderBy: []string{"t.id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"tool_map": func(editor string) md.Query {
			return md.Query{
				Fields: []string{"*", "map_key as field_name", "map_value as value",
					"id as tool_id", "'" + cu.ToString(editor, "tool") + "' as editor", "'maps' as editor_view"},
				From:    "tool_map",
				OrderBy: []string{"id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"tool_events": func(editor string) md.Query {
			return md.Query{
				Fields: []string{"t.*", "t.start_time as start_date", "t.end_time as end_date",
					"t.id as tool_id", "'" + cu.ToString(editor, "tool") + "' as editor", "'events' as editor_view"},
				From:    "tool_events t",
				OrderBy: []string{"t.id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"project_simple": func(editor string) md.Query {
			return md.Query{
				Fields: []string{
					"p.*", "'" + cu.ToString(editor, "project") + "' as editor", "p.id as editor_id", "'project' as editor_view"},
				From:    "project_view p",
				OrderBy: []string{"p.id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"project": func(editor string) md.Query {
			return md.Query{
				Fields: []string{
					"p.*", "'" + cu.ToString(editor, "project") + "' as editor", "p.id as editor_id", "'project' as editor_view"},
				From:    "project_view p",
				OrderBy: []string{"p.id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"project_map": func(editor string) md.Query {
			return md.Query{
				Fields: []string{"*", "map_key as field_name", "map_value as value",
					"id as project_id", "'" + cu.ToString(editor, "project") + "' as editor", "'maps' as editor_view"},
				From:    "project_map",
				OrderBy: []string{"id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"project_addresses": func(editor string) md.Query {
			return md.Query{
				Fields:  []string{"p.*", "p.id as project_id", "'" + cu.ToString(editor, "project") + "' as editor", "'addresses' as editor_view"},
				From:    "project_addresses p",
				OrderBy: []string{"p.id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"project_contacts": func(editor string) md.Query {
			return md.Query{
				Fields:  []string{"p.*", "p.id as project_id", "'" + cu.ToString(editor, "project") + "' as editor", "'contacts' as editor_view"},
				From:    "project_contacts p",
				OrderBy: []string{"p.id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"project_events": func(editor string) md.Query {
			return md.Query{
				Fields: []string{"p.*", "p.start_time as start_date", "p.end_time as end_date",
					"p.id as project_id", "'" + cu.ToString(editor, "project") + "' as editor", "'events' as editor_view"},
				From:    "project_events p",
				OrderBy: []string{"p.id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"employee_simple": func(editor string) md.Query {
			return md.Query{
				Fields: []string{
					"e.*", "'" + cu.ToString(editor, "employee") + "' as editor", "e.id as editor_id", "'employee' as editor_view"},
				From:    "employee_view e",
				OrderBy: []string{"e.id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"employee": func(editor string) md.Query {
			return md.Query{
				Fields: []string{
					"e.*", "'" + cu.ToString(editor, "employee") + "' as editor", "e.id as editor_id", "'employee' as editor_view"},
				From:    "employee_view e",
				OrderBy: []string{"e.id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"employee_map": func(editor string) md.Query {
			return md.Query{
				Fields: []string{"*", "map_key as field_name", "map_value as value",
					"id as employee_id", "'" + cu.ToString(editor, "employee") + "' as editor", "'maps' as editor_view"},
				From:    "employee_map",
				OrderBy: []string{"id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"employee_events": func(editor string) md.Query {
			return md.Query{
				Fields: []string{"e.*", "e.start_time as start_date", "e.end_time as end_date",
					"e.id as employee_id", "'" + cu.ToString(editor, "employee") + "' as editor", "'events' as editor_view"},
				From:    "employee_events e",
				OrderBy: []string{"e.id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"place_simple": func(editor string) md.Query {
			return md.Query{
				Fields: []string{
					"p.*", "'" + cu.ToString(editor, "place") + "' as editor", "p.id as editor_id", "'place' as editor_view"},
				From:    "place_view p",
				OrderBy: []string{"p.id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"place": func(editor string) md.Query {
			return md.Query{
				Fields: []string{
					"p.*", "'" + cu.ToString(editor, "place") + "' as editor", "p.id as editor_id", "'place' as editor_view"},
				From:    "place_view p",
				OrderBy: []string{"p.id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"place_map": func(editor string) md.Query {
			return md.Query{
				Fields: []string{"*", "map_key as field_name", "map_value as value",
					"id as place_id", "'" + cu.ToString(editor, "place") + "' as editor", "'maps' as editor_view"},
				From:    "place_map",
				OrderBy: []string{"id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"place_contacts": func(editor string) md.Query {
			return md.Query{
				Fields:  []string{"p.*", "p.id as place_id", "'" + cu.ToString(editor, "place") + "' as editor", "'contacts' as editor_view"},
				From:    "place_contacts p",
				OrderBy: []string{"p.id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"place_events": func(editor string) md.Query {
			return md.Query{
				Fields: []string{"p.*", "p.start_time as start_date", "p.end_time as end_date",
					"p.id as place_id", "'" + cu.ToString(editor, "place") + "' as editor", "'events' as editor_view"},
				From:    "place_events p",
				OrderBy: []string{"p.id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"office_template_editor": func(editor string) md.Query {
			return md.Query{
				Fields:  []string{"id", "code", "report_key", "report_type", "trans_type", "direction", "report_name", "description", "label"},
				From:    "config_report",
				Filter:  fmt.Sprintf("file_type = '%s'", md.FileTypePDF.String()),
				OrderBy: []string{"id"},
				Limit:   st.BrowserRowLimit,
			}
		},
		"office_rate": func(editor string) md.Query {
			return md.Query{
				Fields: []string{
					"r.*", "p.place_name as place_name",
					"'" + cu.ToString(editor, "rate") + "' as editor", "r.id as editor_id", "'rate' as editor_view"},
				From:    "rate_view r left join place p on r.place_code = p.code",
				OrderBy: []string{"r.id"},
				Limit:   st.BrowserRowLimit,
			}
		},
	}
	query = md.Query{}
	editor := cu.ToString(params["editor"], "")
	if q, found := qMap[key]; found {
		query = q(editor)
	}
	return query
}

func (s *SearchConfig) Filter(view string, filter ct.BrowserFilter, queryFilters []string) (filters []string) {
	result := map[string]func() []string{
		"customer_simple": func() []string {
			return s.filterCustomer(view, filter, queryFilters)
		},
		"customer": func() []string {
			return s.filterCustomer(view, filter, queryFilters)
		},
		"customer_map": func() []string {
			return s.filterCustomer(view, filter, queryFilters)
		},
		"customer_addresses": func() []string {
			return s.filterCustomer(view, filter, queryFilters)
		},
		"customer_contacts": func() []string {
			return s.filterCustomer(view, filter, queryFilters)
		},
		"customer_events": func() []string {
			return s.filterCustomer(view, filter, queryFilters)
		},
		"product_simple": func() []string {
			return s.filterProduct(view, filter, queryFilters)
		},
		"product": func() []string {
			return s.filterProduct(view, filter, queryFilters)
		},
		"product_map": func() []string {
			return s.filterProduct(view, filter, queryFilters)
		},
		"product_events": func() []string {
			return s.filterProduct(view, filter, queryFilters)
		},
		"product_prices": func() []string {
			return s.filterProduct(view, filter, queryFilters)
		},
		"product_components": func() []string {
			return s.filterProduct(view, filter, queryFilters)
		},
		"tool_simple": func() []string {
			return s.filterTool(view, filter, queryFilters)
		},
		"tool": func() []string {
			return s.filterTool(view, filter, queryFilters)
		},
		"tool_map": func() []string {
			return s.filterTool(view, filter, queryFilters)
		},
		"tool_events": func() []string {
			return s.filterTool(view, filter, queryFilters)
		},
		"project_simple": func() []string {
			return s.filterProject(view, filter, queryFilters)
		},
		"project": func() []string {
			return s.filterProject(view, filter, queryFilters)
		},
		"project_map": func() []string {
			return s.filterProject(view, filter, queryFilters)
		},
		"project_addresses": func() []string {
			return s.filterProject(view, filter, queryFilters)
		},
		"project_contacts": func() []string {
			return s.filterProject(view, filter, queryFilters)
		},
		"project_events": func() []string {
			return s.filterProject(view, filter, queryFilters)
		},
		"employee_simple": func() []string {
			return s.filterEmployee(view, filter, queryFilters)
		},
		"employee": func() []string {
			return s.filterEmployee(view, filter, queryFilters)
		},
		"employee_map": func() []string {
			return s.filterEmployee(view, filter, queryFilters)
		},
		"employee_events": func() []string {
			return s.filterEmployee(view, filter, queryFilters)
		},
		"transitem_simple": func() []string {
			return s.filterTransItem(view, filter, queryFilters)
		},
		"invoice_simple": func() []string {
			return s.filterTransItem(view, filter, queryFilters)
		},
		"transitem": func() []string {
			return s.filterTransItem(view, filter, queryFilters)
		},
		"transitem_map": func() []string {
			return s.filterTransItem(view, filter, queryFilters)
		},
		"transitem_item": func() []string {
			return s.filterTransItem(view, filter, queryFilters)
		},
		"transpayment_simple": func() []string {
			return s.filterTransPayment(view, filter, queryFilters)
		},
		"transpayment": func() []string {
			return s.filterTransPayment(view, filter, queryFilters)
		},
		"transpayment_map": func() []string {
			return s.filterTransPayment(view, filter, queryFilters)
		},
		"transpayment_invoice": func() []string {
			return s.filterTransPayment(view, filter, queryFilters)
		},
		"transmovement_simple": func() []string {
			return s.filterTransMovement(view, filter, queryFilters)
		},
		"transmovement_stock": func() []string {
			return s.filterTransMovement(view, filter, queryFilters)
		},
		"transmovement": func() []string {
			return s.filterTransMovement(view, filter, queryFilters)
		},
		"transmovement_waybill": func() []string {
			return s.filterTransMovement(view, filter, queryFilters)
		},
		"transmovement_formula": func() []string {
			return s.filterTransMovement(view, filter, queryFilters)
		},
		"transmovement_map": func() []string {
			return s.filterTransMovement(view, filter, queryFilters)
		},
		"place_simple": func() []string {
			return s.filterPlace(view, filter, queryFilters)
		},
		"place": func() []string {
			return s.filterPlace(view, filter, queryFilters)
		},
		"place_map": func() []string {
			return s.filterPlace(view, filter, queryFilters)
		},
		"place_contacts": func() []string {
			return s.filterPlace(view, filter, queryFilters)
		},
		"place_events": func() []string {
			return s.filterPlace(view, filter, queryFilters)
		},
		"office_template_editor": func() []string {
			return s.filterOffice(view, filter, queryFilters)
		},
		"office_rate": func() []string {
			return s.filterOffice(view, filter, queryFilters)
		},
	}

	filters = []string{}
	if f, found := result[view]; found {
		filters = f()
	}
	return filters
}

func (s *SearchConfig) filterCustomer(view string, filter ct.BrowserFilter, queryFilters []string) []string {
	result := map[string]func() []string{
		"customer_simple": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (CAST(%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"customer": func() []string {
			switch filter.Field {
			case "tax_free", "inactive":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %t)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToBoolean(filter.Value, false)))
			case "code", "customer_name", "customer_type", "tax_number", "account", "notes", "tag_lst":
				return append(queryFilters,
					fmt.Sprintf("%s (CAST(%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			case "id", "terms", "credit_limit", "discount":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %s)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			default:
				return queryFilters
			}

		},
		"customer_map": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"customer_addresses": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"customer_contacts": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"customer_events": func() []string {
			switch filter.Field {
			case "start_time", "end_time":
				return append(queryFilters,
					fmt.Sprintf("%s (start_time %s '%s')", pre(filter.Or), compMap[filter.Comp], cu.ToString(filter.Value, "")))
			default:
				return append(queryFilters,
					fmt.Sprintf("%s (CAST(%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			}
		},
	}
	return result[view]()
}

func (s *SearchConfig) filterProduct(view string, filter ct.BrowserFilter, queryFilters []string) []string {
	result := map[string]func() []string{
		"product_simple": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (CAST(%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"product": func() []string {
			switch filter.Field {
			case "inactive":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %t)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToBoolean(filter.Value, false)))
			case "code", "product_name", "product_type", "tax_code", "unit", "notes", "tag_lst", "barcode_type", "barcode":
				return append(queryFilters,
					fmt.Sprintf("%s (CAST(%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			case "id", "barcode_qty":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %s)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			default:
				return queryFilters
			}
		},
		"product_map": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"product_events": func() []string {
			switch filter.Field {
			case "start_time", "end_time":
				return append(queryFilters,
					fmt.Sprintf("%s (start_time %s '%s')", pre(filter.Or), compMap[filter.Comp], cu.ToString(filter.Value, "")))
			default:
				return append(queryFilters,
					fmt.Sprintf("%s (CAST(%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			}
		},
		"product_prices": func() []string {
			switch filter.Field {
			case "valid_from", "valid_to":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			case "id", "qty", "price_value":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %s)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			default:
				return append(queryFilters,
					fmt.Sprintf("%s (CAST(%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			}
		},
		"product_components": func() []string {
			switch filter.Field {
			case "id", "qty":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %s)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			default:
				return append(queryFilters,
					fmt.Sprintf("%s (CAST(%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			}
		},
	}
	return result[view]()
}

func (s *SearchConfig) filterTool(view string, filter ct.BrowserFilter, queryFilters []string) []string {
	result := map[string]func() []string{
		"tool_simple": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (CAST(%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"tool": func() []string {
			switch filter.Field {
			case "inactive":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %t)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToBoolean(filter.Value, false)))
			case "code", "description", "product_code", "notes", "tag_lst", "serial_number":
				return append(queryFilters,
					fmt.Sprintf("%s (CAST(%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			case "id":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %s)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			default:
				return queryFilters
			}
		},
		"tool_map": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"tool_events": func() []string {
			switch filter.Field {
			case "start_time", "end_time":
				return append(queryFilters,
					fmt.Sprintf("%s (start_time %s '%s')", pre(filter.Or), compMap[filter.Comp], cu.ToString(filter.Value, "")))
			default:
				return append(queryFilters,
					fmt.Sprintf("%s (CAST(%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			}
		},
	}
	return result[view]()
}

func (s *SearchConfig) filterProject(view string, filter ct.BrowserFilter, queryFilters []string) []string {
	result := map[string]func() []string{
		"project_simple": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (CAST(%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"project": func() []string {
			switch filter.Field {
			case "inactive":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %t)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToBoolean(filter.Value, false)))
			case "start_date", "end_date":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			case "code", "project_name", "customer_code", "notes", "tag_lst":
				return append(queryFilters,
					fmt.Sprintf("%s (CAST(%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			case "id":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %s)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			default:
				return queryFilters
			}
		},
		"project_map": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"project_addresses": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"project_contacts": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"project_events": func() []string {
			switch filter.Field {
			case "start_time", "end_time":
				return append(queryFilters,
					fmt.Sprintf("%s (start_time %s '%s')", pre(filter.Or), compMap[filter.Comp], cu.ToString(filter.Value, "")))
			default:
				return append(queryFilters,
					fmt.Sprintf("%s (CAST(%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			}
		},
	}
	return result[view]()
}

func (s *SearchConfig) filterEmployee(view string, filter ct.BrowserFilter, queryFilters []string) []string {
	result := map[string]func() []string{
		"employee_simple": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (CAST(%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"employee": func() []string {
			switch filter.Field {
			case "inactive":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %t)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToBoolean(filter.Value, false)))
			case "start_date", "end_date":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			case "code", "notes", "tag_lst",
				"first_name", "surname", "status", "email", "phone", "mobile",
				"street", "city", "state", "zip_code", "country":
				return append(queryFilters,
					fmt.Sprintf("%s (CAST(%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			case "id":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %s)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			default:
				return queryFilters
			}
		},
		"employee_map": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"employee_events": func() []string {
			switch filter.Field {
			case "start_time", "end_time":
				return append(queryFilters,
					fmt.Sprintf("%s (start_time %s '%s')", pre(filter.Or), compMap[filter.Comp], cu.ToString(filter.Value, "")))
			default:
				return append(queryFilters,
					fmt.Sprintf("%s (CAST(%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			}
		},
	}
	return result[view]()
}

func (s *SearchConfig) filterPlace(view string, filter ct.BrowserFilter, queryFilters []string) []string {
	result := map[string]func() []string{
		"place_simple": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (CAST(%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"place": func() []string {
			switch filter.Field {
			case "inactive":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %t)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToBoolean(filter.Value, false)))
			case "code", "place_name", "place_type", "currency_code", "notes", "tag_lst",
				"country", "state", "zip_code", "city", "street":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			case "id":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %s)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			default:
				return queryFilters
			}
		},
		"place_map": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"place_contacts": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"place_events": func() []string {
			switch filter.Field {
			case "start_time", "end_time":
				return append(queryFilters,
					fmt.Sprintf("%s (start_time %s '%s')", pre(filter.Or), compMap[filter.Comp], cu.ToString(filter.Value, "")))
			default:
				return append(queryFilters,
					fmt.Sprintf("%s (CAST(%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			}
		},
	}
	return result[view]()
}

func (s *SearchConfig) filterTransItem(view string, filter ct.BrowserFilter, queryFilters []string) []string {
	result := map[string]func() []string{
		"transitem_simple": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (CAST(%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"invoice_simple": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (CAST(%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"transitem": func() []string {
			switch filter.Field {
			case "tax_free", "paid", "closed":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %t)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToBoolean(filter.Value, false)))
			case "trans_date", "due_time":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			case "id", "rate", "worksheet_distance", "worksheet_repair", "worksheet_total", "rent_holiday", "rent_bad_tool", "rent_other", "amount":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %s)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			case "customer_name":
				return append(queryFilters,
					fmt.Sprintf("%s (CAST(%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			default:
				return append(queryFilters,
					fmt.Sprintf("%s (CAST(t.%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			}
		},
		"transitem_map": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"transitem_item": func() []string {
			switch filter.Field {
			case "deposit", "action_price":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %t)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToBoolean(filter.Value, false)))
			case "id", "qty", "fx_price", "net_amount", "discount", "vat_amount", "amount", "own_stock":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %s)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			default:
				return append(queryFilters,
					fmt.Sprintf("%s (CAST(%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			}
		},
	}
	return result[view]()
}

func (s *SearchConfig) filterTransPayment(view string, filter ct.BrowserFilter, queryFilters []string) []string {
	result := map[string]func() []string{
		"transpayment_simple": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (CAST(%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"transpayment": func() []string {
			switch filter.Field {
			case "closed":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %t)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToBoolean(filter.Value, false)))
			case "trans_date", "paid_date":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			case "id", "amount":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %s)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			case "code":
				return append(queryFilters,
					fmt.Sprintf("%s (CAST(t.%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			default:
				return append(queryFilters,
					fmt.Sprintf("%s (CAST(%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			}
		},
		"transpayment_map": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"transpayment_invoice": func() []string {
			switch filter.Field {
			case "id", "paid_amount", "paid_rate", "invoice_amount":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %s)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			case "trans_date", "payment_date":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			default:
				return append(queryFilters,
					fmt.Sprintf("%s (CAST(%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			}
		},
	}
	return result[view]()
}

func (s *SearchConfig) filterTransMovement(view string, filter ct.BrowserFilter, queryFilters []string) []string {
	result := map[string]func() []string{
		"transmovement_simple": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (CAST(%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"transmovement_stock": func() []string {
			switch filter.Field {
			case "posdate":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			case "id", "qty":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %s)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			default:
				return append(queryFilters,
					fmt.Sprintf("%s (CAST(%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			}
		},
		"transmovement": func() []string {
			switch filter.Field {
			case "id", "qty":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %s)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			case "shipping_date":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			default:
				return append(queryFilters,
					fmt.Sprintf("%s (CAST(%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			}
		},
		"transmovement_waybill": func() []string {
			switch filter.Field {
			case "id":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %s)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			case "shipping_time":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			default:
				return append(queryFilters,
					fmt.Sprintf("%s (CAST(%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			}
		},
		"transmovement_formula": func() []string {
			switch filter.Field {
			case "id", "qty":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %s)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			case "shared":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %t)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToBoolean(filter.Value, false)))
			default:
				return append(queryFilters,
					fmt.Sprintf("%s (CAST(%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			}
		},
		"transmovement_map": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
	}
	return result[view]()
}

func (s *SearchConfig) filterOffice(view string, filter ct.BrowserFilter, queryFilters []string) []string {
	result := map[string]func() []string{
		"office_template_editor": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (CAST(%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"office_rate": func() []string {
			switch filter.Field {
			case "id", "qty", "rate_value":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %s)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			case "rate_date":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			default:
				return append(queryFilters,
					fmt.Sprintf("%s (CAST(%s as CHAR(255)) %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			}
		},
	}
	return result[view]()
}
