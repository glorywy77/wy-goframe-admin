package SysApi

import (
	"context"
	"wy-goframe-admin/internal/dao"
	"wy-goframe-admin/internal/model"
	"wy-goframe-admin/internal/service"

	_ "github.com/go-sql-driver/mysql"
)

type (
	sSysApi struct{}
)

func New() *sSysApi {
	return &sSysApi{}
}

func init() {
	service.RegisterSysApi(New())
}

// 保存接口
func (s *sSysApi) ApiSave(ctx context.Context, in model.SysApiSaveInput) (err error) {
	_, err = dao.SysApi.Ctx(ctx).Data(in).Save()
	return
}

// 分页展示
func (s *sSysApi) ApiPage(ctx context.Context, in model.SysApiPageInput) (out []*model.SysApiPageOutput, total int, err error) {
	m := dao.SysApi.Ctx(ctx)
	err = m.Fields(`id,path,method,api_group,description,create_at,update_at`).
		WhereLike("path", "%"+in.Path+"%").
		WhereLike("method", "%"+in.Method+"%").
		WhereLike("api_group", "%"+in.ApiGroup+"%").
		OrderAsc("api_group,id,path").Limit((in.CurrentPage-1)*in.PageSize, in.PageSize).
		ScanAndCount(&out, &total, false)

	if err != nil {
		return nil, 0, err
	}
	return out, total, nil

}

// 删
func (s *sSysApi) ApiDelete(ctx context.Context, in model.SysApiDeleteInput) (err error) {
	_, err = dao.SysApi.Ctx(ctx).Where("id", in.Id).Delete()
	return
}
