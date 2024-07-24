package model

type CasbinRuleSaveInput struct {
	Id      int `json:"id,omitempty"` //新增用户时候Id为空
	P_type  string
	V0      string
	V1      string
	V2      string
	V3      string
	V4      string
	V5      string
	Summary string
}

type CasbinRulePageInput struct {
	Id          int
	V0          string
	V1          string
	V2          string
	PageSize    int
	CurrentPage int
}

type CasbinRulePageOutput struct {
	Id      int    `json:"id"`
	P_type  string `d:"p" v:"required" json:"p_type"`
	V0      string `v:"required" dc:"角色" json:"v0"`
	V1      string `v:"required" dc:"接口" json:"v1"`
	V2      string `v:"required" dc:"请求方法" json:"v2"`
	V3      string
	V4      string
	V5      string
	Summary string `dc:"权限描述" json:"summary"`
}

type CasbinRuleDeleteInput struct{ Id int }
