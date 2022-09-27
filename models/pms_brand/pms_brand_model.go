//Package pms_brand 模型
package pms_brand

import (
	"go-mall-api/models"
	"go-mall-api/pkg/database"
)

//PmsBrand 品牌表
type PmsBrand struct {
	models.BaseModel

	Name                string `gorm:"column:name;type:varchar(64)" json:"name"`
	FirstLetter         string `gorm:"column:first_letter;type:varchar(8);comment:首字母" json:"first_letter"`
	Sort                int    `gorm:"column:sort;type:int(11)" json:"sort"`
	FactoryStatus       int    `gorm:"column:factory_status;type:int(1);comment:是否为品牌制造商：0->不是；1->是" json:"factory_status"`
	ShowStatus          int    `gorm:"column:show_status;type:int(1)" json:"show_status"`
	ProductCount        int    `gorm:"column:product_count;type:int(11);comment:产品数量" json:"product_count"`
	ProductCommentCount int    `gorm:"column:product_comment_count;type:int(11);comment:产品评论数量" json:"product_comment_count"`
	Logo                string `gorm:"column:logo;type:varchar(255);comment:品牌logo" json:"logo"`
	BigPic              string `gorm:"column:big_pic;type:varchar(255);comment:专区大图" json:"big_pic"`
	BrandStory          string `gorm:"column:brand_story;type:text;comment:品牌故事" json:"brand_story"`
}

func (pmsBrand *PmsBrand) TableName() string {
	return "pms_brand"
}

func (pmsBrand *PmsBrand) Create() {
	database.DB.Create(&pmsBrand)
}

func (pmsBrand *PmsBrand) Save() (rowsAffected int64) {
	result := database.DB.Save(&pmsBrand)
	return result.RowsAffected
}

func (pmsBrand *PmsBrand) Updates(data map[string]interface{}) (rowsAffected int64) {
	result := database.DB.Model(&pmsBrand).Updates(data)
	return result.RowsAffected
}

func (pmsBrand *PmsBrand) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&pmsBrand)
	return result.RowsAffected
}
