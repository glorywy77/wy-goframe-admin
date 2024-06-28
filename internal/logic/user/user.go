package user

import (
	"context"

	"wy-goframe-admin/internal/dao"
	"wy-goframe-admin/internal/model"
	"wy-goframe-admin/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type (
	sUser struct{}
)

func New() *sUser {
	return &sUser{}
}

func init() {
	service.RegisterUser(New())
}

func (s *sUser) GetUserByUserNamePassword(ctx context.Context, in model.UserLoginInput) map[string]interface{} {
	if in.UserName == "admin" && in.Password == "admin" {
		return g.Map{
			"id":       1,
			"username": "admin",
		}
	}
	return nil
}

func (s *sUser) UserCreate(ctx context.Context, in model.UserCreateInput) (err error) {
	_, err = dao.User.Ctx(ctx).Data(in).Insert()

	return
}

// func(s *sUser) UserGet(ctx context.Context, in model.UserGetInput)(*model.UserGetOutput, error){
// 	// 查询用户
// 	user, err := dao.User.Ctx(ctx).WherePri(in.Id).One()
// 	if err != nil {
// 		return nil, err
// 	}
// 	// 查询用户角色
// }
