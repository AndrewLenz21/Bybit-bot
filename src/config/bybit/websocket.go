package crypto

import (
	"bybitbot/src/controllers/handlers"
	bot_service "bybitbot/src/services/bot"
	"log"

	bybit_connector "github.com/wuhewuhe/bybit.go.api"
)

var ws *bybit_connector.WebSocket

func CreateBybitConfig() {
	keys = LoadCryptoEnvironment()
	Client = NewCryptoClient()
	//Send the client to bot
	bot_service.ObtainBybitClient(Client)
	StartWebsocketStream() // Create Websocket Stream
}

func StartWebsocketStream() {
	ws = bybit_connector.NewBybitPrivateWebSocket(keys.wsURL, keys.apiKey, keys.secretKey, myMessageHandler)
	if err := ws.Connect([]string{"order"}); err != nil {
		log.Fatal("Error connecting to WebSocket:", err)
	}
}

func CloseWebsocketStream() {
	ws.Disconnect()
}

func myMessageHandler(message string) error {
	handlers.BybitOrderHandler(message, Client)
	return nil
}
