package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type SysUserLoginInput struct {
	UserName string
	Password string
	Code     string
}

// type SysUserSaveInput struct {
// 	Id       int
// 	SysUserName string
// 	Password string `json:"Password,omitempty"` //允许不传入password字段
// 	Email    string
// 	Roles    g.Slice
// 	Enable   int
// 	Remark   string
// }

// type SysUserSaveOutput struct {
// 	Result string
// }

type SysUserCreateInput struct {
	UserId   string
	UserName string
	Password string
	Email    string
	Roles    g.Slice
	Enable   int
	Remark   string
}

type SysUserUpdateInput struct {
	Id       int
	UserName string
	Email    string
	Roles    g.Slice
	Enable   int
	Remark   string
}

type SysUserResetPassInput struct {
	Id          int
	UserName string
	Password    string
}

type SysUserPageInput struct {
	UserName    string
	Email       string
	PageSize    int
	CurrentPage int
}

type SysUserPageOutput struct {
	Id        int         `json:"id"`
	UserId    string      `json:"userid"`
	UserName  string      `json:"username"`
	Email     string      `json:"email"`
	Roles     g.Slice     `json:"roles"`
	Enable    int         `json:"enable"`
	Create_at *gtime.Time `json:"create_at"`
	Update_at *gtime.Time `json:"update_at"`
	Remark    string      `json:"remark"`
}

type SysUserDeleteInput struct {
	Id       int
	UserName string
}
