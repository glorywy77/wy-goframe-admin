package login

import (
	"context"
	"errors"
	"wy-goframe-admin/internal/model"
	"wy-goframe-admin/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type (
	sLogin struct{}
)

func New() *sLogin {
	return &sLogin{}
}

func init() {
	service.RegisterLogin(New())
}

// 登录验证码函数必须大写驼峰
func (s *sLogin) LoginCode(ctx context.Context) (Data *model.LoginCodeOutput, err error) {
	//作为测试先写死
	testCode := "test112233"
	Data = &model.LoginCodeOutput{Code: testCode}
	return Data, nil
}

// 用户登录
func (s *sLogin) UserLogin(ctx context.Context, in *model.UserLoginInput) (Data *model.UserLoginOutput, err error) {
	if in.Username == "admin" || in.Username == "editor" && in.Password == "123456" {
		g.Log().Infof(ctx, "用户名在白名单内: %v", in.Username)
		testToekn := "sdsafasdsadsadsadsadagd32313sjdhakhdahdka"
		Data = &model.UserLoginOutput{Token: testToekn}

	} else {
		g.Log().Errorf(ctx, "用户名或密码错误: %v", in.Username)
		Data = &model.UserLoginOutput{Token: ""}
		err = errors.New("用户名或密码错误")

	}
	return
}
