package http

import (
	"errors"
	"net/http"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

// CustomerPost - create new customer
func CustomerPost(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)

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
	err := ds.ConvertFromReader(r.Body, &data)
	if err != nil {
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
	}

	if data.CustomerName == "" {
		err = errors.New("customer name is required")
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
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
	RespondMessage(w, http.StatusCreated, result, http.StatusUnprocessableEntity, err)
}

// CustomerPut - update customer
func CustomerPut(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	customerID := cu.ToInteger(r.PathValue("id_code"), 0)
	customerCode := cu.ToString(r.PathValue("id_code"), "")

	var customer md.Customer
	var inputFields, metaFields []string
	var err error
	if _, inputFields, metaFields, err = ds.GetBodyData("customer", r.Body, &customer); err == nil {
		_, err = ds.UpdateData(md.UpdateDataOptions{
			Model: "customer", IDKey: customerID, Code: customerCode,
			Data: customer, Meta: customer.CustomerMeta, Fields: inputFields, MetaFields: metaFields,
		})
	}
	RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

func CustomerDelete(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	customerID := cu.ToInteger(r.PathValue("id_code"), 0)
	customerCode := cu.ToString(r.PathValue("id_code"), "")
	err := ds.DataDelete("customer", customerID, customerCode)
	RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// CustomerQuery - get customers
func CustomerQuery(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)

	var params cu.IM = cu.IM{
		"model":  "customer",
		"limit":  cu.ToInteger(r.URL.Query().Get("limit"), 0),
		"offset": cu.ToInteger(r.URL.Query().Get("offset"), 0),
	}
	for _, v := range []string{"customer_type", "customer_name", "tag"} {
		if r.URL.Query().Get(v) != "" {
			params[v] = cu.ToString(r.URL.Query().Get(v), "")
		}
	}
	response, err := ds.StoreDataGet(params, false)
	RespondMessage(w, http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// CustomerGet - get customer by id or code
func CustomerGet(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	customerID := cu.ToInteger(r.PathValue("id_code"), 0)
	customerCode := cu.ToString(r.PathValue("id_code"), "")
	var err error
	var customers []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if customers, err = ds.GetDataByID("customer", customerID, customerCode, true); err == nil {
		response = customers[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	RespondMessage(w, http.StatusOK, response, errStatus, err)
}
