package grpc

import (
	"context"
	"errors"
	"net/http"
	"slices"

	cu "github.com/nervatura/component/pkg/util"
	api "github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
	pb "github.com/nervatura/nervatura/v6/proto"
)

func (s *GService) CustomerUpdate(ctx context.Context, req *pb.Customer) (pbCustomer *pb.Customer, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	user := ctx.Value(md.AuthUserCtxKey).(md.Auth)
	if user.UserGroup == md.UserGroupGuest {
		return pbCustomer, errors.New(http.StatusText(http.StatusMethodNotAllowed))
	}

	var updateID int64 = req.Id
	if req.Id > 0 || req.Code != "" {
		var rows []cu.IM
		if rows, err = ds.GetDataByID("customer", req.Id, req.Code, false); err != nil {
			return pbCustomer, err
		}
		if len(rows) > 0 {
			updateID = cu.ToInteger(rows[0]["id"], 0)
		}
	}

	if req.CustomerName == "" {
		return pbCustomer, errors.New(http.StatusText(http.StatusUnprocessableEntity) + ": customer name is required")
	}
	values := cu.IM{
		"customer_type": req.CustomerType.String(),
		"customer_name": req.CustomerName,
	}
	if updateID == 0 && req.Code != "" {
		values["code"] = req.Code
	}

	ut.ConvertByteToIMValue(&req.Contacts, []*pb.Contact{}, values, "contacts")
	ut.ConvertByteToIMValue(&req.Addresses, []*pb.Address{}, values, "addresses")
	ut.ConvertByteToIMValue(&req.Events, []*pb.Event{}, values, "events")
	ut.ConvertByteToIMValue(&req.CustomerMeta, &pb.CustomerMeta{}, values, "customer_meta")
	ut.ConvertByteToIMValue(&req.CustomerMap, &pb.JsonString{}, values, "customer_map")

	update := md.Update{Values: values, Model: "customer"}
	if updateID > 0 {
		update.IDKey = updateID
	}

	if updateID, err = ds.StoreDataUpdate(update); err == nil {
		pbCustomer, err = s.CustomerGet(ctx, &pb.RequestGet{Id: updateID, Code: ""})
	}

	return pbCustomer, err
}

func (s *GService) CustomerGet(ctx context.Context, req *pb.RequestGet) (pbCustomer *pb.Customer, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	var customers []cu.IM
	if customers, err = ds.GetDataByID("customer_view", req.Id, req.Code, true); err == nil {
		if customerJson, found := customers[0]["customer_object"].(string); found {
			err = ds.ConvertFromByte([]byte(customerJson), &pbCustomer)
		} else {
			err = errors.New(http.StatusText(http.StatusUnprocessableEntity))
		}
	}
	return pbCustomer, err
}

func (s *GService) CustomerQuery(ctx context.Context, req *pb.RequestQuery) (pbCustomers *pb.Customers, err error) {
	pbCustomers = &pb.Customers{}
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	var params cu.IM = cu.IM{"model": "customer_view", "limit": req.Limit, "offset": req.Offset}
	for _, filter := range req.Filters {
		if slices.Contains([]string{"customer_type", "customer_name", "tag"}, filter.GetField()) {
			params[filter.GetField()] = filter.GetValue()
		} else {
			return pbCustomers, errors.New(http.StatusText(http.StatusUnprocessableEntity) +
				": invalid filter field. Valid fields are: customer_type, customer_name, tag")
		}
	}
	if rows, err := ds.StoreDataGet(params, false); err == nil {
		for _, row := range rows {
			var pbCustomer *pb.Customer
			if customerJson, found := row["customer_object"].(string); found {
				ds.ConvertFromByte([]byte(customerJson), &pbCustomer)
				pbCustomers.Data = append(pbCustomers.Data, pbCustomer)
			}
		}
	}
	return pbCustomers, err
}
