package test

import (
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

/********GENERAL STRUCT********/
type TestRepo struct {
	pool *pgxpool.Pool
}

func NewTestRepo(pool *pgxpool.Pool) *TestRepo {
	return &TestRepo{pool: pool}
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
type TestTableTodo struct {
	ID         int
	NumberTest int
	NumberPlus int
	InsDate    time.Time
}

func MapTestTableTodo(rows pgx.Rows) ([]TestTableTodo, error) {
	var results []TestTableTodo
	for rows.Next() {
		var result TestTableTodo
		if err := rows.Scan(&result.ID, &result.NumberTest, &result.NumberPlus, &result.InsDate); err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}

/*****************************/
