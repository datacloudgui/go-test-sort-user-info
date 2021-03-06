package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	//Postgres migration driver
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	//migration file
	_ "github.com/golang-migrate/migrate/v4/source/file"
	//Postgres general driver
	_ "github.com/lib/pq"
)

// manage the migration commands for upgrade and downgrade. See readme.md
func main() {

	var action string
	flag.StringVar(&action, "action", "", "Usage")
	flag.Parse()

	fmt.Println("Selected action: " + action + "!")

	m, err := migrate.New(
		os.Getenv("MIGRATIONS_PATH"),
		os.Getenv("CORE_DATABASE_URI"))

	if err != nil {
		fmt.Println("Error starting migration", err)
		return
	}

	if action == "upgrade" {

		if err := m.Up(); err != nil {

			fmt.Println("Error", err)

		} else {

			fmt.Println(action, "run successfuly")
		}

	} else if action == "downgrade" {

		if err := m.Down(); err != nil {

			fmt.Println("Error", err)

		} else {

			fmt.Println(action, "run successfuly")
		}

	} else {

		fmt.Println("action", action, "not valid")

	}

}
