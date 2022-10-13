//Package sms_coupon 模型
package sms_coupon

import (
	"go-mall-api/models"
	"go-mall-api/pkg/database"
	"time"
)

//SmsCoupon 优惠券表
type SmsCoupon struct {
	models.BaseModel
	Type         int       `gorm:"column:type;type:int(1);comment:优惠券类型；0->全场赠券；1->会员赠券；2->购物赠券；3->注册赠券" json:"type"`
	Name         string    `gorm:"column:name;type:varchar(100)" json:"name"`
	Platform     int       `gorm:"column:platform;type:int(1);comment:使用平台：0->全部；1->移动；2->PC" json:"platform"`
	Count        int       `gorm:"column:count;type:int(11);comment:数量" json:"count"`
	Amount       float64   `gorm:"column:amount;type:decimal(10,2);comment:金额" json:"amount"`
	PerLimit     int       `gorm:"column:per_limit;type:int(11);comment:每人限领张数" json:"per_limit"`
	MinPoint     float64   `gorm:"column:min_point;type:decimal(10,2);comment:使用门槛；0表示无门槛" json:"min_point"`
	StartTime    time.Time `gorm:"column:start_time;type:datetime" json:"start_time"`
	EndTime      time.Time `gorm:"column:end_time;type:datetime" json:"end_time"`
	UseType      int       `gorm:"column:use_type;type:int(1);comment:使用类型：0->全场通用；1->指定分类；2->指定商品" json:"use_type"`
	Note         string    `gorm:"column:note;type:varchar(200);comment:备注" json:"note"`
	PublishCount int       `gorm:"column:publish_count;type:int(11);comment:发行数量" json:"publish_count"`
	UseCount     int       `gorm:"column:use_count;type:int(11);comment:已使用数量" json:"use_count"`
	ReceiveCount int       `gorm:"column:receive_count;type:int(11);comment:领取数量" json:"receive_count"`
	EnableTime   time.Time `gorm:"column:enable_time;type:datetime;comment:可以领取的日期" json:"enable_time"`
	Code         string    `gorm:"column:code;type:varchar(64);comment:优惠码" json:"code"`
	MemberLevel  int       `gorm:"column:member_level;type:int(1);comment:可领取的会员类型：0->无限时" json:"member_level"`
}

func (smsCoupon *SmsCoupon) TableName() string {
	return "sms_coupon"
}

func (smsCoupon *SmsCoupon) Create() {
	database.DB.Create(&smsCoupon)
}

func (smsCoupon *SmsCoupon) Save() (rowsAffected int64) {
	result := database.DB.Save(&smsCoupon)
	return result.RowsAffected
}

func (smsCoupon *SmsCoupon) Updates(data map[string]interface{}) (rowsAffected int64) {
	result := database.DB.Model(&smsCoupon).Updates(data)
	return result.RowsAffected
}

func (smsCoupon *SmsCoupon) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&smsCoupon)
	return result.RowsAffected
}
