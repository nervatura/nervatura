package http

import (
	"errors"
	"net/http"
	"time"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

// MovementPost - create new movement
func MovementPost(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)

	// convert request body to struct and schema validation
	var data md.Movement = md.Movement{
		MovementType: md.MovementType(md.MovementTypeInventory),
		MovementMeta: md.MovementMeta{
			Tags: []string{},
		},
		MovementMap: cu.IM{},
	}
	err := ds.ConvertFromReader(r.Body, &data)
	if err != nil {
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
	}

	if data.ShippingTime.IsZero() || data.TransCode == "" {
		err = errors.New("movement shipping_time and trans_code are required")
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
	}

	if data.ProductCode == "" && data.ToolCode == "" {
		err = errors.New("movement product_code or tool_code are required")
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
	}

	// prepare values for database update
	values := cu.IM{
		"shipping_time": data.ShippingTime.Format(time.RFC3339),
		"trans_code":    data.TransCode,
	}

	// Optional fields
	optionalFields := map[string]string{
		"code":          data.Code,
		"product_code":  data.ProductCode,
		"tool_code":     data.ToolCode,
		"place_code":    data.PlaceCode,
		"item_code":     data.ItemCode,
		"movement_code": data.MovementCode,
	}
	for key, value := range optionalFields {
		if value != "" {
			values[key] = value
		}
	}

	ut.ConvertByteToIMData(data.MovementMeta, values, "movement_meta")
	ut.ConvertByteToIMData(data.MovementMap, values, "movement_map")

	// database insert
	var rows []cu.IM
	var result cu.IM
	var movementID int64
	if movementID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "movement"}); err == nil {
		if rows, err = ds.StoreDataGet(cu.IM{"id": movementID, "model": "movement"}, true); err == nil {
			result = rows[0]
		}
	}
	RespondMessage(w, http.StatusCreated, result, http.StatusUnprocessableEntity, err)
}

// MovementPut - update movement
func MovementPut(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	movementID := cu.ToInteger(r.PathValue("id_code"), 0)
	movementCode := cu.ToString(r.PathValue("id_code"), "")

	var movement md.Movement
	var inputFields, metaFields []string
	var err error
	if _, inputFields, metaFields, err = ds.GetBodyData("movement", r.Body, &movement); err == nil {
		_, err = ds.UpdateData(md.UpdateDataOptions{
			Model: "movement", IDKey: movementID, Code: movementCode,
			Data: movement, Meta: movement.MovementMeta, Fields: inputFields, MetaFields: metaFields,
		})
	}
	RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

func MovementDelete(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	movementID := cu.ToInteger(r.PathValue("id_code"), 0)
	movementCode := cu.ToString(r.PathValue("id_code"), "")
	err := ds.DataDelete("movement", movementID, movementCode)
	RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// MovementQuery - get movements
func MovementQuery(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)

	var params cu.IM = cu.IM{
		"model":  "movement",
		"limit":  cu.ToInteger(r.URL.Query().Get("limit"), 0),
		"offset": cu.ToInteger(r.URL.Query().Get("offset"), 0),
	}
	for _, v := range []string{"trans_code", "movement_type", "tag"} {
		if r.URL.Query().Get(v) != "" {
			params[v] = cu.ToString(r.URL.Query().Get(v), "")
		}
	}
	response, err := ds.StoreDataGet(params, false)
	RespondMessage(w, http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// MovementGet - get movement by id or code
func MovementGet(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	movementID := cu.ToInteger(r.PathValue("id_code"), 0)
	movementCode := cu.ToString(r.PathValue("id_code"), "")
	var err error
	var movements []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if movements, err = ds.GetDataByID("movement", movementID, movementCode, true); err == nil {
		response = movements[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	RespondMessage(w, http.StatusOK, response, errStatus, err)
}
