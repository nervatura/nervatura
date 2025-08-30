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

func (s *GService) PriceUpdate(ctx context.Context, req *pb.Price) (pbPrice *pb.Price, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	user := ctx.Value(md.AuthUserCtxKey).(md.Auth)
	if user.UserGroup == md.UserGroupGuest {
		return pbPrice, errors.New(http.StatusText(http.StatusMethodNotAllowed))
	}

	if req.ValidFrom == "" || req.CurrencyCode == "" || req.ProductCode == "" {
		return pbPrice, errors.New(http.StatusText(http.StatusUnprocessableEntity) + ": valid from, currency code and product code are required")
	}

	var updateID int64 = req.Id
	if req.Id > 0 || req.Code != "" {
		var rows []cu.IM
		if rows, err = ds.GetDataByID("price", req.Id, req.Code, false); err != nil {
			return pbPrice, err
		}
		if len(rows) > 0 {
			updateID = cu.ToInteger(rows[0]["id"], 0)
		}
	}

	values := cu.IM{
		"valid_from":    req.ValidFrom,
		"product_code":  req.ProductCode,
		"price_type":    req.PriceType.String(),
		"currency_code": req.CurrencyCode,
		"qty":           req.Qty,
	}
	if updateID == 0 && req.Code != "" {
		values["code"] = req.Code
	}
	if req.ValidTo != "" {
		values["valid_to"] = req.ValidTo
	}
	if req.CustomerCode != "" {
		values["customer_code"] = req.CustomerCode
	}

	ut.ConvertByteToIMValue(&req.PriceMeta, &pb.PriceMeta{}, values, "price_meta")
	ut.ConvertByteToIMValue(&req.PriceMap, &pb.JsonString{}, values, "price_map")

	update := md.Update{Values: values, Model: "price"}
	if updateID > 0 {
		update.IDKey = updateID
	}

	if updateID, err = ds.StoreDataUpdate(update); err == nil {
		pbPrice, err = s.PriceGet(ctx, &pb.RequestGet{Id: updateID, Code: ""})
	}

	return pbPrice, err
}

func (s *GService) PriceGet(ctx context.Context, req *pb.RequestGet) (pbPrice *pb.Price, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	var prices []cu.IM
	if prices, err = ds.GetDataByID("price_view", req.Id, req.Code, true); err == nil {
		if priceJson, found := prices[0]["price_object"].(string); found {
			err = ds.ConvertFromByte([]byte(priceJson), &pbPrice)
		} else {
			err = errors.New(http.StatusText(http.StatusUnprocessableEntity))
		}
	}
	return pbPrice, err
}

func (s *GService) PriceQuery(ctx context.Context, req *pb.RequestQuery) (pbPrices *pb.Prices, err error) {
	pbPrices = &pb.Prices{}
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	var params cu.IM = cu.IM{"model": "price_view", "limit": req.Limit, "offset": req.Offset}
	for _, filter := range req.Filters {
		if slices.Contains([]string{"price_type", "product_code", "currency_code", "customer_code", "tag"}, filter.GetField()) {
			params[filter.GetField()] = filter.GetValue()
		} else {
			return pbPrices, errors.New(http.StatusText(http.StatusUnprocessableEntity) +
				": invalid filter field. Valid fields are: price_type, product_code, currency_code, customer_code, tag")
		}
	}
	if rows, err := ds.StoreDataGet(params, false); err == nil {
		for _, row := range rows {
			var pbPrice *pb.Price
			if priceJson, found := row["price_object"].(string); found {
				ds.ConvertFromByte([]byte(priceJson), &pbPrice)
				pbPrices.Data = append(pbPrices.Data, pbPrice)
			}
		}
	}
	return pbPrices, err
}
