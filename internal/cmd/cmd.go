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
					service.Mm().Session,
				)

				group.Bind(
					// hello.New(),
					// controller.Login,
					controller.Auth,
					// hello.New(),
				)

				// Special handler that needs authentication.
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(
						// service.Mm().Auth,
						service.Mm().Casbin,
					)
					// group.ALLMap(g.Map{
					// 	"/user/info": controller.User.Info,
					// })
					group.Bind(
            hello.New(),
						controller.User,
						controller.Login,
					)
				})
			})

			s.Run()
			return nil
		},
	}
)
