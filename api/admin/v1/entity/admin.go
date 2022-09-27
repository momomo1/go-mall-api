package entity

import (
	"go-mall-api/models/ums_role"
	"go-mall-api/pkg/paginator"
)

type ListRequest struct {
	PageNum  int    `json:"pageNum"`
	PageSize int    `json:"pageSize"`
	Keyword  string `json:"keyword"`
}

type ListReply struct {
	List []AdminList `json:"list"`
	paginator.PagingAdmin
}

type AdminList struct {
	Id         uint64 `json:"id"`
	Status     *int   `json:"status"`
	UserName   string `json:"username"`
	NickName   string `json:"nickName"`
	Email      string `json:"email"`
	CreateTime string `json:"createTime"`
	LoginTime  string `json:"loginTime"`
	Password   string `json:"password"`
	Icon       string `json:"icon"`
	Note       string `json:"note"`
}

type AdminUpdateStatusRequest struct {
	Id     string `json:"id"`
	Status int    `json:"status"`
}

type AdminRegisterRequest struct {
	Username string `json:"username" valid:"username"`
	NickName string `json:"nickName" valid:"nickName"`
	Password string `json:"password" valid:"password"`
	Email    string `json:"email" valid:"email"`
	Note     string `json:"note" valid:"note"`
	Status   int    `json:"status" valid:"status"`
}

type AdminRoleUpdateRequest struct {
	AdminId string `json:"adminId" valid:"adminId"`
	RoleIds string `json:"roleIds" valid:"roleIds"`
}

type AdminRolesRequest struct {
	Id string `json:"id"`
}

type AdminRolesReply struct {
	Data []ums_role.UmsRole `json:"data"`
}

type AdminUpdateRequest struct {
	Id       int    `json:"id" valid:"id"`
	Username string `json:"username" valid:"username"`
	NickName string `json:"nickName" valid:"nickName"`
	Password string `json:"password" valid:"password"`
	Email    string `json:"email" valid:"email"`
	Note     string `json:"note" valid:"note"`
	Status   int    `json:"status" valid:"status"`
}

type AdminDeleteRequest struct {
	Id string `json:"id"`
}
