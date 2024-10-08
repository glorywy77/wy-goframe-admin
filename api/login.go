package api

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type LoginCodeReq struct {
	g.Meta `path:"/api/login/code" tags:"LoginService" method:"get" summary:"登录获取验证码"`
}
// type LoginCodeRes struct {
// 	CodeUrl string `json:"codeUrl"`
// }

type LoginCodeRes struct{
	CaptchaSrc []byte `json:"captchaSrc"`
}

type AuthLoginReq struct {
	g.Meta   `path:"/api/login" tags:"LoginService" method:"post"`
	UserName string `v:"required"`
	Password string `v:"required"`
	Code     string `v:"required"`
}

type AuthLoginRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
	
}

type AuthRefreshTokenReq struct {
	g.Meta `path:"/api/refresh_token" tags:"LoginService"  method:"post"`
}

type AuthRefreshTokenRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type AuthLogoutReq struct {
	g.Meta `path:"/api/logout" tags:"LoginService" method:"post"`
}

type AuthLogoutRes struct {
}
