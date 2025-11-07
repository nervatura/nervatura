package service

import (
	"errors"
	"fmt"
	"math"
	"slices"
	"strings"
	"time"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	api "github.com/nervatura/nervatura/v6/pkg/api"
	cp "github.com/nervatura/nervatura/v6/pkg/component/client/component"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

type TransService struct {
	cls *ClientService
}

func NewTransService(cls *ClientService) *TransService {
	return &TransService{
		cls: cls,
	}
}

var transRowTypeMap = map[string]func(stateData cu.IM) cu.IM{
	"items": func(stateData cu.IM) cu.IM {
		var item cu.IM
		defaultTax := cu.IM{}
		taxCodes := cu.ToIMA(stateData["tax_codes"], []cu.IM{})
		if len(taxCodes) > 0 {
			defaultTax = cu.ToIM(taxCodes[0], cu.IM{})
		}
		ut.ConvertToType(md.Item{
			//TransCode: cu.ToString(trans["code"], ""),
			TaxCode: cu.ToString(defaultTax["code"], ""),
			ItemMeta: md.ItemMeta{
				Qty:  1,
				Tags: []string{},
			},
			ItemMap: cu.IM{},
		}, &item)
		return item
	},
	"payments": func(stateData cu.IM) cu.IM {
		var payment cu.IM
		ut.ConvertToType(md.Payment{
			//TransCode: cu.ToString(trans["code"], ""),
			PaidDate: md.TimeDate{Time: time.Now()},
			PaymentMeta: md.PaymentMeta{
				Amount: 0,
				Tags:   []string{},
			},
			PaymentMap: cu.IM{},
		}, &payment)
		return payment
	},
	"movements": func(stateData cu.IM) cu.IM {
		var movement cu.IM
		ut.ConvertToType(md.Movement{
			MovementType: md.MovementType(md.MovementTypeHead),
			ShippingTime: md.TimeDateTime{Time: time.Now()},
			MovementMeta: md.MovementMeta{
				Qty:    0,
				Notes:  "",
				Shared: false,
				Tags:   []string{},
			},
			MovementMap: cu.IM{},
		}, &movement)
		return movement
	},
	"link": func(stateData cu.IM) cu.IM {
		var link cu.IM
		ut.ConvertToType(md.Link{
			LinkType1: md.LinkType(md.LinkTypePayment),
			LinkType2: md.LinkType(md.LinkTypeTrans),
			LinkMeta: md.LinkMeta{
				Amount: 0,
				Qty:    0,
				Rate:   1,
				Tags:   []string{},
			},
			LinkMap: cu.IM{},
		}, &link)
		return link
	},
}

var transDataQueryBase = map[string]func(trans cu.IM) ([]string, md.Query){
	"items": func(trans cu.IM) ([]string, md.Query) {
		return []string{}, md.Query{
			Fields: []string{"*"}, From: "item",
			Filters: []md.Filter{
				{Field: "trans_code", Comp: "==", Value: cu.ToString(trans["code"], "")},
				{Field: "deleted", Comp: "==", Value: false},
			},
		}
	},
	"movements": func(trans cu.IM) ([]string, md.Query) {
		return []string{}, md.Query{
			Fields: []string{"*"}, From: "movement",
			Filters: []md.Filter{
				{Field: "trans_code", Comp: "==", Value: cu.ToString(trans["code"], "")},
				{Field: "deleted", Comp: "==", Value: false},
			},
		}
	},
	"payments": func(trans cu.IM) ([]string, md.Query) {
		return []string{}, md.Query{
			Fields: []string{"*"}, From: "payment",
			Filters: []md.Filter{
				{Field: "trans_code", Comp: "==", Value: cu.ToString(trans["code"], "")},
				{Field: "deleted", Comp: "==", Value: false},
			},
		}
	},
	"payment_link": func(trans cu.IM) ([]string, md.Query) {
		transType := cu.ToString(trans["trans_type"], "")
		query := md.Query{
			Fields: []string{"l.*", "pl.place_name", "pl.currency_code", "pt.trans_type",
				"pt.direction", "pt.code as pt_code", "pm.paid_date", "it.currency_code as invoice_curr"},
			From: `link l inner join payment_view pm on l.link_code_1 = pm.code inner join trans_view it on l.link_code_2 = it.code 
			inner join trans_view pt on pm.trans_code = pt.code inner join place pl on pt.place_code = pl.code`,
			Filters: []md.Filter{
				{Field: "l.link_type_1", Comp: "==", Value: md.LinkTypePayment.String()},
				{Field: "l.link_type_2", Comp: "==", Value: md.LinkTypeTrans.String()},
				{Field: "it.trans_type", Comp: "in", Value: fmt.Sprintf("%s,%s",
					md.TransTypeInvoice.String(), md.TransTypeReceipt.String())},
				{Field: "it.deleted", Comp: "==", Value: false},
				{Field: "l.deleted", Comp: "==", Value: false},
			},
			OrderBy: []string{"pm.id"},
		}
		if slices.Contains([]string{md.TransTypeBank.String(), md.TransTypeCash.String()}, transType) {
			query.Filters = append(query.Filters, md.Filter{Field: "pm.trans_code", Comp: "==", Value: cu.ToString(trans["code"], "")})
		}
		if slices.Contains([]string{md.TransTypeInvoice.String(), md.TransTypeReceipt.String()}, transType) {
			query.Filters = append(query.Filters, md.Filter{Field: "l.link_code_2", Comp: "==", Value: cu.ToString(trans["code"], "")})
		}
		return []string{md.TransTypeInvoice.String(), md.TransTypeReceipt.String(),
			md.TransTypeBank.String(), md.TransTypeCash.String()}, query
	},
	"trans_cancellation": func(trans cu.IM) ([]string, md.Query) {
		return []string{}, md.Query{
			Fields: []string{"*"}, From: "trans_view",
			Filters: []md.Filter{
				{Field: "status", Comp: "==", Value: md.TransStatusCancellation.String()},
				{Field: "trans_code", Comp: "==", Value: cu.ToString(trans["code"], "")},
				{Field: "deleted", Comp: "==", Value: true},
			},
		}
	},
	"transitem_invoice": func(trans cu.IM) ([]string, md.Query) {
		return []string{md.TransTypeOrder.String(), md.TransTypeWorksheet.String(), md.TransTypeRent.String()},
			md.Query{
				Fields: []string{"i.*", "t.id as trans_id", "t.trans_date", "t.currency_code"},
				From:   "trans t inner join item i on i.trans_code = t.code",
				Filters: []md.Filter{
					{Field: "t.trans_code", Comp: "==", Value: cu.ToString(trans["code"], "")},
					{Field: "t.trans_type", Comp: "in", Value: fmt.Sprintf("%s,%s",
						md.TransTypeInvoice.String(), md.TransTypeReceipt.String())},
					{Field: "t.direction", Comp: "==", Value: cu.ToString(trans["direction"], "")},
					{Field: "t.deleted", Comp: "==", Value: false},
					{Field: "i.deleted", Comp: "==", Value: false},
				},
			}
	},
	"element_count": func(trans cu.IM) ([]string, md.Query) {
		return []string{}, md.Query{
			Fields: []string{"p.*"},
			From:   "item i inner join product p on p.code = i.product_code",
			Filters: []md.Filter{
				{Field: "i.trans_code", Comp: "==", Value: cu.ToString(trans["code"], "")},
				{Field: "p.product_type", Comp: "==", Value: md.ProductTypeItem.String()},
				{Field: "p.deleted", Comp: "==", Value: false},
				{Field: "i.deleted", Comp: "==", Value: false},
			},
		}
	},
	"config_report": func(trans cu.IM) ([]string, md.Query) {
		transType := strings.TrimPrefix(cu.ToString(trans["trans_type"], ""), "TRANS_")
		return []string{}, md.Query{
			Fields: []string{"id", "report_key", "report_name"}, From: "config_report",
			Filters: []md.Filter{
				{Field: "report_type", Comp: "==", Value: "TRANS"},
				{Field: "trans_type", Comp: "==", Value: transType},
				{Field: "direction", Comp: "==", Value: trans["direction"]},
			},
		}
	},
	"movement_inventory": func(trans cu.IM) ([]string, md.Query) {
		return []string{md.TransTypeDelivery.String(), md.TransTypeInventory.String(), md.TransTypeProduction.String()},
			md.Query{
				Fields: []string{"*"},
				From:   "movement_inventory",
				Filters: []md.Filter{
					{Field: "trans_code", Comp: "==", Value: cu.ToString(trans["code"], "")},
				},
			}
	},
	"movement_waybill": func(trans cu.IM) ([]string, md.Query) {
		return []string{md.TransTypeWaybill.String()},
			md.Query{
				Fields: []string{"*"},
				From:   "movement_waybill",
				Filters: []md.Filter{
					{Field: "trans_code", Comp: "==", Value: cu.ToString(trans["code"], "")},
				},
			}
	},
	"movement_formula": func(trans cu.IM) ([]string, md.Query) {
		return []string{md.TransTypeFormula.String()},
			md.Query{
				Fields: []string{"*"},
				From:   "movement_formula",
				Filters: []md.Filter{
					{Field: "trans_code", Comp: "==", Value: cu.ToString(trans["code"], "")},
				},
			}
	},
	"transitem_shipping": func(trans cu.IM) ([]string, md.Query) {
		return []string{md.TransTypeOrder.String(), md.TransTypeWorksheet.String(), md.TransTypeRent.String()},
			md.Query{
				Fields: []string{
					"i.id", "i.code", "i.description", "i.qty", "mv.trans_code",
					"mv.shipping_time as shipping_date", "mv.qty as shipping_qty", "pr.product_name",
					"'trans' as editor",
				},
				From: `item_view i inner join movement_view mv on mv.item_code = i.code inner join product pr on pr.code = mv.product_code`,
				Filters: []md.Filter{
					{Field: "i.trans_code", Comp: "==", Value: cu.ToString(trans["code"], "")},
				},
			}
	},
	"tool_movement": func(trans cu.IM) ([]string, md.Query) {
		return []string{
				md.TransTypeOrder.String(), md.TransTypeWorksheet.String(), md.TransTypeRent.String(),
				md.TransTypeInvoice.String(), md.TransTypeReceipt.String()},
			md.Query{
				Fields: []string{
					"mv.id", "t.code as trans_code", "t.direction", "mv.shipping_time", "tl.serial_number", "tl.description",
					"'trans' as editor",
				},
				From: `trans_view t inner join movement_view mv on mv.trans_code = t.code inner join tool_view tl on tl.code = mv.tool_code`,
				Filters: []md.Filter{
					{Field: "t.trans_type", Comp: "==", Value: md.TransTypeWaybill.String()},
					{Field: "t.trans_code", Comp: "==", Value: cu.ToString(trans["code"], "")},
				},
			}
	},
}

var transDataQueryExt = map[string]md.Query{
	"config_map": {
		Fields: []string{"*"}, From: "config_map",
	},
	"config_data": {
		Fields: []string{"*"}, From: "config_data",
	},
	"tax_codes": {
		Fields: []string{"code", "description", "rate_value"}, From: "tax_view",
		OrderBy: []string{"code"},
	},
	"currencies": {
		Fields: []string{"code", "description", "digit"}, From: "currency_view",
		OrderBy: []string{"code"},
	},
	"places": {
		Fields: []string{"code", "place_type", "place_name", "currency_code", "inactive"}, From: "place_view",
		OrderBy: []string{"code"},
	},
}

func (s *TransService) dataDefault(trans cu.IM, configData []cu.IM) (data cu.IM) {
	transMeta := cu.ToIM(trans["trans_meta"], cu.IM{})
	transType := cu.ToString(trans["trans_type"], "")
	valueMap := map[string]func(value string){
		"default_taxcode": func(value string) {
			trans["tax_code"] = value
		},
		"default_currency": func(value string) {
			if cp.TransIsItem(transType) {
				trans["currency_code"] = value
			}
		},
		"default_deadline": func(value string) {
			if transType == md.TransTypeInvoice.String() {
				transMeta["due_time"] = md.TimeDateTime{Time: time.Now().AddDate(0, 0, int(cu.ToInteger(value, 0)))}
			}
			if !slices.Contains([]string{
				md.TransTypeInvoice.String(), md.TransTypeReceipt.String()}, transType) {
				transMeta["rate"] = cu.ToInteger(value, 0)
			}
		},
		"default_paidtype": func(value string) {
			if cp.TransIsItem(transType) {
				transMeta["paid_type"] = value
			}
		},
		"default_bank": func(value string) {
			if transType == md.TransTypeBank.String() {
				trans["place_code"] = value
			}
		},
		"default_chest": func(value string) {
			if transType == md.TransTypeCash.String() {
				trans["place_code"] = value
			}
		},
	}
	for _, cf := range configData {
		if fn, ok := valueMap[cu.ToString(cf["config_key"], "")]; ok {
			fn(cu.ToString(cf["config_value"], ""))
		}
	}
	trans["trans_meta"] = transMeta
	return trans
}

func (s *TransService) Data(evt ct.ResponseEvent, params cu.IM) (data cu.IM, err error) {
	client := evt.Trigger.(*ct.Client)
	ds := s.cls.getDataStore(client.Ticket.Database)
	user := cu.ToIM(client.Ticket.User, cu.IM{})

	data = cu.IM{
		"trans": cu.IM{
			"trans_type": cu.ToString(params["trans_type"], md.TransType(0).String()),
			"direction":  cu.ToString(params["direction"], md.Direction(0).String()),
			"trans_date": md.TimeDateTime{Time: time.Now()}.Format(time.DateOnly),
			"trans_code": cu.ToString(params["ref_trans_code"], ""),
			"trans_meta": cu.IM{
				"status":      md.TransStatusNormal.String(),
				"trans_state": md.TransStateOK.String(),
				"paid_type":   md.PaidTypeOnline.String(),
				"tags":        []string{},
				"worksheet":   cu.IM{},
				"rent":        cu.IM{},
				"invoice":     cu.IM{},
			},
			"trans_map": cu.IM{},
		},
		"items":              []cu.IM{},
		"movements":          []cu.IM{},
		"payments":           []cu.IM{},
		"payment_link":       []cu.IM{},
		"config_map":         []cu.IM{},
		"config_data":        []cu.IM{},
		"config_report":      []cu.IM{},
		"tax_codes":          []cu.IM{},
		"currencies":         []cu.IM{},
		"places":             []cu.IM{},
		"transitem_invoice":  []cu.IM{},
		"trans_cancellation": []cu.IM{},
		"element_count":      []cu.IM{},
		"movement_inventory": []cu.IM{},
		"movement_waybill":   []cu.IM{},
		"movement_formula":   []cu.IM{},
		"transitem_shipping": []cu.IM{},
		"user":               user,
		"dirty":              false,
		"editor_icon":        cu.ToString(params["editor_icon"], ct.IconFileText),
		"editor_title":       cu.ToString(params["editor_title"], ""),
		"customer_name":      "",
		"project_name":       "",
		"place_name":         "",
	}
	if cu.ToString(params["trans_type"], md.TransType(0).String()) == md.TransTypeCash.String() {
		data["payments"] = []cu.IM{transRowTypeMap["payments"](data)}
	}

	var rows []cu.IM = []cu.IM{}
	if cu.ToString(params["trans_id"], "") != "" || cu.ToString(params["trans_code"], "") != "" {
		var trans []cu.IM = []cu.IM{}
		if trans, err = ds.StoreDataQuery(md.Query{
			Fields: []string{"*"}, From: "trans",
			Filters: []md.Filter{
				//{Field: "deleted", Comp: "==", Value: false},
				{Field: "id", Comp: "==", Value: cu.ToInteger(params["trans_id"], 0)},
				{Or: true, Field: "code", Comp: "==", Value: cu.ToString(params["trans_code"], "")},
			},
		}, false); err != nil {
			return data, err
		}
		if len(trans) > 0 {
			data["trans"] = trans[0]
			transType := cu.ToString(trans[0]["trans_type"], "")
			data["editor_title"] = cu.ToString(trans[0]["code"], "")
			data["editor_icon"] = cp.TransTypeIcon(transType)
			data["customer_name"] = s.cls.codeName(ds, cu.ToString(trans[0]["customer_code"], ""), "customer")

			for key, query := range transDataQueryBase {
				if tt, qr := query(trans[0]); len(tt) == 0 || slices.Contains(tt, transType) {
					if rows, err = ds.StoreDataQuery(qr, false); err != nil {
						return data, err
					}
					data[key] = rows
				}
			}
		}
	}

	trans := cu.ToIM(data["trans"], cu.IM{})
	for key, query := range transDataQueryExt {
		if rows, err = ds.StoreDataQuery(query, false); err != nil {
			return data, err
		}
		data[key] = rows
	}

	if cu.ToInteger(trans["id"], 0) == 0 {
		data["trans"] = s.dataDefault(trans, cu.ToIMA(data["config_data"], []cu.IM{}))
	}

	return data, err
}

func (s *TransService) Response(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	switch evt.Name {
	case ct.FormEventOK:
		return s.formNext(evt)

	case ct.ClientEventForm:
		return s.formEvent(evt)

	case ct.ClientEventSideMenu:
		return s.sideMenu(evt)

	default:
		return s.editorField(evt)
	}
}

func (s *TransService) updateItems(ds *api.DataStore, data cu.IM, transCode string) (err error) {
	for _, it := range cu.ToIMA(data["items"], []cu.IM{}) {
		var item md.Item = md.Item{
			TransCode: transCode,
			ItemMeta: md.ItemMeta{
				Tags: []string{},
			},
			ItemMap: cu.IM{},
		}
		if err = ut.ConvertToType(it, &item); err == nil {
			values := cu.IM{
				"trans_code":   transCode,
				"product_code": item.ProductCode,
				"tax_code":     item.TaxCode,
			}
			ut.ConvertByteToIMData(item.ItemMeta, values, "item_meta")
			ut.ConvertByteToIMData(item.ItemMap, values, "item_map")

			itemID := item.Id
			update := md.Update{Values: values, Model: "item"}
			if itemID > 0 {
				update.IDKey = itemID
			}
			if itemID, err = ds.StoreDataUpdate(update); err != nil {
				return err
			}
			it["id"] = itemID
		}
	}
	return err
}

func (s *TransService) updatePayments(ds *api.DataStore, data cu.IM, transCode string) (err error) {
	for _, it := range cu.ToIMA(data["payments"], []cu.IM{}) {
		var payment md.Payment = md.Payment{
			TransCode: transCode,
			PaidDate:  md.TimeDate{Time: time.Now()},
			PaymentMeta: md.PaymentMeta{
				Tags: []string{},
			},
			PaymentMap: cu.IM{},
		}
		if err = ut.ConvertToType(it, &payment); err == nil {
			values := cu.IM{
				"trans_code": transCode,
				"paid_date":  payment.PaidDate.String(),
			}
			ut.ConvertByteToIMData(payment.PaymentMeta, values, "payment_meta")
			ut.ConvertByteToIMData(payment.PaymentMap, values, "payment_map")

			paymentID := payment.Id
			update := md.Update{Values: values, Model: "payment"}
			if paymentID > 0 {
				update.IDKey = paymentID
			}
			if paymentID, err = ds.StoreDataUpdate(update); err != nil {
				return err
			}
			it["id"] = paymentID
		}
	}
	return err
}

var movementUpdateValidate = []func(trans md.Trans, movement md.Movement, data cu.IM, msgFunc func(labelID string) string) (bool, error){
	func(trans md.Trans, movement md.Movement, data cu.IM, msgFunc func(labelID string) string) (bool, error) {
		errMsg := msgFunc("missing_required_field") + ": " + msgFunc("place_name_movement")
		return trans.TransType == md.TransTypeDelivery && trans.Direction != md.DirectionTransfer &&
			movement.PlaceCode == "", errors.New(errMsg)
	},
	func(trans md.Trans, movement md.Movement, data cu.IM, msgFunc func(labelID string) string) (bool, error) {
		errMsg := msgFunc("missing_required_field") + ": " + msgFunc("place_name_target")
		return trans.TransType == md.TransTypeDelivery && trans.Direction == md.DirectionTransfer &&
			movement.PlaceCode == "", errors.New(errMsg)
	},
	/*
		func(trans md.Trans, movement md.Movement, data cu.IM, msgFunc func(labelID string) string) (bool, error) {
			errMsg := msgFunc("movement_transfer_error")
			return trans.TransType == md.TransTypeDelivery && trans.Direction == md.DirectionTransfer &&
					movement.PlaceCode == trans.PlaceCode,
				errors.New(errMsg)
		},
	*/
	func(trans md.Trans, movement md.Movement, data cu.IM, msgFunc func(labelID string) string) (bool, error) {
		errMsg := msgFunc("missing_required_field") + ": " + msgFunc("product_code")
		return slices.Contains([]string{
				md.TransTypeDelivery.String(), md.TransTypeFormula.String(), md.TransTypeProduction.String(),
				md.TransTypeInventory.String()}, trans.TransType.String()) &&
				movement.ProductCode == "",
			errors.New(errMsg)
	},
	func(trans md.Trans, movement md.Movement, data cu.IM, msgFunc func(labelID string) string) (bool, error) {
		errMsg := msgFunc("missing_required_field") + ": " + msgFunc("place_name_movement")
		return trans.TransType == md.TransTypeProduction && movement.PlaceCode == "",
			errors.New(errMsg)
	},
	func(trans md.Trans, movement md.Movement, data cu.IM, msgFunc func(labelID string) string) (bool, error) {
		errMsg := msgFunc("missing_required_field") + ": " + msgFunc("tool_code")
		return trans.TransType == md.TransTypeWaybill && movement.ToolCode == "",
			errors.New(errMsg)
	},
}

func (s *TransService) updateMovements(ds *api.DataStore, data cu.IM, trans md.Trans, msgFunc func(labelID string) string) (err error) {
	transferPair := func(movement md.Movement) (baseMovement, targetMovement md.Movement, isNewTransfer bool) {
		isNewTransfer = (trans.TransType == md.TransTypeDelivery) && (trans.Direction == md.DirectionTransfer) && (movement.Id < 0)
		if isNewTransfer {
			targetMovement = md.Movement{}
			ut.ConvertToType(movement, &targetMovement)
			movement.MovementMeta.Qty = -movement.MovementMeta.Qty
			movement.PlaceCode = trans.PlaceCode
			return movement, targetMovement, isNewTransfer
		}
		return movement, targetMovement, isNewTransfer
	}

	updateMovement := func(movement md.Movement) (movementID int64, err error) {
		values := cu.IM{
			"movement_type": movement.MovementType.String(),
			"shipping_time": movement.ShippingTime.String(),
			"trans_code":    trans.Code,
		}

		// Optional fields
		optionalFields := map[string]string{
			"product_code":  movement.ProductCode,
			"tool_code":     movement.ToolCode,
			"place_code":    movement.PlaceCode,
			"item_code":     movement.ItemCode,
			"movement_code": movement.MovementCode,
		}

		for key, value := range optionalFields {
			if value != "" {
				values[key] = value
			}
		}

		ut.ConvertByteToIMData(movement.MovementMeta, values, "movement_meta")
		ut.ConvertByteToIMData(movement.MovementMap, values, "movement_map")

		update := md.Update{Values: values, Model: "movement"}
		if movement.Id > 0 {
			update.IDKey = movement.Id
		}
		return ds.StoreDataUpdate(update)
	}

	for _, mv := range cu.ToIMA(data["movements"], []cu.IM{}) {
		var movement md.Movement = md.Movement{
			TransCode: trans.TransCode,
			MovementMeta: md.MovementMeta{
				Tags: []string{},
			},
			MovementMap: cu.IM{},
		}
		if err = ut.ConvertToType(mv, &movement); err == nil {

			//validate movement data
			for _, validate := range movementUpdateValidate {
				if invalid, err := validate(trans, movement, data, msgFunc); invalid {
					return err
				}
			}

			movement, targetMovement, isNewTransfer := transferPair(movement)
			if mv["id"], err = updateMovement(movement); err != nil {
				return err
			}
			if isNewTransfer {
				var rows []cu.IM
				if rows, err = ds.StoreDataGet(cu.IM{"id": mv["id"], "model": "movement"}, true); err == nil {
					targetMovement.MovementCode = cu.ToString(rows[0]["code"], "")
					if _, err = updateMovement(targetMovement); err != nil {
						return err
					}
				}
			}
		}
	}
	return err
}

func (s *TransService) updateLinks(ds *api.DataStore, data cu.IM) (err error) {
	for _, ln := range cu.ToIMA(data["payment_link"], []cu.IM{}) {
		var link md.Link = md.Link{
			LinkType1: md.LinkType(md.LinkTypePayment),
			LinkType2: md.LinkType(md.LinkTypeTrans),
			LinkMeta: md.LinkMeta{
				Tags: []string{},
			},
			LinkMap: cu.IM{},
		}
		if err = ut.ConvertToType(ln, &link); err == nil {
			values := cu.IM{
				"link_type_1": link.LinkType1.String(),
				"link_code_1": link.LinkCode1,
				"link_type_2": link.LinkType2.String(),
				"link_code_2": link.LinkCode2,
			}
			ut.ConvertByteToIMData(link.LinkMeta, values, "link_meta")
			ut.ConvertByteToIMData(link.LinkMap, values, "link_map")

			linkID := link.Id
			update := md.Update{Values: values, Model: "link"}
			if linkID > 0 {
				update.IDKey = linkID
			}
			if linkID, err = ds.StoreDataUpdate(update); err != nil {
				return err
			}
			ln["id"] = linkID
		}
	}
	return err
}

func (s *TransService) updateDeleteRows(ds *api.DataStore, data cu.IM) (err error) {
	deleteRows := func(model, deleteKey string) (err error) {
		for _, it := range cu.ToIMA(data[deleteKey], []cu.IM{}) {
			if err = ds.DataDelete(model, cu.ToInteger(it["id"], 0), ""); err != nil {
				return err
			}
		}
		return err
	}

	deleteMap := cu.SM{
		"items_delete":        "item",
		"payments_delete":     "payment",
		"payment_link_delete": "link",
		"movements_delete":    "movement",
	}

	for deleteKey, model := range deleteMap {
		if err = deleteRows(model, deleteKey); err != nil {
			return err
		}
	}
	return err
}

var transUpdateValidate = []func(trans md.Trans, data cu.IM, msgFunc func(labelID string) string) (bool, error){
	func(trans md.Trans, data cu.IM, msgFunc func(labelID string) string) (bool, error) {
		errMsg := msgFunc("missing_required_field") + ": " + msgFunc("customer_name")
		return cp.TransIsItem(trans.TransType.String()) && trans.CustomerCode == "",
			errors.New(errMsg)
	},
	func(trans md.Trans, data cu.IM, msgFunc func(labelID string) string) (bool, error) {
		errMsg := msgFunc("missing_required_field") + ": " + msgFunc("place_name_payment")
		return (trans.TransType == md.TransTypeCash || trans.TransType == md.TransTypeBank) && trans.PlaceCode == "",
			errors.New(errMsg)
	},
	func(trans md.Trans, data cu.IM, msgFunc func(labelID string) string) (bool, error) {
		errMsg := msgFunc("missing_required_field") + ": " + msgFunc("place_name_movement")
		return (trans.TransType == md.TransTypeInventory || trans.TransType == md.TransTypeProduction ||
				(trans.TransType == md.TransTypeDelivery && trans.Direction == md.DirectionTransfer)) && trans.PlaceCode == "",
			errors.New(errMsg)
	},
	func(trans md.Trans, data cu.IM, msgFunc func(labelID string) string) (bool, error) {
		errMsg := msgFunc("missing_required_field") + ": " + msgFunc("due_time")
		return trans.TransType == md.TransTypeProduction && trans.TransMeta.DueTime.Time.IsZero(),
			errors.New(errMsg)
	},
	func(trans md.Trans, data cu.IM, msgFunc func(labelID string) string) (bool, error) {
		movements := cu.ToIMA(data["movements"], []cu.IM{})
		errMsg := msgFunc("missing_required_field") + ": " + msgFunc("product_code")
		return (trans.TransType == md.TransTypeProduction || trans.TransType == md.TransTypeFormula) &&
				(len(movements) == 0 || cu.ToString(movements[0]["product_code"], "") == ""),
			errors.New(errMsg)
	},
	func(trans md.Trans, data cu.IM, msgFunc func(labelID string) string) (bool, error) {
		errMsg := msgFunc("missing_required_field") + ": " + msgFunc("trans_code") + " or " + msgFunc("customer_name") + " or " + msgFunc("employee_code")
		return trans.TransType == md.TransTypeWaybill && trans.TransCode == "" && trans.CustomerCode == "" && trans.EmployeeCode == "",
			errors.New(errMsg)
	},
}

func (s *TransService) update(ds *api.DataStore, data cu.IM, msgFunc func(labelID string) string) (transID int64, err error) {
	user := cu.ToIM(data["user"], cu.IM{})
	var trans md.Trans = md.Trans{}
	ut.ConvertToType(data["trans"], &trans)

	for _, validate := range transUpdateValidate {
		if invalid, err := validate(trans, data, msgFunc); invalid {
			return 0, err
		}
	}

	values := cu.IM{
		"trans_type":    trans.TransType.String(),
		"direction":     trans.Direction.String(),
		"trans_date":    trans.TransDate.String(),
		"customer_code": nil,
		"employee_code": nil,
		"project_code":  nil,
		"place_code":    nil,
		"trans_code":    nil,
		"currency_code": nil,
	}

	// Optional fields
	optionalFields := map[string]string{
		"code":          trans.Code,
		"customer_code": trans.CustomerCode,
		"employee_code": trans.EmployeeCode,
		"project_code":  trans.ProjectCode,
		"place_code":    trans.PlaceCode,
		"trans_code":    trans.TransCode,
		"currency_code": trans.CurrencyCode,
	}

	for key, value := range optionalFields {
		if value != "" {
			values[key] = value
		}
	}

	ut.ConvertByteToIMData(trans.TransMeta, values, "trans_meta")
	ut.ConvertByteToIMData(trans.TransMap, values, "trans_map")

	//var transID int64
	newTrans := (trans.Id == 0)
	update := md.Update{Values: values, Model: "trans"}
	if !newTrans {
		update.IDKey = trans.Id
	} else {
		values["auth_code"] = user["code"]
	}
	if transID, err = ds.StoreDataUpdate(update); err != nil {
		return 0, err
	}
	if newTrans {
		var rows []cu.IM = []cu.IM{}
		if rows, err = ds.StoreDataGet(cu.IM{"id": transID, "model": "trans"}, true); err == nil {
			data["trans"] = rows[0]
			trans.Code = cu.ToString(cu.ToIM(rows[0], cu.IM{})["code"], "")
			data["editor_title"] = trans.Code
		}
	}

	if err != nil {
		return transID, err
	}

	if err = s.updateItems(ds, data, trans.Code); err != nil {
		return transID, err
	}

	if err = s.updatePayments(ds, data, trans.Code); err != nil {
		return transID, err
	}

	if err = s.updateLinks(ds, data); err != nil {
		return transID, err
	}

	if err = s.updateMovements(ds, data, trans, msgFunc); err != nil {
		return transID, err
	}

	err = s.updateDeleteRows(ds, data)

	return transID, err
}

func (s *TransService) delete(ds *api.DataStore, transID int64) (err error) {
	return ds.DataDelete("trans", transID, "")
}

func (s *TransService) formNext(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	view := cu.ToString(stateData["view"], "")
	ds := s.cls.getDataStore(client.Ticket.Database)

	trans := cu.ToIM(stateData["trans"], cu.IM{})
	transMeta := cu.ToIM(trans["trans_meta"], cu.IM{})
	transMap := cu.ToIM(trans["trans_map"], cu.IM{})

	resultUpdate := func(dirty bool) (re ct.ResponseEvent, err error) {
		trans["trans_meta"] = transMeta
		trans["trans_map"] = transMap
		stateData["trans"] = trans
		if dirty {
			stateData["dirty"] = dirty
		}
		client.SetEditor("trans", view, stateData)
		return evt, err
	}

	frmValues := cu.ToIM(evt.Value, cu.IM{})
	frmData := cu.ToIM(frmValues["data"], cu.IM{})
	frmValue := cu.ToIM(frmValues["value"], cu.IM{})

	nextMap := map[string]func() (re ct.ResponseEvent, err error){
		"editor_cancel": func() (re ct.ResponseEvent, err error) {
			client.ResetEditor()
			return evt, err
		},

		"editor_delete": func() (re ct.ResponseEvent, err error) {
			if err = s.delete(ds, cu.ToInteger(trans["id"], 0)); err != nil {
				return evt, err
			}
			client.ResetEditor()
			return evt, err
		},

		"transitem_new": func() (re ct.ResponseEvent, err error) {
			return s.cls.setEditor(evt, "trans", cu.IM{
				"session_id":   client.Ticket.SessionID,
				"trans_type":   frmValue["create_trans_type"],
				"direction":    frmValue["create_direction"],
				"editor_title": client.Msg("transitem_title"),
				"editor_icon":  ct.IconFileText,
			}), nil
		},

		"transpayment_new": func() (re ct.ResponseEvent, err error) {
			transType := cu.ToString(frmValue["create_trans_type"], "")
			direction := cu.ToString(frmValue["create_direction"], "")
			if transType == md.TransTypeBank.String() {
				direction = md.DirectionTransfer.String()
			}
			return s.cls.setEditor(evt, "trans", cu.IM{
				"session_id":   client.Ticket.SessionID,
				"trans_type":   transType,
				"direction":    direction,
				"editor_title": client.Msg("transpayment_title"),
				"editor_icon":  ct.IconMoney,
			}), nil
		},

		"transmovement_new": func() (re ct.ResponseEvent, err error) {
			transType := cu.ToString(frmValue["create_trans_type"], "")
			direction := cu.ToString(frmValue["create_direction"], md.DirectionTransfer.String())
			return s.cls.setEditor(evt, "trans", cu.IM{
				"session_id":   client.Ticket.SessionID,
				"trans_type":   transType,
				"direction":    direction,
				"editor_title": client.Msg(strings.ToLower(transType + "_new")),
				"editor_icon":  cp.TransTypeIcon(transType),
			}), nil
		},

		"trans_create": func() (re ct.ResponseEvent, err error) {
			return s.createData(evt, frmValue)
		},

		"trans_copy": func() (re ct.ResponseEvent, err error) {
			return s.createData(evt, cu.IM{
				"status": md.TransStatusNormal.String(),
			})
		},

		"trans_corrective": func() (re ct.ResponseEvent, err error) {
			return s.createData(evt, cu.IM{
				"status": md.TransStatusAmendment.String(),
			})
		},

		"trans_cancellation": func() (re ct.ResponseEvent, err error) {
			return s.createData(evt, cu.IM{
				"status": md.TransStatusCancellation.String(),
			})
		},

		"customer": func() (re ct.ResponseEvent, err error) {
			params := cu.ToIM(stateData["params"], cu.IM{})
			return s.cls.setEditor(evt, "customer", params), nil
		},

		"trans": func() (re ct.ResponseEvent, err error) {
			params := cu.ToIM(stateData["params"], cu.IM{})
			return s.cls.setEditor(evt, "trans", params), nil
		},

		"shipping": func() (re ct.ResponseEvent, err error) {
			params := cu.ToIM(stateData["params"], cu.IM{})
			return s.cls.setEditor(evt, "shipping", params), nil
		},

		"employee": func() (re ct.ResponseEvent, err error) {
			params := cu.ToIM(stateData["params"], cu.IM{})
			return s.cls.setEditor(evt, "employee", params), nil
		},

		"project": func() (re ct.ResponseEvent, err error) {
			params := cu.ToIM(stateData["params"], cu.IM{})
			return s.cls.setEditor(evt, "project", params), nil
		},

		"product": func() (re ct.ResponseEvent, err error) {
			params := cu.ToIM(stateData["params"], cu.IM{})
			return s.cls.setEditor(evt, "product", params), nil
		},

		"editor_add_tag": func() (re ct.ResponseEvent, err error) {
			tag := cu.ToString(frmValue["value"], "")
			if tag != "" {
				tags := ut.ToStringArray(transMeta["tags"])
				if !slices.Contains(tags, tag) {
					tags = append(tags, tag)
					transMeta["tags"] = tags
					return resultUpdate(true)
				}
			}
			return evt, nil
		},

		"form_add_tag": func() (re ct.ResponseEvent, err error) {
			tag := cu.ToString(frmValue["value"], "")
			frmKey := cu.ToString(frmData["frm_key"], "")
			frmIndex := cu.ToInteger(frmData["frm_index"], 0)
			row := cu.ToIM(frmData["row"], cu.IM{})
			metaName := ut.MetaName(row, "_meta")
			rowField := cu.ToString(frmData["row_field"], "")
			if tag != "" {
				tags := ut.ToStringArray(row[rowField])
				if metaName != "" {
					tags = ut.ToStringArray(cu.ToIM(row[metaName], cu.IM{})[rowField])
				}
				if !slices.Contains(tags, tag) {
					tags = append(tags, tag)
					if metaName != "" {
						cu.ToIM(row[metaName], cu.IM{})[rowField] = tags
					} else {
						row[rowField] = tags
					}
					client.SetForm(frmKey, row, frmIndex, false)
					return evt, nil
				}
			}
			return evt, nil
		},

		"bookmark_add": func() (re ct.ResponseEvent, err error) {
			label := cu.ToString(frmValue["value"], "")
			bookmark := md.Bookmark{
				BookmarkType: md.BookmarkTypeEditor,
				Label:        label,
				Key:          "trans",
				Code:         cu.ToString(trans["code"], ""),
				Filters:      []any{},
				Columns:      map[string]bool{},
				TimeStamp:    md.TimeDateTime{Time: time.Now()},
			}
			return s.cls.addBookmark(evt, bookmark), nil
		},

		"editor_map_value": func() (re ct.ResponseEvent, err error) {
			code := cu.ToString(frmValue["value"], "")
			model := cu.ToString(frmData["model"], "")
			mapField := cu.ToString(frmData["map_field"], "")
			if _, err := ds.GetDataByID(model, 0, code, true); err != nil {
				frmData["label"] = fmt.Sprintf("%s: %s (%s)", client.Msg("invalid_code"), code, model)
				frmData["default_value"] = code
				frmData["invalid"] = true
				client.SetForm("input_string", frmData, 0, true)
				return evt, nil
			}
			transMap[mapField] = code
			stateData["map_field"] = ""
			return resultUpdate(true)
		},
	}

	if fn, ok := nextMap[cu.ToString(frmData["next"], "")]; ok {
		return fn()
	}
	return evt, err
}

func (s *TransService) getProductPrice(ds *api.DataStore, options cu.IM) (price float64, discount float64) {
	if results, err := ds.ProductPrice(options); err == nil {
		price = cu.ToFloat(results["price"], 0)
		discount = cu.ToFloat(results["discount"], 0)
	}
	return price, discount
}

func (s *TransService) roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func (s *TransService) taxRate(taxCodes []cu.IM, taxCode string) (rate float64) {
	if idx := slices.IndexFunc(taxCodes, func(c cu.IM) bool {
		return cu.ToString(c["code"], "") == taxCode
	}); idx > int(-1) {
		return cu.ToFloat(taxCodes[idx]["rate_value"], 0)
	}
	return rate
}

func (s *TransService) currencyDigit(currencies []cu.IM, currencyCode string) (digit int64) {
	if idx := slices.IndexFunc(currencies, func(c cu.IM) bool {
		return cu.ToString(c["code"], "") == currencyCode
	}); idx > int(-1) {
		return cu.ToInteger(currencies[idx]["digit"], 0)
	}
	return digit
}

func (s *TransService) calcItemPrice(calcMode string, value float64, stateData, formRow cu.IM) cu.IM {
	trans := cu.ToIM(stateData["trans"], cu.IM{})
	taxCodes := cu.ToIMA(stateData["tax_codes"], []cu.IM{})
	rate := s.taxRate(taxCodes, cu.ToString(formRow["tax_code"], ""))
	currencies := cu.ToIMA(stateData["currencies"], []cu.IM{})
	digit := uint(s.currencyDigit(currencies, cu.ToString(trans["currency_code"], "")))
	itemRow := cu.ToIM(formRow["item_meta"], cu.IM{})

	var netAmount, vatAmount, amount, fxPrice float64
	switch calcMode {
	case "net_amount":
		netAmount = value
		if cu.ToFloat(itemRow["qty"], 0) != 0 {
			fxPrice = s.roundFloat(netAmount/(1-cu.ToFloat(itemRow["discount"], 0)/100)/cu.ToFloat(itemRow["qty"], 0), digit)
			vatAmount = s.roundFloat(netAmount*rate, digit)
		}
		amount = s.roundFloat(netAmount+vatAmount, digit)

	case "amount":
		amount = value
		if cu.ToFloat(itemRow["qty"], 0) != 0 {
			netAmount = s.roundFloat(amount/(1+rate), digit)
			vatAmount = s.roundFloat(amount-netAmount, digit)
			fxPrice = s.roundFloat(netAmount/(1-cu.ToFloat(itemRow["discount"], 0)/100)/cu.ToFloat(itemRow["qty"], 0), digit)
		}

	case "fx_price":
		fxPrice = value
		netAmount = s.roundFloat(fxPrice*(1-cu.ToFloat(itemRow["discount"], 0)/100)*cu.ToFloat(itemRow["qty"], 0), digit)
		vatAmount = s.roundFloat(fxPrice*(1-cu.ToFloat(itemRow["discount"], 0)/100)*cu.ToFloat(itemRow["qty"], 0)*rate, digit)
		amount = s.roundFloat(netAmount+vatAmount, digit)
	}

	return cu.IM{
		"net_amount": netAmount, "vat_amount": vatAmount, "amount": amount, "fx_price": fxPrice,
	}
}

func (s *TransService) formEventChangeSelector(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := s.cls.getDataStore(client.Ticket.Database)
	trans := cu.ToIM(stateData["trans"], cu.IM{})

	frmValues := cu.ToIM(evt.Value, cu.IM{})
	frmData := cu.ToIM(frmValues["data"], cu.IM{})
	form := cu.ToIM(frmData["form"], cu.IM{})
	frmBaseValues := cu.ToIM(form["data"], cu.IM{})

	fieldName := cu.ToString(frmValues["name"], "")
	itemMeta := cu.ToIM(frmBaseValues["item_meta"], cu.IM{})

	setPriceValues := func(calcMode string, value float64) {
		priceValues := s.calcItemPrice(calcMode, value, stateData, frmBaseValues)
		itemMeta["net_amount"] = priceValues["net_amount"]
		itemMeta["vat_amount"] = priceValues["vat_amount"]
		itemMeta["amount"] = priceValues["amount"]
		itemMeta["fx_price"] = priceValues["fx_price"]
		frmBaseValues["item_meta"] = itemMeta
	}
	switch fieldName {
	case "product_code":
		return s.cls.editorCodeSelector(evt, "trans", strings.Split(fieldName, "_")[0], frmBaseValues,
			func(params cu.IM) (re ct.ResponseEvent, err error) {
				if cu.ToString(params["event"], "") == ct.SelectorEventSelected {
					selectedValues := cu.ToIM(params["values"], cu.IM{})
					frmBaseValues["product_name"] = selectedValues["product_name"]
					frmBaseValues["product_unit"] = selectedValues["unit"]
					if _, found := frmBaseValues["item_meta"]; found {
						itemMeta["unit"] = selectedValues["unit"]
						itemMeta["description"] = selectedValues["product_name"]
						itemMeta["fx_price"], itemMeta["discount"] = s.getProductPrice(ds,
							cu.IM{"currency_code": trans["currency_code"],
								"product_code":  frmBaseValues["product_code"],
								"customer_code": trans["customer_code"],
								"qty":           itemMeta["qty"]})
						frmBaseValues["tax_code"] = cu.ToIM(params["values"], cu.IM{})["tax_code"]
						frmBaseValues["item_meta"] = itemMeta
						setPriceValues("fx_price", cu.ToFloat(itemMeta["fx_price"], 0))
					}
				}
				client.SetForm(cu.ToString(form["key"], ""),
					cu.MergeIM(frmBaseValues,
						cu.IM{"product_selector": stateData["product_selector"]}),
					cu.ToInteger(form["index"], 0), false)
				return evt, nil
			})

	case "tool_code":
		return s.cls.editorCodeSelector(evt, "trans", strings.Split(fieldName, "_")[0], frmBaseValues,
			func(params cu.IM) (re ct.ResponseEvent, err error) {
				if cu.ToString(params["event"], "") == ct.SelectorEventSelected {
					selectedValues := cu.ToIM(params["values"], cu.IM{})
					frmBaseValues["tool_description"] = selectedValues["description"]
					frmBaseValues["serial_number"] = selectedValues["serial_number"]
				}
				client.SetForm(cu.ToString(form["key"], ""),
					cu.MergeIM(frmBaseValues,
						cu.IM{"tool_selector": stateData["tool_selector"]}),
					cu.ToInteger(form["index"], 0), false)
				return evt, nil
			})

	//case "invoice_code":
	default:
		return s.cls.editorCodeSelector(evt, "trans", strings.Split(fieldName, "_")[0], frmBaseValues,
			func(params cu.IM) (re ct.ResponseEvent, err error) {
				if cu.ToString(params["event"], "") == ct.SelectorEventSelected {
					selectedValues := cu.ToIM(params["values"], cu.IM{})
					frmBaseValues["link_code_2"] = selectedValues["code"]
					frmBaseValues["currency_code"] = strings.Split(cu.ToString(selectedValues["currency_code"], ""), "/")[0] + "/" + cu.ToString(selectedValues["currency_code"], "")
					frmBaseValues["invoice_curr"] = selectedValues["currency_code"]
				}
				client.SetForm(cu.ToString(form["key"], ""),
					cu.MergeIM(frmBaseValues,
						cu.IM{"invoice_selector": stateData["invoice_selector"]}),
					cu.ToInteger(form["index"], 0), false)
				return evt, nil
			})
	}
	//return evt, nil
}

func (s *TransService) formEventChange(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()

	frmValues := cu.ToIM(evt.Value, cu.IM{})
	frmData := cu.ToIM(frmValues["data"], cu.IM{})
	form := cu.ToIM(frmData["form"], cu.IM{})
	frmBaseValues := cu.ToIM(form["data"], cu.IM{})

	fieldName := cu.ToString(frmValues["name"], "")
	itemMeta := cu.ToIM(frmBaseValues["item_meta"], cu.IM{})
	linkMeta := cu.ToIM(frmBaseValues["link_meta"], cu.IM{})
	movementMeta := cu.ToIM(frmBaseValues["movement_meta"], cu.IM{})

	setPriceValues := func(calcMode string, value float64) {
		priceValues := s.calcItemPrice(calcMode, value, stateData, frmBaseValues)
		itemMeta["net_amount"] = priceValues["net_amount"]
		itemMeta["vat_amount"] = priceValues["vat_amount"]
		itemMeta["amount"] = priceValues["amount"]
		itemMeta["fx_price"] = priceValues["fx_price"]
		frmBaseValues["item_meta"] = itemMeta
	}
	switch fieldName {
	case "tags":
		return s.cls.editorFormTags(cu.IM{"row_field": fieldName}, evt)
	case "product_code":
		return s.formEventChangeSelector(evt)

	case "tool_code":
		return s.formEventChangeSelector(evt)

	case "invoice_code":
		return s.formEventChangeSelector(evt)

	case "qty", "discount", "tax_code":
		if fieldName != "tax_code" {
			itemMeta[fieldName] = cu.ToFloat(frmValues["value"], 0)
			frmBaseValues["item_meta"] = itemMeta
		} else {
			frmBaseValues[fieldName] = frmValues["value"]
		}
		setPriceValues("fx_price", cu.ToFloat(itemMeta["fx_price"], 0))
		return s.cls.editorCodeSelector(evt, "trans", strings.Split(fieldName, "_")[0], frmBaseValues,
			func(params cu.IM) (re ct.ResponseEvent, err error) {
				client.SetForm(cu.ToString(form["key"], ""), frmBaseValues, cu.ToInteger(form["index"], 0), false)
				return evt, nil
			})

	case "amount", "net_amount", "fx_price":
		setPriceValues(fieldName, cu.ToFloat(frmValues["value"], 0))
		return s.cls.editorCodeSelector(evt, "trans", strings.Split(fieldName, "_")[0], frmBaseValues,
			func(params cu.IM) (re ct.ResponseEvent, err error) {
				client.SetForm(cu.ToString(form["key"], ""), frmBaseValues, cu.ToInteger(form["index"], 0), false)
				return evt, nil
			})

	case "place_code":
		frmBaseValues[fieldName] = frmValues["value"]
		client.SetForm(cu.ToString(form["key"], ""), frmBaseValues, cu.ToInteger(form["index"], 0), false)

	case "own_stock":
		itemMeta[fieldName] = cu.ToFloat(frmValues["value"], 0)
		frmBaseValues["item_meta"] = itemMeta
		cu.ToSM(evt.Header, cu.SM{})[ct.HeaderReswap] = ct.SwapNone

	case "deposit":
		itemMeta[fieldName] = cu.ToBoolean(frmValues["value"], false)
		frmBaseValues["item_meta"] = itemMeta
		cu.ToSM(evt.Header, cu.SM{})[ct.HeaderReswap] = ct.SwapNone

	case "shared":
		movementMeta[fieldName] = cu.ToBoolean(frmValues["value"], false)
		frmBaseValues["movement_meta"] = movementMeta
		client.SetForm(cu.ToString(form["key"], ""), frmBaseValues, cu.ToInteger(form["index"], 0), false)

	case "link_amount", "link_rate":
		linkMeta[strings.Split(fieldName, "_")[1]] = cu.ToFloat(frmValues["value"], 0)
		frmBaseValues["link_meta"] = linkMeta
		client.SetForm(cu.ToString(form["key"], ""), frmBaseValues, cu.ToInteger(form["index"], 0), false)

	default:
		frmBaseValues[fieldName] = frmValues["value"]
		cu.ToSM(evt.Header, cu.SM{})[ct.HeaderReswap] = ct.SwapNone
	}
	return evt, nil
}

func (s *TransService) formEvent(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	trans := cu.ToIM(stateData["trans"], cu.IM{})

	frmValues := cu.ToIM(evt.Value, cu.IM{})
	frmValue := cu.ToIM(frmValues["value"], cu.IM{})
	frmData := cu.ToIM(frmValues["data"], cu.IM{})
	form := cu.ToIM(frmData["form"], cu.IM{})
	frmIndex := cu.ToInteger(form["index"], 0)
	frmKey := cu.ToString(form["key"], "")
	frmEvent := cu.ToString(frmValues["event"], "")
	rows := cu.ToIMA(trans[frmKey], []cu.IM{})
	if srows, found := stateData[frmKey]; found && (len(rows) == 0) {
		rows = cu.ToIMA(srows, []cu.IM{})
	}
	delete := (cu.ToString(frmValue["form_delete"], "") != "")

	resultUpdate := func() (re ct.ResponseEvent, err error) {
		if _, found := trans[frmKey]; found {
			trans[frmKey] = rows
		} else {
			stateData[frmKey] = rows
		}
		stateData["dirty"] = true
		return evt, err
	}

	eventMap := map[string]func() (re ct.ResponseEvent, err error){
		ct.FormEventOK: func() (re ct.ResponseEvent, err error) {
			rowMeta := cu.ToIM(rows[frmIndex][ut.MetaName(rows[frmIndex], "_meta")], cu.IM{})
			//rowMap := cu.ToIM(rows[frmIndex][ut.MetaName(rows[frmIndex], "_map")], cu.IM{})
			customValues := map[string]func(value any){
				"frm_amount": func(value any) {
					rowMeta["amount"] = cu.ToFloat(value, 0)
				},
				"frm_net_amount": func(value any) {
					rowMeta["net_amount"] = cu.ToFloat(value, 0)
				},
				"frm_vat_amount": func(value any) {
					rowMeta["vat_amount"] = cu.ToFloat(value, 0)
				},
				"frm_fx_price": func(value any) {
					rowMeta["fx_price"] = cu.ToFloat(value, 0)
				},
				"frm_qty": func(value any) {
					rowMeta["qty"] = cu.ToFloat(value, 0)
				},
				"frm_discount": func(value any) {
					rowMeta["discount"] = cu.ToFloat(value, 0)
				},
				"frm_own_stock": func(value any) {
					rowMeta["own_stock"] = cu.ToFloat(value, 0)
				},
				"frm_description": func(value any) {
					rowMeta["description"] = value
				},
				"frm_unit": func(value any) {
					rowMeta["unit"] = value
				},
				"frm_notes": func(value any) {
					rowMeta["notes"] = value
				},
				"frm_link_amount": func(value any) {
					rowMeta["amount"] = cu.ToFloat(value, 0)
				},
				"frm_link_rate": func(value any) {
					rowMeta["rate"] = cu.ToFloat(value, 0)
				},
				"frm_shared": func(value any) {
					rowMeta["shared"] = cu.ToBoolean(value, false)
				},

				"base_item_meta": func(value any) {
					rowMeta["tags"] = cu.ToIM(value, cu.IM{})["tags"]
					rowMeta["deposit"] = cu.ToBoolean(cu.ToIM(value, cu.IM{})["deposit"], false)
				},
				"base_product_code": func(value any) {
					rows[frmIndex]["product_code"] = value
				},
				"base_product_name": func(value any) {
					rows[frmIndex]["product_name"] = value
				},
				"base_product_unit": func(value any) {
					rows[frmIndex]["product_unit"] = value
				},
				"base_tool_code": func(value any) {
					rows[frmIndex]["tool_code"] = value
				},
				"base_tool_description": func(value any) {
					rows[frmIndex]["tool_description"] = value
				},
				"base_serial_number": func(value any) {
					rows[frmIndex]["serial_number"] = value
				},
				"base_tax_code": func(value any) {
					rows[frmIndex]["tax_code"] = value
				},
				"base_place_code": func(value any) {
					rows[frmIndex]["place_code"] = value
				},
				"base_tags": func(value any) {
					rows[frmIndex]["tags"] = value
				},
				"base_link_code_2": func(value any) {
					rows[frmIndex]["link_code_2"] = value
				},
				"base_invoice_curr": func(value any) {
					rows[frmIndex]["invoice_curr"] = value
				},
			}
			return s.cls.editorFormOK(evt, rows, customValues)
		},

		ct.FormEventCancel: func() (re ct.ResponseEvent, err error) {
			if delete {
				row := rows[frmIndex]
				rowId := cu.ToInteger(row["id"], 0)
				refCode := cu.ToString(row["movement_code"], "")
				if _, found := stateData[frmKey]; found && rowId > 0 {
					deleteRows := cu.ToIMA(stateData[frmKey+"_delete"], []cu.IM{})
					deleteRows = append(deleteRows, rows[frmIndex])
					if refCode != "" {
						if idx := slices.IndexFunc(rows, func(refRow cu.IM) bool {
							return cu.ToString(refRow["code"], "") == refCode
						}); idx > -1 {
							deleteRows = append(deleteRows, rows[idx])
						}
					}
					stateData[frmKey+"_delete"] = deleteRows
				}
				rows = append(rows[:frmIndex], rows[frmIndex+1:]...)
				return resultUpdate()
			}
			return evt, err
		},

		ct.FormEventChange: func() (re ct.ResponseEvent, err error) {
			return s.formEventChange(evt)
		},
	}

	if len(rows) > 0 && frmIndex < int64(len(rows)) {
		if fn, ok := eventMap[frmEvent]; ok {
			return fn()
		}
	}

	return evt, err
}

func (s *TransService) sideMenu(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()

	view := cu.ToString(stateData["view"], "")
	trans := cu.ToIM(stateData["trans"], cu.IM{})
	transMeta := cu.ToIM(trans["trans_meta"], cu.IM{})
	user := cu.ToIM(stateData["user"], cu.IM{})
	ds := s.cls.getDataStore(client.Ticket.Database)

	menuMap := map[string]func() (re ct.ResponseEvent, err error){
		"editor_save": func() (re ct.ResponseEvent, err error) {
			var transID int64
			if transID, err = s.update(ds, stateData, client.Msg); err != nil {
				return evt, err
			}
			return s.cls.setEditor(evt, "trans", cu.IM{
				"editor_view": view,
				"trans_id":    transID,
				"session_id":  client.Ticket.SessionID,
			}), nil
		},

		"editor_delete": func() (re ct.ResponseEvent, err error) {
			modal := cu.IM{
				"warning_label":   client.Msg("inputbox_delete"),
				"warning_message": "",
				"next":            "editor_delete",
			}
			client.SetForm("warning", modal, 0, true)
			return evt, err
		},

		"editor_cancel": func() (re ct.ResponseEvent, err error) {
			dirty := cu.ToBoolean(stateData["dirty"], false)
			readonly := (cu.ToString(user["user_group"], "") == md.UserGroupGuest.String()) ||
				cu.ToBoolean(trans["deleted"], false) ||
				(cu.ToBoolean(transMeta["closed"], false) && !dirty)
			if dirty && !readonly {
				modal := cu.IM{
					"warning_label":   client.Msg("inputbox_dirty"),
					"warning_message": client.Msg("inputbox_drop"),
					"next":            "editor_cancel",
				}
				client.SetForm("warning", modal, 0, true)
			} else {
				client.ResetEditor()
			}
			return evt, err
		},

		"transitem_new": func() (re ct.ResponseEvent, err error) {
			params := cu.IM{
				"session_id":   client.Ticket.SessionID,
				"trans_type":   trans["trans_type"],
				"direction":    trans["direction"],
				"editor_title": client.Msg(strings.ToLower(cu.ToString(trans["trans_type"], "") + "_new")),
				"editor_icon":  ct.IconFileText,
			}
			return s.cls.setEditor(evt, "trans", params), nil
		},

		"transpayment_new": func() (re ct.ResponseEvent, err error) {
			params := cu.IM{
				"session_id":   client.Ticket.SessionID,
				"trans_type":   trans["trans_type"],
				"direction":    trans["direction"],
				"editor_title": client.Msg(strings.ToLower(cu.ToString(trans["trans_type"], "") + "_new")),
				"editor_icon":  ct.IconMoney,
			}
			return s.cls.setEditor(evt, "trans", params), nil
		},

		"transmovement_new": func() (re ct.ResponseEvent, err error) {
			params := cu.IM{
				"session_id":   client.Ticket.SessionID,
				"trans_type":   trans["trans_type"],
				"direction":    trans["direction"],
				"editor_title": client.Msg(strings.ToLower(cu.ToString(trans["trans_type"], "") + "_new")),
				"editor_icon":  cp.TransTypeIcon(cu.ToString(trans["trans_type"], "")),
			}
			return s.cls.setEditor(evt, "trans", params), nil
		},

		"trans_copy": func() (re ct.ResponseEvent, err error) {
			modal := cu.IM{
				"warning_label":   client.Msg("trans_copy_text"),
				"warning_message": client.Msg("inputbox_delete_info"),
				"next":            "trans_copy",
			}
			client.SetForm("warning", modal, 0, true)
			return evt, err
		},

		"trans_corrective": func() (re ct.ResponseEvent, err error) {
			modal := cu.IM{
				"warning_label":   client.Msg("trans_copy_text"),
				"warning_message": client.Msg("inputbox_delete_info"),
				"next":            "trans_corrective",
			}
			client.SetForm("warning", modal, 0, true)
			return evt, err
		},

		"trans_cancellation": func() (re ct.ResponseEvent, err error) {
			modal := cu.IM{
				"warning_label":   client.Msg("trans_copy_text"),
				"warning_message": client.Msg("inputbox_delete_info"),
				"next":            "trans_cancellation",
			}
			client.SetForm("warning", modal, 0, true)
			return evt, err
		},

		"trans_create": func() (re ct.ResponseEvent, err error) {
			elementCount := cu.ToIMA(stateData["element_count"], []cu.IM{})
			return s.createModal(evt,
				cu.IM{"state_key": trans["trans_type"], "create_direction": trans["direction"],
					"show_delivery": len(elementCount) > 0, "next": "trans_create"}), nil
		},

		"payment_link_add": func() (re ct.ResponseEvent, err error) {
			payments := cu.ToIMA(stateData["payments"], []cu.IM{})
			cashPayment := cu.ToIM(payments[0], cu.IM{})
			cashPaymentMeta := cu.ToIM(cashPayment["payment_meta"], cu.IM{})
			return s.linkAdd(evt,
				cu.IM{"view": "payment_link", "code1": cu.ToString(cashPayment["code"], ""), "code2": "",
					"amount": math.Abs(cu.ToFloat(cashPaymentMeta["amount"], 0))})
		},

		"editor_report": func() (re ct.ResponseEvent, err error) {
			return s.cls.showReportSelector(evt, "TRANS", cu.ToString(trans["code"], ""))
		},

		"editor_bookmark": func() (re ct.ResponseEvent, err error) {
			modal := cu.IM{
				"title":         client.Msg("bookmark_new"),
				"icon":          ct.IconStar,
				"label":         client.Msg("bookmark_enter"),
				"placeholder":   "",
				"field_name":    "value",
				"default_value": "",
				"required":      false,
				"next":          "bookmark_add",
			}
			client.SetForm("input_string", modal, 0, true)
			return evt, nil
		},
	}

	if fn, ok := menuMap[cu.ToString(evt.Value, "")]; ok {
		return fn()
	}

	return evt, err
}

func (s *TransService) editorFieldExternal(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()

	view := cu.ToString(stateData["view"], "")
	trans := cu.ToIM(stateData["trans"], cu.IM{})
	transType := cu.ToString(trans["trans_type"], "")
	direction := cu.ToString(trans["direction"], "")

	var mName string
	var rows []cu.IM = []cu.IM{}
	var row cu.IM = cu.IM{}
	var rowIndex int = 0
	switch transType {
	case md.TransTypeCash.String():
		mName = "payments"
		rows = cu.ToIMA(stateData["payments"], []cu.IM{})
		if len(rows) == 0 {
			rows = append(rows, transRowTypeMap["payments"](stateData))
		}
		row = rows[0]

	case md.TransTypeFormula.String():
		mName = "movements"
		rows = cu.ToIMA(stateData["movements"], []cu.IM{})
		if idx := slices.IndexFunc(rows, func(movement cu.IM) bool {
			return cu.ToString(movement["movement_type"], "") == md.MovementTypeHead.String()
		}); idx > -1 {
			row = rows[idx]
			rowIndex = idx
		} else {
			rows = append(rows, transRowTypeMap["movements"](stateData))
			rowIndex = len(rows) - 1
			row = rows[rowIndex]
		}

	case md.TransTypeProduction.String():
		mName = "movements"
		rows = cu.ToIMA(stateData["movements"], []cu.IM{})
		if idx := slices.IndexFunc(rows, func(movement cu.IM) bool {
			movementMeta := cu.ToIM(movement["movement_meta"], cu.IM{})
			return cu.ToBoolean(movementMeta["shared"], false)
		}); idx > -1 {
			row = rows[idx]
			rowIndex = idx
		} else {
			rows = append(rows, transRowTypeMap["movements"](stateData))
			rowIndex = len(rows) - 1
			row = rows[rowIndex]
			row["movement_type"] = md.MovementTypeInventory.String()
			cu.ToIM(row["movement_meta"], cu.IM{})["shared"] = true
		}

	default:
		return evt, err
	}

	values := cu.ToIM(evt.Value, cu.IM{})
	fieldName := cu.ToString(values["name"], "")
	value := cu.ToString(values["value"], "")

	switch fieldName {
	case "payment_paid_date":
		row["paid_date"] = value
	case "payment_amount":
		rowMeta := cu.ToIM(row["payment_meta"], cu.IM{})
		rowMeta["amount"] = cu.ToFloat(value, 0)
		if direction == md.DirectionOut.String() {
			rowMeta["amount"] = -cu.ToFloat(rowMeta["amount"], 0)
		}
		row["payment_meta"] = rowMeta

	case "movement_product_code":
		valueData := cu.ToIM(values["value"], cu.IM{})
		selectedRow := cu.ToIM(valueData["row"], cu.IM{})
		row["product_code"] = selectedRow["code"]
	case "movement_notes":
		rowMeta := cu.ToIM(row["movement_meta"], cu.IM{})
		rowMeta["notes"] = value
		row["movement_meta"] = rowMeta
	case "movement_qty":
		rowMeta := cu.ToIM(row["movement_meta"], cu.IM{})
		rowMeta["qty"] = cu.ToFloat(value, 0)
		row["movement_meta"] = rowMeta
	}
	rows[rowIndex] = row
	stateData[mName] = rows
	stateData["dirty"] = true
	client.SetEditor("trans", view, stateData)
	return evt, err
}

func (s *TransService) linkAdd(evt ct.ResponseEvent, params cu.IM) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()

	trans := cu.ToIM(stateData["trans"], cu.IM{})
	currencyCode := cu.ToString(trans["currency_code"], "")
	places := cu.ToIMA(stateData["places"], []cu.IM{})
	if idx := slices.IndexFunc(places, func(place cu.IM) bool {
		return cu.ToString(place["code"], "") == cu.ToString(trans["place_code"], "")
	}); idx != -1 {
		currencyCode = cu.ToString(places[idx]["currency_code"], "")
	}

	view := cu.ToString(params["view"], "")
	code1 := cu.ToString(params["code1"], "")
	code2 := cu.ToString(params["code2"], "")

	rows := cu.ToIMA(stateData[view], []cu.IM{})
	index := len(rows)
	row := transRowTypeMap["link"](cu.IM{})
	row["id"] = -(index + 1)
	if code1 != "" {
		row["link_code_1"] = code1
	}
	if code2 != "" {
		row["link_code_2"] = code2
	}
	cu.ToIM(row["link_meta"], cu.IM{})["amount"] = cu.ToFloat(params["amount"], 0)
	rows = append(rows, row)
	stateData[view] = rows
	stateData["view"] = view
	stateData["dirty"] = true
	client.SetForm(view, cu.MergeIM(row, cu.IM{
		"trans_code": code2, "currency_code": currencyCode,
		"amount": cu.ToFloat(params["amount"], 0), "rate": 1,
	}), int64(index), false)
	return evt, err
}

