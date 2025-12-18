package cli

import (
	"bytes"
	"errors"
	"io"
	"net/http"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

// PaymentInsert - create new payment
func (cli *CLIService) PaymentInsert(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	// convert request body to struct and schema validation
	var data md.Payment = md.Payment{
		PaymentMeta: md.PaymentMeta{
			Tags: []string{},
		},
		PaymentMap: cu.IM{},
	}
	err := ds.ConvertFromByte([]byte(requestData), &data)
	if err != nil {
		return cli.respondString(http.StatusUnprocessableEntity, nil, http.StatusUnprocessableEntity, err)
	}

	if data.PaidDate == "" || data.TransCode == "" {
		err = errors.New("payment paid_date and trans_code are required")
		return cli.respondString(http.StatusUnprocessableEntity, nil, http.StatusUnprocessableEntity, err)
	}

	// prepare values for database update
	values := cu.IM{
		"paid_date":  data.PaidDate,
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
	return cli.respondString(http.StatusCreated, result, http.StatusUnprocessableEntity, err)
}

// PaymentPut - update payment
func (cli *CLIService) PaymentUpdate(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	paymentID := cu.ToInteger(options["id"], 0)
	paymentCode := cu.ToString(options["code"], "")

	var payment md.Payment
	var inputFields, metaFields []string
	var err error
	reader := io.NopCloser(bytes.NewReader([]byte(requestData)))
	if _, inputFields, metaFields, err = ds.GetBodyData("payment", reader, &payment); err == nil {
		_, err = ds.UpdateData(md.UpdateDataOptions{
			Model: "payment", IDKey: paymentID, Code: paymentCode,
			Data: payment, Meta: payment.PaymentMeta, Fields: inputFields, MetaFields: metaFields,
		})
	}
	return cli.respondString(http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

func (cli *CLIService) PaymentDelete(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	paymentID := cu.ToInteger(options["id"], 0)
	paymentCode := cu.ToString(options["code"], "")
	err := ds.DataDelete("payment", paymentID, paymentCode)
	return cli.respondString(http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// PaymentQuery - get payments
func (cli *CLIService) PaymentQuery(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	var params cu.IM = cu.IM{"model": "payment"}
	for _, v := range []string{"trans_code", "paid_date", "tag"} {
		if options[v] != nil {
			params[v] = options[v]
		}
	}
	response, err := ds.StoreDataGet(params, false)
	return cli.respondString(http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// PaymentGet - get payment by id or code
func (cli *CLIService) PaymentGet(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	paymentID := cu.ToInteger(options["id"], 0)
	paymentCode := cu.ToString(options["code"], "")
	var err error
	var payments []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if payments, err = ds.GetDataByID("payment", paymentID, paymentCode, true); err == nil {
		response = payments[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	return cli.respondString(http.StatusOK, response, errStatus, err)
}
