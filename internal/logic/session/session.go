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

func (s *sSession) SetSession(r *ghttp.Request) {
	if r.URL.Path == "/login" {
		Username := r.Get("Username")
		r.Session.MustSet("Username", Username)

	}
}
