package api

import (
	"errors"
	"slices"
	"time"

	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

// Function - call Nervatura server side functions and services
func (ds *DataStore) Function(functionName string, options cu.IM) (interface{}, error) {
	fnMap := map[string]func(options cu.IM) (interface{}, error){
		"report_install": func(options cu.IM) (interface{}, error) {
			return ds.ReportInstall(cu.ToString(options["report_key"], ""), cu.ToString(options["report_dir"], ""))
		},
		"report_list": func(options cu.IM) (interface{}, error) {
			return ds.ReportList(
				cu.ToString(options["report_dir"], cu.ToString(ds.Config["NT_REPORT_DIR"], "")),
				cu.ToString(options["label"], ""))
		},
		"report_get": func(options cu.IM) (interface{}, error) {
			return ds.GetReport(options)
		},
		"email_send": func(options cu.IM) (interface{}, error) {
			return ds.SendEmail(options)
		},
		"product_price": func(options cu.IM) (interface{}, error) {
			return ds.ProductPrice(options)
		},
		"test": func(options cu.IM) (interface{}, error) {
			return "test", nil
		},
	}
	if fn, ok := fnMap[functionName]; ok {
		return fn(options)
	}
	return []byte{}, errors.New(ut.GetMessage("unknown_method"))
}

// ProductPrice - get product price
func (ds *DataStore) ProductPrice(options cu.IM) (results cu.IM, err error) {
	results = cu.IM{"price": float64(0), "discount": float64(0)}
	for _, v := range []string{"currency_code", "product_code"} {
		if _, found := options[v]; !found {
			return results, errors.New(ut.GetMessage("missing_required_field") + ": " + v)
		}
	}
	priceType := cu.ToString(options["price_type"], md.PriceTypeCustomer.String())
	if !slices.Contains([]string{md.PriceTypeCustomer.String(), md.PriceTypeVendor.String()}, priceType) {
		return results, errors.New(ut.GetMessage("invalid_enum_value") + " (price_type): " + priceType)
	}
	qty := cu.ToFloat(options["qty"], 0)
	posdate := cu.ToString(options["posdate"], time.Now().Format(time.DateOnly))
	tag := cu.ToString(options["tag"], "")

	// the best listprice
	query := md.Query{
		Fields: []string{"min(price_value) as mp"},
		From:   "price_view",
		Filters: []md.Filter{
			{Field: "product_code", Comp: "==", Value: options["product_code"]},
			{Field: "currency_code", Comp: "==", Value: options["currency_code"]},
			{Field: "price_type", Comp: "==", Value: priceType},
			{Field: "price_value", Comp: "!=", Value: float64(0)},
			{Field: "qty", Comp: "<=", Value: qty},
			{Field: "valid_from", Comp: "<=", Value: posdate},
			{Field: "customer_code", Comp: "is", Value: "null"},
			{Field: "valid_to", Comp: "is", Value: "null"},
			{Or: true, Field: "valid_to", Comp: ">=", Value: posdate},
		},
	}
	if tag != "" {
		query.Filters = append(query.Filters, md.Filter{Field: "code", Comp: "in", Value: md.Query{
			Fields: []string{"code"},
			From:   "price_tags",
			Filters: []md.Filter{
				{Field: "tag", Comp: "==", Value: tag},
			},
		}})
	}
	var prices []cu.IM
	if prices, err = ds.StoreDataQuery(query, false); err != nil {
		return results, err
	}
	if len(prices) > 0 {
		results["price"] = cu.ToFloat(prices[0]["mp"], 0)
	}

	if _, found := options["customer_code"]; found {
		// customer discount
		query = md.Query{
			Fields: []string{"*"},
			From:   "customer_view",
			Filters: []md.Filter{
				{Field: "code", Comp: "==", Value: options["customer_code"]},
			},
		}
		var customers []cu.IM
		if customers, err = ds.StoreDataQuery(query, false); err != nil {
			return results, err
		}
		if len(customers) > 0 {
			results["discount"] = cu.ToFloat(customers[0]["discount"], 0)
		}

		// customer price
		query = md.Query{
			Fields: []string{"min(price_value) as mp"},
			From:   "price_view",
			Filters: []md.Filter{
				{Field: "product_code", Comp: "==", Value: options["product_code"]},
				{Field: "currency_code", Comp: "==", Value: options["currency_code"]},
				{Field: "price_type", Comp: "==", Value: priceType},
				{Field: "price_value", Comp: "!=", Value: float64(0)},
				{Field: "qty", Comp: "<=", Value: qty},
				{Field: "valid_from", Comp: "<=", Value: posdate},
				{Field: "customer_code", Comp: "==", Value: options["customer_code"]},
				{Field: "valid_to", Comp: "is", Value: "null"},
				{Or: true, Field: "valid_to", Comp: ">=", Value: posdate},
			},
		}
		if tag != "" {
			query.Filters = append(query.Filters, md.Filter{Field: "code", Comp: "in", Value: md.Query{
				Fields: []string{"code"},
				From:   "price_tags",
				Filters: []md.Filter{
					{Field: "tag", Comp: "==", Value: tag},
				},
			}})
		}
		if prices, err = ds.StoreDataQuery(query, false); err != nil {
			return results, err
		}
		if len(prices) > 0 && prices[0]["mp"] != nil {
			if cu.ToFloat(results["price"], 0) > cu.ToFloat(prices[0]["mp"], 0) {
				results["price"] = cu.ToFloat(prices[0]["mp"], 0)
				results["discount"] = float64(0)
			}
		}
	}

	return results, nil
}
