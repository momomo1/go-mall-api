package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type LoginRequest struct {
	Username string `json:"username,omitempty" valid:"username"`
	Password string `json:"password,omitempty" valid:"password"`
}

// Login 验证表单，返回长度等于零即通过
func Login(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"username": []string{"required"},
		"password": []string{"required"},
	}
	messages := govalidator.MapData{
		"username": []string{
			"required:用户名为必填项，参数名称 username",
		},
		"password": []string{
			"required:密码必填",
		},
	}
	errs := validate(data, rules, messages)

	return errs
}
