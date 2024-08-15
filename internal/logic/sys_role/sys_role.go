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

// 此函数用于比较两个切片的不同元素
func Difference(slice1, slice2 []int) (onlyInSlice1, onlyInSlice2 []int) {
	elementMap := make(map[int]bool)
	for _, item := range slice1 {
		elementMap[item] = true
	}

	for _, item := range slice2 {
		if _, exists := elementMap[item]; exists {
			delete(elementMap, item) // 如果在第二个切片中找到该元素，则从 map 中删除它
		} else {
			elementMap[item] = false // 如果只存在于第二个切片中，则添加到 map 中
		}
	}

	for key, exists := range elementMap {
		if exists { // 这些是 slice1 有而 slice2 没有的元素
			onlyInSlice1 = append(onlyInSlice1, key)
		} else { // 这些是 slice2 有而 slice1 没有的元素
			onlyInSlice2 = append(onlyInSlice2, key)
		}
	}

	return onlyInSlice1, onlyInSlice2
}

// 新增或者修改角色
func (s *sSysRole) RoleSave(ctx context.Context, in model.SysRoleSaveInput) (err error) {
	//第一步保存角色

	_, err = dao.SysRole.Ctx(ctx).Data(in).Save()
	if err != nil {
		return
	}
	//第二步分配角色权限
	//查询角色已拥有的权限
	apiKeys, err := service.SysRole().RoleHasApis(ctx, model.SysRoleHasApisInput{
		RoleName: in.RoleName,
	})
	old := apiKeys.Apikeys
	new := in.HasApis

	//对比新旧权限,删除和新增
	remove, add := Difference(old, new)
	//删
	if len(remove) > 0 {

		for _, v := range remove {
			var d *model.CasbinRuleDeleteInput
			err = dao.SysApi.Ctx(ctx).Fields("'"+in.RoleName+"' as v0,path as v1,method as v2").Where("id", v).Scan(&d)
			if err != nil {
				return
			}
			_, err = dao.CasbinRule.Ctx(ctx).Where("v0", d.V0).Where("v1", d.V1).Where("v2", d.V2).Delete()

		}
	}
	//增
	if len(add) > 0 {
		var a []*model.CasbinRuleSaveInput
		err = dao.SysApi.Ctx(ctx).Fields("'p' as p_type,'"+in.RoleName+"' as v0,path as v1,method as v2,description").Where("id IN(?)", add).Scan(&a)
		if err != nil {
			return
		}
		_, err = dao.CasbinRule.Ctx(ctx).Data(a).Save()

	}

	return
}

// 分页展示所有角色
func (s *sSysRole) RolePage(ctx context.Context, in model.SysRolePageInput) (out []*model.SysRolePageOutput, total int, err error) {
	m := dao.SysRole.Ctx(ctx)
	err = m.Fields(`id,rolename,description,create_at,update_at`).
		WhereLike("rolename", "%"+in.RoleName+"%").
		OrderAsc("id").Limit((in.CurrentPage-1)*in.PageSize, in.PageSize).
		ScanAndCount(&out, &total, false)

	if err != nil {
		return nil, 0, err
	}
	return out, total, nil

}

// 删除角色,一并删除对应权限
func (s *sSysRole) RoleDelete(ctx context.Context, in model.SysRoleDeleteInput) (err error) {
	//删除角色
	_, err = dao.SysRole.Ctx(ctx).Where("rolename", in.RoleName).Delete()
	if err != nil {
		return
	}

	//查询到当前此用户的权限，再删除
	apiKeys, err := service.SysRole().RoleHasApis(ctx, model.SysRoleHasApisInput{
		RoleName: in.RoleName,
	})
	if err != nil {
		return
	}
	remove := apiKeys.Apikeys
	if len(remove) > 0 {
		for _, v := range remove {
			var d *model.CasbinRuleDeleteInput
			err = dao.SysApi.Ctx(ctx).Fields("'"+in.RoleName+"' as v0,path as v1,method as v2").Where("id", v).Scan(&d)
			if err != nil {
				return
			}
			_, err = dao.CasbinRule.Ctx(ctx).Where("v0", d.V0).Where("v1", d.V1).Where("v2", d.V2).Delete()
			if err != nil {
				return
			}
		}
	}
	return
}

// 获取所有接口
func (s *sSysRole) GetAllApis(ctx context.Context, in model.SysRoleAllApisInput) (out *model.SysRoleAllApisOutput, err error) {

	//定义一个OgSqlData结构体,后面将以其为基础转换为我们想要的树形结构的数据
	type OgSqlData struct {
		ApiGroup  string
		ApiDetail map[string]string
	}
	var og []OgSqlData
	m := dao.SysApi.Ctx(ctx)
	err = m.Fields(`api_group as ApiGroup,JSON_OBJECTAGG(id,CONCAT(description,"  ",path,":",method)) as ApiDetail`).Group("api_group").Scan(&og)
	if err != nil {
		return nil, err
	}

	// 定义一个切片td用于存储转换后的数据
	var td []map[string]interface{}
	for i := 0; i < len(og); i++ {
		// 创建一个映射来存储当前迭代的元素数据
		item := make(map[string]interface{})
		// 将ApiGroup添加到映射中
		item["label"] = og[i].ApiGroup

		// 创建一个子项列表来存储ApiDetail
		children := make([]map[string]interface{}, 0)

		for k, v := range og[i].ApiDetail {
			// 对于每个ApiDetail创建一个映射
			child := make(map[string]interface{})
			// 将id和label添加到子项映射中
			child["id"] = k
			child["label"] = v
			// 将子项映射添加到子项列表中
			children = append(children, child)
		}
		// 将子项列表添加到item映射中
		item["children"] = children
		// 将item添加到最终的td中
		td = append(td, item)
	}
	out = &model.SysRoleAllApisOutput{TreeData: td}

	return
}

// 获取角色拥有的接口
func (s *sSysRole) RoleHasApis(ctx context.Context, in model.SysRoleHasApisInput) (out *model.SysRoleHasApisOutput, err error) {
	subQuery := dao.CasbinRule.Ctx(ctx).Fields(`v1,v2`).Where("v0", in.RoleName)
	err = dao.SysApi.Ctx(ctx).Fields(`JSON_ARRAYAGG(id) as ApiKeys`).WhereIn("(path,method)", subQuery).Scan(&out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// // 查询角色详情
// func (s *sSysRole) RoleDetail(ctx context.Context, in model.SysRoleDetailInput) (out []*model.SysRoleDetailOutput, err error) {
// 	err = dao.CasbinRule.Ctx(ctx).Fields("id,v1 as '路径',v2 as '方法',description as '描述'").
// 		WhereLike("v0", "%"+in.Role+"%").
// 		OrderAsc("v0,id").
// 		Scan(&out)

// 	if err != nil {
// 		return nil, err
// 	}
// 	return
// }

func (s *sSysRole) RoleList(ctx context.Context, in model.SysRoleListInput) (out []*model.SysRoleListOutput, err error) {
	m := dao.SysRole.Ctx(ctx)
	err = m.Fields("roleName").Group("roleName").Scan(&out)

	if err != nil {
		return nil, err
	}

	return out, nil
}
