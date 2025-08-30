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

func (s *GService) RateUpdate(ctx context.Context, req *pb.Rate) (pbRate *pb.Rate, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	user := ctx.Value(md.AuthUserCtxKey).(md.Auth)
	if user.UserGroup == md.UserGroupGuest {
		return pbRate, errors.New(http.StatusText(http.StatusMethodNotAllowed))
	}

	if req.RateDate == "" || req.PlaceCode == "" || req.CurrencyCode == "" {
		return pbRate, errors.New(http.StatusText(http.StatusUnprocessableEntity) + ": rate date, place code and currency code are required")
	}

	var updateID int64 = req.Id
	if req.Id > 0 || req.Code != "" {
		var rows []cu.IM
		if rows, err = ds.GetDataByID("rate", req.Id, req.Code, false); err != nil {
			return pbRate, err
		}
		if len(rows) > 0 {
			updateID = cu.ToInteger(rows[0]["id"], 0)
		}
	}

	values := cu.IM{
		"rate_type":     req.RateType.String(),
		"rate_date":     req.RateDate,
		"place_code":    req.PlaceCode,
		"currency_code": req.CurrencyCode,
	}
	if updateID == 0 && req.Code != "" {
		values["code"] = req.Code
	}

	ut.ConvertByteToIMValue(&req.RateMeta, &pb.RateMeta{}, values, "rate_meta")
	ut.ConvertByteToIMValue(&req.RateMap, &pb.JsonString{}, values, "rate_map")

	update := md.Update{Values: values, Model: "rate"}
	if updateID > 0 {
		update.IDKey = updateID
	}

	if updateID, err = ds.StoreDataUpdate(update); err == nil {
		pbRate, err = s.RateGet(ctx, &pb.RequestGet{Id: updateID, Code: ""})
	}

	return pbRate, err
}

func (s *GService) RateGet(ctx context.Context, req *pb.RequestGet) (pbRate *pb.Rate, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	var rates []cu.IM
	if rates, err = ds.GetDataByID("rate_view", req.Id, req.Code, true); err == nil {
		if rateJson, found := rates[0]["rate_object"].(string); found {
			err = ds.ConvertFromByte([]byte(rateJson), &pbRate)
		} else {
			err = errors.New(http.StatusText(http.StatusUnprocessableEntity))
		}
	}
	return pbRate, err
}

func (s *GService) RateQuery(ctx context.Context, req *pb.RequestQuery) (pbRates *pb.Rates, err error) {
	pbRates = &pb.Rates{}
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	var params cu.IM = cu.IM{"model": "rate_view", "limit": req.Limit, "offset": req.Offset}
	for _, filter := range req.Filters {
		if slices.Contains([]string{"rate_type", "currency_code", "tag"}, filter.GetField()) {
			params[filter.GetField()] = filter.GetValue()
		} else {
			return pbRates, errors.New(http.StatusText(http.StatusUnprocessableEntity) +
				": invalid filter field. Valid fields are: rate_type, currency_code, tag")
		}
	}
	if rows, err := ds.StoreDataGet(params, false); err == nil {
		for _, row := range rows {
			var pbRate *pb.Rate
			if rateJson, found := row["rate_object"].(string); found {
				ds.ConvertFromByte([]byte(rateJson), &pbRate)
				pbRates.Data = append(pbRates.Data, pbRate)
			}
		}
	}
	return pbRates, err
}
