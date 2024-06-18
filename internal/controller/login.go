package controller

import (
	"context"
	"wy-goframe-admin/api"
	"wy-goframe-admin/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

var Login = cLogin{}

type cLogin struct{}

func (c *cLogin) LoginCode(ctx context.Context, req *api.LoginCodeReq) (res *api.LoginCodeRes, err error) {
	data, err := service.Login().LoginCode(ctx)
	if err != nil {
		return nil, err
	}
	res = &api.LoginCodeRes{
		Code: data,
	}
	g.Dump(res)
	return res,nil
}
