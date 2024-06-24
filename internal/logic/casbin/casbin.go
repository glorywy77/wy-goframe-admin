package casbin

import (
	"context"
	"errors"
	"fmt"
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

func (s *sCasbin) SelectRole(ctx context.Context, r *ghttp.Request) (err error) {
	Username, err := r.Session.Get("Username")
	fmt.Printf("Username: %v\n", Username)
	if err != nil {
		g.Log().Errorf(ctx, "获取用户信息异常: %v", err)
		return
	}
	if Username == nil {
		g.Log().Errorf(ctx, "获取用户为空: %v", err)
		return
	}

	sub := gconv.String(Username)
	obj := gconv.String(r.URL.Path)
	act := gconv.String(r.Method)

	adapter, err := xd.NewAdapter("mysql", "root:root@tcp(192.168.162.129:13306)/wy-goframe-admin?charset=utf8mb4", true)
	if err != nil {
		g.Log().Errorf(ctx, "1error: adapter: %v", err)
	}

	e, err := casbin.NewEnforcer("manifest/config/rbac_models.conf", adapter)
	if err != nil {
		g.Log().Errorf(ctx, "2error: enforcer: %v", err)

	}

	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		g.Log().Errorf(ctx, "3获取用户信息异常: %v", err)
		return
	}

	if ok {
		g.Log().Infof(ctx, "4该用户有权限")
		r.Middleware.Next()
	} else {
		g.Log().Errorf(ctx, "5该用户无权限: %v", sub)
		err = errors.New("6该用户无权限")
	}
	return
}
