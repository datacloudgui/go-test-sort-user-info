// Package orm provide DB connection and custom queries
// this file is used to to define user CRUD operations and custom DB manage for users table
package orm

import (
	"context"
	"errors"
	"fmt"
)

// A User represent an employee of a software company
// User ..
type User struct {
	ID          string
	CompanyName string
	RoleName    string
	Name        string
	LastName    string
	Email       string
	IsActive    bool
}

// GetUser ...
func GetUser(userID string) (*User, error) {

	fmt.Println("Getting user from DB: ")

	response := &User{}

	conn, err := GetDBInstance()

	if err != nil {

		fmt.Println("ERR@GetUser error getting db connection", err)
		return nil, errors.New("error getting db connection")
	}

	defer conn.Close()

	// Obtain the DB info
	err = conn.QueryRowContext(context.Background(), GetUserQuery, userID).Scan(
		&response.ID,
		&response.CompanyName,
		&response.RoleName,
		&response.Name,
		&response.LastName,
		&response.Email,
		&response.IsActive,
	)
	if err != nil {

		fmt.Println("ERR@GetUser error fetching agent record from database", err)
		return nil, errors.New("error fetching agent record from database")
	}

	return response, nil

}
