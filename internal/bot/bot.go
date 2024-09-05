package bot

import (
	"context"
	"time"

	"github.com/baoer/QQbot/internal/config"
	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/openapi"
	"github.com/tencent-connect/botgo/token"
)

func Initbot() (openapi.OpenAPI, context.Context, *token.Token) {
	conf := config.Load()
	botToken := token.BotToken(conf.AppID, conf.Token)
	api := botgo.NewOpenAPI(botToken).WithTimeout(3 * time.Second)
	ctx := context.Background()
	return api, ctx, botToken
}

// var atMessage event.ATMessageEventHandler = func(event *dto.WSPayload, data *dto.WSATMessageData) error {

// }
