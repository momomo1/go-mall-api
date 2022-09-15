package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/pkg/helpers"
	"go-mall-api/pkg/response"
	requests "go-mall-api/requests/admin"
	"strconv"
)

func RoleListAllHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		out, _ := c.RoleListAll(ctx)
		response.OkWithData(ctx, out.Data)
	}
}

func RoleListHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.RoleListRequest
		query := helpers.GetQueryParams(ctx)
		pageSize, _ := strconv.Atoi(query["pageSize"])
		keyword := query["keyword"]
		request.PageSize = pageSize
		request.Keyword = keyword

		out, _ := c.RoleList(ctx, &request)
		response.OkWithData(ctx, out)
	}
}

func RoleUpdateStatusHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.RoleUpdateStatusRequest
		id := ctx.Param("id")
		request.Id = id
		query := helpers.GetQueryParams(ctx)
		status, _ := strconv.Atoi(query["status"])
		request.Status = status

		err := c.RoleUpdateStatus(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func RoleRegisterHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.RoleRegisterRequest
		if ok := requests.Validate(ctx, &request, requests.RoleRegister); !ok {
			return
		}

		err := c.RoleRegister(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func RoleDeleteHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.RoleDeleteRequest
		request.Ids = ctx.PostForm("ids")

		err := c.RoleDelete(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func RoleUpdateHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.RoleUpdateRequest
		id := ctx.Param("id")
		request.Id = id
		j := helpers.GetRequestPayload(ctx)
		json.Unmarshal(j, &request)

		err := c.RoleUpdate(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}
