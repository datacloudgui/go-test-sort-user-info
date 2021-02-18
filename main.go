package main

import (
	"fmt"
	"main/service/sort"
	"main/service/user"
	"os"
	"strings"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

const defaultPort = "8080"

func main() {

	port := os.Getenv("GO_USERS_PORT")

	if strings.TrimSpace(port) == "" {
		port = defaultPort
	}

	router := fasthttprouter.New()

	router.Handle("GET", "/api/v1/user/:id", user.GetUser)
	router.Handle("POST", "/api/v1/sort", sort.SortArray)

	server := fasthttp.Server{
		Name:           "service-controller",
		ReadBufferSize: 4096 * 3,
		Handler:        router.Handler,
	}

	fmt.Println("starting up server on port:", port)
	if err := server.ListenAndServe(":" + port); err != nil {
		fmt.Println("Error starting up the http server")
	}
}
