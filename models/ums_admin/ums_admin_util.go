package ums_admin

import (
	"go-mall-api/pkg/database"
)

func Get(idstr string) (umsAdmin UmsAdmin) {
	database.DB.Where("id", idstr).First(&umsAdmin)
	return
}

func GetBy(field, value string) (umsAdmin UmsAdmin) {
	database.DB.Where("? = ?", field, value).First(&umsAdmin)
	return
}

func All() (umsAdmins []UmsAdmin) {
	database.DB.Find(&umsAdmins)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(UmsAdmin{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

// GetByMulti 通过 用户名 来获取用户
func GetByMulti(username string) (userAdminModel UmsAdmin) {
	database.DB.Where("username = ?", username).First(&userAdminModel)
	return
}
