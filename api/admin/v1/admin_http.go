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

		//商品-商品模块
		admin.GET("/product/list", handlers.ProductListHttpHandler(c))
		admin.POST("/product/update/publishStatus", handlers.ProductUpdatePublishStatusHttpHandler(c))     // todo 未完成
		admin.POST("/product/update/newStatus", handlers.ProductUpdateNewStatusHttpHandler(c))             // todo 未完成
		admin.POST("/product/update/recommendStatus", handlers.ProductUpdateRecommendStatusHttpHandler(c)) // todo 未完成
		admin.OPTIONS("/product/update/deleteStatus", handlers.ProductUpdateDeleteStatusHttpHandler(c))    // todo 未完成
		admin.POST("/product/create", handlers.ProductCreateHttpHandler(c))                                // todo 未完成
		admin.POST("/product/updateInfo", handlers.ProductUpdateInfoHttpHandler(c))                        // todo 未完成
		admin.POST("/product/update", handlers.ProductUpdateHttpHandler(c))                                // todo 未完成
		admin.GET("/product/simpleList", handlers.ProductSimpleListHttpHandler(c))

		//商品-sku
		admin.GET("/sku/:id", Placeholder(c))     // todo 未完成
		admin.POST("/sku/update", Placeholder(c)) // todo 未完成

		//营销-品牌推荐
		admin.GET("/home/brand/list", handlers.HomeBrandListHttpHandler(c))
		admin.POST("/home/brand/update/recommendStatus", handlers.HomeBrandUpdateRecommendStatusHttpHandler(c))
		admin.POST("/home/brand/create", handlers.HomeBrandCreateHttpHandler(c))
		admin.POST("/home/brand/update/sort/:id", handlers.HomeBrandUpdateSortHttpHandler(c))
		admin.POST("/home/brand/delete", handlers.HomeBrandDeleteHttpHandler(c))

		//营销-营销专题推荐
		admin.GET("/home/recommendSubject/list", handlers.HomeRecommendSubjectListHttpHandler(c))
		admin.POST("/home/recommendSubject/update/recommendStatus", handlers.HomeRecommendSubjectUpdateRecommendStatusHttpHandler(c))
		admin.POST("/home/recommendSubject/create", handlers.HomeRecommendSubjectCreateHttpHandler(c))
		admin.POST("/home/recommendSubject/update/sort/:id", handlers.HomeRecommendSubjectUpdateSortHttpHandler(c))
		admin.POST("/home/recommendSubject/delete", handlers.HomeRecommendSubjectDeleteHttpHandler(c))

		//营销-营销广告列表
		admin.GET("/home/advertise/:id", handlers.HomeAdvertiseHttpHandler(c))
		admin.GET("/home/advertise/list", handlers.HomeAdvertiseListHttpHandler(c))
		admin.POST("/home/advertise/update/status/:id", handlers.HomeAdvertiseUpdateStatusHttpHandler(c))
		admin.POST("/home/advertise/create", handlers.HomeAdvertiseCreateHttpHandler(c))
		admin.POST("/home/advertise/update/:id", handlers.HomeAdvertiseUpdateHttpHandler(c))
		admin.POST("/home/advertise/delete", handlers.HomeAdvertiseDeleteHttpHandler(c))

		//营销-营销新品推荐
		admin.GET("/home/newProduct/list", handlers.HomeNewProductListHttpHandler(c))
		admin.POST("/home/newProduct/update/recommendStatus", handlers.HomeNewProductUpdateRecommendStatusHttpHandler(c))
		admin.POST("/home/newProduct/create", handlers.HomeNewProductCreateHttpHandler(c))
		admin.POST("/home/newProduct/update/sort/:id", handlers.HomeNewProductUpdateSortHttpHandler(c))
		admin.POST("/home/newProduct/delete", handlers.HomeNewProductDeleteHttpHandler(c))

		//营销-营销人气推荐
		admin.GET("/home/recommendProduct/list", handlers.HomeRecommendProductListHttpHandler(c))
		admin.POST("/home/recommendProduct/update/recommendStatus", handlers.HomeRecommendProductUpdateRecommendStatusHttpHandler(c))
		admin.POST("/home/recommendProduct/create", handlers.HomeRecommendProductCreateHttpHandler(c))
		admin.POST("/home/recommendProduct/update/sort/:id", handlers.HomeRecommendProductUpdateSortHttpHandler(c))
		admin.POST("/home/recommendProduct/delete", handlers.HomeRecommendProductDeleteHttpHandler(c))

		//营销-营销秒杀活动
		admin.GET("/flash/list", handlers.FlashListHttpHandler(c))
		admin.POST("/flash/update/status/:id", handlers.FlashUpdateStatusHttpHandler(c))
		admin.POST("/flash/create", handlers.FlashCreateHttpHandler(c))
		admin.POST("/flash/update/:id", handlers.FlashUpdateHttpHandler(c))
		admin.POST("/flash/delete/:id", handlers.FlashDeleteHttpHandler(c))
		admin.GET("/flashSession/list", handlers.FlashSessionListHttpHandler(c))
		admin.GET("/flashSession/selectList", handlers.FlashSessionSelectListHttpHandler(c))
		admin.POST("/flashSession/create", handlers.FlashSessionCreateHttpHandler(c))
		admin.POST("/flashSession/update/status/:id", handlers.FlashSessionUpdateStatusHttpHandler(c))
		admin.POST("/flashSession/update/:id", handlers.FlashSessionUpdateHttpHandler(c))
		admin.POST("/flashSession/delete/:id", handlers.FlashSessionDeleteHttpHandler(c))
		admin.GET("/flashProductRelation/list", handlers.FlashProductRelationListHttpHandler(c))          // todo 未完成
		admin.POST("/flashProductRelation/create", handlers.FlashProductRelationCreateHttpHandler(c))     // todo 未完成
		admin.POST("/flashProductRelation/update/:id", handlers.FlashProductRelationUpdateHttpHandler(c)) // todo 未完成
		admin.POST("/flashProductRelation/delete/:id", handlers.FlashProductRelationDeleteHttpHandler(c)) // todo 未完成

		//营销-营销优惠券列表
		admin.GET("/coupon/list", handlers.CouponListHttpHandler(c))
		admin.GET("/coupon/:id", handlers.CouponHttpHandler(c))
		admin.POST("/coupon/create", handlers.CouponCreateHttpHandler(c))
		admin.POST("/coupon/update/:id", handlers.CouponUpdateHttpHandler(c))
		admin.POST("/coupon/delete/:id", handlers.CouponDeleteHttpHandler(c))
		admin.GET("/couponHistory/list", handlers.CouponHistoryHttpHandler(c))

		//订单-订单
		admin.GET("/order/list", Placeholder(c))                 // todo 未完成
		admin.GET("/order/:id", Placeholder(c))                  // todo 未完成
		admin.POST("/order/update/note", Placeholder(c))         // todo 未完成
		admin.POST("/order/delete", Placeholder(c))              // todo 未完成
		admin.POST("/order/update/receiverInfo", Placeholder(c)) // todo 未完成
		admin.POST("/order/update/delivery", Placeholder(c))     // todo 未完成

		//订单-订单设置
		admin.GET("/orderSetting/:id", handlers.OrderSettingHttpHandler(c))
		admin.POST("/orderSetting/update/:id", handlers.OrderSettingUpdateHttpHandler(c))

		//订单-退货申请处理
		admin.GET("/returnApply/list", handlers.ReturnApplyListHttpHandler(c))
		admin.GET("/returnApply/:id", handlers.ReturnApplyHttpHandler(c))
		admin.POST("/returnApply/update/status/:id", handlers.ReturnApplyUpdateStatusHttpHandler(c)) // todo 未完成

		//订单-退货原因设置
		admin.GET("/returnReason/list", handlers.ReturnReasonListHttpHandler(c))
		admin.GET("/returnReason/:id", handlers.ReturnReasonHttpHandler(c))
		admin.POST("/returnReason/update/status", handlers.ReturnReasonUpdateStatusHttpHandler(c))
		admin.POST("/returnReason/create", handlers.ReturnReasonCreateHttpHandler(c))
		admin.POST("/returnReason/update/:id", handlers.ReturnReasonUpdateHttpHandler(c))
		admin.POST("/returnReason/delete", handlers.ReturnReasonDeleteHttpHandler(c))

		//其他
		admin.GET("/subject/list", handlers.SubjectListHttpHandler(c))
		admin.GET("/companyAddress/list", handlers.CompanyAddressListHttpHandler(c))
		admin.GET("/prefrenceArea/listAll", Placeholder(c)) // todo 未完成
		admin.GET("/memberLevel/list", Placeholder(c))      // todo 未完成
		admin.GET("/subject/listAll", Placeholder(c))       // todo 未完成
	}
}

//Placeholder 占位
func Placeholder(c handlers.AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		response.Ok(ctx)
	}
}
