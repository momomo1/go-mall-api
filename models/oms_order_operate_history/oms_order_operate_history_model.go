//Package oms_order_operate_history 模型
package oms_order_operate_history

import (
	"go-mall-api/models"
	"go-mall-api/pkg/database"
	"time"
)

//OmsOrderOperateHistory 订单操作历史记录
type OmsOrderOperateHistory struct {
	models.BaseModel

	OrderId     int64     `gorm:"column:order_id;type:bigint(20);comment:订单id" json:"order_id"`
	OperateMan  string    `gorm:"column:operate_man;type:varchar(100);comment:操作人：用户；系统；后台管理员" json:"operate_man"`
	CreateTime  time.Time `gorm:"column:create_time;type:datetime;comment:操作时间" json:"create_time"`
	OrderStatus int       `gorm:"column:order_status;type:int(1);comment:订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单" json:"order_status"`
	Note        string    `gorm:"column:note;type:varchar(500);comment:备注" json:"note"`
}

func (omsOrderOperateHistory *OmsOrderOperateHistory) TableName() string {
	return "oms_order_operate_history"
}

func (omsOrderOperateHistory *OmsOrderOperateHistory) Create() {
	database.DB.Create(&omsOrderOperateHistory)
}

func (omsOrderOperateHistory *OmsOrderOperateHistory) Save() (rowsAffected int64) {
	result := database.DB.Save(&omsOrderOperateHistory)
	return result.RowsAffected
}

func (omsOrderOperateHistory *OmsOrderOperateHistory) Updates(data map[string]interface{}) (rowsAffected int64) {
	result := database.DB.Model(&omsOrderOperateHistory).Updates(data)
	return result.RowsAffected
}

func (omsOrderOperateHistory *OmsOrderOperateHistory) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&omsOrderOperateHistory)
	return result.RowsAffected
}
