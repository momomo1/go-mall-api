package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/models/sms_home_advertise"
	"go-mall-api/models/sms_home_brand"
	"go-mall-api/models/sms_home_new_product"
	"go-mall-api/models/sms_home_recommend_product"
	"go-mall-api/models/sms_home_recommend_subject"
	"strconv"
	"strings"
)

func (c AdminController) HomeBrandList(ctx *gin.Context, request *entity.HomeBrandListRequest) (*entity.HomeBrandListReply, error) {
	where := ""
	if request.BrandName != "" {
		whereAnd(&where)
		where += fmt.Sprintf("brand_name LIKE %s", "'%"+request.BrandName+"%'")
	}
	if request.RecommendStatus != "" {
		whereAnd(&where)
		where += fmt.Sprintf("recommend_status = '%s'", request.RecommendStatus)
	}

	list, paging := sms_home_brand.Paginate(ctx, request.PageSize, where, "sort", "desc")
	replyList := make([]entity.HomeBrandList, 0, request.PageSize)
	for _, v := range list {
		interposition := entity.HomeBrandList{
			Id:              v.ID,
			Sort:            v.Sort,
			BrandId:         v.BrandId,
			RecommendStatus: v.RecommendStatus,
			BrandName:       v.BrandName,
		}
		replyList = append(replyList, interposition)
	}

	return &entity.HomeBrandListReply{
		List:        replyList,
		PagingAdmin: paging,
	}, nil
}

func (c AdminController) HomeBrandUpdateRecommendStatus(ctx *gin.Context, request *entity.HomeBrandUpdateRecommendStatusRequest) error {
	ids := strings.SplitN(request.Ids, ",", -1)
	for _, v := range ids {
		homeBrand := sms_home_brand.Get(v)
		homeBrand.Updates(map[string]interface{}{
			"recommend_status": request.RecommendStatus,
		})
	}

	return nil
}

func (c AdminController) HomeBrandCreate(ctx *gin.Context, request *[]entity.HomeBrandCreateRequest) error {
	for _, v := range *request {
		homeBrand := sms_home_brand.GetByBrandId(strconv.FormatInt(v.BrandId, 10))
		if homeBrand.ID == 0 {
			homeBrandModel := sms_home_brand.SmsHomeBrand{
				BrandId:   v.BrandId,
				BrandName: v.BrandName,
			}
			homeBrandModel.Create()
		}
	}
	return nil
}

func (c AdminController) HomeBrandUpdateSort(ctx *gin.Context, request *entity.HomeBrandUpdateSortRequest) error {
	homeBrand := sms_home_brand.Get(request.Id)
	homeBrand.Updates(map[string]interface{}{
		"sort": request.Sort,
	})
	return nil
}

func (c AdminController) HomeBrandDelete(ctx *gin.Context, request *entity.HomeBrandDeleteRequest) error {
	ids := strings.SplitN(request.Ids, ",", -1)
	for _, v := range ids {
		homeBrand := sms_home_brand.Get(v)
		homeBrand.Delete()
	}
	return nil
}

func (c AdminController) HomeRecommendSubjectList(ctx *gin.Context, request *entity.HomeRecommendSubjectListRequest) (*entity.HomeRecommendSubjectListReply, error) {
	where := ""
	if request.SubjectName != "" {
		whereAnd(&where)
		where += fmt.Sprintf("subject_name LIKE %s", "'%"+request.SubjectName+"%'")
	}
	if request.RecommendStatus != "" {
		whereAnd(&where)
		where += fmt.Sprintf("recommend_status = '%s'", request.RecommendStatus)
	}

	list, paging := sms_home_recommend_subject.Paginate(ctx, request.PageSize, where, "sort", "desc")
	replyList := make([]entity.HomeRecommendSubjectList, 0, request.PageSize)
	for _, v := range list {
		interposition := entity.HomeRecommendSubjectList{
			Id:              v.ID,
			Sort:            v.Sort,
			SubjectId:       v.SubjectId,
			RecommendStatus: v.RecommendStatus,
			SubjectName:     v.SubjectName,
		}
		replyList = append(replyList, interposition)
	}
	return &entity.HomeRecommendSubjectListReply{
		List:        replyList,
		PagingAdmin: paging,
	}, nil
}

