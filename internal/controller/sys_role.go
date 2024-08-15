package controller

import (
	"context"
	"errors"
	"wy-goframe-admin/api"
	"wy-goframe-admin/internal/model"
	"wy-goframe-admin/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type sysRoleController struct{}

var SysRole = sysRoleController{}

// 分页展示
func (c *sysRoleController) Page(ctx context.Context, req *api.SysRolePageReq) (res *api.SysRolePageRes, err error) {
	data, total, err := service.SysRole().RolePage(ctx, model.SysRolePageInput{
		RoleName:    req.RoleName,
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
		RoleName:    req.RoleName,
		Description: req.Description,
		HasApis:     req.HasApis,
	})
	res = &api.CommonResultRes{Result: "保存成功"}
	return
}

// 删除权限规则
func (c *sysRoleController) Delete(ctx context.Context, req *api.SysRoleDeleteReq) (res *api.CommonResultRes, err error) {
	if req.RoleName == "admin" {
		res = &api.CommonResultRes{
			Result: "admin角色禁止删除",
		}
		err = errors.New("admin角色禁止删除")
		g.Log().Errorf(ctx, "admin角色禁止删除")
	} else {
		err = service.SysRole().RoleDelete(ctx, model.SysRoleDeleteInput{
			RoleName: req.RoleName,
		})
		res = &api.CommonResultRes{Result: "删除成功"}
	}
	return
}

func (c *sysRoleController) HasApis(ctx context.Context, req *api.SysRoleHasApisReq) (res *api.SysRoleHasApisRes, err error) {

	hasApis, err := service.SysRole().RoleHasApis(ctx, model.SysRoleHasApisInput{
		RoleName: req.RoleName,
	})
	res = &api.SysRoleHasApisRes{
		HasApis: hasApis,
	}

	return
}

func (c *sysRoleController) AllApis(ctx context.Context, req *api.SysRoleAllApisReq) (res *api.SysRoleAllApisRes, err error) {
	data, err := service.SysRole().GetAllApis(ctx, model.SysRoleAllApisInput{})
	res = &api.SysRoleAllApisRes{
		AllApis: data,
	}
	return
}

func (c *sysRoleController) RoleList(ctx context.Context, req *api.SysRoleListReq) (res *api.SysRoleListRes, err error) {
	out, err := service.SysRole().RoleList(ctx, model.SysRoleListInput{})
	if err != nil {
		return nil, err
	}
	res = &api.SysRoleListRes{Items: out}
	return
}
