package sms

type Driver interface {
	//Send 发送短信
	Send(phone string, message Message, config map[string]string) bool
}
