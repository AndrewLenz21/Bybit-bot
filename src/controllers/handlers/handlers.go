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

func TelegramMessageHandler(sender string, tgmessage tg.MessageClass) {
	msg := tgmessage.(*tg.Message)

	var channelID int64
	if sender == "User" { //message from user
		peerChannel := msg.PeerID.(*tg.PeerUser)
		channelID = peerChannel.UserID
	}
	if sender == "Channel" { //message from channel
		peerChannel := msg.PeerID.(*tg.PeerChannel)
		channelID = peerChannel.ChannelID
	}

	bot_service.SignalReceived(sender, msg.ID, channelID, msg.Date, msg.Message)
}
