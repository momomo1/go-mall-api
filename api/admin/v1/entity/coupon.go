package entity

import (
	"go-mall-api/pkg/paginator"
	"time"
)

type CouponListRequest struct {
	PageNum  int    `json:"pageNum"`
	PageSize int    `json:"pageSize"`
	Name     string `json:"name"`
	Type     string `json:"type"`
}

type CouponListReply struct {
	List []CouponList `json:"list"`
	paginator.PagingAdmin
}

type CouponList struct {
	Id                          uint64                        `json:"id"`
	Amount                      float64                       `json:"amount"`
	MinPoint                    float64                       `json:"minPoint"`
	Count                       int                           `json:"count"`
	PerLimit                    int                           `json:"perLimit"`
	Platform                    int                           `json:"platform"`
	PublishCount                int                           `json:"publishCount"`
	ReceiveCount                int                           `json:"receiveCount"`
	Type                        int                           `json:"type"`
	UseCount                    int                           `json:"useCount"`
	UseType                     int                           `json:"useType"`
	MemberLevel                 int                           `json:"memberLevel"`
	Name                        string                        `json:"name"`
	Note                        string                        `json:"note"`
	Code                        string                        `json:"code"`
	StartTime                   string                        `json:"startTime"`
	EnableTime                  string                        `json:"enableTime"`
	EndTime                     string                        `json:"endTime"`
	ProductRelationList         []ProductRelationList         `json:"productRelationList"`
	ProductCategoryRelationList []ProductCategoryRelationList `json:"productCategoryRelationList"`
}

type CouponRequest struct {
	Id string `json:"id"`
}

type CouponCreateRequest struct {
	Amount                      float64                       `json:"amount"`
	MinPoint                    float64                       `json:"minPoint"`
	PerLimit                    int                           `json:"perLimit"`
	Platform                    int                           `json:"platform"`
	PublishCount                int                           `json:"publishCount"`
	Type                        int                           `json:"type"`
	UseType                     int                           `json:"useType"`
	Name                        string                        `json:"name"`
	Note                        string                        `json:"note"`
	StartTime                   time.Time                     `json:"startTime"`
	EnableTime                  time.Time                     `json:"enableTime"`
	EndTime                     time.Time                     `json:"endTime"`
	ProductRelationList         []ProductRelationList         `json:"productRelationList"`
	ProductCategoryRelationList []ProductCategoryRelationList `json:"productCategoryRelationList"`
}

type CouponUpdateRequest struct {
	Id                          int                           `json:"id"`
	Amount                      float64                       `json:"amount"`
	MinPoint                    float64                       `json:"minPoint"`
	Platform                    int                           `json:"platform"`
	PublishCount                int                           `json:"publishCount"`
	Type                        int                           `json:"type"`
	UseType                     int                           `json:"useType"`
	PerLimit                    string                        `json:"perLimit"`
	Name                        string                        `json:"name"`
	Note                        string                        `json:"note"`
	StartTime                   string                        `json:"startTime"`
	EnableTime                  string                        `json:"enableTime"`
	EndTime                     string                        `json:"endTime"`
	ProductRelationList         []ProductRelationList         `json:"productRelationList"`
	ProductCategoryRelationList []ProductCategoryRelationList `json:"productCategoryRelationList"`
}

type ProductRelationList struct {
	ProductId   int64  `json:"productId"`
	ProductSn   string `json:"productSn"`
	ProductName string `json:"productName"`
}

type ProductCategoryRelationList struct {
	ProductCategoryId   int64  `json:"productCategoryId"`
	ParentCategoryName  string `json:"parentCategoryName"`
	ProductCategoryName string `json:"productCategoryName"`
}

type CouponDeleteRequest struct {
	Id string `json:"id"`
}

type CouponHistoryRequest struct {
	PageNum  int    `json:"pageNum"`
	PageSize int    `json:"pageSize"`
	CouponId string `json:"couponId"`
}

type CouponHistoryReply struct {
	List []CouponHistory `json:"list"`
	paginator.PagingAdmin
}

type CouponHistory struct {
	Id             uint64 `json:"id"`
	CouponId       int64  `json:"couponId"`
	GetType        int    `json:"getType"`
	MemberId       int64  `json:"memberId"`
	OrderId        int64  `json:"orderId"`
	UseStatus      int    `json:"useStatus"`
	CouponCode     string `json:"couponCode"`
	MemberNickname string `json:"memberNickname"`
	OrderSn        string `json:"orderSn"`
	UseTime        string `json:"useTime"`
	CreateTime     string `json:"createTime"`
}
