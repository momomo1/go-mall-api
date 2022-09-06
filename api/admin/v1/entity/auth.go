package entity

type LoginRequest struct {
	UserName string `json:"username,omitempty" valid:"username"`
	Password string `json:"password,omitempty" valid:"password"`
}

type LoginReply struct {
	Token     string `json:"token"`
	TokenHead string `json:"tokenHead"`
}
