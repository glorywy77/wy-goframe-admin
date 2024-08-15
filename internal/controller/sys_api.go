package controller

import (
	"context"
	"wy-goframe-admin/api"
	"wy-goframe-admin/internal/model"
	"wy-goframe-admin/internal/service"
)

type sysApiController struct{}

var SysApi = sysApiController{}

// 分页展示
func (c *sysApiController) Page(ctx context.Context, req *api.SysApiPageReq) (res *api.SysApiPageRes, err error) {
	data, total, err := service.SysApi().ApiPage(ctx, model.SysApiPageInput{
		Path:        req.Path,
		Method:      req.Method,
		ApiGroup:    req.Api_group,
		CurrentPage: req.CurrentPage,
		PageSize:    req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	res = &api.SysApiPageRes{
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
func (c *sysApiController) Save(ctx context.Context, req *api.SysApiSaveReq) (res *api.CommonResultRes, err error) {
	err = service.SysApi().ApiSave(ctx, model.SysApiSaveInput{
		Id:          req.Id,
		Path:        req.Path,
		Method:      req.Method,
		ApiGroup:    req.Api_group,
		Description: req.Description,
	})
	res = &api.CommonResultRes{Result: "保存成功"}
	return
}

// 删除权限规则
func (c *sysApiController) Delete(ctx context.Context, req *api.SysApiDeleteReq) (res *api.CommonResultRes, err error) {
	err = service.SysApi().ApiDelete(ctx, model.SysApiDeleteInput{
		Path:   req.Path,
		Method: req.Method,
	})
	res = &api.CommonResultRes{Result: "删除成功"}
	return
}

func (c *sysApiController) ApiGroupsList(ctx context.Context, req *api.SysApiGroupsListReq) (res *api.SysApiGroupsListRes, err error) {
	out, err := service.SysApi().ApiGroupsList(ctx, model.SysApiGroupsListInput{})
	if err != nil {
		return nil, err
	}
	res = &api.SysApiGroupsListRes{Items: out}
	return
}