package setting

import (
	"log"
	"time"

	"gopkg.in/ini.v1"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTImeout time.Duration

	JWTSecret string
)

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

type RabbitMQ struct {
	Username string
	PassWord string
	Host     string
}

var RedisSetting = &Redis{}
var RabbitMQSetting = &RabbitMQ{}

func Init() {
	var err error
	Cfg, err = ini.Load("/home/swag/go/src/hmdp/config/app.ini")
	if err != nil {
		log.Fatal("Fail to parse 'config/app.ini' : ", err)
	}
	Load_Base()
	Load_Server()
	Load_APP()
	Load_mq()
}

func Load_mq() {
	mapTo("rabbitmq", RabbitMQSetting)
}

func Load_Base() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func Load_Server() {
	sec := Cfg.Section("server")
	// if err != nil {
	// 	log.Fatal("Fail to get section 'server': %v", err)
	// }

	mapTo("redis", RedisSetting)

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8081)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTImeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

func Load_APP() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'server' : %v", err)
	}
	sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
}

func mapTo(section string, v interface{}) {
	err := Cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.Maptp %s err: %v", section, err)
	}
}
