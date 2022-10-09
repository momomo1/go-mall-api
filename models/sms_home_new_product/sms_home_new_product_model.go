//Package sms_home_new_product 模型
package sms_home_new_product

import (
	"go-mall-api/models"
	"go-mall-api/pkg/database"
)

//SmsHomeNewProduct 新鲜好物表
type SmsHomeNewProduct struct {
	models.BaseModel

	ProductId       int64  `gorm:"column:product_id;type:bigint(20)" json:"product_id"`
	ProductName     string `gorm:"column:product_name;type:varchar(64)" json:"product_name"`
	RecommendStatus int    `gorm:"column:recommend_status;type:int(1)" json:"recommend_status"`
	Sort            int    `gorm:"column:sort;type:int(1)" json:"sort"`
}

func (smsHomeNewProduct *SmsHomeNewProduct) TableName() string {
	return "sms_home_new_product"
}

func (smsHomeNewProduct *SmsHomeNewProduct) Create() {
	database.DB.Create(&smsHomeNewProduct)
}

func (smsHomeNewProduct *SmsHomeNewProduct) Save() (rowsAffected int64) {
	result := database.DB.Save(&smsHomeNewProduct)
	return result.RowsAffected
}

func (smsHomeNewProduct *SmsHomeNewProduct) Updates(data map[string]interface{}) (rowsAffected int64) {
	result := database.DB.Model(&smsHomeNewProduct).Updates(data)
	return result.RowsAffected
}

func (smsHomeNewProduct *SmsHomeNewProduct) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&smsHomeNewProduct)
	return result.RowsAffected
}
