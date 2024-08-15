package api

import (
	"wy-goframe-admin/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type AuditLogPageReq struct {
	g.Meta    `path:"/api/auditLog/page" method:"get" summary:"分页获取审计日志" tags:"AuditLogService"`
	Id        int
	TraceId   string
	UserName  string
	Ip        string
	Path      string
	Method    string
	StartTime *gtime.Time
	EndTime   *gtime.Time
	CommonPaginationReq
}

type AuditLogPageRes struct {
	CommonPaginationReq
	CommonPaginationRes
	Items []*model.AuditLogPageOutput `json:"items"`
}
