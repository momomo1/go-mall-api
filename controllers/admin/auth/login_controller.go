package auth

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/controllers/admin"
	requests "go-mall-api/entity/admin/request"
	"go-mall-api/pkg/cache"
	"go-mall-api/pkg/jwt"
	"go-mall-api/pkg/response"
	auth "go-mall-api/service/auth_admin"
	"time"
)

// LoginController 用户控制器
type LoginController struct {
	admin.BaseAdminController
}

// Login 登录
func (lc *LoginController) Login(c *gin.Context) {
	// 1. 验证表单
	request := requests.LoginRequest{}
	if ok := requests.Validate(c, &request, requests.Login); !ok {
		return
	}

	// 2. 尝试登录
	user, err := auth.Attempt(request.UserName, request.Password)
	if err != nil {
		// 失败,显示错误提示
		response.FailWithMessage(c, "账号不存在或密码错误")
	} else {
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Username)
		response.OkWithData(c, gin.H{
			"token":     token,
			"tokenHead": "Bearer ",
		})
	}
}

// Logout 退出登录
func (lc *LoginController) Logout(c *gin.Context) {
	token, err := jwt.NewJWT().GetTokenFromHeader(c)
	if err != nil {
		response.FailWithMessage(c, "token错误")
	}

	// 设置缓存 key
	cacheKey := "blacklist:token:" + token
	// 设置过期时间
	expireTime := 84 * time.Hour
	cache.Set(cacheKey, token, expireTime)

	response.Ok(c)
}
