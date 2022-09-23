package ums_resource

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/pkg/database"
	"go-mall-api/pkg/paginator"
)

func Get(idstr string) (umsResource UmsResource) {
	database.DB.Where("id", idstr).First(&umsResource)
	return
}

func GetBy(field, value string) (umsResource UmsResource) {
	database.DB.Where("? = ?", field, value).First(&umsResource)
	return
}

func All() (umsResources []UmsResource) {
	database.DB.Find(&umsResources)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(UmsResource{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int, where interface{}) (menus []UmsResource, paging paginator.PagingAdmin) {
	paging = paginator.PaginateAdmin(
		c,
		database.DB.Model(UmsResource{}),
		&menus,
		where,
		perPage,
	)
	return
}
