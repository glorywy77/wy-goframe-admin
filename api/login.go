package api

import (
	"wy-goframe-admin/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type LoginCodeReq struct {
	g.Meta `path:"/login/code" tags:"LoginService" method:"get" summary:"登录获取验证码"`
}
type LoginCodeRes struct {
	Code *model.LoginCodeOutput
}

type UserLoginReq struct {
	g.Meta   `path:"/users/login" tags:"LoginService" method:"post" summary:"登录获取验证码"`
	Username string `json:"username" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password" v:"required#密码不能为空" dc:"密码"`
	Code     string `json:"code" v:"required#验证码不能为空" dc:"验证码"`
}
type UserLoginRes struct {
	Token *model.UserLoginOutput
}
