package entity

import (
	"go-mall-api/models/ums_admin"
	"go-mall-api/pkg/paginator"
)

type ListRequest struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}

type ListReply struct {
	List []ums_admin.UmsAdmin `json:"list"`
	paginator.PagingAdmin
}

type AdminUpdateStatusRequest struct {
	Id     string `json:"id"`
	Status int    `json:"status"`
}
