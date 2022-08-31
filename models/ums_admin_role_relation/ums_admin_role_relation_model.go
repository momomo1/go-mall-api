//Package ums_admin_role_relation 模型
package ums_admin_role_relation

import (
    "go-mall-api/models"
    "go-mall-api/pkg/database"
)

//UmsAdminRoleRelation 后台用户和角色关系表
type UmsAdminRoleRelation struct {
    models.BaseModel

    AdminId int64 `gorm:"column:admin_id;type:bigint(20)" json:"admin_id"`
    RoleId  int64 `gorm:"column:role_id;type:bigint(20)" json:"role_id"`
}

func (umsAdminRoleRelation *UmsAdminRoleRelation) TableName() string {
    return "ums_admin_role_relation"
}

func (umsAdminRoleRelation *UmsAdminRoleRelation) Create() {
    database.DB.Create(&umsAdminRoleRelation)
}

func (umsAdminRoleRelation *UmsAdminRoleRelation) Save() (rowsAffected int64) {
    result := database.DB.Save(&umsAdminRoleRelation)
    return result.RowsAffected
}

func (umsAdminRoleRelation *UmsAdminRoleRelation) Delete() (rowsAffected int64) {
    result := database.DB.Delete(&umsAdminRoleRelation)
    return result.RowsAffected
}