package admin

import (
	"github.com/gin-gonic/gin"
	adminV1 "go-mall-api/api/admin/v1/entity"
	"go-mall-api/models/ums_admin"
	"go-mall-api/models/ums_role"
)

// Info 获取用户信息
func (c AdminController) Info(ctx *gin.Context) (*adminV1.UserReply, error) {
	user, _ := ctx.Get("current_user")
	adminUser := user.(ums_admin.UmsAdmin)
	rolesId := ums_admin.GetUserRoleId(adminUser.ID)
	menus := ums_role.GetRoleMenu(rolesId)
	roles := ums_admin.GetUserRole(adminUser.ID)

	return &adminV1.UserReply{
		Icon:     adminUser.Icon,
		Username: adminUser.Username,
		Menus:    menus,
		Roles:    roles,
	}, nil
}
