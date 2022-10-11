package sms_flash_promotion_product_relation

import (
	"go-mall-api/pkg/database"
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
