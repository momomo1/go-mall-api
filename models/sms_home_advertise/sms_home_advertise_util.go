package sms_home_advertise

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/pkg/database"
	"go-mall-api/pkg/paginator"
)

func Get(idstr string) (smsHomeAdvertise SmsHomeAdvertise) {
	database.DB.Where("id", idstr).First(&smsHomeAdvertise)
	return
}

func GetBy(field, value string) (smsHomeAdvertise SmsHomeAdvertise) {
	database.DB.Where("? = ?", field, value).First(&smsHomeAdvertise)
	return
}

func All() (smsHomeAdvertises []SmsHomeAdvertise) {
	database.DB.Find(&smsHomeAdvertises)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(SmsHomeAdvertise{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int, where interface{}, sort string, order string) (homeAdvertise []SmsHomeAdvertise, paging paginator.PagingAdmin) {
	paging = paginator.PaginateAdmin(
		c,
		database.DB.Model(SmsHomeAdvertise{}),
		&homeAdvertise,
		where,
		perPage,
		sort,
		order,
	)
	return
}
