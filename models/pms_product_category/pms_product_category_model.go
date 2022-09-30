//Package pms_product_category 模型
package pms_product_category

import (
	"go-mall-api/models"
	"go-mall-api/pkg/database"
)

//PmsProductCategory 产品分类
type PmsProductCategory struct {
	models.BaseModel

	ParentId     int64  `gorm:"column:parent_id;type:bigint(20);comment:上机分类的编号：0表示一级分类" json:"parent_id"`
	Name         string `gorm:"column:name;type:varchar(64)" json:"name"`
	Level        int    `gorm:"column:level;type:int(1);comment:分类级别：0->1级；1->2级" json:"level"`
	ProductCount int    `gorm:"column:product_count;type:int(11)" json:"product_count"`
	ProductUnit  string `gorm:"column:product_unit;type:varchar(64)" json:"product_unit"`
	NavStatus    int    `gorm:"column:nav_status;type:int(1);comment:是否显示在导航栏：0->不显示；1->显示" json:"nav_status"`
	ShowStatus   int    `gorm:"column:show_status;type:int(1);comment:显示状态：0->不显示；1->显示" json:"show_status"`
	Sort         int    `gorm:"column:sort;type:int(11)" json:"sort"`
	Icon         string `gorm:"column:icon;type:varchar(255);comment:图标" json:"icon"`
	Keywords     string `gorm:"column:keywords;type:varchar(255)" json:"keywords"`
	Description  string `gorm:"column:description;type:text;comment:描述" json:"description"`
}

func (pmsProductCategory *PmsProductCategory) TableName() string {
	return "pms_product_category"
}

func (pmsProductCategory *PmsProductCategory) Create() {
	database.DB.Create(&pmsProductCategory)
	return
}

func (pmsProductCategory *PmsProductCategory) Save() (rowsAffected int64) {
	result := database.DB.Save(&pmsProductCategory)
	return result.RowsAffected
}

func (pmsProductCategory *PmsProductCategory) Updates(data map[string]interface{}) (rowsAffected int64) {
	result := database.DB.Model(&pmsProductCategory).Updates(data)
	return result.RowsAffected
}

func (pmsProductCategory *PmsProductCategory) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&pmsProductCategory)
	return result.RowsAffected
}