func (s *TransService) editorFieldViewAdd(evt ct.ResponseEvent, transMap cu.IM,
	resultUpdate func(params cu.IM) (re ct.ResponseEvent, err error)) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	view := cu.ToString(stateData["view"], "")
	trans := cu.ToIM(stateData["trans"], cu.IM{})
	transType := cu.ToString(trans["trans_type"], "")

	appendRow := func() (row cu.IM, index int64) {
		rows := cu.ToIMA(stateData[view], []cu.IM{})
		row = transRowTypeMap[view](stateData)
		row["id"] = -(len(rows) + 1)
		rows = append(rows, row)
		stateData[view] = rows
		return row, int64(len(rows) - 1)
	}

	switch view {
	case "items", "payments":
		row, index := appendRow()
		client.SetForm(view,
			cu.MergeIM(row, cu.IM{"tax_codes": stateData["tax_codes"]}),
			index, false)

	case "movements":
		row, index := appendRow()
		switch transType {
		case md.TransTypeDelivery.String():
			// transfer movement
			row["movement_type"] = md.MovementTypeInventory.String()
			row["shipping_time"] = cu.ToString(trans["trans_date"], time.Now().Format("2006-01-02"))[:10] + "T00:00:00"
			movements := cu.ToIMA(stateData["movements"], []cu.IM{})
			if len(movements) > 1 {
				row["place_code"] = cu.ToString(movements[1]["place_code"], "")
			}

		case md.TransTypeWaybill.String():
			row["movement_type"] = md.MovementTypeTool.String()
			row["shipping_time"] = cu.ToString(trans["trans_date"], time.Now().Format("2006-01-02"))[:10] + "T00:00:00"

		case md.TransTypeFormula.String():
			row["movement_type"] = md.MovementTypePlan.String()

		default:
			row["movement_type"] = md.MovementTypeInventory.String()
			row["shipping_time"] = cu.ToString(trans["trans_date"], time.Now().Format("2006-01-02"))[:10] + "T00:00:00"
			row["place_code"] = cu.ToString(trans["place_code"], "")

		}
		client.SetForm(view,
			cu.MergeIM(row, cu.IM{"places": stateData["places"], "trans_type": transType, "index": index}),
			index, false)

	case "tool_movement":
		stateData["params"] = cu.IM{
			"ref_trans_code": cu.ToString(trans["code"], ""),
			"trans_type":     md.TransTypeWaybill.String(),
			"direction":      md.DirectionOut.String(),
			"editor_title":   client.Msg("trans_waybill_new"),
			"editor_icon":    ct.IconBriefcase,
			"session_id":     client.Ticket.SessionID,
		}
		if cu.ToBoolean(stateData["dirty"], false) {
			modal := cu.IM{
				"warning_label":   client.Msg("inputbox_dirty"),
				"warning_message": client.Msg("inputbox_drop"),
				"next":            "trans",
			}
			client.SetForm("warning", modal, 0, true)
		} else {
			return s.cls.setEditor(evt, "trans", stateData["params"].(cu.IM)), nil
		}

	case "transitem_shipping":
		stateData["params"] = cu.IM{
			"trans_code": cu.ToString(trans["code"], ""),
			"session_id": client.Ticket.SessionID}
		if cu.ToBoolean(stateData["dirty"], false) {
			modal := cu.IM{
				"warning_label":   client.Msg("inputbox_dirty"),
				"warning_message": client.Msg("inputbox_drop"),
				"next":            "shipping",
			}
			client.SetForm("warning", modal, 0, true)
		} else {
			return s.cls.setEditor(evt, "shipping", stateData["params"].(cu.IM)), nil
		}

	case "maps":
		return s.cls.addMapField(evt, transMap, resultUpdate)
	}
	return evt, err
}

