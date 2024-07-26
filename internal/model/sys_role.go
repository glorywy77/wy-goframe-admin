package model

import "github.com/gogf/gf/v2/os/gtime"

type SysRoleSaveInput struct {
	Id          int `json:"id,omitempty"` //新增时候Id为空
	Role        string
	Description string
	CreateAt    *gtime.Time
	UpdateAt    *gtime.Time
}

type SysRolePageInput struct {
	Role        string
	PageSize    int
	CurrentPage int
}

type SysRolePageOutput struct {
	Id          int         `json:"id"`
	Role        string      `dc:"角色" json:"role"`
	Description string      `json:"description"`
	CreateAt    *gtime.Time `json:"createAt"`
	UpdateAt    *gtime.Time `json:"updateAt"`
}

type SysRoleDeleteInput struct{ Id int }
