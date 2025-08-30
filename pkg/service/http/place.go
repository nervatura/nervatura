package http

import (
	"errors"
	"net/http"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

// PlacePost - create new place
func PlacePost(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)

	// convert request body to struct and schema validation
	var data md.Place = md.Place{
		PlaceType: md.PlaceType(md.PlaceTypeWarehouse),
		Address:   md.Address{},
		Contacts:  []md.Contact{},
		PlaceMeta: md.PlaceMeta{
			Tags: []string{},
		},
		PlaceMap: cu.IM{},
	}
	err := ds.ConvertFromReader(r.Body, &data)
	if err != nil {
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
	}

	if data.PlaceName == "" {
		err = errors.New("place name is required")
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
	}

	// prepare values for database update
	values := cu.IM{
		"place_type": data.PlaceType.String(),
		"place_name": data.PlaceName,
	}
	if data.Code != "" {
		values["code"] = data.Code
	}
	if data.CurrencyCode != "" {
		values["currency_code"] = data.CurrencyCode
	}

	ut.ConvertByteToIMData(data.Contacts, values, "contacts")
	ut.ConvertByteToIMData(data.Address, values, "address")
	ut.ConvertByteToIMData(data.PlaceMeta, values, "place_meta")
	ut.ConvertByteToIMData(data.PlaceMap, values, "place_map")

	// database insert
	var rows []cu.IM
	var result cu.IM
	var placeID int64
	if placeID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "place"}); err == nil {
		if rows, err = ds.StoreDataGet(cu.IM{"id": placeID, "model": "place"}, true); err == nil {
			result = rows[0]
		}
	}
	RespondMessage(w, http.StatusCreated, result, http.StatusUnprocessableEntity, err)
}

// PlacePut - update place
func PlacePut(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	placeID := cu.ToInteger(r.PathValue("id_code"), 0)
	placeCode := cu.ToString(r.PathValue("id_code"), "")

	var place md.Place
	var inputFields, metaFields []string
	var err error
	if _, inputFields, metaFields, err = ds.GetBodyData("place", r.Body, &place); err == nil {
		err = ds.UpdateData(md.UpdateDataOptions{
			Model: "place", IDKey: placeID, Code: placeCode,
			Data: place, Meta: place.PlaceMeta, Fields: inputFields, MetaFields: metaFields,
		})
	}
	RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

func PlaceDelete(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	placeID := cu.ToInteger(r.PathValue("id_code"), 0)
	placeCode := cu.ToString(r.PathValue("id_code"), "")
	err := ds.DataDelete("place", placeID, placeCode)
	RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// PlaceQuery - get places
func PlaceQuery(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)

	var params cu.IM = cu.IM{
		"model":  "place",
		"limit":  cu.ToInteger(r.URL.Query().Get("limit"), 0),
		"offset": cu.ToInteger(r.URL.Query().Get("offset"), 0),
	}
	for _, v := range []string{"place_type", "place_name", "tag"} {
		if r.URL.Query().Get(v) != "" {
			params[v] = cu.ToString(r.URL.Query().Get(v), "")
		}
	}
	response, err := ds.StoreDataGet(params, false)
	RespondMessage(w, http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// PlaceGet - get place by id or code
func PlaceGet(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	placeID := cu.ToInteger(r.PathValue("id_code"), 0)
	placeCode := cu.ToString(r.PathValue("id_code"), "")
	var err error
	var places []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if places, err = ds.GetDataByID("place", placeID, placeCode, true); err == nil {
		response = places[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	RespondMessage(w, http.StatusOK, response, errStatus, err)
}
