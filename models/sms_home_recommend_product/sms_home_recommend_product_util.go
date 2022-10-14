package sms_home_recommend_product

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/pkg/database"
	"go-mall-api/pkg/paginator"
)

func Get(idstr string) (smsHomeRecommendProduct SmsHomeRecommendProduct) {
	database.DB.Where("id", idstr).First(&smsHomeRecommendProduct)
	return
}

func GetByProductId(value string) (smsHomeRecommendProduct SmsHomeRecommendProduct) {
	database.DB.Where("product_id", value).First(&smsHomeRecommendProduct)
	return
}

func GetBy(field, value string) (smsHomeRecommendProduct SmsHomeRecommendProduct) {
	database.DB.Where("? = ?", field, value).First(&smsHomeRecommendProduct)
	return
}

func All() (smsHomeRecommendProducts []SmsHomeRecommendProduct) {
	database.DB.Find(&smsHomeRecommendProducts)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(SmsHomeRecommendProduct{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int, where interface{}, sort string, order string) (homeRecommendProduct []SmsHomeRecommendProduct, paging paginator.PagingAdmin) {
	paging = paginator.PaginateAdmin(
		c,
		database.DB.Model(SmsHomeRecommendProduct{}),
		&homeRecommendProduct,
		where,
		perPage,
		sort,
		order,
	)
	return
}
