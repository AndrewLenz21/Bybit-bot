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

//our open orders
//Getting order ID
type ServerResponse struct {
	Result OrderResult `json:"result"`
}

// OrderResult define la estructura esperada en el campo `Result`.
type OrderResult struct {
	OrderId     string `json:"orderId"`
	OrderLinkId string `json:"orderLinkId"`
}

// Response
type BybitOpenOrders struct {
	Data []OrderData `json:"list"`
}

// Order Data
type OpenOrder struct {
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
	CreateType  string `json:"createType"`

	ReduceOnly bool `json:"reduceOnly"`
}

//our last Pnl
// Response
type BybitLastPnl struct {
	Data []SymbolLastPnl `json:"list"`
}

// Order Data
type SymbolLastPnl struct {
	Symbol     string `json:"symbol"`
	OrderID    string `json:"orderId"`
	Side       string `json:"side"`
	OrderType  string `json:"orderType"`
	OrderPrice string `json:"orderPrice"`

	Qty       string `json:"qty"`
	ExecType  string `json:"execType"`
	ClosedPnl string `json:"closedPnl"`
}

//For targets
type Target struct {
	Id    int
	Price float64
	Qty   float64
}
