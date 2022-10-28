//Package pms_sku_stock 模型
package pms_sku_stock

import (
	"go-mall-api/models"
	"go-mall-api/pkg/database"
)

//PmsSkuStock sku的库存
type PmsSkuStock struct {
	models.BaseModel
	ProductId      int64   `gorm:"column:product_id;type:bigint(20)" json:"product_id"`
	SkuCode        string  `gorm:"column:sku_code;type:varchar(64);comment:sku编码;NOT NULL" json:"sku_code"`
	Price          float64 `gorm:"column:price;type:decimal(10,2)" json:"price"`
	Stock          int     `gorm:"column:stock;type:int(11);default:0;comment:库存" json:"stock"`
	LowStock       int     `gorm:"column:low_stock;type:int(11);comment:预警库存" json:"low_stock"`
	Pic            string  `gorm:"column:pic;type:varchar(255);comment:展示图片" json:"pic"`
	Sale           int     `gorm:"column:sale;type:int(11);comment:销量" json:"sale"`
	PromotionPrice float64 `gorm:"column:promotion_price;type:decimal(10,2);comment:单品促销价格" json:"promotion_price"`
	LockStock      int     `gorm:"column:lock_stock;type:int(11);default:0;comment:锁定库存" json:"lock_stock"`
	SpData         string  `gorm:"column:sp_data;type:varchar(500);comment:商品销售属性，json格式" json:"sp_data"`
}

func (pmsSkuStock *PmsSkuStock) TableName() string {
	return "pms_sku_stock"
}

func (pmsSkuStock *PmsSkuStock) Create() {
	database.DB.Create(&pmsSkuStock)
}

func (pmsSkuStock *PmsSkuStock) Save() (rowsAffected int64) {
	result := database.DB.Save(&pmsSkuStock)
	return result.RowsAffected
}

func (pmsSkuStock *PmsSkuStock) Updates(data map[string]interface{}) (rowsAffected int64) {
	result := database.DB.Model(&pmsSkuStock).Updates(data)
	return result.RowsAffected
}

func (pmsSkuStock *PmsSkuStock) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&pmsSkuStock)
	return result.RowsAffected
}
