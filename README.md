# go-test-sort-user-info
Sort a Json array and return info about user.

This project show how to manage http request using golang and packages.
Also a simple custom ORM is programmed to access to a postgres DB.

## Requirements

* Go installed
* You need a Postgres DB with the password: secret running in the 5432 port or docker installed.

## Usage

1. Clone this repository by https or SSH

2. build the packages with
```go

go build ./...
```
This can take some minutes according to your internet speed

3. Load the environment variables with
```
$ source .envrc.example
```

4. setup the Postgres DB
If you have a docker installation you only need run this command in a terminal

```
$ docker run --rm --name core-postgres -p 5432:5432 -e POSTGRES_PASSWORD=secret -d postgres
```

5. Run the upgrade migration, this create the needed tables and fill some test data. See config/migrations/01_initial_setup
```go

// upgrade migrations
go run ./migrations --action=upgrade
```

6. Run the app with
```go

go run main.go
```

7. use the endpoints of the app

**NOTE:** You can use the go-test.postman_collection.json file provided by importing to your postman.

* Sort endpoint - POST
```

localhost:8080/api/v1/sort
```

Data to send
```json
{
    "unsorted": [3,5,5,6,8,3,4,4,7,7,1,1,2]
}
```

Data returned
```json
{
"response":{
    "Unsorted":[3,5,5,6,8,3,4,4,7,7,1,1,2],
    "Sorted":[1,2,3,4,5,6,7,8,5,3,4,7,1]
    },
"error":null
}
```

Note the repeated values at the end

* Get information about a user - GET

To get information a query that join information about the company, the role and the user is performed.
```

localhost:8080/api/v1/user/:id
```
two user are stored in the DB you can replace :id by 1 or 2 and get a JSON data like:

```json
{
    "response": {
        "ID": "1",
        "CompanyName": "first software company",
        "RoleName": "backend Dev",
        "Name": "Guillermo",
        "LastName": "Sanchez",
        "Email": "guillermo@datacloudgui.com",
        "IsActive": true
    },
    "error": null
}
```

8. You also can run a simple test created with:
```
go test ./... -coverprofile=coverage.out -v
```

### Migrations

in order to run migrations script apply

```go

// upgrade migrations
go run ./migrations --action=upgrade

// downgrade migrations
go run ./migrations --action=downgrade
```

### Test

A basic test is implemented, you can run it with:
```go
// Run test with verbose
go test ./... -coverprofile=coverage.out -v
// See the coverage report
go tool cover -html=coverage.out
```