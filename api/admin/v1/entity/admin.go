package entity

import (
	"go-mall-api/models/ums_admin"
	"go-mall-api/pkg/paginator"
)

type ListRequest struct {
	PageNum  int    `json:"pageNum"`
	PageSize int    `json:"pageSize"`
	Keyword  string `json:"keyword"`
}

type ListReply struct {
	List []ums_admin.UmsAdmin `json:"list"`
	paginator.PagingAdmin
}

type AdminUpdateStatusRequest struct {
	Id     string `json:"id"`
	Status int    `json:"status"`
}

type RegisterRequest struct {
	Username string `json:"username" valid:"username"`
	NickName string `json:"nickName" valid:"nickName"`
	Password string `json:"password" valid:"password"`
	Email    string `json:"email" valid:"email"`
	Note     string `json:"note" valid:"note"`
	Status   int    `json:"status" valid:"status"`
}