func (s *TransService) editorField(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := s.cls.getDataStore(client.Ticket.Database)
	view := cu.ToString(stateData["view"], "")

	trans := cu.ToIM(stateData["trans"], cu.IM{})
	transMeta := cu.ToIM(trans["trans_meta"], cu.IM{})
	transMetaWorksheet := cu.ToIM(transMeta["worksheet"], cu.IM{})
	transMetaRent := cu.ToIM(transMeta["rent"], cu.IM{})
	transMap := cu.ToIM(trans["trans_map"], cu.IM{})

	resultUpdate := func(params cu.IM) (re ct.ResponseEvent, err error) {
		trans["trans_meta"] = transMeta
		trans["trans_map"] = transMap
		stateData["trans"] = trans
		if cu.ToBoolean(params["dirty"], false) {
			stateData["dirty"] = true
		}
		client.SetEditor("trans", view, stateData)
		return evt, err
	}

	values := cu.ToIM(evt.Value, cu.IM{})
	fieldName := cu.ToString(values["name"], "")
	value := cu.ToString(values["value"], "")
	valueData := cu.ToIM(values["value"], cu.IM{})
	row := cu.ToIM(valueData["row"], cu.IM{})

	fieldMap := map[string]func() (re ct.ResponseEvent, err error){
		ct.TableEventRowSelected: func() (re ct.ResponseEvent, err error) {
			switch view {
			case "transitem_invoice":
				stateData["params"] = cu.IM{
					"trans_code": cu.ToString(row["trans_code"], ""),
					"session_id": client.Ticket.SessionID}
				if cu.ToBoolean(stateData["dirty"], false) {
					modal := cu.IM{
						"warning_label":   client.Msg("inputbox_dirty"),
						"warning_message": client.Msg("inputbox_drop"),
						"next":            "trans",
					}
					client.SetForm("warning", modal, 0, true)
				} else {
					return s.cls.setEditor(evt, "trans", stateData["params"].(cu.IM)), nil
				}
			default:
				client.SetForm(view,
					cu.MergeIM(row, cu.IM{"tax_codes": stateData["tax_codes"], "places": stateData["places"]}),
					cu.ToInteger(row["index"], cu.ToInteger(valueData["index"], 0)), false)
			}
			return evt, nil
		},

		ct.TableEventAddItem: func() (re ct.ResponseEvent, err error) {
			return s.editorFieldViewAdd(evt, transMap, resultUpdate)

		},

		ct.TableEventFormDelete: func() (re ct.ResponseEvent, err error) {
			fieldName := cu.ToString(row["field_name"], "")
			delete(transMap, fieldName)
			return resultUpdate(cu.IM{"dirty": true})
		},

		ct.TableEventFormUpdate: func() (re ct.ResponseEvent, err error) {
			return s.cls.updateMapField(evt, transMap, resultUpdate)
		},

		ct.TableEventFormChange: func() (re ct.ResponseEvent, err error) {
			return evt, nil
		},

		ct.TableEventFormCancel: func() (re ct.ResponseEvent, err error) {
			return evt, nil
		},

		ct.TableEventEditCell: func() (re ct.ResponseEvent, err error) {
			fieldName := cu.ToString(valueData["fieldname"], "")
			if fieldName == "payment_link_add" {
				return s.linkAdd(evt,
					cu.IM{"view": "payment_link", "code1": cu.ToString(row["code"], ""), "code2": "",
						"amount": cu.ToFloat(row["amount"], 0)})
			}
			module := strings.Split(strings.TrimPrefix(fieldName, "ref_"), "_")[0]
			fieldValue := cu.ToString(valueData["value"], "")
			stateData["params"] = cu.IM{
				"editor_view":    module,
				module + "_code": fieldValue,
				"session_id":     client.Ticket.SessionID,
			}
			if cu.ToBoolean(stateData["dirty"], false) {
				modal := cu.IM{
					"warning_label":   client.Msg("inputbox_dirty"),
					"warning_message": client.Msg("inputbox_drop"),
					"next":            module,
				}
				client.SetForm("warning", modal, 0, true)
				return evt, nil
			}
			return s.cls.setEditor(evt, module, stateData["params"].(cu.IM)), nil
		},

		"map_field": func() (re ct.ResponseEvent, err error) {
			stateData["map_field"] = value
			return resultUpdate(cu.IM{"dirty": false})
		},

		"queue": func() (re ct.ResponseEvent, err error) {
			modal := cu.ToIM(client.Data["modal"], cu.IM{})
			modalData := cu.ToIM(modal["data"], cu.IM{})
			if err = s.cls.insertPrintQueue(ds, modalData); err == nil {
				return s.cls.evtMsg(evt.Name, evt.TriggerName, client.Msg("report_add_queue"), ct.ToastTypeSuccess, 5), nil
			}
			return evt, err
		},

		"tags": func() (re ct.ResponseEvent, err error) {
			return s.cls.editorTags(evt, transMeta, resultUpdate)
		},

		"rate": func() (re ct.ResponseEvent, err error) {
			transMeta[fieldName] = cu.ToFloat(value, 0)
			return resultUpdate(cu.IM{"dirty": true})
		},

		"ref_number": func() (re ct.ResponseEvent, err error) {
			transMeta[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"trans_state": func() (re ct.ResponseEvent, err error) {
			transMeta[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"direction": func() (re ct.ResponseEvent, err error) {
			trans[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"trans_date": func() (re ct.ResponseEvent, err error) {
			trans[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"currency_code": func() (re ct.ResponseEvent, err error) {
			trans[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"due_time": func() (re ct.ResponseEvent, err error) {
			transMeta[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"paid_type": func() (re ct.ResponseEvent, err error) {
			transMeta[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"paid": func() (re ct.ResponseEvent, err error) {
			transMeta[fieldName] = cu.ToBoolean(value, false)
			return resultUpdate(cu.IM{"dirty": true})
		},

		"closed": func() (re ct.ResponseEvent, err error) {
			transMeta[fieldName] = cu.ToBoolean(value, false)
			return resultUpdate(cu.IM{"dirty": true})
		},

		"code": func() (re ct.ResponseEvent, err error) {
			trans[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"place_code": func() (re ct.ResponseEvent, err error) {
			trans[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"target_place_code": func() (re ct.ResponseEvent, err error) {
			movements := cu.ToIMA(stateData["movements"], []cu.IM{})
			for _, movement := range movements {
				if cu.ToString(movement["movement_code"], "") != "" {
					movement["place_code"] = value
				}
			}
			stateData["movements"] = movements
			return resultUpdate(cu.IM{"dirty": true})
		},

		"customer_code": func() (re ct.ResponseEvent, err error) {
			return s.cls.editorCodeSelector(evt, "trans", strings.Split(fieldName, "_")[0], trans, resultUpdate)
		},

		"employee_code": func() (re ct.ResponseEvent, err error) {
			return s.cls.editorCodeSelector(evt, "trans", strings.Split(fieldName, "_")[0], trans, resultUpdate)
		},

		"project_code": func() (re ct.ResponseEvent, err error) {
			return s.cls.editorCodeSelector(evt, "trans", strings.Split(fieldName, "_")[0], trans, resultUpdate)
		},

		"transitem_code": func() (re ct.ResponseEvent, err error) {
			return s.cls.editorCodeSelector(evt, "trans", strings.Split(fieldName, "_")[0], trans, resultUpdate)
		},

		"trans_code": func() (re ct.ResponseEvent, err error) {
			stateData["params"] = cu.IM{
				"trans_code": cu.ToString(value, ""),
				"session_id": client.Ticket.SessionID}
			if cu.ToBoolean(stateData["dirty"], false) {
				modal := cu.IM{
					"warning_label":   client.Msg("inputbox_dirty"),
					"warning_message": client.Msg("inputbox_drop"),
					"next":            "trans",
				}
				client.SetForm("warning", modal, 0, true)
				return evt, nil
			}
			return s.cls.setEditor(evt, "trans", stateData["params"].(cu.IM)), nil
		},

		"notes": func() (re ct.ResponseEvent, err error) {
			transMeta[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"internal_notes": func() (re ct.ResponseEvent, err error) {
			transMeta[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"worksheet_distance": func() (re ct.ResponseEvent, err error) {
			transMetaWorksheet["distance"] = cu.ToFloat(value, 0)
			return resultUpdate(cu.IM{"dirty": true})
		},

		"worksheet_repair": func() (re ct.ResponseEvent, err error) {
			transMetaWorksheet["repair"] = cu.ToFloat(value, 0)
			return resultUpdate(cu.IM{"dirty": true})
		},

		"worksheet_total": func() (re ct.ResponseEvent, err error) {
			transMetaWorksheet["total"] = cu.ToFloat(value, 0)
			return resultUpdate(cu.IM{"dirty": true})
		},

		"worksheet_notes": func() (re ct.ResponseEvent, err error) {
			transMetaWorksheet["notes"] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"rent_holiday": func() (re ct.ResponseEvent, err error) {
			transMetaRent["holiday"] = cu.ToFloat(value, 0)
			return resultUpdate(cu.IM{"dirty": true})
		},

		"rent_bad_tool": func() (re ct.ResponseEvent, err error) {
			transMetaRent["bad_tool"] = cu.ToFloat(value, 0)
			return resultUpdate(cu.IM{"dirty": true})
		},

		"rent_other": func() (re ct.ResponseEvent, err error) {
			transMetaRent["other"] = cu.ToFloat(value, 0)
			return resultUpdate(cu.IM{"dirty": true})
		},

		"rent_notes": func() (re ct.ResponseEvent, err error) {
			transMetaRent["notes"] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"payment_amount": func() (re ct.ResponseEvent, err error) {
			return s.editorFieldExternal(evt)
		},

		"payment_paid_date": func() (re ct.ResponseEvent, err error) {
			return s.editorFieldExternal(evt)
		},

		"movement_product_code": func() (re ct.ResponseEvent, err error) {
			return s.cls.editorCodeSelector(evt, "trans", "product", cu.IM{},
				func(params cu.IM) (re ct.ResponseEvent, err error) {
					event := cu.ToString(params["event"], "")
					if event == ct.SelectorEventSelected || event == ct.SelectorEventDelete {
						return s.editorFieldExternal(evt)
					}
					return resultUpdate(params)
				})
		},

		"movement_qty": func() (re ct.ResponseEvent, err error) {
			return s.editorFieldExternal(evt)
		},

		"movement_notes": func() (re ct.ResponseEvent, err error) {
			return s.editorFieldExternal(evt)
		},
	}

	if fn, ok := fieldMap[fieldName]; ok {
		return fn()
	}
	if slices.Contains([]string{"orientation", "template", "paper_size", "copy",
		"create_netto", "create_ref_number", "create_delivery", "create_direction", "create_trans_type",
	}, fieldName) {
		modal := cu.ToIM(client.Data["modal"], cu.IM{})
		modalData := cu.ToIM(modal["data"], cu.IM{})
		modalData[fieldName] = value
		client.SetForm(cu.ToString(modal["key"], ""), modalData, 0, true)
		return evt, nil
	}
	return evt, nil
}

var createValidate = []func(transType, direction, status string, trans md.Trans, msgFunc func(labelID string) string) (bool, string){
	func(transType, direction, status string, trans md.Trans, msgFunc func(labelID string) string) (bool, string) {
		return (transType == md.TransTypeReceipt.String() || transType == md.TransTypeWorksheet.String()) && direction == md.DirectionIn.String(),
			"invalid_trans"
	},
	func(transType, direction, status string, trans md.Trans, msgFunc func(labelID string) string) (bool, string) {
		return trans.TransMeta.Status.String() == md.TransStatusCancellation.String(), "trans_create_cancellation_err1"
	},
	func(transType, direction, status string, trans md.Trans, msgFunc func(labelID string) string) (bool, string) {
		return status == md.TransStatusCancellation.String() &&
				slices.Contains([]string{md.TransTypeReceipt.String(), md.TransTypeInvoice.String()}, transType) && !trans.Deleted,
			"trans_create_cancellation_err2"
	},
	/*
		func(transType, direction, status string, trans md.Trans, msgFunc func(labelID string) string) (bool, string) {
			return status == md.TransStatusCancellation.String() && trans.TransCode != "", "trans_create_cancellation_err3"
		},
	*/
	func(transType, direction, status string, trans md.Trans, msgFunc func(labelID string) string) (bool, string) {
		return status == md.TransStatusAmendment.String() && trans.Deleted, "trans_create_amendment_err"
	},
}

func (s *TransService) createProductQty(items []cu.IM, productCode string, deposit bool) (tqty float64) {
	for _, bi := range items {
		biMeta := cu.ToIM(bi["item_meta"], cu.IM{})
		if cu.ToString(bi["product_code"], "") == productCode && cu.ToBoolean(biMeta["deposit"], false) == deposit {
			tqty += cu.ToFloat(biMeta["qty"], 0)
		}
	}
	return tqty
}

func (s *TransService) createInvoiceItems(evt ct.ResponseEvent, options cu.IM) (items []cu.IM) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()

	trans := cu.ToIM(stateData["trans"], cu.IM{})
	baseItems := cu.ToIMA(stateData["items"], []cu.IM{})
	transitemInvoice := cu.ToIMA(stateData["transitem_invoice"], []cu.IM{})
	transitemShipping := cu.ToIMA(stateData["transitem_shipping"], []cu.IM{})
	taxCodes := cu.ToIMA(stateData["tax_codes"], []cu.IM{})
	currencies := cu.ToIMA(stateData["currencies"], []cu.IM{})
	digit := uint(s.currencyDigit(currencies, cu.ToString(trans["currency_code"], "")))

	deliveryBase := cu.ToBoolean(options["create_delivery"], false)
	nettoInvoice := cu.ToBoolean(options["create_netto"], false)

	items = []cu.IM{}
	products := map[string]bool{}

	recalcItem := func(item cu.IM) cu.IM {
		rate := s.taxRate(taxCodes, cu.ToString(item["tax_code"], ""))
		itemMeta := cu.ToIM(item["item_meta"], cu.IM{})
		itemMeta["net_amount"] = s.roundFloat(cu.ToFloat(itemMeta["fx_price"], 0)*(1-cu.ToFloat(itemMeta["discount"], 0)/100)*cu.ToFloat(itemMeta["qty"], 0), digit)
		itemMeta["vat_amount"] = s.roundFloat(cu.ToFloat(itemMeta["fx_price"], 0)*(1-cu.ToFloat(itemMeta["discount"], 0)/100)*cu.ToFloat(itemMeta["qty"], 0)*rate, digit)
		itemMeta["amount"] = s.roundFloat(cu.ToFloat(itemMeta["net_amount"], 0)+cu.ToFloat(itemMeta["vat_amount"], 0), digit)
		item["item_meta"] = itemMeta
		return item
	}

	appendItem := func(item cu.IM, iqty float64) {
		if _, found := products[cu.ToString(item["product_code"], "")]; !found {
			iqty -= s.createProductQty(transitemInvoice, cu.ToString(item["product_code"], ""), false)
			products[cu.ToString(item["product_code"], "")] = true
		}
		if iqty != 0 {
			itemMeta := cu.ToIM(item["item_meta"], cu.IM{})
			itemMeta["qty"] = iqty
			item["item_meta"] = itemMeta
			item = recalcItem(item)
			items = append(items, item)
		}
	}

	if deliveryBase {
		// create from order,worksheet and rent, on base the delivery rows
		for _, si := range transitemShipping {
			if idx := slices.IndexFunc(baseItems, func(bi cu.IM) bool {
				return cu.ToString(bi["id"], "")+"-"+cu.ToString(bi["product_code"], "") == cu.ToString(si["id"], "")
			}); idx > -1 {
				bi := baseItems[idx]
				iqty := cu.ToFloat(si["sqty"], 0)
				if cu.ToString(trans["direction"], "") == md.DirectionOut.String() {
					iqty = -iqty
				}
				if iqty > 0 {
					appendItem(cu.MergeIM(cu.IM{}, bi), iqty)
				}
			}
		}
	} else {
		for _, bi := range baseItems {
			if nettoInvoice {
				// create from order,worksheet and rent, on base the invoice rows
				biMeta := cu.ToIM(bi["item_meta"], cu.IM{})
				iqty := cu.ToFloat(biMeta["qty"], 0)
				appendItem(cu.MergeIM(cu.IM{}, bi), iqty)
			} else {
				items = append(items, cu.MergeIM(cu.IM{}, bi))
			}
		}
	}

	// put to deposit rows
	for _, it := range items {
		itMeta := cu.ToIM(it["item_meta"], cu.IM{})
		if cu.ToBoolean(itMeta["deposit"], false) {
			dqty := s.createProductQty(transitemInvoice, cu.ToString(it["product_code"], ""), true)
			if dqty != 0 {
				item := cu.MergeIM(cu.IM{}, it)
				item["qty"] = -dqty
				items = append(items, item)
			}
		}
	}

	return items
}

func (s *TransService) createItems(evt ct.ResponseEvent, options cu.IM, transCode string) (err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := s.cls.getDataStore(client.Ticket.Database)

	trans := cu.ToIM(stateData["trans"], cu.IM{})
	transType := cu.ToString(options["create_trans_type"], cu.ToString(trans["trans_type"], ""))
	//direction := cu.ToString(options["direction"], cu.ToString(trans["direction"], ""))
	status := cu.ToString(options["status"], md.TransStatusNormal.String())

	items := cu.ToIMA(stateData["items"], []cu.IM{})
	if transType == md.TransTypeInvoice.String() || transType == md.TransTypeReceipt.String() {
		items = s.createInvoiceItems(evt, options)
	}

	storno := func(it md.Item) md.Item {
		it.ItemMeta.Qty = -it.ItemMeta.Qty
		it.ItemMeta.NetAmount = -it.ItemMeta.NetAmount
		it.ItemMeta.VatAmount = -it.ItemMeta.VatAmount
		it.ItemMeta.Amount = -it.ItemMeta.Amount
		return it
	}

	updateItem := func(it md.Item) (storeID int64, err error) {
		values := cu.IM{
			"trans_code":   transCode,
			"product_code": it.ProductCode,
			"tax_code":     it.TaxCode,
		}
		ut.ConvertByteToIMData(it.ItemMeta, values, "item_meta")
		ut.ConvertByteToIMData(it.ItemMap, values, "item_map")
		return ds.StoreDataUpdate(md.Update{Values: values, Model: "item"})
	}

	for _, it := range items {
		var item md.Item = md.Item{
			TransCode: transCode,
			ItemMeta: md.ItemMeta{
				Tags: []string{},
			},
			ItemMap: cu.IM{},
		}
		if err = ut.ConvertToType(it, &item); err == nil {
			item.ItemMeta.OwnStock = 0
			if transType == md.TransTypeInvoice.String() || transType == md.TransTypeReceipt.String() {
				item.ItemMeta.Deposit = false
			}
			if status == md.TransStatusCancellation.String() {
				item = storno(item)
			}
			if _, err = updateItem(item); err == nil {
				if status == md.TransStatusAmendment.String() {
					_, err = updateItem(storno(item))
				}
			}
			if err != nil {
				return err
			}
		}
	}
	return err
}

func (s *TransService) createPayments(evt ct.ResponseEvent, options cu.IM, transCode string) (err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := s.cls.getDataStore(client.Ticket.Database)

	status := cu.ToString(options["status"], md.TransStatusNormal.String())
	payments := cu.ToIMA(stateData["payments"], []cu.IM{})

	for _, pm := range payments {
		var payment md.Payment = md.Payment{
			PaidDate: md.TimeDate{Time: time.Now()},
			PaymentMeta: md.PaymentMeta{
				Tags: []string{},
			},
			PaymentMap: cu.IM{},
		}
		if err = ut.ConvertToType(pm, &payment); err == nil {
			values := cu.IM{
				"paid_date":  payment.PaidDate.String(),
				"trans_code": transCode,
			}
			if status == md.TransStatusCancellation.String() {
				payment.PaymentMeta.Amount = -payment.PaymentMeta.Amount
			}
			ut.ConvertByteToIMData(payment.PaymentMeta, values, "payment_meta")
			ut.ConvertByteToIMData(payment.PaymentMap, values, "payment_map")
			if _, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "payment"}); err != nil {
				return err
			}
		}
	}
	return err
}

func (s *TransService) createMovements(evt ct.ResponseEvent, options cu.IM, transCode string) (err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := s.cls.getDataStore(client.Ticket.Database)

	trans := cu.ToIM(stateData["trans"], cu.IM{})
	transType := cu.ToString(options["create_trans_type"], cu.ToString(trans["trans_type"], ""))
	status := cu.ToString(options["status"], md.TransStatusNormal.String())
	movementHead := cu.ToIM(stateData["movement_head"], cu.IM{})
	movements := cu.ToIMA(stateData["movements"], []cu.IM{})

	updateMovement := func(mv md.Movement) (storeID int64, err error) {
		values := cu.IM{
			"movement_type": mv.MovementType.String(),
			"shipping_time": mv.ShippingTime.String(),
			"trans_code":    transCode,
		}
		if mv.MovementType == md.MovementTypeTool {
			values["tool_code"] = mv.ToolCode
		} else {
			values["product_code"] = mv.ProductCode
		}
		if mv.PlaceCode != "" {
			values["place_code"] = mv.PlaceCode
		}
		if status == md.TransStatusCancellation.String() {
			if mv.ItemCode != "" {
				values["item_code"] = mv.ItemCode
			}
			if mv.MovementCode != "" {
				values["movement_code"] = mv.MovementCode
			}
		}
		ut.ConvertByteToIMData(mv.MovementMeta, values, "movement_meta")
		ut.ConvertByteToIMData(mv.MovementMap, values, "movement_map")
		return ds.StoreDataUpdate(md.Update{Values: values, Model: "movement"})
	}

	getMovement := func(movement cu.IM) (mv md.Movement, err error) {
		mv = md.Movement{
			MovementType: md.MovementType(md.MovementTypeInventory),
			ShippingTime: md.TimeDateTime{Time: time.Now()},
			MovementMeta: md.MovementMeta{
				Tags: []string{},
			},
			MovementMap: cu.IM{},
		}
		err = ut.ConvertToType(movement, &mv)
		return mv, err
	}

	var mv md.Movement
	if transType == md.TransTypeFormula.String() || transType == md.TransTypeProduction.String() {
		if mv, err = getMovement(movementHead); err == nil {
			_, err = updateMovement(mv)
		}
		if err != nil {
			return err
		}
	}

	for _, movement := range movements {
		if mv, err = getMovement(movement); err == nil {
			if status == md.TransStatusCancellation.String() {
				mv.MovementMeta.Qty = -mv.MovementMeta.Qty
			}
			if _, err = updateMovement(mv); err != nil {
				return err
			}
		}
	}
	return err
}

func (s *TransService) createTrans(evt ct.ResponseEvent, options cu.IM, trans md.Trans) (transCode string, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := s.cls.getDataStore(client.Ticket.Database)

	user := cu.ToIM(stateData["user"], cu.IM{})
	configData := cu.ToIMA(stateData["config_data"], []cu.IM{})

	transType := cu.ToString(options["create_trans_type"], trans.TransType.String())
	direction := cu.ToString(options["create_direction"], trans.Direction.String())
	status := cu.ToString(options["status"], md.TransStatusNormal.String())

	values := cu.IM{
		"trans_type": transType,
		"trans_date": time.Now().Format(time.DateOnly),
		"direction":  direction,
		"auth_code":  user["code"],
	}

	optionalFields := map[string]func() (bool, any){
		"trans_code": func() (bool, any) {
			return cu.ToBoolean(options["create_ref_number"], false) || (status == md.TransStatusCancellation.String()) ||
				(status == md.TransStatusAmendment.String()), trans.Code
		},
		"customer_code": func() (bool, any) {
			return trans.CustomerCode != "" && transType != md.TransTypeReceipt.String(), trans.CustomerCode
		},
		"employee_code": func() (bool, any) {
			return trans.EmployeeCode != "", trans.EmployeeCode
		},
		"project_code": func() (bool, any) {
			return trans.ProjectCode != "", trans.ProjectCode
		},
		"place_code": func() (bool, any) {
			return trans.PlaceCode != "", trans.PlaceCode
		},
		"currency_code": func() (bool, any) {
			return trans.CurrencyCode != "", trans.CurrencyCode
		},
		"trans_date": func() (bool, any) {
			return status == md.TransStatusCancellation.String(), trans.TransDate.Format(time.DateOnly)
		},
		"deleted": func() (bool, any) {
			return status == md.TransStatusCancellation.String() &&
				transType != md.TransTypeDelivery.String() && transType != md.TransTypeInventory.String(), true
		},
	}

	for key, fn := range optionalFields {
		if ok, value := fn(); ok {
			values[key] = value
		}
	}

	meta := md.TransMeta{
		DueTime:       md.TimeDateTime{Time: time.Now()},
		RefNumber:     "",
		PaidType:      trans.TransMeta.PaidType,
		TaxFree:       trans.TransMeta.TaxFree,
		Paid:          false,
		Rate:          trans.TransMeta.Rate,
		Closed:        false,
		Status:        md.TransStatus(0).Get(status),
		TransState:    md.TransStateOK,
		Notes:         trans.TransMeta.Notes,
		InternalNotes: trans.TransMeta.InternalNotes,
		ReportNotes:   trans.TransMeta.ReportNotes,
		Worksheet:     md.TransMetaWorksheet{},
		Rent:          md.TransMetaRent{},
		Invoice:       md.TransMetaInvoice{},
		Tags:          trans.TransMeta.Tags,
	}

	optionalMeta := []func() (bool, func()){
		func() (bool, func()) {
			return (transType == md.TransTypeWorksheet.String()), func() {
				meta.Worksheet = md.TransMetaWorksheet{}
			}
		},
		func() (bool, func()) {
			return (transType == md.TransTypeRent.String()), func() {
				meta.Rent = md.TransMetaRent{}
			}
		},
		func() (bool, func()) {
			return (transType == md.TransTypeInvoice.String()), func() {
				meta.Invoice = md.TransMetaInvoice{}
				if direction == md.DirectionOut.String() {
					if idx := slices.IndexFunc(configData, func(cf cu.IM) bool {
						return cu.ToString(cf["config_key"], "") == "default_deadline"
					}); idx > -1 {
						meta.DueTime = md.TimeDateTime{Time: time.Now().AddDate(0, 0, int(cu.ToInteger(cu.ToString(configData[idx]["config_value"], ""), 0)))}
					}
				}
			}
		},
		func() (bool, func()) {
			return (status == md.TransStatusCancellation.String()), func() {
				meta.DueTime = trans.TransMeta.DueTime
			}
		},
	}
	for _, fn := range optionalMeta {
		if ok, fn := fn(); ok {
			fn()
		}
	}

	ut.ConvertByteToIMData(meta, values, "trans_meta")
	ut.ConvertByteToIMData(trans.TransMap, values, "trans_map")

	var transID int64
	update := md.Update{Values: values, Model: "trans"}
	if transID, err = ds.StoreDataUpdate(update); err == nil {
		var rows []cu.IM = []cu.IM{}
		if rows, err = ds.StoreDataQuery(md.Query{
			Fields: []string{"*"}, From: "trans",
			Filters: []md.Filter{
				{Field: "id", Comp: "==", Value: transID},
			},
		}, true); err == nil {
			transCode = cu.ToString(cu.ToIM(rows[0], cu.IM{})["code"], "")
		}
	}
	return transCode, err
}

func (s *TransService) createData(evt ct.ResponseEvent, options cu.IM) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()

	var trans md.Trans = md.Trans{}
	ut.ConvertToType(stateData["trans"], &trans)

	// to check some things...
	transType := cu.ToString(options["create_trans_type"], trans.TransType.String())
	direction := cu.ToString(options["create_direction"], trans.Direction.String())
	status := cu.ToString(options["status"], md.TransStatusNormal.String())
	for _, validate := range createValidate {
		if invalid, errMsg := validate(transType, direction, status, trans, client.Msg); invalid {
			return s.cls.errorModal(evt, client.Msg("trans_create_title"), client.Msg(errMsg))
		}
	}

	var transCode string
	if transCode, err = s.createTrans(evt, options, trans); err != nil {
		return s.cls.errorModal(evt, client.Msg("trans_create_title"), client.Msg(err.Error()))
	}

	if err = s.createItems(evt, options, transCode); err != nil {
		return s.cls.errorModal(evt, client.Msg("trans_create_title"), client.Msg(err.Error()))
	}

	if err = s.createPayments(evt, options, transCode); err != nil {
		return s.cls.errorModal(evt, client.Msg("trans_create_title"), client.Msg(err.Error()))
	}

	if err = s.createMovements(evt, options, transCode); err != nil {
		return s.cls.errorModal(evt, client.Msg("trans_create_title"), client.Msg(err.Error()))
	}

	return s.cls.setEditor(evt, "trans", cu.IM{
		"session_id": client.Ticket.SessionID,
		"trans_code": transCode,
	}), nil
}

func (s *TransService) createModal(evt ct.ResponseEvent, data cu.IM) (re ct.ResponseEvent) {
	client := evt.Trigger.(*ct.Client)
	baseTransItem := cu.IM{
		"title": client.Msg("trans_create_title"),
		"icon":  ct.IconFileText,
		"trans_types": []string{
			md.TransTypeOffer.String(), md.TransTypeOrder.String(), md.TransTypeWorksheet.String(),
			md.TransTypeRent.String(), md.TransTypeInvoice.String(), md.TransTypeReceipt.String()},
		"create_trans_type": md.TransTypeOrder.String(),
		"create_direction":  md.DirectionOut.String(),
		"status":            md.TransStatusNormal.String(),
		"show_delivery":     false,
		"next":              "transitem_new",
		"module":            "trans",
		"editor_title":      client.Msg("transitem_title"),
		"editor_icon":       ct.IconFileText,
	}
	dataMap := map[string]func() cu.IM{
		"transitem": func() cu.IM {
			return baseTransItem
		},
		"transpayment": func() cu.IM {
			return cu.IM{
				"title": client.Msg("trans_create_title"),
				"icon":  ct.IconMoney,
				"trans_types": []string{
					md.TransTypeCash.String(), md.TransTypeBank.String()},
				"create_trans_type": md.TransTypeCash.String(),
				"create_direction":  md.DirectionOut.String(),
				"status":            md.TransStatusNormal.String(),
				"next":              "transpayment_new",
				"module":            "trans",
			}
		},
		"transmovement": func() cu.IM {
			return cu.IM{
				"title": client.Msg("transmovement_view"),
				"icon":  ct.IconTruck,
				"trans_types": []string{
					md.TransTypeDelivery.String(), md.TransTypeInventory.String(), md.TransTypeProduction.String()},
				"create_trans_type": md.TransTypeDelivery.String(),
				"create_direction":  md.DirectionTransfer.String(),
				"status":            md.TransStatusNormal.String(),
				"next":              "transmovement_new",
				"module":            "trans",
			}
		},
		md.TransTypeOffer.String(): func() cu.IM {
			return cu.MergeIM(baseTransItem, cu.IM{
				"trans_types": []string{
					md.TransTypeOffer.String(), md.TransTypeOrder.String(), md.TransTypeWorksheet.String(),
					md.TransTypeRent.String()},
				"create_trans_type": md.TransTypeOrder.String(),
				"create_direction":  cu.ToString(data["create_direction"], md.DirectionOut.String()),
				"show_delivery":     false, "create_ref_number": true, "create_netto": true,
				"next": "trans_create",
			})
		},
		md.TransTypeOrder.String(): func() cu.IM {
			return cu.MergeIM(baseTransItem, cu.IM{
				"trans_types": []string{
					md.TransTypeOffer.String(), md.TransTypeOrder.String(), md.TransTypeWorksheet.String(),
					md.TransTypeRent.String(), md.TransTypeInvoice.String(), md.TransTypeReceipt.String()},
				"create_trans_type": md.TransTypeInvoice.String(),
				"create_direction":  cu.ToString(data["create_direction"], md.DirectionOut.String()),
				"show_delivery":     false, "create_ref_number": true, "create_netto": true,
				"next": "trans_create",
			})
		},
		md.TransTypeWorksheet.String(): func() cu.IM {
			return cu.MergeIM(baseTransItem, cu.IM{
				"trans_types": []string{
					md.TransTypeOffer.String(), md.TransTypeOrder.String(), md.TransTypeWorksheet.String(),
					md.TransTypeRent.String(), md.TransTypeInvoice.String(), md.TransTypeReceipt.String()},
				"create_trans_type": md.TransTypeInvoice.String(),
				"create_direction":  cu.ToString(data["create_direction"], md.DirectionOut.String()),
				"show_delivery":     false, "create_ref_number": true, "create_netto": true,
				"next": "trans_create",
			})
		},
		md.TransTypeRent.String(): func() cu.IM {
			return cu.MergeIM(baseTransItem, cu.IM{
				"trans_types": []string{
					md.TransTypeOffer.String(), md.TransTypeOrder.String(), md.TransTypeWorksheet.String(),
					md.TransTypeRent.String(), md.TransTypeInvoice.String(), md.TransTypeReceipt.String()},
				"create_trans_type": md.TransTypeInvoice.String(),
				"create_direction":  cu.ToString(data["create_direction"], md.DirectionOut.String()),
				"show_delivery":     false, "create_ref_number": true, "create_netto": true,
				"next": "trans_create",
			})
		},
		md.TransTypeInvoice.String(): func() cu.IM {
			return cu.MergeIM(baseTransItem, cu.IM{
				"trans_types": []string{
					md.TransTypeOrder.String(), md.TransTypeWorksheet.String(),
					md.TransTypeRent.String(), md.TransTypeInvoice.String(), md.TransTypeReceipt.String()},
				"create_trans_type": md.TransTypeOrder.String(),
				"create_direction":  cu.ToString(data["create_direction"], md.DirectionOut.String()),
				"show_delivery":     false, "create_ref_number": true, "create_netto": false,
				"next": "trans_create",
			})
		},
	}
	client.SetForm("trans_create", cu.MergeIM(dataMap[cu.ToString(data["state_key"], "")](), data), 0, true)
	return evt
}
