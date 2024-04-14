package bot_service

// Response
type BybitResponse struct {
	Data []OrderData `json:"data"`
}

// Order Data
type OrderData struct {
	Symbol    string `json:"symbol"`
	OrderID   string `json:"orderId"`
	Side      string `json:"side"`
	OrderType string `json:"orderType"`
	Price     string `json:"price"`

	Qty        string `json:"qty"`
	CumExecQty string `json:"cumExecQty"`

	OrderStatus string `json:"orderStatus"`
	TakeProfit  string `json:"takeProfit"`
	StopLoss    string `json:"stopLoss"`
	Category    string `json:"category"`
	CreateType  string `json:"createType"`

	ReduceOnly bool `json:"reduceOnly"`
}

// our account assets info
type AccountInfo struct {
	List []struct {
		Coin []struct {
			AvailableToWithdraw string `json:"availableToWithdraw"`
			Equity              string `json:"equity"`
			WalletBalance       string `json:"walletBalance"`
		} `json:"coin"`
	} `json:"list"`
}

// our account assets info
type PositionInfo struct {
	List []struct {
		Symbol          string `json:"symbol"`
		Side            string `json:"side"`
		Size            string `json:"size"`
		Leverage        string `json:"leverage"`
		EntryPrice      string `json:"avgPrice"`
		PositionValue   string `json:"positionValue"`   //value with leverage
		PositionBalance string `json:"positionBalance"` //value with no leverage
		StopLoss        string `json:"stopLoss"`
		TakeProfit      string `json:"takeProfit"`
		TrailingStop    string `json:"trailingStop"`
		AutoAddMargin   int    `json:"autoAddMargin"`
		PositionIdx     int    `json:"positionIdx"`
		UnrealizedPnl   string `json:"unrealizedPnl"`
		CurRealizedPnl  string `json:"curRealisedPnl"`
		CumRealisedPnl  string `json:"cumRealisedPnl"`
	} `json:"list"`
}

// the coin tickers info
type TickersInfo struct {
	List []struct {
		Symbol      string `json:"symbol"`
		LastPrice   string `json:"lastPrice"`
		IndexPrice  string `json:"indexPrice"`
		MarketPrice string `json:"markPrice"`
	} `json:"list"`
}
