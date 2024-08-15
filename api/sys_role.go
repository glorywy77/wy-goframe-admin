package api

import (
	"wy-goframe-admin/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type SysRoleSaveReq struct {
	g.Meta      `path:"/api/sysRole/save" method:"post,put" summary:"保存角色" tags:"SysRoleService"`
	Id          int    `json:"id"`
	RoleName        string `v:"required#角色名必填" dc:"角色"`
	Description string `v:"required#接口描述必填" dc:"api接口描述"`
	HasApis     []int  `json:"hasApis"`
}

type SysRolePageReq struct {
	g.Meta `path:"/api/sysRole/page" method:"get" summary:"分页获取角色" tags:"RoleService"`
	RoleName   string
	CommonPaginationReq
}

type SysRolePageRes struct {
	CommonPaginationReq
	CommonPaginationRes
	Items []*model.SysRolePageOutput `json:"items"`
}

type SysRoleDeleteReq struct {
	g.Meta `path:"/api/sysRole/delete" method:"delete" summary:"删除角色" tags:"RoleService"`
	RoleName   string `v:"required" `
}

type SysRoleHasApisReq struct {
	g.Meta `path:"/api/sysRole/hasApis" method:"get" summary:"获取角色拥有的api接口" tags:"RoleService"`
	RoleName   string
}

type SysRoleHasApisRes struct {
	HasApis *model.SysRoleHasApisOutput `json:"hasApis"`
}

type SysRoleAllApisReq struct {
	g.Meta `path:"/api/sysRole/allApis" method:"get" summary:"获取所有api接口" tags:"RoleService"`
}

type SysRoleAllApisRes struct {
	AllApis *model.SysRoleAllApisOutput `json:"allApis"`
}

type SysRoleListReq struct{
	g.Meta `path:"/api/sysRole/list" method:"get" summary:"获取接口组列表" tags:"SysApiService"`
}

type SysRoleListRes struct {
	Items []*model. SysRoleListOutput `json:"items"`
}