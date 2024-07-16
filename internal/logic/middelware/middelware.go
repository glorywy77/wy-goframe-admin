package middelware

import (
	"wy-goframe-admin/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	sMiddelware struct{}
)

func init() {
	service.RegisterMiddelware(New())
}
func New() *sMiddelware {
	return &sMiddelware{}
}

func (s *sMiddelware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func (s *sMiddelware) Auth(r *ghttp.Request) {
	service.Login().Auth().MiddlewareFunc()(r)
	r.Middleware.Next()
}

func (s *sMiddelware) Session(r *ghttp.Request) {
	service.Session().SetSession(r)
	r.Middleware.Next()
}

func (s *sMiddelware) Casbin(r *ghttp.Request) {
	service.Casbin().SelectRole(r)

}
