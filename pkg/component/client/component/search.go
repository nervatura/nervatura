package component

import (
	"slices"
	"strings"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
)

type SearchView struct {
	Title             string
	Icon              string
	Simple            bool
	ReadOnly          bool
	LabelAdd          string
	Fields            []ct.TableField
	VisibleColumns    cu.IM
	HideFilters       cu.IM
	Filters           []any
	FilterPlaceholder string
}

func searchSideBar(labels cu.SM, data cu.IM) (items []ct.SideBarItem) {
	sideGroup := cu.ToString(data["side_group"], "")
	sideElement := func(name string) *ct.SideBarElement {
		return &ct.SideBarElement{
			Name:  name,
			Value: name,
			Label: " " + SearchViewConfig(name, labels).Title,
			Icon:  SearchViewConfig(name, labels).Icon,
			//Selected: (cu.ToString(data["view"], "") == name),
		}
	}
	sideGroupElement := func(name, label string, views []string) *ct.SideBarGroup {
		sg := &ct.SideBarGroup{
			Name:  name,
			Value: name,
			Label: label,
		}
		for _, name := range views {
			se := sideElement(name)
			//se.Selected = (cu.ToString(data["view"], "") == name)
			sg.Items = append(sg.Items, *se)
		}
		return sg
	}
	sb := []ct.SideBarItem{
		&ct.SideBarSeparator{},
	}

	transitemViews := []string{
		"transitem_simple", "transitem", "transitem_map", "transitem_item",
	}
	transitemGroup := sideGroupElement("group_transitem", labels["transitem_title"], transitemViews)
	transitemGroup.Selected = (sideGroup == "group_transitem") ||
		((sideGroup == "") && slices.Contains(transitemViews, cu.ToString(data["view"], "")))
	sb = append(sb, transitemGroup)

	customerViews := []string{
		"customer_simple", "customer", "customer_map", "customer_addresses", "customer_contacts", "customer_events",
	}
	customerGroup := sideGroupElement("group_customer", labels["customer_title"], customerViews)
	customerGroup.Selected = (sideGroup == "group_customer") ||
		((sideGroup == "") && slices.Contains(customerViews, cu.ToString(data["view"], "")))
	sb = append(sb, customerGroup)

	productViews := []string{
		"product_simple", "product", "product_map", "product_events", "product_prices",
	}
	productGroup := sideGroupElement("group_product", labels["product_title"], productViews)
	productGroup.Selected = (sideGroup == "group_product") ||
		((sideGroup == "") && slices.Contains(productViews, cu.ToString(data["view"], "")))
	sb = append(sb, productGroup)

	employeeViews := []string{
		"employee_simple", "employee", "employee_map", "employee_events",
	}
	employeeGroup := sideGroupElement("group_employee", labels["employee_title"], employeeViews)
	employeeGroup.Selected = (sideGroup == "group_employee") ||
		((sideGroup == "") && slices.Contains(employeeViews, cu.ToString(data["view"], "")))
	sb = append(sb, employeeGroup)

	toolViews := []string{
		"tool_simple", "tool", "tool_map", "tool_events",
	}
	toolGroup := sideGroupElement("group_tool", labels["tool_title"], toolViews)
	toolGroup.Selected = (sideGroup == "group_tool") ||
		((sideGroup == "") && slices.Contains(toolViews, cu.ToString(data["view"], "")))
	sb = append(sb, toolGroup)

	projectViews := []string{
		"project_simple", "project", "project_map", "project_addresses", "project_contacts", "project_events",
	}
	projectGroup := sideGroupElement("group_project", labels["project_title"], projectViews)
	projectGroup.Selected = (sideGroup == "group_project") ||
		((sideGroup == "") && slices.Contains(projectViews, cu.ToString(data["view"], "")))
	sb = append(sb, projectGroup)

	//for _, name := range []string{"customer_simple", "customer", "customer_addresses", "customer_contacts"} {
	//	sb = append(sb, sideElement(name))
	//}
	return sb
}

func SearchViewConfig(view string, labels cu.SM) (config SearchView) {
	viewMap := map[string]SearchView{
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
				{Name: "trans_code", Label: labels["trans_code"], FieldType: ct.TableFieldTypeLink},
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
				{Name: "worksheet_notes", Label: labels["trans_worksheet_notes"]},
				{Name: "rent_holiday", Label: labels["trans_rent_holiday"], FieldType: ct.TableFieldTypeNumber},
				{Name: "rent_bad_tool", Label: labels["trans_rent_bad_tool"], FieldType: ct.TableFieldTypeNumber},
				{Name: "rent_other", Label: labels["trans_rent_other"], FieldType: ct.TableFieldTypeNumber},
				{Name: "rent_notes", Label: labels["trans_rent_notes"]},
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
			Icon:        ct.IconFileText,
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
				{Name: "code", Label: labels["trans_code"], FieldType: ct.TableFieldTypeLink},
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
				"code": true, "trans_date": true, "description": true, "qty": true, "currency_code": true,
				"fx_price": true, "net_amount": true, "discount": true, "vat_amount": true,
				"amount": true, "deposit": true,
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
			Icon:        ct.IconUser,
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
			Icon:        ct.IconShoppingCart,
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
			Icon:        ct.IconWrench,
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
			Icon:        ct.IconClock,
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
			Icon:        ct.IconMale,
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
	}
	return viewMap[view]
}
