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

func (s *GService) ItemUpdate(ctx context.Context, req *pb.Item) (pbItem *pb.Item, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	user := ctx.Value(md.AuthUserCtxKey).(md.Auth)
	if user.UserGroup == md.UserGroupGuest {
		return pbItem, errors.New(http.StatusText(http.StatusMethodNotAllowed))
	}

	var updateID int64 = req.Id
	if req.Id > 0 || req.Code != "" {
		var rows []cu.IM
		if rows, err = ds.GetDataByID("item", req.Id, req.Code, false); err != nil {
			return pbItem, err
		}
		if len(rows) > 0 {
			updateID = cu.ToInteger(rows[0]["id"], 0)
		}
	}

	if req.TransCode == "" || req.ProductCode == "" || req.TaxCode == "" {
		return pbItem, errors.New(http.StatusText(http.StatusUnprocessableEntity) + ": item trans_code, product_code and tax_code are required")
	}
	values := cu.IM{
		"trans_code":   req.TransCode,
		"product_code": req.ProductCode,
		"tax_code":     req.TaxCode,
	}
	if updateID == 0 && req.Code != "" {
		values["code"] = req.Code
	}

	ut.ConvertByteToIMValue(&req.ItemMeta, &pb.ItemMeta{}, values, "item_meta")
	ut.ConvertByteToIMValue(&req.ItemMap, &pb.JsonString{}, values, "item_map")

	update := md.Update{Values: values, Model: "item"}
	if updateID > 0 {
		update.IDKey = updateID
	}

	if updateID, err = ds.StoreDataUpdate(update); err == nil {
		pbItem, err = s.ItemGet(ctx, &pb.RequestGet{Id: updateID, Code: ""})
	}

	return pbItem, err
}

func (s *GService) ItemGet(ctx context.Context, req *pb.RequestGet) (pbItem *pb.Item, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	var items []cu.IM
	if items, err = ds.GetDataByID("item_view", req.Id, req.Code, true); err == nil {
		if itemJson, found := items[0]["item_object"].(string); found {
			err = ds.ConvertFromByte([]byte(itemJson), &pbItem)
		} else {
			err = errors.New(http.StatusText(http.StatusUnprocessableEntity))
		}
	}
	return pbItem, err
}

func (s *GService) ItemQuery(ctx context.Context, req *pb.RequestQuery) (pbItems *pb.Items, err error) {
	pbItems = &pb.Items{}
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	var params cu.IM = cu.IM{"model": "item_view", "limit": req.Limit, "offset": req.Offset}
	for _, filter := range req.Filters {
		if slices.Contains([]string{"trans_code", "product_code", "tag"}, filter.GetField()) {
			params[filter.GetField()] = filter.GetValue()
		} else {
			return pbItems, errors.New(http.StatusText(http.StatusUnprocessableEntity) +
				": invalid filter field. Valid fields are: trans_code, product_code, tag")
		}
	}
	if rows, err := ds.StoreDataGet(params, false); err == nil {
		for _, row := range rows {
			var pbItem *pb.Item
			if itemJson, found := row["item_object"].(string); found {
				ds.ConvertFromByte([]byte(itemJson), &pbItem)
				pbItems.Data = append(pbItems.Data, pbItem)
			}
		}
	}
	return pbItems, err
}
