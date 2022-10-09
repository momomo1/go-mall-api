//Package sms_home_recommend_product 模型
package sms_home_recommend_product

import (
	"go-mall-api/models"
	"go-mall-api/pkg/database"
)

//SmsHomeRecommendProduct 人气推荐商品表
type SmsHomeRecommendProduct struct {
	models.BaseModel

	ProductId       int64  `gorm:"column:product_id;type:bigint(20)" json:"product_id"`
	ProductName     string `gorm:"column:product_name;type:varchar(64)" json:"product_name"`
	RecommendStatus int    `gorm:"column:recommend_status;type:int(1)" json:"recommend_status"`
	Sort            int    `gorm:"column:sort;type:int(1)" json:"sort"`
}

func (smsHomeRecommendProduct *SmsHomeRecommendProduct) TableName() string {
	return "sms_home_recommend_product"
}

func (smsHomeRecommendProduct *SmsHomeRecommendProduct) Create() {
	database.DB.Create(&smsHomeRecommendProduct)
}

func (smsHomeRecommendProduct *SmsHomeRecommendProduct) Save() (rowsAffected int64) {
	result := database.DB.Save(&smsHomeRecommendProduct)
	return result.RowsAffected
}

func (smsHomeRecommendProduct *SmsHomeRecommendProduct) Updates(data map[string]interface{}) (rowsAffected int64) {
	result := database.DB.Model(&smsHomeRecommendProduct).Updates(data)
	return result.RowsAffected
}

func (smsHomeRecommendProduct *SmsHomeRecommendProduct) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&smsHomeRecommendProduct)
	return result.RowsAffected
}
