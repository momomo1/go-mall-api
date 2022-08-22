package ums_menu

import (
    "go-mall-api/pkg/database"
)

func Get(idstr string) (umsMenu UmsMenu) {
    database.DB.Where("id", idstr).First(&umsMenu)
    return
}

func GetBy(field, value string) (umsMenu UmsMenu) {
    database.DB.Where("? = ?", field, value).First(&umsMenu)
    return
}

func All() (umsMenus []UmsMenu) {
    database.DB.Find(&umsMenus)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(UmsMenu{}).Where("? = ?", field, value).Count(&count)
    return count > 0
}