package SysRole

import (
	"context"
	"wy-goframe-admin/internal/dao"
	"wy-goframe-admin/internal/model"
	"wy-goframe-admin/internal/service"

	_ "github.com/go-sql-driver/mysql"
)

type (
	sSysRole struct{}
)

func New() *sSysRole {
	return &sSysRole{}
}

func init() {
	service.RegisterSysRole(New())
}

// 保存接口
func (s *sSysRole) RoleSave(ctx context.Context, in model.SysRoleSaveInput) (err error) {
	_, err = dao.SysRole.Ctx(ctx).Data(in).Save()
	return
}

// 分页展示所有权限规则
func (s *sSysRole) RolePage(ctx context.Context, in model.SysRolePageInput) (out []*model.SysRolePageOutput, total int, err error) {
	m := dao.SysRole.Ctx(ctx)
	err = m.Fields(`id,role,description,create_at,update_at`).
		WhereLike("role", "%"+in.Role+"%").
		OrderAsc("id").Limit((in.CurrentPage-1)*in.PageSize, in.PageSize).
		ScanAndCount(&out, &total, false)

	if err != nil {
		return nil, 0, err
	}
	return out, total, nil

}

// 删除角色
func (s *sSysRole) RoleDelete(ctx context.Context, in model.SysRoleDeleteInput) (err error) {
	_, err = dao.SysRole.Ctx(ctx).Where("id", in.Id).Delete()
	return
}
