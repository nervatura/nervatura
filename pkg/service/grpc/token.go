package grpc

import (
	"context"

	cu "github.com/nervatura/component/pkg/util"
	api "github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	pb "github.com/nervatura/nervatura/v6/proto"
)

// Login by access token
func (s *GService) TokenLogin(ctx context.Context, req *pb.RequestEmpty) (pbAuth *pb.Auth, err error) {
	user := ctx.Value(md.AuthUserCtxKey).(md.Auth)
	return s.AuthGet(ctx, &pb.RequestGet{Id: user.Id, Code: user.Code})
}

// Refresh access token
func (s *GService) TokenRefresh(ctx context.Context, req *pb.RequestEmpty) (pbResponse *pb.ResponseAuthLogin, err error) {
	ds := ctx.Value(md.DataStoreCtxKey).(*api.DataStore)
	user := ctx.Value(md.AuthUserCtxKey).(md.Auth)
	var token string
	token, err = ds.TokenRefresh(user)
	pbResponse = &pb.ResponseAuthLogin{Token: token, Version: cu.ToString(ds.Config["version"], "")}
	return pbResponse, err
}

// Decode JWT access token but doesn't validate the signature
func (s *GService) TokenDecode(ctx context.Context, req *pb.RequestTokenDecode) (pbResponse *pb.ResponseTokenDecode, err error) {
	token, err := s.DecodeAuthToken(req.Token)
	pbResponse = &pb.ResponseTokenDecode{
		Code:     cu.ToString(token["sub"], ""),
		UserName: cu.ToString(token["user_name"], ""),
		Database: cu.ToString(token["alias"], ""),
		Exp:      cu.ToString(token["exp"], ""),
		Iss:      cu.ToString(token["iss"], ""),
	}
	return pbResponse, err
}
