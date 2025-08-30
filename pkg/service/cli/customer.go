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

// CustomerInsert - insert a new customer
func (cli *CLIService) CustomerInsert(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	// convert request body to struct and schema validation
	var data md.Customer = md.Customer{
		CustomerType: md.CustomerType(md.CustomerTypeCompany),
		Addresses:    []md.Address{},
		Contacts:     []md.Contact{},
		Events:       []md.Event{},
		CustomerMeta: md.CustomerMeta{
			Tags: []string{},
		},
		CustomerMap: cu.IM{},
	}
	err := ds.ConvertFromByte([]byte(requestData), &data)
	if err != nil {
		return cli.respondString(0, nil, http.StatusUnprocessableEntity, err)
	}

	if data.CustomerName == "" {
		err = errors.New("customer name is required")
		return cli.respondString(0, nil, http.StatusUnprocessableEntity, err)
	}

	// prepare values for database update
	values := cu.IM{
		"customer_type": data.CustomerType.String(),
		"customer_name": data.CustomerName,
	}
	if data.Code != "" {
		values["code"] = data.Code
	}

	ut.ConvertByteToIMData(data.Contacts, values, "contacts")
	ut.ConvertByteToIMData(data.Addresses, values, "addresses")
	ut.ConvertByteToIMData(data.Events, values, "events")
	ut.ConvertByteToIMData(data.CustomerMeta, values, "customer_meta")
	ut.ConvertByteToIMData(data.CustomerMap, values, "customer_map")

	// database insert
	var rows []cu.IM
	var result cu.IM
	var customerID int64
	if customerID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "customer"}); err == nil {
		if rows, err = ds.StoreDataGet(cu.IM{"id": customerID, "model": "customer"}, true); err == nil {
			result = rows[0]
		}
	}
	return cli.respondString(http.StatusCreated, result, http.StatusUnprocessableEntity, err)
}

// CustomerUpdate - update a customer
func (cli *CLIService) CustomerUpdate(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	customerID := cu.ToInteger(options["id"], 0)
	customerCode := cu.ToString(options["code"], "")

	var customer md.Customer
	var inputFields, metaFields []string
	var err error
	reader := io.NopCloser(bytes.NewReader([]byte(requestData)))
	if _, inputFields, metaFields, err = ds.GetBodyData("customer", reader, &customer); err == nil {
		err = ds.UpdateData(md.UpdateDataOptions{
			Model: "customer", IDKey: customerID, Code: customerCode,
			Data: customer, Meta: customer.CustomerMeta, Fields: inputFields, MetaFields: metaFields,
		})
	}
	return cli.respondString(http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// CustomerDelete - delete a customer
func (cli *CLIService) CustomerDelete(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	customerID := cu.ToInteger(options["id"], 0)
	customerCode := cu.ToString(options["code"], "")
	err := ds.DataDelete("customer", customerID, customerCode)
	return cli.respondString(http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// CustomerQuery - query customers
func (cli *CLIService) CustomerQuery(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	var params cu.IM = cu.IM{"model": "customer"}
	for _, v := range []string{"customer_type", "customer_name", "tag"} {
		if options[v] != nil {
			params[v] = options[v]
		}
	}
	response, err := ds.StoreDataGet(params, false)
	return cli.respondString(http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

func (cli *CLIService) CustomerGet(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	customerID := cu.ToInteger(options["id"], 0)
	customerCode := cu.ToString(options["code"], "")
	var err error
	var customers []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if customers, err = ds.GetDataByID("customer", customerID, customerCode, true); err == nil {
		response = customers[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	return cli.respondString(http.StatusOK, response, errStatus, err)
}
