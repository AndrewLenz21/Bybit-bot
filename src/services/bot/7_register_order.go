package bot_service

import (
	"bybitbot/src/config/postgres"
	"bybitbot/src/repositories/orders"
	"encoding/json"
	"fmt"
)

// The bybit wesocket handler will arrive here
func RegisterOrder(message string) {
	print(message + "\n")
	var orderMessage BybitResponse
	if err := json.Unmarshal([]byte(message), &orderMessage); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	//data := orderMessage.Data[0] //ERROR: data can have two elements, the first element usually will be the stop loss
	// For each element in Data we will register the orders
	for _, data := range orderMessage.Data {
		fmt.Println("Symbol:", data.Symbol)
		fmt.Println("OrderID:", data.OrderID)
		fmt.Println("Side:", data.Side)
		fmt.Println("OrderType:", data.OrderType)
		fmt.Println("Price:", data.Price)

		fmt.Println("Qty:", data.Qty)
		fmt.Println("Qty:", data.CumExecQty)

		fmt.Println("OrderStatus:", data.OrderStatus)
		fmt.Println("TakeProfit:", data.TakeProfit)
		fmt.Println("StopLoss:", data.StopLoss)
		fmt.Println("Category:", data.Category)
		fmt.Println("CreateType:", data.CreateType)

		fmt.Println("ReduceOnly:", data.ReduceOnly)

		//Call OrdersRepo
		db := orders.NewOrdersRepo(postgres.GetPool())
		db.InsertNewOrder(data.Symbol, data.OrderID, data.Side, data.OrderType, data.Price, data.Qty, data.CumExecQty, data.OrderStatus, data.TakeProfit, data.StopLoss, data.Category, data.CreateType, data.ReduceOnly)

		//When order is 'filled', 'reduceOnly' is false and 'OrderType' is Limit, then this is an open position
		//Verify the open position of symbol recieved from the websocket connection with the correct direction 'Buy' or 'Sell'
		if data.OrderType == "Limit" && !data.ReduceOnly && data.OrderStatus == "Filled" {
			//We got an open position
			ControlOpenPositions(data.Symbol, data.Side, data.CreateType)
		}

	}

}
