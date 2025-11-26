package cli

import (
	"bytes"
	"io"
	"net/http"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

// EmployeeInsert - create new employee
func (cli *CLIService) EmployeeInsert(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

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
	err := ds.ConvertFromByte([]byte(requestData), &data)
	if err != nil {
		return cli.respondString(http.StatusUnprocessableEntity, nil, http.StatusUnprocessableEntity, err)
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
	return cli.respondString(http.StatusCreated, result, http.StatusUnprocessableEntity, err)
}

// EmployeePut - update employee
func (cli *CLIService) EmployeeUpdate(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	employeeID := cu.ToInteger(options["id"], 0)
	employeeCode := cu.ToString(options["code"], "")

	var employee md.Employee
	var inputFields, metaFields []string
	var err error
	reader := io.NopCloser(bytes.NewReader([]byte(requestData)))
	if _, inputFields, metaFields, err = ds.GetBodyData("employee", reader, &employee); err == nil {
		_, err = ds.UpdateData(md.UpdateDataOptions{
			Model: "employee", IDKey: employeeID, Code: employeeCode,
			Data: employee, Meta: employee.EmployeeMeta, Fields: inputFields, MetaFields: metaFields,
		})
	}
	return cli.respondString(http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

func (cli *CLIService) EmployeeDelete(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	employeeID := cu.ToInteger(options["id"], 0)
	employeeCode := cu.ToString(options["code"], "")
	err := ds.DataDelete("employee", employeeID, employeeCode)
	return cli.respondString(http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// EmployeeQuery - get employees
func (cli *CLIService) EmployeeQuery(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	var params cu.IM = cu.IM{"model": "employee"}
	if tag, found := options["tag"].(string); found {
		params["tag"] = tag
	}
	response, err := ds.StoreDataGet(params, false)
	return cli.respondString(http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// EmployeeGet - get employee by id or code
func (cli *CLIService) EmployeeGet(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	employeeID := cu.ToInteger(options["id"], 0)
	employeeCode := cu.ToString(options["code"], "")
	var err error
	var employees []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if employees, err = ds.GetDataByID("employee", employeeID, employeeCode, true); err == nil {
		response = employees[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	return cli.respondString(http.StatusOK, response, errStatus, err)
}
