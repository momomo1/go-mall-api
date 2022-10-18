package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/models/oms_order_return_apply"
	"go-mall-api/models/oms_order_return_reason"
	"strconv"
	"strings"
	"time"
)

func (c AdminController) ReturnApplyList(ctx *gin.Context, request *entity.ReturnApplyListRequest) (*entity.ReturnApplyListReply, error) {
	where := ""
	if request.HandleMan != "" {
		whereAnd(&where)
		where += fmt.Sprintf("handle_man LIKE %s", "'%"+request.HandleMan+"%'")
	}
	if request.Id != "" {
		whereAnd(&where)
		where += fmt.Sprintf("id = '%s'", request.Id)
	}

	if request.Status != "" {
		whereAnd(&where)
		where += fmt.Sprintf("status = '%s'", request.Status)
	}

	if request.CreateTime != "" {
		whereAnd(&where)
		where += fmt.Sprintf("create_time <= '%s'", request.CreateTime)
	}

	if request.HandleTime != "" {
		whereAnd(&where)
		where += fmt.Sprintf("handle_time <= '%s'", request.HandleTime)
	}
	list, paging := oms_order_return_apply.Paginate(ctx, request.PageSize, where, "order_sn", "asc")
	replyList := make([]entity.ReturnApplyList, 0, request.PageSize)
	for _, v := range list {
		interposition := entity.ReturnApplyList{
			Id:               v.ID,
			Status:           v.Status,
			ProductRealPrice: v.ProductRealPrice,
			ReturnName:       v.ReturnName,
			MemberUsername:   v.MemberUsername,
			CreateTime:       v.CreateTime.Format("2006-01-02 15:04:05"),
			HandleTime:       v.HandleTime.Format("2006-01-02 15:04:05"),
		}
		replyList = append(replyList, interposition)
	}

	return &entity.ReturnApplyListReply{
		List:        replyList,
		PagingAdmin: paging,
	}, nil
}

func (c AdminController) ReturnApply(ctx *gin.Context, request *entity.ReturnApplyRequest) (*entity.ReturnApplyReply, error) {
	data := oms_order_return_apply.Get(request.Id)
	return &entity.ReturnApplyReply{
		Id:               data.ID,
		OrderId:          data.OrderId,
		ProductCount:     data.ProductCount,
		ProductId:        data.ProductId,
		Status:           data.Status,
		ProductPrice:     data.ProductPrice,
		ProductRealPrice: data.ProductRealPrice,
		//CompanyAddress:   data.CompanyAddress,
		CompanyAddressId: data.CompanyAddressId,
		Description:      data.Description,
		HandleMan:        data.HandleMan,
		HandleNote:       data.HandleNote,
		HandleTime:       data.HandleTime.Format("2006-01-02 15:04:05"),
		MemberUsername:   data.MemberUsername,
		OrderSn:          data.OrderSn,
		ProductAttr:      data.ProductAttr,
		ProductBrand:     data.ProductBrand,
		ProductName:      data.ProductName,
		ProductPic:       data.ProductPic,
		ProofPics:        data.ProofPics,
		Reason:           data.Reason,
		ReceiveMan:       data.ReceiveMan,
		ReceiveNote:      data.ReceiveNote,
		ReceiveTime:      data.ReceiveTime.Format("2006-01-02 15:04:05"),
		ReturnAmount:     data.ReturnAmount,
		ReturnName:       data.ReturnName,
		ReturnPhone:      data.ReturnPhone,
	}, nil
}

func (c AdminController) ReturnApplyUpdateStatus(ctx *gin.Context, request *entity.ReturnApplyUpdateStatusRequest) error {
	data := oms_order_return_apply.Get(request.Id)
	data.Updates(map[string]interface{}{
		"Status":           request.Status,
		"CompanyAddressId": request.CompanyAddressId,
		"ReceiveNote":      request.ReceiveNote,
		"HandleNote":       request.HandleNote,
		"ReturnAmount":     request.ReturnAmount,
	})
	return nil
}

func (c AdminController) ReturnReasonList(ctx *gin.Context, request *entity.ReturnReasonListRequest) (*entity.ReturnReasonListReply, error) {
	list, paging := oms_order_return_reason.Paginate(ctx, request.PageSize, "", "sort", "desc")
	roleList := make([]entity.ReturnReasonList, 0, request.PageSize)
	for _, v := range list {
		interposition := entity.ReturnReasonList{
			Id:         v.ID,
			Sort:       v.Sort,
			Status:     v.Status,
			Name:       v.Name,
			CreateTime: v.CreateTime.Format("2006-01-02 15:04:05"),
		}
		roleList = append(roleList, interposition)
	}

	return &entity.ReturnReasonListReply{
		List:        roleList,
		PagingAdmin: paging,
	}, nil
}

func (c AdminController) ReturnReason(ctx *gin.Context, request *entity.ReturnReasonRequest) (*entity.ReturnReasonList, error) {
	returnReason := oms_order_return_reason.Get(request.Id)

	return &entity.ReturnReasonList{
		Id:         returnReason.ID,
		Sort:       returnReason.Sort,
		Status:     returnReason.Status,
		Name:       returnReason.Name,
		CreateTime: returnReason.CreateTime.Format("2006-01-02 15:04:05"),
	}, nil
}

func (c AdminController) ReturnReasonUpdateStatus(ctx *gin.Context, request *entity.ReturnReasonUpdateStatusRequest) error {
	ids := strings.SplitN(request.Ids, ",", -1)
	for _, v := range ids {
		returnReason := oms_order_return_reason.Get(v)
		returnReason.Updates(map[string]interface{}{
			"status": request.Status,
		})
	}
	return nil
}

func (c AdminController) ReturnReasonCreate(ctx *gin.Context, request *entity.ReturnReasonCreateRequest) error {
	sort, _ := strconv.Atoi(request.Sort)
	reason := oms_order_return_reason.OmsOrderReturnReason{
		Name:       request.Name,
		Sort:       sort,
		Status:     request.Status,
		CreateTime: time.Now(),
	}
	reason.Create()
	return nil
}

func (c AdminController) ReturnReasonUpdate(ctx *gin.Context, request *entity.ReturnReasonUpdateRequest) error {
	returnReason := oms_order_return_reason.Get(strconv.Itoa(request.Id))
	data := map[string]interface{}{
		"Name":   request.Name,
		"Status": request.Status,
	}

	if request.Sort != "" {
		data["Sort"] = request.Sort
	}
	returnReason.Updates(data)
	return nil
}

func (c AdminController) ReturnReasonDelete(ctx *gin.Context, request *entity.ReturnReasonDeleteRequest) error {
	ids := strings.SplitN(request.Ids, ",", -1)
	for _, v := range ids {
		returnReason := oms_order_return_reason.Get(v)
		returnReason.Delete()
	}
	return nil
}
