package admin

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/models/oms_order_setting"
)

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
