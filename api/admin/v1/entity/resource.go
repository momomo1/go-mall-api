package entity

import (
	"go-mall-api/pkg/paginator"
)

type ResourceListRequest struct {
	PageNum     int    `json:"pageNum"`
	PageSize    int    `json:"pageSize"`
	CategoryId  int    `json:"categoryId"`
	NameKeyword string `json:"nameKeyword"`
	UrlKeyword  string `json:"urlKeyword"`
}

type ResourceList struct {
	Id          int64  `json:"id"`
	CategoryId  int64  `json:"categoryId"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	Description string `json:"description"`
	CreateTime  string `json:"createTime"`
}

type ResourceListReply struct {
	List []ResourceList `json:"list"`
	paginator.PagingAdmin
}

type ResourceCategoryListAllReply struct {
	Id         int64  `json:"id"`
	Sort       int    `json:"sort"`
	Name       string `json:"name"`
	CreateTime string `json:"createTime"`
}

type ResourceCategoryCreateRequest struct {
	Name string `json:"name,omitempty" valid:"name"`
	Sort string `json:"sort,omitempty" valid:"sort"`
}

type ResourceCategoryUpdateRequest struct {
	Id   int    `json:"id,omitempty" valid:"id"`
	Name string `json:"name,omitempty" valid:"name"`
	Sort string `json:"sort,omitempty" valid:"sort"`
}

type ResourceCategoryDeleteRequest struct {
	Id string `json:"id,omitempty" valid:"id"`
}
