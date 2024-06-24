package mm

import (
	"context"
	"wy-goframe-admin/internal/service"
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	sMm struct{}
)

func init() {
	service.RegisterMm(New())
}
func New() *sMm {
	return &sMm{}
}

func (s *sMm) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func (s *sMm) Auth(r *ghttp.Request) {
	service.Auth().MiddlewareFunc()(r)
	r.Middleware.Next()
}

func (s *sMm) Session(r *ghttp.Request) {
	service.Session().SetSession(r)
	r.Middleware.Next()
}

func (s *sMm) Casbin(r *ghttp.Request) {
	var ctx context.Context
	service.Casbin().SelectRole(ctx, r)

}
