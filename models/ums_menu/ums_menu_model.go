//Package ums_menu 模型
package ums_menu

import (
	"go-mall-api/models"
	"go-mall-api/pkg/database"
	"time"
)

//UmsMenu 后台菜单表
type UmsMenu struct {
	models.BaseModel
	ParentId   int64     `gorm:"column:parent_id;type:bigint(20);comment:父级ID" json:"parentId"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;comment:创建时间" json:"create_time"`
	Title      string    `gorm:"column:title;type:varchar(100);comment:菜单名称" json:"title"`
	Level      int       `gorm:"column:level;type:int(4);comment:菜单级数" json:"level"`
	Sort       int       `gorm:"column:sort;type:int(4);comment:菜单排序" json:"sort"`
	Name       string    `gorm:"column:name;type:varchar(100);comment:前端名称" json:"name"`
	Icon       string    `gorm:"column:icon;type:varchar(200);comment:前端图标" json:"icon"`
	Hidden     int       `gorm:"column:hidden;type:int(1);comment:前端隐藏" json:"hidden"`
}

func (umsMenu *UmsMenu) TableName() string {
	return "ums_menu"
}

func (umsMenu *UmsMenu) Create() {
	database.DB.Create(&umsMenu)
}

func (umsRole *UmsMenu) Updates(data map[string]interface{}) (rowsAffected int64) {
	result := database.DB.Model(&umsRole).Updates(data)
	return result.RowsAffected
}

func (umsMenu *UmsMenu) Save() (rowsAffected int64) {
	result := database.DB.Save(&umsMenu)
	return result.RowsAffected
}

func (umsMenu *UmsMenu) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&umsMenu)
	return result.RowsAffected
}
