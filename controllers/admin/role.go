package admin

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/models/ums_role"
)

func (c AdminController) RoleListAll(ctx *gin.Context) (*entity.RoleListAllReply, error) {
	return &entity.RoleListAllReply{
		Data: ums_role.All(),
	}, nil
}
