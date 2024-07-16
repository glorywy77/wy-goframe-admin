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

// 通过用户名核密码查询用户是否存在，这部分是jwt的源码原本是写死的，改写了下通过数据库查
func (s *sUser) UserCheck(ctx context.Context, in model.UserLoginInput) (userMap map[string]interface{}) {

	// if in.UserName == "admin" && in.Password == "admin" {
	// 	return g.Map{
	// 		"id":       1,
	// 		"username": "admin",
	// 	}
	// }
	// return nil

	//查询用户是否已经被屏蔽
	m := dao.User.Ctx(ctx)
	enable, err := m.Fields("enable").Where("username", in.UserName).Value()
	if err != nil {
		return nil
	}

	if gconv.Int(enable) == 1 {
		//通过ctx获取到r
		r := g.RequestFromCtx(ctx)
		r.Response.WriteJson(g.Map{
			"code":    401,
			"message": "用户已屏蔽，请联系管理员",
		})
	}

	//校验密码是否正确
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

	//查询用户信息
	type User struct {
		UserId   string     `json:"userid"`
		Username string  `json:"username"`
		Roles    g.Slice `json:"roles"`
	}
	user := User{}
	err = m.Fields("userid,username,roles").Where("username", in.UserName).Scan(&user)
	if err != nil {
		return nil
	}
	userMap = gconv.Map(user)
	g.Dump(userMap)

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

// 更新用户
func (s *sUser) UserUpdate(ctx context.Context, in model.UserUpdateInput) (err error) {
	_, err = dao.User.Ctx(ctx).Data(in).Where("userid", in.UserId).Save()
	return
}

//重置用户密码
func (s *sUser) UserResetPass(ctx context.Context, in model.UserResetPassInput) (err error) {
	_, err = dao.User.Ctx(ctx).Data(in).Update()
	return
}



// 分页返回用户信息
func (s *sUser) UserPage(ctx context.Context, in model.UserPageInput) (out []*model.UserPageOutput, total int, err error) {
	m := dao.User.Ctx(ctx)
	err = m.Fields(`id,userid,username,email,roles,enable,create_at,update_at,remark`).WhereLike("username", "%"+in.UserName).OrderAsc("id").Limit((in.CurrentPage-1)*in.PageSize, in.PageSize).ScanAndCount(&out, &total, false)

	if err != nil {
		return nil, 0, err
	}

	return out, total, nil

}
