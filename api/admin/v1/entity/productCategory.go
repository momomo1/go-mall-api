package entity

import "go-mall-api/pkg/paginator"

type ProductCategoryListRequest struct {
	Id       string `json:"id"`
	PageNum  int    `json:"pageNum"`
	PageSize int    `json:"pageSize"`
}

type ProductCategoryListReply struct {
	List []ProductCategoryListList `json:"list"`
	paginator.PagingAdmin
}

type ProductCategoryListList struct {
	Id           uint64 `json:"id"`
	Level        int    `json:"level"`
	ParentId     int64  `json:"parentId"`
	ProductCount int    `json:"productCount"`
	NavStatus    int    `json:"navStatus"`
	ShowStatus   int    `json:"showStatus"`
	Sort         int    `json:"sort"`
	Keywords     string `json:"keywords"`
	Name         string `json:"name"`
	ProductUnit  string `json:"productUnit"`
	Description  string `json:"description"`
	Icon         string `json:"icon"`
}
