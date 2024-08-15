package api

import (
	"wy-goframe-admin/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type SysApiSaveReq struct {
	g.Meta      `path:"/api/sysApi/save" method:"post,put" summary:"保存接口" tags:"SysApiService"`
	Id          int
	Path        string `v:"required" dc:"api路径"`
	Method      string `v:"required" dc:"api方法"`
	Api_group   string `v:"required" dc:"api接口组"`
	Description string `dc:"api接口描述"`
}

type SysApiPageReq struct {
	g.Meta    `path:"/api/sysApi/page" method:"get" summary:"分页获取权限规则" tags:"SysApiService"`
	Path      string
	Method    string
	Api_group string
	CommonPaginationReq
}

type SysApiPageRes struct {
	CommonPaginationReq
	CommonPaginationRes
	Items []*model.SysApiPageOutput `json:"items"`
}

type SysApiDeleteReq struct {
	g.Meta `path:"/api/sysApi/delete" method:"delete" summary:"删除权限规则" tags:"SysApiService"`
	Path   string `v:"required" `
	Method string `v:"required" `
}


type SysApiGroupsListReq struct{
	g.Meta `path:"/api/sysApi/groups" method:"get" summary:"获取接口组列表" tags:"SysApiService"`
}

type SysApiGroupsListRes struct {
	Items []*model.SysApiGroupsListOutput `json:"items"`
}