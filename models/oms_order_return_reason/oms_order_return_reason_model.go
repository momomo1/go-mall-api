//Package oms_order_return_reason 模型
package oms_order_return_reason

import (
	"go-mall-api/models"
	"go-mall-api/pkg/database"
	"time"
)

//OmsOrderReturnReason 退货原因表
type OmsOrderReturnReason struct {
	models.BaseModel
	Name       string    `gorm:"column:name;type:varchar(100);comment:退货类型" json:"name"`
	Sort       int       `gorm:"column:sort;type:int(11)" json:"sort"`
	Status     int       `gorm:"column:status;type:int(1);comment:状态：0->不启用；1->启用" json:"status"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;comment:添加时间" json:"create_time"`
}

func (omsOrderReturnReason *OmsOrderReturnReason) TableName() string {
	return "oms_order_return_reason"
}

func (omsOrderReturnReason *OmsOrderReturnReason) Create() {
	database.DB.Create(&omsOrderReturnReason)
}

func (omsOrderReturnReason *OmsOrderReturnReason) Save() (rowsAffected int64) {
	result := database.DB.Save(&omsOrderReturnReason)
	return result.RowsAffected
}

func (omsOrderReturnReason *OmsOrderReturnReason) Updates(data map[string]interface{}) (rowsAffected int64) {
	result := database.DB.Model(&omsOrderReturnReason).Updates(data)
	return result.RowsAffected
}

func (omsOrderReturnReason *OmsOrderReturnReason) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&omsOrderReturnReason)
	return result.RowsAffected
}
