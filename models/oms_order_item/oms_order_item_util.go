package oms_order_item

import (
    "go-mall-api/pkg/database"
)

func Get(idstr string) (omsOrderItem OmsOrderItem) {
    database.DB.Where("id", idstr).First(&omsOrderItem)
    return
}

func GetBy(field, value string) (omsOrderItem OmsOrderItem) {
    database.DB.Where("? = ?", field, value).First(&omsOrderItem)
    return
}

func All() (omsOrderItems []OmsOrderItem) {
    database.DB.Find(&omsOrderItems)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(OmsOrderItem{}).Where("? = ?", field, value).Count(&count)
    return count > 0
}