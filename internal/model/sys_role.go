package model

import "github.com/gogf/gf/v2/os/gtime"

type SysRoleSaveInput struct {
	Id          int `json:"id,omitempty"` //新增时候Id为空
	RoleName    string
	Description string
	HasApis     []int
	CreateAt    *gtime.Time
	UpdateAt    *gtime.Time
}

type SysRolePageInput struct {
	RoleName    string
	PageSize    int
	CurrentPage int
}

type SysRolePageOutput struct {
	Id          int         `json:"id"`
	RoleName    string      `dc:"角色" json:"roleName"`
	Description string      `json:"description"`
	CreateAt    *gtime.Time `json:"createAt"`
	UpdateAt    *gtime.Time `json:"updateAt"`
}

type SysRoleDeleteInput struct {
	RoleName string
}

type SysRoleHasApisInput struct {
	RoleName string
}

type SysRoleHasApisOutput struct {
	Apikeys []int `json:"apiKeys" dc:"查询角色拥有的api接口"`
}

type SysRoleAllApisInput struct {
}

type SysRoleAllApisOutput struct {
	TreeData []map[string]interface{} `json:"treeData" dc:"所有api权限通过树形数据返回"`
}

type SysRoleListInput struct {
}

type SysRoleListOutput struct {
	RoleName string `json:"roleName"`
}