package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-mall-api/models/user"
	"go-mall-api/pkg/cache"
	"go-mall-api/pkg/config"
	"go-mall-api/pkg/jwt"
	"go-mall-api/pkg/response"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := jwt.NewJWT().GetTokenFromHeader(c)
		if err != nil {
			response.Unauthorized(c, fmt.Sprintf("请查看 %v 相关的接口认证文档", config.GetString("app.name")))
			return
		}

		cacheKey := "blacklist:token:" + token
		getToken := cache.Get(cacheKey)
		if getToken != "" {
			response.Unauthorized(c, "该token已加入黑名单")
			return
		}

		// 从标头 Authorization:Bearer xxxxx 中获取信息，并验证 JWT 的准确性
		claims, err := jwt.NewJWT().ParserToken(c)

		//JWT解析失败
		if err != nil {
			response.Unauthorized(c, fmt.Sprintf("请查看 %v 相关的接口认证文档", config.GetString("app.name")))
			return
		}

		//JWT 解析成功, 设置用户信息
		userModel := user.Get(claims.UserID)
		if userModel.ID == 0 {
			response.Unauthorized(c, "找不到对应用户，用户可能已删除")
			return
		}

		// 将用户信息存入 gin.context 里，后续 auth 包将从这里拿到当前用户数据
		c.Set("current_user_id", userModel.GetStringID())
		c.Set("current_user_name", userModel.Name)
		c.Set("current_user", userModel)

		c.Next()
	}
}
