package admin

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/models/ums_admin"
	"go-mall-api/models/ums_admin_role_relation"
	"strconv"
	"strings"
	"time"
)

func (c AdminController) AdminList(ctx *gin.Context, request *entity.ListRequest) (reply *entity.ListReply, err error) {
	list, paging := ums_admin.Paginate(ctx, request.PageSize, request.Keyword)
	for _, v := range list {
		v.CreateTime.Format("2006-01-02 15:04:05")
	}

	return &entity.ListReply{
		List:        list,
		PagingAdmin: paging,
	}, nil
}

func (c AdminController) AdminUpdateStatus(ctx *gin.Context, request *entity.AdminUpdateStatusRequest) error {
	userModel := ums_admin.Get(request.Id)
	userModel.Updates(map[string]interface{}{"status": request.Status})
	return nil
}

func (c AdminController) AdminRegister(ctx *gin.Context, request *entity.RegisterRequest) error {
	userModel := ums_admin.UmsAdmin{
		NickName:   request.NickName,
		Username:   request.Username,
		Password:   request.Password,
		Email:      request.Email,
		Note:       request.Note,
		Status:     request.Status,
		CreateTime: time.Now(),
		LoginTime:  time.Now(),
	}
	userModel.Create()
	return nil
}

func (c AdminController) AdminRoleUpdate(ctx *gin.Context, request *entity.AdminRoleUpdateRequest) error {
	adminId, _ := strconv.ParseInt(request.AdminId, 10, 64)
	parts := strings.SplitN(request.RoleIds, ",", -1)
	for _, v := range parts {
		roleId, _ := strconv.ParseInt(v, 10, 64)
		relationModel := ums_admin_role_relation.UmsAdminRoleRelation{
			AdminId: adminId,
			RoleId:  roleId,
		}
		relationModel.Create()
	}
	return nil
}
