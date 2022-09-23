package ums_resource_category

import (
	"go-mall-api/pkg/database"
)

func Get(idstr string) (umsResourceCategory UmsResourceCategory) {
	database.DB.Where("id", idstr).First(&umsResourceCategory)
	return
}

func GetBy(field, value string) (umsResourceCategory UmsResourceCategory) {
	database.DB.Where("? = ?", field, value).First(&umsResourceCategory)
	return
}

func All() (umsResourceCategories []UmsResourceCategory, count int64) {
	database.DB.Find(&umsResourceCategories)
	database.DB.Table("ums_resource_category").Count(&count)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(UmsResourceCategory{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}
