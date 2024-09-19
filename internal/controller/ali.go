package controller

import (
	"context"
	"wy-goframe-admin/api"
	"wy-goframe-admin/internal/model"
	"wy-goframe-admin/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

type aliController struct{}

var Ali = aliController{}

// 分页展示所有ecs实例
func (c *aliController) EcsPage(ctx context.Context, req *api.AliEcsPageReq) (res *api.AliEcsPageRes, err error) {
	out, err := service.Ecs().EcsPage(ctx, model.AliEcsPageInput{
		RegionId: req.RegionId,
		// InstanceName:       req.InstanceName,
		// PrivateIpAddresses: req.PrivateIpAddresses,
		// PublicIpAddresses:  req.PublicIpAddresses,
		DynamicSelectKey:   req.DynamicSelectKey,
		DynamicSelectValue: req.DynamicSelectValue,
		CurrentPage:        req.CurrentPage,
		PageSize:           req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	body := gconv.Map(out["body"])

	instances := gconv.Map(body["Instances"])
	instance := gconv.SliceMap(instances["Instance"])
	// //保持代码风格一致故把键值全都改成小写

	// // 定义一个闭包函数，用于将字符串的首字母转换为小写
	// toLowerFirstChar := func(s string) string {
	// 	if s == "" {
	// 		return ""
	// 	}
	// 	r, n := utf8.DecodeRuneInString(s)
	// 	return string(unicode.ToLower(r)) + s[n:]
	// }

	// for i := range instance {
	// 	for k, v := range instance[i] {
	// 		// instance[i][strings.ToLower(k)] = v
	// 		instance[i][toLowerFirstChar(k)] = v
	// 		delete(instance[i], k)
	// 	}
	// }

	res = &api.AliEcsPageRes{
		Total:       gconv.Int(body["TotalCount"]),
		Instance:    instance,
		CurrentPage: req.CurrentPage,
		PageSize:    req.PageSize,
	}

	return
}
