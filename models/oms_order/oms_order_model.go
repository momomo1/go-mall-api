//Package oms_order 模型
package oms_order

import (
	"go-mall-api/models"
	"go-mall-api/models/oms_order_item"
	"go-mall-api/models/oms_order_operate_history"
	"go-mall-api/pkg/database"
	"time"
)

const Obligation = 0    //待付款
const SendGoods = 1     //待发货
const Shipped = 2       //已发货
const Completed = 3     //已完成
const Closed = 4        //已关闭
const InvalidOrders = 5 //无效订单

//OmsOrder 订单表
type OmsOrder struct {
	models.BaseModel

	MemberId              int64                                              `gorm:"column:member_id;type:bigint(20);NOT NULL" json:"member_id"`
	CouponId              int64                                              `gorm:"column:coupon_id;type:bigint(20)" json:"coupon_id"`
	OrderSn               string                                             `gorm:"column:order_sn;type:varchar(64);comment:订单编号" json:"order_sn"`
	CreateTime            time.Time                                          `gorm:"column:create_time;type:datetime;comment:提交时间" json:"create_time"`
	MemberUsername        string                                             `gorm:"column:member_username;type:varchar(64);comment:用户帐号" json:"member_username"`
	TotalAmount           float64                                            `gorm:"column:total_amount;type:decimal(10,2);comment:订单总金额" json:"total_amount"`
	PayAmount             float64                                            `gorm:"column:pay_amount;type:decimal(10,2);comment:应付金额（实际支付金额）" json:"pay_amount"`
	FreightAmount         float64                                            `gorm:"column:freight_amount;type:decimal(10,2);comment:运费金额" json:"freight_amount"`
	PromotionAmount       float64                                            `gorm:"column:promotion_amount;type:decimal(10,2);comment:促销优化金额（促销价、满减、阶梯价）" json:"promotion_amount"`
	IntegrationAmount     float64                                            `gorm:"column:integration_amount;type:decimal(10,2);comment:积分抵扣金额" json:"integration_amount"`
	CouponAmount          float64                                            `gorm:"column:coupon_amount;type:decimal(10,2);comment:优惠券抵扣金额" json:"coupon_amount"`
	DiscountAmount        float64                                            `gorm:"column:discount_amount;type:decimal(10,2);comment:管理员后台调整订单使用的折扣金额" json:"discount_amount"`
	PayType               int                                                `gorm:"column:pay_type;type:int(1);comment:支付方式：0->未支付；1->支付宝；2->微信" json:"pay_type"`
	SourceType            int                                                `gorm:"column:source_type;type:int(1);comment:订单来源：0->PC订单；1->app订单" json:"source_type"`
	Status                int                                                `gorm:"column:status;type:int(1);comment:订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单" json:"status"`
	OrderType             int                                                `gorm:"column:order_type;type:int(1);comment:订单类型：0->正常订单；1->秒杀订单" json:"order_type"`
	DeliveryCompany       string                                             `gorm:"column:delivery_company;type:varchar(64);comment:物流公司(配送方式)" json:"delivery_company"`
	DeliverySn            string                                             `gorm:"column:delivery_sn;type:varchar(64);comment:物流单号" json:"delivery_sn"`
	AutoConfirmDay        int                                                `gorm:"column:auto_confirm_day;type:int(11);comment:自动确认时间（天）" json:"auto_confirm_day"`
	Integration           int                                                `gorm:"column:integration;type:int(11);comment:可以获得的积分" json:"integration"`
	Growth                int                                                `gorm:"column:growth;type:int(11);comment:可以活动的成长值" json:"growth"`
	PromotionInfo         string                                             `gorm:"column:promotion_info;type:varchar(100);comment:活动信息" json:"promotion_info"`
	BillType              int                                                `gorm:"column:bill_type;type:int(1);comment:发票类型：0->不开发票；1->电子发票；2->纸质发票" json:"bill_type"`
	BillHeader            string                                             `gorm:"column:bill_header;type:varchar(200);comment:发票抬头" json:"bill_header"`
	BillContent           string                                             `gorm:"column:bill_content;type:varchar(200);comment:发票内容" json:"bill_content"`
	BillReceiverPhone     string                                             `gorm:"column:bill_receiver_phone;type:varchar(32);comment:收票人电话" json:"bill_receiver_phone"`
	BillReceiverEmail     string                                             `gorm:"column:bill_receiver_email;type:varchar(64);comment:收票人邮箱" json:"bill_receiver_email"`
	ReceiverName          string                                             `gorm:"column:receiver_name;type:varchar(100);comment:收货人姓名;NOT NULL" json:"receiver_name"`
	ReceiverPhone         string                                             `gorm:"column:receiver_phone;type:varchar(32);comment:收货人电话;NOT NULL" json:"receiver_phone"`
	ReceiverPostCode      string                                             `gorm:"column:receiver_post_code;type:varchar(32);comment:收货人邮编" json:"receiver_post_code"`
	ReceiverProvince      string                                             `gorm:"column:receiver_province;type:varchar(32);comment:省份/直辖市" json:"receiver_province"`
	ReceiverCity          string                                             `gorm:"column:receiver_city;type:varchar(32);comment:城市" json:"receiver_city"`
	ReceiverRegion        string                                             `gorm:"column:receiver_region;type:varchar(32);comment:区" json:"receiver_region"`
	ReceiverDetailAddress string                                             `gorm:"column:receiver_detail_address;type:varchar(200);comment:详细地址" json:"receiver_detail_address"`
	Note                  string                                             `gorm:"column:note;type:varchar(500);comment:订单备注" json:"note"`
	ConfirmStatus         int                                                `gorm:"column:confirm_status;type:int(1);comment:确认收货状态：0->未确认；1->已确认" json:"confirm_status"`
	DeleteStatus          int                                                `gorm:"column:delete_status;type:int(1);default:0;comment:删除状态：0->未删除；1->已删除;NOT NULL" json:"delete_status"`
	UseIntegration        int                                                `gorm:"column:use_integration;type:int(11);comment:下单时使用的积分" json:"use_integration"`
	PaymentTime           time.Time                                          `gorm:"column:payment_time;type:datetime;comment:支付时间" json:"payment_time"`
	DeliveryTime          time.Time                                          `gorm:"column:delivery_time;type:datetime;comment:发货时间" json:"delivery_time"`
	ReceiveTime           time.Time                                          `gorm:"column:receive_time;type:datetime;comment:确认收货时间" json:"receive_time"`
	CommentTime           time.Time                                          `gorm:"column:comment_time;type:datetime;comment:评价时间" json:"comment_time"`
	ModifyTime            time.Time                                          `gorm:"column:modify_time;type:datetime;comment:修改时间" json:"modify_time"`
	OrderItemList         []oms_order_item.OmsOrderItem                      `gorm:"foreignKey:OrderId" json:"orderItemList"`
	HistoryList           []oms_order_operate_history.OmsOrderOperateHistory `gorm:"foreignKey:OrderId" json:"historyList"`
}

func (omsOrder *OmsOrder) TableName() string {
	return "oms_order"
}

func (omsOrder *OmsOrder) Create() {
	database.DB.Create(&omsOrder)
}

func (omsOrder *OmsOrder) Save() (rowsAffected int64) {
	result := database.DB.Save(&omsOrder)
	return result.RowsAffected
}

func (omsOrder *OmsOrder) Updates(data map[string]interface{}) (rowsAffected int64) {
	result := database.DB.Model(&omsOrder).Updates(data)
	return result.RowsAffected
}

func (omsOrder *OmsOrder) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&omsOrder)
	return result.RowsAffected
}
