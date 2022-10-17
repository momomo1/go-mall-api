package oms_order_operate_history

import (
	"go-mall-api/pkg/database"
	"time"
)

func Get(idstr string) (omsOrderOperateHistory OmsOrderOperateHistory) {
	database.DB.Where("id", idstr).First(&omsOrderOperateHistory)
	return
}

func GetBy(field, value string) (omsOrderOperateHistory OmsOrderOperateHistory) {
	database.DB.Where("? = ?", field, value).First(&omsOrderOperateHistory)
	return
}

func All() (omsOrderOperateHistories []OmsOrderOperateHistory) {
	database.DB.Find(&omsOrderOperateHistories)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(OmsOrderOperateHistory{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

func AddOperateHistory(orderId int64, orderStatus int, userName string, note string) {
	history := OmsOrderOperateHistory{
		OrderId:     orderId,
		OperateMan:  userName,
		CreateTime:  time.Now(),
		OrderStatus: orderStatus,
		Note:        note,
	}
	history.Create()
}
