//Package ums_resource_category 模型
package ums_resource_category

import (
	"go-mall-api/pkg/database"
	"time"
)

//UmsResourceCategory 资源分类表
type UmsResourceCategory struct {
	Id         int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;comment:创建时间" json:"create_time"`
	Name       string    `gorm:"column:name;type:varchar(200);comment:分类名称" json:"name"`
	Sort       int       `gorm:"column:sort;type:int(4);comment:排序" json:"sort"`
}

func (umsResourceCategory *UmsResourceCategory) TableName() string {
	return "ums_resource_category"
}

func (umsResourceCategory *UmsResourceCategory) Create() {
	database.DB.Create(&umsResourceCategory)
}

func (umsResourceCategory *UmsResourceCategory) Updates(data map[string]interface{}) (rowsAffected int64) {
	result := database.DB.Model(&umsResourceCategory).Updates(data)
	return result.RowsAffected
}

func (umsResourceCategory *UmsResourceCategory) Save() (rowsAffected int64) {
	result := database.DB.Save(&umsResourceCategory)
	return result.RowsAffected
}

func (umsResourceCategory *UmsResourceCategory) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&umsResourceCategory)
	return result.RowsAffected
}
