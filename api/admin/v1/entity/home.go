package entity

import (
	"go-mall-api/pkg/paginator"
	"time"
)

type HomeBrandListRequest struct {
	PageNum         int    `json:"pageNum"`
	PageSize        int    `json:"pageSize"`
	RecommendStatus string `json:"recommendStatus"`
	BrandName       string `json:"brandName"`
}

type HomeBrandListReply struct {
	List []HomeBrandList `json:"list"`
	paginator.PagingAdmin
}

type HomeBrandList struct {
	Id              uint64 `json:"id"`
	BrandId         int64  `json:"brandId"`
	RecommendStatus int    `json:"recommendStatus"`
	Sort            int    `json:"sort"`
	BrandName       string `json:"brandName"`
}

type HomeBrandCreateRequest struct {
	BrandId   int64  `json:"brandId"`
	BrandName string `json:"brandName"`
}

type HomeBrandUpdateRecommendStatusRequest struct {
	Ids             string `json:"ids"`
	RecommendStatus string `json:"recommendStatus"`
}

type HomeBrandUpdateSortRequest struct {
	Id   string `json:"id"`
	Sort string `json:"sort"`
}

type HomeBrandDeleteRequest struct {
	Ids string `json:"id"`
}

type HomeRecommendSubjectListRequest struct {
	PageNum         int    `json:"pageNum"`
	PageSize        int    `json:"pageSize"`
	RecommendStatus string `json:"recommendStatus"`
	SubjectName     string `json:"subjectName"`
}

type HomeRecommendSubjectListReply struct {
	List []HomeRecommendSubjectList `json:"list"`
	paginator.PagingAdmin
}

type HomeRecommendSubjectList struct {
	Id              uint64 `json:"id"`
	SubjectId       int64  `json:"subjectId"`
	RecommendStatus int    `json:"recommendStatus"`
	Sort            int    `json:"sort"`
	SubjectName     string `json:"subjectName"`
}

type HomeRecommendSubjectUpdateRecommendStatusRequest struct {
	Ids             string `json:"ids"`
	RecommendStatus string `json:"recommendStatus"`
}

type HomeRecommendSubjectCreateRequest struct {
	SubjectId   int64  `json:"subjectId"`
	SubjectName string `json:"subjectName"`
}

type HomeRecommendSubjectUpdateSortRequest struct {
	Id   string `json:"id"`
	Sort string `json:"sort"`
}

type HomeRecommendSubjectDeleteRequest struct {
	Ids string `json:"id"`
}

type HomeAdvertiseListRequest struct {
	PageNum  int    `json:"pageNum"`
	PageSize int    `json:"pageSize"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	EndTime  string `json:"endTime"`
}

type HomeAdvertiseListReply struct {
	List []HomeAdvertiseList `json:"list"`
	paginator.PagingAdmin
}

type HomeAdvertiseRequest struct {
	Id string `json:"id"`
}

type HomeAdvertiseList struct {
	Id         uint64    `json:"id"`
	ClickCount int       `json:"clickCount"`
	OrderCount int       `json:"orderCount"`
	Sort       int       `json:"sort"`
	Status     int       `json:"status"`
	Type       int       `json:"type"`
	Name       string    `json:"name"`
	Note       string    `json:"note"`
	Pic        string    `json:"pic"`
	Url        string    `json:"url"`
	StartTime  time.Time `json:"startTime"`
	EndTime    time.Time `json:"endTime"`
}

type HomeAdvertiseUpdateStatusRequest struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

type HomeAdvertiseCreateRequest struct {
	Type      int       `json:"type"`
	Sort      int       `json:"sort"`
	Status    int       `json:"status"`
	Name      string    `json:"name"`
	Note      string    `json:"note"`
	Pic       string    `json:"pic"`
	Url       string    `json:"url"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

type HomeAdvertiseUpdateRequest struct {
	Type      int       `json:"type"`
	Status    int       `json:"status"`
	Sort      string    `json:"sort"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Note      string    `json:"note"`
	Pic       string    `json:"pic"`
	Url       string    `json:"url"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

type HomeAdvertiseDeleteRequest struct {
	Ids string `json:"id"`
}

type HomeNewProductListRequest struct {
	PageNum         int    `json:"pageNum"`
	PageSize        int    `json:"pageSize"`
	RecommendStatus string `json:"recommendStatus"`
	ProductName     string `json:"productName"`
}

type HomeNewProductListReply struct {
	List []HomeNewProductList `json:"list"`
	paginator.PagingAdmin
}

type HomeNewProductList struct {
	Id              uint64 `json:"id"`
	ProductId       int64  `json:"productId"`
	RecommendStatus int    `json:"recommendStatus"`
	Sort            int    `json:"sort"`
	ProductName     string `json:"productName"`
}

type HomeNewProductUpdateRecommendStatusRequest struct {
	Ids             string `json:"ids"`
	RecommendStatus string `json:"recommendStatus"`
}

type HomeNewProductCreateRequest struct {
	ProductId   int64  `json:"productId"`
	ProductName string `json:"productName"`
}

type HomeNewProductUpdateSortRequest struct {
	Id   string `json:"id"`
	Sort string `json:"sort"`
}

type HomeNewProductDeleteRequest struct {
	Ids string `json:"ids"`
}

type HomeRecommendProductListRequest struct {
	PageNum         int    `json:"pageNum"`
	PageSize        int    `json:"pageSize"`
	RecommendStatus string `json:"recommendStatus"`
	ProductName     string `json:"productName"`
}

type HomeRecommendProductListReply struct {
	List []HomeRecommendProductList `json:"list"`
	paginator.PagingAdmin
}

type HomeRecommendProductList struct {
	Id              uint64 `json:"id"`
	ProductId       int64  `json:"productId"`
	RecommendStatus int    `json:"recommendStatus"`
	Sort            int    `json:"sort"`
	ProductName     string `json:"productName"`
}

type HomeRecommendProductUpdateRecommendStatusRequest struct {
	Ids             string `json:"ids"`
	RecommendStatus string `json:"recommendStatus"`
}
type HomeRecommendProductCreateRequest struct {
	ProductId   int64  `json:"productId"`
	ProductName string `json:"productName"`
}

type HomeRecommendProductUpdateSortRequest struct {
	Id   string `json:"id"`
	Sort string `json:"sort"`
}

type HomeRecommendProductDeleteRequest struct {
	Ids string `json:"ids"`
}
