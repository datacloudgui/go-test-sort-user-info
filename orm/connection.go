// Package orm provide DB connection and custom queries
// this file allow the connection to the DB form environment variables
package orm

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"strconv"
)

var instance *sql.DB
var err error

var maxDBConnections, _ = strconv.Atoi(os.Getenv("CORE_DB_CONNECTIONS"))
var maxIdleDBConnections, _ = strconv.Atoi(os.Getenv("CORE_MAX_IDLE_DB_CONNECTIONS"))

// GetDBInstance ...
func GetDBInstance() (*sql.Conn, error) {

	if instance == nil {

		instance, err = sql.Open("postgres", os.Getenv("CORE_DATABASE_URI"))

		if err != nil {
			return nil, err
		}

		instance.SetMaxOpenConns(maxDBConnections)
		instance.SetMaxIdleConns(maxIdleDBConnections)
	}

	con, err := instance.Conn(context.Background())

	if err != nil {

		fmt.Println("Error getting connection from pool", err)
		return nil, errors.New("Error getting conneciton from pool")
	}

	return con, nil
}
