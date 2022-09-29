package pms_product_attribute

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/models/pms_product_attribute_category"
	"go-mall-api/pkg/database"
	"go-mall-api/pkg/paginator"
	"gorm.io/gorm"
	"strconv"
)

func Get(idstr string) (pmsProductAttribute PmsProductAttribute) {
	database.DB.Where("id", idstr).First(&pmsProductAttribute)
	return
}

func GetBy(field, value string) (pmsProductAttribute PmsProductAttribute) {
	database.DB.Where("? = ?", field, value).First(&pmsProductAttribute)
	return
}

func All() (pmsProductAttributes []PmsProductAttribute) {
	database.DB.Find(&pmsProductAttributes)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(PmsProductAttribute{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int, where interface{}, sort string, order string) (attribute []PmsProductAttribute, paging paginator.PagingAdmin) {
	paging = paginator.PaginateAdmin(
		c,
		database.DB.Model(PmsProductAttribute{}),
		&attribute,
		where,
		perPage,
		sort,
		order,
	)
	return
}

//AddAttribute 添加属性
func AddAttribute(types int, attribute PmsProductAttribute, attributeCategory pms_product_attribute_category.PmsProductAttributeCategory) error {
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
func UpdateAttribute(request *entity.ProductAttributeUpdateRequest, attributeModel PmsProductAttribute) error {
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
			categoryOneModel := pms_product_attribute_category.Get(strconv.FormatInt(attributeModel.ProductAttributeCategoryId, 10))
			if err := tx.Model(&categoryOneModel).Updates(dataDelete).Error; err != nil {
				return err
			}
			//增加
			categoryTwoModel := pms_product_attribute_category.Get(strconv.FormatInt(request.ProductAttributeCategoryId, 10))
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
			attribute := Get(v)
			if err := tx.Delete(&attribute).Error; err != nil {
				return err
			}
			tests = attribute.Type
			productAttributeCategoryId = attribute.ProductAttributeCategoryId
		}

		attributeCategory := pms_product_attribute_category.Get(strconv.FormatInt(productAttributeCategoryId, 10))
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
