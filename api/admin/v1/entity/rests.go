package entity

import "go-mall-api/pkg/paginator"

type SubjectListRequest struct {
	PageNum  int    `json:"pageNum"`
	PageSize int    `json:"pageSize"`
	Keyword  string `json:"keyword"`
}

type SubjectListReply struct {
	List []SubjectList `json:"list"`
	paginator.PagingAdmin
}

type SubjectList struct {
	Id           uint64 `json:"id"`
	CategoryId   int64  `json:"categoryId"`
	Title        string `json:"title"`
	CategoryName string `json:"categoryName"`
	CreateTime   string `json:"createTime"`
}

type CompanyAddressListReply struct {
	List []CompanyAddressList `json:"list"`
}

type CompanyAddressList struct {
	Id            uint64 `json:"id"`
	ReceiveStatus int    `json:"receiveStatus"`
	SendStatus    int    `json:"sendStatus"`
	AddressName   string `json:"addressName"`
	City          string `json:"city"`
	DetailAddress string `json:"detailAddress"`
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	Province      string `json:"province"`
	Region        string `json:"region"`
}
