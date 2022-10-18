package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/models/oms_order"
	"go-mall-api/models/oms_order_operate_history"
	"go-mall-api/models/oms_order_setting"
	"strconv"
	"strings"
)

func (c AdminController) OrderList(ctx *gin.Context, request *entity.OrderListRequest) (*entity.OrderListReply, error) {
	where := ""
	if request.ReceiverKeyword != "" {
		whereAnd(&where)
		where += fmt.Sprintf("receiver_name LIKE %s", "'%"+request.ReceiverKeyword+"%'")
	}

	if request.OrderSn != "" {
		whereAnd(&where)
		where += fmt.Sprintf("order_sn = '%s'", request.OrderSn)
	}

	if request.OrderSn != "" {
		whereAnd(&where)
		where += fmt.Sprintf("order_sn = '%s'", request.OrderSn)
	}

	if request.Status != "" {
		whereAnd(&where)
		where += fmt.Sprintf("status = '%s'", request.Status)
	}

	if request.OrderType != "" {
		whereAnd(&where)
		where += fmt.Sprintf("order_type = '%s'", request.OrderType)
	}

	if request.SourceType != "" {
		whereAnd(&where)
		where += fmt.Sprintf("source_type = '%s'", request.SourceType)
	}

	if request.CreateTime != "" {
		whereAnd(&where)
		where += fmt.Sprintf("create_time <= '%s'", request.CreateTime)
	}

	whereAnd(&where)
	where += fmt.Sprintf("delete_status = '%d'", 0)

	list, paging := oms_order.Paginate(ctx, request.PageSize, where, "id", "asc")
	replyList := make([]entity.OrderList, 0, request.PageSize)
	for _, v := range list {
		orderItemList := make([]entity.OrderItemList, 0)
		historyList := make([]entity.HistoryList, 0)
		interposition := OrderFilling(v, orderItemList, historyList)
		replyList = append(replyList, interposition)
	}

	return &entity.OrderListReply{
		List:        replyList,
		PagingAdmin: paging,
	}, nil
}

func (c AdminController) Order(ctx *gin.Context, request *entity.OrderRequest) (*entity.OrderList, error) {
	detail := oms_order.Detail(request.Id)
	orderItemList := make([]entity.OrderItemList, 0, len(detail.OrderItemList))
	for _, v := range detail.OrderItemList {
		productData := entity.OrderItemList{
			Id:              v.ID,
			ProductQuantity: v.ProductQuantity,
			ProductId:       v.ProductId,
			ProductPrice:    v.ProductPrice,
			ProductAttr:     v.ProductAttr,
			ProductBrand:    v.ProductBrand,
			ProductName:     v.ProductName,
			ProductPic:      v.ProductPic,
			ProductSn:       v.ProductSn,
		}
		orderItemList = append(orderItemList, productData)
	}

	historyList := make([]entity.HistoryList, 0, len(detail.HistoryList))
	for _, v := range detail.HistoryList {
		historyData := entity.HistoryList{
			Id:          v.ID,
			OrderStatus: v.OrderStatus,
			Note:        v.Note,
			OperateMan:  v.OperateMan,
			CreateTime:  v.CreateTime.Format("2006-01-02 15:04:05"),
		}
		historyList = append(historyList, historyData)
	}
	data := OrderFilling(detail, orderItemList, historyList)
	return &data, nil
}

func (c AdminController) OrderUpdateNote(ctx *gin.Context, request *entity.OrderUpdateNoteRequest) error {
	order := oms_order.Get(request.Id)
	order.Updates(map[string]interface{}{
		"note": request.Note,
	})
	userName := ctx.GetString("current_user_name")
	oms_order_operate_history.AddOperateHistory(int64(order.ID), order.Status, userName, "修改费用信息:"+request.Note)
	return nil
}

func (c AdminController) OrderDelete(ctx *gin.Context, request *entity.OrderDeleteRequest) error {
	ids := strings.SplitN(request.Ids, ",", -1)
	userName := ctx.GetString("current_user_name")
	for _, v := range ids {
		order := oms_order.Get(v)
		order.Updates(map[string]interface{}{
			"delete_status": 1,
		})
		oms_order_operate_history.AddOperateHistory(int64(order.ID), order.Status, userName, "订单删除")
	}
	return nil
}

func (c AdminController) OrderUpdateReceiverInfo(ctx *gin.Context, request *entity.OrderUpdateReceiverInfoRequest) error {
	order := oms_order.Get(strconv.FormatInt(request.OrderId, 10))
	order.Updates(map[string]interface{}{
		"Status":                request.Status,
		"ReceiverCity":          request.ReceiverCity,
		"ReceiverDetailAddress": request.ReceiverDetailAddress,
		"ReceiverName":          request.ReceiverName,
		"ReceiverPhone":         request.ReceiverPhone,
		"ReceiverPostCode":      request.ReceiverPostCode,
		"ReceiverProvince":      request.ReceiverProvince,
		"ReceiverRegion":        request.ReceiverRegion,
	})
	userName := ctx.GetString("current_user_name")
	oms_order_operate_history.AddOperateHistory(int64(order.ID), order.Status, userName, "修改收货人信息")
	return nil
}

