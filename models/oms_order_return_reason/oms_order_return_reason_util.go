package oms_order_return_reason

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/pkg/database"
	"go-mall-api/pkg/paginator"
)

func Get(idstr string) (omsOrderReturnReason OmsOrderReturnReason) {
	database.DB.Where("id", idstr).First(&omsOrderReturnReason)
	return
}

func GetBy(field, value string) (omsOrderReturnReason OmsOrderReturnReason) {
	database.DB.Where("? = ?", field, value).First(&omsOrderReturnReason)
	return
}

func All() (omsOrderReturnReasons []OmsOrderReturnReason) {
	database.DB.Find(&omsOrderReturnReasons)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(OmsOrderReturnReason{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int, where interface{}, sort string, order string) (orderReturnReason []OmsOrderReturnReason, paging paginator.PagingAdmin) {
	paging = paginator.PaginateAdmin(
		c,
		database.DB.Model(OmsOrderReturnReason{}),
		&orderReturnReason,
		where,
		perPage,
		sort,
		order,
	)
	return
}
