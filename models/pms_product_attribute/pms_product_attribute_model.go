//Package pms_product_attribute 模型
package pms_product_attribute

import (
	"go-mall-api/models"
	"go-mall-api/pkg/database"
)

//PmsProductAttribute 商品属性参数表
type PmsProductAttribute struct {
	models.BaseModel
	ProductAttributeCategoryId int64  `gorm:"column:product_attribute_category_id;type:bigint(20)" json:"product_attribute_category_id"`
	Name                       string `gorm:"column:name;type:varchar(64)" json:"name"`
	SelectType                 int    `gorm:"column:select_type;type:int(1);comment:属性选择类型：0->唯一；1->单选；2->多选" json:"select_type"`
	InputType                  int    `gorm:"column:input_type;type:int(1);comment:属性录入方式：0->手工录入；1->从列表中选取" json:"input_type"`
	InputList                  string `gorm:"column:input_list;type:varchar(255);comment:可选值列表，以逗号隔开" json:"input_list"`
	Sort                       int    `gorm:"column:sort;type:int(11);comment:排序字段：最高的可以单独上传图片" json:"sort"`
	FilterType                 int    `gorm:"column:filter_type;type:int(1);comment:分类筛选样式：1->普通；1->颜色" json:"filter_type"`
	SearchType                 int    `gorm:"column:search_type;type:int(1);comment:检索类型；0->不需要进行检索；1->关键字检索；2->范围检索" json:"search_type"`
	RelatedStatus              int    `gorm:"column:related_status;type:int(1);comment:相同属性产品是否关联；0->不关联；1->关联" json:"related_status"`
	HandAddStatus              int    `gorm:"column:hand_add_status;type:int(1);comment:是否支持手动新增；0->不支持；1->支持" json:"hand_add_status"`
	Type                       int    `gorm:"column:type;type:int(1);comment:属性的类型；0->规格；1->参数" json:"type"`
}

func (pmsProductAttribute *PmsProductAttribute) TableName() string {
	return "pms_product_attribute"
}

func (pmsProductAttribute *PmsProductAttribute) Create() {
	database.DB.Create(&pmsProductAttribute)
}

func (pmsProductAttribute *PmsProductAttribute) Save() (rowsAffected int64) {
	result := database.DB.Save(&pmsProductAttribute)
	return result.RowsAffected
}

func (pmsProductAttribute *PmsProductAttribute) Updates(data map[string]interface{}) (rowsAffected int64) {
	result := database.DB.Model(&pmsProductAttribute).Updates(data)
	return result.RowsAffected
}

func (pmsProductAttribute *PmsProductAttribute) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&pmsProductAttribute)
	return result.RowsAffected
}
