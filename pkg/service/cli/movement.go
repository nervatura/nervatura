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

// MovementInsert - create new movement
func (cli *CLIService) MovementInsert(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	// convert request body to struct and schema validation
	var data md.Movement = md.Movement{
		MovementType: md.MovementType(md.MovementTypeInventory),
		MovementMeta: md.MovementMeta{
			Tags: []string{},
		},
		MovementMap: cu.IM{},
	}
	err := ds.ConvertFromByte([]byte(requestData), &data)
	if err != nil {
		return cli.respondString(http.StatusUnprocessableEntity, nil, http.StatusUnprocessableEntity, err)
	}

	if data.ShippingTime == "" || data.TransCode == "" {
		err = errors.New("movement shipping_time and trans_code are required")
		return cli.respondString(http.StatusUnprocessableEntity, nil, http.StatusUnprocessableEntity, err)
	}

	if data.ProductCode == "" && data.ToolCode == "" {
		err = errors.New("movement product_code or tool_code are required")
		return cli.respondString(http.StatusUnprocessableEntity, nil, http.StatusUnprocessableEntity, err)
	}

	// prepare values for database update
	values := cu.IM{
		"shipping_time": data.ShippingTime,
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
	return cli.respondString(http.StatusCreated, result, http.StatusUnprocessableEntity, err)
}

// MovementPut - update movement
func (cli *CLIService) MovementUpdate(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	movementID := cu.ToInteger(options["id"], 0)
	movementCode := cu.ToString(options["code"], "")

	var movement md.Movement
	var inputFields, metaFields []string
	var err error
	reader := io.NopCloser(bytes.NewReader([]byte(requestData)))
	if _, inputFields, metaFields, err = ds.GetBodyData("movement", reader, &movement); err == nil {
		_, err = ds.UpdateData(md.UpdateDataOptions{
			Model: "movement", IDKey: movementID, Code: movementCode,
			Data: movement, Meta: movement.MovementMeta, Fields: inputFields, MetaFields: metaFields,
		})
	}
	return cli.respondString(http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

func (cli *CLIService) MovementDelete(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	movementID := cu.ToInteger(options["id"], 0)
	movementCode := cu.ToString(options["code"], "")
	err := ds.DataDelete("movement", movementID, movementCode)
	return cli.respondString(http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// MovementQuery - get movements
func (cli *CLIService) MovementQuery(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	var params cu.IM = cu.IM{"model": "movement"}
	for _, v := range []string{"trans_code", "movement_type", "tag"} {
		if options[v] != nil {
			params[v] = options[v]
		}
	}
	response, err := ds.StoreDataGet(params, false)
	return cli.respondString(http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// MovementGet - get movement by id or code
func (cli *CLIService) MovementGet(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	movementID := cu.ToInteger(options["id"], 0)
	movementCode := cu.ToString(options["code"], "")
	var err error
	var movements []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if movements, err = ds.GetDataByID("movement", movementID, movementCode, true); err == nil {
		response = movements[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	return cli.respondString(http.StatusOK, response, errStatus, err)
}
