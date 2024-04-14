package postgres

/*
CONECTION TO OUR POSTGRESQL DATABASE
*/

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Our global pool
var pool *pgxpool.Pool

/*******CONNECTION FILE*******/
// Pool is our pgx database connection to allow multiple connections
// LoadEnviroment: we are going to obtain our enviroment variables

// Create the connection pool
func CreateConnectionPool() {
	creds := LoadEnvironment()

	fmt.Println("Creating Database Connection pool")
	//CREATE POSTGRES URL CONNECTION
	dbService := &DbUrlStructService{c: creds}
	urlStruct, err := dbService.CreateUrl(context.Background())
	if err != nil {
		fmt.Printf("Error creating database URL: %v\n", err)
	}

	//CREATE POOL
	mypool := ConnectPostgresPool(urlStruct.url)
	if mypool != nil {
		fmt.Println("Successfully connected to PostgreSQL pool.")
		// SET THE POOL VALUE
		pool = mypool
	}
}

// PGX CONN
func ConnectPostgresPool(url string) *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		fmt.Printf("Failed to connect to PostgreSQL pool: %v\n", err)
		return nil
	}

	return pool
}

// Obtain the created connection pool
func GetPool() *pgxpool.Pool {
	return pool
}

func NewQueryService(pool *pgxpool.Pool) *QueryService {
	//default timeout 30 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	return &QueryService{pool: pool, ctx: ctx, cancel: cancel}
}
