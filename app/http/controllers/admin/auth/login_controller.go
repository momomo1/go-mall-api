package auth

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/app/http/controllers/admin"
	"go-mall-api/app/http/controllers/admin/requests"
	"go-mall-api/pkg/auth"
	"go-mall-api/pkg/jwt"
	"go-mall-api/pkg/response"
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
	user, err := auth.AttemptAdmin(request.Username, request.Password)
	if err != nil {
		// 失败,显示错误提示
		response.FailWithMessage("账号不存在或密码错误", c)
	} else {
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Username)
		response.OkWithData(gin.H{
			"token":     token,
			"tokenHead": "Bearer ",
		}, c)
	}
}
