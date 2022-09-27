//Package ums_resource 模型
package ums_resource

import (
	"go-mall-api/pkg/database"
	"time"
)

//UmsResource 后台资源表
type UmsResource struct {
	Id          int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`
	CreateTime  time.Time `gorm:"column:create_time;type:datetime;comment:创建时间" json:"create_time"`
	Name        string    `gorm:"column:name;type:varchar(200);comment:资源名称" json:"name"`
	Url         string    `gorm:"column:url;type:varchar(200);comment:资源URL" json:"url"`
	Description string    `gorm:"column:description;type:varchar(500);comment:描述" json:"description"`
	CategoryId  int64     `gorm:"column:category_id;type:bigint(20);comment:资源分类ID" json:"category_id"`
}

func (umsResource *UmsResource) TableName() string {
	return "ums_resource"
}

func (umsResource *UmsResource) Create() {
	database.DB.Create(&umsResource)
}

func (umsResource *UmsResource) Save() (rowsAffected int64) {
	result := database.DB.Save(&umsResource)
	return result.RowsAffected
}

func (umsResource *UmsResource) Updates(data map[string]interface{}) (rowsAffected int64) {
	result := database.DB.Model(&umsResource).Updates(data)
	return result.RowsAffected
}

func (umsResource *UmsResource) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&umsResource)
	return result.RowsAffected
}
