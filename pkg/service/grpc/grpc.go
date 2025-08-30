package grpc

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"slices"
	"strings"

	cu "github.com/nervatura/component/pkg/util"
	api "github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	pb "github.com/nervatura/nervatura/v6/pkg/service/grpc/proto"
)

// GService implements the Nervatura gRPC API service
type GService struct {
	Config          cu.IM
	AppLog          *slog.Logger
	DecodeAuthToken func(tokenString string) (cu.IM, error)
	pb.UnimplementedAPIServer
}

func (s *GService) TokenAuth(authorization []string, parent context.Context) (ctx context.Context, err error) {
	if len(authorization) < 1 {
		return ctx, errors.New(http.StatusText(http.StatusUnauthorized))
	}
	tokenStr := strings.TrimPrefix(authorization[0], "Bearer ")
	if tokenStr == "" {
		return ctx, errors.New(http.StatusText(http.StatusUnauthorized))
	}
	claim, err := s.DecodeAuthToken(tokenStr)
	if err != nil {
		return ctx, err
	}

	alias := ""
	if _, found := claim["alias"]; found {
		alias = claim["alias"].(string)
	}
	ds := api.NewDataStore(s.Config, alias, s.AppLog)
	ctx = context.WithValue(parent, md.DataStoreCtxKey, ds)

	user, err := ds.TokenLogin(tokenStr)
	if err != nil {
		return ctx, err
	}
	ctx = context.WithValue(ctx, md.AuthUserCtxKey, user)
	return ctx, nil
}

func (s *GService) ApiKeyAuth(authorization []string, parent context.Context) (ctx context.Context, err error) {
	if len(authorization) < 1 {
		return ctx, errors.New(http.StatusText(http.StatusUnauthorized))
	}
	apiKey := strings.Trim(authorization[0], " ")
	if apiKey == "" {
		return ctx, errors.New(http.StatusText(http.StatusUnauthorized))
	}
	if cu.ToString(s.Config["NT_API_KEY"], "") != apiKey {
		return ctx, errors.New(http.StatusText(http.StatusUnauthorized))
	}

	alias := cu.ToString(s.Config["NT_DEFAULT_ALIAS"], "")
	userCode := cu.ToString(s.Config["NT_DEFAULT_ADMIN"], "")
	ds := api.NewDataStore(s.Config, alias, s.AppLog)
	var user md.Auth = md.Auth{
		UserGroup: md.UserGroupAdmin,
		Code:      userCode,
		UserName:  "admin",
	}
	ctx = context.WithValue(parent, md.DataStoreCtxKey, ds)
	ctx = context.WithValue(ctx, md.AuthUserCtxKey, user)
	return ctx, nil
}

func (s *GService) Delete(ctx context.Context, req *pb.RequestDelete) (pbStatus *pb.ResponseStatus, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	user := ctx.Value(md.AuthUserCtxKey).(md.Auth)
	if user.UserGroup == md.UserGroupGuest {
		return &pb.ResponseStatus{Success: false, Message: http.StatusText(http.StatusMethodNotAllowed)},
			errors.New(http.StatusText(http.StatusMethodNotAllowed))
	}
	pbStatus = &pb.ResponseStatus{Success: true, Message: http.StatusText(http.StatusOK)}
	switch req.Model {
	case pb.Model_AUTH:
		if req.Id == 1 || req.Code == "USR0000000000N1" {
			return &pb.ResponseStatus{Success: false, Message: "Cannot disable master admin user"}, nil
		}
		var pbAuth *pb.Auth
		if pbAuth, err = s.AuthGet(ctx, &pb.RequestGet{Id: req.Id, Code: req.Code}); err != nil {
			return &pb.ResponseStatus{Success: false, Message: err.Error()}, err
		}
		_, err = ds.StoreDataUpdate(md.Update{Model: "auth", IDKey: pbAuth.Id, Values: cu.IM{"disabled": true}})
	case pb.Model_CONFIG:
		var pbConfig *pb.Config
		if pbConfig, err = s.ConfigGet(ctx, &pb.RequestGet{Id: req.Id, Code: req.Code}); err != nil {
			return &pb.ResponseStatus{Success: false, Message: err.Error()}, err
		}
		if !slices.Contains([]string{"CONFIG_PRINT_QUEUE", "CONFIG_PATTERN"}, pbConfig.ConfigType.String()) &&
			user.UserGroup != md.UserGroupAdmin {
			return &pb.ResponseStatus{Success: false, Message: http.StatusText(http.StatusMethodNotAllowed)},
				errors.New(http.StatusText(http.StatusMethodNotAllowed))
		}
		_, err = ds.StoreDataUpdate(md.Update{Model: "config", IDKey: pbConfig.Id})
	default:
		model := strings.ToLower(req.Model.String())
		err = ds.DataDelete(model, req.Id, req.Code)
	}
	if err != nil {
		pbStatus = &pb.ResponseStatus{Success: false, Message: err.Error()}
	}
	return pbStatus, err
}

// Call a server side function and get the result
// Example: create new PDF invoice, send email or get a product price
func (s *GService) Function(ctx context.Context, req *pb.RequestFunction) (*pb.JsonBytes, error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	options := cu.IM{}
	for key, value := range req.Args {
		options[key] = value
	}
	var err error
	var response interface{}
	var result []byte
	if response, err = ds.Function(req.Function, options); err == nil {
		result, err = cu.ConvertToByte(response)
	}
	return &pb.JsonBytes{Data: result}, err
}

// Create new nervatura database schema
func (s *GService) Database(ctx context.Context, req *pb.RequestDatabase) (*pb.JsonBytes, error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	options := cu.IM{"alias": req.Alias, "demo": req.Demo}
	response := api.CreateDatabase(options, ds.Config)
	var err error
	var result []byte
	result, err = cu.ConvertToByte(response)
	return &pb.JsonBytes{Data: result}, err
}

// Get a predefined view by name
func (s *GService) View(ctx context.Context, req *pb.RequestView) (response *pb.JsonBytes, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)

	query := md.Query{
		Fields:  []string{"*"},
		From:    strings.ToLower(strings.TrimPrefix(req.Name.String(), "VIEW_")),
		Filter:  req.Filter,
		Filters: []md.Filter{},
		OrderBy: req.OrderBy,
		Limit:   req.Limit,
		Offset:  req.Offset,
	}

	var rows []cu.IM
	response = &pb.JsonBytes{Data: []byte{}}
	if rows, err = ds.StoreDataQuery(query, false); err == nil {
		response.Data, err = cu.ConvertToByte(rows)
	}
	return response, err
}
