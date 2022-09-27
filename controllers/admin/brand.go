package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/models/pms_brand"
	"strconv"
	"strings"
)

func (c AdminController) Brand(ctx *gin.Context, request *entity.BrandRequest) (*entity.BrandList, error) {
	brandModel := pms_brand.Get(request.Id)
	return &entity.BrandList{
		Id:                  brandModel.ID,
		Sort:                brandModel.Sort,
		ShowStatus:          brandModel.ShowStatus,
		ProductCount:        brandModel.ProductCount,
		ProductCommentCount: brandModel.ProductCommentCount,
		FactoryStatus:       brandModel.FactoryStatus,
		Name:                brandModel.Name,
		FirstLetter:         brandModel.FirstLetter,
		Logo:                brandModel.Logo,
		BigPic:              brandModel.BigPic,
		BrandStory:          brandModel.BrandStory,
	}, nil
}

func (c AdminController) BrandList(ctx *gin.Context, request *entity.BrandListRequest) (*entity.BrandListReply, error) {
	where := ""
	if request.Keyword != "" {
		where += fmt.Sprintf("name LIKE %s", "'%"+request.Keyword+"%'")
	}
	list, paging := pms_brand.Paginate(ctx, request.PageSize, where, "sort")
	replyList := make([]entity.BrandList, 0, request.PageSize)
	for _, v := range list {
		interposition := entity.BrandList{
			Id:                  v.ID,
			Sort:                v.Sort,
			ShowStatus:          v.ShowStatus,
			ProductCount:        v.ProductCount,
			ProductCommentCount: v.ProductCommentCount,
			FactoryStatus:       v.FactoryStatus,
			Name:                v.Name,
			FirstLetter:         v.FirstLetter,
			Logo:                v.Logo,
			BigPic:              v.BigPic,
			BrandStory:          v.BrandStory,
		}
		replyList = append(replyList, interposition)
	}
	return &entity.BrandListReply{
		List:        replyList,
		PagingAdmin: paging,
	}, nil
}

func (c AdminController) BrandCreate(ctx *gin.Context, request *entity.BrandCreateRequest) error {
	brandModel := pms_brand.PmsBrand{
		Name:                request.Name,
		FirstLetter:         request.FirstLetter,
		Sort:                request.Sort,
		FactoryStatus:       request.FactoryStatus,
		ShowStatus:          request.ShowStatus,
		ProductCount:        0,
		ProductCommentCount: 0,
		Logo:                request.Logo,
		BigPic:              request.BigPic,
		BrandStory:          request.BrandStory,
	}
	brandModel.Create()
	return nil
}

func (c AdminController) BrandUpdate(ctx *gin.Context, request *entity.BrandUpdateRequest) error {
	brandModel := pms_brand.Get(strconv.Itoa(request.Id))
	brandModel.Updates(map[string]interface{}{
		"name":           request.Name,
		"first_letter":   request.FirstLetter,
		"sort":           request.Sort,
		"factory_status": request.FactoryStatus,
		"show_status":    request.ShowStatus,
		"logo":           request.Logo,
		"big_pic":        request.BigPic,
		"brand_story":    request.BrandStory,
	})
	return nil
}

func (c AdminController) BrandDelete(ctx *gin.Context, request *entity.BrandDeleteRequest) error {
	bandModel := pms_brand.Get(request.Id)
	bandModel.Delete()
	return nil
}

func (c AdminController) BrandUpdateFactoryStatus(ctx *gin.Context, request *entity.BrandUpdateFactoryStatusRequest) error {
	parts := strings.SplitN(request.Ids, ",", -1)
	for _, v := range parts {
		brandModel := pms_brand.Get(v)
		brandModel.Updates(map[string]interface{}{
			"factory_status": request.FactoryStatus,
		})
	}
	return nil
}

func (c AdminController) BrandUpdateShowStatus(ctx *gin.Context, request *entity.BrandUpdateShowStatusRequest) error {
	parts := strings.SplitN(request.Ids, ",", -1)
	for _, v := range parts {
		brandModel := pms_brand.Get(v)
		brandModel.Updates(map[string]interface{}{
			"show_status": request.ShowStatus,
		})
	}
	return nil
}
