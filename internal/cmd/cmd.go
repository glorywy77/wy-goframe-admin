package cmd

import (
	"context"
	"wy-goframe-admin/internal/controller"
	"wy-goframe-admin/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gsession"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.SetSessionStorage(gsession.NewStorageMemory())
			s.Group("/", func(group *ghttp.RouterGroup) {

				group.Middleware(
					ghttp.MiddlewareHandlerResponse,
					ghttp.MiddlewareCORS,

					//这里是之前为了测试通过添加session来获取用户名，现在的用户名在jwt中取的所以暂时没有实际作用
					service.Middelware().Session,
				)

				group.Bind(
					controller.Login,
				)

				// Special handler that needs authentication.
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(
						service.Middelware().AuditLog,
						service.Middelware().Auth,
						service.Middelware().Casbin,
					)

					group.Bind(
						controller.SysUser,
						controller.SysApi,
						controller.SysRole,
						controller.AuditLog,
						controller.Ali,
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
