package ums_role_menu_relation

import (
    "go-mall-api/pkg/database"
)

func Get(idstr string) (umsRoleMenuRelation UmsRoleMenuRelation) {
    database.DB.Where("id", idstr).First(&umsRoleMenuRelation)
    return
}

func GetBy(field, value string) (umsRoleMenuRelation UmsRoleMenuRelation) {
    database.DB.Where("? = ?", field, value).First(&umsRoleMenuRelation)
    return
}

func All() (umsRoleMenuRelations []UmsRoleMenuRelation) {
    database.DB.Find(&umsRoleMenuRelations)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(UmsRoleMenuRelation{}).Where("? = ?", field, value).Count(&count)
    return count > 0
}