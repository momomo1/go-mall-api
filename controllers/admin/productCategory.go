package admin

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/models/pms_product_category"
	"go-mall-api/models/pms_product_category_attribute_relation"
	"strconv"
)

func (c AdminController) ProductCategory(ctx *gin.Context, request *entity.ProductCategoryRequest) (*entity.ProductCategoryList, error) {
	categoryModel := pms_product_category.Get(request.Id)
	return &entity.ProductCategoryList{
		Id:           categoryModel.ID,
		Level:        categoryModel.Level,
		ParentId:     categoryModel.ParentId,
		ProductCount: categoryModel.ProductCount,
		NavStatus:    categoryModel.NavStatus,
		ShowStatus:   categoryModel.ShowStatus,
		Sort:         categoryModel.Sort,
		Keywords:     categoryModel.Keywords,
		Name:         categoryModel.Name,
		ProductUnit:  categoryModel.ProductUnit,
		Description:  categoryModel.Description,
		Icon:         categoryModel.Icon,
	}, nil
}

func (c AdminController) ProductCategoryList(ctx *gin.Context, request *entity.ProductCategoryListRequest) (*entity.ProductCategoryListReply, error) {
	where := map[string]interface{}{"parent_id": request.Id}
	list, paging := pms_product_category.Paginate(ctx, request.PageSize, where, "sort", "desc")
	replyList := make([]entity.ProductCategoryList, 0, request.PageSize)
	for _, v := range list {
		interposition := entity.ProductCategoryList{
			Id:           v.ID,
			Level:        v.Level,
			ParentId:     v.ParentId,
			ProductCount: v.ProductCount,
			NavStatus:    v.NavStatus,
			ShowStatus:   v.ShowStatus,
			Sort:         v.Sort,
			Keywords:     v.Keywords,
			Name:         v.Name,
			ProductUnit:  v.ProductUnit,
			Description:  v.Description,
			Icon:         v.Icon,
		}
		replyList = append(replyList, interposition)
	}

	return &entity.ProductCategoryListReply{
		List:        replyList,
		PagingAdmin: paging,
	}, nil
}

func (c AdminController) ProductCategoryListWithChildren(ctx *gin.Context) ([]*entity.ProductCategoryListWithChildrenReply, error) {
	list := pms_product_category.All()
	treeMap := make(map[string]*entity.ProductCategoryListWithChildrenReply, len(list))
	for _, v := range list {
		id := strconv.Itoa(int(v.ID))
		treeMap[id] = &entity.ProductCategoryListWithChildrenReply{
			Id:       v.ID,
			Name:     v.Name,
			Children: []*entity.ProductCategoryListWithChildrenReply{},
		}
	}
	treeData := make([]*entity.ProductCategoryListWithChildrenReply, 0, len(list)) //全部的树状数据
	//无限极分类
	for _, v := range list {
		id := strconv.Itoa(int(v.ID))
		parentId := strconv.FormatInt(v.ParentId, 10)
		if v.ParentId == 0 {
			treeData = append(treeData, treeMap[id])
			continue
		}
		if k, ok := treeMap[parentId]; ok {
			k.Children = append(k.Children, treeMap[id])
		}
	}
	return treeData, nil
}

func (c AdminController) ProductCategoryUpdateNavStatus(ctx *gin.Context, request *entity.ProductCategoryUpdateNavStatusRequest) error {
	categoryModel := pms_product_category.Get(request.Ids)
	categoryModel.Updates(map[string]interface{}{
		"nav_status": request.NavStatus,
	})
	return nil
}

func (c AdminController) ProductCategoryUpdateShowStatus(ctx *gin.Context, request *entity.ProductCategoryUpdateShowStatusRequest) error {
	categoryModel := pms_product_category.Get(request.Ids)
	categoryModel.Updates(map[string]interface{}{
		"show_status": request.ShowStatus,
	})
	return nil
}

func (c AdminController) ProductCategoryCreate(ctx *gin.Context, request *entity.ProductCategoryCreateRequest) error {
	var level int
	if request.ParentId == 0 {
		level = 0
	} else {
		parentCategory := pms_product_category.Get(strconv.FormatInt(request.ParentId, 10))
		level = parentCategory.Level + 1
	}

	categoryModel := pms_product_category.PmsProductCategory{
		ParentId:     request.ParentId,
		Name:         request.Name,
		Level:        level,
		ProductCount: 0,
		ProductUnit:  request.ProductUnit,
		NavStatus:    request.NavStatus,
		ShowStatus:   request.ShowStatus,
		Sort:         request.Sort,
		Icon:         request.Icon,
		Keywords:     request.Keywords,
		Description:  request.Description,
	}
	categoryModel.Create()
	for _, v := range request.ProductAttributeIdList {
		relation := pms_product_category_attribute_relation.PmsProductCategoryAttributeRelation{
			ProductCategoryId:  int64(categoryModel.ID),
			ProductAttributeId: v,
		}
		relation.Create()
	}
	return nil
}

func (c AdminController) ProductCategoryUpdate(ctx *gin.Context, request *entity.ProductCategoryUpdateRequest) error {
	var level int
	if request.ParentId == 0 {
		level = 0
	} else {
		parentCategory := pms_product_category.Get(strconv.FormatInt(request.ParentId, 10))
		level = parentCategory.Level + 1
	}
	categoryModel := pms_product_category.Get(strconv.Itoa(request.Id))
	data := map[string]interface{}{
		"ParentId":    request.ParentId,
		"Name":        request.Name,
		"Level":       level,
		"ProductUnit": request.ProductUnit,
		"NavStatus":   request.NavStatus,
		"ShowStatus":  request.ShowStatus,
		"Sort":        request.Sort,
		"Icon":        request.Icon,
		"Keywords":    request.Keywords,
		"Description": request.Description,
	}
	err := pms_product_category.UpdateProductCategory(data, categoryModel, request.ProductAttributeIdList)
	return err
}

func (c AdminController) ProductCategoryDelete(ctx *gin.Context, request *entity.ProductCategoryDeleteRequest) error {
	categoryModel := pms_product_category.Get(request.Id)
	err := pms_product_category.DeleteRelevance(categoryModel)
	return err
}
