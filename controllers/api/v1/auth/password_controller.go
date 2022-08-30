package auth

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/controllers/api/v1"
	"go-mall-api/models/user"
	"go-mall-api/pkg/response"
	"go-mall-api/requests/api"
)

// PasswordController 用户控制器
type PasswordController struct {
	v1.BaseAPIController
}

// ResetByPhone 使用手机和验证码重置密码
func (pc *PasswordController) ResetByPhone(c *gin.Context) {
	//验证表单
	request := api.ResetByPhoneRequest{}
	if ok := api.Validate(c, &request, api.ResetByPhone); !ok {
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
	request := api.ResetByEmailRequest{}
	if ok := api.Validate(c, &request, api.ResetByEmail); !ok {
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
