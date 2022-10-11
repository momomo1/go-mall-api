//Package sms_flash_promotion 模型
package sms_flash_promotion

import (
	"go-mall-api/models"
	"go-mall-api/pkg/database"
	"time"
)

//SmsFlashPromotion 限时购表
type SmsFlashPromotion struct {
	models.BaseModel

	Title      string    `gorm:"column:title;type:varchar(200)" json:"title"`
	StartDate  time.Time `gorm:"column:start_date;type:date;comment:开始日期" json:"start_date"`
	EndDate    time.Time `gorm:"column:end_date;type:date;comment:结束日期" json:"end_date"`
	Status     int       `gorm:"column:status;type:int(1);comment:上下线状态" json:"status"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;comment:秒杀时间段名称" json:"create_time"`
}

func (smsFlashPromotion *SmsFlashPromotion) TableName() string {
	return "sms_flash_promotion"
}

func (smsFlashPromotion *SmsFlashPromotion) Create() {
	database.DB.Create(&smsFlashPromotion)
}

func (smsFlashPromotion *SmsFlashPromotion) Save() (rowsAffected int64) {
	result := database.DB.Save(&smsFlashPromotion)
	return result.RowsAffected
}

func (smsFlashPromotion *SmsFlashPromotion) Updates(data map[string]interface{}) (rowsAffected int64) {
	result := database.DB.Model(&smsFlashPromotion).Updates(data)
	return result.RowsAffected
}

func (smsFlashPromotion *SmsFlashPromotion) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&smsFlashPromotion)
	return result.RowsAffected
}
