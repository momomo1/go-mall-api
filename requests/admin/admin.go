package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

func Register(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"username": []string{"required"},
		"password": []string{"required"},
	}
	messages := govalidator.MapData{
		"username": []string{
			"required:账号必填",
		},
		"password": []string{
			"required:密码必填",
		},
	}
	errs := validate(data, rules, messages)

	return errs
}
