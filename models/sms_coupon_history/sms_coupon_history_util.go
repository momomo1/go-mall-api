package sms_coupon_history

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/pkg/database"
	"go-mall-api/pkg/paginator"
)

func Get(idstr string) (smsCouponHistory SmsCouponHistory) {
	database.DB.Where("id", idstr).First(&smsCouponHistory)
	return
}

func GetBy(field, value string) (smsCouponHistory SmsCouponHistory) {
	database.DB.Where("? = ?", field, value).First(&smsCouponHistory)
	return
}

func All() (smsCouponHistories []SmsCouponHistory) {
	database.DB.Find(&smsCouponHistories)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(SmsCouponHistory{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int, where interface{}, sort string, order string) (smsCouponHistory []SmsCouponHistory, paging paginator.PagingAdmin) {
	paging = paginator.PaginateAdmin(
		c,
		database.DB.Model(SmsCouponHistory{}),
		&smsCouponHistory,
		where,
		perPage,
		sort,
		order,
	)
	return
}
