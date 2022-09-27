package ums_role_resource_relation

import (
	"go-mall-api/pkg/database"
)

func Get(idstr string) (umsRoleResourceRelation UmsRoleResourceRelation) {
	database.DB.Where("id", idstr).First(&umsRoleResourceRelation)
	return
}

func GetBy(field, value string) (umsRoleResourceRelation UmsRoleResourceRelation) {
	database.DB.Where("? = ?", field, value).First(&umsRoleResourceRelation)
	return
}

func All() (umsRoleResourceRelations []UmsRoleResourceRelation) {
	database.DB.Find(&umsRoleResourceRelations)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(UmsRoleResourceRelation{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

func GetByRoleIdFind(value string) (umsRoleResourceRelations []UmsRoleResourceRelation, count int64) {
	database.DB.Where("role_id = ?", value).Find(&umsRoleResourceRelations)
	database.DB.Table("ums_role_menu_relation").Where("role_id = ?", value).Count(&count)
	return
}

func DeleteByAdminId(roleId string) {
	database.DB.Where("role_id = ?", roleId).Delete(UmsRoleResourceRelation{})
}
