package service

import (
	"context"

	"wy-goframe-admin/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type userService struct{}

var user = userService{}

func User() *userService {
	return &user
}

func (s *userService) GetUserByUserNamePassword(ctx context.Context, in model.UserLoginInput) map[string]interface{} {
	if in.UserName == "admin" && in.Password == "admin" {
		return g.Map{
			"id":       1,
			"username": "admin",
		}
	}
	return nil
}
