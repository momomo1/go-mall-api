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
	RoleListMenu(*gin.Context, *entity.RoleListMenuRequest) ([]*entity.RoleListMenuReply, error)
	RoleAllocMenu(*gin.Context, *entity.RoleAllocMenuRequest) error
	RoleListResource(*gin.Context, *entity.RoleListResourceRequest) ([]*entity.RoleListResourceReply, error)
	RoleAllocResource(*gin.Context, *entity.RoleAllocResourceRequest) error

	MenuTreeList(*gin.Context) ([]*entity.MenuTreeListReply, error)
	MenuList(*gin.Context, *entity.MenuListRequest) (*entity.MenuListReply, error)
	MenuUpdateHidden(*gin.Context, *entity.MenuUpdateHiddenRequest) error
	MenuCreate(*gin.Context, *entity.MenuCreateRequest) error
	Menu(*gin.Context, *entity.MenuRequest) (*entity.MenuReply, error)
	MenuUpdate(*gin.Context, *entity.MenuUpdateRequest) error
	MenuDelete(*gin.Context, *entity.MenuDeleteRequest) error

	ResourceListAll(*gin.Context) ([]*entity.ResourceList, error)
	ResourceList(*gin.Context, *entity.ResourceListRequest) (*entity.ResourceListReply, error)
	ResourceCreate(*gin.Context, *entity.ResourceCreateRequest) error
	ResourceUpdate(*gin.Context, *entity.ResourceUpdateRequest) error
	ResourceDelete(*gin.Context, *entity.ResourceDeleteRequest) error

	ResourceCategoryListAll(*gin.Context) ([]*entity.ResourceCategoryListAllReply, error)
	ResourceCategoryCreate(*gin.Context, *entity.ResourceCategoryCreateRequest) error
	ResourceCategoryUpdate(*gin.Context, *entity.ResourceCategoryUpdateRequest) error
	ResourceCategoryDelete(*gin.Context, *entity.ResourceCategoryDeleteRequest) error

	Brand(*gin.Context, *entity.BrandRequest) (*entity.BrandList, error)
	BrandList(*gin.Context, *entity.BrandListRequest) (*entity.BrandListReply, error)
	BrandCreate(*gin.Context, *entity.BrandCreateRequest) error
	BrandUpdate(*gin.Context, *entity.BrandUpdateRequest) error
	BrandDelete(*gin.Context, *entity.BrandDeleteRequest) error
	BrandUpdateFactoryStatus(*gin.Context, *entity.BrandUpdateFactoryStatusRequest) error
	BrandUpdateShowStatus(*gin.Context, *entity.BrandUpdateShowStatusRequest) error
}
