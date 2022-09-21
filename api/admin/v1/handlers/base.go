package handlers

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
)

type AdminController interface {
	Login(*gin.Context, *entity.LoginRequest) (*entity.LoginReply, error)
	Logout(*gin.Context) error
	Info(*gin.Context) (*entity.UserReply, error)

	AdminList(*gin.Context, *entity.ListRequest) (*entity.ListReply, error)
	AdminUpdateStatus(*gin.Context, *entity.AdminUpdateStatusRequest) error
	AdminRegister(*gin.Context, *entity.AdminRegisterRequest) error
	AdminRoleUpdate(*gin.Context, *entity.AdminRoleUpdateRequest) error
	AdminRoles(*gin.Context, *entity.AdminRolesRequest) (*entity.AdminRolesReply, error)
	AdminUpdate(*gin.Context, *entity.AdminUpdateRequest) error
	AdminDelete(*gin.Context, *entity.AdminDeleteRequest) error

	RoleListAll(*gin.Context) (*entity.RoleListAllReply, error)
	RoleList(*gin.Context, *entity.RoleListRequest) (*entity.RoleListReply, error)
	RoleUpdateStatus(*gin.Context, *entity.RoleUpdateStatusRequest) error
	RoleRegister(*gin.Context, *entity.RoleRegisterRequest) error
	RoleDelete(*gin.Context, *entity.RoleDeleteRequest) error
	RoleUpdate(*gin.Context, *entity.RoleUpdateRequest) error

	MenuTreeList(*gin.Context) ([]*entity.MenuTreeListReply, error)
	MenuList(*gin.Context, *entity.MenuListRequest) (*entity.MenuListReply, error)
	MenuUpdateHidden(*gin.Context, *entity.MenuUpdateHiddenRequest) error
	MenuCreate(*gin.Context, *entity.MenuCreateRequest) error
	Menu(*gin.Context, *entity.MenuRequest) (*entity.MenuReply, error)
	MenuUpdate(*gin.Context, *entity.MenuUpdateRequest) error
	MenuDelete(*gin.Context, *entity.MenuDeleteRequest) error
}
