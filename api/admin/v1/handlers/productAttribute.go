package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/pkg/helpers"
	"go-mall-api/pkg/response"
	"strconv"
)

func ProductAttributeHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ProductAttribute
		request.Id = ctx.Param("id")
		out, _ := c.ProductAttribute(ctx, &request)
		response.OkWithData(ctx, out)
	}
}

func ProductAttributeAttrInfoHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		response.Ok(ctx)
	}
}

func ProductAttributeListHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ProductAttributeListRequest
		query := helpers.GetQueryParams(ctx)
		pageSize, _ := strconv.Atoi(query["pageSize"])
		AttributeType, _ := strconv.Atoi(query["type"])
		request.PageSize = pageSize
		request.Type = AttributeType
		request.Id = ctx.Param("id")

		out, _ := c.ProductAttributeList(ctx, &request)
		response.OkWithData(ctx, out)
	}
}

func ProductAttributeCreateHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ProductAttributeCreateRequest
		j := helpers.GetRequestPayload(ctx)
		json.Unmarshal(j, &request)

		err := c.ProductAttributeCreate(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func ProductAttributeUpdateHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ProductAttributeUpdateRequest
		j := helpers.GetRequestPayload(ctx)
		json.Unmarshal(j, &request)

		err := c.ProductAttributeUpdate(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func ProductAttributeDeleteHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ProductAttributeDeleteRequest
		request.Ids = ctx.PostForm("ids")
		err := c.ProductAttributeDelete(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func ProductAttributeCategoryListHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ProductAttributeCategoryListRequest
		query := helpers.GetQueryParams(ctx)
		pageSize, _ := strconv.Atoi(query["pageSize"])
		request.PageSize = pageSize

		out, _ := c.ProductAttributeCategoryList(ctx, &request)
		response.OkWithData(ctx, out)
	}
}

func ProductAttributeCategoryListWithAttrHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		response.Ok(ctx)
	}
}

func ProductAttributeCategoryCreateHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ProductAttributeCategoryCreateRequest
		request.Name = ctx.PostForm("name")
		err := c.ProductAttributeCategoryCreate(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func ProductAttributeCategoryUpdateHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ProductAttributeCategoryUpdateRequest
		request.Id = ctx.Param("id")
		request.Name = ctx.PostForm("name")
		err := c.ProductAttributeCategoryUpdate(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func ProductAttributeCategoryDeleteHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ProductAttributeCategoryDeleteRequest
		request.Id = ctx.Param("id")
		err := c.ProductAttributeCategoryDelete(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}
