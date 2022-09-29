package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/models/ums_role"
	"go-mall-api/models/ums_role_menu_relation"
	"go-mall-api/models/ums_role_resource_relation"
	"strconv"
	"strings"
	"time"
)

func (c AdminController) RoleListAll(ctx *gin.Context) (*entity.RoleListAllReply, error) {
	return &entity.RoleListAllReply{
		Data: ums_role.All(),
	}, nil
}

func (c AdminController) RoleList(ctx *gin.Context, request *entity.RoleListRequest) (*entity.RoleListReply, error) {
	where := ""
	if request.Keyword != "" {
		where = fmt.Sprintf("name LIKE %s", "'%"+request.Keyword+"%'")
	}
	list, paging := ums_role.Paginate(ctx, request.PageSize, where, "id", "asc")
	roleList := make([]entity.RoleList, 0, request.PageSize)
	for _, v := range list {
		interposition := entity.RoleList{
			Id:          v.ID,
			AdminCount:  v.AdminCount,
			Sort:        v.Sort,
			Status:      v.Status,
			Name:        v.Name,
			Description: v.Description,
			CreateTime:  v.CreateTime.Format("2006-01-02 15:04:05"),
		}
		roleList = append(roleList, interposition)
	}

	return &entity.RoleListReply{
		List:        roleList,
		PagingAdmin: paging,
	}, nil
}

func (c AdminController) RoleUpdateStatus(ctx *gin.Context, request *entity.RoleUpdateStatusRequest) error {
	roleModel := ums_role.Get(request.Id)
	roleModel.Updates(map[string]interface{}{"status": request.Status})
	return nil
}

func (c AdminController) RoleRegister(ctx *gin.Context, request *entity.RoleRegisterRequest) error {
	roleModel := ums_role.UmsRole{
		Name:        request.Name,
		Description: request.Description,
		Status:      &request.Status,
		CreateTime:  time.Now(),
	}
	roleModel.Create()
	return nil
}

func (c AdminController) RoleDelete(ctx *gin.Context, request *entity.RoleDeleteRequest) error {
	roleModel := ums_role.Get(request.Ids)
	roleModel.Delete()
	return nil
}

func (c AdminController) RoleUpdate(ctx *gin.Context, request *entity.RoleUpdateRequest) error {
	roelModel := ums_role.Get(request.Id)
	roelModel.Updates(map[string]interface{}{
		"Name":        request.Name,
		"Description": request.Description,
		"Status":      request.Status,
	})
	return nil
}

func (c AdminController) RoleListMenu(ctx *gin.Context, request *entity.RoleListMenuRequest) ([]*entity.RoleListMenuReply, error) {
	relation, count := ums_role_menu_relation.GetByRoleIdFind(request.Id)
	reply := make([]*entity.RoleListMenuReply, 0, count)
	for _, v := range relation {
		interposition := &entity.RoleListMenuReply{
			Id: v.MenuId,
		}
		reply = append(reply, interposition)
	}
	return reply, nil
}

func (c AdminController) RoleAllocMenu(ctx *gin.Context, request *entity.RoleAllocMenuRequest) error {
	roleId, _ := strconv.ParseInt(request.RoleId, 10, 64)
	menuIds := strings.SplitN(request.MenuIds, ",", -1)

	ums_role_menu_relation.DeleteByAdminId(request.RoleId)
	for _, v := range menuIds {
		MenuId, _ := strconv.ParseInt(v, 10, 64)
		relationModel := ums_role_menu_relation.UmsRoleMenuRelation{
			RoleId: roleId,
			MenuId: MenuId,
		}
		relationModel.Create()
	}

	return nil
}

func (c AdminController) RoleListResource(ctx *gin.Context, request *entity.RoleListResourceRequest) ([]*entity.RoleListResourceReply, error) {
	relation, count := ums_role_resource_relation.GetByRoleIdFind(request.Id)
	reply := make([]*entity.RoleListResourceReply, 0, count)
	for _, v := range relation {
		interposition := &entity.RoleListResourceReply{
			Id: v.ResourceId,
		}
		reply = append(reply, interposition)
	}
	return reply, nil
}

func (c AdminController) RoleAllocResource(ctx *gin.Context, request *entity.RoleAllocResourceRequest) error {
	roleId, _ := strconv.ParseInt(request.RoleId, 10, 64)
	resourceIds := strings.SplitN(request.ResourceIds, ",", -1)
	ums_role_resource_relation.DeleteByAdminId(request.RoleId)

	for _, v := range resourceIds {
		resourceId, _ := strconv.ParseInt(v, 10, 64)
		relationModel := ums_role_resource_relation.UmsRoleResourceRelation{
			RoleId:     roleId,
			ResourceId: resourceId,
		}
		relationModel.Create()
	}
	return nil
}
