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

func (s *GService) TaxUpdate(ctx context.Context, req *pb.Tax) (pbTax *pb.Tax, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	user := ctx.Value(md.AuthUserCtxKey).(md.Auth)
	if user.UserGroup == md.UserGroupGuest {
		return pbTax, errors.New(http.StatusText(http.StatusMethodNotAllowed))
	}

	if req.Code == "" {
		return pbTax, errors.New(http.StatusText(http.StatusUnprocessableEntity) + ": tax code is required")
	}

	var updateID int64 = req.Id
	if req.Id > 0 || req.Code != "" {
		var rows []cu.IM
		if rows, err = ds.GetDataByID("tax", req.Id, req.Code, false); err != nil {
			return pbTax, err
		}
		if len(rows) > 0 {
			updateID = cu.ToInteger(rows[0]["id"], 0)
		}
	}

	values := cu.IM{
		"code": req.Code,
	}

	ut.ConvertByteToIMValue(&req.TaxMeta, &pb.TaxMeta{}, values, "tax_meta")
	ut.ConvertByteToIMValue(&req.TaxMap, &pb.JsonString{}, values, "tax_map")

	update := md.Update{Values: values, Model: "tax"}
	if updateID > 0 {
		update.IDKey = updateID
	}

	if updateID, err = ds.StoreDataUpdate(update); err == nil {
		pbTax, err = s.TaxGet(ctx, &pb.RequestGet{Id: updateID, Code: ""})
	}

	return pbTax, err
}

func (s *GService) TaxGet(ctx context.Context, req *pb.RequestGet) (pbTax *pb.Tax, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	var taxs []cu.IM
	if taxs, err = ds.GetDataByID("tax_view", req.Id, req.Code, true); err == nil {
		if taxJson, found := taxs[0]["tax_object"].(string); found {
			err = ds.ConvertFromByte([]byte(taxJson), &pbTax)
		} else {
			err = errors.New(http.StatusText(http.StatusUnprocessableEntity))
		}
	}
	return pbTax, err
}

func (s *GService) TaxQuery(ctx context.Context, req *pb.RequestQuery) (pbTaxs *pb.Taxes, err error) {
	pbTaxs = &pb.Taxes{}
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	var params cu.IM = cu.IM{"model": "tax_view", "limit": req.Limit, "offset": req.Offset}
	for _, filter := range req.Filters {
		if slices.Contains([]string{"tag"}, filter.GetField()) {
			params[filter.GetField()] = filter.GetValue()
		} else {
			return pbTaxs, errors.New(http.StatusText(http.StatusUnprocessableEntity) +
				": invalid filter field. Valid fields are: tag")
		}
	}
	if rows, err := ds.StoreDataGet(params, false); err == nil {
		for _, row := range rows {
			var pbTax *pb.Tax
			if taxJson, found := row["tax_object"].(string); found {
				ds.ConvertFromByte([]byte(taxJson), &pbTax)
				pbTaxs.Data = append(pbTaxs.Data, pbTax)
			}
		}
	}
	return pbTaxs, err
}
