//Package sms_home_advertise 模型
package sms_home_advertise

import (
	"go-mall-api/models"
	"go-mall-api/pkg/database"
	"time"
)

//SmsHomeAdvertise 首页轮播广告表
type SmsHomeAdvertise struct {
	models.BaseModel

	Name       string    `gorm:"column:name;type:varchar(100)" json:"name"`
	Type       int       `gorm:"column:type;type:int(1);comment:轮播位置：0->PC首页轮播；1->app首页轮播" json:"type"`
	Pic        string    `gorm:"column:pic;type:varchar(500)" json:"pic"`
	StartTime  time.Time `gorm:"column:start_time;type:datetime" json:"start_time"`
	EndTime    time.Time `gorm:"column:end_time;type:datetime" json:"end_time"`
	Status     int       `gorm:"column:status;type:int(1);comment:上下线状态：0->下线；1->上线" json:"status"`
	ClickCount int       `gorm:"column:click_count;type:int(11);comment:点击数" json:"click_count"`
	OrderCount int       `gorm:"column:order_count;type:int(11);comment:下单数" json:"order_count"`
	Url        string    `gorm:"column:url;type:varchar(500);comment:链接地址" json:"url"`
	Note       string    `gorm:"column:note;type:varchar(500);comment:备注" json:"note"`
	Sort       int       `gorm:"column:sort;type:int(11);default:0;comment:排序" json:"sort"`
}

func (smsHomeAdvertise *SmsHomeAdvertise) TableName() string {
	return "sms_home_advertise"
}

func (smsHomeAdvertise *SmsHomeAdvertise) Create() {
	database.DB.Create(&smsHomeAdvertise)
}

func (smsHomeAdvertise *SmsHomeAdvertise) Save() (rowsAffected int64) {
	result := database.DB.Save(&smsHomeAdvertise)
	return result.RowsAffected
}

func (smsHomeAdvertise *SmsHomeAdvertise) Updates(data map[string]interface{}) (rowsAffected int64) {
	result := database.DB.Model(&smsHomeAdvertise).Updates(data)
	return result.RowsAffected
}

func (smsHomeAdvertise *SmsHomeAdvertise) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&smsHomeAdvertise)
	return result.RowsAffected
}