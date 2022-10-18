package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/pkg/helpers"
	"go-mall-api/pkg/response"
	"strconv"
)

func OrderListHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.OrderListRequest
		query := helpers.GetQueryParams(ctx)
		pageSize, _ := strconv.Atoi(query["pageSize"])
		request.PageSize = pageSize
		request.ReceiverKeyword = query["receiverKeyword"]
		request.OrderSn = query["orderSn"]
		request.Status = query["status"]
		request.OrderType = query["orderType"]
		request.SourceType = query["sourceType"]
		request.CreateTime = query["createTime"]

		out, _ := c.OrderList(ctx, &request)
		response.OkWithData(ctx, out)
	}
}

func OrderHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.OrderRequest
		request.Id = ctx.Param("id")

		out, _ := c.Order(ctx, &request)
		response.OkWithData(ctx, out)
	}
}

func OrderUpdateNoteHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.OrderUpdateNoteRequest
		query := helpers.GetQueryParams(ctx)
		request.Id = query["id"]
		request.Note = query["note"]

		err := c.OrderUpdateNote(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func OrderDeleteHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.OrderDeleteRequest
		query := helpers.GetQueryParams(ctx)
		request.Ids = query["ids"]

		err := c.OrderDelete(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func OrderUpdateReceiverInfoHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.OrderUpdateReceiverInfoRequest
		j := helpers.GetRequestPayload(ctx)
		json.Unmarshal(j, &request)
		err := c.OrderUpdateReceiverInfo(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func OrderUpdateMoneyInfoHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.OrderUpdateMoneyInfoRequest
		j := helpers.GetRequestPayload(ctx)
		json.Unmarshal(j, &request)
		err := c.OrderUpdateMoneyInfo(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func OrderUpdateDeliveryHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request []entity.OrderUpdateDeliveryRequest
		j := helpers.GetRequestPayload(ctx)
		json.Unmarshal(j, &request)
		err := c.OrderUpdateDelivery(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func OrderUpdateCloseHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.OrderUpdateCloseRequest
		query := helpers.GetQueryParams(ctx)
		request.Ids = query["ids"]
		request.Note = query["note"]

		err := c.OrderUpdateClose(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

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
