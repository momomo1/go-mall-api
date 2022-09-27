//Package pms_product_category_attribute_relation 模型
package pms_product_category_attribute_relation

import (
	"go-mall-api/models"
	"go-mall-api/pkg/database"
)

//PmsProductCategoryAttributeRelation 产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）
type PmsProductCategoryAttributeRelation struct {
	models.BaseModel
	ProductCategoryId  int64 `gorm:"column:product_category_id;type:bigint(20)" json:"product_category_id"`
	ProductAttributeId int64 `gorm:"column:product_attribute_id;type:bigint(20)" json:"product_attribute_id"`
}

func (pmsProductCategoryAttributeRelation *PmsProductCategoryAttributeRelation) TableName() string {
	return "pms_product_category_attribute_relation"
}

func (pmsProductCategoryAttributeRelation *PmsProductCategoryAttributeRelation) Create() {
	database.DB.Create(&pmsProductCategoryAttributeRelation)
}

func (pmsProductCategoryAttributeRelation *PmsProductCategoryAttributeRelation) Save() (rowsAffected int64) {
	result := database.DB.Save(&pmsProductCategoryAttributeRelation)
	return result.RowsAffected
}

func (pmsProductCategoryAttributeRelation *PmsProductCategoryAttributeRelation) Updates(data map[string]interface{}) (rowsAffected int64) {
	result := database.DB.Model(&pmsProductCategoryAttributeRelation).Updates(data)
	return result.RowsAffected
}

func (pmsProductCategoryAttributeRelation *PmsProductCategoryAttributeRelation) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&pmsProductCategoryAttributeRelation)
	return result.RowsAffected
}
