package sms_coupon_product_category_relation

import (
	"go-mall-api/pkg/database"
)

func Get(idstr string) (smsCouponProductCategoryRelation SmsCouponProductCategoryRelation) {
	database.DB.Where("id", idstr).First(&smsCouponProductCategoryRelation)
	return
}

func GetByCouponId(couponId string) (smsCouponProductCategoryRelation []SmsCouponProductCategoryRelation) {
	database.DB.Where("coupon_id", couponId).Find(&smsCouponProductCategoryRelation)
	return
}

func GetBy(field, value string) (smsCouponProductCategoryRelation SmsCouponProductCategoryRelation) {
	database.DB.Where("? = ?", field, value).First(&smsCouponProductCategoryRelation)
	return
}

func All() (smsCouponProductCategoryRelations []SmsCouponProductCategoryRelation) {
	database.DB.Find(&smsCouponProductCategoryRelations)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(SmsCouponProductCategoryRelation{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}
