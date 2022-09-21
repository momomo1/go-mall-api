package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/pkg/helpers"
	"go-mall-api/pkg/response"
	"strconv"
)

func MenuTreeListHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		out, _ := c.MenuTreeList(ctx)
		response.OkWithData(ctx, out)
	}
}

func MenuListHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.MenuListRequest
		request.Id = ctx.Param("id")
		query := helpers.GetQueryParams(ctx)
		pageSize, _ := strconv.Atoi(query["pageSize"])
		request.PageSize = pageSize

		out, _ := c.MenuList(ctx, &request)
		response.OkWithData(ctx, out)
	}
}

func MenuUpdateHiddenHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.MenuUpdateHiddenRequest
		request.Id = ctx.Param("id")
		query := helpers.GetQueryParams(ctx)
		hidden, _ := strconv.Atoi(query["hidden"])
		request.Hidden = hidden

		err := c.MenuUpdateHidden(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func MenuCreateHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.MenuCreateRequest
		j := helpers.GetRequestPayload(ctx)
		json.Unmarshal(j, &request)

		err := c.MenuCreate(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func MenuHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.MenuRequest
		request.Id = ctx.Param("id")
		out, _ := c.Menu(ctx, &request)
		response.OkWithData(ctx, out.Data)
	}
}

func MenuUpdateHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.MenuUpdateRequest
		j := helpers.GetRequestPayload(ctx)
		json.Unmarshal(j, &request)
		err := c.MenuUpdate(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func MenuDeleteHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.MenuDeleteRequest
		request.Id = ctx.Param("id")
		err := c.MenuDelete(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}
