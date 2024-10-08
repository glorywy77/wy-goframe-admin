// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"wy-goframe-admin/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type (
	IEcs interface {
		EcsPage(ctx context.Context, in model.AliEcsPageInput) (out g.Map, err error)
	}
)

var (
	localEcs IEcs
)

func Ecs() IEcs {
	if localEcs == nil {
		panic("implement not found for interface IEcs, forgot register?")
	}
	return localEcs
}

func RegisterEcs(i IEcs) {
	localEcs = i
}
