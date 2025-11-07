package service

import (
	"fmt"
	"slices"
	"strings"
	"time"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	api "github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

type ProductService struct {
	cls *ClientService
}

func NewProductService(cls *ClientService) *ProductService {
	return &ProductService{
		cls: cls,
	}
}

func (s *ProductService) Data(evt ct.ResponseEvent, params cu.IM) (data cu.IM, err error) {
	client := evt.Trigger.(*ct.Client)
	ds := s.cls.getDataStore(client.Ticket.Database)
	user := cu.ToIM(client.Ticket.User, cu.IM{})

	data = cu.IM{
		"product": cu.IM{
			"product_type": md.ProductType(0),
			"product_meta": cu.IM{
				"barcode_type": md.BarcodeTypeEan13.String(),
			},
		},
		"prices":        cu.IM{},
		"components":    cu.IM{},
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
	}
	var rows []cu.IM = []cu.IM{}
	if cu.ToString(params["product_id"], "") != "" || cu.ToString(params["product_code"], "") != "" {
		var products []cu.IM = []cu.IM{}
		if products, err = ds.StoreDataQuery(md.Query{
			Fields: []string{"*"}, From: "product",
			Filters: []md.Filter{
				{Field: "deleted", Comp: "==", Value: false},
				{Field: "id", Comp: "==", Value: cu.ToInteger(params["product_id"], 0)},
				{Or: true, Field: "code", Comp: "==", Value: cu.ToString(params["product_code"], "")},
			},
		}, false); err != nil {
			return data, err
		}
		if len(products) > 0 {
			data["product"] = products[0]
			data["editor_title"] = cu.ToString(products[0]["code"], "")
		}

		if rows, err = ds.StoreDataQuery(md.Query{
			Fields: []string{"*"}, From: "price",
			Filters: []md.Filter{
				{Field: "product_code", Comp: "==", Value: cu.ToString(products[0]["code"], "")},
				{Field: "deleted", Comp: "==", Value: false},
			},
		}, false); err != nil {
			return data, err
		}
		data["prices"] = rows

		if rows, err = ds.StoreDataQuery(md.Query{
			Fields: []string{"l.*", "p.product_type", "p.product_name", "p.unit"},
			From:   "link l inner join product_view p on l.link_code_2 = p.code",
			Filters: []md.Filter{
				{Field: "l.link_type_1", Comp: "==", Value: "LINK_PRODUCT"},
				{Field: "l.link_type_2", Comp: "==", Value: "LINK_PRODUCT"},
				{Field: "l.link_code_1", Comp: "==", Value: cu.ToString(products[0]["code"], "")},
				{Field: "l.deleted", Comp: "==", Value: false},
			},
			OrderBy: []string{"l.id"},
		}, false); err != nil {
			return data, err
		}
		data["components"] = rows
	}
	product := cu.ToIM(data["product"], cu.IM{})

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
	if cu.ToInteger(product["id"], 0) == 0 {
		if idx := slices.IndexFunc(rows, func(c cu.IM) bool {
			return cu.ToString(c["config_key"], "") == "default_taxcode"
		}); idx > int(-1) {
			product["tax_code"] = cu.ToString(rows[idx]["config_value"], "")
		}
		if idx := slices.IndexFunc(rows, func(c cu.IM) bool {
			return cu.ToString(c["config_key"], "") == "default_unit"
		}); idx > int(-1) {
			productMeta := cu.ToIM(product["product_meta"], cu.IM{})
			productMeta["unit"] = cu.ToString(rows[idx]["config_value"], "")
			product["product_meta"] = productMeta
		}
	}

	if rows, err = ds.StoreDataQuery(md.Query{
		Fields: []string{"id", "report_key", "report_name"}, From: "config_report",
		Filters: []md.Filter{
			{Field: "report_type", Comp: "==", Value: "PRODUCT"},
		},
	}, false); err != nil {
		return data, err
	}
	data["config_report"] = rows

	if rows, err = ds.StoreDataQuery(md.Query{
		Fields: []string{"code", "description"}, From: "tax_view",
		OrderBy: []string{"code"},
	}, false); err != nil {
		return data, err
	}
	data["tax_codes"] = rows

	if rows, err = ds.StoreDataQuery(md.Query{
		Fields: []string{"code", "description"}, From: "currency_view",
		OrderBy: []string{"code"},
	}, false); err != nil {
		return data, err
	}
	data["currencies"] = rows

	return data, err
}

func (s *ProductService) updatePrices(ds *api.DataStore, data cu.IM) (err error) {
	for _, pr := range cu.ToIMA(data["prices"], []cu.IM{}) {
		var price md.Price = md.Price{
			PriceType: md.PriceType(md.PriceTypeCustomer),
			PriceMeta: md.PriceMeta{
				Tags: []string{},
			},
			PriceMap: cu.IM{},
		}
		if err = ut.ConvertToType(pr, &price); err == nil {
			values := cu.IM{
				"valid_from":    price.ValidFrom.Format(time.DateOnly),
				"valid_to":      "",
				"product_code":  price.ProductCode,
				"price_type":    price.PriceType.String(),
				"currency_code": price.CurrencyCode,
				"qty":           price.Qty,
				"customer_code": nil,
			}
			if !price.ValidTo.IsZero() {
				values["valid_to"] = price.ValidTo.Format(time.DateOnly)
			}
			if price.CustomerCode != "" {
				values["customer_code"] = price.CustomerCode
			}
			ut.ConvertByteToIMData(price.PriceMeta, values, "price_meta")
			ut.ConvertByteToIMData(price.PriceMap, values, "price_map")

			priceID := price.Id
			update := md.Update{Values: values, Model: "price"}
			if priceID > 0 {
				update.IDKey = priceID
			}
			if priceID, err = ds.StoreDataUpdate(update); err != nil {
				return err
			}
			pr["id"] = priceID
		}
	}
	return err
}

func (s *ProductService) updateLinks(ds *api.DataStore, data cu.IM) (err error) {
	for _, ln := range cu.ToIMA(data["components"], []cu.IM{}) {
		var link md.Link = md.Link{
			LinkType1: md.LinkType(md.LinkTypeProduct),
			LinkType2: md.LinkType(md.LinkTypeProduct),
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

func (s *ProductService) updateDeleteRows(ds *api.DataStore, data cu.IM) (err error) {
	deleteRows := func(model, deleteKey string) (err error) {
		for _, it := range cu.ToIMA(data[deleteKey], []cu.IM{}) {
			if err = ds.DataDelete(model, cu.ToInteger(it["id"], 0), ""); err != nil {
				return err
			}
		}
		return err
	}

	deleteMap := cu.SM{
		"prices_delete":     "price",
		"components_delete": "link",
	}

	for deleteKey, model := range deleteMap {
		if err = deleteRows(model, deleteKey); err != nil {
			return err
		}
	}
	return err
}

func (s *ProductService) update(ds *api.DataStore, data cu.IM) (productID int64, err error) {
	var product md.Product = md.Product{}
	ut.ConvertToType(data["product"], &product)
	values := cu.IM{
		"product_type": product.ProductType.String(),
		"product_name": product.ProductName,
		"tax_code":     product.TaxCode,
	}
	if product.Code != "" {
		values["code"] = product.Code
	}

	ut.ConvertByteToIMData(product.Events, values, "events")
	ut.ConvertByteToIMData(product.ProductMeta, values, "product_meta")
	ut.ConvertByteToIMData(product.ProductMap, values, "product_map")

	newProduct := (product.Id == 0)
	update := md.Update{Values: values, Model: "product"}
	if !newProduct {
		update.IDKey = product.Id
	}
	if productID, err = ds.StoreDataUpdate(update); err == nil && newProduct {
		var products []cu.IM = []cu.IM{}
		if products, err = ds.StoreDataGet(cu.IM{"id": productID, "model": "product"}, true); err == nil {
			data["product"] = products[0]
			data["editor_title"] = cu.ToString(cu.ToIM(products[0], cu.IM{})["code"], "")
		}
	}

	if err != nil {
		return productID, err
	}

	if err = s.updatePrices(ds, data); err != nil {
		return productID, err
	}

	if err = s.updateLinks(ds, data); err != nil {
		return productID, err
	}

	err = s.updateDeleteRows(ds, data)

	return productID, err
}

func (s *ProductService) delete(ds *api.DataStore, productID int64) (err error) {
	return ds.DataDelete("product", productID, "")
}

func (s *ProductService) formNext(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := s.cls.getDataStore(client.Ticket.Database)
	product := cu.ToIM(stateData["product"], cu.IM{})
	productMeta := cu.ToIM(product["product_meta"], cu.IM{})
	productMap := cu.ToIM(product["product_map"], cu.IM{})

	resultUpdate := func(dirty bool) (re ct.ResponseEvent, err error) {
		product["product_meta"] = productMeta
		product["product_map"] = productMap
		stateData["product"] = product
		if dirty {
			stateData["dirty"] = dirty
		}
		client.SetEditor("product", cu.ToString(stateData["view"], ""), stateData)
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
			if err = s.delete(ds, cu.ToInteger(product["id"], 0)); err != nil {
				return evt, err
			}
			client.ResetEditor()
			return evt, err
		},

		"customer": func() (re ct.ResponseEvent, err error) {
			params := cu.ToIM(stateData["params"], cu.IM{})
			return s.cls.setEditor(evt, "customer", params), nil
		},

		"product": func() (re ct.ResponseEvent, err error) {
			params := cu.ToIM(stateData["params"], cu.IM{})
			return s.cls.setEditor(evt, "product", params), nil
		},

		"editor_add_tag": func() (re ct.ResponseEvent, err error) {
			tag := cu.ToString(frmValue["value"], "")
			if tag != "" {
				tags := ut.ToStringArray(productMeta["tags"])
				if !slices.Contains(tags, tag) {
					tags = append(tags, tag)
					productMeta["tags"] = tags
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
				Key:          "product",
				Code:         cu.ToString(product["code"], ""),
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
			productMap[mapField] = code
			stateData["map_field"] = ""
			return resultUpdate(true)
		},
	}

	if fn, ok := nextMap[cu.ToString(frmData["next"], "")]; ok {
		return fn()
	}
	return evt, err
}

func (s *ProductService) formEvent(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	product := cu.ToIM(stateData["product"], cu.IM{})

	frmValues := cu.ToIM(evt.Value, cu.IM{})
	frmValue := cu.ToIM(frmValues["value"], cu.IM{})
	frmData := cu.ToIM(frmValues["data"], cu.IM{})
	form := cu.ToIM(frmData["form"], cu.IM{})
	frmIndex := cu.ToInteger(form["index"], 0)
	frmKey := cu.ToString(form["key"], "")
	frmBaseValues := cu.ToIM(form["data"], cu.IM{})
	frmEvent := cu.ToString(frmValues["event"], "")
	rows := cu.ToIMA(product[frmKey], []cu.IM{})
	if srows, found := stateData[frmKey]; found && (len(rows) == 0) {
		rows = cu.ToIMA(srows, []cu.IM{})
	}
	delete := (cu.ToString(frmValue["form_delete"], "") != "")

	resultUpdate := func() (re ct.ResponseEvent, err error) {
		if _, found := product[frmKey]; found {
			product[frmKey] = rows
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
				"frm_price_value": func(value any) {
					rowMeta["price_value"] = cu.ToFloat(value, 0)
				},
				"frm_qty": func(value any) {
					rows[frmIndex]["qty"] = cu.ToFloat(value, 0)
				},
				"frm_component_qty": func(value any) {
					rowMeta["qty"] = cu.ToFloat(value, 0)
				},
				"frm_component_notes": func(value any) {
					rowMeta["notes"] = value
				},
				"base_price_meta": func(value any) {
					rowMeta["tags"] = cu.ToIM(value, cu.IM{})["tags"]
				},
				"base_customer_code": func(value any) {
					rows[frmIndex]["customer_code"] = value
				},
				"base_tags": func(value any) {
					rows[frmIndex]["tags"] = value
				},
				"base_link_code_2": func(value any) {
					rows[frmIndex]["link_code_2"] = value
				},
				"base_product_name": func(value any) {
					rows[frmIndex]["product_name"] = value
				},
				"base_unit": func(value any) {
					rows[frmIndex]["unit"] = value
				},
			}
			return s.cls.editorFormOK(evt, rows, customValues)
		},

		ct.FormEventCancel: func() (re ct.ResponseEvent, err error) {
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
			fieldName := cu.ToString(frmValues["name"], "")
			switch fieldName {
			case "tags":
				return s.cls.editorFormTags(cu.IM{"row_field": fieldName}, evt)
			case "customer_code":
				form := cu.ToIM(stateData["form"], cu.IM{})
				formRow := cu.ToIM(form["data"], cu.IM{})
				return s.cls.editorCodeSelector(evt, "product", strings.Split(fieldName, "_")[0], formRow,
					func(params cu.IM) (re ct.ResponseEvent, err error) {
						client.SetForm(cu.ToString(form["key"], ""),
							cu.MergeIM(formRow,
								cu.IM{"currencies": stateData["currencies"], "customer_selector": stateData["customer_selector"]}),
							cu.ToInteger(form["index"], 0), false)
						return evt, nil
					})
			case "product_code":
				form := cu.ToIM(stateData["form"], cu.IM{})
				formRow := cu.ToIM(form["data"], cu.IM{})
				return s.cls.editorCodeSelector(evt, "product", strings.Split(fieldName, "_")[0], formRow,
					func(params cu.IM) (re ct.ResponseEvent, err error) {
						if cu.ToString(params["event"], "") == ct.SelectorEventSelected {
							selectedValues := cu.ToIM(params["values"], cu.IM{})
							frmBaseValues["link_code_2"] = selectedValues["code"]
							frmBaseValues["product_name"] = selectedValues["product_name"]
							frmBaseValues["unit"] = selectedValues["unit"]
						}
						client.SetForm(cu.ToString(form["key"], ""),
							cu.MergeIM(frmBaseValues,
								cu.IM{"product_selector": stateData["product_selector"]}),
							cu.ToInteger(form["index"], 0), false)
						return evt, nil
					})
			case "component_qty", "component_notes":
				cu.ToIM(frmBaseValues["link_meta"], cu.IM{})[strings.Split(fieldName, "_")[1]] = frmValues["value"]
				cu.ToSM(evt.Header, cu.SM{})[ct.HeaderReswap] = ct.SwapNone
			case "price_value":
				cu.ToIM(frmBaseValues["price_meta"], cu.IM{})["price_value"] = frmValues["value"]
				cu.ToSM(evt.Header, cu.SM{})[ct.HeaderReswap] = ct.SwapNone
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

func (s *ProductService) sideMenu(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	view := cu.ToString(stateData["view"], "")
	product := cu.ToIM(stateData["product"], cu.IM{})
	ds := s.cls.getDataStore(client.Ticket.Database)

	menuMap := map[string]func() (re ct.ResponseEvent, err error){
		"editor_save": func() (re ct.ResponseEvent, err error) {
			var productID int64
			if productID, err = s.update(ds, stateData); err != nil {
				return evt, err
			}
			return s.cls.setEditor(evt, "product", cu.IM{
				"editor_view": view,
				"product_id":  productID,
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
			return s.cls.setEditor(evt, "product",
				cu.IM{
					"session_id": client.Ticket.SessionID,
				}), nil
		},

		"editor_report": func() (re ct.ResponseEvent, err error) {
			return s.cls.showReportSelector(evt, "PRODUCT", cu.ToString(product["code"], ""))
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

func (s *ProductService) editorField(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := s.cls.getDataStore(client.Ticket.Database)
	product := cu.ToIM(stateData["product"], cu.IM{})
	productMeta := cu.ToIM(product["product_meta"], cu.IM{})
	productMap := cu.ToIM(product["product_map"], cu.IM{})

	resultUpdate := func(params cu.IM) (re ct.ResponseEvent, err error) {
		product["product_meta"] = productMeta
		product["product_map"] = productMap
		stateData["product"] = product
		if cu.ToBoolean(params["dirty"], false) {
			stateData["dirty"] = true
		}
		client.SetEditor("product", cu.ToString(stateData["view"], ""), stateData)
		return evt, err
	}

	values := cu.ToIM(evt.Value, cu.IM{})
	fieldName := cu.ToString(values["name"], "")
	value := cu.ToString(values["value"], "")
	valueData := cu.ToIM(values["value"], cu.IM{})

	fieldMap := map[string]func() (re ct.ResponseEvent, err error){
		ct.TableEventRowSelected: func() (re ct.ResponseEvent, err error) {
			client.SetForm(cu.ToString(stateData["view"], ""),
				cu.MergeIM(cu.ToIM(valueData["row"], cu.IM{}), cu.IM{"currencies": stateData["currencies"]}),
				cu.ToInteger(valueData["index"], 0), false)
			return evt, nil
		},

		ct.TableEventAddItem: func() (re ct.ResponseEvent, err error) {
			view := cu.ToString(stateData["view"], "")
			typeMap := map[string]func(rc int) cu.IM{
				"prices": func(rc int) cu.IM {
					var price cu.IM
					ut.ConvertToType(md.Price{
						Id:           -int64(rc + 1),
						ProductCode:  cu.ToString(product["code"], ""),
						PriceType:    md.PriceTypeCustomer,
						CurrencyCode: cu.ToString(cu.ToIMA(stateData["currencies"], []cu.IM{})[0]["code"], ""),
						ValidFrom:    md.TimeDate{Time: time.Now()},
						PriceMeta: md.PriceMeta{
							Tags: []string{},
						},
						PriceMap: cu.IM{},
					}, &price)
					return price
				},
				"events": func(rc int) cu.IM {
					var event cu.IM
					ut.ConvertToType(md.Event{
						Tags:     []string{},
						EventMap: cu.IM{},
					}, &event)
					return event
				},
				"components": func(rc int) cu.IM {
					var link cu.IM
					ut.ConvertToType(md.Link{
						Id:        -int64(rc + 1),
						LinkType1: md.LinkTypeProduct,
						LinkType2: md.LinkTypeProduct,
						LinkCode1: cu.ToString(product["code"], ""),
						LinkCode2: "",
						LinkMeta: md.LinkMeta{
							Qty:   1,
							Notes: "",
							Tags:  []string{},
						},
						LinkMap: cu.IM{},
					}, &link)
					return link
				},
			}
			if slices.Contains([]string{"prices", "events", "components"}, view) {
				getBase := func() (base cu.IM) {
					if _, found := product[view]; found {
						return product
					}
					return stateData
				}
				base := getBase()
				rows := cu.ToIMA(base[view], []cu.IM{})
				rows = append(rows, typeMap[view](len(rows)))
				base[view] = rows
				client.SetForm(view,
					cu.MergeIM(typeMap[view](len(rows)),
						cu.IM{"currencies": stateData["currencies"]}),
					cu.ToInteger(len(rows)-1, 0), false)
				return evt, nil
			}
			return s.cls.addMapField(evt, productMap, resultUpdate)
		},

		ct.TableEventFormDelete: func() (re ct.ResponseEvent, err error) {
			row := cu.ToIM(valueData["row"], cu.IM{})
			fieldName := cu.ToString(row["field_name"], "")
			delete(productMap, fieldName)
			return resultUpdate(cu.IM{"dirty": true})
		},

		ct.TableEventFormUpdate: func() (re ct.ResponseEvent, err error) {
			return s.cls.updateMapField(evt, productMap, resultUpdate)
		},

		ct.TableEventFormChange: func() (re ct.ResponseEvent, err error) {
			return evt, nil
		},

		ct.TableEventFormCancel: func() (re ct.ResponseEvent, err error) {
			return evt, nil
		},

		ct.TableEventEditCell: func() (re ct.ResponseEvent, err error) {
			fieldName := cu.ToString(valueData["fieldname"], "")
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
			return s.cls.editorTags(evt, productMeta, resultUpdate)
		},

		"barcode_qty": func() (re ct.ResponseEvent, err error) {
			productMeta[fieldName] = cu.ToFloat(value, 0)
			return resultUpdate(cu.IM{"dirty": true})
		},

		"inactive": func() (re ct.ResponseEvent, err error) {
			productMeta[fieldName] = cu.ToBoolean(value, false)
			return resultUpdate(cu.IM{"dirty": true})
		},

		"code": func() (re ct.ResponseEvent, err error) {
			product[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"product_name": func() (re ct.ResponseEvent, err error) {
			product[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"product_type": func() (re ct.ResponseEvent, err error) {
			product[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"tax_code": func() (re ct.ResponseEvent, err error) {
			product[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"unit": func() (re ct.ResponseEvent, err error) {
			productMeta[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"barcode_type": func() (re ct.ResponseEvent, err error) {
			productMeta[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"barcode": func() (re ct.ResponseEvent, err error) {
			productMeta[fieldName] = value
			return resultUpdate(cu.IM{"dirty": true})
		},

		"notes": func() (re ct.ResponseEvent, err error) {
			productMeta[fieldName] = value
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

func (s *ProductService) Response(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
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
