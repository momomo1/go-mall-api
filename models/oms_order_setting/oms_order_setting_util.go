package oms_order_setting

import (
    "go-mall-api/pkg/database"
)

func Get(idstr string) (omsOrderSetting OmsOrderSetting) {
    database.DB.Where("id", idstr).First(&omsOrderSetting)
    return
}

func GetBy(field, value string) (omsOrderSetting OmsOrderSetting) {
    database.DB.Where("? = ?", field, value).First(&omsOrderSetting)
    return
}

func All() (omsOrderSettings []OmsOrderSetting) {
    database.DB.Find(&omsOrderSettings)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(OmsOrderSetting{}).Where("? = ?", field, value).Count(&count)
    return count > 0
}