package user

import (
	"database/sql"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

/********GENERAL STRUCT********/
type UserRepo struct {
	pool *pgxpool.Pool
}

func NewUserRepo(pool *pgxpool.Pool) *UserRepo {
	return &UserRepo{pool: pool}
}

/******************************/

/***********ENTITIES***********/
/*
On the entities section we are going to use the object and mapper.
- OBJECT type
- MAPPER rows Next() Scan()
After a SELECT QUERY we are going o receive pgx.rows object.
So on the final function we are going to return the mapper.
*/

/*****************************/
type UserTradingConf struct {
	IDConfiguration      int             `db:"id_configuration"`
	UserID               int             `db:"user_id"`
	PairID               int             `db:"pair_id"`
	AmountPerPosition    sql.NullFloat64 `db:"amount_per_position"`     // NUMERIC(6,4) se mapea a float64
	MaxLossPerPosition   sql.NullFloat64 `db:"max_loss_per_position"`   // NUMERIC(6,4)
	TakeProfit           sql.NullFloat64 `db:"take_profit"`             // NUMERIC(6,4)
	StopLoss             sql.NullFloat64 `db:"stop_loss"`               // NUMERIC(6,4)
	TrailingStop         sql.NullFloat64 `db:"trailing_stop"`           // NUMERIC(6,4), considera si realmente debería ser float64
	DefaultFirstTarget   sql.NullFloat64 `db:"default_first_target"`    // NUMERIC(6,4)
	AmountPerFirstTarget sql.NullFloat64 `db:"amount_per_first_target"` // NUMERIC(6,4)
	DefaultNextTarget    sql.NullFloat64 `db:"default_next_target"`     // NUMERIC(6,4)
	AmountPerNextTarget  sql.NullFloat64 `db:"amount_per_next_target"`  // NUMERIC(6,4)
	LongLeverage         sql.NullFloat64 `db:"long_leverage"`           // NUMERIC(4,2)
	ShortLeverage        sql.NullFloat64 `db:"short_leverage"`          // NUMERIC(4,2)
}

func MapUserTradingConf(rows pgx.Rows) ([]UserTradingConf, error) {
	var config []UserTradingConf
	for rows.Next() {
		var result UserTradingConf
		if err := rows.Scan(
			&result.IDConfiguration,
			&result.UserID,
			&result.PairID,
			&result.AmountPerPosition,
			&result.MaxLossPerPosition,
			&result.TakeProfit,
			&result.StopLoss,
			&result.TrailingStop,
			&result.DefaultFirstTarget,
			&result.AmountPerFirstTarget,
			&result.DefaultNextTarget,
			&result.AmountPerNextTarget,
			&result.LongLeverage,
			&result.ShortLeverage,
		); err != nil {
			return nil, err
		}
		config = append(config, result)
	}
	return config, nil
}
