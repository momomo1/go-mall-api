package sms_home_new_product

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/pkg/database"
	"go-mall-api/pkg/paginator"
)

func Get(idstr string) (smsHomeNewProduct SmsHomeNewProduct) {
	database.DB.Where("id", idstr).First(&smsHomeNewProduct)
	return
}

func GetByProductId(value string) (smsHomeNewProduct SmsHomeNewProduct) {
	database.DB.Where("product_id", value).First(&smsHomeNewProduct)
	return
}

func GetBy(field, value string) (smsHomeNewProduct SmsHomeNewProduct) {
	database.DB.Where("? = ?", field, value).First(&smsHomeNewProduct)
	return
}

func All() (smsHomeNewProducts []SmsHomeNewProduct) {
	database.DB.Find(&smsHomeNewProducts)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(SmsHomeNewProduct{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int, where interface{}, sort string, order string) (homeNewProduct []SmsHomeNewProduct, paging paginator.PagingAdmin) {
	paging = paginator.PaginateAdmin(
		c,
		database.DB.Model(SmsHomeNewProduct{}),
		&homeNewProduct,
		where,
		perPage,
		sort,
		order,
	)
	return
}
