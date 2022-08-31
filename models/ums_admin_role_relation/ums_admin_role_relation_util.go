package ums_admin_role_relation

import (
    "go-mall-api/pkg/database"
)

func Get(idstr string) (umsAdminRoleRelation UmsAdminRoleRelation) {
    database.DB.Where("id", idstr).First(&umsAdminRoleRelation)
    return
}

func GetBy(field, value string) (umsAdminRoleRelation UmsAdminRoleRelation) {
    database.DB.Where("? = ?", field, value).First(&umsAdminRoleRelation)
    return
}

func All() (umsAdminRoleRelations []UmsAdminRoleRelation) {
    database.DB.Find(&umsAdminRoleRelations)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(UmsAdminRoleRelation{}).Where("? = ?", field, value).Count(&count)
    return count > 0
}