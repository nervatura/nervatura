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

func (s *GService) LinkUpdate(ctx context.Context, req *pb.Link) (pbLink *pb.Link, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	user := ctx.Value(md.AuthUserCtxKey).(md.Auth)
	if user.UserGroup == md.UserGroupGuest {
		return pbLink, errors.New(http.StatusText(http.StatusMethodNotAllowed))
	}

	var updateID int64 = req.Id
	if req.Id > 0 || req.Code != "" {
		var rows []cu.IM
		if rows, err = ds.GetDataByID("link", req.Id, req.Code, false); err != nil {
			return pbLink, err
		}
		if len(rows) > 0 {
			updateID = cu.ToInteger(rows[0]["id"], 0)
		}
	}

	if req.LinkCode_1 == "" || req.LinkCode_2 == "" {
		return pbLink, errors.New(http.StatusText(http.StatusUnprocessableEntity) + ": link link_code_1 and link_code_2 are required")
	}
	values := cu.IM{
		"link_type_1": req.LinkType_1.String(),
		"link_code_1": req.LinkCode_1,
		"link_type_2": req.LinkType_2.String(),
		"link_code_2": req.LinkCode_2,
	}
	if updateID == 0 && req.Code != "" {
		values["code"] = req.Code
	}

	ut.ConvertByteToIMValue(&req.LinkMeta, &pb.LinkMeta{}, values, "link_meta")
	ut.ConvertByteToIMValue(&req.LinkMap, &pb.JsonString{}, values, "link_map")

	update := md.Update{Values: values, Model: "link"}
	if updateID > 0 {
		update.IDKey = updateID
	}

	if updateID, err = ds.StoreDataUpdate(update); err == nil {
		pbLink, err = s.LinkGet(ctx, &pb.RequestGet{Id: updateID, Code: ""})
	}

	return pbLink, err
}

func (s *GService) LinkGet(ctx context.Context, req *pb.RequestGet) (pbLink *pb.Link, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	var links []cu.IM
	if links, err = ds.GetDataByID("link_view", req.Id, req.Code, true); err == nil {
		if linkJson, found := links[0]["link_object"].(string); found {
			err = ds.ConvertFromByte([]byte(linkJson), &pbLink)
		} else {
			err = errors.New(http.StatusText(http.StatusUnprocessableEntity))
		}
	}
	return pbLink, err
}

func (s *GService) LinkQuery(ctx context.Context, req *pb.RequestQuery) (pbLinks *pb.Links, err error) {
	pbLinks = &pb.Links{}
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	var params cu.IM = cu.IM{"model": "link_view", "limit": req.Limit, "offset": req.Offset}
	for _, filter := range req.Filters {
		if slices.Contains([]string{"link_type_1", "link_code_1", "link_type_2", "link_code_2", "tag"}, filter.GetField()) {
			params[filter.GetField()] = filter.GetValue()
		} else {
			return pbLinks, errors.New(http.StatusText(http.StatusUnprocessableEntity) +
				": invalid filter field. Valid fields are: link_type_1, link_code_1, link_type_2, link_code_2, tag")
		}
	}
	if rows, err := ds.StoreDataGet(params, false); err == nil {
		for _, row := range rows {
			var pbLink *pb.Link
			if linkJson, found := row["link_object"].(string); found {
				ds.ConvertFromByte([]byte(linkJson), &pbLink)
				pbLinks.Data = append(pbLinks.Data, pbLink)
			}
		}
	}
	return pbLinks, err
}
