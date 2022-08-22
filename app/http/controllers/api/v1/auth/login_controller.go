package auth

import (
	"github.com/gin-gonic/gin"
	requests2 "go-mall-api/app/http/controllers/api/requests"
	v1 "go-mall-api/app/http/controllers/api/v1"
	"go-mall-api/pkg/auth"
	"go-mall-api/pkg/jwt"
	"go-mall-api/pkg/response"
)

// LoginController 用户控制器
type LoginController struct {
	v1.BaseAPIController
}

// LoginByPhone 手机登录
func (lc *LoginController) LoginByPhone(c *gin.Context) {
	//1. 验证表单
	request := requests2.LoginByPhoneRequest{}
	if ok := requests2.Validate(c, &request, requests2.LoginByPhone); !ok {
		return
	}

	user, err := auth.LoginByPhone(request.Phone)
	if err != nil {
		// 失败，显示错误提示
		response.Error(c, err, "账号不存在或密码错误")
	} else {
		//登录成功
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)

		response.JSON(c, gin.H{
			"token": token,
		})
	}
}

// LoginByPassword 多种方法登录，支持手机号、email 和用户名
func (lc *LoginController) LoginByPassword(c *gin.Context) {
	// 1. 验证表单
	request := requests2.LoginByPasswordRequest{}
	if ok := requests2.Validate(c, &request, requests2.LoginByPassword); !ok {
		return
	}

	// 2. 尝试登录
	user, err := auth.Attempt(request.LoginID, request.Password)
	if err != nil {
		// 失败,显示错误提示
		response.Unauthorized(c, "账号不存在或密码错误")
	} else {
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)
		response.JSON(c, gin.H{
			"token": token,
		})
	}
}

// RefreshToken 刷新 Access Token
func (lc *LoginController) RefreshToken(c *gin.Context) {
	token, err := jwt.NewJWT().RefreshToken(c)
	if err != nil {
		response.Error(c, err, "令牌刷新失效")
	} else {
		response.JSON(c, gin.H{
			"token": token,
		})
	}
}
