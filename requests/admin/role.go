package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

func RoleRegister(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"name":        []string{"required"},
		"description": []string{"required"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:名称必填",
		},
		"description": []string{
			"required:描述必填",
		},
	}
	errs := validate(data, rules, messages)

	return errs
}
