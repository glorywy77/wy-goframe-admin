package controller

import (
	"context"
	"wy-goframe-admin/api"
	"wy-goframe-admin/internal/model"
	"wy-goframe-admin/internal/service"
)

type sysRoleController struct{}

var SysRole = sysRoleController{}

// 分页展示
func (c *sysRoleController) Page(ctx context.Context, req *api.SysRolePageReq) (res *api.SysRolePageRes, err error) {
	data, total, err := service.SysRole().RolePage(ctx, model.SysRolePageInput{
		Role:        req.Role,
		CurrentPage: req.CurrentPage,
		PageSize:    req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	res = &api.SysRolePageRes{
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
func (c *sysRoleController) Save(ctx context.Context, req *api.SysRoleSaveReq) (res *api.CommonResultRes, err error) {
	err = service.SysRole().RoleSave(ctx, model.SysRoleSaveInput{
		Id:          req.Id,
		Role:        req.Role,
		Description: req.Description,
	})
	res = &api.CommonResultRes{Result: "保存成功"}
	return
}

// 删除权限规则
func (c *sysRoleController) Delete(ctx context.Context, req *api.SysRoleDeleteReq) (res *api.CommonResultRes, err error) {
	err = service.SysRole().RoleDelete(ctx, model.SysRoleDeleteInput{
		Id: req.Id,
	})
	res = &api.CommonResultRes{Result: "删除成功"}
	return
}
