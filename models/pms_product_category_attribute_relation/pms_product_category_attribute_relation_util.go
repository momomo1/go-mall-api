package pms_product_category_attribute_relation

import (
	"go-mall-api/pkg/database"
)

func Get(idstr string) (pmsProductCategoryAttributeRelation PmsProductCategoryAttributeRelation) {
	database.DB.Where("id", idstr).First(&pmsProductCategoryAttributeRelation)
	return
}

func GetBy(field, value string) (pmsProductCategoryAttributeRelation PmsProductCategoryAttributeRelation) {
	database.DB.Where("? = ?", field, value).First(&pmsProductCategoryAttributeRelation)
	return
}

func All() (pmsProductCategoryAttributeRelations []PmsProductCategoryAttributeRelation) {
	database.DB.Find(&pmsProductCategoryAttributeRelations)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(PmsProductCategoryAttributeRelation{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

func GetRelationByCategoryId(productCategoryId string) (pmsProductCategoryAttributeRelations []PmsProductCategoryAttributeRelation) {
	database.DB.Where("product_category_id = ?", productCategoryId).Find(&pmsProductCategoryAttributeRelations)
	return
}
