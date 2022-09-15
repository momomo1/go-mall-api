//Package ums_role 模型
package ums_role

import (
	"go-mall-api/models"
	"go-mall-api/models/ums_menu"
	"go-mall-api/pkg/database"
	"time"
)

//UmsRole 后台用户角色表
type UmsRole struct {
	models.BaseModel

	Name        string             `gorm:"column:name;type:varchar(100);comment:名称" json:"name"`
	Description string             `gorm:"column:description;type:varchar(500);comment:描述" json:"description"`
	AdminCount  int                `gorm:"column:admin_count;type:int(11);comment:后台用户数量" json:"admin_count"`
	CreateTime  time.Time          `gorm:"column:create_time;type:datetime;comment:创建时间" json:"create_time"`
	Status      *int               `gorm:"column:status;type:int(1);default:1;comment:启用状态：0->禁用；1->启用" json:"status"`
	Sort        int                `gorm:"column:sort;type:int(11);default:0" json:"sort"`
	Menus       []ums_menu.UmsMenu `gorm:"many2many:ums_role_menu_relation;joinForeignKey:RoleId;joinReferences:MenuId" json:"menus"`
}

func (umsRole *UmsRole) TableName() string {
	return "ums_role"
}

func (umsRole *UmsRole) Create() {
	database.DB.Create(&umsRole)
}

func (umsRole *UmsRole) Updates(data map[string]interface{}) (rowsAffected int64) {
	result := database.DB.Model(&umsRole).Updates(data)
	return result.RowsAffected
}

func (umsRole *UmsRole) Save() (rowsAffected int64) {
	result := database.DB.Save(&umsRole)
	return result.RowsAffected
}

func (umsRole *UmsRole) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&umsRole)
	return result.RowsAffected
}
