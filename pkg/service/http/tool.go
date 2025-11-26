package http

import (
	"errors"
	"net/http"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

// ToolPost - create new tool
func ToolPost(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)

	// convert request body to struct and schema validation
	var data md.Tool = md.Tool{
		Events: []md.Event{},
		ToolMeta: md.ToolMeta{
			Tags: []string{},
		},
		ToolMap: cu.IM{},
	}
	err := ds.ConvertFromReader(r.Body, &data)
	if err != nil {
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
	}

	if data.ProductCode == "" || data.Description == "" {
		err = errors.New("product code and description are required")
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
	}

	// prepare values for database update
	values := cu.IM{
		"product_code": data.ProductCode,
		"description":  data.Description,
	}
	if data.Code != "" {
		values["code"] = data.Code
	}

	ut.ConvertByteToIMData(data.Events, values, "events")
	ut.ConvertByteToIMData(data.ToolMeta, values, "tool_meta")
	ut.ConvertByteToIMData(data.ToolMap, values, "tool_map")

	// database insert
	var rows []cu.IM
	var result cu.IM
	var toolID int64
	if toolID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "tool"}); err == nil {
		if rows, err = ds.StoreDataGet(cu.IM{"id": toolID, "model": "tool"}, true); err == nil {
			result = rows[0]
		}
	}
	RespondMessage(w, http.StatusCreated, result, http.StatusUnprocessableEntity, err)
}

// ToolPut - update tool
func ToolPut(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	toolID := cu.ToInteger(r.PathValue("id_code"), 0)
	toolCode := cu.ToString(r.PathValue("id_code"), "")

	var tool md.Tool
	var inputFields, metaFields []string
	var err error
	if _, inputFields, metaFields, err = ds.GetBodyData("tool", r.Body, &tool); err == nil {
		_, err = ds.UpdateData(md.UpdateDataOptions{
			Model: "tool", IDKey: toolID, Code: toolCode,
			Data: tool, Meta: tool.ToolMeta, Fields: inputFields, MetaFields: metaFields,
		})
	}
	RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

func ToolDelete(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	toolID := cu.ToInteger(r.PathValue("id_code"), 0)
	toolCode := cu.ToString(r.PathValue("id_code"), "")
	err := ds.DataDelete("tool", toolID, toolCode)
	RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// ToolQuery - get tools
func ToolQuery(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)

	var params cu.IM = cu.IM{
		"model":  "tool",
		"limit":  cu.ToInteger(r.URL.Query().Get("limit"), 0),
		"offset": cu.ToInteger(r.URL.Query().Get("offset"), 0),
	}
	for _, v := range []string{"product_code", "description", "tag"} {
		if r.URL.Query().Get(v) != "" {
			params[v] = cu.ToString(r.URL.Query().Get(v), "")
		}
	}
	response, err := ds.StoreDataGet(params, false)
	RespondMessage(w, http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// ToolGet - get tool by id or code
func ToolGet(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	toolID := cu.ToInteger(r.PathValue("id_code"), 0)
	toolCode := cu.ToString(r.PathValue("id_code"), "")
	var err error
	var tools []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if tools, err = ds.GetDataByID("tool", toolID, toolCode, true); err == nil {
		response = tools[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	RespondMessage(w, http.StatusOK, response, errStatus, err)
}
