package login

import (
	"context"
	"math/rand"
	"time"
	"wy-goframe-admin/internal/model"
	"wy-goframe-admin/internal/service"

	jwt "wy-goframe-admin/internal/logic/login/jwt"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	//"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
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

var authService *jwt.GfJWTMiddleware

func (s *sLogin) Auth() *jwt.GfJWTMiddleware {

	//写法一
	// var redisConfig = &gredis.Config{
	// 	Address: "192.168.162.129:6379",
	// 	Pass:    "canpanscp",
	// 	Db:      0,
	// }

	// redis, err := gredis.New(redisConfig)
	// if err != nil {
	// 	panic(err)
	// }
	// redisAdapter := gcache.NewAdapterRedis(redis)

	//写法二,配置文件
	redisAdapter := gcache.NewAdapterRedis(g.Redis())

	auth := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte("secret key11"),
		Timeout:         time.Hour * 24,
		MaxRefresh:      time.Hour * 24,
		IdentityKey:     "userid",
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		Authenticator:   Authenticator,
		Unauthorized:    Unauthorized,
		PayloadFunc:     PayloadFunc,
		IdentityHandler: IdentityHandler,
		CacheAdapter:    redisAdapter,
	})
	authService = auth
	return authService
}

// PayloadFunc is a callback function that will be called during login.
// Using this function it is possible to add additional payload data to the webtoken.
// The data is then made available during requests via c.Get("JWT_PAYLOAD").
// Note that the payload is not encrypted.
// The attributes mentioned on jwt.io can't be used as keys for the map.
// Optional, by default no additional data will be set.
func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

// IdentityHandler get the identity from JWT and set the identity for every request
// Using this function, by r.GetParam("id") get identity
func IdentityHandler(ctx context.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	return claims[authService.IdentityKey]
}

// Unauthorized is used to define customized Unauthorized callback function.
func Unauthorized(ctx context.Context, code int, message string) {
	r := g.RequestFromCtx(ctx)
	r.Response.WriteJson(g.Map{
		"code":    code,
		"message": message,
	})
	r.ExitAll()
}

// Authenticator is used to validate login parameters.
// It must return user data as user identifier, it will be stored in Claim Array.
// if your identityKey is 'id', your user data must have 'id'
// Check error (e) to determine the appropriate error message.
func Authenticator(ctx context.Context) (interface{}, error) {
	var (
		r  = g.RequestFromCtx(ctx)
		in model.SysUserLoginInput
	)
	if err := r.Parse(&in); err != nil {
		return "", err
	}

	user, err := service.SysUser().UserLogin(ctx, in)

	if err == nil {
		if user["enable"] == 0 {
			return user, nil
		} else {
			return nil, jwt.ErrFailedEnable
		}
	} else {
		return nil, err
	}

}

// 登录验证码
func (s *sLogin) LoginCode(ctx context.Context) (out string, err error) {

	src := rand.NewSource(time.Now().UnixNano())
	// 创建一个新的随机数生成器
	const allowedChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var x [4]byte
	for i := range x {
		x[i] = allowedChars[rand.New(src).Intn(len(allowedChars))]
	}
	code := string(x[:])

	out = "https://dummyimage.com/100x40/dcdfe6/000000.png&text=" + code
	r := g.RequestFromCtx(ctx)
	r.Session.MustSet("code", code)

	return out, nil
}
