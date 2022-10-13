package sms_coupon_product_relation

import (
	"go-mall-api/pkg/database"
)

func Get(idstr string) (smsCouponProductRelation SmsCouponProductRelation) {
	database.DB.Where("id", idstr).First(&smsCouponProductRelation)
	return
}

func GetByCouponId(couponId string) (smsCouponProductRelation []SmsCouponProductRelation) {
	database.DB.Where("coupon_id", couponId).Find(&smsCouponProductRelation)
	return
}

func GetBy(field, value string) (smsCouponProductRelation SmsCouponProductRelation) {
	database.DB.Where("? = ?", field, value).First(&smsCouponProductRelation)
	return
}

func All() (smsCouponProductRelations []SmsCouponProductRelation) {
	database.DB.Find(&smsCouponProductRelations)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(SmsCouponProductRelation{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}
