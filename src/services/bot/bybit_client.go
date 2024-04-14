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
	print(bybit.PrettyPrint(res) + "\n")

	// TODO: When retMsg is not "OK", make telegram bot to send errors
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

	print(bybit.PrettyPrint(res) + "\n")

	return &positionInfo
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
