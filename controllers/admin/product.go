package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/models/pms_product"
)

func (c AdminController) ProductList(ctx *gin.Context, request *entity.ProductListRequest) (*entity.ProductListReply, error) {
	where := ""
	if request.Keyword != "" {
		whereAnd(&where)
		where += fmt.Sprintf("keywords LIKE %s", "'%"+request.Keyword+"%'")
	}
	if request.ProductSn != "" {
		whereAnd(&where)
		where += fmt.Sprintf("product_sn LIKE %s", "'%"+request.ProductSn+"%'")
	}
	if request.PublishStatus != "" {
		whereAnd(&where)
		where += fmt.Sprintf("publish_status = '%s'", request.PublishStatus)
	}
	if request.VerifyStatus != "" {
		whereAnd(&where)
		where += fmt.Sprintf("verify_status = '%s'", request.VerifyStatus)
	}
	if request.ProductCategoryId != "" {
		whereAnd(&where)
		where += fmt.Sprintf("product_category_id = '%s'", request.ProductCategoryId)
	}
	if request.BrandId != "" {
		whereAnd(&where)
		where += fmt.Sprintf("brand_id = '%s'", request.BrandId)
	}

	list, paging := pms_product.Paginate(ctx, request.PageSize, where, "sort", "desc")
	replyList := make([]entity.ProductList, 0, request.PageSize)
	for _, v := range list {
		interposition := productFilling(v)
		replyList = append(replyList, interposition)
	}

	return &entity.ProductListReply{
		List:        replyList,
		PagingAdmin: paging,
	}, nil
}

func (c AdminController) ProductUpdatePublishStatus(ctx *gin.Context) error {
	return nil
}

func (c AdminController) ProductUpdateNewStatus(ctx *gin.Context) error {
	return nil
}

func (c AdminController) ProductUpdateRecommendStatus(ctx *gin.Context) error {
	return nil
}

func (c AdminController) ProductUpdateDeleteStatus(ctx *gin.Context) error {
	return nil
}

func (c AdminController) ProductCreate(ctx *gin.Context) error {
	return nil
}

func (c AdminController) ProductUpdateInfo(ctx *gin.Context) error {
	return nil
}

func (c AdminController) ProductUpdate(ctx *gin.Context) error {
	return nil
}

func (c AdminController) ProductSimpleList(ctx *gin.Context, request *entity.ProductSimpleListRequest) ([]*entity.ProductList, error) {
	where := ""
	if request.Keyword != "" {
		where += fmt.Sprintf("keywords LIKE %s", "'%"+request.Keyword+"%'")
		where += fmt.Sprintf("or name LIKE %s", "'%"+request.Keyword+"%'")
	}
	productData := pms_product.GetByWhereAll(where)
	replyList := make([]*entity.ProductList, 0, len(productData))
	for _, v := range productData {
		interposition := productFilling(v)
		replyList = append(replyList, &interposition)
	}
	return replyList, nil
}

func productFilling(v pms_product.PmsProduct) entity.ProductList {
	interposition := entity.ProductList{
		Id:                         v.ID,
		BrandId:                    v.BrandId,
		FeightTemplateId:           v.FeightTemplateId,
		ProductCategoryId:          v.ProductCategoryId,
		ProductAttributeCategoryId: v.ProductAttributeCategoryId,
		GiftGrowth:                 v.GiftGrowth,
		GiftPoint:                  v.GiftPoint,
		LowStock:                   v.LowStock,
		Sale:                       v.Sale,
		Sort:                       v.Sort,
		Stock:                      v.Stock,
		Weight:                     v.Weight,
		PromotionType:              v.PromotionType,
		NewStatus:                  v.NewStatus,
		PreviewStatus:              v.PreviewStatus,
		DeleteStatus:               v.DeleteStatus,
		PublishStatus:              v.PublishStatus,
		VerifyStatus:               v.VerifyStatus,
		RecommandStatus:            v.RecommandStatus,
		UsePointLimit:              v.UsePointLimit,
		PromotionPerLimit:          v.PromotionPerLimit,
		OriginalPrice:              v.OriginalPrice,
		Price:                      v.Price,
		AlbumPics:                  v.AlbumPics,
		BrandName:                  v.BrandName,
		Description:                v.Description,
		DetailDesc:                 v.DetailDesc,
		DetailHtml:                 v.DetailHtml,
		DetailMobileHtml:           v.DetailMobileHtml,
		DetailTitle:                v.DetailTitle,
		Keywords:                   v.Keywords,
		Name:                       v.Name,
		Note:                       v.Note,
		Pic:                        v.Pic,
		ProductCategoryName:        v.ProductCategoryName,
		ProductSn:                  v.ProductSn,
		ServiceIds:                 v.ServiceIds,
		SubTitle:                   v.SubTitle,
		Unit:                       v.Unit,
		PromotionPrice:             v.PromotionPrice,
		PromotionStartTime:         v.PromotionStartTime.Format("2006-01-02 15:04:05"),
		PromotionEndTime:           v.PromotionEndTime.Format("2006-01-02 15:04:05"),
	}

	return interposition
}
