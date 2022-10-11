package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/pkg/helpers"
	"go-mall-api/pkg/response"
)

func OrderSettingHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		out, _ := c.OrderSetting(ctx)
		response.OkWithData(ctx, out)
	}
}

func OrderSettingUpdateHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.OrderSettingRequest
		j := helpers.GetRequestPayload(ctx)
		json.Unmarshal(j, &request)
		err := c.OrderSettingUpdate(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}
