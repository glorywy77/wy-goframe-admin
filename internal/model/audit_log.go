package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type AuditLogPageInput struct {
	TraceId     string
	UserName    string
	Ip          string
	Path        string
	Method      string
	StartTime   *gtime.Time
	EndTime     *gtime.Time
	PageSize    int
	CurrentPage int
}

type AuditLogPageOutput struct {
	Id           int         `json:"id"`
	TraceId      string      `json:"traceId"`
	UserName     string      `json:"userName"`
	Ip           string      `json:"ip"`
	Path         string      `json:"path"`
	Method       string      `json:"method"`
	Params       g.Map       `json:"params"`
	HttpCode     int         `json:"httpCode"`
	ResponseTime string      `json:"responseTime"`
	CreateAt     *gtime.Time `json:"createAt"`
}

type AuditLogSaveInput struct {
	TraceId      string
	UserName     string
	Ip           string
	Path         string
	Method       string
	Params       g.Map
	HttpCode     int
	ResponseTime int
	CreateAt     *gtime.Time
}
