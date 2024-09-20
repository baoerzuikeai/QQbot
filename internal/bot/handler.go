package bot

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"github.com/baoer/QQbot/internal/pixiv"
	"github.com/baoer/QQbot/my_dto"
)

// func SendAtMessage(ctx context.Context, api openapi.OpenAPI) func(event *dto.WSPayload, data *dto.WSATMessageData) error {
// 	return func(event *dto.WSPayload, data *dto.WSATMessageData) error {
// 		if data.Content == "" {
// 			log.Printf("Message doesn't sent to channel ")
// 			return nil
// 		}
// 		// 获取消息内容和频道ID
// 		channelID := data.ChannelID // 替换为实际的频道ID
// 		messageContent := "hello world"

// 		// 发送消息
// 		// _, err := api.PostMessage(ctx, channelID, message)
// 		err := api.CreateMessageReaction(ctx, channelID, data.ID, dto.Emoji{
// 			ID:   "4",
// 			Type: 1,
// 		})
// 		if err != nil {
// 			log.Printf("Failed to send message: %v", err)
// 			return err
// 		}

// 		log.Printf("Message sent to channel %s: %s", channelID, messageContent)
// 		return nil
// 	}
// }

func SendGroupAtMessage(gm my_dto.GroupMessage) error {
	switch gm.Content {
	case " /帮助 ":
	case " /随机图片 ":
		if err := func() error {
			gm.Content = "🥵🥵🥵🥵"
			gm.MsgType = 7
			gm.Media = PostFile(gm)
			media := gm.Media.(my_dto.Media)
			log.Println(media)
			if media.FileUuid == "" {
				gm.Content = "\n图片发送出错,请重试🤡👉🏻🤡"
				gm.MsgType = 0
				gm.Media = nil
				err := PostGroupMessage(gm)
				if err != nil {
					log.Println("发送出错", err)
					return err
				}
			}
			err := PostGroupMessage(gm)
			if err != nil {
				log.Println("发送出错", err)
				return err
			}
			return nil
		}(); err != nil {
			log.Println(err)
			return err
		}
	case " /聊天 ":
		if err := func() error {
			gm.Content = "你好"
			gm.MsgType = 0
			err := PostGroupMessage(gm)
			if err != nil {
				log.Println(err)
				return err
			}
			return nil
		}(); err != nil {
			log.Println(err)
			return err
		}
	default:
	}

	return nil
}

func PostGroupMessage(gm my_dto.GroupMessage) error {
	jsonData, err := json.Marshal(gm)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(string(jsonData))
	url := "https://sandbox.api.sgroup.qq.com/v2/groups/" + gm.GroupOpenid + "/messages"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("post出错", err)
		return err
	}
	req.Header.Set("Authorization", "Bot 102340632.ER0AW2JMtAA1G8PweWfXMjGZOoOCpbXB")
	req.Header.Set("Content-Type", "application/json")

	// 创建 HTTP 客户端
	client := &http.Client{}

	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to send request: %v", err)
		return err
	}
	log.Println(resp)
	defer resp.Body.Close()
	return nil
}

func PostFile(gm my_dto.GroupMessage) my_dto.Media {
	c := pixiv.InitClient()
	defer c.CloseIdleConnections()
	imagedata := my_dto.PostMedia{
		FileType:   1,
		SrvSendMsg: false,
		FileData:   pixiv.Getimage(c),
	}
	jsonData, err := json.Marshal(imagedata)
	if err != nil {
		log.Println("映射图片json出错:", err)
	}
	url := "https://sandbox.api.sgroup.qq.com/v2/groups/" + gm.GroupOpenid + "/files"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("post图片出错", err)
	}
	req.Header.Set("Authorization", "Bot 102340632.ER0AW2JMtAA1G8PweWfXMjGZOoOCpbXB")
	req.Header.Set("Content-Type", "application/json")

	// 创建 HTTP 客户端
	client := &http.Client{}

	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var media my_dto.Media
	json.Unmarshal(body, &media)
	return media
}
