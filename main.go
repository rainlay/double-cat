package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"os"
	"strings"
)

const fanPage = "https://www.facebook.com/%E9%9B%99%E5%96%B5%E9%A3%B2%E6%96%99%E9%83%A8-104314214253783/"

func main() {

	secret := os.Getenv("LINE_SECRET")
	token := os.Getenv("LINE_TOKEN")

	r := gin.Default()

	bot, err := linebot.New(secret, token)

	if err != nil {
		return
	}

	r.POST("callback", func(c *gin.Context) {
		log.Println(c)

		events, err := bot.ParseRequest(c.Request)

		if err == linebot.ErrInvalidSignature {
			c.Status(400)
		} else {
			c.Status(200)
		}

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				// handle text message
				case *linebot.TextMessage:
					if strings.HasPrefix(message.Text, "[公告]:") {
						//announce(bot, message, event.ReplyToken)
					} else {
						replyMenuMessage(bot, message, event.ReplyToken)
					}
				}
			}
		}
	})

	// run as server
	rErr := r.Run()

	if rErr != nil {
		log.Print(rErr)
	}
}

func replyMenuMessage(bot *linebot.Client, message *linebot.TextMessage, token string) {
	switch message.Text {
	case "粉絲專頁":
		reply := "歡迎加入 Facebook 粉絲專頁 😂"
		str := "點選連結加入粉絲團 👇\n" + fanPage

		if _, err := bot.ReplyMessage(token, linebot.NewTextMessage(reply), linebot.NewTextMessage(str)).Do(); err != nil {
			log.Print(err)
		}
	case "菜單":
		if _, err := bot.ReplyMessage(token,
			linebot.NewImageMessage("https://i.imgur.com/LuTtZtu.jpg", "https://i.imgur.com/LuTtZtum.jpg")).Do(); err != nil {
			log.Print(err)
		}
	}
}

//func announce(bot *linebot.Client, message *linebot.TextMessage, token string) {
//	if _, err := bot.Multicast(linebot.NewTextMessage("大家好 我是邦尼")).Do(); err != nil {
//		log.Print(err)
//	}
//}
