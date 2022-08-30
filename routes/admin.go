package routes

import (
	"github.com/gin-gonic/gin"
	controllers "go-mall-api/controllers/admin"
	"go-mall-api/controllers/admin/auth"
	"go-mall-api/middlewares"
)

// RegisterAdminRoutes 注册网页相关路由
func RegisterAdminRoutes(r *gin.Engine) {
	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	var admin *gin.RouterGroup
	admin = r.Group("/admin")

	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 API 每小时最多 60 个请求（根据 IP）。
	admin.Use(middlewares.LimitIP("208-H"))
	{
		lc := new(auth.LoginController)
		// 登录
		admin.POST("/login", lc.Login)
		admin.POST("/logout", lc.Logout)

		user := new(controllers.UsersController)
		// 获取当前用户
		admin.GET("/info", user.Info)
	}
}
