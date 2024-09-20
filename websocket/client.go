package client

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/tencent-connect/botgo/dto"
)

func Connect() *websocket.Conn {
	serverURL := "wss://api.sgroup.qq.com/websocket/"
	// 建立 WebSocket 连接
	c, _, err := websocket.DefaultDialer.Dial(serverURL, nil)
	if err != nil {
		log.Fatal("Error connecting to WebSocket server:", err)
	}
	return c
}

func Login(c *websocket.Conn) string {
	wspayload := dto.WSPayload{
		WSPayloadBase: dto.WSPayloadBase{
			OPCode: 2,
		},
		Data: dto.WSIdentityData{
			Token:   "Bot 102340632.ER0AW2JMtAA1G8PweWfXMjGZOoOCpbXB",
			Intents: 33554432,
			Shard:   []uint32{0, 1},
		},
	}
	err := c.WriteJSON(wspayload)
	if err != nil {
		log.Println("发送登录OpCode 2 Identify出错", err)
	}
	_, msg, err := c.ReadMessage()
	if err != nil {
		log.Println("接收Ready Event出错", err)
	}
	var readyDate dto.WSPayload
	err = json.Unmarshal(msg, &readyDate)
	if err != nil {
		log.Println("解析Ready Event出错", err)
	}
	log.Println(readyDate)
	readyMap := readyDate.Data.(map[string]interface{})
	session_id := readyMap["session_id"].(string)
	return session_id
}

func SendReLogin(s uint32, session_id string, c *websocket.Conn) {
	wsResume := dto.WSPayload{
		WSPayloadBase: dto.WSPayloadBase{
			OPCode: 6,
		},
		Data: dto.WSResumeData{
			Token:     "Bot 102340632.ER0AW2JMtAA1G8PweWfXMjGZOoOCpbXB",
			SessionID: session_id,
			Seq:       s,
		},
	}
	c.WriteJSON(wsResume)
}

func SendHeartBeat(c *websocket.Conn, ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// 发送心跳消息
			err := c.WriteMessage(websocket.TextMessage, []byte("heartbeat"))
			if err != nil {
				log.Println("error sending heartbeat:", err)
				return
			}
			log.Println("Sent heartbeat")
		case <-ctx.Done():
			// 终止心跳
			log.Println("Stopping heartbeat")
			return
		}
	}
}
