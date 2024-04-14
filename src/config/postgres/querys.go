package postgres

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

/*WE ARE GOING TO CREARE OUR GENERAL QUERYS*/

type QueryService struct {
	pool   *pgxpool.Pool
	ctx    context.Context
	cancel context.CancelFunc
}

// We can change timeout
func (s *QueryService) WithTimeout(timeout time.Duration) *QueryService {
	s.ctx, s.cancel = context.WithTimeout(s.ctx, timeout)
	return s
}

/*****GENERAL QUERYS*****/
//NOTE: Insert, Update and Delete are similar, but we are just change the timeout and the error message
//ATTENTION: We are calling FUNCTIONS from POSTGRES, not STORED PROCEDURES

// INSERT QUERY
func (s *QueryService) InsertCall(function string, args ...any) (string, error) {
	//Begin Transaction
	tx, err := s.pool.Begin(s.ctx)
	if err != nil {
		return "", fmt.Errorf("begin transaction failed => %w", err)
	}

	defer func() {
		if p := recover(); p != nil || err != nil {
			//If there is an error, rollback the transaction
			if rbErr := tx.Rollback(s.ctx); rbErr != nil {
				fmt.Printf("Transaction rollback failed => %v\n", rbErr)
			}
			if p != nil {
				panic(p) // Re-throw panic after rollback
			}
		} else {
			//Else commit the transaction
			err = tx.Commit(s.ctx)
		}
	}()
	//format ($1, $2 ...)
	//construct args
	argsString := constructArgs(args...)

	//prepare the QUERY
	query := fmt.Sprintf("SELECT dbo.%s(%s)", function, argsString)
	var resultMsg string
	err = tx.QueryRow(s.ctx, query, args...).Scan(&resultMsg)
	if err != nil {
		return "", err
	}

	return resultMsg, nil
}

// UPDATE QUERY
func (s *QueryService) UpdateCall(function string, args ...any) (string, error) {
	//Begin Transaction
	tx, err := s.pool.Begin(s.ctx)
	if err != nil {
		return "", fmt.Errorf("begin transaction failed => %w", err)
	}

	defer func() {
		if p := recover(); p != nil || err != nil {
			//If there is an error, rollback the transaction
			if rbErr := tx.Rollback(s.ctx); rbErr != nil {
				fmt.Printf("Transaction rollback failed => %v\n", rbErr)
			}
			if p != nil {
				panic(p) // Re-throw panic after rollback
			}
		} else {
			//Else commit the transaction
			err = tx.Commit(s.ctx)
		}
	}()
	//format ($1, $2 ...)
	//construct args
	argsString := constructArgs(args...)

	//prepare the QUERY
	query := fmt.Sprintf("SELECT dbo.%s(%s)", function, argsString)
	var resultMsg string
	err = tx.QueryRow(s.ctx, query, args...).Scan(&resultMsg)
	if err != nil {
		return "", err
	}

	return resultMsg, nil
}

// SELECT QUERY (we are goin to  need the entity map)
func (s *QueryService) SelectCall(function string, args ...any) (pgx.Rows, error) {
	//for select function we don't need transactions
	argsString := constructArgs(args...)
	query := fmt.Sprintf("SELECT * FROM dbo.%s(%s)", function, argsString)
	//call query
	rows, err := s.pool.Query(s.ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	return rows, nil
}

func constructArgs(args ...any) string {
	argPlaceholders := make([]string, len(args))
	for i := range args {
		argPlaceholders[i] = fmt.Sprintf("$%d", i+1)
	}
	argsString := strings.Join(argPlaceholders, ",")
	return argsString
}
