package main

import (
	"decomposition/internal/app/handlers/routers"
	"decomposition/internal/app/middleware"
	"fmt"
	"net/http"
)

func main() {

	router := http.NewServeMux()
	routers.Router(router)

	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", middleware.Middleware(router))

	if err != nil {
		fmt.Println("Server is not running", err)
	}
}
