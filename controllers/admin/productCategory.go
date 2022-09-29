package admin

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/models/pms_product_category"
)

func (c AdminController) ProductCategory(ctx *gin.Context) error {
	return nil
}

func (c AdminController) ProductCategoryList(ctx *gin.Context, request *entity.ProductCategoryListRequest) (*entity.ProductCategoryListReply, error) {
	where := map[string]interface{}{"parent_id": request.Id}
	list, paging := pms_product_category.Paginate(ctx, request.PageSize, where, "sort", "desc")
	replyList := make([]entity.ProductCategoryListList, 0, request.PageSize)
	for _, v := range list {
		interposition := entity.ProductCategoryListList{
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

func (c AdminController) ProductCategoryListWithChildren(ctx *gin.Context) error {
	return nil
}

func (c AdminController) ProductCategoryUpdateNavStatus(ctx *gin.Context) error {
	return nil
}

func (c AdminController) ProductCategoryUpdateShowStatus(ctx *gin.Context) error {
	return nil
}

func (c AdminController) ProductCategoryCreate(ctx *gin.Context) error {
	return nil
}

func (c AdminController) ProductCategoryUpdate(ctx *gin.Context) error {
	return nil
}

func (c AdminController) ProductCategoryDelete(ctx *gin.Context) error {
	return nil
}
