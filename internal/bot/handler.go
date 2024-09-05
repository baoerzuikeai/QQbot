package bot

import (
	"context"
	"log"

	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/openapi"
)

func SendAtMessage(ctx context.Context, api openapi.OpenAPI) func(event *dto.WSPayload, data *dto.WSATMessageData) error {
	return func(event *dto.WSPayload, data *dto.WSATMessageData) error {
		if data.Content == "" {
			log.Printf("Message doesn't sent to channel ")
			return nil
		}
		// 获取消息内容和频道ID
		channelID := data.ChannelID // 替换为实际的频道ID
		messageContent := "hello world"

		// 发送消息
		// _, err := api.PostMessage(ctx, channelID, message)
		err := api.CreateMessageReaction(ctx, channelID, data.ID, dto.Emoji{
			ID:   "4",
			Type: 1,
		})
		if err != nil {
			log.Printf("Failed to send message: %v", err)
			return err
		}

		log.Printf("Message sent to channel %s: %s", channelID, messageContent)
		return nil
	}
}
