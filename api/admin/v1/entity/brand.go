package entity

import "go-mall-api/pkg/paginator"

type BrandListRequest struct {
	PageNum  int    `json:"pageNum"`
	PageSize int    `json:"pageSize"`
	Keyword  string `json:"keyword"`
}

type BrandList struct {
	Id                  uint64 `json:"id"`
	Sort                int    `json:"sort"`
	ShowStatus          int    `json:"showStatus"`
	ProductCount        int    `json:"productCount"`
	ProductCommentCount int    `json:"productCommentCount"`
	FactoryStatus       int    `json:"factoryStatus"`
	Name                string `json:"name"`
	FirstLetter         string `json:"firstLetter"`
	Logo                string `json:"logo"`
	BigPic              string `json:"bigPic"`
	BrandStory          string `json:"brandStory"`
}

type BrandListReply struct {
	List []BrandList `json:"list"`
	paginator.PagingAdmin
}

type BrandRequest struct {
	Id string `json:"id"`
}

type BrandCreateRequest struct {
	Sort          int    `json:"sort"`
	FactoryStatus int    `json:"factoryStatus"`
	ShowStatus    int    `json:"showStatus"`
	Name          string `json:"name"`
	BrandStory    string `json:"brandStory"`
	FirstLetter   string `json:"firstLetter"`
	Logo          string `json:"logo"`
	BigPic        string `json:"bigPic"`
}

type BrandUpdateRequest struct {
	Id            int    `json:"id"`
	Sort          int    `json:"sort"`
	FactoryStatus int    `json:"factoryStatus"`
	ShowStatus    int    `json:"showStatus"`
	Name          string `json:"name"`
	BrandStory    string `json:"brandStory"`
	FirstLetter   string `json:"firstLetter"`
	Logo          string `json:"logo"`
	BigPic        string `json:"bigPic"`
}

type BrandDeleteRequest struct {
	Id string `json:"id"`
}

type BrandUpdateFactoryStatusRequest struct {
	Ids           string `json:"ids"`
	FactoryStatus string `json:"factoryStatus"`
}

type BrandUpdateShowStatusRequest struct {
	Ids        string `json:"ids"`
	ShowStatus string `json:"showStatus"`
}
