package mq

import (
	"fmt"
	"github.com/streadway/amqp"
	"go-mall-api/pkg/config"
	"go-mall-api/pkg/logger"
	"log"
	"strconv"
	"time"
)

// Message 消息体: DelayTime 仅在 SendDelayMessage 方法有效
type Message struct {
	DelayTime int // desc:延迟时间
	Body      string
}

type MessageQueue struct {
	conn         *amqp.Connection //amqp连接对象
	ch           *amqp.Channel    //channel对象
	ExchangeName string           //交换机名称
	RouteKey     string           //路由名称
	QueueName    string           //队列名称
}

// Consumer 消费者回调方法
type Consumer func(delivery amqp.Delivery)

// NewRabbitMq 新建 rabbitmq 实例
func NewRabbitMq(exchange, route, queue string) MessageQueue {
	var messageQueue = MessageQueue{
		ExchangeName: exchange,
		RouteKey:     route,
		QueueName:    queue,
	}
	conn, err := amqp.Dial(fmt.Sprintf("amqp:%s:%s@%s:%s/", config.GetString("mq.username"), config.GetString("mq.password"), config.GetString("mq.host"), config.GetString("mq.port")))
	failOnError(err, "Dial")
	messageQueue.conn = conn

	//建立channel通道
	ch, err := conn.Channel()
	failOnError(err, "Channel")
	messageQueue.ch = ch

	//声明 exchange 交换器
	return messageQueue
}

// SendMessage 发送普通消息
func (mq *MessageQueue) SendMessage(message Message) {
	q := mq.declareQueue(mq.QueueName, nil)

	err := mq.ch.Publish(
		mq.ExchangeName, //exchange
		q.Name,
		true,
		false,
		amqp.Publishing{
			Timestamp:    time.Now(),
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(message.Body),
		},
	)
	failOnError(err, "send common msg err")
}

// SendDelayMessage 发送延迟消息
func (mq *MessageQueue) SendDelayMessage(message Message) {
	delayQueueName := mq.QueueName + "_delay" + strconv.Itoa(message.DelayTime)
	delayRouteKey := mq.RouteKey + "_delay" + strconv.Itoa(message.DelayTime)

	// 定义延迟队列(死信队列)
	dq := mq.declareQueue(
		delayQueueName,
		amqp.Table{
			"x-dead-letter-exchange":    mq.ExchangeName, // 指定死信交换机
			"x-dead-letter-routing-key": mq.RouteKey,     // 指定死信routing-key
		},
	)

	// 延迟队列绑定到exchange
	mq.bindQueue(dq.Name, delayRouteKey, mq.ExchangeName)

	// 发送消息，将消息发送到延迟队列，到期后自动路由到正常队列中
	err := mq.ch.Publish(
		mq.ExchangeName,
		delayRouteKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message.Body),
			Expiration:  strconv.Itoa(message.DelayTime * 1000),
		},
	)
	failOnError(err, "send delay msg err")
}

// Consume 获取消费消息
func (mq *MessageQueue) Consume(fn Consumer) {
	//声明队列
	q := mq.declareQueue(mq.QueueName, nil)

	//队列绑定到exchange
	mq.bindQueue(q.Name, mq.RouteKey, mq.ExchangeName)

	// 设置Qos
	err := mq.ch.Qos(1, 0, false)
	failOnError(err, "Qos")

	// 监听消息
	msgs, err := mq.ch.Consume(
		q.Name, // queue name,
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Consume")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fn(d)
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}

// Close 关闭链接
func (mq *MessageQueue) Close() {
	mq.ch.Close()
	mq.conn.Close()
}

// declareQueue 定义队列
func (mq *MessageQueue) declareQueue(name string, args amqp.Table) amqp.Queue {
	q, err := mq.ch.QueueDeclare(
		name,
		true,
		false,
		false,
		false,
		args,
	)

	failOnError(err, "QueueDeclare")

	return q
}

// declareExchange 定义交换器
func (mq *MessageQueue) declareExchange(exchange string, kind string, args amqp.Table) {
	err := mq.ch.ExchangeDeclare(
		exchange,
		kind,
		true,
		false,
		false,
		false,
		args,
	)
	failOnError(err, "ExchangeDeclare")
}

// bindQueue 绑定队列
func (mq *MessageQueue) bindQueue(queue, routekey, exchange string) {
	err := mq.ch.QueueBind(
		queue,
		routekey,
		exchange,
		false,
		nil,
	)
	failOnError(err, "QueueBind")
}

// failOnError 错误处理
func failOnError(err error, name string) {
	if err != nil {
		logger.ErrorString("RabbitMQ", name, err.Error())
	}
}
