package sms_home_brand

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/pkg/database"
	"go-mall-api/pkg/paginator"
)

func Get(idstr string) (smsHomeBrand SmsHomeBrand) {
	database.DB.Where("id", idstr).First(&smsHomeBrand)
	return
}

func GetByBrandId(value string) (smsHomeBrand SmsHomeBrand) {
	database.DB.Where("brand_id", value).First(&smsHomeBrand)
	return
}

func GetBy(field, value string) (smsHomeBrand SmsHomeBrand) {
	database.DB.Where("? = ?", field, value).First(&smsHomeBrand)
	return
}

func All() (smsHomeBrands []SmsHomeBrand) {
	database.DB.Find(&smsHomeBrands)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(SmsHomeBrand{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int, where interface{}, sort string, order string) (homeBrand []SmsHomeBrand, paging paginator.PagingAdmin) {
	paging = paginator.PaginateAdmin(
		c,
		database.DB.Model(SmsHomeBrand{}),
		&homeBrand,
		where,
		perPage,
		sort,
		order,
	)
	return
}
