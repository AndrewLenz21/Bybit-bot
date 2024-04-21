package bot_service

import (
	"bybitbot/src/config/postgres"
	"bybitbot/src/repositories/positions"
	"bybitbot/src/repositories/user"
	"context"
	"fmt"
	"math"
	"time"
)

func ControlOpenPositions(symbol string, side string, createType string, targetPrice string, reduceOnly bool) {
	ctx := context.Background()
	Positions := BybitGetPositions(symbol, ctx)
	// Look for the right side position
	if Positions == nil {
		fmt.Println("Failed to get positions")
		return
	}
	db := positions.NewPositionsRepo(postgres.GetPool())

	flg_active := true
	index_position := -1

	var oppositeSide string = ""
	if side == "Sell" {
		oppositeSide = "Buy"
	}
	if side == "Buy" {
		oppositeSide = "Sell"
	}

	switch createType {
	case "CreateByUser":
		if !reduceOnly {
			for index, position := range Positions.List {
				// Side will be "Sell" or "Buy"
				if position.Side == side {
					index_position = index
				}
			}
			if index_position != -1 {
				// When then order is "Filled" and is "CreateByUser", is a new open position
				// We do not need to control if we have a new position
				position := Positions.List[index_position] //obtain the correct index
				if position.AutoAddMargin == 1 {
					BybitSetAutoAddMargin(position.Symbol, 0, position.PositionIdx, ctx)
				}
				db.InsertNewPosition(position.Symbol,
					position.EntryPrice,
					position.Side,
					position.Size,
					position.PositionValue,
					position.PositionBalance,
					position.StopLoss,
					position.TakeProfit,
					position.TrailingStop,
					position.PositionIdx,
					position.UnrealizedPnl,
					position.CurRealizedPnl,
					position.CumRealisedPnl,
					flg_active)
				//OBTAIN AVAILABLE ASSETS TO OPEN NEW TARGETS
				var avl_size float64 = StringToFloat64(position.Size)
				BybitOpenOrders := BybitGetOpenOrders(symbol, side, ctx)
				//SUM THE qty orders
				for _, order := range BybitOpenOrders.Data {
					// reduce only are targets
					if order.ReduceOnly {
						avl_size -= StringToFloat64(order.Qty)
					}
				}
				//SET NEW TARGETS USING THE "targetPrice", "side" and "size"
				//OpenNewTargets(symbol, position.Size, side, targetPrice, ctx)
				OpenNewTargets(symbol, Float64ToString(avl_size), oppositeSide, targetPrice, ctx)
			}
		} else if reduceOnly { //This is a target, we need to use the OppositeSide
			//Position on our database
			user_db_positions, err := db.GetUserPositions(symbol, oppositeSide) //We will have 1, because we just have two sides, Sell or Buy
			if err != nil {
				fmt.Println("Error obtaining user open positions:", err)
				return
			}
			//Bybit opened positions
			positions_on_bybit := 0
			index_position := -1
			for index, position := range Positions.List {
				if position.Side == oppositeSide {
					positions_on_bybit++
					index_position = index
				}
			}
			//Do we have an opened position with this side?
			if len(user_db_positions) == positions_on_bybit && index_position != -1 {
				//just update the current position
				position := Positions.List[index_position] //obtain the correct index
				db.InsertNewPosition(position.Symbol,
					position.EntryPrice,
					position.Side,
					position.Size,
					position.PositionValue,
					position.PositionBalance,
					position.StopLoss,
					position.TakeProfit,
					position.TrailingStop,
					position.PositionIdx,
					position.UnrealizedPnl,
					position.CurRealizedPnl,
					position.CumRealisedPnl,
					flg_active)

				//CONTROL THE TARGETS
				BybitOpenOrders := BybitGetOpenOrders(symbol, oppositeSide, ctx)
				//If there is no targets, open NEW TARGETS
				if len(BybitOpenOrders.Data) == 0 {
					OpenNewTargets(symbol, position.Size, side, targetPrice, ctx)
				}
			}

			//Position has closed but we have still opened on database
			if len(user_db_positions) > positions_on_bybit {
				flg_active = false
				//Control the last Pnl Order
				LastPnl := BybitGetLastPnl(symbol, ctx)
				//We should close the position and update the realized pnl
				if LastPnl != "" {
					db.ClosePosition(symbol, oppositeSide, LastPnl)
				}
			}
		}

	case "CreateByClosing":
		//Position on our database
		user_db_positions, err := db.GetUserPositions(symbol, oppositeSide) //We will have 1, because we just have two sides, Sell or Buy
		if err != nil {
			fmt.Println("Error obtaining user open positions:", err)
			return
		}
		//Bybit opened positions
		positions_on_bybit := 0
		index_position := -1
		for index, position := range Positions.List {
			if position.Side == oppositeSide {
				positions_on_bybit++
				index_position = index
			}
		}
		//Do we have an opened position with this side?
		if len(user_db_positions) == positions_on_bybit && index_position != -1 {
			//just update the current position
			position := Positions.List[index_position] //obtain the correct index
			db.InsertNewPosition(position.Symbol,
				position.EntryPrice,
				position.Side,
				position.Size,
				position.PositionValue,
				position.PositionBalance,
				position.StopLoss,
				position.TakeProfit,
				position.TrailingStop,
				position.PositionIdx,
				position.UnrealizedPnl,
				position.CurRealizedPnl,
				position.CumRealisedPnl,
				flg_active)

			//CONTROL THE TARGETS
			BybitOpenOrders := BybitGetOpenOrders(symbol, oppositeSide, ctx)
			//If there is no targets, open NEW TARGETS
			if len(BybitOpenOrders.Data) == 0 {
				OpenNewTargets(symbol, position.Size, side, targetPrice, ctx)
			}
		}

		if len(user_db_positions) > positions_on_bybit {
			flg_active = false
			//Control the last Pnl Order
			LastPnl := BybitGetLastPnl(symbol, ctx)
			//We should close the position and update the realized pnl
			if LastPnl != "" {
				db.ClosePosition(symbol, oppositeSide, LastPnl)
			}
		}
	case "CreateByStopLoss":
		flg_active = false
		//Control the last Pnl Order
		time.Sleep(500 * time.Millisecond)
		LastPnl := BybitGetLastPnl(symbol, ctx)
		//We should close the position and update the realized pnl
		if LastPnl != "" {
			db.ClosePosition(symbol, oppositeSide, LastPnl)
		}
	default:
		println("error on receiving ''filled'' targets")
	}
}

