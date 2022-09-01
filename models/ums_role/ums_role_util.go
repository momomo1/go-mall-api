package ums_role

import (
	"go-mall-api/pkg/database"
)

func Get(idstr string) (umsRole UmsRole) {
	database.DB.Where("id", idstr).First(&umsRole)
	return
}

func GetBy(field, value string) (umsRole UmsRole) {
	database.DB.Where("? = ?", field, value).First(&umsRole)
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
