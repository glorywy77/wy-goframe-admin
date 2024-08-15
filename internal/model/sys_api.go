package model

import "github.com/gogf/gf/v2/os/gtime"

type SysApiSaveInput struct {
	Id          int 
	Path        string
	Method      string
	ApiGroup    string
	Description string
	CreatedAt   *gtime.Time
	UpdatedAt   *gtime.Time
}

type SysApiPageInput struct {
	Id          int
	Path        string
	ApiGroup    string
	Method      string
	PageSize    int
	CurrentPage int
}

type SysApiPageOutput struct {
	Id          int         `json:"id"`
	Path        string      `dc:"api路径" json:"path"`
	Method      string      `dc:"api方法" json:"method"`
	ApiGroup    string      `dc:"api接口组" json:"apiGroup"`
	Description string      `dc:"api接口描述" json:"description"`
	CreateAt    *gtime.Time `json:"createAt"`
	UpdateAt    *gtime.Time `json:"updateAt"`
}

type SysApiDeleteInput struct{
	Path        string      
	Method      string      
}


type SysApiGroupsListInput struct {
}

type SysApiGroupsListOutput struct {
	ApiGroup string `json:"apiGroup"`
}