func (c AdminController) HomeRecommendSubjectUpdateRecommendStatus(ctx *gin.Context, request *entity.HomeRecommendSubjectUpdateRecommendStatusRequest) error {
	ids := strings.SplitN(request.Ids, ",", -1)
	for _, v := range ids {
		homeRecommendSubject := sms_home_recommend_subject.Get(v)
		homeRecommendSubject.Updates(map[string]interface{}{
			"recommend_status": request.RecommendStatus,
		})
	}
	return nil
}

func (c AdminController) HomeRecommendSubjectCreate(ctx *gin.Context, request *[]entity.HomeRecommendSubjectCreateRequest) error {
	for _, v := range *request {
		homeRecommendSubject := sms_home_recommend_subject.GetBySubjectId(strconv.FormatInt(v.SubjectId, 10))
		if homeRecommendSubject.ID == 0 {
			homeRecommendSubjectModel := sms_home_recommend_subject.SmsHomeRecommendSubject{
				SubjectId:   v.SubjectId,
				SubjectName: v.SubjectName,
			}
			homeRecommendSubjectModel.Create()
		}
	}
	return nil
}

func (c AdminController) HomeRecommendSubjectUpdateSort(ctx *gin.Context, request *entity.HomeRecommendSubjectUpdateSortRequest) error {
	homeRecommendSubject := sms_home_recommend_subject.Get(request.Id)
	homeRecommendSubject.Updates(map[string]interface{}{
		"sort": request.Sort,
	})
	return nil
}

func (c AdminController) HomeRecommendSubjectDelete(ctx *gin.Context, request *entity.HomeRecommendSubjectDeleteRequest) error {
	ids := strings.SplitN(request.Ids, ",", -1)
	for _, v := range ids {
		homeRecommendSubject := sms_home_recommend_subject.Get(v)
		homeRecommendSubject.Delete()
	}
	return nil
}

func (c AdminController) HomeAdvertise(ctx *gin.Context, request *entity.HomeAdvertiseRequest) (*entity.HomeAdvertiseList, error) {
	advertise := sms_home_advertise.Get(request.Id)
	return &entity.HomeAdvertiseList{
		Id:         advertise.ID,
		ClickCount: advertise.ClickCount,
		OrderCount: advertise.OrderCount,
		Sort:       advertise.Sort,
		Status:     advertise.Status,
		Type:       advertise.Type,
		Name:       advertise.Name,
		Note:       advertise.Note,
		Pic:        advertise.Pic,
		Url:        advertise.Url,
		StartTime:  advertise.StartTime,
		EndTime:    advertise.EndTime,
	}, nil
}

func (c AdminController) HomeAdvertiseList(ctx *gin.Context, request *entity.HomeAdvertiseListRequest) (*entity.HomeAdvertiseListReply, error) {
	where := ""
	if request.Name != "" {
		whereAnd(&where)
		where += fmt.Sprintf("name LIKE %s", "'%"+request.Name+"%'")
	}
	if request.Type != "" {
		whereAnd(&where)
		where += fmt.Sprintf("type = '%s'", request.Type)
	}
	if request.EndTime != "" {
		whereAnd(&where)
		where += fmt.Sprintf("end_time <= '%s'", (request.EndTime + " 00:00:00"))
	}

	list, paging := sms_home_advertise.Paginate(ctx, request.PageSize, where, "sort", "desc")
	replyList := make([]entity.HomeAdvertiseList, 0, request.PageSize)
	for _, v := range list {
		interposition := entity.HomeAdvertiseList{
			Id:         v.ID,
			ClickCount: v.ClickCount,
			OrderCount: v.OrderCount,
			Sort:       v.Sort,
			Status:     v.Status,
			Type:       v.Type,
			Name:       v.Name,
			Note:       v.Note,
			Pic:        v.Pic,
			Url:        v.Url,
			StartTime:  v.StartTime,
			EndTime:    v.EndTime,
		}
		replyList = append(replyList, interposition)
	}

	return &entity.HomeAdvertiseListReply{
		List:        replyList,
		PagingAdmin: paging,
	}, nil
}

