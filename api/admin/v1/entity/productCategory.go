package entity

import "go-mall-api/pkg/paginator"

type ProductCategoryRequest struct {
	Id string `json:"id"`
}

type ProductCategoryListRequest struct {
	Id       string `json:"id"`
	PageNum  int    `json:"pageNum"`
	PageSize int    `json:"pageSize"`
}

type ProductCategoryListReply struct {
	List []ProductCategoryList `json:"list"`
	paginator.PagingAdmin
}

type ProductCategoryList struct {
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

type ProductCategoryListWithChildrenReply struct {
	Id       uint64                                  `json:"id"`
	Name     string                                  `json:"name"`
	Children []*ProductCategoryListWithChildrenReply `json:"children"`
}

type ProductCategoryUpdateNavStatusRequest struct {
	Ids       string `json:"ids"`
	NavStatus string `json:"navStatus"`
}

type ProductCategoryUpdateShowStatusRequest struct {
	Ids        string `json:"ids"`
	ShowStatus string `json:"showStatus"`
}

type ProductCategoryCreateRequest struct {
	NavStatus              int     `json:"navStatus"`
	ParentId               int64   `json:"parentId"`
	ShowStatus             int     `json:"showStatus"`
	Sort                   int     `json:"sort"`
	Description            string  `json:"description"`
	Icon                   string  `json:"icon"`
	Keywords               string  `json:"keywords"`
	Name                   string  `json:"name"`
	ProductUnit            string  `json:"productUnit"`
	ProductAttributeIdList []int64 `json:"productAttributeIdList"`
}

type ProductCategoryUpdateRequest struct {
	Id                     int     `json:"id"`
	NavStatus              int     `json:"navStatus"`
	ParentId               int64   `json:"parentId"`
	ShowStatus             int     `json:"showStatus"`
	Sort                   int     `json:"sort"`
	Description            string  `json:"description"`
	Icon                   string  `json:"icon"`
	Keywords               string  `json:"keywords"`
	Name                   string  `json:"name"`
	ProductUnit            string  `json:"productUnit"`
	ProductAttributeIdList []int64 `json:"productAttributeIdList"`
}

type ProductCategoryDeleteRequest struct {
	Id string `json:"id"`
}
