package bot_service

import (
	"context"
	"encoding/json"
	"fmt"

	bybit "github.com/wuhewuhe/bybit.go.api"
)

// Our Bybit Client
var Client *bybit.Client

// This will be called by the opening configuration
func ObtainBybitClient(Bybit_client *bybit.Client) {
	Client = Bybit_client
}

func BybitOpenOrder(symbol string, side string, entry string, qty string, stop_loss string, reduce_only bool, ctx context.Context) {
	Order := Client.NewPlaceOrderService("linear", symbol, side, "limit", qty)
	Order.Price(entry)
	Order.StopLoss(stop_loss)
	Order.ReduceOnly(reduce_only)

	res, err := Order.Do(ctx)
	if err != nil {
		fmt.Println("Error doing Order.Do:", err)
	}

	if res == nil {
		print("ERROR")
	}
	//print(bybit.PrettyPrint(res) + "\n")

	// TODO: When retMsg is not "OK", make telegram bot to send errors
}

func BybitGetOpenOrders(symbol string, side string, ctx context.Context) *BybitOpenOrders {
	params := map[string]interface{}{
		"category":    "linear",
		"symbol":      symbol,
		"orderFilter": "Order",
		"openOnly":    0, //'0' is open only
		"limit":       3,
	}

	Order := Client.NewTradeService(params)
	res, err := Order.GetOpenOrders(ctx)
	if err != nil {
		fmt.Println("Error doing Order.GetOpenOrders:", err)
		return nil
	}

	var OpenOrders BybitOpenOrders
	resultJSON, err := json.Marshal(res.Result)
	if err != nil {
		fmt.Println("Error parsing json Open Orders:", err)
		return nil
	}
	json.Unmarshal(resultJSON, &OpenOrders)

	//filter just the orders we need to test
	//print(bybit.PrettyPrint(res) + "\n")
	filteredOrders := []OrderData{}
	for _, order := range OpenOrders.Data {
		if order.ReduceOnly && order.Side != side {
			filteredOrders = append(filteredOrders, order)
		}
	}
	OpenOrders.Data = filteredOrders

	return &OpenOrders
}

func BybitGetPositions(symbol string, ctx context.Context) *PositionInfo {
	params := map[string]interface{}{
		"category": "linear",
		"symbol":   symbol,
	}
	Position := Client.NewPositionService(params)
	res, err := Position.GetPositionList(ctx)
	if err != nil {
		fmt.Println("Error doing Position.GetPositionList:", err)
		return nil
	}

	var positionInfo PositionInfo
	resultJSON, err := json.Marshal(res.Result)
	if err != nil {
		fmt.Println("Error parsing json Position:", err)
		return nil
	}
	json.Unmarshal(resultJSON, &positionInfo)

	//print(bybit.PrettyPrint(res.Result) + "\n")

	return &positionInfo
}

func BybitGetLastPnl(symbol string, ctx context.Context) string {
	params := map[string]interface{}{
		"category": "linear",
		"symbol":   symbol,
		"limit":    1,
	}
	Pnl := Client.NewPositionService(params)
	res, err := Pnl.GetClosePnl(ctx)
	if err != nil {
		fmt.Println("Error doing Pnl.GetClosePnl:", err)
		return ""
	}

	var lastPnl BybitLastPnl
	resultJSON, err := json.Marshal(res.Result)
	if err != nil {
		fmt.Println("Error parsing json Position:", err)
		return ""
	}
	json.Unmarshal(resultJSON, &lastPnl)

	//print(bybit.PrettyPrint(res) + "\n")

	return lastPnl.Data[0].ClosedPnl
}

func BybitSetLeverage(symbol string, buyLeverage string, sellLeverage string, ctx context.Context) {
	params := map[string]interface{}{
		"category":     "linear",
		"symbol":       symbol,
		"buyLeverage":  buyLeverage,
		"sellLeverage": sellLeverage,
	}
	Position := Client.NewPositionService(params)
	Position.SetPositionLeverage(ctx)
}

func BybitSetAutoAddMargin(symbol string, autoAddMargin int, positionIdx int, ctx context.Context) {
	params := map[string]interface{}{
		"category":      "linear",
		"symbol":        symbol,
		"autoAddMargin": autoAddMargin,
		"positionIdx":   positionIdx,
	}
	Margin := Client.NewPositionService(params)
	res, err := Margin.SetPositionAutoMargin(ctx)
	if err != nil {
		fmt.Println("Error changing positionIdx:", err)
		return
	}
	print(bybit.PrettyPrint(res) + "\n")
}
