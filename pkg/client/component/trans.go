package component

import (
	"slices"
	"strings"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
	st "github.com/nervatura/nervatura/v6/static"
)

type TransEditor struct{}

func TransTypeIcon(transType string) string {
	switch transType {
	case md.TransTypeBank.String(), md.TransTypeCash.String():
		return ct.IconMoney
	case md.TransTypeWaybill.String():
		return ct.IconBriefcase
	case md.TransTypeProduction.String():
		return ct.IconFlask
	case md.TransTypeFormula.String():
		return ct.IconMagic
	case md.TransTypeDelivery.String(), md.TransTypeInventory.String():
		return ct.IconTruck
	}
	return ct.IconFileText
}

func TransIsItem(transType string) bool {
	return slices.Contains([]string{
		md.TransTypeInvoice.String(), md.TransTypeReceipt.String(), md.TransTypeOrder.String(),
		md.TransTypeOffer.String(), md.TransTypeWorksheet.String(), md.TransTypeRent.String()},
		transType,
	)
}

func TransIsPayment(transType string) bool {
	return slices.Contains([]string{
		md.TransTypeBank.String(), md.TransTypeCash.String()},
		transType,
	)
}

func TransIsMovement(transType string) bool {
	return slices.Contains([]string{
		md.TransTypeDelivery.String(), md.TransTypeInventory.String(), md.TransTypeWaybill.String(),
		md.TransTypeProduction.String(), md.TransTypeFormula.String()},
		transType,
	)
}

func (e *TransEditor) Frame(labels cu.SM, data cu.IM) (title, icon string) {
	return cu.ToString(data["editor_title"], labels["trans_title"]),
		cu.ToString(data["editor_icon"], ct.IconFileText)
}

func transSideBarItem(labels cu.SM, data cu.IM, stateOptions map[string]bool) (items []ct.SideBarItem) {
	var trans cu.IM = cu.ToIM(data["trans"], cu.IM{"trans_meta": cu.IM{}})
	transMeta := cu.ToIM(trans["trans_meta"], cu.IM{})
	enabled := !stateOptions["newInput"] && !stateOptions["dirty"] && !stateOptions["readonly"]
	transType := cu.ToString(trans["trans_type"], "")

	items = []ct.SideBarItem{}
	optionalItems := []func() (bool, func()){
		func() (bool, func()) {
			return cu.ToString(transMeta["status"], "") == md.TransStatusNormal.String() &&
					enabled, func() {
					items = append(items,
						&ct.SideBarSeparator{},
						&ct.SideBarElement{
							Name:  "transitem_new",
							Value: "transitem_new",
							Label: labels[strings.ToLower(transType+"_new")],
							Icon:  ct.IconFileText,
						})
				}
		},
		func() (bool, func()) {
			return enabled, func() {
				items = append(items,
					&ct.SideBarElement{
						Name:  "trans_copy",
						Value: "trans_copy",
						Label: labels["trans_copy"],
						Icon:  ct.IconCopy,
					})
			}
		},
		func() (bool, func()) {
			return transType != md.TransTypeReceipt.String() &&
					enabled, func() {
					items = append(items,
						&ct.SideBarElement{
							Name:  "trans_create",
							Value: "trans_create",
							Label: labels["trans_create"],
							Icon:  ct.IconSitemap,
						})
				}
		},
		func() (bool, func()) {
			return slices.Contains([]string{md.TransTypeInvoice.String(), md.TransTypeReceipt.String()}, transType) &&
					enabled, func() {
					items = append(items,
						&ct.SideBarElement{
							Name:  "trans_corrective",
							Value: "trans_corrective",
							Label: labels["trans_corrective"],
							Icon:  ct.IconShare,
						})
				}
		},
		func() (bool, func()) {
			return slices.Contains([]string{md.TransTypeInvoice.String(), md.TransTypeReceipt.String()}, transType) &&
					cu.ToString(trans["direction"], "") == md.DirectionOut.String() &&
					cu.ToBoolean(stateOptions["deleted"], false) && !stateOptions["guest"] && !stateOptions["transCancellations"], func() {
					items = append(items,
						&ct.SideBarElement{
							Name:  "trans_cancellation",
							Value: "trans_cancellation",
							Label: labels["trans_cancellation"],
							Icon:  ct.IconUndo,
						})
				}
		},
	}
	for _, fn := range optionalItems {
		if ok, fn := fn(); ok {
			fn()
		}
	}
	return items
}

func transSideBarPayment(labels cu.SM, data cu.IM, stateOptions map[string]bool) (items []ct.SideBarItem) {
	var trans cu.IM = cu.ToIM(data["trans"], cu.IM{"trans_meta": cu.IM{}})
	transMeta := cu.ToIM(trans["trans_meta"], cu.IM{})
	enabled := !stateOptions["newInput"] && !stateOptions["dirty"] && !stateOptions["readonly"]
	transType := cu.ToString(trans["trans_type"], "")

	items = []ct.SideBarItem{}
	optionalItems := []func() (bool, func()){
		func() (bool, func()) {
			return cu.ToString(transMeta["status"], "") == md.TransStatusNormal.String() &&
					enabled, func() {
					items = append(items,
						&ct.SideBarSeparator{},
						&ct.SideBarElement{
							Name:  "transpayment_new",
							Value: "transpayment_new",
							Label: labels[strings.ToLower(transType+"_new")],
							Icon:  ct.IconMoney,
						})
				}
		},
		func() (bool, func()) {
			return enabled, func() {
				items = append(items,
					&ct.SideBarElement{
						Name:  "trans_copy",
						Value: "trans_copy",
						Label: labels["trans_copy"],
						Icon:  ct.IconCopy,
					})
			}
		},
		func() (bool, func()) {
			return slices.Contains([]string{md.TransTypeCash.String()}, transType) &&
					enabled, func() {
					items = append(items,
						&ct.SideBarElement{
							Name:  "payment_link_add",
							Value: "payment_link_add",
							Label: strings.ToUpper(labels["payment_link_add"]),
							Icon:  ct.IconLink,
						})
				}
		},
		func() (bool, func()) {
			return slices.Contains([]string{md.TransTypeCash.String()}, transType) &&
					cu.ToBoolean(stateOptions["deleted"], false) && !stateOptions["guest"] && !stateOptions["transCancellations"], func() {
					items = append(items,
						&ct.SideBarElement{
							Name:  "trans_cancellation",
							Value: "trans_cancellation",
							Label: labels["trans_cancellation"],
							Icon:  ct.IconUndo,
						})
				}
		},
	}
	for _, fn := range optionalItems {
		if ok, fn := fn(); ok {
			fn()
		}
	}
	return items
}

func transSideBarMovement(labels cu.SM, data cu.IM, stateOptions map[string]bool) (items []ct.SideBarItem) {
	var trans cu.IM = cu.ToIM(data["trans"], cu.IM{"trans_meta": cu.IM{}})
	transMeta := cu.ToIM(trans["trans_meta"], cu.IM{})
	enabled := !stateOptions["newInput"] && !stateOptions["dirty"] && !stateOptions["readonly"]
	transType := cu.ToString(trans["trans_type"], "")
	direction := cu.ToString(trans["direction"], "")
	isDelivery := (transType == md.TransTypeDelivery.String() && direction != md.DirectionTransfer.String())

	items = []ct.SideBarItem{}
	optionalItems := []func() (bool, func()){
		func() (bool, func()) {
			return cu.ToString(transMeta["status"], "") == md.TransStatusNormal.String() &&
					enabled && !isDelivery, func() {
					items = append(items,
						&ct.SideBarSeparator{},
						&ct.SideBarElement{
							Name:  "transmovement_new",
							Value: "transmovement_new",
							Label: labels[strings.ToLower(transType+"_new")],
							Icon:  TransTypeIcon(transType),
						})
				}
		},
		func() (bool, func()) {
			return enabled && !isDelivery, func() {
				items = append(items,
					&ct.SideBarElement{
						Name:  "trans_copy",
						Value: "trans_copy",
						Label: labels["trans_copy"],
						Icon:  ct.IconCopy,
					})
			}
		},
		func() (bool, func()) {
			return slices.Contains([]string{md.TransTypeProduction.String()}, transType) &&
					enabled, func() {
					items = append(items,
						&ct.SideBarElement{
							Name:  "load_formula",
							Value: "load_formula",
							Label: strings.ToUpper(labels["trans_load_formula"]),
							Icon:  ct.IconMagic,
						})
				}
		},
		func() (bool, func()) {
			return slices.Contains([]string{md.TransTypeDelivery.String(), md.TransTypeInventory.String()}, transType) &&
					cu.ToString(transMeta["status"], "") == md.TransStatusNormal.String() &&
					enabled && !stateOptions["transCancellations"], func() {
					items = append(items,
						&ct.SideBarElement{
							Name:  "trans_cancellation",
							Value: "trans_cancellation",
							Label: labels["trans_cancellation"],
							Icon:  ct.IconUndo,
						})
				}
		},
	}
	for _, fn := range optionalItems {
		if ok, fn := fn(); ok {
			fn()
		}
	}
	return items
}

func transSideBarState(labels cu.SM, data cu.IM, stateOptions map[string]bool) (sb *ct.SideBarStatic) {
	var trans cu.IM = cu.ToIM(data["trans"], cu.IM{"trans_meta": cu.IM{}})
	transMeta := cu.ToIM(trans["trans_meta"], cu.IM{})
	state := strings.TrimPrefix(cu.ToString(transMeta["status"], md.TransStatusNormal.String()), "STATUS_")
	if stateOptions["newInput"] {
		state = "NEW"
	}
	if cu.ToBoolean(transMeta["closed"], false) {
		state = "CLOSED"
	}
	stateMap := map[string]*ct.SideBarStatic{
		"NEW": {
			Icon: ct.IconPlus, Label: labels["state_new"], Color: "yellow",
		},
		"DELETED": {
			Icon: ct.IconExclamationTriangle, Label: labels["state_deleted"], Color: "red",
		},
		"CLOSED": {
			Icon: ct.IconLock, Label: labels["state_closed"], Color: "brown",
		},
		"NORMAL": {
			Icon: ct.IconEdit, Label: labels["state_edit"], Color: "green",
		},
		"CANCELLATION": {
			Icon: ct.IconExclamationTriangle, Label: labels["state_cancellation"], Color: "orange",
		},
		"AMENDMENT": {
			Icon: ct.IconEdit, Label: labels["state_amendment"], Color: "orange",
		},
	}
	sb = stateMap["NORMAL"]
	var ok bool
	if _, ok = stateMap[state]; ok {
		sb = stateMap[state]
	}
	return sb
}

func (e *TransEditor) SideBar(labels cu.SM, data cu.IM) (items []ct.SideBarItem) {
	var trans cu.IM = cu.ToIM(data["trans"], cu.IM{"trans_meta": cu.IM{}})
	transMeta := cu.ToIM(trans["trans_meta"], cu.IM{})
	user := cu.ToIM(data["user"], cu.IM{})
	transCancellations := cu.ToIMA(data["trans_cancellation"], []cu.IM{})
	transType := cu.ToString(trans["trans_type"], "")
	isDeleted := (cu.ToString(transMeta["status"], "") == md.TransStatusDeleted.String())

	stateOptions := map[string]bool{
		"newInput": (cu.ToInteger(trans["id"], 0) == 0),
		"dirty":    cu.ToBoolean(data["dirty"], false),
		"deleted":  isDeleted,
		"closed":   cu.ToBoolean(transMeta["closed"], false),
		"readonly": (cu.ToString(user["user_group"], "") == md.UserGroupGuest.String()) || isDeleted ||
			(cu.ToBoolean(transMeta["closed"], false) && !cu.ToBoolean(data["dirty"], false)),
		"guest":              cu.ToString(user["user_group"], "") == md.UserGroupGuest.String(),
		"transCancellations": len(transCancellations) > 0,
	}

	updateLabel := labels["editor_save"]
	if stateOptions["newInput"] {
		updateLabel = labels["editor_create"]
	}

	items = []ct.SideBarItem{
		&ct.SideBarSeparator{},
		&ct.SideBarElement{
			Name:    "editor_cancel",
			Value:   "editor_cancel",
			Label:   labels["browser_title"],
			Icon:    ct.IconReply,
			NotFull: true,
		},
		&ct.SideBarSeparator{},
		transSideBarState(labels, data, stateOptions),
	}

	if !stateOptions["readonly"] {
		items = append(items,
			&ct.SideBarSeparator{},
			&ct.SideBarElement{
				Name:     "editor_save",
				Value:    "editor_save",
				Label:    updateLabel,
				Icon:     ct.IconUpload,
				Selected: stateOptions["dirty"],
			})
	}

	if !isDeleted &&
		!slices.Contains([]string{md.TransTypeDelivery.String(), md.TransTypeInventory.String()}, transType) &&
		(cu.ToString(transMeta["status"], "") == md.TransStatusNormal.String()) &&
		!stateOptions["newInput"] && !stateOptions["readonly"] {
		items = append(items, &ct.SideBarElement{
			Name:  "editor_delete",
			Value: "editor_delete",
			Label: labels["editor_delete"],
			Icon:  ct.IconTimes,
		})
	}

	if TransIsItem(transType) {
		items = append(items, transSideBarItem(labels, data, stateOptions)...)
	}

	if TransIsPayment(transType) {
		items = append(items, transSideBarPayment(labels, data, stateOptions)...)
	}

	if TransIsMovement(transType) {
		items = append(items, transSideBarMovement(labels, data, stateOptions)...)
	}

	items = append(items,
		&ct.SideBarSeparator{},
		&ct.SideBarElement{
			Name:     "editor_report",
			Value:    "editor_report",
			Label:    labels["editor_report"],
			Icon:     ct.IconChartBar,
			Disabled: stateOptions["newInput"] || stateOptions["dirty"],
		},
		&ct.SideBarSeparator{},
		&ct.SideBarElement{
			Name:     "editor_bookmark",
			Value:    "editor_bookmark",
			Label:    labels["editor_bookmark"],
			Icon:     ct.IconStar,
			Disabled: stateOptions["newInput"],
		},
		&ct.SideBarElementLink{
			SideBarElement: ct.SideBarElement{
				Name:  "editor_help",
				Value: "editor_help",
				Label: labels["editor_help"],
				Icon:  ct.IconQuestionCircle,
			},
			Href:       st.DocsClientPath + "#document",
			LinkTarget: "_blank",
		})
	return items
}

