//Package sms_flash_promotion_product_relation 模型
package sms_flash_promotion_product_relation

import (
	"go-mall-api/models"
	"go-mall-api/pkg/database"
)

//SmsFlashPromotionProductRelation 商品限时购与商品关系表
type SmsFlashPromotionProductRelation struct {
	models.BaseModel
	FlashPromotionId        int64   `gorm:"column:flash_promotion_id;type:bigint(20)" json:"flash_promotion_id"`
	FlashPromotionSessionId int64   `gorm:"column:flash_promotion_session_id;type:bigint(20);comment:编号" json:"flash_promotion_session_id"`
	ProductId               int64   `gorm:"column:product_id;type:bigint(20)" json:"product_id"`
	FlashPromotionPrice     float64 `gorm:"column:flash_promotion_price;type:decimal(10,2);comment:限时购价格" json:"flash_promotion_price"`
	FlashPromotionCount     int     `gorm:"column:flash_promotion_count;type:int(11);comment:限时购数量" json:"flash_promotion_count"`
	FlashPromotionLimit     int     `gorm:"column:flash_promotion_limit;type:int(11);comment:每人限购数量" json:"flash_promotion_limit"`
	Sort                    int     `gorm:"column:sort;type:int(11);comment:排序" json:"sort"`
}

func (smsFlashPromotionProductRelation *SmsFlashPromotionProductRelation) TableName() string {
	return "sms_flash_promotion_product_relation"
}

func (smsFlashPromotionProductRelation *SmsFlashPromotionProductRelation) Create() {
	database.DB.Create(&smsFlashPromotionProductRelation)
}

func (smsFlashPromotionProductRelation *SmsFlashPromotionProductRelation) Save() (rowsAffected int64) {
	result := database.DB.Save(&smsFlashPromotionProductRelation)
	return result.RowsAffected
}

func (smsFlashPromotionProductRelation *SmsFlashPromotionProductRelation) Updates(data map[string]interface{}) (rowsAffected int64) {
	result := database.DB.Model(&smsFlashPromotionProductRelation).Updates(data)
	return result.RowsAffected
}

func (smsFlashPromotionProductRelation *SmsFlashPromotionProductRelation) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&smsFlashPromotionProductRelation)
	return result.RowsAffected
}
