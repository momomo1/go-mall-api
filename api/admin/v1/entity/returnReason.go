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

type ReturnApplyRequest struct {
	Id string `json:"id"`
}

type ReturnApplyReply struct {
	Id               uint64  `json:"id"`
	OrderId          int64   `json:"orderId"`
	ProductCount     int     `json:"productCount"`
	ProductId        int64   `json:"productId"`
	Status           int     `json:"status"`
	CompanyAddressId int64   `json:"companyAddressId"`
	ProductPrice     float64 `json:"productPrice"`
	ProductRealPrice float64 `json:"productRealPrice"`
	ReturnAmount     float64 `json:"returnAmount"`
	CompanyAddress   string  `json:"companyAddress"`
	Description      string  `json:"description"`
	HandleMan        string  `json:"handleMan"`
	HandleNote       string  `json:"handleNote"`
	HandleTime       string  `json:"handleTime"`
	MemberUsername   string  `json:"memberUsername"`
	OrderSn          string  `json:"orderSn"`
	ProductAttr      string  `json:"productAttr"`
	ProductBrand     string  `json:"productBrand"`
	ProductName      string  `json:"productName"`
	ProductPic       string  `json:"productPic"`
	ProofPics        string  `json:"proofPics"`
	Reason           string  `json:"reason"`
	ReceiveMan       string  `json:"receiveMan"`
	ReceiveNote      string  `json:"receiveNote"`
	ReceiveTime      string  `json:"receiveTime"`
	ReturnName       string  `json:"returnName"`
	ReturnPhone      string  `json:"returnPhone"`
}

type ReturnApplyUpdateStatusRequest struct {
	CompanyAddressId int    `json:"companyAddressId"`
	Status           int    `json:"status"`
	Id               string `json:"id"`
	ReceiveNote      string `json:"receiveNote"`
	HandleNote       string `json:"handleNote"`
	ReturnAmount     string `json:"returnAmount"`
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
