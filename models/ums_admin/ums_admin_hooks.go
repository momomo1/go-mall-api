package ums_admin

import (
	"go-mall-api/pkg/hash"
	"gorm.io/gorm"
)

// func (umsAdmin *UmsAdmin) BeforeSave(tx *gorm.DB) (err error) {}
// func (umsAdmin *UmsAdmin) BeforeCreate(tx *gorm.DB) (err error) {}
// func (umsAdmin *UmsAdmin) AfterCreate(tx *gorm.DB) (err error) {}
// func (umsAdmin *UmsAdmin) BeforeUpdate(tx *gorm.DB) (err error) {}
// func (umsAdmin *UmsAdmin) AfterUpdate(tx *gorm.DB) (err error) {}
// func (umsAdmin *UmsAdmin) AfterSave(tx *gorm.DB) (err error) {}
// func (umsAdmin *UmsAdmin) BeforeDelete(tx *gorm.DB) (err error) {}
// func (umsAdmin *UmsAdmin) AfterDelete(tx *gorm.DB) (err error) {}
// func (umsAdmin *UmsAdmin) AfterFind(tx *gorm.DB) (err error) {}
func (userModel *UmsAdmin) BeforeSave(tx *gorm.DB) (err error) {
	if !hash.BcryptIsHashed(userModel.Password) {
		userModel.Password = hash.BcryptHash(userModel.Password)
	}
	return
}
