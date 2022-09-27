package ums_role

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/pkg/database"
	"go-mall-api/pkg/paginator"
)

func Get(idstr string) (umsRole UmsRole) {
	database.DB.Where("id", idstr).First(&umsRole)
	return
}

func GetBy(field, value string) (umsRole UmsRole) {
	database.DB.Where("? = ?", field, value).First(&umsRole)
	return
}

func GetByIdInFind(value []string) (umsRole []UmsRole) {
	database.DB.Where("id in (?)", value).Find(&umsRole)
	return
}

func All() (umsRoles []UmsRole) {
	database.DB.Find(&umsRoles)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(UmsRole{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// GetRoleMenu 获取角色菜单
func GetRoleMenu(id []int64) (menus []interface{}) {
	userRoleModel := []UmsRole{}
	idArray := []uint64{}
	database.DB.Preload("Menus").Where("id in (?)", id).Find(&userRoleModel)

	IsContain := func(items []uint64, item uint64) bool {
		for _, eachItem := range items {
			if eachItem == item {
				return true
			}
		}
		return false
	}

	for _, value := range userRoleModel {
		for _, v := range value.Menus {
			ok := IsContain(idArray, v.ID)
			if !ok {
				menus = append(menus, v)
				idArray = append(idArray, v.ID)
			}
		}
	}
	return
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int, where interface{}, sort string) (roles []UmsRole, paging paginator.PagingAdmin) {
	paging = paginator.PaginateAdmin(
		c,
		database.DB.Model(UmsRole{}),
		&roles,
		where,
		perPage,
		sort,
	)
	return
}
