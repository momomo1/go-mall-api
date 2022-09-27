//Package pms_product_attribute_category 模型
package pms_product_attribute_category

import (
	"go-mall-api/models"
	"go-mall-api/pkg/database"
)

//PmsProductAttributeCategory 产品属性分类表
type PmsProductAttributeCategory struct {
	models.BaseModel

	Name           string `gorm:"column:name;type:varchar(64)" json:"name"`
	AttributeCount int    `gorm:"column:attribute_count;type:int(11);default:0;comment:属性数量" json:"attribute_count"`
	ParamCount     int    `gorm:"column:param_count;type:int(11);default:0;comment:参数数量" json:"param_count"`
}

func (pmsProductAttributeCategory *PmsProductAttributeCategory) TableName() string {
	return "pms_product_attribute_category"
}

func (pmsProductAttributeCategory *PmsProductAttributeCategory) Create() {
	database.DB.Create(&pmsProductAttributeCategory)
}

func (pmsProductAttributeCategory *PmsProductAttributeCategory) Save() (rowsAffected int64) {
	result := database.DB.Save(&pmsProductAttributeCategory)
	return result.RowsAffected
}

func (pmsProductAttributeCategory *PmsProductAttributeCategory) Updates(data map[string]interface{}) (rowsAffected int64) {
	result := database.DB.Model(&pmsProductAttributeCategory).Updates(data)
	return result.RowsAffected
}

func (pmsProductAttributeCategory *PmsProductAttributeCategory) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&pmsProductAttributeCategory)
	return result.RowsAffected
}
