package entity

type OrderSettingReply struct {
	Id                  uint64 `json:"id"`
	CommentOvertime     int    `json:"commentOvertime"`
	ConfirmOvertime     int    `json:"confirmOvertime"`
	FinishOvertime      int    `json:"finishOvertime"`
	FlashOrderOvertime  int    `json:"flashOrderOvertime"`
	NormalOrderOvertime int    `json:"normalOrderOvertime"`
}

type OrderSettingRequest struct {
	CommentOvertime     string `json:"commentOvertime"`
	ConfirmOvertime     string `json:"confirmOvertime"`
	FinishOvertime      string `json:"finishOvertime"`
	FlashOrderOvertime  string `json:"flashOrderOvertime"`
	NormalOrderOvertime string `json:"normalOrderOvertime"`
}
