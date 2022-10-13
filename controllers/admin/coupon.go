package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/models/sms_coupon"
	"go-mall-api/models/sms_coupon_history"
	"go-mall-api/models/sms_coupon_product_category_relation"
	"go-mall-api/models/sms_coupon_product_relation"
	"strconv"
)

func (c AdminController) CouponList(ctx *gin.Context, request *entity.CouponListRequest) (*entity.CouponListReply, error) {
	where := ""
	if request.Name != "" {
		whereAnd(&where)
		where += fmt.Sprintf("name LIKE %s", "'%"+request.Name+"%'")
	}
	if request.Type != "" {
		whereAnd(&where)
		where += fmt.Sprintf("type = '%s'", request.Type)
	}

	list, paging := sms_coupon.Paginate(ctx, request.PageSize, where, "id", "asc")
	replyList := make([]entity.CouponList, 0, request.PageSize)
	for _, v := range list {
		interposition := entity.CouponList{
			Id:           v.ID,
			Amount:       v.Amount,
			MinPoint:     v.MinPoint,
			Count:        v.Count,
			PerLimit:     v.PerLimit,
			Platform:     v.Platform,
			PublishCount: v.PublishCount,
			Type:         v.Type,
			UseCount:     v.UseCount,
			UseType:      v.UseType,
			MemberLevel:  v.MemberLevel,
			Name:         v.Name,
			Note:         v.Note,
			Code:         v.Code,
			StartTime:    v.StartTime.Format("2006-01-02 15:04:05"),
			EnableTime:   v.EnableTime.Format("2006-01-02 15:04:05"),
			EndTime:      v.EndTime.Format("2006-01-02 15:04:05"),
		}
		replyList = append(replyList, interposition)
	}

	return &entity.CouponListReply{
		List:        replyList,
		PagingAdmin: paging,
	}, nil
}

func (c AdminController) Coupon(ctx *gin.Context, request *entity.CouponRequest) (*entity.CouponList, error) {
	coupon := sms_coupon.Get(request.Id)
	productRelationData := sms_coupon_product_relation.GetByCouponId(request.Id)
	productCategoryRelationData := sms_coupon_product_category_relation.GetByCouponId(request.Id)
	productRelationList := make([]entity.ProductRelationList, 0, len(productRelationData))
	productCategoryRelationList := make([]entity.ProductCategoryRelationList, 0, len(productRelationData))
	for _, v := range productRelationData {
		interposition := entity.ProductRelationList{
			ProductId:   v.ProductId,
			ProductSn:   v.ProductSn,
			ProductName: v.ProductName,
		}
		productRelationList = append(productRelationList, interposition)
	}

	for _, value := range productCategoryRelationData {
		interposition := entity.ProductCategoryRelationList{
			ProductCategoryId:   value.ProductCategoryId,
			ParentCategoryName:  value.ParentCategoryName,
			ProductCategoryName: value.ProductCategoryName,
		}
		productCategoryRelationList = append(productCategoryRelationList, interposition)
	}

	return &entity.CouponList{
		Id:                          coupon.ID,
		Amount:                      coupon.Amount,
		MinPoint:                    coupon.MinPoint,
		Count:                       coupon.Count,
		PerLimit:                    coupon.PerLimit,
		Platform:                    coupon.Platform,
		PublishCount:                coupon.PublishCount,
		Type:                        coupon.Type,
		UseCount:                    coupon.UseCount,
		UseType:                     coupon.UseType,
		MemberLevel:                 coupon.MemberLevel,
		Name:                        coupon.Name,
		Note:                        coupon.Note,
		Code:                        coupon.Code,
		StartTime:                   coupon.StartTime.Format("2006-01-02 15:04:05"),
		EnableTime:                  coupon.EnableTime.Format("2006-01-02 15:04:05"),
		EndTime:                     coupon.EndTime.Format("2006-01-02 15:04:05"),
		ProductRelationList:         productRelationList,
		ProductCategoryRelationList: productCategoryRelationList,
	}, nil
}

func (c AdminController) CouponCreate(ctx *gin.Context, request *entity.CouponCreateRequest) error {
	coupon := sms_coupon.SmsCoupon{
		Type:         request.Type,
		Name:         request.Name,
		Platform:     request.Platform,
		Count:        request.PublishCount,
		Amount:       request.Amount,
		PerLimit:     request.PerLimit,
		MinPoint:     request.MinPoint,
		StartTime:    request.StartTime,
		EndTime:      request.EndTime,
		UseType:      request.UseType,
		Note:         request.Note,
		PublishCount: request.PublishCount,
		UseCount:     0,
		ReceiveCount: 0,
		EnableTime:   request.EnableTime,
		Code:         "",
		MemberLevel:  0,
	}

	err := sms_coupon.CouponAdd(coupon, request)
	return err
}

func (c AdminController) CouponUpdate(ctx *gin.Context, request *entity.CouponUpdateRequest) error {
	coupon := sms_coupon.Get(strconv.Itoa(request.Id))

	err := sms_coupon.CouponUpdate(coupon, request)
	return err
}

func (c AdminController) CouponDelete(ctx *gin.Context, request *entity.CouponDeleteRequest) error {
	coupon := sms_coupon.Get(request.Id)
	err := sms_coupon.CouponDelete(coupon)
	return err
}

func (c AdminController) CouponHistory(ctx *gin.Context, request *entity.CouponHistoryRequest) (*entity.CouponHistoryReply, error) {
	where := map[string]interface{}{"coupon_id": request.CouponId}
	list, paging := sms_coupon_history.Paginate(ctx, request.PageSize, where, "id", "asc")

	replyList := make([]entity.CouponHistory, 0, request.PageSize)
	var useTime string
	for _, v := range list {
		if v.UseStatus == 1 {
			useTime = v.UseTime.Format("2006-01-02 15:04:05")
		} else {
			useTime = ""
		}

		interposition := entity.CouponHistory{
			Id:             v.ID,
			CouponId:       v.CouponId,
			GetType:        v.GetType,
			MemberId:       v.MemberId,
			OrderId:        v.OrderId,
			UseStatus:      v.UseStatus,
			CouponCode:     v.CouponCode,
			MemberNickname: v.MemberNickname,
			OrderSn:        v.OrderSn,
			UseTime:        useTime,
			CreateTime:     v.CreateTime.Format("2006-01-02 15:04:05"),
		}
		replyList = append(replyList, interposition)
	}

	return &entity.CouponHistoryReply{
		List:        replyList,
		PagingAdmin: paging,
	}, nil
}
