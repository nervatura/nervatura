package http

import (
	"net/http"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

// EmployeePost - create new employee
func EmployeePost(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)

	// convert request body to struct and schema validation
	var data md.Employee = md.Employee{
		Address: md.Address{},
		Contact: md.Contact{},
		Events:  []md.Event{},
		EmployeeMeta: md.EmployeeMeta{
			Tags: []string{},
		},
		EmployeeMap: cu.IM{},
	}
	err := ds.ConvertFromReader(r.Body, &data)
	if err != nil {
		RespondMessage(w, 0, nil, http.StatusUnprocessableEntity, err)
		return
	}

	// prepare values for database update
	values := cu.IM{}
	if data.Code != "" {
		values["code"] = data.Code
	}

	ut.ConvertByteToIMData(data.Contact, values, "contact")
	ut.ConvertByteToIMData(data.Address, values, "address")
	ut.ConvertByteToIMData(data.Events, values, "events")
	ut.ConvertByteToIMData(data.EmployeeMeta, values, "employee_meta")
	ut.ConvertByteToIMData(data.EmployeeMap, values, "employee_map")

	// database insert
	var rows []cu.IM
	var result cu.IM
	var employeeID int64
	if employeeID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "employee"}); err == nil {
		if rows, err = ds.StoreDataGet(cu.IM{"id": employeeID, "model": "employee"}, true); err == nil {
			result = rows[0]
		}
	}
	RespondMessage(w, http.StatusCreated, result, http.StatusUnprocessableEntity, err)
}

// EmployeePut - update employee
func EmployeePut(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	employeeID := cu.ToInteger(r.PathValue("id_code"), 0)
	employeeCode := cu.ToString(r.PathValue("id_code"), "")

	var employee md.Employee
	var inputFields, metaFields []string
	var err error
	if _, inputFields, metaFields, err = ds.GetBodyData("employee", r.Body, &employee); err == nil {
		_, err = ds.UpdateData(md.UpdateDataOptions{
			Model: "employee", IDKey: employeeID, Code: employeeCode,
			Data: employee, Meta: employee.EmployeeMeta, Fields: inputFields, MetaFields: metaFields,
		})
	}
	RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

func EmployeeDelete(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	employeeID := cu.ToInteger(r.PathValue("id_code"), 0)
	employeeCode := cu.ToString(r.PathValue("id_code"), "")
	err := ds.DataDelete("employee", employeeID, employeeCode)
	RespondMessage(w, http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// EmployeeQuery - get employees
func EmployeeQuery(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)

	var params cu.IM = cu.IM{
		"model":  "employee",
		"limit":  cu.ToInteger(r.URL.Query().Get("limit"), 0),
		"offset": cu.ToInteger(r.URL.Query().Get("offset"), 0),
	}
	if r.URL.Query().Get("tag") != "" {
		params["tag"] = cu.ToString(r.URL.Query().Get("tag"), "")
	}
	response, err := ds.StoreDataGet(params, false)
	RespondMessage(w, http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// EmployeeGet - get employee by id or code
func EmployeeGet(w http.ResponseWriter, r *http.Request) {
	ds := r.Context().Value(md.DataStoreCtxKey).(*api.DataStore)
	employeeID := cu.ToInteger(r.PathValue("id_code"), 0)
	employeeCode := cu.ToString(r.PathValue("id_code"), "")
	var err error
	var employees []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if employees, err = ds.GetDataByID("employee", employeeID, employeeCode, true); err == nil {
		response = employees[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	RespondMessage(w, http.StatusOK, response, errStatus, err)
}
