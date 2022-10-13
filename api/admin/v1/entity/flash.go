package entity

import (
	"go-mall-api/pkg/paginator"
	"time"
)

type FlashListRequest struct {
	PageNum  int    `json:"pageNum"`
	PageSize int    `json:"pageSize"`
	Keyword  string `json:"keyword"`
}

type FlashListReply struct {
	List []FlashList `json:"list"`
	paginator.PagingAdmin
}

type FlashList struct {
	Id         uint64 `json:"id"`
	Status     int    `json:"status"`
	Title      string `json:"title"`
	CreateTime string `json:"createTime"`
	StartDate  string `json:"startDate"`
	EndDate    string `json:"endDate"`
}

type FlashUpdateStatusRequest struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

type FlashCreateRequest struct {
	Status    int       `json:"status"`
	Title     string    `json:"title"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}

type FlashUpdateRequest struct {
	Id        int    `json:"id"`
	Status    int    `json:"status"`
	Title     string `json:"title"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type FlashDeleteRequest struct {
	Id string `json:"id"`
}

type FlashSessionListReply struct {
	List []FlashSessionList `json:"list"`
}

type FlashSessionList struct {
	Id         uint64    `json:"id"`
	Status     int       `json:"status"`
	Name       string    `json:"name"`
	StartTime  string    `json:"startTime"`
	EndTime    string    `json:"endTime"`
	CreateTime time.Time `json:"createTime"`
}

type FlashSessionSelectListRequest struct {
	FlashPromotionId string `json:"flashPromotionId"`
}

type FlashSessionSelectListReply struct {
	List []FlashSessionSelectList `json:"list"`
}

type FlashSessionSelectList struct {
	Id           uint64    `json:"id"`
	ProductCount int64     `json:"productCount"`
	Status       int       `json:"status"`
	Name         string    `json:"name"`
	StartTime    string    `json:"startTime"`
	EndTime      string    `json:"endTime"`
	CreateTime   time.Time `json:"createTime"`
}

type FlashSessionCreateRequest struct {
	Name      string    `json:"name"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Status    int       `json:"status"`
}

type FlashSessionUpdateStatusRequest struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

type FlashSessionUpdateRequest struct {
	Id        int       `json:"id"`
	Status    int       `json:"status"`
	Name      string    `json:"name"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

type FlashSessionDeleteRequest struct {
	Id string `json:"id"`
}

type FlashProductRelationListRequest struct {
	PageNum                 int    `json:"pageNum"`
	PageSize                int    `json:"pageSize"`
	FlashPromotionId        string `json:"flashPromotionId"`
	FlashPromotionSessionId string `json:"flashPromotionSessionId"`
}

type FlashProductRelationListReply struct {
	List []FlashProductRelationList `json:"list"`
	paginator.PagingAdmin
}

type FlashProductRelationList struct {
	Id                      uint64        `json:"id"`
	Sort                    int           `json:"sort"`
	ProductId               int           `json:"productId"`
	FlashPromotionId        int           `json:"flashPromotionId"`
	FlashPromotionCount     int           `json:"flashPromotionCount"`
	FlashPromotionLimit     int           `json:"flashPromotionLimit"`
	FlashPromotionPrice     int           `json:"flashPromotionPrice"`
	FlashPromotionSessionId int           `json:"flashPromotionSessionId"`
	Product                 []interface{} `json:"product"`
}
