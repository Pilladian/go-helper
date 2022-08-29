package helper

import (
	"fmt"
	"net/http"
)

// Takes input strings bot_token, chatID and message and sends a Telegram message to a given chat as specific bot
func SendTelegramMessage(bot_token string, chatID string, message string) (int, error) {
	send_text := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&parse_mode=Markdown&text=%s", bot_token, chatID, message)
	re, re_err := http.Get(send_text)
	return re.StatusCode, re_err
}
