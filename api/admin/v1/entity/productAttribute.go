package entity

import (
	"go-mall-api/models/pms_product_attribute"
	"go-mall-api/pkg/paginator"
)

type ProductAttributeRequest struct {
	Id string `json:"id"`
}

type ProductAttributeAttrInfoRequest struct {
	Id string `json:"id"`
}

type ProductAttributeAttrInfoReply struct {
	AttributeCategoryId int64 `json:"attributeCategoryId"`
	AttributeId         int64 `json:"attributeId"`
}

type ProductAttributeListRequest struct {
	Id       string `json:"id"`
	PageNum  int    `json:"pageNum"`
	PageSize int    `json:"pageSize"`
	Type     int    `json:"type"`
}

type ProductAttributeListReply struct {
	List []ProductAttributeList `json:"list"`
	paginator.PagingAdmin
}

type ProductAttributeList struct {
	Id                         uint64 `json:"id"`
	FilterType                 int    `json:"filterType"`
	HandAddStatus              int    `json:"handAddStatus"`
	InputType                  int    `json:"inputType"`
	ProductAttributeCategoryId int64  `json:"productAttributeCategoryId"`
	RelatedStatus              int    `json:"relatedStatus"`
	SearchType                 int    `json:"searchType"`
	SelectType                 int    `json:"selectType"`
	Sort                       int    `json:"sort"`
	Type                       int    `json:"type"`
	InputList                  string `json:"inputList"`
	Name                       string `json:"name"`
}

type ProductAttributeCreateRequest struct {
	FilterType                 int    `json:"filterType"`
	HandAddStatus              int    `json:"handAddStatus"`
	InputType                  int    `json:"inputType"`
	ProductAttributeCategoryId int64  `json:"productAttributeCategoryId"`
	RelatedStatus              int    `json:"relatedStatus"`
	SearchType                 int    `json:"searchType"`
	SelectType                 int    `json:"selectType"`
	Sort                       int    `json:"sort"`
	Type                       int    `json:"type"`
	InputList                  string `json:"inputList"`
	Name                       string `json:"name"`
}

type ProductAttributeUpdateRequest struct {
	Id                         int    `json:"id"`
	FilterType                 int    `json:"filterType"`
	HandAddStatus              int    `json:"handAddStatus"`
	InputType                  int    `json:"inputType"`
	ProductAttributeCategoryId int64  `json:"productAttributeCategoryId"`
	RelatedStatus              int    `json:"relatedStatus"`
	SearchType                 int    `json:"searchType"`
	SelectType                 int    `json:"selectType"`
	Sort                       int    `json:"sort"`
	Type                       int    `json:"type"`
	InputList                  string `json:"inputList"`
	Name                       string `json:"name"`
}

type ProductAttributeDeleteRequest struct {
	Ids string `json:"ids"`
}

type ProductAttributeCategoryListRequest struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}

type ProductAttributeCategoryListReply struct {
	List []ProductAttributeCategoryList `json:"list"`
	paginator.PagingAdmin
}

type ProductAttributeCategoryList struct {
	Id             uint64 `json:"id"`
	ParamCount     int    `json:"paramCount"`
	AttributeCount int    `json:"attributeCount"`
	Name           string `json:"name"`
}

type ProductAttributeCategoryCreateRequest struct {
	Name string `json:"name"`
}

type ProductAttributeCategoryUpdateRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ProductAttributeCategoryDeleteRequest struct {
	Id string `json:"id"`
}

type ProductAttributeCategoryListWithAttrReply struct {
	Id                   uint64                                      `json:"id"`
	AttributeCount       int                                         `json:"attributeCount"`
	ParamCount           int                                         `json:"paramCount"`
	Name                 string                                      `json:"name"`
	ProductAttributeList []pms_product_attribute.PmsProductAttribute `json:"productAttributeList"`
}
