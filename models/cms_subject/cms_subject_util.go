package cms_subject

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/pkg/database"
	"go-mall-api/pkg/paginator"
)

func Get(idstr string) (cmsSubject CmsSubject) {
	database.DB.Where("id", idstr).First(&cmsSubject)
	return
}

func GetBy(field, value string) (cmsSubject CmsSubject) {
	database.DB.Where("? = ?", field, value).First(&cmsSubject)
	return
}

func All() (cmsSubjects []CmsSubject) {
	database.DB.Find(&cmsSubjects)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(CmsSubject{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int, where interface{}, sort string, order string) (subject []CmsSubject, paging paginator.PagingAdmin) {
	paging = paginator.PaginateAdmin(
		c,
		database.DB.Model(CmsSubject{}),
		&subject,
		where,
		perPage,
		sort,
		order,
	)
	return
}
