package main

import (
	"decomposition/internal/app/handlers/echo"
	"decomposition/internal/app/handlers/ping"
	"decomposition/pkg/logger"
	"fmt"
	"net/http"
)

func main() {
	logger.SetLogger(&logger.ConsoleLogger{})
	logger.Info("Starting server")
	http.HandleFunc("/ping", ping.PingHandler)

	http.HandleFunc("POST /echo", echo.EchoHandler)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
