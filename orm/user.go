package orm

import (
	"context"
	"errors"
	"fmt"
)

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

		fmt.Println("ERR@GetAgency error getting db connection", err)
		return nil, errors.New("error getting db connection")
	}

	defer conn.Close()

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

		fmt.Println("ERR@GetAgency error fetching agent record from database", err)
		return nil, errors.New("error fetching agent record from database")
	}

	return response, nil

}
