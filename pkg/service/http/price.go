package http

import (
	"errors"
	"net/http"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

// PricePost - create new price
func PricePost(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)

	// convert request body to struct and schema validation
	var data md.Price = md.Price{
		PriceType: md.PriceType(md.PriceTypeCustomer),
		PriceMeta: md.PriceMeta{
			Tags: []string{},
		},
		PriceMap: cu.IM{},
	}
	err := ds.ConvertFromReader(r.Body, &data)
	if err != nil {
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
	}

	if data.ValidFrom == "" || data.CurrencyCode == "" || data.ProductCode == "" {
		err = errors.New("valid from, currency code and product code are required")
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
	}

	// prepare values for database update
	values := cu.IM{
		"valid_from":    data.ValidFrom,
		"product_code":  data.ProductCode,
		"price_type":    data.PriceType.String(),
		"currency_code": data.CurrencyCode,
		"qty":           data.Qty,
	}
	if data.Code != "" {
		values["code"] = data.Code
	}
	if data.ValidTo != "" {
		values["valid_to"] = data.ValidTo
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
	RespondMessage(w, http.StatusCreated, result, http.StatusUnprocessableEntity, err)
}

// PricePut - update price
func PricePut(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	priceID := cu.ToInteger(r.PathValue("id_code"), 0)
	priceCode := cu.ToString(r.PathValue("id_code"), "")

	var price md.Price
	var inputFields, metaFields []string
	var err error
	if _, inputFields, metaFields, err = ds.GetBodyData("price", r.Body, &price); err == nil {
		_, err = ds.UpdateData(md.UpdateDataOptions{
			Model: "price", IDKey: priceID, Code: priceCode,
			Data: price, Meta: price.PriceMeta, Fields: inputFields, MetaFields: metaFields,
		})
	}
	RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

func PriceDelete(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	priceID := cu.ToInteger(r.PathValue("id_code"), 0)
	priceCode := cu.ToString(r.PathValue("id_code"), "")
	err := ds.DataDelete("price", priceID, priceCode)
	RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// PriceQuery - get prices
func PriceQuery(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)

	var params cu.IM = cu.IM{
		"model":  "price",
		"limit":  cu.ToInteger(r.URL.Query().Get("limit"), 0),
		"offset": cu.ToInteger(r.URL.Query().Get("offset"), 0),
	}
	for _, v := range []string{"price_type", "product_code", "currency_code", "customer_code", "tag"} {
		if r.URL.Query().Get(v) != "" {
			params[v] = cu.ToString(r.URL.Query().Get(v), "")
		}
	}
	response, err := ds.StoreDataGet(params, false)
	RespondMessage(w, http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// PriceGet - get price by id or code
func PriceGet(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	priceID := cu.ToInteger(r.PathValue("id_code"), 0)
	priceCode := cu.ToString(r.PathValue("id_code"), "")
	var err error
	var prices []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if prices, err = ds.GetDataByID("price", priceID, priceCode, true); err == nil {
		response = prices[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	RespondMessage(w, http.StatusOK, response, errStatus, err)
}
