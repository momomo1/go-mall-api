package pms_product_attribute_category

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/pkg/database"
	"go-mall-api/pkg/paginator"
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

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int, where interface{}, sort string, order string) (attributeCategory []PmsProductAttributeCategory, paging paginator.PagingAdmin) {
	paging = paginator.PaginateAdmin(
		c,
		database.DB.Model(PmsProductAttributeCategory{}),
		&attributeCategory,
		where,
		perPage,
		sort,
		order,
	)
	return
}
