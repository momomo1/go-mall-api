package pms_product_attribute

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/pkg/database"
	"go-mall-api/pkg/paginator"
)

func Get(idstr string) (pmsProductAttribute PmsProductAttribute) {
	database.DB.Where("id", idstr).First(&pmsProductAttribute)
	return
}

func GetBy(field, value string) (pmsProductAttribute PmsProductAttribute) {
	database.DB.Where("? = ?", field, value).First(&pmsProductAttribute)
	return
}

func All() (pmsProductAttributes []PmsProductAttribute) {
	database.DB.Find(&pmsProductAttributes)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(PmsProductAttribute{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int, where interface{}, sort string) (menus []PmsProductAttribute, paging paginator.PagingAdmin) {
	paging = paginator.PaginateAdmin(
		c,
		database.DB.Model(PmsProductAttribute{}),
		&menus,
		where,
		perPage,
		sort,
	)
	return
}
