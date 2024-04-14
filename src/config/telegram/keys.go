package telegram

import (
	"os"

	"github.com/joho/godotenv"
)

// Globals
var keys *Credentials

type Credentials struct {
	apiKey    string
	secretKey string
}

func LoadTelegramEnvironment() *Credentials {
	godotenv.Load()
	// LOAD ENVIRONMENT
	keys = &Credentials{
		apiKey:    os.Getenv("APP_ID"),
		secretKey: os.Getenv("APP_HASH"),
	}
	return keys
}
