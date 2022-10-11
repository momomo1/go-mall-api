package sms_flash_promotion

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/pkg/database"
	"go-mall-api/pkg/paginator"
)

func Get(idstr string) (smsFlashPromotion SmsFlashPromotion) {
	database.DB.Where("id", idstr).First(&smsFlashPromotion)
	return
}

func GetBy(field, value string) (smsFlashPromotion SmsFlashPromotion) {
	database.DB.Where("? = ?", field, value).First(&smsFlashPromotion)
	return
}

func All() (smsFlashPromotions []SmsFlashPromotion) {
	database.DB.Find(&smsFlashPromotions)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(SmsFlashPromotion{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int, where interface{}, sort string, order string) (smsFlashPromotion []SmsFlashPromotion, paging paginator.PagingAdmin) {
	paging = paginator.PaginateAdmin(
		c,
		database.DB.Model(SmsFlashPromotion{}),
		&smsFlashPromotion,
		where,
		perPage,
		sort,
		order,
	)
	return
}
