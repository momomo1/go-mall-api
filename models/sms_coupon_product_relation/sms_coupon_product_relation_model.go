//Package sms_coupon_product_relation 模型
package sms_coupon_product_relation

import (
	"go-mall-api/models"
	"go-mall-api/pkg/database"
)

//SmsCouponProductRelation 优惠券和产品的关系表
type SmsCouponProductRelation struct {
	models.BaseModel
	CouponId    int64  `gorm:"column:coupon_id;type:bigint(20)" json:"coupon_id"`
	ProductId   int64  `gorm:"column:product_id;type:bigint(20)" json:"product_id"`
	ProductName string `gorm:"column:product_name;type:varchar(500);comment:商品名称" json:"product_name"`
	ProductSn   string `gorm:"column:product_sn;type:varchar(200);comment:商品编码" json:"product_sn"`
}

func (smsCouponProductRelation *SmsCouponProductRelation) TableName() string {
	return "sms_coupon_product_relation"
}

func (smsCouponProductRelation *SmsCouponProductRelation) Create() {
	database.DB.Create(&smsCouponProductRelation)
}

func (smsCouponProductRelation *SmsCouponProductRelation) Save() (rowsAffected int64) {
	result := database.DB.Save(&smsCouponProductRelation)
	return result.RowsAffected
}

func (smsCouponProductRelation *SmsCouponProductRelation) Updates(data map[string]interface{}) (rowsAffected int64) {
	result := database.DB.Model(&smsCouponProductRelation).Updates(data)
	return result.RowsAffected
}

func (smsCouponProductRelation *SmsCouponProductRelation) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&smsCouponProductRelation)
	return result.RowsAffected
}
