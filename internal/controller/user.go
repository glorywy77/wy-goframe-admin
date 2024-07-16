package controller

import (
	"context"
	"math/rand"
	"time"
	"wy-goframe-admin/api"
	"wy-goframe-admin/internal/model"
	"wy-goframe-admin/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"golang.org/x/crypto/bcrypt"
)

type userController struct{}

var User = userController{}

// Info should be authenticated to view.
// It is the get user data handler
func (c *userController) Info(ctx context.Context, req *api.UserGetInfoReq) (res *api.UserGetInfoRes, err error) {
	return &api.UserGetInfoRes{
		UserId:      gconv.String(service.Login().Auth().GetIdentity(ctx)),
		IdentityKey: service.Login().Auth().IdentityKey,
		Payload:     service.Login().Auth().GetPayload(ctx),
	}, nil
}

// 生成随机的UserId
func GenerateUserID() string {
	// 使用当前时间的纳秒值作为种子创建一个新的随机源
	src := rand.NewSource(time.Now().UnixNano())
	// 创建一个新的随机数生成器
	r := rand.New(src)
	const letterRunes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const digitRunes = "0123456789"
	var letters [3]byte
	for i := range letters {
		letters[i] = letterRunes[r.Intn(len(letterRunes))]
	}
	var digits [5]byte
	for i := range digits {
		digits[i] = digitRunes[r.Intn(len(digitRunes))]
	}
	return string(digits[:]) + string(letters[:])
}

// 创建新用户
func (c *userController) Create(ctx context.Context, req *api.UserCreateReq) (res *api.UserCreateRes, err error) {
	g.DumpJson(req)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	Password := gconv.String(hashedPassword)
	err = service.User().UserCreate(ctx, model.UserCreateInput{
		UserId:   GenerateUserID(),
		UserName: req.UserName,
		Password: Password,
		Email:    req.Email,
		Roles:    req.Roles,
		Enable:   req.Enable,
		Remark:   req.Remark,
	})
	if err != nil {
		return nil, err
	}
	res = &api.UserCreateRes{
		Result: "用户创建成功",
	}
	return
}

// 修改用户基础信息（不包含密码）
func (c *userController) Update(ctx context.Context, req *api.UserUpdateReq) (res *api.UserUpdateRes, err error) {
	g.DumpJson(req)
	err = service.User().UserUpdate(ctx, model.UserUpdateInput{
		Id:       req.Id,
		UserId:   req.UserId,
		UserName: req.UserName,
		Email:    req.Email,
		Roles:    req.Roles,
		Enable:   req.Enable,
		Remark:   req.Remark,
	})
	if err != nil {
		return nil, err
	}
	res = &api.UserUpdateRes{
		Result: "用户更新成功",
	}
	return
}

// 重置用户密码
func (c *userController) ResetPass(ctx context.Context, req *api.UserResetPassReq) (res *api.UserResetPassRes, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	Password := gconv.String(hashedPassword)
	err = service.User().UserResetPass(ctx, model.UserResetPassInput{
		UserId:   req.UserId,
		UserName: req.UserName,
		Password: Password,
	})
	if err != nil {
		return nil, err
	}
	res = &api.UserResetPassRes{
		Result: "用户重置密码成功",
	}
	return res, nil
}

// 分页返回用户信息
func (c *userController) Page(ctx context.Context, req *api.UserPageReq) (res *api.UserPageRes, err error) {
	data, total, err := service.User().UserPage(ctx, model.UserPageInput{
		UserName:    req.UserName,
		CurrentPage: req.CurrentPage,
		PageSize:    req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	res = &api.UserPageRes{
		CommonPaginationReq: api.CommonPaginationReq{
			PageSize:    req.PageSize,
			CurrentPage: req.CurrentPage,
		},
		CommonPaginationRes: api.CommonPaginationRes{
			Total: total,
		},
		Items: data,
	}
	return
}
