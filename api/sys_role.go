package api

import (
	"wy-goframe-admin/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type SysRoleSaveReq struct {
	g.Meta      `path:"/api/sysRole/save" method:"post,put" summary:"保存接口" tags:"SysRoleService"`
	Id          int    `json:"id"`
	Role        string `v:"required" dc:"角色" json:"role"`
	Description string `dc:"api接口描述"`
}

type SysRolePageReq struct {
	g.Meta `path:"/api/sysRole/page" method:"get" summary:"分页获取权限规则" tags:"RoleService"`
	Role   string
	CommonPaginationReq
}

type SysRolePageRes struct {
	CommonPaginationReq
	CommonPaginationRes
	Items []*model.SysRolePageOutput `json:"items"`
}

type SysRoleDeleteReq struct {
	g.Meta `path:"/api/sysRole/delete" method:"delete" summary:"删除权限规则" tags:"RoleService"`
	Id     int `v:"required" dc:"权限规则Id"`
}
