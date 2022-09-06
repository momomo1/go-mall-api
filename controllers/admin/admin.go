package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	adminV1 "go-mall-api/api/admin/v1/entity"
	"go-mall-api/models/ums_admin"
)

func (c AdminController) AdminList(ctx *gin.Context, request *adminV1.ListRequest) (reply *adminV1.ListReply, err error) {
	list, paging := ums_admin.Paginate(ctx, request.PageSize)
	for _, v := range list {
		v.CreateTime.Format("2006-01-02 15:04:05")
	}

	return &adminV1.ListReply{
		List:        list,
		PagingAdmin: paging,
	}, nil
}

func (c AdminController) AdminUpdateStatus(ctx *gin.Context, request *adminV1.AdminUpdateStatusRequest) error {
	fmt.Println(request.Id)
	userModel := ums_admin.Get(request.Id)
	userModel.Updates(map[string]interface{}{"status": request.Status})
	return nil
}
