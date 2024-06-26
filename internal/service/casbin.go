// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	ICasbin interface {
		SelectRole(ctx context.Context, r *ghttp.Request)
	}
)

var (
	localCasbin ICasbin
)

func Casbin() ICasbin {
	if localCasbin == nil {
		panic("implement not found for interface ICasbin, forgot register?")
	}
	return localCasbin
}

func RegisterCasbin(i ICasbin) {
	localCasbin = i
}
