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
	pb "github.com/nervatura/nervatura/v6/protos/go"
)

func (s *GService) ToolUpdate(ctx context.Context, req *pb.Tool) (pbTool *pb.Tool, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	user := ctx.Value(md.AuthUserCtxKey).(md.Auth)
	if user.UserGroup == md.UserGroupGuest {
		return pbTool, errors.New(http.StatusText(http.StatusMethodNotAllowed))
	}

	if req.ProductCode == "" || req.Description == "" {
		return pbTool, errors.New(http.StatusText(http.StatusUnprocessableEntity) + ": product code and description are required")
	}

	var updateID int64 = req.Id
	if req.Id > 0 || req.Code != "" {
		var rows []cu.IM
		if rows, err = ds.GetDataByID("tool", req.Id, req.Code, false); err != nil {
			return pbTool, err
		}
		if len(rows) > 0 {
			updateID = cu.ToInteger(rows[0]["id"], 0)
		}
	}

	values := cu.IM{
		"product_code": req.ProductCode,
		"description":  req.Description,
	}
	if updateID == 0 && req.Code != "" {
		values["code"] = req.Code
	}

	ut.ConvertByteToIMValue(&req.Events, []*pb.Event{}, values, "events")
	ut.ConvertByteToIMValue(&req.ToolMeta, &pb.ToolMeta{}, values, "tool_meta")
	ut.ConvertByteToIMValue(&req.ToolMap, &pb.JsonString{}, values, "tool_map")

	update := md.Update{Values: values, Model: "tool"}
	if updateID > 0 {
		update.IDKey = updateID
	}

	if updateID, err = ds.StoreDataUpdate(update); err == nil {
		pbTool, err = s.ToolGet(ctx, &pb.RequestGet{Id: updateID, Code: ""})
	}

	return pbTool, err
}

func (s *GService) ToolGet(ctx context.Context, req *pb.RequestGet) (pbTool *pb.Tool, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	var tools []cu.IM
	if tools, err = ds.GetDataByID("tool_view", req.Id, req.Code, true); err == nil {
		if toolJson, found := tools[0]["tool_object"].(string); found {
			err = ds.ConvertFromByte([]byte(toolJson), &pbTool)
		} else {
			err = errors.New(http.StatusText(http.StatusUnprocessableEntity))
		}
	}
	return pbTool, err
}

func (s *GService) ToolQuery(ctx context.Context, req *pb.RequestQuery) (pbTools *pb.Tools, err error) {
	pbTools = &pb.Tools{}
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	var params cu.IM = cu.IM{"model": "tool_view", "limit": req.Limit, "offset": req.Offset}
	for _, filter := range req.Filters {
		if slices.Contains([]string{"product_code", "description", "tag"}, filter.GetField()) {
			params[filter.GetField()] = filter.GetValue()
		} else {
			return pbTools, errors.New(http.StatusText(http.StatusUnprocessableEntity) +
				": invalid filter field. Valid fields are: product_code, description, tag")
		}
	}
	if rows, err := ds.StoreDataGet(params, false); err == nil {
		for _, row := range rows {
			var pbTool *pb.Tool
			if toolJson, found := row["tool_object"].(string); found {
				ds.ConvertFromByte([]byte(toolJson), &pbTool)
				pbTools.Data = append(pbTools.Data, pbTool)
			}
		}
	}
	return pbTools, err
}
