package v1

import (
	requests2 "go-mall-api/app/http/controllers/api/requests"
	"go-mall-api/app/models/user"
	"go-mall-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	BaseAPIController
}

// Index 所有用户
func (ctrl *UsersController) Index(c *gin.Context) {
	request := requests2.PaginationRequest{}
	if ok := requests2.Validate(c, &request, requests2.Pagination); !ok {
		return
	}

	data, pager := user.Paginate(c, 10)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}
