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

func AdminListHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.ListRequest
		query := helpers.GetQueryParams(ctx)
		pageSize, _ := strconv.Atoi(query["pageSize"])
		keyword := query["keyword"]
		request.PageSize = pageSize
		request.Keyword = keyword

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

func AdminRegisterHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.AdminRegisterRequest
		if ok := requests.Validate(ctx, &request, requests.Register); !ok {
			return
		}

		err := c.AdminRegister(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func AdminRoleUpdateHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.AdminRoleUpdateRequest
		request.AdminId = ctx.PostForm("adminId")
		request.RoleIds = ctx.PostForm("roleIds")
		if ok := requests.ValidateSimple(ctx, &request, requests.AdminRoleUpdate); !ok {
			return
		}
		err := c.AdminRoleUpdate(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func AdminRolesHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.AdminRolesRequest
		id := ctx.Param("id")
		request.Id = id
		roles, err := c.AdminRoles(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.OkWithData(ctx, roles.Data)
	}
}

func AdminUpdateHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.AdminUpdateRequest
		j := helpers.GetRequestPayload(ctx)
		json.Unmarshal(j, &request)

		err := c.AdminUpdate(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func AdminDeleteHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.AdminDeleteRequest
		id := ctx.Param("id")
		request.Id = id
		err := c.AdminDelete(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}
