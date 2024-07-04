package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type UserLoginInput struct {
	UserName string
	Password string
}

type UserCreateInput struct {
	UserName string
	Password string
	Email    string
	Roles    g.Slice
	Enable   int
}

type UserCreateOutput struct {
	Result string
}

type UserPageInput struct {
	UserName string
	Size     int
	Page     int
}

type UserPageOutput struct {
	UserName  string
	Email     string
	Roles     g.Slice
	Enable    int
	Create_at *gtime.Time
	Update_at *gtime.Time
	Remark    string
}
