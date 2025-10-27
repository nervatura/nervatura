package component

import (
	"slices"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
	st "github.com/nervatura/nervatura/v6/pkg/static"
)

type ProductEditor struct{}

func (e *ProductEditor) SideBar(labels cu.SM, data cu.IM) (items []ct.SideBarItem) {
	var product cu.IM = cu.ToIM(data["product"], cu.IM{"product_meta": cu.IM{}})
	user := cu.ToIM(data["user"], cu.IM{})

	dirty := cu.ToBoolean(data["dirty"], false)
	readonly := (cu.ToString(user["user_group"], "") == md.UserGroupGuest.String())
	newInput := (cu.ToInteger(product["id"], 0) == 0)
	updateLabel := labels["editor_save"]
	if newInput {
		updateLabel = labels["editor_create"]
	}
	updateDisabled := func() (disabled bool) {
		return (cu.ToString(product["product_name"], "") == "") || readonly
	}

	smState := func() *ct.SideBarStatic {
		if cu.ToBoolean(product["inactive"], false) {
			return &ct.SideBarStatic{
				Icon: ct.IconLock, Label: labels["state_closed"], Color: "brown",
			}
		}
		if newInput {
			return &ct.SideBarStatic{
				Icon: ct.IconPlus, Label: labels["state_new"], Color: "yellow",
			}
		}
		return &ct.SideBarStatic{
			Icon: ct.IconEdit, Label: labels["state_edit"], Color: "green",
		}
	}

	return []ct.SideBarItem{
		&ct.SideBarSeparator{},
		&ct.SideBarElement{
			Name:    "editor_cancel",
			Value:   "editor_cancel",
			Label:   labels["browser_title"],
			Icon:    ct.IconReply,
			NotFull: true,
		},
		&ct.SideBarSeparator{},
		smState(),
		&ct.SideBarSeparator{},
		&ct.SideBarElement{
			Name:     "editor_save",
			Value:    "editor_save",
			Label:    updateLabel,
			Icon:     ct.IconUpload,
			Selected: dirty,
			Disabled: updateDisabled(),
		},
		&ct.SideBarElement{
			Name:     "editor_delete",
			Value:    "editor_delete",
			Label:    labels["editor_delete"],
			Icon:     ct.IconTimes,
			Disabled: newInput || readonly,
		},
		&ct.SideBarSeparator{},
		&ct.SideBarElement{
			Name:     "editor_new",
			Value:    "editor_new",
			Label:    labels["product_new"],
			Icon:     ct.IconShoppingCart,
			Disabled: newInput || dirty || readonly,
		},
		&ct.SideBarSeparator{},
		&ct.SideBarElement{
			Name:     "editor_report",
			Value:    "editor_report",
			Label:    labels["editor_report"],
			Icon:     ct.IconChartBar,
			Disabled: newInput || dirty,
		},
		&ct.SideBarSeparator{},
		&ct.SideBarElement{
			Name:     "editor_bookmark",
			Value:    "editor_bookmark",
			Label:    labels["editor_bookmark"],
			Icon:     ct.IconStar,
			Disabled: newInput,
		},
		&ct.SideBarElementLink{
			SideBarElement: ct.SideBarElement{
				Name:  "editor_help",
				Value: "editor_help",
				Label: labels["editor_help"],
				Icon:  ct.IconQuestionCircle,
			},
			Href:       st.DocsClientPath, //+ "/product",
			LinkTarget: "_blank",
		},
	}
}

