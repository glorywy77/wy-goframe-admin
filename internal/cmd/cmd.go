package cmd

import (
	"context"
	"wy-goframe-admin/internal/controller"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	//"wy-goframe-admin/internal/controller/hello"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Use(ghttp.MiddlewareHandlerResponse)
			// s.Group("/", func(group *ghttp.RouterGroup) {
			// 	// group.Middleware(ghttp.MiddlewareHandlerResponse)

			// 	group.Bind(
			// 		// hello.New(),
			// 		controller.Login,
			// 	)
			s.Group("/", func(group *ghttp.RouterGroup) {
				// Group middlewares.
				group.Middleware(

					ghttp.MiddlewareCORS,
				)
				group.Bind(

					controller.Login,
				)

			})
			s.Run()
			return nil
		},
	}
)
