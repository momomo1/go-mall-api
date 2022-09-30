package pms_product_attribute_category

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/models/pms_product_attribute"
	"go-mall-api/pkg/database"
	"go-mall-api/pkg/paginator"
	"gorm.io/gorm"
	"strconv"
)

func Get(idstr string) (pmsProductAttributeCategory PmsProductAttributeCategory) {
	database.DB.Where("id", idstr).First(&pmsProductAttributeCategory)
	return
}

func GetBy(field, value string) (pmsProductAttributeCategory PmsProductAttributeCategory) {
	database.DB.Where("? = ?", field, value).First(&pmsProductAttributeCategory)
	return
}

func All() (pmsProductAttributeCategories []PmsProductAttributeCategory) {
	database.DB.Find(&pmsProductAttributeCategories)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(PmsProductAttributeCategory{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int, where interface{}, sort string, order string) (attributeCategory []PmsProductAttributeCategory, paging paginator.PagingAdmin) {
	paging = paginator.PaginateAdmin(
		c,
		database.DB.Model(PmsProductAttributeCategory{}),
		&attributeCategory,
		where,
		perPage,
		sort,
		order,
	)
	return
}

//AddAttribute 添加属性
func AddAttribute(types int, attribute pms_product_attribute.PmsProductAttribute, attributeCategory PmsProductAttributeCategory) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Create(&attribute).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		data := map[string]interface{}{}
		if types == 0 {
			data = map[string]interface{}{"attribute_count": gorm.Expr("attribute_count + ?", 1)}
		} else {
			data = map[string]interface{}{"param_count": gorm.Expr("param_count + ?", 1)}
		}
		if err := tx.Model(&attributeCategory).Updates(data).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})

	return err
}

//UpdateAttribute 更新属性
func UpdateAttribute(request *entity.ProductAttributeUpdateRequest, attributeModel pms_product_attribute.PmsProductAttribute) error {
	data := map[string]interface{}{
		"filter_type":                   request.FilterType,
		"hand_add_status":               request.HandAddStatus,
		"input_type":                    request.InputType,
		"product_attribute_category_id": request.ProductAttributeCategoryId,
		"related_status":                request.RelatedStatus,
		"search_type":                   request.SearchType,
		"select_type":                   request.SelectType,
		"sort":                          request.Sort,
		"type":                          request.Type,
		"input_list":                    request.InputList,
		"name":                          request.Name,
	}

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if attributeModel.ProductAttributeCategoryId != request.ProductAttributeCategoryId {
			dataDelete := map[string]interface{}{}
			dataAdd := map[string]interface{}{}
			if attributeModel.Type == 0 {
				dataDelete = map[string]interface{}{"attribute_count": gorm.Expr("attribute_count - ?", 1)}
				dataAdd = map[string]interface{}{"attribute_count": gorm.Expr("attribute_count + ?", 1)}
			} else {
				dataDelete = map[string]interface{}{"param_count": gorm.Expr("param_count - ?", 1)}
				dataAdd = map[string]interface{}{"param_count": gorm.Expr("param_count + ?", 1)}
			}

			//减少
			categoryOneModel := Get(strconv.FormatInt(attributeModel.ProductAttributeCategoryId, 10))
			if err := tx.Model(&categoryOneModel).Updates(dataDelete).Error; err != nil {
				return err
			}
			//增加
			categoryTwoModel := Get(strconv.FormatInt(request.ProductAttributeCategoryId, 10))
			if err := tx.Model(&categoryTwoModel).Updates(dataAdd).Error; err != nil {
				return err
			}
		}
		if err := tx.Model(&attributeModel).Updates(data).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
	return err
}

//DeleteAttribute 删除属性
func DeleteAttribute(parts []string) error {
	var tests int
	var productAttributeCategoryId int64

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		for _, v := range parts {
			attribute := pms_product_attribute.Get(v)
			if err := tx.Delete(&attribute).Error; err != nil {
				return err
			}
			tests = attribute.Type
			productAttributeCategoryId = attribute.ProductAttributeCategoryId
		}

		attributeCategory := Get(strconv.FormatInt(productAttributeCategoryId, 10))
		data := map[string]interface{}{}
		count := len(parts)
		if tests == 0 {
			data = map[string]interface{}{"attribute_count": gorm.Expr("attribute_count - ?", count)}
		} else {
			data = map[string]interface{}{"param_count": gorm.Expr("param_count - ?", count)}
		}
		if err := tx.Model(&attributeCategory).Updates(data).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})

	return err
}

//GetClassificationAttribute 获取分类属性
func GetClassificationAttribute() (pmsProductAttributeCategories []PmsProductAttributeCategory, count int64) {
	database.DB.Preload("PmsProductAttributes", "type = ?", 1).Find(&pmsProductAttributeCategories)
	database.DB.Model(PmsProductAttributeCategory{}).Count(&count)
	return
}
