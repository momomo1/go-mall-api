package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/models/cms_subject"
	"go-mall-api/models/oms_company_address"
)

func (c AdminController) SubjectList(ctx *gin.Context, request *entity.SubjectListRequest) (*entity.SubjectListReply, error) {
	where := ""
	if request.Keyword != "" {
		whereAnd(&where)
		where += fmt.Sprintf("title LIKE %s", "'%"+request.Keyword+"%'")
	}

	list, paging := cms_subject.Paginate(ctx, request.PageSize, where, "id", "asc")
	replyList := make([]entity.SubjectList, 0, request.PageSize)
	for _, v := range list {
		interposition := entity.SubjectList{
			Id:           v.ID,
			CategoryId:   v.CategoryId,
			Title:        v.Title,
			CategoryName: v.CategoryName,
			CreateTime:   v.CreateTime.Format("2006-01-02 15:04:05"),
		}
		replyList = append(replyList, interposition)
	}

	return &entity.SubjectListReply{
		List:        replyList,
		PagingAdmin: paging,
	}, nil
}

func (c AdminController) CompanyAddressList(ctx *gin.Context) (*entity.CompanyAddressListReply, error) {
	list := oms_company_address.All()
	replyList := make([]entity.CompanyAddressList, 0, len(list))
	for _, v := range list {
		interposition := entity.CompanyAddressList{
			Id:            v.ID,
			ReceiveStatus: v.ReceiveStatus,
			SendStatus:    v.SendStatus,
			AddressName:   v.AddressName,
			City:          v.City,
			DetailAddress: v.DetailAddress,
			Name:          v.Name,
			Phone:         v.Phone,
			Province:      v.Province,
			Region:        v.Region,
		}
		replyList = append(replyList, interposition)
	}

	return &entity.CompanyAddressListReply{
		List: replyList,
	}, nil
}
