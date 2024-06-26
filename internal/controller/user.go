package controller

import (
	"context"
	"wy-goframe-admin/api"
	"wy-goframe-admin/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

type userController struct{}

var User = userController{}

// Info should be authenticated to view.
// It is the get user data handler
func (c *userController) Info(ctx context.Context, req *api.UserGetInfoReq) (res *api.UserGetInfoRes, err error) {
	return &api.UserGetInfoRes{
		Id:          gconv.Int(service.Login().Auth().GetIdentity(ctx)),
		IdentityKey: service.Login().Auth().IdentityKey,
		Payload:     service.Login().Auth().GetPayload(ctx),
	}, nil
}
