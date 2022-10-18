package entity

import (
	"go-mall-api/pkg/paginator"
)

type OrderListRequest struct {
	PageNum         int    `json:"pageNum"`
	PageSize        int    `json:"pageSize"`
	OrderSn         string `json:"orderSn"`
	ReceiverKeyword string `json:"receiverKeyword"`
	Status          string `json:"status"`
	OrderType       string `json:"orderType"`
	SourceType      string `json:"sourceType"`
	CreateTime      string `json:"createTime"`
}

type OrderListReply struct {
	List []OrderList `json:"list"`
	paginator.PagingAdmin
}

type OrderList struct {
	Id                    uint64          `json:"id"`
	AutoConfirmDay        int             `json:"autoConfirmDay"`
	ConfirmStatus         int             `json:"confirmStatus"`
	CouponId              int64           `json:"couponId"`
	DeleteStatus          int             `json:"deleteStatus"`
	Growth                int             `json:"growth"`
	Integration           int             `json:"integration"`
	MemberId              int64           `json:"memberId"`
	OrderType             int             `json:"orderType"`
	PayType               int             `json:"payType"`
	SourceType            int             `json:"sourceType"`
	Status                int             `json:"status"`
	CouponAmount          float64         `json:"couponAmount"`
	DiscountAmount        float64         `json:"discountAmount"`
	FreightAmount         float64         `json:"freightAmount"`
	TotalAmount           float64         `json:"totalAmount"`
	IntegrationAmount     float64         `json:"integrationAmount"`
	PayAmount             float64         `json:"payAmount"`
	PromotionAmount       float64         `json:"promotionAmount"`
	BillContent           string          `json:"billContent"`
	BillHeader            string          `json:"billHeader"`
	BillReceiverEmail     string          `json:"billReceiverEmail"`
	BillReceiverPhone     string          `json:"billReceiverPhone"`
	BillType              string          `json:"billType"`
	DeliveryCompany       string          `json:"deliveryCompany"`
	DeliverySn            string          `json:"deliverySn"`
	MemberUsername        string          `json:"memberUsername"`
	Note                  string          `json:"note"`
	OrderSn               string          `json:"orderSn"`
	PromotionInfo         string          `json:"promotionInfo"`
	ReceiverCity          string          `json:"receiverCity"`
	ReceiverDetailAddress string          `json:"receiverDetailAddress"`
	ReceiverName          string          `json:"receiverName"`
	ReceiverPhone         string          `json:"receiverPhone"`
	ReceiverPostCode      string          `json:"receiverPostCode"`
	ReceiverProvince      string          `json:"receiverProvince"`
	ReceiverRegion        string          `json:"receiverRegion"`
	CommentTime           string          `json:"commentTime"`
	DeliveryTime          string          `json:"deliveryTime"`
	ModifyTime            string          `json:"modifyTime"`
	CreateTime            string          `json:"createTime"`
	PaymentTime           string          `json:"paymentTime"`
	ReceiveTime           string          `json:"receiveTime"`
	OrderItemList         []OrderItemList `json:"orderItemList"`
	HistoryList           []HistoryList   `json:"historyList"`
}

type OrderRequest struct {
	Id string `json:"id"`
}

type OrderItemList struct {
	Id              uint64  `json:"id"`
	ProductQuantity int     `json:"productQuantity"`
	ProductId       int64   `json:"productId"`
	ProductPrice    float64 `json:"productPrice"`
	ProductAttr     string  `json:"productAttr"`
	ProductBrand    string  `json:"productBrand"`
	ProductName     string  `json:"productName"`
	ProductPic      string  `json:"productPic"`
	ProductSn       string  `json:"productSn"`
}

type HistoryList struct {
	Id          uint64 `json:"id"`
	OrderStatus int    `json:"orderStatus"`
	Note        string `json:"note"`
	OperateMan  string `json:"operateMan"`
	CreateTime  string `json:"createTime"`
}

type OrderUpdateNoteRequest struct {
	Id   string `json:"id"`
	Note string `json:"note"`
}

type OrderDeleteRequest struct {
	Ids string `json:"ids"`
}

type OrderUpdateReceiverInfoRequest struct {
	OrderId               int64  `json:"orderId"`
	Status                int    `json:"status"`
	ReceiverCity          string `json:"receiverCity"`
	ReceiverDetailAddress string `json:"receiverDetailAddress"`
	ReceiverName          string `json:"receiverName"`
	ReceiverPhone         string `json:"receiverPhone"`
	ReceiverPostCode      string `json:"receiverPostCode"`
	ReceiverProvince      string `json:"receiverProvince"`
	ReceiverRegion        string `json:"receiverRegion"`
}

type OrderUpdateMoneyInfoRequest struct {
	OrderId        int64 `json:"orderId"`
	Status         int   `json:"status"`
	DiscountAmount int   `json:"discountAmount"`
	FreightAmount  int   `json:"freightAmount"`
}

type OrderUpdateDeliveryRequest struct {
	Id              int    `json:"id"`
	OrderSn         string `json:"orderSn"`
	DeliveryCompany string `json:"deliveryCompany"`
	DeliverySn      string `json:"deliverySn"`
}

type OrderUpdateCloseRequest struct {
	Ids  string `json:"ids"`
	Note string `json:"note"`
}

type OrderSettingReply struct {
	Id                  uint64 `json:"id"`
	CommentOvertime     int    `json:"commentOvertime"`
	ConfirmOvertime     int    `json:"confirmOvertime"`
	FinishOvertime      int    `json:"finishOvertime"`
	FlashOrderOvertime  int    `json:"flashOrderOvertime"`
	NormalOrderOvertime int    `json:"normalOrderOvertime"`
}

type OrderSettingRequest struct {
	CommentOvertime     string `json:"commentOvertime"`
	ConfirmOvertime     string `json:"confirmOvertime"`
	FinishOvertime      string `json:"finishOvertime"`
	FlashOrderOvertime  string `json:"flashOrderOvertime"`
	NormalOrderOvertime string `json:"normalOrderOvertime"`
}
