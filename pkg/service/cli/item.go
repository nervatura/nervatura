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

// ItemInsert - create new item
func (cli *CLIService) ItemInsert(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	// convert request body to struct and schema validation
	var data md.Item = md.Item{
		ItemMeta: md.ItemMeta{
			Tags: []string{},
		},
		ItemMap: cu.IM{},
	}
	err := ds.ConvertFromByte([]byte(requestData), &data)
	if err != nil {
		return cli.respondString(http.StatusUnprocessableEntity, nil, http.StatusUnprocessableEntity, err)
	}

	if data.TransCode == "" || data.ProductCode == "" || data.TaxCode == "" {
		err = errors.New("item trans_code, product_code and tax_code are required")
		return cli.respondString(http.StatusUnprocessableEntity, nil, http.StatusUnprocessableEntity, err)
	}

	// prepare values for database update
	values := cu.IM{
		"trans_code":   data.TransCode,
		"product_code": data.ProductCode,
		"tax_code":     data.TaxCode,
	}
	if data.Code != "" {
		values["code"] = data.Code
	}

	ut.ConvertByteToIMData(data.ItemMeta, values, "item_meta")
	ut.ConvertByteToIMData(data.ItemMap, values, "item_map")

	// database insert
	var rows []cu.IM
	var result cu.IM
	var itemID int64
	if itemID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "item"}); err == nil {
		if rows, err = ds.StoreDataGet(cu.IM{"id": itemID, "model": "item"}, true); err == nil {
			result = rows[0]
		}
	}
	return cli.respondString(http.StatusCreated, result, http.StatusUnprocessableEntity, err)
}

// ItemUpdate - update item
func (cli *CLIService) ItemUpdate(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	itemID := cu.ToInteger(options["id"], 0)
	itemCode := cu.ToString(options["code"], "")

	var item md.Item
	var inputFields, metaFields []string
	var err error
	reader := io.NopCloser(bytes.NewReader([]byte(requestData)))
	if _, inputFields, metaFields, err = ds.GetBodyData("item", reader, &item); err == nil {
		_, err = ds.UpdateData(md.UpdateDataOptions{
			Model: "item", IDKey: itemID, Code: itemCode,
			Data: item, Meta: item.ItemMeta, Fields: inputFields, MetaFields: metaFields,
		})
	}
	return cli.respondString(http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// ItemDelete - delete item
func (cli *CLIService) ItemDelete(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	itemID := cu.ToInteger(options["id"], 0)
	itemCode := cu.ToString(options["code"], "")
	err := ds.DataDelete("item", itemID, itemCode)
	return cli.respondString(http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// ItemQuery - get items
func (cli *CLIService) ItemQuery(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	var params cu.IM = cu.IM{"model": "item"}
	for _, v := range []string{"trans_code", "product_code", "tag"} {
		if options[v] != nil {
			params[v] = options[v]
		}
	}
	response, err := ds.StoreDataGet(params, false)
	return cli.respondString(http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// ItemGet - get item by id or code
func (cli *CLIService) ItemGet(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	itemID := cu.ToInteger(options["id"], 0)
	itemCode := cu.ToString(options["code"], "")
	var err error
	var items []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if items, err = ds.GetDataByID("item", itemID, itemCode, true); err == nil {
		response = items[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	return cli.respondString(http.StatusOK, response, errStatus, err)
}
