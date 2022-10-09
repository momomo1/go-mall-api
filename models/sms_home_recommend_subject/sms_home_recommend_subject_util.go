package sms_home_recommend_subject

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/pkg/database"
	"go-mall-api/pkg/paginator"
)

func Get(idstr string) (smsHomeRecommendSubject SmsHomeRecommendSubject) {
	database.DB.Where("id", idstr).First(&smsHomeRecommendSubject)
	return
}

func GetBySubjectId(value string) (smsHomeRecommendSubject SmsHomeRecommendSubject) {
	database.DB.Where("subject_id", value).First(&smsHomeRecommendSubject)
	return
}

func GetBy(field, value string) (smsHomeRecommendSubject SmsHomeRecommendSubject) {
	database.DB.Where("? = ?", field, value).First(&smsHomeRecommendSubject)
	return
}

func All() (smsHomeRecommendSubjects []SmsHomeRecommendSubject) {
	database.DB.Find(&smsHomeRecommendSubjects)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(SmsHomeRecommendSubject{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int, where interface{}, sort string, order string) (homeRecommendSubject []SmsHomeRecommendSubject, paging paginator.PagingAdmin) {
	paging = paginator.PaginateAdmin(
		c,
		database.DB.Model(SmsHomeRecommendSubject{}),
		&homeRecommendSubject,
		where,
		perPage,
		sort,
		order,
	)
	return
}
