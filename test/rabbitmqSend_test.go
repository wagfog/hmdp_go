package test

import (
	"log"
	"testing"

	"github.com/streadway/amqp"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s %s", msg, err.Error())
	}
}

var ch *amqp.Channel
var q amqp.Queue

func TestRabbit(t *testing.T) {
	// 1. 尝试连接RabbitMQ，建立连接
	// 该连接抽象了套接字连接，并为我们处理协议版本协商和认证等。
	conn, err := amqp.Dial("amqp://swag:123456@localhost:5672/")
	FailOnError(err, "Failed to open a channel")
	defer conn.Close()

	// 2. 接下来，我们创建一个通道，大多数API都是用过该通道操作的。
	ch, err = conn.Channel()
	FailOnError(err, "Failed to opne a channel")
	defer ch.Close()

	// 3. 声明消息要发送到的队列
	q, err = ch.QueueDeclare(
		"seckill_test",
		false,
		false,
		false,
		false,
		nil,
	)

	FailOnError(err, "Failed to declare a queue")

	body := "hello world!"
	// 4.将消息发布到声明的队列
	err = ch.Publish(
		"",     //""：表示交换机的名称，这里为空字符串表示使用默认的交换机。
		q.Name, //表示路由键，即指定消息要发送到的队列。
		false,  //mandatory,表示是否要求消息必须被写入到至少一个队列中，如果没有符合的队列，消息会被丢弃。
		false,  // 表示是否要求消息必须立即被消费，如果没有消费者立即消费该消息，消息会被丢弃。
		amqp.Publishing{ //表示要发布的消息的属性。
			ContentType: "text/plain", //表示消息的内容类型，这里设置为"text/plain"，表示纯文本类型的消息。
			Body:        []byte(body),
		},
	)
	FailOnError(err, "Failed to publish a message")
}
