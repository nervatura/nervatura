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

func (s *GService) MovementUpdate(ctx context.Context, req *pb.Movement) (pbMovement *pb.Movement, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	user := ctx.Value(md.AuthUserCtxKey).(md.Auth)
	if user.UserGroup == md.UserGroupGuest {
		return pbMovement, errors.New(http.StatusText(http.StatusMethodNotAllowed))
	}

	if req.TransCode == "" || req.ShippingTime == "" {
		return pbMovement, errors.New(http.StatusText(http.StatusUnprocessableEntity) + ": movement trans_code and shipping_time are required")
	}
	if req.ProductCode == "" && req.ToolCode == "" {
		return pbMovement, errors.New(http.StatusText(http.StatusUnprocessableEntity) + ": movement product_code or tool_code are required")
	}

	var updateID int64 = req.Id
	if req.Id > 0 || req.Code != "" {
		var rows []cu.IM
		if rows, err = ds.GetDataByID("movement", req.Id, req.Code, false); err != nil {
			return pbMovement, err
		}
		if len(rows) > 0 {
			updateID = cu.ToInteger(rows[0]["id"], 0)
		}
	}

	values := cu.IM{
		"movement_type": req.MovementType.String(),
		"shipping_time": req.ShippingTime,
		"trans_code":    req.TransCode,
	}

	// Optional fields
	optionalFields := map[string]string{
		"code":          req.Code,
		"product_code":  req.ProductCode,
		"tool_code":     req.ToolCode,
		"place_code":    req.PlaceCode,
		"item_code":     req.ItemCode,
		"movement_code": req.MovementCode,
	}
	for key, value := range optionalFields {
		if value != "" {
			values[key] = value
		}
	}

	ut.ConvertByteToIMValue(&req.MovementMeta, &pb.MovementMeta{}, values, "movement_meta")
	ut.ConvertByteToIMValue(&req.MovementMap, &pb.JsonString{}, values, "movement_map")

	update := md.Update{Values: values, Model: "movement"}
	if updateID > 0 {
		update.IDKey = updateID
	}

	if updateID, err = ds.StoreDataUpdate(update); err == nil {
		pbMovement, err = s.MovementGet(ctx, &pb.RequestGet{Id: updateID, Code: ""})
	}

	return pbMovement, err
}

func (s *GService) MovementGet(ctx context.Context, req *pb.RequestGet) (pbMovement *pb.Movement, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	var movements []cu.IM
	if movements, err = ds.GetDataByID("movement_view", req.Id, req.Code, true); err == nil {
		if movementJson, found := movements[0]["movement_object"].(string); found {
			err = ds.ConvertFromByte([]byte(movementJson), &pbMovement)
		} else {
			err = errors.New(http.StatusText(http.StatusUnprocessableEntity))
		}
	}
	return pbMovement, err
}

func (s *GService) MovementQuery(ctx context.Context, req *pb.RequestQuery) (pbMovements *pb.Movements, err error) {
	pbMovements = &pb.Movements{}
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	var params cu.IM = cu.IM{"model": "movement_view", "limit": req.Limit, "offset": req.Offset}
	for _, filter := range req.Filters {
		if slices.Contains([]string{"movement_type", "trans_code", "tag"}, filter.GetField()) {
			params[filter.GetField()] = filter.GetValue()
		} else {
			return pbMovements, errors.New(http.StatusText(http.StatusUnprocessableEntity) +
				": invalid filter field. Valid fields are: movement_type, trans_code, tag")
		}
	}
	if rows, err := ds.StoreDataGet(params, false); err == nil {
		for _, row := range rows {
			var pbMovement *pb.Movement
			if movementJson, found := row["movement_object"].(string); found {
				ds.ConvertFromByte([]byte(movementJson), &pbMovement)
				pbMovements.Data = append(pbMovements.Data, pbMovement)
			}
		}
	}
	return pbMovements, err
}
