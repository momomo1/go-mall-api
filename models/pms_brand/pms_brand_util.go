package pms_brand

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/pkg/database"
	"go-mall-api/pkg/paginator"
)

func Get(idstr string) (pmsBrand PmsBrand) {
	database.DB.Where("id", idstr).First(&pmsBrand)
	return
}

func GetBy(field, value string) (pmsBrand PmsBrand) {
	database.DB.Where("? = ?", field, value).First(&pmsBrand)
	return
}

func All() (pmsBrands []PmsBrand) {
	database.DB.Find(&pmsBrands)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(PmsBrand{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int, where interface{}, sort string) (menus []PmsBrand, paging paginator.PagingAdmin) {
	paging = paginator.PaginateAdmin(
		c,
		database.DB.Model(PmsBrand{}),
		&menus,
		where,
		perPage,
		sort,
	)
	return
}
