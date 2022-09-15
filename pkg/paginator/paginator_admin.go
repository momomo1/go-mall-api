package paginator

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/pkg/config"
	"go-mall-api/pkg/logger"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// PagingAdmin 分页数据
type PagingAdmin struct {
	PageNum   int   `json:"pageNum"`   // 当前页
	PageSize  int   `json:"pageSize"`  // 每页条数
	Total     int64 `json:"total"`     // 总条数
	TotalPage int   `json:"totalPage"` // 总页数
}

func PaginateAdmin(c *gin.Context, db *gorm.DB, data interface{}, where interface{}, perPage int) PagingAdmin {
	// 初始化 Paginator 实例
	p := &Paginator{
		query: db,
		ctx:   c,
	}
	p.initPropertiesAdmin(perPage, where)

	// 查询数据库
	err := p.query.Preload(clause.Associations). // 读取关联
		Order(p.Sort + " " + p.Order). // 排序
		Where(where).
		Limit(p.PerPage).
		Offset(p.Offset).
		Find(data).
		Error

	// 数据库出错
	if err != nil {
		logger.LogIf(err)
		return PagingAdmin{}
	}

	return PagingAdmin{
		PageNum:   p.Page,
		PageSize:  p.PerPage,
		Total:     p.TotalCount,
		TotalPage: p.TotalPage,
	}
}

// 初始化分页必须用到的属性，基于这些属性查询数据库
func (p *Paginator) initPropertiesAdmin(perPage int, where interface{}) {
	p.PerPage = p.getPerPage(perPage)

	// 排序参数（控制器中以验证过这些参数，可放心使用）
	p.Order = p.ctx.DefaultQuery(config.Get("paging.url_query_order"), "asc")
	p.Sort = p.ctx.DefaultQuery(config.Get("paging.url_query_sort"), "id")

	p.TotalCount = p.getTotalCount(where)
	p.TotalPage = p.getTotalPage()
	p.Page = p.getCurrentPage()
	p.Offset = (p.Page - 1) * p.PerPage
}
