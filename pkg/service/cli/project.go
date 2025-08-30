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

// ProjectInsert - create new project
func (cli *CLIService) ProjectInsert(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

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
	err := ds.ConvertFromByte([]byte(requestData), &data)
	if err != nil {
		return cli.respondString(http.StatusUnprocessableEntity, nil, http.StatusUnprocessableEntity, err)
	}

	if data.ProjectName == "" {
		err = errors.New("project name is required")
		return cli.respondString(http.StatusUnprocessableEntity, nil, http.StatusUnprocessableEntity, err)
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
	return cli.respondString(http.StatusCreated, result, http.StatusUnprocessableEntity, err)
}

// ProjectPut - update project
func (cli *CLIService) ProjectUpdate(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	projectID := cu.ToInteger(options["id"], 0)
	projectCode := cu.ToString(options["code"], "")

	var project md.Project
	var inputFields, metaFields []string
	var err error
	reader := io.NopCloser(bytes.NewReader([]byte(requestData)))
	if _, inputFields, metaFields, err = ds.GetBodyData("project", reader, &project); err == nil {
		err = ds.UpdateData(md.UpdateDataOptions{
			Model: "project", IDKey: projectID, Code: projectCode,
			Data: project, Meta: project.ProjectMeta, Fields: inputFields, MetaFields: metaFields,
		})
	}
	return cli.respondString(http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

func (cli *CLIService) ProjectDelete(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	projectID := cu.ToInteger(options["id"], 0)
	projectCode := cu.ToString(options["code"], "")
	err := ds.DataDelete("project", projectID, projectCode)
	return cli.respondString(http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// ProjectQuery - get projects
func (cli *CLIService) ProjectQuery(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	var params cu.IM = cu.IM{"model": "project"}
	for _, v := range []string{"project_name", "customer_code", "tag"} {
		if options[v] != nil {
			params[v] = options[v]
		}
	}
	response, err := ds.StoreDataGet(params, false)
	return cli.respondString(http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// ProjectGet - get project by id or code
func (cli *CLIService) ProjectGet(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	projectID := cu.ToInteger(options["id"], 0)
	projectCode := cu.ToString(options["code"], "")
	var err error
	var projects []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if projects, err = ds.GetDataByID("project", projectID, projectCode, true); err == nil {
		response = projects[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	return cli.respondString(http.StatusOK, response, errStatus, err)
}
