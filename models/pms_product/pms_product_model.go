//Package pms_product 模型
package pms_product

import (
	"go-mall-api/models"
	"go-mall-api/pkg/database"
	"time"
)

//PmsProduct 商品信息
type PmsProduct struct {
	models.BaseModel

	BrandId                    int64     `gorm:"column:brand_id;type:bigint(20)" json:"brand_id"`
	ProductCategoryId          int64     `gorm:"column:product_category_id;type:bigint(20)" json:"product_category_id"`
	FeightTemplateId           int64     `gorm:"column:feight_template_id;type:bigint(20)" json:"feight_template_id"`
	ProductAttributeCategoryId int64     `gorm:"column:product_attribute_category_id;type:bigint(20)" json:"product_attribute_category_id"`
	Name                       string    `gorm:"column:name;type:varchar(64);NOT NULL" json:"name"`
	Pic                        string    `gorm:"column:pic;type:varchar(255)" json:"pic"`
	ProductSn                  string    `gorm:"column:product_sn;type:varchar(64);comment:货号;NOT NULL" json:"product_sn"`
	DeleteStatus               int       `gorm:"column:delete_status;type:int(1);comment:删除状态：0->未删除；1->已删除" json:"delete_status"`
	PublishStatus              int       `gorm:"column:publish_status;type:int(1);comment:上架状态：0->下架；1->上架" json:"publish_status"`
	NewStatus                  int       `gorm:"column:new_status;type:int(1);comment:新品状态:0->不是新品；1->新品" json:"new_status"`
	RecommandStatus            int       `gorm:"column:recommand_status;type:int(1);comment:推荐状态；0->不推荐；1->推荐" json:"recommand_status"`
	VerifyStatus               int       `gorm:"column:verify_status;type:int(1);comment:审核状态：0->未审核；1->审核通过" json:"verify_status"`
	Sort                       int       `gorm:"column:sort;type:int(11);comment:排序" json:"sort"`
	Sale                       int       `gorm:"column:sale;type:int(11);comment:销量" json:"sale"`
	Price                      float64   `gorm:"column:price;type:decimal(10,2)" json:"price"`
	PromotionPrice             float64   `gorm:"column:promotion_price;type:decimal(10,2);comment:促销价格" json:"promotion_price"`
	GiftGrowth                 int       `gorm:"column:gift_growth;type:int(11);default:0;comment:赠送的成长值" json:"gift_growth"`
	GiftPoint                  int       `gorm:"column:gift_point;type:int(11);default:0;comment:赠送的积分" json:"gift_point"`
	UsePointLimit              int       `gorm:"column:use_point_limit;type:int(11);comment:限制使用的积分数" json:"use_point_limit"`
	SubTitle                   string    `gorm:"column:sub_title;type:varchar(255);comment:副标题" json:"sub_title"`
	Description                string    `gorm:"column:description;type:text;comment:商品描述" json:"description"`
	OriginalPrice              float64   `gorm:"column:original_price;type:decimal(10,2);comment:市场价" json:"original_price"`
	Stock                      int       `gorm:"column:stock;type:int(11);comment:库存" json:"stock"`
	LowStock                   int       `gorm:"column:low_stock;type:int(11);comment:库存预警值" json:"low_stock"`
	Unit                       string    `gorm:"column:unit;type:varchar(16);comment:单位" json:"unit"`
	Weight                     float64   `gorm:"column:weight;type:decimal(10,2);comment:商品重量，默认为克" json:"weight"`
	PreviewStatus              int       `gorm:"column:preview_status;type:int(1);comment:是否为预告商品：0->不是；1->是" json:"preview_status"`
	ServiceIds                 string    `gorm:"column:service_ids;type:varchar(64);comment:以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮" json:"service_ids"`
	Keywords                   string    `gorm:"column:keywords;type:varchar(255)" json:"keywords"`
	Note                       string    `gorm:"column:note;type:varchar(255)" json:"note"`
	AlbumPics                  string    `gorm:"column:album_pics;type:varchar(255);comment:画册图片，连产品图片限制为5张，以逗号分割" json:"album_pics"`
	DetailTitle                string    `gorm:"column:detail_title;type:varchar(255)" json:"detail_title"`
	DetailDesc                 string    `gorm:"column:detail_desc;type:text" json:"detail_desc"`
	DetailHtml                 string    `gorm:"column:detail_html;type:text;comment:产品详情网页内容" json:"detail_html"`
	DetailMobileHtml           string    `gorm:"column:detail_mobile_html;type:text;comment:移动端网页详情" json:"detail_mobile_html"`
	PromotionStartTime         time.Time `gorm:"column:promotion_start_time;type:datetime;comment:促销开始时间" json:"promotion_start_time"`
	PromotionEndTime           time.Time `gorm:"column:promotion_end_time;type:datetime;comment:促销结束时间" json:"promotion_end_time"`
	PromotionPerLimit          int       `gorm:"column:promotion_per_limit;type:int(11);comment:活动限购数量" json:"promotion_per_limit"`
	PromotionType              int       `gorm:"column:promotion_type;type:int(1);comment:促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购" json:"promotion_type"`
	BrandName                  string    `gorm:"column:brand_name;type:varchar(255);comment:品牌名称" json:"brand_name"`
	ProductCategoryName        string    `gorm:"column:product_category_name;type:varchar(255);comment:商品分类名称" json:"product_category_name"`
}

func (pmsProduct *PmsProduct) TableName() string {
	return "pms_product"
}

func (pmsProduct *PmsProduct) Create() {
	database.DB.Create(&pmsProduct)
}

func (pmsProduct *PmsProduct) Save() (rowsAffected int64) {
	result := database.DB.Save(&pmsProduct)
	return result.RowsAffected
}

func (pmsProduct *PmsProduct) Updates(data map[string]interface{}) (rowsAffected int64) {
	result := database.DB.Model(&pmsProduct).Updates(data)
	return result.RowsAffected
}

func (pmsProduct *PmsProduct) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&pmsProduct)
	return result.RowsAffected
}
