package orders

import (
	"bybitbot/src/config/postgres"
	"fmt"
)

func (r *OrdersRepo) InsertNewOrder(
	pairString string,
	orderId string,
	side string,
	orderType string,
	targetPrice string,

	amountOrder string,
	execAmountOrder string,

	orderStatusString string,
	takeProfit string,
	stopLoss string,
	category string,
	createType string,

	reduceOnly bool,
) (string, error) {
	sql := postgres.NewQueryService(r.pool)
	msg, err := sql.InsertCall(
		"insert_new_order_id", // Postgres function
		1,                     //user_id_in we are user 1
		pairString,            // pair_string VARCHAR(50)
		orderId,               // order_id VARCHAR(100)
		side,                  // side VARCHAR(5)
		orderType,             // order_type VARCHAR(6)
		targetPrice,           // target_price VARCHAR(100)
		amountOrder,           // amount_order VARCHAR(100)
		execAmountOrder,       // exec_amount_order VARCHAR(100)
		orderStatusString,     // order_status_string VARCHAR(50)
		takeProfit,            // take_profit VARCHAR(100)
		stopLoss,              // stop_loss VARCHAR(100)
		category,              // category VARCHAR(100)
		createType,            // create by StopLoss, ByUser, By Closing
		reduceOnly,            // reduce_only BOOL
	)
	if err != nil {
		fmt.Printf("Error inserting new order: %v\n", err)
		return "", err
	}
	fmt.Println("PostgreSQL InsertNewOrder:", msg)
	return msg, nil
}