func (e *ProductEditor) View(labels cu.SM, data cu.IM) (views []ct.EditorView) {
	var product cu.IM = cu.ToIM(data["product"], cu.IM{})
	productMap := cu.ToIM(product["product_map"], cu.IM{})
	events := cu.ToIMA(product["events"], []cu.IM{})
	prices := cu.ToIMA(data["prices"], []cu.IM{})
	newInput := (cu.ToInteger(product["id"], 0) == 0)

	if newInput {
		return []ct.EditorView{
			{
				Key:   "product",
				Label: labels["product_view"],
				Icon:  ct.IconShoppingCart,
			},
		}
	}
	return []ct.EditorView{
		{
			Key:   "product",
			Label: labels["product_view"],
			Icon:  ct.IconShoppingCart,
		},
		{
			Key:   "maps",
			Label: labels["map_view"],
			Icon:  ct.IconDatabase,
			Badge: cu.ToString(int64(len(productMap)), "0"),
		},
		{
			Key:   "events",
			Label: labels["event_view"],
			Icon:  ct.IconCalendar,
			Badge: cu.ToString(int64(len(events)), "0"),
		},
		{
			Key:   "prices",
			Label: " " + labels["price_view"],
			Icon:  ct.IconMoney,
			Badge: cu.ToString(int64(len(prices)), "0"),
		},
	}
}

func (e *ProductEditor) Row(view string, labels cu.SM, data cu.IM) (rows []ct.Row) {
	if !slices.Contains([]string{"product", "maps"}, view) {
		return []ct.Row{}
	}

	var product md.Product = md.Product{}
	ut.ConvertToType(data["product"], &product)

	configMap := cu.ToIMA(data["config_map"], []cu.IM{})
	taxCodes := cu.ToIMA(data["tax_codes"], []cu.IM{})
	selectedField := cu.ToString(data["map_field"], "")

	mapFieldOpt := func() (opt []ct.SelectOption) {
		opt = []ct.SelectOption{}
		for _, field := range configMap {
			filter := ut.ToStringArray(field["filter"])
			if slices.Contains(filter, "FILTER_PRODUCT") || len(filter) == 0 {
				if _, ok := product.ProductMap[cu.ToString(field["field_name"], "")]; !ok {
					opt = append(opt, ct.SelectOption{
						Value: cu.ToString(field["field_name"], ""), Text: cu.ToString(field["description"], ""),
					})
				}
			}
		}
		return opt
	}

	productTypeOpt := func() (opt []ct.SelectOption) {
		opt = []ct.SelectOption{}
		for _, ptype := range []md.ProductType{
			md.ProductTypeItem, md.ProductTypeService,
		} {
			opt = append(opt, ct.SelectOption{
				Value: ptype.String(), Text: ptype.String(),
			})
		}
		return opt
	}

	taxCodeOpt := func() (opt []ct.SelectOption) {
		opt = []ct.SelectOption{}
		for _, tcode := range taxCodes {
			opt = append(opt, ct.SelectOption{
				Value: cu.ToString(tcode["code"], ""), Text: cu.ToString(tcode["description"], ""),
			})
		}
		return opt
	}

	barcodeTypeOpt := func() (opt []ct.SelectOption) {
		opt = []ct.SelectOption{}
		for _, btype := range []md.BarcodeType{
			md.BarcodeTypeCode128, md.BarcodeTypeCode39, md.BarcodeTypeEan8, md.BarcodeTypeEan13, md.BarcodeTypeQRCode,
		} {
			opt = append(opt, ct.SelectOption{
				Value: btype.String(), Text: btype.String(),
			})
		}
		return opt
	}

	if view == "maps" {
		return []ct.Row{
			{Columns: []ct.RowColumn{
				{Label: labels["map_fields"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "map_field_" + cu.ToString(product.Id, ""),
					},
					Type: ct.FieldTypeSelect, Value: cu.IM{
						"name":    "map_field",
						"options": mapFieldOpt(),
						"is_null": true,
						"value":   selectedField,
					},
				}},
			}, Full: false, FieldCol: true},
		}
	}

	return []ct.Row{
		{Columns: []ct.RowColumn{
			{Label: labels["product_name"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "product_name_" + cu.ToString(product.Id, ""),
				},
				Type: ct.FieldTypeString, Value: cu.IM{
					"name":        "product_name",
					"invalid":     (product.ProductName == ""),
					"placeholder": labels["mandatory_data"],
					"value":       product.ProductName,
				},
			}},
		}, Full: true, BorderBottom: true, FieldCol: true},
		{Columns: []ct.RowColumn{
			{Label: labels["product_code"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "code_" + cu.ToString(product.Id, ""),
				},
				Type: ct.FieldTypeString, Value: cu.IM{
					"name":     "code",
					"value":    product.Code,
					"disabled": true,
				},
			}},
			{Label: labels["product_type"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "product_type_" + cu.ToString(product.Id, ""),
				},
				Type: ct.FieldTypeSelect, Value: cu.IM{
					"name":    "product_type",
					"options": productTypeOpt(),
					"is_null": false,
					"value":   product.ProductType.String(),
				},
			}},
			{Label: labels["tax_code"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "tax_code_" + cu.ToString(product.Id, ""),
				},
				Type: ct.FieldTypeSelect, Value: cu.IM{
					"name":    "tax_code",
					"options": taxCodeOpt(),
					"is_null": false,
					"value":   product.TaxCode,
				},
			}},
			{Label: labels["product_unit"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "unit_" + cu.ToString(product.Id, ""),
				},
				Type: ct.FieldTypeString, Value: cu.IM{
					"name":  "unit",
					"value": product.ProductMeta.Unit,
				},
			}},
		}, Full: true, BorderBottom: true},
		{Columns: []ct.RowColumn{
			{Label: labels["product_barcode_type"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "barcode_type_" + cu.ToString(product.Id, ""),
				},
				Type: ct.FieldTypeSelect, Value: cu.IM{
					"name":    "barcode_type",
					"options": barcodeTypeOpt(),
					"is_null": false,
					"value":   product.ProductMeta.BarcodeType.String(),
				},
			}},
			{Label: labels["product_barcode"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "barcode_" + cu.ToString(product.Id, ""),
				},
				Type: ct.FieldTypeString, Value: cu.IM{
					"name":  "barcode",
					"value": product.ProductMeta.Barcode,
				},
			}},
			{Label: labels["product_barcode_qty"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "barcode_qty_" + cu.ToString(product.Id, ""),
				},
				Type: ct.FieldTypeNumber, Value: cu.IM{
					"name":  "barcode_qty",
					"value": product.ProductMeta.BarcodeQty,
				},
			}},
			{Label: labels["product_inactive"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "inactive_" + cu.ToString(product.Id, ""),
				},
				Type: ct.FieldTypeBool, Value: cu.IM{
					"name":  "inactive",
					"value": cu.ToBoolean(product.ProductMeta.Inactive, false),
				},
			}},
		}, Full: true, BorderBottom: true},
		{Columns: []ct.RowColumn{
			{Label: labels["product_notes"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "notes_" + cu.ToString(product.Id, ""),
				},
				Type: ct.FieldTypeText, Value: cu.IM{
					"name":  "notes",
					"value": product.ProductMeta.Notes,
					"rows":  4,
				},
			}},
			{
				Label: labels["product_tags"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "tags_" + cu.ToString(product.Id, ""),
					},
					Type: ct.FieldTypeList, Value: cu.IM{
						"name":                "tags",
						"rows":                ut.ToTagList(product.ProductMeta.Tags),
						"label_value":         "tag",
						"pagination":          ct.PaginationTypeBottom,
						"page_size":           5,
						"hide_paginaton_size": true,
						"list_filter":         true,
						"filter_placeholder":  labels["placeholder_filter"],
						"add_item":            true,
						"add_icon":            ct.IconTag,
						"edit_item":           false,
						"delete_item":         true,
						"indicator":           ct.IndicatorSpinner,
					},
				},
			},
		}, Full: true, BorderBottom: true},
	}
}

