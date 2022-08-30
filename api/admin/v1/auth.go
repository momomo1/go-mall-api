package v1

type LoginRequest struct {
	Username string `form:"username" binding:"required`
	Password string `form:"password" binding:"required`
}

type LoginReply struct {
	Data struct {
		Token string
	}
}
