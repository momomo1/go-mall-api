package v1

import "go-mall-api/models/ums_menu"

type UserReply struct {
	Icon     string             `json:"icon"`
	Username string             `json:"username"`
	Roles    []ums_menu.UmsMenu `json:"roles"`
	Menus    string             `json:"menus"`
}
