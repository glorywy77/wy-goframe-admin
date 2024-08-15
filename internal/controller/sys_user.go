package controller

import (
	"context"
	"errors"
	"math/rand"
	"time"
	"wy-goframe-admin/api"
	"wy-goframe-admin/internal/model"
	"wy-goframe-admin/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"golang.org/x/crypto/bcrypt"
)

type sysUserController struct{}

var SysUser = sysUserController{}

// Info should be authenticated to view.
// It is the get user data handler
func (c *sysUserController) Info(ctx context.Context, req *api.SysUserGetInfoReq) (res *api.SysUserGetInfoRes, err error) {

	return &api.SysUserGetInfoRes{
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
	const allowedChars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var x [8]byte
	for i := range x {
		x[i] = allowedChars[r.Intn(len(allowedChars))]
	}

	return string(x[:])
}

// 创建新用户
func (c *sysUserController) Create(ctx context.Context, req *api.SysUserCreateReq) (res *api.SysUserCreateRes, err error) {
	g.DumpJson(req)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	Password := gconv.String(hashedPassword)
	err = service.SysUser().UserCreate(ctx, model.SysUserCreateInput{
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
	res = &api.SysUserCreateRes{
		Result: "用户创建成功",
	}
	return
}

// 修改用户基础信息（不包含密码）
func (c *sysUserController) Update(ctx context.Context, req *api.SysUserUpdateReq) (res *api.SysUserUpdateRes, err error) {
	g.DumpJson(req)
	err = service.SysUser().UserUpdate(ctx, model.SysUserUpdateInput{
		Id:       req.Id,
		UserName: req.UserName,
		Email:    req.Email,
		Roles:    req.Roles,
		Enable:   req.Enable,
		Remark:   req.Remark,
	})
	if err != nil {
		return nil, err
	}
	res = &api.SysUserUpdateRes{
		Result: "用户更新成功",
	}

	return
}

// 重置用户密码
func (c *sysUserController) ResetPass(ctx context.Context, req *api.SysUserResetPassReq) (res *api.SysUserResetPassRes, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	Password := gconv.String(hashedPassword)
	err = service.SysUser().UserResetPass(ctx, model.SysUserResetPassInput{
		Id:       req.Id,
		UserName: req.UserName,
		Password: Password,
	})
	if err != nil {
		return nil, err
	}
	res = &api.SysUserResetPassRes{
		Result: "用户重置密码成功",
	}
	return
}

// 分页返回用户信息
func (c *sysUserController) Page(ctx context.Context, req *api.SysUserPageReq) (res *api.SysUserPageRes, err error) {
	data, total, err := service.SysUser().UserPage(ctx, model.SysUserPageInput{
		UserName:    req.UserName,
		Email:       req.Email,
		CurrentPage: req.CurrentPage,
		PageSize:    req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	res = &api.SysUserPageRes{
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

// 用户删除
func (c *sysUserController) Delete(ctx context.Context, req *api.SysUserDeleteReq) (res *api.CommonResultRes, err error) {
	if req.UserName == "admin" {
		res = &api.CommonResultRes{
			Result: "admin用户禁止删除",
		}
		err = errors.New("admin用户禁止删除")
		g.Log().Errorf(ctx, "admin用户禁止删除")

	} else {
		err = service.SysUser().UserDelete(ctx, model.SysUserDeleteInput{
			Id:       req.Id,
			UserName: req.UserName,
		})
		if err != nil {
			return nil, err
		}
		res = &api.CommonResultRes{
			Result: "用户删除成功",
		}
	}
	return
}