func (e *TransEditor) View(labels cu.SM, data cu.IM) (views []ct.EditorView) {
	var trans cu.IM = cu.ToIM(data["trans"], cu.IM{})
	transMap := cu.ToIM(trans["trans_map"], cu.IM{})
	items := cu.ToIMA(data["items"], []cu.IM{})
	invoiceItems := cu.ToIMA(data["transitem_invoice"], []cu.IM{})
	payments := cu.ToIMA(data["payments"], []cu.IM{})
	paymentLink := cu.ToIMA(data["payment_link"], []cu.IM{})
	movements := transMovementRows(data)
	shippingItems := cu.ToIMA(data["transitem_shipping"], []cu.IM{})
	toolMovement := cu.ToIMA(data["tool_movement"], []cu.IM{})

	newInput := (cu.ToInteger(trans["id"], 0) == 0)
	transType := cu.ToString(trans["trans_type"], "")
	direction := strings.TrimPrefix(cu.ToString(trans["direction"], ""), "DIRECTION_")
	transLabel := cu.ToString(labels[strings.ToLower(transType+"_"+direction)], labels[strings.ToLower(transType)])

	if newInput {
		return []ct.EditorView{
			{
				Key:   "trans",
				Label: transLabel,
				Icon:  TransTypeIcon(transType),
			},
		}
	}
	views = []ct.EditorView{
		{
			Key:   "trans",
			Label: transLabel,
			Icon:  TransTypeIcon(transType),
		},
		{
			Key:   "maps",
			Label: labels["map_view"],
			Icon:  ct.IconDatabase,
			Badge: cu.ToString(int64(len(transMap)), "0"),
		},
	}
	if TransIsItem(transType) {
		views = append(views,
			ct.EditorView{
				Key:   "items",
				Label: labels["items_view"],
				Icon:  ct.IconListOl,
				Badge: cu.ToString(int64(len(items)), "0"),
			}, ct.EditorView{
				Key:   "tool_movement",
				Label: labels["tool_movement_view"],
				Icon:  ct.IconBriefcase,
				Badge: cu.ToString(int64(len(toolMovement)), "0"),
			})
	}
	if slices.Contains([]string{md.TransTypeOrder.String(), md.TransTypeRent.String(), md.TransTypeWorksheet.String()}, transType) {
		views = append(views,
			ct.EditorView{
				Key:   "transitem_invoice",
				Label: labels["transitem_invoice_view"],
				Icon:  ct.IconListOl,
				Badge: cu.ToString(int64(len(invoiceItems)), "0"),
			},
			ct.EditorView{
				Key:   "transitem_shipping",
				Label: labels["transitem_shipping_view"],
				Icon:  ct.IconTruck,
				Badge: cu.ToString(int64(len(shippingItems)), "0"),
			})
	}
	if transType == md.TransTypeBank.String() {
		views = append(views, ct.EditorView{
			Key:   "payments",
			Label: labels["payments_view"],
			Icon:  ct.IconListOl,
			Badge: cu.ToString(int64(len(payments)), "0"),
		})
	}
	if slices.Contains([]string{md.TransTypeBank.String(), md.TransTypeCash.String(),
		md.TransTypeInvoice.String(), md.TransTypeReceipt.String()}, transType) {
		views = append(views, ct.EditorView{
			Key:   "payment_link",
			Label: transTypeLabel(labels, transType, "payment_link_view"),
			Icon:  ct.IconFileText,
			Badge: cu.ToString(int64(len(paymentLink)), "0"),
		})
	}
	if TransIsMovement(transType) {
		views = append(views, ct.EditorView{
			Key:   "movements",
			Label: labels["movements_view"],
			Icon:  ct.IconListOl,
			Badge: cu.ToString(int64(len(movements)), "0"),
		})
	}
	return views
}

func transTypeLabel(labels cu.SM, transType string, key string) string {
	return cu.ToString(labels[key+"_"+strings.ToLower(strings.Split(transType, "_")[1])], labels[key])
}

func transMainHeadRow(trans md.Trans, labels cu.SM, _ cu.IM) (rows []ct.Row) {
	transStateOpt := func() (opt []ct.SelectOption) {
		opt = []ct.SelectOption{}
		for _, state := range []md.TransState{
			md.TransStateOK, md.TransStateNew, md.TransStateBack,
		} {
			opt = append(opt, ct.SelectOption{
				Value: state.String(), Text: state.String(),
			})
		}
		return opt
	}

	directionOpt := func() (opt []ct.SelectOption) {
		opt = []ct.SelectOption{}
		for _, direction := range []md.Direction{
			md.DirectionOut, md.DirectionIn,
		} {
			opt = append(opt, ct.SelectOption{
				Value: direction.String(), Text: direction.String(),
			})
		}
		if slices.Contains([]string{
			md.TransTypeBank.String(), md.TransTypeInventory.String(), md.TransTypeFormula.String(), md.TransTypeProduction.String()}, trans.TransType.String()) ||
			(trans.TransType == md.TransTypeDelivery && trans.Direction == md.DirectionTransfer) {
			opt = []ct.SelectOption{{
				Value: md.DirectionTransfer.String(), Text: md.DirectionTransfer.String(),
			}}
		}
		return opt
	}

	rows = []ct.Row{{Columns: []ct.RowColumn{
		{Label: labels["trans_code"], Value: ct.Field{
			BaseComponent: ct.BaseComponent{
				Name: "code_" + cu.ToString(trans.Id, ""),
			},
			Type: ct.FieldTypeString, Value: cu.IM{
				"name":     "code",
				"value":    trans.Code,
				"disabled": true,
			},
		}},
		{Label: labels["trans_ref_number"], Value: ct.Field{
			BaseComponent: ct.BaseComponent{
				Name: "ref_number_" + cu.ToString(trans.Id, ""),
			},
			Type: ct.FieldTypeString, Value: cu.IM{
				"name":  "ref_number",
				"value": trans.TransMeta.RefNumber,
			},
		}},
		{Label: labels["trans_state"], Value: ct.Field{
			BaseComponent: ct.BaseComponent{
				Name: "trans_state_" + cu.ToString(trans.Id, ""),
			},
			Type: ct.FieldTypeSelect, Value: cu.IM{
				"name":    "trans_state",
				"options": transStateOpt(),
				"is_null": false,
				"value":   trans.TransMeta.TransState.String(),
			},
		}},
	}, Full: true, BorderBottom: true},
		{Columns: []ct.RowColumn{
			{Label: labels["trans_direction"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "trans_direction_" + cu.ToString(trans.Id, ""),
				},
				Type: ct.FieldTypeSelect, Value: cu.IM{
					"name":     "direction",
					"options":  directionOpt(),
					"is_null":  false,
					"value":    trans.Direction.String(),
					"disabled": (trans.Id > 0),
				},
			}},
			{Label: labels["trans_time_stamp"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "time_stamp_" + cu.ToString(trans.Id, ""),
				},
				Type: ct.FieldTypeDate, Value: cu.IM{
					"name":     "time_stamp",
					"is_null":  false,
					"value":    trans.TimeStamp,
					"disabled": true,
				},
			}},
			{Label: transTypeLabel(labels, trans.TransType.String(), "trans_date"),
				Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "trans_date_" + cu.ToString(trans.Id, ""),
					},
					Type: ct.FieldTypeDate, Value: cu.IM{
						"name":    "trans_date",
						"is_null": false,
						"value":   trans.TransDate,
					},
				}},
		}, Full: true, BorderBottom: true}}
	return rows
}

