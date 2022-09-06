package handlers

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/pkg/helpers"
	"go-mall-api/pkg/response"
	"strconv"
)

func AdminListHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ListRequest
		query := helpers.GetQueryParams(ctx)
		pageSize, _ := strconv.Atoi(query["pageSize"])
		request.PageSize = pageSize

		out, _ := c.AdminList(ctx, &request)
		response.OkWithData(ctx, out)
	}
}

func AdminUpdateStatusHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.AdminUpdateStatusRequest
		id := ctx.Param("id")
		request.Id = id
		query := helpers.GetQueryParams(ctx)
		status, _ := strconv.Atoi(query["status"])
		request.Status = status

		err := c.AdminUpdateStatus(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}
