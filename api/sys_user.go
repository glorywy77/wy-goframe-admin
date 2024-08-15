package api

import (
	"wy-goframe-admin/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// 查询当前登录用户的信息
type SysUserGetInfoReq struct {
	g.Meta `path:"/api/sysUser/info" tags:"SysUserService" method:"get"  summary:"当前登录用户信息"`
}

type SysUserGetInfoRes struct {
	UserId      string `json:"userid" summary:"用户id作为唯一标识"` //此字段绝对不能用于传参会覆盖同命参数!!!!!比如你admin用户登录后,用read的userid去请求,会被覆盖成admin的userid
	IdentityKey string `json:"identity_key"`
	Payload     g.Map  `json:"payload"`
}

// 增
type SysUserCreateReq struct {
	g.Meta          `path:"/api/sysUser/create" method:"post"  tags:"SysUserService" summary:"新增用户"`
	UserName        string  `v:"required" `
	Password        string  `v:"required|length:8,30#请输入密码|密码长度不够" `
	ConfirmPassword string  `v:"required|length:8,30|same:Password#请确认密码|密码长度不够|两次密码不一致"`
	Email           string  `v:"required"`
	Roles           g.Slice `d:"dev" v:"required"`
	Enable          int     `d:"0" v:"required"  dc:"用户准入默认为0允许"`
	Remark          string  `dc:"备注"`
}

type SysUserCreateRes struct {
	Result string `json:"result"`
}

// 改
type SysUserUpdateReq struct {
	g.Meta   `path:"/api/sysUser/update" method:"put"  tags:"SysUserService" summary:"修改用户"`
	Id       int     `v:"required" dc:"ID"`
	UserName string  `v:"required" `
	Email    string  `v:"required"`
	Roles    g.Slice `d:"dev" v:"required"`
	Enable   int     `d:"0" v:"required"  dc:"用户准入默认为0允许"`
	Remark   string  `dc:"备注"`
}

type SysUserUpdateRes struct {
	Result string `json:"result"`
}

type SysUserResetPassReq struct {
	g.Meta          `path:"/api/sysUser/resetPass" method:"put"  tags:"SysUserService" summary:"重置密码"`
	Id              int    `v:"required" `
	UserName        string `v:"required" `
	Password        string `v:"required|length:8,30#请输入密码|密码长度不够" `
	ConfirmPassword string `v:"required|length:8,30|same:Password#请确认密码|密码长度不够|两次密码不一致"`
}

type SysUserResetPassRes struct {
	Result string `json:"result"`
}

// 查
type SysUserPageReq struct {
	g.Meta   `path:"/api/sysUser/page" method:"get" tags:"SysUserService" summary:"用户列表"`
	UserName string `d:"%" v:"required"  dc:"用户名可以模糊查询"`
	Email    string `d:"%" v:"required"  dc:"邮箱可以模糊查询"`
	CommonPaginationReq
}

type SysUserPageRes struct {
	CommonPaginationReq
	CommonPaginationRes
	Items []*model.SysUserPageOutput `json:"items"`
}

// 删
type SysUserDeleteReq struct {
	g.Meta   `path:"/api/sysUser/delete" method:"delete"  tags:"SysUserService" summary:"删除用户"`
	Id       int    `v:"required" `
	UserName string `v:"required" `
}
