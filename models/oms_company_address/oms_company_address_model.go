//Package oms_company_address 模型
package oms_company_address

import (
	"go-mall-api/models"
	"go-mall-api/pkg/database"
)

//OmsCompanyAddress 公司收发货地址表
type OmsCompanyAddress struct {
	models.BaseModel

	AddressName   string `gorm:"column:address_name;type:varchar(200);comment:地址名称" json:"address_name"`
	SendStatus    int    `gorm:"column:send_status;type:int(1);comment:默认发货地址：0->否；1->是" json:"send_status"`
	ReceiveStatus int    `gorm:"column:receive_status;type:int(1);comment:是否默认收货地址：0->否；1->是" json:"receive_status"`
	Name          string `gorm:"column:name;type:varchar(64);comment:收发货人姓名" json:"name"`
	Phone         string `gorm:"column:phone;type:varchar(64);comment:收货人电话" json:"phone"`
	Province      string `gorm:"column:province;type:varchar(64);comment:省/直辖市" json:"province"`
	City          string `gorm:"column:city;type:varchar(64);comment:市" json:"city"`
	Region        string `gorm:"column:region;type:varchar(64);comment:区" json:"region"`
	DetailAddress string `gorm:"column:detail_address;type:varchar(200);comment:详细地址" json:"detail_address"`
}

func (omsCompanyAddress *OmsCompanyAddress) TableName() string {
	return "oms_company_address"
}

func (omsCompanyAddress *OmsCompanyAddress) Create() {
	database.DB.Create(&omsCompanyAddress)
}

func (omsCompanyAddress *OmsCompanyAddress) Save() (rowsAffected int64) {
	result := database.DB.Save(&omsCompanyAddress)
	return result.RowsAffected
}

func (omsCompanyAddress *OmsCompanyAddress) Updates(data map[string]interface{}) (rowsAffected int64) {
	result := database.DB.Model(&omsCompanyAddress).Updates(data)
	return result.RowsAffected
}

func (omsCompanyAddress *OmsCompanyAddress) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&omsCompanyAddress)
	return result.RowsAffected
}
