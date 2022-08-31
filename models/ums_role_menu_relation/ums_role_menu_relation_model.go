//Package ums_role_menu_relation 模型
package ums_role_menu_relation

import (
	"go-mall-api/models"
	"go-mall-api/pkg/database"
)

//UmsRoleMenuRelation 后台角色菜单关系表
type UmsRoleMenuRelation struct {
	models.BaseModel
	RoleId int64 `gorm:"column:role_id;type:bigint(20);comment:角色ID" json:"role_id"`
	MenuId int64 `gorm:"column:menu_id;type:bigint(20);comment:菜单ID" json:"menu_id"`
}

func (umsRoleMenuRelation *UmsRoleMenuRelation) TableName() string {
	return "ums_role_menu_relation"
}

func (umsRoleMenuRelation *UmsRoleMenuRelation) Create() {
	database.DB.Create(&umsRoleMenuRelation)
}

func (umsRoleMenuRelation *UmsRoleMenuRelation) Save() (rowsAffected int64) {
	result := database.DB.Save(&umsRoleMenuRelation)
	return result.RowsAffected
}

func (umsRoleMenuRelation *UmsRoleMenuRelation) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&umsRoleMenuRelation)
	return result.RowsAffected
}
