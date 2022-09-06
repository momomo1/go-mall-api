package v1

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/handlers"
	"go-mall-api/middlewares"
	"go-mall-api/pkg/response"
)

func RegisterHTTPServer(r *gin.Engine, c handlers.AdminController) {
	var admin *gin.RouterGroup
	admin = r.Group("")

	// 登录
	admin.POST("admin/login", middlewares.LimitIP("208-H"), handlers.LoginHttpHandler(c))
	// 退出登录
	admin.POST("admin/logout", middlewares.LimitIP("208-H"), handlers.LogoutHttpHandler(c))

	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 API 每小时最多 60 个请求（根据 IP）。
	admin.Use(middlewares.LimitIP("208-H"), middlewares.AuthJWTAdmin())
	{
		// 用户信息
		admin.GET("admin/info", handlers.InfoHttpHandler(c))

		// 用户列表
		admin.GET("/admin/list", handlers.AdminListHttpHandler(c))
		admin.POST("/admin/updateStatus/:id", handlers.AdminUpdateStatusHttpHandler(c))

		// 角色
		admin.GET("/role/listAll", handlers.RoleListAllHttpHandler(c))

		//待实现接口
		//商品
		admin.GET("/productCategory/:id", Placeholder(c))
		admin.GET("/productCategory/list/:id", Placeholder(c))
		admin.GET("/productCategory/list/withChildren", Placeholder(c))
		admin.POST("/productCategory/update/navStatus", Placeholder(c))
		admin.POST("/productCategory/update/showStatus", Placeholder(c))
		admin.POST("/productCategory/create", Placeholder(c))
		admin.POST("/productCategory/update/:id", Placeholder(c))
		admin.POST("/productCategory/delete/:id", Placeholder(c))

		admin.GET("/productAttribute/attrInfo/:id", Placeholder(c))
		admin.GET("/productAttribute/list/:id", Placeholder(c))
		admin.POST("/productAttribute/create", Placeholder(c))
		admin.POST("/productAttribute/update/:id", Placeholder(c))
		admin.POST("/productAttribute/delete", Placeholder(c))
		admin.GET("/productAttribute/category/list", Placeholder(c))
		admin.GET("/productAttribute/category/list/withAttr", Placeholder(c))
		admin.POST("/productAttribute/category/create", Placeholder(c))
		admin.POST("/productAttribute/category/update/:id", Placeholder(c))
		admin.GET("/productAttribute/category/delete/:id", Placeholder(c))

		admin.GET("/brand/:id", Placeholder(c))
		admin.GET("/brand/list", Placeholder(c))
		admin.POST("/brand/update/:id", Placeholder(c))
		admin.GET("/brand/delete/:id", Placeholder(c))
		admin.POST("/brand/update/factoryStatus", Placeholder(c))
		admin.POST("/brand/update/showStatus", Placeholder(c))

		admin.GET("/prefrenceArea/listAll", Placeholder(c))
		admin.GET("/memberLevel/list", Placeholder(c))
		admin.GET("/subject/list", Placeholder(c))
		admin.GET("/subject/listAll", Placeholder(c))

		admin.GET("/product/list", Placeholder(c))
		admin.POST("/product/create", Placeholder(c))
		admin.POST("/product/update/publishStatus", Placeholder(c))
		admin.POST("/product/update/newStatus", Placeholder(c))
		admin.POST("/product/update/recommendStatus", Placeholder(c))
		admin.POST("/product/updateInfo", Placeholder(c))
		admin.POST("/product/update", Placeholder(c))
		admin.OPTIONS("/product/update/deleteStatus", Placeholder(c))

		admin.GET("/sku/:id", Placeholder(c))
		admin.POST("/sku/update", Placeholder(c))

		//订单
		admin.GET("/order/list", Placeholder(c))
		admin.GET("/order/:id", Placeholder(c))
		admin.POST("/order/update/note", Placeholder(c))
		admin.POST("/order/delete", Placeholder(c))
		admin.POST("/order/update/receiverInfo", Placeholder(c))
		admin.POST("/order/update/delivery", Placeholder(c))

		admin.GET("/orderSetting/:id", Placeholder(c))
		admin.POST("/orderSetting/update/:id", Placeholder(c))

		admin.GET("/returnApply/list", Placeholder(c))
		admin.GET("/returnApply/:id", Placeholder(c))
		admin.POST("/returnApply/update/status/:id", Placeholder(c))

		admin.GET("/companyAddress/list", Placeholder(c))
		admin.GET("/returnReason/list", Placeholder(c))
		admin.GET("/returnReason/:id", Placeholder(c))
		admin.POST("/returnReason/update/status", Placeholder(c))
		admin.POST("/returnReason/create", Placeholder(c))
		admin.POST("/returnReason/update/:id", Placeholder(c))
		admin.POST("/returnReason/delete", Placeholder(c))

		//营销活动
		admin.GET("/flash/list", Placeholder(c))
		admin.POST("/flash/update/status/:id", Placeholder(c))
		admin.POST("/flash/create", Placeholder(c))
		admin.POST("/flash/update/:id", Placeholder(c))
		admin.POST("/flash/delete/:id", Placeholder(c))

		admin.GET("/flashSession/list", Placeholder(c))
		admin.GET("/flashSession/selectList", Placeholder(c))
		admin.POST("/flashSession/create", Placeholder(c))
		admin.POST("/flashSession/update/status/:id", Placeholder(c))
		admin.POST("/flashSession/update/:id", Placeholder(c))
		admin.POST("/flashSession/delete/:id", Placeholder(c))

		admin.GET("/flashProductRelation/list", Placeholder(c))
		admin.POST("/flashProductRelation/create", Placeholder(c))
		admin.POST("/flashProductRelation/update/:id", Placeholder(c))
		admin.POST("/flashProductRelation/delete/:id", Placeholder(c))

		admin.GET("/coupon/list", Placeholder(c))
		admin.GET("/coupon/:id", Placeholder(c))
		admin.POST("/coupon/create", Placeholder(c))
		admin.POST("/coupon/update/:id", Placeholder(c))
		admin.POST("/coupon/delete/:id", Placeholder(c))
		admin.GET("/couponHistory/list", Placeholder(c))

		admin.GET("/home/brand/list", Placeholder(c))
		admin.GET("/home/brand/update/recommendStatus", Placeholder(c))
		admin.POST("/home/brand/create", Placeholder(c))
		admin.POST("/home/brand/update/sort/:id", Placeholder(c))
		admin.POST("/home/brand/delete", Placeholder(c))

		admin.GET("/home/newProduct/list", Placeholder(c))
		admin.POST("/home/newProduct/update/recommendStatus", Placeholder(c))
		admin.POST("/home/newProduct/create", Placeholder(c))
		admin.POST("/home/newProduct/update/sort/:id", Placeholder(c))
		admin.POST("/home/newProduct/delete", Placeholder(c))

		admin.GET("/home/recommendProduct/list", Placeholder(c))
		admin.POST("/home/recommendProduct/update/recommendStatus", Placeholder(c))
		admin.POST("/home/recommendProduct/create", Placeholder(c))
		admin.POST("/home/recommendProduct/update/sort/:id", Placeholder(c))
		admin.POST("/home/recommendProduct/delete", Placeholder(c))

		admin.GET("/home/recommendSubject/list", Placeholder(c))
		admin.POST("/home/recommendSubject/update/recommendStatus", Placeholder(c))
		admin.POST("/home/recommendSubject/create", Placeholder(c))
		admin.POST("/home/recommendSubject/update/sort/:id", Placeholder(c))
		admin.POST("/home/recommendSubject/delete", Placeholder(c))

		admin.GET("/home/advertise/:id", Placeholder(c))
		admin.GET("/home/advertise/list", Placeholder(c))
		admin.POST("/home/advertise/update/status/:id", Placeholder(c))
		admin.POST("/home/advertise/create", Placeholder(c))
		admin.POST("/home/advertise/update/:id", Placeholder(c))
		admin.POST("/home/advertise/delete", Placeholder(c))

		admin.POST("/admin/register", Placeholder(c))
		admin.GET("/admin/role/:id", Placeholder(c))
		admin.POST("/admin/role/update", Placeholder(c))
		admin.POST("/admin/update/:id", Placeholder(c))
		admin.POST("/admin/delete/:id", Placeholder(c))

		admin.GET("/role/list", Placeholder(c))
		admin.POST("/role/create", Placeholder(c))
		admin.GET("/role/listMenu/:id", Placeholder(c))
		admin.POST("/role/allocMenu", Placeholder(c))
		admin.GET("/role/listResource/:id", Placeholder(c))
		admin.POST("/role/allocResource", Placeholder(c))
		admin.POST("/role/update/:id", Placeholder(c))
		admin.POST("/role/delete", Placeholder(c))

		admin.GET("/menu/treeList", Placeholder(c))
		admin.GET("/menu/list/:id", Placeholder(c))
		admin.GET("/menu/:id", Placeholder(c))
		admin.POST("/menu/updateHidden/:id", Placeholder(c))
		admin.POST("/menu/create", Placeholder(c))
		admin.POST("/menu/update/:id", Placeholder(c))
		admin.POST("/menu/delete/:id", Placeholder(c))

		admin.GET("/resource/listAll", Placeholder(c))
		admin.GET("/resource/list", Placeholder(c))
		admin.POST("/resource/create", Placeholder(c))
		admin.POST("/resource/update/:id", Placeholder(c))
		admin.POST("/resource/delete/:id", Placeholder(c))

		admin.GET("/resourceCategory/listAll", Placeholder(c))
		admin.POST("/resourceCategory/create", Placeholder(c))
		admin.POST("/resourceCategory/update/:id", Placeholder(c))
		admin.POST("/resourceCategory/delete/:id", Placeholder(c))
	}
}

//Placeholder 占位
func Placeholder(c handlers.AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		response.Ok(ctx)
	}
}
