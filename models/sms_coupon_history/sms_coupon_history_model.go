//Package sms_coupon_history 模型
package sms_coupon_history

import (
	"go-mall-api/models"
	"go-mall-api/pkg/database"
	"time"
)

//SmsCouponHistory 优惠券使用、领取历史表
type SmsCouponHistory struct {
	models.BaseModel
	CouponId       int64     `gorm:"column:coupon_id;type:bigint(20)" json:"coupon_id"`
	MemberId       int64     `gorm:"column:member_id;type:bigint(20)" json:"member_id"`
	CouponCode     string    `gorm:"column:coupon_code;type:varchar(64)" json:"coupon_code"`
	MemberNickname string    `gorm:"column:member_nickname;type:varchar(64);comment:领取人昵称" json:"member_nickname"`
	GetType        int       `gorm:"column:get_type;type:int(1);comment:获取类型：0->后台赠送；1->主动获取" json:"get_type"`
	CreateTime     time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`
	UseStatus      int       `gorm:"column:use_status;type:int(1);comment:使用状态：0->未使用；1->已使用；2->已过期" json:"use_status"`
	UseTime        time.Time `gorm:"column:use_time;type:datetime;comment:使用时间" json:"use_time"`
	OrderId        int64     `gorm:"column:order_id;type:bigint(20);comment:订单编号" json:"order_id"`
	OrderSn        string    `gorm:"column:order_sn;type:varchar(100);comment:订单号码" json:"order_sn"`
}

func (smsCouponHistory *SmsCouponHistory) TableName() string {
	return "sms_coupon_history"
}

func (smsCouponHistory *SmsCouponHistory) Create() {
	database.DB.Create(&smsCouponHistory)
}

func (smsCouponHistory *SmsCouponHistory) Save() (rowsAffected int64) {
	result := database.DB.Save(&smsCouponHistory)
	return result.RowsAffected
}

func (smsCouponHistory *SmsCouponHistory) Updates(data map[string]interface{}) (rowsAffected int64) {
	result := database.DB.Model(&smsCouponHistory).Updates(data)
	return result.RowsAffected
}

func (smsCouponHistory *SmsCouponHistory) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&smsCouponHistory)
	return result.RowsAffected
}
