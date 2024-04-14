package positions

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

/********GENERAL STRUCT********/
type PositionsRepo struct {
	pool *pgxpool.Pool
}

func NewPositionsRepo(pool *pgxpool.Pool) *PositionsRepo {
	return &PositionsRepo{pool: pool}
}

/***********ENTITIES***********/
/*
On the entities section we are going to use the object and mapper.
- OBJECT type
- MAPPER rows Next() Scan()
After a SELECT QUERY we are going o receive pgx.rows object.
So on the final function we are going to return the mapper.
*/

/*****************************/
type UserOpenPositions struct {
	IdPosition     int    `db:"id_position"`
	Side           string `db:"side"`
	PairId         int    `db:"pair_id"`
	AmountContract string `db:"amount_contract"` // VARCHAR (100)
}

func MapUserOpenPositions(rows pgx.Rows) ([]UserOpenPositions, error) {
	var config []UserOpenPositions
	for rows.Next() {
		var result UserOpenPositions
		if err := rows.Scan(
			&result.IdPosition,
			&result.Side,
			&result.PairId,
			&result.AmountContract,
		); err != nil {
			return nil, err
		}
		config = append(config, result)
	}
	return config, nil
}
