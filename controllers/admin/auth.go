package admin

import (
	"errors"
	"github.com/gin-gonic/gin"
	adminV1 "go-mall-api/api/admin/v1"
	"go-mall-api/pkg/cache"
	"go-mall-api/pkg/jwt"
	auth "go-mall-api/service/auth_admin"
	"time"
)

func (c AdminController) Login(ctx *gin.Context, request *adminV1.LoginRequest) (*adminV1.LoginReply, error) {
	// 2. 尝试登录
	user, err := auth.Attempt(request.UserName, request.Password)
	if err != nil {
		// 失败,显示错误提示
		return &adminV1.LoginReply{}, err
	}
	token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Username)
	return &adminV1.LoginReply{
		Token:     token,
		TokenHead: "Bearer ",
	}, nil
}

func (c AdminController) Logout(ctx *gin.Context) error {
	token, err := jwt.NewJWT().GetTokenFromHeader(ctx)
	if err != nil {
		return errors.New("token错误")
	}

	// 设置缓存 key
	cacheKey := "blacklist:token:" + token
	// 设置过期时间
	expireTime := 84 * time.Hour
	cache.Set(cacheKey, token, expireTime)
	return nil
}
