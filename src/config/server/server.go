package server

import (
	"context"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
)

var server *echo.Echo

// creating server with echo package
func NewServer() {
	fmt.Println("Starting server with echo...")
	server = echo.New()
}

// select the port to start server
func StartServer(port int) error {
	portStr := fmt.Sprintf(":%d", port)
	err := server.Start(portStr)
	return err
}

func GetServer() *echo.Echo {
	return server
}

// not implemented yet
func StopServer(seconds time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), seconds)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		server.Logger.Fatal(err)
		fmt.Printf("Error during server shutdown: %s\n", err)
	} else {
		fmt.Println("Server shutdown gracefully")
	}

	fmt.Println("Graceful Shutdown done!")
}
