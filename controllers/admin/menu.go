package admin

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/models/ums_menu"
	"strconv"
	"time"
)

func (c AdminController) MenuTreeList(ctx *gin.Context) ([]*entity.MenuTreeListReply, error) {
	list, total := ums_menu.All()
	treeMap := make(map[string]*entity.MenuTreeListReply, total)
	for _, v := range list {
		id := strconv.Itoa(int(v.ID))
		treeMap[id] = &entity.MenuTreeListReply{
			Id:         v.ID,
			Level:      v.Level,
			ParentId:   v.ParentId,
			Hidden:     v.Hidden,
			Sort:       v.Sort,
			Name:       v.Name,
			Title:      v.Title,
			Icon:       v.Icon,
			CreateTime: v.CreateTime,
			Children:   []*entity.MenuTreeListReply{},
		}
	}
	treeData := make([]*entity.MenuTreeListReply, 0, total) //全部的树状数据
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

func (c AdminController) MenuList(ctx *gin.Context, request *entity.MenuListRequest) (*entity.MenuListReply, error) {
	where := map[string]interface{}{"parent_id": request.Id}

	list, paging := ums_menu.Paginate(ctx, request.PageSize, where, "id", "asc")
	return &entity.MenuListReply{
		List:        list,
		PagingAdmin: paging,
	}, nil
}

func (c AdminController) MenuUpdateHidden(ctx *gin.Context, request *entity.MenuUpdateHiddenRequest) error {
	menuModel := ums_menu.Get(request.Id)
	menuModel.Updates(map[string]interface{}{
		"Hidden": request.Hidden,
	})
	return nil
}

func (c AdminController) MenuCreate(ctx *gin.Context, request *entity.MenuCreateRequest) error {
	var level int
	if request.ParentId == 0 {
		level = 0
	} else {
		parentId := strconv.FormatInt(request.ParentId, 10)
		menu := ums_menu.Get(parentId)
		level = menu.Level + 1
	}
	menuModel := ums_menu.UmsMenu{
		Title:      request.Title,
		Name:       request.Name,
		Icon:       request.Icon,
		ParentId:   request.ParentId,
		Hidden:     request.Hidden,
		Sort:       request.Sort,
		Level:      level,
		CreateTime: time.Now(),
	}
	menuModel.Create()
	return nil
}

func (c AdminController) Menu(ctx *gin.Context, request *entity.MenuRequest) (*entity.MenuReply, error) {
	menu := ums_menu.Get(request.Id)
	return &entity.MenuReply{
		Data: menu,
	}, nil
}

func (c AdminController) MenuUpdate(ctx *gin.Context, request *entity.MenuUpdateRequest) error {
	menuId := strconv.Itoa(request.Id)
	menuModel := ums_menu.Get(menuId)
	menuModel.Updates(map[string]interface{}{
		"Title":    request.Title,
		"ParentId": request.ParentId,
		"Name":     request.Name,
		"Icon":     request.Icon,
		"Hidden":   request.Hidden,
		"Sort":     request.Sort,
	})
	return nil
}

func (c AdminController) MenuDelete(ctx *gin.Context, request *entity.MenuDeleteRequest) error {
	menuModel := ums_menu.Get(request.Id)
	menuModel.Delete()
	return nil
}
