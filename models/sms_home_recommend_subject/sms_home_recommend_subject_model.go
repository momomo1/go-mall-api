//Package sms_home_recommend_subject 模型
package sms_home_recommend_subject

import (
	"go-mall-api/models"
	"go-mall-api/pkg/database"
)

//SmsHomeRecommendSubject 首页推荐专题表
type SmsHomeRecommendSubject struct {
	models.BaseModel
	SubjectId       int64  `gorm:"column:subject_id;type:bigint(20)" json:"subject_id"`
	SubjectName     string `gorm:"column:subject_name;type:varchar(64)" json:"subject_name"`
	RecommendStatus int    `gorm:"column:recommend_status;type:int(1)" json:"recommend_status"`
	Sort            int    `gorm:"column:sort;type:int(11)" json:"sort"`
}

func (smsHomeRecommendSubject *SmsHomeRecommendSubject) TableName() string {
	return "sms_home_recommend_subject"
}

func (smsHomeRecommendSubject *SmsHomeRecommendSubject) Create() {
	database.DB.Create(&smsHomeRecommendSubject)
}

func (smsHomeRecommendSubject *SmsHomeRecommendSubject) Save() (rowsAffected int64) {
	result := database.DB.Save(&smsHomeRecommendSubject)
	return result.RowsAffected
}

func (smsHomeRecommendSubject *SmsHomeRecommendSubject) Updates(data map[string]interface{}) (rowsAffected int64) {
	result := database.DB.Model(&smsHomeRecommendSubject).Updates(data)
	return result.RowsAffected
}

func (smsHomeRecommendSubject *SmsHomeRecommendSubject) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&smsHomeRecommendSubject)
	return result.RowsAffected
}
