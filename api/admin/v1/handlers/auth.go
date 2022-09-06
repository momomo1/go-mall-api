package handlers

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/pkg/response"
	requests "go-mall-api/requests/admin"
)

func LoginHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.LoginRequest
		if ok := requests.Validate(ctx, &request, requests.Login); !ok {
			return
		}

		h := func(ctx *gin.Context, req interface{}) (interface{}, error) {
			return c.Login(ctx, req.(*entity.LoginRequest))
		}
		out, err := h(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}

		reply := out.(*entity.LoginReply)
		response.OkWithData(ctx, *reply)
	}
}

func LogoutHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		err := c.Logout(ctx)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func InfoHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		out, _ := c.Info(ctx)
		response.OkWithData(ctx, out)
	}
}