func OpenNewTargets(symbol string, size string, side string, entryPrice string, ctx context.Context) {
	//OBTAIN THE NEW THREE TARGETS
	targets := GetNewTargets(symbol, size, side, entryPrice, ctx)

	//open new orders
	for _, target := range targets {
		price := Float64ToString(target.Price)
		qty := Float64ToString(target.Qty)
		// Reduce Only on true
		BybitOpenOrder(symbol, side, price, qty, "", true, ctx)
	}
}

func GetNewTargets(symbol string, size_string string, side string, entryPrice_string string, ctx context.Context) []Target {
	//OBTAIN USER CONFIGURATION
	db := user.NewUserRepo(postgres.GetPool())
	UserConf, err := db.GetUserConfiguration(1) // we are using configuration bot 1
	if err != nil {
		fmt.Println("Error obtaining user configuration:", err)
		return nil
	}
	//OBTAIN COIN RULES
	service := Client.NewInstrumentsInfoService().Category("linear").Symbol(symbol)
	result, err := service.Do(ctx)
	if err != nil {
		fmt.Println("Error al llamar a la API:", err)
		return nil
	}

	//GET VALUES TO USE
	User := UserConf[0]

	MinOrder := StringToFloat64(result.List[0].LotSizeFilter.MinOrderQty)

	Size := StringToFloat64(size_string)
	EntryPrice := StringToFloat64(entryPrice_string)

	//Targets
	var targets []Target

	var tg float64 = EntryPrice
	var QtyLeft float64 = Size

	for i := 0; i < 3 && QtyLeft > 0; i++ {
		defaultTarget := User.DefaultNextTarget.Float64
		amountPerTarget := User.AmountPerNextTarget.Float64

		if i == 0 { //Just for first target
			defaultTarget = User.DefaultFirstTarget.Float64
			amountPerTarget = User.AmountPerFirstTarget.Float64
		}

		//fmt.Printf("Perc Target %7f: Amount Target = %7f, CURRENT tg = %5f\n", defaultTarget, amountPerTarget, tg)

		next_price, qty := ObtainTarget(tg, side, QtyLeft, defaultTarget, amountPerTarget, MinOrder)
		tg = next_price
		fmt.Printf("Target %d: Price = %7f, Qty = %5f\n", i, tg, qty)
		if tg == 0 || qty == 0 {
			//error
			return nil
		}
		//ADD ON A TUPLE
		targets = append(targets, Target{
			Id:    i,
			Price: tg,
			Qty:   qty,
		})
		//END TUPLE
		QtyLeft -= qty
	}

	return targets
}

func ObtainTarget(entryPrice float64, side string, position_size float64, target float64, target_size float64, order_rule float64) (float64, float64) {
	var Price float64 = 0
	//OBTAIN  TARGET
	if side == "Buy" {
		Price = entryPrice * (1 - target/100)
	}
	if side == "Sell" {
		Price = entryPrice * (1 + target/100)
	}
	//fmt.Printf("Our next price is Price = %7f\n", Price)
	//OBTAIN AMOUNT
	Qty := position_size * (target_size / 100)
	Amount := math.Floor(Qty/order_rule) * order_rule

	//Control qty
	if Amount == 0 {
		Amount = position_size //This will close position (hidden take profit)
	}

	return Price, Amount
}
