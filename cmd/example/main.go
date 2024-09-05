package main

import (
	"log"

	"github.com/baoer/QQbot/internal/bot"
	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/event"
	"github.com/tencent-connect/botgo/websocket"
)

func main() {
	api, ctx, botToken := bot.Initbot()

	// 获取 WebSocket 连接
	ws, err := api.WS(ctx, nil, "")
	if err != nil {
		log.Fatalf("Failed to get WebSocket connection: %v", err)
	}
	//定义事件
	var atmessage event.ATMessageEventHandler = bot.SendAtMessage(ctx, api)
	// 注册事件处理器
	intent := websocket.RegisterHandlers(atmessage)
	// 启动会话管理器并处理 WebSocket 连接
	manager := botgo.NewSessionManager()
	if err := manager.Start(ws, botToken, &intent); err != nil {
		log.Fatalf("Failed to start session manager: %v", err)
	}

	// 阻止主线程退出
	select {}

}
