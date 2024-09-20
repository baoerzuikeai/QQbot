package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type BotInfo struct {
	AppID uint64 `yaml:"appid"`
	Token string `yaml:"token"`
}

func Load() BotInfo {
	data, err := os.ReadFile("../../config/config.yml")
	if err != nil {
		log.Fatalf("无法读取配置文件，%v", err)
	}
	var botinfo BotInfo

	err = yaml.Unmarshal(data, &botinfo)
	if err != nil {
		log.Fatalf("解析配置文件出错，%v", err)
	}
	return botinfo
}
