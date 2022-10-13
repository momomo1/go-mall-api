package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/models/sms_flash_promotion"
	"go-mall-api/models/sms_flash_promotion_product_relation"
	"go-mall-api/models/sms_flash_promotion_session"
	"go-mall-api/pkg/helpers"
	"strconv"
	"time"
)

func (c AdminController) FlashList(ctx *gin.Context, request *entity.FlashListRequest) (*entity.FlashListReply, error) {
	where := ""
	if request.Keyword != "" {
		whereAnd(&where)
		where += fmt.Sprintf("title LIKE %s", "'%"+request.Keyword+"%'")
	}

	list, paging := sms_flash_promotion.Paginate(ctx, request.PageSize, where, "id", "asc")
	replyList := make([]entity.FlashList, 0, request.PageSize)
	for _, v := range list {
		interposition := entity.FlashList{
			Id:         v.ID,
			Status:     v.Status,
			Title:      v.Title,
			CreateTime: v.CreateTime.Format("2006-01-02 15:04:05"),
			StartDate:  v.StartDate.Format("2006-01-02 15:04:05"),
			EndDate:    v.EndDate.Format("2006-01-02 15:04:05"),
		}
		replyList = append(replyList, interposition)
	}

	return &entity.FlashListReply{
		List:        replyList,
		PagingAdmin: paging,
	}, nil
}

func (c AdminController) FlashUpdateStatus(ctx *gin.Context, request *entity.FlashUpdateStatusRequest) error {
	flashPromotion := sms_flash_promotion.Get(request.Id)
	flashPromotion.Updates(map[string]interface{}{
		"status": request.Status,
	})
	return nil
}

func (c AdminController) FlashCreate(ctx *gin.Context, request *entity.FlashCreateRequest) error {
	flashPromotion := sms_flash_promotion.SmsFlashPromotion{
		Title:      request.Title,
		StartDate:  request.StartDate,
		EndDate:    request.EndDate,
		Status:     request.Status,
		CreateTime: time.Now(),
	}

	flashPromotion.Create()
	return nil
}

func (c AdminController) FlashUpdate(ctx *gin.Context, request *entity.FlashUpdateRequest) error {
	flashPromotion := sms_flash_promotion.Get(strconv.Itoa(request.Id))
	startDate := helpers.TimeStringToGoTime(request.StartDate)
	endDate := helpers.TimeStringToGoTime(request.EndDate)

	flashPromotion.Updates(map[string]interface{}{
		"status":     request.Status,
		"title":      request.Title,
		"start_date": startDate,
		"end_date":   endDate,
	})
	return nil
}

func (c AdminController) FlashDelete(ctx *gin.Context, request *entity.FlashDeleteRequest) error {
	flashPromotion := sms_flash_promotion.Get(request.Id)
	flashPromotion.Delete()
	return nil
}

func (c AdminController) FlashSessionList(ctx *gin.Context) (*entity.FlashSessionListReply, error) {
	flashPromotionSession := sms_flash_promotion_session.All()
	replyList := make([]entity.FlashSessionList, 0, len(flashPromotionSession))

	for _, v := range flashPromotionSession {
		interposition := entity.FlashSessionList{
			Id:         v.ID,
			Status:     v.Status,
			Name:       v.Name,
			CreateTime: v.CreateTime,
			StartTime:  "1970-01-01" + helpers.TimeStringToGoTime(v.StartTime).Format("T15:04:05"),
			EndTime:    "1970-01-01" + helpers.TimeStringToGoTime(v.EndTime).Format("T15:04:05"),
		}
		replyList = append(replyList, interposition)
	}
	return &entity.FlashSessionListReply{
		List: replyList,
	}, nil
}

func (c AdminController) FlashSessionSelectList(ctx *gin.Context, request *entity.FlashSessionSelectListRequest) (*entity.FlashSessionSelectListReply, error) {
	relation := sms_flash_promotion_session.AllByStartUsing()
	replyList := make([]entity.FlashSessionSelectList, 0, len(relation))
	for _, v := range relation {
		count := sms_flash_promotion_product_relation.GetWhereCount(request.FlashPromotionId, strconv.FormatUint(v.ID, 10))
		interposition := entity.FlashSessionSelectList{
			Id:           v.ID,
			ProductCount: count,
			Name:         v.Name,
			Status:       v.Status,
			StartTime:    "1970-01-01" + helpers.TimeStringToGoTime(v.StartTime).Format("T15:04:05"),
			EndTime:      "1970-01-01" + helpers.TimeStringToGoTime(v.EndTime).Format("T15:04:05"),
			CreateTime:   v.CreateTime,
		}
		replyList = append(replyList, interposition)
	}

	return &entity.FlashSessionSelectListReply{List: replyList}, nil
}

func (c AdminController) FlashSessionCreate(ctx *gin.Context, request *entity.FlashSessionCreateRequest) error {
	startTime := time.Unix(request.StartTime.Unix(), 0).Format("15:04:05")
	endTime := time.Unix(request.EndTime.Unix(), 0).Format("15:04:05")

	flashPromotionSession := sms_flash_promotion_session.SmsFlashPromotionSession{
		Name:       request.Name,
		StartTime:  startTime,
		EndTime:    endTime,
		Status:     request.Status,
		CreateTime: time.Now(),
	}
	flashPromotionSession.Create()
	return nil
}

func (c AdminController) FlashSessionUpdateStatus(ctx *gin.Context, request *entity.FlashSessionUpdateStatusRequest) error {
	flashPromotionSession := sms_flash_promotion_session.Get(request.Id)
	flashPromotionSession.Updates(map[string]interface{}{
		"status": request.Status,
	})
	return nil
}

func (c AdminController) FlashSessionUpdate(ctx *gin.Context, request *entity.FlashSessionUpdateRequest) error {
	startTime := time.Unix(request.StartTime.Unix(), 0).Format("15:04:05")
	endTime := time.Unix(request.EndTime.Unix(), 0).Format("15:04:05")
	flashPromotionSession := sms_flash_promotion_session.Get(strconv.Itoa(request.Id))
	flashPromotionSession.Updates(map[string]interface{}{
		"Name":      request.Name,
		"StartTime": startTime,
		"EndTime":   endTime,
		"Status":    request.Status,
	})
	return nil
}

func (c AdminController) FlashSessionDelete(ctx *gin.Context, request *entity.FlashSessionDeleteRequest) error {
	flashPromotionSession := sms_flash_promotion_session.Get(request.Id)
	flashPromotionSession.Delete()
	return nil
}

func (c AdminController) FlashProductRelationList(ctx *gin.Context) error {
	return nil
}

func (c AdminController) FlashProductRelationCreate(ctx *gin.Context) error {
	return nil
}

func (c AdminController) FlashProductRelationUpdate(ctx *gin.Context) error {
	return nil
}

func (c AdminController) FlashProductRelationDelete(ctx *gin.Context) error {
	return nil
}
