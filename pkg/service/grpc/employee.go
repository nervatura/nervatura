package grpc

import (
	"context"
	"errors"
	"net/http"
	"slices"

	cu "github.com/nervatura/component/pkg/util"
	api "github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	pb "github.com/nervatura/nervatura/v6/pkg/service/grpc/proto"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

func (s *GService) EmployeeUpdate(ctx context.Context, req *pb.Employee) (pbEmployee *pb.Employee, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	user := ctx.Value(md.AuthUserCtxKey).(md.Auth)
	if user.UserGroup == md.UserGroupGuest {
		return pbEmployee, errors.New(http.StatusText(http.StatusMethodNotAllowed))
	}

	var updateID int64 = req.Id
	if req.Id > 0 || req.Code != "" {
		var rows []cu.IM
		if rows, err = ds.GetDataByID("employee", req.Id, req.Code, false); err != nil {
			return pbEmployee, err
		}
		if len(rows) > 0 {
			updateID = cu.ToInteger(rows[0]["id"], 0)
		}
	}

	values := cu.IM{}
	if updateID == 0 && req.Code != "" {
		values["code"] = req.Code
	}

	ut.ConvertByteToIMValue(&req.Contact, &pb.Contact{}, values, "contact")
	ut.ConvertByteToIMValue(&req.Address, &pb.Address{}, values, "address")
	ut.ConvertByteToIMValue(&req.Events, []*pb.Event{}, values, "events")
	ut.ConvertByteToIMValue(&req.EmployeeMeta, &pb.EmployeeMeta{}, values, "employee_meta")
	ut.ConvertByteToIMValue(&req.EmployeeMap, &pb.JsonString{}, values, "employee_map")

	update := md.Update{Values: values, Model: "employee"}
	if updateID > 0 {
		update.IDKey = updateID
	}

	if updateID, err = ds.StoreDataUpdate(update); err == nil {
		pbEmployee, err = s.EmployeeGet(ctx, &pb.RequestGet{Id: updateID, Code: ""})
	}

	return pbEmployee, err
}

func (s *GService) EmployeeGet(ctx context.Context, req *pb.RequestGet) (pbEmployee *pb.Employee, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	var employees []cu.IM
	if employees, err = ds.GetDataByID("employee_view", req.Id, req.Code, true); err == nil {
		if employeeJson, found := employees[0]["employee_object"].(string); found {
			err = ds.ConvertFromByte([]byte(employeeJson), &pbEmployee)
		} else {
			err = errors.New(http.StatusText(http.StatusUnprocessableEntity))
		}
	}
	return pbEmployee, err
}

func (s *GService) EmployeeQuery(ctx context.Context, req *pb.RequestQuery) (pbEmployees *pb.Employees, err error) {
	pbEmployees = &pb.Employees{}
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	var params cu.IM = cu.IM{"model": "employee_view", "limit": req.Limit, "offset": req.Offset}
	for _, filter := range req.Filters {
		if slices.Contains([]string{"tag"}, filter.GetField()) {
			params[filter.GetField()] = filter.GetValue()
		} else {
			return pbEmployees, errors.New(http.StatusText(http.StatusUnprocessableEntity) +
				": invalid filter field. Valid fields are: tag")
		}
	}
	if rows, err := ds.StoreDataGet(params, false); err == nil {
		for _, row := range rows {
			var pbEmployee *pb.Employee
			if employeeJson, found := row["employee_object"].(string); found {
				ds.ConvertFromByte([]byte(employeeJson), &pbEmployee)
				pbEmployees.Data = append(pbEmployees.Data, pbEmployee)
			}
		}
	}
	return pbEmployees, err
}
