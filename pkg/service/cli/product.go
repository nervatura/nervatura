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

// ProductInsert - create new product
func (cli *CLIService) ProductInsert(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	// convert request body to struct and schema validation
	var data md.Product = md.Product{
		ProductType: md.ProductType(md.ProductTypeItem),
		Events:      []md.Event{},
		ProductMeta: md.ProductMeta{
			Tags: []string{},
		},
		ProductMap: cu.IM{},
	}
	err := ds.ConvertFromByte([]byte(requestData), &data)
	if err != nil {
		return cli.respondString(http.StatusUnprocessableEntity, nil, http.StatusUnprocessableEntity, err)
	}

	if data.ProductName == "" || data.TaxCode == "" {
		err = errors.New("product name and tax code are required")
		return cli.respondString(http.StatusUnprocessableEntity, nil, http.StatusUnprocessableEntity, err)
	}

	// prepare values for database update
	values := cu.IM{
		"product_type": data.ProductType.String(),
		"product_name": data.ProductName,
		"tax_code":     data.TaxCode,
	}
	if data.Code != "" {
		values["code"] = data.Code
	}

	ut.ConvertByteToIMData(data.Events, values, "events")
	ut.ConvertByteToIMData(data.ProductMeta, values, "product_meta")
	ut.ConvertByteToIMData(data.ProductMap, values, "product_map")

	// database insert
	var rows []cu.IM
	var result cu.IM
	var productID int64
	if productID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "product"}); err == nil {
		if rows, err = ds.StoreDataGet(cu.IM{"id": productID, "model": "product"}, true); err == nil {
			result = rows[0]
		}
	}
	return cli.respondString(http.StatusCreated, result, http.StatusUnprocessableEntity, err)
}

// ProductPut - update product
func (cli *CLIService) ProductUpdate(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	productID := cu.ToInteger(options["id"], 0)
	productCode := cu.ToString(options["code"], "")

	var product md.Product
	var inputFields, metaFields []string
	var err error
	reader := io.NopCloser(bytes.NewReader([]byte(requestData)))
	if _, inputFields, metaFields, err = ds.GetBodyData("product", reader, &product); err == nil {
		err = ds.UpdateData(md.UpdateDataOptions{
			Model: "product", IDKey: productID, Code: productCode,
			Data: product, Meta: product.ProductMeta, Fields: inputFields, MetaFields: metaFields,
		})
	}
	return cli.respondString(http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

func (cli *CLIService) ProductDelete(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	productID := cu.ToInteger(options["id"], 0)
	productCode := cu.ToString(options["code"], "")
	err := ds.DataDelete("product", productID, productCode)
	return cli.respondString(http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// ProductQuery - get products
func (cli *CLIService) ProductQuery(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	var params cu.IM = cu.IM{"model": "product"}
	for _, v := range []string{"product_type", "product_name", "tag"} {
		if options[v] != nil {
			params[v] = options[v]
		}
	}
	response, err := ds.StoreDataGet(params, false)
	return cli.respondString(http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// ProductGet - get product by id or code
func (cli *CLIService) ProductGet(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	productID := cu.ToInteger(options["id"], 0)
	productCode := cu.ToString(options["code"], "")
	var err error
	var products []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if products, err = ds.GetDataByID("product", productID, productCode, true); err == nil {
		response = products[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	return cli.respondString(http.StatusOK, response, errStatus, err)
}
