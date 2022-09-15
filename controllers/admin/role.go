package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/models/ums_role"
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
	
	list, paging := ums_role.Paginate(ctx, request.PageSize, where)
	for _, v := range list {
		v.CreateTime.Format("2006-01-02 15:04:05")
	}

	return &entity.RoleListReply{
		List:        list,
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
