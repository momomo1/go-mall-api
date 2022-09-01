package auth_admin

import (
	"errors"
	"go-mall-api/models/ums_admin"
)

// Attempt 尝试登录
func Attempt(username string, password string) (ums_admin.UmsAdmin, error) {
	userAdminModel := ums_admin.GetByMulti(username)

	if userAdminModel.ID == 0 {
		return ums_admin.UmsAdmin{}, errors.New("账号不存在")
	}
	if !userAdminModel.ComparePassword(password) {
		return ums_admin.UmsAdmin{}, errors.New("密码错误")
	}

	//更新登录时间
	ums_admin.UpdateLoginTime(userAdminModel)
	return userAdminModel, nil
}
