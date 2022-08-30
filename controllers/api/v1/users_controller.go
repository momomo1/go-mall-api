package v1

import (
	"go-mall-api/models/user"
	"go-mall-api/pkg/response"
	"go-mall-api/requests/api"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	BaseAPIController
}

// Index 所有用户
func (ctrl *UsersController) Index(c *gin.Context) {
	request := api.PaginationRequest{}
	if ok := api.Validate(c, &request, api.Pagination); !ok {
		return
	}

	data, pager := user.Paginate(c, 10)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}
