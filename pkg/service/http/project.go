package http

import (
	"errors"
	"net/http"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

// ProjectPost - create new project
func ProjectPost(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)

	// convert request body to struct and schema validation
	var data md.Project = md.Project{
		Addresses: []md.Address{},
		Contacts:  []md.Contact{},
		Events:    []md.Event{},
		ProjectMeta: md.ProjectMeta{
			Tags: []string{},
		},
		ProjectMap: cu.IM{},
	}
	err := ds.ConvertFromReader(r.Body, &data)
	if err != nil {
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
	}

	if data.ProjectName == "" {
		err = errors.New("project name is required")
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
	}

	// prepare values for database update
	values := cu.IM{
		"project_name": data.ProjectName,
	}
	if data.Code != "" {
		values["code"] = data.Code
	}
	if data.CustomerCode != "" {
		values["customer_code"] = data.CustomerCode
	}

	ut.ConvertByteToIMData(data.Contacts, values, "contacts")
	ut.ConvertByteToIMData(data.Addresses, values, "addresses")
	ut.ConvertByteToIMData(data.Events, values, "events")
	ut.ConvertByteToIMData(data.ProjectMeta, values, "project_meta")
	ut.ConvertByteToIMData(data.ProjectMap, values, "project_map")

	// database insert
	var rows []cu.IM
	var result cu.IM
	var projectID int64
	if projectID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "project"}); err == nil {
		if rows, err = ds.StoreDataGet(cu.IM{"id": projectID, "model": "project"}, true); err == nil {
			result = rows[0]
		}
	}
	RespondMessage(w, http.StatusCreated, result, http.StatusUnprocessableEntity, err)
}

// ProjectPut - update project
func ProjectPut(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	projectID := cu.ToInteger(r.PathValue("id_code"), 0)
	projectCode := cu.ToString(r.PathValue("id_code"), "")

	var project md.Project
	var inputFields, metaFields []string
	var err error
	if _, inputFields, metaFields, err = ds.GetBodyData("project", r.Body, &project); err == nil {
		_, err = ds.UpdateData(md.UpdateDataOptions{
			Model: "project", IDKey: projectID, Code: projectCode,
			Data: project, Meta: project.ProjectMeta, Fields: inputFields, MetaFields: metaFields,
		})
	}
	RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

func ProjectDelete(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	projectID := cu.ToInteger(r.PathValue("id_code"), 0)
	projectCode := cu.ToString(r.PathValue("id_code"), "")
	err := ds.DataDelete("project", projectID, projectCode)
	RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// ProjectQuery - get projects
func ProjectQuery(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)

	var params cu.IM = cu.IM{
		"model":  "project",
		"limit":  cu.ToInteger(r.URL.Query().Get("limit"), 0),
		"offset": cu.ToInteger(r.URL.Query().Get("offset"), 0),
	}
	for _, v := range []string{"project_name", "customer_code", "tag"} {
		if r.URL.Query().Get(v) != "" {
			params[v] = cu.ToString(r.URL.Query().Get(v), "")
		}
	}
	response, err := ds.StoreDataGet(params, false)
	RespondMessage(w, http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// ProjectGet - get project by id or code
func ProjectGet(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	projectID := cu.ToInteger(r.PathValue("id_code"), 0)
	projectCode := cu.ToString(r.PathValue("id_code"), "")
	var err error
	var projects []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if projects, err = ds.GetDataByID("project", projectID, projectCode, true); err == nil {
		response = projects[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	RespondMessage(w, http.StatusOK, response, errStatus, err)
}
