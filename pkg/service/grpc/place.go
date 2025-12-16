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

func (s *GService) PlaceUpdate(ctx context.Context, req *pb.Place) (pbPlace *pb.Place, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	user := ctx.Value(md.AuthUserCtxKey).(md.Auth)
	if user.UserGroup == md.UserGroupGuest {
		return pbPlace, errors.New(http.StatusText(http.StatusMethodNotAllowed))
	}

	if req.PlaceName == "" {
		return pbPlace, errors.New(http.StatusText(http.StatusUnprocessableEntity) + ": place name is required")
	}

	var updateID int64 = req.Id
	if req.Id > 0 || req.Code != "" {
		var rows []cu.IM
		if rows, err = ds.GetDataByID("place", req.Id, req.Code, false); err != nil {
			return pbPlace, err
		}
		if len(rows) > 0 {
			updateID = cu.ToInteger(rows[0]["id"], 0)
		}
	}

	values := cu.IM{
		"place_type": req.PlaceType.String(),
		"place_name": req.PlaceName,
	}
	if updateID == 0 && req.Code != "" {
		values["code"] = req.Code
	}
	if req.CurrencyCode != "" {
		values["currency_code"] = req.CurrencyCode
	}

	ut.ConvertByteToIMValue(&req.Contacts, []*pb.Contact{}, values, "contacts")
	ut.ConvertByteToIMValue(&req.Address, &pb.Address{}, values, "address")
	ut.ConvertByteToIMValue(&req.Events, []*pb.Event{}, values, "events")
	ut.ConvertByteToIMValue(&req.PlaceMeta, &pb.PlaceMeta{}, values, "place_meta")
	ut.ConvertByteToIMValue(&req.PlaceMap, &pb.JsonString{}, values, "place_map")

	update := md.Update{Values: values, Model: "place"}
	if updateID > 0 {
		update.IDKey = updateID
	}

	if updateID, err = ds.StoreDataUpdate(update); err == nil {
		pbPlace, err = s.PlaceGet(ctx, &pb.RequestGet{Id: updateID, Code: ""})
	}

	return pbPlace, err
}

func (s *GService) PlaceGet(ctx context.Context, req *pb.RequestGet) (pbPlace *pb.Place, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	var places []cu.IM
	if places, err = ds.GetDataByID("place_view", req.Id, req.Code, true); err == nil {
		if placeJson, found := places[0]["place_object"].(string); found {
			err = ds.ConvertFromByte([]byte(placeJson), &pbPlace)
		} else {
			err = errors.New(http.StatusText(http.StatusUnprocessableEntity))
		}
	}
	return pbPlace, err
}

func (s *GService) PlaceQuery(ctx context.Context, req *pb.RequestQuery) (pbPlaces *pb.Places, err error) {
	pbPlaces = &pb.Places{}
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	var params cu.IM = cu.IM{"model": "place_view", "limit": req.Limit, "offset": req.Offset}
	for _, filter := range req.Filters {
		if slices.Contains([]string{"place_type", "place_name", "tag"}, filter.GetField()) {
			params[filter.GetField()] = filter.GetValue()
		} else {
			return pbPlaces, errors.New(http.StatusText(http.StatusUnprocessableEntity) +
				": invalid filter field. Valid fields are: place_type, place_name, tag")
		}
	}
	if rows, err := ds.StoreDataGet(params, false); err == nil {
		for _, row := range rows {
			var pbPlace *pb.Place
			if placeJson, found := row["place_object"].(string); found {
				ds.ConvertFromByte([]byte(placeJson), &pbPlace)
				pbPlaces.Data = append(pbPlaces.Data, pbPlace)
			}
		}
	}
	return pbPlaces, err
}
