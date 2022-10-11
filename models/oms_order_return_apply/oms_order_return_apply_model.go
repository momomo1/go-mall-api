//Package oms_order_return_apply 模型
package oms_order_return_apply

import (
	"go-mall-api/models"
	"go-mall-api/pkg/database"
	"time"
)

//OmsOrderReturnApply 订单退货申请
type OmsOrderReturnApply struct {
	models.BaseModel
	OrderId          int64     `gorm:"column:order_id;type:bigint(20);comment:订单id" json:"order_id"`
	CompanyAddressId int64     `gorm:"column:company_address_id;type:bigint(20);comment:收货地址表id" json:"company_address_id"`
	ProductId        int64     `gorm:"column:product_id;type:bigint(20);comment:退货商品id" json:"product_id"`
	OrderSn          string    `gorm:"column:order_sn;type:varchar(64);comment:订单编号" json:"order_sn"`
	CreateTime       time.Time `gorm:"column:create_time;type:datetime;comment:申请时间" json:"create_time"`
	MemberUsername   string    `gorm:"column:member_username;type:varchar(64);comment:会员用户名" json:"member_username"`
	ReturnAmount     float64   `gorm:"column:return_amount;type:decimal(10,2);comment:退款金额" json:"return_amount"`
	ReturnName       string    `gorm:"column:return_name;type:varchar(100);comment:退货人姓名" json:"return_name"`
	ReturnPhone      string    `gorm:"column:return_phone;type:varchar(100);comment:退货人电话" json:"return_phone"`
	Status           int       `gorm:"column:status;type:int(1);comment:申请状态：0->待处理；1->退货中；2->已完成；3->已拒绝" json:"status"`
	HandleTime       time.Time `gorm:"column:handle_time;type:datetime;comment:处理时间" json:"handle_time"`
	ProductPic       string    `gorm:"column:product_pic;type:varchar(500);comment:商品图片" json:"product_pic"`
	ProductName      string    `gorm:"column:product_name;type:varchar(200);comment:商品名称" json:"product_name"`
	ProductBrand     string    `gorm:"column:product_brand;type:varchar(200);comment:商品品牌" json:"product_brand"`
	ProductAttr      string    `gorm:"column:product_attr;type:varchar(500);comment:商品销售属性：颜色：红色；尺码：xl;" json:"product_attr"`
	ProductCount     int       `gorm:"column:product_count;type:int(11);comment:退货数量" json:"product_count"`
	ProductPrice     float64   `gorm:"column:product_price;type:decimal(10,2);comment:商品单价" json:"product_price"`
	ProductRealPrice float64   `gorm:"column:product_real_price;type:decimal(10,2);comment:商品实际支付单价" json:"product_real_price"`
	Reason           string    `gorm:"column:reason;type:varchar(200);comment:原因" json:"reason"`
	Description      string    `gorm:"column:description;type:varchar(500);comment:描述" json:"description"`
	ProofPics        string    `gorm:"column:proof_pics;type:varchar(1000);comment:凭证图片，以逗号隔开" json:"proof_pics"`
	HandleNote       string    `gorm:"column:handle_note;type:varchar(500);comment:处理备注" json:"handle_note"`
	HandleMan        string    `gorm:"column:handle_man;type:varchar(100);comment:处理人员" json:"handle_man"`
	ReceiveMan       string    `gorm:"column:receive_man;type:varchar(100);comment:收货人" json:"receive_man"`
	ReceiveTime      time.Time `gorm:"column:receive_time;type:datetime;comment:收货时间" json:"receive_time"`
	ReceiveNote      string    `gorm:"column:receive_note;type:varchar(500);comment:收货备注" json:"receive_note"`
}

func (omsOrderReturnApply *OmsOrderReturnApply) TableName() string {
	return "oms_order_return_apply"
}

func (omsOrderReturnApply *OmsOrderReturnApply) Create() {
	database.DB.Create(&omsOrderReturnApply)
}

func (omsOrderReturnApply *OmsOrderReturnApply) Save() (rowsAffected int64) {
	result := database.DB.Save(&omsOrderReturnApply)
	return result.RowsAffected
}

func (omsOrderReturnApply *OmsOrderReturnApply) Updates(data map[string]interface{}) (rowsAffected int64) {
	result := database.DB.Model(&omsOrderReturnApply).Updates(data)
	return result.RowsAffected
}

func (omsOrderReturnApply *OmsOrderReturnApply) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&omsOrderReturnApply)
	return result.RowsAffected
}
