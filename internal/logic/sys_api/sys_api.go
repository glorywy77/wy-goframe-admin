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

