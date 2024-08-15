package session

import (
	"wy-goframe-admin/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	sSession struct{}
)

func New() *sSession {
	return &sSession{}
}

func init() {
	service.RegisterSession(New())
}


//暂时无用
func (s *sSession) SetSession(r *ghttp.Request) {
	if r.URL.Path == "/api/xxxxx" {
		loginCode := "4321"
		r.Session.MustSet("loginCode", loginCode)
		// g.Dump(r.Session.Id())
	}
}