func transMainFooterRow(trans md.Trans, labels cu.SM, _ cu.IM) (rows []ct.Row) {
	rows = []ct.Row{
		{Columns: []ct.RowColumn{
			{Label: labels["trans_notes"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "notes_" + cu.ToString(trans.Id, ""),
				},
				Type: ct.FieldTypeText, Value: cu.IM{
					"name":  "notes",
					"value": trans.TransMeta.Notes,
					"rows":  4,
				},
			}},
			{
				Label: labels["trans_tags"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "tags_" + cu.ToString(trans.Id, ""),
					},
					Type: ct.FieldTypeList, Value: cu.IM{
						"name":                "tags",
						"rows":                ut.ToTagList(trans.TransMeta.Tags),
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
		{Columns: []ct.RowColumn{
			{Label: labels["trans_internal_notes"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "internal_notes_" + cu.ToString(trans.Id, ""),
				},
				Type: ct.FieldTypeText, Value: cu.IM{
					"name":  "internal_notes",
					"value": trans.TransMeta.InternalNotes,
					"rows":  4,
				},
			}},
			{Label: labels["trans_report_notes"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "report_notes_" + cu.ToString(trans.Id, ""),
				},
				Type: ct.FieldTypeText, Value: cu.IM{
					"name":  "report_notes",
					"value": trans.TransMeta.ReportNotes,
					"rows":  4,
				},
			}},
		}, Full: true, BorderBottom: true},
	}
	return rows
}

func transMainItemRow(trans md.Trans, labels cu.SM, data cu.IM) (rows []ct.Row) {
	currencies := cu.ToIMA(data["currencies"], []cu.IM{})
	customerSelectorRows := cu.ToIMA(data["customer_selector"], []cu.IM{})
	customerName := cu.ToString(data["customer_name"], "")
	transitemSelectorRows := cu.ToIMA(data["transitem_selector"], []cu.IM{})
	employeeSelectorRows := cu.ToIMA(data["employee_selector"], []cu.IM{})
	projectSelectorRows := cu.ToIMA(data["project_selector"], []cu.IM{})

	currencyOpt := func() (opt []ct.SelectOption) {
		opt = []ct.SelectOption{}
		for _, currency := range currencies {
			opt = append(opt, ct.SelectOption{
				Value: cu.ToString(currency["code"], ""), Text: cu.ToString(currency["code"], ""),
			})
		}
		return opt
	}

	paidTypeOpt := func() (opt []ct.SelectOption) {
		opt = []ct.SelectOption{}
		for _, paidType := range []md.PaidType{
			md.PaidTypeOnline, md.PaidTypeCard, md.PaidTypeTransfer, md.PaidTypeCash, md.PaidTypeOther,
		} {
			opt = append(opt, ct.SelectOption{
				Value: paidType.String(), Text: paidType.String(),
			})
		}
		return opt
	}

	var customerSelectorFields []ct.TableField = []ct.TableField{
		{Name: "code", Label: labels["customer_code"]},
		{Name: "customer_name", Label: labels["customer_name"]},
		{Name: "tax_number", Label: labels["customer_tax_number"]},
	}

	var transitemSelectorFields []ct.TableField = []ct.TableField{
		{Name: "code", Label: labels["trans_code"]},
		{Name: "trans_date", Label: labels["trans_date"]},
		{Name: "trans_type", Label: labels["trans_type"]},
		{Name: "direction", Label: labels["trans_direction"]},
	}

	var employeeSelectorFields []ct.TableField = []ct.TableField{
		{Name: "code", Label: labels["employee_code"]},
		{Name: "first_name", Label: labels["contact_first_name"]},
		{Name: "surname", Label: labels["contact_surname"]},
		{Name: "status", Label: labels["contact_status"]},
	}

	var projectSelectorFields []ct.TableField = []ct.TableField{
		{Name: "code", Label: labels["project_code"]},
		{Name: "project_name", Label: labels["project_name"]},
		{Name: "customer_code", Label: labels["customer_code"]},
	}

	rows = transMainHeadRow(trans, labels, data)
	rows[1].Columns = append(rows[1].Columns, ct.RowColumn{Label: transTypeLabel(labels, trans.TransType.String(), "trans_due_time"),
		Value: ct.Field{
			BaseComponent: ct.BaseComponent{
				Name: "due_time_" + cu.ToString(trans.Id, ""),
			},
			Type: ct.FieldTypeDate, Value: cu.IM{
				"name":    "due_time",
				"is_null": false,
				"value":   trans.TransMeta.DueTime,
			},
		}})

	transCodeValue := func() ct.Field {
		if trans.TransMeta.Status != md.TransStatusNormal {
			return ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "trans_code_" + cu.ToString(trans.Id, ""),
				},
				Type: ct.FieldTypeLink,
				Value: cu.IM{
					"name":  "trans_code",
					"value": trans.TransCode,
				},
			}
		}
		return ct.Field{
			BaseComponent: ct.BaseComponent{
				Name: "trans_code_" + cu.ToString(trans.Id, ""),
			},
			Type: ct.FieldTypeSelector,
			Value: cu.IM{
				"name":  "transitem_code",
				"title": labels["transitem_view"],
				"value": ct.SelectOption{
					Value: trans.TransCode,
					Text:  trans.TransCode,
				},
				"fields":  transitemSelectorFields,
				"rows":    transitemSelectorRows,
				"link":    true,
				"is_null": true,
			},
			FormTrigger: true,
		}
	}

	if trans.TransType == md.TransTypeReceipt {
		rows = append(rows, ct.Row{Columns: []ct.RowColumn{
			{Label: labels["trans_paid_type"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "paid_type_" + cu.ToString(trans.Id, ""),
				},
				Type: ct.FieldTypeSelect, Value: cu.IM{
					"name":    "paid_type",
					"options": paidTypeOpt(),
					"is_null": false,
					"value":   trans.TransMeta.PaidType.String(),
				},
			}},
			{
				Label: labels["trans_trans_code"],
				Value: transCodeValue(),
			},
		},
			Full: true, BorderBottom: true},
			ct.Row{Columns: []ct.RowColumn{
				{
					Label: labels["employee_code"], Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "employee_code_" + cu.ToString(trans.Id, ""),
						},
						Type: ct.FieldTypeSelector, Value: cu.IM{
							"name":  "employee_code",
							"title": labels["employee_view"],
							"value": ct.SelectOption{
								Value: trans.EmployeeCode,
								Text:  trans.EmployeeCode,
							},
							"fields":  employeeSelectorFields,
							"rows":    employeeSelectorRows,
							"link":    true,
							"is_null": true,
						},
						FormTrigger: true,
					},
				},
				{
					Label: labels["project_code"], Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "project_code_" + cu.ToString(trans.Id, ""),
						},
						Type: ct.FieldTypeSelector, Value: cu.IM{
							"name":  "project_code",
							"title": labels["project_view"],
							"value": ct.SelectOption{
								Value: trans.ProjectCode,
								Text:  trans.ProjectCode,
							},
							"fields":  projectSelectorFields,
							"rows":    projectSelectorRows,
							"link":    true,
							"is_null": true,
						},
						FormTrigger: true,
					},
				},
			},
				Full: true, BorderBottom: true})
	} else {
		rows = append(rows, ct.Row{Columns: []ct.RowColumn{
			{
				Label: labels["customer_name"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "customer_code_" + cu.ToString(trans.Id, ""),
					},
					Type: ct.FieldTypeSelector, Value: cu.IM{
						"name":  "customer_code",
						"title": labels["view_customer"],
						"value": ct.SelectOption{
							Value: trans.CustomerCode,
							Text:  customerName,
						},
						"fields":  customerSelectorFields,
						"rows":    customerSelectorRows,
						"link":    true,
						"is_null": false,
					},
					FormTrigger: true,
				},
			},
			{
				Label: labels["trans_trans_code"],
				Value: transCodeValue(),
			},
		},
			Full: true, BorderBottom: true},
			ct.Row{Columns: []ct.RowColumn{
				{
					Label: labels["employee_code"], Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "employee_code_" + cu.ToString(trans.Id, ""),
						},
						Type: ct.FieldTypeSelector, Value: cu.IM{
							"name":  "employee_code",
							"title": labels["employee_view"],
							"value": ct.SelectOption{
								Value: trans.EmployeeCode,
								Text:  trans.EmployeeCode,
							},
							"fields":  employeeSelectorFields,
							"rows":    employeeSelectorRows,
							"link":    true,
							"is_null": true,
						},
						FormTrigger: true,
					},
				},
				{
					Label: labels["project_code"], Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "project_code_" + cu.ToString(trans.Id, ""),
						},
						Type: ct.FieldTypeSelector, Value: cu.IM{
							"name":  "project_code",
							"title": labels["project_view"],
							"value": ct.SelectOption{
								Value: trans.ProjectCode,
								Text:  trans.ProjectCode,
							},
							"fields":  projectSelectorFields,
							"rows":    projectSelectorRows,
							"link":    true,
							"is_null": true,
						},
						FormTrigger: true,
					},
				},
				{Label: labels["trans_paid_type"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "paid_type_" + cu.ToString(trans.Id, ""),
					},
					Type: ct.FieldTypeSelect, Value: cu.IM{
						"name":    "paid_type",
						"options": paidTypeOpt(),
						"is_null": false,
						"value":   trans.TransMeta.PaidType.String(),
					},
				}},
			},
				Full: true, BorderBottom: true})
	}
	if trans.TransType == md.TransTypeWorksheet {
		rows = append(rows, ct.Row{
			Columns: []ct.RowColumn{{
				Label: labels["trans_worksheet_distance"],
				Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "worksheet_distance_" + cu.ToString(trans.Id, ""),
					},
					Type: ct.FieldTypeNumber, Value: cu.IM{
						"name":  "worksheet_distance",
						"value": trans.TransMeta.Worksheet.Distance,
					},
				}},
				{
					Label: labels["trans_worksheet_repair"],
					Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "worksheet_repair_" + cu.ToString(trans.Id, ""),
						},
						Type: ct.FieldTypeNumber, Value: cu.IM{
							"name":  "worksheet_repair",
							"value": trans.TransMeta.Worksheet.Repair,
						},
					}},
				{
					Label: labels["trans_worksheet_total"],
					Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "worksheet_total_" + cu.ToString(trans.Id, ""),
						},
						Type: ct.FieldTypeNumber, Value: cu.IM{
							"name":  "worksheet_total",
							"value": trans.TransMeta.Worksheet.Total,
						},
					}},
			}, Full: true, BorderBottom: true,
		},
			ct.Row{
				Columns: []ct.RowColumn{
					{
						Label: labels["trans_worksheet_justification"],
						Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "worksheet_justification_" + cu.ToString(trans.Id, ""),
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":  "worksheet_justification",
								"value": trans.TransMeta.Worksheet.Justification,
							},
						}},
				}, Full: true, BorderBottom: true,
			})
	}
	if trans.TransType == md.TransTypeRent {
		rows = append(rows, ct.Row{
			Columns: []ct.RowColumn{{
				Label: labels["trans_rent_holiday"],
				Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "rent_holiday_" + cu.ToString(trans.Id, ""),
					},
					Type: ct.FieldTypeNumber, Value: cu.IM{
						"name":  "rent_holiday",
						"value": trans.TransMeta.Rent.Holiday,
					},
				}},
				{
					Label: labels["trans_rent_bad_tool"],
					Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "rent_bad_tool_" + cu.ToString(trans.Id, ""),
						},
						Type: ct.FieldTypeNumber, Value: cu.IM{
							"name":  "rent_bad_tool",
							"value": trans.TransMeta.Rent.BadTool,
						},
					}},
				{
					Label: labels["trans_rent_other"],
					Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "rent_other_" + cu.ToString(trans.Id, ""),
						},
						Type: ct.FieldTypeNumber, Value: cu.IM{
							"name":  "rent_other",
							"value": trans.TransMeta.Rent.Other,
						},
					}},
			}, Full: true, BorderBottom: true,
		},
			ct.Row{
				Columns: []ct.RowColumn{
					{
						Label: labels["trans_rent_justification"],
						Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "rent_justification_" + cu.ToString(trans.Id, ""),
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":  "rent_justification",
								"value": trans.TransMeta.Rent.Justification,
							},
						}},
				}, Full: true, BorderBottom: true,
			})
	}
	rows = append(rows,
		ct.Row{Columns: []ct.RowColumn{
			{Label: labels["currency_code"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "currency_code_" + cu.ToString(trans.Id, ""),
				},
				Type: ct.FieldTypeSelect, Value: cu.IM{
					"name":    "currency_code",
					"options": currencyOpt(),
					"is_null": false,
					"value":   trans.CurrencyCode,
				},
			}},
			{Label: transTypeLabel(labels, trans.TransType.String(), "trans_rate"), Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "rate_" + cu.ToString(trans.Id, ""),
				},
				Type: ct.FieldTypeNumber, Value: cu.IM{
					"name":  "rate",
					"value": trans.TransMeta.Rate,
				},
			}},
			{Label: transTypeLabel(labels, trans.TransType.String(), "trans_paid"), Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "paid_" + cu.ToString(trans.Id, ""),
				},
				Type: ct.FieldTypeBool, Value: cu.IM{
					"name":  "paid",
					"value": cu.ToBoolean(trans.TransMeta.Paid, false),
				},
			}},
			{Label: labels["trans_closed"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "closed_" + cu.ToString(trans.Id, ""),
				},
				Type: ct.FieldTypeBool, Value: cu.IM{
					"name":  "closed",
					"value": cu.ToBoolean(trans.TransMeta.Closed, false),
				},
			}},
		}, Full: true, BorderBottom: true})

	rows = append(rows, transMainFooterRow(trans, labels, data)...)
	return rows
}

func transMainPaymentRow(trans md.Trans, labels cu.SM, data cu.IM) (rows []ct.Row) {
	payments := cu.ToIMA(data["payments"], []cu.IM{})
	cashPayment := cu.ToIM(payments[0], cu.IM{})
	cashPaymentMeta := cu.ToIM(cashPayment["payment_meta"], cu.IM{})
	employeeSelectorRows := cu.ToIMA(data["employee_selector"], []cu.IM{})
	places := cu.ToIMA(data["places"], []cu.IM{})

	cashAmount := func() (amount float64) {
		amount = cu.ToFloat(cashPaymentMeta["amount"], 0)
		if trans.Direction == md.DirectionOut {
			amount = -amount
		}
		return amount
	}

	placeOpt := func(placeType md.PlaceType) (opt []ct.SelectOption) {
		opt = []ct.SelectOption{}
		for _, place := range places {
			if cu.ToString(place["place_type"], "") == placeType.String() {
				opt = append(opt, ct.SelectOption{
					Value: cu.ToString(place["code"], ""), Text: cu.ToString(place["place_name"], ""),
				})
			}
		}
		return opt
	}

	var employeeSelectorFields []ct.TableField = []ct.TableField{
		{Name: "code", Label: labels["employee_code"]},
		{Name: "first_name", Label: labels["contact_first_name"]},
		{Name: "surname", Label: labels["contact_surname"]},
		{Name: "status", Label: labels["contact_status"]},
	}

	rows = transMainHeadRow(trans, labels, data)
	if trans.TransType == md.TransTypeBank {
		rows[1].Columns[0] = ct.RowColumn{
			Label: transTypeLabel(labels, trans.TransType.String(), "trans_place_code"),
			Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "place_code_" + cu.ToString(trans.Id, ""),
				},
				Type: ct.FieldTypeSelect, Value: cu.IM{
					"name":    "place_code",
					"options": placeOpt(md.PlaceTypeBank),
					"is_null": true,
					"value":   trans.PlaceCode,
				},
			}}
	}
	if trans.TransType == md.TransTypeCash {
		rows[0].Columns[1] = ct.RowColumn{
			Label: labels["employee_code"], Value: ct.Field{
				BaseComponent: ct.BaseComponent{
					Name: "employee_code_" + cu.ToString(trans.Id, ""),
				},
				Type: ct.FieldTypeSelector, Value: cu.IM{
					"name":  "employee_code",
					"title": labels["employee_view"],
					"value": ct.SelectOption{
						Value: trans.EmployeeCode,
						Text:  trans.EmployeeCode,
					},
					"fields":  employeeSelectorFields,
					"rows":    employeeSelectorRows,
					"link":    true,
					"is_null": true,
				},
				FormTrigger: true,
			},
		}
		rows = append(rows,
			ct.Row{Columns: []ct.RowColumn{
				{
					Label: transTypeLabel(labels, trans.TransType.String(), "trans_place_code"),
					Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "place_code_" + cu.ToString(trans.Id, ""),
						},
						Type: ct.FieldTypeSelect, Value: cu.IM{
							"name":    "place_code",
							"options": placeOpt(md.PlaceTypeCash),
							"is_null": true,
							"value":   trans.PlaceCode,
						},
					}},
				{Label: labels["payment_paid_date"],
					Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "payment_paid_date_" + cu.ToString(trans.Id, ""),
						},
						Type: ct.FieldTypeDate, Value: cu.IM{
							"name":    "payment_paid_date",
							"is_null": false,
							"value":   cashPayment["paid_date"],
						},
					}},
				{
					Label: labels["payment_amount"],
					Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "payment_amount_" + cu.ToString(trans.Id, ""),
						},
						Type: ct.FieldTypeNumber, Value: cu.IM{
							"name":  "payment_amount",
							"value": cashAmount(),
						},
					}},
			},
				Full: true, BorderBottom: true})
	}
	rows = append(rows, transMainFooterRow(trans, labels, data)...)
	return rows
}

