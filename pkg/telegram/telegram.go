package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"pingocean-front/users-service/internal/config"
	"pingocean-front/users-service/pkg/logger"
	"strconv"
)

type Message struct {
	Header      string
	Datetime    string
	Payload     string
	UserAgent   string
	Ip          string
	CompanyName string
	Email       string
}

type Telegram struct {
	messages chan Message
	cfg      *config.Config
	bot      *tgbotapi.BotAPI
}

func NewTelegram(cfg *config.Config) (*Telegram, error) {
	bot, err := tgbotapi.NewBotAPI(cfg.Telegram.BotToken)
	if err != nil {
		return nil, err
	}

	t := &Telegram{
		messages: make(chan Message, 50),
		cfg:      cfg,
		bot:      bot,
	}

	go t.sender()

	return t, nil
}

func (t *Telegram) Send(msg Message) {
	t.messages <- msg
}

func (t *Telegram) sender() {
	chatId, err := strconv.Atoi(t.cfg.Telegram.ChatId)
	if err != nil {
		logger.Error(fmt.Sprintf("telegram: invalid chat id - %d", chatId))
		return
	}

	for message := range t.messages {
		msg := tgbotapi.NewMessage(int64(chatId), getFormattedMessage(message))

		_, err = t.bot.Send(msg)
		if err != nil {
			logger.Error(fmt.Sprintf("telegram sender: can`t to send to channel new message, message - %v", message))
		}
	}
}

func getFormattedMessage(message Message) string {
	return fmt.Sprintf("Выявлено подозрительное действие (%s):\n🔸 TIME: %s\n🔸 PAYLOAD: %s\n🔸 USER-AGENT: %s\n🔸 IP: %s\n🔸 Company: %s\n🔸 Email: %s",
		message.Header, message.Datetime, message.Payload, message.UserAgent, message.Ip, message.CompanyName, message.Email)
}
