// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AuditLog is the golang structure of table audit_log for DAO operations like Where/Data.
type AuditLog struct {
	g.Meta       `orm:"table:audit_log, do:true"`
	Id           interface{} //
	TraceId      interface{} //
	Username     interface{} //
	Ip           interface{} //
	Path         interface{} //
	Method       interface{} //
	Params       interface{} //
	HttpCode     interface{} //
	ResponseTime interface{} //
	CreateAt     *gtime.Time //
}