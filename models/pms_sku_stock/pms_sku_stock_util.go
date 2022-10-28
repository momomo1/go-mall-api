package pms_sku_stock

import (
	"go-mall-api/pkg/database"
)

func Get(idstr string) (pmsSkuStock PmsSkuStock) {
	database.DB.Where("id", idstr).First(&pmsSkuStock)
	return
}

func GetBy(field, value string) (pmsSkuStock PmsSkuStock) {
	database.DB.Where("? = ?", field, value).First(&pmsSkuStock)
	return
}

func All() (pmsSkuStocks []PmsSkuStock) {
	database.DB.Find(&pmsSkuStocks)
	return
}

func GetByProductIdAll(where map[string]interface{}) (pmsSkuStocks []PmsSkuStock) {
	database.DB.Where(where).Find(&pmsSkuStocks)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(PmsSkuStock{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}
