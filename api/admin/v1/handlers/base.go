package handlers

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
)

type AdminController interface {
	Login(*gin.Context, *entity.LoginRequest) (*entity.LoginReply, error)
	Logout(*gin.Context) error
	Info(*gin.Context) (*entity.UserReply, error)
	RoleListAll(*gin.Context) (*entity.RoleListAllReply, error)
	AdminList(*gin.Context, *entity.ListRequest) (*entity.ListReply, error)
	AdminUpdateStatus(*gin.Context, *entity.AdminUpdateStatusRequest) error
}