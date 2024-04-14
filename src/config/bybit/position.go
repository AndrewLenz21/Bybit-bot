package crypto

import (
	bybit "github.com/wuhewuhe/bybit.go.api"
)

var Client *bybit.Client
var Service *bybit.InstrumentsInfoService

func NewCryptoClient() *bybit.Client {
	return bybit.NewBybitHttpClient(keys.apiKey, keys.secretKey, bybit.WithBaseURL(bybit.MAINNET))
}
