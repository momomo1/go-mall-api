package oms_order_return_apply

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/pkg/database"
	"go-mall-api/pkg/paginator"
)

func Get(idstr string) (omsOrderReturnApply OmsOrderReturnApply) {
	database.DB.Where("id", idstr).First(&omsOrderReturnApply)
	return
}

func GetBy(field, value string) (omsOrderReturnApply OmsOrderReturnApply) {
	database.DB.Where("? = ?", field, value).First(&omsOrderReturnApply)
	return
}

func All() (omsOrderReturnApplies []OmsOrderReturnApply) {
	database.DB.Find(&omsOrderReturnApplies)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(OmsOrderReturnApply{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int, where interface{}, sort string, order string) (orderReturnApply []OmsOrderReturnApply, paging paginator.PagingAdmin) {
	paging = paginator.PaginateAdmin(
		c,
		database.DB.Model(OmsOrderReturnApply{}),
		&orderReturnApply,
		where,
		perPage,
		sort,
		order,
	)
	return
}
