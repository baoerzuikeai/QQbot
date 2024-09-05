package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type BotInfo struct {
	AppID uint64 `yaml:"appid"`
	Token string `yaml:"token"`
}

func Load() BotInfo {
	data, err := os.ReadFile("/home/baoer/go/src/QQbot/config/config.yml")
	if err != nil {
		log.Fatalf("无法读取配置文件，%v", err)
	}
	var botinfo BotInfo

	err = yaml.Unmarshal(data, &botinfo)
	if err != nil {
		log.Fatalf("解析配置文件出错，%v", err)
	}
	fmt.Println(botinfo)
	return botinfo
}
