package grpc

import (
	"context"
	"errors"
	"net/http"

	cu "github.com/nervatura/component/pkg/util"
	api "github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
	pb "github.com/nervatura/nervatura/v6/proto"
)

// Login by username and password and get access token
func (s *GService) AuthLogin(ctx context.Context, req *pb.RequestAuthLogin) (login *pb.ResponseAuthLogin, err error) {
	ds := api.NewDataStore(s.Config, cu.ToString(req.Database, ""), s.AppLog)
	token, err := ds.UserLogin(req.UserName, req.Password, true)
	login = &pb.ResponseAuthLogin{
		Token:   token,
		Version: cu.ToString(s.Config["version"], ""),
	}
	return login, err
}

// Update or create user account
// If id or existing code is set, the user is updated, otherwise a new user is created.
// If user is not admin, the user id and code are set to the current user id and code
func (s *GService) AuthUpdate(ctx context.Context, req *pb.Auth) (pbAuth *pb.Auth, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	user := ctx.Value(md.AuthUserCtxKey).(md.Auth)

	if req.UserName == "" {
		return pbAuth, errors.New("auth user_name and user_group are required")
	}

	var updateID int64 = req.Id
	if req.Id > 0 || req.Code != "" {
		var rows []cu.IM
		if rows, err = ds.GetDataByID("auth", req.Id, req.Code, false); err != nil {
			return pbAuth, err
		}
		if len(rows) > 0 {
			updateID = cu.ToInteger(rows[0]["id"], 0)
		}
	}

	if user.UserGroup != md.UserGroupAdmin && (updateID == 0 || updateID != user.Id) {
		return pbAuth, errors.New(http.StatusText(http.StatusMethodNotAllowed))
	}

	// prepare values for database update
	values := cu.IM{
		"user_name":  req.UserName,
		"user_group": req.UserGroup.String(),
	}
	if updateID == 0 && req.Code != "" {
		values["code"] = req.Code
	}

	ut.ConvertByteToIMValue(&req.AuthMeta, &pb.AuthMeta{}, values, "auth_meta")
	ut.ConvertByteToIMValue(&req.AuthMap, &pb.JsonString{}, values, "auth_map")

	update := md.Update{
		Values: values,
		Model:  "auth",
	}
	if updateID > 0 {
		update.IDKey = updateID
	}

	if updateID, err = ds.StoreDataUpdate(update); err == nil {
		pbAuth, err = s.AuthGet(ctx, &pb.RequestGet{Id: updateID, Code: ""})
	}

	return pbAuth, err
}

// Get user account by database id or code (unique key)
// If user is not admin, the user id and code are set to the current user id and code
func (s *GService) AuthGet(ctx context.Context, req *pb.RequestGet) (pbAuth *pb.Auth, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	user := ctx.Value(md.AuthUserCtxKey).(md.Auth)
	if user.UserGroup != md.UserGroupAdmin || (req.Id == 0 && req.Code == "") {
		req.Id = user.Id
		req.Code = user.Code
	}
	var users []cu.IM
	if users, err = ds.GetDataByID("auth_view", req.Id, req.Code, true); err == nil {
		if userJson, found := users[0]["auth_object"].(string); found {
			err = ds.ConvertFromByte([]byte(userJson), &pbAuth)
		} else {
			err = errors.New(http.StatusText(http.StatusUnprocessableEntity))
		}
	}
	return pbAuth, err
}

// Change token user password
func (s *GService) AuthPassword(ctx context.Context, req *pb.RequestPasswordChange) (pbStatus *pb.ResponseStatus, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	user := ctx.Value(md.AuthUserCtxKey).(md.Auth)
	pbStatus = &pb.ResponseStatus{Success: true, Message: http.StatusText(http.StatusOK)}
	err = ds.UserPassword(user.Code, req.Password, req.Confirm)
	if err != nil {
		pbStatus = &pb.ResponseStatus{Success: false, Message: err.Error()}
	}
	return pbStatus, err
}

// Reset a user password and result a new password
// Allow admin group user or the token user itself to reset password.
// If request id and code are not set, the token user id and code are used.
func (s *GService) AuthPasswordReset(ctx context.Context, req *pb.RequestGet) (pbStatus *pb.ResponseStatus, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	tokenUser := ctx.Value(md.AuthUserCtxKey).(md.Auth)
	if tokenUser.UserGroup != md.UserGroupAdmin && (req.Id != tokenUser.Id || (req.Id == 0 && req.Code != tokenUser.Code)) {
		req.Id = tokenUser.Id
		req.Code = tokenUser.Code
	}
	password := cu.RandString(16)
	pbStatus = &pb.ResponseStatus{Success: true, Message: password}
	var user *pb.Auth
	if user, err = s.AuthGet(ctx, &pb.RequestGet{Id: req.Id, Code: req.Code}); err == nil {
		if tokenUser.UserGroup.String() != md.UserGroupAdmin.String() && tokenUser.Code != user.Code {
			return pbStatus, errors.New(http.StatusText(http.StatusMethodNotAllowed))
		}
		err = ds.UserPassword(user.Code, password, password)
	}
	if err != nil {
		pbStatus = &pb.ResponseStatus{Success: false, Message: err.Error()}
	}
	return pbStatus, err
}
