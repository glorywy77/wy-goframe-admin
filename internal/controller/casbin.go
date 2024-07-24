package controller

import (
	"context"
	"wy-goframe-admin/api"
	"wy-goframe-admin/internal/model"
	"wy-goframe-admin/internal/service"
)

type casbinController struct{}

var Casbin = casbinController{}

// 分页展示
func (c *casbinController) Page(ctx context.Context, req *api.CasbinRulePageReq) (res *api.CasbinRulePageRes, err error) {
	data, total, err := service.Casbin().RulePage(ctx, model.CasbinRulePageInput{
		Id:          req.Id,
		V0:          req.V0,
		V1:          req.V1,
		V2:          req.V2,
		CurrentPage: req.CurrentPage,
		PageSize:    req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	res = &api.CasbinRulePageRes{
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

// 新增或者保存
func (c *casbinController) Save(ctx context.Context, req *api.CasbinRuleSaveReq) (res *api.CommonResultRes, err error) {
	err = service.Casbin().RuleSave(ctx, model.CasbinRuleSaveInput{
		Id:      req.Id,
		V0:      req.V0,
		V1:      req.V1,
		V2:      req.V2,
		V3:      req.V3,
		V4:      req.V4,
		V5:      req.V5,
		P_type:  req.P_type,
		Summary: req.Summary,
	})
	res = &api.CommonResultRes{Result: "保存成功"}
	return
}

// 删除权限规则
func (c *casbinController) Delete(ctx context.Context, req *api.CasbinRuleDeleteReq) (res *api.CommonResultRes, err error) {
	err = service.Casbin().RuleDelete(ctx, model.CasbinRuleDeleteInput{
		Id: req.Id,
	})
	res = &api.CommonResultRes{Result: "删除成功"}
	return
}
