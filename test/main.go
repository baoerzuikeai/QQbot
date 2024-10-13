package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/baoer/QQbot/internal/bot"
	"github.com/baoer/QQbot/my_dto"
	client "github.com/baoer/QQbot/websocket"
	"github.com/tencent-connect/botgo/dto"
)

var s uint32
var session_id string

func main() {
	// WebSocket 服务器的 URL
	c := client.Connect()
	defer c.Close()
	_, msg, err := c.ReadMessage()
	fmt.Printf("Received message: %s\n", msg)
	if err != nil {
		log.Println("read:", err)
		return
	}
	// #登录鉴权获得 Session
	session_id = client.Login(c)
	log.Println(session_id)
	//发送心跳 Ack
	ctx, cancel := context.WithCancel(context.Background())
	go client.SendHeartBeat(c, ctx)
	//读取响应消息
	fmt.Println("读取响应消息")
	for {
		var wspayload dto.WSPayload
		_, msg, err = c.ReadMessage()
		log.Printf("Received message:%s \n", msg)
		if err != nil {
			log.Println("read:", err)
			return
		}
		err = json.Unmarshal(msg, &wspayload)
		if err != nil {
			log.Println("解析群监听消息出错:", err)
			return
		}
		if wspayload.OPCode == dto.WSReconnect {
			cancel()
			c.Close()
			c = client.Connect()
			client.SendReLogin(s, session_id, c)
			ctx, cancel = context.WithCancel(context.Background())
			go client.SendHeartBeat(c, ctx)
			continue
		}
		if wspayload.Data == nil {
			continue
		}

		switch v := wspayload.Data.(type) {
		case string:
			log.Println(v)
			continue
		default:
		}
		s = wspayload.Seq
		groupmessagemap := wspayload.Data.(map[string]interface{})
		if groupmessagemap["id"] == nil || groupmessagemap["group_openid"] == nil {
			continue
		}
		var gm my_dto.GroupMessage
		id := groupmessagemap["id"].(string)
		group_openid := groupmessagemap["group_openid"].(string)
		content := groupmessagemap["content"].(string)
		gm.MsgID = id
		gm.GroupOpenid = group_openid
		gm.Content = content
		go bot.SendGroupAtMessage(gm)
	}
}