func (c AdminController) HomeAdvertiseUpdateStatus(ctx *gin.Context, request *entity.HomeAdvertiseUpdateStatusRequest) error {
	homeAdvertise := sms_home_advertise.Get(request.Id)
	homeAdvertise.Updates(map[string]interface{}{
		"status": request.Status,
	})
	return nil
}

func (c AdminController) HomeAdvertiseCreate(ctx *gin.Context, request *entity.HomeAdvertiseCreateRequest) error {
	advertiseModel := sms_home_advertise.SmsHomeAdvertise{
		Name:       request.Name,
		Type:       request.Type,
		Pic:        request.Pic,
		StartTime:  request.StartTime,
		EndTime:    request.EndTime,
		Status:     request.Status,
		ClickCount: 0,
		OrderCount: 0,
		Url:        request.Url,
		Note:       request.Note,
		Sort:       request.Sort,
	}

	advertiseModel.Create()
	return nil
}

func (c AdminController) HomeAdvertiseUpdate(ctx *gin.Context, request *entity.HomeAdvertiseUpdateRequest) error {
	advertise := sms_home_advertise.Get(request.Id)
	advertise.Updates(map[string]interface{}{
		"Name":      request.Name,
		"Type":      request.Type,
		"Pic":       request.Pic,
		"StartTime": request.StartTime,
		"EndTime":   request.EndTime,
		"Status":    request.Status,
		"Url":       request.Url,
		"Note":      request.Note,
		"Sort":      request.Sort,
	})
	return nil
}

func (c AdminController) HomeAdvertiseDelete(ctx *gin.Context, request *entity.HomeAdvertiseDeleteRequest) error {
	ids := strings.SplitN(request.Ids, ",", -1)
	for _, v := range ids {
		homeAdvertise := sms_home_advertise.Get(v)
		homeAdvertise.Delete()
	}
	return nil
}

func (c AdminController) HomeNewProductList(ctx *gin.Context, request *entity.HomeNewProductListRequest) (*entity.HomeNewProductListReply, error) {
	where := ""
	if request.ProductName != "" {
		whereAnd(&where)
		where += fmt.Sprintf("product_name LIKE %s", "'%"+request.ProductName+"%'")
	}
	if request.RecommendStatus != "" {
		whereAnd(&where)
		where += fmt.Sprintf("recommend_status = '%s'", request.RecommendStatus)
	}

	list, paging := sms_home_new_product.Paginate(ctx, request.PageSize, where, "sort", "desc")
	replyList := make([]entity.HomeNewProductList, 0, request.PageSize)
	for _, v := range list {
		interposition := entity.HomeNewProductList{
			Id:              v.ID,
			ProductId:       v.ProductId,
			RecommendStatus: v.RecommendStatus,
			Sort:            v.Sort,
			ProductName:     v.ProductName,
		}
		replyList = append(replyList, interposition)
	}

	return &entity.HomeNewProductListReply{
		List:        replyList,
		PagingAdmin: paging,
	}, nil
}

func (c AdminController) HomeNewProductUpdateRecommendStatus(ctx *gin.Context, request *entity.HomeNewProductUpdateRecommendStatusRequest) error {
	ids := strings.SplitN(request.Ids, ",", -1)
	for _, v := range ids {
		homeNewProduct := sms_home_new_product.Get(v)
		homeNewProduct.Updates(map[string]interface{}{
			"recommend_status": request.RecommendStatus,
		})
	}
	return nil
}

