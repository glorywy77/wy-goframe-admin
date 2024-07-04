package controller

import (
	"context"
	"wy-goframe-admin/api"
	"wy-goframe-admin/internal/model"
	"wy-goframe-admin/internal/service"

	"github.com/gogf/gf/v2/util/gconv"

	"golang.org/x/crypto/bcrypt"
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

func (c *userController) Create(ctx context.Context, req *api.UserCreateReq) (res *api.UserCreateRes, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	Password := gconv.String(hashedPassword)
	err = service.User().UserCreate(ctx, model.UserCreateInput{
		UserName: req.UserName,
		Password: Password,
		Email:    req.Email,
		Roles:     req.Roles,
		Enable:   req.Enable,
	})
	if err != nil {
		return nil, err
	}

	res = &api.UserCreateRes{
		Result: "用户创建成功",
	}
	return res, nil
}

func (c *userController) Page(ctx context.Context, req *api.UserPageReq) (res *api.UserPageRes, err error) {
	data, total, err := service.User().UserPage(ctx, model.UserPageInput{
		UserName: req.UserName,
		Page:     req.Page,
		Size:     req.Size,
	})
	if err != nil {
		return nil, err
	}
	res = &api.UserPageRes{
		CommonPaginationReq: api.CommonPaginationReq{
			Size: req.Size,
			Page: req.Page,
		},
		CommonPaginationRes: api.CommonPaginationRes{
			Total: total,
		},
		Items: data,
	}
	return
}
