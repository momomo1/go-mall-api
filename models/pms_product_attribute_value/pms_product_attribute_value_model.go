//Package pms_product_attribute_value 模型
package pms_product_attribute_value

import (
	"go-mall-api/models"
	"go-mall-api/pkg/database"
)

//PmsProductAttributeValue 存储产品参数信息的表
type PmsProductAttributeValue struct {
	models.BaseModel

	ProductId          int64  `gorm:"column:product_id;type:bigint(20)" json:"product_id"`
	ProductAttributeId int64  `gorm:"column:product_attribute_id;type:bigint(20)" json:"product_attribute_id"`
	Value              string `gorm:"column:value;type:varchar(64);comment:手动添加规格或参数的值，参数单值，规格有多个时以逗号隔开" json:"value"`
}

func (pmsProductAttributeValue *PmsProductAttributeValue) TableName() string {
	return "pms_product_attribute_value"
}

func (pmsProductAttributeValue *PmsProductAttributeValue) Create() {
	database.DB.Create(&pmsProductAttributeValue)
}

func (pmsProductAttributeValue *PmsProductAttributeValue) Save() (rowsAffected int64) {
	result := database.DB.Save(&pmsProductAttributeValue)
	return result.RowsAffected
}

func (pmsProductAttributeValue *PmsProductAttributeValue) Updates(data map[string]interface{}) (rowsAffected int64) {
	result := database.DB.Model(&pmsProductAttributeValue).Updates(data)
	return result.RowsAffected
}

func (pmsProductAttributeValue *PmsProductAttributeValue) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&pmsProductAttributeValue)
	return result.RowsAffected
}