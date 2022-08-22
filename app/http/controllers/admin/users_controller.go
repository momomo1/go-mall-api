package admin

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/app/models/ums_menu"
	"go-mall-api/pkg/response"
)

type UsersController struct {
	BaseAdminController
}

// Info 获取用户信息
func (ctrl *UsersController) Info(c *gin.Context) {
	response.OkWithData(c, gin.H{
		"icon":     "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180607/timg.jpg",
		"username": "admin",
		"roles":    ums_menu.All(),
		"menus":    ums_menu.All(),
	})
}
