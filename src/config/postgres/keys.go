package postgres

/*******OUR PGX CONFIGURATION*******/

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Globals
type Credentials struct {
	db_name string
	db_port string
	db_user string
	db_pwd  string
}

type DbUrlStructService struct {
	c    *Credentials
	name string
	url  string
}

/*******CONFIGURATION FILE********/
// Create the URL structure with the enviroment (.env) variables
// Create the connection with pgx pool

// URL STRUCT
func (d *DbUrlStructService) CreateUrl(ctx context.Context) (client *DbUrlStructService, err error) {
	missing := []string{}
	//verify that we have all values
	if d.c.db_name == "" {
		missing = append(missing, "DATABASE_NAME")
	}
	if d.c.db_port == "" {
		missing = append(missing, "DATABASE_PORT")
	}
	if d.c.db_user == "" {
		missing = append(missing, "DATABASE_USER")
	}
	if d.c.db_pwd == "" {
		missing = append(missing, "DATABASE_PWD")
	}
	//if we have one missing = ERROR
	if len(missing) > 0 {
		return nil, fmt.Errorf("missing required fields in .env file: %v", missing)
	}
	//else return the URL
	d.name = d.c.db_name
	d.url = fmt.Sprintf("postgresql://%s:%s@%s/%s", d.c.db_user, d.c.db_pwd, d.c.db_port, d.c.db_name)
	//print(d.url)
	return d, nil
}

// Load enviroment variables
func LoadEnvironment() *Credentials {
	godotenv.Load()
	// LOAD ENVIROMENT
	client := &Credentials{
		db_name: os.Getenv("DATABASE_NAME"),
		db_port: os.Getenv("DATABASE_PORT"),
		db_user: os.Getenv("DATABASE_USER"),
		db_pwd:  os.Getenv("DATABASE_PWD"),
	}

	return client
}
