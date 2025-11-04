package component

import (
	"slices"
	"strings"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	st "github.com/nervatura/nervatura/v6/pkg/static"
)

type ShippingEditor struct{}

func (e *ShippingEditor) SideBar(labels cu.SM, data cu.IM) (items []ct.SideBarItem) {
	user := cu.ToIM(data["user"], cu.IM{})
	shipping := cu.ToIM(data["shipping"], cu.IM{})

	dirty := cu.ToBoolean(data["dirty"], false)
	createItems := e.createItems(data)
	readonly := (cu.ToString(user["user_group"], "") == md.UserGroupGuest.String()) ||
		(cu.ToString(shipping["place_code"], "") == "") || createItems == 0

	return []ct.SideBarItem{
		&ct.SideBarSeparator{},
		&ct.SideBarElement{
			Name:    "editor_cancel",
			Value:   "editor_cancel",
			Label:   labels["transitem_title"],
			Icon:    ct.IconReply,
			NotFull: true,
		},
		&ct.SideBarSeparator{},
		&ct.SideBarElement{
			Name:     "create_all",
			Value:    "create_all",
			Label:    labels["shipping_all_label"],
			Icon:     ct.IconPlus,
			Selected: !dirty,
			Disabled: readonly || dirty,
		},
		&ct.SideBarElement{
			Name:     "create_item",
			Value:    "create_item",
			Label:    labels["shipping_create_label"],
			Icon:     ct.IconCheck,
			Selected: dirty,
			Disabled: readonly || !dirty,
		},
		&ct.SideBarSeparator{},
		&ct.SideBarElementLink{
			SideBarElement: ct.SideBarElement{
				Name:  "editor_help",
				Value: "editor_help",
				Label: labels["editor_help"],
				Icon:  ct.IconQuestionCircle,
			},
			Href:       st.DocsClientPath, //+ "/shipping",
			LinkTarget: "_blank",
		},
	}
}

func (e *ShippingEditor) createItems(data cu.IM) (rc int64) {
	dirty := cu.ToBoolean(data["dirty"], false)
	items := cu.ToIMA(data["items"], []cu.IM{})
	for _, item := range items {
		flag := func() int64 {
			if dirty && cu.ToFloat(item["qty"], 0) != 0 {
				return 1
			}
			if !dirty && cu.ToFloat(item["difference"], 0) != 0 {
				return 1
			}
			return 0
		}
		rc += flag()
	}
	return rc
}

func (e *ShippingEditor) View(labels cu.SM, data cu.IM) (views []ct.EditorView) {
	movements := cu.ToIMA(data["movements"], []cu.IM{})

	return []ct.EditorView{
		{
			Key:   "shipping",
			Label: labels["shipping_view"],
			Icon:  ct.IconTruck,
		},
		{
			Key:   "items",
			Label: labels["shipping_create"],
			Icon:  ct.IconPlus,
			Badge: cu.ToString(e.createItems(data), "0"),
		},
		{
			Key:   "movements",
			Label: labels["shipping_delivery"],
			Icon:  ct.IconTruck,
			Badge: cu.ToString(int64(len(movements)), "0"),
		},
	}
}

