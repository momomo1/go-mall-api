package pms_product

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/pkg/database"
	"go-mall-api/pkg/paginator"
)

func Get(idstr string) (pmsProduct PmsProduct) {
	database.DB.Where("id", idstr).First(&pmsProduct)
	return
}

func GetByWhereAll(where string) (pmsProduct []PmsProduct) {
	database.DB.Where(where).Find(&pmsProduct)
	return
}

func GetBy(field, value string) (pmsProduct PmsProduct) {
	database.DB.Where("? = ?", field, value).First(&pmsProduct)
	return
}

func All() (pmsProducts []PmsProduct) {
	database.DB.Find(&pmsProducts)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(PmsProduct{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int, where interface{}, sort string, order string) (pmsProducts []PmsProduct, paging paginator.PagingAdmin) {
	paging = paginator.PaginateAdmin(
		c,
		database.DB.Model(PmsProduct{}),
		&pmsProducts,
		where,
		perPage,
		sort,
		order,
	)
	return
}