func transMainMovementRow(trans md.Trans, labels cu.SM, data cu.IM) (rows []ct.Row) {
	movements := cu.ToIMA(data["movements"], []cu.IM{})
	newInput := (trans.Id == 0)
	movementInventory := cu.ToIMA(data["movement_inventory"], []cu.IM{})
	places := cu.ToIMA(data["places"], []cu.IM{})

	transitemSelectorRows := cu.ToIMA(data["transitem_selector"], []cu.IM{})
	employeeSelectorRows := cu.ToIMA(data["employee_selector"], []cu.IM{})
	productSelectorRows := cu.ToIMA(data["product_selector"], []cu.IM{})
	customerSelectorRows := cu.ToIMA(data["customer_selector"], []cu.IM{})
	customerName := cu.ToString(data["customer_name"], "")

	refTransCode := func() (refTransCode string) {
		if len(movementInventory) > 0 {
			refTransCode = cu.ToString(movementInventory[0]["ref_trans_code"], "")
		}
		return refTransCode
	}

	targetPlaceCode := func() (targetPlaceCode string) {
		for _, movement := range movements {
			if cu.ToString(movement["movement_code"], "") != "" || cu.ToInteger(movement["id"], 0) < 0 {
				targetPlaceCode = cu.ToString(movement["place_code"], "")
				break
			}
		}
		return targetPlaceCode
	}

	productionHeadRow := func() (productCode, batchNo string, qty float64) {
		if idx := slices.IndexFunc(movements, func(movement cu.IM) bool {
			movementMeta := cu.ToIM(movement["movement_meta"], cu.IM{})
			return cu.ToBoolean(movementMeta["shared"], false)
		}); idx > -1 {
			movementMeta := cu.ToIM(movements[idx]["movement_meta"], cu.IM{})
			productCode = cu.ToString(movements[idx]["product_code"], "")
			batchNo = cu.ToString(movementMeta["notes"], "")
			qty = cu.ToFloat(movementMeta["qty"], 0)
			return productCode, batchNo, qty
		}
		return productCode, batchNo, qty
	}

	formulaHeadRow := func() (productCode, batchNo string, qty float64) {
		if idx := slices.IndexFunc(movements, func(movement cu.IM) bool {
			return cu.ToString(movement["movement_type"], "") == md.MovementTypeHead.String()
		}); idx > -1 {
			movementMeta := cu.ToIM(movements[idx]["movement_meta"], cu.IM{})
			productCode = cu.ToString(movements[idx]["product_code"], "")
			batchNo = cu.ToString(movementMeta["notes"], "")
			qty = cu.ToFloat(movementMeta["qty"], 0)
			return productCode, batchNo, qty
		}
		return productCode, batchNo, qty
	}

	placeOpt := func() (opt []ct.SelectOption) {
		opt = []ct.SelectOption{}
		for _, place := range places {
			if cu.ToString(place["place_type"], "") == md.PlaceTypeWarehouse.String() {
				opt = append(opt, ct.SelectOption{
					Value: cu.ToString(place["code"], ""), Text: cu.ToString(place["place_name"], ""),
				})
			}
		}
		return opt
	}

	var transitemSelectorFields []ct.TableField = []ct.TableField{
		{Name: "code", Label: labels["trans_code"]},
		{Name: "trans_date", Label: labels["trans_date"]},
		{Name: "trans_type", Label: labels["trans_type"]},
		{Name: "direction", Label: labels["trans_direction"]},
	}

	var customerSelectorFields []ct.TableField = []ct.TableField{
		{Name: "code", Label: labels["customer_code"]},
		{Name: "customer_name", Label: labels["customer_name"]},
		{Name: "tax_number", Label: labels["customer_tax_number"]},
	}

	var employeeSelectorFields []ct.TableField = []ct.TableField{
		{Name: "code", Label: labels["employee_code"]},
		{Name: "first_name", Label: labels["contact_first_name"]},
		{Name: "surname", Label: labels["contact_surname"]},
		{Name: "status", Label: labels["contact_status"]},
	}

	var productSelectorFields []ct.TableField = []ct.TableField{
		{Name: "code", Label: labels["product_code"]},
		{Name: "product_name", Label: labels["product_name"]},
		{Name: "product_type", Label: labels["product_type"]},
		{Name: "tag_lst", Label: labels["product_tags"]},
	}

	var rowMap = map[md.TransType]func(){
		md.TransTypeDelivery: func() {
			if trans.Direction != md.DirectionTransfer {
				rows[0].Columns[1] = ct.RowColumn{
					Label: labels["trans_trans_code"],
					Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "trans_code_" + cu.ToString(trans.Id, ""),
						},
						Type: ct.FieldTypeLink,
						Value: cu.IM{
							"name":  "trans_code",
							"value": refTransCode(),
						},
					}}
				rows[1].Columns[2].Value.Value["disabled"] = true
				rows[1].Columns = append(rows[1].Columns, ct.RowColumn{
					Label: labels["place_name_movement"],
					Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "place_code_" + cu.ToString(trans.Id, ""),
						},
						Type: ct.FieldTypeSelect, Value: cu.IM{
							"name":     "place_code",
							"options":  placeOpt(),
							"is_null":  true,
							"value":    trans.PlaceCode,
							"disabled": true,
						},
					}})
				return
			}
			rows[1].Columns[0].Value.Value["disabled"] = true
			rows[1].Columns = append(rows[1].Columns, ct.RowColumn{
				Label: labels["trans_closed"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "closed_" + cu.ToString(trans.Id, ""),
					},
					Type: ct.FieldTypeBool, Value: cu.IM{
						"name":  "closed",
						"value": cu.ToBoolean(trans.TransMeta.Closed, false),
					},
				}})
			rows = append(rows, ct.Row{Columns: []ct.RowColumn{
				{
					Label: labels["place_name_movement"],
					Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "place_code_" + cu.ToString(trans.Id, ""),
						},
						Type: ct.FieldTypeSelect, Value: cu.IM{
							"name":     "place_code",
							"options":  placeOpt(),
							"is_null":  true,
							"value":    trans.PlaceCode,
							"disabled": !newInput,
						},
					}},
				{
					Label: labels["place_name_target"],
					Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "target_place_code_" + cu.ToString(trans.Id, ""),
						},
						Type: ct.FieldTypeSelect, Value: cu.IM{
							"name":     "target_place_code",
							"options":  placeOpt(),
							"is_null":  true,
							"value":    targetPlaceCode(),
							"disabled": len(movements) == 0,
						},
					}},
			}, Full: true, BorderBottom: true})
		},
		md.TransTypeInventory: func() {
			rows[1].Columns[0].Value.Value["disabled"] = true
			rows[0].Columns = append(rows[0].Columns, ct.RowColumn{
				Label: labels["place_name_movement"],
				Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "place_code_" + cu.ToString(trans.Id, ""),
					},
					Type: ct.FieldTypeSelect, Value: cu.IM{
						"name":    "place_code",
						"options": placeOpt(),
						"is_null": true,
						"value":   trans.PlaceCode,
					},
				}})
			rows[1].Columns = append(rows[1].Columns, ct.RowColumn{
				Label: labels["trans_closed"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "closed_" + cu.ToString(trans.Id, ""),
					},
					Type: ct.FieldTypeBool, Value: cu.IM{
						"name":  "closed",
						"value": cu.ToBoolean(trans.TransMeta.Closed, false),
					},
				}})
		},
		md.TransTypeWaybill: func() {
			rows[1].Columns[2] = ct.RowColumn{
				Label: labels["trans_closed"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "closed_" + cu.ToString(trans.Id, ""),
					},
					Type: ct.FieldTypeBool, Value: cu.IM{
						"name":  "closed",
						"value": cu.ToBoolean(trans.TransMeta.Closed, false),
					},
				}}
			rows = append(rows, ct.Row{Columns: []ct.RowColumn{
				{
					Label: labels["trans_trans_code"], Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "trans_code_" + cu.ToString(trans.Id, ""),
						},
						Type: ct.FieldTypeSelector, Value: cu.IM{
							"name":  "transitem_code",
							"title": labels["view_transitem"],
							"value": ct.SelectOption{
								Value: trans.TransCode,
								Text:  trans.TransCode,
							},
							"fields":  transitemSelectorFields,
							"rows":    transitemSelectorRows,
							"link":    true,
							"is_null": true,
						},
					},
				},
				{
					Label: labels["customer_name"], Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "customer_code_" + cu.ToString(trans.Id, ""),
						},
						Type: ct.FieldTypeSelector, Value: cu.IM{
							"name":  "customer_code",
							"title": labels["view_customer"],
							"value": ct.SelectOption{
								Value: trans.CustomerCode,
								Text:  customerName,
							},
							"fields":  customerSelectorFields,
							"rows":    customerSelectorRows,
							"link":    true,
							"is_null": true,
						},
					},
				},
				{
					Label: labels["employee_code"], Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "employee_code_" + cu.ToString(trans.Id, ""),
						},
						Type: ct.FieldTypeSelector, Value: cu.IM{
							"name":  "employee_code",
							"title": labels["employee_view"],
							"value": ct.SelectOption{
								Value: trans.EmployeeCode,
								Text:  trans.EmployeeCode,
							},
							"fields":  employeeSelectorFields,
							"rows":    employeeSelectorRows,
							"link":    true,
							"is_null": true,
						},
					},
				},
			}, Full: true, BorderBottom: true})
		},
		md.TransTypeProduction: func() {
			rows[1].Columns[0].Value.Value["disabled"] = true
			rows[0].Columns = append(rows[0].Columns, ct.RowColumn{
				Label: labels["place_name_movement"],
				Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "place_code_" + cu.ToString(trans.Id, ""),
					},
					Type: ct.FieldTypeSelect, Value: cu.IM{
						"name":    "place_code",
						"options": placeOpt(),
						"is_null": true,
						"value":   trans.PlaceCode,
					},
				}})
			rows[1].Columns = append(rows[1].Columns,
				ct.RowColumn{
					Label: transTypeLabel(labels, trans.TransType.String(), "trans_due_time"),
					Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "due_time_" + cu.ToString(trans.Id, ""),
						},
						Type: ct.FieldTypeDate, Value: cu.IM{
							"name":    "due_time",
							"is_null": false,
							"value":   trans.TransMeta.DueTime,
						},
					}})
			productCode, batchNo, qty := productionHeadRow()
			rows = append(rows, ct.Row{Columns: []ct.RowColumn{
				{
					Label: labels["product_code"], Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "product_code_" + cu.ToString(trans.Id, ""),
						},
						Type: ct.FieldTypeSelector, Value: cu.IM{
							"name":  "movement_product_code",
							"title": labels["product_view"],
							"value": ct.SelectOption{
								Value: productCode,
								Text:  productCode,
							},
							"fields":  productSelectorFields,
							"rows":    productSelectorRows,
							"link":    true,
							"is_null": false,
						},
					},
				},
				{Label: labels["movement_batchnumber"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "movement_notes_" + cu.ToString(trans.Id, ""),
					},
					Type: ct.FieldTypeString, Value: cu.IM{
						"name":  "movement_notes",
						"value": batchNo,
					},
				}},
				{
					Label: labels["movement_qty"],
					Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "movement_qty_" + cu.ToString(trans.Id, ""),
						},
						Type: ct.FieldTypeNumber, Value: cu.IM{
							"name":  "movement_qty",
							"value": qty,
						},
					}},
			},
				Full: true, BorderBottom: true})
		},
		md.TransTypeFormula: func() {
			rows[1].Columns[0].Value.Value["disabled"] = true
			rows[1].Columns[2] = ct.RowColumn{
				Label: labels["trans_closed"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "closed_" + cu.ToString(trans.Id, ""),
					},
					Type: ct.FieldTypeBool, Value: cu.IM{
						"name":  "closed",
						"value": cu.ToBoolean(trans.TransMeta.Closed, false),
					},
				}}
			productCode, batchNo, qty := formulaHeadRow()
			rows = append(rows, ct.Row{Columns: []ct.RowColumn{
				{
					Label: labels["product_code"], Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "product_code_" + cu.ToString(trans.Id, ""),
						},
						Type: ct.FieldTypeSelector, Value: cu.IM{
							"name":  "movement_product_code",
							"title": labels["product_view"],
							"value": ct.SelectOption{
								Value: productCode,
								Text:  productCode,
							},
							"fields":  productSelectorFields,
							"rows":    productSelectorRows,
							"link":    true,
							"is_null": false,
						},
					},
				},
				{Label: labels["movement_batchnumber"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "movement_notes_" + cu.ToString(trans.Id, ""),
					},
					Type: ct.FieldTypeString, Value: cu.IM{
						"name":  "movement_notes",
						"value": batchNo,
					},
				}},
				{
					Label: labels["movement_qty"],
					Value: ct.Field{
						BaseComponent: ct.BaseComponent{
							Name: "movement_qty_" + cu.ToString(trans.Id, ""),
						},
						Type: ct.FieldTypeNumber, Value: cu.IM{
							"name":  "movement_qty",
							"value": qty,
						},
					}},
			},
				Full: true, BorderBottom: true})
		},
	}

	rows = transMainHeadRow(trans, labels, data)
	if fn, ok := rowMap[trans.TransType]; ok {
		fn()
	}
	rows = append(rows, transMainFooterRow(trans, labels, data)...)
	return rows
}

