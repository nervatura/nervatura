package http

import (
	"errors"
	"net/http"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

// TaxPost - create new tax
func TaxPost(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)

	// convert request body to struct and schema validation
	var data md.Tax = md.Tax{
		TaxMeta: md.TaxMeta{
			Tags: []string{},
		},
		TaxMap: cu.IM{},
	}
	err := ds.ConvertFromReader(r.Body, &data)
	if err != nil {
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
	}

	if data.Code == "" {
		err = errors.New("tax code is required")
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
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
	RespondMessage(w, http.StatusCreated, result, http.StatusUnprocessableEntity, err)
}

// TaxPut - update tax
func TaxPut(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	taxID := cu.ToInteger(r.PathValue("id_code"), 0)
	taxCode := cu.ToString(r.PathValue("id_code"), "")

	var tax md.Tax
	var inputFields, metaFields []string
	var err error
	if _, inputFields, metaFields, err = ds.GetBodyData("tax", r.Body, &tax); err == nil {
		err = ds.UpdateData(md.UpdateDataOptions{
			Model: "tax", IDKey: taxID, Code: taxCode,
			Data: tax, Meta: tax.TaxMeta, Fields: inputFields, MetaFields: metaFields,
		})
	}
	RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

func TaxDelete(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	taxID := cu.ToInteger(r.PathValue("id_code"), 0)
	taxCode := cu.ToString(r.PathValue("id_code"), "")
	err := ds.DataDelete("tax", taxID, taxCode)
	RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// TaxQuery - get taxes
func TaxQuery(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)

	var params cu.IM = cu.IM{
		"model":  "tax",
		"limit":  cu.ToInteger(r.URL.Query().Get("limit"), 0),
		"offset": cu.ToInteger(r.URL.Query().Get("offset"), 0),
	}
	if r.URL.Query().Get("tag") != "" {
		params["tag"] = cu.ToString(r.URL.Query().Get("tag"), "")
	}
	response, err := ds.StoreDataGet(params, false)
	RespondMessage(w, http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// TaxGet - get tax by id or code
func TaxGet(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	taxID := cu.ToInteger(r.PathValue("id_code"), 0)
	taxCode := cu.ToString(r.PathValue("id_code"), "")
	var err error
	var taxs []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if taxs, err = ds.GetDataByID("tax", taxID, taxCode, true); err == nil {
		response = taxs[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	RespondMessage(w, http.StatusOK, response, errStatus, err)
}
