package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/pkg/helpers"
	"go-mall-api/pkg/response"
	"strconv"
)

func ProductCategoryHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ProductCategoryRequest
		request.Id = ctx.Param("id")
		out, _ := c.ProductCategory(ctx, &request)
		response.OkWithData(ctx, out)
	}
}

func ProductCategoryListHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ProductCategoryListRequest
		query := helpers.GetQueryParams(ctx)
		pageSize, _ := strconv.Atoi(query["pageSize"])
		request.PageSize = pageSize
		request.Id = ctx.Param("id")

		out, _ := c.ProductCategoryList(ctx, &request)
		response.OkWithData(ctx, out)
	}
}

func ProductCategoryListWithChildrenHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		out, _ := c.ProductCategoryListWithChildren(ctx)
		response.OkWithData(ctx, out)
	}
}

func ProductCategoryUpdateNavStatusHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ProductCategoryUpdateNavStatusRequest
		request.Ids = ctx.PostForm("ids")
		request.NavStatus = ctx.PostForm("navStatus")
		err := c.ProductCategoryUpdateNavStatus(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func ProductCategoryUpdateShowStatusHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ProductCategoryUpdateShowStatusRequest
		request.Ids = ctx.PostForm("ids")
		request.ShowStatus = ctx.PostForm("showStatus")
		err := c.ProductCategoryUpdateShowStatus(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func ProductCategoryCreateHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ProductCategoryCreateRequest
		j := helpers.GetRequestPayload(ctx)
		json.Unmarshal(j, &request)
		err := c.ProductCategoryCreate(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func ProductCategoryUpdateHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ProductCategoryUpdateRequest
		j := helpers.GetRequestPayload(ctx)
		json.Unmarshal(j, &request)
		err := c.ProductCategoryUpdate(ctx, &request)
		fmt.Println(err, ".........")
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func ProductCategoryDeleteHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ProductCategoryDeleteRequest
		request.Id = ctx.Param("id")

		err := c.ProductCategoryDelete(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}