func (c AdminController) OrderUpdateMoneyInfo(ctx *gin.Context, request *entity.OrderUpdateMoneyInfoRequest) error {
	order := oms_order.Get(strconv.FormatInt(request.OrderId, 10))
	order.Updates(map[string]interface{}{
		"Status":         request.Status,
		"DiscountAmount": request.DiscountAmount,
		"FreightAmount":  request.FreightAmount,
	})
	userName := ctx.GetString("current_user_name")
	oms_order_operate_history.AddOperateHistory(int64(order.ID), order.Status, userName, "修改费用信息")
	return nil
}

func (c AdminController) OrderUpdateDelivery(ctx *gin.Context, request *[]entity.OrderUpdateDeliveryRequest) error {
	userName := ctx.GetString("current_user_name")
	for _, v := range *request {
		order := oms_order.GetByOrderSn(v.OrderSn)
		order.Updates(map[string]interface{}{
			"delivery_company": v.DeliveryCompany,
			"delivery_sn":      v.DeliverySn,
			"status":           oms_order.Shipped,
		})

		oms_order_operate_history.AddOperateHistory(int64(order.ID), oms_order.Shipped, userName, "完成发货")
	}
	return nil
}

func (c AdminController) OrderUpdateClose(ctx *gin.Context, request *entity.OrderUpdateCloseRequest) error {
	userName := ctx.GetString("current_user_name")
	ids := strings.SplitN(request.Ids, ",", -1)
	for _, v := range ids {
		order := oms_order.Get(v)
		order.Updates(map[string]interface{}{
			"status": oms_order.Closed,
		})
		oms_order_operate_history.AddOperateHistory(int64(order.ID), oms_order.Closed, userName, "订单关闭:"+request.Note)
	}
	return nil
}

func (c AdminController) OrderSetting(ctx *gin.Context) (*entity.OrderSettingReply, error) {
	setting := oms_order_setting.Get("1")

	return &entity.OrderSettingReply{
		Id:                  setting.ID,
		CommentOvertime:     setting.CommentOvertime,
		ConfirmOvertime:     setting.ConfirmOvertime,
		FinishOvertime:      setting.FinishOvertime,
		FlashOrderOvertime:  setting.FlashOrderOvertime,
		NormalOrderOvertime: setting.NormalOrderOvertime,
	}, nil
}

func (c AdminController) OrderSettingUpdate(ctx *gin.Context, request *entity.OrderSettingRequest) error {
	setting := oms_order_setting.Get("1")
	data := map[string]interface{}{}
	if request.CommentOvertime != "" {
		data["CommentOvertime"] = request.CommentOvertime
	}
	if request.ConfirmOvertime != "" {
		data["ConfirmOvertime"] = request.ConfirmOvertime
	}
	if request.FinishOvertime != "" {
		data["FinishOvertime"] = request.FinishOvertime
	}
	if request.FlashOrderOvertime != "" {
		data["FlashOrderOvertime"] = request.FlashOrderOvertime
	}
	if request.NormalOrderOvertime != "" {
		data["NormalOrderOvertime"] = request.NormalOrderOvertime
	}
	setting.Updates(data)
	return nil
}

func OrderFilling(v oms_order.OmsOrder, orderItemList []entity.OrderItemList, historyList []entity.HistoryList) entity.OrderList {
	interposition := entity.OrderList{
		Id:                    v.ID,
		AutoConfirmDay:        v.AutoConfirmDay,
		ConfirmStatus:         v.ConfirmStatus,
		CouponAmount:          v.CouponAmount,
		CouponId:              v.CouponId,
		DeleteStatus:          v.DeleteStatus,
		DiscountAmount:        v.DiscountAmount,
		FreightAmount:         v.FreightAmount,
		Growth:                v.Growth,
		Integration:           v.Integration,
		IntegrationAmount:     v.IntegrationAmount,
		MemberId:              v.MemberId,
		OrderType:             v.OrderType,
		PayAmount:             v.PayAmount,
		PayType:               v.PayType,
		PromotionAmount:       v.PromotionAmount,
		SourceType:            v.SourceType,
		Status:                v.Status,
		TotalAmount:           v.TotalAmount,
		BillContent:           v.BillContent,
		BillHeader:            v.BillHeader,
		BillReceiverEmail:     v.BillReceiverEmail,
		BillReceiverPhone:     v.BillReceiverPhone,
		BillType:              strconv.Itoa(v.BillType),
		DeliveryCompany:       v.DeliveryCompany,
		DeliverySn:            v.DeliverySn,
		MemberUsername:        v.MemberUsername,
		Note:                  v.Note,
		OrderSn:               v.OrderSn,
		PromotionInfo:         v.PromotionInfo,
		ReceiverCity:          v.ReceiverCity,
		ReceiverDetailAddress: v.ReceiverDetailAddress,
		ReceiverName:          v.ReceiverName,
		ReceiverPhone:         v.ReceiverPhone,
		ReceiverPostCode:      v.ReceiverPostCode,
		ReceiverProvince:      v.ReceiverProvince,
		ReceiverRegion:        v.ReceiverRegion,
		CommentTime:           v.CommentTime.Format("2006-01-02 15:04:05"),
		DeliveryTime:          v.DeliveryTime.Format("2006-01-02 15:04:05"),
		ModifyTime:            v.ModifyTime.Format("2006-01-02 15:04:05"),
		CreateTime:            v.CreateTime.Format("2006-01-02 15:04:05"),
		PaymentTime:           v.PaymentTime.Format("2006-01-02 15:04:05"),
		ReceiveTime:           v.ReceiveTime.Format("2006-01-02 15:04:05"),
		OrderItemList:         orderItemList,
		HistoryList:           historyList,
	}

	return interposition
}
