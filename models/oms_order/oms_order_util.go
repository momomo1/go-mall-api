package oms_order

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/pkg/database"
	"go-mall-api/pkg/paginator"
)

func Get(idstr string) (omsOrder OmsOrder) {
	database.DB.Where("id", idstr).First(&omsOrder)
	return
}

func Detail(idstr string) (omsOrder OmsOrder) {
	database.DB.Preload("OrderItemList").Preload("HistoryList").Where("id", idstr).First(&omsOrder)
	return
}

func GetByOrderSn(orderSn string) (omsOrder OmsOrder) {
	database.DB.Where("order_sn", orderSn).First(&omsOrder)
	return
}

func GetBy(field, value string) (omsOrder OmsOrder) {
	database.DB.Where("? = ?", field, value).First(&omsOrder)
	return
}

func All() (omsOrders []OmsOrder) {
	database.DB.Find(&omsOrders)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(OmsOrder{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int, where interface{}, sort string, order string) (omsOrder []OmsOrder, paging paginator.PagingAdmin) {
	paging = paginator.PaginateAdmin(
		c,
		database.DB.Model(OmsOrder{}),
		&omsOrder,
		where,
		perPage,
		sort,
		order,
	)
	return
}
