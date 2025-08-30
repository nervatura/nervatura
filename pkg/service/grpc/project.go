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

func (s *GService) ProjectUpdate(ctx context.Context, req *pb.Project) (pbProject *pb.Project, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	user := ctx.Value(md.AuthUserCtxKey).(md.Auth)
	if user.UserGroup == md.UserGroupGuest {
		return pbProject, errors.New(http.StatusText(http.StatusMethodNotAllowed))
	}

	if req.ProjectName == "" {
		return pbProject, errors.New(http.StatusText(http.StatusUnprocessableEntity) + ": project name is required")
	}

	var updateID int64 = req.Id
	if req.Id > 0 || req.Code != "" {
		var rows []cu.IM
		if rows, err = ds.GetDataByID("project", req.Id, req.Code, false); err != nil {
			return pbProject, err
		}
		if len(rows) > 0 {
			updateID = cu.ToInteger(rows[0]["id"], 0)
		}
	}
	values := cu.IM{
		"project_name": req.ProjectName,
	}
	if updateID == 0 && req.Code != "" {
		values["code"] = req.Code
	}
	if req.CustomerCode != "" {
		values["customer_code"] = req.CustomerCode
	}

	ut.ConvertByteToIMValue(&req.Contacts, []*pb.Contact{}, values, "contacts")
	ut.ConvertByteToIMValue(&req.Addresses, []*pb.Address{}, values, "addresses")
	ut.ConvertByteToIMValue(&req.Events, []*pb.Event{}, values, "events")
	ut.ConvertByteToIMValue(&req.ProjectMeta, &pb.ProjectMeta{Tags: []string{}}, values, "project_meta")
	ut.ConvertByteToIMValue(&req.ProjectMap, &pb.JsonString{}, values, "project_map")

	update := md.Update{Values: values, Model: "project"}
	if updateID > 0 {
		update.IDKey = updateID
	}

	if updateID, err = ds.StoreDataUpdate(update); err == nil {
		pbProject, err = s.ProjectGet(ctx, &pb.RequestGet{Id: updateID, Code: ""})
	}

	return pbProject, err
}

func (s *GService) ProjectGet(ctx context.Context, req *pb.RequestGet) (pbProject *pb.Project, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	var projects []cu.IM
	if projects, err = ds.GetDataByID("project_view", req.Id, req.Code, true); err == nil {
		if projectJson, found := projects[0]["project_object"].(string); found {
			err = ds.ConvertFromByte([]byte(projectJson), &pbProject)
		} else {
			err = errors.New(http.StatusText(http.StatusUnprocessableEntity))
		}
	}
	return pbProject, err
}

func (s *GService) ProjectQuery(ctx context.Context, req *pb.RequestQuery) (pbProjects *pb.Projects, err error) {
	pbProjects = &pb.Projects{}
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	var params cu.IM = cu.IM{"model": "project_view", "limit": req.Limit, "offset": req.Offset}
	for _, filter := range req.Filters {
		if slices.Contains([]string{"project_name", "customer_code", "tag"}, filter.GetField()) {
			params[filter.GetField()] = filter.GetValue()
		} else {
			return pbProjects, errors.New(http.StatusText(http.StatusUnprocessableEntity) +
				": invalid filter field. Valid fields are: project_name, customer_code, tag")
		}
	}
	if rows, err := ds.StoreDataGet(params, false); err == nil {
		for _, row := range rows {
			var pbProject *pb.Project
			if projectJson, found := row["project_object"].(string); found {
				ds.ConvertFromByte([]byte(projectJson), &pbProject)
				pbProjects.Data = append(pbProjects.Data, pbProject)
			}
		}
	}
	return pbProjects, err
}
