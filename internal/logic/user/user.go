package user

import (
	"context"
	"errors"
	"wy-goframe-admin/internal/dao"
	"wy-goframe-admin/internal/model"
	"wy-goframe-admin/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"golang.org/x/crypto/bcrypt"
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

// UserCheck 用户登录校验
func (s *sUser) UserCheck(ctx context.Context, in model.UserLoginInput) (userMap map[string]interface{}) {

	//校验密码是否正确
	m := dao.User.Ctx(ctx)
	hashedPassword, err := m.Fields("password").Where("username", in.UserName).Value()
	if err != nil {
		return nil
	}

	if hashedPassword == nil {
		err = errors.New("用户名不存在")
		g.Log().Errorf(ctx, "%v", err)
		return nil
	}
	err = bcrypt.CompareHashAndPassword([]byte(gconv.String(hashedPassword)), []byte(in.Password))
	if err != nil {
		err = errors.New("密码不正确")
		g.Log().Errorf(ctx, "%v", err)
		return nil
	}

	//查询当前登录用户信息
	type User struct {
		UserId   string  `json:"userid"`
		Username string  `json:"username"`
		Roles    g.Slice `json:"roles"`
		Enable   int     `json:"enable"`
	}
	user := User{}
	err = m.Fields("userid,username,roles,enable").Where("username", in.UserName).Scan(&user)
	if err != nil {
		return nil
	}
	userMap = gconv.Map(user)
	// g.Dump(userMap)

	return userMap
}

// 新增用户
func (s *sUser) UserCreate(ctx context.Context, in model.UserCreateInput) (err error) {
	_, err = dao.User.Ctx(ctx).Data(in).Insert()
	//这里对用户名重复进行一个错误改写前端比较好看
	if gstr.Contains(gconv.String(err), "Duplicate entry") && gstr.Contains(gconv.String(err), "username") {
		err = errors.New("用户已存在")
		g.Log().Errorf(ctx, "%v", err)
	}
	return
}

// 更新用户基础信息,不修改密码
func (s *sUser) UserUpdate(ctx context.Context, in model.UserUpdateInput) (err error) {
	_, err = dao.User.Ctx(ctx).Data("email", in.Email, "roles", in.Roles, "enable", in.Enable, "remark", in.Remark).
		Where("id", in.Id).
		Where("username", in.UserName).
		Update()
	return
}

// 重置用户密码
func (s *sUser) UserResetPass(ctx context.Context, in model.UserResetPassInput) (err error) {
	_, err = dao.User.Ctx(ctx).Data("password", in.Password).
		Where("id", in.Id).
		Where("username", in.UserName).
		Update()
	return
}

// 分页返回用户信息
func (s *sUser) UserPage(ctx context.Context, in model.UserPageInput) (out []*model.UserPageOutput, total int, err error) {
	m := dao.User.Ctx(ctx)
	err = m.Fields(`id,userid,username,email,roles,enable,create_at,update_at,remark`).
		WhereLike("username", "%"+in.UserName+"%").
		WhereLike("email", "%"+in.Email+"%").
		OrderAsc("id").Limit((in.CurrentPage-1)*in.PageSize, in.PageSize).
		ScanAndCount(&out, &total, false)

	if err != nil {
		return nil, 0, err
	}
	return out, total, nil

}


// 删除用户
func (s *sUser) UserDelete(ctx context.Context, in model.UserDeleteInput) (err error) {
	_, err = dao.User.Ctx(ctx).
	Where("id", in.Id).
	Where("username", in.UserName).
	Delete()
	return
}