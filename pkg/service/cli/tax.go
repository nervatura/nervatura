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

// TaxInsert - create new tax
func (cli *CLIService) TaxInsert(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	// convert request body to struct and schema validation
	var data md.Tax = md.Tax{
		TaxMeta: md.TaxMeta{
			Tags: []string{},
		},
		TaxMap: cu.IM{},
	}
	err := ds.ConvertFromByte([]byte(requestData), &data)
	if err != nil {
		return cli.respondString(http.StatusUnprocessableEntity, nil, http.StatusUnprocessableEntity, err)
	}

	if data.Code == "" {
		err = errors.New("tax code is required")
		return cli.respondString(http.StatusUnprocessableEntity, nil, http.StatusUnprocessableEntity, err)
	}

	// prepare values for database update
	values := cu.IM{
		"code": data.Code,
	}

	ut.ConvertByteToIMData(data.TaxMeta, values, "tax_meta")
	ut.ConvertByteToIMData(data.TaxMap, values, "tax_map")

	// database insert
	var rows []cu.IM
	var result cu.IM
	var taxID int64
	if taxID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "tax"}); err == nil {
		if rows, err = ds.StoreDataGet(cu.IM{"id": taxID, "model": "tax"}, true); err == nil {
			result = rows[0]
		}
	}
	return cli.respondString(http.StatusCreated, result, http.StatusUnprocessableEntity, err)
}

// TaxPut - update tax
func (cli *CLIService) TaxUpdate(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	taxID := cu.ToInteger(options["id"], 0)
	taxCode := cu.ToString(options["code"], "")

	var tax md.Tax
	var inputFields, metaFields []string
	var err error
	reader := io.NopCloser(bytes.NewReader([]byte(requestData)))
	if _, inputFields, metaFields, err = ds.GetBodyData("tax", reader, &tax); err == nil {
		err = ds.UpdateData(md.UpdateDataOptions{
			Model: "tax", IDKey: taxID, Code: taxCode,
			Data: tax, Meta: tax.TaxMeta, Fields: inputFields, MetaFields: metaFields,
		})
	}
	return cli.respondString(http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

func (cli *CLIService) TaxDelete(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	taxID := cu.ToInteger(options["id"], 0)
	taxCode := cu.ToString(options["code"], "")
	err := ds.DataDelete("tax", taxID, taxCode)
	return cli.respondString(http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// TaxQuery - get taxes
func (cli *CLIService) TaxQuery(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	var params cu.IM = cu.IM{"model": "tax"}
	if tag, found := options["tag"].(string); found && tag != "" {
		params["tag"] = tag
	}
	response, err := ds.StoreDataGet(params, false)
	return cli.respondString(http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// TaxGet - get tax by id or code
func (cli *CLIService) TaxGet(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	taxID := cu.ToInteger(options["id"], 0)
	taxCode := cu.ToString(options["code"], "")
	var err error
	var taxs []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if taxs, err = ds.GetDataByID("tax", taxID, taxCode, true); err == nil {
		response = taxs[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	return cli.respondString(http.StatusOK, response, errStatus, err)
}
