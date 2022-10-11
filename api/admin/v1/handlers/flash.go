package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/pkg/helpers"
	"go-mall-api/pkg/response"
	"strconv"
)

func FlashListHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.FlashListRequest
		query := helpers.GetQueryParams(ctx)
		pageSize, _ := strconv.Atoi(query["pageSize"])
		request.PageSize = pageSize
		request.Keyword = query["keyword"]

		out, _ := c.FlashList(ctx, &request)
		response.OkWithData(ctx, out)
	}
}

func FlashUpdateStatusHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.FlashUpdateStatusRequest
		query := helpers.GetQueryParams(ctx)
		request.Id = ctx.Param("id")
		request.Status = query["status"]
		err := c.FlashUpdateStatus(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func FlashCreateHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.FlashCreateRequest
		j := helpers.GetRequestPayload(ctx)
		json.Unmarshal(j, &request)

		err := c.FlashCreate(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func FlashUpdateHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.FlashUpdateRequest
		j := helpers.GetRequestPayload(ctx)
		json.Unmarshal(j, &request)
		err := c.FlashUpdate(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func FlashDeleteHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.FlashDeleteRequest
		request.Id = ctx.Param("id")

		err := c.FlashDelete(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func FlashSessionListHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		out, _ := c.FlashSessionList(ctx)
		response.OkWithData(ctx, out.List)
	}
}

func FlashSessionSelectListHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.FlashSessionSelectListRequest
		query := helpers.GetQueryParams(ctx)
		request.FlashPromotionId = query["flashPromotionId"]
		out, _ := c.FlashSessionSelectList(ctx, &request)
		response.OkWithData(ctx, out.List)
	}
}

func FlashSessionCreateHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.FlashSessionCreateRequest
		j := helpers.GetRequestPayload(ctx)
		json.Unmarshal(j, &request)
		err := c.FlashSessionCreate(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func FlashSessionUpdateStatusHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.FlashSessionUpdateStatusRequest
		request.Id = ctx.Param("id")
		query := helpers.GetQueryParams(ctx)
		request.Status = query["status"]
		err := c.FlashSessionUpdateStatus(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func FlashSessionUpdateHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.FlashSessionUpdateRequest
		j := helpers.GetRequestPayload(ctx)
		json.Unmarshal(j, &request)
		err := c.FlashSessionUpdate(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func FlashSessionDeleteHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.FlashSessionDeleteRequest
		request.Id = ctx.Param("id")
		err := c.FlashSessionDelete(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}
