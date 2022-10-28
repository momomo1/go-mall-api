package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/models/pms_product"
	"go-mall-api/models/pms_sku_stock"
	"strings"
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

	whereAnd(&where)
	where += fmt.Sprintf("delete_status = '%d'", 0)

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

func (c AdminController) ProductSimpleList(ctx *gin.Context, request *entity.ProductSimpleListRequest) ([]*entity.ProductList, error) {
	where := ""
	if request.Keyword != "" {
		where += fmt.Sprintf("keywords LIKE %s", "'%"+request.Keyword+"%'")
		where += fmt.Sprintf("or name LIKE %s", "'%"+request.Keyword+"%'")
	}
	whereAnd(&where)
	where += fmt.Sprintf("delete_status = '%d'", 0)

	productData := pms_product.GetByWhereAll(where)
	replyList := make([]*entity.ProductList, 0, len(productData))
	for _, v := range productData {
		interposition := productFilling(v)
		replyList = append(replyList, &interposition)
	}
	return replyList, nil
}

func (c AdminController) ProductUpdatePublishStatus(ctx *gin.Context, request *entity.ProductUpdatePublishStatusRequest) error {
	ids := strings.SplitN(request.Ids, ",", -1)
	for _, v := range ids {
		productData := pms_product.Get(v)
		productData.Updates(map[string]interface{}{
			"PublishStatus": request.PublishStatus,
		})
	}
	return nil
}

func (c AdminController) ProductUpdateNewStatus(ctx *gin.Context, request *entity.ProductUpdateNewStatusRequest) error {
	ids := strings.SplitN(request.Ids, ",", -1)
	for _, v := range ids {
		productData := pms_product.Get(v)
		productData.Updates(map[string]interface{}{
			"NewStatus": request.NewStatus,
		})
	}
	return nil
}

func (c AdminController) ProductUpdateRecommendStatus(ctx *gin.Context, request *entity.ProductUpdateRecommendStatusRequest) error {
	ids := strings.SplitN(request.Ids, ",", -1)
	for _, v := range ids {
		productData := pms_product.Get(v)
		productData.Updates(map[string]interface{}{
			"RecommandStatus": request.RecommendStatus,
		})
	}
	return nil
}

func (c AdminController) ProductUpdateDeleteStatus(ctx *gin.Context, request *entity.ProductUpdateDeleteStatusRequest) error {
	ids := strings.SplitN(request.Ids, ",", -1)
	for _, v := range ids {
		productData := pms_product.Get(v)
		productData.Updates(map[string]interface{}{
			"DeleteStatus": request.DeleteStatus,
		})
	}
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

func (c AdminController) Sku(ctx *gin.Context, request *entity.SkuRequest) (*[]entity.SkuReply, error) {
	where := map[string]interface{}{
		"product_id": request.Id,
	}
	if request.Keyword != "" {
		where["sku_code"] = request.Keyword
	}
	skuData := pms_sku_stock.GetByProductIdAll(where)
	reply := make([]entity.SkuReply, 0, len(skuData))
	for _, v := range skuData {
		interposition := entity.SkuReply{
			Id:             v.ID,
			LockStock:      v.LockStock,
			Price:          v.Price,
			ProductId:      v.ProductId,
			Stock:          v.Stock,
			LowStock:       v.LowStock,
			Pic:            v.Pic,
			PromotionPrice: v.PromotionPrice,
			Sale:           v.Sale,
			SkuCode:        v.SkuCode,
			SpData:         v.SpData,
		}
		reply = append(reply, interposition)
	}
	return &reply, nil
}

func (c AdminController) SkuUpdate(ctx *gin.Context) error {
	return nil
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
