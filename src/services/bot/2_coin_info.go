package bot_service

import (
	"context"
	"fmt"
)

type CoinRules struct {
	Ticker    float64
	MinAmount float64
	MaxAmount float64
}

func NewTradingPosition(symbol string, side string, entry string) *TradingOrder {
	fmt.Println("/**/")
	fmt.Println("/************STEP 2*************/")

	return &TradingOrder{
		ctx:          context.Background(),
		flg_approval: true,
		Symbol:       symbol,
		Side:         side,
		Entry:        entry,
		CoinRules:    &CoinRules{},
	}
}

func (s *TradingOrder) ObtainCoinInfo() (*TradingOrder, error) {
	//from Bybit, obtain the minimun and maximum order
	service := Client.NewInstrumentsInfoService().Category("linear").Symbol(s.Symbol)
	result, err := service.Do(s.ctx)
	if err != nil {
		fmt.Println("Error al llamar a la API:", err)
		return nil, err
	}

	Ticker := result.List[0].PriceScale
	MinOrder := result.List[0].LotSizeFilter.MinOrderQty
	MaxOrder := result.List[0].LotSizeFilter.MaxOrderQty

	//fmt.Println("Decimals for coin:", Ticker)
	//fmt.Println("Minimun Order qty string:", MinOrder)
	//fmt.Println("Maximun Order qty string:", MaxOrder)

	FloatTicker := StringToFloat64(Ticker)
	Float64MinOrder := StringToFloat64(MinOrder)
	Float64MaxOrder := StringToFloat64(MaxOrder)

	// fmt.Printf("Decimals for coin: %.4f\n", FloatTicker)
	// fmt.Printf("Minimun Order qty: %.4f\n", Float64MinOrder)
	// fmt.Printf("Maximun Order qty: %.4f\n", Float64MaxOrder)

	s.CoinRules.Ticker = FloatTicker
	s.CoinRules.MinAmount = Float64MinOrder
	s.CoinRules.MaxAmount = Float64MaxOrder

	return s, nil
}
