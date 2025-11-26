package http

import (
	"errors"
	"net/http"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

// CurrencyPost - create new currency
func CurrencyPost(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)

	// convert request body to struct and schema validation
	var data md.Currency = md.Currency{
		CurrencyMeta: md.CurrencyMeta{
			Tags: []string{},
		},
		CurrencyMap: cu.IM{},
	}
	err := ds.ConvertFromReader(r.Body, &data)
	if err != nil {
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
	}

	if data.Code == "" {
		err = errors.New("code is required")
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
	}
	// prepare values for database update
	values := cu.IM{
		"code": data.Code,
	}

	ut.ConvertByteToIMData(data.CurrencyMeta, values, "currency_meta")
	ut.ConvertByteToIMData(data.CurrencyMap, values, "currency_map")

	// database insert
	var rows []cu.IM
	var result cu.IM
	var currencyID int64
	if currencyID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "currency"}); err == nil {
		if rows, err = ds.StoreDataGet(cu.IM{"id": currencyID, "model": "currency"}, true); err == nil {
			result = rows[0]
		}
	}
	RespondMessage(w, http.StatusCreated, result, http.StatusUnprocessableEntity, err)
}

// CurrencyPut - update currency
func CurrencyPut(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	currencyID := cu.ToInteger(r.PathValue("id_code"), 0)
	currencyCode := cu.ToString(r.PathValue("id_code"), "")

	var currency md.Currency
	var inputFields, metaFields []string
	var err error
	if _, inputFields, metaFields, err = ds.GetBodyData("currency", r.Body, &currency); err == nil {
		_, err = ds.UpdateData(md.UpdateDataOptions{
			Model: "currency", IDKey: currencyID, Code: currencyCode,
			Data: currency, Meta: currency.CurrencyMeta, Fields: inputFields, MetaFields: metaFields,
		})
	}
	RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

func CurrencyDelete(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	currencyID := cu.ToInteger(r.PathValue("id_code"), 0)
	currencyCode := cu.ToString(r.PathValue("id_code"), "")
	err := ds.DataDelete("currency", currencyID, currencyCode)
	RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// CurrencyQuery - get currencys
func CurrencyQuery(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)

	var params cu.IM = cu.IM{
		"model":  "currency",
		"limit":  cu.ToInteger(r.URL.Query().Get("limit"), 0),
		"offset": cu.ToInteger(r.URL.Query().Get("offset"), 0),
	}
	if r.URL.Query().Get("tag") != "" {
		params["tag"] = cu.ToString(r.URL.Query().Get("tag"), "")
	}
	response, err := ds.StoreDataGet(params, false)
	RespondMessage(w, http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// CurrencyGet - get currency by id or code
func CurrencyGet(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	currencyID := cu.ToInteger(r.PathValue("id_code"), 0)
	currencyCode := cu.ToString(r.PathValue("id_code"), "")
	var err error
	var currencys []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if currencys, err = ds.GetDataByID("currency", currencyID, currencyCode, true); err == nil {
		response = currencys[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	RespondMessage(w, http.StatusOK, response, errStatus, err)
}
