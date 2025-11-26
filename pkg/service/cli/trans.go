package cli

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"slices"
	"time"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

// TransInsert - create new trans
func (cli *CLIService) TransInsert(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	// convert request body to struct and schema validation
	now, _ := time.Parse(time.DateOnly, time.Now().Format(time.DateOnly))
	var data md.Trans = md.Trans{
		TransType: md.TransType(md.TransTypeInvoice),
		Direction: md.Direction(md.DirectionOut),
		TransMeta: md.TransMeta{
			DueTime:    md.TimeDateTime{Time: now},
			Status:     md.TransStatus(md.TransStatusNormal),
			TransState: md.TransState(md.TransStateOK),
			Worksheet:  md.TransMetaWorksheet{},
			Rent:       md.TransMetaRent{},
			Invoice:    md.TransMetaInvoice{},
			Tags:       []string{},
		},
		TransMap: cu.IM{},
	}
	err := ds.ConvertFromByte([]byte(requestData), &data)
	if err != nil {
		return cli.respondString(http.StatusUnprocessableEntity, nil, http.StatusUnprocessableEntity, err)
	}

	if data.TransDate.IsZero() {
		err = errors.New("trans date is required")
		return cli.respondString(http.StatusUnprocessableEntity, nil, http.StatusUnprocessableEntity, err)
	}

	if slices.Contains([]md.TransType{
		md.TransTypeInvoice, md.TransTypeReceipt, md.TransTypeOffer, md.TransTypeOrder, md.TransTypeWorksheet, md.TransTypeRent}, data.TransType,
	) && (data.CustomerCode == "" || data.CurrencyCode == "") {
		err = errors.New("invoice, receipt, offer, order, worksheet and rent must have customer code and currency code")
		return cli.respondString(http.StatusUnprocessableEntity, nil, http.StatusUnprocessableEntity, err)
	}

	// prepare values for database update
	values := cu.IM{
		"trans_type": data.TransType.String(),
		"direction":  data.Direction.String(),
		"trans_date": data.TransDate.Format(time.DateOnly),
		"auth_code":  data.AuthCode,
	}

	// Optional fields
	optionalFields := map[string]string{
		"code":          data.Code,
		"customer_code": data.CustomerCode,
		"employee_code": data.EmployeeCode,
		"project_code":  data.ProjectCode,
		"place_code":    data.PlaceCode,
		"trans_code":    data.TransCode,
		"currency_code": data.CurrencyCode,
	}

	for key, value := range optionalFields {
		if value != "" {
			values[key] = value
		}
	}

	ut.ConvertByteToIMData(data.TransMeta, values, "trans_meta")
	ut.ConvertByteToIMData(data.TransMap, values, "trans_map")

	// database insert
	var rows []cu.IM
	var result cu.IM
	var transID int64
	if transID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "trans"}); err == nil {
		if rows, err = ds.StoreDataGet(cu.IM{"id": transID, "model": "trans"}, true); err == nil {
			result = rows[0]
		}
	}
	return cli.respondString(http.StatusCreated, result, http.StatusUnprocessableEntity, err)
}

// TransPut - update trans
func (cli *CLIService) TransUpdate(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	transID := cu.ToInteger(options["id"], 0)
	transCode := cu.ToString(options["code"], "")

	var trans md.Trans
	var inputFields, metaFields []string
	var err error
	reader := io.NopCloser(bytes.NewReader([]byte(requestData)))
	if _, inputFields, metaFields, err = ds.GetBodyData("trans", reader, &trans); err == nil {
		_, err = ds.UpdateData(md.UpdateDataOptions{
			Model: "trans", IDKey: transID, Code: transCode,
			Data: trans, Meta: trans.TransMeta, Fields: inputFields, MetaFields: metaFields,
		})
	}
	return cli.respondString(http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

func (cli *CLIService) TransDelete(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	transID := cu.ToInteger(options["id"], 0)
	transCode := cu.ToString(options["code"], "")
	err := ds.DataDelete("trans", transID, transCode)
	return cli.respondString(http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// TransQuery - get transs
func (cli *CLIService) TransQuery(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	var params cu.IM = cu.IM{"model": "trans"}
	for _, v := range []string{"trans_type", "direction", "trans_date", "tag"} {
		if options[v] != nil {
			params[v] = options[v]
		}
	}
	response, err := ds.StoreDataGet(params, false)
	return cli.respondString(http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// TransGet - get trans by id or code
func (cli *CLIService) TransGet(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	transID := cu.ToInteger(options["id"], 0)
	transCode := cu.ToString(options["code"], "")
	var err error
	var transs []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if transs, err = ds.GetDataByID("trans", transID, transCode, true); err == nil {
		response = transs[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	return cli.respondString(http.StatusOK, response, errStatus, err)
}
