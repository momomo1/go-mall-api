//Package oms_order_item 模型
package oms_order_item

import (
	"go-mall-api/models"
	"go-mall-api/pkg/database"
)

//OmsOrderItem 订单中所包含的商品
type OmsOrderItem struct {
	models.BaseModel
	
	OrderId           int64   `gorm:"column:order_id;type:bigint(20);comment:订单id" json:"order_id"`
	OrderSn           string  `gorm:"column:order_sn;type:varchar(64);comment:订单编号" json:"order_sn"`
	ProductId         int64   `gorm:"column:product_id;type:bigint(20)" json:"product_id"`
	ProductPic        string  `gorm:"column:product_pic;type:varchar(500)" json:"product_pic"`
	ProductName       string  `gorm:"column:product_name;type:varchar(200)" json:"product_name"`
	ProductBrand      string  `gorm:"column:product_brand;type:varchar(200)" json:"product_brand"`
	ProductSn         string  `gorm:"column:product_sn;type:varchar(64)" json:"product_sn"`
	ProductPrice      float64 `gorm:"column:product_price;type:decimal(10,2);comment:销售价格" json:"product_price"`
	ProductQuantity   int     `gorm:"column:product_quantity;type:int(11);comment:购买数量" json:"product_quantity"`
	ProductSkuId      int64   `gorm:"column:product_sku_id;type:bigint(20);comment:商品sku编号" json:"product_sku_id"`
	ProductSkuCode    string  `gorm:"column:product_sku_code;type:varchar(50);comment:商品sku条码" json:"product_sku_code"`
	ProductCategoryId int64   `gorm:"column:product_category_id;type:bigint(20);comment:商品分类id" json:"product_category_id"`
	PromotionName     string  `gorm:"column:promotion_name;type:varchar(200);comment:商品促销名称" json:"promotion_name"`
	PromotionAmount   float64 `gorm:"column:promotion_amount;type:decimal(10,2);comment:商品促销分解金额" json:"promotion_amount"`
	CouponAmount      float64 `gorm:"column:coupon_amount;type:decimal(10,2);comment:优惠券优惠分解金额" json:"coupon_amount"`
	IntegrationAmount float64 `gorm:"column:integration_amount;type:decimal(10,2);comment:积分优惠分解金额" json:"integration_amount"`
	RealAmount        float64 `gorm:"column:real_amount;type:decimal(10,2);comment:该商品经过优惠后的分解金额" json:"real_amount"`
	GiftIntegration   int     `gorm:"column:gift_integration;type:int(11);default:0" json:"gift_integration"`
	GiftGrowth        int     `gorm:"column:gift_growth;type:int(11);default:0" json:"gift_growth"`
	ProductAttr       string  `gorm:"column:product_attr;type:varchar(500);comment:商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]" json:"product_attr"`
}

func (omsOrderItem *OmsOrderItem) TableName() string {
	return "oms_order_item"
}

func (omsOrderItem *OmsOrderItem) Create() {
	database.DB.Create(&omsOrderItem)
}

func (omsOrderItem *OmsOrderItem) Save() (rowsAffected int64) {
	result := database.DB.Save(&omsOrderItem)
	return result.RowsAffected
}

func (omsOrderItem *OmsOrderItem) Updates(data map[string]interface{}) (rowsAffected int64) {
	result := database.DB.Model(&omsOrderItem).Updates(data)
	return result.RowsAffected
}

func (omsOrderItem *OmsOrderItem) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&omsOrderItem)
	return result.RowsAffected
}
