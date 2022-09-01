package admin

import (
	"github.com/gin-gonic/gin"
	adminV1 "go-mall-api/api/admin/v1"
	"go-mall-api/models/ums_role"
)

func (c AdminController) RoleListAll(ctx *gin.Context) (*adminV1.RoleListAllReply, error) {
	return &adminV1.RoleListAllReply{
		Data: ums_role.All(),
	}, nil
}
