package http

import (
	"errors"
	"net/http"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

// ItemPost - create new item
func ItemPost(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)

	// convert request body to struct and schema validation
	var data md.Item = md.Item{
		ItemMeta: md.ItemMeta{
			Tags: []string{},
		},
		ItemMap: cu.IM{},
	}
	err := ds.ConvertFromReader(r.Body, &data)
	if err != nil {
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
	}

	if data.TransCode == "" || data.ProductCode == "" || data.TaxCode == "" {
		err = errors.New("item trans_code, product_code and tax_code are required")
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
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
	RespondMessage(w, http.StatusCreated, result, http.StatusUnprocessableEntity, err)
}

// ItemPut - update item
func ItemPut(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	itemID := cu.ToInteger(r.PathValue("id_code"), 0)
	itemCode := cu.ToString(r.PathValue("id_code"), "")

	var item md.Item
	var inputFields, metaFields []string
	var err error
	if _, inputFields, metaFields, err = ds.GetBodyData("item", r.Body, &item); err == nil {
		err = ds.UpdateData(md.UpdateDataOptions{
			Model: "item", IDKey: itemID, Code: itemCode,
			Data: item, Meta: item.ItemMeta, Fields: inputFields, MetaFields: metaFields,
		})
	}
	RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

func ItemDelete(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	itemID := cu.ToInteger(r.PathValue("id_code"), 0)
	itemCode := cu.ToString(r.PathValue("id_code"), "")
	err := ds.DataDelete("item", itemID, itemCode)
	RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// ItemQuery - get items
func ItemQuery(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)

	var params cu.IM = cu.IM{
		"model":  "item",
		"limit":  cu.ToInteger(r.URL.Query().Get("limit"), 0),
		"offset": cu.ToInteger(r.URL.Query().Get("offset"), 0),
	}
	for _, v := range []string{"trans_code", "product_code", "tag"} {
		if r.URL.Query().Get(v) != "" {
			params[v] = cu.ToString(r.URL.Query().Get(v), "")
		}
	}
	response, err := ds.StoreDataGet(params, false)
	RespondMessage(w, http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// ItemGet - get item by id or code
func ItemGet(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	itemID := cu.ToInteger(r.PathValue("id_code"), 0)
	itemCode := cu.ToString(r.PathValue("id_code"), "")
	var err error
	var items []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if items, err = ds.GetDataByID("item", itemID, itemCode, true); err == nil {
		response = items[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	RespondMessage(w, http.StatusOK, response, errStatus, err)
}
