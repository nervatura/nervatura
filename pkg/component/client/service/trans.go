package service

import (
	"fmt"
	"math"
	"slices"
	"strings"
	"time"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	api "github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

func (cls *ClientService) codeName(ds *api.DataStore, code, model string) (name string) {
	if code != "" {
		if rows, err := ds.StoreDataQuery(md.Query{
			Fields: []string{"*"}, From: model,
			Filters: []md.Filter{
				{Field: "code", Comp: "==", Value: code},
			},
		}, true); err == nil {
			name = cu.ToString(rows[0][model+"_name"], "")
		}
	}
	return name
}

var transDataQueryBase = map[string]func(trans cu.IM) md.Query{
	"items": func(trans cu.IM) md.Query {
		return md.Query{
			Fields: []string{"*"}, From: "item",
			Filters: []md.Filter{
				{Field: "trans_code", Comp: "==", Value: cu.ToString(trans["code"], "")},
				{Field: "deleted", Comp: "==", Value: false},
			},
		}
	},
	"movements": func(trans cu.IM) md.Query {
		return md.Query{
			Fields: []string{"*"}, From: "movement",
			Filters: []md.Filter{
				{Field: "trans_code", Comp: "==", Value: cu.ToString(trans["code"], "")},
				{Field: "deleted", Comp: "==", Value: false},
			},
		}
	},
	"payments": func(trans cu.IM) md.Query {
		return md.Query{
			Fields: []string{"*"}, From: "payment",
			Filters: []md.Filter{
				{Field: "trans_code", Comp: "==", Value: cu.ToString(trans["code"], "")},
				{Field: "deleted", Comp: "==", Value: false},
			},
		}
	},
	"links": func(trans cu.IM) md.Query {
		return md.Query{
			Fields: []string{"*"}, From: "link",
			Filters: []md.Filter{
				{Field: "link_type_1", Comp: "==", Value: md.LinkTypeTrans.String()},
				{Field: "link_code_1", Comp: "==", Value: cu.ToString(trans["code"], "")},
				{Field: "deleted", Comp: "==", Value: false},
			},
		}
	},
	"trans_cancellation": func(trans cu.IM) md.Query {
		return md.Query{
			Fields: []string{"*"}, From: "trans_view",
			Filters: []md.Filter{
				{Field: "status", Comp: "==", Value: md.TransStatusCancellation.String()},
				{Field: "trans_code", Comp: "==", Value: cu.ToString(trans["code"], "")},
				{Field: "deleted", Comp: "==", Value: true},
			},
		}
	},
	"transitem_invoice": func(trans cu.IM) md.Query {
		/*
			return md.Query{
				Fields: []string{"i.*", "t.id as trans_id", "t.trans_date", "t.currency_code"},
				From:   "link l inner join trans t on t.code = l.link_code_1 inner join item i on i.trans_code = t.code",
				Filters: []md.Filter{
					{Field: "l.link_code_2", Comp: "==", Value: cu.ToString(trans["code"], "")},
					{Field: "l.link_type_2", Comp: "==", Value: md.LinkTypeTrans.String()},
					{Field: "l.link_type_1", Comp: "==", Value: md.LinkTypeTrans.String()},
					{Field: "t.trans_type", Comp: "in", Value: fmt.Sprintf("%s,%s",
						md.TransTypeInvoice.String(), md.TransTypeReceipt.String())},
					{Field: "t.direction", Comp: "==", Value: cu.ToString(trans["direction"], "")},
					{Field: "l.deleted", Comp: "==", Value: false},
					{Field: "t.deleted", Comp: "==", Value: false},
					{Field: "i.deleted", Comp: "==", Value: false},
				},
			}
		*/
		return md.Query{
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
	"element_count": func(trans cu.IM) md.Query {
		return md.Query{
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
	"config_report": func(trans cu.IM) md.Query {
		return md.Query{
			Fields: []string{"id", "report_key", "report_name"}, From: "config_report",
			Filters: []md.Filter{
				{Field: "report_type", Comp: "==", Value: "TRANS"},
				{Field: "trans_type", Comp: "==", Value: strings.Split(cu.ToString(trans["trans_type"], ""), "_")[1]},
				{Field: "direction", Comp: "==", Value: trans["direction"]},
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
}

func (cls *ClientService) transDataDefault(trans cu.IM, configData []cu.IM) (data cu.IM) {
	transMeta := cu.ToIM(trans["trans_meta"], cu.IM{})
	isItems := slices.Contains([]string{
		md.TransTypeOffer.String(), md.TransTypeOrder.String(), md.TransTypeWorksheet.String(),
		md.TransTypeRent.String(), md.TransTypeInvoice.String(), md.TransTypeReceipt.String()}, cu.ToString(trans["trans_type"], ""))
	valueMap := map[string]func(value string){
		"default_taxcode": func(value string) {
			trans["tax_code"] = value
		},
		"default_currency": func(value string) {
			if isItems {
				transMeta["currency_code"] = value
			}
		},
		"default_deadline": func(value string) {
			if cu.ToString(trans["trans_type"], "") == md.TransTypeInvoice.String() {
				transMeta["due_time"] = md.TimeDateTime{Time: time.Now().AddDate(0, 0, int(cu.ToInteger(value, 0)))}
			}
			if !slices.Contains([]string{
				md.TransTypeInvoice.String(), md.TransTypeReceipt.String()}, cu.ToString(trans["trans_type"], "")) {
				transMeta["rate"] = cu.ToInteger(value, 0)
			}
		},
		"default_paidtype": func(value string) {
			if isItems {
				transMeta["paid_type"] = value
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

func (cls *ClientService) transData(ds *api.DataStore, user, params cu.IM) (data cu.IM, err error) {
	data = cu.IM{
		"trans": cu.IM{
			"trans_type": cu.ToString(params["trans_type"], md.TransType(0).String()),
			"direction":  cu.ToString(params["direction"], md.Direction(0).String()),
			"trans_meta": cu.IM{
				"status":      md.TransStatusNormal.String(),
				"trans_state": md.TransStateOK.String(),
				"paid_type":   md.PaidTypeOnline.String(),
				"tags":        []string{},
				"worksheet":   cu.IM{},
				"rent":        cu.IM{},
				"invoice":     cu.IM{},
			},
		},
		"items":              []cu.IM{},
		"movements":          []cu.IM{},
		"payments":           []cu.IM{},
		"links":              []cu.IM{},
		"config_map":         []cu.IM{},
		"config_data":        []cu.IM{},
		"config_report":      []cu.IM{},
		"tax_codes":          []cu.IM{},
		"currencies":         []cu.IM{},
		"transitem_invoice":  []cu.IM{},
		"trans_cancellation": []cu.IM{},
		"element_count":      []cu.IM{},
		"user":               user,
		"dirty":              false,
		"editor_icon":        cu.ToString(params["editor_icon"], ct.IconFileText),
		"editor_title":       cu.ToString(params["editor_title"], ""),
		"customer_name":      "",
		"project_name":       "",
		"place_name":         "",
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
			data["editor_title"] = cu.ToString(trans[0]["code"], "")
			data["customer_name"] = cls.codeName(ds, cu.ToString(trans[0]["customer_code"], ""), "customer")
		}

		for key, query := range transDataQueryBase {
			if rows, err = ds.StoreDataQuery(query(trans[0]), false); err != nil {
				return data, err
			}
			data[key] = rows
		}

		if rows, err = ds.StoreDataQuery(md.Query{
			Fields: []string{"*"}, From: "link",
			Filters: []md.Filter{
				{Field: "link_type_2", Comp: "==", Value: md.LinkTypeTrans.String()},
				{Field: "link_code_2", Comp: "==", Value: cu.ToString(trans[0]["code"], "")},
				{Field: "deleted", Comp: "==", Value: false},
			},
		}, false); err != nil {
			return data, err
		}
		data["links"] = append(cu.ToIMA(data["links"], []cu.IM{}), rows...)
	}

	trans := cu.ToIM(data["trans"], cu.IM{})
	for key, query := range transDataQueryExt {
		if rows, err = ds.StoreDataQuery(query, false); err != nil {
			return data, err
		}
		data[key] = rows
	}

	if cu.ToInteger(trans["id"], 0) == 0 {
		data["trans"] = cls.transDataDefault(trans, cu.ToIMA(data["config_data"], []cu.IM{}))
	}

	return data, err
}

func (cls *ClientService) transUpdateItems(ds *api.DataStore, data cu.IM, transCode string) (editor cu.IM, err error) {
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
				return data, err
			}
			it["id"] = itemID
		}
	}
	return data, err
}

func (cls *ClientService) transUpdate(ds *api.DataStore, data cu.IM) (editor cu.IM, err error) {
	user := cu.ToIM(data["user"], cu.IM{})
	var trans md.Trans = md.Trans{}
	ut.ConvertToType(data["trans"], &trans)
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

	var transID int64
	newTrans := (trans.Id == 0)
	update := md.Update{Values: values, Model: "trans"}
	if !newTrans {
		update.IDKey = trans.Id
	} else {
		values["auth_code"] = user["code"]
	}
	if transID, err = ds.StoreDataUpdate(update); err == nil && newTrans {
		var rows []cu.IM = []cu.IM{}
		if rows, err = ds.StoreDataGet(cu.IM{"id": transID, "model": "trans"}, true); err == nil {
			data["trans"] = rows[0]
			trans.Code = cu.ToString(cu.ToIM(rows[0], cu.IM{})["code"], "")
			data["editor_title"] = trans.Code
		}
	}

	if data, err = cls.transUpdateItems(ds, data, trans.Code); err != nil {
		return data, err
	}

	for _, it := range cu.ToIMA(data["items_delete"], []cu.IM{}) {
		if err = ds.DataDelete("item", cu.ToInteger(it["id"], 0), ""); err != nil {
			return data, err
		}
	}

	return data, err
}

func (cls *ClientService) transDelete(ds *api.DataStore, transID int64) (err error) {
	return ds.DataDelete("trans", transID, "")
}

func (cls *ClientService) transResponseFormNext(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := cls.getDataStore(client.Ticket.Database)
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
		client.SetEditor("trans", cu.ToString(stateData["view"], ""), stateData)
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
			if err = cls.transDelete(ds, cu.ToInteger(trans["id"], 0)); err != nil {
				return evt, err
			}
			client.ResetEditor()
			return evt, err
		},

		"trans_new": func() (re ct.ResponseEvent, err error) {
			return cls.setEditor(evt, "trans", cu.IM{
				"session_id":   client.Ticket.SessionID,
				"trans_type":   frmValue["trans_type"],
				"direction":    frmValue["direction"],
				"editor_title": client.Msg("transitem_title"),
				"editor_icon":  ct.IconFileText,
			}), nil
		},

		"trans_create": func() (re ct.ResponseEvent, err error) {
			return cls.transCreateData(evt, frmValue)
		},

		"trans_copy": func() (re ct.ResponseEvent, err error) {
			return cls.transCreateData(evt, cu.IM{
				"status": md.TransStatusNormal.String(),
			})
		},

		"trans_corrective": func() (re ct.ResponseEvent, err error) {
			return cls.transCreateData(evt, cu.IM{
				"status": md.TransStatusAmendment.String(),
			})
		},

		"trans_cancellation": func() (re ct.ResponseEvent, err error) {
			return cls.transCreateData(evt, cu.IM{
				"status": md.TransStatusCancellation.String(),
			})
		},

		"customer": func() (re ct.ResponseEvent, err error) {
			params := cu.ToIM(stateData["params"], cu.IM{})
			return cls.setEditor(evt, "customer", params), nil
		},

		"trans": func() (re ct.ResponseEvent, err error) {
			params := cu.ToIM(stateData["params"], cu.IM{})
			return cls.setEditor(evt, "trans", params), nil
		},

		"employee": func() (re ct.ResponseEvent, err error) {
			params := cu.ToIM(stateData["params"], cu.IM{})
			return cls.setEditor(evt, "employee", params), nil
		},

		"project": func() (re ct.ResponseEvent, err error) {
			params := cu.ToIM(stateData["params"], cu.IM{})
			return cls.setEditor(evt, "project", params), nil
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
			return cls.addBookmark(evt, bookmark), nil
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

func (cls *ClientService) getProductPrice(ds *api.DataStore, options cu.IM) (price float64, discount float64) {
	if results, err := ds.ProductPrice(options); err == nil {
		price = cu.ToFloat(results["price"], 0)
		discount = cu.ToFloat(results["discount"], 0)
	}
	return price, discount
}

func (cls *ClientService) roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func (cls *ClientService) taxRate(taxCodes []cu.IM, taxCode string) (rate float64) {
	if idx := slices.IndexFunc(taxCodes, func(c cu.IM) bool {
		return cu.ToString(c["code"], "") == taxCode
	}); idx > int(-1) {
		return cu.ToFloat(taxCodes[idx]["rate_value"], 0)
	}
	return rate
}

func (cls *ClientService) currencyDigit(currencies []cu.IM, currencyCode string) (digit int64) {
	if idx := slices.IndexFunc(currencies, func(c cu.IM) bool {
		return cu.ToString(c["code"], "") == currencyCode
	}); idx > int(-1) {
		return cu.ToInteger(currencies[idx]["digit"], 0)
	}
	return digit
}

func (cls *ClientService) calcItemPrice(calcMode string, value float64, stateData, formRow cu.IM) cu.IM {
	trans := cu.ToIM(stateData["trans"], cu.IM{})
	taxCodes := cu.ToIMA(stateData["tax_codes"], []cu.IM{})
	rate := cls.taxRate(taxCodes, cu.ToString(formRow["tax_code"], ""))
	currencies := cu.ToIMA(stateData["currencies"], []cu.IM{})
	digit := uint(cls.currencyDigit(currencies, cu.ToString(trans["currency_code"], "")))
	itemRow := cu.ToIM(formRow["item_meta"], cu.IM{})

	var netAmount, vatAmount, amount, fxPrice float64
	switch calcMode {
	case "net_amount":
		netAmount = value
		if cu.ToFloat(itemRow["qty"], 0) != 0 {
			fxPrice = cls.roundFloat(netAmount/(1-cu.ToFloat(itemRow["discount"], 0)/100)/cu.ToFloat(itemRow["qty"], 0), digit)
			vatAmount = cls.roundFloat(netAmount*rate, digit)
		}
		amount = cls.roundFloat(netAmount+vatAmount, digit)

	case "amount":
		amount = value
		if cu.ToFloat(itemRow["qty"], 0) != 0 {
			netAmount = cls.roundFloat(amount/(1+rate), digit)
			vatAmount = cls.roundFloat(amount-netAmount, digit)
			fxPrice = cls.roundFloat(netAmount/(1-cu.ToFloat(itemRow["discount"], 0)/100)/cu.ToFloat(itemRow["qty"], 0), digit)
		}

	case "fx_price":
		fxPrice = value
		netAmount = cls.roundFloat(fxPrice*(1-cu.ToFloat(itemRow["discount"], 0)/100)*cu.ToFloat(itemRow["qty"], 0), digit)
		vatAmount = cls.roundFloat(fxPrice*(1-cu.ToFloat(itemRow["discount"], 0)/100)*cu.ToFloat(itemRow["qty"], 0)*rate, digit)
		amount = cls.roundFloat(netAmount+vatAmount, digit)
	}

	return cu.IM{
		"net_amount": netAmount, "vat_amount": vatAmount, "amount": amount, "fx_price": fxPrice,
	}
}

func (cls *ClientService) transResponseFormEventChange(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := cls.getDataStore(client.Ticket.Database)
	trans := cu.ToIM(stateData["trans"], cu.IM{})

	frmValues := cu.ToIM(evt.Value, cu.IM{})
	frmData := cu.ToIM(frmValues["data"], cu.IM{})
	form := cu.ToIM(frmData["form"], cu.IM{})
	frmBaseValues := cu.ToIM(form["data"], cu.IM{})

	fieldName := cu.ToString(frmValues["name"], "")
	rowMeta := cu.ToIM(frmBaseValues["item_meta"], cu.IM{})
	setPriceValues := func(calcMode string, value float64) {
		priceValues := cls.calcItemPrice(calcMode, value, stateData, frmBaseValues)
		rowMeta["net_amount"] = priceValues["net_amount"]
		rowMeta["vat_amount"] = priceValues["vat_amount"]
		rowMeta["amount"] = priceValues["amount"]
		rowMeta["fx_price"] = priceValues["fx_price"]
		frmBaseValues["item_meta"] = rowMeta
	}
	switch fieldName {
	case "tags":
		return cls.editorFormTags(cu.IM{"row_field": fieldName}, evt)
	case "product_code":
		return cls.editorCodeSelector(evt, strings.Split(fieldName, "_")[0], frmBaseValues,
			func(params cu.IM) (re ct.ResponseEvent, err error) {
				if cu.ToString(params["event"], "") == ct.SelectorEventSelected {
					selectedValues := cu.ToIM(params["values"], cu.IM{})
					rowMeta["unit"] = selectedValues["unit"]
					rowMeta["description"] = selectedValues["product_name"]
					rowMeta["fx_price"], rowMeta["discount"] = cls.getProductPrice(ds,
						cu.IM{"currency_code": trans["currency_code"],
							"product_code":  frmBaseValues["product_code"],
							"customer_code": trans["customer_code"],
							"qty":           rowMeta["qty"]})
					frmBaseValues["tax_code"] = cu.ToIM(params["values"], cu.IM{})["tax_code"]
					frmBaseValues["item_meta"] = rowMeta
					setPriceValues("fx_price", cu.ToFloat(rowMeta["fx_price"], 0))
				}
				client.SetForm(cu.ToString(form["key"], ""),
					cu.MergeIM(frmBaseValues,
						cu.IM{"tax_codes": stateData["tax_codes"], "product_selector": stateData["product_selector"]}),
					cu.ToInteger(form["index"], 0), false)
				return evt, nil
			})

	case "qty", "discount", "tax_code":
		if fieldName != "tax_code" {
			rowMeta[fieldName] = cu.ToFloat(frmValues["value"], 0)
			frmBaseValues["item_meta"] = rowMeta
		} else {
			frmBaseValues[fieldName] = frmValues["value"]
		}
		setPriceValues("fx_price", cu.ToFloat(rowMeta["fx_price"], 0))
		return cls.editorCodeSelector(evt, strings.Split(fieldName, "_")[0], frmBaseValues,
			func(params cu.IM) (re ct.ResponseEvent, err error) {
				client.SetForm(cu.ToString(form["key"], ""),
					cu.MergeIM(frmBaseValues,
						cu.IM{"tax_codes": stateData["tax_codes"], "product_selector": stateData["product_selector"]}),
					cu.ToInteger(form["index"], 0), false)
				return evt, nil
			})

	case "amount", "net_amount", "fx_price":
		setPriceValues(fieldName, cu.ToFloat(frmValues["value"], 0))
		return cls.editorCodeSelector(evt, strings.Split(fieldName, "_")[0], frmBaseValues,
			func(params cu.IM) (re ct.ResponseEvent, err error) {
				client.SetForm(cu.ToString(form["key"], ""),
					cu.MergeIM(frmBaseValues,
						cu.IM{"tax_codes": stateData["tax_codes"], "product_selector": stateData["product_selector"]}),
					cu.ToInteger(form["index"], 0), false)
				return evt, nil
			})

	case "own_stock":
		rowMeta[fieldName] = cu.ToFloat(frmValues["value"], 0)
		frmBaseValues["item_meta"] = rowMeta
		cu.ToSM(evt.Header, cu.SM{})[ct.HeaderReswap] = ct.SwapNone

	case "deposit":
		rowMeta[fieldName] = cu.ToBoolean(frmValues["value"], false)
		frmBaseValues["item_meta"] = rowMeta
		cu.ToSM(evt.Header, cu.SM{})[ct.HeaderReswap] = ct.SwapNone

	default:
		frmBaseValues[fieldName] = frmValues["value"]
		cu.ToSM(evt.Header, cu.SM{})[ct.HeaderReswap] = ct.SwapNone
	}
	return evt, nil
}

func (cls *ClientService) transResponseFormEvent(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	trans := cu.ToIM(stateData["trans"], cu.IM{})

	frmValues := cu.ToIM(evt.Value, cu.IM{})
	frmValue := cu.ToIM(frmValues["value"], cu.IM{})
	frmData := cu.ToIM(frmValues["data"], cu.IM{})
	form := cu.ToIM(frmData["form"], cu.IM{})
	frmIndex := cu.ToInteger(form["index"], 0)
	frmKey := cu.ToString(form["key"], "")
	frmBaseValues := cu.ToIM(form["data"], cu.IM{})
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

				"base_item_meta": func(value any) {
					rowMeta["tags"] = cu.ToIM(value, cu.IM{})["tags"]
					rowMeta["deposit"] = cu.ToBoolean(cu.ToIM(value, cu.IM{})["deposit"], false)
				},
				"base_product_code": func(value any) {
					rows[frmIndex]["product_code"] = value
				},
				"base_tax_code": func(value any) {
					rows[frmIndex]["tax_code"] = value
				},
				"base_tags": func(value any) {
					rows[frmIndex]["tags"] = value
				},
			}
			return cls.editorFormOK(evt, rows, customValues)
		},

		ct.FormEventCancel: func() (re ct.ResponseEvent, err error) {
			if frmKey == "items" && frmBaseValues["product_code"] == "" {
				delete = true
			}
			if delete {
				if _, found := stateData[frmKey]; found {
					deleteRows := cu.ToIMA(stateData[frmKey+"_delete"], []cu.IM{})
					deleteRows = append(deleteRows, rows[frmIndex])
					stateData[frmKey+"_delete"] = deleteRows
				}
				rows = append(rows[:frmIndex], rows[frmIndex+1:]...)
				return resultUpdate()
			}
			return evt, err
		},

		ct.FormEventChange: func() (re ct.ResponseEvent, err error) {
			return cls.transResponseFormEventChange(evt)
		},
	}

	if len(rows) > 0 && frmIndex < int64(len(rows)) {
		if fn, ok := eventMap[frmEvent]; ok {
			return fn()
		}
	}

	return evt, err
}

func (cls *ClientService) transResponseSideMenu(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	trans := cu.ToIM(stateData["trans"], cu.IM{})
	transMeta := cu.ToIM(trans["trans_meta"], cu.IM{})
	user := cu.ToIM(stateData["user"], cu.IM{})
	ds := cls.getDataStore(client.Ticket.Database)

	menuMap := map[string]func() (re ct.ResponseEvent, err error){
		"editor_save": func() (re ct.ResponseEvent, err error) {
			if stateData, err = cls.transUpdate(ds, stateData); err != nil {
				return evt, err
			}
			stateData["dirty"] = false
			client.SetEditor("trans", cu.ToString(stateData["view"], ""), stateData)
			return evt, err
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

		"editor_new": func() (re ct.ResponseEvent, err error) {
			params := cu.IM{
				"session_id":   client.Ticket.SessionID,
				"trans_type":   trans["trans_type"],
				"direction":    trans["direction"],
				"editor_title": client.Msg("transitem_title"),
				"editor_icon":  ct.IconFileText,
			}
			return cls.setEditor(evt, "trans", params), nil
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
			return cls.transCreateModal(evt,
				cu.IM{"state_key": trans["trans_type"], "create_direction": trans["direction"],
					"show_delivery": len(elementCount) > 0, "next": "trans_create"}), nil
		},

		"editor_report": func() (re ct.ResponseEvent, err error) {
			return cls.showReportSelector(evt, "TRANS", cu.ToString(trans["code"], ""))
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

func (cls *ClientService) transResponseEditorField(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := cls.getDataStore(client.Ticket.Database)
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
		client.SetEditor("trans", cu.ToString(stateData["view"], ""), stateData)
		return evt, err
	}

	values := cu.ToIM(evt.Value, cu.IM{})
	fieldName := cu.ToString(values["name"], "")
	value := cu.ToString(values["value"], "")

	fieldMap := map[string]func() (re ct.ResponseEvent, err error){
		ct.TableEventRowSelected: func() (re ct.ResponseEvent, err error) {
			valueData := cu.ToIM(values["value"], cu.IM{})
			row := cu.ToIM(valueData["row"], cu.IM{})
			switch cu.ToString(stateData["view"], "") {
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
					return cls.setEditor(evt, "trans", stateData["params"].(cu.IM)), nil
				}
			default:
				client.SetForm(cu.ToString(stateData["view"], ""),
					cu.MergeIM(cu.ToIM(valueData["row"], cu.IM{}),
						cu.IM{"tax_codes": stateData["tax_codes"], "product_selector": stateData["product_selector"]}),
					cu.ToInteger(valueData["index"], 0), false)
			}
			return evt, nil
		},

		ct.TableEventAddItem: func() (re ct.ResponseEvent, err error) {
			view := cu.ToString(stateData["view"], "")
			typeMap := map[string]func() cu.IM{
				"items": func() cu.IM {
					var item cu.IM
					ut.ConvertToType(md.Item{
						TransCode: cu.ToString(trans["code"], ""),
						TaxCode:   cu.ToString(cu.ToIMA(stateData["tax_codes"], []cu.IM{})[0]["code"], ""),
						ItemMeta: md.ItemMeta{
							Qty:  1,
							Tags: []string{},
						},
						ItemMap: cu.IM{},
					}, &item)
					return item
				},
				"events": func() cu.IM {
					var event cu.IM
					ut.ConvertToType(md.Event{
						Tags:     []string{},
						EventMap: cu.IM{},
					}, &event)
					return event
				},
			}
			if slices.Contains([]string{"items", "events"}, view) {
				getBase := func() (base cu.IM) {
					if _, found := trans[view]; found {
						return trans
					}
					return stateData
				}
				base := getBase()
				rows := cu.ToIMA(base[view], []cu.IM{})
				rows = append(rows, typeMap[view]())
				base[view] = rows
				client.SetForm(view,
					cu.MergeIM(typeMap[view](),
						cu.IM{"tax_codes": stateData["tax_codes"], "product_selector": stateData["product_selector"]}),
					cu.ToInteger(len(rows)-1, 0), false)
				return evt, nil
			}
			return cls.addMapField(evt, transMap, resultUpdate)
		},

		ct.TableEventFormDelete: func() (re ct.ResponseEvent, err error) {
			valueData := cu.ToIM(values["value"], cu.IM{})
			row := cu.ToIM(valueData["row"], cu.IM{})
			fieldName := cu.ToString(row["field_name"], "")
			delete(transMap, fieldName)
			return resultUpdate(cu.IM{"dirty": true})
		},

		ct.TableEventFormUpdate: func() (re ct.ResponseEvent, err error) {
			return cls.updateMapField(evt, transMap, resultUpdate)
		},

		ct.TableEventFormChange: func() (re ct.ResponseEvent, err error) {
			return evt, nil
		},

		ct.TableEventFormCancel: func() (re ct.ResponseEvent, err error) {
			return evt, nil
		},

		"map_field": func() (re ct.ResponseEvent, err error) {
			stateData["map_field"] = value
			return resultUpdate(cu.IM{"dirty": false})
		},

		"queue": func() (re ct.ResponseEvent, err error) {
			modal := cu.ToIM(client.Data["modal"], cu.IM{})
			modalData := cu.ToIM(modal["data"], cu.IM{})
			if err = cls.insertPrintQueue(ds, modalData); err == nil {
				return cls.evtMsg(evt.Name, evt.TriggerName, client.Msg("report_add_queue"), ct.ToastTypeSuccess, 5), nil
			}
			return evt, err
		},

		"tags": func() (re ct.ResponseEvent, err error) {
			return cls.editorTags(evt, transMeta, resultUpdate)
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

		"customer_code": func() (re ct.ResponseEvent, err error) {
			return cls.editorCodeSelector(evt, strings.Split(fieldName, "_")[0], trans, resultUpdate)
		},

		"employee_code": func() (re ct.ResponseEvent, err error) {
			return cls.editorCodeSelector(evt, strings.Split(fieldName, "_")[0], trans, resultUpdate)
		},

		"project_code": func() (re ct.ResponseEvent, err error) {
			return cls.editorCodeSelector(evt, strings.Split(fieldName, "_")[0], trans, resultUpdate)
		},

		"transitem_code": func() (re ct.ResponseEvent, err error) {
			return cls.editorCodeSelector(evt, strings.Split(fieldName, "_")[0], trans, resultUpdate)
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
			return cls.setEditor(evt, "trans", stateData["params"].(cu.IM)), nil
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

func (cls *ClientService) transCreateError(evt ct.ResponseEvent, msg string) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	modal := cu.IM{
		"title":      client.Msg("trans_create_title"),
		"info_label": msg,
		"icon":       ct.IconExclamationTriangle,
	}
	client.SetForm("info", modal, 0, true)
	return evt, nil
}

func (cls *ClientService) transCreateCheck(options cu.IM, trans md.Trans) (errMsg string) {
	transType := cu.ToString(options["create_trans_type"], trans.TransType.String())
	direction := cu.ToString(options["create_direction"], trans.Direction.String())
	status := cu.ToString(options["status"], md.TransStatusNormal.String())

	if (transType == md.TransTypeReceipt.String() || transType == md.TransTypeWorksheet.String()) && direction == md.DirectionIn.String() {
		return "invalid_trans"
	}
	if trans.TransMeta.Status.String() == md.TransStatusCancellation.String() {
		return "trans_create_cancellation_err1"
	}
	if status == md.TransStatusCancellation.String() &&
		slices.Contains([]string{md.TransTypeReceipt.String(), md.TransTypeInvoice.String()}, transType) && !trans.Deleted {
		return "trans_create_cancellation_err2"
	}
	//if status == md.TransStatusCancellation.String() && trans.TransCode != "" {
	//	return "trans_create_cancellation_err3"
	//	}
	if status == md.TransStatusAmendment.String() && trans.Deleted {
		return "trans_create_amendment_err"
	}
	return ""
}

func (cls *ClientService) transCreateProductQty(items []cu.IM, productCode string, deposit bool) (tqty float64) {
	for _, bi := range items {
		biMeta := cu.ToIM(bi["item_meta"], cu.IM{})
		if cu.ToString(bi["product_code"], "") == productCode && cu.ToBoolean(biMeta["deposit"], false) == deposit {
			tqty += cu.ToFloat(biMeta["qty"], 0)
		}
	}
	return tqty
}

func (cls *ClientService) transCreateInvoiceItems(evt ct.ResponseEvent, options cu.IM) (items []cu.IM) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()

	trans := cu.ToIM(stateData["trans"], cu.IM{})
	baseItems := cu.ToIMA(stateData["items"], []cu.IM{})
	transitemInvoice := cu.ToIMA(stateData["transitem_invoice"], []cu.IM{})
	transitemShipping := cu.ToIMA(stateData["transitem_shipping"], []cu.IM{})
	taxCodes := cu.ToIMA(stateData["tax_codes"], []cu.IM{})
	currencies := cu.ToIMA(stateData["currencies"], []cu.IM{})
	digit := uint(cls.currencyDigit(currencies, cu.ToString(trans["currency_code"], "")))

	deliveryBase := cu.ToBoolean(options["create_delivery"], false)
	nettoInvoice := cu.ToBoolean(options["create_netto"], false)

	items = []cu.IM{}
	products := map[string]bool{}

	recalcItem := func(item cu.IM) cu.IM {
		rate := cls.taxRate(taxCodes, cu.ToString(item["tax_code"], ""))
		itemMeta := cu.ToIM(item["item_meta"], cu.IM{})
		itemMeta["net_amount"] = cls.roundFloat(cu.ToFloat(itemMeta["fx_price"], 0)*(1-cu.ToFloat(itemMeta["discount"], 0)/100)*cu.ToFloat(itemMeta["qty"], 0), digit)
		itemMeta["vat_amount"] = cls.roundFloat(cu.ToFloat(itemMeta["fx_price"], 0)*(1-cu.ToFloat(itemMeta["discount"], 0)/100)*cu.ToFloat(itemMeta["qty"], 0)*rate, digit)
		itemMeta["amount"] = cls.roundFloat(cu.ToFloat(itemMeta["net_amount"], 0)+cu.ToFloat(itemMeta["vat_amount"], 0), digit)
		item["item_meta"] = itemMeta
		return item
	}

	appendItem := func(item cu.IM, iqty float64) {
		if _, found := products[cu.ToString(item["product_code"], "")]; !found {
			iqty -= cls.transCreateProductQty(transitemInvoice, cu.ToString(item["product_code"], ""), false)
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
		if cu.ToBoolean(it["deposit"], false) {
			dqty := cls.transCreateProductQty(transitemInvoice, cu.ToString(it["product_code"], ""), true)
			if dqty != 0 {
				item := cu.MergeIM(cu.IM{}, it)
				item["qty"] = -dqty
				items = append(items, item)
			}
		}
	}

	return items
}

func (cls *ClientService) transCreateItems(evt ct.ResponseEvent, options cu.IM, transCode string) (err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := cls.getDataStore(client.Ticket.Database)

	trans := cu.ToIM(stateData["trans"], cu.IM{})
	transType := cu.ToString(options["create_trans_type"], cu.ToString(trans["trans_type"], ""))
	//direction := cu.ToString(options["direction"], cu.ToString(trans["direction"], ""))
	status := cu.ToString(options["status"], md.TransStatusNormal.String())

	items := cu.ToIMA(stateData["items"], []cu.IM{})
	if transType == md.TransTypeInvoice.String() || transType == md.TransTypeReceipt.String() {
		items = cls.transCreateInvoiceItems(evt, options)
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

func (cls *ClientService) transCreatePayments(evt ct.ResponseEvent, options cu.IM, transCode string) (err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := cls.getDataStore(client.Ticket.Database)

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

func (cls *ClientService) transCreateMovements(evt ct.ResponseEvent, options cu.IM, transCode string) (err error) {
	// TODO: movement-item links

	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := cls.getDataStore(client.Ticket.Database)

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

func (cls *ClientService) transCreateTrans(evt ct.ResponseEvent, options cu.IM, trans md.Trans) (transCode string, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := cls.getDataStore(client.Ticket.Database)

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

func (cls *ClientService) transCreateData(evt ct.ResponseEvent, options cu.IM) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	//ds := cls.getDataStore(client.Ticket.Database)

	var trans md.Trans = md.Trans{}
	ut.ConvertToType(stateData["trans"], &trans)

	//transType := cu.ToString(options["trans_type"], trans.TransType.String())
	//direction := cu.ToString(options["direction"], trans.Direction.String())
	//status := cu.ToString(options["status"], md.TransStatusNormal.String())

	// to check some things...
	if errMsg := cls.transCreateCheck(options, trans); errMsg != "" {
		return cls.transCreateError(evt, client.Msg(errMsg))
	}

	var transCode string
	if transCode, err = cls.transCreateTrans(evt, options, trans); err != nil {
		return cls.transCreateError(evt, client.Msg(err.Error()))
	}

	if err = cls.transCreateItems(evt, options, transCode); err != nil {
		return cls.transCreateError(evt, client.Msg(err.Error()))
	}

	if err = cls.transCreatePayments(evt, options, transCode); err != nil {
		return cls.transCreateError(evt, client.Msg(err.Error()))
	}

	if err = cls.transCreateMovements(evt, options, transCode); err != nil {
		return cls.transCreateError(evt, client.Msg(err.Error()))
	}

	return cls.setEditor(evt, "trans", cu.IM{
		"session_id": client.Ticket.SessionID,
		"trans_code": transCode,
	}), nil
}

func (cls *ClientService) transCreateModal(evt ct.ResponseEvent, data cu.IM) (re ct.ResponseEvent) {
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
		"next":              "trans_new",
		"editor_title":      client.Msg("transitem_title"),
		"editor_icon":       ct.IconFileText,
	}
	dataMap := map[string]func() cu.IM{
		"transitem": func() cu.IM {
			return baseTransItem
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

func (cls *ClientService) transResponse(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	switch evt.Name {
	case ct.FormEventOK:
		return cls.transResponseFormNext(evt)

	case ct.ClientEventForm:
		return cls.transResponseFormEvent(evt)

	case ct.ClientEventSideMenu:
		return cls.transResponseSideMenu(evt)

	default:
		return cls.transResponseEditorField(evt)
	}
}
