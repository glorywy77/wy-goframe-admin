package api

import (
	"wy-goframe-admin/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)


//查询当前登录用户的信息
type UserGetInfoReq struct {
	g.Meta `path:"/api/user/info" tags:"UserService" method:"get"  summary:"当前登录用户信息"`
}

type UserGetInfoRes struct {
	UserId      string `json:"userid" summary:"用户id作为唯一标识"` //此字段绝对不能用于传参会覆盖同命参数!!!!!比如你admin用户登录后,用read的userid去请求,会被覆盖成admin的userid
	IdentityKey string `json:"identity_key"`
	Payload     g.Map  `json:"payload"`
}


//增
type UserCreateReq struct {
	g.Meta          `path:"/api/user/create" method:"post"  tags:"UserService" summary:"新增用户"`
	UserName        string  `v:"required" `
	Password        string  `v:"required|length:8,30#请输入密码|密码长度不够" `
	ConfirmPassword string  `v:"required|length:8,30|same:Password#请确认密码|密码长度不够|两次密码不一致"`
	Email           string  `v:"required"`
	Roles           g.Slice `d:"dev" v:"required"`
	Enable          int     `d:"0" v:"required"  dc:"用户准入默认为0允许"`
	Remark          string  `dc:"备注"`
}

type UserCreateRes struct {
	Result string `json:"result"`
}



//改
type UserUpdateReq struct {
	g.Meta   `path:"/api/user/update" method:"put"  tags:"UserService" summary:"修改用户"`
	Id       int     `v:"required" dc:"ID"`
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
	g.Meta          `path:"/api/user/resetPass" method:"put"  tags:"UserService" summary:"重置密码"`
	Id              int    `v:"required" `
	UserName        string `v:"required" `
	Password        string `v:"required|length:8,30#请输入密码|密码长度不够" `
	ConfirmPassword string `v:"required|length:8,30|same:Password#请确认密码|密码长度不够|两次密码不一致"`
}

type UserResetPassRes struct {
	Result string `json:"result"`
}

//查
type UserPageReq struct {
	g.Meta   `path:"/api/user/page" method:"get" tags:"UserService" summary:"用户列表"`
	UserName string `d:"%" v:"required"  dc:"用户名可以模糊查询"`
	Email    string `d:"%" v:"required"  dc:"邮箱可以模糊查询"`
	CommonPaginationReq
}

type UserPageRes struct {
	CommonPaginationReq
	CommonPaginationRes
	Items []*model.UserPageOutput `json:"items"`
}


//删
type UserDeleteReq struct {
	g.Meta          `path:"/api/user/delete" method:"delete"  tags:"UserService" summary:"删除用户"`
	Id              int    `v:"required" `
	UserName        string `v:"required" `
}
type UserDeleteRes struct {
	Result string `json:"result"`
}

	