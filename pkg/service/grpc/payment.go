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
	pb "github.com/nervatura/nervatura/v6/protos/go"
)

func (s *GService) PaymentUpdate(ctx context.Context, req *pb.Payment) (pbPayment *pb.Payment, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	user := ctx.Value(md.AuthUserCtxKey).(md.Auth)
	if user.UserGroup == md.UserGroupGuest {
		return pbPayment, errors.New(http.StatusText(http.StatusMethodNotAllowed))
	}

	if req.PaidDate == "" || req.TransCode == "" {
		return pbPayment, errors.New(http.StatusText(http.StatusUnprocessableEntity) + ": payment paid_date and trans_code are required")
	}

	var updateID int64 = req.Id
	if req.Id > 0 || req.Code != "" {
		var rows []cu.IM
		if rows, err = ds.GetDataByID("payment", req.Id, req.Code, false); err != nil {
			return pbPayment, err
		}
		if len(rows) > 0 {
			updateID = cu.ToInteger(rows[0]["id"], 0)
		}
	}

	values := cu.IM{
		"paid_date":  req.PaidDate,
		"trans_code": req.TransCode,
	}
	if updateID == 0 && req.Code != "" {
		values["code"] = req.Code
	}

	ut.ConvertByteToIMValue(&req.PaymentMeta, &pb.PaymentMeta{}, values, "payment_meta")
	ut.ConvertByteToIMValue(&req.PaymentMap, &pb.JsonString{}, values, "payment_map")

	update := md.Update{Values: values, Model: "payment"}
	if updateID > 0 {
		update.IDKey = updateID
	}

	if updateID, err = ds.StoreDataUpdate(update); err == nil {
		pbPayment, err = s.PaymentGet(ctx, &pb.RequestGet{Id: updateID, Code: ""})
	}

	return pbPayment, err
}

func (s *GService) PaymentGet(ctx context.Context, req *pb.RequestGet) (pbPayment *pb.Payment, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	var payments []cu.IM
	if payments, err = ds.GetDataByID("payment_view", req.Id, req.Code, true); err == nil {
		if paymentJson, found := payments[0]["payment_object"].(string); found {
			err = ds.ConvertFromByte([]byte(paymentJson), &pbPayment)
		} else {
			err = errors.New(http.StatusText(http.StatusUnprocessableEntity))
		}
	}
	return pbPayment, err
}

func (s *GService) PaymentQuery(ctx context.Context, req *pb.RequestQuery) (pbPayments *pb.Payments, err error) {
	pbPayments = &pb.Payments{}
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	var params cu.IM = cu.IM{"model": "payment_view", "limit": req.Limit, "offset": req.Offset}
	for _, filter := range req.Filters {
		if slices.Contains([]string{"paid_date", "trans_code", "tag"}, filter.GetField()) {
			params[filter.GetField()] = filter.GetValue()
		} else {
			return pbPayments, errors.New(http.StatusText(http.StatusUnprocessableEntity) +
				": invalid filter field. Valid fields are: paid_date, trans_code, tag")
		}
	}
	if rows, err := ds.StoreDataGet(params, false); err == nil {
		for _, row := range rows {
			var pbPayment *pb.Payment
			if paymentJson, found := row["payment_object"].(string); found {
				ds.ConvertFromByte([]byte(paymentJson), &pbPayment)
				pbPayments.Data = append(pbPayments.Data, pbPayment)
			}
		}
	}
	return pbPayments, err
}
