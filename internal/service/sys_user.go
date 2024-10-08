// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"wy-goframe-admin/internal/model"
)

type (
	ISysUser interface {
		UserLogin(ctx context.Context, in model.SysUserLoginInput) (userMap map[string]interface{}, err error)
		UserCreate(ctx context.Context, in model.SysUserCreateInput) (err error)
		UserUpdate(ctx context.Context, in model.SysUserUpdateInput) (err error)
		UserResetPass(ctx context.Context, in model.SysUserResetPassInput) (err error)
		UserPage(ctx context.Context, in model.SysUserPageInput) (out []*model.SysUserPageOutput, total int, err error)
		UserDelete(ctx context.Context, in model.SysUserDeleteInput) (err error)
	}
)

var (
	localSysUser ISysUser
)

func SysUser() ISysUser {
	if localSysUser == nil {
		panic("implement not found for interface ISysUser, forgot register?")
	}
	return localSysUser
}

func RegisterSysUser(i ISysUser) {
	localSysUser = i
}
