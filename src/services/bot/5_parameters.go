package bot_service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"math"
)

type TradingOrder struct {
	ctx          context.Context
	flg_approval bool

	Symbol string
	Side   string //Sell or buy
	Entry  string

	EntryFloat64 float64

	MaxLoss   sql.NullFloat64
	MaxAmount sql.NullFloat64

	FullPositions  int
	PositionAmount sql.NullFloat64

	TakeProfit   sql.NullFloat64
	StopLoss     sql.NullFloat64
	TrailingStop sql.NullFloat64

	FirstTarget       sql.NullFloat64
	FirstTargetAmount sql.NullFloat64

	NextTarget       sql.NullFloat64
	NextTargetAmount sql.NullFloat64

	LongLeverage  float64
	ShortLeverage float64

	CoinRules *CoinRules
}

func (s *TradingOrder) SetOrderParameters() (*TradingOrder, error) {
	fmt.Println("/**/")
	fmt.Println("/************STEP 5*************/")
	//Obtain User USDT ammount
	USDT_acc := ObtainUserAssets(s.ctx)
	//Transform the Entry Price to float64
	Entry := StringToFloat64(s.Entry)
	//How maximun we are allowed to lose
	MaxLoss := (USDT_acc * s.MaxLoss.Float64) / 100
	//Consider that we have fees
	Taker := 0.055 //Taker: 0.0550%
	Maker := 0.020 //Maker: 0.0200%

	//OBTAIN MAKER FEE -> OBTAIN TAKER FEE -> OBTAIN LOSING AMOUNT
	var n float64 = 0
	if s.Side == "Sell" {
		n = (MaxLoss * 100) / (Entry * (Maker + (1+s.StopLoss.Float64/100)*Taker + s.StopLoss.Float64))
	}
	if s.Side == "Buy" {
		n = (MaxLoss * 100) / (Entry * (Maker + (1-s.StopLoss.Float64/100)*Taker + s.StopLoss.Float64))
	}
	// n is the exact ammount to open

	// With 'n' we have the exact number of coins to enter but we have rules
	MaxAmmountPositions := math.Floor(n / s.CoinRules.MaxAmount)         //Positions with max amount - could be zero
	PositionAmmount := n - (MaxAmmountPositions * s.CoinRules.MaxAmount) //Normal position

	if MaxAmmountPositions > 0 { //  almost always will be zero
		PositionAmmount = s.CoinRules.MaxAmount //If we are over the maxAmount, we are going to open just the max amount
	}

	AdjustedPosition := math.Floor(PositionAmmount/s.CoinRules.MinAmount) * s.CoinRules.MinAmount //Our exact position

	s.FullPositions = int(MaxAmmountPositions)
	s.PositionAmount = sql.NullFloat64{Float64: AdjustedPosition, Valid: true}
	//fmt.Printf("Our position value: %.7f\n", s.PositionAmount.Float64)

	//fmt.Println("USDT Available:", USDT_acc)
	//fmt.Printf("Entry price: %.7f \nMax Loss: %.9f\n", Entry, MaxLoss)
	// The AdjustedPosition could be zero '0'
	if AdjustedPosition > 0 {
		// OPEN POSITIONS
		fmt.Printf("Qty to open position: %.7f\n", AdjustedPosition)
	}

	return s, nil
}

func ObtainUserAssets(ctx context.Context) float64 {
	wallet_string := ""

	// Obtain how much USDT we have on our account
	params := map[string]interface{}{"accountType": "UNIFIED", "coin": "USDT", "category": "linear"}
	accountResult, err := Client.NewAccountService(params).GetAccountWallet(ctx)
	if err != nil {
		return 0
	}

	var accountInfo AccountInfo
	resultJSON, err := json.Marshal(accountResult.Result)
	if err != nil {
		fmt.Println("Error getting JSON Assets:", err)
		return 0
	}
	json.Unmarshal(resultJSON, &accountInfo)
	coinInfo := accountInfo.List[0].Coin[0] // First Coin Element
	wallet_string = coinInfo.WalletBalance
	wallet := StringToFloat64(wallet_string)

	return wallet
}
