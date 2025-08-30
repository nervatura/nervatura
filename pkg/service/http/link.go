package http

import (
	"errors"
	"net/http"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

// LinkPost - create new link
func LinkPost(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)

	// convert request body to struct and schema validation
	var data md.Link = md.Link{
		LinkType1: md.LinkType(md.LinkTypeTrans),
		LinkType2: md.LinkType(md.LinkTypeTrans),
		LinkMeta: md.LinkMeta{
			Tags: []string{},
		},
		LinkMap: cu.IM{},
	}
	err := ds.ConvertFromReader(r.Body, &data)
	if err != nil {
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
	}

	if data.LinkCode1 == "" || data.LinkCode2 == "" {
		err = errors.New("link link_code_1 and link_code_2 are required")
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
	}

	// prepare values for database update
	values := cu.IM{
		"link_type_1": data.LinkType1.String(),
		"link_code_1": data.LinkCode1,
		"link_type_2": data.LinkType2.String(),
		"link_code_2": data.LinkCode2,
	}
	if data.Code != "" {
		values["code"] = data.Code
	}

	ut.ConvertByteToIMData(data.LinkMeta, values, "link_meta")
	ut.ConvertByteToIMData(data.LinkMap, values, "link_map")

	// database insert
	var rows []cu.IM
	var result cu.IM
	var linkID int64
	if linkID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "link"}); err == nil {
		if rows, err = ds.StoreDataGet(cu.IM{"id": linkID, "model": "link"}, true); err == nil {
			result = rows[0]
		}
	}
	RespondMessage(w, http.StatusCreated, result, http.StatusUnprocessableEntity, err)
}

// LinkPut - update link
func LinkPut(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	linkID := cu.ToInteger(r.PathValue("id_code"), 0)
	linkCode := cu.ToString(r.PathValue("id_code"), "")

	var link md.Link
	var inputFields, metaFields []string
	var err error
	if _, inputFields, metaFields, err = ds.GetBodyData("link", r.Body, &link); err == nil {
		err = ds.UpdateData(md.UpdateDataOptions{
			Model: "link", IDKey: linkID, Code: linkCode,
			Data: link, Meta: link.LinkMeta, Fields: inputFields, MetaFields: metaFields,
		})
	}
	RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

func LinkDelete(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	linkID := cu.ToInteger(r.PathValue("id_code"), 0)
	linkCode := cu.ToString(r.PathValue("id_code"), "")
	err := ds.DataDelete("link", linkID, linkCode)
	RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// LinkQuery - get links
func LinkQuery(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)

	var params cu.IM = cu.IM{
		"model":  "link",
		"limit":  cu.ToInteger(r.URL.Query().Get("limit"), 0),
		"offset": cu.ToInteger(r.URL.Query().Get("offset"), 0),
	}
	for _, v := range []string{"link_type_1", "link_code_1", "link_type_2", "link_code_2", "tag"} {
		if r.URL.Query().Get(v) != "" {
			params[v] = cu.ToString(r.URL.Query().Get(v), "")
		}
	}
	response, err := ds.StoreDataGet(params, false)
	RespondMessage(w, http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// LinkGet - get link by id or code
func LinkGet(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	linkID := cu.ToInteger(r.PathValue("id_code"), 0)
	linkCode := cu.ToString(r.PathValue("id_code"), "")
	var err error
	var links []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if links, err = ds.GetDataByID("link", linkID, linkCode, true); err == nil {
		response = links[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	RespondMessage(w, http.StatusOK, response, errStatus, err)
}
