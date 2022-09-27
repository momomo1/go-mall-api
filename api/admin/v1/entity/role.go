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
	List []RoleList `json:"list"`
	paginator.PagingAdmin
}

type RoleList struct {
	Id          uint64 `json:"id"`
	AdminCount  int    `json:"adminCount"`
	Sort        int    `json:"sort"`
	Status      *int   `json:"status"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreateTime  string `json:"createTime"`
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

type RoleListMenuRequest struct {
	Id string `json:"id"`
}

type RoleListMenuReply struct {
	Id int64 `json:"id"`
}

type RoleAllocMenuRequest struct {
	RoleId  string `json:"roleId"`
	MenuIds string `json:"menuIds"`
}

type RoleListResourceRequest struct {
	Id string `json:"id"`
}

type RoleListResourceReply struct {
	Id int64 `json:"id"`
}

type RoleAllocResourceRequest struct {
	RoleId      string `json:"roleId"`
	ResourceIds string `json:"resourceIds"`
}
