package admin

import (
	"github.com/gin-gonic/gin"
	adminV1 "go-mall-api/api/admin/v1"
	"go-mall-api/pkg/jwt"
	auth "go-mall-api/service/auth_admin"
)

func (c AdminController) Login(ctx *gin.Context, request *adminV1.LoginRequest) (*adminV1.LoginReply, error) {
	// 2. 尝试登录
	user, err := auth.Attempt(request.UserName, request.Password)
	if err != nil {
		// 失败,显示错误提示
		return &adminV1.LoginReply{}, err
	}
	token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Username)
	return &adminV1.LoginReply{
		Token:     token,
		TokenHead: "Bearer ",
	}, nil
}
