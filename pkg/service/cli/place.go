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

// PlaceInsert - create new place
func (cli *CLIService) PlaceInsert(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	// convert request body to struct and schema validation
	var data md.Place = md.Place{
		PlaceType: md.PlaceType(md.PlaceTypeWarehouse),
		Address:   md.Address{},
		Contacts:  []md.Contact{},
		Events:    []md.Event{},
		PlaceMeta: md.PlaceMeta{
			Tags: []string{},
		},
		PlaceMap: cu.IM{},
	}
	err := ds.ConvertFromByte([]byte(requestData), &data)
	if err != nil {
		return cli.respondString(http.StatusUnprocessableEntity, nil, http.StatusUnprocessableEntity, err)
	}

	if data.PlaceName == "" {
		err = errors.New("place name is required")
		return cli.respondString(http.StatusUnprocessableEntity, nil, http.StatusUnprocessableEntity, err)
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
	ut.ConvertByteToIMData(data.Events, values, "events")
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
	return cli.respondString(http.StatusCreated, result, http.StatusUnprocessableEntity, err)
}

// PlacePut - update place
func (cli *CLIService) PlaceUpdate(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	placeID := cu.ToInteger(options["id"], 0)
	placeCode := cu.ToString(options["code"], "")

	var place md.Place
	var inputFields, metaFields []string
	var err error
	reader := io.NopCloser(bytes.NewReader([]byte(requestData)))
	if _, inputFields, metaFields, err = ds.GetBodyData("place", reader, &place); err == nil {
		_, err = ds.UpdateData(md.UpdateDataOptions{
			Model: "place", IDKey: placeID, Code: placeCode,
			Data: place, Meta: place.PlaceMeta, Fields: inputFields, MetaFields: metaFields,
		})
	}
	return cli.respondString(http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

func (cli *CLIService) PlaceDelete(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	placeID := cu.ToInteger(options["id"], 0)
	placeCode := cu.ToString(options["code"], "")
	err := ds.DataDelete("place", placeID, placeCode)
	return cli.respondString(http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// PlaceQuery - get places
func (cli *CLIService) PlaceQuery(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	var params cu.IM = cu.IM{"model": "place"}
	for _, v := range []string{"place_type", "place_name", "tag"} {
		if options[v] != nil {
			params[v] = options[v]
		}
	}
	response, err := ds.StoreDataGet(params, false)
	return cli.respondString(http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// PlaceGet - get place by id or code
func (cli *CLIService) PlaceGet(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	placeID := cu.ToInteger(options["id"], 0)
	placeCode := cu.ToString(options["code"], "")
	var err error
	var places []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if places, err = ds.GetDataByID("place", placeID, placeCode, true); err == nil {
		response = places[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	return cli.respondString(http.StatusOK, response, errStatus, err)
}
