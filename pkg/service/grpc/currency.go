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

func (s *GService) CurrencyUpdate(ctx context.Context, req *pb.Currency) (pbCurrency *pb.Currency, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	user := ctx.Value(md.AuthUserCtxKey).(md.Auth)
	if user.UserGroup == md.UserGroupGuest {
		return pbCurrency, errors.New(http.StatusText(http.StatusMethodNotAllowed))
	}

	var updateID int64 = req.Id
	if req.Id > 0 || req.Code != "" {
		var rows []cu.IM
		if rows, err = ds.GetDataByID("currency_view", req.Id, req.Code, false); err != nil {
			return pbCurrency, err
		}
		if len(rows) > 0 {
			updateID = cu.ToInteger(rows[0]["id"], 0)
		}
	}

	if req.Code == "" {
		return pbCurrency, errors.New(http.StatusText(http.StatusUnprocessableEntity) + ": code is required")
	}
	values := cu.IM{
		"code": req.Code,
	}

	ut.ConvertByteToIMValue(&req.CurrencyMeta, &pb.CurrencyMeta{}, values, "currency_meta")
	ut.ConvertByteToIMValue(&req.CurrencyMap, &pb.JsonString{}, values, "currency_map")

	update := md.Update{Values: values, Model: "currency"}
	if updateID > 0 {
		update.IDKey = updateID
	}

	if updateID, err = ds.StoreDataUpdate(update); err == nil {
		pbCurrency, err = s.CurrencyGet(ctx, &pb.RequestGet{Id: updateID, Code: ""})
	}

	return pbCurrency, err
}

func (s *GService) CurrencyGet(ctx context.Context, req *pb.RequestGet) (pbCurrency *pb.Currency, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	var currencys []cu.IM
	if currencys, err = ds.GetDataByID("currency_view", req.Id, req.Code, true); err == nil {
		if currencyJson, found := currencys[0]["currency_object"].(string); found {
			err = ds.ConvertFromByte([]byte(currencyJson), &pbCurrency)
		} else {
			err = errors.New(http.StatusText(http.StatusUnprocessableEntity))
		}
	}
	return pbCurrency, err
}

func (s *GService) CurrencyQuery(ctx context.Context, req *pb.RequestQuery) (pbCurrencies *pb.Currencies, err error) {
	pbCurrencies = &pb.Currencies{}
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	var params cu.IM = cu.IM{"model": "currency_view", "limit": req.Limit, "offset": req.Offset}
	for _, filter := range req.Filters {
		if slices.Contains([]string{"tag"}, filter.GetField()) {
			params[filter.GetField()] = filter.GetValue()
		} else {
			return pbCurrencies, errors.New(http.StatusText(http.StatusUnprocessableEntity) +
				": invalid filter field. Valid fields are: tag")
		}
	}
	if rows, err := ds.StoreDataGet(params, false); err == nil {
		for _, row := range rows {
			var pbCurrency *pb.Currency
			if currencyJson, found := row["currency_object"].(string); found {
				ds.ConvertFromByte([]byte(currencyJson), &pbCurrency)
				pbCurrencies.Data = append(pbCurrencies.Data, pbCurrency)
			}
		}
	}
	return pbCurrencies, err
}
