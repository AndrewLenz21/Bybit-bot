package bot_service

import (
	"bybitbot/src/config/postgres"
	"bybitbot/src/repositories/user"
	"fmt"
)

func (s *TradingOrder) ObtainUserConfiguration() (*TradingOrder, error) {
	fmt.Println("/**/")
	fmt.Println("/************STEP 3*************/")
	db := user.NewUserRepo(postgres.GetPool())
	UserConf, err := db.GetUserConfiguration(1) // we are user 1
	if err != nil {
		fmt.Println("Error obtaining user configuration:", err)
		return nil, err
	}
	user := UserConf[0]
	// Define the Trading Position
	s.TakeProfit = user.TakeProfit
	s.StopLoss = user.StopLoss
	s.TrailingStop = user.TrailingStop
	// Max Loss
	s.MaxLoss = user.MaxLossPerPosition
	// Targets
	s.FirstTarget = user.DefaultFirstTarget
	s.FirstTargetAmount = user.AmountPerFirstTarget
	s.NextTarget = user.DefaultNextTarget
	s.NextTargetAmount = user.AmountPerNextTarget
	// Leverage
	s.ShortLeverage = user.ShortLeverage.Float64
	s.LongLeverage = user.LongLeverage.Float64

	fmt.Println("Configuration Ready!!")

	return s, nil
}
