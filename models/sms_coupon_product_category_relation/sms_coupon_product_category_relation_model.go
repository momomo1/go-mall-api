//Package sms_coupon_product_category_relation 模型
package sms_coupon_product_category_relation

import (
	"go-mall-api/models"
	"go-mall-api/pkg/database"
)

//SmsCouponProductCategoryRelation 优惠券和产品分类关系表
type SmsCouponProductCategoryRelation struct {
	models.BaseModel
	CouponId            int64  `gorm:"column:coupon_id;type:bigint(20)" json:"coupon_id"`
	ProductCategoryId   int64  `gorm:"column:product_category_id;type:bigint(20)" json:"product_category_id"`
	ProductCategoryName string `gorm:"column:product_category_name;type:varchar(200);comment:产品分类名称" json:"product_category_name"`
	ParentCategoryName  string `gorm:"column:parent_category_name;type:varchar(200);comment:父分类名称" json:"parent_category_name"`
}

func (smsCouponProductCategoryRelation *SmsCouponProductCategoryRelation) TableName() string {
	return "sms_coupon_product_category_relation"
}

func (smsCouponProductCategoryRelation *SmsCouponProductCategoryRelation) Create() {
	database.DB.Create(&smsCouponProductCategoryRelation)
}

func (smsCouponProductCategoryRelation *SmsCouponProductCategoryRelation) Save() (rowsAffected int64) {
	result := database.DB.Save(&smsCouponProductCategoryRelation)
	return result.RowsAffected
}

func (smsCouponProductCategoryRelation *SmsCouponProductCategoryRelation) Updates(data map[string]interface{}) (rowsAffected int64) {
	result := database.DB.Model(&smsCouponProductCategoryRelation).Updates(data)
	return result.RowsAffected
}

func (smsCouponProductCategoryRelation *SmsCouponProductCategoryRelation) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&smsCouponProductCategoryRelation)
	return result.RowsAffected
}
