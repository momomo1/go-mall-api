package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/pkg/helpers"
	"go-mall-api/pkg/response"
	"strconv"
)

func ReturnApplyListHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ReturnApplyListRequest
		query := helpers.GetQueryParams(ctx)
		pageSize, _ := strconv.Atoi(query["pageSize"])
		request.PageSize = pageSize
		request.Id = query["id"]
		request.Status = query["status"]
		request.HandleMan = query["handleMan"]
		request.CreateTime = query["createTime"]
		request.HandleTime = query["handleTime"]

		out, _ := c.ReturnApplyList(ctx, &request)
		response.OkWithData(ctx, out)
	}
}

func ReturnApplyHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		response.Ok(ctx)
	}
}

func ReturnApplyUpdateStatusHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		response.Ok(ctx)
	}
}

func ReturnReasonListHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ReturnReasonListRequest
		query := helpers.GetQueryParams(ctx)
		pageSize, _ := strconv.Atoi(query["pageSize"])
		request.PageSize = pageSize
		out, _ := c.ReturnReasonList(ctx, &request)
		response.OkWithData(ctx, out)
	}
}

func ReturnReasonHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ReturnReasonRequest
		request.Id = ctx.Param("id")
		out, _ := c.ReturnReason(ctx, &request)
		response.OkWithData(ctx, out)
	}
}

func ReturnReasonUpdateStatusHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ReturnReasonUpdateStatusRequest
		query := helpers.GetQueryParams(ctx)
		request.Ids = query["ids"]
		request.Status = query["status"]
		err := c.ReturnReasonUpdateStatus(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func ReturnReasonCreateHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ReturnReasonCreateRequest
		j := helpers.GetRequestPayload(ctx)
		json.Unmarshal(j, &request)
		err := c.ReturnReasonCreate(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func ReturnReasonUpdateHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ReturnReasonUpdateRequest
		j := helpers.GetRequestPayload(ctx)
		json.Unmarshal(j, &request)
		err := c.ReturnReasonUpdate(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func ReturnReasonDeleteHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ReturnReasonDeleteRequest
		query := helpers.GetQueryParams(ctx)
		request.Ids = query["ids"]
		err := c.ReturnReasonDelete(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}
