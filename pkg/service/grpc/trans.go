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

func (s *GService) TransUpdate(ctx context.Context, req *pb.Trans) (pbTrans *pb.Trans, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	user := ctx.Value(md.AuthUserCtxKey).(md.Auth)
	if user.UserGroup == md.UserGroupGuest {
		return pbTrans, errors.New(http.StatusText(http.StatusMethodNotAllowed))
	}

	if req.TransDate == "" {
		return pbTrans, errors.New(http.StatusText(http.StatusUnprocessableEntity) + ": trans date is required")
	}

	if slices.Contains([]pb.TransType{
		pb.TransType_TRANS_INVOICE, pb.TransType_TRANS_RECEIPT, pb.TransType_TRANS_OFFER, pb.TransType_TRANS_ORDER, pb.TransType_TRANS_WORKSHEET, pb.TransType_TRANS_RENT}, req.TransType,
	) && (req.CustomerCode == "" || req.CurrencyCode == "") {
		err = errors.New("invoice, receipt, offer, order, worksheet and rent must have customer code and currency code")
		return pbTrans, err
	}

	var updateID int64 = req.Id
	if req.Id > 0 || req.Code != "" {
		var rows []cu.IM
		if rows, err = ds.GetDataByID("trans", req.Id, req.Code, false); err != nil {
			return pbTrans, err
		}
		if len(rows) > 0 {
			updateID = cu.ToInteger(rows[0]["id"], 0)
		}
	}

	values := cu.IM{
		"trans_type": req.TransType.String(),
		"direction":  req.Direction.String(),
		"trans_date": req.TransDate,
		"auth_code":  user.Code,
	}

	// Optional fields
	optionalFields := map[string]string{
		"code":          req.Code,
		"customer_code": req.CustomerCode,
		"employee_code": req.EmployeeCode,
		"project_code":  req.ProjectCode,
		"place_code":    req.PlaceCode,
		"trans_code":    req.TransCode,
		"currency_code": req.CurrencyCode,
	}

	for key, value := range optionalFields {
		if value != "" {
			values[key] = value
		}
	}

	ut.ConvertByteToIMValue(&req.TransMeta, &pb.TransMeta{}, values, "trans_meta")
	ut.ConvertByteToIMValue(&req.TransMap, &pb.JsonString{}, values, "trans_map")

	update := md.Update{Values: values, Model: "trans"}
	if updateID > 0 {
		update.IDKey = updateID
	}

	if updateID, err = ds.StoreDataUpdate(update); err == nil {
		pbTrans, err = s.TransGet(ctx, &pb.RequestGet{Id: updateID, Code: ""})
	}

	return pbTrans, err
}

func (s *GService) TransGet(ctx context.Context, req *pb.RequestGet) (pbTrans *pb.Trans, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	var transs []cu.IM
	if transs, err = ds.GetDataByID("trans_view", req.Id, req.Code, true); err == nil {
		if transJson, found := transs[0]["trans_object"].(string); found {
			err = ds.ConvertFromByte([]byte(transJson), &pbTrans)
		} else {
			err = errors.New(http.StatusText(http.StatusUnprocessableEntity))
		}
	}
	return pbTrans, err
}

func (s *GService) TransQuery(ctx context.Context, req *pb.RequestQuery) (pbTransactions *pb.Transactions, err error) {
	pbTransactions = &pb.Transactions{}
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	var params cu.IM = cu.IM{"model": "trans_view", "limit": req.Limit, "offset": req.Offset}
	for _, filter := range req.Filters {
		if slices.Contains([]string{"trans_type", "direction", "trans_date", "tag"}, filter.GetField()) {
			params[filter.GetField()] = filter.GetValue()
		} else {
			return pbTransactions, errors.New(http.StatusText(http.StatusUnprocessableEntity) +
				": invalid filter field. Valid fields are: trans_type, direction, trans_date, tag")
		}
	}
	if rows, err := ds.StoreDataGet(params, false); err == nil {
		for _, row := range rows {
			var pbTrans *pb.Trans
			if transJson, found := row["trans_object"].(string); found {
				ds.ConvertFromByte([]byte(transJson), &pbTrans)
				pbTransactions.Data = append(pbTransactions.Data, pbTrans)
			}
		}
	}
	return pbTransactions, err
}
