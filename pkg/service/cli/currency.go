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

// CurrencyInsert - create new currency
func (cli *CLIService) CurrencyInsert(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	// convert request body to struct and schema validation
	var data md.Currency = md.Currency{
		CurrencyMeta: md.CurrencyMeta{
			Tags: []string{},
		},
		CurrencyMap: cu.IM{},
	}
	err := ds.ConvertFromByte([]byte(requestData), &data)
	if err != nil {
		return cli.respondString(http.StatusUnprocessableEntity, nil, http.StatusUnprocessableEntity, err)
	}

	if data.Code == "" {
		err = errors.New("code is required")
		return cli.respondString(http.StatusUnprocessableEntity, nil, http.StatusUnprocessableEntity, err)
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
	return cli.respondString(http.StatusCreated, result, http.StatusUnprocessableEntity, err)
}

// CurrencyUpdate - update currency
func (cli *CLIService) CurrencyUpdate(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	currencyID := cu.ToInteger(options["id"], 0)
	currencyCode := cu.ToString(options["code"], "")

	var currency md.Currency
	var inputFields, metaFields []string
	var err error
	reader := io.NopCloser(bytes.NewReader([]byte(requestData)))
	if _, inputFields, metaFields, err = ds.GetBodyData("currency", reader, &currency); err == nil {
		err = ds.UpdateData(md.UpdateDataOptions{
			Model: "currency", IDKey: currencyID, Code: currencyCode,
			Data: currency, Meta: currency.CurrencyMeta, Fields: inputFields, MetaFields: metaFields,
		})
	}
	return cli.respondString(http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

func (cli *CLIService) CurrencyDelete(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	currencyID := cu.ToInteger(options["id"], 0)
	currencyCode := cu.ToString(options["code"], "")
	err := ds.DataDelete("currency", currencyID, currencyCode)
	return cli.respondString(http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// CurrencyQuery - get currencys
func (cli *CLIService) CurrencyQuery(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	var params cu.IM = cu.IM{"model": "currency"}
	if tag, found := options["tag"].(string); found {
		params["tag"] = tag
	}
	response, err := ds.StoreDataGet(params, false)
	return cli.respondString(http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// CurrencyGet - get currency by id or code
func (cli *CLIService) CurrencyGet(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	currencyID := cu.ToInteger(options["id"], 0)
	currencyCode := cu.ToString(options["code"], "")
	var err error
	var currencys []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if currencys, err = ds.GetDataByID("currency", currencyID, currencyCode, true); err == nil {
		response = currencys[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	return cli.respondString(http.StatusOK, response, errStatus, err)
}
