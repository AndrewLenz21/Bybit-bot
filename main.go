package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	crypto "bybitbot/src/config/bybit"
	"bybitbot/src/config/postgres"
	"bybitbot/src/config/server"
	"bybitbot/src/config/telegram"
)

// START server
// ONLY FOR EDUCATIONAL PURPOSE, TRADING IS DANGEROUS
func main() {
	postgres.CreateConnectionPool()
	go ctrlC() //closing signal with Ctrl + C

	crypto.CreateBybitConfig()      //client and websocket bybit
	telegram.CreateTelegramConfig() //client and websocket telegram

	server.NewServer()
	server.StartRoutes()
	server.StartServer(8080) //port 8080
}

// SHUTDOWN server
func ctrlC() {
	sigChannel := make(chan os.Signal, 1)

	signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM)
	<-sigChannel
	fmt.Println("Recieved interrupt signal, closing connections...")

	postgres.GetPool().Close()     //close postgres connection
	crypto.CloseWebsocketStream()  //close websocket connection
	telegram.CloseTelegramClient() //close telegram client
	//server.StopServer(5)

	os.Exit(0)
}
