package sms_flash_promotion_session

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/pkg/database"
	"go-mall-api/pkg/paginator"
)

func Get(idstr string) (smsFlashPromotionSession SmsFlashPromotionSession) {
	database.DB.Where("id", idstr).First(&smsFlashPromotionSession)
	return
}

func GetBy(field, value string) (smsFlashPromotionSession SmsFlashPromotionSession) {
	database.DB.Where("? = ?", field, value).First(&smsFlashPromotionSession)
	return
}

func All() (smsFlashPromotionSessions []SmsFlashPromotionSession) {
	database.DB.Find(&smsFlashPromotionSessions)
	return
}

func AllByStartUsing() (smsFlashPromotionSessions []SmsFlashPromotionSession) {
	database.DB.Where("status", 1).Find(&smsFlashPromotionSessions)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(SmsFlashPromotionSession{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int, where interface{}, sort string, order string) (flashPromotionSession []SmsFlashPromotionSession, paging paginator.PagingAdmin) {
	paging = paginator.PaginateAdmin(
		c,
		database.DB.Model(SmsFlashPromotionSession{}),
		&flashPromotionSession,
		where,
		perPage,
		sort,
		order,
	)
	return
}
