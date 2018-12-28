package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func getStrEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("Key not found: " + key))
	}
	return val
}

func getIntEnv(key string) int64 {
	val := getStrEnv(key)
	ret, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		panic(err)
	}
	return ret
}

func handler(w http.ResponseWriter, r *http.Request, bot *tgbotapi.BotAPI) {
	messages, ok := r.URL.Query()["message"]

	if !ok || len(messages[0]) < 1 {
		return
	}

	message := messages[0]

	chatID := getIntEnv("CHAT_ID")

	msg := tgbotapi.NewMessage(chatID, message)
	bot.Send(msg)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\": \"ok\"}"))

}

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.RemoveWebhook()

	http.HandleFunc("/handler", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, bot)
	})

	if err := http.ListenAndServe(":5050", nil); err != nil {
		log.Panic(err)
	}

	log.Print("Server started")
}
