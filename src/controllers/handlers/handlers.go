package handlers

import (
	bot_service "bybitbot/src/services/bot"

	"github.com/gotd/td/tg"
	bybit_connector "github.com/wuhewuhe/bybit.go.api"
)

//package to recieve the calls and parse the object inside messages

func BybitOrderHandler(message string, Client *bybit_connector.Client) {
	bot_service.RegisterOrder(message)
}

func TelegramMessageHandler(userId int64, tgmessage *tg.Message, username string) {
	// ID is channel ID
	bot_service.SignalReceived(userId, username, tgmessage.ID, tgmessage.Message, tgmessage.Date)
}
