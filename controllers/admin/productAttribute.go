package admin

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/models/pms_product_attribute"
	"go-mall-api/models/pms_product_attribute_category"
	"strconv"
	"strings"
)

func (c AdminController) ProductAttribute(ctx *gin.Context, request *entity.ProductAttribute) (*entity.ProductAttributeList, error) {
	attributeModel := pms_product_attribute.Get(request.Id)

	return &entity.ProductAttributeList{
		Id:                         attributeModel.ID,
		FilterType:                 attributeModel.FilterType,
		HandAddStatus:              attributeModel.HandAddStatus,
		InputType:                  attributeModel.InputType,
		ProductAttributeCategoryId: attributeModel.ProductAttributeCategoryId,
		RelatedStatus:              attributeModel.RelatedStatus,
		SearchType:                 attributeModel.SearchType,
		SelectType:                 attributeModel.SelectType,
		Sort:                       attributeModel.Sort,
		Type:                       attributeModel.Type,
		InputList:                  attributeModel.InputList,
		Name:                       attributeModel.Name,
	}, nil
}

func (c AdminController) ProductAttributeAttrInfo(ctx *gin.Context) error {
	return nil
}

func (c AdminController) ProductAttributeList(ctx *gin.Context, request *entity.ProductAttributeListRequest) (*entity.ProductAttributeListReply, error) {
	where := map[string]interface{}{"type": request.Type, "product_attribute_category_id": request.Id}

	list, paging := pms_product_attribute.Paginate(ctx, request.PageSize, where, "sort", "desc")
	replyList := make([]entity.ProductAttributeList, 0, request.PageSize)

	for _, v := range list {
		interposition := entity.ProductAttributeList{
			Id:                         v.ID,
			FilterType:                 v.FilterType,
			HandAddStatus:              v.HandAddStatus,
			InputType:                  v.InputType,
			ProductAttributeCategoryId: v.ProductAttributeCategoryId,
			RelatedStatus:              v.RelatedStatus,
			SearchType:                 v.SearchType,
			SelectType:                 v.SelectType,
			Sort:                       v.Sort,
			Type:                       v.Type,
			InputList:                  v.InputList,
			Name:                       v.Name,
		}
		replyList = append(replyList, interposition)
	}

	return &entity.ProductAttributeListReply{
		List:        replyList,
		PagingAdmin: paging,
	}, nil
}

func (c AdminController) ProductAttributeCreate(ctx *gin.Context, request *entity.ProductAttributeCreateRequest) error {
	attributeModel := pms_product_attribute.PmsProductAttribute{
		ProductAttributeCategoryId: request.ProductAttributeCategoryId,
		Name:                       request.Name,
		SelectType:                 request.SelectType,
		InputType:                  request.InputType,
		InputList:                  request.InputList,
		Sort:                       request.Sort,
		FilterType:                 request.FilterType,
		SearchType:                 request.SearchType,
		RelatedStatus:              request.RelatedStatus,
		HandAddStatus:              request.HandAddStatus,
		Type:                       request.Type,
	}
	attributeCategoryModel := pms_product_attribute_category.Get(strconv.FormatInt(request.ProductAttributeCategoryId, 10))
	err := pms_product_attribute.AddAttribute(request.Type, attributeModel, attributeCategoryModel)
	return err
}

func (c AdminController) ProductAttributeUpdate(ctx *gin.Context, request *entity.ProductAttributeUpdateRequest) error {
	attributeModel := pms_product_attribute.Get(strconv.Itoa(request.Id))
	err := pms_product_attribute.UpdateAttribute(request, attributeModel)
	return err
}

func (c AdminController) ProductAttributeDelete(ctx *gin.Context, request *entity.ProductAttributeDeleteRequest) error {
	parts := strings.SplitN(request.Ids, ",", -1)
	err := pms_product_attribute.DeleteAttribute(parts)
	return err
}

func (c AdminController) ProductAttributeCategoryList(ctx *gin.Context, request *entity.ProductAttributeCategoryListRequest) (*entity.ProductAttributeCategoryListReply, error) {
	where := ""
	list, paging := pms_product_attribute_category.Paginate(ctx, request.PageSize, where, "id", "asc")
	replyList := make([]entity.ProductAttributeCategoryList, 0, request.PageSize)
	for _, v := range list {
		interposition := entity.ProductAttributeCategoryList{
			Id:             v.ID,
			Name:           v.Name,
			ParamCount:     v.ParamCount,
			AttributeCount: v.AttributeCount,
		}
		replyList = append(replyList, interposition)
	}

	return &entity.ProductAttributeCategoryListReply{
		List:        replyList,
		PagingAdmin: paging,
	}, nil
}

func (c AdminController) ProductAttributeCategoryListWithAttr(ctx *gin.Context) error {
	return nil
}

func (c AdminController) ProductAttributeCategoryCreate(ctx *gin.Context, request *entity.ProductAttributeCategoryCreateRequest) error {
	attributeCategoryModel := pms_product_attribute_category.PmsProductAttributeCategory{
		Name: request.Name,
	}
	attributeCategoryModel.Create()
	return nil
}

func (c AdminController) ProductAttributeCategoryUpdate(ctx *gin.Context, request *entity.ProductAttributeCategoryUpdateRequest) error {
	attributeCategoryModel := pms_product_attribute_category.Get(request.Id)
	attributeCategoryModel.Updates(map[string]interface{}{
		"name": request.Name,
	})
	return nil
}

func (c AdminController) ProductAttributeCategoryDelete(ctx *gin.Context, request *entity.ProductAttributeCategoryDeleteRequest) error {
	attributeCategoryModel := pms_product_attribute_category.Get(request.Id)
	attributeCategoryModel.Delete()
	return nil
}
