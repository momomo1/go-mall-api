package ums_role

import (
	"fmt"
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
	database.DB.Preload("Menus").Where("id in (?)", id).Find(&userRoleModel)
	for _, value := range userRoleModel {
		for _, v := range value.Menus {
			menus = append(menus, v)
		}
	}
	fmt.Println(menus)
	return
}
