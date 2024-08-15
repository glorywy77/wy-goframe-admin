package casbin

import (
	"errors"
	"wy-goframe-admin/internal/service"

	"github.com/casbin/casbin/v2"
	xd "github.com/casbin/xorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

type (
	sCasbin struct{}
)

func New() *sCasbin {
	return &sCasbin{}
}

func init() {
	service.RegisterCasbin(New())
}

// 这里规定一个casbin的错误响应
func ErrReponse(r *ghttp.Request, err error) {
	r.Response.WriteJson(g.Map{
		"code":    403,
		"message": gconv.String(err),
		"data":    nil,
	})
}

// RoleCheck用于查询和校验用户权限
func (s *sCasbin) RoleCheck(r *ghttp.Request) {
	// fmt.Printf("r: %v\n", *r)
	
	ctx := r.GetCtx()
	//方法一 在先在session中去加然后再session中去取
	//Username, err := r.Session.Get("Username")

	//方法二 由于之前拦截器已经附加了一些信息在r中,所以可以在直接取JWT_PAYLOAD
	payload := gconv.Map(r.Get("JWT_PAYLOAD"))
	// g.Dump(payload)
	//获取用户权限

	roles := payload["roles"]
	username := payload["username"]
	// fmt.Printf("Username: %v\n", Username)

	if len(gconv.Map(roles)) == 0 {
		err := errors.New("鉴权失败,获取用户权限为空")
		g.Log().Errorf(ctx, "%v", err)
		ErrReponse(r, err)
		return
	}

	//连接到数据库
	adapter, err := xd.NewAdapter("mysql", "root:root@tcp(192.168.162.129:13306)/wy-goframe-admin?charset=utf8mb4", true)
	if err != nil {
		g.Log().Errorf(ctx, "error: adapter: %v", err)
		ErrReponse(r, err)
		return
	}

	//读取rbac并配置
	e, err := casbin.NewEnforcer("manifest/config/rbac_models.conf", adapter)
	if err != nil {
		g.Log().Errorf(ctx, "error: enforcer: %v", err)
		ErrReponse(r, err)
		return
	}

	//对鉴权结果进行判断
	obj := gconv.String(r.URL.Path)
	act := gconv.String(r.Method)
	var hasRole bool
	hasRole = false
	for sub := range gconv.Map(roles) {
		ok, err := e.Enforce(sub, obj, act)
		if err != nil {
			g.Log().Errorf(ctx, "鉴权出错: %v", err)
			err = errors.New("鉴权出错")
			ErrReponse(r, err)
		}
		if ok {
			hasRole = true
			break
		}
	}

	if hasRole {
		g.Log().Infof(ctx, "鉴权成功,<用户:%v> <角色:%v> <接口:%v> <方法:%v>", username, roles, obj, act)
		r.Middleware.Next()
	} else {
		g.Log().Errorf(ctx, "鉴权失败,<用户:%v> <角色:%v> <接口:%v> <方法:%v>", username, roles, obj, act)
		err = errors.New("鉴权失败,用户无权限")
		ErrReponse(r, err)
		return
	}

}

// 保存和新增权限规则
// func (s *sCasbin) RuleSave(ctx context.Context, in model.CasbinRuleSaveInput) (err error) {
// 	_, err = dao.CasbinRule.Ctx(ctx).Data(in).Save()
// 	return
// }

// 分页展示所有权限规则
// func (s *sCasbin) RulePage(ctx context.Context, in model.CasbinRulePageInput) (out []*model.CasbinRulePageOutput, total int, err error) {
// 	m := dao.CasbinRule.Ctx(ctx)
// 	err = m.Fields(`id,p_type,v0,v1,v2,v3,v4,v5,description`).
// 		WhereLike("v0", "%"+in.V0+"%").
// 		WhereLike("v1", "%"+in.V1+"%").
// 		WhereLike("v2", "%"+in.V2+"%").
// 		OrderAsc("v0,id").Limit((in.CurrentPage-1)*in.PageSize, in.PageSize).
// 		ScanAndCount(&out, &total, false)

// 	if err != nil {
// 		return nil, 0, err
// 	}
// 	return out, total, nil

// }

// 删除权限规则
// func (s *sCasbin) RuleDelete(ctx context.Context, in model.CasbinRuleDeleteInput) (err error) {
// 	_, err = dao.CasbinRule.Ctx(ctx).Where("id", in.Id).Delete()
// 	return
// }
