package model

import "github.com/gogf/gf/v2/os/gtime"

type SysApiSaveInput struct {
	Id          int64
	Path        string
	Description string
	ApiGroup    string
	Method      string
	CreatedAt   *gtime.Time
	UpdatedAt   *gtime.Time
	DeletedAt   *gtime.Time
}
