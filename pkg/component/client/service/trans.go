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

func (cls *ClientService) transData(ds *api.DataStore, user, params cu.IM) (data cu.IM, err error) {
	data = cu.IM{
		"trans": cu.IM{
			"trans_type": md.TransType(0),
			"trans_meta": cu.IM{
				"barcode_type": md.BarcodeTypeEan13.String(),
			},
		},
		"items":         cu.IM{},
		"movements":     cu.IM{},
		"payments":      cu.IM{},
		"links":         cu.IM{},
		"config_map":    cu.IM{},
		"config_data":   cu.IM{},
		"config_report": cu.IM{},
		"tax_codes":     cu.IM{},
		"currencies":    cu.IM{},
		"units":         cu.IM{},
		"user":          user,
		"dirty":         false,
		"editor_icon":   ct.IconShoppingCart,
		"editor_title":  "",
		"customer_name": "",
		"project_name":  "",
		"place_name":    "",
	}
	var rows []cu.IM = []cu.IM{}
	if cu.ToString(params["trans_id"], "") != "" || cu.ToString(params["trans_code"], "") != "" {
		var trans []cu.IM = []cu.IM{}
		if trans, err = ds.StoreDataQuery(md.Query{
			Fields: []string{"*"}, From: "trans",
			Filters: []md.Filter{
				{Field: "deleted", Comp: "==", Value: false},
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

		if rows, err = ds.StoreDataQuery(md.Query{
			Fields: []string{"*"}, From: "item",
			Filters: []md.Filter{
				{Field: "trans_code", Comp: "==", Value: cu.ToString(trans[0]["code"], "")},
				{Field: "deleted", Comp: "==", Value: false},
			},
		}, false); err != nil {
			return data, err
		}
		data["items"] = rows

		if rows, err = ds.StoreDataQuery(md.Query{
			Fields: []string{"*"}, From: "movement",
			Filters: []md.Filter{
				{Field: "trans_code", Comp: "==", Value: cu.ToString(trans[0]["code"], "")},
				{Field: "deleted", Comp: "==", Value: false},
			},
		}, false); err != nil {
			return data, err
		}
		data["movements"] = rows

		if rows, err = ds.StoreDataQuery(md.Query{
			Fields: []string{"*"}, From: "payment",
			Filters: []md.Filter{
				{Field: "trans_code", Comp: "==", Value: cu.ToString(trans[0]["code"], "")},
				{Field: "deleted", Comp: "==", Value: false},
			},
		}, false); err != nil {
			return data, err
		}
		data["payments"] = rows

		if rows, err = ds.StoreDataQuery(md.Query{
			Fields: []string{"*"}, From: "link",
			Filters: []md.Filter{
				{Field: "link_type_1", Comp: "==", Value: md.LinkTypeTrans.String()},
				{Field: "link_code_1", Comp: "==", Value: cu.ToString(trans[0]["code"], "")},
				{Field: "deleted", Comp: "==", Value: false},
			},
		}, false); err != nil {
			return data, err
		}
		data["links"] = rows

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

	if rows, err = ds.StoreDataQuery(md.Query{
		Fields: []string{"*"}, From: "config_map",
	}, false); err != nil {
		return data, err
	}
	data["config_map"] = rows

	if rows, err = ds.StoreDataQuery(md.Query{
		Fields: []string{"*"}, From: "config_data",
	}, false); err != nil {
		return data, err
	}
	data["config_data"] = rows
	if cu.ToInteger(trans["id"], 0) == 0 {
		if idx := slices.IndexFunc(rows, func(c cu.IM) bool {
			return cu.ToString(c["config_key"], "") == "default_taxcode"
		}); idx > int(-1) {
			trans["tax_code"] = cu.ToString(rows[idx]["config_value"], "")
		}
		if idx := slices.IndexFunc(rows, func(c cu.IM) bool {
			return cu.ToString(c["config_key"], "") == "default_unit"
		}); idx > int(-1) {
			transMeta := cu.ToIM(trans["trans_meta"], cu.IM{})
			transMeta["unit"] = cu.ToString(rows[idx]["config_value"], "")
			trans["trans_meta"] = transMeta
		}
	}

	if rows, err = ds.StoreDataQuery(md.Query{
		Fields: []string{"id", "report_key", "report_name"}, From: "config_report",
		Filters: []md.Filter{
			{Field: "report_type", Comp: "==", Value: "TRANS"},
			{Field: "trans_type", Comp: "==", Value: strings.Split(cu.ToString(trans["trans_type"], ""), "_")[1]},
			{Field: "direction", Comp: "==", Value: trans["direction"]},
		},
	}, false); err != nil {
		return data, err
	}
	data["config_report"] = rows

	if rows, err = ds.StoreDataQuery(md.Query{
		Fields: []string{"code", "description", "rate_value"}, From: "tax_view",
		OrderBy: []string{"code"},
	}, false); err != nil {
		return data, err
	}
	data["tax_codes"] = rows

	if rows, err = ds.StoreDataQuery(md.Query{
		Fields: []string{"code", "description", "digit"}, From: "currency_view",
		OrderBy: []string{"code"},
	}, false); err != nil {
		return data, err
	}
	data["currencies"] = rows

	return data, err
}

func (cls *ClientService) transUpdate(ds *api.DataStore, data cu.IM) (editor cu.IM, err error) {
	var trans md.Trans = md.Trans{}
	ut.ConvertToType(data["trans"], &trans)
	values := cu.IM{
		"trans_type": trans.TransType.String(),
	}
	if trans.Code != "" {
		values["code"] = trans.Code
	}

	ut.ConvertByteToIMData(trans.TransMeta, values, "trans_meta")
	ut.ConvertByteToIMData(trans.TransMap, values, "trans_map")

	var transID int64
	newTrans := (trans.Id == 0)
	update := md.Update{Values: values, Model: "trans"}
	if !newTrans {
		update.IDKey = trans.Id
	}
	if transID, err = ds.StoreDataUpdate(update); err == nil && newTrans {
		var rows []cu.IM = []cu.IM{}
		if rows, err = ds.StoreDataGet(cu.IM{"id": transID, "model": "trans"}, true); err == nil {
			data["trans"] = rows[0]
			trans.Code = cu.ToString(cu.ToIM(rows[0], cu.IM{})["code"], "")
			data["editor_title"] = trans.Code
		}
	}

	for _, it := range cu.ToIMA(data["items"], []cu.IM{}) {
		var item md.Item = md.Item{
			TransCode: trans.Code,
			ItemMeta: md.ItemMeta{
				Tags: []string{},
			},
			ItemMap: cu.IM{},
		}
		if err = ut.ConvertToType(it, &item); err == nil {
			values = cu.IM{
				"trans_code":   trans.Code,
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
			if tag != "" {
				tags := ut.ToStringArray(row["tags"])
				if metaName != "" {
					tags = ut.ToStringArray(cu.ToIM(row[metaName], cu.IM{})["tags"])
				}
				if !slices.Contains(tags, tag) {
					tags = append(tags, tag)
					if metaName != "" {
						cu.ToIM(row[metaName], cu.IM{})["tags"] = tags
					} else {
						row["tags"] = tags
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

func (cls *ClientService) calcItemPrice(calcMode string, value float64, stateData, formRow cu.IM) cu.IM {
	roundFloat := func(val float64, precision uint) float64 {
		ratio := math.Pow(10, float64(precision))
		return math.Round(val*ratio) / ratio
	}

	trans := cu.ToIM(stateData["trans"], cu.IM{})
	taxCodes := cu.ToIMA(stateData["tax_codes"], []cu.IM{})
	rate := float64(0)
	if idx := slices.IndexFunc(taxCodes, func(c cu.IM) bool {
		return cu.ToString(c["code"], "") == cu.ToString(formRow["tax_code"], "")
	}); idx > int(-1) {
		rate = cu.ToFloat(taxCodes[idx]["rate_value"], 0)
	}
	currencies := cu.ToIMA(stateData["currencies"], []cu.IM{})
	digit := uint(0)
	if idx := slices.IndexFunc(currencies, func(c cu.IM) bool {
		return cu.ToString(c["code"], "") == cu.ToString(trans["currency_code"], "")
	}); idx > int(-1) {
		digit = uint(cu.ToInteger(currencies[idx]["digit"], 0))
	}
	itemRow := cu.ToIM(formRow["item_meta"], cu.IM{})

	var netAmount, vatAmount, amount, fxPrice float64
	switch calcMode {
	case "net_amount":
		netAmount = value
		if cu.ToFloat(itemRow["qty"], 0) != 0 {
			fxPrice = roundFloat(netAmount/(1-cu.ToFloat(itemRow["discount"], 0)/100)/cu.ToFloat(itemRow["qty"], 0), digit)
			vatAmount = roundFloat(netAmount*rate, digit)
		}
		amount = roundFloat(netAmount+vatAmount, digit)

	case "amount":
		amount = value
		if cu.ToFloat(itemRow["qty"], 0) != 0 {
			netAmount = roundFloat(amount/(1+rate), digit)
			vatAmount = roundFloat(amount-netAmount, digit)
			fxPrice = roundFloat(netAmount/(1-cu.ToFloat(itemRow["discount"], 0)/100)/cu.ToFloat(itemRow["qty"], 0), digit)
		}

	case "fx_price":
		fxPrice = value
		netAmount = roundFloat(fxPrice*(1-cu.ToFloat(itemRow["discount"], 0)/100)*cu.ToFloat(itemRow["qty"], 0), digit)
		vatAmount = roundFloat(fxPrice*(1-cu.ToFloat(itemRow["discount"], 0)/100)*cu.ToFloat(itemRow["qty"], 0)*rate, digit)
		amount = roundFloat(netAmount+vatAmount, digit)
	}

	return cu.IM{
		"net_amount": netAmount, "vat_amount": vatAmount, "amount": amount, "fx_price": fxPrice,
	}
}

func (cls *ClientService) transResponseFormEvent(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := cls.getDataStore(client.Ticket.Database)
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
			form := cu.ToIM(stateData["form"], cu.IM{})
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
				return cls.editorFormTags(evt)
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

			//case "price_value":
			//	cu.ToIM(frmBaseValues["price_meta"], cu.IM{})["price_value"] = frmValues["value"]
			//	cu.ToSM(evt.Header, cu.SM{})[ct.HeaderReswap] = ct.SwapNone

			default:
				frmBaseValues[fieldName] = frmValues["value"]
				cu.ToSM(evt.Header, cu.SM{})[ct.HeaderReswap] = ct.SwapNone
			}
			return evt, nil
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
			if cu.ToBoolean(stateData["dirty"], false) {
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
			return cls.setEditor(evt, "trans",
				cu.IM{
					"session_id": client.Ticket.SessionID,
				}), nil
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
			client.SetForm(cu.ToString(stateData["view"], ""),
				cu.MergeIM(cu.ToIM(valueData["row"], cu.IM{}),
					cu.IM{"tax_codes": stateData["tax_codes"], "product_selector": stateData["product_selector"]}),
				cu.ToInteger(valueData["index"], 0), false)
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
	if slices.Contains([]string{"orientation", "template", "paper_size", "copy"}, fieldName) {
		modal := cu.ToIM(client.Data["modal"], cu.IM{})
		modalData := cu.ToIM(modal["data"], cu.IM{})
		modalData[fieldName] = value
		client.SetForm("report", modalData, 0, true)
		return evt, nil
	}
	return evt, nil
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
