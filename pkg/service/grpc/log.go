package grpc

import (
	"context"
	"errors"
	"net/http"
	"slices"

	cu "github.com/nervatura/component/pkg/util"
	api "github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	pb "github.com/nervatura/nervatura/v6/protos/go"
)

func (s *GService) LogGet(ctx context.Context, req *pb.RequestGet) (pbLog *pb.Log, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	var logs []cu.IM
	if logs, err = ds.GetDataByID("log_view", req.Id, req.Code, true); err == nil {
		if logJson, found := logs[0]["log_object"].(string); found {
			err = ds.ConvertFromByte([]byte(logJson), &pbLog)
		} else {
			err = errors.New(http.StatusText(http.StatusUnprocessableEntity))
		}
	}
	return pbLog, err
}

func (s *GService) LogQuery(ctx context.Context, req *pb.RequestQuery) (pbLogs *pb.Logs, err error) {
	pbLogs = &pb.Logs{}
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	var params cu.IM = cu.IM{"model": "log_view", "limit": req.Limit, "offset": req.Offset}
	for _, filter := range req.Filters {
		if slices.Contains([]string{"log_type", "ref_type", "ref_code", "tag"}, filter.GetField()) {
			params[filter.GetField()] = filter.GetValue()
		} else {
			return pbLogs, errors.New(http.StatusText(http.StatusUnprocessableEntity) +
				": invalid filter field. Valid fields are: log_type, ref_type, ref_code, tag")
		}
	}
	if rows, err := ds.StoreDataGet(params, false); err == nil {
		for _, row := range rows {
			var pbLog *pb.Log
			if logJson, found := row["log_object"].(string); found {
				ds.ConvertFromByte([]byte(logJson), &pbLog)
				pbLogs.Data = append(pbLogs.Data, pbLog)
			}
		}
	}
	return pbLogs, err
}