func (e *ShippingEditor) Row(view string, labels cu.SM, data cu.IM) (rows []ct.Row) {
	if !slices.Contains([]string{"shipping"}, view) {
		return []ct.Row{}
	}

	shipping := cu.ToIM(data["shipping"], cu.IM{})
	transType := strings.TrimPrefix(cu.ToString(shipping["trans_type"], ""), "TRANS_")
	places := cu.ToIMA(data["places"], []cu.IM{})

	placeOpt := func() (opt []ct.SelectOption) {
		opt = []ct.SelectOption{}
		for _, place := range places {
			opt = append(opt, ct.SelectOption{
				Value: cu.ToString(place["code"], ""), Text: cu.ToString(place["place_name"], ""),
			})
		}
		return opt
	}

	return []ct.Row{
		{Columns: []ct.RowColumn{
			{Label: labels["trans_type"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "trans_type_" + cu.ToString(shipping["id"], ""),
				},
				Type: ct.FieldTypeString, Value: cu.IM{
					"name":     "trans_type",
					"value":    transType,
					"disabled": true,
				},
			}},
			{Label: labels["customer_name"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "customer_name_" + cu.ToString(shipping["id"], ""),
				},
				Type: ct.FieldTypeString, Value: cu.IM{
					"name":     "customer_name",
					"value":    shipping["customer_name"],
					"disabled": true,
				},
			}},
		}},
		{Columns: []ct.RowColumn{
			{Label: labels["trans_direction"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "direction_" + cu.ToString(shipping["id"], ""),
				},
				Type: ct.FieldTypeString, Value: cu.IM{
					"name":     "direction",
					"value":    shipping["direction"],
					"disabled": true,
				},
			}},
			{Label: labels["movement_shipping_time"],
				Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "shipping_time_" + cu.ToString(shipping["id"], ""),
					},
					Type: ct.FieldTypeDateTime, Value: cu.IM{
						"name":    "shipping_time",
						"is_null": false,
						"value":   cu.ToString(shipping["shipping_time"], ""),
					},
				}},
			{
				Label: labels["place_name_movement"],
				Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "place_code_" + cu.ToString(shipping["id"], ""),
					},
					Type: ct.FieldTypeSelect, Value: cu.IM{
						"name":    "place_code",
						"options": placeOpt(),
						"is_null": true,
						"value":   cu.ToString(shipping["place_code"], ""),
					},
				}},
		}, Full: true, BorderTop: true},
	}
}

func (e *ShippingEditor) Table(view string, labels cu.SM, data cu.IM) []ct.Table {
	if !slices.Contains([]string{"movements", "items"}, view) {
		return []ct.Table{}
	}

	tblMap := map[string]func() []ct.Table{
		"items": func() []ct.Table {
			items := cu.ToIMA(data["items"], []cu.IM{})
			return []ct.Table{
				{
					Fields: []ct.TableField{
						{Name: "product_code", Label: labels["product_code"], FieldType: ct.TableFieldTypeLink, ReadOnly: true},
						{Name: "product_name", Label: labels["product_name"], ReadOnly: true},
						{Name: "item_qty", Label: labels["shipping_qty"], FieldType: ct.TableFieldTypeNumber, ReadOnly: true},
						{Name: "turnover", Label: labels["shipping_turnover"], FieldType: ct.TableFieldTypeNumber, ReadOnly: true},
						{Name: "difference", Label: labels["shipping_diff"], FieldType: ct.TableFieldTypeNumber, Format: true, ReadOnly: true},
						{Name: "batch_no", Label: labels["movement_batchnumber"]},
						{Name: "qty", Label: labels["movement_qty"], FieldType: ct.TableFieldTypeNumber, TriggerEvent: true},
					},
					Rows:               items,
					Pagination:         ct.PaginationTypeTop,
					PageSize:           10,
					HidePaginatonSize:  true,
					RowSelected:        true,
					TableFilter:        true,
					FilterPlaceholder:  labels["placeholder_filter"],
					Editable:           true,
					EditDeleteDisabled: true,
					Unsortable:         true,
				},
			}
		},
		"movements": func() []ct.Table {
			items := cu.ToIMA(data["movements"], []cu.IM{})
			return []ct.Table{
				{
					Fields: []ct.TableField{
						{Name: "trans_code", Label: labels["trans_code"], FieldType: ct.TableFieldTypeLink},
						{Name: "shipping_date", Label: labels["shipping_date"], FieldType: ct.TableFieldTypeDate},
						{Name: "product_code", Label: labels["product_code"], FieldType: ct.TableFieldTypeLink},
						{Name: "product_name", Label: labels["product_name"]},
						{Name: "unit", Label: labels["product_unit"]},
						{Name: "shipping_qty", Label: labels["movement_qty"], FieldType: ct.TableFieldTypeNumber},
					},
					Rows:              items,
					Pagination:        ct.PaginationTypeTop,
					PageSize:          10,
					HidePaginatonSize: true,
					RowSelected:       false,
					TableFilter:       true,
					FilterPlaceholder: labels["placeholder_filter"],
				},
			}
		},
	}
	return tblMap[view]()
}

func (e *ShippingEditor) Form(formKey string, labels cu.SM, data cu.IM) (form ct.Form) {
	return ct.Form{}
}
