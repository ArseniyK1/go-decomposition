package main

import (
	"decomposition/internal/app/handlers/echo"
	"decomposition/internal/app/handlers/ping"
	"decomposition/internal/app/middleware"
	"fmt"
	"net/http"
)

func main() {

	router := http.NewServeMux()
	router.HandleFunc("GET /ping", ping.PingHandler)
	router.HandleFunc("POST /echo", echo.EchoHandler)

	// http.HandleFunc("/ping", middleware.Middleware(ping.PingHandler))

	// http.HandleFunc("POST /echo", middleware.Middleware(echo.EchoHandler))

	err := http.ListenAndServe(":8080", middleware.Middleware(router))
	if err != nil {
		fmt.Println("Server is not running", err)
	} else {
		fmt.Println("Server is running on http://localhost:8080")
	}
}
