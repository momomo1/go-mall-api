package admin

import (
	"github.com/gin-gonic/gin"
	adminV1 "go-mall-api/api/admin/v1"
	"go-mall-api/pkg/cache"
	"go-mall-api/pkg/jwt"
	"time"
)

func (c AdminController) Login(ctx *gin.Context, request *adminV1.LoginRequest) (*adminV1.LoginReply, error) {
	token, err := jwt.NewJWT().GetTokenFromHeader(ctx)
	if err != nil {
		return nil, err
	}

	// 设置缓存 key
	cacheKey := "blacklist:token:" + token
	// 设置过期时间
	expireTime := 84 * time.Hour
	cache.Set(cacheKey, token, expireTime)

	return &adminV1.LoginReply{
		Data: struct{ Token string }{
			Token: token,
		},
	}, nil
}
