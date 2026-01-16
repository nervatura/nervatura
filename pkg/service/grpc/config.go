package grpc

import (
	"context"
	"errors"
	"net/http"
	"slices"

	cu "github.com/nervatura/component/pkg/util"
	api "github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	pb "github.com/nervatura/nervatura/v6/proto"
)

func (s *GService) ConfigUpdate(ctx context.Context, req *pb.Config) (pbConfig *pb.Config, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	user := ctx.Value(md.AuthUserCtxKey).(md.Auth)

	if user.UserGroup == md.UserGroupGuest {
		return pbConfig, errors.New(http.StatusText(http.StatusMethodNotAllowed))
	}

	if !slices.Contains([]string{"CONFIG_PRINT_QUEUE", "CONFIG_PATTERN"}, req.ConfigType.String()) &&
		user.UserGroup != md.UserGroupAdmin {
		return pbConfig, errors.New(http.StatusText(http.StatusMethodNotAllowed))
	}

	var updateID int64 = req.Id
	if req.Id > 0 || req.Code != "" {
		if pbConfig, err = s.ConfigGet(ctx, &pb.RequestGet{Id: req.Id, Code: req.Code}); err != nil {
			return pbConfig, err
		}
		updateID = pbConfig.Id
	}

	values := cu.IM{
		"config_type": req.ConfigType.String(),
	}
	if updateID == 0 && req.Code != "" {
		values["code"] = req.Code
	}

	if req.Data == nil {
		return pbConfig, errors.New(http.StatusText(http.StatusUnprocessableEntity) +
			": data is required")
	}
	if metaData, err := ds.ConvertToByte(req.Data); err == nil {
		values["data"] = string(metaData[:])
	}

	update := md.Update{
		Values: values,
		Model:  "config",
	}
	if updateID > 0 {
		update.IDKey = updateID
	}

	if updateID, err = ds.StoreDataUpdate(update); err == nil {
		pbConfig, err = s.ConfigGet(ctx, &pb.RequestGet{Id: updateID, Code: ""})
	}

	return pbConfig, err
}

func (s *GService) ConfigGet(ctx context.Context, req *pb.RequestGet) (pbConfig *pb.Config, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	query := md.Query{
		Fields: []string{"*"}, From: "config", Filters: []md.Filter{},
	}
	if req.Id > 0 {
		query.Filters = append(query.Filters, md.Filter{Field: "id", Comp: "==", Value: req.Id})
	}
	if req.Code != "" {
		query.Filters = append(query.Filters, md.Filter{Field: "code", Comp: "==", Value: req.Code})
	}
	var configs []cu.IM
	if configs, err = ds.StoreDataQuery(query, false); err == nil {
		if len(configs) == 0 {
			return pbConfig, errors.New(http.StatusText(http.StatusNotFound) + ": config not found")
		}
		err = ds.ConvertToType(configs[0], &pbConfig)
	}
	return pbConfig, err
}

func (s *GService) ConfigQuery(ctx context.Context, req *pb.RequestQuery) (pbConfigs *pb.Configs, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	pbConfigs = &pb.Configs{}
	query := md.Query{
		Fields: []string{"*"}, From: "config", Filters: []md.Filter{},
		Limit: req.Limit, Offset: req.Offset,
	}
	for _, filter := range req.Filters {
		if slices.Contains([]string{"config_type"}, filter.GetField()) {
			query.Filters = append(query.Filters, md.Filter{
				Field: filter.GetField(), Comp: "==", Value: filter.GetValue(),
			})
		} else {
			return pbConfigs, errors.New(http.StatusText(http.StatusUnprocessableEntity) +
				": invalid filter field. Valid fields are: config_type")
		}
	}
	var rows []cu.IM
	if rows, err = ds.StoreDataQuery(query, false); err == nil {
		var row cu.IM
		for _, row = range rows {
			var pbConfig *pb.Config
			if err = ds.ConvertToType(row, &pbConfig); err == nil {
				pbConfigs.Data = append(pbConfigs.Data, pbConfig)
			}
		}
	}
	return pbConfigs, err
}
