package v1

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/middlewares"
	"go-mall-api/pkg/response"
	requests "go-mall-api/requests/admin"
)

type adminController interface {
	Login(*gin.Context, *LoginRequest) (*LoginReply, error)
}

func RegisterHTTPServer(r *gin.Engine, c adminController) {
	var admin *gin.RouterGroup
	admin = r.Group("/admin")

	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 API 每小时最多 60 个请求（根据 IP）。
	admin.Use(middlewares.LimitIP("208-H"))
	{
		// 登录
		admin.POST("/login", LoginHttpHandler(c))
	}
}

func LoginHttpHandler(c adminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var in LoginRequest
		if ok := requests.Validate(ctx, &in, requests.Login); !ok {
			return
		}

		h := func(ctx *gin.Context, req interface{}) (interface{}, error) {
			return c.Login(ctx, req.(*LoginRequest))
		}
		out, err := h(ctx, &in)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}

		reply := out.(*LoginReply)
		response.OkWithData(ctx, *reply)
	}
}
