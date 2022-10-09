//Package cms_subject 模型
package cms_subject

import (
	"go-mall-api/models"
	"go-mall-api/pkg/database"
	"time"
)

//CmsSubject 专题表
type CmsSubject struct {
	models.BaseModel

	CategoryId      int64     `gorm:"column:category_id;type:bigint(20)" json:"category_id"`
	Title           string    `gorm:"column:title;type:varchar(100)" json:"title"`
	Pic             string    `gorm:"column:pic;type:varchar(500);comment:专题主图" json:"pic"`
	ProductCount    int       `gorm:"column:product_count;type:int(11);comment:关联产品数量" json:"product_count"`
	RecommendStatus int       `gorm:"column:recommend_status;type:int(1)" json:"recommend_status"`
	CreateTime      time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`
	CollectCount    int       `gorm:"column:collect_count;type:int(11)" json:"collect_count"`
	ReadCount       int       `gorm:"column:read_count;type:int(11)" json:"read_count"`
	CommentCount    int       `gorm:"column:comment_count;type:int(11)" json:"comment_count"`
	AlbumPics       string    `gorm:"column:album_pics;type:varchar(1000);comment:画册图片用逗号分割" json:"album_pics"`
	Description     string    `gorm:"column:description;type:varchar(1000)" json:"description"`
	ShowStatus      int       `gorm:"column:show_status;type:int(1);comment:显示状态：0->不显示；1->显示" json:"show_status"`
	Content         string    `gorm:"column:content;type:text" json:"content"`
	ForwardCount    int       `gorm:"column:forward_count;type:int(11);comment:转发数" json:"forward_count"`
	CategoryName    string    `gorm:"column:category_name;type:varchar(200);comment:专题分类名称" json:"category_name"`
}

func (cmsSubject *CmsSubject) TableName() string {
	return "cms_subject"
}

func (cmsSubject *CmsSubject) Create() {
	database.DB.Create(&cmsSubject)
}

func (cmsSubject *CmsSubject) Save() (rowsAffected int64) {
	result := database.DB.Save(&cmsSubject)
	return result.RowsAffected
}

func (cmsSubject *CmsSubject) Updates(data map[string]interface{}) (rowsAffected int64) {
	result := database.DB.Model(&cmsSubject).Updates(data)
	return result.RowsAffected
}

func (cmsSubject *CmsSubject) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&cmsSubject)
	return result.RowsAffected
}
