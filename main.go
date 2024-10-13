package main

import (
	"fmt"

	"github.com/baoer/QQbot/internal/voice"
)

func main() {
	str := voice.GetVoiceBase64()
	fmt.Println(str)
}
