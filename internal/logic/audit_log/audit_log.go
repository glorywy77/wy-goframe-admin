package auditLog

import (
	"context"
	"time"
	"wy-goframe-admin/internal/dao"
	"wy-goframe-admin/internal/model"
	"wy-goframe-admin/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type (
	sAuditLog struct{}
)

func New() *sAuditLog {
	return &sAuditLog{}
}

func init() {
	service.RegisterAuditLog(New())
}

// 操作日志记录到数据库
func (s *sAuditLog) Save(r *ghttp.Request) (err error) {
	ctx := r.GetCtx()
	now := time.Now()
	r.Middleware.Next()
	RespTime := gconv.Int(time.Since(now).Milliseconds())
	payload := gconv.Map(r.Get("JWT_PAYLOAD"))
	username := gconv.String(payload["username"])
	in := model.AuditLogSaveInput{
		TraceId:      gctx.CtxId(r.Context()),
		UserName:     username,
		Ip:           r.GetClientIp(),
		Path:         r.URL.Path,
		Method:       r.Method,
		Params:       gconv.Map(r.GetQueryMap()),
		ResponseTime: RespTime,
		HttpCode:     r.Response.Status,
	}

	// g.Dump(in)
	_, err = dao.AuditLog.Ctx(ctx).Data(in).Insert()
	return
}

func (s *sAuditLog) Page(ctx context.Context, in model.AuditLogPageInput) (out []*model.AuditLogPageOutput, total int, err error) {
	m := dao.AuditLog.Ctx(ctx)
	err = m.Fields("trace_id,username,ip,path,method,params,http_code,response_time,create_at").
		WhereLike("trace_id", "%"+in.TraceId+"%").
		WhereLike("username", "%"+in.UserName+"%").
		WhereLike("ip", "%"+in.Ip+"%").
		WhereLike("path", "%"+in.Path+"%").
		WhereLike("method", "%"+in.Method+"%").
		WhereBetween("create_at", in.StartTime, in.EndTime).
		OrderDesc("id").Limit((in.CurrentPage-1)*in.PageSize, in.PageSize).
		ScanAndCount(&out, &total, false)

	if err != nil {
		return nil, 0, err
	}

	return out, total, nil

}
