package pms_product

import (
    "go-mall-api/pkg/database"
)

func Get(idstr string) (pmsProduct PmsProduct) {
    database.DB.Where("id", idstr).First(&pmsProduct)
    return
}

func GetBy(field, value string) (pmsProduct PmsProduct) {
    database.DB.Where("? = ?", field, value).First(&pmsProduct)
    return
}

func All() (pmsProducts []PmsProduct) {
    database.DB.Find(&pmsProducts)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(PmsProduct{}).Where("? = ?", field, value).Count(&count)
    return count > 0
}