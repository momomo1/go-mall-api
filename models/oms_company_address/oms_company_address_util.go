package oms_company_address

import (
    "go-mall-api/pkg/database"
)

func Get(idstr string) (omsCompanyAddress OmsCompanyAddress) {
    database.DB.Where("id", idstr).First(&omsCompanyAddress)
    return
}

func GetBy(field, value string) (omsCompanyAddress OmsCompanyAddress) {
    database.DB.Where("? = ?", field, value).First(&omsCompanyAddress)
    return
}

func All() (omsCompanyAddresses []OmsCompanyAddress) {
    database.DB.Find(&omsCompanyAddresses)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(OmsCompanyAddress{}).Where("? = ?", field, value).Count(&count)
    return count > 0
}