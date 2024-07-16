package api

import (
	"wy-goframe-admin/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type UserGetInfoReq struct {
	g.Meta `path:"/api/user/info" method:"get"`
}

type UserGetInfoRes struct {
	UserId      string    `json:"userid"`
	IdentityKey string `json:"identity_key"`
	Payload     g.Map  `json:"payload"`
}

type UserCreateReq struct {
	g.Meta   `path:"/api/user/create" method:"post,put"  tags:"UserService" summary:"保存用户"`
	UserName string  `v:"required" `
	Password string  `v:"required|length:6,30#密码长度不够" `
	Email    string  `v:"required"`
	Roles    g.Slice `d:"dev" v:"required"`
	Enable   int     `d:"0" v:"required"  dc:"用户准入默认为0允许"`
	Remark   string  `dc:"备注"`
}

type UserCreateRes struct {
	Result string `json:"result"`
}

type UserUpdateReq struct {
	g.Meta   `path:"/api/user/update" method:"post,put"  tags:"UserService" summary:"保存用户"`
	Id       int     `v:"required" dc:"ID"`
	UserId   string  `v:"required" dc:"用户ID"`
	UserName string  `v:"required" `
	Email    string  `v:"required"`
	Roles    g.Slice `d:"dev" v:"required"`
	Enable   int     `d:"0" v:"required"  dc:"用户准入默认为0允许"`
	Remark   string  `dc:"备注"`
}

type UserUpdateRes struct {
	Result string `json:"result"`
}

type UserResetPassReq struct {
	g.Meta   `path:"/api/user/resetPass" method:"put"  tags:"UserService" summary:"重置密码"`
	UserId   string `v:"required" dc:"用户ID"`
	UserName string `v:"required" `
	Password string `v:"required length:6,30#密码长度不够" `
}

type UserResetPassRes struct {
	Result string `json:"result"`
}

type UserPageReq struct {
	g.Meta   `path:"/api/user/page" method:"post" tags:"UserService" summary:"用户列表"`
	UserName string `d:"%" v:"required"  dc:"用户名可以模糊查询"`
	CommonPaginationReq
}

type UserPageRes struct {
	CommonPaginationReq
	CommonPaginationRes
	Items []*model.UserPageOutput `json:"items"`
}
