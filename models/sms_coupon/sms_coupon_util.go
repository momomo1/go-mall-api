package sms_coupon

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/models/sms_coupon_product_category_relation"
	"go-mall-api/models/sms_coupon_product_relation"
	"go-mall-api/pkg/database"
	"go-mall-api/pkg/helpers"
	"go-mall-api/pkg/paginator"
)

func Get(idstr string) (smsCoupon SmsCoupon) {
	database.DB.Where("id", idstr).First(&smsCoupon)
	return
}

func GetBy(field, value string) (smsCoupon SmsCoupon) {
	database.DB.Where("? = ?", field, value).First(&smsCoupon)
	return
}

func All() (smsCoupons []SmsCoupon) {
	database.DB.Find(&smsCoupons)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(SmsCoupon{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int, where interface{}, sort string, order string) (smsCoupon []SmsCoupon, paging paginator.PagingAdmin) {
	paging = paginator.PaginateAdmin(
		c,
		database.DB.Model(SmsCoupon{}),
		&smsCoupon,
		where,
		perPage,
		sort,
		order,
	)
	return
}

//CouponAdd  新增优惠券
func CouponAdd(coupon SmsCoupon, request *entity.CouponCreateRequest) error {
	// 开始事务
	tx := database.DB.Begin()
	err := tx.Create(&coupon).Error
	if err != nil {
		// 遇到错误时回滚事务
		tx.Rollback()
		return err
	}

	if len(request.ProductRelationList) != 0 {
		for _, v := range request.ProductRelationList {
			productRelation := sms_coupon_product_relation.SmsCouponProductRelation{
				CouponId:    int64(coupon.ID),
				ProductId:   v.ProductId,
				ProductName: v.ProductName,
				ProductSn:   v.ProductSn,
			}
			err := tx.Create(&productRelation).Error
			if err != nil {
				// 遇到错误时回滚事务
				tx.Rollback()
				return err
			}
		}
	}

	if len(request.ProductCategoryRelationList) != 0 {
		for _, v := range request.ProductCategoryRelationList {
			productCategoryRelation := sms_coupon_product_category_relation.SmsCouponProductCategoryRelation{
				CouponId:            int64(coupon.ID),
				ProductCategoryId:   v.ProductCategoryId,
				ProductCategoryName: v.ProductCategoryName,
				ParentCategoryName:  v.ParentCategoryName,
			}
			err := tx.Create(&productCategoryRelation).Error
			if err != nil {
				// 遇到错误时回滚事务
				tx.Rollback()
				return err
			}
		}
	}

	//提交事务
	tx.Commit()

	return nil
}

//CouponUpdate 更新优惠券
func CouponUpdate(coupon SmsCoupon, request *entity.CouponUpdateRequest) error {
	data := map[string]interface{}{
		"Type":         request.Type,
		"Name":         request.Name,
		"Platform":     request.Platform,
		"Amount":       request.Amount,
		"MinPoint":     request.MinPoint,
		"StartTime":    helpers.TimeStringToGoTime(request.StartTime),
		"EndTime":      helpers.TimeStringToGoTime(request.EndTime),
		"UseType":      request.UseType,
		"Note":         request.Note,
		"PublishCount": request.PublishCount,
		"EnableTime":   helpers.TimeStringToGoTime(request.EnableTime),
	}
	if request.PerLimit != "" {
		data["PerLimit"] = request.PerLimit
	}
	// 开始事务
	tx := database.DB.Begin()
	err := tx.Model(coupon).Updates(&data).Error
	if err != nil {
		tx.Rollback() //回滚
		return err
	}

	//删除商品分类关联
	categoryRelation := sms_coupon_product_category_relation.SmsCouponProductCategoryRelation{}
	err = tx.Delete(categoryRelation, "coupon_id = ?", coupon.ID).Error
	if err != nil {
		tx.Rollback() //回滚
		return err
	}
	//删除商品关联
	productRelation := sms_coupon_product_relation.SmsCouponProductRelation{}
	err = tx.Delete(productRelation, "coupon_id = ?", coupon.ID).Error
	if err != nil {
		tx.Rollback() //回滚
		return err
	}

	if request.UseType == 1 {
		if len(request.ProductCategoryRelationList) != 0 {
			for _, v := range request.ProductCategoryRelationList {
				productCategoryRelation := sms_coupon_product_category_relation.SmsCouponProductCategoryRelation{
					CouponId:            int64(coupon.ID),
					ProductCategoryId:   v.ProductCategoryId,
					ProductCategoryName: v.ProductCategoryName,
					ParentCategoryName:  v.ParentCategoryName,
				}
				fmt.Println(productCategoryRelation, "........")
				err := tx.Create(&productCategoryRelation).Error
				if err != nil {
					// 遇到错误时回滚事务
					tx.Rollback()
					return err
				}
			}
		}
	}

	if request.UseType == 2 {
		for _, v := range request.ProductCategoryRelationList {
			productCategoryRelation := sms_coupon_product_category_relation.SmsCouponProductCategoryRelation{
				CouponId:            int64(coupon.ID),
				ProductCategoryId:   v.ProductCategoryId,
				ProductCategoryName: v.ProductCategoryName,
				ParentCategoryName:  v.ParentCategoryName,
			}
			err := tx.Create(&productCategoryRelation).Error
			if err != nil {
				// 遇到错误时回滚事务
				tx.Rollback()
				return err
			}
		}
	}

	//提交事务
	tx.Commit()

	return nil
}

//CouponDelete 删除优惠券
func CouponDelete(coupon SmsCoupon) error {
	// 开始事务
	tx := database.DB.Begin()
	err := tx.Delete(&coupon).Error
	if err != nil {
		tx.Rollback() //回滚
		return err
	}
	if coupon.UseType == 1 {
		categoryRelation := sms_coupon_product_category_relation.SmsCouponProductCategoryRelation{}
		err = tx.Delete(categoryRelation, "coupon_id = ?", coupon.ID).Error
		if err != nil {
			tx.Rollback() //回滚
			return err
		}
	}

	if coupon.UseType == 2 {
		productRelation := sms_coupon_product_relation.SmsCouponProductRelation{}
		err = tx.Delete(productRelation, "coupon_id = ?", coupon.ID).Error
		if err != nil {
			tx.Rollback() //回滚
			return err
		}
	}

	//提交事务
	tx.Commit()

	return nil
}
