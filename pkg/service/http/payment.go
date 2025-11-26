package http

import (
	"errors"
	"net/http"
	"time"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

// PaymentPost - create new payment
func PaymentPost(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)

	// convert request body to struct and schema validation
	var data md.Payment = md.Payment{
		PaymentMeta: md.PaymentMeta{
			Tags: []string{},
		},
		PaymentMap: cu.IM{},
	}
	err := ds.ConvertFromReader(r.Body, &data)
	if err != nil {
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
	}

	if data.PaidDate.IsZero() || data.TransCode == "" {
		err = errors.New("payment paid_date and trans_code are required")
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
	}

	// prepare values for database update
	values := cu.IM{
		"paid_date":  data.PaidDate.Format(time.DateOnly),
		"trans_code": data.TransCode,
	}
	if data.Code != "" {
		values["code"] = data.Code
	}

	ut.ConvertByteToIMData(data.PaymentMeta, values, "payment_meta")
	ut.ConvertByteToIMData(data.PaymentMap, values, "payment_map")

	// database insert
	var rows []cu.IM
	var result cu.IM
	var paymentID int64
	if paymentID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "payment"}); err == nil {
		if rows, err = ds.StoreDataGet(cu.IM{"id": paymentID, "model": "payment"}, true); err == nil {
			result = rows[0]
		}
	}
	RespondMessage(w, http.StatusCreated, result, http.StatusUnprocessableEntity, err)
}

// PaymentPut - update payment
func PaymentPut(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	paymentID := cu.ToInteger(r.PathValue("id_code"), 0)
	paymentCode := cu.ToString(r.PathValue("id_code"), "")

	var payment md.Payment
	var inputFields, metaFields []string
	var err error
	if _, inputFields, metaFields, err = ds.GetBodyData("payment", r.Body, &payment); err == nil {
		_, err = ds.UpdateData(md.UpdateDataOptions{
			Model: "payment", IDKey: paymentID, Code: paymentCode,
			Data: payment, Meta: payment.PaymentMeta, Fields: inputFields, MetaFields: metaFields,
		})
	}
	RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

func PaymentDelete(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	paymentID := cu.ToInteger(r.PathValue("id_code"), 0)
	paymentCode := cu.ToString(r.PathValue("id_code"), "")
	err := ds.DataDelete("payment", paymentID, paymentCode)
	RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// PaymentQuery - get payments
func PaymentQuery(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)

	var params cu.IM = cu.IM{
		"model":  "payment",
		"limit":  cu.ToInteger(r.URL.Query().Get("limit"), 0),
		"offset": cu.ToInteger(r.URL.Query().Get("offset"), 0),
	}
	for _, v := range []string{"trans_code", "paid_date", "tag"} {
		if r.URL.Query().Get(v) != "" {
			params[v] = cu.ToString(r.URL.Query().Get(v), "")
		}
	}
	response, err := ds.StoreDataGet(params, false)
	RespondMessage(w, http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// PaymentGet - get payment by id or code
func PaymentGet(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	paymentID := cu.ToInteger(r.PathValue("id_code"), 0)
	paymentCode := cu.ToString(r.PathValue("id_code"), "")
	var err error
	var payments []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if payments, err = ds.GetDataByID("payment", paymentID, paymentCode, true); err == nil {
		response = payments[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	RespondMessage(w, http.StatusOK, response, errStatus, err)
}
