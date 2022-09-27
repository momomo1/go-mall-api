package pms_product_attribute_category

import (
	"go-mall-api/pkg/database"
)

func Get(idstr string) (pmsProductAttributeCategory PmsProductAttributeCategory) {
	database.DB.Where("id", idstr).First(&pmsProductAttributeCategory)
	return
}

func GetBy(field, value string) (pmsProductAttributeCategory PmsProductAttributeCategory) {
	database.DB.Where("? = ?", field, value).First(&pmsProductAttributeCategory)
	return
}

func All() (pmsProductAttributeCategories []PmsProductAttributeCategory) {
	database.DB.Find(&pmsProductAttributeCategories)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(PmsProductAttributeCategory{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}
