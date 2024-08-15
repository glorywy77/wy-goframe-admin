package controller

import (
	"context"
	"wy-goframe-admin/api"
	"wy-goframe-admin/internal/model"
	"wy-goframe-admin/internal/service"
)

type auditLogController struct{}

var AuditLog = auditLogController{}

// 分页展示
func (c *auditLogController) Page(ctx context.Context, req *api.AuditLogPageReq) (res *api.AuditLogPageRes, err error) {
	data, total, err := service.AuditLog().Page(ctx, model.AuditLogPageInput{
		TraceId:     req.TraceId,
		UserName:    req.UserName,
		Ip:          req.Ip,
		Path:        req.Path,
		Method:      req.Method,
		CurrentPage: req.CurrentPage,
		PageSize:    req.PageSize,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
	})
	if err != nil {
		return nil, err
	}
	res = &api.AuditLogPageRes{
		CommonPaginationReq: api.CommonPaginationReq{
			PageSize:    req.PageSize,
			CurrentPage: req.CurrentPage,
		},
		CommonPaginationRes: api.CommonPaginationRes{
			Total: total,
		},
		Items: data,
	}

	return
}

// 新增或者保存
func (c *auditLogController) Save(ctx context.Context, req *api.SysApiSaveReq) (res *api.CommonResultRes, err error) {
	err = service.SysApi().ApiSave(ctx, model.SysApiSaveInput{
		Id:          req.Id,
		Path:        req.Path,
		Method:      req.Method,
		ApiGroup:    req.Api_group,
		Description: req.Description,
	})
	res = &api.CommonResultRes{Result: "保存成功"}
	return
}
