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

func ErrReponse(r *ghttp.Request, err error) {
	r.Response.WriteJson(g.Map{
		"code":    401,
		"message": gconv.String(err),
	})
}

func (s *sCasbin) SelectRole(r *ghttp.Request) {
	// fmt.Printf("r: %v\n", *r)
	ctx := r.GetCtx()
	//方法一 在先在session中去加然后再session中去取
	//Username, err := r.Session.Get("Username")

	//方法二 由于之前拦截器已经附加了一些信息在r中,所以可以在直接取JWT_PAYLOAD
	Payload := gconv.Map(r.Get("JWT_PAYLOAD"))
	// fmt.Printf("Payload: %v\n", Payload)
	//获取用户名
	Username := Payload["username"]
	// fmt.Printf("Username: %v\n", Username)

	if Username == nil || gconv.String(Username) == "" {
		err := errors.New("鉴权失败，获取用户名为空")
		g.Log().Errorf(ctx, "%v", err)
		ErrReponse(r, err)
		return
	}

	sub := gconv.String(Username)
	obj := gconv.String(r.URL.Path)
	act := gconv.String(r.Method)

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
	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		g.Log().Errorf(ctx, "鉴权出错: %v", err)
		ErrReponse(r, err)
		return
	}
	if ok {
		g.Log().Infof(ctx, "鉴权成功，用户有权限：%v", sub)
		r.Middleware.Next()
	} else {
		g.Log().Errorf(ctx, "鉴权失败，用户无权限: %v", sub)
		err = errors.New("鉴权失败，用户无权限")
		ErrReponse(r, err)
		return
	}

}
