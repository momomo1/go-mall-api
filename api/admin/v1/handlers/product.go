package handlers

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/pkg/helpers"
	"go-mall-api/pkg/response"
	"strconv"
)

func ProductListHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ProductListRequest
		query := helpers.GetQueryParams(ctx)
		pageSize, _ := strconv.Atoi(query["pageSize"])
		request.PageSize = pageSize
		request.Keyword = query["keyword"]
		request.PublishStatus = query["publishStatus"]
		request.VerifyStatus = query["verifyStatus"]
		request.ProductSn = query["productSn"]
		request.ProductCategoryId = query["productCategoryId"]
		request.BrandId = query["brandId"]

		out, _ := c.ProductList(ctx, &request)
		response.OkWithData(ctx, out)
	}
}

func ProductSimpleListHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ProductSimpleListRequest
		query := helpers.GetQueryParams(ctx)
		request.Keyword = query["keyword"]

		out, _ := c.ProductSimpleList(ctx, &request)
		response.OkWithData(ctx, out)
	}
}

func ProductUpdatePublishStatusHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ProductUpdatePublishStatusRequest
		query := helpers.GetQueryParams(ctx)
		request.Ids = query["ids"]
		request.PublishStatus = query["publishStatus"]
		err := c.ProductUpdatePublishStatus(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func ProductUpdateNewStatusHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ProductUpdateNewStatusRequest
		query := helpers.GetQueryParams(ctx)
		request.Ids = query["ids"]
		request.NewStatus = query["newStatus"]
		err := c.ProductUpdateNewStatus(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func ProductUpdateRecommendStatusHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ProductUpdateRecommendStatusRequest
		query := helpers.GetQueryParams(ctx)
		request.Ids = query["ids"]
		request.RecommendStatus = query["recommendStatus"]
		err := c.ProductUpdateRecommendStatus(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func ProductUpdateDeleteStatusHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ProductUpdateDeleteStatusRequest
		query := helpers.GetQueryParams(ctx)
		request.Ids = query["ids"]
		request.DeleteStatus = query["deleteStatus"]
		err := c.ProductUpdateDeleteStatus(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func ProductCreateHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		response.Ok(ctx)
	}
}

func ProductUpdateInfoHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		response.Ok(ctx)
	}
}

func ProductUpdateHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		response.Ok(ctx)
	}
}

func SkuHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.SkuRequest
		query := helpers.GetQueryParams(ctx)
		request.Keyword = query["keyword"]
		request.Id = ctx.Param("id")

		out, _ := c.Sku(ctx, &request)
		response.OkWithData(ctx, out)
	}
}

func SkuUpdateHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		response.Ok(ctx)
	}
}