func (e *ProductEditor) Table(view string, labels cu.SM, data cu.IM) []ct.Table {
	if !slices.Contains([]string{"maps", "events", "prices"}, view) {
		return []ct.Table{}
	}

	var product cu.IM = cu.ToIM(data["product"], cu.IM{})
	newInput := (cu.ToInteger(product["id"], 0) == 0)
	tblMap := map[string]func() []ct.Table{
		"maps": func() []ct.Table {
			configMap := cu.ToIMA(data["config_map"], []cu.IM{})
			productMap := cu.ToIM(product["product_map"], cu.IM{})
			return []ct.Table{
				{
					Fields: []ct.TableField{
						{Name: "description", Label: labels["map_description"], ReadOnly: true},
						{Name: "value", Label: labels["map_value"], FieldType: ct.TableFieldTypeMeta, Required: true},
					},
					Rows:              mapTableRows(productMap, configMap),
					Pagination:        ct.PaginationTypeTop,
					PageSize:          5,
					HidePaginatonSize: true,
					RowSelected:       true,
					TableFilter:       true,
					FilterPlaceholder: labels["placeholder_filter"],
					AddItem:           !newInput && (cu.ToString(data["map_field"], "") != ""),
					LabelAdd:          labels["map_new"],
					Editable:          true,
					Unsortable:        true,
				},
			}
		},
		"events": func() []ct.Table {
			event := cu.ToIMA(product["events"], []cu.IM{})
			return []ct.Table{
				{
					Fields: []ct.TableField{
						{Name: "subject", Label: labels["event_subject"]},
						{Name: "start_time", Label: labels["event_start_time"], FieldType: ct.TableFieldTypeDateTime},
						{Name: "end_time", Label: labels["event_end_time"], FieldType: ct.TableFieldTypeDateTime},
						{Name: "place", Label: labels["event_place"]},
						{Name: "description", Label: labels["event_description"]},
						//{Name: "tag_lst", Label: labels["event_tags"]},
					},
					Rows:              event,
					Pagination:        ct.PaginationTypeTop,
					PageSize:          5,
					HidePaginatonSize: true,
					RowSelected:       true,
					TableFilter:       true,
					FilterPlaceholder: labels["placeholder_filter"],
					AddItem:           !newInput,
					LabelAdd:          labels["event_new"],
				},
			}
		},
		"prices": func() []ct.Table {
			priceRows := func() []cu.IM {
				rows := []cu.IM{}
				prices := cu.ToIMA(data["prices"], []cu.IM{})
				for _, price := range prices {
					rows = append(rows, cu.IM{
						"price_type":    price["price_type"],
						"valid_from":    price["valid_from"],
						"valid_to":      price["valid_to"],
						"customer_code": price["customer_code"],
						"currency_code": price["currency_code"],
						"qty":           price["qty"],
						"price_value":   cu.ToIM(price["price_meta"], cu.IM{})["price_value"],
						"price_meta":    price["price_meta"],
					})
				}
				return rows
			}
			return []ct.Table{
				{
					Fields: []ct.TableField{
						{Name: "price_type", Label: labels["price_type"]},
						{Name: "valid_from", Label: labels["price_valid_from"], FieldType: ct.TableFieldTypeDate},
						{Name: "valid_to", Label: labels["price_valid_to"], FieldType: ct.TableFieldTypeDate},
						{Name: "customer_code", Label: labels["customer_code"], FieldType: ct.TableFieldTypeLink},
						{Name: "currency_code", Label: labels["currency_code"]},
						{Name: "qty", Label: labels["price_qty"], FieldType: ct.TableFieldTypeNumber},
						{Name: "price_value", Label: labels["price_value"], FieldType: ct.TableFieldTypeNumber},
					},
					Rows:              priceRows(),
					Pagination:        ct.PaginationTypeTop,
					PageSize:          5,
					HidePaginatonSize: true,
					RowSelected:       true,
					TableFilter:       true,
					FilterPlaceholder: labels["placeholder_filter"],
					AddItem:           !newInput,
					LabelAdd:          labels["price_new"],
				},
			}
		},
	}
	return tblMap[view]()
}

