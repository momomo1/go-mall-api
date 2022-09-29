package pms_product_category

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/pkg/database"
	"go-mall-api/pkg/paginator"
)

func Get(idstr string) (pmsProductCategory PmsProductCategory) {
	database.DB.Where("id", idstr).First(&pmsProductCategory)
	return
}

func GetBy(field, value string) (pmsProductCategory PmsProductCategory) {
	database.DB.Where("? = ?", field, value).First(&pmsProductCategory)
	return
}

func All() (pmsProductCategories []PmsProductCategory) {
	database.DB.Find(&pmsProductCategories)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(PmsProductCategory{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int, where interface{}, sort string, order string) (productCategory []PmsProductCategory, paging paginator.PagingAdmin) {
	paging = paginator.PaginateAdmin(
		c,
		database.DB.Model(PmsProductCategory{}),
		&productCategory,
		where,
		perPage,
		sort,
		order,
	)
	return
}
