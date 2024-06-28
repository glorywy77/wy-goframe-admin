package api

import (
	"time"
	"wy-goframe-admin/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type LoginCodeReq struct {
	g.Meta `path:"/login/code" tags:"LoginService" method:"get" summary:"登录获取验证码"`
}
type LoginCodeRes struct {
	Code *model.LoginCodeOutput
}

type AuthLoginReq struct {
	g.Meta `path:"/login" method:"post"`
}

type AuthLoginRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type AuthRefreshTokenReq struct {
	g.Meta `path:"/refresh_token" method:"post"`
}

type AuthRefreshTokenRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type AuthLogoutReq struct {
	g.Meta `path:"/logout" method:"post"`
}

type AuthLogoutRes struct {
	Result string
}
