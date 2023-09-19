package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/streadway/amqp"
	lua "github.com/wagfog/hmdp_go/config/Lua"
	"github.com/wagfog/hmdp_go/config/gredis"
	"github.com/wagfog/hmdp_go/config/rabbitmq"
	"github.com/wagfog/hmdp_go/dto/result"
	"github.com/wagfog/hmdp_go/models"
	"github.com/wagfog/hmdp_go/utils"
)

type VoucherOrderService struct {
}

var seckillHash string

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func publish(body []byte) {
	err := rabbitmq.Ch.Publish(
		"",                // exchange
		rabbitmq.Que.Name, // routing key
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	failOnError(err, "Failed to publish a message")
}

func consume() {
	msgs, err := rabbitmq.Ch.Consume(
		rabbitmq.Que.Name, // queue
		"",                // consumer
		true,              // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	failOnError(err, "Failed to register a consumer")

	for msg := range msgs {
		mesg := make(map[string]string)
		json.Unmarshal(msg.Body, &mesg)
		id, err := strconv.Atoi(mesg["id"])
		failOnError(err, "exchange error!")
		Voucherid, err := strconv.Atoi(mesg["voucherId"])
		failOnError(err, "exchange error!")
		userid, err := strconv.Atoi(mesg["userId"])
		failOnError(err, "exchange error!")
		vOrder := models.VoucherOrder{
			ID:        int64(id),
			VoucherID: int64(Voucherid),
			UserID:    int64(userid),
		}
		HandelVoucher(vOrder)
	}
}

func HandelVoucher(voucherOrder models.VoucherOrder) {

}

func CreateSript() {
	//返回的脚本会产生一个sha1哈希值,下次用的时候可以直接使用这个值
	var err error
	seckillHash, err = gredis.Client.ScriptLoad(context.Background(), lua.Seckkill).Result()
	if err != nil {
		panic(err)
	}
}

func NewVoucherOrderService(vs *VoucherOrderService) bool {
	if vs != nil {
		return false
	}
	vs = &VoucherOrderService{}
	return true
}

func (v *VoucherOrderService) SeckillVoucher(voucherId int, phone string) result.Result {
	//获取用户
	u := models.GetUserByPhone(phone)
	//获取订单
	orderID := utils.NextId("order")
	id := strconv.Itoa(int(u.ID))
	vid := strconv.Itoa(voucherId)
	sorderID := strconv.Itoa(int(orderID))
	n, err := gredis.Client.EvalSha(context.Background(), seckillHash, []string{id, vid, sorderID}).Result()

	if err != nil {
		panic(err)
	}
	if n == 1 {
		return *result.Fail("库存不足")
	} else if n == 2 {
		return *result.Fail("重复下单")
	}
	mesg := make(map[string]string)
	mesg["userId"] = id
	mesg["voucherId"] = vid
	mesg["id"] = sorderID
	body, _ := json.Marshal(mesg)
	publish(body)
	return *result.OkWithData(orderID)
}

func (v *VoucherOrderService) CreateVoucherOrder(voucherOrder models.VoucherOrder, phone string) result.Result {
	u := models.GetUserByPhone(phone)
	flag := models.CreateVoucherOrder(int(u.ID), voucherOrder)
	if flag {
		return *result.Ok()
	}
	fmt.Println("库存不足！")
	return *result.Fail("库存不足！")
}
