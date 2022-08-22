package auth

import (
	"github.com/gin-gonic/gin"
	requests2 "go-mall-api/app/http/controllers/api/requests"
	v1 "go-mall-api/app/http/controllers/api/v1"
	"go-mall-api/app/models/user"
	"go-mall-api/pkg/response"
)

// PasswordController 用户控制器
type PasswordController struct {
	v1.BaseAPIController
}

// ResetByPhone 使用手机和验证码重置密码
func (pc *PasswordController) ResetByPhone(c *gin.Context) {
	//验证表单
	request := requests2.ResetByPhoneRequest{}
	if ok := requests2.Validate(c, &request, requests2.ResetByPhone); !ok {
		return
	}
	//更新密码
	userModel := user.GetByPhone(request.Phone)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()

		response.Success(c)
	}
}

// ResetByEmail 使用 Email 和验证码重置密码
func (pc *PasswordController) ResetByEmail(c *gin.Context) {
	//验证表单
	request := requests2.ResetByEmailRequest{}
	if ok := requests2.Validate(c, &request, requests2.ResetByEmail); !ok {
		return
	}

	// 2. 更新密码
	userModel := user.GetByEmail(request.Email)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()
		response.Success(c)
	}
}
