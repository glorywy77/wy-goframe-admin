// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure for table sys_user.
type SysUser struct {
	Id       int         `json:"id"       description:""`
	Userid   string      `json:"userid"   description:""`
	Username string      `json:"username" description:""`
	Password string      `json:"password" description:""`
	Email    string      `json:"email"    description:""`
	Roles    string      `json:"roles"    description:""`
	Enable   int         `json:"enable"   description:""`
	CreateAt *gtime.Time `json:"createAt" description:""`
	UpdateAt *gtime.Time `json:"updateAt" description:""`
	Remark   string      `json:"remark"   description:""`
}