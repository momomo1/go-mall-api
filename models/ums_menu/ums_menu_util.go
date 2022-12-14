package ums_menu

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/pkg/database"
	"go-mall-api/pkg/paginator"
)

func Get(idstr string) (umsMenu UmsMenu) {
	database.DB.Where("id", idstr).First(&umsMenu)
	return
}

func GetBy(field, value string) (umsMenu UmsMenu) {
	database.DB.Where("? = ?", field, value).First(&umsMenu)
	return
}

func All() (umsMenus []UmsMenu, count int64) {
	database.DB.Find(&umsMenus)
	database.DB.Table("ums_menu").Count(&count)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(UmsMenu{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int, where interface{}, sort string, order string) (menus []UmsMenu, paging paginator.PagingAdmin) {
	paging = paginator.PaginateAdmin(
		c,
		database.DB.Model(UmsMenu{}),
		&menus,
		where,
		perPage,
		sort,
		order,
	)
	return
}
