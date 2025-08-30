package http

import (
	"errors"
	"net/http"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

// ProductPost - create new product
func ProductPost(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)

	// convert request body to struct and schema validation
	var data md.Product = md.Product{
		ProductType: md.ProductType(md.ProductTypeItem),
		Events:      []md.Event{},
		ProductMeta: md.ProductMeta{
			Tags: []string{},
		},
		ProductMap: cu.IM{},
	}
	err := ds.ConvertFromReader(r.Body, &data)
	if err != nil {
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
	}

	if data.ProductName == "" || data.TaxCode == "" {
		err = errors.New("product name and tax code are required")
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
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
	RespondMessage(w, http.StatusCreated, result, http.StatusUnprocessableEntity, err)
}

// ProductPut - update product
func ProductPut(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	productID := cu.ToInteger(r.PathValue("id_code"), 0)
	productCode := cu.ToString(r.PathValue("id_code"), "")

	var product md.Product
	var inputFields, metaFields []string
	var err error
	if _, inputFields, metaFields, err = ds.GetBodyData("product", r.Body, &product); err == nil {
		err = ds.UpdateData(md.UpdateDataOptions{
			Model: "product", IDKey: productID, Code: productCode,
			Data: product, Meta: product.ProductMeta, Fields: inputFields, MetaFields: metaFields,
		})
	}
	RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

func ProductDelete(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	productID := cu.ToInteger(r.PathValue("id_code"), 0)
	productCode := cu.ToString(r.PathValue("id_code"), "")
	err := ds.DataDelete("product", productID, productCode)
	RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// ProductQuery - get products
func ProductQuery(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)

	var params cu.IM = cu.IM{
		"model":  "product",
		"limit":  cu.ToInteger(r.URL.Query().Get("limit"), 0),
		"offset": cu.ToInteger(r.URL.Query().Get("offset"), 0),
	}
	for _, v := range []string{"product_type", "product_name", "tag"} {
		if r.URL.Query().Get(v) != "" {
			params[v] = cu.ToString(r.URL.Query().Get(v), "")
		}
	}
	response, err := ds.StoreDataGet(params, false)
	RespondMessage(w, http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// ProductGet - get product by id or code
func ProductGet(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	productID := cu.ToInteger(r.PathValue("id_code"), 0)
	productCode := cu.ToString(r.PathValue("id_code"), "")
	var err error
	var products []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if products, err = ds.GetDataByID("product", productID, productCode, true); err == nil {
		response = products[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	RespondMessage(w, http.StatusOK, response, errStatus, err)
}
