package entity

import (
	"go-mall-api/models/ums_menu"
	"go-mall-api/pkg/paginator"
	"time"
)

type MenuTreeListReply struct {
	Id         uint64               `json:"id"`
	Level      int                  `json:"level"`
	ParentId   int64                `json:"parentId"`
	Hidden     int                  `json:"hidden"`
	Sort       int                  `json:"sort"`
	Name       string               `json:"name"`
	Title      string               `json:"title"`
	Icon       string               `json:"icon"`
	CreateTime time.Time            `json:"createTime"`
	Children   []*MenuTreeListReply `json:"children"`
}

type MenuListRequest struct {
	Id       string `json:"id"`
	PageNum  int    `json:"pageNum"`
	PageSize int    `json:"pageSize"`
}

type MenuListReply struct {
	List []ums_menu.UmsMenu `json:"list"`
	paginator.PagingAdmin
}

type MenuUpdateHiddenRequest struct {
	Id     string `json:"id"`
	Hidden int    `json:"hidden"`
}

type MenuCreateRequest struct {
	Title    string `json:"title"`
	Name     string `json:"name"`
	Icon     string `json:"icon"`
	ParentId int64  `json:"parentId"`
	Hidden   int    `json:"hidden"`
	Sort     int    `json:"sort"`
}

type MenuRequest struct {
	Id string `json:"id"`
}

type MenuReply struct {
	Data ums_menu.UmsMenu `json:"data"`
}

type MenuUpdateRequest struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	ParentId int64  `json:"parentId"`
	Name     string `json:"name"`
	Icon     string `json:"icon"`
	Hidden   int    `json:"hidden"`
	Sort     string `json:"sort"`
}

type MenuDeleteRequest struct {
	Id string `json:"id"`
}
