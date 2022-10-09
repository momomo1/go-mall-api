//Package sms_home_brand 模型
package sms_home_brand

import (
	"go-mall-api/models"
	"go-mall-api/pkg/database"
)

//SmsHomeBrand 首页推荐品牌表
type SmsHomeBrand struct {
	models.BaseModel
	BrandId         int64  `gorm:"column:brand_id;type:bigint(20)" json:"brand_id"`
	BrandName       string `gorm:"column:brand_name;type:varchar(64)" json:"brand_name"`
	RecommendStatus int    `gorm:"column:recommend_status;type:int(1)" json:"recommend_status"`
	Sort            int    `gorm:"column:sort;type:int(11)" json:"sort"`
}

func (smsHomeBrand *SmsHomeBrand) TableName() string {
	return "sms_home_brand"
}

func (smsHomeBrand *SmsHomeBrand) Create() {
	database.DB.Create(&smsHomeBrand)
}

func (smsHomeBrand *SmsHomeBrand) Save() (rowsAffected int64) {
	result := database.DB.Save(&smsHomeBrand)
	return result.RowsAffected
}

func (smsHomeBrand *SmsHomeBrand) Updates(data map[string]interface{}) (rowsAffected int64) {
	result := database.DB.Model(&smsHomeBrand).Updates(data)
	return result.RowsAffected
}

func (smsHomeBrand *SmsHomeBrand) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&smsHomeBrand)
	return result.RowsAffected
}
