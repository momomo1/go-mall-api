package handlers

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/pkg/helpers"
	"go-mall-api/pkg/response"
	"strconv"
)

func SubjectListHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.SubjectListRequest
		query := helpers.GetQueryParams(ctx)
		pageSize, _ := strconv.Atoi(query["pageSize"])
		request.PageSize = pageSize
		request.Keyword = query["keyword"]

		out, _ := c.SubjectList(ctx, &request)
		response.OkWithData(ctx, out)
	}
}
func CompanyAddressListHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		out, _ := c.CompanyAddressList(ctx)
		response.OkWithData(ctx, out.List)
	}
}
