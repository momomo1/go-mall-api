package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/pkg/helpers"
	"go-mall-api/pkg/response"
	"strconv"
)

func ResourceListAllHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		out, _ := c.ResourceListAll(ctx)
		response.OkWithData(ctx, out)
	}
}

func ResourceListHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ResourceListRequest
		query := helpers.GetQueryParams(ctx)
		pageSize, _ := strconv.Atoi(query["pageSize"])
		CategoryId, _ := strconv.Atoi(query["categoryId"])
		request.PageSize = pageSize
		request.CategoryId = CategoryId
		request.NameKeyword = query["nameKeyword"]
		request.UrlKeyword = query["urlKeyword"]

		out, _ := c.ResourceList(ctx, &request)
		response.OkWithData(ctx, out)
	}
}

func ResourceCreateHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ResourceCreateRequest
		j := helpers.GetRequestPayload(ctx)
		json.Unmarshal(j, &request)

		err := c.ResourceCreate(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func ResourceUpdateHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ResourceUpdateRequest
		j := helpers.GetRequestPayload(ctx)
		json.Unmarshal(j, &request)
		err := c.ResourceUpdate(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func ResourceDeleteHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ResourceDeleteRequest
		request.Id = ctx.Param("id")
		err := c.ResourceDelete(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func ResourceCategoryListAllHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		out, _ := c.ResourceCategoryListAll(ctx)
		response.OkWithData(ctx, out)
	}
}

func ResourceCategoryCreateHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ResourceCategoryCreateRequest
		j := helpers.GetRequestPayload(ctx)
		json.Unmarshal(j, &request)

		err := c.ResourceCategoryCreate(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func ResourceCategoryUpdateHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ResourceCategoryUpdateRequest
		j := helpers.GetRequestPayload(ctx)
		json.Unmarshal(j, &request)
		err := c.ResourceCategoryUpdate(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func ResourceCategoryDeleteHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ResourceCategoryDeleteRequest
		request.Id = ctx.Param("id")
		err := c.ResourceCategoryDelete(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}
