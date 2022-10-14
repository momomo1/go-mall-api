package entity

import "go-mall-api/pkg/paginator"

type ProductListRequest struct {
	PageNum           int    `json:"pageNum"`
	PageSize          int    `json:"pageSize"`
	Keyword           string `json:"keyword"`
	PublishStatus     string `json:"publishStatus"`
	VerifyStatus      string `json:"verifyStatus"`
	ProductSn         string `json:"productSn"`
	ProductCategoryId string `json:"productCategoryId"`
	BrandId           string `json:"brandId"`
}

type ProductListReply struct {
	List []ProductList `json:"list"`
	paginator.PagingAdmin
}

type ProductList struct {
	Id                         uint64  `json:"id"`
	BrandId                    int64   `json:"brandId"`
	FeightTemplateId           int64   `json:"feightTemplateId"`
	ProductCategoryId          int64   `json:"productCategoryId"`
	ProductAttributeCategoryId int64   `json:"productAttributeCategoryId"`
	GiftGrowth                 int     `json:"giftGrowth"`
	GiftPoint                  int     `json:"giftPoint"`
	LowStock                   int     `json:"lowStock"`
	Sale                       int     `json:"sale"`
	Sort                       int     `json:"sort"`
	Stock                      int     `json:"stock"`
	Weight                     float64 `json:"weight"`
	PromotionType              int     `json:"promotionType"`
	NewStatus                  int     `json:"newStatus"`
	PreviewStatus              int     `json:"previewStatus"`
	DeleteStatus               int     `json:"deleteStatus"`
	PublishStatus              int     `json:"publishStatus"`
	VerifyStatus               int     `json:"verifyStatus"`
	RecommandStatus            int     `json:"recommandStatus"`
	UsePointLimit              int     `json:"usePointLimit"`
	PromotionPerLimit          int     `json:"promotionPerLimit"`
	OriginalPrice              float64 `json:"originalPrice"`
	Price                      float64 `json:"price"`
	AlbumPics                  string  `json:"albumPics"`
	BrandName                  string  `json:"brandName"`
	Description                string  `json:"description"`
	DetailDesc                 string  `json:"detailDesc"`
	DetailHtml                 string  `json:"detailHtml"`
	DetailMobileHtml           string  `json:"detailMobileHtml"`
	DetailTitle                string  `json:"detailTitle"`
	Keywords                   string  `json:"keywords"`
	Name                       string  `json:"name"`
	Note                       string  `json:"note"`
	Pic                        string  `json:"pic"`
	ProductCategoryName        string  `json:"productCategoryName"`
	ProductSn                  string  `json:"productSn"`
	ServiceIds                 string  `json:"serviceIds"`
	SubTitle                   string  `json:"subTitle"`
	Unit                       string  `json:"unit"`
	PromotionPrice             float64 `json:"promotionPrice"`
	PromotionStartTime         string  `json:"promotionStartTime"`
	PromotionEndTime           string  `json:"promotionEndTime"`
}

type ProductSimpleListRequest struct {
	Keyword string `json:"keyword"`
}
