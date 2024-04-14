package bot_service

import (
	"bybitbot/src/config/postgres"
	"bybitbot/src/repositories/positions"
	"context"
	"fmt"
)

func ControlOpenPositions(symbol string, side string, createType string) {
	ctx := context.Background()
	Positions := BybitGetPositions(symbol, ctx)
	// Look for the right side position
	if Positions == nil {
		fmt.Println("Failed to get positions")
		return
	}

	db := positions.NewPositionsRepo(postgres.GetPool())
	flg_active := true
	if createType == "CreateByUser" {
		// In this case we will have an openned position
		for _, position := range Positions.List {
			if position.Side == side {
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
			}
		}
	}

	if createType == "CreateByClosing" {
		user_db_positions, err := db.GetUserPositions(symbol, side) //We will have 1, because we just have two sides, Sell or Buy
		if err != nil {
			fmt.Println("Error obtaining user open positions:", err)
			return
		}
		//Bybit opened positions
		positions_on_bybit := 0
		index_position := -1
		for index, position := range Positions.List {
			if position.Side == side {
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
			//If there is no targets, open another two
		}

		if len(user_db_positions) > positions_on_bybit && index_position == -1 {
			//Control the last Pnl Order
			//We should close the position and update the realized pnl
			print("UPDATE DB AND SET flg_active FALSE and UPDATE THE REALIZED PNL")
		}

	}
	//On Database, using this position, look for the 'reduceOnly' 'New' orders ('New' means that they are not 'filled')
	//Plus that reduceOnly orders qty, and obtain the qty available RETURN the coin with side and qty available to get the targets

	//With qty value, obtain the 3 targets

}
