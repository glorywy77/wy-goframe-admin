package controller

import (
	"context"
	"wy-goframe-admin/api"
	"wy-goframe-admin/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type cLogin struct{}

var Login = cLogin{}

func (c *cLogin) Login(ctx context.Context, req *api.AuthLoginReq) (res *api.AuthLoginRes, err error) {
	res = &api.AuthLoginRes{}
	res.Token, res.Expire = service.Login().Auth().LoginHandler(ctx)
	return
}

func (c *cLogin) RefreshToken(ctx context.Context, req *api.AuthRefreshTokenReq) (res *api.AuthRefreshTokenRes, err error) {
	res = &api.AuthRefreshTokenRes{}
	res.Token, res.Expire = service.Login().Auth().RefreshHandler(ctx)
	return
}

func (c *cLogin) Logout(ctx context.Context, req *api.AuthLogoutReq) (res *api.AuthLogoutRes, err error) {
	service.Login().Auth().LogoutHandler(ctx)
	return
}

func (c *cLogin) LoginCode(ctx context.Context, req *api.LoginCodeReq) (res *api.LoginCodeRes, err error) {
	data, err := service.Login().LoginCode(ctx)
	if err != nil {
		return nil, err
	}
	res = &api.LoginCodeRes{
		Code: data,
	}
	g.Dump(res)
	return res, nil
}
