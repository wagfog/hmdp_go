package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
	"github.com/wagfog/hmdp_go/config/setting"
)

var Mq_Conn *amqp.Connection
var Ch *amqp.Channel
var Que amqp.Queue

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Init_Rabbitmq() {
	url := fmt.Sprintf("amqp://%s:%s@%s/",
		setting.RabbitMQSetting.Username,
		setting.RabbitMQSetting.PassWord,
		setting.RabbitMQSetting.Host,
	)
	var err error
	Mq_Conn, err = amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")
	Ch, err = Mq_Conn.Channel()
	failOnError(err, "Failed to declare a channel")

	Que, err = Ch.QueueDeclare(
		"voucher_seckill", // name
		false,             // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	failOnError(err, "Failed to declare a queue")
}
