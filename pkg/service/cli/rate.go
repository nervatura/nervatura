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

// RatePost - create new rate
func (cli *CLIService) RateInsert(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	// convert request body to struct and schema validation
	var data md.Rate = md.Rate{
		RateType: md.RateType(md.RateTypeRate),
		RateMeta: md.RateMeta{
			Tags: []string{},
		},
		RateMap: cu.IM{},
	}
	err := ds.ConvertFromByte([]byte(requestData), &data)
	if err != nil {
		return cli.respondString(http.StatusUnprocessableEntity, nil, http.StatusUnprocessableEntity, err)
	}

	if data.RateDate.IsZero() || data.PlaceCode == "" || data.CurrencyCode == "" {
		err = errors.New("rate date, place code and currency code are required")
		return cli.respondString(http.StatusUnprocessableEntity, nil, http.StatusUnprocessableEntity, err)
	}

	// prepare values for database update
	values := cu.IM{
		"rate_type":     data.RateType.String(),
		"rate_date":     data.RateDate.Format(time.DateOnly),
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
	return cli.respondString(http.StatusCreated, result, http.StatusUnprocessableEntity, err)
}

// RatePut - update rate
func (cli *CLIService) RateUpdate(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	rateID := cu.ToInteger(options["id"], 0)
	rateCode := cu.ToString(options["code"], "")

	var rate md.Rate
	var inputFields, metaFields []string
	var err error
	reader := io.NopCloser(bytes.NewReader([]byte(requestData)))
	if _, inputFields, metaFields, err = ds.GetBodyData("rate", reader, &rate); err == nil {
		_, err = ds.UpdateData(md.UpdateDataOptions{
			Model: "rate", IDKey: rateID, Code: rateCode,
			Data: rate, Meta: rate.RateMeta, Fields: inputFields, MetaFields: metaFields,
		})
	}
	return cli.respondString(http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

func (cli *CLIService) RateDelete(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	rateID := cu.ToInteger(options["id"], 0)
	rateCode := cu.ToString(options["code"], "")
	err := ds.DataDelete("rate", rateID, rateCode)
	return cli.respondString(http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// RateQuery - get rates
func (cli *CLIService) RateQuery(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	var params cu.IM = cu.IM{"model": "rate"}
	for _, v := range []string{"rate_type", "currency_code", "tag"} {
		if options[v] != nil {
			params[v] = options[v]
		}
	}
	response, err := ds.StoreDataGet(params, false)
	return cli.respondString(http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// RateGet - get rate by id or code
func (cli *CLIService) RateGet(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	rateID := cu.ToInteger(options["id"], 0)
	rateCode := cu.ToString(options["code"], "")
	var err error
	var rates []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if rates, err = ds.GetDataByID("rate", rateID, rateCode, true); err == nil {
		response = rates[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	return cli.respondString(http.StatusOK, response, errStatus, err)
}
