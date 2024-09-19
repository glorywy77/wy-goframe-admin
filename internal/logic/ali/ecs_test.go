package ecs

import (
	"log"
	"testing"
	"wy-goframe-admin/internal/model"
	"wy-goframe-admin/internal/service"

	"github.com/gogf/gf/v2/os/gctx"
)

func TestEcsPage(t *testing.T) {
	ctx := gctx.New()
	out, err := service.Ecs().EcsPage(ctx, model.AliEcsPageInput{
		RegionId:   "cn-zhangjiakou",
		CurrentPage: 1,
		PageSize:   10,
		// InstanceIds:       []string{"i-8vbdn64nyz3ew6yo6z6d"},
		// PublicIpAddresses: []string{"39.100.66.142"},
	})
	log.Println(out, err)
}
