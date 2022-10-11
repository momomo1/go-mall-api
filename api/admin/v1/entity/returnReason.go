package entity

import "go-mall-api/pkg/paginator"

type ReturnApplyListRequest struct {
	PageNum    int    `json:"pageNum"`
	PageSize   int    `json:"pageSize"`
	Id         string `json:"id"`
	Status     string `json:"status"`
	HandleMan  string `json:"handleMan"`
	CreateTime string `json:"createTime"`
	HandleTime string `json:"handleTime"`
}

type ReturnApplyListReply struct {
	List []ReturnApplyList `json:"list"`
	paginator.PagingAdmin
}

type ReturnApplyList struct {
	Id               uint64  `json:"id"`
	Status           int     `json:"status"`
	ProductRealPrice float64 `json:"productRealPrice"`
	ReturnName       string  `json:"returnName"`
	MemberUsername   string  `json:"memberUsername"`
	CreateTime       string  `json:"createTime"`
	HandleTime       string  `json:"handleTime"`
}

type ReturnReasonListRequest struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}

type ReturnReasonListReply struct {
	List []ReturnReasonList `json:"list"`
	paginator.PagingAdmin
}

type ReturnReasonList struct {
	Id         uint64 `json:"id"`
	Sort       int    `json:"sort"`
	Status     int    `json:"status"`
	Name       string `json:"name"`
	CreateTime string `json:"createTime"`
}

type ReturnReasonRequest struct {
	Id string `json:"id"`
}

type ReturnReasonUpdateStatusRequest struct {
	Ids    string `json:"ids"`
	Status string `json:"status"`
}

type ReturnReasonCreateRequest struct {
	Status int    `json:"status"`
	Name   string `json:"name"`
	Sort   string `json:"sort"`
}

type ReturnReasonUpdateRequest struct {
	Id     int    `json:"id"`
	Status int    `json:"status"`
	Name   string `json:"name"`
	Sort   string `json:"sort"`
}

type ReturnReasonDeleteRequest struct {
	Ids string `json:"ids"`
}
