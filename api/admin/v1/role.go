package v1

import "go-mall-api/models/ums_role"

type RoleListAllReply struct {
	Data []ums_role.UmsRole `json:"data"`
}
