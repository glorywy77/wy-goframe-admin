package api

import (
	"wy-goframe-admin/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type CasbinRuleSaveReq struct {
	g.Meta  `path:"/api/casbin-rule/save" method:"post,put" summary:"保存权限规则" tags:"CasbinService"`
	Id      int
	P_type  string `d:"p" v:"required"`
	V0      string `v:"required" dc:"角色"`
	V1      string `v:"required" dc:"接口"`
	V2      string `v:"required" dc:"请求方法"`
	V3      string
	V4      string
	V5      string
	Summary string `dc:"权限描述"`
}

type CasbinRulePageReq struct {
	g.Meta `path:"/api/casbin-rule/page" method:"get" summary:"分页获取权限规则" tags:"CasbinService"`
	Id     int
	V0     string
	V1     string
	V2     string
	CommonPaginationReq
}

type CasbinRulePageRes struct {
	CommonPaginationReq
	CommonPaginationRes
	Items []*model.CasbinRulePageOutput `json:"items"`
}

type CasbinRuleDeleteReq struct {
	g.Meta `path:"/api/casbin-rule/delete" method:"delete" summary:"删除权限规则" tags:"CasbinService"`
	Id     int `v:"required" dc:"权限规则Id"`
}
