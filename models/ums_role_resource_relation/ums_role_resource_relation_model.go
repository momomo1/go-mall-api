//Package ums_role_resource_relation 模型
package ums_role_resource_relation

import (
	"go-mall-api/pkg/database"
)

//UmsRoleResourceRelation 后台角色资源关系表
type UmsRoleResourceRelation struct {
	Id         int64 `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`
	RoleId     int64 `gorm:"column:role_id;type:bigint(20);comment:角色ID" json:"role_id"`
	ResourceId int64 `gorm:"column:resource_id;type:bigint(20);comment:资源ID" json:"resource_id"`
}

func (umsRoleResourceRelation *UmsRoleResourceRelation) TableName() string {
	return "ums_role_resource_relation"
}

func (umsRoleResourceRelation *UmsRoleResourceRelation) Create() {
	database.DB.Create(&umsRoleResourceRelation)
}

func (umsRoleResourceRelation *UmsRoleResourceRelation) Save() (rowsAffected int64) {
	result := database.DB.Save(&umsRoleResourceRelation)
	return result.RowsAffected
}

func (umsRoleResourceRelation *UmsRoleResourceRelation) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&umsRoleResourceRelation)
	return result.RowsAffected
}
