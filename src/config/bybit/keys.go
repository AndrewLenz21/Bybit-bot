package crypto

import (
	"os"

	"github.com/joho/godotenv"
	bybit_connector "github.com/wuhewuhe/bybit.go.api"
)

// Globals
var keys *Credentials

type Credentials struct {
	apiKey    string
	secretKey string
	wsURL     string
}

func LoadCryptoEnvironment() *Credentials {
	godotenv.Load()
	// LOAD ENVIROMENT
	keys := &Credentials{
		apiKey:    os.Getenv("BYBIT_API_KEY"),
		secretKey: os.Getenv("BYBIT_SECRET_KEY"),
		wsURL:     bybit_connector.WEBSOCKET_PRIVATE_MAINNET,
	}

	return keys
}
