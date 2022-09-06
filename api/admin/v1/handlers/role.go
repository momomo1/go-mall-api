package handlers

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/pkg/response"
)

func RoleListAllHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		out, _ := c.RoleListAll(ctx)
		response.OkWithData(ctx, out.Data)
	}
}
