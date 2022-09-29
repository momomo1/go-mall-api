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
		//用户信息
		admin.GET("admin/info", handlers.InfoHttpHandler(c))
		//上传图片
		admin.GET("aliyun/oss/policy", Placeholder(c))

		//权限-用户列表
		admin.GET("/admin/list", handlers.AdminListHttpHandler(c))
		admin.POST("/admin/updateStatus/:id", handlers.AdminUpdateStatusHttpHandler(c))
		admin.GET("/admin/role/:id", handlers.AdminRolesHttpHandler(c))
		admin.POST("/admin/register", handlers.AdminRegisterHttpHandler(c))
		admin.POST("/admin/role/update", handlers.AdminRoleUpdateHttpHandler(c))
		admin.POST("/admin/update/:id", handlers.AdminUpdateHttpHandler(c))
		admin.POST("/admin/delete/:id", handlers.AdminDeleteHttpHandler(c))

		//权限-角色
		admin.GET("/role/listAll", handlers.RoleListAllHttpHandler(c))
		admin.GET("/role/list", handlers.RoleListHttpHandler(c))
		admin.POST("/role/updateStatus/:id", handlers.RoleUpdateStatusHttpHandler(c))
		admin.POST("/role/create", handlers.RoleRegisterHttpHandler(c))
		admin.POST("/role/delete", handlers.RoleDeleteHttpHandler(c))
		admin.POST("/role/update/:id", handlers.RoleUpdateHttpHandler(c))
		admin.GET("/role/listMenu/:id", handlers.RoleListMenuHttpHandler(c))
		admin.POST("/role/allocMenu", handlers.RoleAllocMenuHttpHandler(c))
		admin.GET("/role/listResource/:id", handlers.RoleListResourceHttpHandler(c))
		admin.POST("/role/allocResource", handlers.RoleAllocResourceHttpHandler(c))

		//权限-菜单
		admin.GET("/menu/treeList", handlers.MenuTreeListHandler(c))
		admin.GET("/menu/list/:id", handlers.MenuListHttpHandler(c))
		admin.POST("/menu/updateHidden/:id", handlers.MenuUpdateHiddenHttpHandler(c))
		admin.POST("/menu/create", handlers.MenuCreateHttpHandler(c))
		admin.GET("/menu/:id", handlers.MenuHttpHandler(c))
		admin.POST("/menu/update/:id", handlers.MenuUpdateHandler(c))
		admin.POST("/menu/delete/:id", handlers.MenuDeleteHandler(c))

		//权限-资源
		admin.GET("/resource/listAll", handlers.ResourceListAllHttpHandler(c))
		admin.GET("/resource/list", handlers.ResourceListHttpHandler(c))
		admin.POST("/resource/create", handlers.ResourceCreateHttpHandler(c))
		admin.POST("/resource/update/:id", handlers.ResourceUpdateHttpHandler(c))
		admin.POST("/resource/delete/:id", handlers.ResourceDeleteHttpHandler(c))

		//权限-资源分类
		admin.GET("/resourceCategory/listAll", handlers.ResourceCategoryListAllHttpHandler(c))
		admin.POST("/resourceCategory/create", handlers.ResourceCategoryCreateHttpHandler(c))
		admin.POST("/resourceCategory/update/:id", handlers.ResourceCategoryUpdateHttpHandler(c))
		admin.POST("/resourceCategory/delete/:id", handlers.ResourceCategoryDeleteHttpHandler(c))

		//商品-品牌管理
		admin.GET("/brand/:id", handlers.BrandHttpHandler(c))
		admin.GET("/brand/list", handlers.BrandListHandler(c))
		admin.POST("/brand/create", handlers.BrandCreateHandler(c))
		admin.POST("/brand/update/:id", handlers.BrandUpdateHandler(c))
		admin.GET("/brand/delete/:id", handlers.BrandDeleteHandler(c))
		admin.POST("/brand/update/factoryStatus", handlers.BrandUpdateFactoryStatusHandler(c))
		admin.POST("/brand/update/showStatus", handlers.BrandUpdateShowStatusHttpHandler(c))

		//商品-商品类型
		admin.GET("/productAttribute/:id", handlers.ProductAttributeHttpHandler(c))
		admin.GET("/productAttribute/attrInfo/:id", handlers.ProductAttributeAttrInfoHttpHandler(c))
		admin.GET("/productAttribute/list/:id", handlers.ProductAttributeListHttpHandler(c))
		admin.POST("/productAttribute/create", handlers.ProductAttributeCreateHttpHandler(c))
		admin.POST("/productAttribute/update/:id", handlers.ProductAttributeUpdateHttpHandler(c))
		admin.POST("/productAttribute/delete", handlers.ProductAttributeDeleteHttpHandler(c))
		admin.GET("/productAttribute/category/list", handlers.ProductAttributeCategoryListHttpHandler(c))
		admin.GET("/productAttribute/category/list/withAttr", handlers.ProductAttributeCategoryListWithAttrHttpHandler(c))
		admin.POST("/productAttribute/category/create", handlers.ProductAttributeCategoryCreateHttpHandler(c))
		admin.POST("/productAttribute/category/update/:id", handlers.ProductAttributeCategoryUpdateHttpHandler(c))
		admin.GET("/productAttribute/category/delete/:id", handlers.ProductAttributeCategoryDeleteHttpHandler(c))

		//商品-商品分类
		admin.GET("/productCategory/:id", handlers.ProductCategoryHttpHandler(c))
		admin.GET("/productCategory/list/:id", handlers.ProductCategoryListHttpHandler(c))
		admin.GET("/productCategory/list/withChildren", handlers.ProductCategoryListWithChildrenHttpHandler(c))
		admin.POST("/productCategory/update/navStatus", handlers.ProductCategoryUpdateNavStatusHttpHandler(c))
		admin.POST("/productCategory/update/showStatus", handlers.ProductCategoryUpdateShowStatusHttpHandler(c))
		admin.POST("/productCategory/create", handlers.ProductCategoryCreateHttpHandler(c))
		admin.POST("/productCategory/update/:id", handlers.ProductCategoryUpdateHttpHandler(c))
		admin.POST("/productCategory/delete/:id", handlers.ProductCategoryDeleteHttpHandler(c))

		//待实现接口
		//商品
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
	}
}

//Placeholder 占位
func Placeholder(c handlers.AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		response.Ok(ctx)
	}
}
