package positions

import (
	"bybitbot/src/config/postgres"
	"fmt"
)

/*INSERT OR UPDATE NEW POSITION*/
func (r *PositionsRepo) InsertNewPosition(
	symbol string,
	entryPrice string,
	side string,
	size string,

	positionValueIn string,
	positionBalanceIn string,

	stopLoss string,
	takeProfit string,
	trailingStop string,

	positionIdx int,
	unrealizedPnlIn string,
	curRealisedPnlIn string,
	cumRealisedPnlIn string,

	flgActive bool,
) (string, error) {
	sql := postgres.NewQueryService(r.pool)
	msg, err := sql.InsertCall(
		"upsert_position", // Postgres function
		symbol,            // pair_string VARCHAR(50)
		entryPrice,        // entry_price VARCHAR(100)
		side,              // side VARCHAR(5)
		size,              // size VARCHAR(6)

		positionValueIn,
		positionBalanceIn,

		stopLoss,     // target_price VARCHAR(100)
		takeProfit,   // amount_order VARCHAR(100)
		trailingStop, // exec_amount_order VARCHAR(100)

		positionIdx,      // order_status_string VARCHAR(50)
		unrealizedPnlIn,  // take_profit VARCHAR(100)
		curRealisedPnlIn, // stop_loss VARCHAR(100)
		cumRealisedPnlIn, // category VARCHAR(100)

		flgActive, // reduce_only BOOL
	)
	if err != nil {
		fmt.Printf("Error inserting new position: %v on %s\n", err, symbol)
		return "", err
	}
	fmt.Println("PostgreSQL InsertNewPosition:", msg)
	return msg, nil
}

/*CLOSE POSITION*/
func (r *PositionsRepo) ClosePosition(symbol string, side string, last_pnl string) (string, error) {
	sql := postgres.NewQueryService(r.pool)
	msg, err := sql.UpdateCall(
		"close_position", // Postgres function
		symbol,           // pair_string VARCHAR(50)
		side,             // side VARCHAR(5)
		last_pnl,         // size VARCHAR(6)
	)
	if err != nil {
		fmt.Printf("Error closing position: %v on %s\n", err, symbol)
		return "", err
	}
	fmt.Println("PostgreSQL ClosePosition:", msg)
	return msg, nil
}

/*OBTAIN POSITIONS*/
func (r *PositionsRepo) GetUserPositions(symbol string, side string) ([]UserOpenPositions, error) {
	sql := postgres.NewQueryService(r.pool)
	rows, err := sql.SelectCall("get_user_open_positions", symbol, side)
	if err != nil {
		return nil, err
	}
	//MAP THE ROWS TO THE ENTITY
	return MapUserOpenPositions(rows)
}
