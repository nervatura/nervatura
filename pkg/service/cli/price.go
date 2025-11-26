package cli

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"time"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

// PriceInsert - create new price
func (cli *CLIService) PriceInsert(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	// convert request body to struct and schema validation
	var data md.Price = md.Price{
		PriceType: md.PriceType(md.PriceTypeCustomer),
		PriceMeta: md.PriceMeta{
			Tags: []string{},
		},
		PriceMap: cu.IM{},
	}
	err := ds.ConvertFromByte([]byte(requestData), &data)
	if err != nil {
		return cli.respondString(http.StatusUnprocessableEntity, nil, http.StatusUnprocessableEntity, err)
	}

	if data.ValidFrom.IsZero() || data.CurrencyCode == "" || data.ProductCode == "" {
		err = errors.New("valid from, currency code and product code are required")
		return cli.respondString(http.StatusUnprocessableEntity, nil, http.StatusUnprocessableEntity, err)
	}

	// prepare values for database update
	values := cu.IM{
		"valid_from":    data.ValidFrom.Format(time.DateOnly),
		"product_code":  data.ProductCode,
		"price_type":    data.PriceType.String(),
		"currency_code": data.CurrencyCode,
		"qty":           data.Qty,
	}
	if data.Code != "" {
		values["code"] = data.Code
	}
	if !data.ValidTo.IsZero() {
		values["valid_to"] = data.ValidTo.Format(time.DateOnly)
	}
	if data.CustomerCode != "" {
		values["customer_code"] = data.CustomerCode
	}

	ut.ConvertByteToIMData(data.PriceMeta, values, "price_meta")
	ut.ConvertByteToIMData(data.PriceMap, values, "price_map")

	// database insert
	var rows []cu.IM
	var result cu.IM
	var priceID int64
	if priceID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "price"}); err == nil {
		if rows, err = ds.StoreDataGet(cu.IM{"id": priceID, "model": "price"}, true); err == nil {
			result = rows[0]
		}
	}
	return cli.respondString(http.StatusCreated, result, http.StatusUnprocessableEntity, err)
}

// PricePut - update price
func (cli *CLIService) PriceUpdate(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	priceID := cu.ToInteger(options["id"], 0)
	priceCode := cu.ToString(options["code"], "")

	var price md.Price
	var inputFields, metaFields []string
	var err error
	reader := io.NopCloser(bytes.NewReader([]byte(requestData)))
	if _, inputFields, metaFields, err = ds.GetBodyData("price", reader, &price); err == nil {
		_, err = ds.UpdateData(md.UpdateDataOptions{
			Model: "price", IDKey: priceID, Code: priceCode,
			Data: price, Meta: price.PriceMeta, Fields: inputFields, MetaFields: metaFields,
		})
	}
	return cli.respondString(http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

func (cli *CLIService) PriceDelete(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	priceID := cu.ToInteger(options["id"], 0)
	priceCode := cu.ToString(options["code"], "")
	err := ds.DataDelete("price", priceID, priceCode)
	return cli.respondString(http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// PriceQuery - get prices
func (cli *CLIService) PriceQuery(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	var params cu.IM = cu.IM{"model": "price"}
	for _, v := range []string{"price_type", "product_code", "currency_code", "customer_code", "tag"} {
		if options[v] != nil {
			params[v] = options[v]
		}
	}
	response, err := ds.StoreDataGet(params, false)
	return cli.respondString(http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// PriceGet - get price by id or code
func (cli *CLIService) PriceGet(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	priceID := cu.ToInteger(options["id"], 0)
	priceCode := cu.ToString(options["code"], "")
	var err error
	var prices []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if prices, err = ds.GetDataByID("price", priceID, priceCode, true); err == nil {
		response = prices[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	return cli.respondString(http.StatusOK, response, errStatus, err)
}
