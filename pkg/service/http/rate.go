package http

import (
	"errors"
	"net/http"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

// RatePost - create new rate
func RatePost(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)

	// convert request body to struct and schema validation
	var data md.Rate = md.Rate{
		RateType: md.RateType(md.RateTypeRate),
		RateMeta: md.RateMeta{
			Tags: []string{},
		},
		RateMap: cu.IM{},
	}
	err := ds.ConvertFromReader(r.Body, &data)
	if err != nil {
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
	}

	if data.RateDate == "" || data.PlaceCode == "" || data.CurrencyCode == "" {
		err = errors.New("rate date, place code and currency code are required")
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
	}

	// prepare values for database update
	values := cu.IM{
		"rate_type":     data.RateType.String(),
		"rate_date":     data.RateDate,
		"place_code":    data.PlaceCode,
		"currency_code": data.CurrencyCode,
	}
	if data.Code != "" {
		values["code"] = data.Code
	}

	ut.ConvertByteToIMData(data.RateMeta, values, "rate_meta")
	ut.ConvertByteToIMData(data.RateMap, values, "rate_map")

	// database insert
	var rows []cu.IM
	var result cu.IM
	var rateID int64
	if rateID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "rate"}); err == nil {
		if rows, err = ds.StoreDataGet(cu.IM{"id": rateID, "model": "rate"}, true); err == nil {
			result = rows[0]
		}
	}
	RespondMessage(w, http.StatusCreated, result, http.StatusUnprocessableEntity, err)
}

// RatePut - update rate
func RatePut(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	rateID := cu.ToInteger(r.PathValue("id_code"), 0)
	rateCode := cu.ToString(r.PathValue("id_code"), "")

	var rate md.Rate
	var inputFields, metaFields []string
	var err error
	if _, inputFields, metaFields, err = ds.GetBodyData("rate", r.Body, &rate); err == nil {
		_, err = ds.UpdateData(md.UpdateDataOptions{
			Model: "rate", IDKey: rateID, Code: rateCode,
			Data: rate, Meta: rate.RateMeta, Fields: inputFields, MetaFields: metaFields,
		})
	}
	RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

func RateDelete(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	rateID := cu.ToInteger(r.PathValue("id_code"), 0)
	rateCode := cu.ToString(r.PathValue("id_code"), "")
	err := ds.DataDelete("rate", rateID, rateCode)
	RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// RateQuery - get rates
func RateQuery(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)

	var params cu.IM = cu.IM{
		"model":  "rate",
		"limit":  cu.ToInteger(r.URL.Query().Get("limit"), 0),
		"offset": cu.ToInteger(r.URL.Query().Get("offset"), 0),
	}
	for _, v := range []string{"rate_type", "currency_code", "tag"} {
		if r.URL.Query().Get(v) != "" {
			params[v] = cu.ToString(r.URL.Query().Get(v), "")
		}
	}
	response, err := ds.StoreDataGet(params, false)
	RespondMessage(w, http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// RateGet - get rate by id or code
func RateGet(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	rateID := cu.ToInteger(r.PathValue("id_code"), 0)
	rateCode := cu.ToString(r.PathValue("id_code"), "")
	var err error
	var rates []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if rates, err = ds.GetDataByID("rate", rateID, rateCode, true); err == nil {
		response = rates[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	RespondMessage(w, http.StatusOK, response, errStatus, err)
}
