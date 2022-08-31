//Package ums_admin 模型
package ums_admin

import (
	"go-mall-api/models"
	"go-mall-api/models/ums_admin_role_relation"
	"go-mall-api/models/ums_role"
	"go-mall-api/pkg/database"
	"go-mall-api/pkg/hash"
	"time"
)

//UmsAdmin 后台用户表
type UmsAdmin struct {
	models.BaseModel
	Username      string                                         `gorm:"column:username;type:varchar(64)" json:"username"`
	Password      string                                         `gorm:"column:password;type:varchar(64)" json:"password"`
	Icon          string                                         `gorm:"column:icon;type:varchar(500);comment:头像" json:"icon"`
	Email         string                                         `gorm:"column:email;type:varchar(100);comment:邮箱" json:"email"`
	NickName      string                                         `gorm:"column:nick_name;type:varchar(200);comment:昵称" json:"nick_name"`
	Note          string                                         `gorm:"column:note;type:varchar(500);comment:备注信息" json:"note"`
	CreateTime    time.Time                                      `gorm:"column:create_time;type:datetime;comment:创建时间" json:"create_time"`
	LoginTime     time.Time                                      `gorm:"column:login_time;type:datetime;comment:最后登录时间" json:"login_time"`
	RolesRelation []ums_admin_role_relation.UmsAdminRoleRelation `gorm:"foreignKey:AdminId" json:"roles_relation"`
	Roles         []ums_role.UmsRole                             `gorm:"many2many:ums_admin_role_relation;joinForeignKey:AdminId;joinReferences:RoleId" json:"Roles"`
}

func (umsAdmin *UmsAdmin) TableName() string {
	return "ums_admin"
}

func (umsAdmin *UmsAdmin) Create() {
	database.DB.Create(&umsAdmin)
}

func (umsAdmin *UmsAdmin) Save() (rowsAffected int64) {
	result := database.DB.Save(&umsAdmin)
	return result.RowsAffected
}

func (umsAdmin *UmsAdmin) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&umsAdmin)
	return result.RowsAffected
}

// ComparePassword 密码是否正确
func (umsAdmin *UmsAdmin) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, umsAdmin.Password)
}
