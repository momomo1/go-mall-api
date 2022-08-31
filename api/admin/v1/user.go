package v1

type UserReply struct {
	Icon     string        `json:"icon"`
	Username string        `json:"username"`
	Menus    []interface{} `json:"menus"`
	Roles    []string      `json:"roles"`
}
