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

func (s *GService) ProductUpdate(ctx context.Context, req *pb.Product) (pbProduct *pb.Product, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	user := ctx.Value(md.AuthUserCtxKey).(md.Auth)
	if user.UserGroup == md.UserGroupGuest {
		return pbProduct, errors.New(http.StatusText(http.StatusMethodNotAllowed))
	}

	if req.ProductName == "" || req.TaxCode == "" {
		return pbProduct, errors.New(http.StatusText(http.StatusUnprocessableEntity) + ": product name and tax code are required")
	}

	var updateID int64 = req.Id
	if req.Id > 0 || req.Code != "" {
		var rows []cu.IM
		if rows, err = ds.GetDataByID("product", req.Id, req.Code, false); err != nil {
			return pbProduct, err
		}
		if len(rows) > 0 {
			updateID = cu.ToInteger(rows[0]["id"], 0)
		}
	}

	values := cu.IM{
		"product_type": req.ProductType.String(),
		"product_name": req.ProductName,
		"tax_code":     req.TaxCode,
	}
	if updateID == 0 && req.Code != "" {
		values["code"] = req.Code
	}

	ut.ConvertByteToIMValue(&req.Events, []*pb.Event{}, values, "events")
	ut.ConvertByteToIMValue(&req.ProductMeta, &pb.ProductMeta{}, values, "product_meta")
	ut.ConvertByteToIMValue(&req.ProductMap, &pb.JsonString{}, values, "product_map")

	update := md.Update{Values: values, Model: "product"}
	if updateID > 0 {
		update.IDKey = updateID
	}

	if updateID, err = ds.StoreDataUpdate(update); err == nil {
		pbProduct, err = s.ProductGet(ctx, &pb.RequestGet{Id: updateID, Code: ""})
	}

	return pbProduct, err
}

func (s *GService) ProductGet(ctx context.Context, req *pb.RequestGet) (pbProduct *pb.Product, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	var products []cu.IM
	if products, err = ds.GetDataByID("product_view", req.Id, req.Code, true); err == nil {
		if productJson, found := products[0]["product_object"].(string); found {
			err = ds.ConvertFromByte([]byte(productJson), &pbProduct)
		} else {
			err = errors.New(http.StatusText(http.StatusUnprocessableEntity))
		}
	}
	return pbProduct, err
}

func (s *GService) ProductQuery(ctx context.Context, req *pb.RequestQuery) (pbProducts *pb.Products, err error) {
	pbProducts = &pb.Products{}
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	var params cu.IM = cu.IM{"model": "product_view", "limit": req.Limit, "offset": req.Offset}
	for _, filter := range req.Filters {
		if slices.Contains([]string{"product_type", "product_name", "tag"}, filter.GetField()) {
			params[filter.GetField()] = filter.GetValue()
		} else {
			return pbProducts, errors.New(http.StatusText(http.StatusUnprocessableEntity) +
				": invalid filter field. Valid fields are: product_type, product_name, tag")
		}
	}
	if rows, err := ds.StoreDataGet(params, false); err == nil {
		for _, row := range rows {
			var pbProduct *pb.Product
			if productJson, found := row["product_object"].(string); found {
				ds.ConvertFromByte([]byte(productJson), &pbProduct)
				pbProducts.Data = append(pbProducts.Data, pbProduct)
			}
		}
	}
	return pbProducts, err
}
