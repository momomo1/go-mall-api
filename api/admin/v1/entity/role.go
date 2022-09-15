package entity

import (
	"go-mall-api/models/ums_role"
	"go-mall-api/pkg/paginator"
)

type RoleListAllReply struct {
	Data []ums_role.UmsRole `json:"data"`
}

type RoleListRequest struct {
	PageNum  int    `json:"pageNum"`
	PageSize int    `json:"pageSize"`
	Keyword  string `json:"keyword"`
}

type RoleListReply struct {
	List []ums_role.UmsRole `json:"list"`
	paginator.PagingAdmin
}

type RoleUpdateStatusRequest struct {
	Id     string `json:"id"`
	Status int    `json:"status"`
}

type RoleRegisterRequest struct {
	Name        string `json:"name" valid:"name"`
	Description string `json:"description" valid:"description"`
	Status      int    `json:"status" valid:"status"`
}

type RoleDeleteRequest struct {
	Ids string `json:"ids"`
}

type RoleUpdateRequest struct {
	Id          string `json:"id" valid:"id"`
	Name        string `json:"name" valid:"name"`
	Description string `json:"description" valid:"description"`
	Status      int    `json:"status" valid:"status"`
}
