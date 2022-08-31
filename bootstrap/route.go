package bootstrap

import (
	"github.com/gin-gonic/gin"
	v1 "go-mall-api/api/admin/v1"
	"go-mall-api/controllers/admin"
	"go-mall-api/middlewares"
	"net/http"
	"strings"
)

// SetupRoute 路由初始化
func SetupRoute(router *gin.Engine) {
	//注册全局中间件
	registerGlobalMiddleWare(router)

	//注册Admin路由
	//routes.RegisterAdminRoutes(router)
	controller := admin.NewAdminController()
	v1.RegisterHTTPServer(router, controller)

	//配置404路由
	setup404Handler(router)
}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
		middlewares.ForceUA(),
		middlewares.Cors())
}

func setup404Handler(router *gin.Engine) {
	// 处理 404 请求
	router.NoRoute(func(c *gin.Context) {
		// 获取标头信息的Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			//如果是HTML
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			//默认返回 JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}