func (c AdminController) HomeNewProductCreate(ctx *gin.Context, request *[]entity.HomeNewProductCreateRequest) error {
	for _, v := range *request {
		homeNewProduct := sms_home_new_product.GetByProductId(strconv.FormatInt(v.ProductId, 10))
		if homeNewProduct.ID == 0 {
			homeNewProductModel := sms_home_new_product.SmsHomeNewProduct{
				ProductId:   v.ProductId,
				ProductName: v.ProductName,
			}
			homeNewProductModel.Create()
		}
	}
	return nil
}

func (c AdminController) HomeNewProductUpdateSort(ctx *gin.Context, request *entity.HomeNewProductUpdateSortRequest) error {
	newProductModel := sms_home_new_product.Get(request.Id)
	newProductModel.Updates(map[string]interface{}{
		"sort": request.Sort,
	})
	return nil
}

func (c AdminController) HomeNewProductDelete(ctx *gin.Context, request *entity.HomeNewProductDeleteRequest) error {
	ids := strings.SplitN(request.Ids, ",", -1)
	for _, v := range ids {
		homeNewProduct := sms_home_new_product.Get(v)
		homeNewProduct.Delete()
	}
	return nil
}

func (c AdminController) HomeRecommendProductList(ctx *gin.Context, request *entity.HomeRecommendProductListRequest) (*entity.HomeRecommendProductListReply, error) {
	where := ""
	if request.ProductName != "" {
		whereAnd(&where)
		where += fmt.Sprintf("product_name LIKE %s", "'%"+request.ProductName+"%'")
	}
	if request.RecommendStatus != "" {
		whereAnd(&where)
		where += fmt.Sprintf("recommend_status = '%s'", request.RecommendStatus)
	}

	list, paging := sms_home_recommend_product.Paginate(ctx, request.PageSize, where, "sort", "desc")
	replyList := make([]entity.HomeRecommendProductList, 0, request.PageSize)
	for _, v := range list {
		interposition := entity.HomeRecommendProductList{
			Id:              v.ID,
			ProductId:       v.ProductId,
			RecommendStatus: v.RecommendStatus,
			Sort:            v.Sort,
			ProductName:     v.ProductName,
		}
		replyList = append(replyList, interposition)
	}

	return &entity.HomeRecommendProductListReply{
		List:        replyList,
		PagingAdmin: paging,
	}, nil
}

func (c AdminController) HomeRecommendProductUpdateRecommendStatus(ctx *gin.Context, request *entity.HomeRecommendProductUpdateRecommendStatusRequest) error {
	ids := strings.SplitN(request.Ids, ",", -1)
	for _, v := range ids {
		homeRecommendProduct := sms_home_recommend_product.Get(v)
		homeRecommendProduct.Updates(map[string]interface{}{
			"recommend_status": request.RecommendStatus,
		})
	}
	return nil
}

func (c AdminController) HomeRecommendProductCreate(ctx *gin.Context, request *[]entity.HomeRecommendProductCreateRequest) error {
	for _, v := range *request {
		homeRecommendProduct := sms_home_recommend_product.GetByProductId(strconv.FormatInt(v.ProductId, 10))
		if homeRecommendProduct.ID == 0 {
			homeRecommendProductModel := sms_home_recommend_product.SmsHomeRecommendProduct{
				ProductId:   v.ProductId,
				ProductName: v.ProductName,
			}
			homeRecommendProductModel.Create()
		}
	}
	return nil
}

func (c AdminController) HomeRecommendProductUpdateSort(ctx *gin.Context, request *entity.HomeRecommendProductUpdateSortRequest) error {
	homeRecommendProduct := sms_home_recommend_product.Get(request.Id)
	homeRecommendProduct.Updates(map[string]interface{}{
		"sort": request.Sort,
	})
	return nil
}

func (c AdminController) HomeRecommendProductDelete(ctx *gin.Context, request *entity.HomeRecommendProductDeleteRequest) error {
	ids := strings.SplitN(request.Ids, ",", -1)
	for _, v := range ids {
		homeRecommendProduct := sms_home_recommend_product.Get(v)
		homeRecommendProduct.Delete()
	}
	return nil
}
