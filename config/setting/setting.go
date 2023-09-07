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

func init() {
	var err error
	Cfg, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatal("Fail to parse 'config/app.ini' : %v ", err)
	}
	Load_Base()
	Load_Server()
	Load_APP()
}

func Load_Base() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func Load_Server() {
	sec, err := Cfg.Section("server")
	if err != nil {
		log.Fatal("Fail to get section 'server': %v", err)
	}
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8081)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTImeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func Load_APP() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatal("Fail to get section 'server' : %v", err)
	}
	sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
}
