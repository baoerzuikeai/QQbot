package bot

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/baoer/QQbot/internal/pixiv"
	"github.com/baoer/QQbot/internal/voice"
	"github.com/baoer/QQbot/my_dto"
)

func SendGroupAtMessage(gm my_dto.GroupMessage) error {
	switch {
	case strings.HasPrefix(gm.Content, " /帮助 "):
	case strings.HasPrefix(gm.Content, " /随机图片 "):
		if err := func() error {
			gm.Content = "✌️🥵✌️"
			gm.MsgType = 7
			gm.Media = PostFile(gm)
			media := gm.Media.(my_dto.Media)
			PostGroupMessage(gm)
			if media.FileUuid == "" {
				gm.Content = "\n图片发送出错,请重试🤡👉🏻🤡"
				gm.MsgType = 0
				gm.Media = nil
				err := PostGroupMessage(gm)
				if err != nil {
					log.Println("发送出错", err)
					return err
				}
				// gm.Media = PostFile(gm)
				// media = gm.Media.(my_dto.Media)
				// PostGroupMessage(gm)
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
	case strings.HasPrefix(gm.Content, " /聊天 "):
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
	case strings.HasPrefix(gm.Content, " /搜索图片 "):
		if err := func() error {
			sep := strings.Split(gm.Content, " ")
			if sep[2] == " " {
				gm.Content = "参数非法,请输入正确格式 /搜索图片 (keywords)"
				gm.MsgType = 0
				err := PostGroupMessage(gm)
				if err != nil {
					log.Println(err)
					return err
				}
				return nil
			}
			gm.Media = PostSerchFile(gm, sep[2]+" 10000users入り")
			gm.Content = "\n✌️🥵✌️"
			gm.MsgType = 7
			PostGroupMessage(gm)
			media := gm.Media.(my_dto.Media)
			if media.FileUuid == "" {
				// gm.Media = PostSerchFile(gm, sep[1])
				// media = gm.Media.(*my_dto.Media)
				// PostGroupMessage(gm)
				gm.Content = "\n未找到匹配图片，请更换关键词🤡👉🏻🤡"
				gm.MsgType = 0
				gm.Media = nil
				err := PostGroupMessage(gm)
				if err != nil {
					log.Println("发送出错", err)
					return err
				}
			}
			return nil
		}(); err != nil {
			log.Println(err)
			return err
		}
	case strings.HasPrefix(gm.Content, " /晚安 "):
		if err := func() error {
			gm.Content = ""
			gm.Media = PostVoiceFile(gm)
			gm.MsgType = 7
			PostGroupMessage(gm)
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
	}
	filedata := pixiv.Getimage(c)
	for len(filedata) < 100 {
		filedata = pixiv.Getimage(c)
	}
	imagedata.FileData = filedata
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

func PostSerchFile(gm my_dto.GroupMessage, name string) my_dto.Media {
	c := pixiv.InitClient()
	defer c.CloseIdleConnections()
	imagedata := my_dto.PostMedia{
		FileType:   1,
		SrvSendMsg: false,
	}
	filedata := pixiv.GetSerchimage(c, name)
	for len(filedata) < 100 {
		filedata = pixiv.GetSerchimage(c, name)
		if len(filedata) == 0 {
			break
		}
	}
	imagedata.FileData = filedata
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

func PostVoiceFile(gm my_dto.GroupMessage) my_dto.Media {
	voicedata := my_dto.PostMedia{
		FileType: 3,
	}
	voicedata.FileData = voice.GetVoiceBase64()
	jsonData, err := json.Marshal(voicedata)
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
