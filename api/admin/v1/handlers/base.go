package handlers

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
)

type AdminController interface {
	Login(*gin.Context, *entity.LoginRequest) (*entity.LoginReply, error)
	Logout(*gin.Context) error
	Info(*gin.Context) (*entity.UserReply, error)

	AdminList(*gin.Context, *entity.ListRequest) (*entity.ListReply, error)
	AdminUpdateStatus(*gin.Context, *entity.AdminUpdateStatusRequest) error
	AdminRegister(*gin.Context, *entity.AdminRegisterRequest) error
	AdminRoleUpdate(*gin.Context, *entity.AdminRoleUpdateRequest) error
	AdminRoles(*gin.Context, *entity.AdminRolesRequest) (*entity.AdminRolesReply, error)
	AdminUpdate(*gin.Context, *entity.AdminUpdateRequest) error
	AdminDelete(*gin.Context, *entity.AdminDeleteRequest) error

	RoleListAll(*gin.Context) (*entity.RoleListAllReply, error)
	RoleList(*gin.Context, *entity.RoleListRequest) (*entity.RoleListReply, error)
	RoleUpdateStatus(*gin.Context, *entity.RoleUpdateStatusRequest) error
	RoleRegister(*gin.Context, *entity.RoleRegisterRequest) error
	RoleDelete(*gin.Context, *entity.RoleDeleteRequest) error
	RoleUpdate(*gin.Context, *entity.RoleUpdateRequest) error
	RoleListMenu(*gin.Context, *entity.RoleListMenuRequest) ([]*entity.RoleListMenuReply, error)
	RoleAllocMenu(*gin.Context, *entity.RoleAllocMenuRequest) error
	RoleListResource(*gin.Context, *entity.RoleListResourceRequest) ([]*entity.RoleListResourceReply, error)
	RoleAllocResource(*gin.Context, *entity.RoleAllocResourceRequest) error

	MenuTreeList(*gin.Context) ([]*entity.MenuTreeListReply, error)
	MenuList(*gin.Context, *entity.MenuListRequest) (*entity.MenuListReply, error)
	MenuUpdateHidden(*gin.Context, *entity.MenuUpdateHiddenRequest) error
	MenuCreate(*gin.Context, *entity.MenuCreateRequest) error
	Menu(*gin.Context, *entity.MenuRequest) (*entity.MenuReply, error)
	MenuUpdate(*gin.Context, *entity.MenuUpdateRequest) error
	MenuDelete(*gin.Context, *entity.MenuDeleteRequest) error

	ResourceListAll(*gin.Context) ([]*entity.ResourceList, error)
	ResourceList(*gin.Context, *entity.ResourceListRequest) (*entity.ResourceListReply, error)
	ResourceCreate(*gin.Context, *entity.ResourceCreateRequest) error
	ResourceUpdate(*gin.Context, *entity.ResourceUpdateRequest) error
	ResourceDelete(*gin.Context, *entity.ResourceDeleteRequest) error

	ResourceCategoryListAll(*gin.Context) ([]*entity.ResourceCategoryListAllReply, error)
	ResourceCategoryCreate(*gin.Context, *entity.ResourceCategoryCreateRequest) error
	ResourceCategoryUpdate(*gin.Context, *entity.ResourceCategoryUpdateRequest) error
	ResourceCategoryDelete(*gin.Context, *entity.ResourceCategoryDeleteRequest) error

	Brand(*gin.Context, *entity.BrandRequest) (*entity.BrandList, error)
	BrandList(*gin.Context, *entity.BrandListRequest) (*entity.BrandListReply, error)
	BrandCreate(*gin.Context, *entity.BrandCreateRequest) error
	BrandUpdate(*gin.Context, *entity.BrandUpdateRequest) error
	BrandDelete(*gin.Context, *entity.BrandDeleteRequest) error
	BrandUpdateFactoryStatus(*gin.Context, *entity.BrandUpdateFactoryStatusRequest) error
	BrandUpdateShowStatus(*gin.Context, *entity.BrandUpdateShowStatusRequest) error

	ProductAttribute(*gin.Context, *entity.ProductAttributeRequest) (*entity.ProductAttributeList, error)
	ProductAttributeAttrInfo(*gin.Context, *entity.ProductAttributeAttrInfoRequest) ([]*entity.ProductAttributeAttrInfoReply, error)
	ProductAttributeList(*gin.Context, *entity.ProductAttributeListRequest) (*entity.ProductAttributeListReply, error)
	ProductAttributeCreate(*gin.Context, *entity.ProductAttributeCreateRequest) error
	ProductAttributeUpdate(*gin.Context, *entity.ProductAttributeUpdateRequest) error
	ProductAttributeDelete(*gin.Context, *entity.ProductAttributeDeleteRequest) error
	ProductAttributeCategoryList(*gin.Context, *entity.ProductAttributeCategoryListRequest) (*entity.ProductAttributeCategoryListReply, error)
	ProductAttributeCategoryListWithAttr(*gin.Context) ([]*entity.ProductAttributeCategoryListWithAttrReply, error)
	ProductAttributeCategoryCreate(*gin.Context, *entity.ProductAttributeCategoryCreateRequest) error
	ProductAttributeCategoryUpdate(*gin.Context, *entity.ProductAttributeCategoryUpdateRequest) error
	ProductAttributeCategoryDelete(*gin.Context, *entity.ProductAttributeCategoryDeleteRequest) error

	ProductCategory(*gin.Context, *entity.ProductCategoryRequest) (*entity.ProductCategoryList, error)
	ProductCategoryList(*gin.Context, *entity.ProductCategoryListRequest) (*entity.ProductCategoryListReply, error)
	ProductCategoryListWithChildren(*gin.Context) ([]*entity.ProductCategoryListWithChildrenReply, error)
	ProductCategoryUpdateNavStatus(*gin.Context, *entity.ProductCategoryUpdateNavStatusRequest) error
	ProductCategoryUpdateShowStatus(*gin.Context, *entity.ProductCategoryUpdateShowStatusRequest) error
	ProductCategoryCreate(*gin.Context, *entity.ProductCategoryCreateRequest) error
	ProductCategoryUpdate(*gin.Context, *entity.ProductCategoryUpdateRequest) error
	ProductCategoryDelete(*gin.Context, *entity.ProductCategoryDeleteRequest) error

	ProductList(*gin.Context, *entity.ProductListRequest) (*entity.ProductListReply, error)
	ProductUpdatePublishStatus(*gin.Context) error
	ProductUpdateNewStatus(*gin.Context) error
	ProductUpdateRecommendStatus(*gin.Context) error
	ProductUpdateDeleteStatus(*gin.Context) error
	ProductCreate(*gin.Context) error
	ProductUpdateInfo(*gin.Context) error
	ProductUpdate(*gin.Context) error
	ProductSimpleList(*gin.Context, *entity.ProductSimpleListRequest) ([]*entity.ProductList, error)
	HomeBrandList(*gin.Context, *entity.HomeBrandListRequest) (*entity.HomeBrandListReply, error)
	HomeBrandUpdateRecommendStatus(*gin.Context, *entity.HomeBrandUpdateRecommendStatusRequest) error
	HomeBrandCreate(*gin.Context, *[]entity.HomeBrandCreateRequest) error
	HomeBrandUpdateSort(*gin.Context, *entity.HomeBrandUpdateSortRequest) error
	HomeBrandDelete(*gin.Context, *entity.HomeBrandDeleteRequest) error

	HomeRecommendSubjectList(*gin.Context, *entity.HomeRecommendSubjectListRequest) (*entity.HomeRecommendSubjectListReply, error)
	HomeRecommendSubjectUpdateRecommendStatus(*gin.Context, *entity.HomeRecommendSubjectUpdateRecommendStatusRequest) error
	HomeRecommendSubjectCreate(*gin.Context, *[]entity.HomeRecommendSubjectCreateRequest) error
	HomeRecommendSubjectUpdateSort(*gin.Context, *entity.HomeRecommendSubjectUpdateSortRequest) error
	HomeRecommendSubjectDelete(*gin.Context, *entity.HomeRecommendSubjectDeleteRequest) error

	HomeAdvertise(*gin.Context, *entity.HomeAdvertiseRequest) (*entity.HomeAdvertiseList, error)
	HomeAdvertiseList(*gin.Context, *entity.HomeAdvertiseListRequest) (*entity.HomeAdvertiseListReply, error)
	HomeAdvertiseUpdateStatus(*gin.Context, *entity.HomeAdvertiseUpdateStatusRequest) error
	HomeAdvertiseCreate(*gin.Context, *entity.HomeAdvertiseCreateRequest) error
	HomeAdvertiseUpdate(*gin.Context, *entity.HomeAdvertiseUpdateRequest) error
	HomeAdvertiseDelete(*gin.Context, *entity.HomeAdvertiseDeleteRequest) error

	HomeNewProductList(*gin.Context, *entity.HomeNewProductListRequest) (*entity.HomeNewProductListReply, error)
	HomeNewProductUpdateRecommendStatus(*gin.Context, *entity.HomeNewProductUpdateRecommendStatusRequest) error
	HomeNewProductCreate(*gin.Context, *[]entity.HomeNewProductCreateRequest) error
	HomeNewProductUpdateSort(*gin.Context, *entity.HomeNewProductUpdateSortRequest) error
	HomeNewProductDelete(*gin.Context, *entity.HomeNewProductDeleteRequest) error

	HomeRecommendProductList(*gin.Context, *entity.HomeRecommendProductListRequest) (*entity.HomeRecommendProductListReply, error)
	HomeRecommendProductUpdateRecommendStatus(*gin.Context, *entity.HomeRecommendProductUpdateRecommendStatusRequest) error
	HomeRecommendProductCreate(*gin.Context, *[]entity.HomeRecommendProductCreateRequest) error
	HomeRecommendProductUpdateSort(*gin.Context, *entity.HomeRecommendProductUpdateSortRequest) error
	HomeRecommendProductDelete(*gin.Context, *entity.HomeRecommendProductDeleteRequest) error

	FlashList(*gin.Context, *entity.FlashListRequest) (*entity.FlashListReply, error)
	FlashUpdateStatus(*gin.Context, *entity.FlashUpdateStatusRequest) error
	FlashCreate(*gin.Context, *entity.FlashCreateRequest) error
	FlashUpdate(*gin.Context, *entity.FlashUpdateRequest) error
	FlashDelete(*gin.Context, *entity.FlashDeleteRequest) error

	FlashSessionList(*gin.Context) (*entity.FlashSessionListReply, error)
	FlashSessionSelectList(*gin.Context, *entity.FlashSessionSelectListRequest) (*entity.FlashSessionSelectListReply, error)
	FlashSessionCreate(*gin.Context, *entity.FlashSessionCreateRequest) error
	FlashSessionUpdateStatus(*gin.Context, *entity.FlashSessionUpdateStatusRequest) error
	FlashSessionUpdate(*gin.Context, *entity.FlashSessionUpdateRequest) error
	FlashSessionDelete(*gin.Context, *entity.FlashSessionDeleteRequest) error

	FlashProductRelationList(*gin.Context, *entity.FlashProductRelationListRequest) (*entity.FlashProductRelationListReply, error)
	FlashProductRelationCreate(*gin.Context, *[]entity.FlashProductRelationCreateRequest) error
	FlashProductRelationUpdate(*gin.Context, *entity.FlashProductRelationUpdateRequest) error
	FlashProductRelationDelete(*gin.Context, *entity.FlashProductRelationDeleteRequest) error

	CouponList(*gin.Context, *entity.CouponListRequest) (*entity.CouponListReply, error)
	Coupon(*gin.Context, *entity.CouponRequest) (*entity.CouponList, error)
	CouponCreate(*gin.Context, *entity.CouponCreateRequest) error
	CouponUpdate(*gin.Context, *entity.CouponUpdateRequest) error
	CouponDelete(*gin.Context, *entity.CouponDeleteRequest) error
	CouponHistory(*gin.Context, *entity.CouponHistoryRequest) (*entity.CouponHistoryReply, error)

	OrderList(*gin.Context, *entity.OrderListRequest) (*entity.OrderListReply, error)
	Order(*gin.Context, *entity.OrderRequest) (*entity.OrderList, error)
	OrderUpdateNote(*gin.Context, *entity.OrderUpdateNoteRequest) error
	OrderDelete(*gin.Context, *entity.OrderDeleteRequest) error
	OrderUpdateReceiverInfo(*gin.Context, *entity.OrderUpdateReceiverInfoRequest) error
	OrderUpdateMoneyInfo(*gin.Context) error
	OrderUpdateDelivery(*gin.Context, *[]entity.OrderUpdateDeliveryRequest) error
	OrderUpdateClose(*gin.Context, *entity.OrderUpdateCloseRequest) error

	OrderSetting(*gin.Context) (*entity.OrderSettingReply, error)
	OrderSettingUpdate(*gin.Context, *entity.OrderSettingRequest) error

	ReturnApplyList(*gin.Context, *entity.ReturnApplyListRequest) (*entity.ReturnApplyListReply, error)
	ReturnApply(*gin.Context, *entity.ReturnApplyRequest) (*entity.ReturnApplyReply, error)
	ReturnApplyUpdateStatus(*gin.Context) error

	ReturnReasonList(*gin.Context, *entity.ReturnReasonListRequest) (*entity.ReturnReasonListReply, error)
	ReturnReason(*gin.Context, *entity.ReturnReasonRequest) (*entity.ReturnReasonList, error)
	ReturnReasonUpdateStatus(*gin.Context, *entity.ReturnReasonUpdateStatusRequest) error
	ReturnReasonCreate(*gin.Context, *entity.ReturnReasonCreateRequest) error
	ReturnReasonUpdate(*gin.Context, *entity.ReturnReasonUpdateRequest) error
	ReturnReasonDelete(*gin.Context, *entity.ReturnReasonDeleteRequest) error

	SubjectList(*gin.Context, *entity.SubjectListRequest) (*entity.SubjectListReply, error)
	CompanyAddressList(*gin.Context) (*entity.CompanyAddressListReply, error)
}
