package http

import (
	"errors"
	"net/http"
	"slices"
	"time"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

// TransPost - create new trans
func TransPost(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	user := r.Context().Value(md.AuthUserCtxKey).(*md.Auth)

	// convert request body to struct and schema validation
	now, _ := time.Parse(time.DateOnly, time.Now().Format(time.DateOnly))
	var data md.Trans = md.Trans{
		TransType: md.TransType(md.TransTypeInvoice),
		Direction: md.Direction(md.DirectionOut),
		AuthCode:  user.Code,
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
	err := ds.ConvertFromReader(r.Body, &data)
	if err != nil {
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
	}

	if data.TransDate.IsZero() {
		err = errors.New("trans date is required")
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
	}

	if slices.Contains([]md.TransType{
		md.TransTypeInvoice, md.TransTypeReceipt, md.TransTypeOffer, md.TransTypeOrder, md.TransTypeWorksheet, md.TransTypeRent}, data.TransType,
	) && (data.CustomerCode == "" || data.CurrencyCode == "") {
		err = errors.New("invoice, receipt, offer, order, worksheet and rent must have customer code and currency code")
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
	}

	// prepare values for database update
	values := cu.IM{
		"trans_type": data.TransType.String(),
		"direction":  data.Direction.String(),
		"trans_date": data.TransDate.Format(time.DateOnly),
		"auth_code":  user.Code,
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
	RespondMessage(w, http.StatusCreated, result, http.StatusUnprocessableEntity, err)
}

// TransPut - update trans
func TransPut(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	transID := cu.ToInteger(r.PathValue("id_code"), 0)
	transCode := cu.ToString(r.PathValue("id_code"), "")

	var trans md.Trans
	var inputFields, metaFields []string
	var err error
	if _, inputFields, metaFields, err = ds.GetBodyData("trans", r.Body, &trans); err == nil {
		err = ds.UpdateData(md.UpdateDataOptions{
			Model: "trans", IDKey: transID, Code: transCode,
			Data: trans, Meta: trans.TransMeta, Fields: inputFields, MetaFields: metaFields,
		})
	}
	RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

func TransDelete(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	transID := cu.ToInteger(r.PathValue("id_code"), 0)
	transCode := cu.ToString(r.PathValue("id_code"), "")
	err := ds.DataDelete("trans", transID, transCode)
	RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// TransQuery - get transs
func TransQuery(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)

	var params cu.IM = cu.IM{
		"model":  "trans",
		"limit":  cu.ToInteger(r.URL.Query().Get("limit"), 0),
		"offset": cu.ToInteger(r.URL.Query().Get("offset"), 0),
	}
	for _, v := range []string{"trans_type", "direction", "trans_date", "tag"} {
		if r.URL.Query().Get(v) != "" {
			params[v] = cu.ToString(r.URL.Query().Get(v), "")
		}
	}
	response, err := ds.StoreDataGet(params, false)
	RespondMessage(w, http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// TransGet - get trans by id or code
func TransGet(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	transID := cu.ToInteger(r.PathValue("id_code"), 0)
	transCode := cu.ToString(r.PathValue("id_code"), "")
	var err error
	var transs []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if transs, err = ds.GetDataByID("trans", transID, transCode, true); err == nil {
		response = transs[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	RespondMessage(w, http.StatusOK, response, errStatus, err)
}