func transRowItemsTotal(_ string, labels cu.SM, data cu.IM) (rows []ct.Row) {
	items := cu.ToIMA(data["items"], []cu.IM{})
	itemTotal := func() (netAmount, vatAmount, amount float64) {
		for _, item := range items {
			netAmount += cu.ToFloat(cu.ToIM(item["item_meta"], cu.IM{})["net_amount"], 0)
			vatAmount += cu.ToFloat(cu.ToIM(item["item_meta"], cu.IM{})["vat_amount"], 0)
			amount += cu.ToFloat(cu.ToIM(item["item_meta"], cu.IM{})["amount"], 0)
		}
		return netAmount, vatAmount, amount
	}
	netAmount, vatAmount, amount := itemTotal()
	return []ct.Row{
		{Columns: []ct.RowColumn{
			{
				Label: labels["item_net_amount"],
				Value: ct.Field{
					Type: ct.FieldTypeNumber, Value: cu.IM{
						"name":     "net_amount",
						"value":    netAmount,
						"disabled": true,
						"style": cu.SM{
							"opacity": "1",
						},
					},
				}},
			{
				Label: labels["item_vat_amount"],
				Value: ct.Field{
					Type: ct.FieldTypeNumber, Value: cu.IM{
						"name":     "vat_amount",
						"value":    vatAmount,
						"disabled": true,
						"style": cu.SM{
							"opacity": "1",
						},
					},
				}},
			{
				Label: labels["item_amount"],
				Value: ct.Field{
					Type: ct.FieldTypeNumber, Value: cu.IM{
						"name":     "amount",
						"value":    amount,
						"disabled": true,
						"style": cu.SM{
							"opacity": "1",
						},
					},
				}},
		}, Full: false},
	}
}

func transRowPaymentTotal(_ string, labels cu.SM, data cu.IM) (rows []ct.Row) {
	payments := cu.ToIMA(data["payments"], []cu.IM{})
	itemTotal := func() (expense, income, balance float64) {
		for _, payment := range payments {
			amount := cu.ToFloat(cu.ToIM(payment["payment_meta"], cu.IM{})["amount"], 0)
			if amount > 0 {
				income += amount
			} else {
				expense += amount
			}
			balance += amount
		}
		return expense, income, balance
	}
	expense, income, balance := itemTotal()
	return []ct.Row{
		{Columns: []ct.RowColumn{
			{
				Label: labels["payment_expense"],
				Value: ct.Field{
					Type: ct.FieldTypeNumber, Value: cu.IM{
						"name":     "expense",
						"value":    expense,
						"disabled": true,
						"style": cu.SM{
							"opacity": "1",
						},
					},
				}},
			{
				Label: labels["payment_income"],
				Value: ct.Field{
					Type: ct.FieldTypeNumber, Value: cu.IM{
						"name":     "income",
						"value":    income,
						"disabled": true,
						"style": cu.SM{
							"opacity": "1",
						},
					},
				}},
			{
				Label: labels["payment_balance"],
				Value: ct.Field{
					Type: ct.FieldTypeNumber, Value: cu.IM{
						"name":     "balance",
						"value":    balance,
						"disabled": true,
						"style": cu.SM{
							"opacity": "1",
						},
					},
				}},
		}, Full: false},
	}
}

