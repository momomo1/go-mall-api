package sms_flash_promotion_product_relation

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/pkg/database"
	"go-mall-api/pkg/paginator"
)

func Get(idstr string) (smsFlashPromotionProductRelation SmsFlashPromotionProductRelation) {
	database.DB.Where("id", idstr).First(&smsFlashPromotionProductRelation)
	return
}

func GetBy(field, value string) (smsFlashPromotionProductRelation SmsFlashPromotionProductRelation) {
	database.DB.Where("? = ?", field, value).First(&smsFlashPromotionProductRelation)
	return
}

func All() (smsFlashPromotionProductRelations []SmsFlashPromotionProductRelation) {
	database.DB.Find(&smsFlashPromotionProductRelations)
	return
}

func GetWhereCount(flashPromotionId, flashPromotionSessionId string) (count int64) {
	database.DB.Model(SmsFlashPromotionProductRelation{}).Where("flash_promotion_id = ? and flash_promotion_session_id = ?", flashPromotionId, flashPromotionSessionId).Count(&count)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(SmsFlashPromotionProductRelation{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int, where interface{}, sort string, order string) (smsFlashPromotionProductRelations []SmsFlashPromotionProductRelation, paging paginator.PagingAdmin) {
	paging = paginator.PaginateAdmin(
		c,
		database.DB.Model(SmsFlashPromotionProductRelation{}),
		&smsFlashPromotionProductRelations,
		where,
		perPage,
		sort,
		order,
	)
	return
}
