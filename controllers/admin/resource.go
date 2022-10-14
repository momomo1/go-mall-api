package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/models/ums_resource"
	"go-mall-api/models/ums_resource_category"
	"strconv"
	"time"
)

func (c AdminController) ResourceListAll(ctx *gin.Context) ([]*entity.ResourceList, error) {
	list, count := ums_resource.All()
	replyList := make([]*entity.ResourceList, 0, count)
	for _, v := range list {
		interposition := &entity.ResourceList{
			Id:          v.Id,
			CategoryId:  v.CategoryId,
			Name:        v.Name,
			Url:         v.Url,
			Description: v.Description,
			CreateTime:  v.CreateTime.Format("2006-01-02 15:04:05"),
		}
		replyList = append(replyList, interposition)
	}
	return replyList, nil
}

func (c AdminController) ResourceList(ctx *gin.Context, request *entity.ResourceListRequest) (*entity.ResourceListReply, error) {
	where := ""
	if request.NameKeyword != "" {
		whereAnd(&where)
		where += fmt.Sprintf("name LIKE %s", "'%"+request.NameKeyword+"%'")
	}
	if request.UrlKeyword != "" {
		whereAnd(&where)
		where += fmt.Sprintf("url LIKE %s", "'%"+request.UrlKeyword+"%'")
	}
	if request.CategoryId != 0 {
		whereAnd(&where)
		where += fmt.Sprintf("category_id = %d", request.CategoryId)
	}

	list, paging := ums_resource.Paginate(ctx, request.PageSize, where, "id", "asc")
	replyList := make([]entity.ResourceList, 0, request.PageSize)
	for _, v := range list {
		interposition := entity.ResourceList{
			Id:          v.Id,
			CategoryId:  v.CategoryId,
			Name:        v.Name,
			Url:         v.Url,
			Description: v.Description,
			CreateTime:  v.CreateTime.Format("2006-01-02 15:04:05"),
		}
		replyList = append(replyList, interposition)
	}
	return &entity.ResourceListReply{
		List:        replyList,
		PagingAdmin: paging,
	}, nil
}

func (c AdminController) ResourceCreate(ctx *gin.Context, request *entity.ResourceCreateRequest) error {
	categoryModel := ums_resource.UmsResource{
		Name:        request.Name,
		Url:         request.Url,
		Description: request.Description,
		CategoryId:  request.CategoryId,
		CreateTime:  time.Now(),
	}
	categoryModel.Create()
	return nil
}

func (c AdminController) ResourceUpdate(ctx *gin.Context, request *entity.ResourceUpdateRequest) error {
	resourceModel := ums_resource.Get(strconv.Itoa(request.Id))
	resourceModel.Updates(map[string]interface{}{
		"Name":        request.Name,
		"Url":         request.Url,
		"Description": request.Description,
		"CategoryId":  request.CategoryId,
	})
	return nil
}

func (c AdminController) ResourceDelete(ctx *gin.Context, request *entity.ResourceDeleteRequest) error {
	resourceModel := ums_resource.Get(request.Id)
	resourceModel.Delete()
	return nil
}

func (c AdminController) ResourceCategoryListAll(ctx *gin.Context) ([]*entity.ResourceCategoryListAllReply, error) {
	category, count := ums_resource_category.All()
	reply := make([]*entity.ResourceCategoryListAllReply, 0, count)
	for _, v := range category {
		interposition := &entity.ResourceCategoryListAllReply{
			Id:         v.Id,
			Sort:       v.Sort,
			Name:       v.Name,
			CreateTime: v.CreateTime.Format("2006-01-02 15:04:05"),
		}
		reply = append(reply, interposition)
	}
	return reply, nil
}

func (c AdminController) ResourceCategoryCreate(ctx *gin.Context, request *entity.ResourceCategoryCreateRequest) error {
	sort, _ := strconv.Atoi(request.Sort)
	categoryModel := ums_resource_category.UmsResourceCategory{
		Name:       request.Name,
		Sort:       sort,
		CreateTime: time.Now(),
	}
	categoryModel.Create()
	return nil
}

func (c AdminController) ResourceCategoryUpdate(ctx *gin.Context, request *entity.ResourceCategoryUpdateRequest) error {
	categoryModel := ums_resource_category.Get(strconv.Itoa(request.Id))
	categoryModel.Updates(map[string]interface{}{
		"Name": request.Name,
		"Sort": request.Sort,
	})
	return nil
}

func (c AdminController) ResourceCategoryDelete(ctx *gin.Context, request *entity.ResourceCategoryDeleteRequest) error {
	categoryModel := ums_resource_category.Get(request.Id)
	categoryModel.Delete()
	return nil
}