func (e *TransEditor) Row(view string, labels cu.SM, data cu.IM) (rows []ct.Row) {
	if !slices.Contains([]string{"trans", "maps", "items", "payments"}, view) {
		return []ct.Row{}
	}

	var trans md.Trans = md.Trans{}
	ut.ConvertToType(data["trans"], &trans)

	configMap := cu.ToIMA(data["config_map"], []cu.IM{})
	selectedField := cu.ToString(data["map_field"], "")

	mapFieldOpt := func() (opt []ct.SelectOption) {
		opt = []ct.SelectOption{}
		for _, field := range configMap {
			filter := ut.ToStringArray(field["filter"])
			if slices.Contains(filter, "FILTER_TRANS") || len(filter) == 0 {
				if _, ok := trans.TransMap[cu.ToString(field["field_name"], "")]; !ok {
					opt = append(opt, ct.SelectOption{
						Value: cu.ToString(field["field_name"], ""), Text: cu.ToString(field["description"], ""),
					})
				}
			}
		}
		return opt
	}

	if view == "maps" {
		return []ct.Row{
			{Columns: []ct.RowColumn{
				{Label: labels["map_fields"], Value: ct.Field{
					BaseComponent: ct.BaseComponent{
						Name: "map_field_" + cu.ToString(trans.Id, ""),
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

	if view == "items" {
		return transRowItemsTotal(view, labels, data)
	}
	if view == "payments" {
		return transRowPaymentTotal(view, labels, data)
	}

	if TransIsItem(trans.TransType.String()) {
		return transMainItemRow(trans, labels, data)
	}

	if TransIsPayment(trans.TransType.String()) {
		return transMainPaymentRow(trans, labels, data)
	}

	//if TransIsMovement(trans.TransType.String()) {
	return transMainMovementRow(trans, labels, data)
	//}

	//return rows
}

func transMovementRows(data cu.IM) (rows []cu.IM) {
	var trans cu.IM = cu.ToIM(data["trans"], cu.IM{})
	transType := cu.ToString(trans["trans_type"], "")
	direction := cu.ToString(trans["direction"], "")
	isTransfer := (transType == md.TransTypeDelivery.String() && direction == md.DirectionTransfer.String())

	validMap := map[string]func(movement cu.IM) bool{
		md.TransTypeDelivery.String(): func(movement cu.IM) bool {
			return (isTransfer && cu.ToString(movement["movement_code"], "") != "") || !isTransfer || cu.ToInteger(movement["id"], 0) < 0
		},
		md.TransTypeInventory.String(): func(movement cu.IM) bool {
			return true
		},
		md.TransTypeProduction.String(): func(movement cu.IM) bool {
			movementMeta := cu.ToIM(movement["movement_meta"], cu.IM{})
			return !cu.ToBoolean(movementMeta["shared"], false) || cu.ToInteger(movement["id"], 0) < 0
		},
		md.TransTypeWaybill.String(): func(movement cu.IM) bool {
			return true
		},
		md.TransTypeFormula.String(): func(movement cu.IM) bool {
			return cu.ToString(movement["movement_type"], "") != md.MovementTypeHead.String() || cu.ToInteger(movement["id"], 0) < 0
		},
	}

	viewMap := cu.SM{
		md.TransTypeDelivery.String():   "movement_inventory",
		md.TransTypeInventory.String():  "movement_inventory",
		md.TransTypeProduction.String(): "movement_inventory",
		md.TransTypeWaybill.String():    "movement_waybill",
		md.TransTypeFormula.String():    "movement_formula",
	}

	rows = []cu.IM{}
	movements := cu.ToIMA(data["movements"], []cu.IM{})
	for idx, movement := range movements {
		if validMap[transType](movement) {
			movementMeta := cu.ToIM(movement["movement_meta"], cu.IM{})
			row := cu.IM{
				"id":               movement["id"],
				"index":            idx,
				"trans_type":       transType,
				"code":             movement["code"],
				"shipping_time":    movement["shipping_time"],
				"movement_type":    movement["movement_type"],
				"product_code":     movement["product_code"],
				"product_name":     movement["product_name"],
				"product_unit":     movement["product_unit"],
				"tool_code":        movement["tool_code"],
				"tool_description": movement["tool_description"],
				"serial_number":    movement["serial_number"],
				"place_code":       movement["place_code"],
				"item_code":        movement["item_code"],
				"movement_code":    movement["movement_code"],
				"qty":              movementMeta["qty"],
				"notes":            movementMeta["notes"],
				"shared":           movementMeta["shared"],
				"tags":             movementMeta["tags"],
				"movement_meta":    movement["movement_meta"],
				"movement_map":     movement["movement_map"],
				"editor":           "trans",
			}
			if row["movement_code"] != "" {
				if idx := slices.IndexFunc(movements, func(refMovement cu.IM) bool {
					return cu.ToString(refMovement["code"], "") == cu.ToString(row["movement_code"], "")
				}); idx > -1 {
					row["ref_index"] = idx
					row["ref_id"] = cu.ToInteger(movements[idx]["id"], 0)
				}
			}
			if view, ok := viewMap[transType]; ok {
				inventories := cu.ToIMA(data[view], []cu.IM{})
				if idx := slices.IndexFunc(inventories, func(inventory cu.IM) bool {
					return cu.ToString(inventory["code"], "") == cu.ToString(movement["code"], "")
				}); idx > -1 {
					row["product_name"] = cu.ToString(movement["product_name"], cu.ToString(inventories[idx]["product_name"], ""))
					row["place_name"] = cu.ToString(movement["place_name"], cu.ToString(inventories[idx]["place_name"], ""))
					row["product_unit"] = cu.ToString(movement["product_unit"], cu.ToString(inventories[idx]["unit"], ""))
					row["serial_number"] = cu.ToString(movement["serial_number"], cu.ToString(inventories[idx]["serial_number"], ""))
					row["tool_description"] = cu.ToString(movement["tool_description"], cu.ToString(inventories[idx]["description"], ""))
				}
			}
			rows = append(rows, row)
		}
	}
	return rows
}

func (e *TransEditor) Table(view string, labels cu.SM, data cu.IM) []ct.Table {
	if !slices.Contains([]string{"maps", "items", "transitem_invoice", "payments", "payment_link", "movements",
		"transitem_shipping", "tool_movement"}, view) {
		return []ct.Table{}
	}

	var trans cu.IM = cu.ToIM(data["trans"], cu.IM{})
	transType := cu.ToString(trans["trans_type"], "")
	direction := cu.ToString(trans["direction"], "")
	newInput := (cu.ToInteger(trans["id"], 0) == 0)
	tblMap := map[string]func() []ct.Table{
		"maps": func() []ct.Table {
			configMap := cu.ToIMA(data["config_map"], []cu.IM{})
			transMap := cu.ToIM(trans["trans_map"], cu.IM{})
			return []ct.Table{
				{
					Fields: []ct.TableField{
						{Name: "description", Label: labels["map_description"], ReadOnly: true},
						{Name: "value", Label: labels["map_value"], FieldType: ct.TableFieldTypeMeta, Required: true},
					},
					Rows:              mapTableRows(transMap, configMap),
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
		"items": func() []ct.Table {
			itemRows := func() []cu.IM {
				rows := []cu.IM{}
				items := cu.ToIMA(data["items"], []cu.IM{})
				for _, item := range items {
					rows = append(rows, cu.IM{
						"id":           item["id"],
						"product_code": item["product_code"],
						"tax_code":     item["tax_code"],
						"description":  cu.ToIM(item["item_meta"], cu.IM{})["description"],
						"unit":         cu.ToIM(item["item_meta"], cu.IM{})["unit"],
						"qty":          cu.ToFloat(cu.ToIM(item["item_meta"], cu.IM{})["qty"], 0),
						"amount":       cu.ToFloat(cu.ToIM(item["item_meta"], cu.IM{})["amount"], 0),
						"item_meta":    item["item_meta"],
					})
				}
				return rows
			}
			return []ct.Table{
				{
					Fields: []ct.TableField{
						{Name: "description", Label: labels["item_description"]},
						{Name: "unit", Label: labels["item_unit"]},
						{Name: "qty", Label: labels["item_qty"], FieldType: ct.TableFieldTypeNumber},
						{Name: "amount", Label: labels["item_amount"], FieldType: ct.TableFieldTypeNumber},
					},
					Rows:              itemRows(),
					Pagination:        ct.PaginationTypeTop,
					PageSize:          5,
					HidePaginatonSize: true,
					RowSelected:       true,
					TableFilter:       true,
					FilterPlaceholder: labels["placeholder_filter"],
					AddItem:           !newInput,
					LabelAdd:          labels["item_new"],
				},
			}
		},
		"transitem_invoice": func() []ct.Table {
			itemRows := func() []cu.IM {
				rows := []cu.IM{}
				items := cu.ToIMA(data["transitem_invoice"], []cu.IM{})
				for _, item := range items {
					rows = append(rows, cu.IM{
						"trans_code":    item["trans_code"],
						"trans_date":    item["trans_date"],
						"description":   cu.ToIM(item["item_meta"], cu.IM{})["description"],
						"unit":          cu.ToIM(item["item_meta"], cu.IM{})["unit"],
						"currency_code": item["currency_code"],
						"qty":           cu.ToFloat(cu.ToIM(item["item_meta"], cu.IM{})["qty"], 0),
						"amount":        cu.ToFloat(cu.ToIM(item["item_meta"], cu.IM{})["amount"], 0),
						"deposit":       cu.ToBoolean(cu.ToIM(item["item_meta"], cu.IM{})["deposit"], false),
					})
				}
				return rows
			}
			return []ct.Table{
				{
					Fields: []ct.TableField{
						{Name: "trans_code", Label: labels["trans_code"]},
						//{Name: "trans_date", Label: labels["trans_date"]},
						{Name: "description", Label: labels["item_description"]},
						{Name: "unit", Label: labels["item_unit"]},
						{Name: "currency_code", Label: labels["currency_code"]},
						{Name: "qty", Label: labels["item_qty"], FieldType: ct.TableFieldTypeNumber},
						{Name: "amount", Label: labels["item_amount"], FieldType: ct.TableFieldTypeNumber},
						{Name: "deposit", Label: labels["item_deposit"], FieldType: ct.TableFieldTypeBool, TextAlign: ct.TextAlignCenter},
					},
					Rows:              itemRows(),
					Pagination:        ct.PaginationTypeTop,
					PageSize:          5,
					HidePaginatonSize: true,
					RowSelected:       true,
					TableFilter:       true,
					FilterPlaceholder: labels["placeholder_filter"],
				},
			}
		},
		"payments": func() []ct.Table {
			itemRows := func() []cu.IM {
				rows := []cu.IM{}
				payments := cu.ToIMA(data["payments"], []cu.IM{})
				for _, payment := range payments {
					paymentMeta := cu.ToIM(payment["payment_meta"], cu.IM{})
					rows = append(rows, cu.IM{
						"id":               payment["id"],
						"code":             payment["code"],
						"paid_date":        payment["paid_date"],
						"amount":           paymentMeta["amount"],
						"notes":            paymentMeta["notes"],
						"payment_link_add": labels["payment_link_add"],
						"editor":           "trans",
						"payment_meta":     paymentMeta,
					})
				}
				return rows
			}
			return []ct.Table{
				{
					Fields: []ct.TableField{
						{Name: "code", Label: labels["payment_code"], ReadOnly: true},
						{Name: "paid_date", Label: labels["payment_paid_date"], FieldType: ct.TableFieldTypeDate},
						{Name: "amount", Label: labels["payment_amount"], FieldType: ct.TableFieldTypeNumber},
						{Name: "notes", Label: labels["payment_notes"]},
						{Name: "payment_link_add", Label: labels["payment_link_view_bank"], FieldType: ct.TableFieldTypeLink, ReadOnly: true},
					},
					Rows:              itemRows(),
					Pagination:        ct.PaginationTypeTop,
					PageSize:          10,
					HidePaginatonSize: true,
					RowSelected:       true,
					TableFilter:       true,
					FilterPlaceholder: labels["placeholder_filter"],
					AddItem:           !newInput,
					LabelAdd:          labels["payment_new"],
				},
			}
		},
		"movements": func() []ct.Table {
			isDelivery := (transType == md.TransTypeDelivery.String() && direction != md.DirectionTransfer.String())
			fields := []ct.TableField{
				{Name: "code", Label: labels["movement_code"], ReadOnly: true},
				{Name: "product_code", Label: labels["product_code"], FieldType: ct.TableFieldTypeLink},
				{Name: "product_name", Label: labels["product_name"]},
				{Name: "product_unit", Label: labels["product_unit"]},
				{Name: "notes", Label: labels["movement_batchnumber"]},
				{Name: "qty", Label: labels["movement_qty"], FieldType: ct.TableFieldTypeNumber},
			}
			fieldMap := map[string][]ct.TableField{
				md.TransTypeDelivery.String() + "_" + md.DirectionTransfer.String(): {
					{Name: "code", Label: labels["movement_code"], ReadOnly: true},
					{Name: "movement_code", Label: labels["movement_movement_code"]},
					{Name: "product_name", Label: labels["product_name"]},
					{Name: "product_unit", Label: labels["product_unit"]},
					{Name: "notes", Label: labels["movement_batchnumber"]},
					{Name: "qty", Label: labels["movement_qty"], FieldType: ct.TableFieldTypeNumber},
				},
				md.TransTypeFormula.String() + "_" + md.DirectionTransfer.String(): {
					{Name: "code", Label: labels["movement_code"], ReadOnly: true},
					{Name: "product_name", Label: labels["product_name"]},
					{Name: "product_unit", Label: labels["product_unit"]},
					{Name: "shared", Label: labels["movement_shared"], FieldType: ct.TableFieldTypeBool},
					{Name: "qty", Label: labels["movement_qty"], FieldType: ct.TableFieldTypeNumber},
				},
				md.TransTypeWaybill.String() + "_" + md.DirectionIn.String(): {
					{Name: "code", Label: labels["movement_code"], ReadOnly: true},
					{Name: "shipping_time", Label: labels["movement_shipping_time"], FieldType: ct.TableFieldTypeDateTime},
					{Name: "serial_number", Label: labels["tool_serial_number"]},
					{Name: "tool_description", Label: labels["tool_description"]},
				},
				md.TransTypeWaybill.String() + "_" + md.DirectionOut.String(): {
					{Name: "code", Label: labels["movement_code"], ReadOnly: true},
					{Name: "shipping_time", Label: labels["movement_shipping_time"], FieldType: ct.TableFieldTypeDateTime},
					{Name: "serial_number", Label: labels["tool_serial_number"]},
					{Name: "tool_description", Label: labels["tool_description"]},
				},
			}
			if field, ok := fieldMap[transType+"_"+direction]; ok {
				fields = field
			}
			return []ct.Table{
				{
					Fields:            fields,
					Rows:              transMovementRows(data),
					Pagination:        ct.PaginationTypeTop,
					PageSize:          10,
					HidePaginatonSize: true,
					RowSelected:       !isDelivery,
					TableFilter:       true,
					FilterPlaceholder: labels["placeholder_filter"],
					AddItem:           !isDelivery,
					LabelAdd:          labels["item_new"],
				},
			}
		},
		"payment_link": func() []ct.Table {
			isInvoice := slices.Contains([]string{md.TransTypeInvoice.String(), md.TransTypeReceipt.String()}, cu.ToString(trans["trans_type"], ""))
			itemRows := func() []cu.IM {
				rows := []cu.IM{}
				items := cu.ToIMA(data["payment_link"], []cu.IM{})
				for _, item := range items {
					row := cu.IM{
						"id":            item["id"],
						"code":          item["code"],
						"link_code_1":   item["link_code_1"],
						"link_code_2":   item["link_code_2"],
						"trans_code":    item["link_code_2"],
						"currency_code": cu.ToString(item["currency_code"], "") + "/" + cu.ToString(item["invoice_curr"], ""),
						"amount":        cu.ToFloat(cu.ToIM(item["link_meta"], cu.IM{})["amount"], 0),
						"rate":          cu.ToFloat(cu.ToIM(item["link_meta"], cu.IM{})["rate"], 0),
						"link_meta":     item["link_meta"],
						"place_name":    item["place_name"],
						"paid_date":     item["paid_date"],
						"editor":        "trans",
					}
					if isInvoice {
						row["trans_code"] = item["pt_code"]
					}
					rows = append(rows, row)
				}
				return rows
			}
			fields := []ct.TableField{
				{Name: "code", Label: labels["link_code"]},
				{Name: "link_code_1", Label: labels["payment_code"]},
				{Name: "currency_code", Label: labels["currency_code"]},
				{Name: "amount", Label: labels["payment_amount"], FieldType: ct.TableFieldTypeNumber},
				{Name: "rate", Label: labels["payment_rate"], FieldType: ct.TableFieldTypeNumber},
				{Name: "trans_code", Label: labels["trans_code"], FieldType: ct.TableFieldTypeLink},
			}
			if isInvoice {
				fields = []ct.TableField{
					//{Name: "code", Label: labels["link_code"]},
					{Name: "trans_code", Label: labels["trans_code"], FieldType: ct.TableFieldTypeLink},
					{Name: "place_name", Label: labels["place_name_payment"]},
					{Name: "link_code_1", Label: labels["payment_code"]},
					{Name: "paid_date", Label: labels["payment_paid_date"]},
					{Name: "currency_code", Label: labels["currency_code"]},
					{Name: "amount", Label: labels["payment_amount"], FieldType: ct.TableFieldTypeNumber},
					{Name: "rate", Label: labels["payment_rate"], FieldType: ct.TableFieldTypeNumber},
				}
			}
			return []ct.Table{
				{
					Fields:            fields,
					Rows:              itemRows(),
					Pagination:        ct.PaginationTypeTop,
					PageSize:          5,
					HidePaginatonSize: true,
					RowSelected:       !isInvoice,
					TableFilter:       true,
					FilterPlaceholder: labels["placeholder_filter"],
				},
			}
		},
		"transitem_shipping": func() []ct.Table {
			items := cu.ToIMA(data["transitem_shipping"], []cu.IM{})
			return []ct.Table{
				{
					Fields: []ct.TableField{
						//{Name: "code", Label: labels["item_code"]},
						{Name: "description", Label: labels["shipping_item_product"]},
						{Name: "qty", Label: labels["shipping_qty"], FieldType: ct.TableFieldTypeNumber},
						{Name: "trans_code", Label: labels["trans_code"], FieldType: ct.TableFieldTypeLink},
						{Name: "product_name", Label: labels["shipping_movement_product"]},
						//{Name: "shipping_date", Label: labels["shipping_date"], FieldType: ct.TableFieldTypeDate},
						{Name: "shipping_qty", Label: labels["shipping_sqty"], FieldType: ct.TableFieldTypeNumber},
					},
					Rows:              items,
					Pagination:        ct.PaginationTypeTop,
					PageSize:          5,
					HidePaginatonSize: true,
					RowSelected:       false,
					TableFilter:       true,
					FilterPlaceholder: labels["placeholder_filter"],
					AddItem:           true,
					LabelAdd:          labels["transitem_shipping_view"],
					AddIcon:           ct.IconTruck,
				},
			}
		},
		"tool_movement": func() []ct.Table {
			items := cu.ToIMA(data["tool_movement"], []cu.IM{})
			return []ct.Table{
				{
					Fields: []ct.TableField{
						{Name: "trans_code", Label: labels["trans_code"], ReadOnly: true, FieldType: ct.TableFieldTypeLink},
						{Name: "direction", Label: labels["movement_direction"], ReadOnly: true},
						{Name: "shipping_time", Label: labels["movement_shipping_time"], FieldType: ct.TableFieldTypeDateTime},
						{Name: "serial_number", Label: labels["tool_serial_number"]},
						{Name: "description", Label: labels["tool_description"]},
					},
					Rows:              items,
					Pagination:        ct.PaginationTypeTop,
					PageSize:          10,
					HidePaginatonSize: true,
					RowSelected:       false,
					TableFilter:       true,
					FilterPlaceholder: labels["placeholder_filter"],
					AddItem:           true,
					LabelAdd:          labels["trans_waybill_new"],
					AddIcon:           ct.IconBriefcase,
				},
			}
		},
	}
	return tblMap[view]()
}

func (e *TransEditor) Form(formKey string, labels cu.SM, data cu.IM) (form ct.Form) {
	var productSelectorFields []ct.TableField = []ct.TableField{
		{Name: "code", Label: labels["product_code"]},
		{Name: "product_name", Label: labels["product_name"]},
		{Name: "product_type", Label: labels["product_type"]},
		{Name: "tag_lst", Label: labels["product_tags"]},
	}
	var toolSelectorFields []ct.TableField = []ct.TableField{
		{Name: "code", Label: labels["tool_code"]},
		{Name: "description", Label: labels["tool_description"]},
		{Name: "product_code", Label: labels["product_code"]},
		{Name: "serial_number", Label: labels["tool_serial_number"]},
		{Name: "tag_lst", Label: labels["tool_tags"]},
	}
	formData := cu.ToIM(data, cu.IM{})
	footerRows := func(disabled bool) []ct.Row {
		rows := []ct.Row{
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
							//"auto_focus":   true,
							"selected": true,
							"disabled": disabled,
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
							"disabled":     cu.ToInteger(formData["id"], 0) < 0,
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
		return rows
	}
	frmMap := map[string]func() ct.Form{
		"items": func() ct.Form {
			var item md.Item = md.Item{}
			ut.ConvertToType(formData, &item)
			taxCodes := cu.ToIMA(data["tax_codes"], []cu.IM{})
			productSelectorRows := cu.ToIMA(data["product_selector"], []cu.IM{})
			taxCodeOpt := func() (opt []ct.SelectOption) {
				opt = []ct.SelectOption{}
				for _, taxCode := range taxCodes {
					opt = append(opt, ct.SelectOption{
						Value: cu.ToString(taxCode["code"], ""), Text: cu.ToString(taxCode["description"], ""),
					})
				}
				return opt
			}
			return ct.Form{
				Title: labels["item_view"],
				Icon:  ct.IconListOl,
				BodyRows: []ct.Row{
					{Columns: []ct.RowColumn{
						{
							Label: labels["product_code"], Value: ct.Field{
								BaseComponent: ct.BaseComponent{
									Name: "product_code_" + cu.ToString(item.Id, ""),
								},
								Type: ct.FieldTypeSelector, Value: cu.IM{
									"name":  "product_code",
									"title": labels["view_product"],
									"value": ct.SelectOption{
										Value: item.ProductCode,
										Text:  item.ProductCode,
									},
									"fields":  productSelectorFields,
									"rows":    productSelectorRows,
									"link":    true,
									"is_null": false,
								},
								FormTrigger: true,
							},
						},
					}, Full: true, BorderBottom: true, FieldCol: true},
					{Columns: []ct.RowColumn{
						{Label: labels["item_unit"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "unit",
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":  "unit",
								"value": item.ItemMeta.Unit,
							},
							FormTrigger: true,
						}},
						{Label: labels["item_own_stock"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "own_stock_" + cu.ToString(item.Id, ""),
							},
							Type: ct.FieldTypeNumber, Value: cu.IM{
								"name":  "own_stock",
								"value": item.ItemMeta.OwnStock,
							},
							FormTrigger: true,
						}},
						{Label: labels["item_deposit"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "deposit_" + cu.ToString(item.Id, ""),
							},
							Type: ct.FieldTypeBool, Value: cu.IM{
								"name":  "deposit",
								"value": cu.ToBoolean(item.ItemMeta.Deposit, false),
							},
							FormTrigger: true,
						}},
					}, Full: true, BorderBottom: true},
					{Columns: []ct.RowColumn{
						{Label: labels["item_qty"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "qty_" + cu.ToString(item.Id, ""),
							},
							Type: ct.FieldTypeNumber, Value: cu.IM{
								"name":  "qty",
								"value": item.ItemMeta.Qty,
							},
							FormTrigger: true,
						}},
						{Label: labels["item_discount"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "discount_" + cu.ToString(item.Id, ""),
							},
							Type: ct.FieldTypeInteger, Value: cu.IM{
								"name":      "discount",
								"set_max":   true,
								"max_value": 100,
								"set_min":   true,
								"min_value": 0,
								"value":     cu.ToInteger(item.ItemMeta.Discount, 0),
							},
							FormTrigger: true,
						}},
						{Label: labels["item_fx_price"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "fx_price_" + cu.ToString(item.Id, ""),
							},
							Type: ct.FieldTypeNumber, Value: cu.IM{
								"name":  "fx_price",
								"value": item.ItemMeta.FxPrice,
							},
							FormTrigger: true,
						}},
					}, Full: true, BorderBottom: true},
					{Columns: []ct.RowColumn{
						{Label: labels["item_net_amount"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "net_amount_" + cu.ToString(item.Id, ""),
							},
							Type: ct.FieldTypeNumber, Value: cu.IM{
								"name":  "net_amount",
								"value": item.ItemMeta.NetAmount,
							},
							FormTrigger: true,
						}},
						{Label: labels["tax_code"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "tax_code_" + cu.ToString(item.Id, ""),
							},
							Type: ct.FieldTypeSelect, Value: cu.IM{
								"name":    "tax_code",
								"options": taxCodeOpt(),
								"is_null": false,
								"value":   item.TaxCode,
							},
							FormTrigger: true,
						}},
						{Label: labels["item_amount"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "amount_" + cu.ToString(item.Id, ""),
							},
							Type: ct.FieldTypeNumber, Value: cu.IM{
								"name":  "amount",
								"value": item.ItemMeta.Amount,
							},
							FormTrigger: true,
						}},
					}, Full: true, BorderBottom: true},
					{Columns: []ct.RowColumn{
						{Label: labels["item_description"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "description",
							},
							Type: ct.FieldTypeText, Value: cu.IM{
								"name":  "description",
								"value": item.ItemMeta.Description,
								"rows":  3,
							},
							FormTrigger: true,
						}},
						{
							Label: labels["item_tags"], Value: ct.Field{
								BaseComponent: ct.BaseComponent{
									Name: "tags",
								},
								Type: ct.FieldTypeList, Value: cu.IM{
									"name":                "tags",
									"rows":                ut.ToTagList(item.ItemMeta.Tags),
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
					}, Full: true},
				},
				FooterRows: footerRows(item.ProductCode == ""),
			}
		},
		"payments": func() ct.Form {
			var payment md.Payment = md.Payment{}
			ut.ConvertToType(formData, &payment)
			return ct.Form{
				Title: labels["payment_view"],
				Icon:  ct.IconListOl,
				BodyRows: []ct.Row{
					{Columns: []ct.RowColumn{
						{Label: labels["payment_code"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "code_" + cu.ToString(formData["id"], ""),
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":     "payment_code",
								"value":    payment.Code,
								"disabled": true,
							},
						}},
						{Label: labels["payment_paid_date"],
							Value: ct.Field{
								BaseComponent: ct.BaseComponent{
									Name: "paid_date_" + cu.ToString(payment.Id, ""),
								},
								Type: ct.FieldTypeDate, Value: cu.IM{
									"name":    "paid_date",
									"is_null": false,
									"value":   payment.PaidDate,
								},
							}},
						{Label: labels["payment_amount"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "amount_" + cu.ToString(payment.Id, ""),
							},
							Type: ct.FieldTypeNumber, Value: cu.IM{
								"name":  "amount",
								"value": payment.PaymentMeta.Amount,
							},
						}},
					}, Full: true, BorderBottom: true},
					{Columns: []ct.RowColumn{
						{Label: labels["payment_notes"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "notes_" + cu.ToString(payment.Id, ""),
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":  "notes",
								"value": payment.PaymentMeta.Notes,
							},
						}},
					}, Full: true, BorderBottom: true, FieldCol: true},
				},
				FooterRows: footerRows(false),
			}
		},
		"payment_link": func() ct.Form {
			var link md.Link = md.Link{}
			ut.ConvertToType(formData, &link)
			invoiceSelectorRows := cu.ToIMA(formData["invoice_selector"], []cu.IM{})
			var invoiceSelectorFields []ct.TableField = []ct.TableField{
				{Name: "code", Label: labels["invoice_code"]},
				//{Name: "trans_date", Label: labels["trans_date"]},
				{Name: "customer_name", Label: labels["customer_name"]},
				{Name: "currency_code", Label: labels["currency_code"]},
				{Name: "amount", Label: labels["item_amount"], FieldType: ct.TableFieldTypeNumber},
				//{Name: "trans_type", Label: labels["trans_type"]},
				//{Name: "direction", Label: labels["trans_direction"]},
			}
			return ct.Form{
				Title: labels["payment_link_add"],
				Icon:  ct.IconLink,
				BodyRows: []ct.Row{
					{Columns: []ct.RowColumn{
						{Label: labels["link_code"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "code_" + cu.ToString(formData["id"], ""),
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":     "code",
								"value":    link.Code,
								"disabled": true,
							},
						}},
						{Label: labels["payment_code"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "payment_code_" + cu.ToString(formData["id"], ""),
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":     "payment_code",
								"value":    link.LinkCode1,
								"disabled": true,
							},
						}},
						{
							Label: labels["invoice_code"], Value: ct.Field{
								BaseComponent: ct.BaseComponent{
									Name: "trans_code_" + cu.ToString(formData["id"], ""),
								},
								Type: ct.FieldTypeSelector, Value: cu.IM{
									"name":  "invoice_code",
									"title": labels["view_invoice"],
									"value": ct.SelectOption{
										Value: link.LinkCode2,
										Text:  link.LinkCode2,
									},
									"fields":  invoiceSelectorFields,
									"rows":    invoiceSelectorRows,
									"link":    true,
									"is_null": false,
								},
								FormTrigger: true,
							},
						},
					}, Full: true, BorderBottom: true},
					{Columns: []ct.RowColumn{
						{Label: labels["currency_code"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "currency_code_" + cu.ToString(formData["id"], ""),
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":     "currency_code",
								"value":    cu.ToString(formData["currency_code"], ""),
								"disabled": true,
							},
						}},
						{Label: labels["payment_amount"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "amount_" + cu.ToString(formData["id"], ""),
							},
							Type: ct.FieldTypeNumber, Value: cu.IM{
								"name":  "link_amount",
								"value": cu.ToFloat(link.LinkMeta.Amount, link.LinkMeta.Qty),
							},
							FormTrigger: true,
						}},
						{Label: labels["payment_rate"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "rate_" + cu.ToString(formData["id"], ""),
							},
							Type: ct.FieldTypeNumber, Value: cu.IM{
								"name":  "link_rate",
								"value": link.LinkMeta.Rate,
							},
							FormTrigger: true,
						}},
					}, Full: true, BorderBottom: true},
				},
				FooterRows: footerRows(link.LinkCode2 == ""),
			}
		},
		"movements": func() ct.Form {
			var movement md.Movement = md.Movement{}
			ut.ConvertToType(formData, &movement)
			transType := cu.ToString(formData["trans_type"], "")
			productSelectorRows := cu.ToIMA(data["product_selector"], []cu.IM{})
			toolSelectorRows := cu.ToIMA(data["tool_selector"], []cu.IM{})
			places := cu.ToIMA(formData["places"], []cu.IM{})
			placeOpt := func() (opt []ct.SelectOption) {
				opt = []ct.SelectOption{}
				for _, place := range places {
					if cu.ToString(place["place_type"], "") == md.PlaceTypeWarehouse.String() {
						opt = append(opt, ct.SelectOption{
							Value: cu.ToString(place["code"], ""), Text: cu.ToString(place["place_name"], ""),
						})
					}
				}
				return opt
			}
			rowMap := map[string][]ct.Row{
				md.TransTypeDelivery.String(): {
					{Columns: []ct.RowColumn{
						{Label: labels["movement_code"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "code_" + cu.ToString(movement.Id, ""),
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":     "movement_code",
								"value":    movement.Code,
								"disabled": true,
							},
						}},
						{Label: labels["movement_movement_code"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "movement_code_" + cu.ToString(movement.Id, ""),
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":     "movement_reference_code",
								"value":    movement.MovementCode,
								"disabled": true,
							},
						}},
						{
							Label: labels["place_name_target"],
							Value: ct.Field{
								BaseComponent: ct.BaseComponent{
									Name: "place_code_" + cu.ToString(movement.Id, ""),
								},
								Type: ct.FieldTypeSelect, Value: cu.IM{
									"name":    "place_code",
									"options": placeOpt(),
									"is_null": true,
									"value":   movement.PlaceCode,
								},
								FormTrigger: true,
							}},
					}, Full: true, BorderBottom: true},
					{Columns: []ct.RowColumn{
						{
							Label: labels["product_code"], Value: ct.Field{
								BaseComponent: ct.BaseComponent{
									Name: "product_code_" + cu.ToString(movement.Id, ""),
								},
								Type: ct.FieldTypeSelector, Value: cu.IM{
									"name":  "product_code",
									"title": labels["view_product"],
									"value": ct.SelectOption{
										Value: movement.ProductCode,
										Text:  movement.ProductCode,
									},
									"fields":  productSelectorFields,
									"rows":    productSelectorRows,
									"link":    true,
									"is_null": true,
								},
								FormTrigger: true,
							},
						},
						{Label: labels["product_name"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "product_name_" + cu.ToString(movement.Id, ""),
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":     "product_name",
								"value":    cu.ToString(formData["product_name"], ""),
								"disabled": true,
							},
						}},
					}, Full: true, BorderBottom: true},
					{Columns: []ct.RowColumn{
						{Label: labels["movement_shipping_date"],
							Value: ct.Field{
								BaseComponent: ct.BaseComponent{
									Name: "shipping_time_" + cu.ToString(movement.Id, ""),
								},
								Type: ct.FieldTypeDate, Value: cu.IM{
									"name":    "shipping_time",
									"is_null": false,
									"value":   movement.ShippingTime,
								},
							}},
						{Label: labels["movement_batchnumber"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "notes_" + cu.ToString(movement.Id, ""),
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":  "notes",
								"value": movement.MovementMeta.Notes,
							},
						}},
						{Label: labels["movement_qty"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "qty_" + cu.ToString(movement.Id, ""),
							},
							Type: ct.FieldTypeNumber, Value: cu.IM{
								"name":  "qty",
								"value": movement.MovementMeta.Qty,
							},
						}},
					}, Full: true, BorderBottom: false},
				},
				md.TransTypeInventory.String(): {
					{Columns: []ct.RowColumn{
						{Label: labels["movement_code"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "code_" + cu.ToString(movement.Id, ""),
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":     "movement_code",
								"value":    movement.Code,
								"disabled": true,
							},
						}},
						{
							Label: labels["product_code"], Value: ct.Field{
								BaseComponent: ct.BaseComponent{
									Name: "product_code_" + cu.ToString(movement.Id, ""),
								},
								Type: ct.FieldTypeSelector, Value: cu.IM{
									"name":  "product_code",
									"title": labels["view_product"],
									"value": ct.SelectOption{
										Value: movement.ProductCode,
										Text:  movement.ProductCode,
									},
									"fields":  productSelectorFields,
									"rows":    productSelectorRows,
									"link":    true,
									"is_null": false,
								},
								FormTrigger: true,
							},
						},
						{Label: labels["product_name"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "product_name_" + cu.ToString(movement.Id, ""),
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":     "product_name",
								"value":    cu.ToString(formData["product_name"], ""),
								"disabled": true,
							},
						}},
					}, Full: true, BorderBottom: true},
					{Columns: []ct.RowColumn{
						{Label: labels["movement_shipping_date"],
							Value: ct.Field{
								BaseComponent: ct.BaseComponent{
									Name: "shipping_time_" + cu.ToString(movement.Id, ""),
								},
								Type: ct.FieldTypeDate, Value: cu.IM{
									"name":     "shipping_time",
									"is_null":  false,
									"value":    movement.ShippingTime,
									"disabled": true,
								},
							}},
						{Label: labels["movement_batchnumber"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "notes_" + cu.ToString(movement.Id, ""),
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":  "notes",
								"value": movement.MovementMeta.Notes,
							},
						}},
						{Label: labels["movement_qty"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "qty_" + cu.ToString(movement.Id, ""),
							},
							Type: ct.FieldTypeNumber, Value: cu.IM{
								"name":  "qty",
								"value": movement.MovementMeta.Qty,
							},
						}},
					}, Full: true, BorderBottom: false},
				},
				md.TransTypeProduction.String(): {
					{Columns: []ct.RowColumn{
						{Label: labels["movement_code"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "code_" + cu.ToString(movement.Id, ""),
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":     "movement_code",
								"value":    movement.Code,
								"disabled": true,
							},
						}},
						{
							Label: labels["product_code"], Value: ct.Field{
								BaseComponent: ct.BaseComponent{
									Name: "product_code_" + cu.ToString(movement.Id, ""),
								},
								Type: ct.FieldTypeSelector, Value: cu.IM{
									"name":  "product_code",
									"title": labels["view_product"],
									"value": ct.SelectOption{
										Value: movement.ProductCode,
										Text:  movement.ProductCode,
									},
									"fields":  productSelectorFields,
									"rows":    productSelectorRows,
									"link":    true,
									"is_null": false,
								},
								FormTrigger: true,
							},
						},
						{Label: labels["product_name"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "product_name_" + cu.ToString(movement.Id, ""),
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":     "product_name",
								"value":    cu.ToString(formData["product_name"], ""),
								"disabled": true,
							},
						}},
					}, Full: true, BorderBottom: true},
					{Columns: []ct.RowColumn{
						{Label: labels["movement_shipping_time"],
							Value: ct.Field{
								BaseComponent: ct.BaseComponent{
									Name: "shipping_time_" + cu.ToString(movement.Id, ""),
								},
								Type: ct.FieldTypeDateTime, Value: cu.IM{
									"name":    "shipping_time",
									"is_null": false,
									"value":   movement.ShippingTime,
								},
							}},
						{
							Label: labels["place_name_movement"],
							Value: ct.Field{
								BaseComponent: ct.BaseComponent{
									Name: "place_code_" + cu.ToString(movement.Id, ""),
								},
								Type: ct.FieldTypeSelect, Value: cu.IM{
									"name":    "place_code",
									"options": placeOpt(),
									"is_null": true,
									"value":   movement.PlaceCode,
								},
								FormTrigger: true,
							}},
						{Label: labels["movement_batchnumber"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "notes_" + cu.ToString(movement.Id, ""),
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":  "notes",
								"value": movement.MovementMeta.Notes,
							},
						}},
						{Label: labels["movement_qty"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "qty_" + cu.ToString(movement.Id, ""),
							},
							Type: ct.FieldTypeNumber, Value: cu.IM{
								"name":  "qty",
								"value": movement.MovementMeta.Qty,
							},
						}},
					}, Full: true, BorderBottom: false},
				},
				md.TransTypeFormula.String(): {
					{Columns: []ct.RowColumn{
						{Label: labels["movement_code"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "code_" + cu.ToString(movement.Id, ""),
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":     "movement_code",
								"value":    movement.Code,
								"disabled": true,
							},
						}},
						{
							Label: labels["product_code"], Value: ct.Field{
								BaseComponent: ct.BaseComponent{
									Name: "product_code_" + cu.ToString(movement.Id, ""),
								},
								Type: ct.FieldTypeSelector, Value: cu.IM{
									"name":  "product_code",
									"title": labels["view_product"],
									"value": ct.SelectOption{
										Value: movement.ProductCode,
										Text:  movement.ProductCode,
									},
									"fields":  productSelectorFields,
									"rows":    productSelectorRows,
									"link":    true,
									"is_null": false,
								},
								FormTrigger: true,
							},
						},
						{Label: labels["product_name"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "product_name_" + cu.ToString(movement.Id, ""),
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":     "product_name",
								"value":    cu.ToString(formData["product_name"], ""),
								"disabled": true,
							},
						}},
					}, Full: true, BorderBottom: true},
					{Columns: []ct.RowColumn{
						{
							Label: labels["place_name_movement"],
							Value: ct.Field{
								BaseComponent: ct.BaseComponent{
									Name: "place_code_" + cu.ToString(movement.Id, ""),
								},
								Type: ct.FieldTypeSelect, Value: cu.IM{
									"name":    "place_code",
									"options": placeOpt(),
									"is_null": true,
									"value":   movement.PlaceCode,
								},
								FormTrigger: true,
							}},
						{Label: labels["movement_qty"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "qty_" + cu.ToString(movement.Id, ""),
							},
							Type: ct.FieldTypeNumber, Value: cu.IM{
								"name":  "qty",
								"value": movement.MovementMeta.Qty,
							},
						}},
						{Label: labels["movement_shared"],
							Value: ct.Field{
								BaseComponent: ct.BaseComponent{
									Name: "shared_" + cu.ToString(movement.Id, ""),
								},
								Type: ct.FieldTypeBool, Value: cu.IM{
									"name":  "shared",
									"value": movement.MovementMeta.Shared,
								},
								FormTrigger: true,
							}},
					}, Full: true, BorderBottom: true},
					{Columns: []ct.RowColumn{
						{Label: labels["movement_notes"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "notes_" + cu.ToString(movement.Id, ""),
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":  "notes",
								"value": movement.MovementMeta.Notes,
							},
						}},
					}, Full: true, BorderBottom: false, FieldCol: true},
				},
				md.TransTypeWaybill.String(): {
					{Columns: []ct.RowColumn{
						{Label: labels["movement_code"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "code_" + cu.ToString(movement.Id, ""),
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":     "movement_code",
								"value":    movement.Code,
								"disabled": true,
							},
						}},
						{
							Label: labels["tool_code"], Value: ct.Field{
								BaseComponent: ct.BaseComponent{
									Name: "tool_code_" + cu.ToString(movement.Id, ""),
								},
								Type: ct.FieldTypeSelector, Value: cu.IM{
									"name":  "tool_code",
									"title": labels["view_tool"],
									"value": ct.SelectOption{
										Value: movement.ToolCode,
										Text:  movement.ToolCode,
									},
									"fields":  toolSelectorFields,
									"rows":    toolSelectorRows,
									"link":    true,
									"is_null": false,
								},
								FormTrigger: true,
							},
						},
						{Label: labels["movement_shipping_time"],
							Value: ct.Field{
								BaseComponent: ct.BaseComponent{
									Name: "shipping_time_" + cu.ToString(movement.Id, ""),
								},
								Type: ct.FieldTypeDateTime, Value: cu.IM{
									"name":    "shipping_time",
									"is_null": false,
									"value":   movement.ShippingTime,
								},
							}},
					}, Full: true, BorderBottom: true},
					{Columns: []ct.RowColumn{
						{Label: labels["tool_description"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "tool_description_" + cu.ToString(movement.Id, ""),
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":     "tool_description",
								"value":    cu.ToString(formData["tool_description"], ""),
								"disabled": true,
							},
						}},
						{Label: labels["movement_notes"], Value: ct.Field{
							BaseComponent: ct.BaseComponent{
								Name: "notes_" + cu.ToString(movement.Id, ""),
							},
							Type: ct.FieldTypeString, Value: cu.IM{
								"name":  "notes",
								"value": movement.MovementMeta.Notes,
							},
						}},
					}, Full: true, BorderBottom: false},
				},
			}
			diasbledMap := map[string]bool{
				md.TransTypeDelivery.String():   movement.PlaceCode == "" || movement.ProductCode == "",
				md.TransTypeInventory.String():  movement.ProductCode == "",
				md.TransTypeProduction.String(): movement.PlaceCode == "" || movement.ProductCode == "",
				md.TransTypeFormula.String():    movement.ProductCode == "",
				md.TransTypeWaybill.String():    movement.ToolCode == "",
			}
			frm := ct.Form{
				Title:      labels["movement_view"],
				Icon:       ct.IconListOl,
				BodyRows:   rowMap[transType],
				FooterRows: footerRows(diasbledMap[transType]),
			}
			return frm
		},
	}

	if frm, found := frmMap[formKey]; found {
		return frm()
	}
	return ct.Form{}
}
