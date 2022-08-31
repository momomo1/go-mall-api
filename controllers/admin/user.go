package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	adminV1 "go-mall-api/api/admin/v1"
	"go-mall-api/models/ums_menu"
)

// Info 获取用户信息
func (c AdminController) Info(ctx *gin.Context) (*adminV1.UserReply, error) {
	username, _ := ctx.Get("current_user_name")

	return &adminV1.UserReply{
		Icon:     "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180607/timg.jpg",
		Username: cast.ToString(username),
		Roles:    ums_menu.All(),
		Menus:    "admin",
	}, nil
}