func (e *ProductEditor) Form(formKey string, labels cu.SM, data cu.IM) (form ct.Form) {
	formData := cu.ToIM(data, cu.IM{})
	footerRows := []ct.Row{
		{
			Columns: []ct.RowColumn{
				{Value: ct.Field{
					Type: ct.FieldTypeButton,
					Value: cu.IM{
						"name":         ct.FormEventOK,
						"type":         ct.ButtonTypeSubmit,
						"button_style": ct.ButtonStylePrimary,
						"icon":         ct.IconCheck,
						"label":        labels["editor_save"],
						"auto_focus":   true,
						"selected":     true,
					},
				}},
				{Value: ct.Field{
					Type: ct.FieldTypeButton,
					Value: cu.IM{
						"name":         ct.FormEventCancel,
						"type":         ct.ButtonTypeSubmit,
						"button_style": ct.ButtonStyleDefault,
						"icon":         ct.IconReply,
						"label":        labels["editor_back"],
					},
				}},
				{Value: ct.Field{
					Type:  ct.FieldTypeLabel,
					Value: cu.IM{},
				}},
				{Value: ct.Field{
					Type: ct.FieldTypeButton,
					Value: cu.IM{
						"name":         "form_delete",
						"type":         ct.ButtonTypeSubmit,
						"button_style": ct.ButtonStyleBorder,
						"icon":         ct.IconTimes,
						"label":        labels["editor_delete"],
						"style":        cu.SM{"color": "red", "fill": "red"},
					},
				}},
			},
			Full:         true,
			FieldCol:     false,
			BorderTop:    false,
			BorderBottom: false,
		},
	}
	frmMap := map[string]func() ct.Form{
		"prices": func() ct.Form {
			var price md.Price = md.Price{}
			ut.ConvertToType(formData, &price)
			currencies := cu.ToIMA(data["currencies"], []cu.IM{})
			customerSelectorRows := cu.ToIMA(data["customer_selector"], []cu.IM{})
			priceTypeOpt := func() (opt []ct.SelectOption) {
				opt = []ct.SelectOption{}
				for _, ptype := range []md.PriceType{
					md.PriceTypeCustomer, md.PriceTypeVendor,
				} {
					opt = append(opt, ct.SelectOption{
						Value: ptype.String(), Text: ptype.String(),
					})
				}
				return opt
			}
			currencyOpt := func() (opt []ct.SelectOption) {
				opt = []ct.SelectOption{}
				for _, currency := range currencies {
					opt = append(opt, ct.SelectOption{
						Value: cu.ToString(currency["code"], ""), Text: cu.ToString(currency["code"], ""),
					})
				}
				return opt
			}
			var customerSelectorFields []ct.TableField = []ct.TableField{
				{Name: "code", Label: labels["customer_code"]},
				{Name: "customer_name", Label: labels["customer_name"]},
				{Name: "tax_number", Label: labels["customer_tax_number"]},
			}
			return ct.Form{
				Title: labels["price_view"],
				Icon:  ct.IconMoney,
				BodyRows: []ct.Row{
					{Columns: []ct.RowColumn{
						{Label: labels["price_type"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "price_type_" + cu.ToString(price.Id, ""),
							},
							Type: ct.FieldTypeSelect, Value: cu.IM{
								"name":    "price_type",
								"options": priceTypeOpt(),
								"is_null": false,
								"value":   price.PriceType.String(),
							},
							FormTrigger: true,
						}},
						{Label: labels["price_valid_from"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "valid_from",
							},
							Type: ct.FieldTypeDate, Value: cu.IM{
								"name":    "valid_from",
								"value":   price.ValidFrom.String(),
								"is_null": false,
							},
							FormTrigger: true,
						}},
						{Label: labels["price_valid_to"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "valid_to",
							},
							Type: ct.FieldTypeDate, Value: cu.IM{
								"name":    "valid_to",
								"is_null": true,
								"value":   price.ValidTo.String(),
							},
							FormTrigger: true,
						}},
					}, Full: true, BorderBottom: true},
					{Columns: []ct.RowColumn{
						{Label: labels["currency_code"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "currency_code_" + cu.ToString(price.Id, ""),
							},
							Type: ct.FieldTypeSelect, Value: cu.IM{
								"name":    "currency_code",
								"options": currencyOpt(),
								"is_null": false,
								"value":   price.CurrencyCode,
							},
							FormTrigger: true,
						}},
						{Label: labels["price_qty"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "qty_" + cu.ToString(price.Id, ""),
							},
							Type: ct.FieldTypeNumber, Value: cu.IM{
								"name":  "qty",
								"value": price.Qty,
							},
							FormTrigger: true,
						}},
						{Label: labels["price_value"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "price_value_" + cu.ToString(price.Id, ""),
							},
							Type: ct.FieldTypeNumber, Value: cu.IM{
								"name":  "price_value",
								"value": price.PriceMeta.PriceValue,
							},
							FormTrigger: true,
						}},
					}, Full: true, BorderBottom: true},
					{Columns: []ct.RowColumn{
						{
							Label: labels["customer_code"], Value: ct.Field{
								BaseComponent: ct.BaseComponent{
									Name: "customer_code_" + cu.ToString(price.Id, ""),
								},
								Type: ct.FieldTypeSelector, Value: cu.IM{
									"name":  "customer_code",
									"title": labels["view_customer"],
									"value": ct.SelectOption{
										Value: price.CustomerCode,
										Text:  price.CustomerCode,
									},
									"fields":  customerSelectorFields,
									"rows":    customerSelectorRows,
									"link":    true,
									"is_null": true,
								},
								FormTrigger: true,
							},
						},
						{
							Label: labels["price_tags"], Value: ct.Field{
								BaseComponent: ct.BaseComponent{
									Name: "tags",
								},
								Type: ct.FieldTypeList, Value: cu.IM{
									"name":                "tags",
									"rows":                ut.ToTagList(price.PriceMeta.Tags),
									"label_value":         "tag",
									"pagination":          ct.PaginationTypeBottom,
									"page_size":           5,
									"hide_paginaton_size": true,
									"list_filter":         true,
									"filter_placeholder":  labels["placeholder_filter"],
									"add_item":            true,
									"add_icon":            ct.IconTag,
									"edit_item":           false,
									"delete_item":         true,
									"indicator":           ct.IndicatorSpinner,
								},
								FormTrigger: true,
							},
						},
					}, Full: true, BorderBottom: true},
				},
				FooterRows: footerRows,
			}
		},
		"events": func() ct.Form {
			var event md.Event = md.Event{}
			ut.ConvertToType(formData, &event)
			return ct.Form{
				Title: labels["event_view"],
				Icon:  ct.IconCalendar,
				BodyRows: []ct.Row{
					{Columns: []ct.RowColumn{
						{Label: labels["event_subject"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "subject",
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":  "subject",
								"value": event.Subject,
							},
							FormTrigger: true,
						}},
					}, Full: true, BorderBottom: true, FieldCol: true},
					{Columns: []ct.RowColumn{
						{Label: labels["event_start_time"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "start_time",
							},
							Type: ct.FieldTypeDateTime, Value: cu.IM{
								"name":    "start_time",
								"value":   event.StartTime.String(),
								"is_null": false,
							},
							FormTrigger: true,
						}},
						{Label: labels["event_end_time"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "end_time",
							},
							Type: ct.FieldTypeDateTime, Value: cu.IM{
								"name":    "end_time",
								"is_null": true,
								"value":   event.EndTime.String(),
							},
							FormTrigger: true,
						}},
						{Label: labels["event_place"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "place",
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":  "place",
								"value": event.Place,
							},
							FormTrigger: true,
						}},
					}, Full: true, BorderBottom: true},
					{Columns: []ct.RowColumn{
						{Label: labels["event_description"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "description",
							},
							Type: ct.FieldTypeText, Value: cu.IM{
								"name":  "description",
								"value": event.Description,
								"rows":  4,
							},
							FormTrigger: true,
						}},
						{
							Label: labels["event_tags"], Value: ct.Field{
								BaseComponent: ct.BaseComponent{
									Name: "tags",
								},
								Type: ct.FieldTypeList, Value: cu.IM{
									"name":                "tags",
									"rows":                ut.ToTagList(event.Tags),
									"label_value":         "tag",
									"pagination":          ct.PaginationTypeBottom,
									"page_size":           5,
									"hide_paginaton_size": true,
									"list_filter":         true,
									"filter_placeholder":  labels["placeholder_filter"],
									"add_item":            true,
									"add_icon":            ct.IconTag,
									"edit_item":           false,
									"delete_item":         true,
									"indicator":           ct.IndicatorSpinner,
								},
								FormTrigger: true,
							},
						},
					}, Full: true, BorderBottom: true},
				},
				FooterRows: footerRows,
			}
		},
	}

	if frm, found := frmMap[formKey]; found {
		return frm()
	}
	return ct.Form{}
}
