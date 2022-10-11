//Package sms_flash_promotion_session 模型
package sms_flash_promotion_session

import (
	"go-mall-api/models"
	"go-mall-api/pkg/database"
	"time"
)

//SmsFlashPromotionSession 限时购场次表
type SmsFlashPromotionSession struct {
	models.BaseModel
	
	Name       string    `gorm:"column:name;type:varchar(200);comment:场次名称" json:"name"`
	StartTime  string    `gorm:"column:start_time;type:varchar(100);comment:每日开始时间" json:"start_time"`
	EndTime    string    `gorm:"column:end_time;type:varchar(100);comment:每日结束时间" json:"end_time"`
	Status     int       `gorm:"column:status;type:int(1);comment:启用状态：0->不启用；1->启用" json:"status"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;comment:创建时间" json:"create_time"`
}

func (m *SmsFlashPromotionSession) TableName() string {
	return "sms_flash_promotion_session"
}

func (smsFlashPromotionSession *SmsFlashPromotionSession) Create() {
	database.DB.Create(&smsFlashPromotionSession)
}

func (smsFlashPromotionSession *SmsFlashPromotionSession) Save() (rowsAffected int64) {
	result := database.DB.Save(&smsFlashPromotionSession)
	return result.RowsAffected
}

func (smsFlashPromotionSession *SmsFlashPromotionSession) Updates(data map[string]interface{}) (rowsAffected int64) {
	result := database.DB.Model(&smsFlashPromotionSession).Updates(data)
	return result.RowsAffected
}

func (smsFlashPromotionSession *SmsFlashPromotionSession) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&smsFlashPromotionSession)
	return result.RowsAffected
}
