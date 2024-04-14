package orders

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

/********GENERAL STRUCT********/
type OrdersRepo struct {
	pool *pgxpool.Pool
}

func NewOrdersRepo(pool *pgxpool.Pool) *OrdersRepo {
	return &OrdersRepo{pool: pool}
}
