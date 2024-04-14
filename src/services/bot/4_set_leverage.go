package bot_service

import (
	"fmt"
	"math"
	"strconv"
)

func (s *TradingOrder) SetLeverage() *TradingOrder {
	Positions := BybitGetPositions(s.Symbol, s.ctx)
	if Positions == nil {
		fmt.Println("Failed to get positions")
		return nil
	}
	//Control the list positions
	for _, pos := range Positions.List {
		//Verify the leverage - None when there is no positions
		if StringToFloat64(pos.Leverage) != s.LongLeverage && StringToFloat64(pos.Leverage) != s.ShortLeverage {
			// float 64 to string
			buyLeverage := strconv.FormatFloat(s.LongLeverage, 'f', -1, 64)
			sellLeverage := strconv.FormatFloat(s.ShortLeverage, 'f', -1, 64)
			BybitSetLeverage(s.Symbol, buyLeverage, sellLeverage, s.ctx)
		}
	}

	// params := map[string]interface{}{
	// 	"category":     "linear",
	// 	"symbol":       s.Symbol,
	// 	"buyLeverage":  buyLeverage,
	// 	"sellLeverage": sellLeverage,
	// }

	// BybitService := Client.NewPositionService(params)

	// coinInfo, err := BybitService.GetPositionList(s.ctx) //obtain actual leverage
	// if err != nil {
	// 	fmt.Printf("Error getting coin leverage: %v\n", err)
	// 	return nil, err
	// }

	// var positionInfo PositionInfo
	// resultJSON, err := json.Marshal(coinInfo.Result)
	// if err == nil {
	// 	json.Unmarshal(resultJSON, &positionInfo)
	// 	coinInfo := positionInfo.List[0]

	// 	if coinInfo.Leverage != buyLeverage && coinInfo.Leverage != sellLeverage {
	// 		BybitService.SetPositionLeverage(s.ctx) // If the code entried here, there is no error
	// 	}
	// }

	return s
}

func (s *TradingOrder) CompareEntryPrice() (*TradingOrder, error) {
	fmt.Println("/**/")
	fmt.Println("/************STEP 4*************/")
	/*
		ON Binance Futures Premium there are signals with no float entry point
		So our objective is to get the right entry price
	*/
	service := Client.NewTickersService().Category("linear").Symbol(s.Symbol)
	result, err := service.Do(s.ctx)
	if err != nil {
		fmt.Println("Error getting coin tickers:", err)
		return nil, err
	}

	// var tickersInfo TickersInfo
	// resultJSON, err := json.Marshal(result)
	// if err != nil {
	// 	fmt.Println("Error getting JSON of Last Price:", err)
	// 	return nil, err
	// }
	// json.Unmarshal(resultJSON, &tickersInfo)
	coinInfo := result.List[0]

	// particular cases
	// ticker is less than len(entry)
	if int(s.CoinRules.Ticker) < len(s.Entry) {
		s.Entry = s.Entry[:int(s.CoinRules.Ticker)]
	}

	//almost every ticker will be more than len(entry)
	EntryFloat64 := StringToFloat64(s.Entry)
	LastPriceFloat64 := StringToFloat64(coinInfo.LastPrice)
	//percetual diff
	perc_diff := math.Abs((LastPriceFloat64 - EntryFloat64) / EntryFloat64 * 100)

	if EntryFloat64 > 1 && LastPriceFloat64 < 1 && perc_diff > 15 {
		fmt.Println("We got decimals problems...resolving")

		EntryFloat64 = EntryFloat64 / math.Pow(10, s.CoinRules.Ticker)
		//We will have particular cases
		var multiplications_price float64 = 0
		for LastPriceFloat64 < 1 {
			LastPriceFloat64 *= 10
			multiplications_price++
		}
		LastPriceFloat64 = LastPriceFloat64 / math.Pow(10, multiplications_price) //return price before for

		var multiplications_entry float64 = 0
		for EntryFloat64 < 1 {
			EntryFloat64 *= 10
			multiplications_entry++
		}
		EntryFloat64 = EntryFloat64 / math.Pow(10, multiplications_entry) //return price before for

		if multiplications_entry > multiplications_price {
			EntryFloat64 = EntryFloat64 * math.Pow(10, float64(math.Abs(multiplications_price-multiplications_entry)))
		}
		if multiplications_entry < multiplications_price {
			EntryFloat64 = EntryFloat64 / math.Pow(10, float64(math.Abs(multiplications_price-multiplications_entry)))
		}
	} else if EntryFloat64 > 1 && LastPriceFloat64 > 1 && perc_diff > 15 {
		// If entry price is less than 0 but actual price is greater than 0
		// Do not open position, tickers and decimals will be wrong
		s.flg_approval = false
	}
	s.EntryFloat64 = EntryFloat64
	s.Entry = Float64ToString(EntryFloat64)
	fmt.Printf("Our Real Entry:  \nEntry -> %s, \nEntryFloat -> %7f, \nLastPriceFloat ->  %7f\n", s.Entry, s.EntryFloat64, LastPriceFloat64)

	return s, nil
}
