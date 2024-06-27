package api

import "github.com/gogf/gf/v2/frame/g"

type UserGetInfoReq struct {
	g.Meta `path:"/user/info" method:"get"`
}

type UserGetInfoRes struct {
	Id          int    `json:"id"`
	IdentityKey string `json:"identity_key"`
	Payload     string `json:"payload"`
}


type UserCreateReq struct {
	g.Meta `path:"/user/create" method:"post"  tags:"UserService" summary:"创建用户"`
  UserName string `v:"required"`
	Password string `v:"required"`
	Email string `v:"required"`
	Role string `v:"required"`
	Status int `d:"0" v:"required"  dc:"用户状态默认正常"`
  
}




