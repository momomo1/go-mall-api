package pms_product_attribute_value

import (
    "go-mall-api/pkg/database"
)

func Get(idstr string) (pmsProductAttributeValue PmsProductAttributeValue) {
    database.DB.Where("id", idstr).First(&pmsProductAttributeValue)
    return
}

func GetBy(field, value string) (pmsProductAttributeValue PmsProductAttributeValue) {
    database.DB.Where("? = ?", field, value).First(&pmsProductAttributeValue)
    return
}

func All() (pmsProductAttributeValues []PmsProductAttributeValue) {
    database.DB.Find(&pmsProductAttributeValues)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(PmsProductAttributeValue{}).Where("? = ?", field, value).Count(&count)
    return count > 0
}