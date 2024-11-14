package gemini

import (
	"context"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func InitClient() *genai.Client {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatal("1", err)
	}
	return client
}

func Chat(client *genai.Client, send_msg_chan, recive_msg_chan chan string) {
	ctx := context.Background()
	model := client.GenerativeModel("gemini-1.5-flash")
	cs := model.StartChat()
	cs.History = []*genai.Content{
		{
			Parts: []genai.Part{
				genai.Text("你好，你能跟我聊一会天吗？"),
			},
			Role: "user",
		},
		{
			Parts: []genai.Part{
				genai.Text("当然可以，你希望跟我聊什么话题呢？"),
			},
			Role: "model",
		},
	}
	for msg := range send_msg_chan {
		resp, err := cs.SendMessage(ctx, genai.Text(msg))
		if err != nil {
			log.Fatal(err)
		}
		var reply []genai.Part
		for _, candidate := range resp.Candidates {
			reply = append(reply, candidate.Content.Parts...)
		}
		text, ok := reply[0].(genai.Text)
		if !ok {
			log.Println("不是text")
		}
		strtext := string(text)
		recive_msg_chan <- strtext
		cs.History = append(cs.History, []*genai.Content{
			{
				Parts: []genai.Part{
					genai.Text(msg),
				},
				Role: "user",
			},
			{
				Parts: reply,
				Role:  "model",
			},
		}...)
		if len(cs.History) >= 20 {
			cs.History = cs.History[len(cs.History)-20 : len(cs.History)-1]
		}
	}
}
