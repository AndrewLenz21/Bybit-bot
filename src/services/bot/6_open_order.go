package bot_service

import (
	"fmt"
)

// WE ARE GOING TO OPEN NEW ORDER USING BYBIT CLIENT
func (s *TradingOrder) OpenNewOrder() (*TradingOrder, error) {
	fmt.Println("/**/")
	fmt.Println("/********STEP 6: open order *********/")
	//Define Stop Loss String
	var SL_float64 float64
	if s.Side == "Buy" {
		SL_float64 = s.EntryFloat64 * ((100 - s.StopLoss.Float64) / 100)
	}
	if s.Side == "Sell" {
		SL_float64 = s.EntryFloat64 * ((100 + s.StopLoss.Float64) / 100)
	}
	Stop_Loss := Float64ToString(SL_float64)
	Qty := Float64ToString(s.PositionAmount.Float64)

	//Always sue a stop loss
	if s.flg_approval && StringToFloat64(Stop_Loss) > 0 {
		// fmt.Println("WE HAVE OUR POSITION READY:")
		// fmt.Println("Symbol:", s.Symbol)
		// fmt.Println("Side:", s.Side)
		// fmt.Println("Entry price:", s.Entry)
		// fmt.Println("Stop Loss:", Stop_Loss)
		// fmt.Println("Quantity:", Qty)
		// fmt.Printf("LongLeverage: %f\nShortLeverage: %f\n", s.LongLeverage, s.ShortLeverage)

		//WE CAN CALL THE BYBIT API - reduce_only is false
		order_id := BybitOpenOrder(s.Symbol, s.Side, s.Entry, Qty, Stop_Loss, false, s.ctx) //we will recieve the order ID
		fmt.Printf("ORDER -> %s \n", order_id)
		//After call the bybit API we will recieve msg from websocket
	} else {
		fmt.Printf("Order on -> %s \n Side -> %s \n NOT APPROVED", s.Symbol, s.Side)
		//fmt.Println("POSSITION NOT APPROVED")
	}

	return s, nil

}
