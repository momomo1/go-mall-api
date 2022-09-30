package pms_product_category

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/models/pms_product_category_attribute_relation"
	"go-mall-api/pkg/database"
	"go-mall-api/pkg/paginator"
	"gorm.io/gorm"
)

func Get(idstr string) (pmsProductCategory PmsProductCategory) {
	database.DB.Where("id", idstr).First(&pmsProductCategory)
	return
}

func GetBy(field, value string) (pmsProductCategory PmsProductCategory) {
	database.DB.Where("? = ?", field, value).First(&pmsProductCategory)
	return
}

func All() (pmsProductCategories []PmsProductCategory) {
	database.DB.Find(&pmsProductCategories)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(PmsProductCategory{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int, where interface{}, sort string, order string) (productCategory []PmsProductCategory, paging paginator.PagingAdmin) {
	paging = paginator.PaginateAdmin(
		c,
		database.DB.Model(PmsProductCategory{}),
		&productCategory,
		where,
		perPage,
		sort,
		order,
	)
	return
}

func DeleteRelevance(category PmsProductCategory) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&category).Error; err != nil {
			return err
		}
		//批量删除
		if err := tx.Where("product_category_id = ?", category.ID).Delete(&pms_product_category_attribute_relation.PmsProductCategoryAttributeRelation{}).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})

	return err
}

func UpdateProductCategory(data map[string]interface{}, category PmsProductCategory, productAttributeIdList []int64) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&category).Updates(data).Error; err != nil {
			return err
		}
		//批量删除
		if err := tx.Where("product_category_id = ?", category.ID).Delete(&pms_product_category_attribute_relation.PmsProductCategoryAttributeRelation{}).Error; err != nil {
			return err
		}
		//批量新增
		for _, v := range productAttributeIdList {
			relation := pms_product_category_attribute_relation.PmsProductCategoryAttributeRelation{
				ProductCategoryId:  int64(category.ID),
				ProductAttributeId: v,
			}
			if err := tx.Create(&relation).Error; err != nil {
				return err
			}
		}

		// 返回 nil 提交事务
		return nil
	})

	return err
}
