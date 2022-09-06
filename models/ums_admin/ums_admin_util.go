package ums_admin

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/pkg/database"
	"go-mall-api/pkg/paginator"
	"gorm.io/gorm"
	"time"
)

func Get(idstr string) (umsAdmin UmsAdmin) {
	database.DB.Where("id", idstr).First(&umsAdmin)
	return
}

func GetBy(field, value string) (umsAdmin UmsAdmin) {
	database.DB.Where("? = ?", field, value).First(&umsAdmin)
	return
}

func All() (umsAdmins []UmsAdmin) {
	database.DB.Find(&umsAdmins)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(UmsAdmin{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// GetByMulti 通过 用户名 来获取用户
func GetByMulti(username string) (userAdminModel UmsAdmin) {
	database.DB.Where("username = ? and status = 1", username).First(&userAdminModel)
	return
}

// UpdateLoginTime 更新登录时间
func UpdateLoginTime(userAdminModel UmsAdmin) {
	now := time.Now()
	userAdminModel.LoginTime = now
	userAdminModel.Save()
}

// GetUserRoleId 获取全部角色id
func GetUserRoleId(id uint64) (roles []int64) {
	userAdminModel := UmsAdmin{}
	database.DB.Preload("RolesRelation").Where("id = ?", id).First(&userAdminModel)

	for _, v := range userAdminModel.RolesRelation {
		roles = append(roles, v.RoleId)
	}
	return
}

// GetUserRole 获取全部角色
func GetUserRole(id uint64) (roles []string) {
	userAdminModel := UmsAdmin{}
	database.DB.Preload("Roles", func(db *gorm.DB) *gorm.DB {
		return db.Select("id,name")
	}).Where("id = ?", id).First(&userAdminModel)

	for _, v := range userAdminModel.Roles {
		roles = append(roles, v.Name)
	}
	return
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int) (users []UmsAdmin, paging paginator.PagingAdmin) {
	paging = paginator.PaginateAdmin(
		c,
		database.DB.Model(UmsAdmin{}),
		&users,
		perPage,
	)
	return
}
