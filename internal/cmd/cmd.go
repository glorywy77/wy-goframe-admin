package cmd

import (
	"context"
	"wy-goframe-admin/internal/controller"
	"wy-goframe-admin/internal/controller/hello"
	"wy-goframe-admin/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
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
						service.Middelware().Auth,
						service.Middelware().Casbin,
					)
					// group.ALLMap(g.Map{
					// 	"/user/info": controller.User.Info,
					// })
					group.Bind(
						hello.New(),
						controller.User,
					)
				})
			})

			s.Run()
			return nil
		},
	}
)
