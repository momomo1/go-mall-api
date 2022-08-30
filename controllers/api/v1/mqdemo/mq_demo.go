package mqdemo

import (
	"github.com/gin-gonic/gin"
	"go-mall-api/controllers/api/v1"
	"go-mall-api/pkg/mq"
	"go-mall-api/pkg/response"
)

type MqController struct {
	v1.BaseAPIController
}

func (mqDemo *MqController) PushMsg(c *gin.Context) {
	rabbit := mq.NewRabbitMq("", "", "test1")
	defer rabbit.Close()
	rabbit.SendMessage(mq.Message{Body: "这是一条普通消息"})
	response.JSON(c, gin.H{})
}
