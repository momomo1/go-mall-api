//Package oms_order_setting 模型
package oms_order_setting

import (
	"go-mall-api/models"
	"go-mall-api/pkg/database"
)

//OmsOrderSetting 订单设置表
type OmsOrderSetting struct {
	models.BaseModel

	FlashOrderOvertime  int `gorm:"column:flash_order_overtime;type:int(11);comment:秒杀订单超时关闭时间(分)" json:"flash_order_overtime"`
	NormalOrderOvertime int `gorm:"column:normal_order_overtime;type:int(11);comment:正常订单超时时间(分)" json:"normal_order_overtime"`
	ConfirmOvertime     int `gorm:"column:confirm_overtime;type:int(11);comment:发货后自动确认收货时间（天）" json:"confirm_overtime"`
	FinishOvertime      int `gorm:"column:finish_overtime;type:int(11);comment:自动完成交易时间，不能申请售后（天）" json:"finish_overtime"`
	CommentOvertime     int `gorm:"column:comment_overtime;type:int(11);comment:订单完成后自动好评时间（天）" json:"comment_overtime"`
}

func (omsOrderSetting *OmsOrderSetting) TableName() string {
	return "oms_order_setting"
}

func (omsOrderSetting *OmsOrderSetting) Create() {
	database.DB.Create(&omsOrderSetting)
}

func (omsOrderSetting *OmsOrderSetting) Save() (rowsAffected int64) {
	result := database.DB.Save(&omsOrderSetting)
	return result.RowsAffected
}

func (omsOrderSetting *OmsOrderSetting) Updates(data map[string]interface{}) (rowsAffected int64) {
	result := database.DB.Model(&omsOrderSetting).Updates(data)
	return result.RowsAffected
}

func (omsOrderSetting *OmsOrderSetting) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&omsOrderSetting)
	return result.RowsAffected
}
