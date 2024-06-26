package session

import (
	"wy-goframe-admin/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
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
		gconv.String(Username)
		r.Session.MustSet("Username", Username)

	}
}
