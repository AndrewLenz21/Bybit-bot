package crypto

import (
	"bybitbot/src/controllers/handlers"
	bot_service "bybitbot/src/services/bot"
	"fmt"
	"log"
	"time"

	bybit_connector "github.com/wuhewuhe/bybit.go.api"
)

var ws *bybit_connector.WebSocket

func CreateBybitConfig() {
	keys = LoadCryptoEnvironment()
	Client = NewCryptoClient()

	bot_service.ObtainBybitClient(Client) //Send the client to bot
	StartWebsocketStream()                // Create Websocket Stream
	RestartWebsocketStream()              // Restart Websocket Stream
}

func RestartWebsocketStream() {
	//Every 9 minutes, restart the channel websocket
	ticker := time.NewTicker(9 * time.Minute)
	go func() {
		for {
			<-ticker.C
			CloseWebsocketStream()
			fmt.Println("Restarting websocket")
			StartWebsocketStream() // Restart WebSocket Stream
		}
	}()
}

func StartWebsocketStream() {
	ws = bybit_connector.NewBybitPrivateWebSocket(keys.wsURL, keys.apiKey, keys.secretKey, myMessageHandler, bybit_connector.WithMaxAliveTime("600s"))
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
